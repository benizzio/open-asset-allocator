package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/domain/service"
)

type PortfolioAnalysisConfigurationAppService struct {
	portfolioAllocationDomService *service.PortfolioAllocationDomService
	allocationPlanDomService      *service.AllocationPlanDomService
}

func (service *PortfolioAnalysisConfigurationAppService) GetDivergenceAnalysisOptions(portfolioId int64) (
	*domain.AnalysisOptions,
	error,
) {

	var observationTimestampsLimit = 10

	availableTimestamps, err := service.portfolioAllocationDomService.GetAvailableObservationTimestamps(
		portfolioId,
		observationTimestampsLimit,
	)
	if err != nil {
		return nil, err
	}

	var planType = allocation.AssetAllocationPlan
	planIdentifiers, err := service.allocationPlanDomService.GetAllAllocationPlanIdentifiers(portfolioId, &planType)
	if err != nil {
		return nil, err
	}

	var analysisOptions = &domain.AnalysisOptions{
		AvailableObservationTimestamps: availableTimestamps,
		AvailablePlans:                 planIdentifiers,
	}

	return analysisOptions, nil
}

func BuildPortfolioAnalysisConfigurationAppService(
	portfolioAllocationDomService *service.PortfolioAllocationDomService,
	allocationPlanDomService *service.AllocationPlanDomService,
) *PortfolioAnalysisConfigurationAppService {
	return &PortfolioAnalysisConfigurationAppService{
		portfolioAllocationDomService,
		allocationPlanDomService,
	}
}
