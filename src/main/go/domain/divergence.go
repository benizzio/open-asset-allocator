package domain

type PotentialDivergence struct {
	HierarchyLevelKey          string
	HierarchicalId             string
	TotalMarketValue           int64
	TotalMarketValueDivergence int64
	InternalDivergences        []*PotentialDivergence
}

func (divergence *PotentialDivergence) AddInternalDivergence(internalDivergence *PotentialDivergence) {
	divergence.InternalDivergences = append(divergence.InternalDivergences, internalDivergence)
}

type DivergenceAnalysis struct {
	PortfolioId               int
	TimeFrameTag              TimeFrameTag
	AllocationPlanId          int
	PortfolioTotalMarketValue int64
	Root                      []*PotentialDivergence
}

func (analysis *DivergenceAnalysis) AddRootDivergence(rootDivergence *PotentialDivergence) {
	analysis.Root = append(analysis.Root, rootDivergence)
}
