package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type PortfolioAnalysisService struct {
	portfolioRepository      domain.PortfolioRepository
	allocationPlanRepository domain.AllocationPlanRepository
}

func (service *PortfolioAnalysisService) GeneratePortfolioDivergenceAnalysis(
	id int,
	timeFrameTag domain.TimeFrameTag,
	allocationPlanId int,
) (*domain.DivergenceAnalysis, error) {

	portfolio, err := service.portfolioRepository.GetPortfolio(id)
	if err != nil {
		return nil, err
	}

	portfolioAllocations, err := service.portfolioRepository.FindPortfolioAllocations(id, timeFrameTag)
	if err != nil {
		return nil, err
	}

	allocationPlan, err := service.allocationPlanRepository.GetAllocationPlan(allocationPlanId)
	if err != nil {
		return nil, err
	}

	//TODO continue

	return nil, nil
}
