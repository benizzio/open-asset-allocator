import htmx from "htmx.org";
import api from "../api/api";
import { Asset } from "../domain/asset";

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

class RowAssetElements {

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

    handleAssetActionButtonClick() {
        if(this.isInSearchMode()) {
            getAsset(this);
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

function getAsset(rowAssetElements: RowAssetElements) {

    const { assetTickerInput } = rowAssetElements;

    assetTickerInput.setCustomValidity("");
    assetTickerInput.reportValidity();

    const assetTicker = assetTickerInput.value;

    if(!assetTicker) {
        assetTickerInput.setCustomValidity("Required for search");
        assetTickerInput.reportValidity();
        return;
    }

    api.getAsset(assetTicker)
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
        const rowAssetElements = new RowAssetElements(formUniqueId, formRowIndex);
        rowAssetElements.handleAssetActionButtonClick();
    },
    validateAssetElementsForPost(formRowIndex: number, formUniqueId: string) {
        const rowAssetElements = new RowAssetElements(formUniqueId, formRowIndex);
        rowAssetElements.validateForPost();
    },
};

export default portfolioHistoryManagement;