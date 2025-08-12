package service

import (
	"context"

	"github.com/benizzio/open-asset-allocator/domain"
)

type AssetDomService struct {
	assetRepository domain.AssetRepository
}

func (service *AssetDomService) GetKnownAssets() ([]*domain.Asset, error) {
	return service.assetRepository.GetKnownAssets()
}

func (service *AssetDomService) FindAssetByUniqueIdentifier(uniqueIdentifier string) (*domain.Asset, error) {
	return service.assetRepository.FindAssetByUniqueIdentifier(uniqueIdentifier)
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

func BuildAssetDomService(assetRepository domain.AssetRepository) *AssetDomService {
	return &AssetDomService{
		assetRepository: assetRepository,
	}
}
