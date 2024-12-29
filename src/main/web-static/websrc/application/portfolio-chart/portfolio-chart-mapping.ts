import { PortfolioAtTime } from "../../domain/portfolio";
import { AppliedAllocationHierarchyLevel, MappedChartData } from "./portfolio-chart-model";
import { AllocationStructure } from "../../domain/allocation";

export function mapChartData(
    portfolioAtTime: PortfolioAtTime,
    portfolioStructure: AllocationStructure,
    hierarchyLevelIndex: number,
    appliedHierarchyLevels?: AppliedAllocationHierarchyLevel[],
): MappedChartData {

    const dataset = { data: [], label: portfolioAtTime.timeFrameTag };
    const chartData = { labels: [], keys: [], datasets: [dataset] };

    const reducedPortfolioAtTime =
        getAccumulatedSlicesPerProperty(
            portfolioAtTime,
            portfolioStructure,
            hierarchyLevelIndex,
            appliedHierarchyLevels,
        );

    reducedPortfolioAtTime.forEach((value, key) => {
        chartData.labels.push(key);
        chartData.keys.push(key);
        dataset.data.push(value);
    });

    return chartData;
}

function getAccumulatedSlicesPerProperty(
    portfolioAtTime: PortfolioAtTime,
    portfolioStructure: AllocationStructure,
    hierarchyLevelIndex: number,
    appliedHierarchyLevels: AppliedAllocationHierarchyLevel[] = [],
): Map<string, number> {

    const accumulationProperty = portfolioStructure.hierarchy[hierarchyLevelIndex].field;

    const filteredSlices = portfolioAtTime.allocations.filter((slice) => {
        return appliedHierarchyLevels.every((filter) => {
            return slice[filter.field] === filter.value;
        });
    });

    const mappedSliceMarketValues = filteredSlices.map((slice) => {
        return {
            label: slice[accumulationProperty],
            data: slice.totalMarketValue,
        };
    });

    return mappedSliceMarketValues.reduce((accumulator, slice) => {
        const currentKey = slice.label;
        const currentValue = accumulator.get(currentKey);
        accumulator.set(currentKey, !currentValue ? slice.data : currentValue + slice.data);
        return accumulator;
    }, new Map<string, number>());
}