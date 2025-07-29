package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioAllocationManagementAppService struct {
	transactionManager  infra.TransactionManager
	portfolioDomService *service.PortfolioDomService
	assetDomService     *service.AssetDomService
}

// TODO clean
func (service *PortfolioAllocationManagementAppService) MergePortfolioAllocations(
	portfolioId int,
	allocations []*domain.PortfolioAllocation,
) error {

	var err = service.transactionManager.RunInTransaction(
		func(transContext *infra.TransactionalContext) error {
			// TODO:
			// 2. observation timestamp insertions

			err := service.persistNewAssets(transContext, allocations)
			if err != nil {
				return err
			}

			return service.portfolioDomService.MergePortfolioAllocations(
				transContext,
				portfolioId,
				allocations,
			)
		},
	)

	return infra.PropagateAsAppErrorWithNewMessage(err, "Failed to merge portfolio allocations", service)
}

func (service *PortfolioAllocationManagementAppService) persistNewAssets(
	transContext *infra.TransactionalContext,
	allocations []*domain.PortfolioAllocation,
) error {

	assetsToInsertPerTicker := mapNewAssetsPerTicker(allocations)

	if len(assetsToInsertPerTicker) > 0 {

		persistedAssetsPerTicker, err := service.assetDomService.InsertMappedAssetsWithinTransaction(
			transContext,
			assetsToInsertPerTicker,
		)
		if err != nil {
			return err
		}

		replacePersistedAssetsOnAllocations(allocations, persistedAssetsPerTicker)
	}
	return nil
}

func mapNewAssetsPerTicker(allocations []*domain.PortfolioAllocation) domain.AssetsPerTicker {
	var assetsToInsertPerTicker = make(domain.AssetsPerTicker, len(allocations))
	for _, allocation := range allocations {
		var asset = allocation.Asset
		if langext.IsZeroValue(asset.Id) {
			assetsToInsertPerTicker[asset.Ticker] = &asset
		}
	}
	return assetsToInsertPerTicker
}

func replacePersistedAssetsOnAllocations(
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
	transactionManager infra.TransactionManager,
	portfolioDomService *service.PortfolioDomService,
	assetDomService *service.AssetDomService,
) *PortfolioAllocationManagementAppService {
	return &PortfolioAllocationManagementAppService{
		transactionManager,
		portfolioDomService,
		assetDomService,
	}
}
