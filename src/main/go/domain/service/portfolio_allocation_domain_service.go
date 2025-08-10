package service

import (
	"time"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioAllocationDomService struct {
	allocationHierarchyFieldExtractorMap map[string]func(*domain.PortfolioAllocation) string
	portfolioAllocationRepository        domain.PortfolioAllocationRepository
}

func (service *PortfolioAllocationDomService) GetPortfolioAllocationHistory(id int) (
	[]*domain.PortfolioAllocation,
	error,
) {
	return service.portfolioAllocationRepository.FindAllPortfolioAllocationsWithinObservationTimestampsLimit(id, 10)
}

func (service *PortfolioAllocationDomService) FindPortfolioAllocationsByObservationTimestamp(
	id int,
	observationTimestampId int,
) ([]*domain.PortfolioAllocation, error) {
	return service.portfolioAllocationRepository.FindPortfolioAllocationsByObservationTimestamp(
		id,
		observationTimestampId,
	)
}

func (service *PortfolioAllocationDomService) GetAvailableObservationTimestamps(
	portfolioId int,
	observationTimestampsLimit int,
) ([]*domain.PortfolioObservationTimestamp, error) {
	return service.portfolioAllocationRepository.FindAvailableObservationTimestamps(
		portfolioId,
		observationTimestampsLimit,
	)
}

func (service *PortfolioAllocationDomService) GenerateHierarchicalId(
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

func (service *PortfolioAllocationDomService) GetIdSegment(
	allocation *domain.PortfolioAllocation,
	hierarchyLevel *domain.AllocationHierarchyLevel,
) (string, error) {

	var hierarchyLevelKey, err = service.getHierarchyLevelFieldValue(hierarchyLevel, allocation)
	if err != nil {
		return "", err
	}

	return hierarchyLevelKey, nil
}

func (service *PortfolioAllocationDomService) getHierarchyLevelFieldValue(
	level *domain.AllocationHierarchyLevel,
	allocation *domain.PortfolioAllocation,
) (string, error) {
	extractorFunction, ok := service.allocationHierarchyFieldExtractorMap[level.Field]
	if !ok {
		return "", infra.BuildAppErrorFormatted("No extractor registered for field: %s", level.Field)
	}
	return extractorFunction(allocation), nil
}

func (service *PortfolioAllocationDomService) FindAvailablePortfolioAllocationClasses(
	portfolioId int,
) ([]string, error) {
	return service.portfolioAllocationRepository.FindAvailablePortfolioAllocationClasses(portfolioId)
}

func (service *PortfolioAllocationDomService) MergePortfolioAllocationsInTransaction(
	transContext *infra.TransactionalContext,
	portfolioId int,
	observationTimestamp *domain.PortfolioObservationTimestamp,
	allocations []*domain.PortfolioAllocation,
) error {
	return service.portfolioAllocationRepository.MergePortfolioAllocationsInTransaction(
		transContext,
		portfolioId,
		observationTimestamp,
		allocations,
	)
}

func (service *PortfolioAllocationDomService) InsertObservationTimestampInTransaction(
	transContext *infra.TransactionalContext,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) (*domain.PortfolioObservationTimestamp, error) {

	if langext.IsZeroValue(observationTimestamp.TimeTag) {
		observationTimestamp.TimeTag = observationTimestamp.Timestamp.Format(time.RFC3339)
	}

	if langext.IsZeroValue(observationTimestamp.Timestamp) {
		observationTimestamp.Timestamp = time.Now()
	}

	return service.portfolioAllocationRepository.InsertObservationTimestampInTransaction(
		transContext,
		observationTimestamp,
	)
}

func BuildPortfolioAllocationDomService(portfolioAllocationRepository domain.PortfolioAllocationRepository) *PortfolioAllocationDomService {
	return &PortfolioAllocationDomService{
		allocationHierarchyFieldExtractorMap: map[string]func(*domain.PortfolioAllocation) string{
			"assetTicker": func(allocation *domain.PortfolioAllocation) string {
				return allocation.Asset.Ticker
			},
			"class": func(allocation *domain.PortfolioAllocation) string {
				return allocation.Class
			},
		},
		portfolioAllocationRepository: portfolioAllocationRepository,
	}
}
