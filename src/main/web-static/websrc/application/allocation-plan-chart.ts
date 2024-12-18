import {
    Allocation,
    AllocationPlan,
    AllocationPlanDTO,
    AllocationPlanHierarchyLevel,
    FractalAllocation,
    FractalAllocationHierarchy,
} from "../domain/allocation";
import BigNumber from "bignumber.js";
import { ChartContent, MultiChartDataSource } from "../infra/chart/chart-types";
import { ChartData } from "chart.js";

//TODO clean code
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

        const fractalHierarchy = extractAllocationPlanFractalHierarchy(allocationPlan);
        console.log(fractalHierarchy);

        const chartDataMap = toChartDataMap(fractalHierarchy);
        console.log(chartDataMap);

        const allocationPlanHierarchy = allocationPlan.structure.hierarchy;
        const topLevelIndex = allocationPlanHierarchy.length - 1;
        const topLevelKey = allocationPlanHierarchy[topLevelIndex].name;

        return { chartDataSource: new MultiChartDataSource(chartDataMap, topLevelKey) };
    },
};

function mapDataSet(
    dataSetLabel: string,
    fractalAllocations: FractalAllocation[],
    chartDataMap: Map<string, ChartData>,
) {

    const dataSet = { data: [], label: dataSetLabel };
    const chartData = { labels: [], datasets: [dataSet] };

    fractalAllocations.forEach((fractalAllocation) => {

        dataSet.data.push(fractalAllocation.allocation.sliceSizePercentage.toNumber());
        chartData.labels.push(fractalAllocation.key);

        const subAllocations = fractalAllocation.subAllocations;

        if (subAllocations && subAllocations.length > 0) {
            mapDataSet(fractalAllocation.key, subAllocations, chartDataMap);
        }
    });
    chartDataMap.set(dataSetLabel, chartData);
}

function toChartDataMap(fractalHierarchy: FractalAllocationHierarchy): Map<string, ChartData> {

    const chartDataMap = new Map<string, ChartData>();
    const fractalAllocations = fractalHierarchy.topAllocations;

    const dataSetLabel = fractalAllocations[0].level.name;
    mapDataSet(dataSetLabel, fractalAllocations, chartDataMap);

    return chartDataMap;
}

function isFromLevel(allocation: Allocation, hierarchyLevelIndex: number) {
    return allocation.structuralId.every(
        (value, index) => index < hierarchyLevelIndex ? value == null : value != null,
    );
}

function extractAllocationPlanFractalHierarchy(allocationPlan: AllocationPlan): FractalAllocationHierarchy {

    const allocationsPerHierarchyLevel = getAllocationsPerHierarchyLevel(allocationPlan);

    const rootNodes =
        mapAllocationPlanHierarchyNodes(
            allocationPlan.structure.hierarchy,
            allocationsPerHierarchyLevel,
            allocationPlan.structure.hierarchy.length - 1,
        );

    return { topAllocations: rootNodes };
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
): FractalAllocation[] {

    let allocationsAtHierarchyLevel = allocationsPerHierarchyLevel[hierarchyLevelIndex];

    if (parentStructuralId) {
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
            key: allocation.structuralId[hierarchyLevelIndex],
            allocation: allocation,
            subAllocations: subNodes,
        } as FractalAllocation;
    });
}

export default allocationPlanChart;

