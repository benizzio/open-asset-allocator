import { PortfolioDTO } from "../../../portfolio";
import {
    AllocationPlanDTO,
    CompleteAllocationPlan,
    SerializableCompleteAllocationPlan,
} from "../../../allocation-plan";
import { mapToSerializableFractalHierarchicalAllocationPlan } from "./serializable-fractal-allocation-plan-mapping";
import { mapToCompleteAllocationPlans } from "./complete-allocation-plan-mapping";

export function mapToSerializableCompleteAllocationPlans(
    portfolioDTO: PortfolioDTO,
    allocationPlanDTOs: AllocationPlanDTO[],
): SerializableCompleteAllocationPlan[] {
    const completeAllocationPlans = mapToCompleteAllocationPlans(
        portfolioDTO,
        allocationPlanDTOs,
    );
    return completeAllocationPlans.map((completeAllocationPlan: CompleteAllocationPlan) =>
        mapToSerializableCompleteAllocationPlan(completeAllocationPlan));
}

export function mapToSerializableCompleteAllocationPlan(
    completeAllocationPlan: CompleteAllocationPlan,
): SerializableCompleteAllocationPlan {
    return {
        allocationPlan: completeAllocationPlan.allocationPlan,
        fractalHierarchicalPlan:
            mapToSerializableFractalHierarchicalAllocationPlan(
                completeAllocationPlan.fractalHierarchicalPlan,
            ),
        topLevelKey: completeAllocationPlan.topLevelKey,
    };
}