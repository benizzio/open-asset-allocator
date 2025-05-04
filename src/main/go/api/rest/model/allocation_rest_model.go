package model

import (
	"github.com/benizzio/open-asset-allocator/domain"
)

// ================================================
// TYPES
// ================================================

type AllocationHierarchyLevelDTS struct {
	Name  string `json:"name,omitempty"`
	Field string `json:"field,omitempty"`
}

type AllocationStructureDTS struct {
	Hierarchy []AllocationHierarchyLevelDTS `json:"hierarchy,omitempty"`
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func mapAllocationStructure(allocationPlanStructure domain.AllocationStructure) AllocationStructureDTS {
	var hierarchyLevels = mapHierarchyLevels(allocationPlanStructure.Hierarchy)
	return AllocationStructureDTS{Hierarchy: hierarchyLevels}
}

func mapHierarchyLevels(levels []domain.AllocationHierarchyLevel) []AllocationHierarchyLevelDTS {

	var dtsLevels = make([]AllocationHierarchyLevelDTS, 0)

	for _, level := range levels {
		var levelDTS = AllocationHierarchyLevelDTS{
			Name:  level.Name,
			Field: level.Field,
		}
		dtsLevels = append(dtsLevels, levelDTS)
	}

	return dtsLevels
}

func mapAllocationStructureDTS(allocationPlanStructureDTS *AllocationStructureDTS) domain.AllocationStructure {
	var hierarchyLevels = mapHierarchyLevelsDTS(allocationPlanStructureDTS.Hierarchy)
	return domain.AllocationStructure{Hierarchy: hierarchyLevels}
}

func mapHierarchyLevelsDTS(levels []AllocationHierarchyLevelDTS) []domain.AllocationHierarchyLevel {

	var domainLevels = make([]domain.AllocationHierarchyLevel, 0)

	for _, level := range levels {
		var levelDomain = domain.AllocationHierarchyLevel{
			Name:  level.Name,
			Field: level.Field,
		}
		domainLevels = append(domainLevels, levelDomain)
	}

	return domainLevels
}
