import BigNumber from "bignumber.js";

export enum AllocationType {
    ASSET_ALLOCATION_PLAN = "ALLOCATION_PLAN",
    BALANCING_EXECUTION_PLAN = "EXECUTION_PLAN",
}

export type AllocationPlanHierarchyLevel = {
    name: string;
    field: string;
};

export type AllocationPlanStructure = { hierarchy: AllocationPlanHierarchyLevel[]; };

export type Allocation = {
    structuralId: string[],
    chashReserve: boolean,
    sliceSizePercentage: BigNumber,
};

export type AllocationPlan = {
    id: number;
    name: string;
    type: AllocationType;
    structure: AllocationPlanStructure,
    plannedExecutionDate?: Date,
    details: Allocation[],
};

export type AllocationDTO = Omit<Allocation, "sliceSizePercentage"> & { sliceSizePercentage: string, };
export type AllocationPlanDTO = Omit<AllocationPlan, "details"> & { details: AllocationDTO[], };

export type FractalAllocation = {
    level: AllocationPlanHierarchyLevel;
    key: string;
    allocation: Allocation;
    subAllocations?: FractalAllocation[];
};

export type FractalAllocationHierarchy = {
    topAllocations: FractalAllocation[];
    aggregatorAllocationMap: Map<string, FractalAllocation>;
};