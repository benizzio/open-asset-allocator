/**
 * Coordinates cursor-based HTMX list requests without replacing committed records during append operations.
 *
 * @module components/paginated-list
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */

import htmx from "htmx.org";
import { HtmxInfra, type BeforeSwapEventDetail } from "../../infra/htmx";
import { focusOnAcceptedRecord, renderControls } from "./controls";

const PAGINATED_LIST_ATTRIBUTE = "data-paginated-list";
const LOAD_MORE_ATTRIBUTE = "data-paginated-list-more";
const LOAD_MORE_EVENT = "paginated-list-more";
const PAGE_CURSOR_PARAMETER = "pageCursor";

/** Selects whether a list request replaces the current page or appends another page.
 *
 * Use `PaginationMode.Initial` for first loads and refreshes, and `PaginationMode.Append`
 * for cursor-based continuation requests.
 * @example `mode === PaginationMode.Initial` identifies a request that resets the list.
 * @author GPT-6 Sol
 */
export enum PaginationMode {
    Initial = "initial",
    Append = "append",
}

type PageEnvelope<T> = {
    items: T[];
    nextCursor: string | null;
};

type ListState = {
    generation: number;
    portfolioId: string;
    cursor: string | null;
    recordIds: Set<string>;
    initialized: boolean;
    loading: boolean;
    nextMode?: PaginationMode;
    retryMode?: PaginationMode;
    controlRequestPending: boolean;
    pendingFocusButton?: HTMLButtonElement;
    preparingRequest?: RequestContext;
    activeRequest?: RequestContext;
};

type RequestContext = {
    owner: HTMLElement;
    generation: number;
    portfolioId: string;
    mode: PaginationMode;
    cursor: string;
    acceptedIds: string[];
    nextCursor: string | null;
    prepared: boolean;
    rejected: boolean;
    stale: boolean;
    committed: boolean;
    focusButton?: HTMLButtonElement;
    focusChanged: boolean;
    xhr?: XMLHttpRequest;
};

const states = new WeakMap<HTMLElement, ListState>();
const requests = new WeakMap<XMLHttpRequest, RequestContext>();
const focusTrackedRequests = new Set<RequestContext>();
let initialized = false;

type HistoryPageItem = { observationTimestamp?: { id?: number } };
type AllocationPlanPageItem = { id?: number };

/** Returns the state associated with a list owner, creating its initial state when needed. */
function getState(owner: HTMLElement): ListState {
    let state = states.get(owner);

    if(!state) {
        state = {
            generation: 0,
            portfolioId: getPortfolioId(),
            cursor: null,
            recordIds: new Set<string>(),
            initialized: false,
            loading: false,
            controlRequestPending: false,
        };
        states.set(owner, state);
    }

    return state;
}

/** Reads the active portfolio identity from the page's shared context field. */
function getPortfolioId(): string {
    return document.querySelector<HTMLInputElement>("[name='portfolioId']")?.value ?? "";
}

/** Resolves a list owner from an HTMX event target. */
function getOwner(target: EventTarget | null): HTMLElement | undefined {
    if(!(target instanceof Element)) {
        return undefined;
    }

    const owner = target.closest<HTMLElement>(`[${ PAGINATED_LIST_ATTRIBUTE }]`);
    return owner ?? undefined;
}

/** Renders owner-scoped controls after a pagination state transition. */
function updateControls(owner: HTMLElement, state: ListState, message = "", isError = false): void {
    renderControls(owner, {
        initialized: state.initialized,
        cursor: state.cursor,
        loading: state.loading,
        retryAvailable: state.retryMode !== undefined,
        controlRequestPending: state.controlRequestPending,
    }, message, isError);
}

/** Marks a pending keyboard continuation abandoned when focus moves elsewhere. */
function trackPaginationFocusChange(event: FocusEvent): void {
    const target = event.target;

    for(const context of focusTrackedRequests) {
        if(target !== context.focusButton && target !== document.body) {
            context.focusChanged = true;
        }
    }
}

/** Returns whether a response still belongs to the connected list generation and portfolio. */
function isCurrentRequest(context: RequestContext): boolean {
    const state = states.get(context.owner);
    return context.owner.isConnected
        && !!state
        && state.generation === context.generation
        && getPortfolioId() === context.portfolioId;
}

