/**
 * Public API for asset search autocomplete and composed asset-row behavior.
 *
 * @module components/asset-composed-columns-input
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */

import htmx from "htmx.org";
import { AssetRowController } from "./asset-row-controller";
import { AssetSearchAutocompleteController } from "./autocomplete-controller";
import { handleAssetSearchKeydown } from "./autocomplete-interactions";
import {
    ASSET_ID_INPUT_SELECTOR,
    ASSET_NAME_INPUT_SELECTOR,
    ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE,
    ASSET_TICKER_INPUT_ATTRIBUTE,
    INVALID_INPUT_CLASS,
    TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE,
    UNABLE_TO_RESOLVE_ASSET_ROW_ERROR,
} from "./constants";
import { installAssetFormValidationGuards, invalidateSelectedAsset, validateAssetRowsForPost } from "./form-validation";

export { AssetSelectionState } from "./constants";
export { ASSET_ROW_SELECTION_CHANGE_EVENT, AssetRowSelectionChangeState } from "./selection-events";
export type {
    AssetRowSelectionChangeDetail,
    AssetRowSelectionChangeEvent,
} from "./selection-events";

/** Resolves an option's owning row and starts the same ticker lookup used by its action button.
 *
 * @author GPT-6 Luna
 */
function selectAssetTicker(input: HTMLInputElement, ticker: string): void {
    const row = input.closest<HTMLTableRowElement>("tr");
    const assetIdInput = row?.querySelector<HTMLInputElement>(ASSET_ID_INPUT_SELECTOR);
    const assetTickerInput = row?.querySelector<HTMLInputElement>(`[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`);
    const assetNameInput = row?.querySelector<HTMLInputElement>(ASSET_NAME_INPUT_SELECTOR);

    if(!row?.id || !assetIdInput || !assetTickerInput || !assetNameInput) {
        console.error(UNABLE_TO_RESOLVE_ASSET_ROW_ERROR);
        return;
    }

    new AssetRowController(
        row.id,
        assetIdInput.name,
        assetTickerInput.name,
        assetNameInput.name,
        autocompleteController,
    ).handleAssetActionButtonClick(ticker);
}

const autocompleteController = new AssetSearchAutocompleteController(selectAssetTicker);

/** Loads the shared class and asset datalists used by composed rows.
 *
 * @author GPT-6 Luna
 */
function loadDatalists(): void {
    const classesDatalistElement = window["datalist-classes"];
    htmx.trigger(classesDatalistElement, "load-classes");

    const assetsDatalistElement: HTMLElement = window["datalist-assets"];
    autocompleteController.ensureDatalistLifecycleListener();
    assetsDatalistElement.dataset.assetsInitialized = "false";
    htmx.trigger(assetsDatalistElement, "load-assets");
}

/**
 * Browser handlers and row helpers consumed by templates and component controllers.
 *
 * @example
 * ```html
 * <input onfocus="AssetComposedColumnsInput.handleAssetSearchFocus(event)"
 *        oninput="AssetComposedColumnsInput.handleAssetSearchInput(event)"
 *        onkeydown="AssetComposedColumnsInput.handleAssetSearchKeydown(event)">
 * ```
 *
 * @author GPT-6 Luna
 * @author benizzio
 * @author GPT-6.1 Sol
 * @author GPT-6 Sol
 */
