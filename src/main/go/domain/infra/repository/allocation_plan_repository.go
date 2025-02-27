package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/sqlext"
	"github.com/benizzio/open-asset-allocator/infra/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/shopspring/decimal"
)

type PlannedAllocationJoinedRowDTS struct {
	AllocationPlanId     int
	Name                 string
	Type                 allocation.PlanType
	PlannedExecutionDate sqlext.NullTime
	StructuralId         sqlext.NullStringSlice
	CashReserve          bool
	SliceSizePercentage  decimal.Decimal
}

const (
	allocationPlanIdentifierSQL = `
		SELECT ap.id, ap.name
		FROM allocation_plan ap
	` + infra.WhereClausePlaceholder
	allocationPlanSQL = `
		SELECT 
		    ap.id as allocation_plan_id,
		    ap.name, 
		    ap.type, 
		    ap.planned_execution_date, 
		    pa.structural_id, 
		    pa.cash_reserve, 
		    pa.slice_size_percentage
		FROM planned_allocation pa 
		JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
	` + infra.WhereClausePlaceholder + `
		ORDER BY pa.cash_reserve DESC, pa.slice_size_percentage DESC
`
)

type AllocationPlanRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans(portfolioId int, planType *allocation.PlanType) (
	[]*domain.AllocationPlan,
	error,
) {

	var queryBuilder = repository.dbAdapter.BuildQuery(allocationPlanSQL)

	queryBuilder.AddWhereClauseAndParam("AND ap.portfolio_id = {:portfolioId}", "portfolioId", portfolioId)

	if planType != nil {
		queryBuilder.AddWhereClauseAndParam("AND ap.type = {:planType}", "planType", planType.String())
	}

	var queryResult []PlannedAllocationJoinedRowDTS
	queryError := queryBuilder.Build().FindInto(&queryResult)
	if queryError != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(queryError, "Error querying allocation plans", repository)
	}

	return repository.mapPlannedAllocationRows(queryResult)
}

func (repository *AllocationPlanRDBMSRepository) GetAllocationPlan(id int) (*domain.AllocationPlan, error) {

	var queryBuilder = repository.dbAdapter.BuildQuery(allocationPlanSQL)
	queryBuilder.AddWhereClauseAndParam("AND ap.id = {:id}", "id", id)

	var queryResult []PlannedAllocationJoinedRowDTS
	queryError := queryBuilder.Build().FindInto(&queryResult)
	if queryError != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(queryError, "Error querying allocation plan", repository)
	}

	var refs, mappingError = repository.mapPlannedAllocationRows(queryResult)
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

	return util.ToPointerSlice(queryResult), nil
}

func (repository *AllocationPlanRDBMSRepository) mapPlannedAllocationRows(rows []PlannedAllocationJoinedRowDTS) (
	[]*domain.AllocationPlan,
	error,
) {

	var allocationPlanCacheMap = make(map[int]*domain.AllocationPlan)
	var allocationPlans = make([]*domain.AllocationPlan, 0)

	for _, row := range rows {

		allocationPlan, err := repository.mapRow(&row, allocationPlanCacheMap)
		if err != nil {
			return nil, err
		}
		if allocationPlan != nil {
			allocationPlans = append(allocationPlans, allocationPlan)
		}
	}

	return allocationPlans, nil
}

func (repository *AllocationPlanRDBMSRepository) mapRow(
	row *PlannedAllocationJoinedRowDTS,
	allocationPlanCacheMap map[int]*domain.AllocationPlan,
) (*domain.AllocationPlan, error) {

	allocationPlanUnit := mapPlannedAllocationFromRow(row)

	if cachedAllocationPlan, exists := allocationPlanCacheMap[row.AllocationPlanId]; exists {
		cachedAllocationPlan.AddDetail(allocationPlanUnit)
	} else {
		allocationPlan := mapPlanFromRow(row, allocationPlanUnit)
		allocationPlanCacheMap[row.AllocationPlanId] = allocationPlan
		return allocationPlan, nil
	}

	return nil, nil
}

func (repository *AllocationPlanRDBMSRepository) scanRow(
	rows *dbx.Rows,
) (*PlannedAllocationJoinedRowDTS, error) {

	var rowDTS PlannedAllocationJoinedRowDTS
	err := rows.ScanStruct(&rowDTS)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error mapping allocation plan unit from rowDTS",
			repository,
		)
	}
	return &rowDTS, nil
}

func mapPlanFromRow(
	rowDTS *PlannedAllocationJoinedRowDTS,
	plannedAllocation *domain.PlannedAllocation,
) *domain.AllocationPlan {

	var allocationPlan = domain.AllocationPlan{
		AllocationPlanIdentifier: domain.AllocationPlanIdentifier{
			Id:   rowDTS.AllocationPlanId,
			Name: rowDTS.Name,
		},
		PlanType:             rowDTS.Type,
		PlannedExecutionDate: rowDTS.PlannedExecutionDate.ToTimeReference(),
	}
	allocationPlan.AddDetail(plannedAllocation)

	return &allocationPlan
}

func mapPlannedAllocationFromRow(rowDTS *PlannedAllocationJoinedRowDTS) *domain.PlannedAllocation {
	return &domain.PlannedAllocation{
		StructuralId:        rowDTS.StructuralId.ToStringSlice(),
		CashReserve:         rowDTS.CashReserve,
		SliceSizePercentage: rowDTS.SliceSizePercentage,
	}
}

func BuildAllocationPlanRepository(dbAdapter *infra.RDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
