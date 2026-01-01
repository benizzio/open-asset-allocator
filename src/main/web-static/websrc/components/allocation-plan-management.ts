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

type FormRowHierarchicalStructure = {
    formRowHierarchicalFields: HTMLInputElement[];
    formRowHierarchicalId: string
};

const FORM_LAST_ROW_INDEX_INPUT_NAME = "last-planned-allocation-row-index";
const FORM_FIELD_DEPENDENT_ATTRIBUTE = "data-bind-to-name";

const ALLOCATION_HIERARCHY_LEVEL_MANAGING_FIELD_TEMP_PROPERTY_NAME = "currentManagingFieldName";
const ALLOCATION_PLAN_MANAGEMENT_FORM_PREFIX = "allocation-plan-management-form-";

const ASSET_TICKER_FIELD_NAME_SUFFIX = "[asset][ticker]";

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

function copyAssetTickersToHierarchicalIdFields(form: HTMLFormElement) {

    DomInfra.DomUtils.queryAllInDescendants(
        form,
        `input[name$='${ ASSET_TICKER_FIELD_NAME_SUFFIX }']`,
    ).forEach((assetTickerInput: HTMLInputElement) => {

        const assetTickerValue = assetTickerInput.value;

        const parentTr = assetTickerInput.closest("tr");
        const allocationIndexString = parentTr.getAttribute("data-allocation-index");

        const assetIdInput = form.elements.namedItem(
            `details[${ allocationIndexString }][hierarchicalId][0]`,
        ) as HTMLInputElement;
        assetIdInput.value = assetTickerValue;
    });
}


function addHierarchicalFieldMapping(
    formTableRow: HTMLTableRowElement,
    hierarchyLevelIndex: number,
    hierarchicalField: HTMLInputElement,
    formRowHierarchicalStructure: FormRowHierarchicalStructure,
) {

    if(hierarchyLevelIndex > 0 && hierarchicalField.type !== "hidden") {
        formRowHierarchicalStructure.formRowHierarchicalFields.push(hierarchicalField);
    }
    else if(hierarchyLevelIndex == 0) {

        const assertSearchField = DomInfra.DomUtils.queryFirstInDescendants(
            formTableRow,
            `input[name$='${ ASSET_TICKER_FIELD_NAME_SUFFIX }']`,
        ) as HTMLInputElement;

        if(assertSearchField) {
            formRowHierarchicalStructure.formRowHierarchicalFields.push(assertSearchField);
        }
    }
}

function getHierarchicalFieldForValidation(
    formTableRow: HTMLTableRowElement,
    hierarchyLevelIndex: number,
    hierarchicalField: HTMLInputElement,
) {

    if(hierarchyLevelIndex > 0 && hierarchicalField.type !== "hidden") {
        return hierarchicalField;
    }
    else if(hierarchyLevelIndex == 0) {
        return DomInfra.DomUtils.queryFirstInDescendants(
            formTableRow,
            `input[name$='${ ASSET_TICKER_FIELD_NAME_SUFFIX }']`,
        ) as HTMLInputElement;
    }
}

function mapFormRowHierarchicalStructure(formTableRow: HTMLTableRowElement, hierarchySize: number) {

    const lastHierarchyLevelIndex = hierarchySize - 1;

    const formRowHierarchicalStructure = {
        formRowHierarchicalFields: [] as HTMLInputElement[],
        formRowHierarchicalId: "",
    };

    for(let hierarchyLevelIndex = lastHierarchyLevelIndex; hierarchyLevelIndex >= 0; hierarchyLevelIndex--) {

        const allocationHierarchyLevelFieldNameSuffix = `[hierarchicalId][${ hierarchyLevelIndex }]`;

        const hierarchicalField = DomInfra.DomUtils.queryFirstInDescendants(
            formTableRow,
            `input[name$='${ allocationHierarchyLevelFieldNameSuffix }']`,
        ) as HTMLInputElement;

        if(!hierarchicalField) {
            continue;
        }

        if(!hierarchicalField.value.trim()) {
            // hierarchical fields are required, let the form validation handle this
            return undefined;
        }

        const separator = hierarchyLevelIndex < lastHierarchyLevelIndex ? "|" : "";
        formRowHierarchicalStructure.formRowHierarchicalId += `${ separator }${ hierarchicalField.value }`;

        const field = getHierarchicalFieldForValidation(formTableRow, hierarchyLevelIndex, hierarchicalField);

        if(field) {
            formRowHierarchicalStructure.formRowHierarchicalFields.push(field);
        }
    }

    return formRowHierarchicalStructure.formRowHierarchicalId.trim() ? formRowHierarchicalStructure : undefined;
}

function mapPlannedAllocationFormEntriesPerHierarchicalKey(form: HTMLFormElement, hierarchySize: number) {

    const formEntriesPerHierarchicalKey = new Map<string, AllocationPlanningHierarchicalFormEntry>();

    DomInfra.DomUtils.queryAllInDescendants(
        form,
        "tr",
    ).forEach((formTableRow: HTMLTableRowElement) => {

        const formRowHirarchicalStructure = mapFormRowHierarchicalStructure(
            formTableRow,
            hierarchySize,
        );

        if(!formRowHirarchicalStructure) {
            return;
        }

        const { formRowHierarchicalId, formRowHierarchicalFields } = formRowHirarchicalStructure;

        let formEntry = formEntriesPerHierarchicalKey.get(formRowHierarchicalId);

        if(!formEntry) {
            formEntry = {
                inputFields: [],
                occurences: 0,
            };
            formEntriesPerHierarchicalKey.set(formRowHierarchicalId, formEntry);
        }

        formEntry.occurences++;
        formEntry.inputFields.push(...formRowHierarchicalFields);
    });

    return formEntriesPerHierarchicalKey;
}

function validateDuplicateEntries(
    hierarchicalFormEntriesPerHierarchicalKey: Map<string, AllocationPlanningHierarchicalFormEntry>,
) {

    let containsDuplicates = false;

    hierarchicalFormEntriesPerHierarchicalKey.forEach((formEntriesForKey) => {

        if(formEntriesForKey.occurences > 1) {

            formEntriesForKey.inputFields.forEach((field) => {

                if(field.name.endsWith(ASSET_TICKER_FIELD_NAME_SUFFIX)) {
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
    return containsDuplicates;
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
        const assetTickerFieldName = `details[${ formRowIndex }]${ ASSET_TICKER_FIELD_NAME_SUFFIX }`;
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
        const assetTickerFieldName = `details[${ formRowIndex }]${ ASSET_TICKER_FIELD_NAME_SUFFIX }`;
        const assetNameFieldName = `details[${ formRowIndex }][asset][name]`;

        AssetComposedColumnsInput.validateAssetElementsForPost(
            formRowId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
    },

    preSubmitHandler(form: HTMLFormElement, hierarchySize: number) {

        copyAssetTickersToHierarchicalIdFields(form);

        const hierarchicalFormEntriesPerHierarchicalKey = mapPlannedAllocationFormEntriesPerHierarchicalKey(
            form,
            hierarchySize,
        );

        const containsDuplicateEntries = validateDuplicateEntries(hierarchicalFormEntriesPerHierarchicalKey);

        if(containsDuplicateEntries) {
            return;
        }

        form.requestSubmit();
    },

    postRequestHandler() {
        //TODO
    },
};

export default allocationPlanManagement;