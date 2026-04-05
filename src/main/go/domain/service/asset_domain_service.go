package service

import (
	"context"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
)

type AssetIntegrationServicesPerSource map[domain.AssetExternalSource]domain.AssetIntegrationService

type AssetDomService struct {
	assetRepository                   domain.AssetRepository
	assetIntegrationServicesPerSource AssetIntegrationServicesPerSource
}

func (service *AssetDomService) GetKnownAssets() ([]*domain.Asset, error) {
	return service.assetRepository.GetKnownAssets()
}

func (service *AssetDomService) FindAssetByUniqueIdentifier(uniqueIdentifier string) (*domain.Asset, error) {
	return service.assetRepository.FindAssetByUniqueIdentifier(uniqueIdentifier)
}

// UpdateAsset delegates the update of an asset's ticker and name to the repository.
//
// Authored by: GitHub Copilot
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
// Co-authored by: GitHub Copilot (claude-opus-4.6) and benizzio
func (service *AssetDomService) SearchExternalAssets(query string) ([]*domain.ExternalAsset, error) {

	var integrationServices = collectIntegrationServices(service.assetIntegrationServicesPerSource)

	var searchAssetsOnService = func(integrationService domain.AssetIntegrationService) ([]*domain.ExternalAsset, error) {
		return integrationService.SearchAssets(query)
	}

	return langext.FlatMapConcurrently(integrationServices, searchAssetsOnService)
}

func BuildAssetDomService(assetRepository domain.AssetRepository, integrationServices AssetIntegrationServicesPerSource) *AssetDomService {
	return &AssetDomService{
		assetRepository:                   assetRepository,
		assetIntegrationServicesPerSource: integrationServices,
	}
}
