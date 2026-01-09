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
	allocationStructure *domain.AllocationStructure,
) error {

	var validationErrors = service.validateAllocationPlan(plan, allocationStructure)
	if validationErrors != nil {
		return infra.BuildDomainValidationError("Allocation plan validation failed", validationErrors)
	}

	if langext.IsZeroValue(plan.Id) {
		return service.allocationPlanRepository.InsertAllocationPlanInTransaction(transContext, plan)
	}

	return service.allocationPlanRepository.UpdateAllocationPlanInTransaction(transContext, plan)

}

type allocationPlanValidationValidationData struct {
	hierarchicalIdCounts map[string]int
	levelSliceSizes      map[string]*levelSliceSizeValidadationData
}

type levelSliceSizeValidadationData struct {
	levelIndex int
	sliceSize  decimal.Decimal
}

func (validationData *levelSliceSizeValidadationData) describeLevel(
	hierarchyLevels domain.AllocationHierarchy,
	levelId string,
) string {
	if validationData.levelIndex == -1 {
		return hierarchyLevels[len(hierarchyLevels)-1].Name + " (TOP)"
	}
	return hierarchyLevels[validationData.levelIndex].Name + ": " + levelId
}

// TODO validate plan before persisting
// - hierarchy matches portfolio hierarchy
// - hierarchical ids with asset match asset ticker from reference
// - slice sizes sum to 100% per hierarchy level
func (service *AllocationPlanDomService) validateAllocationPlan(
	plan *domain.AllocationPlan,
	allocationStructure *domain.AllocationStructure,
) []*infra.AppError {

	var hierarchyLevels = allocationStructure.Hierarchy

	var validationData = &allocationPlanValidationValidationData{
		hierarchicalIdCounts: make(map[string]int),
		levelSliceSizes:      make(map[string]*levelSliceSizeValidadationData),
	}
	// Use empty string key to track TOP-level percentage aggregation
	validationData.levelSliceSizes[""] = &levelSliceSizeValidadationData{
		levelIndex: hierarchyLevels.Size() - 1,
		sliceSize:  decimal.Zero,
	}

	readValidationData(plan, validationData)

	var errors = make([]*infra.AppError, 0)
	errors = service.validateHierarchicalIdUniqueness(validationData, errors)
	errors = service.validateHierarchyLevelsSliceSizeSums(hierarchyLevels, validationData, errors)

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func readValidationData(plan *domain.AllocationPlan, validation *allocationPlanValidationValidationData) {
	for _, plannedAllocation := range plan.Details {
		readPlannedAllocationForRepeatedValidationData(plannedAllocation, validation)
		readPlannedAllocationForSliceSizeTotals(plannedAllocation, validation)
	}
}

// readPlannedAllocationForRepeatedValidationData reads counts of hierarchical ids to validate repetitions
func readPlannedAllocationForRepeatedValidationData(
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationValidationData,
) {
	var hierarchicalIdString = plannedAllocation.HierarchicalId.String()
	var count, exists = validation.hierarchicalIdCounts[hierarchicalIdString]
	if !exists {
		validation.hierarchicalIdCounts[hierarchicalIdString] = 1
	} else {
		validation.hierarchicalIdCounts[hierarchicalIdString] = count + 1
	}
}

// readPlannedAllocationForSliceSizeTotals aggregates slice size percentages per hierarchy level for validations
func readPlannedAllocationForSliceSizeTotals(
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationValidationData,
) {

	// Aggregate percentages per hierarchy level:
	// - TOP level: accumulate into "" key
	// - Child levels: accumulate into the parent hierarchical id (drop the lowest level element)
	var parentHierarchicalId = plannedAllocation.HierarchicalId.ParentLevelId()
	if parentHierarchicalId == nil {

		var topLevelValidationData = validation.levelSliceSizes[""]
		topLevelValidationData.sliceSize = topLevelValidationData.sliceSize.Add(plannedAllocation.SliceSizePercentage)
		topLevelValidationData.levelIndex = plannedAllocation.HierarchicalId.GetParentLevelIndex()

	} else {

		var parentHierarchicalIdString = parentHierarchicalId.String()
		var levelValidationData, levelExists = validation.levelSliceSizes[parentHierarchicalIdString]

		if !levelExists {
			validation.levelSliceSizes[parentHierarchicalIdString] = &levelSliceSizeValidadationData{
				levelIndex: plannedAllocation.HierarchicalId.GetParentLevelIndex(),
				sliceSize:  plannedAllocation.SliceSizePercentage,
			}
		} else {
			levelValidationData.sliceSize = levelValidationData.sliceSize.Add(plannedAllocation.SliceSizePercentage)
		}
	}
}

func (service *AllocationPlanDomService) validateHierarchicalIdUniqueness(
	validationData *allocationPlanValidationValidationData,
	errors []*infra.AppError,
) []*infra.AppError {

	var repeatedHierarchicalIds = make(langext.CustomSlice[string], 0)
	for hierarchicalId, count := range validationData.hierarchicalIdCounts {
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

func (service *AllocationPlanDomService) validateHierarchyLevelsSliceSizeSums(
	hierarchyLevels domain.AllocationHierarchy,
	validation *allocationPlanValidationValidationData,
	errors []*infra.AppError,
) []*infra.AppError {

	var exceededLevelDescriptions = make(langext.CustomSlice[string], 0)
	for hierarchicalId, validationData := range validation.levelSliceSizes {
		if validationData.sliceSize.GreaterThan(decimal.NewFromInt(1)) {
			var levelDescription = validationData.describeLevel(hierarchyLevels, hierarchicalId)
			exceededLevelDescriptions = append(exceededLevelDescriptions, levelDescription)
		}
	}

	if len(exceededLevelDescriptions) > 0 {
		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations slice sizes exceed 100%% within hierarchy level(s): %s",
				exceededLevelDescriptions.PrettyString(),
			),
		)
	}
	return errors
}

func BuildAllocationPlanDomService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanDomService {
	return &AllocationPlanDomService{allocationPlanRepository}
}