/** Stores request metadata against the XHR before its response can arrive. */
function bindRequestToXhr(event: CustomEvent): void {
    const detail = event.detail as { elt?: Element; xhr?: XMLHttpRequest };
    const requestElement = detail.elt ?? event.target;
    const owner = getOwner(requestElement);

    if(!owner || requestElement !== owner || !detail.xhr) {
        return;
    }

    const state = getState(owner);
    const context = state.preparingRequest;

    if(!context) {
        return;
    }

    context.xhr = detail.xhr;
    state.activeRequest = context;
    state.preparingRequest = undefined;

    if(context.focusButton) {
        focusTrackedRequests.add(context);
    }
    requests.set(detail.xhr, context);
}

/** Adds the explicit initial or continuation cursor to a paginated HTMX request.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function configureRequest(event: CustomEvent): void {
    const requestDetail = event.detail as { elt?: Element };
    const requestElement = requestDetail.elt ?? event.target;
    const owner = getOwner(requestElement);

    if(!owner || requestElement !== owner) {
        return;
    }

    const detail = event.detail as {
        parameters?: Record<string, string | string[]>;
        verb?: string;
    };

    if(detail.verb?.toLowerCase() !== "get" || !detail.parameters) {
        return;
    }

    const state = getState(owner);
    const portfolioId = getPortfolioId();
    let requestedMode = state.nextMode ?? PaginationMode.Initial;
    state.nextMode = undefined;

    const portfolioChanged = portfolioId !== state.portfolioId;

    if(portfolioChanged) {
        requestedMode = PaginationMode.Initial;
        state.portfolioId = portfolioId;
        state.cursor = null;
        state.recordIds.clear();
        state.initialized = false;
    }

    if(requestedMode === PaginationMode.Append && (state.loading || !state.initialized || state.cursor === null)) {
        event.preventDefault();
        return;
    }

    if(portfolioChanged || requestedMode === PaginationMode.Initial) {
        state.generation++;
    }

    const cursor = requestedMode === PaginationMode.Append ? state.cursor : "";

    const context: RequestContext = {
        owner,
        generation: state.generation,
        portfolioId,
        mode: requestedMode,
        cursor,
        acceptedIds: [],
        nextCursor: null,
        prepared: false,
        rejected: false,
        stale: false,
        committed: false,
        focusButton: state.pendingFocusButton,
        focusChanged: false,
    };

    state.pendingFocusButton = undefined;
    state.preparingRequest = context;
    state.loading = true;
    detail.parameters[PAGE_CURSOR_PARAMETER] = cursor;
    updateControls(owner, state);
}

/** Releases the active request after a completed append or full list refresh. */
function finishRequest(context: RequestContext, message?: string, isError = false): void {
    const state = states.get(context.owner);

    if(state?.activeRequest !== context || !isCurrentRequest(context)) {
        return;
    }

    state.loading = false;
    state.controlRequestPending = false;
    state.pendingFocusButton = undefined;
    state.activeRequest = undefined;
    state.preparingRequest = undefined;
    focusTrackedRequests.delete(context);
    updateControls(context.owner, state, message, isError);
}

/** Commits a validated page only after HTMX has completed its DOM swap and settle lifecycle.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function commitSettledPage(event: CustomEvent): void {
    const detail = event.detail as { xhr?: XMLHttpRequest };
    const context = detail.xhr ? requests.get(detail.xhr) : undefined;

    if(!context || !context.prepared || context.rejected || context.stale) {
        return;
    }

    if(!isCurrentRequest(context)) {
        context.stale = true;
        focusTrackedRequests.delete(context);
        return;
    }

    const state = states.get(context.owner);

    if(state?.activeRequest !== context) {
        return;
    }

    if(context.mode === PaginationMode.Initial) {
        state.recordIds = new Set(context.acceptedIds);
    }
    else {
        context.acceptedIds.forEach(id => state.recordIds.add(id));
    }

    state.cursor = context.nextCursor;
    state.portfolioId = context.portfolioId;
    state.initialized = true;
    state.retryMode = undefined;
    context.committed = true;

    const shouldMoveFocus = shouldMoveFocusAfterTerminalAppend(context);

    let message = "";

    if(context.nextCursor === null) {
        message = state.recordIds.size === 0 ? "No records found." : "";
    }

    finishRequest(context, message);

    if(shouldMoveFocus) {
        focusOnAcceptedRecord(context.owner, context.acceptedIds);
    }
}

/** Returns whether a terminal append should restore focus to its first persisted record. */
function shouldMoveFocusAfterTerminalAppend(context: RequestContext): boolean {
    if(context.mode !== PaginationMode.Append || context.nextCursor !== null) {
        return false;
    }

    if(!context.focusButton || context.focusChanged) {
        return false;
    }

    return document.activeElement === context.focusButton
        || document.activeElement === document.body
        || document.activeElement === document.documentElement;
}

