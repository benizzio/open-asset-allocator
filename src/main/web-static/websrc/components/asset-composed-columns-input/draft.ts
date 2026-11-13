/**
 * Captures and restores editable asset selection state without performing a lookup.
 * Pending lookups are represented as editable searches in the independent draft.
 * @module components/asset-composed-columns-input/draft
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
import { renderCommittedAssetSelection } from "./committed-selection-renderer";
import {
    ASSET_ACTION_BUTTON_SELECTOR,
    ASSET_ID_INPUT_SELECTOR,
    ASSET_NAME_INPUT_SELECTOR,
    ASSET_SEARCH_INPUT_ATTRIBUTE,
    ASSET_TICKER_INPUT_ATTRIBUTE,
    NEW_ASSET_TICKER_MESSAGE_SELECTOR,
    AssetSelectionState,
} from "./constants";

/** Independent selection state for one asset row.
 * @author GPT-6.1 Terra
 */
export type AssetDraft =
    | { state: AssetSelectionState.EXISTING; id: string; ticker: string; name: string }
    | { state: AssetSelectionState.NEW; ticker: string; name: string }
    | { state: AssetSelectionState.SEARCH; query: string };

/** Reads a required control from a composed asset row.
 * @author GPT-6.1 Terra
 */
function input(row: HTMLElement, selector: string): HTMLInputElement {
    const element = row.querySelector(selector);

    if(!(element instanceof HTMLInputElement)) {
        throw new Error(`Missing asset input: ${ selector }`);
    }

    return element;
}

/** Captures static or interactive asset controls without triggering lookup or validation.
 * @param row - The source allocation row.
 * @returns An independent snapshot of its current selection.
 * @example `const asset = captureAssetDraft(sourceRow);`
 * @author GPT-6.1 Terra
 */
export function captureAssetDraft(row: HTMLElement): AssetDraft {
    const id = input(row, ASSET_ID_INPUT_SELECTOR).value;
    const wrapper = row.querySelector<HTMLElement>("[data-asset-search-autocomplete]");

    const name = row.querySelector<HTMLInputElement>(ASSET_NAME_INPUT_SELECTOR)?.value
        ?? row.querySelector<HTMLElement>("[data-asset-static-name]")?.textContent ?? "";

    if(!wrapper) {
        return {
            state: AssetSelectionState.EXISTING,
            id,
            ticker: row.querySelector<HTMLElement>("[data-asset-static-ticker]")?.textContent ?? "",
            name,
        };
    }

    const state = wrapper.dataset.assetSelectionState;

    if(state === AssetSelectionState.EXISTING || state === AssetSelectionState.NEW) {
        const ticker = input(row, `[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`).value;
        return state === AssetSelectionState.EXISTING
            ? { state: AssetSelectionState.EXISTING, id, ticker, name }
            : { state: AssetSelectionState.NEW, ticker, name };
    }

    return { state: AssetSelectionState.SEARCH, query: input(row, `[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`).value };
}

/** Restores an asset snapshot into a fresh interactive row without events, focus, or network work.
 * @param row - A newly rendered editable allocation row.
 * @param draft - The snapshot to restore.
 * @example `restoreAssetDraft(destinationRow, asset);`
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
export function restoreAssetDraft(row: HTMLElement, draft: AssetDraft): void {
    const wrapper = row.querySelector<HTMLElement>("[data-asset-search-autocomplete]");
    const button = row.querySelector<HTMLButtonElement>(ASSET_ACTION_BUTTON_SELECTOR);
    const message = row.querySelector<HTMLElement>(NEW_ASSET_TICKER_MESSAGE_SELECTOR);
    const search = input(row, `[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`);
    const ticker = input(row, `[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`);
    const name = input(row, ASSET_NAME_INPUT_SELECTOR);
    const id = input(row, ASSET_ID_INPUT_SELECTOR);

    if(!wrapper || !button || !message) {
        throw new Error("Missing asset selection controls.");
    }

    if(draft.state === AssetSelectionState.SEARCH) {
        wrapper.dataset.assetSelectionState = draft.state;
        search.value = draft.query;
        return;
    }

    renderCommittedAssetSelection({ wrapper, button, message, search, ticker, name, id }, {
        state: draft.state,
        id: draft.state === AssetSelectionState.EXISTING ? draft.id : "",
        ticker: draft.ticker,
        name: draft.name,
    });
}
