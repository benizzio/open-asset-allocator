import BigNumber from "bignumber.js";
import { AllocationHierarchyLevel, AllocationType } from "./allocation";

export type PlannedAllocation = {
    structuralId: string[],
    chashReserve: boolean,
    sliceSizePercentage: BigNumber,
};

export type AllocationPlan = {
    id: number;
    name: string;
    type: AllocationType;
    plannedExecutionDate?: Date,
    details: PlannedAllocation[],
};

export type PlannedAllocationDTO = Omit<PlannedAllocation, "sliceSizePercentage"> & { sliceSizePercentage: string, };

export type AllocationPlanDTO = Omit<AllocationPlan, "details"> & { details: PlannedAllocationDTO[], };

export type FractalPlannedAllocation = {
    key: string;
    level: AllocationHierarchyLevel;
    subLevel?: AllocationHierarchyLevel;
    allocation: PlannedAllocation;
    subAllocations?: FractalPlannedAllocation[];
    superAllocation?: FractalPlannedAllocation;
};

export type FractalPlannedAllocationHierarchy = {
    subLevel: AllocationHierarchyLevel;
    topAllocations: FractalPlannedAllocation[];
    aggregatorAllocationMap: Map<string, FractalPlannedAllocation>;
};