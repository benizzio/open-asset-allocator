import DomUtils from "../infra/dom/dom-utils";
import { PortfolioDTO } from "../domain/portfolio";
import { AllocationPlanDTO } from "../domain/allocation-plan";
import { DomainService } from "../domain/service";
import { htmxInfra } from "../infra/htmx/htmx";

function mapToCompleteAllocationPlans(originalServerResponseJSON: string): string {

    const portfolioDTO = DomUtils.getContextDataFromRoot("#portfolio-context #portfolio") as PortfolioDTO;
    const allocationPlanDTOs = JSON.parse(originalServerResponseJSON) as AllocationPlanDTO[];

    const completeAllocationPlan = DomainService.mapping.mapToSerializableCompleteAllocationPlans(
        portfolioDTO,
        allocationPlanDTOs,
    );
    return JSON.stringify(completeAllocationPlan);
}

const allocationPlanManagement = {
    init() {
        htmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "mapToCompleteAllocationPlans",
            mapToCompleteAllocationPlans,
        );
    },
};

export default allocationPlanManagement;