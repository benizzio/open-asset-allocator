import { Chart, ChartData, ChartDataset, ChartOptions, ChartType } from "chart.js";
import format from "./format";

const chartData = new Map<string, ChartData>;
const UNIDIMENSIONAL_DATASET_SUM_FIELD = "sum";

const PIE_CHART_OPTIONS: ChartOptions<"pie"> = {
    plugins: {
        legend: {
            position: "right",
            labels: {
                boxHeight: 20,
                font: {
                    weight: "bolder",
                    size: 15,
                },
            },
        },
        tooltip: {
            callbacks: {
                label: (context) => {

                    let label = "";
                    const datasetLabel = context.dataset.label;

                    if(context.chart.data.datasets.length > 1 && datasetLabel)  {
                        label = datasetLabel + ": ";
                    }

                    return label +
                        format.formatCurrency(context.parsed);
                },
            },
        },
        datalabels: {
            font: {
                weight: "bolder",
                size: 17,
            },
            formatter: (value: number, context) => {
                const total = getDatasetSum(context.chart.data.datasets[0] as ChartDataset<"pie">);
                return format.calculateAndFormatPercent(value, total);
            },
        },
    },
    responsive: true, maintainAspectRatio: true,
};

function getDatasetSum(dataSet: ChartDataset<"pie"|"doughnut">) {
    if(!dataSet[UNIDIMENSIONAL_DATASET_SUM_FIELD]) {
        const data = dataSet.data;
        dataSet[UNIDIMENSIONAL_DATASET_SUM_FIELD] = data.reduce((accumulator, value) => accumulator + value, 0);
    }
    return dataSet[UNIDIMENSIONAL_DATASET_SUM_FIELD];
}

function getPieChartOptions(): ChartOptions<"pie"> {
    return PIE_CHART_OPTIONS;
}

const chart = {
    saveChartData(chartId: string, data: ChartData): void {
        chartData.set(chartId, data);
    },
    getChartData(chartId: string): ChartData {
        return chartData.get(chartId);
    },
    loadChart(canvas: HTMLCanvasElement): void {

        const id = canvas.id;
        const data = this.getChartData(id);
        const chartType = canvas.getAttribute("data-chart-type");

        let options = {};

        if(chartType === "pie") {
            options = getPieChartOptions();
        }

        if (data) {
            new Chart(canvas, {
                type: chartType as ChartType,
                data: data,
                options: options,
            });
        }
    },
};

export default chart;

