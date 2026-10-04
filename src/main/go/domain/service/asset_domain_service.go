package service

import (
	"context"
	"strings"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
)

type AssetIntegrationServicesPerSource map[domain.AssetExternalSource]domain.AssetIntegrationService

type AssetDomService struct {
	assetRepository                   domain.AssetRepository
	assetIntegrationServicesPerSource AssetIntegrationServicesPerSource
}

// GetKnownAssets returns all known assets when textSearch is blank. For a nonblank search, every
// parsed term must occur as a case-insensitive substring of either the asset ticker or name, and
// the result is limited to maxAssetTextSearchResults.
//
// Example:
//
//	assets, err := assetService.GetKnownAssets(`spdr "bloomberg"`)
//
// Authored by: OpenCode
func (service *AssetDomService) GetKnownAssets(textSearch string) ([]*domain.Asset, error) {
	var searchTerms = parseAssetTextSearch(textSearch)
	if len(searchTerms) == 0 {
		if strings.TrimSpace(textSearch) == "" {
			return service.assetRepository.GetKnownAssets()
		}
		return []*domain.Asset{}, nil
	}
	if len(searchTerms) > maxAssetTextSearchTerms {
		var validationError = infra.BuildAppErrorFormattedUnconverted(
			service,
			"Text search must not exceed %d terms",
			maxAssetTextSearchTerms,
		)
		return nil, infra.BuildDomainValidationError(
			"Asset text search validation failed",
			[]*infra.AppError{validationError},
		)
	}

	return service.assetRepository.FindAssetsByTextSearchTerms(searchTerms, maxAssetTextSearchResults)
}

func (service *AssetDomService) FindAssetByUniqueIdentifier(uniqueIdentifier string) (*domain.Asset, error) {
	return service.assetRepository.FindAssetByUniqueIdentifier(uniqueIdentifier)
}

// QuoteExternalAssetLastClosePrice quotes a persisted external-asset association for the given
// local asset ticker. It returns a nil quote without error when the asset or association is missing.
//
// Example:
//
//	quote, err := assetService.QuoteExternalAssetLastClosePrice(ctx, "ARCA:BIL", externalAsset)
//
// Authored by: OpenCode
func (service *AssetDomService) QuoteExternalAssetLastClosePrice(
	requestContext context.Context,
	assetTicker string,
	requestedExternalAsset *domain.ExternalAsset,
) (*domain.ExternalAssetQuote, error) {
	if requestContext == nil {
		return nil, infra.BuildAppError("Request context is required to retrieve an asset quote", service)
	}
	if strings.TrimSpace(assetTicker) == "" {
		return nil, buildExternalAssetQuoteValidationError(service, "Asset ticker is required")
	}
	if requestedExternalAsset == nil {
		return nil, buildExternalAssetQuoteValidationError(service, "External asset identifiers are required")
	}
	if strings.TrimSpace(requestedExternalAsset.Ticker) == "" {
		return nil, buildExternalAssetQuoteValidationError(service, "External asset ticker is required")
	}
	if strings.TrimSpace(requestedExternalAsset.ExchangeId) == "" {
		return nil, buildExternalAssetQuoteValidationError(service, "External asset exchange ID is required")
	}
	if err := requestedExternalAsset.Source.Validate(); err != nil {
		return nil, err
	}

	asset, err := service.assetRepository.FindAssetByTicker(assetTicker)
	if err != nil {
		return nil, err
	}
	if asset == nil || asset.ExternalData == nil {
		return nil, nil
	}

	for index := range asset.ExternalData.Data {
		var persistedExternalAsset = &asset.ExternalData.Data[index]
		if persistedExternalAsset.Source != requestedExternalAsset.Source ||
			persistedExternalAsset.ExchangeId != requestedExternalAsset.ExchangeId ||
			persistedExternalAsset.Ticker != requestedExternalAsset.Ticker {
			continue
		}

		var integrationService, exists = service.assetIntegrationServicesPerSource[persistedExternalAsset.Source]
		if !exists || integrationService == nil {
			return nil, infra.BuildAppErrorFormatted(
				service,
				"No integration service is configured for external asset source %s",
				persistedExternalAsset.Source,
			)
		}

		quote, quoteErr := integrationService.QuoteAssetLastClosePrice(requestContext, persistedExternalAsset)
		if quoteErr != nil {
			return nil, quoteErr
		}
		if quote == nil {
			return nil, infra.BuildAppError("External asset integration returned an empty quote", service)
		}
		return quote, nil
	}

	return nil, nil
}

