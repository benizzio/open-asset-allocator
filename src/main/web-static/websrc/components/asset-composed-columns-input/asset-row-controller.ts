/**
 * Controls asset lookup and selection transitions for one composed asset row.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */

import { Asset } from "../../domain/asset";
import api from "../../api/api";
import notifications from "../notifications";
import { ASSET_ROW_SELECTION_CHANGE_EVENT, AssetRowSelectionChangeState } from "./selection-events";
import type {
    AssetRowSelectionChangeDetail,
    AssetRowSelectionChangePayload,
} from "./selection-events";
import {
    ASSET_ACTION_BUTTON_IDENTITIES,
    ASSET_ACTION_BUTTON_SELECTOR,
    ASSET_LOOKUP_PENDING_ERROR_MESSAGE,
    ASSET_NOT_FOUND_ERROR_MESSAGE,
    ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE,
    ASSET_SEARCH_INPUT_ATTRIBUTE,
    ASSET_SELECTION_ERROR_MESSAGE,
    FAILED_TO_FETCH_ASSET_ERROR_PREFIX,
    INVALID_INPUT_CLASS,
    INVALID_TICKER_ERROR_MESSAGE,
    NEW_ASSET_TICKER_MESSAGE_SELECTOR,
    REQUIRED_FOR_SEARCH_ERROR_MESSAGE,
    TICKER_CANDIDATE_PATTERN,
    TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE,
    TICKER_TOO_LONG_ERROR_MESSAGE,
    AssetSelectionState,
} from "./constants";
import type { AutocompleteControllerActions } from "./types";

