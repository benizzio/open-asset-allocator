import { Chart, ChartType } from "chart.js";
import { buildChartOptions } from "./chart-options";
import { CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent, LocalChartOptions } from "./chart-types";
import DomUtils from "../dom/dom-utils";
import { convertUnidimensionalDatasetBackgroundToPattern, getPieDoughnutChartColorScale } from "./chart-utils";

const chartContentRepo = new Map<string, ChartContent>;

function saveChartContent(chartId: string, content: ChartContent): void {
    chartContentRepo.set(chartId, content);
}

function getChartContent(chartId: string): ChartContent {
    return chartContentRepo.get(chartId);
}

function getChartContentFromChart(chart: Chart) {
    const chartId = chart.canvas.id;
    const chartContent = getChartContent(chartId);
    return { chartId, chartContent };
}

function loadChart(canvas: HTMLCanvasElement): void {

    const id = canvas.id;
    const content = getChartContent(id);

    const interactionObserverCallback = content.interactionObserverCallback;
    const chartInteractions = content.chartInteractions;

    const localOptionsElementId = canvas.getAttribute(CHART_OPTIONS_JSON_ELEMENT_ID);
    const localOptions = DomUtils.getContextDataFromRoot("#" + localOptionsElementId) as LocalChartOptions;
    const chartType = localOptions.type;


    const options = buildChartOptions(
        localOptions,
        chartInteractions,
        interactionObserverCallback,
    );

    if(content) {
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
    getChartContentFromChart,
    loadDescendantCharts(element: HTMLElement) {
        element.querySelectorAll(`canvas[${ CHART_ATTRIBUTE }]`).forEach((canvas: HTMLCanvasElement) => {
            loadChart(canvas);
        });
    },
    getPieDoughnutChartColorScale,
    convertUnidimensionalDatasetBackgroundToPattern,
};

export default chart;