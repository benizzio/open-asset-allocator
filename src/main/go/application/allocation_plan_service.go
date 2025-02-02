package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
)

// TODO move this to a domain service
type AllocationPlanAppService struct {
	allocationPlanRepository domain.AllocationPlanRepository
}

func (service *AllocationPlanAppService) GetAllocationPlans(portfolioId int, planType *allocation.PlanType) (
	[]*domain.AllocationPlan,
	error,
) {
	return service.allocationPlanRepository.GetAllAllocationPlans(portfolioId, planType)
}

func BuildAllocationPlanAppService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanAppService {
	return &AllocationPlanAppService{allocationPlanRepository}
}
