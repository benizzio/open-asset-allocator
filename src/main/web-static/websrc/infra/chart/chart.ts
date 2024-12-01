import { Chart, ChartType } from "chart.js";
import { buildChartOptions } from "./chart-options";
import { CHART_DATA_TYPE_ATTRIBUTE, ChartContent, MEASURAMENT_UNIT_ATTRIBUTE, MeasuramentUnit } from "./chart-types";

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
    const chartType = canvas.getAttribute(CHART_DATA_TYPE_ATTRIBUTE);
    const measuramentUnit = canvas.getAttribute(MEASURAMENT_UNIT_ATTRIBUTE) as MeasuramentUnit;
    const chartInteractions = content.chartInteractions;

    let options = {};

    switch (chartType) {
        case "pie":
            options = buildChartOptions(
                "pie",
                measuramentUnit,
                chartInteractions,
                interactionObserverCallback,
            );
            break;
        case "doughnut":
            options = buildChartOptions(
                "doughnut",
                measuramentUnit,
                chartInteractions,
                interactionObserverCallback,
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
    getChartContentFromChart,
    loadDescendantCharts(element: HTMLElement) {
        element.querySelectorAll(`canvas[${ CHART_DATA_TYPE_ATTRIBUTE }]`).forEach((canvas: HTMLCanvasElement) => {
            loadChart(canvas);
        });
    },
};

export default chart;