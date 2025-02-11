import { AllocationHierarchyLevel, AllocationStructure, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX } from "./allocation";
import {
    getAllocationHierarchySize,
    getHierarchyLevelIndex,
    getPlannedAllocationStructuralIdAsString,
    getStructuralIdAsString,
    getTopLevelHierarchyIndexFromAllocationStructure,
} from "./allocation-utils";
import BigNumber from "bignumber.js";
import {
    AllocationPlan,
    AllocationPlanDTO,
    FractalPlannedAllocation,
    FractalPlannedAllocationHierarchy,
    PlannedAllocation,
} from "./allocation-plan";

export function mapAllocationPlanFractalHierarchy(
    allocationPlan: AllocationPlan,
    allocationStructure: AllocationStructure,
): FractalPlannedAllocationHierarchy {

    const allocationsPerHierarchyLevel = mapAllocationsPerHierarchyLevel(allocationPlan);

    const topHierarchyLevelIndex = getTopLevelHierarchyIndexFromAllocationStructure(allocationStructure);
    const aggregatorAllocationMap = new Map<string, FractalPlannedAllocation>();

    const fractalAllocationsPerHierarchyLevel = mapFractalAllocationsPerHierarchyLevel(
        allocationsPerHierarchyLevel,
        allocationStructure,
        aggregatorAllocationMap,
    );

    connectFractalStructure(fractalAllocationsPerHierarchyLevel, aggregatorAllocationMap);

    return {
        subLevel: allocationStructure.hierarchy[topHierarchyLevelIndex],
        topAllocations: fractalAllocationsPerHierarchyLevel[topHierarchyLevelIndex],
        aggregatorAllocationMap: aggregatorAllocationMap,
    };
}

function mapAllocationsPerHierarchyLevel(allocationPlan: AllocationPlan) {

    const allocationsPerHierarchyLevel: PlannedAllocation[][] = [];

    allocationPlan.details.forEach((allocation) => {

        const hierarchyLevelIndex = getHierarchyLevelIndex(allocation);

        if(!allocationsPerHierarchyLevel[hierarchyLevelIndex]) {
            allocationsPerHierarchyLevel[hierarchyLevelIndex] = [];
        }
        allocationsPerHierarchyLevel[hierarchyLevelIndex].push(allocation);

    });

    return allocationsPerHierarchyLevel;
}

function mapFractalAllocationsPerHierarchyLevel(
    allocationsPerHierarchyLevel: PlannedAllocation[][],
    allocationStructure: AllocationStructure,
    aggregatorAllocationMap: Map<string, FractalPlannedAllocation>,
) {

    const fractalAllocationsPerHierarchyLevel: FractalPlannedAllocation[][] = [];
    const hierachySize = getAllocationHierarchySize(allocationStructure);

    for(
        let hierarchyLevelIndex = hierachySize - 1;
        hierarchyLevelIndex >= LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX;
        hierarchyLevelIndex--
    ) {

        fractalAllocationsPerHierarchyLevel[hierarchyLevelIndex] = [];
        const allocationsAtCurrentLevel = allocationsPerHierarchyLevel[hierarchyLevelIndex];

        mapFractalAllocationsAtHierarchyLevel(
            allocationStructure.hierarchy,
            hierarchyLevelIndex,
            allocationsAtCurrentLevel,
            fractalAllocationsPerHierarchyLevel,
            aggregatorAllocationMap,
        );
    }
    return fractalAllocationsPerHierarchyLevel;
}

function mapFractalAllocationsAtHierarchyLevel(
    hierarchy: AllocationHierarchyLevel[],
    hierarchyLevelIndex: number,
    allocationsAtCurrentLevel: PlannedAllocation[],
    fractalAllocationsPerHierarchyLevel: FractalPlannedAllocation[][],
    aggregatorAllocationMap: Map<string, FractalPlannedAllocation>,
) {

    allocationsAtCurrentLevel.forEach((allocation) => {

        const fractalAllocationKey = getStructuralIdAsString(allocation);

        const hierarchySublevel = hierarchyLevelIndex - 1 >= LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX
            ? hierarchy[hierarchyLevelIndex - 1]
            : null;

        const fractalAggregationAllocation = {
            key: fractalAllocationKey,
            level: hierarchy[hierarchyLevelIndex],
            subLevel: hierarchySublevel,
            allocation: allocation,
        };

        fractalAllocationsPerHierarchyLevel[hierarchyLevelIndex].push(fractalAggregationAllocation);

        if(hierarchyLevelIndex > LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX) {
            aggregatorAllocationMap.set(fractalAllocationKey, fractalAggregationAllocation);
        }
    });
}

function connectFractalStructure(
    fractalAllocationsPerHierarchyLevel: FractalPlannedAllocation[][],
    aggregatorAllocationMap: Map<string, FractalPlannedAllocation>,
) {

    const hierachySize = fractalAllocationsPerHierarchyLevel.length;

    for(
        let hierarchyLevelIndex = hierachySize - 2; // Top hierarchy level has no upper level
        hierarchyLevelIndex >= LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX;
        hierarchyLevelIndex--
    ) {

        const fractalAllocationsAtCurrentLevel = fractalAllocationsPerHierarchyLevel[hierarchyLevelIndex];

        connectAllocationsToFractalStructure(
            hierarchyLevelIndex,
            fractalAllocationsAtCurrentLevel,
            aggregatorAllocationMap,
        );
    }
}

function connectAllocationsToFractalStructure(
    hierarchyLevelIndex: number,
    fractalAllocationsAtLevel: FractalPlannedAllocation[],
    aggregatorAllocationMap: Map<string, FractalPlannedAllocation>,
) {

    fractalAllocationsAtLevel.forEach((fractalAllocation) => {

        const allocation = fractalAllocation.allocation;

        const parentStructuralId = allocation.structuralId.slice(hierarchyLevelIndex + 1);
        const parentKey = getPlannedAllocationStructuralIdAsString(parentStructuralId);
        const parent = aggregatorAllocationMap.get(parentKey);

        if(!parent) {
            throw new Error(`Parent not found for allocation ${ allocation.structuralId }`);
        }

        parent.subAllocations = parent.subAllocations || [];
        parent.subAllocations.push(fractalAllocation);
        fractalAllocation.superAllocation = parent;
    });
}

export function mapToAllocationPlan(
    allocationPlanDTO: AllocationPlanDTO,
): AllocationPlan {

    const allocations = allocationPlanDTO.details.map((allocation) => {
        return {
            ...allocation,
            sliceSizePercentage: new BigNumber(allocation.sliceSizePercentage),
        } as PlannedAllocation;
    });

    return {
        ...allocationPlanDTO,
        details: allocations,
    };
}