package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/langext"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

const (
	availableObservationTimestampsSQL = `
		SELECT DISTINCT paot.id, paot.observation_time_tag AS time_tag, paot.observation_timestamp AS "timestamp"
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
		    coalesce(paot.observation_time_tag, '') AS "observation_timestamp.time_tag",
		    paot.observation_timestamp AS "observation_timestamp.timestamp"
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
	portfolioAllocationClassesSQL = `
		SELECT DISTINCT pa.class 
		FROM portfolio_allocation_fact pa ` + infra.WhereClausePlaceholder + `
		ORDER BY pa.class ASC
	`
	portfolioAllocationsTempTableName   = `portfolio_allocation_fact_merge_temp`
	portfolioAllocationsTempTableDDLSQL = `
		CREATE TEMPORARY TABLE ` + portfolioAllocationsTempTableName + `
		(LIKE portfolio_allocation_fact INCLUDING DEFAULTS) 
		ON COMMIT DROP
	`
	portfolioAllocationsMergeSQL = `
		MERGE INTO portfolio_allocation_fact paf
		USING ` + portfolioAllocationsTempTableName + ` pafmt
		ON 
			paf.asset_id = pafmt.asset_id
			AND paf."class" = pafmt."class"
			AND paf.cash_reserve = pafmt.cash_reserve
			AND paf.portfolio_id = pafmt.portfolio_id
			AND paf.observation_time_id = pafmt.observation_time_id
		WHEN NOT MATCHED BY TARGET THEN
    		INSERT (
				asset_id, 
				"class", 
				cash_reserve, 
				asset_quantity, 
				asset_market_price, 
				total_market_value, 
				portfolio_id, 
				observation_time_id
			)
			VALUES (
				pafmt.asset_id, 
				pafmt."class", 
				pafmt.cash_reserve, 
				pafmt.asset_quantity, 
				pafmt.asset_market_price, 
				pafmt.total_market_value, 
				pafmt.portfolio_id, 
				pafmt.observation_time_id
			)
		WHEN MATCHED AND (
				paf.asset_quantity != pafmt.asset_quantity
				OR paf.asset_market_price != pafmt.asset_market_price
				OR paf.total_market_value != pafmt.total_market_value
			) THEN
				UPDATE SET 
					asset_quantity = pafmt.asset_quantity,
					asset_market_price = pafmt.asset_market_price,
					total_market_value = pafmt.total_market_value
		WHEN NOT MATCHED BY SOURCE AND (
				paf.portfolio_id = $1
				AND paf.observation_time_id = $2
			) THEN
    			DELETE
	`
	observationTimestampInsertSQL = `
		INSERT INTO portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
		VALUES ($1, $2)
		RETURNING id
    `
)

const (
	portfolioIdWhereClause = "AND pa.portfolio_id = {:portfolioId}"
)

const (
	queryAllocationsError           = "Error querying portfolio allocations"
	queryPortfoliosError            = "Error querying portfolios"
	queryPortfolioError             = "Error querying single portfolio"
	queryObservationTimestampsError = "Error querying observation timestamps"
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

// TODO rename to FindPortfolio
func (repository *PortfolioRDBMSRepository) GetPortfolio(id int) (*domain.Portfolio, error) {

	var query = portfolioSQL + `
		WHERE p.id = {:id}
	`

	var result domain.Portfolio
	err := repository.dbAdapter.BuildQuery(query).AddParam("id", id).Build().GetInto(&result)

	return &result, infra.PropagateAsAppErrorWithNewMessage(err, queryPortfolioError, repository)
}

// TODO rename to FindAllPortfolioAllocationsWithinObservationTimestampsLimit
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

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
	}

	langext.UnifyStructPointers(queryResult)
	var result = langext.ToPointerSlice(queryResult)

	return result, nil
}

// TODO rename to FindAvailableObservationTimestamps
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

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, queryObservationTimestampsError, repository)
	}

	var result = langext.ToPointerSlice(queryResult)

	return result, nil
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

func (repository *PortfolioRDBMSRepository) FindAvailablePortfolioAllocationClasses(portfolioId int) ([]string, error) {

	var query = portfolioAllocationClassesSQL

	rows, err := repository.findAvailablePortfolioAllocationClassesRows(portfolioId, query)
	if err != nil {
		return nil, err
	}

	return repository.scanAvailablePortfolioAllocationClassesRows(rows, err)
}

func (repository *PortfolioRDBMSRepository) findAvailablePortfolioAllocationClassesRows(
	portfolioId int,
	query string,
) (*dbx.Rows, error) {

	rows, err := repository.dbAdapter.BuildQuery(query).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", portfolioId).
		Build().GetRows()

	return rows, infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error querying portfolio allocation classes",
		repository,
	)
}

func (repository *PortfolioRDBMSRepository) scanAvailablePortfolioAllocationClassesRows(
	rows *dbx.Rows,
	err error,
) ([]string, error) {

	var queryResult = make([]string, 0)
	for rows.Next() {

		var class string
		err = rows.Scan(&class)
		if err != nil {
			return nil, infra.PropagateAsAppErrorWithNewMessage(
				err,
				"Error scanning portfolio allocation class",
				repository,
			)
		}

		queryResult = append(queryResult, class)
	}
	return queryResult, nil
}

func (repository *PortfolioRDBMSRepository) MergePortfolioAllocationsInTransaction(
	transContext *infra.TransactionalContext,
	portfolioId int,
	observationTimestamp *domain.PortfolioObservationTimestamp,
	allocations []*domain.PortfolioAllocation,
) error {

	if len(allocations) == 0 {
		return nil
	}

	err := repository.insertPortfolioAllocationsInTempTable(transContext, portfolioId, allocations)
	if err != nil {
		return err
	}

	return repository.mergePortfolioAllocations(transContext, portfolioId, observationTimestamp)
}

func (repository *PortfolioRDBMSRepository) insertPortfolioAllocationsInTempTable(
	transContext *infra.TransactionalContext,
	portfolioId int,
	allocations []*domain.PortfolioAllocation,
) error {

	// Create temporary table for merging allocations
	_, err := repository.dbAdapter.ExecuteInTransaction(transContext, portfolioAllocationsTempTableDDLSQL)
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error creating temporary table for portfolio allocation merge",
			repository,
		)
	}

	// Insert values in the temporary table
	columns, insertValues := preparePortfolioAllocationTempInserts(allocations, portfolioId)
	err = repository.dbAdapter.InsertBulkInTransaction(
		transContext,
		portfolioAllocationsTempTableName,
		columns,
		insertValues,
	)
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error merging portfolio allocations",
			repository,
		)
	}

	return nil
}

func preparePortfolioAllocationTempInserts(
	allocations []*domain.PortfolioAllocation,
	id int,
) ([]string, [][]any) {

	var columns = []string{
		"portfolio_id",
		"asset_id",
		"class",
		"cash_reserve",
		"observation_time_id",
		"asset_quantity",
		"asset_market_price",
		"total_market_value",
	}

	var insertValues = make([][]any, len(allocations))
	for i, allocation := range allocations {
		insertValues[i] = []any{
			id,
			allocation.Asset.Id,
			allocation.Class,
			allocation.CashReserve,
			allocation.ObservationTimestamp.Id,
			allocation.AssetQuantity,
			allocation.AssetMarketPrice,
			allocation.TotalMarketValue,
		}
	}

	return columns, insertValues
}

func (repository *PortfolioRDBMSRepository) mergePortfolioAllocations(
	transContext *infra.TransactionalContext,
	portfolioId int,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) error {
	_, err := repository.dbAdapter.ExecuteInTransaction(
		transContext,
		portfolioAllocationsMergeSQL,
		portfolioId,
		observationTimestamp.Id,
	)
	return infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error merging portfolio allocations",
		repository,
	)
}

func (repository *PortfolioRDBMSRepository) InsertObservationTimestampInTransaction(
	transContext *infra.TransactionalContext,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) (*domain.PortfolioObservationTimestamp, error) {

	ids, err := infra.BuildQueryInTransaction[int64](transContext, observationTimestampInsertSQL).
		AddParams(observationTimestamp.TimeTag, observationTimestamp.Timestamp).
		Build().
		Find(infra.ReturningIntIdRowScanner)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting portfolio observation timestamp",
			repository,
		)
	}

	return &domain.PortfolioObservationTimestamp{
		Id:        int(ids[0]),
		TimeTag:   observationTimestamp.TimeTag,
		Timestamp: observationTimestamp.Timestamp,
	}, nil
}

func BuildPortfolioRepository(dbAdapter infra.RepositoryRDBMSAdapter) *PortfolioRDBMSRepository {
	return &PortfolioRDBMSRepository{dbAdapter: dbAdapter}
}
