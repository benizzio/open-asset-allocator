package domain

type Portfolio struct {
	Id                  int
	Name                string
	AllocationStructure AllocationStructure
}

type AnalysisOptions struct {
	AvailableObservationTimestamps []*PortfolioObservationTimestamp
	AvailablePlans                 []*AllocationPlanIdentifier
}

type PortfolioRepository interface {
	GetAllPortfolios() ([]*Portfolio, error)
	FindPortfolio(id int) (*Portfolio, error)
	InsertPortfolio(portfolio *Portfolio) (*Portfolio, error)
	UpdatePortfolio(portfolio *Portfolio) (*Portfolio, error)
}