// buildExternalAssetQuoteValidationError converts an invalid quote identifier into a domain
// validation error that the REST layer maps to HTTP 400.
//
// Authored by: OpenCode
func buildExternalAssetQuoteValidationError(origin any, message string) error {
	var validationError = infra.BuildAppErrorFormattedUnconverted(origin, "%s", message)
	var validationErrors = []*infra.AppError{validationError}
	return infra.BuildDomainValidationError("External asset quote validation failed", validationErrors)
}

// CreateAsset delegates insertion of one asset, including its optional external data, to the
// repository and returns the generated persisted asset.
//
// Example:
//
//	createdAsset, err := assetService.CreateAsset(asset)
//
// Authored by: OpenCode
func (service *AssetDomService) CreateAsset(asset *domain.Asset) (*domain.Asset, error) {
	return service.assetRepository.InsertAsset(asset)
}

// UpdateAsset delegates the update of an asset's ticker, name, and external data to the repository.
//
// Co-authored by: GitHub Copilot and OpenCode
func (service *AssetDomService) UpdateAsset(asset *domain.Asset) (*domain.Asset, error) {
	return service.assetRepository.UpdateAsset(asset)
}

func (service *AssetDomService) InsertAssetsInTransaction(
	transContext context.Context,
	assets []*domain.Asset,
) ([]*domain.Asset, error) {
	return service.assetRepository.InsertAssetsInTransaction(transContext, assets)
}

func (service *AssetDomService) InsertMappedAssetsInTransaction(
	transContext context.Context,
	assetsPerTicker domain.AssetsPerTicker,
) (domain.AssetsPerTicker, error) {

	var assets = make([]*domain.Asset, 0, len(assetsPerTicker))
	for _, asset := range assetsPerTicker {
		assets = append(assets, asset)
	}

	persistedAssets, err := service.InsertAssetsInTransaction(transContext, assets)
	if err != nil {
		return nil, err
	}

	var persistedAssetsPerTicker = make(domain.AssetsPerTicker, len(persistedAssets))
	for _, persistedAsset := range persistedAssets {
		persistedAssetsPerTicker[persistedAsset.Ticker] = persistedAsset
	}

	return persistedAssetsPerTicker, nil
}

// collectIntegrationServices extracts the integration service values from the source-keyed map
// into a slice suitable for concurrent processing.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func collectIntegrationServices(servicesPerSource AssetIntegrationServicesPerSource) []domain.AssetIntegrationService {
	var services = make([]domain.AssetIntegrationService, 0, len(servicesPerSource))
	for _, integrationService := range servicesPerSource {
		services = append(services, integrationService)
	}
	return services
}

// SearchExternalAssets queries all configured external asset integration services concurrently
// for assets matching the given query, and returns the aggregated results.
//
// Parameters:
//   - query: the search term to query across all configured external sources
//
// Returns:
//   - []*domain.ExternalAsset: the aggregated external assets from all sources
//   - error: the first error encountered from any source, or nil if all succeeded
//
// Co-authored by: OpenCode and benizzio
func (service *AssetDomService) SearchExternalAssets(
	requestContext context.Context,
	query string,
) ([]*domain.ExternalAsset, error) {

	var integrationServices = collectIntegrationServices(service.assetIntegrationServicesPerSource)

	var searchAssetsOnService = func(
		searchContext context.Context,
		integrationService domain.AssetIntegrationService,
	) ([]*domain.ExternalAsset, error) {
		return integrationService.SearchAssets(searchContext, query)
	}

	return langext.FlatMapConcurrentlyCtx(requestContext, integrationServices, searchAssetsOnService)
}

func BuildAssetDomService(
	assetRepository domain.AssetRepository,
	integrationServices AssetIntegrationServicesPerSource,
) *AssetDomService {
	return &AssetDomService{
		assetRepository:                   assetRepository,
		assetIntegrationServicesPerSource: integrationServices,
	}
}
