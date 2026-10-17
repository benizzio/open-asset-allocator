package model

import (
	"errors"
	"time"

	"github.com/shopspring/decimal"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/langext"
)

// ================================================
// TYPES
// ================================================

// AllocationPlanAssetDTS is the asset data transfer struct used within allocation plan endpoints.
// Name and Ticker are optional here, unlike AssetDTS used in the asset CRUD endpoints.
//
// Authored by: GitHub Copilot
type AllocationPlanAssetDTS struct {
	Id     *langext.ParseableInt64 `json:"id"`
	Name   string                  `json:"name" validate:"max=100"`
	Ticker string                  `json:"ticker" validate:"max=40"`
}

type PlannedAllocationDTS struct {
	Id                  *langext.ParseableInt64 `json:"id,omitempty"`
	HierarchicalId      []*string               `json:"hierarchicalId,omitempty" validate:"required,min=1"`
	CashReserve         bool                    `json:"cashReserve"`
	SliceSizePercentage decimal.Decimal         `json:"sliceSizePercentage,omitempty"`
	Asset               *AllocationPlanAssetDTS `json:"asset,omitempty"`
}

// AllocationPlanDTS is the REST request and response representation of an allocation plan.
// Create and update requests must include plannedDate in YYYY-MM-DD format, for example:
// `{"name":"Long-term allocation","plannedDate":"2026-10-17","details":[...]}`.
//
// Co-authored by: OpenCode
type AllocationPlanDTS struct {
	Id                   *langext.ParseableInt64 `json:"id,omitempty"`
	Name                 string                  `json:"name,omitempty" validate:"required,max=100"`
	Type                 string                  `json:"type,omitempty" validate:"max=50"`
	PlannedExecutionDate *time.Time              `json:"plannedExecutionDate,omitempty"`
	PlannedDate          AllocationPlanDate      `json:"plannedDate" validate:"required,datetime=2006-01-02"`
	Details              []*PlannedAllocationDTS `json:"details,omitempty" validate:"required,min=1"`
}

// AllocationPlanDate is a calendar date serialized as YYYY-MM-DD without a timezone.
// Use it for allocation-plan plannedDate values, for example `{"plannedDate":"2026-10-17"}`;
// it keeps calendar days independent of browser and server timezones.
//
// Authored by: OpenCode
type AllocationPlanDate string

// newAllocationPlanDate wraps a stored calendar date for JSON serialization.
func newAllocationPlanDate(value time.Time) AllocationPlanDate {
	return AllocationPlanDate(value.Format(time.DateOnly))
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

// mapToAllocationPlanDTS maps a domain allocation plan to the REST representation, preserving its calendar date.
//
// Co-authored by: OpenCode
func mapToAllocationPlanDTS(allocationPlan *domain.AllocationPlan) *AllocationPlanDTS {
	var allocations = mapToPlannedAllocationDTSs(allocationPlan)
	var allocationPlanId = langext.ParseableInt64(allocationPlan.Id)
	return &AllocationPlanDTS{
		Id:                   &allocationPlanId,
		Name:                 allocationPlan.Name,
		Type:                 allocationPlan.PlanType.String(),
		PlannedExecutionDate: allocationPlan.PlannedExecutionDate,
		PlannedDate:          newAllocationPlanDate(allocationPlan.PlannedDate),
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
	var assetDTS = mapToAllocationPlanAssetDTS(allocation.Asset)
	return &PlannedAllocationDTS{
		Id:                  new(langext.ParseableInt64(allocation.Id)),
		HierarchicalId:      allocation.HierarchicalId,
		CashReserve:         allocation.CashReserve,
		SliceSizePercentage: allocation.SliceSizePercentage,
		Asset:               assetDTS,
	}
}

// MapToAllocationPlan converts a validated REST allocation-plan request into a domain plan.
// A request must provide plannedDate as YYYY-MM-DD; for example, `{"plannedDate":"2026-10-17"}`
// maps to a UTC-midnight value that persists as a SQL DATE without timezone conversion.
//
// Co-authored by: OpenCode
func MapToAllocationPlan(
	allocationPlanDTS *AllocationPlanDTS,
	portfolioId int64,
	planType allocation.PlanType,
) (*domain.AllocationPlan, error) {

	if allocationPlanDTS == nil {
		return nil, nil
	}
	plannedDate, err := time.Parse(time.DateOnly, string(allocationPlanDTS.PlannedDate))
	if err != nil || plannedDate.IsZero() || plannedDate.Year() == 0 {
		return nil, errors.New("plannedDate is required and must be a non-zero YYYY-MM-DD date")
	}

	var allocations = mapToPlannedAllocations(allocationPlanDTS.Details)

	var plannedExecutionDate *time.Time
	if planType == allocation.BalancingExecutionPlan {
		plannedExecutionDate = allocationPlanDTS.PlannedExecutionDate
	}

	var allocationPlanId int64
	if allocationPlanDTS.Id != nil {
		allocationPlanId = int64(*allocationPlanDTS.Id)
	}

	return &domain.AllocationPlan{
		AllocationPlanIdentifier: domain.AllocationPlanIdentifier{
			Id:          allocationPlanId,
			Name:        allocationPlanDTS.Name,
			PlannedDate: plannedDate,
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

	var asset = mapToAssetFromAllocationPlan(allocationDTS.Asset)

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

// mapToAllocationPlanAssetDTS maps a domain Asset to an AllocationPlanAssetDTS.
//
// Authored by: GitHub Copilot
func mapToAllocationPlanAssetDTS(asset *domain.Asset) *AllocationPlanAssetDTS {

	if asset == nil {
		return nil
	}

	return &AllocationPlanAssetDTS{
		Id:     new(langext.ParseableInt64(asset.Id)),
		Name:   asset.Name,
		Ticker: asset.Ticker,
	}
}

// mapToAssetFromAllocationPlan maps an AllocationPlanAssetDTS to a domain Asset.
//
// Authored by: GitHub Copilot
func mapToAssetFromAllocationPlan(assetDTS *AllocationPlanAssetDTS) *domain.Asset {

	if assetDTS == nil {
		return nil
	}

	var assetId int64
	if assetDTS.Id != nil {
		assetId = int64(*assetDTS.Id)
	}
	return &domain.Asset{
		Id:     assetId,
		Name:   assetDTS.Name,
		Ticker: assetDTS.Ticker,
	}
}
