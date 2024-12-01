import * as handlebars from "handlebars";
import { PortfolioAtTime } from "../../domain/portfolio";
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

const handlebarsChartHelper = (
    source: PortfolioAtTime | AllocationPlanDTO,
    chartDataType: ChartDataType,
    options: CanvasChartOptions,
    idPrefix: string,
    idSuffix: string,
) => {

    let chartContent: ChartContent;
    let multiChart = false;

    //TODO apply visitor pattern
    switch (chartDataType) {

        case ChartDataType.PORTFOLIO_AT_TIME_1D:
            chartContent = portfolioChart.toUnidimensionalChartContent(source as PortfolioAtTime);
            multiChart = true;
            break;

        case ChartDataType.ASSET_ALLOCATION_PLAN_1D:
            chartContent = allocationPlanChart.toUnidimensionalChartContent(source as AllocationPlanDTO);
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
