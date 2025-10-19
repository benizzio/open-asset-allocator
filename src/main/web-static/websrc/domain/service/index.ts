import { AllocationDomainService } from "./allocation-service";
import {
    mapToAllocationPlan,
    mapToFractalHierarchicalAllocationPlan,
    mapToSerializableCompleteAllocationPlan,
} from "./mapping/allocation-plan-mapping";
import { mapToPortfolio } from "./mapping/portfolio-mapping";
import { AllocationPlanDTO, CompleteAllocationPlan, SerializableCompleteAllocationPlan } from "../allocation-plan";
import { PortfolioDTO } from "../portfolio";

export const DomainService = {
    allocation: AllocationDomainService,
    mapping: {

        mapToPortfolio,

        mapToCompleteAllocationPlan(
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
                portfolio,
                allocationPlan,
                fractalHierarchicalPlan,
                topLevelKey,
            };
        },

        mapToCompleteAllocationPlans(
            portfolioDTO: PortfolioDTO,
            allocationPlanDTOs: AllocationPlanDTO[],
        ): CompleteAllocationPlan[] {
            return allocationPlanDTOs.map((allocationPlanDTO) =>
                this.mapToCompleteAllocationPlan(
                    portfolioDTO,
                    allocationPlanDTO,
                ));
        },

        mapToSerializableCompleteAllocationPlans(
            portfolioDTO: PortfolioDTO,
            allocationPlanDTOs: AllocationPlanDTO[],
        ): SerializableCompleteAllocationPlan[] {
            const completeAllocationPlans = this.mapToCompleteAllocationPlans(
                portfolioDTO,
                allocationPlanDTOs,
            );
            return completeAllocationPlans.map((completeAllocationPlan: CompleteAllocationPlan) =>
                mapToSerializableCompleteAllocationPlan(completeAllocationPlan));
        },
    },
};