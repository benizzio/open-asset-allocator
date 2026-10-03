/**
 * Owns autocomplete state, rendering, accessibility state, and open-widget lifecycle.
 *
 * @author GPT-6 Luna
 */

import {
    ARIA_ACTIVE_DESCENDANT_ATTRIBUTE,
    ARIA_CONTROLS_ATTRIBUTE,
    ARIA_EXPANDED_ATTRIBUTE,
    ARIA_SELECTED_ATTRIBUTE,
    ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE,
    ASSET_SEARCH_LISTBOX_ATTRIBUTE,
    ASSET_SEARCH_OPTION_CLASS,
    ASSET_SEARCH_OPTION_SELECTOR,
    ASSET_SEARCH_STATUS_ATTRIBUTE,
    ASSET_SEARCH_SUGGESTIONS_ATTRIBUTE,
    VISUALLY_HIDDEN_CLASS,
} from "./constants";
import { registerAutocompleteInteractions } from "./autocomplete-interactions";
import type {
    AssetSearchAutocompleteState,
    AssetTickerSelectionHandler,
    AutocompleteControllerActions,
} from "./types";

/**
 * Manages a single active asset-search autocomplete while retaining state for each input.
 *
 * @example
 * ```ts
 * const autocomplete = new AssetSearchAutocompleteController(resolveSelectedTicker);
 * autocomplete.open(searchInput);
 * ```
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
export class AssetSearchAutocompleteController implements AutocompleteControllerActions {
    private readonly autocompleteStates = new WeakMap<HTMLInputElement, AssetSearchAutocompleteState>();
    private activeAutocompleteInput: HTMLInputElement | null = null;
    private nextAutocompleteId = 0;
    private assetDatalistListenerRegistered = false;

    /** Creates the shared autocomplete controller and its ticker-selection callback.
     *
     * @author GPT-6 Luna
     */
    constructor(private readonly onTickerSelected: AssetTickerSelectionHandler) {}

    /** Lazily creates state and installs interaction listeners for a component search input.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    ensureState(input: HTMLInputElement): AssetSearchAutocompleteState | null {
        const currentState = this.autocompleteStates.get(input);

        if(currentState) {
            return currentState;
        }

        const state = this.createState(input);

        if(!state) {
            return null;
        }

        registerAutocompleteInteractions(input, state, this);
        this.autocompleteStates.set(input, state);
        return state;
    }

    /** Returns state only when it has already been initialized for the input.
     *
     * @author GPT-6 Luna
     */
    getState(input: HTMLInputElement): AssetSearchAutocompleteState | undefined {
        return this.autocompleteStates.get(input);
    }

    /** Opens one editable autocomplete and synchronizes its suggestions with the local datalist.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    open(input: HTMLInputElement): void {
        if(input.readOnly || input.disabled) {
            return;
        }

        const state = this.ensureState(input);

        if(!state || state.isLookupPending || state.isOpen) {
            return;
        }

        this.ensureDatalistLifecycleListener();

        if(this.activeAutocompleteInput && this.activeAutocompleteInput !== input) {
            this.close(this.activeAutocompleteInput);
        }

        this.activeAutocompleteInput = input;
        state.isOpen = true;
        state.suggestions.hidden = false;
        input.setAttribute(ARIA_EXPANDED_ATTRIBUTE, "true");

        state.outsidePointerHandler = event => {
            if(!state.wrapper.contains(event.target as Node)) {
                this.close(input);
            }
        };
        document.addEventListener("pointerdown", state.outsidePointerHandler, true);

        state.removalObserver = new MutationObserver(() => {
            if(!input.isConnected || input.getClientRects().length === 0) {
                this.close(input);
            }
        });

        state.removalObserver.observe(document.body, {
            childList: true,
            subtree: true,
            attributes: true,
            attributeFilter: ["class", "hidden", "style"],
        });

        this.render(input);
    }

    /** Closes one autocomplete without changing the asset search query.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    close(input: HTMLInputElement): void {
        const state = this.autocompleteStates.get(input);

        if(!state) {
            return;
        }

        state.isOpen = false;
        state.activeOptionIndex = -1;
        state.activeOptionSource = null;
        state.pointerSelection = undefined;
        state.suggestions.hidden = true;
        state.status.hidden = true;
        input.setAttribute(ARIA_EXPANDED_ATTRIBUTE, "false");
        input.removeAttribute(ARIA_ACTIVE_DESCENDANT_ATTRIBUTE);

        if(state.outsidePointerHandler) {
            document.removeEventListener("pointerdown", state.outsidePointerHandler, true);
        }
        state.outsidePointerHandler = undefined;
        state.removalObserver?.disconnect();
        state.removalObserver = undefined;

        if(this.activeAutocompleteInput === input) {
            this.activeAutocompleteInput = null;
        }
    }

    /** Renders filtered asset suggestions and a create action for an unmatched non-empty query.
     *
     * @author GPT-6 Luna
     * @author GPT-6.1 Sol
     * @author GPT-6 Sol
     */
    render(input: HTMLInputElement): void {
        const state = this.ensureState(input);

        if(!state?.isOpen) {
            return;
        }

        const sourceId = input.dataset.assetSearchSource;
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

        if(matchingAssets.length > 0) {
            state.options = matchingAssets.map(({ ticker, label }) => ({ kind: "asset", label, ticker }));
        }
        else if(searchQuery) {
            state.options = [{ kind: "create", label: "No matching assets - create new", ticker: searchQuery }];
        }
        else {
            state.options = [];
        }

        state.activeOptionIndex = -1;
        state.activeOptionSource = null;
        state.pointerSelection = undefined;
        state.listbox.replaceChildren();
        input.removeAttribute(ARIA_ACTIVE_DESCENDANT_ATTRIBUTE);

        state.options.forEach((autocompleteOption, index) => {
            const option = document.createElement("div");
            option.id = `${ state.listbox.id }-option-${ index }`;
            option.className = ASSET_SEARCH_OPTION_CLASS;
            option.setAttribute("role", "option");
            option.setAttribute(ARIA_SELECTED_ATTRIBUTE, "false");
            option.dataset.assetSearchOption = "";
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

    /** Highlights one option and optionally scrolls it into view for keyboard navigation.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    highlight(
        input: HTMLInputElement,
        optionIndex: number,
        source: "keyboard" | "pointer",
        shouldScroll: boolean,
    ): void {
        const state = this.autocompleteStates.get(input);

        if(!state?.options.length) {
            return;
        }

        state.activeOptionIndex = Math.max(0, Math.min(optionIndex, state.options.length - 1));
        state.activeOptionSource = source;

        const options = state.listbox.querySelectorAll<HTMLElement>(ASSET_SEARCH_OPTION_SELECTOR);

        options.forEach((option, index) => {
            const isActive = index === state.activeOptionIndex;
            option.setAttribute(ARIA_SELECTED_ATTRIBUTE, isActive.toString());

            if(isActive) {
                input.setAttribute(ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, option.id);

                if(shouldScroll) {
                    option.scrollIntoView({ block: "nearest" });
                }
            }
        });
    }

    /** Clears the active option and its corresponding accessibility state.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    clearHighlight(input: HTMLInputElement): void {
        const state = this.autocompleteStates.get(input);

        if(!state) {
            return;
        }

        state.activeOptionIndex = -1;
        state.activeOptionSource = null;
        input.removeAttribute(ARIA_ACTIVE_DESCENDANT_ATTRIBUTE);

        state.listbox.querySelectorAll<HTMLElement>(ASSET_SEARCH_OPTION_SELECTOR).forEach(option => {
            option.setAttribute(ARIA_SELECTED_ATTRIBUTE, "false");
        });
    }

    /** Writes a selected ticker to the input, closes the widget, and delegates its row lookup.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    selectOption(input: HTMLInputElement, ticker: string): void {
        const state = this.autocompleteStates.get(input);

        if(input.readOnly || state?.isLookupPending) {
            return;
        }

        input.value = ticker;
        this.close(input);
        this.onTickerSelected(input, ticker);
    }

    /** Ensures datalist refreshes rerender whichever autocomplete is currently open.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    ensureDatalistLifecycleListener(): void {
        if(this.assetDatalistListenerRegistered) {
            return;
        }

        document.addEventListener("htmx:afterSettle", event => {
            const htmxEvent = event as CustomEvent<{ elt?: Element }>;
            const datalist = document.getElementById("datalist-assets");

            if(!datalist || (htmxEvent.detail?.elt !== datalist && event.target !== datalist)) {
                return;
            }

            if(this.activeAutocompleteInput && this.autocompleteStates.get(this.activeAutocompleteInput)?.isOpen) {
                this.render(this.activeAutocompleteInput);
            }
        });

        this.assetDatalistListenerRegistered = true;
    }

    /** Builds an autocomplete state from its wrapper and accessibility controls.
     *
     * @author GPT-6 Luna
     * @author GPT-6 Sol
     */
    private createState(input: HTMLInputElement): AssetSearchAutocompleteState | null {
        const wrapper = input.closest<HTMLElement>(`[${ ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE }]`);
        const suggestions = wrapper?.querySelector<HTMLElement>(`[${ ASSET_SEARCH_SUGGESTIONS_ATTRIBUTE }]`);
        const listbox = wrapper?.querySelector<HTMLElement>(`[${ ASSET_SEARCH_LISTBOX_ATTRIBUTE }]`);
        const status = wrapper?.querySelector<HTMLElement>(`[${ ASSET_SEARCH_STATUS_ATTRIBUTE }]`);

        if(!wrapper || !suggestions || !listbox || !status) {
            return null;
        }

        const listboxId = `asset-search-listbox-${ ++this.nextAutocompleteId }`;
        listbox.id = listboxId;
        input.setAttribute(ARIA_CONTROLS_ATTRIBUTE, listboxId);

        const state: AssetSearchAutocompleteState = {
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

        status.classList.add(VISUALLY_HIDDEN_CLASS);
        return state;
    }
}
