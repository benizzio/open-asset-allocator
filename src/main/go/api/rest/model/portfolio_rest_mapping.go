package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
	"sort"
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
	var portfolioId = langext.ParseableInt(portfolio.Id)
	var portfolioDTS = PortfolioDTS{
		Id:                  &portfolioId,
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
		portfolioId = int(*portfolioDTS.Id)
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

func AggregateAndMapToPortfolioHistoryDTSs(portfolioHistory []*domain.PortfolioAllocation) []*PortfolioSnapshotDTS {
	portfolioAllocationsPerTimeFrame := aggregateHistoryAsDTSMap(portfolioHistory)
	aggregatedPortfolioHistory := buildHistoryDTS(portfolioAllocationsPerTimeFrame)
	return aggregatedPortfolioHistory
}

func aggregateHistoryAsDTSMap(portfolioHistory []*domain.PortfolioAllocation) portfolioAllocationsPerObservationTimestamp {

	var portfolioAllocationsPerTimestamp = make(portfolioAllocationsPerObservationTimestamp)
	for _, portfolioAllocation := range portfolioHistory {
		var observationTimestampDTS = mapToAvailableObservationTimestampDTS(portfolioAllocation.ObservationTimestamp)
		var allocationDTS = portfolioAllocationToAllocationDTS(portfolioAllocation)
		portfolioAllocationsPerTimestamp.aggregate(observationTimestampDTS, allocationDTS)
	}

	return portfolioAllocationsPerTimestamp
}

func portfolioAllocationToAllocationDTS(portfolioAllocation *domain.PortfolioAllocation) *PortfolioAllocationDTS {
	return &PortfolioAllocationDTS{
		AssetId:          portfolioAllocation.Asset.Id,
		AssetName:        portfolioAllocation.Asset.Name,
		AssetTicker:      portfolioAllocation.Asset.Ticker,
		Class:            portfolioAllocation.Class,
		CashReserve:      portfolioAllocation.CashReserve,
		TotalMarketValue: portfolioAllocation.TotalMarketValue,
	}
}

func buildHistoryDTS(
	portfolioAllocationsPerObservationTimestamp portfolioAllocationsPerObservationTimestamp,
) []*PortfolioSnapshotDTS {

	var aggregatedPortfoliohistory = make([]*PortfolioSnapshotDTS, 0)
	for observationTimestamp, allocations := range portfolioAllocationsPerObservationTimestamp {
		var totalMarketValue = portfolioAllocationsPerObservationTimestamp.getAggregatedMarketValue(observationTimestamp)
		//TODO remove
		var timeFrameTag = domain.TimeFrameTag(observationTimestamp.ObservationTimeTag)
		portfolioSnapshot := buildSnapshotDTS(timeFrameTag, observationTimestamp, allocations, totalMarketValue)
		aggregatedPortfoliohistory = append(aggregatedPortfoliohistory, portfolioSnapshot)
	}

	// Sort the aggregated portfolio history by observation timestamp, in descending order
	sort.Slice(
		aggregatedPortfoliohistory, func(i, j int) bool {
			return aggregatedPortfoliohistory[i].ObservationTimestamp.ObservationTimestamp.After(
				aggregatedPortfoliohistory[j].ObservationTimestamp.ObservationTimestamp,
			)
		},
	)

	return aggregatedPortfoliohistory
}

func buildSnapshotDTS(
	// Deprecated: use observationTimestamp
	timeFrameTag domain.TimeFrameTag,
	observationTimestamp *PortfolioObservationTimestampDTS,
	allocations []*PortfolioAllocationDTS,
	totalMarketValue int64,
) *PortfolioSnapshotDTS {
	return &PortfolioSnapshotDTS{
		TimeFrameTag:         timeFrameTag,
		ObservationTimestamp: observationTimestamp,
		Allocations:          allocations,
		TotalMarketValue:     totalMarketValue,
	}
}

func mapToAvailableObservationTimestampsDTS(
	availableObservationTimestamps []*domain.PortfolioObservationTimestamp,
) []*PortfolioObservationTimestampDTS {
	var observationTimestampsDTS = make([]*PortfolioObservationTimestampDTS, 0)
	for _, observationTimestamp := range availableObservationTimestamps {
		observationTimestampDTS := mapToAvailableObservationTimestampDTS(observationTimestamp)
		observationTimestampsDTS = append(observationTimestampsDTS, observationTimestampDTS)
	}
	return observationTimestampsDTS
}

func mapToAvailableObservationTimestampDTS(observationTimestamp *domain.PortfolioObservationTimestamp) *PortfolioObservationTimestampDTS {
	var observationTimestampDTS = &PortfolioObservationTimestampDTS{
		Id:                   observationTimestamp.Id,
		ObservationTimeTag:   observationTimestamp.ObservationTimeTag,
		ObservationTimestamp: observationTimestamp.ObservationTimestamp,
	}
	return observationTimestampDTS
}

// ==========================================
// ANALYSIS OPTIONS
// ==========================================

func MapToAnalysisOptionsDTS(analysisOptions *domain.AnalysisOptions) *AnalysisOptionsDTS {
	var plans = mapToAllocationPlanIdentifierDTSs(analysisOptions.AvailablePlans)
	var availableHistory = mapToAvailableObservationTimestampsDTS(analysisOptions.AvailableObservationTimestamps)
	return &AnalysisOptionsDTS{
		AvailableHistory:         analysisOptions.AvailableHistory,
		AvailableObservedHistory: availableHistory,
		AvailablePlans:           plans,
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
