package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"sort"
	"strings"
)

// ==========================================
// PORTFOLIO
// ==========================================

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

// ==========================================
// PORTFOLIO HISTORY
// ==========================================

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

	// Sort the aggregated portfolio history by time frame tag, in descending order
	sort.Slice(
		aggregatedPortfoliohistory, func(i, j int) bool {
			return strings.Compare(
				string(aggregatedPortfoliohistory[i].TimeFrameTag),
				string(aggregatedPortfoliohistory[j].TimeFrameTag),
			) > 0
		},
	)

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

// ==========================================
// ANALYSIS OPTIONS
// ==========================================

func MapAnalysisOptions(analysisOptions *domain.AnalysisOptions) *AnalysisOptionsDTS {
	var plans = mapAllocationPlanIdentifiers(analysisOptions.AvailablePlans)
	return &AnalysisOptionsDTS{
		AvailableHistory: analysisOptions.AvailableHistory,
		AvailablePlans:   plans,
	}
}

func mapAllocationPlanIdentifiers(plans []*domain.AllocationPlanIdentifier) []*AllocationPlanIdentifierDTS {
	var plansDTS = make([]*AllocationPlanIdentifierDTS, 0)
	for _, plan := range plans {
		var planDTS = mapAllocationPlanIdentifier(plan)
		plansDTS = append(plansDTS, planDTS)
	}
	return plansDTS
}

func mapAllocationPlanIdentifier(plan *domain.AllocationPlanIdentifier) *AllocationPlanIdentifierDTS {
	return &AllocationPlanIdentifierDTS{
		Id:   plan.Id,
		Name: plan.Name,
	}
}

// ==========================================
// DIVERGENCE ANALYSIS
// ==========================================

func MapDivergenceAnalysis(analysis *domain.DivergenceAnalysis) *DivergenceAnalysisDTS {
	var rootDivergences = mapPotentialDivergences(analysis.Root, 0)
	var analysisDTS = DivergenceAnalysisDTS{
		PortfolioId:               analysis.PortfolioId,
		TimeFrameTag:              analysis.TimeFrameTag,
		AllocationPlanId:          analysis.AllocationPlanId,
		PortfolioTotalMarketValue: analysis.PortfolioTotalMarketValue,
		Root:                      rootDivergences,
	}
	return &analysisDTS
}

func mapPotentialDivergences(divergences []*domain.PotentialDivergence, depth int) []*PotentialDivergenceDTS {
	var divergencesDTS = make([]*PotentialDivergenceDTS, 0)
	for _, divergence := range divergences {
		var divergenceDTS = mapPotentialDivergence(divergence, depth)
		divergencesDTS = append(divergencesDTS, divergenceDTS)
	}
	return divergencesDTS
}

func mapPotentialDivergence(divergence *domain.PotentialDivergence, depth int) *PotentialDivergenceDTS {
	var internalDivergences = mapPotentialDivergences(divergence.InternalDivergences, depth+1)
	var divergenceDTS = PotentialDivergenceDTS{
		HierarchyLevelKey:          divergence.HierarchyLevelKey,
		HierarchicalId:             divergence.HierarchicalId,
		TotalMarketValue:           divergence.TotalMarketValue,
		TotalMarketValueDivergence: divergence.TotalMarketValueDivergence,
		Depth:                      depth,
		InternalDivergences:        internalDivergences,
	}
	return &divergenceDTS
}
