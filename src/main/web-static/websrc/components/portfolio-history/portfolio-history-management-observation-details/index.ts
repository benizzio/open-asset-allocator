/**
 * Controls lazy observation-detail loading for portfolio-history management forms.
 *
 * Bind one header and its accordion panel after each management-list render. The loader requests details once when
 * the panel opens, exposes a retry after failure, and keeps allocation controls unavailable until the response settles.
 * Example: `bindObservationDetailsLoader(header, panel, form)` wires one persisted observation.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */

import htmx from "htmx.org";

const LOADER_ATTRIBUTE = "data-observation-details-loader";
const BOUND_ATTRIBUTE = "data-observation-details-loader-bound";
const DETAILS_EVENT = "load-observation-details";
const PANEL_SHOWN_EVENT = "shown.bs.collapse";
const STATE_CHANGED_EVENT = "observation-details-state-changed";
const requests = new WeakMap<HTMLElement, XMLHttpRequest>();

/** Lifecycle values persisted in observation detail loader datasets.
 * @author GPT-6.1 Terra
 */
enum ObservationDetailsState {
    UNLOADED = "unloaded",
    LOADING = "loading",
    READY = "ready",
    FAILED = "failed",
}

type RequestEventDetail = {
    requestConfig?: {
        elt?: Element;
        verb?: string;
    };
    successful?: boolean;
    xhr?: XMLHttpRequest;
    target?: Element;
    shouldSwap?: boolean;
};

/** Binds one observation header and panel to an idempotent lazy-details lifecycle.
 *
 * @param header - HTMX source configured for the observation detail GET request.
 * @param panel - Accordion panel whose first opening should initiate the request.
 * @param form - Form whose allocation controls depend on loaded observation details.
 * @example `bindObservationDetailsLoader(header, panel, form)` enables lazy, draft-safe detail loading.
 * @author GPT-6 Luna
 */
export function bindObservationDetailsLoader(
    header: HTMLElement,
    panel: HTMLElement,
    form: HTMLFormElement,
): void {
    if(header.hasAttribute(BOUND_ATTRIBUTE)) {
        return;
    }

    header.setAttribute(BOUND_ATTRIBUTE, "true");
    header.setAttribute(LOADER_ATTRIBUTE, "true");
    form.setAttribute(LOADER_ATTRIBUTE, "true");
    setObservationDetailsState(header, form, readObservationDetailsState(header));
    header.addEventListener("htmx:beforeRequest", event => handleBeforeRequest(event, header, form));
    header.addEventListener("htmx:afterRequest", event => handleAfterRequest(event, header, form));
    form.addEventListener("htmx:beforeSwap", event => handleBeforeSwap(event, header, form));
    form.addEventListener("htmx:afterSettle", event => handleAfterSettle(event, header, form));
    const fail = () => setObservationDetailsState(header, form, ObservationDetailsState.FAILED);
    header.addEventListener("htmx:swapError", fail);
    header.addEventListener("htmx:onLoadError", fail);

    panel.addEventListener(PANEL_SHOWN_EVENT, event => {
        if(event.target === panel) {
            htmx.trigger(header, DETAILS_EVENT);
        }
    });
}

/** Fails a matching detail load when another handler suppresses its swap.
 * @author GPT-6.1 Terra
 */
function handleBeforeSwap(event: Event, header: HTMLElement, form: HTMLFormElement): void {
    const detail = (event as CustomEvent<RequestEventDetail>).detail;
    const xhr = detail?.xhr;

    if(!xhr || requests.get(header) !== xhr || detail.target !== form.querySelector("tbody")) {
        return;
    }

    queueMicrotask(() => {
        if(requests.get(header) === xhr && (event.defaultPrevented || detail.shouldSwap === false)) {
            setObservationDetailsState(header, form, ObservationDetailsState.FAILED);
        }
    });
}

/** Ensures that an observation's matching detail request has settled and its row bindings are usable.
 * Existing requests are shared; cancelling this waiter does not abort the source request.
 * @param header - The bound HTMX detail request source.
 * @param form - The corresponding observation form.
 * @param signal - Optional cancellation signal for this waiter.
 * @example `await ensureObservationDetailsReady(header, form, abortController.signal);`
 * @author GPT-6.1 Terra
 */
