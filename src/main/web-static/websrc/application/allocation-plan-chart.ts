import { AllocationPlanDTO, FractalAllocation, FractalAllocationHierarchy } from "../domain/allocation";
import { ChartContent, MultiChartDataSource } from "../infra/chart/chart-types";
import { ChartData } from "chart.js";
import { allocationDomainService } from "../domain/allocation-service";

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
    });

    chartDataMap.set(dataSetLabel, chartData);
}

function mapChildDatasets(fractalHierarchy: FractalAllocationHierarchy, chartDataMap: Map<string, ChartData>) {

    fractalHierarchy.aggregatorAllocationMap.forEach((fractalAllocation) => {

        const subAllocations = fractalAllocation.subAllocations;

        if (subAllocations && subAllocations.length > 0) {
            mapDataSet(fractalAllocation.key, subAllocations, chartDataMap);
        }
    });
}

function toChartDataMap(fractalHierarchy: FractalAllocationHierarchy): Map<string, ChartData> {

    const chartDataMap = new Map<string, ChartData>();
    const topFractalAllocations = fractalHierarchy.topAllocations;

    const dataSetLabel = topFractalAllocations[0].level.name;
    mapDataSet(dataSetLabel, topFractalAllocations, chartDataMap);

    mapChildDatasets(fractalHierarchy, chartDataMap);

    return chartDataMap;
}

const allocationPlanChart = {

    toUnidimensionalChartContent(allocationPlanDTO: AllocationPlanDTO): ChartContent {

        const allocationPlan = allocationDomainService.mapToAllocationPlan(allocationPlanDTO);
        const fractalHierarchy = allocationDomainService.mapAllocationPlanFractalHierarchy(allocationPlan);

        const chartDataMap = toChartDataMap(fractalHierarchy);
        const topLevelKey = allocationDomainService.getTopLevelKey(allocationPlan);

        return { chartDataSource: new MultiChartDataSource(chartDataMap, topLevelKey) };
    },
};

export default allocationPlanChart;

