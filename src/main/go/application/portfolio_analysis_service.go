package application

import (
	"context"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra/util"
	"github.com/shopspring/decimal"
)

type PortfolioAnalysisAppService struct {
	portfolioDomService      *service.PortfolioDomService
	allocationPlanDomService *service.AllocationPlanDomService
}

type potentialDivercencesPerHierarchicalId map[string]*domain.PotentialDivergence

func (service *PortfolioAnalysisAppService) GeneratePortfolioDivergenceAnalysis(
	portfolioId int,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) (*domain.DivergenceAnalysis, error) {

	var portfolio, portfolioAllocations, err = service.portfolioDomService.GetPortfolioSnapshot(
		portfolioId,
		timeFrameTag,
	)
	if err != nil {
		return nil, err
	}

	divergenceAnalysis, potentialDivergenceMap, err := service.generateDivergenceAnalysisFromPortfolioAllocationSet(
		portfolio,
		portfolioAllocations,
		timeFrameTag,
		allocationPlanId,
	)
	if err != nil {
		return nil, err
	}

	plannedAllocationMap, err := service.allocationPlanDomService.GetPlannedAllocationsPerHyerarchicalIdMap(allocationPlanId)
	if err != nil {
		return nil, err
	}

	calculateCurrentDivergenceValuesFromReferencedPlan(
		divergenceAnalysis.Root,
		plannedAllocationMap,
		divergenceAnalysis.PortfolioTotalMarketValue,
	)

	generatePotentialDivergencesFromAllocationPlanSetDifference(
		portfolio,
		divergenceAnalysis,
		potentialDivergenceMap,
		plannedAllocationMap,
	)

	return divergenceAnalysis, nil
}

func (service *PortfolioAnalysisAppService) generateDivergenceAnalysisFromPortfolioAllocationSet(
	portfolio *domain.Portfolio,
	portfolioAllocations []*domain.PortfolioAllocation,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) (*domain.DivergenceAnalysis, potentialDivercencesPerHierarchicalId, error) {

	var divergenceAnalysis = buildDivergenceAnalysis(portfolio, timeFrameTag, allocationPlanId)
	var analysisContext = buildDivergenceAnalysisContext(context.Background(), divergenceAnalysis)

	potentialDivergenceMap, err := service.mapPotentialDivergencesFromPortfolioAllocations(
		analysisContext,
		divergenceAnalysis,
		portfolio,
		portfolioAllocations,
	)
	if err != nil {
		return nil, nil, err
	}

	return divergenceAnalysis, potentialDivergenceMap, nil
}

func (service *PortfolioAnalysisAppService) mapPotentialDivergencesFromPortfolioAllocations(
	analysisContext context.Context,
	divergenceAnalysis *domain.DivergenceAnalysis,
	portfolio *domain.Portfolio,
	portfolioAllocations []*domain.PortfolioAllocation,
) (potentialDivercencesPerHierarchicalId, error) {

	var allocationHierarchy = portfolio.AllocationStructure.Hierarchy
	var potentialDivergenceMap = make(potentialDivercencesPerHierarchicalId)

	var iterationContextValue = &allocationIterationMappingContextValue{
		potentialDivergenceMap:       potentialDivergenceMap,
		portfolioAllocationsIterator: util.NewIterator(portfolioAllocations),
	}

	var allocationIterationMappingContext = buildAllocationIterationContext(analysisContext, iterationContextValue)

	for iterationContextValue.portfolioAllocationsIterator.HasNext() {

		var allocation, _ = iterationContextValue.portfolioAllocationsIterator.Next()

		err := service.mapPotentialDivergenceFromPortfolioAllocation(
			allocationIterationMappingContext,
			allocationHierarchy,
		)
		if err != nil {
			return nil, err
		}

		divergenceAnalysis.PortfolioTotalMarketValue += allocation.TotalMarketValue
	}

	return potentialDivergenceMap, nil
}

func (service *PortfolioAnalysisAppService) mapPotentialDivergenceFromPortfolioAllocation(
	analysisContext context.Context,
	allocationHierarchy domain.AllocationHierarchy,
) error {

	var allocationIterationContextValue = getAllocationIterationContextValue(analysisContext)

	var allocationHierarchyIterator = util.NewIterator(allocationHierarchy)
	var hierarchySubIterationMappingContext = buildHierarchySubIterationContext(
		analysisContext,
		allocationHierarchyIterator,
	)

	var potentialDivergencesInAllocationHierarchy = make([]*domain.PotentialDivergence, allocationHierarchy.Size())
	var lowerLevelPotentialDivergenceCreated = false

	for allocationHierarchyIterator.HasNext() {

		var _, allocationHierarchyLevelIndex = allocationHierarchyIterator.Next()

		var lowerLevelDivergence *domain.PotentialDivergence
		if lowerLevelPotentialDivergenceCreated {
			lowerLevelDivergence = potentialDivergencesInAllocationHierarchy[allocationHierarchyLevelIndex-1]
		}

		potentialDivergence, potentialDivergenceCreated, err := service.buildAndConnectPotentialDivergenceIfNotExists(
			hierarchySubIterationMappingContext,
			allocationHierarchy,
			lowerLevelDivergence,
		)
		if err != nil {
			return err
		}

		if potentialDivergenceCreated {
			potentialDivergencesInAllocationHierarchy[allocationHierarchyLevelIndex] = potentialDivergence
		}

		var currentAllocation, _ = allocationIterationContextValue.portfolioAllocationsIterator.Current()
		potentialDivergence.TotalMarketValue += currentAllocation.TotalMarketValue

		lowerLevelPotentialDivergenceCreated = potentialDivergenceCreated
	}

	return nil
}

