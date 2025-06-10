package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	"time"
)

// Deprecated: Use domain.PortfolioObservationTimestamp
type portfolioTimeFrame struct {
	TimeFrameTag    domain.TimeFrameTag
	CreateTimestamp time.Time
}

const (
	// Deprecated: Use availableObservationTimeTagsSQL
	timeFrameTagsSQL = `
		SELECT DISTINCT ON (time_frame_tag) time_frame_tag, create_timestamp
		FROM portfolio_allocation_fact pa
		` + infra.WhereClausePlaceholder + `
		ORDER BY time_frame_tag DESC, create_timestamp DESC LIMIT {:timeFrameLimit}
	`
	availableObservationTimestampsSQL = `
		SELECT DISTINCT paot.*
		FROM portfolio_allocation_fact pa
		JOIN portfolio_allocation_obs_time paot ON pa.observation_time_id = paot.id
		` + infra.WhereClausePlaceholder + `
		ORDER BY paot.observation_time_tag DESC, paot.observation_timestamp DESC LIMIT {:observationTimestampLimit}
	`
	availableObservationTimestampsComplement = `
		WITH observation_timestamps
			AS (
				SELECT DISTINCT paot.*
				FROM portfolio_allocation_fact pa
				JOIN portfolio_allocation_obs_time paot ON pa.observation_time_id = paot.id
				ORDER BY paot.observation_timestamp DESC 
				LIMIT {:observationTimestampLimit}
			)
	`
	portfolioAllocationsSQL = `
		SELECT 
		    pa.*, 
		    ass.id AS "asset.id", 
		    ass.ticker AS "asset.ticker", 
		    coalesce(ass.name, '') AS "asset.name", 
		    paot.id AS "observation_timestamp.id",
		    coalesce(paot.observation_time_tag, '') AS "observation_timestamp.observation_time_tag",
		    paot.observation_timestamp AS "observation_timestamp.observation_timestamp"
		FROM portfolio_allocation_fact pa
		JOIN asset ass ON ass.id = pa.asset_id
		JOIN portfolio_allocation_obs_time paot ON pa.observation_time_id = paot.id 
		` + infra.WhereClausePlaceholder + `
		ORDER BY paot.observation_timestamp DESC, pa.class ASC, pa.cash_reserve DESC, pa.total_market_value DESC
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
	queryAllocationsError           = "Error querying portfolio allocations"
	queryPortfoliosError            = "Error querying portfolios"
	queryPortfolioError             = "Error querying single portfolio"
	queryTimeFrameTagsError         = "Error querying time frame tags"
	queryObservationTimestampsError = "Error querying observation timestamps"
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

func (repository *PortfolioRDBMSRepository) GetAllPortfolioAllocationsWithinObservationTimestampsLimit(
	id int,
	observationTimestampsLimit int,
) (
	[]*domain.PortfolioAllocation,
	error,
) {
	var query = availableObservationTimestampsComplement + portfolioAllocationsSQL

	var queryResult []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(query).
		AddParam("observationTimestampLimit", observationTimestampsLimit).
		AddWhereClause("AND pa.observation_time_id IN (SELECT id FROM observation_timestamps)").
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		Build().FindInto(&queryResult)

	//TODO proper error handling
	langext.UnifyStructPointers(queryResult)
	var result = langext.ToPointerSlice(queryResult)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
}

// Deprecated: use FindPortfolioAllocationsByObservationTimestamp
func (repository *PortfolioRDBMSRepository) FindPortfolioAllocations(id int, timeFrameTag domain.TimeFrameTag) (
	[]*domain.PortfolioAllocation,
	error,
) {
	var queryResult []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(portfolioAllocationsSQL).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		AddWhereClauseAndParam("AND pa.time_frame_tag = {:timeFrameTag}", "timeFrameTag", timeFrameTag).
		Build().FindInto(&queryResult)

	//TODO proper error handling
	langext.UnifyStructPointers(queryResult)
	var result = langext.ToPointerSlice(queryResult)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
}

func (repository *PortfolioRDBMSRepository) FindPortfolioAllocationsByObservationTimestamp(
	id int,
	observationTimestampId int,
) (
	[]*domain.PortfolioAllocation,
	error,
) {
	var queryResult []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(portfolioAllocationsSQL).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		AddWhereClauseAndParam(
			"AND pa.observation_time_id = {:observationTimestampId}",
			"observationTimestampId",
			observationTimestampId,
		).
		Build().FindInto(&queryResult)

	//TODO proper error handling
	langext.UnifyStructPointers(queryResult)
	var result = langext.ToPointerSlice(queryResult)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
}

// Deprecated: use GetAvailableObservationTimestamps
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

	//TODO proper error handling
	var result = make([]domain.TimeFrameTag, len(queryResult))
	for i, timeFrame := range queryResult {
		result[i] = timeFrame.TimeFrameTag
	}

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryTimeFrameTagsError, repository)
}

func (repository *PortfolioRDBMSRepository) GetAvailableObservationTimestamps(
	portfolioId int,
	observationTimestampsLimit int,
) ([]*domain.PortfolioObservationTimestamp, error) {

	var query = availableObservationTimestampsSQL

	var queryResult []domain.PortfolioObservationTimestamp
	err := repository.dbAdapter.BuildQuery(query).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", portfolioId).
		AddParam("observationTimestampLimit", observationTimestampsLimit).
		Build().FindInto(&queryResult)

	//TODO proper error handling
	var result = langext.ToPointerSlice(queryResult)

	return result, infra.PropagateAsAppErrorWithNewMessage(err, queryObservationTimestampsError, repository)
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