/** Releases failed or malformed requests without consuming their continuation cursor. */
function handleRequestCompletion(event: CustomEvent): void {
    const detail = event.detail as { successful?: boolean; xhr?: XMLHttpRequest };
    const context = detail.xhr ? requests.get(detail.xhr) : undefined;

    if(!context) {
        return;
    }

    if(!isCurrentRequest(context)) {
        context.stale = true;
        focusTrackedRequests.delete(context);

        if(detail.xhr && !detail.successful) {
            HtmxInfra.suppressAfterRequestErrorNotification(detail.xhr);
        }

        return;
    }

    if(context.committed) {
        return;
    }

    if(detail.successful && context.prepared && !context.rejected) {
        // HTMX emits afterRequest before afterSettle; keep the request active until the swap commits.
        return;
    }

    const state = states.get(context.owner);

    if(state?.activeRequest !== context) {
        return;
    }

    state.retryMode = context.mode;
    finishRequest(context, "Could not load records. Try again.", true);
}

/** Starts a continuation request or retries a failed initial/continuation request.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function requestNextPage(event: MouseEvent): void {
    const button = (event.target as Element | null)?.closest<HTMLButtonElement>(`[${ LOAD_MORE_ATTRIBUTE }]`);

    if(!button || button.disabled || button.hidden) {
        return;
    }

    const owner = document.getElementById(button.getAttribute(LOAD_MORE_ATTRIBUTE) ?? "");

    if(!(owner instanceof HTMLElement) || !owner.hasAttribute(PAGINATED_LIST_ATTRIBUTE)) {
        return;
    }

    const state = getState(owner);

    if(state.loading) {
        return;
    }

    const mode = state.retryMode ?? PaginationMode.Append;

    if(mode === PaginationMode.Append && (!state.initialized || state.cursor === null)) {
        return;
    }

    event.preventDefault();
    state.nextMode = mode;
    state.retryMode = undefined;
    state.controlRequestPending = true;
    state.pendingFocusButton = document.activeElement === button ? button : undefined;
    updateControls(owner, state);
    htmx.trigger(owner, LOAD_MORE_EVENT);
}

/** Invalidates a list when HTMX removes its owner from the document. */
function invalidateRemovedList(event: CustomEvent): void {
    const detail = event.detail as { elt?: Element };
    const removedElement = detail.elt ?? event.target;

    if(!(removedElement instanceof HTMLElement) || !removedElement.hasAttribute(PAGINATED_LIST_ATTRIBUTE)) {
        return;
    }

    const state = states.get(removedElement);

    if(state) {
        if(state.activeRequest) {
            focusTrackedRequests.delete(state.activeRequest);
        }

        if(state.preparingRequest) {
            focusTrackedRequests.delete(state.preparingRequest);
        }

        state.generation++;
        state.loading = false;
        state.controlRequestPending = false;
        state.pendingFocusButton = undefined;
        state.activeRequest = undefined;
        state.preparingRequest = undefined;
    }
}

