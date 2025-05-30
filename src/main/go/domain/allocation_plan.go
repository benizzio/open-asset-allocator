package domain

import (
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/shopspring/decimal"
	"time"
)

type PlannedAllocationsPerHierarchicalId map[string]*PlannedAllocation

func (plannedAllocationMap PlannedAllocationsPerHierarchicalId) Get(hierarchicalId string) *PlannedAllocation {
	return plannedAllocationMap[hierarchicalId]
}

func (plannedAllocationMap PlannedAllocationsPerHierarchicalId) Remove(hierarchicalId string) {
	delete(plannedAllocationMap, hierarchicalId)
}

type PlannedAllocation struct {
	StructuralId        HierarchicalId //TODO rename this to HierarchicalId in all stack
	CashReserve         bool
	SliceSizePercentage decimal.Decimal
}

type AllocationPlanIdentifier struct {
	Id   int
	Name string
}

type AllocationPlan struct {
	AllocationPlanIdentifier
	PlanType             allocation.PlanType
	PlannedExecutionDate *time.Time
	Details              []*PlannedAllocation
}

func (allocationPlan *AllocationPlan) AddDetail(detail *PlannedAllocation) {
	allocationPlan.Details = append(allocationPlan.Details, detail)
}

type AllocationPlanRepository interface {
	GetAllAllocationPlans(portfolioId int, planType *allocation.PlanType) ([]*AllocationPlan, error)
	GetAllocationPlan(id int) (*AllocationPlan, error)
	GetAllAllocationPlanIdentifiers(portfolioId int, planType *allocation.PlanType) ([]*AllocationPlanIdentifier, error)
}
