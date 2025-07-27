package service

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
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

func (service *AssetDomService) InsertAssetsWithinTransaction(
	transContext *infra.TransactionalContext,
	assets []*domain.Asset,
) ([]*domain.Asset, error) {
	return service.assetRepository.InsertAssetsWithinTransaction(transContext, assets)
}

func (service *AssetDomService) InsertMappedAssetsWithinTransaction(
	transContext *infra.TransactionalContext,
	assetsPerTicker domain.AssetsPerTicker,
) (domain.AssetsPerTicker, error) {

	var assets = make([]*domain.Asset, 0, len(assetsPerTicker))
	for _, asset := range assetsPerTicker {
		assets = append(assets, asset)
	}

	persistedAssets, err := service.InsertAssetsWithinTransaction(transContext, assets)
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