func (service *PortfolioAnalysisAppService) buildAndConnectPotentialDivergenceIfNotExists(
	hierarchySubIterationMappingContext context.Context,
	allocationHierarchy domain.AllocationHierarchy,
	lowerLevelDivergence *domain.PotentialDivergence,
) (*domain.PotentialDivergence, bool, error) {

	var hierarchicalId, hierarchyLevelKey, err = service.generatePotentialDivergenceIdentifiers(
		hierarchySubIterationMappingContext,
		allocationHierarchy,
	)
	if err != nil {
		return nil, false, err
	}

	var potentialDivergence, potentialDivergenceCreated = buildAndAttachPotentialDivergenceIfNotExists(
		hierarchySubIterationMappingContext,
		hierarchicalId,
		hierarchyLevelKey,
	)

	if lowerLevelDivergence != nil {
		potentialDivergence.AddInternalDivergence(lowerLevelDivergence)
	}

	return potentialDivergence, potentialDivergenceCreated, nil
}

func (service *PortfolioAnalysisAppService) generatePotentialDivergenceIdentifiers(
	analysisContext context.Context,
	allocationHierarchy domain.AllocationHierarchy,
) (string, string, error) {

	var allocationIterationContextValue = getAllocationIterationContextValue(analysisContext)
	var currentAllocation, _ = allocationIterationContextValue.portfolioAllocationsIterator.Current()
	var currentHierarchyLevelIterator = getHierarchySubIterationContextValue(analysisContext)
	var currentHierarchyLevel, currentHierarchyLevelIndex = currentHierarchyLevelIterator.CurrentPointer()

	var hierarchicalId, err = service.portfolioDomService.GenerateHierarchicalId(
		currentAllocation,
		allocationHierarchy,
		currentHierarchyLevelIndex,
	)
	if err != nil {
		return "", "", err
	}
	//hierarchicalId += "a"

	hierarchyLevelKey, err := service.portfolioDomService.GetIdSegment(
		currentAllocation,
		currentHierarchyLevel,
	)
	if err != nil {
		return "", "", err
	}
	return hierarchicalId, hierarchyLevelKey, nil
}

func buildDivergenceAnalysis(
	portfolio *domain.Portfolio,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) *domain.DivergenceAnalysis {
	return &domain.DivergenceAnalysis{
		PortfolioId:               portfolio.Id,
		TimeFrameTag:              timeFrameTag,
		AllocationPlanId:          allocationPlanId,
		PortfolioTotalMarketValue: 0,
		Root:                      make([]*domain.PotentialDivergence, 0),
	}
}

func buildAndAttachPotentialDivergenceIfNotExists(
	analysisContext context.Context,
	hierarchicalId string,
	hierarchyLevelKey string,
) (*domain.PotentialDivergence, bool) {

	var allocationIterationContextValue = getAllocationIterationContextValue(analysisContext)
	var currentHierarchyLevelIterator = getHierarchySubIterationContextValue(analysisContext)
	var _, currentHierarchyLevelIndex = currentHierarchyLevelIterator.Current()

	var potentialDivergence = allocationIterationContextValue.potentialDivergenceMap[hierarchicalId]
	if potentialDivergence == nil {

		var isLowestLevel = currentHierarchyLevelIndex == 0
		potentialDivergence = newPotentialDivergence(hierarchyLevelKey, hierarchicalId, isLowestLevel)

		attachToRootIfTopLevel(
			analysisContext,
			potentialDivergence,
		)

		allocationIterationContextValue.potentialDivergenceMap[hierarchicalId] = potentialDivergence

		return potentialDivergence, true
	}

	return potentialDivergence, false
}

func newPotentialDivergence(
	hierarchyLevelKey string,
	hierarchicalId string,
	isLowestLevel bool,
) *domain.PotentialDivergence {

	var potentialDivergence = &domain.PotentialDivergence{
		HierarchyLevelKey:          hierarchyLevelKey,
		HierarchicalId:             hierarchicalId,
		TotalMarketValue:           0,
		TotalMarketValueDivergence: 0,
		InternalDivergences:        nil,
	}

	if !isLowestLevel {
		potentialDivergence.InternalDivergences = make([]*domain.PotentialDivergence, 0)
	}

	return potentialDivergence
}

