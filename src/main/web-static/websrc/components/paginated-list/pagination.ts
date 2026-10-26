/**
 * Coordinates cursor-based HTMX list requests without replacing committed records during append operations.
 *
 * @module components/paginated-list
 * @author GPT-6 Luna
 */

import htmx from "htmx.org";
import { HtmxInfra, type BeforeSwapEventDetail } from "../../infra/htmx";

const PAGINATED_LIST_ATTRIBUTE = "data-paginated-list";
const LOAD_MORE_ATTRIBUTE = "data-paginated-list-more";
const STATUS_ATTRIBUTE = "data-paginated-list-status";
const LOAD_MORE_EVENT = "paginated-list-more";
const PAGE_CURSOR_PARAMETER = "pageCursor";

type PaginationMode = "initial" | "append";

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
    xhr?: XMLHttpRequest;
};

const states = new WeakMap<HTMLElement, ListState>();
const requests = new WeakMap<XMLHttpRequest, RequestContext>();
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

/** Updates the control associated with the list after its request state changes. */
function updateControls(owner: HTMLElement, state: ListState, message?: string, isError = false): void {
    const button = Array.from(document.querySelectorAll<HTMLButtonElement>(`[${ LOAD_MORE_ATTRIBUTE }]`))
        .find(candidate => candidate.getAttribute(LOAD_MORE_ATTRIBUTE) === owner.id);

    const status = Array.from(document.querySelectorAll<HTMLElement>(`[${ STATUS_ATTRIBUTE }]`))
        .find(candidate => candidate.getAttribute(STATUS_ATTRIBUTE) === owner.id);

    if(button) {
        const retryAvailable = state.retryMode !== undefined;
        button.hidden = !(retryAvailable || (state.initialized && state.cursor !== null));
        button.disabled = state.loading;
        button.textContent = "Show more";

        if(state.loading) {
            button.textContent = "Loading…";
        }
        else if(retryAvailable) {
            button.textContent = "Retry";
        }
    }

    if(status) {
        status.textContent = message ?? "";
        status.setAttribute("role", isError ? "alert" : "status");
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
    requests.set(detail.xhr, context);
}

/** Adds the explicit initial or continuation cursor to a paginated HTMX request. */
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
    let requestedMode = state.nextMode ?? "initial";
    state.nextMode = undefined;

    const portfolioChanged = portfolioId !== state.portfolioId;

    if(portfolioChanged) {
        requestedMode = "initial";
        state.portfolioId = portfolioId;
        state.cursor = null;
        state.recordIds.clear();
        state.initialized = false;
    }

    if(requestedMode === "append" && (state.loading || !state.initialized || state.cursor === null)) {
        event.preventDefault();
        return;
    }

    if(portfolioChanged || requestedMode === "initial") {
        state.generation++;
    }

    const cursor = requestedMode === "append" ? state.cursor : "";

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
    };

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
    state.activeRequest = undefined;
    state.preparingRequest = undefined;
    updateControls(context.owner, state, message, isError);
}

/** Commits a validated page only after HTMX has completed its DOM swap and settle lifecycle. */
function commitSettledPage(event: CustomEvent): void {
    const detail = event.detail as { xhr?: XMLHttpRequest };
    const context = detail.xhr ? requests.get(detail.xhr) : undefined;

    if(!context || !context.prepared || context.rejected || context.stale) {
        return;
    }

    if(!isCurrentRequest(context)) {
        context.stale = true;
        return;
    }

    const state = states.get(context.owner);

    if(state?.activeRequest !== context) {
        return;
    }

    if(context.mode === "initial") {
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

    const loadMoreButton = Array.from(document.querySelectorAll<HTMLButtonElement>(`[${ LOAD_MORE_ATTRIBUTE }]`))
        .find(button => button.getAttribute(LOAD_MORE_ATTRIBUTE) === context.owner.id);
    const shouldMoveFocus = context.nextCursor === null && document.activeElement === loadMoreButton;

    let message = "";

    if(context.nextCursor === null) {
        message = state.recordIds.size === 0 ? "No records found." : "All records loaded.";
    }

    finishRequest(context, message);

    if(shouldMoveFocus) {
        const status = Array.from(document.querySelectorAll<HTMLElement>(`[${ STATUS_ATTRIBUTE }]`))
            .find(candidate => candidate.getAttribute(STATUS_ATTRIBUTE) === context.owner.id);
        status?.focus();
    }
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

/** Starts a continuation request or retries a failed initial/continuation request. */
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

    const mode = state.retryMode ?? "append";

    if(mode === "append" && (!state.initialized || state.cursor === null)) {
        return;
    }

    event.preventDefault();
    state.nextMode = mode;
    state.retryMode = undefined;
    button.disabled = true;
    button.textContent = "Loading…";
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
        state.generation++;
        state.loading = false;
        state.activeRequest = undefined;
        state.preparingRequest = undefined;
    }
}

/** Prepares complete portfolio-history snapshots for the viewing template. */
function transformPortfolioHistoryView(responseBody: string, eventDetail: BeforeSwapEventDetail): string {
    return transformPageResponse<HistoryPageItem, { items: HistoryPageItem[]; initial: boolean }>(
        responseBody,
        eventDetail,
        item => item.observationTimestamp?.id,
        (items, mode) => ({ items, initial: mode === "initial" }),
    );
}

/** Prepares complete allocation-plan records for the viewing template. */
function transformAllocationPlanView(responseBody: string, eventDetail: BeforeSwapEventDetail): string {
    return transformPageResponse<AllocationPlanPageItem, { items: AllocationPlanPageItem[]; initial: boolean }>(
        responseBody,
        eventDetail,
        item => item.id,
        (items, mode) => ({ items, initial: mode === "initial" }),
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

        const seenIds = context.mode === "append"
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
        eventDetail.swapOverride = context.mode === "append" ? "beforeend" : "innerHTML";

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
