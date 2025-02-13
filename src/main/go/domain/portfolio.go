package domain

type TimeFrameTag string

type PortfolioAllocation struct {
	Asset            Asset
	Class            string
	CashReserve      bool
	TimeFrameTag     TimeFrameTag
	TotalMarketValue int64
}

type Portfolio struct {
	Id                  int
	Name                string
	AllocationStructure AllocationStructure
}

type AnalysisOptions struct {
	AvailableHistory []TimeFrameTag
	AvailablePlans   []*AllocationPlanIdentifier
}

type PortfolioRepository interface {
	GetAllPortfolios() ([]*Portfolio, error)
	GetPortfolio(id int) (*Portfolio, error)
	GetAllPortfolioAllocations(id int, limit int) ([]*PortfolioAllocation, error)
	FindPortfolioAllocations(id int, timeFrameTag TimeFrameTag) ([]*PortfolioAllocation, error)
	GetAllTimeFrameTags(portfolioId int, timeFrameLimit int) ([]TimeFrameTag, error)
}
