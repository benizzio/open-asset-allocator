package repository

import (
	"github.com/shopspring/decimal"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
)

// portfolioAllocationJoinedRowDTS represents a joined portfolio allocation read-model row,
// including the first persisted external asset reference already projected by SQL.
//
// Authored by: OpenCode
type portfolioAllocationJoinedRowDTS struct {
	Asset                           *domain.Asset
	Class                           string
	CashReserve                     bool
	ObservationTimestamp            *domain.PortfolioObservationTimestamp
	TotalMarketValue                int64
	AssetQuantity                   decimal.Decimal
	AssetMarketPrice                decimal.Decimal
	SelectedExternalAssetSource     string
	SelectedExternalAssetTicker     string
	SelectedExternalAssetExchangeId string
}

// mapPortfolioAllocationRows converts joined query rows into domain portfolio allocations while
// preserving shared nested pointers for assets and observation timestamps.
//
// Authored by: OpenCode
func mapPortfolioAllocationRows(rows []portfolioAllocationJoinedRowDTS) []*domain.PortfolioAllocation {
	langext.UnifyStructPointers(rows)

	var allocations = make([]*domain.PortfolioAllocation, len(rows))
	for index, row := range rows {
		allocations[index] = mapPortfolioAllocationRow(&row)
	}

	return allocations
}

// mapPortfolioAllocationRow maps a single joined repository row into the portfolio allocation
// domain model.
//
// Authored by: OpenCode
func mapPortfolioAllocationRow(rowDTS *portfolioAllocationJoinedRowDTS) *domain.PortfolioAllocation {
	var asset domain.Asset
	if rowDTS.Asset != nil {
		asset = *rowDTS.Asset
	}

	return &domain.PortfolioAllocation{
		Asset:                 asset,
		SelectedExternalAsset: buildSelectedExternalAsset(rowDTS),
		Class:                 rowDTS.Class,
		CashReserve:           rowDTS.CashReserve,
		ObservationTimestamp:  rowDTS.ObservationTimestamp,
		TotalMarketValue:      rowDTS.TotalMarketValue,
		AssetQuantity:         rowDTS.AssetQuantity,
		AssetMarketPrice:      rowDTS.AssetMarketPrice,
	}
}

// buildSelectedExternalAsset creates the projected external asset only when the query extracted
// persisted reference fields from the asset record.
//
// Authored by: OpenCode
func buildSelectedExternalAsset(rowDTS *portfolioAllocationJoinedRowDTS) *domain.ExternalAsset {
	if langext.IsZeroValue(rowDTS.SelectedExternalAssetSource) &&
		langext.IsZeroValue(rowDTS.SelectedExternalAssetTicker) &&
		langext.IsZeroValue(rowDTS.SelectedExternalAssetExchangeId) {
		return nil
	}

	return &domain.ExternalAsset{
		Source:     domain.AssetExternalSource(rowDTS.SelectedExternalAssetSource),
		Ticker:     rowDTS.SelectedExternalAssetTicker,
		ExchangeId: rowDTS.SelectedExternalAssetExchangeId,
	}
}
