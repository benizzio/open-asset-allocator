/**
 * Presents an already replaced management draft after its accordion transition,
 * without stealing focus when ownership or user intent changes.
 * @module components/management-draft-presentation/focus
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
import { Collapse } from "bootstrap";
import type { CloneOperation } from "./operation";

const COLLAPSE_SHOWN = "shown.bs.collapse";
const COLLAPSE_HIDDEN = "hidden.bs.collapse";

/** Waits for a Bootstrap collapse to open, including transitions of siblings.
 * @author GPT-6.1 Terra
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
                    reject(new Error("Draft destination could not be opened."));
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

/** Opens the destination and focuses its field only while the operation owns focus.
 * @param destination - Item containing the replaced draft.
 * @param field - Destination name or tag input.
 * @param operation - Ownership and lifetime checks.
 * @example `await focusDraftDestination(item, nameInput, operation);`
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
export async function focusDraftDestination(
    destination: HTMLElement, field: HTMLInputElement, operation: CloneOperation,
): Promise<void> {
    const panel = destination.querySelector<HTMLElement>(".accordion-collapse");

    if(!panel) {
        throw new Error("Draft destination panel is missing.");
    }
    await openPanel(panel, operation.controller.signal);

    if(operation.owns() && (operation.mayFocus() || document.activeElement === field)) {
        destination.scrollIntoView({ block: "center" });
        field.focus();
    }
}
