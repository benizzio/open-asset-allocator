package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

// ================================================
// TYPES
// ================================================

type PortfolioDTS struct {
	Id                  int                    `json:"id"`
	Name                string                 `json:"name"`
	AllocationStructure AllocationStructureDTS `json:"allocationStructure"`
}

type PortfolioAllocationDTS struct {
	AssetName        string `json:"assetName"`
	AssetTicker      string `json:"assetTicker"`
	Class            string `json:"class"`
	CashReserve      bool   `json:"cashReserve"`
	TotalMarketValue int64  `json:"totalMarketValue"`
}

type PortfolioAtTimeDTS struct {
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

type PotentialDivergenceDTS struct {
	HierarchyLevelKey          string                    `json:"hierarchyLevelKey"`
	HierarchicalId             string                    `json:"hierarchicalId"`
	TotalMarketValue           int64                     `json:"totalMarketValue"`
	TotalMarketValueDivergence int64                     `json:"totalMarketValueDivergence"`
	InternalDivergences        []*PotentialDivergenceDTS `json:"internalDivergences,omitempty"`
}

type DivergenceAnalysisDTS struct {
	PortfolioId               int                       `json:"portfolioId"`
	TimeFrameTag              domain.TimeFrameTag       `json:"timeFrameTag"`
	AllocationPlanId          int                       `json:"allocationPlanId"`
	PortfolioTotalMarketValue int64                     `json:"portfolioTotalMarketValue"`
	Root                      []*PotentialDivergenceDTS `json:"root"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func MapPortfolios(portfolios []*domain.Portfolio) []PortfolioDTS {
	var portfoliosDTS = make([]PortfolioDTS, 0)
	for _, portfolio := range portfolios {
		var portfolioDTS = MapPortfolio(portfolio)
		portfoliosDTS = append(portfoliosDTS, *portfolioDTS)
	}
	return portfoliosDTS
}

func MapPortfolio(portfolio *domain.Portfolio) *PortfolioDTS {
	var structure = mapAllocationStructure(portfolio.AllocationStructure)
	var portfolioDTS = PortfolioDTS{
		Id:                  portfolio.Id,
		Name:                portfolio.Name,
		AllocationStructure: structure,
	}
	return &portfolioDTS
}

func AggregateAndMapPortfolioHistory(portfolioHistory []*domain.PortfolioAllocation) []PortfolioAtTimeDTS {
	portfolioAllocationsPerTimeFrame := aggregateHistoryAsDTSMap(portfolioHistory)
	aggregatedPortfolioHistory := buildHistoryDTS(portfolioAllocationsPerTimeFrame)
	return aggregatedPortfolioHistory
}

func aggregateHistoryAsDTSMap(portfolioHistory []*domain.PortfolioAllocation) portfolioAllocationsPerTimeFrameMap {

	var portfolioAllocationsPerTimeFrame = make(portfolioAllocationsPerTimeFrameMap)
	for _, portfolioAllocation := range portfolioHistory {
		var aggregationTimeFrame = portfolioAllocation.TimeFrameTag
		var allocationJSON = portfolioAllocationToAllocationDTS(portfolioAllocation)
		portfolioAllocationsPerTimeFrame.aggregate(aggregationTimeFrame, allocationJSON)
	}

	return portfolioAllocationsPerTimeFrame
}

func portfolioAllocationToAllocationDTS(portfolioAllocation *domain.PortfolioAllocation) PortfolioAllocationDTS {
	return PortfolioAllocationDTS{
		AssetName:        portfolioAllocation.Asset.Name,
		AssetTicker:      portfolioAllocation.Asset.Ticker,
		Class:            portfolioAllocation.Class,
		CashReserve:      portfolioAllocation.CashReserve,
		TotalMarketValue: portfolioAllocation.TotalMarketValue,
	}
}

func buildHistoryDTS(portfolioAllocationsPerTimeFrame portfolioAllocationsPerTimeFrameMap) []PortfolioAtTimeDTS {

	var aggregatedPortfoliohistory = make([]PortfolioAtTimeDTS, 0)
	for timeFrameTag, allocations := range portfolioAllocationsPerTimeFrame {
		var totalMarketValue = portfolioAllocationsPerTimeFrame.getAggregatedMarketValue(timeFrameTag)
		portfolioSnapshot := buildSnapshotDTS(timeFrameTag, allocations, totalMarketValue)
		aggregatedPortfoliohistory = append(aggregatedPortfoliohistory, portfolioSnapshot)
	}

	return aggregatedPortfoliohistory
}

func buildSnapshotDTS(
	timeFrameTag domain.TimeFrameTag,
	allocations []PortfolioAllocationDTS,
	totalMarketValue int64,
) PortfolioAtTimeDTS {
	return PortfolioAtTimeDTS{
		TimeFrameTag:     timeFrameTag,
		Allocations:      allocations,
		TotalMarketValue: totalMarketValue,
	}
}

func MapDivergenceAnalysis(analysis *domain.DivergenceAnalysis) *DivergenceAnalysisDTS {
	var rootDivergences = mapPotentialDivergences(analysis.Root)
	var analysisDTS = DivergenceAnalysisDTS{
		PortfolioId:               analysis.PortfolioId,
		TimeFrameTag:              analysis.TimeFrameTag,
		AllocationPlanId:          analysis.AllocationPlanId,
		PortfolioTotalMarketValue: analysis.PortfolioTotalMarketValue,
		Root:                      rootDivergences,
	}
	return &analysisDTS
}

func mapPotentialDivergences(divergences []*domain.PotentialDivergence) []*PotentialDivergenceDTS {
	var divergencesDTS = make([]*PotentialDivergenceDTS, 0)
	for _, divergence := range divergences {
		var divergenceDTS = mapPotentialDivergence(divergence)
		divergencesDTS = append(divergencesDTS, divergenceDTS)
	}
	return divergencesDTS
}

func mapPotentialDivergence(divergence *domain.PotentialDivergence) *PotentialDivergenceDTS {
	var internalDivergences = mapPotentialDivergences(divergence.InternalDivergences)
	var divergenceDTS = PotentialDivergenceDTS{
		HierarchyLevelKey:          divergence.HierarchyLevelKey,
		HierarchicalId:             divergence.HierarchicalId,
		TotalMarketValue:           divergence.TotalMarketValue,
		TotalMarketValueDivergence: divergence.TotalMarketValueDivergence,
		InternalDivergences:        internalDivergences,
	}
	return &divergenceDTS
}
