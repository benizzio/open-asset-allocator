import { ActiveElement, Chart, ChartData, ChartEvent, ChartOptions, ChartType, CoreChartOptions } from "chart.js";

export enum MeasuramentUnit {
    CURRENCY = "CURRENCY",
    PERCENTAGE = "PERCENTAGE",
}

export enum ChartDataType {
    PORTFOLIO_HISTORY_1D = "PORTFOLIO_HISTORY_1D",
    ASSET_ALLOCATION_PLAN_1D = "ASSET_ALLOCATION_PLAN_1D",
}

export type LocalChartOptions = { type: string, measuramentUnit?: MeasuramentUnit } & ChartOptions;

export const CHART_ATTRIBUTE = "data-chart";
export const CHART_OPTIONS_JSON_ELEMENT_ID = "chart-options-element";
export const UNIDIMENSIONAL_DATASET_SUM_FIELD = "sum";

export type ChartInteractions = Partial<Pick<CoreChartOptions<ChartType>, "onClick" | "onHover">>;
export type ChartInteraction = (event: ChartEvent, elements: ActiveElement[], chart: Chart) => void;

export type ChartContent = {
    chartDataSource: ChartDataSource,
    chartInteractions?: ChartInteractions,
    interactionObserverCallback?: ChartInteraction,
    multiChart?: boolean,
};

export interface ChartDataSource {

    getChartData(dataKey?: string): ChartData;

    hasChartData(dataKey?: string): boolean;

    accept(visitor: ChartDataSourceVisitor): void;
}

export class SingleChartDataSource implements ChartDataSource {

    constructor(private readonly chartData: ChartData) {
    }

    getChartData(): ChartData {
        return this.chartData;
    }

    accept(visitor: ChartDataSourceVisitor): void {
        visitor.visitSingleChartDataSource(this);
    }

    hasChartData(): boolean {
        return !!this.chartData;
    }
}

export class MultiChartDataSource implements ChartDataSource {

    constructor(protected readonly chartDataMap: Map<string, ChartData>, private readonly initialDataKey: string) {
    }

    getChartData(dataKey?: string): ChartData {
        return this.chartDataMap.get(dataKey || this.initialDataKey);
    }

    accept(visitor: ChartDataSourceVisitor): void {
        visitor.visitMultiChartDataSource(this);
    }

    hasChartData(dataKey?: string): boolean {
        return this.chartDataMap.has(dataKey || this.initialDataKey);
    }
}

export interface ChartDataSourceVisitor {

    visitSingleChartDataSource(dataSource: SingleChartDataSource): void;

    visitMultiChartDataSource(dataSource: MultiChartDataSource): void;
}