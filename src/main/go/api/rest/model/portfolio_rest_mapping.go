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
	portfolioAllocationsPerObsTimestamp := aggregateHistoryAsDTSMap(portfolioHistory)
	aggregatedPortfolioHistory := buildHistoryDTS(portfolioAllocationsPerObsTimestamp)
	return aggregatedPortfolioHistory
}

func aggregateHistoryAsDTSMap(portfolioHistory []*domain.PortfolioAllocation) portfolioAllocationsPerObservationTimestamp {

	var portfolioAllocationsPerTimestamp = make(portfolioAllocationsPerObservationTimestamp)
	for _, portfolioAllocation := range portfolioHistory {
		var observationTimestampDTS = mapToObservationTimestampDTS(portfolioAllocation.ObservationTimestamp)
		var allocationDTS = mapToPortfolioAllocationDTS(portfolioAllocation)
		portfolioAllocationsPerTimestamp.aggregate(*observationTimestampDTS, allocationDTS)
	}

	return portfolioAllocationsPerTimestamp
}

func mapToPortfolioAllocationDTS(portfolioAllocation *domain.PortfolioAllocation) *PortfolioAllocationDTS {
	return &PortfolioAllocationDTS{
		AssetId:          portfolioAllocation.Asset.Id,
		AssetName:        portfolioAllocation.Asset.Name,
		AssetTicker:      portfolioAllocation.Asset.Ticker,
		Class:            portfolioAllocation.Class,
		CashReserve:      portfolioAllocation.CashReserve,
		TotalMarketValue: portfolioAllocation.TotalMarketValue,
		AssetQuantity:    portfolioAllocation.AssetQuantity,
		AssetMarketPrice: portfolioAllocation.AssetMarketPrice,
	}
}

func buildHistoryDTS(
	portfolioAllocationsPerObservationTimestamp portfolioAllocationsPerObservationTimestamp,
) []*PortfolioSnapshotDTS {

	var aggregatedPortfoliohistory = make([]*PortfolioSnapshotDTS, 0)
	for observationTimestamp, allocations := range portfolioAllocationsPerObservationTimestamp {
		var totalMarketValue = portfolioAllocationsPerObservationTimestamp.getAggregatedMarketValue(observationTimestamp)
		//TODO remove
		var timeFrameTag = domain.TimeFrameTag(observationTimestamp.TimeTag)
		obs := observationTimestamp
		portfolioSnapshot := buildSnapshotDTS(timeFrameTag, &obs, allocations, totalMarketValue)
		aggregatedPortfoliohistory = append(aggregatedPortfoliohistory, portfolioSnapshot)
	}

	// Sort the aggregated portfolio history by observation timestamp, in descending order
	sort.Slice(
		aggregatedPortfoliohistory, func(i, j int) bool {
			return aggregatedPortfoliohistory[i].ObservationTimestamp.Timestamp.After(
				aggregatedPortfoliohistory[j].ObservationTimestamp.Timestamp,
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

func MapToPortfolioObservationTimestampDTSs(
	availableObservationTimestamps []*domain.PortfolioObservationTimestamp,
) []*PortfolioObservationTimestampDTS {
	var observationTimestampsDTS = make([]*PortfolioObservationTimestampDTS, 0)
	for _, observationTimestamp := range availableObservationTimestamps {
		observationTimestampDTS := mapToObservationTimestampDTS(observationTimestamp)
		observationTimestampsDTS = append(observationTimestampsDTS, observationTimestampDTS)
	}
	return observationTimestampsDTS
}

func mapToObservationTimestampDTS(observationTimestamp *domain.PortfolioObservationTimestamp) *PortfolioObservationTimestampDTS {
	var observationTimestampDTS = &PortfolioObservationTimestampDTS{
		Id:        observationTimestamp.Id,
		TimeTag:   observationTimestamp.TimeTag,
		Timestamp: observationTimestamp.Timestamp,
	}
	return observationTimestampDTS
}

func MapToPortfolioObservationTimestamp(
	observationTimestampDTS *PortfolioObservationTimestampDTS,
) *domain.PortfolioObservationTimestamp {
	return &domain.PortfolioObservationTimestamp{
		Id:        observationTimestampDTS.Id,
		TimeTag:   observationTimestampDTS.TimeTag,
		Timestamp: observationTimestampDTS.Timestamp,
	}
}

func MapToPortfolioAllocations(
	portfolioAllocationDTSs []*PortfolioAllocationDTS,
	observationTimestampId int,
) []*domain.PortfolioAllocation {
	var portfolioAllocations = make([]*domain.PortfolioAllocation, 0)
	for _, portfolioAllocationDTS := range portfolioAllocationDTSs {
		var portfolioAllocation = MapToPortfolioAllocation(portfolioAllocationDTS, observationTimestampId)
		portfolioAllocations = append(portfolioAllocations, portfolioAllocation)
	}
	return portfolioAllocations
}

func MapToPortfolioAllocation(
	portfolioAllocationDTS *PortfolioAllocationDTS,
	observationTimestampId int,
) *domain.PortfolioAllocation {
	return &domain.PortfolioAllocation{
		Asset: domain.Asset{
			Id:     portfolioAllocationDTS.AssetId,
			Name:   portfolioAllocationDTS.AssetName,
			Ticker: portfolioAllocationDTS.AssetTicker,
		},
		Class:            portfolioAllocationDTS.Class,
		CashReserve:      portfolioAllocationDTS.CashReserve,
		TotalMarketValue: portfolioAllocationDTS.TotalMarketValue,
		AssetQuantity:    portfolioAllocationDTS.AssetQuantity,
		AssetMarketPrice: portfolioAllocationDTS.AssetMarketPrice,
		ObservationTimestamp: &domain.PortfolioObservationTimestamp{
			Id: observationTimestampId,
		},
	}
}

// ==========================================
// ANALYSIS OPTIONS
// ==========================================

func MapToAnalysisOptionsDTS(analysisOptions *domain.AnalysisOptions) *AnalysisOptionsDTS {
	var plans = mapToAllocationPlanIdentifierDTSs(analysisOptions.AvailablePlans)
	var availableHistory = MapToPortfolioObservationTimestampDTSs(analysisOptions.AvailableObservationTimestamps)
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
	// TODO clean if after removing deprecated implementations
	var observationTimestamp *PortfolioObservationTimestampDTS
	if analysis.ObservationTimestamp != nil {
		observationTimestamp = mapToObservationTimestampDTS(analysis.ObservationTimestamp)
	}
	var analysisDTS = DivergenceAnalysisDTS{
		PortfolioId:               analysis.PortfolioId,
		TimeFrameTag:              analysis.TimeFrameTag,
		ObservationTimestamp:      observationTimestamp,
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
