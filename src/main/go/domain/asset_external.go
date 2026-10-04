package domain

import (
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/shopspring/decimal"
	"golang.org/x/text/currency"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms/sqlext"
)

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

type ExternalAssetData struct {
	Data []ExternalAsset `json:"data"`
}

func (externalData *ExternalAssetData) Scan(value any) error {
	return sqlext.ScanJsonColumn(value, externalData)
}

func (externalData ExternalAssetData) Value() (driver.Value, error) {
	return sqlext.ValueJsonColumn(externalData)
}

type ExternalAsset struct {
	Source       AssetExternalSource `json:"source" validate:"required"`
	Ticker       string              `json:"ticker" validate:"required"`
	ExchangeId   string              `json:"exchangeId" validate:"required"`
	Name         string              `json:"-"`
	ExchangeName string              `json:"-"`
}

type ExternalAssetQuote struct {
	Ticker         string
	ExchangeId     string
	Currency       currency.Unit
	LastCloseQuote decimal.Decimal
	LastCloseDate  time.Time
}
