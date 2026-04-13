package domain

import (
	"database/sql/driver"
	"strings"

	"github.com/benizzio/open-asset-allocator/infra/rdbms/sqlext"
)

type HierarchicalId []*string

// String returns the hierarchical identifier as a single string using
// HierarchicalIdLevelSeparator between non-nil levels.
func (hierarchicalId HierarchicalId) String() string {
	var result strings.Builder
	for index, level := range hierarchicalId {
		if level != nil {
			result.WriteString(*level)
			if index < len(hierarchicalId)-1 {
				result.WriteString(HierarchicalIdLevelSeparator)
			}
		}
	}
	return result.String()
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

	var levelIndex = hierarchicalId.GetLevelIndex()
	return hierarchicalId[levelIndex+1:]
}

func (hierarchicalId HierarchicalId) GetParentLevelIndex() int {

	var levelIndex = hierarchicalId.GetLevelIndex()

	if levelIndex == -1 {
		return -2
	}

	if levelIndex == len(hierarchicalId)-1 {
		return -1
	}

	return levelIndex + 1
}
