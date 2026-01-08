import { ChartOptions } from "chart.js";
import Format from "../format";
import { ChartInteraction, ChartInteractions, LocalChartOptions } from "./chart-types";
import { LABEL_CALLBACKS, valueAsPercentageOfDataset } from "./chart-utils";
import { BOOTSTRAP_BODY_BACKGROUND_COLOR } from "../color";
import { Context } from "chartjs-plugin-datalabels/types/context";

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
            color: BOOTSTRAP_BODY_BACKGROUND_COLOR,
            formatter: (value: number, context: Context) => {

                const percent = valueAsPercentageOfDataset(value, context);

                if(percent.isLessThan(0.03)) {
                    return "";
                }

                return Format.formatPercent(percent.toNumber());
            },
        },
    },
    responsive: true, maintainAspectRatio: true,
};

function getPieDoughnutChartOptions(
    localChartOptions: LocalChartOptions,
    chartInteractions?: ChartInteractions,
): ChartOptions<"pie" | "doughnut"> {
    return {
        ...PIE_DOUGHNUT_CHART_OPTIONS,
        ...localChartOptions,
        ...chartInteractions,
        plugins: {
            ...PIE_DOUGHNUT_CHART_OPTIONS.plugins,
            tooltip: { callbacks: { label: LABEL_CALLBACKS[localChartOptions.measuramentUnit] } },
        },
    } as ChartOptions<"pie" | "doughnut">;
}

export function buildChartOptions(
    localChartOptions: LocalChartOptions,
    chartInteractions?: ChartInteractions,
    interactionsCallback?: ChartInteraction,
): ChartOptions {

    const interactions = buildChartInteractions(chartInteractions, interactionsCallback);

    switch(localChartOptions.type) {
        case "pie":
        case "doughnut":
            return getPieDoughnutChartOptions(localChartOptions, interactions) as ChartOptions;
    }
}

export function buildChartInteractions(
    interactions?: ChartInteractions,
    oberverCallback?: ChartInteraction,
): ChartInteractions {

    return {
        onClick: (event, elements, chart) => {

            if(interactions?.onClick) {
                interactions.onClick(event, elements, chart);
            }

            if(oberverCallback) {
                oberverCallback(event, elements, chart);
            }
        },
        onHover: (event, elements, chart) => {

            if(interactions?.onHover) {
                interactions.onHover(event, elements, chart);
            }

            if(oberverCallback) {
                oberverCallback(event, elements, chart);
            }
        },
    };
}