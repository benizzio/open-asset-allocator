package domain

import (
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/shopspring/decimal"
	"time"
)

type AllocationPlanUnit struct {
	StructuralId        []string
	CashReserve         bool
	SliceSizePercentage decimal.Decimal
}

type AllocationPlanHierarchyLevel struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationPlanStructure struct {
	Hierarchy []AllocationPlanHierarchyLevel `json:"hierarchy,omitempty"`
}

type AllocationPlan struct {
	Name                 string
	PlanType             allocation.PlanType
	Structure            AllocationPlanStructure
	PlannedExecutionDate time.Time
	Details              []AllocationPlanUnit
}
