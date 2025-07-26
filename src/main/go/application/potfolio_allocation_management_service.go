package application

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
)

type PortfolioAllocationManagementService struct {
	portfolioDomService *service.PortfolioDomService
	transactionManager  infra.TransactionManager
}

func (service *PortfolioAllocationManagementService) MergePortfolioAllocations(
	portfolioId int,
	allocations []*domain.PortfolioAllocation,
) error {

	var err = service.transactionManager.RunInTransaction(
		func(transContext *infra.TransactionalContext) error {
			return service.portfolioDomService.MergePortfolioAllocations(
				transContext,
				portfolioId,
				allocations,
			)
		},
	)

	return infra.PropagateAsAppErrorWithNewMessage(err, "Failed to merge portfolio allocations", service)
}

func BuildPortfolioAllocationManagementService(
	portfolioDomService *service.PortfolioDomService,
) *PortfolioAllocationManagementService {
	return &PortfolioAllocationManagementService{
		portfolioDomService: portfolioDomService,
	}
}
