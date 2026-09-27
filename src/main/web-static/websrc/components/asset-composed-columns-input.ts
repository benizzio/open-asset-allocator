/**
 * Provides shared asset ticker autocomplete and composed asset-field behavior for portfolio forms.
 *
 * @module components/asset-composed-columns-input
 *
 * @author GPT-6 Luna
 * @author benizzio
 */
import { Asset } from "../domain/asset";
import { BootstrapClasses, BootstrapIconClasses } from "../infra/bootstrap/constants";
import api from "../api/api";
import htmx from "htmx.org";
import notifications from "./notifications";

const TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE = "data-asset-ticker-extra-error-message";
const ASSET_TICKER_AUTOCOMPLETE_ATTRIBUTE = "data-asset-ticker-autocomplete";
const ASSET_TICKER_SUGGESTIONS_ATTRIBUTE = "data-asset-ticker-suggestions";
const ASSET_TICKER_LISTBOX_ATTRIBUTE = "data-asset-ticker-listbox";
const ASSET_TICKER_STATUS_ATTRIBUTE = "data-asset-ticker-status";

/**
 * Stores per-input DOM references and transient state for keyboard and pointer interaction.
 *
 * @author GPT-6 Luna
 */
interface AssetTickerAutocompleteState {
    wrapper: HTMLElement;
    suggestions: HTMLElement;
    listbox: HTMLElement;
    status: HTMLElement;
    optionValues: string[];
    activeOptionIndex: number;
    isOpen: boolean;
    isLookupPending: boolean;
    pointerSelection?: { optionIndex: number; pointerId: number; x: number; y: number };
    outsidePointerHandler?: (event: PointerEvent) => void;
    removalObserver?: MutationObserver;
}

const autocompleteStates = new WeakMap<HTMLInputElement, AssetTickerAutocompleteState>();
let activeAutocompleteInput: HTMLInputElement | null = null;
let nextAutocompleteId = 0;
let assetDatalistListenerRegistered = false;

/**
 * Resolves or creates the DOM state for one asset ticker autocomplete.
 *
 * @param input - Ticker input associated with the autocomplete markup.
 * @returns The widget state, or `null` when the input is not in the component template.
 *
 * @author GPT-6 Luna
 */
function getAutocompleteState(input: HTMLInputElement): AssetTickerAutocompleteState | null {

    const currentState = autocompleteStates.get(input);

    if(currentState) {
        return currentState;
    }

    const wrapper = input.closest<HTMLElement>(`[${ ASSET_TICKER_AUTOCOMPLETE_ATTRIBUTE }]`);
    const suggestions = wrapper?.querySelector<HTMLElement>(`[${ ASSET_TICKER_SUGGESTIONS_ATTRIBUTE }]`);
    const listbox = wrapper?.querySelector<HTMLElement>(`[${ ASSET_TICKER_LISTBOX_ATTRIBUTE }]`);
    const status = wrapper?.querySelector<HTMLElement>(`[${ ASSET_TICKER_STATUS_ATTRIBUTE }]`);

    if(!wrapper || !suggestions || !listbox || !status) {
        return null;
    }

    const listboxId = `asset-ticker-listbox-${ ++nextAutocompleteId }`;
    listbox.id = listboxId;
    input.setAttribute("aria-controls", listboxId);

    const state: AssetTickerAutocompleteState = {
        wrapper,
        suggestions,
        listbox,
        status,
        optionValues: [],
        activeOptionIndex: -1,
        isOpen: false,
        isLookupPending: false,
    };

    listbox.addEventListener("pointerdown", event => {
        const pointerEvent = event as PointerEvent;
        const option = (pointerEvent.target as HTMLElement).closest<HTMLElement>("[data-asset-ticker-option]");

        if(!option || pointerEvent.button !== 0) {
            state.pointerSelection = undefined;
            return;
        }

        pointerEvent.preventDefault();

        state.pointerSelection = {
            optionIndex: Number(option.dataset.optionIndex),
            pointerId: pointerEvent.pointerId,
            x: pointerEvent.clientX,
            y: pointerEvent.clientY,
        };
    });

    listbox.addEventListener("pointerup", event => {
        const pointerEvent = event as PointerEvent;
        const pointerSelection = state.pointerSelection;
        state.pointerSelection = undefined;

        if(pointerSelection?.pointerId !== pointerEvent.pointerId) {
            return;
        }

        const option = (pointerEvent.target as HTMLElement).closest<HTMLElement>("[data-asset-ticker-option]");
        const optionIndex = option ? Number(option.dataset.optionIndex) : -1;
        const movedX = pointerEvent.clientX - pointerSelection.x;
        const movedY = pointerEvent.clientY - pointerSelection.y;

        if(optionIndex !== pointerSelection.optionIndex || Math.hypot(movedX, movedY) > 8) {
            return;
        }

        const ticker = state.optionValues[optionIndex];

        if(ticker !== undefined) {
            selectAssetTicker(input, ticker);
        }
    });

    listbox.addEventListener("pointercancel", () => {
        state.pointerSelection = undefined;
    });

    listbox.addEventListener("click", event => {
        const clickEvent = event as MouseEvent;

        if(clickEvent.detail !== 0) {
            return;
        }

        const option = (clickEvent.target as HTMLElement).closest<HTMLElement>("[data-asset-ticker-option]");
        const optionValue = option ? state.optionValues[Number(option.dataset.optionIndex)] : undefined;

        if(optionValue !== undefined) {
            selectAssetTicker(input, optionValue);
        }
    });

    wrapper.addEventListener("focusout", event => {
        const focusEvent = event as FocusEvent;

        if(focusEvent.relatedTarget && wrapper.contains(focusEvent.relatedTarget as Node)) {
            return;
        }

        closeAssetTickerAutocomplete(input);
    });

    autocompleteStates.set(input, state);
    return state;
}

