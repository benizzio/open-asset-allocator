package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
)

type AllocationPlanService struct {
	allocationPlanRepository domain.AllocationPlanRepository
}

func (service *AllocationPlanService) GetAllocationPlans(planType *allocation.PlanType) (
	[]*domain.AllocationPlan,
	error,
) {
	return service.allocationPlanRepository.GetAllAllocationPlans(planType)
}

func BuildAllocationPlanService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanService {
	return &AllocationPlanService{allocationPlanRepository}
}
