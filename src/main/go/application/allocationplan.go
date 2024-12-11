package application

import "github.com/benizzio/open-asset-allocator/domain"

type AllocationPlanService struct {
	allocationPlanRepository domain.AllocationPlanRepository
}

func (service *AllocationPlanService) GetAllocationPlans() ([]*domain.AllocationPlan, error) {
	return service.allocationPlanRepository.GetAllAllocationPlans()
}

func BuildAllocationPlanService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanService {
	return &AllocationPlanService{allocationPlanRepository}
}