/**
 * Installs the shared listener that refreshes an open widget after its HTMX datalist reloads.
 *
 * @author GPT-6 Luna
 */
function ensureAssetDatalistLifecycleListener(): void {

    if(assetDatalistListenerRegistered) {
        return;
    }

    document.addEventListener("htmx:afterSettle", event => {
        const htmxEvent = event as CustomEvent<{ elt?: Element }>;
        const datalist = document.getElementById("datalist-assets");

        if(!datalist || (htmxEvent.detail?.elt !== datalist && event.target !== datalist)) {
            return;
        }

        if(activeAutocompleteInput && autocompleteStates.get(activeAutocompleteInput)?.isOpen) {
            renderAssetTickerSuggestions(activeAutocompleteInput);
        }
    });

    assetDatalistListenerRegistered = true;
}

/**
 * Opens one autocomplete and keeps its visible results synchronized with the local datalist.
 *
 * @param input - Editable asset ticker input to open.
 *
 * @author GPT-6 Luna
 */
function openAssetTickerAutocomplete(input: HTMLInputElement): void {

    if(input.readOnly || input.disabled) {
        return;
    }

    const state = getAutocompleteState(input);

    if(!state || state.isLookupPending) {
        return;
    }

    if(state.isOpen) {
        return;
    }

    ensureAssetDatalistLifecycleListener();

    if(activeAutocompleteInput && activeAutocompleteInput !== input) {
        closeAssetTickerAutocomplete(activeAutocompleteInput);
    }

    activeAutocompleteInput = input;
    state.isOpen = true;
    state.suggestions.hidden = false;
    input.setAttribute("aria-expanded", "true");

    state.outsidePointerHandler = event => {
        if(!state.wrapper.contains(event.target as Node)) {
            closeAssetTickerAutocomplete(input);
        }
    };
    document.addEventListener("pointerdown", state.outsidePointerHandler, true);

    state.removalObserver = new MutationObserver(() => {
        if(!input.isConnected || input.getClientRects().length === 0) {
            closeAssetTickerAutocomplete(input);
        }
    });

    state.removalObserver.observe(document.body, {
        childList: true,
        subtree: true,
        attributes: true,
        attributeFilter: ["class", "hidden", "style"],
    });

    renderAssetTickerSuggestions(input);
}

/**
 * Closes one autocomplete without changing the ticker currently typed by the user.
 *
 * @param input - Ticker input whose suggestions should be dismissed.
 *
 * @author GPT-6 Luna
 */
