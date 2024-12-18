import { PortfolioAtTime } from "../domain/portfolio";
import { ChartContent, SingleChartDataSource } from "../infra/chart/chart-types";

function getAccumulatedSlicesPerProperty(portfolioAtTime: PortfolioAtTime, dimensionProperty: string) {
    return portfolioAtTime.slices.map((slice) => {
        return {
            label: slice[dimensionProperty],
            data: slice.totalMarketValue,
        };
    }).reduce((accumulator, slice) => {
        const currentKey = slice.label;
        const currentValue = accumulator.get(currentKey);
        accumulator.set(currentKey, !currentValue ? slice.data : currentValue + slice.data);
        return accumulator;
    }, new Map<string, number>());
}

const portfolioChart = {

    toUnidimensionalChartContent(
        portfolioAtTime: PortfolioAtTime,
        dimensionProperty: string,
    ): ChartContent {

        const dataSet = { data: [], label: portfolioAtTime.timeFrameTag };
        const chartData = { labels: [], datasets: [dataSet] };

        const reducedPortfolioAtTime =
            getAccumulatedSlicesPerProperty(portfolioAtTime, dimensionProperty);

        reducedPortfolioAtTime.forEach((value, key) => {
            chartData.labels.push(key);
            dataSet.data.push(value);
        });

        return { chartDataSource: new SingleChartDataSource(chartData) };
    },
};

export default portfolioChart;
