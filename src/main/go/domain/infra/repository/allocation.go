package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type AllocationPlanRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans() ([]domain.AllocationPlan, error) {
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
	`

	result, err := repository.dbAdapter.BuildQuery(query).GetRows()

	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "Error querying allocation plans", repository)
	}

	return repository.mapAllocationPlans(result)
}

func (repository *AllocationPlanRDBMSRepository) mapAllocationPlans(rows *dbx.Rows) ([]domain.AllocationPlan, error) {
	//TODO clean
	var allocationPlanMap = make(map[int]domain.AllocationPlan)
	var allocationPlans []domain.AllocationPlan

	for rows.Next() {

		var tempAllocationPlan domain.AllocationPlan
		var allocationPlanUnit domain.AllocationPlanUnit
		var planId int
		var planTypeName string

		err := rows.Scan(
			&planId,
			&tempAllocationPlan.Name,
			&planTypeName,
			&tempAllocationPlan.Structure, //TODO json mapping not working
			&tempAllocationPlan.PlannedExecutionDate,
			&allocationPlanUnit.StructuralId,
			&allocationPlanUnit.CashReserve,
			&allocationPlanUnit.SliceSizePercentage,
		)

		if err != nil {
			return nil, infra.PropagateAsAppErrorWithNewMessage(
				err,
				"Error mapping allocation plan unit from row",
				repository,
			)
		}

		if allocationPlan, exists := allocationPlanMap[planId]; exists {
			allocationPlan.Details = append(allocationPlan.Details, allocationPlanUnit)
		} else {

			planType, err := allocation.GetPlanType(planTypeName)

			if err != nil {
				return nil, infra.PropagateAsAppErrorWithNewMessage(
					err,
					"Error mapping allocation plan type",
					repository,
				)
			}

			tempAllocationPlan.PlanType = planType
			tempAllocationPlan.Details = make([]domain.AllocationPlanUnit, 0)
			allocationPlanMap[planId] = tempAllocationPlan
		}

		allocationPlans = append(allocationPlans, tempAllocationPlan)
	}
	return allocationPlans, nil
}

func BuildAllocationPlanRepository(dbAdapter *infra.RDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
