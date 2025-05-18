package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type PortfolioDTS struct {
	Id                  *int                    `json:"id"`
	Name                string                  `json:"name" binding:"required"`
	AllocationStructure *AllocationStructureDTS `json:"allocationStructure"`
}

type PortfolioAllocationDTS struct {
	AssetName        string `json:"assetName"`
	AssetTicker      string `json:"assetTicker"`
	Class            string `json:"class"`
	CashReserve      bool   `json:"cashReserve"`
	TotalMarketValue int64  `json:"totalMarketValue"`
}

type PortfolioSnapshotDTS struct {
	TimeFrameTag     domain.TimeFrameTag      `json:"timeFrameTag"`
	Allocations      []PortfolioAllocationDTS `json:"allocations"`
	TotalMarketValue int64                    `json:"totalMarketValue"`
}

type portfolioAllocationsPerTimeFrameMap map[domain.TimeFrameTag][]PortfolioAllocationDTS

func (
	aggregationMap portfolioAllocationsPerTimeFrameMap,
) getOrBuild(timeFrameTag domain.TimeFrameTag) []PortfolioAllocationDTS {
	var allocationAggregation = aggregationMap[timeFrameTag]
	if allocationAggregation == nil {
		allocationAggregation = make([]PortfolioAllocationDTS, 0)
	}
	return allocationAggregation
}

func (aggregationMap portfolioAllocationsPerTimeFrameMap) aggregate(
	timeFrameTag domain.TimeFrameTag,
	allocationDTS PortfolioAllocationDTS,
) {
	var allocationAggregation = aggregationMap.getOrBuild(timeFrameTag)
	allocationAggregation = append(allocationAggregation, allocationDTS)
	aggregationMap[timeFrameTag] = allocationAggregation
}

func (aggregationMap portfolioAllocationsPerTimeFrameMap) getAggregatedMarketValue(timeFrame domain.TimeFrameTag) int64 {
	var allocationAggregation = aggregationMap[timeFrame]
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

type AnalysisOptionsDTS struct {
	AvailableHistory []domain.TimeFrameTag          `json:"availableHistory"`
	AvailablePlans   []*AllocationPlanIdentifierDTS `json:"availablePlans"`
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
	PortfolioId               int                       `json:"portfolioId"`
	TimeFrameTag              domain.TimeFrameTag       `json:"timeFrameTag"`
	AllocationPlanId          int                       `json:"allocationPlanId"`
	PortfolioTotalMarketValue int64                     `json:"portfolioTotalMarketValue"`
	Root                      []*PotentialDivergenceDTS `json:"root"`
}
