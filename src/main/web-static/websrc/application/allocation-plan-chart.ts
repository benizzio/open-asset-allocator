import {
    Allocation,
    AllocationPlan,
    AllocationPlanDTO,
    AllocationPlanFractalHierarchy,
    AllocationPlanHierarchyLevel,
    AllocationPlanHierarchyNode,
} from "../domain/allocation";
import BigNumber from "bignumber.js";
import { ChartContent } from "../infra/chart/chart-types";

const allocationPlanChart = {
    toUnidimensionalChartContent(allocationPlanDTO: AllocationPlanDTO): ChartContent {

        const allocations = allocationPlanDTO.details.map((allocation) => {
            return {
                ...allocation,
                sliceSizePercentage: new BigNumber(allocation.sliceSizePercentage),
            } as Allocation;
        });

        const allocationPlan = {
            ...allocationPlanDTO,
            details: allocations,
        };

        const dataSet = { data: [], label: allocationPlan.name };
        const chartData = { labels: [], datasets: [dataSet], test: "test" };

        allocationPlan.details.filter(allocation => allocation.structuralId[0] == null).forEach((allocation) => {
            chartData.labels.push(allocation.structuralId[1]);
            dataSet.data.push(allocation.sliceSizePercentage.toNumber());
        });

        const fractalHierarchy = extractAllocationPlanFractalHierarchy(allocationPlan);
        console.log(fractalHierarchy);

        //TODO continue (click on the chart section to view next hierarchy dataset)

        return { chartData };
    },
};

function isFromLevel(allocation: Allocation, hierarchyLevelIndex: number) {
    return allocation.structuralId.every(
        (value, index) => index < hierarchyLevelIndex ? value == null : value != null,
    );
}

function extractAllocationPlanFractalHierarchy(allocationPlan: AllocationPlan): AllocationPlanFractalHierarchy {
    
    const allocationsPerHierarchyLevel = getAllocationsPerHierarchyLevel(allocationPlan);

    const rootNodes =
        mapAllocationPlanHierarchyNodes(
            allocationPlan.structure.hierarchy,
            allocationsPerHierarchyLevel,
            allocationPlan.structure.hierarchy.length - 1,
        );

    return { nodes: rootNodes };
}

function getAllocationsPerHierarchyLevel(allocationPlan: AllocationPlan) {
    //TODO map per parent structuralId
    const allocationsPerHierarchyLevel: Allocation[][] = [];

    allocationPlan.structure.hierarchy.forEach((hierarchyLevel, index) => {
        allocationsPerHierarchyLevel.push(
            allocationPlan.details.filter(allocation => isFromLevel(allocation, index)),
        );
    });
    return allocationsPerHierarchyLevel;
}

function mapAllocationPlanHierarchyNodes(
    hierarchy: AllocationPlanHierarchyLevel[],
    allocationsPerHierarchyLevel: Allocation[][],
    hierarchyLevelIndex: number,
    parentStructuralId?: string,
): AllocationPlanHierarchyNode[] {

    let allocationsAtHierarchyLevel = allocationsPerHierarchyLevel[hierarchyLevelIndex];

    if(parentStructuralId) {
        allocationsAtHierarchyLevel = allocationsAtHierarchyLevel.filter(
            allocation => allocation.structuralId[hierarchyLevelIndex + 1] == parentStructuralId);
    }

    return allocationsAtHierarchyLevel.map((allocation) => {

        const subNodes =
            hierarchyLevelIndex > 0 ?
                mapAllocationPlanHierarchyNodes(
                    hierarchy,
                    allocationsPerHierarchyLevel,
                    hierarchyLevelIndex - 1,
                    allocation.structuralId[hierarchyLevelIndex],
                )
                : undefined;

        return {
            level: hierarchy[hierarchyLevelIndex],
            value: allocation.structuralId[hierarchyLevelIndex],
            subNodes,
        } as AllocationPlanHierarchyNode;
    });
}

export default allocationPlanChart;

