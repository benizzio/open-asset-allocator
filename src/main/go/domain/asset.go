package domain

import "github.com/benizzio/open-asset-allocator/infra"

type Asset struct {
	Id     int
	Name   string
	Ticker string
}

type AssetsPerTicker map[string]*Asset

type AssetRepository interface {
	GetKnownAssets() ([]*Asset, error)
	FindAssetByUniqueIdentifier(uniqueIdentifier string) (*Asset, error)
	InsertAssetsInTransaction(transContext *infra.TransactionalContext, assets []*Asset) ([]*Asset, error)
	FindAssetsByTickers(transContext *infra.TransactionalContext, tickers []string) ([]*Asset, error)
}
