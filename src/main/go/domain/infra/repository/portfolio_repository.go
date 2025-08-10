package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
)

const (
	portfolioSQL = `
		SELECT p.id, p.name, p.allocation_structure
		FROM portfolio p
	`
)

const (
	queryPortfoliosError = "Error querying portfolios"
	queryPortfolioError  = "Error querying single portfolio"
)

type PortfolioRDBMSRepository struct {
	dbAdapter infra.RepositoryRDBMSAdapter
}

func (repository *PortfolioRDBMSRepository) GetAllPortfolios() ([]*domain.Portfolio, error) {

	var result []domain.Portfolio
	err := repository.dbAdapter.BuildQuery(portfolioSQL).Build().FindInto(&result)

	return langext.ToPointerSlice(result), infra.PropagateAsAppErrorWithNewMessage(
		err,
		queryPortfoliosError,
		repository,
	)
}

func (repository *PortfolioRDBMSRepository) FindPortfolio(id int) (*domain.Portfolio, error) {

	var query = portfolioSQL + `
		WHERE p.id = {:id}
	`

	var result domain.Portfolio
	err := repository.dbAdapter.BuildQuery(query).AddParam("id", id).Build().GetInto(&result)

	return &result, infra.PropagateAsAppErrorWithNewMessage(err, queryPortfolioError, repository)
}

func (repository *PortfolioRDBMSRepository) InsertPortfolio(portfolio *domain.Portfolio) (*domain.Portfolio, error) {

	var insertingCopyPortfolio = *portfolio
	err := repository.dbAdapter.Insert(&insertingCopyPortfolio)
	//TODO simplify error handling
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error inserting portfolio", repository)
	}

	return &insertingCopyPortfolio, nil
}

func (repository *PortfolioRDBMSRepository) UpdatePortfolio(portfolio *domain.Portfolio) (*domain.Portfolio, error) {

	var updatingCopyPortfolio = *portfolio
	err := repository.dbAdapter.UpdateListedFields(&updatingCopyPortfolio, "Name")
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error updating portfolio", repository)
	}

	var updatedPortfolio domain.Portfolio
	err = repository.dbAdapter.Read(&updatedPortfolio, portfolio.Id)
	//TODO simplify error handling
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error retrieving updated portfolio", repository)
	}

	return &updatedPortfolio, nil
}

func BuildPortfolioRepository(dbAdapter infra.RepositoryRDBMSAdapter) *PortfolioRDBMSRepository {
	return &PortfolioRDBMSRepository{dbAdapter: dbAdapter}
}
