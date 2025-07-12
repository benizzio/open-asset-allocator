package application

import (
	"context"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra/util"
)

const (
	allocationIterationContextKey    = "allocationIteration"
	hierarchySubIterationContextKey  = "hierarchySubIteration"
	divergenceAnalysisContextKey     = "divergenceAnalysis"
	potentialDivergenceMapContextKey = "potentialDivergenceMap"
)

type divergenceAnalysisContextValue struct {
	portfolio            *domain.Portfolio
	portfolioAllocations []*domain.PortfolioAllocation
	divergenceAnalysis   *domain.DivergenceAnalysis
}

type allocationIterationMappingContextValue struct {
	potentialDivergenceMap       potentialDivergencesPerHierarchicalId
	portfolioAllocationsIterator *util.Iterator[*domain.PortfolioAllocation]
}

func buildDivergenceAnalysisContext(
	ctx context.Context,
	divergenceAnalysis *divergenceAnalysisContextValue,
) context.Context {
	return context.WithValue(ctx, divergenceAnalysisContextKey, divergenceAnalysis)
}

func getDivergenceAnalysisContextValue(ctx context.Context) *divergenceAnalysisContextValue {
	return ctx.Value(divergenceAnalysisContextKey).(*divergenceAnalysisContextValue)
}

func buildAllocationIterationContext(
	ctx context.Context,
	value *allocationIterationMappingContextValue,
) context.Context {
	return context.WithValue(ctx, allocationIterationContextKey, value)
}

func getAllocationIterationContextValue(ctx context.Context) *allocationIterationMappingContextValue {
	return ctx.Value(allocationIterationContextKey).(*allocationIterationMappingContextValue)
}

func buildHierarchySubIterationContext(
	ctx context.Context,
	value *util.Iterator[domain.AllocationHierarchyLevel],
) context.Context {
	return context.WithValue(ctx, hierarchySubIterationContextKey, value)
}

func getHierarchySubIterationContextValue(ctx context.Context) *util.Iterator[domain.AllocationHierarchyLevel] {
	return ctx.Value(hierarchySubIterationContextKey).(*util.Iterator[domain.AllocationHierarchyLevel])
}

func buildPotentialDivergenceMapContext(
	ctx context.Context,
	value potentialDivergencesPerHierarchicalId,
) context.Context {
	return context.WithValue(ctx, potentialDivergenceMapContextKey, value)
}

func getPotentialDivergenceMapContextValue(ctx context.Context) potentialDivergencesPerHierarchicalId {
	return ctx.Value(potentialDivergenceMapContextKey).(potentialDivergencesPerHierarchicalId)
}
