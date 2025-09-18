package model

import (
	"time"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/shopspring/decimal"
)

// ================================================
// TYPES
// ================================================

type PlannedAllocationDTS struct {
	Id                  int64           `json:"id,omitempty"`
	HierarchicalId      []*string       `json:"hierarchicalId,omitempty"`
	CashReserve         bool            `json:"cashReserve"`
	SliceSizePercentage decimal.Decimal `json:"sliceSizePercentage,omitempty"`
}

type AllocationPlanDTS struct {
	Id                   int64                  `json:"id,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	PlannedExecutionDate *time.Time             `json:"plannedExecutionDate,omitempty"`
	Details              []PlannedAllocationDTS `json:"details,omitempty"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func MapToAllocationPlanDTSs(allocationPlans []*domain.AllocationPlan) []AllocationPlanDTS {
	var allocationPlansDTS = make([]AllocationPlanDTS, 0)
	for _, allocationPlan := range allocationPlans {
		var allocationPlanDTS = mapToAllocationPlanDTS(allocationPlan)
		allocationPlansDTS = append(allocationPlansDTS, *allocationPlanDTS)
	}
	return allocationPlansDTS
}

func mapToAllocationPlanDTS(allocationPlan *domain.AllocationPlan) *AllocationPlanDTS {

	var allocations = mapToPlannedAllocationDTSs(allocationPlan)

	var allocationPlanDTS = AllocationPlanDTS{
		Id:                   allocationPlan.Id,
		Name:                 allocationPlan.Name,
		Type:                 allocationPlan.PlanType.String(),
		PlannedExecutionDate: allocationPlan.PlannedExecutionDate,
		Details:              allocations,
	}

	return &allocationPlanDTS
}

func mapToPlannedAllocationDTSs(allocationPlan *domain.AllocationPlan) []PlannedAllocationDTS {
	var allocations = make([]PlannedAllocationDTS, 0)
	for _, detail := range allocationPlan.Details {
		var allocation = mapToPlannedAllocationDTS(detail)
		allocations = append(allocations, allocation)
	}
	return allocations
}

func mapToPlannedAllocationDTS(allocation *domain.PlannedAllocation) PlannedAllocationDTS {
	return PlannedAllocationDTS{
		Id:                  allocation.Id,
		HierarchicalId:      allocation.HierarchicalId,
		CashReserve:         allocation.CashReserve,
		SliceSizePercentage: allocation.SliceSizePercentage,
	}
}
