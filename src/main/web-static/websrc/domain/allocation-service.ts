import BigNumber from "bignumber.js";
import {
    Allocation,
    AllocationPlan,
    AllocationPlanDTO,
    AllocationPlanHierarchyLevel,
    FractalAllocation,
    FractalAllocationHierarchy,
} from "./allocation";

function getHierarchyLevel(allocation: Allocation): number {

    const hierarchyTopLevelIndex = allocation.structuralId.length - 1;

    for (let i = hierarchyTopLevelIndex - 1; i >= 0; i--) {
        if (!allocation.structuralId[i]) {
            return i + 1;
        }
    }

    return 0;
}

function mapAllocationsPerHierarchyLevel(allocationPlan: AllocationPlan) {

    const allocationsPerHierarchyLevel: Allocation[][] = [];

    allocationPlan.details.forEach((allocation) => {

        const hierarchyLevelIndex = getHierarchyLevel(allocation);

        if (!allocationsPerHierarchyLevel[hierarchyLevelIndex]) {
            allocationsPerHierarchyLevel[hierarchyLevelIndex] = [];
        }
        allocationsPerHierarchyLevel[hierarchyLevelIndex].push(allocation);

    });

    return allocationsPerHierarchyLevel;
}

function mapFractalAllocationsAtHierarchyLevel(
    hierarchyLevel: AllocationPlanHierarchyLevel,
    hierarchyLevelIndex: number,
    allocationsAtCurrentLevel: Allocation[],
    fractalAllocationsPerHierarchyLevel: FractalAllocation[][],
    aggregatorAllocationMap: Map<string, FractalAllocation>,
) {

    allocationsAtCurrentLevel.forEach((allocation) => {

        const fractalAllocationKey = allocation.structuralId[hierarchyLevelIndex];

        const fractalAggregationAllocation = {
            level: hierarchyLevel,
            key: fractalAllocationKey,
            allocation: allocation,
        };

        fractalAllocationsPerHierarchyLevel[hierarchyLevelIndex].push(fractalAggregationAllocation);

        if (hierarchyLevelIndex > 0) {
            aggregatorAllocationMap.set(fractalAllocationKey, fractalAggregationAllocation);
        }
    });
}

function mapFractalAllocationsPerHierarchyLevel(
    allocationsPerHierarchyLevel: Allocation[][],
    allocationPlan: AllocationPlan,
    aggregatorAllocationMap: Map<string, FractalAllocation>,
) {

    const fractalAllocationsPerHierarchyLevel: FractalAllocation[][] = [];
    const hierachySize = allocationPlan.structure.hierarchy.length;

    for (let hierarchyLevelIndex = hierachySize - 1; hierarchyLevelIndex >= 0; hierarchyLevelIndex--) {

        fractalAllocationsPerHierarchyLevel[hierarchyLevelIndex] = [];
        const allocationsAtCurrentLevel = allocationsPerHierarchyLevel[hierarchyLevelIndex];
        const hierarchyLevel = allocationPlan.structure.hierarchy[hierarchyLevelIndex];

        mapFractalAllocationsAtHierarchyLevel(
            hierarchyLevel,
            hierarchyLevelIndex,
            allocationsAtCurrentLevel,
            fractalAllocationsPerHierarchyLevel,
            aggregatorAllocationMap,
        );
    }
    return fractalAllocationsPerHierarchyLevel;
}

function connectAllocationsToFractalStructure(
    hierarchyLevelIndex: number,
    fractalAllocationsAtLevel: FractalAllocation[],
    aggregatorAllocationMap: Map<string, FractalAllocation>,
) {

    fractalAllocationsAtLevel.forEach((fractalAllocation) => {

        const allocation = fractalAllocation.allocation;
        const parentKey = allocation.structuralId[hierarchyLevelIndex + 1];
        const parent = aggregatorAllocationMap.get(parentKey);

        if (!parent) {
            throw new Error(`Parent not found for allocation ${ allocation.structuralId }`);
        }

        parent.subAllocations = parent.subAllocations || [];
        parent.subAllocations.push(fractalAllocation);
    });
}

function connectFractalStructure(
    fractalAllocationsPerHierarchyLevel: FractalAllocation[][],
    aggregatorAllocationMap: Map<string, FractalAllocation>,
) {

    const hierachySize = fractalAllocationsPerHierarchyLevel.length;

    for (let hierarchyLevelIndex = hierachySize - 2; hierarchyLevelIndex >= 0; hierarchyLevelIndex--) {

        const fractalAllocationsAtCurrentLevel = fractalAllocationsPerHierarchyLevel[hierarchyLevelIndex];

        connectAllocationsToFractalStructure(
            hierarchyLevelIndex,
            fractalAllocationsAtCurrentLevel,
            aggregatorAllocationMap,
        );
    }
}

export const allocationDomainService = {

    getTopLevelKey(allocationPlan: AllocationPlan): string {
        const allocationPlanHierarchy = allocationPlan.structure.hierarchy;
        const topLevelIndex = allocationPlanHierarchy.length - 1;
        return allocationPlanHierarchy[topLevelIndex].name;
    },

    mapToAllocationPlan(allocationPlanDTO: AllocationPlanDTO): AllocationPlan {

        const allocations = allocationPlanDTO.details.map((allocation) => {
            return {
                ...allocation,
                sliceSizePercentage: new BigNumber(allocation.sliceSizePercentage),
            } as Allocation;
        });

        return {
            ...allocationPlanDTO,
            details: allocations,
        };
    },

    mapAllocationPlanFractalHierarchy(allocationPlan: AllocationPlan): FractalAllocationHierarchy {

        const allocationsPerHierarchyLevel = mapAllocationsPerHierarchyLevel(allocationPlan);
        const hierachySize = allocationPlan.structure.hierarchy.length;
        const topHierarchyLevelIndex = hierachySize - 1;
        const aggregatorAllocationMap = new Map<string, FractalAllocation>();

        const fractalAllocationsPerHierarchyLevel = mapFractalAllocationsPerHierarchyLevel(
            allocationsPerHierarchyLevel,
            allocationPlan,
            aggregatorAllocationMap,
        );

        connectFractalStructure(fractalAllocationsPerHierarchyLevel, aggregatorAllocationMap);

        return {
            topAllocations: fractalAllocationsPerHierarchyLevel[topHierarchyLevelIndex],
            aggregatorAllocationMap: aggregatorAllocationMap,
        };
    },
};