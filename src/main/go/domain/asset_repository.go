package domain

import "context"

type AssetRepository interface {
	GetKnownAssets() ([]*Asset, error)
	FindAssetByUniqueIdentifier(uniqueIdentifier string) (*Asset, error)
	UpdateAsset(asset *Asset) (*Asset, error)
	InsertAssetsInTransaction(transContext context.Context, assets []*Asset) ([]*Asset, error)
	FindAssetsByTickersInTransaction(transContext context.Context, tickers []string) ([]*Asset, error)
}
