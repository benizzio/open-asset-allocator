package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type PortfolioHistoryService struct {
	portfolioRepository domain.PortfolioRepository
}

func (service *PortfolioHistoryService) GetPortfolioAllocationHistory(id int) ([]domain.PortfolioAllocation, error) {
	return service.portfolioRepository.GetAllPortfolioAllocations(id, 10)
}

func BuildPortfolioHistoryService(portfolioRepository domain.PortfolioRepository) *PortfolioHistoryService {
	return &PortfolioHistoryService{portfolioRepository}
}
