import { ChartContent, ChartDataType } from "../infra/chart/chart-types";
import portfolioChart from "./portfolio-chart/portfolio-chart";
import { PortfolioAtTime, PortfolioDTO } from "../domain/portfolio";
import allocationPlanChart from "./allocation-plan-chart";
import { AllocationPlanDTO } from "../domain/allocation-plan";

function toUnidimensionalMultiChartContent(
    chartDataType: string,
    sourceData: object,
    contextData: unknown,
): ChartContent {

    let baseChartContent: ChartContent;

    if(chartDataType === ChartDataType.PORTFOLIO_HISTORY_1D) {
        baseChartContent = portfolioChart.toUnidimensionalChartContent(
            sourceData as PortfolioAtTime,
            contextData as PortfolioDTO,
        );
    }
    else if(chartDataType === ChartDataType.ASSET_ALLOCATION_PLAN_1D) {
        baseChartContent = allocationPlanChart.toUnidimensionalChartContent(
            sourceData as AllocationPlanDTO,
            contextData as PortfolioDTO,
        );
    }

    return {
        ...baseChartContent,
        multiChart: true,
    };
}


export function toChartContent(chartDataType: string, sourceData: object, contextData: unknown): ChartContent {

    let chartContent: ChartContent;

    switch(chartDataType) {

        case ChartDataType.PORTFOLIO_HISTORY_1D:
        case ChartDataType.ASSET_ALLOCATION_PLAN_1D:
            chartContent = toUnidimensionalMultiChartContent(chartDataType, sourceData, contextData);
            break;

        default: {
            throw new Error("Unsupported chart data type");
        }
    }

    return chartContent;
}

