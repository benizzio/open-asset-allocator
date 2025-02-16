package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/domain/service"
)

type PortfolioAnalysisConfigurationAppService struct {
	portfolioDomService      *service.PortfolioDomService
	allocationPlanDomService *service.AllocationPlanDomService
}

func (service *PortfolioAnalysisConfigurationAppService) GetDivergenceAnalysisOptions(portfolioId int) (
	*domain.AnalysisOptions,
	error,
) {

	var timeFrameLimit = 10
	var timeFrameTags, err = service.portfolioDomService.GetAllTimeFrameTags(portfolioId, timeFrameLimit)
	if err != nil {
		return nil, err
	}

	var planType = allocation.AssetAllocationPlan
	planIdentifiers, err := service.allocationPlanDomService.GetAllAllocationPlanIdentifiers(portfolioId, &planType)
	if err != nil {
		return nil, err
	}

	var analysisOptions = &domain.AnalysisOptions{
		AvailableHistory: timeFrameTags,
		AvailablePlans:   planIdentifiers,
	}

	return analysisOptions, nil
}

func BuildPortfolioAnalysisConfigurationAppService(
	portfolioDomService *service.PortfolioDomService,
	allocationPlanDomService *service.AllocationPlanDomService,
) *PortfolioAnalysisConfigurationAppService {
	return &PortfolioAnalysisConfigurationAppService{
		portfolioDomService,
		allocationPlanDomService,
	}
}
