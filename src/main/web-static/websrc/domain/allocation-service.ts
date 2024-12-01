import { AllocationHierarchyLevel, AllocationStructure, LOWEST_AVAILABLE_HIERARCHY_LEVEL } from "./allocation";
import { mapAllocationPlanFractalHierarchy, mapToAllocationPlan } from "./allocation-plan-mapping";
import { AllocationPlan } from "./allocation-plan";
import { getTopLevelHierarchyIndexFromAllocationStructure } from "./allocation-utils";


export const allocationDomainService = {

    getTopLevelHierarchyKeyFromAllocationPlan(allocationPlan: AllocationPlan): string {
        const hierarchy = allocationPlan.structure.hierarchy;
        const topLevelIndex = getTopLevelHierarchyIndexFromAllocationStructure(allocationPlan.structure);
        return hierarchy[topLevelIndex].name;
    },

    getTopLevelHierarchyIndexFromAllocationStructure,

    getLowerHierarchyLevelFromAllocationStructure(
        allocationPlanStructure: AllocationStructure,
        currentLevelIndex: number,
    ): AllocationHierarchyLevel {
        return allocationPlanStructure.hierarchy[currentLevelIndex - 1];
    },

    getHierarchyLevelFromAllocationStructure(
        allocationStructure: AllocationStructure,
        levelIndex: number,
    ): AllocationHierarchyLevel {
        return allocationStructure.hierarchy[levelIndex];
    },

    getTopHierarchyLevelInfoFromAllocationStructure(allocationStructure: AllocationStructure): {
        topHierarchyLevel: AllocationHierarchyLevel,
        topLevelIndex: number
    } {

        const hierarchy = allocationStructure.hierarchy;

        let topHierarchyLevel = LOWEST_AVAILABLE_HIERARCHY_LEVEL;
        let topLevelIndex = 0;

        if (allocationStructure) {
            topLevelIndex = getTopLevelHierarchyIndexFromAllocationStructure(allocationStructure);
            topHierarchyLevel = hierarchy[topLevelIndex];
        }

        return { topHierarchyLevel, topLevelIndex };
    },

    mapToAllocationPlan,
    mapAllocationPlanFractalHierarchy,
};