import htmx from "htmx.org";
import api from "../api/api";
import { Asset } from "../domain/asset";
import BigNumber from "bignumber.js";
import { BootstrapClasses, BootstrapIconClasses } from "../infra/bootstrap/constants";
import { AfterRequestEventDetail, htmxInfra } from "../infra/htmx/htmx";
import { ObservationTimestamp } from "../domain/portfolio-allocation";
import router from "../infra/routing/router";

const PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER = "accordion-portfolio-history-management";
const PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX = "portfolio-history-management-form-";
const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-history-management-form-tbody-";

const ASSET_ACTION_BUTTON_IDENTITIES = {
    search: {
        classes: `${ BootstrapClasses.BUTTON_PRIMARY } btn-xs`,
        iconClasses: `${ BootstrapIconClasses.SEARCH }`,
    },
    reset: {
        classes: `${ BootstrapClasses.BUTTON_DANGER } btn-xs`,
        iconClasses: `${ BootstrapIconClasses.RESET }`,
    },
};

class FormRowAssetElements {

    assetIdInput: HTMLInputElement;
    assetTickerInput: HTMLInputElement;
    assetActionButton: HTMLButtonElement;
    newAssetTickerMessage: HTMLDivElement;
    assetNameInput: HTMLInputElement;

    constructor(formUniqueId: string, formRowIndex: number) {

        const formRowId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ formUniqueId }-row-${ formRowIndex }`;
        const formRow = window[formRowId] as HTMLElement;

        this.assetIdInput = formRow.querySelector("[name$='[assetId]']");
        this.assetTickerInput = formRow.querySelector("[name$='[assetTicker]']");
        this.assetActionButton = formRow.querySelector("[data-asset-action-button]");
        this.newAssetTickerMessage = formRow.querySelector("[data-new-asset-ticker-message]");
        this.assetNameInput = formRow.querySelector("[name$='[assetName]']");
    }

    isInSearchMode(): boolean {
        return this.assetActionButton.className === ASSET_ACTION_BUTTON_IDENTITIES.search.classes;
    }

    isInResetMode(): boolean {
        return this.assetActionButton.className === ASSET_ACTION_BUTTON_IDENTITIES.reset.classes;
    }

    switchAssetActionButtonIdentity(identity: typeof ASSET_ACTION_BUTTON_IDENTITIES.search) {
        this.assetActionButton.className = identity.classes;
        this.assetActionButton.innerHTML = `<span class="${ identity.iconClasses }"></span>`;
    }

    activateExistingAssetMode(asset: Asset) {

        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.reset);

        this.assetTickerInput.readOnly = true;
        this.assetTickerInput.value = asset.ticker;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = true;
        this.assetNameInput.value = asset.name;

        this.assetIdInput.value = asset.id.toString();

        this.newAssetTickerMessage.style.display = "none";
    }

    activateNewAssetMode() {

        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.reset);

        this.assetTickerInput.readOnly = false;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = true;

        this.newAssetTickerMessage.style.display = "";
    }

    resetToSearchMode() {

        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.search);

        this.assetTickerInput.value = "";
        this.assetTickerInput.focus();
        this.assetTickerInput.readOnly = false;

        this.assetNameInput.value = "";
        this.assetNameInput.style.display = "none";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = false;

        this.assetIdInput.value = "";

        this.newAssetTickerMessage.style.display = "none";
    }

    clearSearchFieldValidation() {
        this.assetTickerInput.setCustomValidity("");
        this.assetTickerInput.reportValidity();
    }

    validateSearchUniqueIdentifier(): string {

        const assetUniqueIdentifier = this.assetTickerInput.value.trim();

        if(!assetUniqueIdentifier) {
            this.assetTickerInput.setCustomValidity("Required for search");
            this.assetTickerInput.reportValidity();
        }

        return assetUniqueIdentifier;
    }

    handleAssetActionButtonClick() {

        if(this.isInSearchMode()) {

            this.clearSearchFieldValidation();
            const searchUniqueIdentifier = this.validateSearchUniqueIdentifier();

            if(searchUniqueIdentifier) {
                getAsset(this, searchUniqueIdentifier);
            }
        }
        else if(this.isInResetMode()) {
            this.resetToSearchMode();
        }
    }

    validateForPost() {
        if(this.isInSearchMode()) {
            this.assetTickerInput.setCustomValidity("Reference an existing asset or create a new one");
            this.assetTickerInput.reportValidity();
        }
    }
}

class FormRowValueElements {

    quantityInput: HTMLInputElement;
    marketPriceInput: HTMLInputElement;
    totalMarketValueInput: HTMLInputElement;

    constructor(formUniqueId: string, formRowIndex: number) {

        const formRowId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ formUniqueId }-row-${ formRowIndex }`;
        const formRow = window[formRowId] as HTMLElement;

