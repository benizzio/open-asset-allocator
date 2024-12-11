package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type PortfolioHistoryService struct {
	portfolioRepository domain.PortfolioRepository
}

func (service *PortfolioHistoryService) GetPortfolioHistory() ([]domain.PortfolioSliceAtTime, error) {
	return service.portfolioRepository.GetAllPortfolioSlices(10)
}

func BuildPortfolioHistoryService(portfolioRepository domain.PortfolioRepository) *PortfolioHistoryService {
	return &PortfolioHistoryService{portfolioRepository}
}
