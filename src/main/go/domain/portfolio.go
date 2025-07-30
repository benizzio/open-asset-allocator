package domain

import (
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/shopspring/decimal"
	"time"
)

// Deprecated: Use PortfolioObservationTimestamp
type TimeFrameTag string

// TODO separate PortfolioAllocation and related components into new files
type PortfolioAllocation struct {
	Asset       Asset
	Class       string
	CashReserve bool
	// Deprecated: use PortfolioObservationTimestamp
	TimeFrameTag         TimeFrameTag
	ObservationTimestamp *PortfolioObservationTimestamp
	TotalMarketValue     int64
	AssetQuantity        decimal.Decimal
	AssetMarketPrice     decimal.Decimal
}

type Portfolio struct {
	Id                  int
	Name                string
	AllocationStructure AllocationStructure
}

type AnalysisOptions struct {
	// Deprecated: use AvailableObservationTimestamps
	AvailableHistory               []TimeFrameTag
	AvailableObservationTimestamps []*PortfolioObservationTimestamp
	AvailablePlans                 []*AllocationPlanIdentifier
}

type PortfolioObservationTimestamp struct {
	Id        int
	TimeTag   string
	Timestamp time.Time
}

type PortfolioRepository interface {
	GetAllPortfolios() ([]*Portfolio, error)
	GetPortfolio(id int) (*Portfolio, error)
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
		[]*PortfolioObservationTimestamp,
		error,
	)
	InsertPortfolio(portfolio *Portfolio) (*Portfolio, error)
	UpdatePortfolio(portfolio *Portfolio) (*Portfolio, error)
	FindAvailablePortfolioAllocationClasses(portfolioId int) ([]string, error)
	MergePortfolioAllocationsInTransaction(
		transContext *infra.TransactionalContext,
		id int,
		allocations []*PortfolioAllocation,
	) error
	InsertObservationTimestampInTransaction(
		transContext *infra.TransactionalContext,
		observationTimestamp *PortfolioObservationTimestamp,
	) (*PortfolioObservationTimestamp, error)
}
