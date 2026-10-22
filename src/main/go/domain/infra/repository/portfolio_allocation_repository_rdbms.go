package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
)

// Co-authored by: GPT-6 Luna and benizzio
const (
	availableObservationTimestampsSQL = `
		SELECT DISTINCT paot.id, paot.portfolio_id, paot.observation_time_tag AS time_tag,
			paot.observation_timestamp AS "timestamp"
		FROM portfolio_allocation_fact pa
		JOIN portfolio_allocation_obs_time paot
		  ON pa.portfolio_id = paot.portfolio_id AND pa.observation_time_id = paot.id
		` + rdbms.WhereClausePlaceholder + `
		ORDER BY paot.observation_timestamp DESC, paot.id DESC LIMIT {:observationTimestampLimit}
	`
	availableObservationTimestampsComplement = `
		WITH observation_timestamps
			AS (
				SELECT DISTINCT paot.*
				FROM portfolio_allocation_fact pa
				JOIN portfolio_allocation_obs_time paot
				  ON pa.portfolio_id = paot.portfolio_id AND pa.observation_time_id = paot.id
				WHERE pa.portfolio_id = {:portfolioId}
				ORDER BY paot.observation_timestamp DESC, paot.id DESC
				LIMIT {:observationTimestampLimit}
			)
	`
	portfolioAllocationsSQL = `
		SELECT
		    pa.*,
		    ass.id AS "asset.id",
		    ass.ticker AS "asset.ticker",
		    coalesce(ass.name, '') AS "asset.name",
		    coalesce(ass.external_data #>> '{data,0,source}', '') AS selected_external_asset_source,
		    coalesce(ass.external_data #>> '{data,0,ticker}', '') AS selected_external_asset_ticker,
		    coalesce(ass.external_data #>> '{data,0,exchangeId}', '') AS selected_external_asset_exchange_id,
		    paot.id AS "observation_timestamp.id",
		    paot.portfolio_id AS "observation_timestamp.portfolio_id",
		    coalesce(paot.observation_time_tag, '') AS "observation_timestamp.time_tag",
		    paot.observation_timestamp AS "observation_timestamp.timestamp"
		FROM portfolio_allocation_fact pa
		JOIN asset ass ON ass.id = pa.asset_id
		JOIN portfolio_allocation_obs_time paot
		  ON pa.portfolio_id = paot.portfolio_id AND pa.observation_time_id = paot.id
		` + rdbms.WhereClausePlaceholder + `
		ORDER BY paot.observation_timestamp DESC, paot.id DESC,
			pa.class ASC, pa.cash_reserve DESC, pa.total_market_value DESC
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
		INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		VALUES ($1, $2, $3)
		RETURNING id
    `
	observationTimestampUpdateSQL = `
		UPDATE portfolio_allocation_obs_time
		SET observation_time_tag = $1,
		    observation_timestamp = $2
		WHERE id = $3 AND portfolio_id = $4
	`
)

const (
	portfolioIdWhereClause = "AND pa.portfolio_id = {:portfolioId}"
)

const (
	queryAllocationsError           = "Error querying portfolio allocations"
	queryObservationTimestampsError = "Error querying observation timestamps"
	observationTagUniqueConstraint  = "portfolio_allocation_obs_time_portfolio_tag_uk"
)

type PortfolioAllocationRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
}

// FindAllPortfolioAllocationsWithinObservationTimestampsLimit retrieves portfolio allocations for
// the most recent observation timestamps up to the provided limit, projecting the first persisted
// external asset reference into the returned allocation read model.
//
// Example:
//
//	allocations, err := repository.FindAllPortfolioAllocationsWithinObservationTimestampsLimit(1, 12)
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func (repository *PortfolioAllocationRDBMSRepository) FindAllPortfolioAllocationsWithinObservationTimestampsLimit(
	id int64,
	observationTimestampsLimit int,
) ([]*domain.PortfolioAllocation, error) {

	var query = availableObservationTimestampsComplement + portfolioAllocationsSQL

	var queryResult []portfolioAllocationJoinedRowDTS
	err := rdbms.BuildQuery[portfolioAllocationJoinedRowDTS](repository.dbAdapter, query).
		AddParam("observationTimestampLimit", observationTimestampsLimit).
		AddWhereClause("AND pa.observation_time_id IN (SELECT id FROM observation_timestamps)").
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", id).
		Build().FindInto(&queryResult)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, queryAllocationsError, repository)
	}

	return mapPortfolioAllocationRows(queryResult), nil
}

// FindPortfolioAllocationsByObservationTimestamp retrieves portfolio allocations for a single
// observation timestamp, projecting the first persisted external asset reference into the returned
// allocation read model.
//
// Example:
//
//	allocations, err := repository.FindPortfolioAllocationsByObservationTimestamp(1, 3)
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func (repository *PortfolioAllocationRDBMSRepository) FindPortfolioAllocationsByObservationTimestamp(
	id int64,
	observationTimestampId int64,
) ([]*domain.PortfolioAllocation, error) {

	var queryResult []portfolioAllocationJoinedRowDTS
	err := rdbms.BuildQuery[portfolioAllocationJoinedRowDTS](repository.dbAdapter, portfolioAllocationsSQL).
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

	return mapPortfolioAllocationRows(queryResult), nil
}

