package domain

type PotentialDivergence struct {
	HierarchyLevelKey          string
	HierarchicalId             string
	TotalMarketValue           int
	TotalMarketValueDivergence int
	InternalDivergences        []*PotentialDivergence
}

type DivergenceAnalysis struct {
	PortfolioId               int
	TimeFrameTag              TimeFrameTag
	AllocationPlanId          int
	PortfolioTotalMarketValue int
	Root                      []*PotentialDivergence
}
