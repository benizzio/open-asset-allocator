package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
)

type AllocationPlanRDBMSRepository struct {
	dbAdapter *infra.RDBMSAdapter
}

func (repository *AllocationPlanRDBMSRepository) GetAllAllocationPlans() ([]domain.AllocationPlan, error) {
	var query = `
		SELECT 
		    ap.name, 
		    ap.type, 
		    ap.structure, 
		    ap.planned_execution_date, 
		    apu.structural_id, 
		    apu.cash_reserve, 
		    apu.slice_size_percentage
		FROM allocation_plan ap
		JOIN allocation_plan_unit apu ON apu.allocation_plan_id = ap.id
	`

	var result dbx.NullStringMap
	err := repository.dbAdapter.BuildQuery(query).FindInto(&result)

	//TODO implement mapping
	for _, row := range result {
		println(row)
	}

	var allocationPlans []domain.AllocationPlan
	return allocationPlans, infra.PropagateAsAppErrorWithNewMessage(err, "Error querying allocation plans", repository)
}

func BuildAllocationPlanRepository(dbAdapter *infra.RDBMSAdapter) *AllocationPlanRDBMSRepository {
	return &AllocationPlanRDBMSRepository{dbAdapter: dbAdapter}
}