func attachToRootIfTopLevel(
	analysisContext context.Context,
	potentialDivergence *domain.PotentialDivergence,
) {

	var currentHierarchyLevelIterator = getHierarchySubIterationContextValue(analysisContext)
	var topAllocationHierarchyLevelIndex = currentHierarchyLevelIterator.Size() - 1
	var _, currentHierarchyLevelIndex = currentHierarchyLevelIterator.Current()
	var divergenceAnalysis = getDivergenceAnalysisContextValue(analysisContext)

	if currentHierarchyLevelIndex == topAllocationHierarchyLevelIndex {
		divergenceAnalysis.AddRootDivergence(potentialDivergence)
	}
}

func calculateCurrentDivergenceValuesFromReferencedPlan(
	potentialDivergences []*domain.PotentialDivergence,
	plannedAllocationMap domain.PlannedAllocationsPerHierarchicalId,
	levelTotalMarketValue int64,
) {
	for _, potentialDivergence := range potentialDivergences {

		var plannedAllocation = plannedAllocationMap.Get(potentialDivergence.HierarchicalId)
		if plannedAllocation != nil {

			calculateDivergenceValue(potentialDivergence, plannedAllocation, levelTotalMarketValue)

			if potentialDivergence.InternalDivergences != nil {
				calculateCurrentDivergenceValuesFromReferencedPlan(
					potentialDivergence.InternalDivergences, plannedAllocationMap, potentialDivergence.TotalMarketValue,
				)
			}

			//To allow for planned side set difference
			plannedAllocationMap.Remove(potentialDivergence.HierarchicalId)
		} else {
			potentialDivergence.TotalMarketValueDivergence = potentialDivergence.TotalMarketValue
		}
	}
}

func calculateDivergenceValue(
	potentialDivergence *domain.PotentialDivergence,
	plannedAllocation *domain.PlannedAllocation,
	levelTotalMarketValue int64,
) {
	var plannedAllocationValue = plannedAllocation.SliceSizePercentage.
		Mul(decimal.NewFromInt(levelTotalMarketValue)).
		Round(0).IntPart()
	potentialDivergence.TotalMarketValueDivergence = potentialDivergence.TotalMarketValue - plannedAllocationValue
}

// TODO clean code
func generatePotentialDivergencesFromAllocationPlanSetDifference(
	portfolio *domain.Portfolio,
	divergenceAnalysis *domain.DivergenceAnalysis,
	potentialDivergenceMap potentialDivercencesPerHierarchicalId,
	plannedAllocationMap domain.PlannedAllocationsPerHierarchicalId,
) {

	var hierarchySize = portfolio.AllocationStructure.Hierarchy.Size()
	var hierarchytopLevelIndex = hierarchySize - 1

	for _, plannedAllocation := range plannedAllocationMap {

		var currentPlannedStructuralId = plannedAllocation.StructuralId

		for i := hierarchytopLevelIndex; i >= 0; i-- {

			var currentLevelHierarchicalId = currentPlannedStructuralId[i:hierarchySize]
			var currentLevelingHierarchicalIdKey = currentLevelHierarchicalId.String()
			var _, currentLevelExists = potentialDivergenceMap[currentLevelingHierarchicalIdKey]
			if !currentLevelExists {

				var isLowestLevel = i == 0
				var isTopLevel = i == hierarchytopLevelIndex

				var potentialDivergence = newPotentialDivergence(
					*currentPlannedStructuralId[i],
					currentLevelingHierarchicalIdKey,
					isLowestLevel,
				)

				var parentTotalMarketValue int64 = 0
				if isTopLevel {
					divergenceAnalysis.AddRootDivergence(potentialDivergence)
					parentTotalMarketValue = divergenceAnalysis.PortfolioTotalMarketValue
				} else {
					var parentLevelHierarchicalId = currentPlannedStructuralId[i+1 : hierarchySize]
					var parentLevelHierarchicalIdKey = parentLevelHierarchicalId.String()
					var parentPotentialDivergence = potentialDivergenceMap[parentLevelHierarchicalIdKey]
					parentPotentialDivergence.AddInternalDivergence(potentialDivergence)
					parentTotalMarketValue = parentPotentialDivergence.TotalMarketValue
				}

				calculateDivergenceValue(potentialDivergence, plannedAllocation, parentTotalMarketValue)
			}
		}
	}
}

func BuildPortfolioAnalysisAppService(
	portfolioDomService *service.PortfolioDomService,
	allocationPlanDomService *service.AllocationPlanDomService,
) *PortfolioAnalysisAppService {
	return &PortfolioAnalysisAppService{
		allocationPlanDomService: allocationPlanDomService,
		portfolioDomService:      portfolioDomService,
	}
}
