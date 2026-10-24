package domain

import "time"

// ObservationPageCursor identifies the last emitted observation in a descending timestamp/ID traversal.
// The timestamp and ID must be supplied together when requesting a subsequent page.
// For example, a repository can use `ObservationPageCursor{Timestamp: timestamp, Id: 42}` as a keyset boundary.
// Authored by: OpenCode
type ObservationPageCursor struct {
	Timestamp time.Time
	Id        int64
}

// PortfolioObservationPage contains observation headers and the boundary for its next page, if one exists.
// For example, an empty Items slice with a nil NextCursor represents an empty paginated result.
// Authored by: OpenCode
type PortfolioObservationPage struct {
	Items      []*PortfolioObservationTimestamp
	NextCursor *ObservationPageCursor
}

// PortfolioAllocationHistoryPage contains complete allocation rows for the emitted observations.
// NextCursor is based on the last emitted observation, not the look-ahead observation.
// For example, consumers can aggregate Allocations by observation and request the next page with NextCursor.
// Authored by: OpenCode
type PortfolioAllocationHistoryPage struct {
	Allocations []*PortfolioAllocation
	NextCursor  *ObservationPageCursor
}

// AllocationPlanPage contains complete plans and the boundary for the next descending planned-date/ID page.
// For example, an empty Items slice and nil NextCursor represents a terminal empty page.
// Authored by: OpenCode
type AllocationPlanPage struct {
	Items      []*AllocationPlan
	NextCursor *AllocationPlanPageCursor
}

// AllocationPlanPageCursor identifies the last emitted plan in a descending planned-date/ID traversal.
// For example, a repository can use `AllocationPlanPageCursor{PlannedDate: date, Id: 9}` as a keyset boundary.
// Authored by: OpenCode
type AllocationPlanPageCursor struct {
	PlannedDate time.Time
	Id          int64
}
