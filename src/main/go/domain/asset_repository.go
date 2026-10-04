package domain

import "context"

type AssetRepository interface {
	// GetKnownAssets returns all persisted assets ordered by ticker.
	//
	// Example:
	//
	//	assets, err := assetRepository.GetKnownAssets()
	//
	// Authored by: OpenCode
	GetKnownAssets() ([]*Asset, error)
	// FindAssetsByTextSearchTerms returns persisted assets matching every search term
	// against either the ticker or name, ordered by ticker and limited by limit.
	// A term containing spaces is treated as an exact ordered phrase.
	//
	// Example:
	//
	//	assets, err := assetRepository.FindAssetsByTextSearchTerms([]string{"spdr", "bloomberg"}, 20)
	//
	// Authored by: OpenCode
	FindAssetsByTextSearchTerms(searchTerms []string, limit int) ([]*Asset, error)
	FindAssetByUniqueIdentifier(uniqueIdentifier string) (*Asset, error)
	// InsertAsset inserts one asset and returns the generated persisted record.
	//
	// Example:
	//
	//	createdAsset, err := assetRepository.InsertAsset(asset)
	//
	// Authored by: OpenCode
	InsertAsset(asset *Asset) (*Asset, error)
	UpdateAsset(asset *Asset) (*Asset, error)
	InsertAssetsInTransaction(transContext context.Context, assets []*Asset) ([]*Asset, error)
	FindAssetsByTickersInTransaction(transContext context.Context, tickers []string) ([]*Asset, error)
}
