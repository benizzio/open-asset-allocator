package domain

import (
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/shopspring/decimal"
	"time"
)

type PlannedAllocation struct {
	StructuralId        []*string
	CashReserve         bool
	SliceSizePercentage decimal.Decimal
}

type AllocationPlan struct {
	Id                   int
	Name                 string
	PlanType             allocation.PlanType
	PlannedExecutionDate *time.Time
	Details              []*PlannedAllocation
}

func (allocationPlan *AllocationPlan) AddDetail(detail *PlannedAllocation) {
	allocationPlan.Details = append(allocationPlan.Details, detail)
}

type AllocationPlanRepository interface {
	GetAllAllocationPlans(portfolioId int, planType *allocation.PlanType) ([]*AllocationPlan, error)
}
