package application

import (
	"context"
	"errors"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioAllocationManagementAppService struct {
	transactionManager            rdbms.TransactionManager
	portfolioAllocationDomService *service.PortfolioAllocationDomService
	assetDomService               *service.AssetDomService
}

// MergePortfolioAllocations updates observation metadata and merges its asset allocations atomically.
// For example, `service.MergePortfolioAllocations(42, observation, allocations)` writes to portfolio 42.
// Co-authored by: OpenCode and benizzio
func (service *PortfolioAllocationManagementAppService) MergePortfolioAllocations(
	portfolioId int64,
	observationTimestamp *domain.PortfolioObservationTimestamp,
	allocations []*domain.PortfolioAllocation,
) error {

	var err = service.transactionManager.RunInTransaction(
		func(transContext *rdbms.SQLTransactionalContext) error {

			managedObservationTimestamp, err := service.manageObservationTimestamp(
				transContext,
				portfolioId,
				observationTimestamp,
				allocations,
			)
			if err != nil {
				return err
			}

			err = service.persistNewAssets(transContext, allocations)
			if err != nil {
				return err
			}

			return service.portfolioAllocationDomService.MergePortfolioAllocationsInTransaction(
				transContext,
				portfolioId,
				managedObservationTimestamp,
				allocations,
			)
		},
	)

	var validationError *infra.DomainValidationError
	var uniqueConstraintError *infra.UniqueConstraintViolationError
	if errors.As(err, &validationError) || errors.As(err, &uniqueConstraintError) {
		return err
	}
	return infra.PropagateAsAppErrorWithNewMessage(err, "Failed to merge portfolio allocations", service)
}

// manageObservationTimestamp creates or owner-checks an observation, then attaches it to every allocation.
// Co-authored by: GPT-6.1 Sol, OpenCode and benizzio
func (service *PortfolioAllocationManagementAppService) manageObservationTimestamp(
	transContext context.Context,
	portfolioId int64,
	observationTimestamp *domain.PortfolioObservationTimestamp,
	allocations []*domain.PortfolioAllocation,
) (*domain.PortfolioObservationTimestamp, error) {

	var managedObservationTimestamp, err = service.portfolioAllocationDomService.
		PersistObservationTimestampInTransaction(transContext, portfolioId, observationTimestamp)
	if err != nil {
		return nil, err
	}

	for _, allocation := range allocations {
		allocation.ObservationTimestamp = managedObservationTimestamp
	}
	return managedObservationTimestamp, nil
}

func (service *PortfolioAllocationManagementAppService) persistNewAssets(
	transContext context.Context,
	allocations []*domain.PortfolioAllocation,
) error {

	assetsToInsertPerTicker := mapNewAssetsPerTickerFromPortfolioAllocations(allocations)

	if len(assetsToInsertPerTicker) == 0 {
		return nil
	}

	persistedAssetsPerTicker, err := service.assetDomService.InsertMappedAssetsInTransaction(
		transContext,
		assetsToInsertPerTicker,
	)
	if err != nil {
		return err
	}

	replacePersistedAssetsOnPortfolioAllocations(allocations, persistedAssetsPerTicker)

	return nil
}

func mapNewAssetsPerTickerFromPortfolioAllocations(allocations []*domain.PortfolioAllocation) domain.AssetsPerTicker {
	var assetsToInsertPerTicker = make(domain.AssetsPerTicker)
	for _, allocation := range allocations {
		var asset = allocation.Asset
		if langext.IsZeroValue(asset.Id) {
			assetsToInsertPerTicker[asset.Ticker] = &asset
		}
	}
	return assetsToInsertPerTicker
}

func replacePersistedAssetsOnPortfolioAllocations(
	allocations []*domain.PortfolioAllocation,
	persistedAssetsPerTicker domain.AssetsPerTicker,
) {
	for _, allocation := range allocations {
		if langext.IsZeroValue(allocation.Asset.Id) {
			persitedAsset := persistedAssetsPerTicker[allocation.Asset.Ticker]
			allocation.Asset = *persitedAsset
		}
	}
}

func BuildPortfolioAllocationManagementAppService(
	transactionManager rdbms.TransactionManager,
	portfolioAllocationDomService *service.PortfolioAllocationDomService,
	assetDomService *service.AssetDomService,
) *PortfolioAllocationManagementAppService {
	return &PortfolioAllocationManagementAppService{
		transactionManager,
		portfolioAllocationDomService,
		assetDomService,
	}
}
