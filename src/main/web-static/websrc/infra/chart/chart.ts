import { Chart, ChartData, ChartType } from "chart.js";
import { getDoughnutChartOptions, getPieChartOptions } from "./chart-options";
import { CHART_DATA_TYPE_ATTRIBUTE, MEASURAMENT_UNIT_ATTRIBUTE, MeasuramentUnit } from "./chart-types";

const chartData = new Map<string, ChartData>;

function saveChartData(chartId: string, data: ChartData): void {
    chartData.set(chartId, data);
}

function getChartData(chartId: string): ChartData {
    return chartData.get(chartId);
}

function loadChart(canvas: HTMLCanvasElement): void {

    const id = canvas.id;
    const data = getChartData(id);
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

    if (data) {
        new Chart(canvas, {
            type: chartType as ChartType,
            data: data,
            options: options,
        });
    }
}

const chart = {
    saveChartData,
    getChartData,
    loadDescendantCharts(element: HTMLElement)  {
        element.querySelectorAll(`canvas[${CHART_DATA_TYPE_ATTRIBUTE}]`).forEach((canvas: HTMLCanvasElement) => {
            loadChart(canvas);
        });
    },
};

export default chart;

