import htmx from "htmx.org";
import api from "../api/api";

type RowAssetElements = {
    assetIdInput: HTMLInputElement,
    assetTickerInput: HTMLInputElement,
    assetActionButton: HTMLButtonElement,
    assetTickerMessage: HTMLDivElement,
    assetNameInput: HTMLInputElement,
};

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
    isInSearchMode(actionButton: HTMLButtonElement): boolean {
        return actionButton.className === ASSET_ACTION_BUTTON_IDENTITIES.search.classes;
    },
    isInResetMode(actionButton: HTMLButtonElement): boolean {
        return actionButton.className === ASSET_ACTION_BUTTON_IDENTITIES.reset.classes;
    },
};

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

    const { assetTickerInput, assetNameInput, assetIdInput, assetTickerMessage } = rowAssetElements;

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
                    switchAssetActionButtonIdentity(rowAssetElements, ASSET_ACTION_BUTTON_IDENTITIES.reset);
                    assetTickerInput.readOnly = false;
                    assetNameInput.style.display = "";
                    assetNameInput.readOnly = false;
                    assetNameInput.required = true;
                    assetTickerMessage.style.display = "";
                }

                return;
            }

            switchAssetActionButtonIdentity(rowAssetElements, ASSET_ACTION_BUTTON_IDENTITIES.reset);
            assetTickerInput.readOnly = true;
            assetTickerInput.value = responseBody.ticker;
            assetNameInput.style.display = "";
            assetNameInput.readOnly = true;
            assetNameInput.value = responseBody.name;
            assetIdInput.value = responseBody.id.toString();
            assetTickerMessage.style.display = "none";
        })
        .catch(error => {
            // TODO add toast for errors
            console.error("Error fetching asset:", error);
        });
}

function switchAssetActionButtonIdentity(
    rowAssetElements: RowAssetElements,
    identity: typeof ASSET_ACTION_BUTTON_IDENTITIES.search,
) {
    const { assetActionButton } = rowAssetElements;
    assetActionButton.className = identity.classes;
    assetActionButton.innerHTML = `<span class="${ identity.iconClasses }"></span>`;
}

function resetAsset(rowAssetElements: RowAssetElements) {
    const { assetTickerInput, assetNameInput, assetIdInput, assetTickerMessage } = rowAssetElements;
    assetTickerInput.value = "";
    assetTickerInput.focus();
    assetTickerInput.readOnly = false;
    assetNameInput.value = "";
    assetNameInput.style.display = "none";
    assetNameInput.readOnly = false;
    assetNameInput.required = false;
    assetIdInput.value = "";
    assetTickerMessage.style.display = "none";
    switchAssetActionButtonIdentity(rowAssetElements, ASSET_ACTION_BUTTON_IDENTITIES.search);
}

function getFormRowAssetElements(formUniqueId: string, formRowIndex: number): RowAssetElements {

    const formRow =
        window[`portfolio-history-management-form-${ formUniqueId }-row-${ formRowIndex }`] as HTMLElement;

    return {
        assetIdInput: formRow.querySelector("[name$='[assetId]']"),
        assetTickerInput: formRow.querySelector("[name$='[assetTicker]']"),
        assetActionButton: formRow.querySelector("[data-asset-action-button]"),
        assetTickerMessage: formRow.querySelector("[data-asset-ticker-message]"),
        assetNameInput: formRow.querySelector("[name$='[assetName]']"),
    };
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

        const rowAssetElements = getFormRowAssetElements(formUniqueId, formRowIndex);

        if(ASSET_ACTION_BUTTON_IDENTITIES.isInSearchMode(rowAssetElements.assetActionButton)) {
            getAsset(rowAssetElements);
        }
        else if(ASSET_ACTION_BUTTON_IDENTITIES.isInResetMode(rowAssetElements.assetActionButton)) {
            resetAsset(rowAssetElements);
        }
    },
    validateAssetTicker(formRowIndex: number, formUniqueId: string) {

        const rowAssetElements = getFormRowAssetElements(formUniqueId, formRowIndex);

        if(ASSET_ACTION_BUTTON_IDENTITIES.isInSearchMode(rowAssetElements.assetActionButton)) {
            rowAssetElements.assetTickerInput.setCustomValidity("Reference an existing asset or create a new one");
            rowAssetElements.assetTickerInput.reportValidity();
        }
    },
};

export default portfolioHistoryManagement;