function closeAssetTickerAutocomplete(input: HTMLInputElement): void {

    const state = autocompleteStates.get(input);

    if(!state) {
        return;
    }

    state.isOpen = false;
    state.activeOptionIndex = -1;
    state.pointerSelection = undefined;
    state.suggestions.hidden = true;
    state.status.hidden = true;
    input.setAttribute("aria-expanded", "false");
    input.removeAttribute("aria-activedescendant");

    if(state.outsidePointerHandler) {
        document.removeEventListener("pointerdown", state.outsidePointerHandler, true);
    }
    state.outsidePointerHandler = undefined;
    state.removalObserver?.disconnect();
    state.removalObserver = undefined;

    if(activeAutocompleteInput === input) {
        activeAutocompleteInput = null;
    }
}

/**
 * Filters the current prefetched ticker options and renders every matching result.
 *
 * @param input - Ticker input whose current value is used as the filter query.
 *
 * @author GPT-6 Luna
 */
function renderAssetTickerSuggestions(input: HTMLInputElement): void {

    const state = getAutocompleteState(input);

    if(!state?.isOpen) {
        return;
    }

    const sourceId = input.dataset.assetTickerSource;
    const source = sourceId ? document.getElementById(sourceId) as HTMLDataListElement | null : null;
    const query = input.value.toLocaleLowerCase();
    const seenTickers = new Set<string>();

    state.optionValues = source
        ? Array.from(source.options)
            .map(option => option.value.trim())
            .filter(ticker => {
                const normalizedTicker = ticker.toLocaleLowerCase();

                if(!ticker || !normalizedTicker.includes(query) || seenTickers.has(normalizedTicker)) {
                    return false;
                }

                seenTickers.add(normalizedTicker);
                return true;
            })
        : [];
    state.activeOptionIndex = -1;
    state.listbox.replaceChildren();
    input.removeAttribute("aria-activedescendant");

    state.optionValues.forEach((ticker, index) => {
        const option = document.createElement("div");
        option.id = `${ state.listbox.id }-option-${ index }`;
        option.className = "asset-ticker-autocomplete__option";
        option.setAttribute("role", "option");
        option.setAttribute("aria-selected", "false");
        option.dataset.assetTickerOption = "";
        option.dataset.optionIndex = index.toString();
        option.textContent = ticker;
        state.listbox.append(option);
    });

    state.listbox.hidden = state.optionValues.length === 0;
    state.status.textContent = "No matching assets";
    state.status.hidden = state.optionValues.length > 0;
}

/**
 * Highlights a suggestion while keeping the input focused for keyboard navigation.
 *
 * @param input - Input that owns the active listbox.
 * @param optionIndex - Index of the option to highlight.
 *
 * @author GPT-6 Luna
 */
function highlightAssetTickerOption(input: HTMLInputElement, optionIndex: number): void {

    const state = autocompleteStates.get(input);

    if(!state?.optionValues.length) {
        return;
    }

    state.activeOptionIndex = Math.max(0, Math.min(optionIndex, state.optionValues.length - 1));

    const options = state.listbox.querySelectorAll<HTMLElement>("[data-asset-ticker-option]");

    options.forEach((option, index) => {
        const isActive = index === state.activeOptionIndex;
        option.setAttribute("aria-selected", isActive.toString());

        if(isActive) {
            input.setAttribute("aria-activedescendant", option.id);
            option.scrollIntoView({ block: "nearest" });
        }
    });
}

/**
 * Selects an exact ticker and delegates resolution to the existing row search action.
 *
 * @param input - Input associated with the selected suggestion.
 * @param ticker - Exact ticker value from the datalist.
 *
 * @author GPT-6 Luna
 */
function selectAssetTicker(input: HTMLInputElement, ticker: string): void {

    const state = autocompleteStates.get(input);

    if(input.readOnly || state?.isLookupPending) {
        return;
    }

    input.value = ticker;
    closeAssetTickerAutocomplete(input);

    const row = input.closest<HTMLTableRowElement>("tr");
    const assetIdInput = row?.querySelector<HTMLInputElement>("input[type='hidden'][data-null-if-empty]");
    const assetNameInput = row?.querySelector<HTMLInputElement>("input[aria-label='Asset name']");

    if(!row?.id || !assetIdInput || !assetNameInput) {
        console.error("Unable to resolve the asset row for the selected ticker.");
        return;
    }

    new AssetComposedColumnInput(row.id, assetIdInput.name, input.name, assetNameInput.name)
        .handleAssetActionButtonClick();
}

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

