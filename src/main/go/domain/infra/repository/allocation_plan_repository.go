package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/rdbms"
	"github.com/benizzio/open-asset-allocator/langext"
)

const (
	allocationPlanIdentifierSQL = `
		SELECT ap.id, ap.name
		FROM allocation_plan ap
	` + rdbms.WhereClausePlaceholder + `
		ORDER BY ap.create_timestamp DESC
	`
	allocationPlanSQL = `
		SELECT 
		    ap.id AS allocation_plan_id,
		    ap.name, 
		    ap.type, 
		    ap.planned_execution_date,
		    pa.id AS planned_allocation_id,
		    pa.hierarchical_id, 
		    pa.cash_reserve, 
		    pa.slice_size_percentage
		FROM planned_allocation pa 
		JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
	` + rdbms.WhereClausePlaceholder + `
		ORDER BY ap.create_timestamp DESC, pa.cash_reserve DESC, pa.slice_size_percentage DESC
	`
	allocationPlanInsertSQL = `
		INSERT INTO allocation_plan (portfolio_id, name, type)
		VALUES ($1, $2, 'ALLOCATION_PLAN')
    `
	allocationPlanUpdateSQL = `
		UPDATE allocation_plan 
		SET name = $1
		WHERE id = $2
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
		ON pa.id = temp.id
		WHEN NOT MATCHED BY TARGET THEN
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
				pa.cash_reserve = temp.cash_reserve,
				pa.slice_size_percentage = temp.slice_size_percentage,
		WHEN NOT MATCHED BY SOURCE AND pa.allocation_plan_id = $1 THEN
			DELETE
	`
)

type AllocationPlanRDBMSRepository struct {
	dbAdapter rdbms.RepositoryRDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans(
	portfolioId int,
	planType *allocation.PlanType,
) ([]*domain.AllocationPlan, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(allocationPlanSQL)

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

func (repository *AllocationPlanRDBMSRepository) GetAllocationPlan(id int) (*domain.AllocationPlan, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(allocationPlanSQL)
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
	portfolioId int,
	planType *allocation.PlanType,
) ([]*domain.AllocationPlanIdentifier, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(allocationPlanIdentifierSQL)

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

func BuildAllocationPlanRepository(dbAdapter rdbms.RepositoryRDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
