/**
 * Validates composed asset rows and installs native-submit and HTMX form guards.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */

import {
    ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE,
    ASSET_SEARCH_INPUT_ATTRIBUTE,
    ASSET_SELECTION_ERROR_MESSAGE,
    ASSET_LOOKUP_PENDING_ERROR_MESSAGE,
    ASSET_TICKER_INPUT_ATTRIBUTE,
    INVALID_INPUT_CLASS,
    TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE,
} from "./constants";
import { AssetSelectionState } from "./constants";

/** Validates one asset row and returns its search control when the row is invalid.
 *
 * @author GPT-6 Sol
 */
function validateAssetRowForPost(wrapper: HTMLElement): {
    isValid: boolean;
    invalidSearchInput: HTMLInputElement | null;
} {
    const searchInput = wrapper.querySelector<HTMLInputElement>(`[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`);
    const tickerInput = wrapper.querySelector<HTMLInputElement>(`[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`);

    const errorMessage = wrapper.parentElement?.querySelector<HTMLDivElement>(
        `[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }]`,
    );
    const selectionState = wrapper.dataset.assetSelectionState;

    const hasCommittedTicker = (selectionState === AssetSelectionState.EXISTING
        || selectionState === AssetSelectionState.NEW)
        && Boolean(tickerInput?.value.trim());

    if(hasCommittedTicker) {
        searchInput?.setCustomValidity("");
        searchInput?.classList.remove(INVALID_INPUT_CLASS);
        return { isValid: true, invalidSearchInput: null };
    }

    if(!searchInput) {
        return { isValid: false, invalidSearchInput: null };
    }

    const message = selectionState === AssetSelectionState.PENDING
        ? ASSET_LOOKUP_PENDING_ERROR_MESSAGE
        : ASSET_SELECTION_ERROR_MESSAGE;
    searchInput.setCustomValidity(message);
    searchInput.classList.add(INVALID_INPUT_CLASS);

    if(errorMessage) {
        errorMessage.textContent = message;
        errorMessage.style.display = "contents";
    }

    return { isValid: false, invalidSearchInput: searchInput };
}

/** Validates every editable asset row in one form before native or HTMX submission.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
export function validateAssetRowsForPost(form: HTMLFormElement, reportFeedback: boolean): boolean {
    let isValid = true;
    let firstInvalidSearchInput: HTMLInputElement | null = null;

    form.querySelectorAll<HTMLElement>(`[${ ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE }]`).forEach(wrapper => {
        const rowValidation = validateAssetRowForPost(wrapper);

        if(!rowValidation.isValid) {
            firstInvalidSearchInput ??= rowValidation.invalidSearchInput;
            isValid = false;
        }
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

let assetFormValidationGuardsInstalled = false;

/** Installs one-time guards for native submits and HTMX requests with unresolved asset rows.
 *
 * @author GPT-6 Luna
 */
export function installAssetFormValidationGuards(): void {
    if(assetFormValidationGuardsInstalled) {
        return;
    }

    assetFormValidationGuardsInstalled = true;

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
 * Marks a selected ticker invalid and displays its associated validation message.
 *
 * @param field - Selected ticker input to invalidate.
 * @param errorMessage - Validation message shown next to the ticker input.
 * @example `AssetComposedColumnsInput.invalidateSelectedAsset(tickerInput, "Ticker is invalid");`
 * @author GPT-6 Luna
 */
export function invalidateSelectedAsset(field: HTMLInputElement, errorMessage: string): void {
    field.classList.add(INVALID_INPUT_CLASS);

    const parentColumn = field.closest("td");

    const extraErrorMessageDiv = parentColumn.querySelector(
        `[${ TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE }="${ field.name }"]`,
    ) as HTMLDivElement;

    extraErrorMessageDiv.textContent = errorMessage;
    extraErrorMessageDiv.style.display = "contents";
    field.setCustomValidity(errorMessage);
    field.reportValidity();
}
