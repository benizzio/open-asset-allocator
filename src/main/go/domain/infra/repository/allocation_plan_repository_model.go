package repository

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra/sqlext"
	"github.com/shopspring/decimal"
)

type plannedAllocationJoinedRowDTS struct {
	AllocationPlanId     int64
	Name                 string
	Type                 allocation.PlanType
	PlannedExecutionDate sqlext.NullTime
	PlannedAllocationId  int64
	HierarchicalId       sqlext.NullStringSlice
	CashReserve          bool
	SliceSizePercentage  decimal.Decimal
}

func mapPlannedAllocationRows(rows []plannedAllocationJoinedRowDTS) ([]*domain.AllocationPlan, error) {

	var allocationPlanCacheMap = make(map[int64]*domain.AllocationPlan)
	var allocationPlans = make([]*domain.AllocationPlan, 0)

	for _, row := range rows {

		allocationPlan := mapPlannedAllocationRow(&row, allocationPlanCacheMap)
		if allocationPlan != nil {
			allocationPlans = append(allocationPlans, allocationPlan)
		}
	}

	return allocationPlans, nil
}

func mapPlannedAllocationRow(
	row *plannedAllocationJoinedRowDTS,
	allocationPlanCacheMap map[int64]*domain.AllocationPlan,
) *domain.AllocationPlan {

	allocationPlanUnit := buildPlannedAllocationFromRow(row)

	if cachedAllocationPlan, exists := allocationPlanCacheMap[row.AllocationPlanId]; exists {
		cachedAllocationPlan.AddDetail(allocationPlanUnit)
	} else {
		allocationPlan := buildAllocationPlanFromRow(row, allocationPlanUnit)
		allocationPlanCacheMap[row.AllocationPlanId] = allocationPlan
		return allocationPlan
	}

	return nil
}

func buildPlannedAllocationFromRow(rowDTS *plannedAllocationJoinedRowDTS) *domain.PlannedAllocation {
	return &domain.PlannedAllocation{
		Id:                  int64(rowDTS.PlannedAllocationId),
		HierarchicalId:      rowDTS.HierarchicalId.ToStringSlice(),
		CashReserve:         rowDTS.CashReserve,
		SliceSizePercentage: rowDTS.SliceSizePercentage,
	}
}

func buildAllocationPlanFromRow(
	rowDTS *plannedAllocationJoinedRowDTS,
	plannedAllocation *domain.PlannedAllocation,
) *domain.AllocationPlan {

	var allocationPlan = domain.AllocationPlan{
		AllocationPlanIdentifier: domain.AllocationPlanIdentifier{
			Id:   int64(rowDTS.AllocationPlanId),
			Name: rowDTS.Name,
		},
		PlanType:             rowDTS.Type,
		PlannedExecutionDate: rowDTS.PlannedExecutionDate.ToTimeReference(),
	}
	allocationPlan.AddDetail(plannedAllocation)

	return &allocationPlan
}
