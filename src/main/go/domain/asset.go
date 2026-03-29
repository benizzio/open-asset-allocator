package domain

import (
	"database/sql/driver"
	"fmt"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms/sqlext"
)

type Asset struct {
	Id           int64
	Name         string
	Ticker       string
	ExternalData *ExternalData
}

type ExternalData struct {
	Data []ExternalAssetData `json:"data"`
}

func (externalData *ExternalData) Scan(value interface{}) error {
	return sqlext.ScanJsonColumn(value, externalData)
}

func (externalData ExternalData) Value() (driver.Value, error) {
	return sqlext.ValueJsonColumn(externalData)
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
