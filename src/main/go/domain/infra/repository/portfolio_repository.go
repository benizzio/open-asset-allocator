package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	"time"
)

type portfolioTimeFrame struct {
	TimeFrameTag    domain.TimeFrameTag
	CreateTimestamp time.Time
}

const (
	timeFrameTagsSQL = `
		SELECT DISTINCT ON (time_frame_tag) time_frame_tag, create_timestamp
		FROM portfolio_allocation_fact pa
		` + infra.WhereClausePlaceholder + `
		ORDER BY time_frame_tag DESC, create_timestamp DESC LIMIT {:timeFrameLimit}
	`
	timeFrameTagsComplement = `
		WITH time_frame_tags
			AS (SELECT DISTINCT time_frame_tag, create_timestamp::date FROM portfolio_allocation_fact pa ORDER BY create_timestamp DESC LIMIT {:timeFrameLimit})
	`
	portfolioAllocationsSQL = `
		SELECT pa.*, ass.ticker AS "asset.ticker", coalesce(ass.name, '') AS "asset.name"
		FROM portfolio_allocation_fact pa
		JOIN asset ass ON ass.id = pa.asset_id
		` + infra.WhereClausePlaceholder + `
		ORDER BY pa.time_frame_tag DESC, pa.class ASC, pa.cash_reserve DESC, pa.total_market_value DESC
	`
	portfolioSQL = `
		SELECT p.id, p.name, p.allocation_structure
		FROM portfolio p
	`
)

const (
	portfolioIdWhereClause = "AND pa.portfolio_id = {:portfolioId}"
)

const (
	queryAllocationsError   = "Error querying portfolio allocations"
	queryPortfoliosError    = "Error querying portfolios"
	queryPortfolioError     = "Error querying single portfolio"
	queryTimeFrameTagsError = "Error querying time frame tags"
)

type PortfolioRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
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

func (repository *PortfolioRDBMSRepository) GetPortfolio(id int) (*domain.Portfolio, error) {

	var query = portfolioSQL + `
		WHERE p.id = {:id}
	`

	var result domain.Portfolio
	err := repository.dbAdapter.BuildQuery(query).AddParam("id", id).Build().GetInto(&result)

	return &result, infra.PropagateAsAppErrorWithNewMessage(err, queryPortfolioError, repository)
}

func (repository *PortfolioRDBMSRepository) GetAllPortfolioAllocations(id int, timeFrameLimit int) (
	[]*domain.PortfolioAllocation,
	error,
) {
	var query = timeFrameTagsComplement + portfolioAllocationsSQL

	var queryResult []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(query).
		AddParam("timeFrameLimit", timeFrameLimit).
		AddWhereClause("AND pa.time_frame_tag IN (SELECT time_frame_tag FROM time_frame_tags)").
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		Build().FindInto(&queryResult)

	var result = langext.ToPointerSlice(queryResult)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
}

func (repository *PortfolioRDBMSRepository) FindPortfolioAllocations(id int, timeFrameTag domain.TimeFrameTag) (
	[]*domain.PortfolioAllocation,
	error,
) {
	var queryResult []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(portfolioAllocationsSQL).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		AddWhereClauseAndParam("AND pa.time_frame_tag = {:timeFrameTag}", "timeFrameTag", timeFrameTag).
		Build().FindInto(&queryResult)

	var result = langext.ToPointerSlice(queryResult)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
}

func (repository *PortfolioRDBMSRepository) GetAllTimeFrameTags(
	portfolioId int,
	timeFrameLimit int,
) ([]domain.TimeFrameTag, error) {

	var query = timeFrameTagsSQL

	var queryResult []portfolioTimeFrame
	err := repository.dbAdapter.BuildQuery(query).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", portfolioId).
		AddParam("timeFrameLimit", timeFrameLimit).
		Build().FindInto(&queryResult)

	var result = make([]domain.TimeFrameTag, len(queryResult))
	for i, timeFrame := range queryResult {
		result[i] = timeFrame.TimeFrameTag
	}

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryTimeFrameTagsError, repository)
}

func (repository *PortfolioRDBMSRepository) InsertPortfolio(portfolio *domain.Portfolio) (*domain.Portfolio, error) {

	var insertingCopyPortfolio = *portfolio
	err := repository.dbAdapter.Insert(&insertingCopyPortfolio)
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
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error retrieving updated portfolio", repository)
	}

	return &updatedPortfolio, nil
}

func BuildPortfolioRepository(dbAdapter *infra.RDBMSAdapter) *PortfolioRDBMSRepository {
	return &PortfolioRDBMSRepository{dbAdapter: dbAdapter}
}
