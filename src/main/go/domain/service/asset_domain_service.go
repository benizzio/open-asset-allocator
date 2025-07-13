package service

import "github.com/benizzio/open-asset-allocator/domain"

type AssetDomService struct {
	assetRepository domain.AssetRepository
}

func (service *AssetDomService) GetKnownAssets() ([]*domain.Asset, error) {
	return service.assetRepository.GetKnownAssets()
}

func (service *AssetDomService) FindAssetById(id int) (*domain.Asset, error) {
	return service.assetRepository.FindAssetById(id)
}

func BuildAssetDomService(assetRepository domain.AssetRepository) *AssetDomService {
	return &AssetDomService{
		assetRepository: assetRepository,
	}
}
