package model

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioDTS struct {
	Id                  *langext.ParseableInt64 `json:"id"`
	Name                string                  `json:"name" validate:"required,max=100"`
	AllocationStructure *AllocationStructureDTS `json:"allocationStructure"`
}

type PortfolioAllocationDTS struct {
	AssetId          *langext.ParseableInt64 `json:"assetId"`
	AssetName        string                  `json:"assetName" validate:"max=100"`
	AssetTicker      string                  `json:"assetTicker" validate:"max=40"`
	ExternalAssetDTS *ExternalAssetDTS       `json:"externalAsset,omitempty"`
	Class            string                  `json:"class" validate:"required,max=100"`
	CashReserve      bool                    `json:"cashReserve"`
	TotalMarketValue *decimal.Decimal        `json:"totalMarketValue" validate:"required"`
	// TODO AssetQuantity should be a pointer downstream but we get errors
	AssetQuantity decimal.Decimal `json:"assetQuantity"`
	// TODO AssetMarketPrice should be a pointer downstream but we get errors
	AssetMarketPrice decimal.Decimal `json:"assetMarketPrice"`
}

type PortfolioSnapshotDTS struct {
	ObservationTimestamp *PortfolioObservationTimestampDTS `json:"observationTimestamp" validate:"required"`
	Allocations          []*PortfolioAllocationDTS         `json:"allocations" validate:"required,min=1"`
	TotalMarketValue     *decimal.Decimal                  `json:"totalMarketValue"`
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
) *decimal.Decimal {
	var allocationAggregation = aggregationMap[observationTimestamp]
	var totalMarketValue = decimal.NewFromInt(0)
	for _, allocation := range allocationAggregation {
		if allocation.TotalMarketValue != nil {
			totalMarketValue = totalMarketValue.Add(*allocation.TotalMarketValue)
		}
	}
	return &totalMarketValue
}

// AllocationPlanIdentifierDTS is the plan-summary REST model returned by analysis-option reads.
// Each plan includes its YYYY-MM-DD plannedDate, for example `{"id":1,"name":"Balanced","plannedDate":"2026-10-17"}`.
//
// Co-authored by: OpenCode and benizzio
type AllocationPlanIdentifierDTS struct {
	Id          int64              `json:"id"`
	Name        string             `json:"name" validate:"max=100"`
	PlannedDate AllocationPlanDate `json:"plannedDate"`
}

// PortfolioObservationTimestampDTS carries the observation identifier, required nonblank tag, and timestamp.
// Create and update requests provide a tag of at most 100 characters and an RFC3339 timestamp, for example
// `{"id":7,"timeTag":"Quarter-end","timestamp":"2026-10-17T00:00:00Z"}`.
//
// Co-authored by: GPT-6 Luna and benizzio
type PortfolioObservationTimestampDTS struct {
	Id        langext.ParseableInt64 `json:"id"`
	TimeTag   string                 `json:"timeTag" validate:"required,max=100"`
	Timestamp time.Time              `json:"timestamp" validate:"required"`
}

type AnalysisOptionsDTS struct {
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
	PortfolioId               int64                             `json:"portfolioId"`
	ObservationTimestamp      *PortfolioObservationTimestampDTS `json:"observationTimestamp"`
	AllocationPlanId          int64                             `json:"allocationPlanId"`
	PortfolioTotalMarketValue int64                             `json:"portfolioTotalMarketValue"`
	Root                      []*PotentialDivergenceDTS         `json:"root"`
}
