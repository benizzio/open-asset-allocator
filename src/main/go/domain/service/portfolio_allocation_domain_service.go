package service

import (
	"context"
	"strings"
	"unicode/utf8"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioAllocationDomService struct {
	allocationHierarchyFieldExtractorMap map[string]func(*domain.PortfolioAllocation) string
	portfolioAllocationRepository        domain.PortfolioAllocationRepository
}

func (service *PortfolioAllocationDomService) GetPortfolioAllocationHistory(portfolioId int64) (
	[]*domain.PortfolioAllocation,
	error,
) {
	return service.portfolioAllocationRepository.
		FindAllPortfolioAllocationsWithinObservationTimestampsLimit(portfolioId, 10)
}

func (service *PortfolioAllocationDomService) FindPortfolioAllocationsByObservationTimestamp(
	id int64,
	observationTimestampId int64,
) ([]*domain.PortfolioAllocation, error) {
	return service.portfolioAllocationRepository.FindPortfolioAllocationsByObservationTimestamp(
		id,
		observationTimestampId,
	)
}

func (service *PortfolioAllocationDomService) GetAvailableObservationTimestamps(
	portfolioId int64,
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
		return "", infra.BuildAppErrorFormatted(service, "No extractor registered for field: %s", level.Field)
	}
	return extractorFunction(allocation), nil
}

func (service *PortfolioAllocationDomService) MergePortfolioAllocationsInTransaction(
	transContext context.Context,
	portfolioId int64,
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

// InsertObservationTimestampInTransaction persists a new observation for the supplied portfolio, or updates an
// existing observation only when it belongs to that portfolio. Callers should pass the observation selected by the
// client, for example `InsertObservationTimestampInTransaction(ctx, 42, observation)`; its required tag and instant
// are validated before persistence.
//
// Co-authored by: OpenCode and benizzio
func (service *PortfolioAllocationDomService) InsertObservationTimestampInTransaction(
	transContext context.Context,
	portfolioId int64,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) (*domain.PortfolioObservationTimestamp, error) {
	validationError := service.validateObservationTimestamp(portfolioId, observationTimestamp)
	if validationError != nil {
		return nil, validationError
	}

	observationTimestamp.PortfolioId = portfolioId
	if !langext.IsZeroValue(observationTimestamp.Id) {
		err := service.portfolioAllocationRepository.UpdateObservationTimestampInTransaction(
			transContext,
			portfolioId,
			observationTimestamp,
		)
		if err != nil {
			return nil, err
		}
		return observationTimestamp, nil
	}

	return service.portfolioAllocationRepository.InsertObservationTimestampInTransaction(
		transContext,
		portfolioId,
		observationTimestamp,
	)
}

// validateObservationTimestamp rejects invalid metadata before the repository writes it.
// Authored by: OpenCode
func (service *PortfolioAllocationDomService) validateObservationTimestamp(
	portfolioId int64,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) error {
	var causes = make([]*infra.AppError, 0)
	if observationTimestamp == nil {
		causes = append(causes, infra.BuildAppErrorFormattedUnconverted(service, "Observation timestamp is required"))
	} else {
		if portfolioId <= 0 {
			causes = append(causes, infra.BuildAppErrorFormattedUnconverted(service, "Portfolio ID must be positive"))
		}
		if observationTimestamp.TimeTag == "" && observationTimestamp.Id == 0 {
			causes = append(causes, infra.BuildAppErrorFormattedUnconverted(
				service,
				"Observation time tag is required for new observations",
			))
		} else if observationTimestamp.TimeTag != "" && strings.TrimSpace(observationTimestamp.TimeTag) == "" {
			causes = append(causes, infra.BuildAppErrorFormattedUnconverted(
				service,
				"Observation time tag must not contain only whitespace",
			))
		} else if utf8.RuneCountInString(observationTimestamp.TimeTag) > 100 {
			causes = append(causes, infra.BuildAppErrorFormattedUnconverted(
				service,
				"Observation time tag must not exceed 100 characters",
			))
		}
		if langext.IsZeroValue(observationTimestamp.Timestamp) {
			causes = append(causes, infra.BuildAppErrorFormattedUnconverted(
				service,
				"Observation timestamp is required and must not be zero",
			))
		}
	}
	if len(causes) == 0 {
		return nil
	}
	return infra.BuildDomainValidationError("Observation timestamp validation failed", causes)
}

func BuildPortfolioAllocationDomService(
	portfolioAllocationRepository domain.PortfolioAllocationRepository,
) *PortfolioAllocationDomService {
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
