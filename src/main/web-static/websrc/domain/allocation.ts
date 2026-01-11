export enum AllocationPlanType {
    ASSET_ALLOCATION_PLAN = "ALLOCATION_PLAN",
    BALANCING_EXECUTION_PLAN = "EXECUTION_PLAN",
}

export const LOWEST_AVAILABLE_HIERARCHY_LEVEL = { name: "Assets", field: "assetTicker", index: 0 };
export const LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX = 0;

export type AllocationHierarchyLevel = {
    name: string;
    field: string;
    index: number;
};

export type AllocationStructure = { hierarchy: AllocationHierarchyLevel[]; };

export type AllocationHierarchyLevelDTO = Omit<AllocationHierarchyLevel, "index">;

export type AllocationStructureDTO = { hierarchy: AllocationHierarchyLevelDTO[]; };
