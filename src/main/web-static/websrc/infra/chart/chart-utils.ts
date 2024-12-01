import { Chart, ChartDataset, ChartType, TooltipItem } from "chart.js";
import { ChartContent, MeasuramentUnit, MultiChartDataSource, UNIDIMENSIONAL_DATASET_SUM_FIELD } from "./chart-types";
import format from "../format";

export function getDatasetSum(dataset: ChartDataset<"pie" | "doughnut">) {
    if (!dataset[UNIDIMENSIONAL_DATASET_SUM_FIELD]) {
        const data = dataset.data;
        dataset[UNIDIMENSIONAL_DATASET_SUM_FIELD] = data.reduce((accumulator, value) => accumulator + value, 0);
    }
    return dataset[UNIDIMENSIONAL_DATASET_SUM_FIELD];
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

export function changeChartDataOnDatasource(
    chart: Chart,
    content: ChartContent,
    dataKey: string,
) {
    content.chartDataSource.accept({
        visitMultiChartDataSource(dataSource: MultiChartDataSource) {
            const chartData = dataSource.getChartData(dataKey);

            if (chartData) {
                chart.data = chartData;
                chart.update();
            }

        },
        visitSingleChartDataSource() {
            // no op
        },
    });
}