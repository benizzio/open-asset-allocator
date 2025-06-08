package domain

import "time"

// Deprecated: Use PortfolioObservationTimestamp
type TimeFrameTag string

type PortfolioAllocation struct {
	Asset       Asset
	Class       string
	CashReserve bool
	// Deprecated: use PortfolioObservationTimestamp
	TimeFrameTag            TimeFrameTag
	ObservationTimestampTag PortfolioObservationTimestamp
	TotalMarketValue        int64
}

type Portfolio struct {
	Id                  int
	Name                string
	AllocationStructure AllocationStructure
}

type AnalysisOptions struct {
	// Deprecated: use AvailableObservationTimestamps
	AvailableHistory               []TimeFrameTag
	AvailableObservationTimestamps []PortfolioObservationTimestamp
	AvailablePlans                 []*AllocationPlanIdentifier
}

type PortfolioObservationTimestamp struct {
	id                   int
	ObservationTimeTag   string
	ObservationTimestamp time.Time
}

type PortfolioRepository interface {
	GetAllPortfolios() ([]*Portfolio, error)
	GetPortfolio(id int) (*Portfolio, error)
	// Deprecated: use GetAllPortfolioAllocationsWithinObservationTimestampsLimit
	GetAllPortfolioAllocations(id int, timeFrameLimit int) ([]*PortfolioAllocation, error)
	GetAllPortfolioAllocationsWithinObservationTimestampsLimit(id int, observationTimestampsLimit int) (
		[]*PortfolioAllocation,
		error,
	)
	// Deprecated: use FindPortfolioAllocationsByObservationTimestamp
	FindPortfolioAllocations(id int, timeFrameTag TimeFrameTag) ([]*PortfolioAllocation, error)
	FindPortfolioAllocationsByObservationTimestamp(id int, observationTimestampId int) ([]*PortfolioAllocation, error)
	// Deprecated: use GetAvailableObservationTimestamps
	GetAllTimeFrameTags(portfolioId int, timeFrameLimit int) ([]TimeFrameTag, error)
	GetAvailableObservationTimestamps(portfolioId int, observationTimestampsLimit int) (
		[]PortfolioObservationTimestamp,
		error,
	)
	InsertPortfolio(portfolio *Portfolio) (*Portfolio, error)
	UpdatePortfolio(portfolio *Portfolio) (*Portfolio, error)
}
