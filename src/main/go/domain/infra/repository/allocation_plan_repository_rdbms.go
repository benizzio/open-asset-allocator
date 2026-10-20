package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lib/pq"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
)

// Co-authored by: OpenCode
const (
	allocationPlanIdentifierSQL = `
		SELECT ap.id, ap.name, ap.planned_date
		FROM allocation_plan ap
	` + rdbms.WhereClausePlaceholder + `
		ORDER BY ap.planned_date DESC, ap.id DESC
	`
	allocationPlanSQL = `
		SELECT 
		    ap.id AS allocation_plan_id,
		    ap.name, 
		    ap.type, 
		    ap.planned_date,
		    ap.planned_execution_date,
		    pa.id AS planned_allocation_id,
		    pa.hierarchical_id, 
		    pa.cash_reserve, 
		    pa.slice_size_percentage,
		    coalesce(ass.id, 0) AS "asset.id", 
		    coalesce(ass.ticker, '') AS "asset.ticker", 
		    coalesce(ass.name, '') AS "asset.name"
		FROM planned_allocation pa 
		JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
		LEFT JOIN asset ass ON ass.id = pa.asset_id
	` + rdbms.WhereClausePlaceholder + `
		ORDER BY ap.planned_date DESC, ap.id DESC, pa.cash_reserve DESC, pa.slice_size_percentage DESC
	`
	allocationPlanInsertSQL = `
		INSERT INTO allocation_plan (portfolio_id, name, type, planned_date)
		VALUES ($1, $2, $3, $4) RETURNING id
    `
	allocationPlanUpdateSQL = `
		UPDATE allocation_plan 
		SET name = $1, planned_date = $2
		WHERE id = $3 AND portfolio_id = $4 AND type = $5
	`
	plannedAllocationTempTableName   = "planned_allocation_merge_temp"
	plannedAllocationTempTableDDLSQL = `
		CREATE TEMPORARY TABLE ` + plannedAllocationTempTableName + `
		(LIKE planned_allocation INCLUDING DEFAULTS)
		ON COMMIT DROP
	`
	plannedAllocationMergeSQL = `
		MERGE INTO planned_allocation pa
		USING ` + plannedAllocationTempTableName + ` temp
		ON pa.id = temp.id AND pa.allocation_plan_id = temp.allocation_plan_id
		WHEN NOT MATCHED BY TARGET AND temp.id = 0 THEN
			INSERT (allocation_plan_id, hierarchical_id, cash_reserve, slice_size_percentage, asset_id)
			VALUES (
				temp.allocation_plan_id, 
				temp.hierarchical_id, 
				temp.cash_reserve, 
				temp.slice_size_percentage, 
				temp.asset_id
			)
		WHEN MATCHED THEN
			UPDATE SET 
				hierarchical_id = temp.hierarchical_id, 
				cash_reserve = temp.cash_reserve,
				slice_size_percentage = temp.slice_size_percentage
		WHEN NOT MATCHED BY SOURCE AND pa.allocation_plan_id = $1 THEN
			DELETE
	`
)

type AllocationPlanRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans(
	portfolioId int64,
	planType *allocation.PlanType,
) ([]*domain.AllocationPlan, error) {

	var queryBuilder = rdbms.BuildQuery[plannedAllocationJoinedRowDTS](repository.dbAdapter, allocationPlanSQL)

	queryBuilder.AddWhereClauseAndParam("AND ap.portfolio_id = {:portfolioId}", "portfolioId", portfolioId)

	if planType != nil {
		queryBuilder.AddWhereClauseAndParam("AND ap.type = {:planType}", "planType", planType.String())
	}

	var queryResult []plannedAllocationJoinedRowDTS
	queryError := queryBuilder.Build().FindInto(&queryResult)
	if queryError != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(queryError, "Error querying allocation plans", repository)
	}

	return mapPlannedAllocationRows(queryResult)
}

