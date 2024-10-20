package domain

type TimeFrameTag string

type PortfolioSliceAtTime struct {
	Asset            Asset
	Class            string
	CashReserve      bool
	TimeFrameTag     TimeFrameTag
	TotalMarketValue int
}

type PortfolioRepository interface {
	GetAllPortfolioSlices(limit int) ([]PortfolioSliceAtTime, error)
}