        this.quantityInput = formRow.querySelector("[name$='[assetQuantity]']");
        this.marketPriceInput = formRow.querySelector("[name$='[assetMarketPrice]']");
        this.totalMarketValueInput = formRow.querySelector("[name$='[totalMarketValue]']");
    }

    handleInputQuantityOrMarketPrice() {

        const quantity = this.quantityInput.value || 0;
        const marketPrice = this.marketPriceInput.value || 0;

        if(quantity && marketPrice) {
            const totalMarketValue = new BigNumber(quantity)
                .times(marketPrice)
                .decimalPlaces(0, BigNumber.ROUND_HALF_UP);
            this.totalMarketValueInput.value = totalMarketValue.toString();
        }
    }

    handleInputTotalMarketValue() {
        this.quantityInput.value = "";
        this.marketPriceInput.value = "";
    }
}

function focusOnNewLine(newRow: HTMLElement) {

    const firstFieldOfNewRow: HTMLInputElement = newRow.querySelector("input[type='text']");

    if(firstFieldOfNewRow) {
        firstFieldOfNewRow.scrollIntoView({ behavior: "smooth", block: "center" });
        firstFieldOfNewRow.focus();
    }
}

function getNextPortfolioHistoryManagementIndex(tbody: HTMLElement): number {
    const rows = tbody.querySelectorAll("tr");
    const lastRow = rows[rows.length - 1] as HTMLElement;
    const lastRowId = lastRow?.id;
    const lastRowIdIndex = lastRowId?.split("-").pop();
    return lastRowIdIndex ? parseInt(lastRowIdIndex, 10) + 1 : 0;
}

function getAsset(rowAssetElements: FormRowAssetElements, searchUniqueIdentifier: string) {

    api.getAsset(searchUniqueIdentifier)
        .then(responseBody => {

            if(api.isAPIErrorResponse(responseBody)) {
                if(responseBody.errorMessage === "Data not found") {
                    rowAssetElements.activateNewAssetMode();
                }
                else {
                    // TODO add toast for errors
                    console.error("Error fetching asset:", responseBody.errorMessage);
                }
                return;
            }

            rowAssetElements.activateExistingAssetMode(responseBody as Asset);
        })
        .catch(error => {
            // TODO add toast for errors
            console.error("Error fetching asset:", error);
        });
}

function modifyObservationsResponse(originalServerResponseJSON: string): string {

    const originalServerResponse = JSON.parse(originalServerResponseJSON) as ObservationTimestamp[];

    const modifiedServerResponse: ObservationTimestamp[] = [
        {
            id: 0,
            timeTag: "New Observation *",
        },
        ...originalServerResponse,
    ];

    return JSON.stringify(modifiedServerResponse);
}

