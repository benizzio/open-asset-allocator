package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/shopspring/decimal"
	"time"
)

// ================================================
// TYPES
// ================================================

type PlannedAllocationDTS struct {
	StructuralId        []*string       `json:"structuralId,omitempty"`
	CashReserve         bool            `json:"cashReserve"`
	SliceSizePercentage decimal.Decimal `json:"sliceSizePercentage,omitempty"`
}

type AllocationPlanDTS struct {
	Id                   int                    `json:"id,omitempty"`
	Name                 string                 `json:"name,omitempty"`
	Type                 string                 `json:"type,omitempty"`
	PlannedExecutionDate *time.Time             `json:"plannedExecutionDate,omitempty"`
	Details              []PlannedAllocationDTS `json:"details,omitempty"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func MapAllocationPlans(allocationPlans []*domain.AllocationPlan) []AllocationPlanDTS {
	var allocationPlansDTS = make([]AllocationPlanDTS, 0)
	for _, allocationPlan := range allocationPlans {
		var allocationPlanDTS = mapAllocationPlan(allocationPlan)
		allocationPlansDTS = append(allocationPlansDTS, *allocationPlanDTS)
	}
	return allocationPlansDTS
}

func mapAllocationPlan(allocationPlan *domain.AllocationPlan) *AllocationPlanDTS {

	var allocations = mapPlannedAllocations(allocationPlan)

	var allocationPlanDTS = AllocationPlanDTS{
		Id:                   allocationPlan.Id,
		Name:                 allocationPlan.Name,
		Type:                 allocationPlan.PlanType.String(),
		PlannedExecutionDate: allocationPlan.PlannedExecutionDate,
		Details:              allocations,
	}

	return &allocationPlanDTS
}

func mapPlannedAllocations(allocationPlan *domain.AllocationPlan) []PlannedAllocationDTS {
	var allocations = make([]PlannedAllocationDTS, 0)
	for _, detail := range allocationPlan.Details {
		var allocation = mapPlannedAllocation(detail)
		allocations = append(allocations, allocation)
	}
	return allocations
}

func mapPlannedAllocation(allocation *domain.PlannedAllocation) PlannedAllocationDTS {
	return PlannedAllocationDTS{
		StructuralId:        allocation.StructuralId,
		CashReserve:         allocation.CashReserve,
		SliceSizePercentage: allocation.SliceSizePercentage,
	}
}
