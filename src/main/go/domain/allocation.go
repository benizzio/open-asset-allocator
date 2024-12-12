package domain

import (
	"encoding/json"
	"fmt"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/shopspring/decimal"
	"time"
)

type AllocationPlanUnit struct {
	StructuralId        []*string
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

func (allocationPlanStructure *AllocationPlanStructure) Scan(value interface{}) error {

	if value == nil {
		*allocationPlanStructure = AllocationPlanStructure{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scanned value is incompatible with AllocationPlanStructure (not a []byte): %v", value)
	}

	return json.Unmarshal(bytes, allocationPlanStructure)
}

type AllocationPlan struct {
	Id                   int
	Name                 string
	PlanType             allocation.PlanType
	Structure            AllocationPlanStructure
	PlannedExecutionDate *time.Time
	Details              []*AllocationPlanUnit
}

func (allocationPlan *AllocationPlan) AddDetail(detail *AllocationPlanUnit) {
	allocationPlan.Details = append(allocationPlan.Details, detail)
}

type AllocationPlanRepository interface {
	GetAllAllocationPlans(planType *allocation.PlanType) ([]*AllocationPlan, error)
}
