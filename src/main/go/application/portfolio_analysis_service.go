package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/shopspring/decimal"
)

// TODO create utility methods to iterate on context
type PortfolioAllocationIterationMappingContext struct {
	allocation            *domain.PortfolioAllocation
	currentlevelIteration *HierarchyLevelIterationMappingContext
}

type HierarchyLevelIterationMappingContext struct {
	hierarchyLevelIndex int
}

type DivergenceAnalysisMappingContext struct {
	potentialDivergenceMap           map[string]*domain.PotentialDivergence
	allocationHierarchy              domain.AllocationHierarchy
	allocationHierarchySize          int
	topAllocationHierarchyLevelIndex int
	currentAllocationIteration       *PortfolioAllocationIterationMappingContext
}

func (context *DivergenceAnalysisMappingContext) getCurrentAllocation() *domain.PortfolioAllocation {
	if context.currentAllocationIteration != nil {
		return context.currentAllocationIteration.allocation
	}
	return nil
}

func (context *DivergenceAnalysisMappingContext) getCurrentHirearchicalLevel() *domain.AllocationHierarchyLevel {
	if context.currentAllocationIteration != nil && context.currentAllocationIteration.currentlevelIteration != nil {
		return &context.allocationHierarchy[context.currentAllocationIteration.currentlevelIteration.hierarchyLevelIndex]
	}
	return nil
}

func (context *DivergenceAnalysisMappingContext) getCurrentHirearchicalLevelIndex() int {
	if context.currentAllocationIteration != nil && context.currentAllocationIteration.currentlevelIteration != nil {
		return context.currentAllocationIteration.currentlevelIteration.hierarchyLevelIndex
	}
	return -1
}

func buildDivergenceAnalysisMappingContext(
	potentialDivergenceMap map[string]*domain.PotentialDivergence,
	allocationHierarchy domain.AllocationHierarchy,
) *DivergenceAnalysisMappingContext {
	var allocationHierarchySize = len(allocationHierarchy)
	return &DivergenceAnalysisMappingContext{
		potentialDivergenceMap:           potentialDivergenceMap,
		allocationHierarchy:              allocationHierarchy,
		allocationHierarchySize:          allocationHierarchySize,
		topAllocationHierarchyLevelIndex: allocationHierarchySize - 1,
	}
}

type PortfolioAnalysisAppService struct {
	portfolioDomService      *service.PortfolioDomService
	allocationPlanDomService *service.AllocationPlanDomService
}

// TODO clean code
func (service *PortfolioAnalysisAppService) GeneratePortfolioDivergenceAnalysis(
	id int,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) (*domain.DivergenceAnalysis, error) {

	portfolio, portfolioAllocations, err := service.portfolioDomService.GetPortfolioSnapshot(id, timeFrameTag)
	if err != nil {
		return nil, err
	}

	var divergenceAnalysis = buildDivergenceAnalysis(portfolio, timeFrameTag, allocationPlanId)

	err = service.mapPotentialDivergencesFromPortfolioAllocations(divergenceAnalysis, portfolio, portfolioAllocations)
	if err != nil {
		return nil, err
	}

	plannedAllocationMap, err := service.allocationPlanDomService.GetPlannedAllocationsPerHyerarchicalIdMap(allocationPlanId)
	if err != nil {
		return nil, err
	}

	setDivergenceValues(divergenceAnalysis.Root, plannedAllocationMap, divergenceAnalysis.PortfolioTotalMarketValue)

	//TODO calculate planned side set difference (still in plannedAllocationMap, use it to create potential divergences)

	return divergenceAnalysis, nil
}

func (service *PortfolioAnalysisAppService) mapPotentialDivergencesFromPortfolioAllocations(
	divergenceAnalysis *domain.DivergenceAnalysis,
	portfolio *domain.Portfolio,
	portfolioAllocations []*domain.PortfolioAllocation,
) error {

	var mappingContext = buildDivergenceAnalysisMappingContext(
		make(map[string]*domain.PotentialDivergence),
		portfolio.AllocationStructure.Hierarchy,
	)

	for _, allocation := range portfolioAllocations {

		mappingContext.currentAllocationIteration = &PortfolioAllocationIterationMappingContext{
			allocation:            allocation,
			currentlevelIteration: nil,
		}

		err := service.mapPotentialDivergenceFromPortfolioAllocation(divergenceAnalysis, mappingContext)
		if err != nil {
			return err
		}

		divergenceAnalysis.PortfolioTotalMarketValue += allocation.TotalMarketValue

		mappingContext.currentAllocationIteration = nil
	}

	return nil
}

