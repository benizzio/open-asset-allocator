import { AppliedAllocationHierarchyLevel, MappedChartData } from "./portfolio-chart-model";
import { AllocationStructure } from "../../domain/allocation";
import { ChartDataset } from "chart.js";
import chartModule from "../../infra/chart/chart";
import { PortfolioAllocation, PortfolioSnapshot } from "../../domain/portfolio-allocation";

type ReducedAllocation = Pick<PortfolioAllocation, "totalMarketValue" | "cashReserve">;

export function mapChartData(
    portfolioSnapshot: PortfolioSnapshot,
    portfolioStructure: AllocationStructure,
    hierarchyLevelIndex: number,
    appliedHierarchyLevels?: AppliedAllocationHierarchyLevel[],
): MappedChartData {

    const dataset: ChartDataset = {
        data: [],
        label: portfolioSnapshot.observationTimestamp.id.toString(),
        backgroundColor: [],
    };
    const chartData: MappedChartData = { labels: [], keys: [], datasets: [dataset] };
    const colorScale = chartModule.getPieDoughnutChartColorScale();

    const reducedPortfolioAtTime =
        getAccumulatedAllocationsPerProperty(
            portfolioSnapshot,
            portfolioStructure,
            hierarchyLevelIndex,
            appliedHierarchyLevels,
        );

    dataset.backgroundColor = colorScale.colors(reducedPortfolioAtTime.size);

    let index = 0;

    reducedPortfolioAtTime.forEach((value, key) => {

        chartData.labels.push(key + (value.cashReserve ? " (cash reserve)" : ""));
        chartData.keys.push(key);
        dataset.data.push(value.totalMarketValue);

        if(value.cashReserve) {
            chartModule.convertUnidimensionalDatasetBackgroundToPattern(dataset, index);
            index++;
        }
    });

    return chartData;
}

function getAccumulatedAllocationsPerProperty(
    portfolioAtTime: PortfolioSnapshot,
    portfolioStructure: AllocationStructure,
    hierarchyLevelIndex: number,
    appliedHierarchyLevels: AppliedAllocationHierarchyLevel[] = [],
): Map<string, ReducedAllocation> {

    const accumulationProperty = portfolioStructure.hierarchy[hierarchyLevelIndex].field;

    const filteredAllocations = portfolioAtTime.allocations.filter((slice) => {
        return appliedHierarchyLevels.every((filter) => {
            return slice[filter.field] === filter.value;
        });
    });

    const mappedAllocationMarketValues = filteredAllocations.map((allocation) => {
        return {
            label: allocation[accumulationProperty],
            data: allocation.totalMarketValue,
            cashReserve: allocation.cashReserve,
        };
    });

    return mappedAllocationMarketValues.reduce((accumulator, allocation) => {

        const currentKey = allocation.label;
        const currentValue = accumulator.get(currentKey);

        const reducedAllocation = {
            totalMarketValue: !currentValue
                ? allocation.data
                : currentValue.totalMarketValue + allocation.data,
            cashReserve: allocation.cashReserve,
        };
        accumulator.set(currentKey, reducedAllocation);

        return accumulator;
    }, new Map<string, ReducedAllocation>());
}