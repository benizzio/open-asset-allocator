package domain

import (
	"context"
	"fmt"

	"github.com/benizzio/open-asset-allocator/infra"
)

type Asset struct {
	Id     int64
	Name   string
	Ticker string
}

// TODO json marshalling for DB (Scan and Value)
type ExternalData struct {
	Data []ExternalAssetData `json:"data"`
}

type ExternalAssetData struct {
	Source AssetExternalSource `json:"source"`
	Ticker string              `json:"ticker"`
}

type AssetExternalSource string

const (
	YahooFinanceSource AssetExternalSource = "YAHOO_FINANCE"
)

func (externalSource AssetExternalSource) Validate() error {
	switch externalSource {
	case YahooFinanceSource:
		return nil
	}
	return infra.BuildDomainValidationError(fmt.Sprintf("Invalid AssetExternalSource %s", externalSource), nil)
}

type AssetsPerTicker map[string]*Asset

type AssetRepository interface {
	GetKnownAssets() ([]*Asset, error)
	FindAssetByUniqueIdentifier(uniqueIdentifier string) (*Asset, error)
	UpdateAsset(asset *Asset) (*Asset, error)
	InsertAssetsInTransaction(transContext context.Context, assets []*Asset) ([]*Asset, error)
	FindAssetsByTickersInTransaction(transContext context.Context, tickers []string) ([]*Asset, error)
}
