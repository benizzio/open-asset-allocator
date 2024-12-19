import { ChartDataset, ChartOptions } from "chart.js";
import format from "../format";
import { ChartInteraction, ChartInteractions, MeasuramentUnit } from "./chart-types";
import { getDatasetSum, LABEL_CALLBACKS } from "./chart-utils";
import BigNumber from "bignumber.js";

const PIE_DOUGHNUT_CHART_OPTIONS: ChartOptions<"pie" | "doughnut"> = {
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
                const total = getDatasetSum(context.chart.data.datasets[0] as ChartDataset<"pie" | "doughnut">);
                const totalBigNumber = new BigNumber(total);
                return format.calculateAndFormatPercent(valueBigNumber, totalBigNumber);
            },
        },
    },
    responsive: true, maintainAspectRatio: true,
};

function getPieDoughnutChartOptions(
    dataType: MeasuramentUnit,
    chartInteractions?: ChartInteractions,
): ChartOptions<"pie" | "doughnut"> {
    return {
        ...PIE_DOUGHNUT_CHART_OPTIONS,
        ...chartInteractions,
        plugins: {
            ...PIE_DOUGHNUT_CHART_OPTIONS.plugins,
            tooltip: { callbacks: { label: LABEL_CALLBACKS[dataType] } },
        },
    };
}

export function buildChartOptions(
    chartType: string,
    dataType: MeasuramentUnit,
    chartInteractions?: ChartInteractions,
    interactionsCallback?: ChartInteraction,
): ChartOptions<"pie" | "doughnut"> {

    const interactions = buildChartInteractions(chartInteractions, interactionsCallback);

    switch (chartType) {
        case "pie":
        case "doughnut":
            return getPieDoughnutChartOptions(dataType, interactions);
    }
}

export function buildChartInteractions(
    interactions?: ChartInteractions,
    oberverCallback?: ChartInteraction,
): ChartInteractions {

    return {
        onClick: (event, elements, chart) => {

            if (interactions?.onClick) {
                interactions.onClick(event, elements, chart);
            }

            if (oberverCallback) {
                oberverCallback(event, elements, chart);
            }
        },
        onHover: (event, elements, chart) => {

            if (interactions?.onHover) {
                interactions.onHover(event, elements, chart);
            }

            if (oberverCallback) {
                oberverCallback(event, elements, chart);
            }
        },
    };
}