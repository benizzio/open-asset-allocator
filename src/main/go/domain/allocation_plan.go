package domain

import (
	"context"
	"time"

	"github.com/shopspring/decimal"

	"github.com/benizzio/open-asset-allocator/domain/allocation"
)

type PlannedAllocationsPerHierarchicalId map[string]*PlannedAllocation

func (plannedAllocationMap PlannedAllocationsPerHierarchicalId) Get(hierarchicalId string) *PlannedAllocation {
	return plannedAllocationMap[hierarchicalId]
}

func (plannedAllocationMap PlannedAllocationsPerHierarchicalId) Remove(hierarchicalId string) {
	delete(plannedAllocationMap, hierarchicalId)
}

type PlannedAllocation struct {
	Id             int64
	HierarchicalId HierarchicalId
	CashReserve    bool
	//TODO remove the percentage from this name in all the stack
	SliceSizePercentage decimal.Decimal
	Asset               *Asset
}

// AllocationPlanIdentifier identifies an allocation plan and carries its required calendar date for list reads.
// For example, a repository identifier can be mapped to an API object with id, name, and plannedDate.
//
// Co-authored by: OpenCode and benizzio
type AllocationPlanIdentifier struct {
	Id          int64
	Name        string
	PlannedDate time.Time
}

type AllocationPlan struct {
	AllocationPlanIdentifier
	PlanType             allocation.PlanType
	PlannedExecutionDate *time.Time
	PortfolioId          int64
	Details              []*PlannedAllocation
}

func (allocationPlan *AllocationPlan) AddDetail(detail *PlannedAllocation) {
	allocationPlan.Details = append(allocationPlan.Details, detail)
}

type AllocationPlanRepository interface {
	GetAllAllocationPlans(portfolioId int64, planType *allocation.PlanType) ([]*AllocationPlan, error)
	GetAllocationPlan(id int64) (*AllocationPlan, error)
	GetAllAllocationPlanIdentifiers(
		portfolioId int64,
		planType *allocation.PlanType,
	) ([]*AllocationPlanIdentifier, error)
	InsertAllocationPlanInTransaction(transContext context.Context, plan *AllocationPlan) error
	UpdateAllocationPlanInTransaction(transContext context.Context, plan *AllocationPlan) error
}
