import * as handlebars from "handlebars";
import { PortfolioAtTime } from "../../domain/portfolio";
import chart from "../chart/chart";
import portfolioChart from "../../application/portfolio-chart";
import { AllocationPlanDTO } from "../../domain/allocation";
import allocationPlanChart from "../../application/allocation-plan-chart";
import {
    CanvasChartOptions,
    CHART_DATA_TYPE_ATTRIBUTE,
    ChartContent,
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

    let chartContent: ChartContent;

    //TODO apply visitor pattern
    switch (chartDataType) {

        case ChartDataType.PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS:
            chartContent = portfolioChart.toUnidimensionalChartContent(source as PortfolioAtTime, "class");
            break;

        case ChartDataType.PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET:
            chartContent = portfolioChart.toUnidimensionalChartContent(source as PortfolioAtTime, "assetTicker");
            break;

        case ChartDataType.ASSET_ALLOCATION_PLAN_1D_TOP_LEVEL:
            chartContent = allocationPlanChart.toUnidimensionalChartContent(source as AllocationPlanDTO);
            break;
    }

    const id = `${idPrefix}-${idSuffix}`;
    chart.saveChartContent(id, chartContent);

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
