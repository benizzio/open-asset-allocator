/**
 * Coordinates observation clone preparation without replacing the management list.
 * A clone owns its source load and destination revision until it either applies or is cancelled.
 * @module components/portfolio-history/portfolio-history-management-observation-clone
 * @author GPT-6.1 Terra
 */
import htmx from "htmx.org";
import { Collapse } from "bootstrap";
import type { TemplateDelegate } from "handlebars";
import { subscribeRouteExit } from "../../../infra/routing";
import { calendarDateToUtcTimestamp, getBrowserLocalCalendarDate } from "../../../utils/calendar-date";
import { ensureObservationDetailsReady } from "../portfolio-history-management-observation-details";
import { captureObservationDraft, restoreObservationRows, type ObservationDraft } from "./draft";

const ROUTE = "/portfolio/:portfolioId/history/manage";
const DESTINATION = "#portfolio-history-management-container-0";
const COLLAPSE_SHOWN = "shown.bs.collapse";
const COLLAPSE_HIDDEN = "hidden.bs.collapse";
const boundLists = new WeakSet<HTMLElement>();

/** Captures form-control values that mutation observers cannot see.
 * @author GPT-6.1 Terra
 */
function snapshot(item: HTMLElement): string {
    return JSON.stringify(Array.from(item.querySelectorAll<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>(
        "input, select, textarea",
    ), input => [input.name, input.value, input instanceof HTMLInputElement && input.checked, input.disabled]));
}

/** Builds and checks destination-specific row markup before touching the existing draft.
 * @author GPT-6.1 Terra
 */
function renderRows(template: TemplateDelegate, draft: ObservationDraft): string {
    const html = draft.rows.map((_, allocationIndex) =>
        template({ allocationIndex, observationTimestampId: 0 })).join("");
    const documentFragment = document.createElement("template");
    documentFragment.innerHTML = `<table><tbody>${ html }</tbody></table>`;
    const rows = documentFragment.content.querySelectorAll("tbody > tr");

    if(rows.length !== draft.rows.length) {
        throw new Error("Could not render observation draft rows.");
    }

    for(const row of rows) {
        const requiredSelectors = [
            "[data-asset-search-autocomplete]", "[data-asset-search-input]", "[data-asset-ticker-input]",
            "[data-asset-action-button]", "[data-new-asset-ticker-message]", "[data-external-asset-keys]",
            "[aria-label='Asset name']", "[name$='[assetId]']", "[name$='[class]']", "[name$='[cashReserve]']",
            "[name$='[assetQuantity]']", "[data-financial-input$='[assetMarketPrice]']",
            "[data-financial-input$='[totalMarketValue]']",
        ];

        if(requiredSelectors.some(selector => !row.querySelector(selector))) {
            throw new Error("Rendered observation draft row is incomplete.");
        }
    }
    return html;
}

/** Waits for a destination collapse transition without retaining listeners after cancellation.
 * @author GPT-6.1 Terra
 */
function openPanel(panel: HTMLElement, signal: AbortSignal): Promise<void> {
    if(panel.classList.contains("show") && !panel.classList.contains("collapsing")) {
        return Promise.resolve();
    }
    return new Promise((resolve, reject) => {
        const accordion = panel.parentElement?.parentElement;

        const cleanup = () => {
            panel.removeEventListener(COLLAPSE_SHOWN, shown);
            panel.removeEventListener(COLLAPSE_HIDDEN, hidden);
            panel.removeEventListener("show.bs.collapse", showing);
            accordion?.removeEventListener(COLLAPSE_SHOWN, siblingSettled);
            accordion?.removeEventListener(COLLAPSE_HIDDEN, siblingSettled);
            signal.removeEventListener("abort", aborted);
        };
        const shown = () => { cleanup(); resolve(); };

        const showWhenReady = () => {
            const transitioning = accordion?.querySelectorAll<HTMLElement>(
                ":scope > .accordion-item > .accordion-collapse.collapsing",
            );

            if(Array.from(transitioning ?? []).some(element => element !== panel)
                || panel.classList.contains("collapsing")) {
                return;
            }
            Collapse.getOrCreateInstance(panel, { toggle: false }).show();
        };
        const hidden = () => showWhenReady();

        const siblingSettled = (event: Event) => {
            if(event.target !== panel) {
                showWhenReady();
            }
        };

        const showing = (event: Event) => {
            queueMicrotask(() => {
                if(event.defaultPrevented) {
                    cleanup();
                    reject(new Error("Observation destination could not be opened."));
                }
            });
        };
        const aborted = () => { cleanup(); reject(new DOMException("Clone cancelled.", "AbortError")); };
        panel.addEventListener(COLLAPSE_SHOWN, shown);
        panel.addEventListener(COLLAPSE_HIDDEN, hidden);
        panel.addEventListener("show.bs.collapse", showing);
        accordion?.addEventListener(COLLAPSE_SHOWN, siblingSettled);
        accordion?.addEventListener(COLLAPSE_HIDDEN, siblingSettled);
        signal.addEventListener("abort", aborted, { once: true });
        showWhenReady();
    });
}

