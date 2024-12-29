package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type PortfolioHistoryService struct {
	portfolioRepository domain.PortfolioRepository
}

func (service *PortfolioHistoryService) GetPortfolios() ([]domain.Portfolio, error) {
	return service.portfolioRepository.GetAllPortfolios()
}

func (service *PortfolioHistoryService) GetPortfolio(id int) (domain.Portfolio, error) {
	return service.portfolioRepository.GetPortfolio(id)
}

func (service *PortfolioHistoryService) GetPortfolioAllocationHistory(id int) ([]domain.PortfolioAllocation, error) {
	return service.portfolioRepository.GetAllPortfolioAllocations(id, 10)
}

func BuildPortfolioHistoryService(portfolioRepository domain.PortfolioRepository) *PortfolioHistoryService {
	return &PortfolioHistoryService{portfolioRepository}
}
