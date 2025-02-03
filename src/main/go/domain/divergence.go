package domain

type PotentialDivergence struct {
	HierarchyLevelKey          string
	HierarchicalId             string
	TotalMarketValue           int64
	TotalMarketValueDivergence int64
	InternalDivergences        []*PotentialDivergence
}

type DivergenceAnalysis struct {
	PortfolioId               int
	TimeFrameTag              TimeFrameTag
	AllocationPlanId          int
	PortfolioTotalMarketValue int64
	Root                      []*PotentialDivergence
}
