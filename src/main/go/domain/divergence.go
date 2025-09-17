package domain

// PotentialDivergence maps a point in the allocation hierarchy, within any level, where the portfolio allocation
// and the plan's values are compared and can potentially diverge.
// Example hierarchy: Asset (level 0 - bottom) -> Class (level 1 - top)
type PotentialDivergence struct {

	// HierarchyLevelKey is the key of the point of comparison inside the hierarchy level
	// (e.g., "BTC-USD" as an Asset or "CRYPTO" as a Class).
	HierarchyLevelKey string

	// HierarchicalId is the unique identifier within the hierarchy of the point of comparison
	// (e.g., "BTC-USD|CRYPTO" for a BTC-USD asset in the CRYPTO class, or CRYPTO for the entire CRYPTO class).
	HierarchicalId string

	// TotalMarketValue informs allocated value in the portfolio for the point of comparison.
	// For upper levels within the hierarchy (containing lower levels), it is a sum of all lower levels' points of comparison.
	TotalMarketValue int64

	// TotalMarketValueDivergence informs the difference between the allocated value and the planned value at the point of comparison.
	TotalMarketValueDivergence int64

	// InternalDivergences references to points of comparison in the lower levels of the hierarchy, when they exist.
	InternalDivergences []*PotentialDivergence
}

func (divergence *PotentialDivergence) AddInternalDivergence(internalDivergence *PotentialDivergence) {
	divergence.InternalDivergences = append(divergence.InternalDivergences, internalDivergence)
}

type DivergenceAnalysis struct {
	PortfolioId               int64
	ObservationTimestamp      *PortfolioObservationTimestamp
	AllocationPlanId          int64
	PortfolioTotalMarketValue int64
	Root                      []*PotentialDivergence
}

func (analysis *DivergenceAnalysis) AddRootDivergence(rootDivergence *PotentialDivergence) {
	analysis.Root = append(analysis.Root, rootDivergence)
}
