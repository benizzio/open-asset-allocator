package domain

import (
	"time"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/shopspring/decimal"
)

type PortfolioAllocation struct {
	Asset                Asset
	Class                string
	CashReserve          bool
	ObservationTimestamp *PortfolioObservationTimestamp
	TotalMarketValue     int64
	AssetQuantity        decimal.Decimal
	AssetMarketPrice     decimal.Decimal
}

type PortfolioObservationTimestamp struct {
	Id        int
	TimeTag   string
	Timestamp time.Time
}

type PortfolioAllocationRepository interface {
	FindAllPortfolioAllocationsWithinObservationTimestampsLimit(id int, observationTimestampsLimit int) (
		[]*PortfolioAllocation,
		error,
	)
	FindPortfolioAllocationsByObservationTimestamp(id int, observationTimestampId int) ([]*PortfolioAllocation, error)
	FindAvailableObservationTimestamps(portfolioId int, observationTimestampsLimit int) (
		[]*PortfolioObservationTimestamp,
		error,
	)
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
