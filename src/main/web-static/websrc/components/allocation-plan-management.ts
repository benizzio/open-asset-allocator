import PortfolioPage from "../pages/portfolio";
import { AllocationPlanDTO, SerializableFractalPlannedAllocation } from "../domain/allocation-plan";
import { DomainService } from "../domain/service";
import { Asset } from "../domain/asset";
import { htmxInfra } from "../infra/htmx/htmx";
import DomInfra from "../infra/dom";
import BigNumber from "bignumber.js";
import * as handlebars from "handlebars";
import { isNullish, toInt } from "../utils/lang";
import { Portfolio } from "../domain/portfolio";
import { AllocationHierarchyLevel } from "../domain/allocation";
import AssetComposedColumnsInput from "./asset-composed-columns-input";
import htmx from "htmx.org";

type AllocationPlanningHierarchicalFormEntry = {
    occurences: number;
    inputFields: HTMLInputElement[];
};

const FORM_LAST_ROW_INDEX_INPUT_NAME = "last-planned-allocation-row-index";
const FORM_FIELD_DEPENDENT_ATTRIBUTE = "data-bind-to-name";

const ALLOCATION_HIERARCHY_LEVEL_MANAGING_FIELD_TEMP_PROPERTY_NAME = "currentManagingFieldName";
const ALLOCATION_PLAN_MANAGEMENT_FORM_PREFIX = "allocation-plan-management-form-";

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

    const spansToUpdate =
        ancestorTable.querySelectorAll<HTMLSpanElement>(`[data-label-for-name='${ targetElementName }']`);

    spansToUpdate.forEach((span) => {
        span.textContent = targetElement.value;
    });
}

function handleRemovePlannedAllocationRow(targetElement: HTMLElement) {
    const row = targetElement.closest("tr");
    const rowId = row.id;
    row.closest("table").querySelectorAll(`[data-parent-row-id=${ rowId }]`).forEach(row => row.remove());
    row.remove();
}

function createNewPlannedAllocation(
    portfolio: Portfolio,
    parentHierarchyLevelIndex: number,
): SerializableFractalPlannedAllocation {

    const portfolioAllocationHierarchy = portfolio.allocationStructure.hierarchy;
    const newPlannedAllocationLevel = portfolioAllocationHierarchy[parentHierarchyLevelIndex - 1];
    const hierarchySize = portfolioAllocationHierarchy.length;

    let newPlannedAllocationAsset: Asset;

    if(newPlannedAllocationLevel.field === "assetTicker") {
        newPlannedAllocationAsset = { ticker: "" };
    }

    const subLevel = newPlannedAllocationLevel.index > 0
        ? portfolioAllocationHierarchy[newPlannedAllocationLevel.index - 1]
        : null;

    const newPlannedAllocation: SerializableFractalPlannedAllocation = {
        key: "",
        targetLevelKey: "",
        level: newPlannedAllocationLevel,
        subLevel: subLevel,
        allocation: {
            hierarchicalId: new Array(hierarchySize),
            cashReserve: false,
            sliceSizePercentage: new BigNumber(0),
            asset: newPlannedAllocationAsset,
        },
    };

    newPlannedAllocation.allocation.hierarchicalId[newPlannedAllocationLevel.index] = null;

    return newPlannedAllocation;
}

function setHierarchicalIdFromParentRow(
    newPlannedAllocation: SerializableFractalPlannedAllocation,
    formElement: HTMLFormElement,
    parentRowIndex: number,
    portfolioAllocationHierarchy: AllocationHierarchyLevel[],
) {

    DomInfra.DomUtils.queryAllInDescendants(formElement, `[name^='details[${ parentRowIndex }][hierarchicalId]']`)
        .forEach((hierarchicalIdInput: HTMLInputElement) => {

            const hierarchicalIdValue = hierarchicalIdInput.value;
            const hierarchicalIdFieldName = hierarchicalIdInput.getAttribute("name");

            const hierarchicalIdIndexString = hierarchicalIdFieldName
                .substring(
                    hierarchicalIdFieldName.indexOf("[hierarchicalId][") + "[hierarchicalId][".length,
                    hierarchicalIdFieldName.length - 1,
                );
            const hierarchicalIdIndex = toInt(hierarchicalIdIndexString);

            newPlannedAllocation.allocation.hierarchicalId[hierarchicalIdIndex] = hierarchicalIdValue;

            const bindedFieldName = hierarchicalIdInput.getAttribute(FORM_FIELD_DEPENDENT_ATTRIBUTE);
            const allocationHierarchyLevelForInput = portfolioAllocationHierarchy[hierarchicalIdIndex];

            allocationHierarchyLevelForInput[ALLOCATION_HIERARCHY_LEVEL_MANAGING_FIELD_TEMP_PROPERTY_NAME] =
                bindedFieldName ?? hierarchicalIdFieldName;
        });
}

