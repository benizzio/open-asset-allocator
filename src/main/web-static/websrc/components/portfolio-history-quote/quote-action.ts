/**
 * Owns the row-scoped HTMX lifecycle for fetching and applying a portfolio-history closing quote.
 *
 * @author GPT-6 Luna
 */

import { BigNumber } from "bignumber.js";
import { HtmxInfra } from "../../infra/htmx";
import type { AfterRequestEventDetail, RequestConfigEventDetail } from "../../infra/htmx";
import notifications from "../notifications";

const EXTERNAL_ASSET_KEYS_SELECTOR = "[data-external-asset-keys]";
const EXTERNAL_ASSET_SOURCE_SELECTOR = "[data-external-asset-key=\"source\"]";
const EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR = "[data-external-asset-key=\"exchangeId\"]";
const EXTERNAL_ASSET_TICKER_SELECTOR = "[data-external-asset-key=\"ticker\"]";
const QUOTE_INPUT_GROUP_SELECTOR = ".portfolio-history-market-price";
const QUOTE_ACTION_SELECTOR = "[data-quote-action]";
const QUOTE_ICON_SELECTOR = "[data-quote-icon]";
const QUOTE_SPINNER_SELECTOR = "[data-quote-spinner]";
const QUOTE_MARKET_PRICE_DISPLAY_SELECTOR = "[data-financial-input][data-financial-input-decimals]";
const QUOTE_MARKET_PRICE_RAW_SELECTOR = "input[type=\"hidden\"][name$=\"[assetMarketPrice]\"]";
const QUOTE_QUANTITY_SELECTOR = "[name$=\"[assetQuantity]\"]";
const QUOTE_TOTAL_MARKET_VALUE_RAW_SELECTOR = "input[type=\"hidden\"][name$=\"[totalMarketValue]\"]";
const QUOTE_RESPONSE_ERROR_MESSAGE = "The latest closing price response was invalid.";
const QUOTE_INPUT_ERROR_MESSAGE = "The market price input is unavailable.";
const QUOTE_PATH_ERROR_MESSAGE = "The selected external-asset identifiers could not be encoded.";

/** Recalculates the existing quantity × market-price total for one allocation row.
 *
 * @param allocationIndex - Index encoded in the allocation row ID.
 * @param observationTimestampId - Observation ID encoded in the allocation row ID.
 * @example `recalculatePortfolioAllocation(0, 12)`
 * @author GPT-6 Luna
 */
export type RecalculatePortfolioAllocation = (
    allocationIndex: number,
    observationTimestampId: number,
) => void;

/** Public operations for installing quote listeners and synchronizing row action visibility.
 *
 * @example
 * ```ts
 * const quoteAction = createPortfolioHistoryQuoteAction(recalculateAllocation);
 * quoteAction.bindRow(row);
 * quoteAction.updateButtonVisibility(row);
 * ```
 *
 * @author GPT-6 Luna
 */
export type PortfolioHistoryQuoteAction = {
    /** Binds idempotent HTMX lifecycle listeners to the quote button owned by one portfolio-history row.
     *
     * @param row - Portfolio-history allocation row containing the quote action.
     * @example `quoteAction.bindRow(row)`
     *
     * @author GPT-6 Luna
     */
    bindRow(row: HTMLTableRowElement): void;
    /** Shows the row's quote action only when its persisted asset and provider keys are complete.
     *
     * @author GPT-6 Luna
     */
    updateButtonVisibility(row: HTMLTableRowElement): void;
};

/** Creates the quote lifecycle used by one portfolio-history management component.
 *
 * The returned controller derives the URL and eligibility from the triggering row at request time, ignores stale
 * responses after row replacement or selection changes, and delegates total recalculation to the existing row logic.
 *
 * @param recalculateAllocation - Existing portfolio-history quantity × market-price calculation.
 * @returns A controller for binding one row's quote action and refreshing its visibility.
 * @example
 * ```ts
 * const quoteAction = createPortfolioHistoryQuoteAction((index, observationId) => {
 *     portfolioHistoryManagement.handleInputQuantityOrMarketPrice(index, observationId);
 * });
 * quoteAction.bindRow(row);
 * quoteAction.updateButtonVisibility(row);
 * ```
 * @author GPT-6 Luna
 */
