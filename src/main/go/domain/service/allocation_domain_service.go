package service

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type AllocationDomService struct {
	allocationRepository domain.AllocationRepository
}

// FindAvailableAllocationClassesFromAllSources retrieves allocation classes from
// both portfolio allocation history and planned allocations in allocation plans.
//
// Authored by: GitHub Copilot
func (service *AllocationDomService) FindAvailableAllocationClassesFromAllSources(
	portfolioId int64,
) ([]string, error) {
	return service.allocationRepository.FindAvailableAllocationClassesFromAllSources(portfolioId)
}

// BuildAllocationDomService creates a new AllocationDomService instance.
//
// Authored by: GitHub Copilot
func BuildAllocationDomService(allocationRepository domain.AllocationRepository) *AllocationDomService {
	return &AllocationDomService{
		allocationRepository,
	}
}
