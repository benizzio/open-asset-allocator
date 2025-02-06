package application

import (
	"context"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra/util"
)

const (
	allocationIterationContextKey   = "allocationIteration"
	hierarchySubIterationContextKey = "hierarchySubIteration"
	divergenceAnalysisContextKey    = "divergenceAnalysis"
)

type allocationIterationMappingContextValue struct {
	potentialDivergenceMap       map[string]*domain.PotentialDivergence
	portfolioAllocationsIterator *util.Iterator[*domain.PortfolioAllocation]
}

func buildDivergenceAnalysisContext(
	ctx context.Context,
	divergenceAnalysis *domain.DivergenceAnalysis,
) context.Context {
	return context.WithValue(ctx, divergenceAnalysisContextKey, divergenceAnalysis)
}

func getDivergenceAnalysisContextValue(ctx context.Context) *domain.DivergenceAnalysis {
	return ctx.Value(divergenceAnalysisContextKey).(*domain.DivergenceAnalysis)
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
