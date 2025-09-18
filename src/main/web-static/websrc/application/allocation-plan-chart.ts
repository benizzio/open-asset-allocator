import { ChartContent, MultiChartDataSource } from "../infra/chart/chart-types";
import { ActiveElement, Chart, ChartData, ChartDataset, ChartEvent } from "chart.js";
import { allocationDomainService } from "../domain/service/allocation-service";
import { changeChartDataOnDatasource } from "../infra/chart/chart-utils";
import chartModule from "../infra/chart/chart";
import DomUtils from "../infra/dom/dom-utils";
import {
    AllocationPlanDTO,
    FractalPlannedAllocation,
    FractalPlannedAllocationHierarchy,
} from "../domain/allocation-plan";
import { PortfolioDTO } from "../domain/portfolio";

class FractalPlannedAllocationMultiChartDataSource extends MultiChartDataSource {

    private currentAggregator: FractalPlannedAllocation;

    constructor(
        chartDataMap: Map<string, ChartData>,
        initialDataKey: string,
        private readonly fractalHierarchy: FractalPlannedAllocationHierarchy,
    ) {
        super(chartDataMap, initialDataKey);
    }

    getChartData(dataKey?: string): ChartData {
        this.currentAggregator = dataKey ? this.fractalHierarchy.aggregatorAllocationMap.get(dataKey) : null;
        return super.getChartData(dataKey);
    }

    getSubAllocationKeyByIndex(index: number): string {
        return !this.currentAggregator
            ? this.fractalHierarchy.topAllocations[index].key
            : this.currentAggregator.subAllocations[index].key;
    }

    getFractalHierarchy(): FractalPlannedAllocationHierarchy {
        return this.fractalHierarchy;
    }

    getCurrentAggregator(): FractalPlannedAllocation {
        return this.currentAggregator;
    }
}

function mapChildDatasets(fractalHierarchy: FractalPlannedAllocationHierarchy, chartDataMap: Map<string, ChartData>) {

    fractalHierarchy.aggregatorAllocationMap.forEach((fractalAllocation) => {

        const subAllocations = fractalAllocation.subAllocations;

        if(subAllocations && subAllocations.length > 0) {
            mapDataset(fractalAllocation.key, subAllocations, chartDataMap);
        }
    });
}

function mapDataset(
    datasetLabel: string,
    fractalAllocations: FractalPlannedAllocation[],
    chartDataMap: Map<string, ChartData>,
) {

    const dataset: ChartDataset = { data: [], label: datasetLabel, backgroundColor: [] };
    const chartData = { labels: [], datasets: [dataset] };
    const colorScale = chartModule.getPieDoughnutChartColorScale();
    dataset.backgroundColor = colorScale.colors(fractalAllocations.length);

    fractalAllocations.forEach((fractalAllocation, index) => {

        const allocation = fractalAllocation.allocation;

        dataset.data.push(allocation.sliceSizePercentage.toNumber());
        chartData.labels.push(allocation.hierarchicalId[fractalAllocation.level.index]);

        if(allocation.cashReserve) {
            chartModule.convertUnidimensionalDatasetBackgroundToPattern(dataset, index);
        }
    });

    chartDataMap.set(datasetLabel, chartData);
}

function toChartDataMap(fractalHierarchy: FractalPlannedAllocationHierarchy): Map<string, ChartData> {

    const chartDataMap = new Map<string, ChartData>();
    const topFractalAllocations = fractalHierarchy.topAllocations;

    const datasetLabel = topFractalAllocations[0].level.name;
    mapDataset(datasetLabel, topFractalAllocations, chartDataMap);

    mapChildDatasets(fractalHierarchy, chartDataMap);

    return chartDataMap;
}

function chartDataSelectionEventHandler(_: ChartEvent, elements: ActiveElement[], chart: Chart) {

    const dataIndex = elements.length > 0 ? elements[0].index : null;

    const { chartContent, chartDataSource } = getChartContent(chart);
    const dataKey = getSelectedDataKey(chartDataSource, dataIndex);

    if(chartDataSource.hasChartData(dataKey)) {
        changeChartDataOnDatasource(chart, chartContent, dataKey);
    }
}

function getChartContent(chart: Chart) {
    const identifiedContent = chartModule.getChartContentFromChart(chart);

    const chartDataSource =
        identifiedContent.chartContent.chartDataSource as FractalPlannedAllocationMultiChartDataSource;
    return { ...identifiedContent, chartDataSource };
}

function getSelectedDataKey(chartDataSource: FractalPlannedAllocationMultiChartDataSource, dataIndex: number) {

    let dataKey: string;
    const currentAggregator = chartDataSource.getCurrentAggregator();

    if(dataIndex !== null) {
        dataKey = chartDataSource.getSubAllocationKeyByIndex(dataIndex);
    }
    else if(currentAggregator?.superAllocation) {
        dataKey = currentAggregator.superAllocation.key;
    }

    return dataKey;
}

function interactionObserverCallback(event: ChartEvent, elements: ActiveElement[], chart: Chart) {

    const { chartId, chartDataSource } = getChartContent(chart);

    const currentAggregator = chartDataSource.getCurrentAggregator();

    const currentFractalAllocationLevelName = currentAggregator
        ? currentAggregator.subLevel.name
        : chartDataSource.getFractalHierarchy().subLevel.name;
    const currentFractalAllocationLevelValue = currentAggregator ? "for " + currentAggregator.key : "";
    const levelLabel = currentFractalAllocationLevelName + " " + currentFractalAllocationLevelValue;

    if(event.type === "click") {
        const labelId = `#hierarchy-level-${ chartId }`;
        DomUtils.queryFirst(labelId).textContent = levelLabel;
    }
}

const allocationPlanChart = {

    toUnidimensionalChartContent(allocationPlanDTO: AllocationPlanDTO, portfolioDTO: PortfolioDTO): ChartContent {

        const portfolio = allocationDomainService.mapToPortfolio(portfolioDTO);
        const allocationStructure = portfolio.allocationStructure;

        const allocationPlan = allocationDomainService.mapToAllocationPlan(allocationPlanDTO);

        const fractalHierarchy = allocationDomainService.mapToAllocationPlanFractalHierarchy(
            allocationPlan,
            allocationStructure,
        );

        const chartDataMap = toChartDataMap(fractalHierarchy);

        const topLevelKey = allocationDomainService.getTopLevelHierarchyKeyFromAllocationPlan(
            allocationStructure,
        );

        const dataSource = new FractalPlannedAllocationMultiChartDataSource(
            chartDataMap,
            topLevelKey,
            fractalHierarchy,
        );

        return {
            chartDataSource: dataSource,
            chartInteractions: { onClick: chartDataSelectionEventHandler },
            interactionObserverCallback,
        };
    },
};

export default allocationPlanChart;
