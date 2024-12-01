import BigNumber from "bignumber.js";
import { AllocationHierarchyLevel, AllocationStructure, AllocationType } from "./allocation";

export type PlannedAllocation = {
    structuralId: string[],
    chashReserve: boolean,
    sliceSizePercentage: BigNumber,
};

export type AllocationPlan = {
    id: number;
    name: string;
    type: AllocationType;
    structure: AllocationStructure,
    plannedExecutionDate?: Date,
    details: PlannedAllocation[],
};

export type PlannedAllocationDTO = Omit<PlannedAllocation, "sliceSizePercentage"> & { sliceSizePercentage: string, };

export type AllocationPlanDTO = Omit<AllocationPlan, "details" | "structure"> & {
    details: PlannedAllocationDTO[],
    structure: AllocationPlanStructureDTO,
};

export type AllocationPlanHierarchyLevelDTO = Omit<AllocationHierarchyLevel, "index">;

export type AllocationPlanStructureDTO = { hierarchy: AllocationPlanHierarchyLevelDTO[]; };

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