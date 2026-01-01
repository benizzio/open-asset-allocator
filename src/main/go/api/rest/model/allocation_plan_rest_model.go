package model

import (
	"time"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/shopspring/decimal"
)

// ================================================
// TYPES
// ================================================

type PlannedAllocationDTS struct {
	Id                  *langext.ParseableInt64 `json:"id,omitempty"`
	HierarchicalId      []*string               `json:"hierarchicalId,omitempty" validate:"required"`
	CashReserve         bool                    `json:"cashReserve"`
	SliceSizePercentage decimal.Decimal         `json:"sliceSizePercentage,omitempty"`
	Asset               *AssetDTS               `json:"asset,omitempty"`
}

type AllocationPlanDTS struct {
	Id                   *langext.ParseableInt64 `json:"id,omitempty"`
	Name                 string                  `json:"name,omitempty" validate:"required"`
	Type                 string                  `json:"type,omitempty"`
	PlannedExecutionDate *time.Time              `json:"plannedExecutionDate,omitempty"`
	Details              []*PlannedAllocationDTS `json:"details,omitempty" validate:"required,min=1"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func MapToAllocationPlanDTSs(allocationPlans []*domain.AllocationPlan) []*AllocationPlanDTS {
	var allocationPlansDTS = make([]*AllocationPlanDTS, 0)
	for _, allocationPlan := range allocationPlans {
		var allocationPlanDTS = mapToAllocationPlanDTS(allocationPlan)
		allocationPlansDTS = append(allocationPlansDTS, allocationPlanDTS)
	}
	return allocationPlansDTS
}

func mapToAllocationPlanDTS(allocationPlan *domain.AllocationPlan) *AllocationPlanDTS {
	var allocations = mapToPlannedAllocationDTSs(allocationPlan)
	var allocationPlanId = langext.ParseableInt64(allocationPlan.Id)
	return &AllocationPlanDTS{
		Id:                   &allocationPlanId,
		Name:                 allocationPlan.Name,
		Type:                 allocationPlan.PlanType.String(),
		PlannedExecutionDate: allocationPlan.PlannedExecutionDate,
		Details:              allocations,
	}
}

func mapToPlannedAllocationDTSs(allocationPlan *domain.AllocationPlan) []*PlannedAllocationDTS {

	var plannedAllocations = allocationPlan.Details
	var plannedAllocationDTSs = make([]*PlannedAllocationDTS, len(plannedAllocations))

	for index, detail := range plannedAllocations {
		var plannedAllocationDTS = mapToPlannedAllocationDTS(detail)
		plannedAllocationDTSs[index] = plannedAllocationDTS
	}

	return plannedAllocationDTSs
}

func mapToPlannedAllocationDTS(allocation *domain.PlannedAllocation) *PlannedAllocationDTS {
	var assetDTS = MapToAssetDTS(allocation.Asset)
	var plannedAllocationId = langext.ParseableInt64(allocation.Id)
	return &PlannedAllocationDTS{
		Id:                  &plannedAllocationId,
		HierarchicalId:      allocation.HierarchicalId,
		CashReserve:         allocation.CashReserve,
		SliceSizePercentage: allocation.SliceSizePercentage,
		Asset:               assetDTS,
	}
}

func MapToAllocationPlan(
	allocationPlanDTS *AllocationPlanDTS,
	portfolioId int64,
	planType allocation.PlanType,
) (*domain.AllocationPlan, error) {

	if allocationPlanDTS == nil {
		return nil, nil
	}

	var allocations = mapToPlannedAllocations(allocationPlanDTS.Details)

	var plannedExecutionDate *time.Time
	if planType == allocation.BalancingExecutionPlan {
		plannedExecutionDate = allocationPlanDTS.PlannedExecutionDate
	}

	var plannedAllocationId int64
	if allocationPlanDTS.Id != nil {
		plannedAllocationId = int64(*allocationPlanDTS.Id)
	}

	return &domain.AllocationPlan{
		AllocationPlanIdentifier: domain.AllocationPlanIdentifier{
			Id:   plannedAllocationId,
			Name: allocationPlanDTS.Name,
		},
		PortfolioId:          portfolioId,
		PlanType:             planType,
		PlannedExecutionDate: plannedExecutionDate,
		Details:              allocations,
	}, nil
}

func mapToPlannedAllocations(allocationDTSs []*PlannedAllocationDTS) []*domain.PlannedAllocation {
	var allocations = make([]*domain.PlannedAllocation, len(allocationDTSs))
	for index, allocationDTS := range allocationDTSs {
		var plannedAllocation = mapToPlannedAllocation(allocationDTS)
		allocations[index] = plannedAllocation
	}
	return allocations
}

func mapToPlannedAllocation(allocationDTS *PlannedAllocationDTS) *domain.PlannedAllocation {

	if allocationDTS == nil {
		return nil
	}

	var asset = MapToAsset(allocationDTS.Asset)

	var plannedAllocationId int64
	if allocationDTS.Id != nil {
		plannedAllocationId = int64(*allocationDTS.Id)
	}
	return &domain.PlannedAllocation{
		Id:                  plannedAllocationId,
		HierarchicalId:      allocationDTS.HierarchicalId,
		CashReserve:         allocationDTS.CashReserve,
		SliceSizePercentage: allocationDTS.SliceSizePercentage,
		Asset:               asset,
	}
}
