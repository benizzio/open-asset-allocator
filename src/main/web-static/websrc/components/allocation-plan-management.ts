import PortfolioPage from "../pages/portfolio";
import { AllocationPlanDTO, SerializableFractalPlannedAllocation } from "../domain/allocation-plan";
import { DomainService } from "../domain/service";
import { Asset } from "../domain/asset";
import { htmxInfra } from "../infra/htmx/htmx";
import DomInfra from "../infra/dom";
import BigNumber from "bignumber.js";
import * as handlebars from "handlebars";

function mapToCompleteAllocationPlans(originalServerResponseJSON: string): string {

    const portfolioDTO = PortfolioPage.getContextPortfolio();
    const allocationPlanDTOs = JSON.parse(originalServerResponseJSON) as AllocationPlanDTO[];

    const completeAllocationPlanSet = DomainService.mapping.mapToSerializablePortfolioCompleteAllocationPlanSet(
        portfolioDTO,
        allocationPlanDTOs,
    );
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

function handleRemovePlannedAllocationRow(targetElement: HTMLElement) {
    const row = targetElement.closest("tr");
    const rowId = row.id;
    row.closest("table").querySelectorAll(`[data-parent-row-id=${ rowId }]`).forEach(row => row.remove());
    row.remove();
}


const allocationPlanManagement = {

    handlebarsAllocationPlanManagementRowTemplate: null as handlebars.TemplateDelegate,

    init() {
        htmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "mapToCompleteAllocationPlans",
            mapToCompleteAllocationPlans,
        );
    },

    handleHierarchicalIdLevelChange,
    handleRemovePlannedAllocationRow,

    // TODO clean
    handleAddPlannedAllocationRow(targetButton: HTMLButtonElement, parentRowIndex: number, parentLevelIndex: number) {

        const portfolio = PortfolioPage.getContextPortfolio();

        const newPlannedAllocationLevel = portfolio.allocationStructure.hierarchy[parentLevelIndex - 1];
        const portfolioAllocationHierarchy = portfolio.allocationStructure.hierarchy;
        const hierarchySize = portfolioAllocationHierarchy.length;

        const formElement = targetButton.closest("form");
        const parentRowElement = targetButton.closest("tr");

        let newPlannedAllocationAsset: Asset;

        if(newPlannedAllocationLevel.field === "assetTicker") {
            newPlannedAllocationAsset = { ticker: "" };
        }

        const newPlannedAllocation: SerializableFractalPlannedAllocation = {
            key: "",
            targetLevelKey: "",
            level: newPlannedAllocationLevel,
            allocation: {
                hierarchicalId: new Array(hierarchySize),
                cashReserve: false,
                sliceSizePercentage: new BigNumber(0),
                asset: newPlannedAllocationAsset,
            },
        };

        newPlannedAllocation.allocation.hierarchicalId[newPlannedAllocationLevel.index] = null;

        DomInfra.DomUtils.queryAllInDescendants(formElement, `[name^='details[${ parentRowIndex }][hierarchicalId]']`)
            .forEach((hierarchicalIdInput: HTMLInputElement) => {

                const hierarchicalIdValue = hierarchicalIdInput.value;
                const hierarchicalIdFieldName = hierarchicalIdInput.getAttribute("name");

                const hierarchicalIdIndexString = hierarchicalIdFieldName
                    .substring(
                        hierarchicalIdFieldName.indexOf("[hierarchicalId][") + "[hierarchicalId][".length,
                        hierarchicalIdFieldName.length - 1,
                    );
                const hierarchicalIdIndex = parseInt(hierarchicalIdIndexString, 10);

                newPlannedAllocation.allocation.hierarchicalId[hierarchicalIdIndex] = hierarchicalIdValue;
            });

        const lastRowIndexElement = DomInfra.DomUtils.queryFirstInDescendants(
            formElement,
            "[name='last-planned-allocation-row-index']",
        ) as HTMLInputElement;
        const lastRowIndex = parseInt(lastRowIndexElement.value, 10);
        const newRowIndex = lastRowIndex + 1;

        const newRowHtml = allocationPlanManagement.handlebarsAllocationPlanManagementRowTemplate({
            fractalPlannedAllocation: newPlannedAllocation,
            allocationIndex: newRowIndex,
            hierarchy: portfolioAllocationHierarchy,
            parentRowIndex,
        });

        parentRowElement.insertAdjacentHTML("afterend", newRowHtml);

        lastRowIndexElement.value = newRowIndex.toString();

        console.log(newPlannedAllocation);
        console.log(portfolioAllocationHierarchy);
    },
};

export default allocationPlanManagement;