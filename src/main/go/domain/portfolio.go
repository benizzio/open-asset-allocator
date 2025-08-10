package domain

import (
	"time"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/shopspring/decimal"
)

// TODO separate PortfolioAllocation and related components into new files
type PortfolioAllocation struct {
	Asset                Asset
	Class                string
	CashReserve          bool
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
	FindPortfolioAllocationsByObservationTimestamp(id int, observationTimestampId int) ([]*PortfolioAllocation, error)
	GetAvailableObservationTimestamps(portfolioId int, observationTimestampsLimit int) (
		[]*PortfolioObservationTimestamp,
		error,
	)
	InsertPortfolio(portfolio *Portfolio) (*Portfolio, error)
	UpdatePortfolio(portfolio *Portfolio) (*Portfolio, error)
	FindAvailablePortfolioAllocationClasses(portfolioId int) ([]string, error)
	MergePortfolioAllocationsInTransaction(
		transContext *infra.TransactionalContext,
		portfolioId int,
		observationTimestamp *PortfolioObservationTimestamp,
		allocations []*PortfolioAllocation,
	) error
	InsertObservationTimestampInTransaction(
		transContext *infra.TransactionalContext,
		observationTimestamp *PortfolioObservationTimestamp,
	) (*PortfolioObservationTimestamp, error)
}