/** Writes the new-record metadata after the fresh rows have settled.
 * @author GPT-6.1 Terra
 */
function applyMetadata(
    form: HTMLFormElement, tag: HTMLInputElement, dateInput: HTMLInputElement, draft: ObservationDraft,
): void {
    tag.value = draft.tag;
    tag.setCustomValidity(Array.from(draft.tag).length > 100 ? "Observation tag must not exceed 100 characters." : "");
    tag.classList.remove("is-invalid", "is-valid");
    const payload = form.elements.namedItem("observationTimestamp.timeTag");

    if(!(payload instanceof HTMLInputElement)) {
        throw new Error("Observation tag payload is missing.");
    }
    payload.value = draft.tag;
    dateInput.value = draft.date;
    dateInput.setCustomValidity("");
    dateInput.classList.remove("is-invalid", "is-valid");
    const timestamp = form.elements.namedItem("observationTimestamp.timestamp");

    if(!(timestamp instanceof HTMLInputElement)) {
        throw new Error("Observation timestamp payload is missing.");
    }
    timestamp.value = calendarDateToUtcTimestamp(draft.date);
    form.classList.remove("was-validated");
}

/** Opens the destination and focuses its tag only while the operation is still current.
 * @author GPT-6.1 Terra
 */
async function focusDestination(
    destination: HTMLElement, tag: HTMLInputElement, signal: AbortSignal, owns: () => boolean, mayFocus: () => boolean,
): Promise<void> {
    const panel = destination.querySelector<HTMLElement>(".accordion-collapse");

    if(!panel) {
        throw new Error("Observation destination panel is missing.");
    }
    await openPanel(panel, signal);

    if(owns() && mayFocus()) {
        destination.scrollIntoView({ block: "center" });
        tag.focus();
    }
}

/** Binds one delegated clone controller to the management list.
 * @param list - The stable paginated list owner.
 * @param rowTemplate - Returns the initialized allocation-row template.
 * @param normalize - Rebinds quote state after restoration.
 * @example `bindObservationCloning(list, () => rowTemplate, normalizeRow);`
 * @author GPT-6.1 Terra
 */
