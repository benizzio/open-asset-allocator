package domain

// AllocationRepository provides operations for querying allocation data across
// portfolio allocations and planned allocations.
//
// Authored by: GitHub Copilot
type AllocationRepository interface {
	FindAvailableAllocationClassesFromAllSources(portfolioId int64) ([]string, error)
}
