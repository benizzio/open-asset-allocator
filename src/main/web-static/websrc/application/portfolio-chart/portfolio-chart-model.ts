import { ChartData } from "chart.js";

export type AppliedAllocationHierarchyLevel = {
    field: string;
    value: string | number;
};

export type MappedChartData = ChartData & { keys: string[] };