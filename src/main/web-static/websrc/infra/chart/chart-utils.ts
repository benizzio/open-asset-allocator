import { ChartDataset, ChartType, TooltipItem } from "chart.js";
import { MeasuramentUnit, UNIDIMENSIONAL_DATASET_SUM_FIELD } from "./chart-types";
import format from "../format";

export function getDatasetSum(dataSet: ChartDataset<"pie"|"doughnut">) {
    if(!dataSet[UNIDIMENSIONAL_DATASET_SUM_FIELD]) {
        const data = dataSet.data;
        dataSet[UNIDIMENSIONAL_DATASET_SUM_FIELD] = data.reduce((accumulator, value) => accumulator + value, 0);
    }
    return dataSet[UNIDIMENSIONAL_DATASET_SUM_FIELD];
}

function buildLabelPrefix(context: TooltipItem<ChartType>) {
    let label = "";
    const datasetLabel = context.dataset.label;

    if (context.chart.data.datasets.length > 1 && datasetLabel) {
        label = datasetLabel + ": ";
    }
    return label;
}

export const LABEL_CALLBACKS = {
    [MeasuramentUnit.CURRENCY]: (context: TooltipItem<ChartType>) => {
        const label = buildLabelPrefix(context);
        return label + format.formatCurrency(context.parsed);
    },
    [MeasuramentUnit.PERCENTAGE]: (context: TooltipItem<ChartType>) => {
        const label = buildLabelPrefix(context);
        return label + format.formatPercent(context.parsed);
    },
};