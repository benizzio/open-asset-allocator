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
const ASSET_SEARCH_INPUT_ATTRIBUTE = "data-asset-search-input";
const ASSET_TICKER_INPUT_ATTRIBUTE = "data-asset-ticker-input";
const ASSET_SELECTION_ERROR_MESSAGE = "Reference an existing asset or create a new one";
const ASSET_LOOKUP_PENDING_ERROR_MESSAGE = "Asset lookup is still in progress";
const TICKER_CANDIDATE_PATTERN = /^[A-Za-z0-9_:.-]+$/;

/** Describes one selectable asset suggestion or missing-ticker create action.
 *
 * @author GPT-6 Luna
 */
type AssetTickerAutocompleteOption = {
    kind: "asset" | "create";
    label: string;
    ticker: string;
};

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
    options: AssetTickerAutocompleteOption[];
    activeOptionIndex: number;
    activeOptionSource: "keyboard" | "pointer" | null;
    isOpen: boolean;
    isLookupPending: boolean;
    pointerSelection?: { optionIndex: number; pointerId: number; x: number; y: number; isCancelled: boolean };
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
        options: [],
        activeOptionIndex: -1,
        activeOptionSource: null,
        isOpen: false,
        isLookupPending: false,
    };

    status.classList.add("visually-hidden");

    listbox.addEventListener("pointerover", event => {
        const pointerEvent = event as PointerEvent;
        const option = (pointerEvent.target as HTMLElement).closest<HTMLElement>("[data-asset-ticker-option]");

        if(pointerEvent.pointerType === "mouse" && option) {
            highlightAssetTickerOption(input, Number(option.dataset.optionIndex), "pointer", false);
        }
    });

    listbox.addEventListener("pointerout", event => {
        const pointerEvent = event as PointerEvent;

        if(pointerEvent.pointerType !== "mouse"
            || (pointerEvent.relatedTarget instanceof Node && listbox.contains(pointerEvent.relatedTarget))) {
            return;
        }

        if(state.activeOptionSource === "pointer" && !state.pointerSelection) {
            clearAssetTickerOptionHighlight(input);
        }
    });

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
            isCancelled: false,
        };
        highlightAssetTickerOption(input, state.pointerSelection.optionIndex, "pointer", false);
    });

    listbox.addEventListener("pointermove", event => {
        const pointerEvent = event as PointerEvent;
        const pointerSelection = state.pointerSelection;

        if(!pointerSelection || pointerSelection.pointerId !== pointerEvent.pointerId || pointerSelection.isCancelled) {
            return;
        }

        const movedX = pointerEvent.clientX - pointerSelection.x;
        const movedY = pointerEvent.clientY - pointerSelection.y;

        if(Math.hypot(movedX, movedY) > 8) {
            pointerSelection.isCancelled = true;
            clearAssetTickerOptionHighlight(input);
        }
    });

    listbox.addEventListener("pointerup", event => {
        const pointerEvent = event as PointerEvent;
        const pointerSelection = state.pointerSelection;
        state.pointerSelection = undefined;

        if(!pointerSelection || pointerSelection.pointerId !== pointerEvent.pointerId || pointerSelection.isCancelled) {
            return;
        }

        const pointerTarget = document.elementFromPoint(pointerEvent.clientX, pointerEvent.clientY);
        const option = pointerTarget?.closest<HTMLElement>("[data-asset-ticker-option]");
        const optionIndex = option ? Number(option.dataset.optionIndex) : -1;
        const movedX = pointerEvent.clientX - pointerSelection.x;
        const movedY = pointerEvent.clientY - pointerSelection.y;

        if(optionIndex !== pointerSelection.optionIndex || Math.hypot(movedX, movedY) > 8) {
            return;
        }

        const selectedOption = state.options[optionIndex];

        if(selectedOption) {
            selectAssetTicker(input, selectedOption.ticker);
        }
    });

    listbox.addEventListener("pointercancel", () => {
        state.pointerSelection = undefined;
        clearAssetTickerOptionHighlight(input);
    });

    listbox.addEventListener("click", event => {
        const clickEvent = event as MouseEvent;

        if(clickEvent.detail !== 0) {
            return;
        }

        const option = (clickEvent.target as HTMLElement).closest<HTMLElement>("[data-asset-ticker-option]");
        const selectedOption = option ? state.options[Number(option.dataset.optionIndex)] : undefined;

        if(selectedOption) {
            selectAssetTicker(input, selectedOption.ticker);
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
    state.activeOptionSource = null;
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
 * Filters prefetched tickers and adds a selectable create action for an unmatched query.
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
    const searchQuery = input.value.trim();
    const query = searchQuery.toLocaleLowerCase();
    const seenTickers = new Set<string>();

    const matchingAssets = source
        ? Array.from(source.options)
            .map(option => ({
                ticker: option.value.trim(),
                label: (option.textContent ?? option.value).trim(),
            }))
            .filter(({ ticker, label }) => {
                const normalizedTicker = ticker.toLocaleLowerCase();
                const normalizedLabel = label.toLocaleLowerCase();

                if(!ticker || !normalizedLabel.includes(query) || seenTickers.has(normalizedTicker)) {
                    return false;
                }

                seenTickers.add(normalizedTicker);
                return true;
            })
        : [];

    state.options = matchingAssets.length > 0
        ? matchingAssets.map(({ ticker, label }) => ({ kind: "asset", label, ticker }))
        : searchQuery
            ? [{ kind: "create", label: "No matching assets - create new", ticker: searchQuery }]
            : [];
    state.activeOptionIndex = -1;
    state.activeOptionSource = null;
    state.pointerSelection = undefined;
    state.listbox.replaceChildren();
    input.removeAttribute("aria-activedescendant");

    state.options.forEach((autocompleteOption, index) => {
        const option = document.createElement("div");
        option.id = `${ state.listbox.id }-option-${ index }`;
        option.className = "asset-ticker-autocomplete__option";
        option.setAttribute("role", "option");
        option.setAttribute("aria-selected", "false");
        option.dataset.assetTickerOption = "";
        option.dataset.optionIndex = index.toString();
        option.dataset.optionKind = autocompleteOption.kind;
        option.textContent = autocompleteOption.label;
        state.listbox.append(option);
    });

    const hasMatchingTickers = matchingAssets.length > 0;
    const hasCreateOption = state.options.some(option => option.kind === "create");
    state.listbox.hidden = state.options.length === 0;
    state.status.textContent = "No matching assets. Select the create option to continue.";
    state.status.hidden = hasMatchingTickers || !hasCreateOption;
}

/**
 * Highlights one option and optionally scrolls it into view for keyboard navigation.
 *
 * @param input - Input that owns the active listbox.
 * @param optionIndex - Index of the option to highlight.
 * @param source - Interaction source used to clear transient pointer highlighting.
 * @param shouldScroll - Whether keyboard navigation should scroll the active option into view.
 *
 * @author GPT-6 Luna
 */
function highlightAssetTickerOption(
    input: HTMLInputElement,
    optionIndex: number,
    source: "keyboard" | "pointer",
    shouldScroll: boolean,
): void {

    const state = autocompleteStates.get(input);

    if(!state?.options.length) {
        return;
    }

    state.activeOptionIndex = Math.max(0, Math.min(optionIndex, state.options.length - 1));
    state.activeOptionSource = source;

    const options = state.listbox.querySelectorAll<HTMLElement>("[data-asset-ticker-option]");

    options.forEach((option, index) => {
        const isActive = index === state.activeOptionIndex;
        option.setAttribute("aria-selected", isActive.toString());

        if(isActive) {
            input.setAttribute("aria-activedescendant", option.id);

            if(shouldScroll) {
                option.scrollIntoView({ block: "nearest" });
            }
        }
    });
}

/** Clears the active option and its corresponding accessibility state.
 *
 * @author GPT-6 Luna
 */
function clearAssetTickerOptionHighlight(input: HTMLInputElement): void {

    const state = autocompleteStates.get(input);

    if(!state) {
        return;
    }

    state.activeOptionIndex = -1;
    state.activeOptionSource = null;
    input.removeAttribute("aria-activedescendant");

    state.listbox.querySelectorAll<HTMLElement>("[data-asset-ticker-option]").forEach(option => {
        option.setAttribute("aria-selected", "false");
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
    const assetTickerInput = row?.querySelector<HTMLInputElement>(`[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`);
    const assetNameInput = row?.querySelector<HTMLInputElement>("input[aria-label='Asset name']");

    if(!row?.id || !assetIdInput || !assetTickerInput || !assetNameInput) {
        console.error("Unable to resolve the asset row for the selected ticker.");
        return;
    }

    new AssetComposedColumnInput(row.id, assetIdInput.name, assetTickerInput.name, assetNameInput.name)
        .handleAssetActionButtonClick(ticker);
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
 * @author GPT-6 Luna
 */
class AssetComposedColumnInput {

    container: HTMLElement;
    autocompleteWrapper: HTMLElement;
    assetSearchInput: HTMLInputElement;
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

        const container = document.getElementById(containerId);

        if(!container) {
            throw new Error(`Unable to find asset row container '${ containerId }'.`);
        }

        this.container = container;
        this.autocompleteWrapper = container.querySelector<HTMLElement>(`[${ ASSET_TICKER_AUTOCOMPLETE_ATTRIBUTE }]`);
        this.assetSearchInput = container.querySelector<HTMLInputElement>(`[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`);
        this.assetIdInput = container.querySelector<HTMLInputElement>(`[name='${ assetIdHiddenFieldName }']`);
        this.assetTickerInput = container.querySelector<HTMLInputElement>(`[name='${ assetTickerFieldName }']`);
        this.assetActionButton = container.querySelector<HTMLButtonElement>("[data-asset-action-button]");
        this.newAssetTickerMessage = container.querySelector<HTMLDivElement>("[data-new-asset-ticker-message]");
        this.assetNameInput = container.querySelector<HTMLInputElement>(`[name='${ assetNameFieldName }']`);

        this.assetTickerExtraErrorMessageDiv =
            container.querySelector<HTMLDivElement>(`[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`);

        if(!this.autocompleteWrapper || !this.assetSearchInput || !this.assetIdInput || !this.assetTickerInput
            || !this.assetActionButton || !this.newAssetTickerMessage || !this.assetNameInput
            || !this.assetTickerExtraErrorMessageDiv) {
            throw new Error(`Asset row '${ containerId }' is missing required controls.`);
        }
    }

    /** Returns whether this row is waiting for a search or asset lookup.
     *
     * @author GPT-6 Luna
     */
    isInSearchMode(): boolean {
        return this.autocompleteWrapper.dataset.assetSelectionState === "search";
    }

    /** Returns whether the committed asset state can be cleared with the action button.
     *
     * @author GPT-6 Luna
     */
    isInResetMode(): boolean {
        const selectionState = this.autocompleteWrapper.dataset.assetSelectionState;
        return selectionState === "existing" || selectionState === "new";
    }

    /** Updates the search/reset styling and icon without changing row selection state.
     *
     * @author GPT-6 Luna
     */
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
    activateExistingAssetMode(asset: Asset): void {

        this.completeAssetLookup();
        this.assetActionButton.focus();
        this.autocompleteWrapper.dataset.assetSelectionState = "existing";
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
    }

    /**
     * Changes the row to require the name of an asset that is not yet stored.
     *
     * @author benizzio
     * @author GPT-6 Luna
     */
    activateNewAssetMode(ticker: string): void {

        this.completeAssetLookup();
        this.assetActionButton.focus();
        this.autocompleteWrapper.dataset.assetSelectionState = "new";
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
    }

    /**
     * Clears the current asset and restores editable ticker search mode.
     *
     * @author GPT-6 Luna
     * @author benizzio
     */
    resetToSearchMode() {

        closeAssetTickerAutocomplete(this.assetSearchInput);
        const autocompleteState = autocompleteStates.get(this.assetSearchInput);

        if(autocompleteState) {
            autocompleteState.isLookupPending = false;
        }

        this.autocompleteWrapper.dataset.assetSelectionState = "search";
        this.switchAssetActionButtonIdentity(ASSET_ACTION_BUTTON_IDENTITIES.search);
        this.assetActionButton.disabled = false;

        this.assetSearchInput.value = "";
        this.assetSearchInput.setCustomValidity("");
        this.assetSearchInput.readOnly = false;
        this.assetSearchInput.disabled = false;
        this.assetSearchInput.hidden = false;
        this.assetSearchInput.classList.remove("is-invalid");

        this.assetTickerInput.value = "";
        this.assetTickerInput.setCustomValidity("");
        this.assetTickerInput.readOnly = true;
        this.assetTickerInput.disabled = true;
        this.assetTickerInput.hidden = true;
        this.assetTickerInput.classList.remove("is-invalid");

        this.assetNameInput.value = "";
        this.assetNameInput.style.display = "none";
        this.assetNameInput.readOnly = false;
        this.assetNameInput.required = false;

        this.assetIdInput.value = "";

        this.newAssetTickerMessage.style.display = "none";
        this.clearSearchErrorFeedback();

        this.assetSearchInput.focus();
        openAssetTickerAutocomplete(this.assetSearchInput);
    }

    /** Clears search validation feedback before a new lookup attempt.
     *
     * @author GPT-6 Luna
     */
    clearSearchFieldValidation() {
        this.assetSearchInput.setCustomValidity("");
        this.assetSearchInput.classList.remove("is-invalid");
        this.clearSearchErrorFeedback();
    }

    /** Trims and validates a literal ticker or ID before it is sent to the lookup API.
     *
     * @author GPT-6 Luna
     */
    validateSearchUniqueIdentifier(): string {

        const assetUniqueIdentifier = this.assetSearchInput.value.trim();

        if(!assetUniqueIdentifier) {
            this.assetSearchInput.setCustomValidity("Required for search");
            this.assetSearchInput.reportValidity();
            return "";
        }

        if(assetUniqueIdentifier.length > 40) {
            this.assetSearchInput.setCustomValidity("Asset ticker cannot exceed 40 characters");
            this.assetSearchInput.reportValidity();
            return "";
        }

        if(!TICKER_CANDIDATE_PATTERN.test(assetUniqueIdentifier)) {
            this.assetSearchInput.setCustomValidity("Ticker may contain only letters, numbers, '-', '_', ':', and '.'");
            this.assetSearchInput.reportValidity();
            return "";
        }

        this.assetSearchInput.value = assetUniqueIdentifier;
        this.assetSearchInput.setCustomValidity("");
        return assetUniqueIdentifier;
    }

    /**
     * Searches for the typed ticker or resets a previously resolved asset.
     *
     * @author GPT-6 Luna
     * @author benizzio
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

                getAsset(this, searchUniqueIdentifier);
            }
        }
        else if(this.autocompleteWrapper.dataset.assetSelectionState === "pending") {
            return;
        }
        else if(this.isInResetMode()) {
            this.resetToSearchMode();
        }
    }

    /** Sets custom validity when the row has not committed an existing or new ticker.
     *
     * @author GPT-6 Luna
     */
    validateForPost(): void {
        const selectionState = this.autocompleteWrapper.dataset.assetSelectionState;

        if(selectionState === "search" || selectionState === "pending") {
            const message = selectionState === "pending"
                ? ASSET_LOOKUP_PENDING_ERROR_MESSAGE
                : ASSET_SELECTION_ERROR_MESSAGE;
            this.assetSearchInput.setCustomValidity(message);
            this.assetSearchInput.classList.add("is-invalid");
            this.assetTickerExtraErrorMessageDiv.textContent = message;
            this.assetTickerExtraErrorMessageDiv.style.display = "contents";
        }
        else {
            this.assetSearchInput.setCustomValidity("");
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

        const state = getAutocompleteState(this.assetSearchInput);

        if(state?.isLookupPending) {
            return false;
        }

        closeAssetTickerAutocomplete(this.assetSearchInput);

        if(state) {
            state.isLookupPending = true;
        }

        this.autocompleteWrapper.dataset.assetSelectionState = "pending";
        this.assetSearchInput.setCustomValidity(ASSET_LOOKUP_PENDING_ERROR_MESSAGE);
        this.assetSearchInput.readOnly = true;
        this.assetActionButton.disabled = true;
        return true;
    }

    /**
     * Releases lookup state after an existing asset or a new-asset path is resolved.
     *
     * @author GPT-6 Luna
     */
    completeAssetLookup(): void {

        const state = autocompleteStates.get(this.assetSearchInput);

        if(state) {
            state.isLookupPending = false;
        }

        this.assetActionButton.disabled = false;
        closeAssetTickerAutocomplete(this.assetSearchInput);
        this.assetSearchInput.setCustomValidity("");
    }

    /**
     * Restores editable search mode after a lookup fails for a reason other than a missing asset.
     *
     * @author GPT-6 Luna
     */
    restoreSearchModeAfterLookup(): void {

        this.completeAssetLookup();
        this.autocompleteWrapper.dataset.assetSelectionState = "search";
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

    /** Clears the validation feedback attached to the transient search control.
     *
     * @author GPT-6 Luna
     */
    clearSearchErrorFeedback(): void {
        this.assetSearchInput.classList.remove("is-invalid");
        this.assetTickerExtraErrorMessageDiv.textContent = "";
        this.assetTickerExtraErrorMessageDiv.style.display = "none";
    }
}

/** Validates every editable asset row in one form before native or HTMX submission.
 *
 * @param form - Form whose asset selection states should be checked.
 * @param reportFeedback - Whether to expose the first invalid row to the user.
 * @returns `true` when every editable asset row has a committed ticker.
 *
 * @author GPT-6 Luna
 */
function validateAssetRowsForPost(form: HTMLFormElement, reportFeedback: boolean): boolean {

    let isValid = true;
    let firstInvalidSearchInput: HTMLInputElement | null = null;

    form.querySelectorAll<HTMLElement>(`[${ ASSET_TICKER_AUTOCOMPLETE_ATTRIBUTE }]`).forEach(wrapper => {
        const searchInput = wrapper.querySelector<HTMLInputElement>(`[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`);
        const tickerInput = wrapper.querySelector<HTMLInputElement>(`[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`);

        const errorMessage = wrapper.parentElement?.querySelector<HTMLDivElement>(
            `[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`,
        );
        const selectionState = wrapper.dataset.assetSelectionState;

        const hasCommittedTicker = (selectionState === "existing" || selectionState === "new")
            && Boolean(tickerInput?.value.trim());

        if(hasCommittedTicker) {
            searchInput?.setCustomValidity("");
            searchInput?.classList.remove("is-invalid");
            return;
        }

        if(!searchInput) {
            isValid = false;
            return;
        }

        const message = selectionState === "pending"
            ? ASSET_LOOKUP_PENDING_ERROR_MESSAGE
            : ASSET_SELECTION_ERROR_MESSAGE;
        searchInput.setCustomValidity(message);
        searchInput.classList.add("is-invalid");

        if(errorMessage) {
            errorMessage.textContent = message;
            errorMessage.style.display = "contents";
        }
        firstInvalidSearchInput ??= searchInput;
        isValid = false;
    });

    if(!isValid && reportFeedback && firstInvalidSearchInput) {
        if(firstInvalidSearchInput.willValidate) {
            firstInvalidSearchInput.reportValidity();
        }
        else {
            firstInvalidSearchInput.focus();
        }
    }

    return isValid;
}

/** Installs form guards for native submits and HTMX requests containing editable asset rows.
 *
 * @author GPT-6 Luna
 */
function installAssetFormValidationGuards(): void {

    document.addEventListener("submit", event => {
        const form = event.target;

        if(!(form instanceof HTMLFormElement) || validateAssetRowsForPost(form, true)) {
            return;
        }

        event.preventDefault();
        event.stopImmediatePropagation();
    }, true);

    document.addEventListener("htmx:beforeRequest", event => {
        const htmxEvent = event as CustomEvent<{ elt?: Element }>;

        const requestElement = htmxEvent.detail?.elt
            ?? (event.target instanceof Element ? event.target : null);

        const form = requestElement instanceof HTMLFormElement
            ? requestElement
            : requestElement?.closest("form");

        if(form instanceof HTMLFormElement && !validateAssetRowsForPost(form, true)) {
            event.preventDefault();
        }
    });
}

/**
 * Resolves an asset ticker and updates its owning row with the API result.
 *
 * @author GPT-6 Luna
 */
function getAsset(rowAssetElements: AssetComposedColumnInput, searchUniqueIdentifier: string): void {

    api.getAsset(searchUniqueIdentifier)
        .then(responseBody => {

            if(!rowAssetElements.container.isConnected) {
                return;
            }

            if(api.isAPIErrorResponse(responseBody)) {
                if(responseBody.errorMessage === "Data not found") {
                    rowAssetElements.activateNewAssetMode(searchUniqueIdentifier);
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
            if(!rowAssetElements.container.isConnected) {
                return;
            }

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
 * `handleAssetTickerKeydown` directly to each generated search input. Existing row-specific
 * controllers continue to call the asset-action and validation methods with their own field names.
 *
 * @example
 * ```html
 * <input onfocus="AssetComposedColumnsInput.handleAssetTickerFocus(event)"
 *        oninput="AssetComposedColumnsInput.handleAssetTickerInput(event)"
 *        onkeydown="AssetComposedColumnsInput.handleAssetTickerKeydown(event)">
 * ```
 *
 * @author GPT-6 Luna
 * @author benizzio
 */
const AssetComposedColumnsInput = {

    /**
     * Opens suggestions when the editable asset search input receives focus.
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
     * Refreshes local asset suggestions after the search query changes.
     *
     * @param event - Input event emitted by the ticker field.
     *
     * @example oninput="AssetComposedColumnsInput.handleAssetTickerInput(event)"
     *
     * @author GPT-6 Luna
     */
    handleAssetTickerInput(event: Event): void {

        const input = event.target as HTMLInputElement;
        input.setCustomValidity("");
        input.classList.remove("is-invalid");
        const autocompleteWrapper = input.closest<HTMLElement>(`[${ ASSET_TICKER_AUTOCOMPLETE_ATTRIBUTE }]`);

        const errorMessage = autocompleteWrapper?.parentElement?.querySelector<HTMLDivElement>(
            `[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`,
        );

        if(errorMessage) {
            errorMessage.textContent = "";
            errorMessage.style.display = "none";
        }

        openAssetTickerAutocomplete(input);
        renderAssetTickerSuggestions(input);
    },

    /**
     * Handles autocomplete navigation, suggestion selection, and literal Enter-to-search fallback.
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

            if(state?.options.length) {
                event.preventDefault();

                const nextIndex = state.activeOptionIndex < 0
                    ? (event.key === "ArrowDown" ? 0 : state.options.length - 1)
                    : state.activeOptionIndex + (event.key === "ArrowDown" ? 1 : -1);
                highlightAssetTickerOption(inputElement, nextIndex, "keyboard", true);
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
                const selectedOption = state.options[state.activeOptionIndex];

                if(selectedOption) {
                    selectAssetTicker(inputElement, selectedOption.ticker);
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

    /**
     * Validates committed asset selections before submitting a form outside an HTMX field-validation event.
     *
     * @param form - Form containing one or more editable asset rows.
     * @returns `true` when every editable row is resolved or in confirmed new-asset mode.
     *
     * @example
     * ```ts
     * if(AssetComposedColumnsInput.validateFormBeforePost(form)) {
     *     form.requestSubmit();
     * }
     * ```
     *
     * @author GPT-6 Luna
     */
    validateFormBeforePost(form: HTMLFormElement): boolean {
        return validateAssetRowsForPost(form, true);
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

installAssetFormValidationGuards();

export default AssetComposedColumnsInput;
