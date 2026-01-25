package service

import (
	"context"
	"sort"
	"strconv"
	"strings"

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

type allocationPlanValidationData struct {
	hierarchicalIdCounts           map[string]int
	levelSliceSizes                map[string]*levelSliceSizeValidationData
	hierarchicalAllocationPlanTree *langext.MapTreeNode[string]
	invalidSizeHierarchyBranches   langext.CustomSliceTable[string]
	noParentHierarchyBranches      langext.CustomSliceTable[string]
	childlessHierarchyBranches     langext.CustomSliceTable[string]
}

type levelSliceSizeValidationData struct {
	levelIndex int
	sliceSize  decimal.Decimal
}

func (validationData *levelSliceSizeValidationData) describeLevel(
	hierarchyLevels domain.AllocationHierarchy,
	levelId string,
) string {
	var hierarchySize = len(hierarchyLevels)
	if validationData.levelIndex == -1 {
		return hierarchyLevels[hierarchySize-1].Name + " (TOP)"
	}
	if validationData.levelIndex < 0 || validationData.levelIndex >= hierarchySize {
		return "Unknown hierarchy level" + " = " + levelId + " (index " + strconv.Itoa(validationData.levelIndex) + ")"
	}
	return hierarchyLevels[validationData.levelIndex].Name + " = " + levelId
}

// TODO validate plan before persisting
// - hierarchical ids with asset match asset ticker from reference
func (service *AllocationPlanDomService) validateAllocationPlan(
	plan *domain.AllocationPlan,
	allocationStructure *domain.AllocationStructure,
) []*infra.AppError {

	var hierarchyLevels = allocationStructure.Hierarchy

	// TODO add a matrix representing the hierarchy structure to validate completeness of levels
	// "no parent" can be added to the matrix with a tag
	var validationData = &allocationPlanValidationData{
		hierarchicalIdCounts:           make(map[string]int),
		levelSliceSizes:                make(map[string]*levelSliceSizeValidationData),
		hierarchicalAllocationPlanTree: langext.NewMapTree("ALLOCATION_PLAN_ROOT"),
		invalidSizeHierarchyBranches:   make(langext.CustomSliceTable[string], 0),
		childlessHierarchyBranches:     make(langext.CustomSliceTable[string], 0),
		noParentHierarchyBranches:      make(langext.CustomSliceTable[string], 0),
	}
	// Use empty string key to track TOP-level percentage aggregation
	validationData.levelSliceSizes[""] = &levelSliceSizeValidationData{
		levelIndex: -1,
		sliceSize:  decimal.Zero,
	}

	readValidationData(hierarchyLevels, plan, validationData)

	var errors = make([]*infra.AppError, 0)
	errors = service.validateHierarchicalIdUniqueness(validationData, errors)
	errors = service.validateHierarchyLevelsSliceSizeSums(hierarchyLevels, validationData, errors)
	errors = service.validateHierarchyBranchesCompleteness(hierarchyLevels, validationData, errors)

	if len(errors) > 0 {
		return errors
	}

	return nil
}

func readValidationData(
	hierarchyLevels domain.AllocationHierarchy,
	plan *domain.AllocationPlan,
	validation *allocationPlanValidationData,
) {

	for _, plannedAllocation := range plan.Details {
		readPlannedAllocationForRepeatedValidationData(plannedAllocation, validation)
		readPlannedAllocationForSliceSizeTotalsValidationData(plannedAllocation, validation)
		readPlannedAllocationHierarchicalBranchValidationData(hierarchyLevels, plannedAllocation, validation)
	}

	readPlannedAllocationChildlessHierarchyBranchesValidationData(hierarchyLevels, validation)
}

// readPlannedAllocationForRepeatedValidationData reads counts of hierarchical ids to validate repetitions
func readPlannedAllocationForRepeatedValidationData(
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationData,
) {
	var hierarchicalIdString = plannedAllocation.HierarchicalId.String()
	var count, exists = validation.hierarchicalIdCounts[hierarchicalIdString]
	if !exists {
		validation.hierarchicalIdCounts[hierarchicalIdString] = 1
	} else {
		validation.hierarchicalIdCounts[hierarchicalIdString] = count + 1
	}
}

// readPlannedAllocationForSliceSizeTotalsValidationData aggregates slice size percentages per hierarchy level for validations
func readPlannedAllocationForSliceSizeTotalsValidationData(
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationData,
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
			validation.levelSliceSizes[parentHierarchicalIdString] = &levelSliceSizeValidationData{
				levelIndex: plannedAllocation.HierarchicalId.GetParentLevelIndex(),
				sliceSize:  plannedAllocation.SliceSizePercentage,
			}
		} else {
			levelValidationData.sliceSize = levelValidationData.sliceSize.Add(plannedAllocation.SliceSizePercentage)
		}
	}
}

func readPlannedAllocationHierarchicalBranchValidationData(
	hierarchy domain.AllocationHierarchy,
	plannedAllocation *domain.PlannedAllocation,
	validation *allocationPlanValidationData,
) {

	var hierarchySize = len(hierarchy)
	var plannedAllocationBranchInverted = langext.DereferenceSliceContent(plannedAllocation.HierarchicalId)
	var plannedAllocationBranch = langext.ReverseSlice(plannedAllocationBranchInverted)

	validateHierarchySize(validation, hierarchySize, plannedAllocationBranch)

	validateOrphanBranch(validation, plannedAllocationBranch)
}

func validateHierarchySize(
	validation *allocationPlanValidationData,
	hierarchySize int,
	plannedAllocationBranch []string,
) {
	var branchSize = len(plannedAllocationBranch)
	if branchSize != hierarchySize {
		validation.invalidSizeHierarchyBranches = append(
			validation.invalidSizeHierarchyBranches,
			plannedAllocationBranch,
		)
	}
}

