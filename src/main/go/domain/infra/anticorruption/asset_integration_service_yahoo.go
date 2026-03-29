package anticorruption

import "github.com/benizzio/open-asset-allocator/domain"

type YahooFinanceAssetIntegrationService struct {
}

func (service *YahooFinanceAssetIntegrationService) SearchAssets(queryValue string) ([]*domain.ExternalAsset, error) {
	//TODO
	return nil, nil
}
