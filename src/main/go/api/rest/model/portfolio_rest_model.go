package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/shopspring/decimal"
	"time"
)

type PortfolioDTS struct {
	Id                  *langext.ParseableInt   `json:"id"`
	Name                string                  `json:"name" binding:"required"`
	AllocationStructure *AllocationStructureDTS `json:"allocationStructure"`
}

type PortfolioAllocationDTS struct {
	AssetId          int             `json:"assetId"`
	AssetName        string          `json:"assetName" binding:"required"`
	AssetTicker      string          `json:"assetTicker" binding:"required"`
	Class            string          `json:"class" binding:"required"`
	CashReserve      bool            `json:"cashReserve"`
	TotalMarketValue int64           `json:"totalMarketValue" binding:"required"`
	AssetQuantity    decimal.Decimal `json:"assetQuantity"`
	AssetMarketPrice decimal.Decimal `json:"assetMarketPrice"`
}

type PortfolioSnapshotDTS struct {
	// Deprecated: use PortfolioObservationTimestampDTS
	TimeFrameTag         domain.TimeFrameTag               `json:"timeFrameTag"`
	ObservationTimestamp *PortfolioObservationTimestampDTS `json:"observationTimestamp" binding:"required"`
	Allocations          []*PortfolioAllocationDTS         `json:"allocations"`
	TotalMarketValue     int64                             `json:"totalMarketValue"`
}

type portfolioAllocationsPerObservationTimestamp map[PortfolioObservationTimestampDTS][]*PortfolioAllocationDTS

func (aggregationMap portfolioAllocationsPerObservationTimestamp) getOrBuild(
	observationTimestamp PortfolioObservationTimestampDTS,
) []*PortfolioAllocationDTS {
	var allocationAggregation = aggregationMap[observationTimestamp]
	if allocationAggregation == nil {
		allocationAggregation = make([]*PortfolioAllocationDTS, 0)
	}
	return allocationAggregation
}

func (aggregationMap portfolioAllocationsPerObservationTimestamp) aggregate(
	observationTimestamp PortfolioObservationTimestampDTS,
	allocationDTS *PortfolioAllocationDTS,
) {
	var allocationAggregation = aggregationMap.getOrBuild(observationTimestamp)
	allocationAggregation = append(allocationAggregation, allocationDTS)
	aggregationMap[observationTimestamp] = allocationAggregation
}

func (aggregationMap portfolioAllocationsPerObservationTimestamp) getAggregatedMarketValue(
	observationTimestamp PortfolioObservationTimestampDTS,
) int64 {
	var allocationAggregation = aggregationMap[observationTimestamp]
	var totalMarketValue = int64(0)
	for _, allocation := range allocationAggregation {
		totalMarketValue += allocation.TotalMarketValue
	}
	return totalMarketValue
}

type AllocationPlanIdentifierDTS struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type PortfolioObservationTimestampDTS struct {
	Id        int       `json:"id"`
	TimeTag   string    `json:"timeTag"`
	Timestamp time.Time `json:"timestamp"`
}

type AnalysisOptionsDTS struct {
	// Deprecated: use AvailableObservedHistory
	AvailableHistory         []domain.TimeFrameTag               `json:"availableHistory"`
	AvailableObservedHistory []*PortfolioObservationTimestampDTS `json:"availableObservedHistory"`
	AvailablePlans           []*AllocationPlanIdentifierDTS      `json:"availablePlans"`
}

type PotentialDivergenceDTS struct {
	HierarchyLevelKey          string                    `json:"hierarchyLevelKey"`
	HierarchicalId             string                    `json:"hierarchicalId"`
	TotalMarketValue           int64                     `json:"totalMarketValue"`
	TotalMarketValueDivergence int64                     `json:"totalMarketValueDivergence"`
	Depth                      int                       `json:"depth"`
	InternalDivergences        []*PotentialDivergenceDTS `json:"internalDivergences,omitempty"`
}

type DivergenceAnalysisDTS struct {
	PortfolioId int `json:"portfolioId"`
	// Deprecated: use ObservationTimestamp
	TimeFrameTag              domain.TimeFrameTag               `json:"timeFrameTag"`
	ObservationTimestamp      *PortfolioObservationTimestampDTS `json:"observationTimestamp"`
	AllocationPlanId          int                               `json:"allocationPlanId"`
	PortfolioTotalMarketValue int64                             `json:"portfolioTotalMarketValue"`
	Root                      []*PotentialDivergenceDTS         `json:"root"`
}
