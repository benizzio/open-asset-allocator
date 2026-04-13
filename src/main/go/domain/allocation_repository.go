package domain

// AllocationRepository provides read operations for querying allocation data across portfolio
// allocations and planned allocations.
//
// Co-authored by: OpenCode and benizzio
type AllocationRepository interface {
	// FindAvailableAllocationClassesFromAllSources returns unique allocation class identifiers for
	// the given portfolio, sorted in ascending lexicographic order. It returns an empty slice and a
	// nil error when no classes are available. Query, scan, and iteration failures are returned as
	// errors.
	FindAvailableAllocationClassesFromAllSources(portfolioId int64) ([]string, error)
}
