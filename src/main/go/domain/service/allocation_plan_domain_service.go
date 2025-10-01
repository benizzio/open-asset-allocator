package service

import (
	"context"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/shopspring/decimal"
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

	var validationErrors = service.validateAllocationPlan(plan)
	if validationErrors != nil {
		return infra.BuildDomainValidationError("Allocation plan validation failed", validationErrors)
	}

	if langext.IsZeroValue(plan.Id) {
		return service.allocationPlanRepository.InsertAllocationPlanInTransaction(transContext, plan)
	} else {
		return service.allocationPlanRepository.UpdateAllocationPlanInTransaction(transContext, plan)
	}

}

type AllocationPlanValidationValidation struct {
	hirerarchicalIdCounts map[string]int
	levelPercentages      map[string]decimal.Decimal
}

// TODO clean code
// TODO validate plan before persisting
// - hierarchy matches portfolio hierarchy
func (service *AllocationPlanDomService) validateAllocationPlan(plan *domain.AllocationPlan) []*infra.AppError {

	var validation = &AllocationPlanValidationValidation{
		hirerarchicalIdCounts: make(map[string]int),
		levelPercentages:      make(map[string]decimal.Decimal),
	}

	// Use empty string key to track TOP-level percentage aggregation
	validation.levelPercentages[""] = decimal.Zero

	for _, plannedAllocation := range plan.Details {

		// Track duplicates by the full hierarchical id path
		var hierarchicalIdString = plannedAllocation.HierarchicalId.String()

		var count int
		var exists bool
		count, exists = validation.hirerarchicalIdCounts[hierarchicalIdString]
		if !exists {
			validation.hirerarchicalIdCounts[hierarchicalIdString] = 1
		} else {
			validation.hirerarchicalIdCounts[hierarchicalIdString] = count + 1
		}

		// Aggregate percentages per hierarchy level:
		// - TOP level: accumulate into "" key
		// - Child levels: accumulate into the parent hierarchical id (drop the lowest level element)
		var parentHierarchicalId = plannedAllocation.HierarchicalId.ParentLevelId()
		if parentHierarchicalId == nil {
			validation.levelPercentages[""] = validation.levelPercentages[""].Add(plannedAllocation.SliceSizePercentage)
		} else {

			var parentHierarchicalIdString = parentHierarchicalId.String()

			var percentageSum decimal.Decimal
			var levelExists bool
			percentageSum, levelExists = validation.levelPercentages[parentHierarchicalIdString]
			if !levelExists {
				validation.levelPercentages[parentHierarchicalIdString] = plannedAllocation.SliceSizePercentage
			} else {
				validation.levelPercentages[parentHierarchicalIdString] = percentageSum.Add(plannedAllocation.SliceSizePercentage)
			}
		}
	}

	var repeatedHierarchicalIds = make(langext.CustomSlice[string], 0)
	var exceededLevels = make(langext.CustomSlice[string], 0)
	var errors = make([]*infra.AppError, 0)

	for hierarchicalId, count := range validation.hirerarchicalIdCounts {
		if count > 1 {
			repeatedHierarchicalIds = append(repeatedHierarchicalIds, hierarchicalId)
		}
	}

	if len(repeatedHierarchicalIds) > 0 {
		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations contain duplicated hierarchical IDs: %s",
				repeatedHierarchicalIds.PrettyString(),
			),
		)
	}

	for hierarchicalId, percentageSum := range validation.levelPercentages {
		if percentageSum.GreaterThan(decimal.NewFromInt(1)) {
			if hierarchicalId == "" {
				hierarchicalId = "TOP"
			}
			exceededLevels = append(exceededLevels, hierarchicalId)
		}
	}

	if len(exceededLevels) > 0 {
		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations slice sizes exceed 100%% within hierarchy level(s): %s",
				exceededLevels.PrettyString(),
			),
		)
	}

	if len(errors) > 0 {
		return errors
	} else {
		return nil
	}
}

func BuildAllocationPlanDomService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanDomService {
	return &AllocationPlanDomService{allocationPlanRepository}
}
