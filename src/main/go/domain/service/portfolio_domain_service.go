package service

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioDomService struct {
	portfolioRepository domain.PortfolioRepository
}

func (service *PortfolioDomService) GetPortfolios() ([]*domain.Portfolio, error) {
	return service.portfolioRepository.GetAllPortfolios()
}

func (service *PortfolioDomService) GetPortfolio(id int) (*domain.Portfolio, error) {
	return service.portfolioRepository.FindPortfolio(id)
}

func (service *PortfolioDomService) PersistPortfolio(portfolio *domain.Portfolio) (*domain.Portfolio, error) {

	var persistedPortfolio *domain.Portfolio
	var err error
	if langext.IsZeroValue(portfolio.Id) {
		persistedPortfolio, err = service.portfolioRepository.InsertPortfolio(portfolio)
	} else {
		persistedPortfolio, err = service.portfolioRepository.UpdatePortfolio(portfolio)
	}

	if err != nil {
		return nil, err
	}

	return persistedPortfolio, nil
}

func BuildPortfolioDomService(portfolioRepository domain.PortfolioRepository) *PortfolioDomService {
	return &PortfolioDomService{
		portfolioRepository,
	}
}
