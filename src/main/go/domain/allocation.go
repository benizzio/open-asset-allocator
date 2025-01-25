package domain

import (
	"encoding/json"
	"fmt"
)

const HierarchicalIdLevelSeparator = "|"

type AllocationHierarchyLevel struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationHierarchy []AllocationHierarchyLevel

// TODO add as generic slice method
func (hierarchy AllocationHierarchy) Size() int {
	return len(hierarchy)
}

type AllocationStructure struct {
	Hierarchy AllocationHierarchy `json:"hierarchy,omitempty"`
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

type HierarchicalId []*string

func (hierarchicalId HierarchicalId) String() string {
	var result = ""
	for index, level := range hierarchicalId {
		if level != nil {
			result += *level
			if index < len(hierarchicalId)-1 {
				result += HierarchicalIdLevelSeparator
			}
		}
	}
	return result
}
