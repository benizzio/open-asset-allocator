import * as handlebars from "handlebars";
import { PortfolioAtTime } from "../../domain/portfolio";
import { ChartData } from "chart.js";
import chart from "../chart/chart";
import portfolioChart from "../../application/portfolio-chart";
import { AllocationPlanDTO } from "../../domain/allocation";
import allocationPlanChart from "../../application/allocation-plan-chart";
import {
    CanvasChartOptions,
    CHART_DATA_TYPE_ATTRIBUTE,
    ChartDataType,
    MEASURAMENT_UNIT_ATTRIBUTE,
    MeasuramentUnit,
} from "../chart/chart-types";

const handlebarsChartHelper = (
    source: unknown,
    chartDataType: ChartDataType,
    options: CanvasChartOptions,
    idPrefix: string,
    idSuffix: string,
) => {

    let chartData: ChartData;

    //TODO apply visitor pattern
    switch (chartDataType) {

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

    const measuramentUnit = options.measuramentUnit || MeasuramentUnit.CURRENCY;

    return `<canvas 
                id="${id}" 
                ${CHART_DATA_TYPE_ATTRIBUTE}="${options.type}" 
                ${MEASURAMENT_UNIT_ATTRIBUTE}="${measuramentUnit}">
            </canvas>`;
};

export function registerHandlebarsChartHelper() {
    handlebars.registerHelper("chart", handlebarsChartHelper);
}
