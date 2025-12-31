import { Asset } from "../domain/asset";
import { BootstrapClasses, BootstrapIconClasses } from "../infra/bootstrap/constants";
import api from "../api/api";
import htmx from "htmx.org";

const TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE = "data-asset-ticker-extra-error-message";

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

class AssetComposedColumnInput {

    assetIdInput: HTMLInputElement;
    assetTickerInput: HTMLInputElement;
    assetActionButton: HTMLButtonElement;
    newAssetTickerMessage: HTMLDivElement;
    assetTickerExtraErrorMessageDiv: HTMLDivElement;
    assetNameInput: HTMLInputElement;

    constructor(
        containerId: string,
        assetIdHiddenFieldName: string,
        assetTickerFieldName: string,
        assetNameFieldName: string,
    ) {

        const container = window[containerId] as HTMLElement;

        this.assetIdInput = container.querySelector(`[name='${ assetIdHiddenFieldName }']`);
        this.assetTickerInput = container.querySelector(`[name='${ assetTickerFieldName }']`);
        this.assetActionButton = container.querySelector("[data-asset-action-button]");
        this.newAssetTickerMessage = container.querySelector("[data-new-asset-ticker-message]");
        this.assetNameInput = container.querySelector(`[name='${ assetNameFieldName }']`);
        this.assetTickerExtraErrorMessageDiv = container.querySelector(`[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`);
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

        this.assetTickerInput.readOnly = true;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = true;

        this.newAssetTickerMessage.style.display = "";
    }

    resetToSearchMode() {

        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.search);

        this.assetTickerInput.value = "";
        this.assetTickerInput.setCustomValidity("");
        this.assetTickerInput.reportValidity();
        this.assetTickerInput.focus();
        this.assetTickerInput.readOnly = false;
        this.assetTickerInput.classList.remove("is-invalid");

        this.assetNameInput.value = "";
        this.assetNameInput.style.display = "none";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = false;

        this.assetIdInput.value = "";

        this.newAssetTickerMessage.style.display = "none";
        this.assetTickerExtraErrorMessageDiv.textContent = "";
        this.assetTickerExtraErrorMessageDiv.style.display = "none";
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

function getAsset(rowAssetElements: AssetComposedColumnInput, searchUniqueIdentifier: string) {

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

function loadClassesDatalist() {
    const datalistElement = window["datalist-classes"];
    htmx.trigger(datalistElement, "load-classes");
}

function loadAssetsDatalist() {
    const datalistElement: HTMLElement = window["datalist-assets"];
    datalistElement.dataset.assetsInitialized = "false";
    htmx.trigger(datalistElement, "load-assets");
}

const AssetComposedColumnsInput = {

    assetActionButtonClickHandler(
        containerId: string,
        assetIdHiddenFieldName: string,
        assetTickerFieldName: string,
        assetNameFieldName: string,
    ) {
        const rowAssetElements = new AssetComposedColumnInput(
            containerId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
        rowAssetElements.handleAssetActionButtonClick();
    },

    validateAssetElementsForPost(
        containerId: string,
        assetIdHiddenFieldName: string,
        assetTickerFieldName: string,
        assetNameFieldName: string,
    ) {
        const rowAssetElements = new AssetComposedColumnInput(
            containerId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
        );
        rowAssetElements.validateForPost();
    },

    loadDatalists() {
        loadClassesDatalist();
        loadAssetsDatalist();
    },

    invalidateSelectedAsset(field: HTMLInputElement, errorMessage: string) {

        field.classList.add("is-invalid");

        const parentColumn = field.closest("td");

        const extraErrorMessageDiv =
            parentColumn.querySelector(`[data-asset-ticker-extra-error-message="${ field.name }"]`) as HTMLDivElement;

        extraErrorMessageDiv.textContent = errorMessage;
        extraErrorMessageDiv.style.display = "contents";
        field.setCustomValidity("errorMessage");
        field.reportValidity();
    },
};

export default AssetComposedColumnsInput;