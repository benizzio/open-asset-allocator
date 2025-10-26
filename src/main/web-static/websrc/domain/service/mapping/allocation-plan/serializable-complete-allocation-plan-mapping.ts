import { PortfolioDTO } from "../../../portfolio";
import {
    AllocationPlanDTO,
    CompleteAllocationPlan,
    SerializableCompleteAllocationPlan,
    SerializablePortfolioCompleteAllocationPlanSet,
} from "../../../allocation-plan";
import { mapToSerializableFractalHierarchicalAllocationPlan } from "./serializable-fractal-allocation-plan-mapping";
import { mapToCompleteAllocationPlans } from "./complete-allocation-plan-mapping";
import { mapToPortfolio } from "../portfolio-mapping";

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

export function mapToSerializablePortfolioCompleteAllocationPlanSet(
    portfolioDTO: PortfolioDTO,
    allocationPlanDTOs: AllocationPlanDTO[],
): SerializablePortfolioCompleteAllocationPlanSet {
    const portfolio = mapToPortfolio(portfolioDTO);
    const completeAllocationPlans = mapToSerializableCompleteAllocationPlans(portfolioDTO, allocationPlanDTOs);
    return {
        portfolio,
        completeAllocationPlans,
    };
}