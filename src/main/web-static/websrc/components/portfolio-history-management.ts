import htmx from "htmx.org";
import api from "../api/api";
import { Asset } from "../domain/asset";
import BigNumber from "bignumber.js";
import { BootstrapClasses, BootstrapIconClasses } from "../infra/bootstrap/constants";
import { BeforeSwapEventDetail } from "../infra/htmx";
import { ObservationTimestamp } from "../domain/portfolio-allocation";

const PORTFOLIO_ALLOCATION_MANAGEMENT_TBODY_PREFIX = "portfolio-history-management-form-tbody-";
const PORTFOLIO_ALLOCATION_MANAGEMENT_HISTORY_CONTAINER = "accordion-portfolio-history-management";
// const PORTFOLIO_ALLOCATION_MANAGEMENT_HISTORY_CONTAINER = "accordion-portfolio-history-management-items";
// const PORTFOLIO_ALLOCATION_MANAGEMENT_HISTORY_OBS_CONTAINER_PREFIX = "portfolio-history-management-container-";

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

const portfolioHistoryManagement = {

    handlebarPortfolioHistoryManagementRowTemplate: null,
    handlebarPortfolioHistoryManagementContainerTemplate: null,

    modifyObservationsResponse(
        originalServerResponseJSON: string,
        eventDetail: BeforeSwapEventDetail,
    ) {
        const originalServerResponse = JSON.parse(originalServerResponseJSON) as ObservationTimestamp[];

        const modifiedServerResponse: ObservationTimestamp[] = [
            {
                id: 0,
                timeTag: "New Observation *",
            },
            ...originalServerResponse,
        ];
        eventDetail.serverResponse = JSON.stringify(modifiedServerResponse);
    },

    configEventsAfterSettling() {

        const historyManagementContainer = window[PORTFOLIO_ALLOCATION_MANAGEMENT_HISTORY_CONTAINER] as HTMLElement;

        //TODO generalize this code as "data-modify-response" referring to a function
        historyManagementContainer.addEventListener("htmx:beforeSwap", (event: CustomEvent) => {
            console.log("========>", event);

            const eventDetail = event.detail as BeforeSwapEventDetail;
            const eventRequestPath = eventDetail.pathInfo.finalRequestPath;

            // only if request was to /api/portfolio/:portfolioId/history/observation
            if(!eventDetail.isError && eventRequestPath.match(
                RegExp("^\\/api\\/portfolio\\/.+\\/history\\/observation$"))) {
                const originalServerResponseJSON = eventDetail.serverResponse;
                this.modifyObservationsResponse(originalServerResponseJSON, eventDetail);
            }
        });
    },

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

    reloadObservationHistory(formRowIndex: number) {

        const triggerELement = window[`portfolio-history-management-trigger-${ formRowIndex }`] as HTMLElement;

        const form = window[`portfolio-history-management-form-${ formRowIndex }`] as HTMLFormElement;
        form.reset();

        htmx.trigger(triggerELement, "reload-history-observation");
        window["loadPortfolioHistortDatalists"]();
    },

    //NOT WORKING! JSON FORM PLUGIN DOES NOT ACTIVATE
    // addHistoryObservation() {
    //
    //     const historyContainer = window[PORTFOLIO_ALLOCATION_MANAGEMENT_HISTORY_CONTAINER] as HTMLElement;
    //
    //     const newObservationContainer =
    //         window[PORTFOLIO_ALLOCATION_MANAGEMENT_HISTORY_OBS_CONTAINER_PREFIX + "0"] as HTMLElement;
    //
    //     if(newObservationContainer) {
    //         return;
    //     }
    //
    //     const newObservationHtml = this.handlebarPortfolioHistoryManagementContainerTemplate([
    //         {
    //             id: 0,
    //             timeTag: "New Observation *",
    //         },
    //     ]);
    //     historyContainer.insertAdjacentHTML("afterbegin", newObservationHtml);
    //
    //     const newObservation = historyContainer.firstElementChild as HTMLElement;
    //
    //     // Process the entire new observation element first
    //     htmx.process(newObservation);
    // },
};

export default portfolioHistoryManagement;
