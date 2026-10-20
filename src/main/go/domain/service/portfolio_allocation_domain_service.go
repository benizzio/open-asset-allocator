package service

import (
	"context"
	"errors"
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
// client, for example `InsertObservationTimestampInTransaction(ctx, 42, observation)`. New observations require a tag;
// updates may omit it to preserve the stored tag. Metadata and ownership are validated before persistence.
//
// Co-authored by: GPT-6.1 Sol, OpenCode and benizzio
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
		if err := service.validateObservationOwnership(transContext, portfolioId, observationTimestamp.Id); err != nil {
			return nil, err
		}
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

	persistedObservation, err := service.portfolioAllocationRepository.InsertObservationTimestampInTransaction(
		transContext,
		portfolioId,
		observationTimestamp,
	)
	var foreignKeyError *infra.ForeignKeyConstraintViolationError
	if errors.As(err, &foreignKeyError) &&
		foreignKeyError.ConstraintName == "portfolio_allocation_obs_time_portfolio_fk" {
		return nil, service.observationValidationError("Portfolio does not exist")
	}
	return persistedObservation, err
}

// validateObservationOwnership checks the locked persisted observation before metadata is written.
// Authored by: GPT-6.1 Sol
func (service *PortfolioAllocationDomService) validateObservationOwnership(
	transContext context.Context,
	portfolioId, observationId int64,
) error {
	observation, err := service.portfolioAllocationRepository.FindAndLockObservationTimestampInTransaction(
		transContext, observationId,
	)
	if err != nil {
		return err
	}
	if observation == nil || observation.PortfolioId != portfolioId {
		return service.observationValidationError("Observation does not exist in the requested portfolio")
	}
	return nil
}

// observationValidationError builds a domain validation failure for observation persistence rules.
// Authored by: GPT-6.1 Sol
func (service *PortfolioAllocationDomService) observationValidationError(message string) error {
	return infra.BuildDomainValidationError(
		"Observation validation failed",
		[]*infra.AppError{infra.BuildAppErrorFormattedUnconverted(service, "%s", message)},
	)
}

// validateObservationTimestamp rejects invalid metadata before the repository writes it.
// Co-authored by: GPT-6.1 Sol and OpenCode
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
		if langext.IsZeroValue(observationTimestamp.TimeTag) && langext.IsZeroValue(observationTimestamp.Id) {
			causes = append(causes, infra.BuildAppErrorFormattedUnconverted(
				service,
				"Observation time tag is required for new observations",
			))
		} else if !langext.IsZeroValue(observationTimestamp.TimeTag) &&
			langext.IsEmptyString(observationTimestamp.TimeTag) {
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
