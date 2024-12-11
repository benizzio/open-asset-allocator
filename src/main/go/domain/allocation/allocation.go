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

func (planTypeId *PlanType) String() string {
	return planTypeNames[*planTypeId]
}

func GetPlanType(name string) (PlanType, error) {
	for key, value := range planTypeNames {
		if value == name {
			return key, nil
		}
	}
	return -1, errors.New("invalid plan type '" + name + "'")
}

func (planTypeId *PlanType) Scan(value interface{}) error {

	if value == nil {
		*planTypeId = -1
		return nil
	}

	name, ok := value.(string)
	if !ok {
		return errors.New("scanned value is incompatible with PlanType (not a string)")
	}

	planType, err := GetPlanType(name)
	if err != nil {
		return err
	}

	*planTypeId = planType
	return nil
}
