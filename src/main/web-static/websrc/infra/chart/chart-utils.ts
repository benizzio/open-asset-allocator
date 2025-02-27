import { Chart, ChartDataset, ChartType, TooltipItem } from "chart.js";
import { ChartContent, MeasuramentUnit, MultiChartDataSource, UNIDIMENSIONAL_DATASET_SUM_FIELD } from "./chart-types";
import format from "../format";
import BigNumber from "bignumber.js";
import { Context } from "chartjs-plugin-datalabels/types/context";
import { PAUL_TOL_PALETTE } from "../color";
import chroma from "chroma-js";
import * as pattern from "patternomaly";

export function getDatasetSum(dataset: ChartDataset<"pie" | "doughnut">) {
    if(!dataset[UNIDIMENSIONAL_DATASET_SUM_FIELD]) {
        const data = dataset.data;
        dataset[UNIDIMENSIONAL_DATASET_SUM_FIELD] = data.reduce((accumulator, value) => accumulator + value, 0);
    }
    return dataset[UNIDIMENSIONAL_DATASET_SUM_FIELD];
}

function buildLabelPrefix(context: TooltipItem<ChartType>) {

    let label = "";
    const datasetLabel = context.dataset.label;

    if(context.chart.data.datasets.length > 1 && datasetLabel) {
        label = datasetLabel + ": ";
    }
    return label;
}

export const LABEL_CALLBACKS = {
    [MeasuramentUnit.CURRENCY]: (context: TooltipItem<ChartType>) => {

        const label = buildLabelPrefix(context);
        const percent = valueAsPercentageOfDataset(context.parsed, context);
        const chartType = (context.chart.config as { type: ChartType }).type;

        let percentLabel = "";

        if(chartType === "pie" || chartType === "doughnut") {
            percentLabel = percent.isLessThan(0.03) ? " (" + format.formatPercent(percent.toNumber()) + ")" : "";
        }

        return label + format.formatCurrency(context.parsed) + percentLabel;
    },
    [MeasuramentUnit.PERCENTAGE]: (context: TooltipItem<ChartType>) => {
        const label = buildLabelPrefix(context);
        return label + format.formatPercent(context.parsed);
    },
};

export function changeChartDataOnDatasource(
    chart: Chart,
    content: ChartContent,
    dataKey: string,
) {
    content.chartDataSource.accept({
        visitMultiChartDataSource(dataSource: MultiChartDataSource) {

            const chartData = dataSource.getChartData(dataKey);

            if(chartData) {
                chart.data = chartData;
                chart.update();
            }
        },
        visitSingleChartDataSource() {
            // no op
        },
    });
}

export function valueAsPercentageOfDataset(value: number, context: Context | TooltipItem<"pie" | "doughnut">) {
    const valueBigNumber = new BigNumber(value);
    const total = getDatasetSum(context.dataset as ChartDataset<"pie" | "doughnut">);
    const totalBigNumber = new BigNumber(total);
    return valueBigNumber.div(totalBigNumber);
}

export function getPieDoughnutChartColorScale() {
    return chroma.scale(PAUL_TOL_PALETTE.qualitative).mode("lrgb");
}

export function convertUnidimensionalDatasetBackgroundToPattern(dataset: ChartDataset, index: number) {

    const backgroundColor = dataset.backgroundColor[index];
    const patternColor = chroma(backgroundColor).darken(0.5).hex();

    dataset.backgroundColor[index] = pattern.draw(
        "diagonal",
        backgroundColor,
        patternColor,
    );
}