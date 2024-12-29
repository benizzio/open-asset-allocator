package domain

import (
	"encoding/json"
	"fmt"
)

type AllocationHierarchyLevel struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationStructure struct {
	Hierarchy []AllocationHierarchyLevel `json:"hierarchy,omitempty"`
}

func (allocationStructure *AllocationStructure) Scan(value interface{}) error {

	if value == nil {
		*allocationStructure = AllocationStructure{}
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scanned value is incompatible with AllocationStructure (not a []byte): %v", value)
	}

	return json.Unmarshal(bytes, allocationStructure)
}
