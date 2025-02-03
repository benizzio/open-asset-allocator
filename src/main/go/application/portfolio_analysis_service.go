package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
)

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

	portfolio, err := service.portfolioDomService.GetPortfolio(id)
	if err != nil {
		return nil, err
	}

	portfolioAllocations, err := service.portfolioDomService.FindPortfolioAllocations(id, timeFrameTag)
	if err != nil {
		return nil, err
	}

	//allocationPlan, err := service.allocationPlanDomService.GetAllocationPlan(allocationPlanId)
	//if err != nil {
	//	return nil, err
	//}

	var divergenceAnalysis = buildDivergenceAnalysis(
		portfolio,
		timeFrameTag,
		allocationPlanId,
	)

	var potentialDivergenceMap = make(map[string]*domain.PotentialDivergence)
	var allocationHierarchy = portfolio.AllocationStructure.Hierarchy

	//TODO calculate divergence according to allocation plan
	for _, allocation := range portfolioAllocations {

		var allocationHierarchySize = len(allocationHierarchy)
		var lastHierarchyLevelIndex = allocationHierarchySize - 1
		var allocationPotentialDivergenceHierarchy = make([]*domain.PotentialDivergence, allocationHierarchySize)
		var lowerLevelPotentialDivergenceCreated = false

		for hierarchyLevelIndex, hierarchyLevel := range allocationHierarchy {

			hierarchicalId, err := service.portfolioDomService.GenerateHierarchicalId(
				allocation,
				allocationHierarchy,
				hierarchyLevelIndex,
			)
			if err != nil {
				return nil, err
			}

			hierarchyLevelKey, err := service.portfolioDomService.GetIdSegment(
				allocation,
				hierarchyLevel,
			)
			if err != nil {
				return nil, err
			}

			var potentialDivergence = potentialDivergenceMap[hierarchicalId]
			var potentialDivergenceCreated = false
			if potentialDivergence == nil {

				potentialDivergence = &domain.PotentialDivergence{
					HierarchyLevelKey:          hierarchyLevelKey,
					HierarchicalId:             hierarchicalId,
					TotalMarketValue:           0,
					TotalMarketValueDivergence: 0,
					InternalDivergences:        nil,
				}

				if hierarchyLevelIndex > 0 {
					potentialDivergence.InternalDivergences = make([]*domain.PotentialDivergence, 0)
				}

				potentialDivergenceMap[hierarchicalId] = potentialDivergence
				if hierarchyLevelIndex == lastHierarchyLevelIndex {
					divergenceAnalysis.Root = append(
						divergenceAnalysis.Root,
						potentialDivergence,
					)
				}
				potentialDivergenceCreated = true
			}

			if lowerLevelPotentialDivergenceCreated {
				previousLevelIndex := hierarchyLevelIndex - 1
				lowerLevelPotentialDivergence := allocationPotentialDivergenceHierarchy[previousLevelIndex]
				potentialDivergence.InternalDivergences = append(
					potentialDivergence.InternalDivergences,
					lowerLevelPotentialDivergence,
				)
			}

			allocationPotentialDivergenceHierarchy[hierarchyLevelIndex] = potentialDivergence

			potentialDivergence.TotalMarketValue += allocation.TotalMarketValue

			lowerLevelPotentialDivergenceCreated = potentialDivergenceCreated
		}

		divergenceAnalysis.PortfolioTotalMarketValue += allocation.TotalMarketValue
	}

	return divergenceAnalysis, nil
}

func buildDivergenceAnalysis(
	portfolio domain.Portfolio,
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

func BuildPortfolioAnalysisAppService(
	portfolioDomService *service.PortfolioDomService,
	allocationPlanDomService *service.AllocationPlanDomService,
) *PortfolioAnalysisAppService {
	return &PortfolioAnalysisAppService{
		allocationPlanDomService: allocationPlanDomService,
		portfolioDomService:      portfolioDomService,
	}
}
