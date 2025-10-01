package service

import (
	"context"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/langext"
)

type AllocationPlanDomService struct {
	allocationPlanRepository domain.AllocationPlanRepository
}

func (service *AllocationPlanDomService) GetAllocationPlans(portfolioId int64, planType *allocation.PlanType) (
	[]*domain.AllocationPlan,
	error,
) {
	return service.allocationPlanRepository.GetAllAllocationPlans(portfolioId, planType)
}

func (service *AllocationPlanDomService) GetAllocationPlan(id int64) (*domain.AllocationPlan, error) {
	return service.allocationPlanRepository.GetAllocationPlan(id)
}

func (service *AllocationPlanDomService) GetAllAllocationPlanIdentifiers(
	portfolioId int64,
	planType *allocation.PlanType,
) ([]*domain.AllocationPlanIdentifier, error) {
	return service.allocationPlanRepository.GetAllAllocationPlanIdentifiers(portfolioId, planType)
}

func (service *AllocationPlanDomService) GetPlannedAllocationsPerHyerarchicalIdMap(allocationPlanId int64) (
	domain.PlannedAllocationsPerHierarchicalId,
	error,
) {
	allocationPlan, err := service.GetAllocationPlan(allocationPlanId)
	if err != nil {
		return nil, err
	}

	plannedAllocationMap := make(domain.PlannedAllocationsPerHierarchicalId)

	for _, plannedAllocation := range allocationPlan.Details {
		hierarchicalId := plannedAllocation.HierarchicalId.String()
		plannedAllocationMap[hierarchicalId] = plannedAllocation
	}

	return plannedAllocationMap, nil
}

func (service *AllocationPlanDomService) PersistAllocationPlanInTransaction(
	transContext context.Context,
	plan *domain.AllocationPlan,
) error {

	// TODO validate plan before persisting
	// - unique hierarchical ids
	// - sum of slice size percentages inside parent in hierarchy <= 100%
	// - hierarchy matches portfolio hierarchy

	if langext.IsZeroValue(plan.Id) {
		return service.allocationPlanRepository.InsertAllocationPlanInTransaction(transContext, plan)
	} else {
		return service.allocationPlanRepository.UpdateAllocationPlanInTransaction(transContext, plan)
	}
}

func BuildAllocationPlanDomService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanDomService {
	return &AllocationPlanDomService{allocationPlanRepository}
}
