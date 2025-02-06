package application

import (
	"context"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra/util"
	"github.com/shopspring/decimal"
)

const (
	allocationIterationMappingContextKey   = "allocationIterationMappingContext"
	hierarchySubIterationMappingContextKey = "hierarchySubIterationMappingContext"
)

type allocationIterationMappingContextValue struct {
	potentialDivergenceMap       map[string]*domain.PotentialDivergence
	portfolioAllocationsIterator *util.Iterator[*domain.PortfolioAllocation]
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

	var analysisContext = context.Background()

	portfolio, portfolioAllocations, err := service.portfolioDomService.GetPortfolioSnapshot(id, timeFrameTag)
	if err != nil {
		return nil, err
	}

	var divergenceAnalysis = buildDivergenceAnalysis(portfolio, timeFrameTag, allocationPlanId)

	err = service.mapPotentialDivergencesFromPortfolioAllocations(
		analysisContext,
		divergenceAnalysis,
		portfolio,
		portfolioAllocations,
	)
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
	analysisContext context.Context,
	divergenceAnalysis *domain.DivergenceAnalysis,
	portfolio *domain.Portfolio,
	portfolioAllocations []*domain.PortfolioAllocation,
) error {

	var allocationHierarchy = portfolio.AllocationStructure.Hierarchy

	var iterationContextValue = &allocationIterationMappingContextValue{
		potentialDivergenceMap:       make(map[string]*domain.PotentialDivergence),
		portfolioAllocationsIterator: util.NewIterator(portfolioAllocations),
	}

	var allocationIterationMappingContext = context.WithValue(
		analysisContext,
		allocationIterationMappingContextKey,
		iterationContextValue,
	)

	for iterationContextValue.portfolioAllocationsIterator.HasNext() {

		var allocation, _ = iterationContextValue.portfolioAllocationsIterator.Next()

		err := service.mapPotentialDivergenceFromPortfolioAllocation(
			allocationIterationMappingContext,
			divergenceAnalysis,
			allocationHierarchy,
		)
		if err != nil {
			return err
		}

		divergenceAnalysis.PortfolioTotalMarketValue += allocation.TotalMarketValue
	}

	return nil
}

func (service *PortfolioAnalysisAppService) mapPotentialDivergenceFromPortfolioAllocation(
	analysisContext context.Context,
	divergenceAnalysis *domain.DivergenceAnalysis,
	allocationHierarchy domain.AllocationHierarchy,
) error {

	//TODO make a context type with methods do do this
	var allocationIterationContextValue = analysisContext.
		Value(allocationIterationMappingContextKey).(*allocationIterationMappingContextValue)

	var potentialDivergencesInAllocationHierarchy = make([]*domain.PotentialDivergence, allocationHierarchy.Size())
	var lowerLevelPotentialDivergenceCreated = false

	var allocationHierarchyIterator = util.NewIterator(allocationHierarchy)
	//TODO make a context type with methods do do this
	var hierarchySubIterationMappingContext = context.WithValue(
		analysisContext,
		hierarchySubIterationMappingContextKey,
		allocationHierarchyIterator,
	)

	for allocationHierarchyIterator.HasNext() {

		var _, allocationHierarchyLevelIndex = allocationHierarchyIterator.Next()

		hierarchicalId, hierarchyLevelKey, err := service.generatePotentialDivergenceIdentifiers(
			hierarchySubIterationMappingContext,
			allocationHierarchy,
		)
		if err != nil {
			return err
		}

		var potentialDivergence, potentialDivergenceCreated = buildPotentialDivergenceIfNotExists(
			hierarchySubIterationMappingContext,
			divergenceAnalysis,
			hierarchicalId,
			hierarchyLevelKey,
		)

		if lowerLevelPotentialDivergenceCreated {
			connectLowerLevelDivergence(
				potentialDivergence,
				potentialDivergencesInAllocationHierarchy[allocationHierarchyLevelIndex-1],
			)
		}

		potentialDivergencesInAllocationHierarchy[allocationHierarchyLevelIndex] = potentialDivergence

		var currentAllocation, _ = allocationIterationContextValue.portfolioAllocationsIterator.Current()
		potentialDivergence.TotalMarketValue += currentAllocation.TotalMarketValue

		lowerLevelPotentialDivergenceCreated = potentialDivergenceCreated

	}

	return nil
}

func (service *PortfolioAnalysisAppService) generatePotentialDivergenceIdentifiers(
	analysisContext context.Context,
	allocationHierarchy domain.AllocationHierarchy,
) (string, string, error) {

	//TODO make a context type with methods do do this
	var allocationIterationContextValue = analysisContext.
		Value(allocationIterationMappingContextKey).(*allocationIterationMappingContextValue)

	var currentAllocation, _ = allocationIterationContextValue.portfolioAllocationsIterator.Current()

	//TODO make a context type with methods do do this
	var currentHierarchyLevelIterator = analysisContext.
		Value(hierarchySubIterationMappingContextKey).(*util.Iterator[domain.AllocationHierarchyLevel])
	var currentHierarchyLevel, currentHierarchyLevelIndex = currentHierarchyLevelIterator.CurrentPointer()

	hierarchicalId, err := service.portfolioDomService.GenerateHierarchicalId(
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

// TODO clean code
func buildPotentialDivergenceIfNotExists(
	analysisContext context.Context,
	divergenceAnalysis *domain.DivergenceAnalysis,
	hierarchicalId string,
	hierarchyLevelKey string,
) (*domain.PotentialDivergence, bool) {

	//TODO make a context type with methods do do this
	var allocationIterationContextValue = analysisContext.
		Value(allocationIterationMappingContextKey).(*allocationIterationMappingContextValue)

	//TODO make a context type with methods do do this
	var currentHierarchyLevelIterator = analysisContext.
		Value(hierarchySubIterationMappingContextKey).(*util.Iterator[domain.AllocationHierarchyLevel])
	var _, currentHierarchyLevelIndex = currentHierarchyLevelIterator.Current()

	var potentialDivergence = allocationIterationContextValue.potentialDivergenceMap[hierarchicalId]
	if potentialDivergence == nil {

		potentialDivergence = &domain.PotentialDivergence{
			HierarchyLevelKey:          hierarchyLevelKey,
			HierarchicalId:             hierarchicalId,
			TotalMarketValue:           0,
			TotalMarketValueDivergence: 0,
			InternalDivergences:        nil,
		}

		if currentHierarchyLevelIndex > 0 {
			potentialDivergence.InternalDivergences = make([]*domain.PotentialDivergence, 0)
		}

		var topAllocationHierarchyLevelIndex = currentHierarchyLevelIterator.Size() - 1
		if currentHierarchyLevelIndex == topAllocationHierarchyLevelIndex {
			divergenceAnalysis.Root = append(
				divergenceAnalysis.Root,
				potentialDivergence,
			)
		}

		allocationIterationContextValue.potentialDivergenceMap[hierarchicalId] = potentialDivergence

		return potentialDivergence, true
	}

	return potentialDivergence, false
}

// TODO remove this for an add metod in slice
func connectLowerLevelDivergence(
	potentialDivergence *domain.PotentialDivergence,
	lowerLevelDivergence *domain.PotentialDivergence,
) {
	potentialDivergence.InternalDivergences = append(
		potentialDivergence.InternalDivergences,
		lowerLevelDivergence,
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
