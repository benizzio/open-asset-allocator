package model

import "github.com/benizzio/open-asset-allocator/domain"

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

func mapToAllocationStructureDTS(allocationPlanStructure domain.AllocationStructure) AllocationStructureDTS {
	var hierarchyLevels = mapToAllocationHierarchyLevelDTSs(allocationPlanStructure.Hierarchy)
	return AllocationStructureDTS{Hierarchy: hierarchyLevels}
}

func mapToAllocationHierarchyLevelDTSs(levels []domain.AllocationHierarchyLevel) []AllocationHierarchyLevelDTS {

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

func mapToAllocationStructure(allocationPlanStructureDTS *AllocationStructureDTS) domain.AllocationStructure {
	var hierarchyLevels = mapToAllocationHierarchyLevels(allocationPlanStructureDTS.Hierarchy)
	return domain.AllocationStructure{Hierarchy: hierarchyLevels}
}

func mapToAllocationHierarchyLevels(levels []AllocationHierarchyLevelDTS) []domain.AllocationHierarchyLevel {

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
