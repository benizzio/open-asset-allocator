package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"sort"
	"strings"
)

// ==========================================
// PORTFOLIO
// ==========================================

var (
	defaultAllocationStructure = domain.AllocationStructure{
		Hierarchy: []domain.AllocationHierarchyLevel{
			{
				Name:  "Assets",
				Field: "assetTicker",
			},
			{
				Name:  "Classes",
				Field: "class",
			},
		},
	}
)

func MapToPortfolioDTSs(portfolios []*domain.Portfolio) []PortfolioDTS {
	var portfoliosDTS = make([]PortfolioDTS, 0)
	for _, portfolio := range portfolios {
		var portfolioDTS = MapToPortfolioDTS(portfolio)
		portfoliosDTS = append(portfoliosDTS, *portfolioDTS)
	}
	return portfoliosDTS
}

func MapToPortfolioDTS(portfolio *domain.Portfolio) *PortfolioDTS {
	var structure = mapToAllocationStructureDTS(portfolio.AllocationStructure)
	var portfolioDTS = PortfolioDTS{
		Id:                  &portfolio.Id,
		Name:                portfolio.Name,
		AllocationStructure: &structure,
	}
	return &portfolioDTS
}

func MapToPortfolio(portfolioDTS *PortfolioDTS) *domain.Portfolio {

	var allocationStructure domain.AllocationStructure
	if portfolioDTS.AllocationStructure == nil {
		allocationStructure = defaultAllocationStructure
	} else {
		allocationStructure = mapToAllocationStructure(portfolioDTS.AllocationStructure)
	}

	var portfolioId int
	if portfolioDTS.Id != nil {
		portfolioId = *portfolioDTS.Id
	}

	return &domain.Portfolio{
		Id:                  portfolioId,
		Name:                portfolioDTS.Name,
		AllocationStructure: allocationStructure,
	}
}

// ==========================================
// PORTFOLIO HISTORY
// ==========================================

func AggregateAndMapToPortfolioHistoryDTSs(portfolioHistory []*domain.PortfolioAllocation) []PortfolioSnapshotDTS {
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

func buildHistoryDTS(portfolioAllocationsPerTimeFrame portfolioAllocationsPerTimeFrameMap) []PortfolioSnapshotDTS {

	var aggregatedPortfoliohistory = make([]PortfolioSnapshotDTS, 0)
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
) PortfolioSnapshotDTS {
	return PortfolioSnapshotDTS{
		TimeFrameTag:     timeFrameTag,
		Allocations:      allocations,
		TotalMarketValue: totalMarketValue,
	}
}

// ==========================================
// ANALYSIS OPTIONS
// ==========================================

func MapToAnalysisOptionsDTS(analysisOptions *domain.AnalysisOptions) *AnalysisOptionsDTS {
	var plans = mapToAllocationPlanIdentifierDTSs(analysisOptions.AvailablePlans)
	return &AnalysisOptionsDTS{
		AvailableHistory: analysisOptions.AvailableHistory,
		AvailablePlans:   plans,
	}
}

func mapToAllocationPlanIdentifierDTSs(plans []*domain.AllocationPlanIdentifier) []*AllocationPlanIdentifierDTS {
	var plansDTS = make([]*AllocationPlanIdentifierDTS, 0)
	for _, plan := range plans {
		var planDTS = mapToAllocationPlanIdentifierDTS(plan)
		plansDTS = append(plansDTS, planDTS)
	}
	return plansDTS
}

func mapToAllocationPlanIdentifierDTS(plan *domain.AllocationPlanIdentifier) *AllocationPlanIdentifierDTS {
	return &AllocationPlanIdentifierDTS{
		Id:   plan.Id,
		Name: plan.Name,
	}
}

// ==========================================
// DIVERGENCE ANALYSIS
// ==========================================

func MapToDivergenceAnalysisDTS(analysis *domain.DivergenceAnalysis) *DivergenceAnalysisDTS {
	var rootDivergences = mapToPotentialDivergenceDTSs(analysis.Root, 0)
	var analysisDTS = DivergenceAnalysisDTS{
		PortfolioId:               analysis.PortfolioId,
		TimeFrameTag:              analysis.TimeFrameTag,
		AllocationPlanId:          analysis.AllocationPlanId,
		PortfolioTotalMarketValue: analysis.PortfolioTotalMarketValue,
		Root:                      rootDivergences,
	}
	return &analysisDTS
}

func mapToPotentialDivergenceDTSs(divergences []*domain.PotentialDivergence, depth int) []*PotentialDivergenceDTS {
	var divergencesDTS = make([]*PotentialDivergenceDTS, 0)
	for _, divergence := range divergences {
		var divergenceDTS = mapToPotentialDivergenceDTS(divergence, depth)
		divergencesDTS = append(divergencesDTS, divergenceDTS)
	}
	return divergencesDTS
}

func mapToPotentialDivergenceDTS(divergence *domain.PotentialDivergence, depth int) *PotentialDivergenceDTS {
	var internalDivergences = mapToPotentialDivergenceDTSs(divergence.InternalDivergences, depth+1)
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