func (repository *PortfolioAllocationRDBMSRepository) FindAvailableObservationTimestamps(
	portfolioId int64,
	observationTimestampsLimit int,
) ([]*domain.PortfolioObservationTimestamp, error) {

	var query = availableObservationTimestampsSQL

	var queryResult []domain.PortfolioObservationTimestamp
	err := rdbms.BuildQuery[domain.PortfolioObservationTimestamp](repository.dbAdapter, query).
		AddWhereClauseAndParam(portfolioIdWhereClause, "portfolioId", portfolioId).
		AddParam("observationTimestampLimit", observationTimestampsLimit).
		Build().FindInto(&queryResult)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, queryObservationTimestampsError, repository)
	}

	var result = langext.ToPointerSlice(queryResult)

	return result, nil
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

// InsertObservationTimestampInTransaction creates a new portfolio-owned observation timestamp.
// For example, `repository.InsertObservationTimestampInTransaction(ctx, 42, observation)` inserts for portfolio 42.
// Co-authored by: GPT-6.1 Sol, OpenCode and benizzio
func (repository *PortfolioAllocationRDBMSRepository) InsertObservationTimestampInTransaction(
	transContext context.Context,
	portfolioId int64,
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
		AddParams(portfolioId, observationTimestamp.TimeTag, observationTimestamp.Timestamp).
		Build().
		Get(rdbms.ReturningIntIdSingleRowScanner)

	if err != nil {
		if classifiedError := classifyObservationWriteError(err); classifiedError != nil {
			return nil, classifiedError
		}
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error inserting portfolio observation timestamp",
			repository,
		)
	}

	return &domain.PortfolioObservationTimestamp{
		Id:          id,
		PortfolioId: portfolioId,
		TimeTag:     observationTimestamp.TimeTag,
		Timestamp:   observationTimestamp.Timestamp,
	}, nil
}

// UpdateObservationTimestampInTransaction persists metadata already validated by the domain service.
// A guarded update affecting no rows returns a persistence error rather than inserting an observation implicitly.
// For example, `repository.UpdateObservationTimestampInTransaction(ctx, 42, observation)` updates portfolio 42.
// Co-authored by: GPT-6.1 Sol and OpenCode
func (repository *PortfolioAllocationRDBMSRepository) UpdateObservationTimestampInTransaction(
	transContext context.Context,
	portfolioId int64,
	observationTimestamp *domain.PortfolioObservationTimestamp,
) error {
	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return infra.BuildAppError("Context is not a SQL transactional context", repository)
	}

	result, err := repository.dbAdapter.ExecuteInTransaction(
		transactionalContext,
		observationTimestampUpdateSQL,
		observationTimestamp.TimeTag,
		observationTimestamp.Timestamp,
		observationTimestamp.Id,
		portfolioId,
	)
	if err != nil {
		if classifiedError := classifyObservationWriteError(err); classifiedError != nil {
			return classifiedError
		}
		return infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error updating portfolio observation timestamp",
			repository,
		)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error checking updated portfolio observation timestamp",
			repository,
		)
	}
	if rowsAffected != 1 {
		return &infra.UnexpectedRowsAffectedError{
			Operation: "Update portfolio observation timestamp", Expected: 1, Actual: rowsAffected,
		}
	}
	return nil
}

// classifyObservationWriteError classifies database uniqueness and foreign-key failures at the persistence boundary.
// Co-authored by: GPT-6.1 Sol and OpenCode
func classifyObservationWriteError(err error) error {
	if rdbms.IsUniqueConstraintViolation(err, observationTagUniqueConstraint) {
		return infra.BuildUniqueConstraintViolationError(
			observationTagUniqueConstraint,
			"Observation tag already exists in this portfolio.",
			[]string{"Choose a different observationTimestamp.timeTag."},
		)
	}

	var violation = rdbms.ClassifyForeignKeyViolation(err)
	if violation != nil {
		return violation
	}
	return nil
}

// FindAndLockObservationTimestampInTransaction locks and reads an observation without deciding domain ownership rules.
// For example, repository.FindAndLockObservationTimestampInTransaction(ctx, 2) returns nil if observation 2 is missing.
// Authored by: GPT-6.1 Sol
func (repository *PortfolioAllocationRDBMSRepository) FindAndLockObservationTimestampInTransaction(
	transContext context.Context,
	observationId int64,
) (*domain.PortfolioObservationTimestamp, error) {
	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return nil, infra.BuildAppError("Context is not a SQL transactional context", repository)
	}
	var observation domain.PortfolioObservationTimestamp
	var err = transactionalContext.GetTransaction().QueryRowContext(
		transContext,
		`SELECT id, portfolio_id, observation_time_tag, observation_timestamp
		 FROM portfolio_allocation_obs_time WHERE id = $1 FOR UPDATE`,
		observationId,
	).Scan(&observation.Id, &observation.PortfolioId, &observation.TimeTag, &observation.Timestamp)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err, "Error reading portfolio observation timestamp", repository,
		)
	}
	return &observation, nil
}

func BuildPortfolioAllocationRepository(dbAdapter rdbms.RepositoryRDBMSAdapter) *PortfolioAllocationRDBMSRepository {
	return &PortfolioAllocationRDBMSRepository{dbAdapter: dbAdapter}
}
