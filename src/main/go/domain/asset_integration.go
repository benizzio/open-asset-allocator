package domain

import "context"

// AssetIntegrationService defines the contract for external asset providers used by the
// domain layer. Implementations must return normalized assets for their source and must
// respect cancellation in SearchAssets so request-scoped fan-out work can stop promptly.
//
// Co-authored by: OpenCode and benizzio
type AssetIntegrationService interface {

	// SearchAssets searches the provider for assets matching the given free-text query.
	// It returns a normalized slice for the provider source or an error when the provider
	// cannot complete the request. The method must honor ctx cancellation.
	SearchAssets(ctx context.Context, queryValue string) ([]*ExternalAsset, error)

	// QuoteAssetLastClosePrice fetches the latest close quote for one normalized external asset.
	// It returns the provider quote translated to the domain model or an error when quoting fails.
	QuoteAssetLastClosePrice(asset *ExternalAsset) (*ExternalAssetQuote, error)
}
