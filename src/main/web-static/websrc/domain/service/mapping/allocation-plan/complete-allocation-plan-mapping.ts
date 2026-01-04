import { PortfolioDTO } from "../../../portfolio";
import { AllocationPlanDTO, CompleteAllocationPlan } from "../../../allocation-plan";
import { mapToPortfolio } from "../portfolio-mapping";
import { mapToAllocationPlan } from "./allocation-plan-mapping";
import { mapToFractalHierarchicalAllocationPlan } from "./fractal-allocation-plan-mapping";
import { AllocationDomainService } from "../../allocation-service";

export function mapToCompleteAllocationPlan(
    portfolioDTO: PortfolioDTO,
    allocationPlanDTO: AllocationPlanDTO,
): CompleteAllocationPlan {

    const portfolio = mapToPortfolio(portfolioDTO);
    const allocationStructure = portfolio.allocationStructure;

    const allocationPlan = mapToAllocationPlan(allocationPlanDTO);

    const fractalHierarchicalPlan = mapToFractalHierarchicalAllocationPlan(
        allocationPlan,
        allocationStructure,
    );

    const topLevelKey = AllocationDomainService.getTopLevelHierarchyKeyFromAllocationPlan(
        allocationStructure,
    );

    return {
        allocationPlan,
        fractalHierarchicalPlan,
        topLevelKey,
    };
}

export function mapToCompleteAllocationPlans(
    portfolioDTO: PortfolioDTO,
    allocationPlanDTOs: AllocationPlanDTO[],
): CompleteAllocationPlan[] {
    return allocationPlanDTOs.map((allocationPlanDTO) =>
        mapToCompleteAllocationPlan(
            portfolioDTO,
            allocationPlanDTO,
        ));
}