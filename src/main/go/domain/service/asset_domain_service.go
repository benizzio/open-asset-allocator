package service

import (
	"context"

	"github.com/benizzio/open-asset-allocator/domain"
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

func (service *AssetDomService) SearchExternalAssets(query string) ([]*domain.ExternalAsset, error) {

	var assets = make([]*domain.ExternalAsset, 0)

	for _, assetIntegrationService := range service.assetIntegrationServicesPerSource {
		//TODO search in parallel goroutines and aggregate results
		externalAssets, err := assetIntegrationService.SearchAssets(query)
		if err != nil {
			return nil, err
		}
		assets = append(assets, externalAssets...)
	}

	return assets, nil
}

func BuildAssetDomService(assetRepository domain.AssetRepository, integrationServices AssetIntegrationServicesPerSource) *AssetDomService {
	return &AssetDomService{
		assetRepository:                   assetRepository,
		assetIntegrationServicesPerSource: integrationServices,
	}
}