func (repository *AllocationPlanRDBMSRepository) GetAllocationPlan(id int64) (*domain.AllocationPlan, error) {

	var queryBuilder = rdbms.BuildQuery[plannedAllocationJoinedRowDTS](repository.dbAdapter, allocationPlanSQL)
	queryBuilder.AddWhereClauseAndParam("AND ap.id = {:id}", "id", id)

	var queryResult []plannedAllocationJoinedRowDTS
	queryError := queryBuilder.Build().FindInto(&queryResult)
	if queryError != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(queryError, "Error querying allocation plan", repository)
	}

	var refs, mappingError = mapPlannedAllocationRows(queryResult)
	if mappingError != nil {
		return nil, mappingError
	}

	if len(refs) == 0 {
		return nil, infra.BuildAppErrorFormatted(repository, "Allocation plan with id %d not found", id)
	}

	return refs[0], nil
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlanIdentifiers(
	portfolioId int64,
	planType *allocation.PlanType,
) ([]*domain.AllocationPlanIdentifier, error) {

	var queryBuilder = rdbms.BuildQuery[domain.AllocationPlanIdentifier](
		repository.dbAdapter, allocationPlanIdentifierSQL,
	)

	queryBuilder.AddWhereClauseAndParam("AND ap.portfolio_id = {:portfolioId}", "portfolioId", portfolioId)

	if planType != nil {
		queryBuilder.AddWhereClauseAndParam("AND ap.type = {:planType}", "planType", planType.String())
	}

	var queryResult []domain.AllocationPlanIdentifier
	queryError := queryBuilder.Build().FindInto(&queryResult)
	if queryError != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			queryError,
			"Error querying allocation plan identifiers",
			repository,
		)
	}

	return langext.ToPointerSlice(queryResult), nil
}

// InsertAllocationPlanInTransaction creates a plan and its details as one transaction.
// For example, `repository.InsertAllocationPlanInTransaction(ctx, plan)` creates the plan for plan.PortfolioId.
// Co-authored by: GPT-6.1 Sol, OpenCode and benizzio
func (repository *AllocationPlanRDBMSRepository) InsertAllocationPlanInTransaction(
	transContext context.Context,
	plan *domain.AllocationPlan,
) error {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

	id, err := rdbms.BuildQueryInTransaction[int64](transactionalContext, allocationPlanInsertSQL).
		AddParams(plan.PortfolioId, plan.Name, plan.PlanType.String(), plan.PlannedDate).
		Build().
		Get(rdbms.ReturningIntIdSingleRowScanner)
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(err, "Error inserting allocation plan", repository)
	}

	return repository.mergePlannedAllocationsInTransaction(transactionalContext, id, plan.Details)
}

// UpdateAllocationPlanInTransaction persists a plan already validated by the domain service.
// For example, `repository.UpdateAllocationPlanInTransaction(ctx, plan)` updates plan.Id only in plan.PortfolioId.
// Co-authored by: GPT-6.1 Sol, OpenCode and benizzio
func (repository *AllocationPlanRDBMSRepository) UpdateAllocationPlanInTransaction(
	transContext context.Context,
	plan *domain.AllocationPlan,
) error {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

	result, err := repository.dbAdapter.ExecuteInTransaction(
		transactionalContext,
		allocationPlanUpdateSQL,
		plan.Name,
		plan.PlannedDate,
		plan.Id,
		plan.PortfolioId,
		plan.PlanType.String(),
	)
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(err, "Error updating allocation plan", repository)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(err, "Error checking updated allocation plan", repository)
	}
	if rowsAffected != 1 {
		return &infra.UnexpectedRowsAffectedError{
			Operation: "Update allocation plan", Expected: 1, Actual: rowsAffected,
		}
	}

	return repository.mergePlannedAllocationsInTransaction(transactionalContext, int64(plan.Id), plan.Details)
}

