package domain

type AssetIntegrationService interface {
	SearchAssets(queryValue string) ([]*ExternalAsset, error)
	QuoteAssetLastClosePrice(asset *ExternalAsset) (*ExternalAssetQuote, error)
}
