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
    console.log(completeAllocationPlanSet);
    return JSON.stringify(completeAllocationPlanSet);
}

function handleHierarchicalIdLevelChange(targetElement: HTMLInputElement) {

    const ancestorTable = targetElement.closest("form");

    if(!ancestorTable) {
        return;
    }

    const targetElementName = targetElement.getAttribute("name");

    const fieldsToUpdate =
        ancestorTable.querySelectorAll<HTMLInputElement>(`[data-bind-to-name$='${ targetElementName }']`);

    fieldsToUpdate.forEach((field) => {
        field.value = targetElement.value;
    });

    const spansToUpdat =
        ancestorTable.querySelectorAll<HTMLSpanElement>(`[data-label-for-name='${ targetElementName }']`);

    spansToUpdat.forEach((span) => {
        span.textContent = targetElement.value;
    });
}


const allocationPlanManagement = {
    init() {
        htmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "mapToCompleteAllocationPlans",
            mapToCompleteAllocationPlans,
        );
    },

    handleHierarchicalIdLevelChange,
};

export default allocationPlanManagement;