export function bindObservationCloning(
    list: HTMLElement, rowTemplate: () => TemplateDelegate, normalize: (row: HTMLTableRowElement) => void,
): void {
    if(boundLists.has(list)) {
        return;
    }
    boundLists.add(list);
    let generation = 0;
    let active: AbortController | undefined;
    let pendingSource: HTMLElement | undefined;
    let postInFlight = false;
    const buttonOwners = new WeakMap<HTMLButtonElement, AbortController>();
    const cancel = () => { generation++; active?.abort(); active = undefined; pendingSource = undefined; };

    list.addEventListener("htmx:beforeRequest", event => {
        const detail = (event as CustomEvent).detail;
        const source = detail?.requestConfig?.elt;

        if(source instanceof HTMLFormElement && source.id === "portfolio-history-management-form-0"
            && detail.requestConfig.verb?.toLowerCase() === "post") {
            postInFlight = true;
            cancel();

            queueMicrotask(() => {
                if(event.defaultPrevented) {
                    postInFlight = false;
                }
            });
        }
        else if(source === list && detail.requestConfig.verb?.toLowerCase() === "get") {
            cancel();
        }
    });

    list.addEventListener("htmx:afterRequest", event => {
        const detail = (event as CustomEvent).detail;
        const source = detail?.requestConfig?.elt;

        if(source instanceof HTMLFormElement && source.id === "portfolio-history-management-form-0"
            && detail.requestConfig.verb?.toLowerCase() === "post") {
            postInFlight = false;
        }
    });

    list.addEventListener("click", event => {
        const button = (event.target as Element)?.closest<HTMLButtonElement>("[data-observation-clone]");

        if(!button || !list.contains(button)) {
            return;
        }
        event.preventDefault();
        event.stopPropagation();
        const item = button.closest<HTMLElement>(".accordion-item");
        const destination = list.querySelector<HTMLElement>(DESTINATION);
        const status = item?.querySelector<HTMLElement>("[data-observation-clone-status]");

        if(!item || !destination || !status || item === destination) {
            return;
        }

        if(postInFlight) {
            status.textContent = "Wait for the current save to finish, then try again.";
            return;
        }

        if(active && pendingSource === item) {
            return;
        }
        cancel();
        const controller = new AbortController();
        active = controller;
        buttonOwners.set(button, controller);
        pendingSource = item;

        const cancelIfOwner = () => {
            if(active === controller) {
                cancel();
            }
        };
        const current = generation;
        const date = getBrowserLocalCalendarDate();
        const portfolio = document.querySelector<HTMLInputElement>("[name='portfolioId']")?.value;
        const path = location.pathname;
        const initial = snapshot(destination);
        const observer = new MutationObserver(cancelIfOwner);

        observer.observe(destination, {
            subtree: true,
            childList: true,
            attributes: true,
            characterData: true,
            attributeFilter: ["value", "checked", "disabled", "readonly", "data-asset-selection-state",
                "data-asset-selection-generation"],
        });

        const ownerObserver = new MutationObserver(() => {
            if(!list.isConnected || !destination.isConnected || !item.isConnected) {
                cancelIfOwner();
            }
        });
        ownerObserver.observe(document, { childList: true, subtree: true });
        const onEdit = () => cancelIfOwner();
        let focusMoved = false;
        const onFocus = () => { focusMoved = true; };
        document.addEventListener("focusin", onFocus, true);

        const onPortfolioChange = (changed: Event) => {
            if(changed.target instanceof HTMLInputElement && changed.target.name === "portfolioId") {
                cancelIfOwner();
            }
        };
        document.addEventListener("change", onPortfolioChange, true);

        for(const name of ["input", "change", "submit", "invalid"]) {
            destination.addEventListener(name, onEdit, true);
        }
        const unsubscribe = subscribeRouteExit(ROUTE, cancelIfOwner);
        status.textContent = "Loading observation for cloning…";
        button.setAttribute("aria-busy", "true");
        const header = item.querySelector<HTMLElement>("[data-observation-details-loader]");
        const source = item.querySelector<HTMLFormElement>("form");
        const form = destination.querySelector<HTMLFormElement>("form");
        const tbody = form?.querySelector<HTMLElement>("tbody");
        const tag = destination.querySelector<HTMLInputElement>("[data-observation-editor-tag]");
        const dateInput = destination.querySelector<HTMLInputElement>("#portfolio-history-management-observation-date");

        const owns = () => {
            const sameOperation = !controller.signal.aborted && generation === current && list.isConnected;

            const sameElements = list.contains(item) && item.contains(source)
                && list.querySelector(DESTINATION) === destination && destination.contains(form);
            const currentPortfolio = document.querySelector<HTMLInputElement>("[name='portfolioId']")?.value;
            return sameOperation && sameElements && location.pathname === path && currentPortfolio === portfolio;
        };

        void (async () => {
            try {
                if(!header || !source) {
                    throw new Error("Observation clone source controls are unavailable.");
                }

                if(!form || !tbody || !tag || !dateInput) {
                    throw new Error("Observation clone controls are unavailable.");
                }
                const tagPayload = form.elements.namedItem("observationTimestamp.timeTag");
                const timestamp = form.elements.namedItem("observationTimestamp.timestamp");

                if(!(tagPayload instanceof HTMLInputElement) || !(timestamp instanceof HTMLInputElement)) {
                    throw new Error("Observation clone metadata payload is unavailable.");
                }
                await ensureObservationDetailsReady(header, source, controller.signal);
                const draft = captureObservationDraft(source, date);
                const html = renderRows(rowTemplate(), draft);

                if(!owns() || observer.takeRecords().length || snapshot(destination) !== initial) {
                    if(active === controller) {
                        cancel();
                    }
                    return;
                }
                observer.disconnect();

                htmx.swap(tbody, html, { swapStyle: "innerHTML", swapDelay: 0, settleDelay: 0 }, {
                    afterSettleCallback: () => {
                        restoreObservationRows(tbody, draft);
                        tbody.querySelectorAll<HTMLTableRowElement>("tr").forEach(normalize);
                    },
                });
                applyMetadata(form, tag, dateInput, draft);

                await focusDestination(destination, tag, controller.signal, owns,
                    () => !focusMoved || document.activeElement === tag);

                if(owns()) {
                    status.textContent = "Observation copied to the new draft.";
                }
            } catch(error) {
                if(owns()) {
                    status.textContent = "Unable to clone observation. Please try again.";
                    console.error(error);
                }
            } finally {
                observer.disconnect();
                ownerObserver.disconnect();
                document.removeEventListener("change", onPortfolioChange, true);
                document.removeEventListener("focusin", onFocus, true);

                for(const name of ["input", "change", "submit", "invalid"]) {
                    destination.removeEventListener(name, onEdit, true);
                }
                unsubscribe();

                if(buttonOwners.get(button) === controller) {
                    button.removeAttribute("aria-busy");
                    buttonOwners.delete(button);

                    if(controller.signal.aborted) {
                        status.textContent = "Clone cancelled. Try again when ready.";
                    }
                }

                if(active === controller) {
                    active = undefined;
                    pendingSource = undefined;
                }
            }
        })();
    });
}
