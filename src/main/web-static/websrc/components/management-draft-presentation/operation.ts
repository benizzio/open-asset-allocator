/**
 * Tracks the lifetime of destination opening after a management draft has been applied.
 * Route exits, replaced owners and destination edits cancel pending focus, not the synchronous snapshot.
 * @module components/management-draft-presentation/operation
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */
import { subscribeRouteExit } from "../../infra/routing";

const EDIT_EVENTS = ["input", "change", "submit", "invalid"];

/** Lifetime checks and cleanup for one destination-opening operation.
 * @author GPT-6.1 Sol
 */
export type CloneOperation = {
    controller: AbortController;
    owns: () => boolean;
    mayFocus: () => boolean;
    dispose: () => void;
};

/** Reads the currently selected portfolio without retaining a replaced global input.
 * @author GPT-6.1 Sol
 */
function currentPortfolio(): string | undefined {
    return document.querySelector<HTMLInputElement>("[name='portfolioId']")?.value;
}

/**
 * Guards one asynchronous destination opening and releases all listeners when finished.
 * @param list - Stable management-list owner.
 * @param item - Source item whose current editable snapshot has been copied.
 * @param destination - ID-0 destination item.
 * @param formId - ID of the destination creation form.
 * @param route - Management route whose exit cancels presentation.
 * @returns Cancellation, ownership and focus checks plus mandatory cleanup.
 * @example `const operation = trackCloneOpening(list, sourceItem, destination, formId, route);`
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */
export function trackCloneOpening(
    list: HTMLElement, item: HTMLElement, destination: HTMLElement, formId: string, route: string,
): CloneOperation {
    const controller = new AbortController();
    const cancel = () => controller.abort();
    const path = location.pathname;
    const portfolio = currentPortfolio();
    const form = destination.querySelector<HTMLFormElement>(`#${ formId }`);
    const source = item.querySelector<HTMLFormElement>("form");
    let focusMoved = false;
    const onFocus = () => { focusMoved = true; };

    const owns = () => {
        const sameOperation = !controller.signal.aborted && list.isConnected;
        const sameSource = list.contains(item) && item.contains(source);
        const sameDestination = list.contains(destination) && destination.contains(form);
        const sameContext = location.pathname === path && currentPortfolio() === portfolio;
        return sameOperation && sameSource && sameDestination && sameContext;
    };

    const observer = new MutationObserver(() => {
        if(!owns()) {
            cancel();
        }
    });
    observer.observe(document, { childList: true, subtree: true });
    const draftObserver = new MutationObserver(cancel);

    draftObserver.observe(destination, {
        subtree: true,
        childList: true,
        attributes: true,
        attributeFilter: ["value", "checked", "disabled", "readonly", "data-asset-selection-state",
            "data-asset-selection-generation"],
    });
    document.addEventListener("focusin", onFocus, true);

    const onPortfolioChange = (event: Event) => {
        if(event.target instanceof HTMLInputElement && event.target.name === "portfolioId") {
            cancel();
        }
    };
    document.addEventListener("change", onPortfolioChange, true);

    for(const name of EDIT_EVENTS) {
        destination.addEventListener(name, cancel, true);
    }
    const unsubscribe = subscribeRouteExit(route, cancel);

    return {
        controller,
        owns,
        mayFocus: () => !focusMoved,
        dispose: () => {
            observer.disconnect();
            draftObserver.disconnect();
            document.removeEventListener("focusin", onFocus, true);
            document.removeEventListener("change", onPortfolioChange, true);

            for(const name of EDIT_EVENTS) {
                destination.removeEventListener(name, cancel, true);
            }
            unsubscribe();
        },
    };
}

/**
 * Cancels pending destination focus on list replacement or creation and tracks creation requests.
 * @param list - Stable management-list owner.
 * @param cancel - Cancels the current clone operation.
 * @param formId - ID of the destination creation form.
 * @returns Whether creation is currently in flight.
 * @example `const isSaving = bindCloneRequestGuards(list, cancel, formId);`
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */
export function bindCloneRequestGuards(list: HTMLElement, cancel: () => void, formId: string): () => boolean {
    let saving = false;

    const isCreation = (source: unknown, verb?: string) => source instanceof HTMLFormElement
        && source.getAttribute("id") === formId && verb?.toLowerCase() === "post";

    list.addEventListener("htmx:beforeRequest", event => {
        const detail = (event as CustomEvent).detail;
        const source = detail?.requestConfig?.elt;

        if(isCreation(source, detail?.requestConfig?.verb)) {
            saving = true;
            cancel();

            queueMicrotask(() => {
                if(event.defaultPrevented) {
                    saving = false;
                }
            });
        }
        else if(source === list && detail?.requestConfig?.verb?.toLowerCase() === "get") {
            cancel();
        }
    });

    list.addEventListener("htmx:afterRequest", event => {
        const detail = (event as CustomEvent).detail;

        if(isCreation(detail?.requestConfig?.elt, detail?.requestConfig?.verb)) {
            saving = false;
        }
    });
    return () => saving;
}
