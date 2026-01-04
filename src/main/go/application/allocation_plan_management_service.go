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

type AllocationPlanManagementAppService struct {
	transactionManager       rdbms.TransactionManager
	allocationPlanDomService *service.AllocationPlanDomService
	assetDomService          *service.AssetDomService
}

func (service *AllocationPlanManagementAppService) PersistAllocationPlan(plan *domain.AllocationPlan) error {

	var err = service.transactionManager.RunInTransaction(
		func(transContext *rdbms.SQLTransactionalContext) error {

			var err = service.persistNewAssets(transContext, plan.Details)
			if err != nil {
				return err
			}

			return service.allocationPlanDomService.PersistAllocationPlanInTransaction(
				transContext,
				plan,
			)
		},
	)

	// if error is DomainValidationError, sent it as is, otherwise propagate
	var validationErr *infra.DomainValidationError
	if errors.As(err, &validationErr) {
		return err
	}

	return infra.PropagateAsAppErrorWithNewMessage(err, "Failed to persist allocation plan", service)
}

func (service *AllocationPlanManagementAppService) persistNewAssets(
	transContext context.Context,
	allocations []*domain.PlannedAllocation,
) error {

	var newAssetsPerTicker = mapNewAssetsPerTickerFromPlannedAllocations(allocations)

	if len(newAssetsPerTicker) == 0 {
		return nil
	}

	persistedAssetsPerTicker, err := service.assetDomService.InsertMappedAssetsInTransaction(
		transContext,
		newAssetsPerTicker,
	)
	if err != nil {
		return err
	}

	replacePersistedAssetsOnPlannedAllocations(allocations, persistedAssetsPerTicker)

	return nil
}

func mapNewAssetsPerTickerFromPlannedAllocations(allocations []*domain.PlannedAllocation) domain.AssetsPerTicker {
	newAssetsPerTicker := make(domain.AssetsPerTicker)
	for _, allocation := range allocations {
		var asset = allocation.Asset
		if asset != nil && langext.IsZeroValue(asset.Id) {
			newAssetsPerTicker[asset.Ticker] = asset
		}
	}
	return newAssetsPerTicker
}

func replacePersistedAssetsOnPlannedAllocations(
	allocations []*domain.PlannedAllocation,
	persistedAssetsPerTicker domain.AssetsPerTicker,
) {
	for _, allocation := range allocations {
		if allocation.Asset != nil && langext.IsZeroValue(allocation.Asset.Id) {
			var persistedAsset = persistedAssetsPerTicker[allocation.Asset.Ticker]
			allocation.Asset = persistedAsset
		}
	}
}

func BuildAllocationPlanManagementAppService(
	transactionManager rdbms.TransactionManager,
	allocationPlanDomService *service.AllocationPlanDomService,
	assetDomService *service.AssetDomService,
) *AllocationPlanManagementAppService {
	return &AllocationPlanManagementAppService{
		transactionManager:       transactionManager,
		allocationPlanDomService: allocationPlanDomService,
		assetDomService:          assetDomService,
	}
}
