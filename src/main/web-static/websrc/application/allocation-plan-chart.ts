import { ChartData, ChartType } from "chart.js";
import { Allocation, AllocationPlanDTO } from "../domain/allocation";
import BigNumber from "bignumber.js";

const allocationPlanChart = {
    toUnidimensionalChartData(allocationPlanDTO: AllocationPlanDTO): ChartData<ChartType, number[], string> {

        //TODO continue
        const allocations = allocationPlanDTO.details.map((allocation) => {
            return {
                ...allocation,
                sliceSizePercentage: new BigNumber(allocation.sliceSizePercentage),
            } as Allocation;
        });

        const allocationPlan = {
            ...allocationPlanDTO,
            details: allocations,
        };
        
        const dataSet = { data: [], label: allocationPlan.name };
        const chartData = { labels: [], datasets: [dataSet] };

        allocationPlan.details.filter(allocation => allocation.structuralId[0] == null).forEach((allocation) => {
            chartData.labels.push(allocation.structuralId[1]);
            dataSet.data.push(allocation.sliceSizePercentage.toNumber());
        });

        return chartData;
    },
};

export default allocationPlanChart;