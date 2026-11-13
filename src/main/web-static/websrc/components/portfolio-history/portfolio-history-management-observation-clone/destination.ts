/**
 * Validates fresh allocation markup and presents the cloned draft through Bootstrap's accordion lifecycle.
 * @module components/portfolio-history/portfolio-history-management-observation-clone/destination
 * @author GPT-6.1 Sol
 */
import { Collapse } from "bootstrap";
import type { TemplateDelegate } from "handlebars";
import type { ObservationDraft } from "./draft";
import type { CloneOperation } from "./operation";

const COLLAPSE_SHOWN = "shown.bs.collapse";
const COLLAPSE_HIDDEN = "hidden.bs.collapse";

/**
 * Builds and validates destination-specific row markup before the existing draft is touched.
 * @param template - Initialized allocation row template.
 * @param draft - Detached current source snapshot.
 * @returns Complete ID-0 row markup.
 * @example `const html = renderRows(template, draft);`
 * @author GPT-6.1 Sol
 */
export function renderRows(template: TemplateDelegate, draft: ObservationDraft): string {
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

/** Waits for the destination transition without retaining listeners after cancellation.
 * @author GPT-6.1 Sol
 */
function openPanel(panel: HTMLElement, signal: AbortSignal): Promise<void> {
    if(signal.aborted) {
        return Promise.reject(new DOMException("Clone cancelled.", "AbortError"));
    }

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

/**
 * Opens the destination and focuses its tag only if the operation still owns the UI and focus.
 * @param destination - ID-0 item containing the newly applied draft.
 * @param tag - New observation's tag input.
 * @param operation - Ownership and lifetime checks for this opening.
 * @example `await focusDestination(destination, tag, operation);`
 * @author GPT-6.1 Sol
 */
export async function focusDestination(
    destination: HTMLElement, tag: HTMLInputElement, operation: CloneOperation,
): Promise<void> {
    const panel = destination.querySelector<HTMLElement>(".accordion-collapse");

    if(!panel) {
        throw new Error("Observation destination panel is missing.");
    }
    await openPanel(panel, operation.controller.signal);

    if(operation.owns() && (operation.mayFocus() || document.activeElement === tag)) {
        destination.scrollIntoView({ block: "center" });
        tag.focus();
    }
}
