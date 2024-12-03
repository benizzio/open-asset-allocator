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

func GetPlanType(name string) (PlanType, error) {
	for key, value := range planTypeNames {
		if value == name {
			return key, nil
		}
	}
	return -1, errors.New("invalid plan type")
}
