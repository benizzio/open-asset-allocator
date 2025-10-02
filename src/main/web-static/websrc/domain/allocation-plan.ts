import BigNumber from "bignumber.js";
import { AllocationHierarchyLevel, AllocationType } from "./allocation";

export type PlannedAllocation = {
    hierarchicalId: string[];
    cashReserve: boolean;
    sliceSizePercentage: BigNumber;
};

export type AllocationPlan = {
    id: number;
    name: string;
    type: AllocationType;
    plannedExecutionDate?: Date,
    details: PlannedAllocation[],
};

/**
 * Type that defines a Planned Allocation as it's tranferred via API,
 * to allow conversion to property types needed locally
 */
export type PlannedAllocationDTO = Omit<PlannedAllocation, "sliceSizePercentage"> & { sliceSizePercentage: string, };

/**
 * Type that defines a Allocation Plan as it's tranferred via API,
 * to allow conversion to property types needed locally
 */
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