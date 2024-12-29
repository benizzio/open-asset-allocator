package domain

type TimeFrameTag string

type PortfolioAllocation struct {
	Asset            Asset
	Class            string
	CashReserve      bool
	TimeFrameTag     TimeFrameTag
	TotalMarketValue int
}

type Portfolio struct {
	Id                  int
	Name                string
	AllocationStructure AllocationStructure
}

type PortfolioRepository interface {
	GetAllPortfolios() ([]Portfolio, error)
	GetPortfolio(id int) (Portfolio, error)
	GetAllPortfolioAllocations(id int, limit int) ([]PortfolioAllocation, error)
}
