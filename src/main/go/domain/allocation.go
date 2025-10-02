package domain

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/benizzio/open-asset-allocator/infra/sqlext"
)

const HierarchicalIdLevelSeparator = "|"

type AllocationHierarchyLevel struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationHierarchy []AllocationHierarchyLevel

func (hierarchy AllocationHierarchy) Size() int {
	return len(hierarchy)
}

type AllocationStructure struct {
	Hierarchy AllocationHierarchy `json:"hierarchy,omitempty"`
}

func (allocationStructure *AllocationStructure) Scan(value interface{}) error {

	if value == nil {
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("scanned value is incompatible with AllocationStructure (not a []byte): %v", value)
	}

	return json.Unmarshal(bytes, allocationStructure)
}

func (allocationStructure AllocationStructure) Value() (driver.Value, error) {

	bytes, err := json.Marshal(allocationStructure)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal AllocationStructure: %v", err)
	}

	return bytes, nil
}

type HierarchicalId []*string

// String returns the hierarchical identifier as a single string using
// HierarchicalIdLevelSeparator between non-nil levels.
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

// Value implements driver.Valuer so HierarchicalId can be used directly as a
// SQL parameter with database/sql and github.com/lib/pq. It encodes the
// hierarchical levels as a PostgreSQL text[] array, preserving NULLs for any
// nil entries.
//
// Co-authored by: GitHub Copilot
func (hierarchicalId HierarchicalId) Value() (driver.Value, error) {
	return sqlext.BuildNullStringSlice(hierarchicalId).Value()
}

func (hierarchicalId HierarchicalId) IsTopLevel() bool {

	var length = len(hierarchicalId)
	var lastIndex = length - 1

	if lastIndex == 0 {
		return true
	}

	return hierarchicalId[lastIndex] != nil && hierarchicalId[lastIndex-1] == nil
}

func (hierarchicalId HierarchicalId) GetLevelIndex() int {
	for index := range hierarchicalId {
		if hierarchicalId[index] != nil {
			return index
		}
	}
	return -1
}

func (hierarchicalId HierarchicalId) ParentLevelId() HierarchicalId {

	if hierarchicalId.IsTopLevel() {
		return nil
	}

	levelIndex := hierarchicalId.GetLevelIndex()
	return hierarchicalId[levelIndex+1:]
}
