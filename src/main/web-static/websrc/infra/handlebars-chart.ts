import * as handlebars from "handlebars";
import { PortfolioAtTime } from "../domain/portfolio";
import { ChartData } from "chart.js";
import chart from "./chart";
import portfolioChart from "../application/portfolio-chart";

export enum ChartDataType {
    PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS = "PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS",
    PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET = "PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET",
}

const handlebarsChartHelper =
    (source: unknown, dataType: ChartDataType, type: string, idPrefix: string, idSuffix: string) => {

        let chartData: ChartData;

        //TODO apply visitor pattern
        switch (dataType) {
            case ChartDataType.PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS:
                chartData = portfolioChart.toUnidimensionalChartData(source as PortfolioAtTime, "class");
                break;
            case ChartDataType.PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET:
                chartData = portfolioChart.toUnidimensionalChartData(source as PortfolioAtTime, "assetTicker");
                break;
        }

        const id = `${idPrefix}-${idSuffix}`;
        chart.saveChartData(id, chartData);

        return `<canvas id="${id}" data-chart-type="${type}"></canvas>`;
    };

const handlebarsChart = {
    registerHandlebarsChartHelper() {
        handlebars.registerHelper("chart", handlebarsChartHelper);
    },
};

export default handlebarsChart;