export function ensureObservationDetailsReady(
    header: HTMLElement, form: HTMLFormElement, signal?: AbortSignal,
): Promise<void> {
    if(!header.isConnected || !form.isConnected || signal?.aborted) {
        return Promise.reject(new DOMException("Observation details wait cancelled.", "AbortError"));
    }

    if(areObservationDetailsReady(form)) {
        return Promise.resolve();
    }

    return new Promise((resolve, reject) => {
        const observer = new MutationObserver(() => {
            if(!header.isConnected || !form.isConnected) {
                cancel();
            }
        });

        const cleanup = () => {
            observer.disconnect();
            header.removeEventListener(STATE_CHANGED_EVENT, stateChanged);
            signal?.removeEventListener("abort", cancel);
        };

        const cancel = () => {
            cleanup();
            reject(new DOMException("Observation details wait cancelled.", "AbortError"));
        };

        const stateChanged = () => {
            const state = readObservationDetailsState(header);

            if(state === ObservationDetailsState.READY) {
                cleanup();
                resolve();
            }
            else if(state === ObservationDetailsState.FAILED) {
                cleanup();
                reject(new Error("Observation details could not be loaded."));
            }
        };
        header.addEventListener(STATE_CHANGED_EVENT, stateChanged);
        signal?.addEventListener("abort", cancel, { once: true });
        observer.observe(document, { childList: true, subtree: true });

        if(readObservationDetailsState(header) !== ObservationDetailsState.LOADING) {
            htmx.trigger(header, DETAILS_EVENT);

            if(readObservationDetailsState(header) !== ObservationDetailsState.LOADING) {
                cleanup();
                reject(new Error("Observation detail request did not start."));
            }
        }
    });
}

/** Retries a failed detail request without changing the accordion's expanded state.
 *
 * @param header - The observation header configured as the detail-request source.
 * @example `retryObservationDetails(header)` retries a failed lazy load.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
export function retryObservationDetails(header: HTMLElement): void {
    if(readObservationDetailsState(header) === ObservationDetailsState.FAILED) {
        htmx.trigger(header, DETAILS_EVENT);
    }
}

/** Invalidates loaded details and reloads them after a successful allocation-detail save.
 *
 * @param header - The observation header configured as the detail-request source.
 * @param form - The form whose allocation details are being refreshed.
 * @example `reloadObservationDetails(header, form)` refreshes rows without reloading the observation list.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
export function reloadObservationDetails(header: HTMLElement, form: HTMLFormElement): void {
    setObservationDetailsState(header, form, ObservationDetailsState.UNLOADED);
    htmx.trigger(header, DETAILS_EVENT);
}

/** Reports whether allocation details are ready for an observation form.
 *
 * @param form - Observation form whose allocation state is being checked.
 * @returns `true` for a new-record form or a loaded existing observation.
 * @example `areObservationDetailsReady(form)` guards allocation edits and form submission.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
export function areObservationDetailsReady(form: HTMLFormElement): boolean {
    if(!form.hasAttribute(LOADER_ATTRIBUTE)) {
        return true;
    }

    return form.dataset.observationDetailsState === ObservationDetailsState.READY;
}

/** Handles a request before it starts, suppressing duplicate and already-completed detail loads.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function handleBeforeRequest(event: Event, header: HTMLElement, form: HTMLFormElement): void {
    const detail = (event as CustomEvent<RequestEventDetail>).detail;

    if(detail?.requestConfig?.elt !== header || detail.requestConfig.verb?.toLowerCase() !== "get") {
        return;
    }

    const state = readObservationDetailsState(header);

    if(state === ObservationDetailsState.LOADING || state === ObservationDetailsState.READY) {
        event.preventDefault();
        return;
    }

    setObservationDetailsState(header, form, ObservationDetailsState.LOADING);

    if(detail.xhr) {
        requests.set(header, detail.xhr);
        const xhr = detail.xhr;

        queueMicrotask(() => {
            if(event.defaultPrevented && requests.get(header) === xhr) {
                setObservationDetailsState(header, form, ObservationDetailsState.FAILED);
            }
        });
    }
}

/** Completes the matching detail request and updates the status and allocation controls.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function handleAfterRequest(event: Event, header: HTMLElement, form: HTMLFormElement): void {
    const detail = (event as CustomEvent<RequestEventDetail>).detail;

    if(detail?.requestConfig?.elt !== header || detail.requestConfig.verb?.toLowerCase() !== "get"
        || !detail.xhr || requests.get(header) !== detail.xhr) {
        return;
    }

    if(!detail.successful || detail.target !== form.querySelector("tbody") || !validDetailsResponse(detail.xhr, form)) {
        setObservationDetailsState(header, form, ObservationDetailsState.FAILED);
    }
}

/** Checks the response identity and allocation shape before accepting a settled swap.
 * @author GPT-6.1 Terra
 */
