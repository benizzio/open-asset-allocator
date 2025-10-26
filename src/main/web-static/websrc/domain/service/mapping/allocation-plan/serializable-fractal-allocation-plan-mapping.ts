import {
    FractalHierarchicalAllocationPlan,
    FractalPlannedAllocation,
    SerializableFractalHierarchicalAllocationPlan,
    SerializableFractalPlannedAllocation,
} from "../../../allocation-plan";

export function mapToSerializableFractalHierarchicalAllocationPlan(
    fractalHierarchicalPlan: FractalHierarchicalAllocationPlan,
): SerializableFractalHierarchicalAllocationPlan {

    function mapToSerializableFractalPlannedAllocation(
        fractalPlannedAllocation: FractalPlannedAllocation,
    ): SerializableFractalPlannedAllocation {

        return {
            key: fractalPlannedAllocation.key,
            level: fractalPlannedAllocation.level,
            subLevel: fractalPlannedAllocation.subLevel,
            allocation: fractalPlannedAllocation.allocation,
            subAllocations: fractalPlannedAllocation.subAllocations
                ? fractalPlannedAllocation.subAllocations.map(mapToSerializableFractalPlannedAllocation)
                : undefined,
        };
    }

    return {
        subLevel: fractalHierarchicalPlan.subLevel,
        topAllocations: fractalHierarchicalPlan.topAllocations.map(mapToSerializableFractalPlannedAllocation),
    };
}