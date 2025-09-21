package domain

import (
	"context"
	"time"

	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/shopspring/decimal"
)

type PlannedAllocationsPerHierarchicalId map[string]*PlannedAllocation

func (plannedAllocationMap PlannedAllocationsPerHierarchicalId) Get(hierarchicalId string) *PlannedAllocation {
	return plannedAllocationMap[hierarchicalId]
}

func (plannedAllocationMap PlannedAllocationsPerHierarchicalId) Remove(hierarchicalId string) {
	delete(plannedAllocationMap, hierarchicalId)
}

type PlannedAllocation struct {
	Id                  int64
	HierarchicalId      HierarchicalId
	CashReserve         bool
	SliceSizePercentage decimal.Decimal
	Asset               *Asset
}

type AllocationPlanIdentifier struct {
	Id   int64
	Name string
}

type AllocationPlan struct {
	AllocationPlanIdentifier
	PlanType             allocation.PlanType
	PlannedExecutionDate *time.Time
	PortfolioId          int
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
