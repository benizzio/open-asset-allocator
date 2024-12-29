package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/sqlext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/shopspring/decimal"
)

type PlannedAllocationJoinedRow struct {
	AllocationPlanId     int
	Name                 string
	Type                 allocation.PlanType
	PlannedExecutionDate sqlext.NullTime
	StructuralId         sqlext.NullStringSlice
	CashReserve          bool
	SliceSizePercentage  decimal.Decimal
}

type AllocationPlanRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans(portfolioId int, planType *allocation.PlanType) (
	[]*domain.AllocationPlan,
	error,
) {
	var query = `
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
	` + infra.WhereClausePlaceholder

	var queryBuilder = repository.dbAdapter.BuildQuery(query)

	queryBuilder.AddWhereClauseAndParam("AND ap.portfolio_id = {:portfolioId}", "portfolioId", portfolioId)

	if planType != nil {
		queryBuilder.AddWhereClauseAndParam("AND ap.type = {:planType}", "planType", planType.String())
	}

	result, err := queryBuilder.Build().GetRows()
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error querying allocation plans", repository)
	}

	var refs, err2 = repository.mapPlannedAllocationRows(result)
	var vals []*domain.AllocationPlan
	for _, ref := range refs {
		vals = append(vals, ref)
	}
	return vals, err2
}

func (repository *AllocationPlanRDBMSRepository) mapPlannedAllocationRows(rows *dbx.Rows) (
	[]*domain.AllocationPlan,
	error,
) {

	var allocationPlanCacheMap = make(map[int]*domain.AllocationPlan)
	var allocationPlans = make([]*domain.AllocationPlan, 0)

	for rows.Next() {

		allocationPlan, err := repository.mapRow(rows, allocationPlanCacheMap)
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
	rows *dbx.Rows,
	allocationPlanCacheMap map[int]*domain.AllocationPlan,
) (*domain.AllocationPlan, error) {

	row, err := repository.scanRow(rows)
	if err != nil {
		return nil, err
	}

	allocationPlanUnit := mapPlannedAllocationFromRow(row)

	if cachedAllocationPlan, exists := allocationPlanCacheMap[row.AllocationPlanId]; exists {
		cachedAllocationPlan.AddDetail(allocationPlanUnit)
	} else {
		allocationPlan := mapPlanFromRow(row, allocationPlanUnit)
		allocationPlanCacheMap[row.AllocationPlanId] = &allocationPlan
		return &allocationPlan, nil
	}

	return nil, nil
}

func (repository *AllocationPlanRDBMSRepository) scanRow(
	rows *dbx.Rows,
) (*PlannedAllocationJoinedRow, error) {

	var row PlannedAllocationJoinedRow
	err := rows.ScanStruct(&row)

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(
			err,
			"Error mapping allocation plan unit from row",
			repository,
		)
	}
	return &row, nil
}

func mapPlanFromRow(
	row *PlannedAllocationJoinedRow,
	plannedAllocation *domain.PlannedAllocation,
) domain.AllocationPlan {

	var allocationPlan = domain.AllocationPlan{
		Id:                   row.AllocationPlanId,
		Name:                 row.Name,
		PlanType:             row.Type,
		PlannedExecutionDate: row.PlannedExecutionDate.ToTimeReference(),
	}
	allocationPlan.AddDetail(plannedAllocation)

	return allocationPlan
}

func mapPlannedAllocationFromRow(row *PlannedAllocationJoinedRow) *domain.PlannedAllocation {
	return &domain.PlannedAllocation{
		StructuralId:        row.StructuralId.ToStringSlice(),
		CashReserve:         row.CashReserve,
		SliceSizePercentage: row.SliceSizePercentage,
	}
}

func BuildAllocationPlanRepository(dbAdapter *infra.RDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
