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

// ObservationPageCursor identifies the last emitted observation in a descending timestamp/ID traversal.
// The timestamp and ID must be supplied together when requesting a subsequent page.
// For example, a repository can use `ObservationPageCursor{Timestamp: timestamp, Id: 42}` as a keyset boundary.
// Co-authored by: GPT-6.1 Sol and OpenCode
type ObservationPageCursor struct {
	Timestamp time.Time
	Id        int64
}

// PortfolioObservationPage contains observation headers and the boundary for its next page, if one exists.
// For example, an empty Items slice with a nil NextCursor represents an empty paginated result.
// Co-authored by: GPT-6.1 Sol and OpenCode
type PortfolioObservationPage struct {
	Items      []*PortfolioObservationTimestamp
	NextCursor *ObservationPageCursor
}

// PortfolioAllocationHistoryPage contains complete allocation rows for the emitted observations.
// NextCursor is based on the last emitted observation, not the look-ahead observation.
// For example, consumers can aggregate Allocations by observation and request the next page with NextCursor.
// Co-authored by: GPT-6.1 Sol and OpenCode
type PortfolioAllocationHistoryPage struct {
	Allocations []*PortfolioAllocation
	NextCursor  *ObservationPageCursor
}

// PortfolioAllocationRepository provides persistence operations used by the portfolio allocation domain service.
// Co-authored by: GPT-6.1 Sol and benizzio
type PortfolioAllocationRepository interface {
	// FindAndLockObservationTimestampInTransaction reads and locks an observation for domain ownership validation.
	// For example, FindAndLockObservationTimestampInTransaction(ctx, 2) returns nil if observation 2 is missing.
	// Authored by: GPT-6.1 Sol
	FindAndLockObservationTimestampInTransaction(context.Context, int64) (*PortfolioObservationTimestamp, error)
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
	// FindAvailableObservationTimestampsPage returns up to ten headers using a timestamp/ID boundary.
	// For example, pass nil for the first page and the last response's cursor for continuation.
	FindAvailableObservationTimestampsPage(
		portfolioId int64,
		cursor *ObservationPageCursor,
	) (*PortfolioObservationPage, error)
	// FindPortfolioAllocationHistoryPage returns complete snapshots for up to ten observations.
	// For example, pass nil for the first page and the last response's cursor for continuation.
	FindPortfolioAllocationHistoryPage(
		portfolioId int64,
		cursor *ObservationPageCursor,
	) (*PortfolioAllocationHistoryPage, error)
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
