import DomUtils from "../infra/dom/dom-utils";
import { PortfolioDTO } from "../domain/portfolio";
import { AllocationPlanDTO } from "../domain/allocation-plan";
import { DomainService } from "../domain/service";
import { htmxInfra } from "../infra/htmx/htmx";

function mapToCompleteAllocationPlans(originalServerResponseJSON: string): string {

    const portfolioDTO = DomUtils.getContextDataFromRoot("#portfolio-context #portfolio") as PortfolioDTO;
    const allocationPlanDTOs = JSON.parse(originalServerResponseJSON) as AllocationPlanDTO[];

    const completeAllocationPlanSet = DomainService.mapping.mapToSerializablePortfolioCompleteAllocationPlanSet(
        portfolioDTO,
        allocationPlanDTOs,
    );
    return JSON.stringify(completeAllocationPlanSet);
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