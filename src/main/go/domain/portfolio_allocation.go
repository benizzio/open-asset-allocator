package domain

import (
	"context"
	"time"

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
	FindAllPortfolioAllocationsWithinObservationTimestampsLimit(
		id int,
		observationTimestampsLimit int,
	) ([]*PortfolioAllocation, error)
	FindPortfolioAllocationsByObservationTimestamp(id int, observationTimestampId int) ([]*PortfolioAllocation, error)
	FindAvailableObservationTimestamps(
		portfolioId int,
		observationTimestampsLimit int,
	) ([]*PortfolioObservationTimestamp, error)
	FindAvailablePortfolioAllocationClasses(portfolioId int) ([]string, error)
	MergePortfolioAllocationsInTransaction(
		transContext context.Context,
		portfolioId int,
		observationTimestamp *PortfolioObservationTimestamp,
		allocations []*PortfolioAllocation,
	) error
	InsertObservationTimestampInTransaction(
		transContext context.Context,
		observationTimestamp *PortfolioObservationTimestamp,
	) (*PortfolioObservationTimestamp, error)
}
