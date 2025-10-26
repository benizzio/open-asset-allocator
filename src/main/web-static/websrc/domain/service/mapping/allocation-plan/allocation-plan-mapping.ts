import BigNumber from "bignumber.js";
import { AllocationPlan, AllocationPlanDTO, PlannedAllocation } from "../../../allocation-plan";

export function mapToAllocationPlan(
    allocationPlanDTO: AllocationPlanDTO,
): AllocationPlan {

    const allocations = allocationPlanDTO.details.map((allocation) => {
        return {
            ...allocation,
            sliceSizePercentage: new BigNumber(allocation.sliceSizePercentage),
        } as PlannedAllocation;
    });

    return {
        ...allocationPlanDTO,
        details: allocations,
    };
}
