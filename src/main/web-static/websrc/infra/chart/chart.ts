import { Chart, ChartType } from "chart.js";
import { getDoughnutChartOptions, getPieChartOptions } from "./chart-options";
import { CHART_DATA_TYPE_ATTRIBUTE, ChartContent, MEASURAMENT_UNIT_ATTRIBUTE, MeasuramentUnit } from "./chart-types";

const chartContentRepo = new Map<string, ChartContent>;

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

    let options = {};

    switch (chartType) {
        case "pie":
            options = getPieChartOptions(measuramentUnit);
            break;
        case "doughnut":
            options = getDoughnutChartOptions(measuramentUnit);
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

