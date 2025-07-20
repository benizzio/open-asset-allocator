import htmx from "htmx.org";
import api from "../api/api";
import { Asset } from "../domain/asset";
import BigNumber from "bignumber.js";

const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-history-management-form-tbody-";

const ASSET_ACTION_BUTTON_IDENTITIES = {
    search: {
        classes: "btn btn-primary btn-xs",
        iconClasses: "bi bi-search",
    },
    reset: {
        classes: "btn btn-danger btn-xs",
        iconClasses: "bi bi-x-circle",
    },
};

class FormRowAssetElements {

    assetIdInput: HTMLInputElement;
    assetTickerInput: HTMLInputElement;
    assetActionButton: HTMLButtonElement;
    newAssetTickerMessage: HTMLDivElement;
    assetNameInput: HTMLInputElement;

    constructor(formUniqueId: string, formRowIndex: number) {
        const formRow =
            window[`portfolio-history-management-form-${ formUniqueId }-row-${ formRowIndex }`] as HTMLElement;
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
        const formRow =
            window[`portfolio-history-management-form-${ formUniqueId }-row-${ formRowIndex }`] as HTMLElement;
        this.quantityInput = formRow.querySelector("[name$='[assetQuantity]']");
        this.marketPriceInput = formRow.querySelector("[name$='[assetMarketPrice]']");
        this.totalMarketValueInput = formRow.querySelector("[name$='[totalMarketValue]']");
    }

    handleInputQuantityOrMarketPrice() {

        const quantity = this.quantityInput.value || 0;
        const marketPrice = this.marketPriceInput.value || 0;

        if(quantity && marketPrice) {
            const totalMarketValue = new BigNumber(quantity).times(marketPrice);
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
    return rows.length;
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

const portfolioHistoryManagement = {

    handlebarPortfolioHistoryManagementRowTemplate: null,

    addPortfolioHistoryManagementRow(observationTimestampId: string) {

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
};

export default portfolioHistoryManagement;