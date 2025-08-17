import { AllocationStructure } from "../allocation";
import { PlannedAllocation } from "../allocation-plan";

export function getHierarchyLevelIndex(allocation: PlannedAllocation): number {

    const hierarchyTopLevelIndex = getTopLevelHierarchyIndexFromPlannedAllocation(allocation);

    for(let i = hierarchyTopLevelIndex - 1; i >= 0; i--) {
        if(!allocation.structuralId[i]) {
            return i + 1;
        }
    }

    return 0;
}

export function getStructuralIdAsString(plannedAllocation: PlannedAllocation): string {
    return getPlannedAllocationStructuralIdAsString(plannedAllocation.structuralId);
}

export function getPlannedAllocationStructuralIdAsString(structuralId: string[]): string {
    return structuralId.filter(value => value != null).join("|");
}

export function getAllocationHierarchySize(allocationStructure: AllocationStructure): number {
    return allocationStructure.hierarchy.length;
}

export function getTopLevelHierarchyIndexFromAllocationStructure(allocationStructure: AllocationStructure) {
    return allocationStructure.hierarchy.length - 1;
}

export function getTopLevelHierarchyIndexFromPlannedAllocation(plannedAllocation: PlannedAllocation) {
    return plannedAllocation.structuralId.length - 1;
}