package allocation

import "errors"

type PlanType int

const (
	AssetAllocationPlan PlanType = iota
	BalancingExecutionPlan
)

var planTypeNames = map[PlanType]string{
	AssetAllocationPlan:    "ALLOCATION_PLAN",
	BalancingExecutionPlan: "EXECUTION_PLAN",
}

func (planTypeId PlanType) String() string {
	return planTypeNames[planTypeId]
}

func (planTypeId PlanType) Get(id string) (PlanType, error) {
	for key, value := range planTypeNames {
		if value == id {
			return key, nil
		}
	}
	return -1, errors.New("invalid plan type")
}
