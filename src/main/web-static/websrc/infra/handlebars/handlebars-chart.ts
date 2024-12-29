import * as handlebars from "handlebars";
import { PortfolioAtTime, PortfolioDTO } from "../../domain/portfolio";
import chart from "../chart/chart";
import portfolioChart from "../../application/portfolio-chart/portfolio-chart";
import allocationPlanChart from "../../application/allocation-plan-chart";
import {
    CanvasChartOptions,
    CHART_DATA_TYPE_ATTRIBUTE,
    ChartContent,
    ChartDataType,
    MEASURAMENT_UNIT_ATTRIBUTE,
    MeasuramentUnit,
    MULTI_CHART_DATA_ATTRIBUTE,
} from "../chart/chart-types";
import { AllocationPlanDTO } from "../../domain/allocation-plan";
import DomUtils from "../dom/dom-utils";

const handlebarsChartHelper = (
    source: PortfolioAtTime | AllocationPlanDTO,
    chartDataType: ChartDataType,
    options: CanvasChartOptions,
    idPrefix: string,
    idSuffix: string,
    contextDataSelector: string,
) => {

    let chartContent: ChartContent;
    let multiChart = false;

    const contextData = DomUtils.getContextDataFromRoot(contextDataSelector);

    //TODO apply visitor pattern
    //TODO generalize code
    switch(chartDataType) {

        case ChartDataType.PORTFOLIO_HISTORY_1D:
            chartContent = portfolioChart.toUnidimensionalChartContent(
                source as PortfolioAtTime,
                contextData as PortfolioDTO,
            );
            multiChart = true;
            break;

        case ChartDataType.ASSET_ALLOCATION_PLAN_1D:
            chartContent = allocationPlanChart.toUnidimensionalChartContent(
                source as AllocationPlanDTO,
                contextData as PortfolioDTO,
            );
            multiChart = true;
            break;
    }

    const id = `${ idPrefix }-${ idSuffix }`;
    chart.saveChartContent(id, chartContent);

    const measuramentUnit = options.measuramentUnit || MeasuramentUnit.CURRENCY;

    return `<canvas 
                id="${ id }" 
                ${ CHART_DATA_TYPE_ATTRIBUTE }="${ options.type }" 
                ${ MEASURAMENT_UNIT_ATTRIBUTE }="${ measuramentUnit }"
                ${ multiChart ? MULTI_CHART_DATA_ATTRIBUTE : "" }
            ></canvas>`;
};

export function registerHandlebarsChartHelper() {
    handlebars.registerHelper("chart", handlebarsChartHelper);
}
