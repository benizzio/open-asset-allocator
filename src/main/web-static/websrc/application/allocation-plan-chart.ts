import { AllocationPlanDTO, FractalAllocation, FractalAllocationHierarchy } from "../domain/allocation";
import { ChartContent, MultiChartDataSource } from "../infra/chart/chart-types";
import { ActiveElement, Chart, ChartData, ChartEvent } from "chart.js";
import { allocationDomainService } from "../domain/allocation-service";
import { changeChartDataSource } from "../infra/chart/chart-utils";
import chartModule from "../infra/chart/chart";

class FractalAllocationMultiChartDataSource extends MultiChartDataSource {

    private currentFractalAllocation: FractalAllocation;

    constructor(
        chartDataMap: Map<string, ChartData>,
        initialDataKey: string,
        private readonly fractalHierarchy: FractalAllocationHierarchy,
    ) {
        super(chartDataMap, initialDataKey);
    }

    getChartData(dataKey?: string): ChartData {
        this.currentFractalAllocation = this.fractalHierarchy.aggregatorAllocationMap.get(dataKey);
        return super.getChartData(dataKey);
    }

    getCurrentFractalSubAllocationKeyByIndex(index: number): string {
        return !this.currentFractalAllocation
            ? this.fractalHierarchy.topAllocations[index].key
            : this.currentFractalAllocation.subAllocations[index].key;
    }
}

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

function chartDataSelectionEventHandler(_: ChartEvent, elements: ActiveElement[], chart: Chart) {

    if (!elements.length) {
        return;
    }

    const dataIndex = elements[0].index;
    const chartId = chart.canvas.id;

    const content = chartModule.getChartContent(chartId);
    const chartDataSource = content.chartDataSource as FractalAllocationMultiChartDataSource;
    const dataKey = chartDataSource.getCurrentFractalSubAllocationKeyByIndex(dataIndex);
    changeChartDataSource(chart, content, dataKey);
}

const allocationPlanChart = {

    toUnidimensionalChartContent(allocationPlanDTO: AllocationPlanDTO): ChartContent {

        const allocationPlan = allocationDomainService.mapToAllocationPlan(allocationPlanDTO);
        const fractalHierarchy = allocationDomainService.mapAllocationPlanFractalHierarchy(allocationPlan);

        const chartDataMap = toChartDataMap(fractalHierarchy);
        const topLevelKey = allocationDomainService.getTopLevelKey(allocationPlan);

        const dataSource = new FractalAllocationMultiChartDataSource(chartDataMap, topLevelKey, fractalHierarchy);
        return { chartDataSource: dataSource, chartInteractions: { onClick: chartDataSelectionEventHandler } };
    },
};

export default allocationPlanChart;