/** Prepares complete portfolio-history snapshots for the viewing template.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function transformPortfolioHistoryView(responseBody: string, eventDetail: BeforeSwapEventDetail): string {
    return transformPageResponse<HistoryPageItem, { items: HistoryPageItem[]; initial: boolean }>(
        responseBody,
        eventDetail,
        item => item.observationTimestamp?.id,
        (items, mode) => ({ items, initial: mode === PaginationMode.Initial }),
    );
}

/** Prepares complete allocation-plan records for the viewing template.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function transformAllocationPlanView(responseBody: string, eventDetail: BeforeSwapEventDetail): string {
    return transformPageResponse<AllocationPlanPageItem, { items: AllocationPlanPageItem[]; initial: boolean }>(
        responseBody,
        eventDetail,
        item => item.id,
        (items, mode) => ({ items, initial: mode === PaginationMode.Initial }),
    );
}

/** Initializes the delegated HTMX lifecycle and control handlers once per page. */
function init(): void {
    if(initialized) {
        return;
    }

    initialized = true;

    HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
        "paginatePortfolioHistoryView",
        transformPortfolioHistoryView,
    );

    HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
        "paginateAllocationPlanView",
        transformAllocationPlanView,
    );
    document.addEventListener("click", requestNextPage);
    document.addEventListener("htmx:configRequest", configureRequest);
    document.addEventListener("htmx:beforeRequest", bindRequestToXhr);
    document.addEventListener("htmx:afterSettle", commitSettledPage);
    document.addEventListener("htmx:afterRequest", handleRequestCompletion);
    document.addEventListener("htmx:beforeCleanupElement", invalidateRemovedList);
    document.addEventListener("focusin", trackPaginationFocusChange);
}

/**
 * Converts one opt-in cursor response for a Handlebars template and stages its cursor until settle succeeds.
 *
 * @param responseBody - Raw API JSON returned by the paginated endpoint.
 * @param eventDetail - HTMX before-swap details for the response being transformed.
 * @param getId - Extracts a persisted record ID from each response item.
 * @param mapItems - Converts accepted records into the shape expected by the Handlebars template.
 * @returns JSON for the template extension, or an empty list when the response is stale/invalid.
 * @example `transformPageResponse(body, event, item => item.id, items => items)` prepares a plan page.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
function transformPageResponse<T, R>(
    responseBody: string,
    eventDetail: BeforeSwapEventDetail,
    getId: (item: T) => string | number | null | undefined,
    mapItems: (items: T[], mode: PaginationMode) => R,
): string {
    const context = requests.get(eventDetail.xhr);

    if(!context) {
        return responseBody;
    }

    if(!isCurrentRequest(context)) {
        context.stale = true;
        eventDetail.shouldSwap = false;
        eventDetail.swapOverride = "innerHTML";
        return JSON.stringify([]);
    }

    try {
        const envelope = JSON.parse(responseBody) as PageEnvelope<T>;

        if(!envelope || !Array.isArray(envelope.items)) {
            throw new Error("The paginated response envelope is malformed.");
        }

        const hasValidNextCursor = envelope.nextCursor === null
            || (typeof envelope.nextCursor === "string" && envelope.nextCursor.length > 0);

        if(!hasValidNextCursor) {
            throw new Error("The paginated response envelope is malformed.");
        }

        const seenIds = context.mode === PaginationMode.Append
            ? new Set(states.get(context.owner)?.recordIds ?? [])
            : new Set<string>();
        const acceptedItems: T[] = [];
        const acceptedIds: string[] = [];

        envelope.items.forEach(item => {
            const value = getId(item);

            if(value === null || value === undefined || String(value).length === 0) {
                throw new Error("A paginated record is missing its ID.");
            }

            const id = String(value);

            if(!seenIds.has(id)) {
                seenIds.add(id);
                acceptedItems.push(item);
                acceptedIds.push(id);
            }
        });

        context.acceptedIds = acceptedIds;
        context.nextCursor = envelope.nextCursor;
        context.prepared = true;
        eventDetail.swapOverride = context.mode === PaginationMode.Append ? "beforeend" : "innerHTML";

        return JSON.stringify(mapItems(acceptedItems, context.mode));
    }
    catch {
        context.rejected = true;
        eventDetail.shouldSwap = false;
        eventDetail.isError = true;
        eventDetail.swapOverride = "innerHTML";
        return JSON.stringify([]);
    }
}

/**
 * Provides independent HTMX cursor state and template response preparation for each paginated list.
 *
 * @example
 * ```ts
 * paginatedList.init();
 * const response = paginatedList.transformPageResponse(body, event, item => item.id, items => items);
 * ```
 * @author GPT-6 Luna
 */
const paginatedList = { init, transformPageResponse };

export default paginatedList;
