package application

import (
	"context"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra/util"
	"github.com/shopspring/decimal"
)

type PortfolioDivergenceAnalysisAppService struct {
	portfolioDomService      *service.PortfolioDomService
	allocationPlanDomService *service.AllocationPlanDomService
}

type potentialDivergencesPerHierarchicalId map[string]*domain.PotentialDivergence

// Deprecated: use GeneratePortfolioDivergenceAnalysisNew
func (service *PortfolioDivergenceAnalysisAppService) GeneratePortfolioDivergenceAnalysis(
	portfolioId int,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) (*domain.DivergenceAnalysis, error) {

	var analysisContext, err = service.initializeAnalysisContext(portfolioId, timeFrameTag, allocationPlanId)
	if err != nil {
		return nil, err
	}

	potentialDivergenceMap, err := service.generateDivergenceAnalysisFromPortfolioAllocationSet(analysisContext)
	if err != nil {
		return nil, err
	}

	analysisContext = buildPotentialDivergenceMapContext(analysisContext, potentialDivergenceMap)

	err = service.complementAnalysisWithAllocationPlanSetDifference(analysisContext, allocationPlanId)
	if err != nil {
		return nil, err
	}

	var divergenceAnalysis = getDivergenceAnalysisContextValue(analysisContext).divergenceAnalysis
	return divergenceAnalysis, nil
}

// TODO rename to GeneratePortfolioDivergenceAnalysis after removal of old method
func (service *PortfolioDivergenceAnalysisAppService) GeneratePortfolioDivergenceAnalysisNew(
	portfolioId int,
	observationTimestampId int,
	allocationPlanId int,
) (*domain.DivergenceAnalysis, error) {

	var analysisContext, err = service.initializeAnalysisContextForObservationTimestamp(
		portfolioId,
		observationTimestampId,
		allocationPlanId,
	)
	if err != nil {
		return nil, err
	}

	potentialDivergenceMap, err := service.generateDivergenceAnalysisFromPortfolioAllocationSet(analysisContext)
	if err != nil {
		return nil, err
	}

	analysisContext = buildPotentialDivergenceMapContext(analysisContext, potentialDivergenceMap)

	err = service.complementAnalysisWithAllocationPlanSetDifference(analysisContext, allocationPlanId)
	if err != nil {
		return nil, err
	}

	var divergenceAnalysis = getDivergenceAnalysisContextValue(analysisContext).divergenceAnalysis
	return divergenceAnalysis, nil
}

// Deprecated: use initializeAnalysisContextForObservationTimestamp
func (service *PortfolioDivergenceAnalysisAppService) initializeAnalysisContext(
	portfolioId int,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) (context.Context, error) {

	var portfolio, portfolioAllocations, err = service.portfolioDomService.GetPortfolioSnapshot(
		portfolioId,
		timeFrameTag,
	)
	if err != nil {
		return nil, err
	}

	var divergenceAnalysis = buildDivergenceAnalysis(portfolio, timeFrameTag, nil, allocationPlanId)

	var analysisContextValue = &divergenceAnalysisContextValue{
		portfolio:            portfolio,
		portfolioAllocations: portfolioAllocations,
		divergenceAnalysis:   divergenceAnalysis,
	}

	var analysisContext = buildDivergenceAnalysisContext(context.Background(), analysisContextValue)
	return analysisContext, nil
}

func (service *PortfolioDivergenceAnalysisAppService) initializeAnalysisContextForObservationTimestamp(
	portfolioId int,
	observationTimestampId int,
	allocationPlanId int,
) (context.Context, error) {

	var portfolio, portfolioAllocations, err = service.portfolioDomService.GetPortfolioAtObservationTimestamp(
		portfolioId,
		observationTimestampId,
	)
	if err != nil {
		return nil, err
	}

	// Getting a pointer of PortfolioObservationTimestamp to populate the divergence analysis
	var observationTimestamp *domain.PortfolioObservationTimestamp
	if len(portfolioAllocations) > 0 {
		observationTimestamp = portfolioAllocations[0].ObservationTimestamp
	}

	var divergenceAnalysis = buildDivergenceAnalysis(portfolio, "", observationTimestamp, allocationPlanId)

	var analysisContextValue = &divergenceAnalysisContextValue{
		portfolio:            portfolio,
		portfolioAllocations: portfolioAllocations,
		divergenceAnalysis:   divergenceAnalysis,
	}

	var analysisContext = buildDivergenceAnalysisContext(context.Background(), analysisContextValue)
	return analysisContext, nil
}

