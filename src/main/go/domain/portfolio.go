package domain

type PortfolioSliceAtTime struct {
	Asset            Asset
	Class            string
	CashReserve      bool
	TimeFrameTag     string
	TotalMarketValue int
}
