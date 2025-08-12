package domain

import (
	"context"
)

type Asset struct {
	Id     int
	Name   string
	Ticker string
}

type AssetsPerTicker map[string]*Asset

type AssetRepository interface {
	GetKnownAssets() ([]*Asset, error)
	FindAssetByUniqueIdentifier(uniqueIdentifier string) (*Asset, error)
	InsertAssetsInTransaction(transContext context.Context, assets []*Asset) ([]*Asset, error)
	FindAssetsByTickersInTransaction(transContext context.Context, tickers []string) ([]*Asset, error)
}
