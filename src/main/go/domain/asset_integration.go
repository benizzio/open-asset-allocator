package domain

type AssetIntegrationService interface {
	SearchAssets(queryValue string) ([]*ExternalAsset, error)
	QuoteAsset(asset *ExternalAsset) (*ExternalAssetQuote, error)
}