/**
 * Coordinates an asset ticker, its lookup button, and the related form fields for one row.
 *
 * @author benizzio
 */
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

    /**
     * Changes the row to display the selected existing asset as read-only fields.
     *
     * @author benizzio
     * @author GPT-6 Luna
     */
    activateExistingAssetMode(asset: Asset) {

        this.completeAssetLookup();
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.reset);

        this.assetTickerInput.readOnly = true;
        this.assetTickerInput.value = asset.ticker;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = true;
        this.assetNameInput.value = asset.name;

        this.assetIdInput.value = asset.id.toString();

        this.newAssetTickerMessage.style.display = "none";
    }

    /**
     * Changes the row to require the name of an asset that is not yet stored.
     *
     * @author benizzio
     * @author GPT-6 Luna
     */
    activateNewAssetMode() {

        this.completeAssetLookup();
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.reset);

        this.assetTickerInput.readOnly = true;

        this.assetNameInput.style.display = "";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = true;

        this.newAssetTickerMessage.style.display = "";
    }

    /**
     * Clears the current asset and restores editable ticker search mode.
     *
     * @author GPT-6 Luna
     * @author benizzio
     */
    resetToSearchMode() {

        closeAssetTickerAutocomplete(this.assetTickerInput);
        const autocompleteState = autocompleteStates.get(this.assetTickerInput);

        if(autocompleteState) {
            autocompleteState.isLookupPending = false;
        }

        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.search);
        this.assetActionButton.disabled = false;

        this.assetTickerInput.value = "";
        this.assetTickerInput.setCustomValidity("");
        this.assetTickerInput.reportValidity();
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

        this.assetTickerInput.focus();
        openAssetTickerAutocomplete(this.assetTickerInput);
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

    /**
     * Searches for the typed ticker or resets a previously resolved asset.
     *
     * @author GPT-6 Luna
     * @author benizzio
     */
    handleAssetActionButtonClick() {

        if(this.isInSearchMode()) {

            this.clearSearchFieldValidation();
            const searchUniqueIdentifier = this.validateSearchUniqueIdentifier();

            if(searchUniqueIdentifier) {
                if(!this.beginAssetLookup()) {
                    return;
                }

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

    /**
     * Marks a row lookup as pending and locks the ticker against duplicate requests.
     *
     * @returns `false` when a lookup is already pending; otherwise `true`.
     *
     * @author GPT-6 Luna
     */
    beginAssetLookup(): boolean {

        const state = getAutocompleteState(this.assetTickerInput);

        if(state?.isLookupPending) {
            return false;
        }

        closeAssetTickerAutocomplete(this.assetTickerInput);

        if(state) {
            state.isLookupPending = true;
        }

        this.assetTickerInput.readOnly = true;
        this.assetActionButton.disabled = true;
        return true;
    }

    /**
     * Releases lookup state after an existing asset or a new-asset path is resolved.
     *
     * @author GPT-6 Luna
     */
    completeAssetLookup(): void {

        const state = autocompleteStates.get(this.assetTickerInput);

        if(state) {
            state.isLookupPending = false;
        }

        this.assetActionButton.disabled = false;
        closeAssetTickerAutocomplete(this.assetTickerInput);
    }

    /**
     * Restores editable search mode after a lookup fails for a reason other than a missing asset.
     *
     * @author GPT-6 Luna
     */
    restoreSearchModeAfterLookup(): void {

        this.completeAssetLookup();
        this.assetTickerInput.readOnly = false;
    }
}

/**
 * Resolves an asset ticker and updates its owning row with the API result.
 *
 * @author GPT-6 Luna
 */
function getAsset(rowAssetElements: AssetComposedColumnInput, searchUniqueIdentifier: string) {

    api.getAsset(searchUniqueIdentifier)
        .then(responseBody => {

            if(api.isAPIErrorResponse(responseBody)) {
                if(responseBody.errorMessage === "Data not found") {
                    rowAssetElements.activateNewAssetMode();
                }
                else {
                    rowAssetElements.restoreSearchModeAfterLookup();
                    notifications.notifyErrorResponse(responseBody);
                }
                return;
            }

            rowAssetElements.activateExistingAssetMode(responseBody as Asset);
        })
        .catch(error => {
            rowAssetElements.restoreSearchModeAfterLookup();
            console.error("Error fetching asset:", error);
            notifications.notifyErrorResponse({ errorMessage: "Failed to fetch asset data: " + error.message });
        });
}

function loadClassesDatalist() {
    const datalistElement = window["datalist-classes"];
    htmx.trigger(datalistElement, "load-classes");
}

/**
 * Reloads the shared prefetched asset ticker source.
 *
 * @author GPT-6 Luna
 * @author benizzio
 */
function loadAssetsDatalist() {
    const datalistElement: HTMLElement = window["datalist-assets"];

    ensureAssetDatalistLifecycleListener();
    datalistElement.dataset.assetsInitialized = "false";
    htmx.trigger(datalistElement, "load-assets");
}

/**
 * Public browser handlers for the shared ticker input template and its row lookup actions.
 *
 * The HTML partial binds `handleAssetTickerFocus`, `handleAssetTickerInput`, and
 * `handleAssetTickerKeydown` directly to each generated ticker input. Existing row-specific
 * controllers continue to call the asset-action and validation methods with their own field names.
 *
 * @example
 * ```html
 * <input onfocus="AssetComposedColumnsInput.handleAssetTickerFocus(event)"
 *        oninput="maskTickerInput(this); AssetComposedColumnsInput.handleAssetTickerInput(event)"
 *        onkeydown="AssetComposedColumnsInput.handleAssetTickerKeydown(event)">
 * ```
 *
 * @author GPT-6 Luna
 * @author benizzio
 */
const AssetComposedColumnsInput = {

    /**
     * Opens suggestions when an editable asset ticker input receives focus.
     *
     * @param event - Focus event emitted by the ticker input.
     *
     * @example onfocus="AssetComposedColumnsInput.handleAssetTickerFocus(event)"
     *
     * @author GPT-6 Luna
     */
    handleAssetTickerFocus(event: FocusEvent): void {
        openAssetTickerAutocomplete(event.target as HTMLInputElement);
    },

    /**
     * Refreshes local ticker suggestions after the ticker input has been masked.
     *
     * @param event - Input event emitted by the ticker field.
     *
     * @example oninput="maskTickerInput(this); AssetComposedColumnsInput.handleAssetTickerInput(event)"
     *
     * @author GPT-6 Luna
     */
    handleAssetTickerInput(event: Event): void {

        const input = event.target as HTMLInputElement;
        openAssetTickerAutocomplete(input);
        renderAssetTickerSuggestions(input);
    },

    /**
     * Handles autocomplete navigation, suggestion selection, and the existing Enter-to-search action.
     *
     * @param event - The keyboard event from the asset ticker input
     *
     * @example onkeydown="AssetComposedColumnsInput.handleAssetTickerKeydown(event)"
     *
     * @author GPT-6 Luna
     */
    handleAssetTickerKeydown(event: KeyboardEvent) {

        if(event.isComposing) {
            return;
        }

        const inputElement = event.target as HTMLInputElement;
        let state = getAutocompleteState(inputElement);

        if(event.key === "ArrowDown" || event.key === "ArrowUp") {

            if(inputElement.readOnly) {
                return;
            }

            openAssetTickerAutocomplete(inputElement);
            state = getAutocompleteState(inputElement);

            if(state?.optionValues.length) {
                event.preventDefault();

                const nextIndex = state.activeOptionIndex < 0
                    ? (event.key === "ArrowDown" ? 0 : state.optionValues.length - 1)
                    : state.activeOptionIndex + (event.key === "ArrowDown" ? 1 : -1);
                highlightAssetTickerOption(inputElement, nextIndex);
            }
            return;
        }

        if(event.key === "Escape") {

            if(state?.isOpen) {
                event.preventDefault();
                closeAssetTickerAutocomplete(inputElement);
            }
            return;
        }

        if(event.key === "Tab") {
            closeAssetTickerAutocomplete(inputElement);
            return;
        }

        if(event.key === "Enter") {

            if(state?.activeOptionIndex >= 0) {
                event.preventDefault();
                const selectedTicker = state.optionValues[state.activeOptionIndex];

                if(selectedTicker !== undefined) {
                    selectAssetTicker(inputElement, selectedTicker);
                }
                return;
            }

            const actionButton = inputElement
                .closest(".input-group")
                ?.querySelector<HTMLButtonElement>("[data-asset-action-button]");

            if(actionButton?.className === ASSET_ACTION_BUTTON_IDENTITIES.search.classes) {
                event.preventDefault();
                actionButton.click();
            }
        }
    },

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
        field.setCustomValidity(errorMessage);
        field.reportValidity();
    },
};

export default AssetComposedColumnsInput;
