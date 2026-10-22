/**
 * Controls lazy observation-detail loading for portfolio-history management forms.
 *
 * Bind one header and its accordion panel after each management-list render. The loader requests details once when
 * the panel opens, exposes a retry after failure, and keeps allocation controls unavailable until the response settles.
 * Example: `bindObservationDetailsLoader(header, panel, form)` wires one persisted observation.
 * @author GPT-6 Luna
 */

import htmx from "htmx.org";

const LOADER_ATTRIBUTE = "data-observation-details-loader";
const BOUND_ATTRIBUTE = "data-observation-details-loader-bound";
const DETAILS_EVENT = "load-observation-details";
const PANEL_SHOWN_EVENT = "shown.bs.collapse";

type ObservationDetailsState = "unloaded" | "loading" | "ready" | "failed";

type RequestEventDetail = {
    requestConfig?: {
        elt?: Element;
        verb?: string;
    };
    successful?: boolean;
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

    panel.addEventListener(PANEL_SHOWN_EVENT, event => {
        if(event.target === panel) {
            htmx.trigger(header, DETAILS_EVENT);
        }
    });
}

/** Retries a failed detail request without changing the accordion's expanded state.
 *
 * @param header - The observation header configured as the detail-request source.
 * @example `retryObservationDetails(header)` retries a failed lazy load.
 * @author GPT-6 Luna
 */
export function retryObservationDetails(header: HTMLElement): void {
    if(readObservationDetailsState(header) === "failed") {
        htmx.trigger(header, DETAILS_EVENT);
    }
}

/** Invalidates loaded details and reloads them after a successful allocation-detail save.
 *
 * @param header - The observation header configured as the detail-request source.
 * @param form - The form whose allocation details are being refreshed.
 * @example `reloadObservationDetails(header, form)` refreshes rows without reloading the observation list.
 * @author GPT-6 Luna
 */
export function reloadObservationDetails(header: HTMLElement, form: HTMLFormElement): void {
    setObservationDetailsState(header, form, "unloaded");
    htmx.trigger(header, DETAILS_EVENT);
}

/** Reports whether allocation details are ready for an observation form.
 *
 * @param form - Observation form whose allocation state is being checked.
 * @returns `true` for a new-record form or a loaded existing observation.
 * @example `areObservationDetailsReady(form)` guards allocation edits and form submission.
 * @author GPT-6 Luna
 */
export function areObservationDetailsReady(form: HTMLFormElement): boolean {
    if(!form.hasAttribute(LOADER_ATTRIBUTE)) {
        return true;
    }

    return form.dataset.observationDetailsState === "ready";
}

/** Handles a request before it starts, suppressing duplicate and already-completed detail loads.
 * @author GPT-6 Luna
 */
function handleBeforeRequest(event: Event, header: HTMLElement, form: HTMLFormElement): void {
    const detail = (event as CustomEvent<RequestEventDetail>).detail;

    if(detail?.requestConfig?.elt !== header || detail.requestConfig.verb?.toLowerCase() !== "get") {
        return;
    }

    const state = readObservationDetailsState(header);

    if(state === "loading" || state === "ready") {
        event.preventDefault();
        return;
    }

    setObservationDetailsState(header, form, "loading");
}

/** Completes the matching detail request and updates the status and allocation controls.
 * @author GPT-6 Luna
 */
function handleAfterRequest(event: Event, header: HTMLElement, form: HTMLFormElement): void {
    const detail = (event as CustomEvent<RequestEventDetail>).detail;

    if(detail?.requestConfig?.elt !== header || detail.requestConfig.verb?.toLowerCase() !== "get") {
        return;
    }

    setObservationDetailsState(header, form, detail.successful ? "ready" : "failed");
}

/** Reads the persisted lifecycle value, defaulting an uninitialized loader to its unloaded state.
 * @author GPT-6 Luna
 */
function readObservationDetailsState(header: HTMLElement): ObservationDetailsState {
    const state = header.dataset.observationDetailsState;

    return state === "loading" || state === "ready" || state === "failed" ? state : "unloaded";
}

/** Updates loader status, editor feedback and detail-dependent form controls together.
 * @author GPT-6 Luna
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
        loadingMessage.hidden = state !== "loading";
    }

    if(errorMessage) {
        errorMessage.hidden = state !== "failed";
    }

    form.querySelectorAll<HTMLButtonElement>("[data-requires-observation-details]").forEach(button => {
        button.disabled = state !== "ready";
    });
}
