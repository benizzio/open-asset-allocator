package domain

type PortfolioSliceAtTime struct {
	Asset            Asset
	Class            string
	CashReserve      bool
	TimeFrameTag     string
	TotalMarketValue int
}

type PortfolioRepository interface {
	GetAllPortfolioSlices(limit int) ([]PortfolioSliceAtTime, error)
}
