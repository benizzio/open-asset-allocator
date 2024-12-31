package domain

type TimeFrameTag string

type PortfolioAllocation struct {
	Asset            Asset
	Class            string
	CashReserve      bool
	TimeFrameTag     TimeFrameTag
	TotalMarketValue int
}

type PortfolioRepository interface {
	GetAllPortfolioAllocations(limit int) ([]PortfolioAllocation, error)
}