function addPlannedAllocationRow(
    allocationPlanId: number,
    newPlannedAllocation: SerializableFractalPlannedAllocation,
    formElement: HTMLFormElement,
    portfolioAllocationHierarchy: AllocationHierarchyLevel[],
    parentRowIndex?: number,
    parentRowElement?: HTMLTableRowElement,
) {

    const lastRowIndexElement = DomInfra.DomUtils.queryFirstInDescendants(
        formElement,
        `[name="${ FORM_LAST_ROW_INDEX_INPUT_NAME }"]`,
    ) as HTMLInputElement;

    const lastRowIndex = toInt(lastRowIndexElement.value);
    const newRowIndex = lastRowIndex + 1;

    const newRowHtml = allocationPlanManagement.handlebarsAllocationPlanManagementRowTemplate({
        allocationPlanId,
        fractalPlannedAllocation: newPlannedAllocation,
        allocationIndex: newRowIndex,
        hierarchy: portfolioAllocationHierarchy,
        parentRowIndex,
    });

    let newRow: HTMLElement;

    if(parentRowElement) {
        parentRowElement.insertAdjacentHTML("afterend", newRowHtml);
        newRow = parentRowElement.nextElementSibling as HTMLElement;
    }
    else {
        const tbody = formElement.querySelector("tbody");
        tbody.insertAdjacentHTML("beforeend", newRowHtml);
        newRow = tbody.lastElementChild as HTMLElement;
    }

    lastRowIndexElement.value = newRowIndex.toString();

    htmx.process(newRow);
    DomInfra.bindDescendants(newRow);
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

    handleAddPlannedAllocationRow(
        allocationPlanId: number,
        targetButton: HTMLButtonElement,
        parentHierarchyLevelIndex: number,
        parentRowIndex?: number,
    ) {

        const portfolio = PortfolioPage.getContextPortfolio();
        const portfolioAllocationHierarchy = portfolio.allocationStructure.hierarchy;

        const newPlannedAllocation = createNewPlannedAllocation(
            portfolio,
            parentHierarchyLevelIndex,
        );

        const formElement = targetButton.closest("form");

        let parentRowElement: HTMLTableRowElement;

        if(!isNullish(parentRowIndex)) {

            setHierarchicalIdFromParentRow(
                newPlannedAllocation,
                formElement,
                parentRowIndex,
                portfolioAllocationHierarchy,
            );

            parentRowElement = targetButton.closest("tr");
        }

        addPlannedAllocationRow(
            allocationPlanId,
            newPlannedAllocation,
            formElement,
            portfolioAllocationHierarchy,
            parentRowIndex,
            parentRowElement,
        );
    },

    assetActionButtonClickHandler(allocationPlanId: number, formRowIndex: number) {

        const formRowId = `${ ALLOCATION_PLAN_MANAGEMENT_FORM_PREFIX }${ allocationPlanId }-row-${ formRowIndex }`;
        const assetIdHiddenFieldName = `details[${ formRowIndex }][asset][id]`;
        const assetTickerFieldName = `details[${ formRowIndex }][asset][ticker]`;
        const assetNameFieldName = `details[${ formRowIndex }][asset][name]`;

        AssetComposedColumnsInput.assetActionButtonClickHandler(
            formRowId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
    },

    validateAssetElementsForPost(allocationPlanId: string, formRowIndex: number) {

        const formRowId = `${ ALLOCATION_PLAN_MANAGEMENT_FORM_PREFIX }${ allocationPlanId }-row-${ formRowIndex }`;
        const assetIdHiddenFieldName = `details[${ formRowIndex }][asset][id]`;
        const assetTickerFieldName = `details[${ formRowIndex }][asset][ticker]`;
        const assetNameFieldName = `details[${ formRowIndex }][asset][name]`;

        AssetComposedColumnsInput.validateAssetElementsForPost(
            formRowId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
    },

    //TODO clean code
    preSubmitHandler(form: HTMLFormElement, hierarchySize: number) {

        DomInfra.DomUtils.queryAllInDescendants(
            form,
            "input[name$='[asset][ticker]']",
        ).forEach((assetTickerInput: HTMLInputElement) => {

            const assetTickerValue = assetTickerInput.value;

            const parentTr = assetTickerInput.closest("tr");
            const allocationIndexString = parentTr.getAttribute("data-allocation-index");

            const assetIdInput = form.elements.namedItem(
                `details[${ allocationIndexString }][hierarchicalId][0]`,
            ) as HTMLInputElement;
            assetIdInput.value = assetTickerValue;
        });

        const hirarchicalFormEntriesPerHierarchicalKey = new Map<string, AllocationPlanningHierarchicalFormEntry>();

        DomInfra.DomUtils.queryAllInDescendants(
            form,
            "tr",
        ).forEach((tableRow: HTMLTableRowElement) => {

            const lastHierarchyLevelIndex = hierarchySize - 1;
            const hierarchicalFieldsForKey: HTMLInputElement[] = [];
            let fieldHierarchicalId = "";

            for(let hierarchyLevelIndex = lastHierarchyLevelIndex; hierarchyLevelIndex >= 0; hierarchyLevelIndex--) {

                const allocationHierarchyLevelFieldNameSuffix = `[hierarchicalId][${ hierarchyLevelIndex }]`;

                const hierarchicalField = DomInfra.DomUtils.queryFirstInDescendants(
                    tableRow,
                    `input[name$='${ allocationHierarchyLevelFieldNameSuffix }']`,
                ) as HTMLInputElement;

                if(hierarchicalField) {

                    if(!hierarchicalField.value.trim()) {
                        return;
                    }

                    const separator = hierarchyLevelIndex < lastHierarchyLevelIndex ? "|" : "";
                    fieldHierarchicalId += `${ separator }${ hierarchicalField.value }`;

                    // if the field is not on the last level and its not a hidden input, add to map
                    if(hierarchyLevelIndex > 0 && hierarchicalField.type !== "hidden") {
                        hierarchicalFieldsForKey.push(hierarchicalField);
                    }
                    else if(hierarchyLevelIndex == 0) {
                        const assertSearchField = DomInfra.DomUtils.queryFirstInDescendants(
                            tableRow,
                            "input[name$='[asset][ticker]']",
                        ) as HTMLInputElement;

                        if(assertSearchField) {
                            hierarchicalFieldsForKey.push(assertSearchField);
                        }
                    }
                }
            }

            if(!fieldHierarchicalId) {
                return;
            }

            let formEntry = hirarchicalFormEntriesPerHierarchicalKey.get(fieldHierarchicalId);

            if(!formEntry) {
                formEntry = {
                    inputFields: [],
                    occurences: 0,
                };
                hirarchicalFormEntriesPerHierarchicalKey.set(fieldHierarchicalId, formEntry);
            }

            formEntry.occurences++;
            formEntry.inputFields.push(...hierarchicalFieldsForKey);
        });

        let containsDuplicates = false;

        hirarchicalFormEntriesPerHierarchicalKey.forEach((formEntriesForKey, key) => {

            if(formEntriesForKey.occurences > 1) {

                formEntriesForKey.inputFields.forEach((field) => {

                    if(field.name.endsWith("[asset][ticker]")) {
                        AssetComposedColumnsInput.invalidateSelectedAsset(
                            field,
                            "Duplicate hierarchical id in allocation plan",
                        );
                    }
                    else {
                        field.setCustomValidity("Duplicate hierarchical id in allocation plan");
                        field.reportValidity();
                    }
                });
                containsDuplicates = true;
            }
            else if(formEntriesForKey.inputFields.length > 0) {
                formEntriesForKey.inputFields[0].setCustomValidity("");
                formEntriesForKey.inputFields[0].reportValidity();
            }
        });

        if(containsDuplicates) {
            return;
        }

        form.requestSubmit();
    },

    postRequestHandler() {
        //TODO
    },
};

export default allocationPlanManagement;