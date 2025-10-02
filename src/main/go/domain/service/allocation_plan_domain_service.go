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

type allocationPlanValidationValidationData struct {
	hirerarchicalIdCounts map[string]int
	levelPercentages      map[string]decimal.Decimal
}

// TODO validate plan before persisting
// - hierarchy matches portfolio hierarchy
func (service *AllocationPlanDomService) validateAllocationPlan(plan *domain.AllocationPlan) []*infra.AppError {

	var validation = &allocationPlanValidationValidationData{
		hirerarchicalIdCounts: make(map[string]int),
		levelPercentages:      make(map[string]decimal.Decimal),
	}
	// Use empty string key to track TOP-level percentage aggregation
	validation.levelPercentages[""] = decimal.Zero

	readValidationData(plan, validation)

	var errors = make([]*infra.AppError, 0)
	errors = service.validateHirearchicalIdUniqueness(validation, errors)
	errors = service.validateHierarchyLevelsPercentageSums(validation, errors)

	if len(errors) > 0 {
		return errors
	} else {
		return nil
	}
}

func (service *AllocationPlanDomService) validateHierarchyLevelsPercentageSums(
	validation *allocationPlanValidationValidationData,
	errors []*infra.AppError,
) []*infra.AppError {

	// TODO obtain hierarchy level names to print proper messages
	var exceededLevels = make(langext.CustomSlice[string], 0)
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
	return errors
}

func readValidationData(plan *domain.AllocationPlan, validation *allocationPlanValidationValidationData) {
	for _, plannedAllocation := range plan.Details {
		readPlannedAllocationForRepeatedValidationData(plannedAllocation, validation)
		readPlannedAllocationForPercentageTotals(plannedAllocation, validation)
	}
}

// readPlannedAllocationForRepeatedValidationData reads counts of hierarchical ids to validate repetitions
func readPlannedAllocationForRepeatedValidationData(
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationValidationData,
) {
	var hierarchicalIdString = plannedAllocation.HierarchicalId.String()
	var count, exists = validation.hirerarchicalIdCounts[hierarchicalIdString]
	if !exists {
		validation.hirerarchicalIdCounts[hierarchicalIdString] = 1
	} else {
		validation.hirerarchicalIdCounts[hierarchicalIdString] = count + 1
	}
}

// readPlannedAllocationForPercentageTotals aggregates slice size percentages per hierarchy level to validate they do not exceed 100%
func readPlannedAllocationForPercentageTotals(
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationValidationData,
) {

	// Aggregate percentages per hierarchy level:
	// - TOP level: accumulate into "" key
	// - Child levels: accumulate into the parent hierarchical id (drop the lowest level element)
	var parentHierarchicalId = plannedAllocation.HierarchicalId.ParentLevelId()
	if parentHierarchicalId == nil {
		validation.levelPercentages[""] = validation.levelPercentages[""].Add(plannedAllocation.SliceSizePercentage)
	} else {

		var parentHierarchicalIdString = parentHierarchicalId.String()
		var percentageSum, levelExists = validation.levelPercentages[parentHierarchicalIdString]

		if !levelExists {
			validation.levelPercentages[parentHierarchicalIdString] = plannedAllocation.SliceSizePercentage
		} else {
			validation.levelPercentages[parentHierarchicalIdString] = percentageSum.Add(plannedAllocation.SliceSizePercentage)
		}
	}
}

func (service *AllocationPlanDomService) validateHirearchicalIdUniqueness(
	validationData *allocationPlanValidationValidationData,
	errors []*infra.AppError,
) []*infra.AppError {

	var repeatedHierarchicalIds = make(langext.CustomSlice[string], 0)
	for hierarchicalId, count := range validationData.hirerarchicalIdCounts {
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

	return errors
}

func BuildAllocationPlanDomService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanDomService {
	return &AllocationPlanDomService{allocationPlanRepository}
}