/**
 * Connects an asset row's lookup action, selected ticker, and asset-name controls.
 *
 * @example
 * ```ts
 * const row = new AssetRowController(rowId, idField, tickerField, nameField, autocomplete);
 * row.handleAssetActionButtonClick();
 * ```
 *
 * @author benizzio
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
export class AssetRowController {
    private container: HTMLElement;
    private autocompleteWrapper: HTMLElement;
    private assetSearchInput: HTMLInputElement;
    private assetIdInput: HTMLInputElement;
    private assetTickerInput: HTMLInputElement;
    private assetActionButton: HTMLButtonElement;
    private newAssetTickerMessage: HTMLDivElement;
    private assetTickerExtraErrorMessageDiv: HTMLDivElement;
    private assetNameInput: HTMLInputElement;

    /** Resolves the row controls and rejects incomplete composed-row markup.
     *
     * @param containerId - ID of the composed row container.
     * @param assetIdHiddenFieldName - Name of the row's hidden asset ID input.
     * @param assetTickerFieldName - Name of the row's ticker input.
     * @param assetNameFieldName - Name of the row's asset-name input.
     * @param autocomplete - Shared autocomplete controller used by the row.
     * @example `new AssetRowController(rowId, idName, tickerName, assetName, autocomplete)`
     * @author GPT-6 Luna
     */
    constructor(
        containerId: string,
        assetIdHiddenFieldName: string,
        assetTickerFieldName: string,
        assetNameFieldName: string,
        private readonly autocomplete: AutocompleteControllerActions,
    ) {
        const container = document.getElementById(containerId);

        if(!container) {
            throw new Error(`Unable to find asset row container '${ containerId }'.`);
        }

        this.container = container;
        this.ensureSelectionGeneration();
        this.autocompleteWrapper = container.querySelector<HTMLElement>(`[${ ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE }]`);
        this.assetSearchInput = container.querySelector<HTMLInputElement>(`[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`);
        this.assetIdInput = container.querySelector<HTMLInputElement>(`[name='${ assetIdHiddenFieldName }']`);
        this.assetTickerInput = container.querySelector<HTMLInputElement>(`[name='${ assetTickerFieldName }']`);
        this.assetActionButton = container.querySelector<HTMLButtonElement>(ASSET_ACTION_BUTTON_SELECTOR);
        this.newAssetTickerMessage = container.querySelector<HTMLDivElement>(NEW_ASSET_TICKER_MESSAGE_SELECTOR);
        this.assetNameInput = container.querySelector<HTMLInputElement>(`[name='${ assetNameFieldName }']`);

        this.assetTickerExtraErrorMessageDiv = container.querySelector<HTMLDivElement>(
            `[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`,
        );

        const hasMissingSearchControls = !this.autocompleteWrapper || !this.assetSearchInput;
        const hasMissingAssetControls = !this.assetIdInput || !this.assetTickerInput;
        const hasMissingActionControls = !this.assetActionButton || !this.newAssetTickerMessage;
        const hasMissingNameAndValidationControls = !this.assetNameInput || !this.assetTickerExtraErrorMessageDiv;

        const hasMissingRequiredControls = hasMissingSearchControls || hasMissingAssetControls
            || hasMissingActionControls || hasMissingNameAndValidationControls;

        if(hasMissingRequiredControls) {
            throw new Error(`Asset row '${ containerId }' is missing required controls.`);
        }
    }

    /** Returns whether this row is waiting for a search or asset lookup.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private isInSearchMode(): boolean {
        return this.autocompleteWrapper.dataset.assetSelectionState === AssetSelectionState.SEARCH;
    }

    /** Returns whether the committed asset state can be cleared with the action button.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private isInResetMode(): boolean {
        const selectionState = this.autocompleteWrapper.dataset.assetSelectionState;
        return selectionState === AssetSelectionState.EXISTING || selectionState === AssetSelectionState.NEW;
    }

    /** Updates search/reset styling and icon without changing row selection state.
     *
     * @author GPT-6 Luna
     */
    private switchAssetActionButtonIdentity(identity: typeof ASSET_ACTION_BUTTON_IDENTITIES.search): void {
        this.assetActionButton.className = identity.classes;
        this.assetActionButton.innerHTML = `<span class="${ identity.iconClasses }"></span>`;
    }

    /** Changes the row to show a selected existing asset as read-only fields.
     *
     * @author benizzio
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private activateExistingAssetMode(asset: Asset): void {
        this.completeAssetLookup();
        this.assetActionButton.focus();
        this.autocompleteWrapper.dataset.assetSelectionState = AssetSelectionState.EXISTING;
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.reset);

        this.assetSearchInput.setCustomValidity("");
        this.assetSearchInput.readOnly = false;
        this.assetSearchInput.disabled = true;
        this.assetSearchInput.hidden = true;

        this.assetTickerInput.hidden = false;
        this.assetTickerInput.disabled = false;
        this.assetTickerInput.readOnly = true;
        this.assetTickerInput.value = asset.ticker;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = true;
        this.assetNameInput.required = false;
        this.assetNameInput.value = asset.name ?? "";

        this.assetIdInput.value = asset.id?.toString() ?? "";
        this.newAssetTickerMessage.style.display = "none";
        this.clearSearchErrorFeedback();
        this.publishSelectionChange({ state: AssetRowSelectionChangeState.EXISTING, asset });
    }

    /** Changes the row to require a name for an asset that is not yet stored.
     *
     * @author benizzio
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private activateNewAssetMode(ticker: string): void {
        this.completeAssetLookup();
        this.assetActionButton.focus();
        this.autocompleteWrapper.dataset.assetSelectionState = AssetSelectionState.NEW;
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.reset);

        this.assetSearchInput.setCustomValidity("");
        this.assetSearchInput.readOnly = false;
        this.assetSearchInput.disabled = true;
        this.assetSearchInput.hidden = true;

        this.assetTickerInput.value = ticker;
        this.assetTickerInput.hidden = false;
        this.assetTickerInput.disabled = false;
        this.assetTickerInput.readOnly = true;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = true;
        this.assetNameInput.value = "";

        this.assetIdInput.value = "";
        this.newAssetTickerMessage.style.display = "";
        this.clearSearchErrorFeedback();
        this.assetNameInput.focus();
        this.publishSelectionChange({ state: AssetRowSelectionChangeState.NEW });
    }

    /** Clears the committed asset and restores editable search mode.
     *
     * @author GPT-6 Luna
     * @author benizzio
     * @author GPT-6 Sol
     */
    private resetToSearchMode(): void {
        this.autocomplete.close(this.assetSearchInput);
        const autocompleteState = this.autocomplete.getState(this.assetSearchInput);

        if(autocompleteState) {
            autocompleteState.isLookupPending = false;
        }

        this.autocompleteWrapper.dataset.assetSelectionState = AssetSelectionState.SEARCH;
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.search);
        this.assetActionButton.disabled = false;

        this.assetSearchInput.value = "";
        this.assetSearchInput.setCustomValidity("");
        this.assetSearchInput.readOnly = false;
        this.assetSearchInput.disabled = false;
        this.assetSearchInput.hidden = false;
        this.assetSearchInput.classList.remove(INVALID_INPUT_CLASS);

        this.assetTickerInput.value = "";
        this.assetTickerInput.setCustomValidity("");
        this.assetTickerInput.readOnly = true;
        this.assetTickerInput.disabled = true;
        this.assetTickerInput.hidden = true;
        this.assetTickerInput.classList.remove(INVALID_INPUT_CLASS);

        this.assetNameInput.value = "";
        this.assetNameInput.style.display = "none";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = false;

        this.assetIdInput.value = "";
        this.newAssetTickerMessage.style.display = "none";
        this.clearSearchErrorFeedback();

        this.assetSearchInput.focus();
        this.autocomplete.open(this.assetSearchInput);
        this.publishSelectionChange({ state: AssetRowSelectionChangeState.SEARCH });
    }

    /** Initializes a missing or invalid generation without resetting a valid row's history.
     *
     * @author GPT-6 Luna
     */
    private ensureSelectionGeneration(): number {
        const currentGeneration = Number(this.container.dataset.assetSelectionGeneration);

        if(!Number.isSafeInteger(currentGeneration) || currentGeneration < 0) {
            this.container.dataset.assetSelectionGeneration = "0";
            return 0;
        }

        return currentGeneration;
    }

    /** Increments the row generation and publishes its completed selection transition.
     *
     * @author GPT-6 Luna
     */
    private publishSelectionChange(selection: AssetRowSelectionChangePayload): void {
        const currentGeneration = this.ensureSelectionGeneration();

        if(currentGeneration === Number.MAX_SAFE_INTEGER) {
            throw new Error("Asset row selection generation has reached its maximum safe integer.");
        }

        const generation = currentGeneration + 1;
        this.container.dataset.assetSelectionGeneration = String(generation);
        const detail: AssetRowSelectionChangeDetail = { ...selection, generation };

        this.container.dispatchEvent(new CustomEvent<AssetRowSelectionChangeDetail>(
            ASSET_ROW_SELECTION_CHANGE_EVENT,
            { bubbles: true, detail },
        ));
    }

    /** Clears search validation feedback before a new lookup attempt.
     *
     * @author GPT-6 Luna
     */
    private clearSearchFieldValidation(): void {
        this.assetSearchInput.setCustomValidity("");
        this.assetSearchInput.classList.remove(INVALID_INPUT_CLASS);
        this.clearSearchErrorFeedback();
    }

    /** Trims and validates a literal ticker or ID before sending it to the lookup API.
     *
     * @author GPT-6 Luna
     */
    private validateSearchUniqueIdentifier(): string {
        const assetUniqueIdentifier = this.assetSearchInput.value.trim();

        if(!assetUniqueIdentifier) {
            this.assetSearchInput.setCustomValidity(REQUIRED_FOR_SEARCH_ERROR_MESSAGE);
            this.assetSearchInput.reportValidity();
            return "";
        }

        if(assetUniqueIdentifier.length > 40) {
            this.assetSearchInput.setCustomValidity(TICKER_TOO_LONG_ERROR_MESSAGE);
            this.assetSearchInput.reportValidity();
            return "";
        }

        if(!TICKER_CANDIDATE_PATTERN.test(assetUniqueIdentifier)) {
            this.assetSearchInput.setCustomValidity(INVALID_TICKER_ERROR_MESSAGE);
            this.assetSearchInput.reportValidity();
            return "";
        }

        this.assetSearchInput.value = assetUniqueIdentifier;
        this.assetSearchInput.setCustomValidity("");
        return assetUniqueIdentifier;
    }

    /**
     * Searches the typed ticker or resets an asset that was already resolved.
     *
     * @param selectedTicker - Optional exact ticker chosen from autocomplete.
     * @example `row.handleAssetActionButtonClick("ABC")`
     * @author GPT-6 Luna
     * @author benizzio
     * @author GPT-6 Sol
     */
    handleAssetActionButtonClick(selectedTicker?: string): void {
        if(this.isInSearchMode()) {
            this.clearSearchFieldValidation();

            if(selectedTicker !== undefined) {
                this.assetSearchInput.value = selectedTicker;
            }
            const searchUniqueIdentifier = this.validateSearchUniqueIdentifier();

            if(searchUniqueIdentifier) {
                if(!this.beginAssetLookup()) {
                    return;
                }

                this.resolveAssetLookup(searchUniqueIdentifier);
            }
        }
        else if(this.autocompleteWrapper.dataset.assetSelectionState === AssetSelectionState.PENDING) {
            return;
        }
        else if(this.isInResetMode()) {
            this.resetToSearchMode();
        }
    }

    /**
     * Sets custom validity when this row has not committed an existing or new ticker.
     *
     * @example `row.validateForPost()`
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    validateForPost(): void {
        const selectionState = this.autocompleteWrapper.dataset.assetSelectionState;

        if(selectionState === AssetSelectionState.SEARCH || selectionState === AssetSelectionState.PENDING) {
            const message = selectionState === AssetSelectionState.PENDING
                ? ASSET_LOOKUP_PENDING_ERROR_MESSAGE
                : ASSET_SELECTION_ERROR_MESSAGE;
            this.assetSearchInput.setCustomValidity(message);
            this.assetSearchInput.classList.add(INVALID_INPUT_CLASS);
            this.assetTickerExtraErrorMessageDiv.textContent = message;
            this.assetTickerExtraErrorMessageDiv.style.display = "contents";
        }
        else {
            this.assetSearchInput.setCustomValidity("");
        }
    }

    /** Marks a row lookup pending and locks its ticker against duplicate requests.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private beginAssetLookup(): boolean {
        const state = this.autocomplete.ensureState(this.assetSearchInput);

        if(state?.isLookupPending) {
            return false;
        }

        this.autocomplete.close(this.assetSearchInput);

        if(state) {
            state.isLookupPending = true;
        }

        this.autocompleteWrapper.dataset.assetSelectionState = AssetSelectionState.PENDING;
        this.assetSearchInput.setCustomValidity(ASSET_LOOKUP_PENDING_ERROR_MESSAGE);
        this.assetSearchInput.readOnly = true;
        this.assetActionButton.disabled = true;
        return true;
    }

    /** Releases lookup state after an existing asset or a new-asset path resolves.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private completeAssetLookup(): void {
        const state = this.autocomplete.getState(this.assetSearchInput);

        if(state) {
            state.isLookupPending = false;
        }

        this.assetActionButton.disabled = false;
        this.autocomplete.close(this.assetSearchInput);
        this.assetSearchInput.setCustomValidity("");
    }

    /** Restores editable search mode after a lookup fails for a reason other than a missing asset.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private restoreSearchModeAfterLookup(): void {
        this.completeAssetLookup();
        this.autocompleteWrapper.dataset.assetSelectionState = AssetSelectionState.SEARCH;
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.search);
        this.assetSearchInput.readOnly = false;
        this.assetSearchInput.disabled = false;
        this.assetSearchInput.hidden = false;
        this.assetTickerInput.value = "";
        this.assetTickerInput.hidden = true;
        this.assetTickerInput.disabled = true;
        this.assetTickerInput.readOnly = true;
        this.assetIdInput.value = "";
        this.assetNameInput.value = "";
        this.assetNameInput.style.display = "none";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = false;
        this.newAssetTickerMessage.style.display = "none";
        this.clearSearchFieldValidation();
        this.assetSearchInput.focus();
    }

    /** Clears validation feedback attached to the transient search control.
     *
     * @author GPT-6 Luna
     */
    private clearSearchErrorFeedback(): void {
        this.assetSearchInput.classList.remove(INVALID_INPUT_CLASS);
        this.assetTickerExtraErrorMessageDiv.textContent = "";
        this.assetTickerExtraErrorMessageDiv.style.display = "none";
    }
    /** Resolves a ticker and updates this row or restores search mode after an error.
     *
     * @author GPT-6 Luna
     */
    private resolveAssetLookup(searchUniqueIdentifier: string): void {
        api.getAsset(searchUniqueIdentifier)
            .then(responseBody => {
                if(!this.container.isConnected) {
                    return;
                }

                if(api.isAPIErrorResponse(responseBody)) {
                    if(responseBody.errorMessage === ASSET_NOT_FOUND_ERROR_MESSAGE) {
                        this.activateNewAssetMode(searchUniqueIdentifier);
                    }
                    else {
                        this.restoreSearchModeAfterLookup();
                        notifications.notifyErrorResponse(responseBody);
                    }
                    return;
                }

                this.activateExistingAssetMode(responseBody as Asset);
            })
            .catch(error => {
                if(!this.container.isConnected) {
                    return;
                }

                this.restoreSearchModeAfterLookup();
                console.error("Error fetching asset:", error);
                notifications.notifyErrorResponse({ errorMessage: FAILED_TO_FETCH_ASSET_ERROR_PREFIX + error.message });
            });
    }
}
