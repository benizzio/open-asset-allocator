package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/sqlext"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/shopspring/decimal"
)

type AllocationPlanUnitJoinedRow struct {
	AllocationPlanId     int
	Name                 string
	Type                 allocation.PlanType
	Structure            domain.AllocationPlanStructure
	PlannedExecutionDate sqlext.NullTime
	StructuralId         sqlext.NullStringSlice
	CashReserve          bool
	SliceSizePercentage  decimal.Decimal
}

type AllocationPlanRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans(planType *allocation.PlanType) (
	[]*domain.AllocationPlan,
	error,
) {
	var query = `
		SELECT 
		    ap.id as allocation_plan_id,
		    ap.name, 
		    ap.type, 
		    ap.structure, 
		    ap.planned_execution_date, 
		    apu.structural_id, 
		    apu.cash_reserve, 
		    apu.slice_size_percentage
		FROM allocation_plan_unit apu 
		JOIN allocation_plan ap ON apu.allocation_plan_id = ap.id
		/*WHERE+PARAMS*/
	`

	var queryBuilder = repository.dbAdapter.BuildQuery(query)
	if planType != nil {
		queryBuilder.AddWhereClauseAndParam("AND ap.type = {:planType}", "planType", planType.String())
	}

	result, err := queryBuilder.Build().GetRows()
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error querying allocation plans", repository)
	}

	var refs, err2 = repository.mapAllocationPlanUnitRows(result)
	var vals []*domain.AllocationPlan
	for _, ref := range refs {
		vals = append(vals, ref)
	}
	return vals, err2
}

func (repository *AllocationPlanRDBMSRepository) mapAllocationPlanUnitRows(rows *dbx.Rows) (
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

	allocationPlanUnit := mapUnitFromRow(row)

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
) (*AllocationPlanUnitJoinedRow, error) {

	var row AllocationPlanUnitJoinedRow
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
	row *AllocationPlanUnitJoinedRow,
	allocationPlanUnit *domain.AllocationPlanUnit,
) domain.AllocationPlan {

	var allocationPlan = domain.AllocationPlan{
		Id:                   row.AllocationPlanId,
		Name:                 row.Name,
		PlanType:             row.Type,
		Structure:            row.Structure,
		PlannedExecutionDate: row.PlannedExecutionDate.ToTimeReference(),
	}
	allocationPlan.AddDetail(allocationPlanUnit)

	return allocationPlan
}

func mapUnitFromRow(row *AllocationPlanUnitJoinedRow) *domain.AllocationPlanUnit {
	return &domain.AllocationPlanUnit{
		StructuralId:        row.StructuralId.ToStringSlice(),
		CashReserve:         row.CashReserve,
		SliceSizePercentage: row.SliceSizePercentage,
	}
}

func BuildAllocationPlanRepository(dbAdapter *infra.RDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
