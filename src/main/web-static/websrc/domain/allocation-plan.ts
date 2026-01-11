import BigNumber from "bignumber.js";
import { AllocationHierarchyLevel, AllocationPlanType } from "./allocation";
import { Portfolio } from "./portfolio";
import { Asset } from "./asset";

export type PlannedAllocation = {
    hierarchicalId: string[];
    cashReserve: boolean;
    sliceSizePercentage: BigNumber;
    asset?: Asset;
};

export type AllocationPlan = {
    id: number;
    name: string;
    type: AllocationPlanType;
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
    allocation: PlannedAllocation;
    key: string;
    targetLevelKey: string;
    level: AllocationHierarchyLevel;
    subLevel?: AllocationHierarchyLevel;
    subAllocations?: FractalPlannedAllocation[];
    superAllocation?: FractalPlannedAllocation;
};

export type FractalHierarchicalAllocationPlan = {
    subLevel: AllocationHierarchyLevel;
    topAllocations: FractalPlannedAllocation[];
    aggregatorAllocationMap: Map<string, FractalPlannedAllocation>;
};

export type CompleteAllocationPlan = {
    allocationPlan: AllocationPlan;
    fractalHierarchicalPlan: FractalHierarchicalAllocationPlan;
    topLevelKey: string;
};

export type SerializableFractalPlannedAllocation =
    Omit<FractalPlannedAllocation, "subAllocations" | "superAllocation">
    & { subAllocations?: SerializableFractalPlannedAllocation[]; };

export type SerializableFractalHierarchicalAllocationPlan =
    Omit<FractalHierarchicalAllocationPlan, "topAllocations" | "aggregatorAllocationMap">
    & { topAllocations: SerializableFractalPlannedAllocation[]; };

export type SerializableCompleteAllocationPlan =
    Omit<CompleteAllocationPlan, "fractalHierarchicalPlan">
    & { fractalHierarchicalPlan: SerializableFractalHierarchicalAllocationPlan; };

export type SerializablePortfolioCompleteAllocationPlanSet = {
    portfolio: Portfolio;
    completeAllocationPlans: SerializableCompleteAllocationPlan[];
};