function validDetailsResponse(xhr: XMLHttpRequest, form: HTMLFormElement): boolean {
    try {
        if(xhr.status < 200 || xhr.status >= 300) {
            return false;
        }
        const body: unknown = JSON.parse(xhr.responseText);

        if(!Array.isArray(body) || body.length !== 1) {
            return false;
        }
        const snapshot = body[0] as { observationTimestamp?: { id?: unknown }; allocations?: unknown };
        return String(snapshot.observationTimestamp?.id) === form.id.slice("portfolio-history-management-form-".length)
            && Array.isArray(snapshot.allocations);
    } catch {
        return false;
    }
}

/** Verifies that the settled row count and required controls match the response.
 * @author GPT-6.1 Terra
 */
function usableDetailsRows(xhr: XMLHttpRequest, target: HTMLElement): boolean {
    const body = JSON.parse(xhr.responseText) as Array<{ allocations: unknown[] }>;
    const rows = Array.from(target.querySelectorAll<HTMLTableRowElement>(":scope > tr"));

    if(rows.length !== body[0].allocations.length) {
        return false;
    }
    return rows.every(row => {
        const financial = row.querySelectorAll<HTMLInputElement>("[data-financial-input]");

        const financialReady = financial.length === 2
            && Array.from(financial).every(input => input.dataset.financialInputBound === "true");

        const basicFieldsReady = !!row.querySelector("input[name$='[assetId]']")
            && !!row.querySelector("input[name$='[class]']")
            && !!row.querySelector("input[name$='[assetQuantity]']");
        return financialReady && basicFieldsReady;
    });
}

/** Marks only the matching settled and initialized detail target ready.
 * @author GPT-6.1 Terra
 */
function handleAfterSettle(event: Event, header: HTMLElement, form: HTMLFormElement): void {
    const detail = (event as CustomEvent<RequestEventDetail>).detail;
    const target = event.target;
    const xhr = detail?.xhr;

    const matchesRequest = xhr && requests.get(header) === xhr && detail.requestConfig?.elt === header;

    if(!(target instanceof HTMLElement) || target !== form.querySelector("tbody") || !matchesRequest) {
        return;
    }

    queueMicrotask(() => {
        if(requests.get(header) !== xhr || !header.isConnected || !form.contains(target)) {
            return;
        }

        const valid = validDetailsResponse(xhr, form) && usableDetailsRows(xhr, target);
        const state = valid ? ObservationDetailsState.READY : ObservationDetailsState.FAILED;
        setObservationDetailsState(header, form, state);
    });
}

/** Reads the persisted lifecycle value, defaulting an uninitialized loader to its unloaded state.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function readObservationDetailsState(header: HTMLElement): ObservationDetailsState {
    const state = header.dataset.observationDetailsState;

    switch(state) {
        case ObservationDetailsState.LOADING:
            return ObservationDetailsState.LOADING;
        case ObservationDetailsState.READY:
            return ObservationDetailsState.READY;
        case ObservationDetailsState.FAILED:
            return ObservationDetailsState.FAILED;
        default:
            return ObservationDetailsState.UNLOADED;
    }
}

/** Updates loader status, editor feedback and detail-dependent form controls together.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function setObservationDetailsState(
    header: HTMLElement,
    form: HTMLFormElement,
    state: ObservationDetailsState,
): void {
    header.dataset.observationDetailsState = state;
    form.dataset.observationDetailsState = state;

    const item = header.closest<HTMLElement>(".accordion-item");
    const loadingMessage = item?.querySelector<HTMLElement>("[data-observation-details-loading]");
    const errorMessage = item?.querySelector<HTMLElement>("[data-observation-details-error]");

    if(loadingMessage) {
        loadingMessage.hidden = state !== ObservationDetailsState.LOADING;
    }

    if(errorMessage) {
        errorMessage.hidden = state !== ObservationDetailsState.FAILED;
    }

    form.querySelectorAll<HTMLButtonElement>("[data-requires-observation-details]").forEach(button => {
        button.disabled = state !== ObservationDetailsState.READY;
    });
    header.dispatchEvent(new Event(STATE_CHANGED_EVENT));
}