func (service *PortfolioAnalysisAppService) mapPotentialDivergenceFromPortfolioAllocation(
	divergenceAnalysis *domain.DivergenceAnalysis,
	context *DivergenceAnalysisMappingContext,
) error {

	var potentialDivergencesInAllocationHierarchy = make([]*domain.PotentialDivergence, context.allocationHierarchySize)
	var lowerLevelPotentialDivergenceCreated = false

	for hierarchyLevelIndex, _ := range context.allocationHierarchy {

		context.currentAllocationIteration.currentlevelIteration = &HierarchyLevelIterationMappingContext{
			hierarchyLevelIndex: hierarchyLevelIndex,
		}

		hierarchicalId, hierarchyLevelKey, err := service.generatePotentialDivergenceIdentifiers(context)
		if err != nil {
			return err
		}

		var potentialDivergence, potentialDivergenceCreated = buildPotentialDivergenceIfNotExists(
			divergenceAnalysis,
			hierarchicalId,
			hierarchyLevelKey,
			context,
		)

		if lowerLevelPotentialDivergenceCreated {
			connectLowerLevelDivergence(
				hierarchyLevelIndex,
				potentialDivergencesInAllocationHierarchy,
				potentialDivergence,
			)
		}

		potentialDivergencesInAllocationHierarchy[hierarchyLevelIndex] = potentialDivergence

		potentialDivergence.TotalMarketValue += context.getCurrentAllocation().TotalMarketValue

		lowerLevelPotentialDivergenceCreated = potentialDivergenceCreated

		context.currentAllocationIteration.currentlevelIteration = nil
	}

	return nil
}

func (service *PortfolioAnalysisAppService) generatePotentialDivergenceIdentifiers(
	context *DivergenceAnalysisMappingContext,
) (string, string, error) {

	hierarchicalId, err := service.portfolioDomService.GenerateHierarchicalId(
		context.getCurrentAllocation(),
		context.allocationHierarchy,
		context.getCurrentHirearchicalLevelIndex(),
	)
	if err != nil {
		return "", "", err
	}
	//hierarchicalId += "a"

	hierarchyLevelKey, err := service.portfolioDomService.GetIdSegment(
		context.getCurrentAllocation(),
		context.getCurrentHirearchicalLevel(),
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

func buildPotentialDivergenceIfNotExists(
	divergenceAnalysis *domain.DivergenceAnalysis,
	hierarchicalId string,
	hierarchyLevelKey string,
	context *DivergenceAnalysisMappingContext,
) (*domain.PotentialDivergence, bool) {

	var potentialDivergence = context.potentialDivergenceMap[hierarchicalId]
	if potentialDivergence == nil {

		potentialDivergence = &domain.PotentialDivergence{
			HierarchyLevelKey:          hierarchyLevelKey,
			HierarchicalId:             hierarchicalId,
			TotalMarketValue:           0,
			TotalMarketValueDivergence: 0,
			InternalDivergences:        nil,
		}

		hierarchyLevelIndex := context.getCurrentHirearchicalLevelIndex()
		if hierarchyLevelIndex > 0 {
			potentialDivergence.InternalDivergences = make([]*domain.PotentialDivergence, 0)
		}

		if hierarchyLevelIndex == context.topAllocationHierarchyLevelIndex {
			divergenceAnalysis.Root = append(
				divergenceAnalysis.Root,
				potentialDivergence,
			)
		}

		return potentialDivergence, true
	}

	return potentialDivergence, false
}

// TODO use context instead of parameters
func connectLowerLevelDivergence(
	hierarchyLevelIndex int,
	potentialDivergencesInAllocationHierarchy []*domain.PotentialDivergence,
	potentialDivergence *domain.PotentialDivergence,
) {
	previousLevelIndex := hierarchyLevelIndex - 1
	lowerLevelPotentialDivergence := potentialDivergencesInAllocationHierarchy[previousLevelIndex]
	potentialDivergence.InternalDivergences = append(
		potentialDivergence.InternalDivergences,
		lowerLevelPotentialDivergence,
	)
}

// TODO clean code
func setDivergenceValues(
	potentialDivergences []*domain.PotentialDivergence,
	plannedAllocationMap domain.PlannedAllocationsPerHierarchicalId,
	levelTotalMarketValue int64,
) {
	for _, potentialDivergence := range potentialDivergences {

		plannedAllocation := plannedAllocationMap.Get(potentialDivergence.HierarchicalId)
		if plannedAllocation != nil {

			var plannedAllocationValue = plannedAllocation.SliceSizePercentage.
				Mul(decimal.NewFromInt(levelTotalMarketValue)).
				Round(0).IntPart()
			potentialDivergence.TotalMarketValueDivergence = potentialDivergence.TotalMarketValue - plannedAllocationValue

			if potentialDivergence.InternalDivergences != nil {
				setDivergenceValues(
					potentialDivergence.InternalDivergences, plannedAllocationMap, potentialDivergence.TotalMarketValue,
				)
			}

			plannedAllocationMap.Remove(potentialDivergence.HierarchicalId)
		} else {
			potentialDivergence.TotalMarketValueDivergence = potentialDivergence.TotalMarketValue
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
