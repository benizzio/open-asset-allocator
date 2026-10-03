/**
 * Internal contracts shared by asset autocomplete and row controllers.
 *
 * @author GPT-6 Luna
 */

/** Describes one selectable asset suggestion or missing-ticker create action.
 *
 * @author GPT-6 Luna
 */
export type AssetSearchAutocompleteOption = {
    kind: "asset" | "create";
    label: string;
    ticker: string;
};

/** Stores DOM references and transient keyboard and pointer state for one autocomplete input.
 *
 * @author GPT-6 Luna
 */
export interface AssetSearchAutocompleteState {
    wrapper: HTMLElement;
    suggestions: HTMLElement;
    listbox: HTMLElement;
    status: HTMLElement;
    options: AssetSearchAutocompleteOption[];
    activeOptionIndex: number;
    activeOptionSource: "keyboard" | "pointer" | null;
    isOpen: boolean;
    isLookupPending: boolean;
    pointerSelection?: { optionIndex: number; pointerId: number; x: number; y: number; isCancelled: boolean };
    outsidePointerHandler?: (event: PointerEvent) => void;
    removalObserver?: MutationObserver;
}

/** Defines the operations autocomplete interactions can delegate to their shared controller.
 *
 * @author GPT-6 Luna
 */
export interface AutocompleteControllerActions {
    ensureState(input: HTMLInputElement): AssetSearchAutocompleteState | null;
    getState(input: HTMLInputElement): AssetSearchAutocompleteState | undefined;
    open(input: HTMLInputElement): void;
    close(input: HTMLInputElement): void;
    highlight(
        input: HTMLInputElement,
        optionIndex: number,
        source: "keyboard" | "pointer",
        shouldScroll: boolean,
    ): void;
    clearHighlight(input: HTMLInputElement): void;
    selectOption(input: HTMLInputElement, ticker: string): void;
}

/** Receives a selected ticker so the owning asset row can perform its usual lookup.
 *
 * @author GPT-6 Luna
 */
export type AssetTickerSelectionHandler = (input: HTMLInputElement, ticker: string) => void;