const AssetComposedColumnsInput = {

    /**
     * Opens suggestions when the editable asset search input receives focus.
     *
     * @param event - Focus event emitted by the asset search input.
     * @example `AssetComposedColumnsInput.handleAssetSearchFocus(event)`
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    handleAssetSearchFocus(event: FocusEvent): void {
        autocompleteController.open(event.target as HTMLInputElement);
    },

    /**
     * Clears transient validation feedback and refreshes suggestions after the search query changes.
     *
     * @param event - Input event emitted by the asset search field.
     * @example `AssetComposedColumnsInput.handleAssetSearchInput(event)`
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    handleAssetSearchInput(event: Event): void {
        const input = event.target as HTMLInputElement;
        input.setCustomValidity("");
        input.classList.remove(INVALID_INPUT_CLASS);
        const autocompleteWrapper = input.closest<HTMLElement>(`[${ ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE }]`);

        const errorMessage = autocompleteWrapper?.parentElement?.querySelector<HTMLDivElement>(
            `[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`,
        );

        if(errorMessage) {
            errorMessage.textContent = "";
            errorMessage.style.display = "none";
        }

        autocompleteController.open(input);
        autocompleteController.render(input);
    },

    /**
     * Handles autocomplete navigation, selection, literal Enter-to-search fallback, Escape, and Tab.
     *
     * @param event - Keyboard event emitted by the asset search field.
     * @example `AssetComposedColumnsInput.handleAssetSearchKeydown(event)`
     * @author GPT-6 Luna
     * @author GPT-6.1 Sol
     * @author GPT-6 Sol
     */
    handleAssetSearchKeydown(event: KeyboardEvent): void {
        handleAssetSearchKeydown(event.target as HTMLInputElement, event, autocompleteController);
    },

    /**
     * Handles the row's action button by searching or resetting its committed asset state.
     *
     * @param containerId - ID of the composed row container.
     * @param assetIdHiddenFieldName - Name of the row's hidden asset ID input.
     * @param assetTickerFieldName - Name of the row's ticker input.
     * @param assetNameFieldName - Name of the row's asset-name input.
     * @example `AssetComposedColumnsInput.assetActionButtonClickHandler(rowId, idName, tickerName, assetName)`
     * @author GPT-6 Luna
     * @author benizzio
     */
    assetActionButtonClickHandler(
        containerId: string,
        assetIdHiddenFieldName: string,
        assetTickerFieldName: string,
        assetNameFieldName: string,
    ): void {
        const rowAssetElements = new AssetRowController(
            containerId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
            autocompleteController,
        );
        rowAssetElements.handleAssetActionButtonClick();
    },

    /**
     * Applies row-level validity feedback before an asset form is posted.
     *
     * @param containerId - ID of the composed row container.
     * @param assetIdHiddenFieldName - Name of the row's hidden asset ID input.
     * @param assetTickerFieldName - Name of the row's ticker input.
     * @param assetNameFieldName - Name of the row's asset-name input.
     * @example `AssetComposedColumnsInput.validateAssetElementsForPost(rowId, idName, tickerName, assetName)`
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    validateAssetElementsForPost(
        containerId: string,
        assetIdHiddenFieldName: string,
        assetTickerFieldName: string,
        assetNameFieldName: string,
    ): void {
        const rowAssetElements = new AssetRowController(
            containerId,
            assetIdHiddenFieldName,
            assetTickerFieldName,
            assetNameFieldName,
            autocompleteController,
        );
        rowAssetElements.validateForPost();
    },

    /**
     * Validates committed asset selections before submitting outside HTMX field-validation handling.
     *
     * @param form - Form containing one or more editable asset rows.
     * @returns `true` when every editable row is resolved or in confirmed new-asset mode.
     * @example `if(AssetComposedColumnsInput.validateFormBeforePost(form)) form.requestSubmit();`
     * @author GPT-6 Luna
     */
    validateFormBeforePost(form: HTMLFormElement): boolean {
        return validateAssetRowsForPost(form, true);
    },

    /**
     * Triggers reloads for the shared class and asset datalists.
     *
     * @example `AssetComposedColumnsInput.loadDatalists()`
     * @author GPT-6 Luna
     * @author benizzio
     */
    loadDatalists,

    /**
     * Marks a selected ticker invalid and displays its associated validation message.
     *
     * @param field - Selected ticker input to invalidate.
     * @param errorMessage - Validation message shown next to the ticker input.
     * @example `AssetComposedColumnsInput.invalidateSelectedAsset(tickerInput, "Ticker is invalid")`
     * @author GPT-6 Luna
     */
    invalidateSelectedAsset,
};

installAssetFormValidationGuards();

export default AssetComposedColumnsInput;
