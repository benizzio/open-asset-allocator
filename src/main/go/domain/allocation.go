package domain

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/lib/pq"
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
// Usage:
//
//	// given: var id domain.HierarchicalId
//	_, err := db.Exec("INSERT INTO table(col) VALUES($1)", id)
//	// pq will receive a proper text[] representation
//
// Notes:
//   - Each non-nil level becomes a valid array element.
//   - Nil levels are encoded as SQL NULL within the array.
//   - Escaping and formatting is delegated to pq.Array to ensure correctness.
//
// Authored by: GitHub Copilot
func (hierarchicalId HierarchicalId) Value() (driver.Value, error) {
	// Map []*string -> []sql.NullString to preserve NULLs.
	var arr = make([]sql.NullString, len(hierarchicalId))
	for i, ptr := range hierarchicalId {
		if ptr == nil {
			arr[i] = sql.NullString{String: "", Valid: false}
		} else {
			arr[i] = sql.NullString{String: *ptr, Valid: true}
		}
	}

	// Delegate array formatting to pq to ensure proper escaping/quoting.
	return pq.Array(arr).Value()
}
