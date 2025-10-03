package repository

import (
	"context"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

const (
	availableObservationTimestampsSQL = `
		SELECT DISTINCT paot.id, paot.observation_time_tag AS time_tag, paot.observation_timestamp AS "timestamp"
		FROM portfolio_allocation_fact pa
		JOIN portfolio_allocation_obs_time paot ON pa.observation_time_id = paot.id
		` + rdbms.WhereClausePlaceholder + `
		ORDER BY paot.observation_timestamp DESC LIMIT {:observationTimestampLimit}
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
		` + rdbms.WhereClausePlaceholder + `
		ORDER BY paot.observation_timestamp DESC, pa.class ASC, pa.cash_reserve DESC, pa.total_market_value DESC
	`
	portfolioAllocationClassesSQL = `
		SELECT DISTINCT pa.class 
		FROM portfolio_allocation_fact pa ` + rdbms.WhereClausePlaceholder + `
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
		USING ` + portfolioAllocationsTempTableName + ` temp
		ON 
			paf.asset_id = temp.asset_id
			AND paf."class" = temp."class"
			AND paf.cash_reserve = temp.cash_reserve
			AND paf.portfolio_id = temp.portfolio_id
			AND paf.observation_time_id = temp.observation_time_id
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
				temp.asset_id, 
				temp."class", 
				temp.cash_reserve, 
				temp.asset_quantity, 
				temp.asset_market_price, 
				temp.total_market_value, 
				temp.portfolio_id, 
				temp.observation_time_id
			)
		WHEN MATCHED AND (
				paf.asset_quantity != temp.asset_quantity
				OR paf.asset_market_price != temp.asset_market_price
				OR paf.total_market_value != temp.total_market_value
			) THEN
				UPDATE SET 
					asset_quantity = temp.asset_quantity,
					asset_market_price = temp.asset_market_price,
					total_market_value = temp.total_market_value
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
	queryObservationTimestampsError = "Error querying observation timestamps"
)

type PortfolioAllocationRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
}

func (repository *PortfolioAllocationRDBMSRepository) FindAllPortfolioAllocationsWithinObservationTimestampsLimit(
	id int64,
	observationTimestampsLimit int,
) ([]*domain.PortfolioAllocation, error) {

	var query = availableObservationTimestampsComplement + portfolioAllocationsSQL

	var queryResult []domain.PortfolioAllocation
	err := repository.dbAdapter.BuildQuery(query).
		AddParam("observationTimestampLimit", observationTimestampsLimit).
		AddWhereClause("AND pa.observation_time_id IN (SELECT id FROM observation_timestamps)").
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		Build().FindInto(&queryResult)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
	}

	langext.UnifyStructPointers(queryResult)
	var result = langext.ToPointerSlice(queryResult)

	return result, nil
}

func (repository *PortfolioAllocationRDBMSRepository) FindPortfolioAllocationsByObservationTimestamp(
	id int64,
	observationTimestampId int64,
) ([]*domain.PortfolioAllocation, error) {

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

func (repository *PortfolioAllocationRDBMSRepository) FindAvailableObservationTimestamps(
	portfolioId int64,
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

func (repository *PortfolioAllocationRDBMSRepository) FindAvailablePortfolioAllocationClasses(portfolioId int64) (
	[]string,
	error,
) {

	var query = portfolioAllocationClassesSQL

	rows, err := repository.findAvailablePortfolioAllocationClassesRows(portfolioId, query)
	if err != nil {
		return nil, err
	}

	return repository.scanAvailablePortfolioAllocationClassesRows(rows, err)
}

// FindAvailablePortfolioAllocationClassesFromAllSources retrieves unique allocation classes
// from both portfolio_allocation_fact and planned_allocation tables. For planned_allocation,
// it extracts the class value from the hierarchical_id array using the position defined in
// the portfolio's allocation_structure.
//
// Authored by: GitHub Copilot
func (repository *PortfolioAllocationRDBMSRepository) FindAvailablePortfolioAllocationClassesFromAllSources(
	portfolioId int64,
) ([]string, error) {

	var query = `
		WITH class_hierarchy_position AS (
			SELECT
				p.id AS portfolio_id,
				(
					SELECT pos - 1
					FROM jsonb_array_elements(p.allocation_structure->'hierarchy') WITH ORDINALITY AS t(elem, pos)
					WHERE elem->>'field' = 'class'
					LIMIT 1
				) AS class_position
			FROM portfolio p
			WHERE p.id = {:portfolioId}
		)
		SELECT DISTINCT class
		FROM (
			SELECT pa.class
			FROM portfolio_allocation_fact pa
			WHERE pa.portfolio_id = {:portfolioId}
			UNION
			SELECT pa.hierarchical_id[chp.class_position + 1] AS class
			FROM planned_allocation pa
			JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
			CROSS JOIN class_hierarchy_position chp
			WHERE ap.portfolio_id = {:portfolioId}
				AND chp.class_position IS NOT NULL
				AND pa.hierarchical_id[chp.class_position + 1] IS NOT NULL
		) AS combined_classes
		ORDER BY class ASC
	`

	rows, err := repository.dbAdapter.BuildQuery(query).
		AddParam("portfolioId", portfolioId).
		Build().GetRows()

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error querying portfolio allocation classes from all sources",
			repository,
		)
	}

	return repository.scanAvailablePortfolioAllocationClassesRows(rows, err)
}

func (repository *PortfolioAllocationRDBMSRepository) findAvailablePortfolioAllocationClassesRows(
	portfolioId int64,
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

func (repository *PortfolioAllocationRDBMSRepository) scanAvailablePortfolioAllocationClassesRows(
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

func (repository *PortfolioAllocationRDBMSRepository) MergePortfolioAllocationsInTransaction(
	transContext context.Context,
	portfolioId int64,
	observationTimestamp *domain.PortfolioObservationTimestamp,
	allocations []*domain.PortfolioAllocation,
) error {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

	if len(allocations) == 0 {
		return nil
	}

	err := repository.insertPortfolioAllocationsInTempTable(transactionalContext, portfolioId, allocations)
	if err != nil {
		return err
	}

	return repository.mergePortfolioAllocations(transactionalContext, portfolioId, observationTimestamp)
}

func (repository *PortfolioAllocationRDBMSRepository) insertPortfolioAllocationsInTempTable(
	transContext *rdbms.SQLTransactionalContext,
	portfolioId int64,
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
	id int64,
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

func (repository *PortfolioAllocationRDBMSRepository) mergePortfolioAllocations(
	transContext *rdbms.SQLTransactionalContext,
	portfolioId int64,
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

func (repository *PortfolioAllocationRDBMSRepository) InsertObservationTimestampInTransaction(
	transContext context.Context,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) (*domain.PortfolioObservationTimestamp, error) {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return nil, infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

	id, err := rdbms.BuildQueryInTransaction[int64](transactionalContext, observationTimestampInsertSQL).
		AddParams(observationTimestamp.TimeTag, observationTimestamp.Timestamp).
		Build().
		Get(rdbms.ReturningIntIdSingleRowScanner)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting portfolio observation timestamp",
			repository,
		)
	}

	return &domain.PortfolioObservationTimestamp{
		Id:        id,
		TimeTag:   observationTimestamp.TimeTag,
		Timestamp: observationTimestamp.Timestamp,
	}, nil
}

func BuildPortfolioAllocationRepository(dbAdapter rdbms.RepositoryRDBMSAdapter) *PortfolioAllocationRDBMSRepository {
	return &PortfolioAllocationRDBMSRepository{dbAdapter: dbAdapter}
}