export function createPortfolioHistoryQuoteAction(
    recalculateAllocation: RecalculatePortfolioAllocation,
): PortfolioHistoryQuoteAction {
    /** Exact row and persisted provider identity captured when HTMX prepares the GET request.
     *
     * @author GPT-6 Luna
     */
    type QuoteRequestSnapshot = {
        row: HTMLTableRowElement;
        assetId: string;
        source: string;
        exchangeId: string;
        ticker: string;
        generation: number;
        observationTimestampId: number;
        allocationIndex: number;
    };

    /** Associates the exact HTMX request and row snapshot with its initiating quote button.
     *
     * @author GPT-6 Luna
     */
    type ActiveQuoteRequest = {
        snapshot: QuoteRequestSnapshot;
        requestConfig: RequestConfigEventDetail;
        xhr?: XMLHttpRequest;
        sent: boolean;
    };

    const activeQuoteRequests = new WeakMap<HTMLButtonElement, ActiveQuoteRequest>();
    const boundQuoteButtons = new WeakSet<HTMLButtonElement>();

    /** Reads only this row's persisted asset ID, enabled first-provider keys, and selection generation.
     *
     * @author GPT-6 Luna
     */
    function getQuoteRequestSnapshot(row: HTMLTableRowElement): QuoteRequestSnapshot | null {
        const assetIdInput = row.querySelector<HTMLInputElement>("input[type=\"hidden\"][name$=\"[assetId]\"]");
        const fieldset = row.querySelector<HTMLFieldSetElement>(EXTERNAL_ASSET_KEYS_SELECTOR);
        const sourceInput = fieldset?.querySelector<HTMLInputElement>(EXTERNAL_ASSET_SOURCE_SELECTOR);
        const exchangeIdInput = fieldset?.querySelector<HTMLInputElement>(EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR);
        const tickerInput = fieldset?.querySelector<HTMLInputElement>(EXTERNAL_ASSET_TICKER_SELECTOR);
        const allocationCoordinates = /^portfolio-history-management-form-(\d+)-row-(\d+)$/.exec(row.id);
        const generation = Number(row.dataset.assetSelectionGeneration);

        if(!assetIdInput) {
            return null;
        }

        const assetId = assetIdInput.value.trim();

        if(!assetId || assetIdInput.value !== assetId) {
            return null;
        }

        if(!fieldset || fieldset.disabled) {
            return null;
        }

        if(!sourceInput || !exchangeIdInput || !tickerInput) {
            return null;
        }

        const source = sourceInput.value.trim();
        const exchangeId = exchangeIdInput.value.trim();
        const ticker = tickerInput.value.trim();

        if(!source || !exchangeId || !ticker) {
            return null;
        }

        if(sourceInput.value !== source || exchangeIdInput.value !== exchangeId || tickerInput.value !== ticker) {
            return null;
        }

        if(!Number.isSafeInteger(generation)
            || generation < 0
            || row.dataset.assetSelectionGeneration !== String(generation)
            || !allocationCoordinates) {
            return null;
        }

        const observationTimestampId = Number(allocationCoordinates[1]);
        const allocationIndex = Number(allocationCoordinates[2]);

        if(!Number.isSafeInteger(observationTimestampId) || !Number.isSafeInteger(allocationIndex)) {
            return null;
        }

        return {
            row,
            assetId,
            source,
            exchangeId,
            ticker,
            generation,
            observationTimestampId,
            allocationIndex,
        };
    }

    /** Updates one button and its containing input group from that row's current eligibility.
     *
     * @author GPT-6 Luna
     */
    function updateButtonVisibility(row: HTMLTableRowElement): void {
        const quoteButton = row.querySelector<HTMLButtonElement>(QUOTE_ACTION_SELECTOR);

        if(!quoteButton) {
            return;
        }

        const isAvailable = getQuoteRequestSnapshot(row) !== null;
        quoteButton.hidden = !isAvailable;

        quoteButton.closest<HTMLElement>(QUOTE_INPUT_GROUP_SELECTOR)
            ?.classList.toggle("has-quote-action", isAvailable);
    }

    /** Confirms that the response still belongs to its connected row and unchanged selection generation.
     *
     * @author GPT-6 Luna
     */
    function isSnapshotCurrent(quoteButton: HTMLButtonElement, snapshot: QuoteRequestSnapshot): boolean {
        const currentSnapshot = getQuoteRequestSnapshot(snapshot.row);

        if(!snapshot.row.isConnected || !quoteButton.isConnected || !snapshot.row.contains(quoteButton)) {
            return false;
        }

        if(quoteButton.closest("tr") !== snapshot.row || !currentSnapshot) {
            return false;
        }

        const sameProviderKeys = currentSnapshot.source === snapshot.source
            && currentSnapshot.exchangeId === snapshot.exchangeId
            && currentSnapshot.ticker === snapshot.ticker;

        return currentSnapshot.assetId === snapshot.assetId
            && currentSnapshot.generation === snapshot.generation
            && sameProviderKeys;
    }

    /** Restores accessible loading indicators after a request finishes or is stopped before sending.
     *
     * @author GPT-6 Luna
     */
    function restoreLoadingState(quoteButton: HTMLButtonElement): void {
        quoteButton.removeAttribute("aria-busy");

        const icon = quoteButton.querySelector<HTMLElement>(QUOTE_ICON_SELECTOR);
        const spinner = quoteButton.querySelector<HTMLElement>(QUOTE_SPINNER_SELECTOR);

        if(icon) {
            icon.hidden = false;
        }

        if(spinner) {
            spinner.hidden = true;
        }
    }

    /** Removes the in-flight snapshot and restores its action button.
     *
     * @author GPT-6 Luna
     */
    function releaseQuoteRequest(quoteButton: HTMLButtonElement, request: ActiveQuoteRequest): void {
        if(activeQuoteRequests.get(quoteButton) !== request) {
            return;
        }

        activeQuoteRequests.delete(quoteButton);
        restoreLoadingState(quoteButton);
    }

    /** Releases only the exact quote request that HTMX halted during validation or path verification.
     *
     * @author GPT-6 Luna
     */
    function releaseUnsentQuoteRequest(event: Event, quoteButton: HTMLButtonElement): void {
        const detail = (event as CustomEvent<RequestConfigEventDetail>).detail;
        const request = activeQuoteRequests.get(quoteButton);

        if(request?.requestConfig && detail === request.requestConfig) {
            releaseQuoteRequest(quoteButton, request);
        }
    }

    /** Confirms an afterRequest event belongs to the request currently tracked for this quote button.
     *
     * @author GPT-6 Luna
     */
    function isOwnedAfterRequest(
        request: ActiveQuoteRequest | undefined,
        detail: AfterRequestEventDetail,
    ): request is ActiveQuoteRequest {
        if(!request?.requestConfig || detail.requestConfig !== request.requestConfig) {
            return false;
        }

        return !request.xhr || detail.xhr === request.xhr;
    }

    /** Marks the exact quote request sent and activates its loading indicators.
     *
     * @author GPT-6 Luna
     */
    function handleBeforeSend(event: Event, quoteButton: HTMLButtonElement): void {
        const detail = (event as CustomEvent<AfterRequestEventDetail>).detail;
        const request = activeQuoteRequests.get(quoteButton);

        if(!detail || !request?.requestConfig || detail.requestConfig !== request.requestConfig) {
            return;
        }

        request.sent = true;
        request.xhr = detail.xhr;
        quoteButton.setAttribute("aria-busy", "true");

        const icon = quoteButton.querySelector<HTMLElement>(QUOTE_ICON_SELECTOR);
        const spinner = quoteButton.querySelector<HTMLElement>(QUOTE_SPINNER_SELECTOR);

        if(icon) {
            icon.hidden = true;
        }

        if(spinner) {
            spinner.hidden = false;
        }
    }

    /** Sets the quote endpoint and captures ownership before HTMX issues the GET.
     *
     * @author GPT-6 Luna
     */
    function handleRequestConfiguration(event: Event, quoteButton: HTMLButtonElement): void {
        const requestConfig = (event as CustomEvent<RequestConfigEventDetail>).detail;

        if(event.target !== quoteButton
            || requestConfig?.elt !== quoteButton
            || requestConfig.verb.toLowerCase() !== "get") {
            return;
        }

        if(activeQuoteRequests.has(quoteButton)) {
            event.preventDefault();
            return;
        }

        const row = quoteButton.closest<HTMLTableRowElement>("tr");
        const snapshot = row ? getQuoteRequestSnapshot(row) : null;

        if(!snapshot) {
            if(row) {
                updateButtonVisibility(row);
            }

            event.preventDefault();
            return;
        }

        let requestPath: string;

        try {
            requestPath = `/api/asset/${ encodeURIComponent(snapshot.assetId) }/external-asset/`
                + `${ encodeURIComponent(snapshot.source) }/${ encodeURIComponent(snapshot.exchangeId) }/`
                + `${ encodeURIComponent(snapshot.ticker) }/quote`;
        }
        catch {
            event.preventDefault();
            notifications.notifyError(new Error(QUOTE_PATH_ERROR_MESSAGE));
            return;
        }

        const request: ActiveQuoteRequest = { snapshot, requestConfig, sent: false };
        activeQuoteRequests.set(quoteButton, request);
        requestConfig.path = requestPath;

        // HTMX completes configuration, beforeRequest, and beforeSend synchronously. If any later listener
        // cancels the request, this checkpoint releases its prepared state without depending on listener order.
        queueMicrotask(() => {
            if(activeQuoteRequests.get(quoteButton) === request && !request.sent) {
                releaseQuoteRequest(quoteButton, request);
            }
        });
    }

    /** Validates the entire response before returning a finite non-negative decimal quote.
     *
     * @author GPT-6 Luna
     */
    function parseLatestClosingQuote(responseText: string, snapshot: QuoteRequestSnapshot): BigNumber | null {
        let response: unknown;

        try {
            response = JSON.parse(responseText) as unknown;
        }
        catch {
            return null;
        }

        if(typeof response !== "object" || response === null || Array.isArray(response)) {
            return null;
        }

        const quoteResponse = response as Record<string, unknown>;

        if(typeof quoteResponse.ticker !== "string" || quoteResponse.ticker !== snapshot.ticker) {
            return null;
        }

        if(typeof quoteResponse.exchangeId !== "string" || quoteResponse.exchangeId !== snapshot.exchangeId) {
            return null;
        }

        if(typeof quoteResponse.currency !== "string" || !quoteResponse.currency.trim()) {
            return null;
        }

        if(typeof quoteResponse.lastCloseDate !== "string" || !quoteResponse.lastCloseDate.trim()) {
            return null;
        }

        if(!Number.isFinite(Date.parse(quoteResponse.lastCloseDate))) {
            return null;
        }

        if(typeof quoteResponse.lastCloseQuote !== "string"
            || !/^(?:\d+(?:\.\d*)?|\.\d+)$/.test(quoteResponse.lastCloseQuote)) {
            return null;
        }

        try {
            const quote = new BigNumber(quoteResponse.lastCloseQuote);

            return quote.isFinite() && !quote.isNegative() ? quote : null;
        }
        catch {
            return null;
        }
    }

    /** Writes the normalized raw price and invokes the existing row total calculation without changing quantity.
     *
     * @author GPT-6 Luna
     */
    function applyLatestClosingQuote(
        quoteButton: HTMLButtonElement,
        snapshot: QuoteRequestSnapshot,
        quote: BigNumber,
    ): boolean {
        if(!isSnapshotCurrent(quoteButton, snapshot)) {
            return false;
        }

        const row = snapshot.row;
        const displayInput = row.querySelector<HTMLInputElement>(QUOTE_MARKET_PRICE_DISPLAY_SELECTOR);
        const rawInput = row.querySelector<HTMLInputElement>(QUOTE_MARKET_PRICE_RAW_SELECTOR);
        const quantityInput = row.querySelector<HTMLInputElement>(QUOTE_QUANTITY_SELECTOR);
        const totalMarketValueInput = row.querySelector<HTMLInputElement>(QUOTE_TOTAL_MARKET_VALUE_RAW_SELECTOR);
        const decimalPlaces = Number(displayInput?.dataset.financialInputDecimals);

        if(!displayInput || !rawInput) {
            return false;
        }

        if(!quantityInput || !totalMarketValueInput) {
            return false;
        }

        if(!Number.isSafeInteger(decimalPlaces) || decimalPlaces < 0) {
            return false;
        }

        const normalizedQuote = quote.decimalPlaces(decimalPlaces, BigNumber.ROUND_HALF_UP).toFixed(decimalPlaces);

        if(document.activeElement === displayInput) {
            displayInput.blur();
        }

        rawInput.value = normalizedQuote;
        recalculateAllocation(snapshot.allocationIndex, snapshot.observationTimestampId);
        return true;
    }

    /** Handles stale responses and ensures failed current requests reach the global HTMX error handler.
     *
     * @author GPT-6 Luna
     */
    function handleAfterRequest(event: Event, quoteButton: HTMLButtonElement): void {
        const detail = (event as CustomEvent<AfterRequestEventDetail>).detail;
        const request = activeQuoteRequests.get(quoteButton);

        if(!detail || !isOwnedAfterRequest(request, detail)) {
            return;
        }

        const { snapshot } = request;
        const isCurrent = isSnapshotCurrent(quoteButton, snapshot);

        try {
            if(!detail.successful) {
                if(!isCurrent) {
                    HtmxInfra.suppressAfterRequestErrorNotification(detail.xhr);
                }

                return;
            }

            if(!isCurrent) {
                return;
            }

            const quote = parseLatestClosingQuote(detail.xhr.responseText, snapshot);

            if(!quote) {
                notifications.notifyError(new Error(QUOTE_RESPONSE_ERROR_MESSAGE));
            }
            else if(!applyLatestClosingQuote(quoteButton, snapshot, quote)) {
                notifications.notifyError(new Error(QUOTE_INPUT_ERROR_MESSAGE));
            }
        }
        finally {
            releaseQuoteRequest(quoteButton, request);
        }
    }

    /** Binds one row's quote button once, keeping events isolated from other HTMX components.
     *
     * @author GPT-6 Luna
     */
    function bindRow(row: HTMLTableRowElement): void {
        const quoteButton = row.querySelector<HTMLButtonElement>(QUOTE_ACTION_SELECTOR);

        if(!quoteButton || boundQuoteButtons.has(quoteButton)) {
            return;
        }

        boundQuoteButtons.add(quoteButton);
        quoteButton.addEventListener("htmx:configRequest", event => handleRequestConfiguration(event, quoteButton));
        quoteButton.addEventListener("htmx:beforeSend", event => handleBeforeSend(event, quoteButton));
        quoteButton.addEventListener("htmx:validation:halted", event => releaseUnsentQuoteRequest(event, quoteButton));
        quoteButton.addEventListener("htmx:invalidPath", event => releaseUnsentQuoteRequest(event, quoteButton));
        quoteButton.addEventListener("htmx:afterRequest", event => handleAfterRequest(event, quoteButton));
    }

    return {
        /** Binds stable, idempotent lifecycle listeners to the quote button in one row.
         *
         * @author GPT-6 Luna
         */
        bindRow(row: HTMLTableRowElement): void {
            bindRow(row);
        },

        /** Updates only the quote action inside the supplied allocation row.
         *
         * @author GPT-6 Luna
         */
        updateButtonVisibility(row: HTMLTableRowElement): void {
            updateButtonVisibility(row);
        },
    };
}