function propagateRefreshDataAfterPost(observationTimestampId: number) {

    if(observationTimestampId !== 0) {

        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = window[formId] as HTMLFormElement;
        form.reset();

        const observationManagementTriggerElement =
            window[`portfolio-history-management-trigger-${ observationTimestampId }`] as HTMLElement;
        htmx.trigger(observationManagementTriggerElement, "reload-portfolio-history-management-observation");
    }
    else {
        const portfolioHistoryManagementContainerElement =
            window[PORTFOLIO_ALLOCATION_MANAGEMENT_PARENT_CONTAINER] as HTMLElement;
        htmx.trigger(portfolioHistoryManagementContainerElement, "reload-portfolio-history-management");
    }

    const portfolioHistoryViewContainerElement = window["accordion-portfolio-history"];
    htmx.trigger(portfolioHistoryViewContainerElement, "reload-portfolio-history");

    loadPortfolioHistoryDatalists();
}

function loadClassesDatalist() {
    const datalist = window["datalist-classes"];
    htmx.trigger(datalist, "load-classes");
}

function loadAssetsDatalist() {
    const datalist = window["datalist-assets"];
    htmx.trigger(datalist, "load-assets");
}

function loadPortfolioHistoryDatalists() {
    loadClassesDatalist();
    loadAssetsDatalist();
}

const portfolioHistoryManagement = {

    handlebarPortfolioHistoryManagementRowTemplate: null,
    handlebarPortfolioHistoryManagementContainerTemplate: null,

    init() {
        htmxInfra.htmxTransformResponse.registerTransformResponseFunction(
            "addObservationZero",
            modifyObservationsResponse,
        );
    },

    addPortfolioHistoryManagementRow(observationTimestampId: number) {

        const tbodyId = PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX + observationTimestampId;
        const tbody: HTMLElement = window[tbodyId];
        const nextIndex = getNextPortfolioHistoryManagementIndex(tbody);

        const newRowHtml = this.handlebarPortfolioHistoryManagementRowTemplate({
            allocationIndex: nextIndex,
            observationTimestampId: observationTimestampId,
        });
        tbody.insertAdjacentHTML("beforeend", newRowHtml);

        const newRow = tbody.lastElementChild as HTMLElement;

        focusOnNewLine(newRow);
        // Process the newly added row with htmx to enable bindings
        htmx.process(newRow);
    },

    assetActionButtonClickHandler(formRowIndex: number, formUniqueId: string) {
        const rowAssetElements = new FormRowAssetElements(formUniqueId, formRowIndex);
        rowAssetElements.handleAssetActionButtonClick();
    },

    validateAssetElementsForPost(formRowIndex: number, formUniqueId: string) {
        const rowAssetElements = new FormRowAssetElements(formUniqueId, formRowIndex);
        rowAssetElements.validateForPost();
    },

    handleInputQuantityOrMarketPrice(formRowIndex: number, formUniqueId: string) {
        const rowValueElements = new FormRowValueElements(formUniqueId, formRowIndex);
        rowValueElements.handleInputQuantityOrMarketPrice();
    },

    handleInputTotalMarketValue(formRowIndex: number, formUniqueId: string) {
        const rowValueElements = new FormRowValueElements(formUniqueId, formRowIndex);
        rowValueElements.handleInputTotalMarketValue();
    },

    handleInputObservationTimeTag(newTimeTagInput: HTMLInputElement, observationTimestampId: number) {
        const formId = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }${ observationTimestampId }`;
        const form = window[formId] as HTMLFormElement;
        const observationTimeTagInput = form.elements.namedItem("observationTimestamp.timeTag") as HTMLInputElement;
        observationTimeTagInput.value = newTimeTagInput.value;
    },

    handleAfterPostObservationHistory(event: CustomEvent, observationTimestampId: number) {

        const eventDetail = event.detail as AfterRequestEventDetail;

        if(!eventDetail.successful) {
            return;
        }

        propagateRefreshDataAfterPost(observationTimestampId);
    },

    navigateToPortfolioAllocationViewing() {
        const globalPortfolioIdField = document.querySelector("[name='portfolioId']") as HTMLInputElement;
        const portfolioId = globalPortfolioIdField.value;
        router.navigateTo(`/portfolio/${ portfolioId }/history`);
    },
};

export default portfolioHistoryManagement;
