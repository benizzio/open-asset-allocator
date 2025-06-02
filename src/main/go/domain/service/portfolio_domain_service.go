package service

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioDomService struct {
	allocationHierarchyFieldExtractorMap map[string]func(*domain.PortfolioAllocation) string
	portfolioRepository                  domain.PortfolioRepository
}

func (service *PortfolioDomService) GetPortfolios() ([]*domain.Portfolio, error) {
	return service.portfolioRepository.GetAllPortfolios()
}

func (service *PortfolioDomService) GetPortfolio(id int) (*domain.Portfolio, error) {
	return service.portfolioRepository.GetPortfolio(id)
}

func (service *PortfolioDomService) GetPortfolioAllocationHistory(id int) ([]*domain.PortfolioAllocation, error) {
	return service.portfolioRepository.GetAllPortfolioAllocations(id, 10)
}

func (service *PortfolioDomService) FindPortfolioAllocations(
	id int,
	timeFrameTag domain.TimeFrameTag,
) ([]*domain.PortfolioAllocation, error) {
	return service.portfolioRepository.FindPortfolioAllocations(id, timeFrameTag)
}

func (service *PortfolioDomService) GetPortfolioSnapshot(
	id int,
	timeFrameTag domain.TimeFrameTag,
) (*domain.Portfolio, []*domain.PortfolioAllocation, error) {

	portfolio, err := service.GetPortfolio(id)
	if err != nil {
		return nil, nil, err
	}

	portfolioAllocations, err := service.FindPortfolioAllocations(id, timeFrameTag)
	if err != nil {
		return nil, nil, err
	}

	return portfolio, portfolioAllocations, nil
}

func (service *PortfolioDomService) GetAllTimeFrameTags(
	portfolioId int,
	timeFrameLimit int,
) ([]domain.TimeFrameTag, error) {
	return service.portfolioRepository.GetAllTimeFrameTags(portfolioId, timeFrameLimit)
}

func (service *PortfolioDomService) GenerateHierarchicalId(
	allocation *domain.PortfolioAllocation,
	hierarchy domain.AllocationHierarchy,
	hierarchyLevelIndex int,
) (string, error) {

	var hierarchicalId string
	var highestHierarchyIndex = len(hierarchy) - 1

	for i := highestHierarchyIndex; i >= hierarchyLevelIndex; i-- {

		idSegment, err := service.GetIdSegment(allocation, &hierarchy[i])
		if err != nil {
			return "", err
		}

		if i <= highestHierarchyIndex-1 {
			idSegment += domain.HierarchicalIdLevelSeparator
		}

		hierarchicalId = idSegment + hierarchicalId
	}

	return hierarchicalId, nil
}

func (service *PortfolioDomService) GetIdSegment(
	allocation *domain.PortfolioAllocation,
	hierarchyLevel *domain.AllocationHierarchyLevel,
) (string, error) {

	var hierarchyLevelKey, err = service.getHierarchyLevelFieldValue(hierarchyLevel, allocation)
	if err != nil {
		return "", err
	}

	return hierarchyLevelKey, nil
}

func (service *PortfolioDomService) getHierarchyLevelFieldValue(
	level *domain.AllocationHierarchyLevel,
	allocation *domain.PortfolioAllocation,
) (string, error) {
	extractorFunction, ok := service.allocationHierarchyFieldExtractorMap[level.Field]
	if !ok {
		return "", infra.BuildAppErrorFormatted("No extractor registered for field: %s", level.Field)
	}
	return extractorFunction(allocation), nil
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
		allocationHierarchyFieldExtractorMap: map[string]func(*domain.PortfolioAllocation) string{
			"assetTicker": func(allocation *domain.PortfolioAllocation) string {
				return allocation.Asset.Ticker
			},
			"class": func(allocation *domain.PortfolioAllocation) string {
				return allocation.Class
			},
		},
		portfolioRepository: portfolioRepository,
	}
}
