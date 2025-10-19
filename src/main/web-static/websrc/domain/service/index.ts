import { AllocationDomainService } from "./allocation-service";
import { mapToAllocationPlan, mapToAllocationPlanFractalHierarchy } from "./mapping/allocation-plan-mapping";
import { mapToPortfolio } from "./mapping/portfolio-mapping";
import { AllocationPlanDTO, CompleteAllocationPlan } from "../allocation-plan";
import { PortfolioDTO } from "../portfolio";

export const DomainService = {
    allocation: AllocationDomainService,
    mapping: {
        mapToAllocationPlan,
        mapToAllocationPlanFractalHierarchy,
        mapToPortfolio,
        mapToCompleteAllocationPlan(
            portfolioDTO: PortfolioDTO,
            allocationPlanDTO: AllocationPlanDTO,
        ): CompleteAllocationPlan {

            const portfolio = DomainService.mapping.mapToPortfolio(portfolioDTO);
            const allocationStructure = portfolio.allocationStructure;

            const allocationPlan = DomainService.mapping.mapToAllocationPlan(allocationPlanDTO);

            const fractalHierarchy = DomainService.mapping.mapToAllocationPlanFractalHierarchy(
                allocationPlan,
                allocationStructure,
            );

            const topLevelKey = DomainService.allocation.getTopLevelHierarchyKeyFromAllocationPlan(
                allocationStructure,
            );

            return {
                portfolio,
                allocationPlan,
                fractalHierarchy,
                topLevelKey,
            };
        },
    },
};