import BigNumber from "bignumber.js";
import { AllocationPlan, AllocationPlanDTO, PlannedAllocation } from "../../../allocation-plan";

export function mapToAllocationPlan(
    allocationPlanDTO: AllocationPlanDTO,
): AllocationPlan {

    const allocations = allocationPlanDTO.details.map((allocation, index) => {

        try {
            const sliceSizePercentage = new BigNumber(allocation.sliceSizePercentage);

            if(sliceSizePercentage.isNaN()) {
                throw new Error("BigNumber resolved to NaN");
            }

            return {
                ...allocation,
                sliceSizePercentage: sliceSizePercentage,
            } as PlannedAllocation;
        } catch(error) {
            throw new Error(
                `Invalid sliceSizePercentage "${ allocation.sliceSizePercentage }" for allocation` +
                ` [${ index }] (id: ${ allocation.hierarchicalId })`,
                { cause: error },
            );
        }
    });

    return {
        ...allocationPlanDTO,
        details: allocations,
    };
}
