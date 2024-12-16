import { ChartDataset, ChartOptions } from "chart.js";
import format from "../format";
import { MeasuramentUnit } from "./chart-types";
import { getDatasetSum, LABEL_CALLBACKS } from "./chart-utils";
import BigNumber from "bignumber.js";

const PIE_DOUGHNUT_CHART_OPTIONS: ChartOptions<"pie"|"doughnut"> = {
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
        datalabels: {
            font: {
                weight: "bolder",
                size: 17,
            },
            formatter: (value: number, context) => {
                const valueBigNumber = new BigNumber(value);
                const total = getDatasetSum(context.chart.data.datasets[0] as ChartDataset<"pie"|"doughnut">);
                const totalBigNumber = new BigNumber(total);
                return format.calculateAndFormatPercent(valueBigNumber, totalBigNumber);
            },
        },
    },
    responsive: true, maintainAspectRatio: true,
};

function getPieDoughnutChartOptions(dataType: MeasuramentUnit) {
    return {
        ...PIE_DOUGHNUT_CHART_OPTIONS,
        plugins: {
            ...PIE_DOUGHNUT_CHART_OPTIONS.plugins,
            tooltip: { callbacks: { label: LABEL_CALLBACKS[dataType] } },
        },
    };
}

export function getPieChartOptions(dataType: MeasuramentUnit): ChartOptions<"pie"> {
    return getPieDoughnutChartOptions(dataType);
}

export function getDoughnutChartOptions(dataType: MeasuramentUnit): ChartOptions<"doughnut"> {
    return getPieDoughnutChartOptions(dataType);
}