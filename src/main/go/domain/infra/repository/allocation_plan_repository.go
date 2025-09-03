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
		    pa.structural_id, 
		    pa.cash_reserve, 
		    pa.slice_size_percentage
		FROM planned_allocation pa 
		JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
	` + rdbms.WhereClausePlaceholder + `
		ORDER BY ap.create_timestamp DESC, pa.cash_reserve DESC, pa.slice_size_percentage DESC
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
