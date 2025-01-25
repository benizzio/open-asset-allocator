package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

type PortfolioService struct {
	portfolioRepository domain.PortfolioRepository
}

func (service *PortfolioService) GetPortfolios() ([]domain.Portfolio, error) {
	return service.portfolioRepository.GetAllPortfolios()
}

func (service *PortfolioService) GetPortfolio(id int) (domain.Portfolio, error) {
	return service.portfolioRepository.GetPortfolio(id)
}

func BuildPortfolioService(portfolioRepository domain.PortfolioRepository) *PortfolioService {
	return &PortfolioService{portfolioRepository}
}
