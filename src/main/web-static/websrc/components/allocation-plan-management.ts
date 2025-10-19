import DomUtils from "../infra/dom/dom-utils";
import { PortfolioDTO } from "../domain/portfolio";
import { AllocationPlanDTO } from "../domain/allocation-plan";
import { DomainService } from "../domain/service";
import { htmxInfra } from "../infra/htmx/htmx";

function mapToCompleteAllocationPlan(originalServerResponseJSON: string): string {
    const portfolioDTO = DomUtils.getContextDataFromRoot("#portfolio-context #portfolio") as PortfolioDTO;
    const allocationPlanDTO = JSON.parse(originalServerResponseJSON) as AllocationPlanDTO;
    const completeAllocationPlan = DomainService.mapping.mapToCompleteAllocationPlan(portfolioDTO, allocationPlanDTO);
    return JSON.stringify(completeAllocationPlan);
}

const allocationPlanManagement = {
    init() {
        htmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "mapToCompleteAllocationPlan",
            mapToCompleteAllocationPlan,
        );
    },
};

export default allocationPlanManagement;