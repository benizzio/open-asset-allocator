package domain

import (
	"database/sql/driver"

	"github.com/benizzio/open-asset-allocator/infra/rdbms/sqlext"
)

const HierarchicalIdLevelSeparator = "|"

type AllocationHierarchyLevel struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationHierarchy []AllocationHierarchyLevel

func (allocationHierarchy AllocationHierarchy) PrettyString() string {

	var result = ""

	for index, level := range allocationHierarchy {
		result += level.Name
		if index < len(allocationHierarchy)-1 {
			result += " -> "
		}
	}

	return result
}

type AllocationStructure struct {
	Hierarchy AllocationHierarchy `json:"hierarchy,omitempty"`
}

func (allocationStructure *AllocationStructure) Scan(value interface{}) error {
	return sqlext.ScanJsonColumn(value, allocationStructure)
}

func (allocationStructure AllocationStructure) Value() (driver.Value, error) {
	return sqlext.ValueJsonColumn(allocationStructure)
}
