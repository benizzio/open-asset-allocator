package application

import (
	"fmt"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
)

type PortfolioAnalysisAppService struct {
	portfolioDomService      *service.PortfolioDomService
	allocationPlanRepository domain.AllocationPlanRepository
}

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

	//allocationPlan, err := service.allocationPlanRepository.GetAllocationPlan(allocationPlanId)
	//if err != nil {
	//	return nil, err
	//}

	//var potentialDivergenceMap = make(map[string]domain.PotentialDivergence)
	var allocationHierarchy = portfolio.AllocationStructure.Hierarchy

	for _, allocation := range portfolioAllocations {
		//_ is level
		for levelIndex, _ := range allocationHierarchy {

			hierarchicalId, err := service.portfolioDomService.GenerateHierarchicalId(
				allocation,
				allocationHierarchy,
				levelIndex,
			)
			if err != nil {
				return nil, err
			}

			//TODO remove and continue
			fmt.Println("============>" + hierarchicalId)
		}
	}

	var divergenceAnalysis = buildDivergenceAnalysis(
		portfolio,
		timeFrameTag,
		allocationPlanId,
	)

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
		Root:                      make([]domain.PotentialDivergence, 0),
	}
}

func BuildPortfolioAnalysisAppService(
	portfolioDomService *service.PortfolioDomService,
	allocationPlanRepository domain.AllocationPlanRepository,
) *PortfolioAnalysisAppService {
	return &PortfolioAnalysisAppService{
		allocationPlanRepository: allocationPlanRepository,
		portfolioDomService:      portfolioDomService,
	}
}
