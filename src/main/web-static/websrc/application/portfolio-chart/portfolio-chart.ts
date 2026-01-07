import { PortfolioDTO } from "../../domain/portfolio";
import { ChartContent } from "../../infra/chart/chart-types";
import { ActiveElement, Chart, ChartEvent } from "chart.js";
import chartModule from "../../infra/chart/chart";
import { changeChartDataOnDatasource } from "../../infra/chart/chart-utils";
import { FractalPortfolioMultiChartDataSource } from "./portfolio-chart-datasource";
import { MappedChartData } from "./portfolio-chart-model";
import { mapChartData } from "./portfolio-chart-mapping";
import { PortfolioSnapshot } from "../../domain/portfolio-allocation";
import { DomainService } from "../../domain/service";

function changeChartData(
    dataKey: string,
    dataSource: FractalPortfolioMultiChartDataSource,
    chart: Chart,
    chartContent: ChartContent,
) {
    if(dataKey && dataSource.hasChartData(dataKey)) {
        changeChartDataOnDatasource(chart, chartContent, dataKey);
    }
}

function chartDataSelectionEventHandler(_: ChartEvent, elements: ActiveElement[], chart: Chart) {

    const dataIndex = elements.length > 0 ? elements[0].index : null;
    const currentChartData = chart.data as MappedChartData;

    const { chartContent } = chartModule.getChartContentFromChart(chart);
    const dataSource = chartContent.chartDataSource as FractalPortfolioMultiChartDataSource;

    let dataKey: string;

    if(dataIndex !== null) {
        dataKey = dataSource.toNextLevel(currentChartData.keys[dataIndex]);
        changeChartData(dataKey, dataSource, chart, chartContent);
    }
    else {
        dataKey = dataSource.toPreviousLevel();
        changeChartData(dataKey, dataSource, chart, chartContent);
    }
}

function interactionObserverCallback(event: ChartEvent, elements: ActiveElement[], chart: Chart) {

    const { chartId, chartDataSource } = getChartContent(chart);

    const currentHierarchyLevel = chartDataSource.getCurrentHierarchyLevel();

    const currentAggregatorLevelValue = chartDataSource.getLastAppliedHierarchyLevel();

    const currentAggregatorLevelValueLabel = currentAggregatorLevelValue
        ? "for " + currentAggregatorLevelValue.value
        : "";
    const levelLabel = currentHierarchyLevel.name + " " + currentAggregatorLevelValueLabel;

    if(event.type === "click") {
        const labelId = `#hierarchy-level-${ chartId }`;
        const labelElement = window[labelId] as HTMLElement;
        labelElement.textContent = levelLabel;
    }
}

function getChartContent(chart: Chart) {

    const identifiedContent = chartModule.getChartContentFromChart(chart);

    const chartDataSource =
        identifiedContent.chartContent.chartDataSource as FractalPortfolioMultiChartDataSource;

    return { ...identifiedContent, chartDataSource };
}

const portfolioChart = {

    toUnidimensionalChartContent(
        portfolioAtTime: PortfolioSnapshot,
        portfolioDTO: PortfolioDTO,
    ): ChartContent {

        const portfolio = DomainService.mapping.mapToPortfolio(portfolioDTO);
        const portfolioAllocationStructure = portfolio.allocationStructure;

        const { topLevelIndex } =
            DomainService.allocation.getTopHierarchyLevelInfoFromAllocationStructure(portfolioAllocationStructure);

        const chartData = mapChartData(portfolioAtTime, portfolio.allocationStructure, topLevelIndex);

        return {
            chartDataSource: new FractalPortfolioMultiChartDataSource(
                chartData,
                portfolioAtTime,
                portfolioAllocationStructure,
            ),
            chartInteractions: { onClick: chartDataSelectionEventHandler },
            interactionObserverCallback,
        };
    },
};

export default portfolioChart;
