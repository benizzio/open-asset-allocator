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
	Id        int64
	TimeTag   string
	Timestamp time.Time
}

type PortfolioAllocationRepository interface {
	FindAllPortfolioAllocationsWithinObservationTimestampsLimit(
		id int64,
		observationTimestampsLimit int,
	) ([]*PortfolioAllocation, error)
	FindPortfolioAllocationsByObservationTimestamp(id int64, observationTimestampId int64) ([]*PortfolioAllocation, error)
	FindAvailableObservationTimestamps(
		portfolioId int64,
		observationTimestampsLimit int,
	) ([]*PortfolioObservationTimestamp, error)
	FindAvailablePortfolioAllocationClasses(portfolioId int64) ([]string, error)
	MergePortfolioAllocationsInTransaction(
		transContext context.Context,
		portfolioId int64,
		observationTimestamp *PortfolioObservationTimestamp,
		allocations []*PortfolioAllocation,
	) error
	InsertObservationTimestampInTransaction(
		transContext context.Context,
		observationTimestamp *PortfolioObservationTimestamp,
	) (*PortfolioObservationTimestamp, error)
}
