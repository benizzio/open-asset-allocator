import { ChartData, ChartType, CoreChartOptions } from "chart.js";

export enum MeasuramentUnit {
    CURRENCY = "CURRENCY",
    PERCENTAGE = "PERCENTAGE",
}

export enum ChartDataType {
    PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS = "PORTFOLIO_AT_TIME_1D_PROPERTY_CLASS",
    PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET = "PORTFOLIO_AT_TIME_1D_PROPERTY_ASSET",
    ASSET_ALLOCATION_PLAN_1D_TOP_LEVEL = "ASSET_ALLOCATION_PLAN_1D_TOP_LEVEL",
}

export type CanvasChartOptions = { type: string, measuramentUnit?: MeasuramentUnit };

export const CHART_DATA_TYPE_ATTRIBUTE = "data-chart-type";
export const MEASURAMENT_UNIT_ATTRIBUTE = "data-measurament";
export const UNIDIMENSIONAL_DATASET_SUM_FIELD = "sum";

export type ChartContent = {
    chartDataSource: ChartDataSource
    chartInteractions?: Pick<CoreChartOptions<ChartType>, "onClick">
};

export interface ChartDataSource {getChartData(dataKey?: string): ChartData;}

export class SingleChartDataSource implements ChartDataSource {

    constructor(private readonly chartData: ChartData) {
    }

    getChartData(): ChartData {
        return this.chartData;
    }
}

export class MultiChartDataSource implements ChartDataSource {

    constructor(private readonly chartDataMap: Map<string, ChartData>, private readonly initialDataKey: string) {
    }

    getChartData(dataKey?: string): ChartData {
        return this.chartDataMap.get(dataKey || this.initialDataKey);
    }
}