package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"
)

// PortfolioAllocation represents the observed allocation of a single asset in a portfolio
// snapshot. SelectedExternalAsset stores the first persisted external reference projected for
// read operations, keeping it separate from the full persisted asset external data.
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
type PortfolioAllocation struct {
	Asset                 Asset
	SelectedExternalAsset *ExternalAsset
	Class                 string
	CashReserve           bool
	ObservationTimestamp  *PortfolioObservationTimestamp
	TotalMarketValue      int64
	AssetQuantity         decimal.Decimal
	AssetMarketPrice      decimal.Decimal
}

// PortfolioObservationTimestamp identifies a portfolio-owned observation and its editable metadata.
//
// Co-authored by: OpenCode and benizzio
type PortfolioObservationTimestamp struct {
	Id          int64
	PortfolioId int64
	TimeTag     string
	Timestamp   time.Time
}

type PortfolioAllocationRepository interface {
	FindAllPortfolioAllocationsWithinObservationTimestampsLimit(
		id int64,
		observationTimestampsLimit int,
	) ([]*PortfolioAllocation, error)
	FindPortfolioAllocationsByObservationTimestamp(id int64, observationTimestampId int64) (
		[]*PortfolioAllocation,
		error,
	)
	FindAvailableObservationTimestamps(
		portfolioId int64,
		observationTimestampsLimit int,
	) ([]*PortfolioObservationTimestamp, error)
	MergePortfolioAllocationsInTransaction(
		transContext context.Context,
		portfolioId int64,
		observationTimestamp *PortfolioObservationTimestamp,
		allocations []*PortfolioAllocation,
	) error
	InsertObservationTimestampInTransaction(
		transContext context.Context,
		portfolioId int64,
		observationTimestamp *PortfolioObservationTimestamp,
	) (*PortfolioObservationTimestamp, error)
	// UpdateObservationTimestampInTransaction changes metadata for an observation already owned by the portfolio.
	// For example, `repository.UpdateObservationTimestampInTransaction(ctx, 42, observation)` updates portfolio 42.
	// Authored by: OpenCode
	UpdateObservationTimestampInTransaction(
		transContext context.Context,
		portfolioId int64,
		observationTimestamp *PortfolioObservationTimestamp,
	) error
}