func (service *PortfolioDivergenceAnalysisAppService) generateDivergenceAnalysisFromPortfolioAllocationSet(
	analysisContext context.Context,
) (potentialDivergencesPerHierarchicalId, error) {

	potentialDivergenceMap, err := service.mapPotentialDivergencesFromPortfolioAllocations(analysisContext)
	if err != nil {
		return nil, err
	}

	return potentialDivergenceMap, nil
}

func (service *PortfolioDivergenceAnalysisAppService) mapPotentialDivergencesFromPortfolioAllocations(
	analysisContext context.Context,
) (potentialDivergencesPerHierarchicalId, error) {

	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)
	var divergenceAnalysis = analysisContextValue.divergenceAnalysis
	var potentialDivergenceMap = make(potentialDivergencesPerHierarchicalId)
	var portfolioAllocations = analysisContextValue.portfolioAllocations

	var iterationContextValue = &allocationIterationMappingContextValue{
		potentialDivergenceMap:       potentialDivergenceMap,
		portfolioAllocationsIterator: util.NewIterator(portfolioAllocations),
	}

	var allocationIterationMappingContext = buildAllocationIterationContext(analysisContext, iterationContextValue)

	for iterationContextValue.portfolioAllocationsIterator.HasNext() {

		var allocation, _ = iterationContextValue.portfolioAllocationsIterator.Next()

		err := service.mapPotentialDivergenceFromPortfolioAllocation(allocationIterationMappingContext)
		if err != nil {
			return nil, err
		}

		divergenceAnalysis.PortfolioTotalMarketValue += allocation.TotalMarketValue
	}

	return potentialDivergenceMap, nil
}

func (service *PortfolioDivergenceAnalysisAppService) complementAnalysisWithAllocationPlanSetDifference(
	analysisContext context.Context,
	allocationPlanId int,
) error {

	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)
	var divergenceAnalysis = analysisContextValue.divergenceAnalysis

	plannedAllocationMap, err := service.allocationPlanDomService.GetPlannedAllocationsPerHyerarchicalIdMap(allocationPlanId)
	if err != nil {
		return err
	}

	calculateCurrentDivergenceValuesFromReferencedPlan(
		divergenceAnalysis.Root,
		plannedAllocationMap,
		divergenceAnalysis.PortfolioTotalMarketValue,
	)

	generatePotentialDivergencesFromAllocationPlanSetDifference(
		analysisContext,
		plannedAllocationMap,
	)

	return nil
}

func (service *PortfolioDivergenceAnalysisAppService) mapPotentialDivergenceFromPortfolioAllocation(
	analysisContext context.Context,
) error {

	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)
	var allocationIterationContextValue = getAllocationIterationContextValue(analysisContext)
	var allocationHierarchy = analysisContextValue.portfolio.AllocationStructure.Hierarchy

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

