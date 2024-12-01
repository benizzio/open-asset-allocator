package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/shopspring/decimal"
	"time"
)

// ================================================
// TYPES
// ================================================

type AllocationDTS struct {
	StructuralId        []*string       `json:"structuralId,omitempty"`
	CashReserve         bool            `json:"cashReserve,omitempty"`
	SliceSizePercentage decimal.Decimal `json:"sliceSizePercentage,omitempty"`
}

type AllocationPlanHierarchyLevelDTS struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationPlanStructureDTS struct {
	Hierarchy []AllocationPlanHierarchyLevelDTS `json:"hierarchy,omitempty"`
}

type AllocationPlanDTS struct {
	Id                   int                        `json:"id,omitempty"`
	Name                 string                     `json:"name,omitempty"`
	Type                 string                     `json:"type,omitempty"`
	Structure            AllocationPlanStructureDTS `json:"structure,omitempty"`
	PlannedExecutionDate *time.Time                 `json:"plannedExecutionDate,omitempty"`
	Details              []AllocationDTS            `json:"details,omitempty"`
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

	var structure = mapAllocationPlanStructure(allocationPlan.Structure)
	var allocations = mapAllocations(allocationPlan)

	var allocationPlanDTS = AllocationPlanDTS{
		Id:                   allocationPlan.Id,
		Name:                 allocationPlan.Name,
		Type:                 allocationPlan.PlanType.String(),
		Structure:            structure,
		PlannedExecutionDate: allocationPlan.PlannedExecutionDate,
		Details:              allocations,
	}

	return &allocationPlanDTS
}

func mapAllocationPlanStructure(allocationPlanStructure domain.AllocationPlanStructure) AllocationPlanStructureDTS {
	var hierarchyLevels = mapHierarchyLevels(allocationPlanStructure.Hierarchy)
	return AllocationPlanStructureDTS{Hierarchy: hierarchyLevels}
}

func mapHierarchyLevels(
	levels []domain.AllocationPlanHierarchyLevel,
) []AllocationPlanHierarchyLevelDTS {

	var dtsLevels = make([]AllocationPlanHierarchyLevelDTS, 0)

	for _, level := range levels {
		var levelDTS = AllocationPlanHierarchyLevelDTS{
			Name:  level.Name,
			Field: level.Field,
		}
		dtsLevels = append(dtsLevels, levelDTS)
	}

	return dtsLevels
}

func mapAllocations(allocationPlan *domain.AllocationPlan) []AllocationDTS {
	var allocations = make([]AllocationDTS, 0)
	for _, detail := range allocationPlan.Details {
		var allocation = mapAllocation(detail)
		allocations = append(allocations, allocation)
	}
	return allocations
}

func mapAllocation(allocation *domain.AllocationPlanUnit) AllocationDTS {
	return AllocationDTS{
		StructuralId:        allocation.StructuralId,
		CashReserve:         allocation.CashReserve,
		SliceSizePercentage: allocation.SliceSizePercentage,
	}
}
