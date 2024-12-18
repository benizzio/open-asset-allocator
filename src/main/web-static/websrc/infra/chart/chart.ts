import { ActiveElement, Chart, ChartEvent, ChartType } from "chart.js";
import { buildChartOptions } from "./chart-options";
import {
    CHART_DATA_TYPE_ATTRIBUTE,
    ChartContent,
    MEASURAMENT_UNIT_ATTRIBUTE,
    MeasuramentUnit,
    MULTI_CHART_DATA_ATTRIBUTE,
} from "./chart-types";
import { visitMultiChartDataSource } from "./chart-utils";

const MULTI_CHART_INTERACTIONS = { onClick: multiChartDataSelectionEventHandler };

const chartContentRepo = new Map<string, ChartContent>;

function multiChartDataSelectionEventHandler(event: ChartEvent, elements: ActiveElement[], chart: Chart) {

    if (!elements.length) {
        return;
    }

    const dataIndex = elements[0].index;
    const dataKey = chart.data.labels[dataIndex] as string;
    const chartId = chart.canvas.id;
    const content = getChartContent(chartId);
    visitMultiChartDataSource(content, dataKey, chart);
}

function saveChartContent(chartId: string, content: ChartContent): void {
    chartContentRepo.set(chartId, content);
}

function getChartContent(chartId: string): ChartContent {
    return chartContentRepo.get(chartId);
}

function loadChart(canvas: HTMLCanvasElement): void {

    const id = canvas.id;
    const content = getChartContent(id);
    const chartType = canvas.getAttribute(CHART_DATA_TYPE_ATTRIBUTE);
    const measuramentUnit = canvas.getAttribute(MEASURAMENT_UNIT_ATTRIBUTE) as MeasuramentUnit;
    const multiChart = canvas.hasAttribute(MULTI_CHART_DATA_ATTRIBUTE);

    let options = {};

    switch (chartType) {
        case "pie":
            options = buildChartOptions("pie", measuramentUnit, multiChart ? MULTI_CHART_INTERACTIONS : null);
            break;
        case "doughnut":
            options = buildChartOptions(
                "doughnut",
                measuramentUnit,
                multiChart ? MULTI_CHART_INTERACTIONS : null,
            );
            break;
    }

    if (content) {
        new Chart(canvas, {
            type: chartType as ChartType,
            data: content.chartDataSource.getChartData(),
            options: options,
        });
    }
}

const chart = {
    saveChartContent,
    getChartContent,
    loadDescendantCharts(element: HTMLElement) {
        element.querySelectorAll(`canvas[${ CHART_DATA_TYPE_ATTRIBUTE }]`).forEach((canvas: HTMLCanvasElement) => {
            loadChart(canvas);
        });
    },
};

export default chart;