func (service *PortfolioDivergenceAnalysisAppService) buildAndConnectPotentialDivergenceIfNotExists(
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

func (service *PortfolioDivergenceAnalysisAppService) generatePotentialDivergenceIdentifiers(
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

	hierarchyLevelKey, err := service.portfolioDomService.GetIdSegment(currentAllocation, currentHierarchyLevel)
	if err != nil {
		return "", "", err
	}
	return hierarchicalId, hierarchyLevelKey, nil
}

func buildDivergenceAnalysis(
	portfolio *domain.Portfolio,
	timeFrameTag domain.TimeFrameTag,
	observationTimestamp *domain.PortfolioObservationTimestamp,
	allocationPlanId int,
) *domain.DivergenceAnalysis {
	return &domain.DivergenceAnalysis{
		PortfolioId:               portfolio.Id,
		TimeFrameTag:              timeFrameTag,
		ObservationTimestamp:      observationTimestamp,
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

		attachToRootIfTopLevel(analysisContext, potentialDivergence)

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
	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)

	if currentHierarchyLevelIndex == topAllocationHierarchyLevelIndex {
		analysisContextValue.divergenceAnalysis.AddRootDivergence(potentialDivergence)
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
					potentialDivergence.InternalDivergences,
					plannedAllocationMap,
					potentialDivergence.TotalMarketValue,
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

func generatePotentialDivergencesFromAllocationPlanSetDifference(
	analysisContext context.Context,
	plannedAllocationMap domain.PlannedAllocationsPerHierarchicalId,
) {

	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)
	var hierarchySize = analysisContextValue.portfolio.AllocationStructure.Hierarchy.Size()

	var hierarchytopLevelIndex = hierarchySize - 1

	for _, plannedAllocation := range plannedAllocationMap {

		var currentPlannedStructuralId = plannedAllocation.StructuralId

		checkAndGeneratePotentialDivergencesOnHierarchy(
			analysisContext,
			plannedAllocation,
			hierarchytopLevelIndex,
			currentPlannedStructuralId,
		)
	}
}

func checkAndGeneratePotentialDivergencesOnHierarchy(
	analysisContext context.Context,
	plannedAllocation *domain.PlannedAllocation,
	hierarchytopLevelIndex int,
	currentPlannedStructuralId domain.HierarchicalId,
) {

	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)
	var potentialDivergenceMap = getPotentialDivergenceMapContextValue(analysisContext)
	var hierarchySize = analysisContextValue.portfolio.AllocationStructure.Hierarchy.Size()

	for i := hierarchytopLevelIndex; i >= 0; i-- {

		var currentLevelHierarchicalId = currentPlannedStructuralId[i:hierarchySize]
		var currentLevelHierarchicalIdString = currentLevelHierarchicalId.String()
		var _, currentLevelExists = potentialDivergenceMap[currentLevelHierarchicalIdString]
		if !currentLevelExists {

			var parentLevelHierarchicalId = currentPlannedStructuralId[i+1 : hierarchySize]
			var isLowestHierarchyLevel = i == 0
			var isTopHierarchyLevel = i == hierarchytopLevelIndex

			generateAndAttachPotentialDivergenceForPlannedAllocation(
				analysisContext,
				plannedAllocation,
				*currentPlannedStructuralId[i],
				parentLevelHierarchicalId.String(),
				currentLevelHierarchicalIdString,
				isTopHierarchyLevel,
				isLowestHierarchyLevel,
			)
		}
	}
}

func generateAndAttachPotentialDivergenceForPlannedAllocation(
	analysisContext context.Context,
	plannedAllocation *domain.PlannedAllocation,
	hierarchyLevelkey string,
	parentLevelHierarchicalId string,
	currentLevelHierarchicalId string,
	isTopHierarchyLevel bool,
	isLowestHierarchyLevel bool,
) {

	var potentialDivergenceMap = getPotentialDivergenceMapContextValue(analysisContext)
	var analysisContextValue = getDivergenceAnalysisContextValue(analysisContext)
	var divergenceAnalysis = analysisContextValue.divergenceAnalysis

	var potentialDivergence = newPotentialDivergence(
		hierarchyLevelkey,
		currentLevelHierarchicalId,
		isLowestHierarchyLevel,
	)

	var parentTotalMarketValue int64 = 0
	if isTopHierarchyLevel {
		divergenceAnalysis.AddRootDivergence(potentialDivergence)
		parentTotalMarketValue = divergenceAnalysis.PortfolioTotalMarketValue
	} else {
		var parentPotentialDivergence = potentialDivergenceMap[parentLevelHierarchicalId]
		parentPotentialDivergence.AddInternalDivergence(potentialDivergence)
		parentTotalMarketValue = parentPotentialDivergence.TotalMarketValue
	}

	calculateDivergenceValue(potentialDivergence, plannedAllocation, parentTotalMarketValue)
}

func BuildPortfolioDivergenceAnalysisAppService(
	portfolioDomService *service.PortfolioDomService,
	allocationPlanDomService *service.AllocationPlanDomService,
) *PortfolioDivergenceAnalysisAppService {
	return &PortfolioDivergenceAnalysisAppService{
		allocationPlanDomService: allocationPlanDomService,
		portfolioDomService:      portfolioDomService,
	}
}
