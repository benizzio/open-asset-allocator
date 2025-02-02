package service

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
)

type PortfolioDomService struct {
	allocationHierarchyFieldExtractorMap map[string]func(*domain.PortfolioAllocation) string
	portfolioRepository                  domain.PortfolioRepository
}

func (service *PortfolioDomService) GetPortfolios() ([]domain.Portfolio, error) {
	return service.portfolioRepository.GetAllPortfolios()
}

func (service *PortfolioDomService) GetPortfolio(id int) (domain.Portfolio, error) {
	return service.portfolioRepository.GetPortfolio(id)
}

func (service *PortfolioDomService) GetPortfolioAllocationHistory(id int) ([]domain.PortfolioAllocation, error) {
	return service.portfolioRepository.GetAllPortfolioAllocations(id, 10)
}

func (service *PortfolioDomService) FindPortfolioAllocations(
	id int,
	timeFrameTag domain.TimeFrameTag,
) ([]domain.PortfolioAllocation, error) {
	return service.portfolioRepository.FindPortfolioAllocations(id, timeFrameTag)
}

func (service *PortfolioDomService) GenerateHierarchicalId(
	allocation domain.PortfolioAllocation,
	hierarchy domain.AllocationHierarchy,
	hierarchyLevelIndex int,
) (string, error) {

	var hierarchicalId string
	var highestHierarchyIndex = len(hierarchy) - 1

	for i := highestHierarchyIndex; i >= hierarchyLevelIndex; i-- {

		idSegment, err := service.getIdSegment(allocation, hierarchy, i)
		if err != nil {
			return "", err
		}

		hierarchicalId = idSegment + hierarchicalId
	}

	return hierarchicalId, nil
}

func (service *PortfolioDomService) getIdSegment(
	allocation domain.PortfolioAllocation,
	hierarchy domain.AllocationHierarchy,
	hierarchyIndex int,
) (string, error) {

	var hierarchyLevel = hierarchy[hierarchyIndex]
	var highestHierarchyIndex = len(hierarchy) - 1

	var hierarchyLevelValue, err = service.getHierarchyLevelFieldValue(hierarchyLevel, allocation)
	if err != nil {
		return "", err
	}

	if hierarchyIndex <= highestHierarchyIndex-1 {
		hierarchyLevelValue += domain.HierarchicalIdLevelSeparator
	}

	return hierarchyLevelValue, nil
}

func (service *PortfolioDomService) getHierarchyLevelFieldValue(
	level domain.AllocationHierarchyLevel,
	allocation domain.PortfolioAllocation,
) (string, error) {
	extractorFunction, ok := service.allocationHierarchyFieldExtractorMap[level.Field]
	if !ok {
		return "", infra.BuildAppErrorFormatted("No extractor registered for field: %s", level.Field)
	}
	return extractorFunction(&allocation), nil
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