// FindAndLockAllocationPlanOwnershipInTransaction locks a plan and reads its ownership and detail identifiers.
// For example, repository.FindAndLockAllocationPlanOwnershipInTransaction(ctx, 6) returns nil if plan 6 is missing.
// Authored by: GPT-6.1 Sol
func (repository *AllocationPlanRDBMSRepository) FindAndLockAllocationPlanOwnershipInTransaction(
	transContext context.Context,
	allocationPlanId int64,
) (*domain.AllocationPlanOwnership, error) {
	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return nil, infra.BuildAppError("Context is not a SQL transactional context", repository)
	}
	var ownership domain.AllocationPlanOwnership
	var err = transactionalContext.GetTransaction().QueryRowContext(
		transContext,
		`SELECT portfolio_id, type FROM allocation_plan WHERE id = $1 FOR UPDATE`,
		allocationPlanId,
	).Scan(&ownership.PortfolioId, &ownership.PlanType)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error reading allocation plan ownership", repository)
	}
	err = transactionalContext.GetTransaction().QueryRowContext(
		transContext,
		`SELECT COALESCE(array_agg(id::BIGINT), '{}'::BIGINT[]) FROM planned_allocation WHERE allocation_plan_id = $1`,
		allocationPlanId,
	).Scan(pq.Array(&ownership.DetailIds))
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err, "Error reading planned allocation identifiers", repository,
		)
	}
	return &ownership, nil
}

func (repository *AllocationPlanRDBMSRepository) mergePlannedAllocationsInTransaction(
	transContext *rdbms.SQLTransactionalContext,
	allocationPlanId int64,
	plannedAllocations []*domain.PlannedAllocation,
) error {

	var transactionalContext, ok = rdbms.ToSQLTransactionalContext(transContext)
	if !ok {
		return infra.BuildAppError(
			"Context is not a SQL transactional context",
			repository,
		)
	}

	err := repository.insertPlannedAllocationsInTempTable(
		transactionalContext,
		plannedAllocations,
		allocationPlanId,
	)
	if err != nil {
		return err
	}

	return repository.mergePlannedAllocations(transactionalContext, allocationPlanId)
}

func (repository *AllocationPlanRDBMSRepository) insertPlannedAllocationsInTempTable(
	transactionalContext *rdbms.SQLTransactionalContext,
	plannedAllocations []*domain.PlannedAllocation,
	allocationPlanId int64,
) error {

	_, err := repository.dbAdapter.ExecuteInTransaction(transactionalContext, plannedAllocationTempTableDDLSQL)
	if err != nil {
		return infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error creating temporary table for planned allocations",
			repository,
		)
	}

	columns, insertValues := repository.preparePlannedAllocationTempInserts(plannedAllocations, allocationPlanId)

	err = repository.dbAdapter.InsertBulkInTransaction(
		transactionalContext,
		plannedAllocationTempTableName,
		columns,
		insertValues,
	)
	return infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error inserting planned allocations into temporary table",
		repository,
	)
}

func (repository *AllocationPlanRDBMSRepository) preparePlannedAllocationTempInserts(
	plannedAllocations []*domain.PlannedAllocation,
	allocationPlanId int64,
) ([]string, [][]any) {

	var columns = []string{
		"id",
		"allocation_plan_id",
		"hierarchical_id",
		"cash_reserve",
		"slice_size_percentage",
		"asset_id",
	}

	var insertValues = make([][]any, len(plannedAllocations))
	for i, plannedAllocation := range plannedAllocations {

		var assetId *int64
		if plannedAllocation.Asset != nil {
			assetId = &plannedAllocation.Asset.Id
		}

		insertValues[i] = []any{
			plannedAllocation.Id,
			allocationPlanId,
			plannedAllocation.HierarchicalId,
			plannedAllocation.CashReserve,
			plannedAllocation.SliceSizePercentage,
			assetId,
		}
	}

	return columns, insertValues
}

func (repository *AllocationPlanRDBMSRepository) mergePlannedAllocations(
	transactionalContext *rdbms.SQLTransactionalContext,
	allocationPlanId int64,
) error {
	_, err := repository.dbAdapter.ExecuteInTransaction(
		transactionalContext,
		plannedAllocationMergeSQL,
		allocationPlanId,
	)

	return infra.PropagateAsAppErrorWithNewMessage(
		err,
		"Error merging planned allocations",
		repository,
	)
}

func BuildAllocationPlanRepository(dbAdapter rdbms.RepositoryRDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