func validateOrphanBranch(validation *allocationPlanValidationData, plannedAllocationBranch []string) {

	var branchSize = len(plannedAllocationBranch)
	var previousLevelWasZeroValue = langext.IsZeroValue(plannedAllocationBranch[0])

	if previousLevelWasZeroValue {
		validation.noParentHierarchyBranches = append(validation.noParentHierarchyBranches, plannedAllocationBranch)
	} else {

		for i := 1; i < branchSize; i++ {

			var currentLevelIsZeroValue = langext.IsZeroValue(plannedAllocationBranch[i])

			if previousLevelWasZeroValue && !currentLevelIsZeroValue {
				validation.noParentHierarchyBranches = append(
					validation.noParentHierarchyBranches,
					plannedAllocationBranch,
				)
				break
			}

			previousLevelWasZeroValue = currentLevelIsZeroValue
		}
	}

	validation.hierarchicalAllocationPlanTree.AddBranchBreakingOnZeroValues(plannedAllocationBranch)
}

func readPlannedAllocationChildlessHierarchyBranchesValidationData(
	hierarchy domain.AllocationHierarchy,
	validation *allocationPlanValidationData,
) {

	var hierarchySize = len(hierarchy)
	var allBranches = validation.hierarchicalAllocationPlanTree.ExtractBranches()

	sort.Slice(
		allBranches, func(i, j int) bool {
			return strings.Join(allBranches[i], "|") < strings.Join(allBranches[j], "|")
		},
	)

	for _, branch := range allBranches {
		// Branch length includes the root node, so a complete branch has hierarchySize + 1 elements
		if len(branch) <= hierarchySize {
			validation.childlessHierarchyBranches = append(validation.childlessHierarchyBranches, branch)
		}
	}
}

func (service *AllocationPlanDomService) validateHierarchicalIdUniqueness(
	validationData *allocationPlanValidationData,
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
	validation *allocationPlanValidationData,
	errors []*infra.AppError,
) []*infra.AppError {

	var exceededLevelDescriptions = make(langext.CustomSlice[string], 0)
	var insufficientLevelDescriptions = make(langext.CustomSlice[string], 0)

	for hierarchicalId, validationData := range validation.levelSliceSizes {

		if validationData.sliceSize.GreaterThan(decimal.NewFromInt(1)) {
			exceededLevelDescriptions = appendLevelDescription(
				validationData,
				exceededLevelDescriptions,
				hierarchyLevels,
				hierarchicalId,
			)
		} else if validationData.sliceSize.LessThan(decimal.NewFromInt(1)) {
			insufficientLevelDescriptions = appendLevelDescription(
				validationData,
				insufficientLevelDescriptions,
				hierarchyLevels,
				hierarchicalId,
			)
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

	if len(insufficientLevelDescriptions) > 0 {
		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations slice sizes sum to less than 100%% within hierarchy level(s): %s",
				insufficientLevelDescriptions.PrettyString(),
			),
		)
	}

	return errors
}

func (service *AllocationPlanDomService) validateHierarchyBranchesCompleteness(
	hierarchyLevels domain.AllocationHierarchy,
	validationData *allocationPlanValidationData,
	errors []*infra.AppError,
) []*infra.AppError {

	var userFriendlyHierarchyLevels = domain.AllocationHierarchy(langext.ReverseSlice(hierarchyLevels))

	if len(validationData.invalidSizeHierarchyBranches) > 0 {
		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations contain hierarchy branches with invalid size: \n%s\n for portfolio hierarchy: %s",
				validationData.invalidSizeHierarchyBranches.ArrowString(),
				userFriendlyHierarchyLevels.PrettyString(),
			),
		)
	}

	if len(validationData.noParentHierarchyBranches) > 0 {
		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations contain hierarchy branches with missing parent levels: \n%s\n for portfolio hierarchy: %s",
				validationData.noParentHierarchyBranches.ArrowString(),
				userFriendlyHierarchyLevels.PrettyString(),
			),
		)
	}

	if len(validationData.childlessHierarchyBranches) > 0 {

		var userFriendlyChildlessBranches = stripRootFromBranches(validationData.childlessHierarchyBranches)

		errors = append(
			errors,
			infra.BuildAppErrorFormattedUnconverted(
				service,
				"Planned allocations contain hierarchy branches with missing child levels: \n%s\n for portfolio hierarchy: %s",
				userFriendlyChildlessBranches.ArrowString(),
				userFriendlyHierarchyLevels.PrettyString(),
			),
		)
	}

	return errors
}

// stripRootFromBranches removes the first element (root node) from each branch for user-friendly display.
//
// Authored by: GitHub Copilot
func stripRootFromBranches(branches langext.CustomSliceTable[string]) langext.CustomSliceTable[string] {

	var result = make(langext.CustomSliceTable[string], 0, len(branches))

	for _, branch := range branches {
		if len(branch) > 1 {
			result = append(result, branch[1:])
		}
	}

	return result
}

func appendLevelDescription(
	validationData *levelSliceSizeValidationData,
	levelDescriptionsSlice langext.CustomSlice[string],
	hierarchyLevels domain.AllocationHierarchy,
	hierarchicalId string,
) langext.CustomSlice[string] {
	var levelDescription = validationData.describeLevel(hierarchyLevels, hierarchicalId)
	var sliceSizePercentage = validationData.sliceSize.Mul(decimal.NewFromInt(100)).String()
	return append(levelDescriptionsSlice, levelDescription+" ("+sliceSizePercentage+"%)")
}

func BuildAllocationPlanDomService(allocationPlanRepository domain.AllocationPlanRepository) *AllocationPlanDomService {
	return &AllocationPlanDomService{allocationPlanRepository}
}
