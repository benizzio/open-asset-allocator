import * as handlebars from "handlebars";
import { PortfolioAtTime } from "../domain/portfolio";
import { ChartData } from "chart.js";
import chart from "./chart";
import portfolioChart from "../application/portfolio-chart";
import { AllocationPlanDTO } from "../domain/allocation";
import allocationPlanChart from "../application/allocation-plan-chart";

export enum ChartDataType {
    PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS = "PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS",
    PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET = "PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET",
    ASSET_ALLOCATION_PLAN_1D_TOP_LEVEL = "ASSET_ALLOCATION_PLAN_1D_TOP_LEVEL",
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

            case ChartDataType.ASSET_ALLOCATION_PLAN_1D_TOP_LEVEL:
                chartData = allocationPlanChart.toUnidimensionalChartData(source as AllocationPlanDTO);
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
