/**
 * Binds pagination event handlers to individual list owners and their associated controls.
 * @module components/paginated-list
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */

const OWNER = "data-paginated-list";
const CONTROL = "data-paginated-list-more";
const AFTER_REQUEST = "htmx:afterRequest";

/**
 * Identifies control-driven continuation and retry requests, which must keep their selected pagination mode.
 * @example `htmx.trigger(owner, LOAD_MORE_EVENT)` requests the next page or retries the selected page.
 * @author GPT-6.1 Sol
 */
export const LOAD_MORE_EVENT = "paginated-list-more";

/**
 * Tracks completion on the original owner, including after it has been detached by HTMX.
 * A cancelled beforeRequest does not emit afterRequest, so its listener is released separately.
 * @param owner - The request owner.
 * @param xhr - The request to match.
 * @param beforeRequest - The cancellable HTMX event.
 * @param complete - Handles the matching completion before global notification listeners.
 * @param cancel - Releases state when the request is cancelled before sending.
 * @example `trackRequestCompletion(owner, xhr, event, complete, cancel)` tracks one request.
 * @author GPT-6.1 Terra
 */
export function trackRequestCompletion(
    owner: HTMLElement,
    xhr: XMLHttpRequest,
    beforeRequest: Event,
    complete: (event: CustomEvent) => void,
    cancel: () => void,
): void {
    const completion: EventListener = event => {
        if((event as CustomEvent<{ xhr?: XMLHttpRequest }>).detail?.xhr !== xhr) {
            return;
        }

        owner.removeEventListener(AFTER_REQUEST, completion);
        complete(event as CustomEvent);
    };
    owner.addEventListener(AFTER_REQUEST, completion);

    queueMicrotask(() => {
        if(beforeRequest.defaultPrevented) {
            owner.removeEventListener(AFTER_REQUEST, completion);
            cancel();
        }
    });
}

/**
 * Supplies the pagination handlers installed for each discovered owner.
 * @example `bindPaginationLifecycle({ reset, configure, beforeRequest, settle, cleanup, click })`.
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
type LifecycleHandlers = {
    reset: (owner: HTMLElement) => void;
    configure: EventListener;
    beforeRequest: EventListener;
    settle: EventListener;
    cleanup: EventListener;
    click: EventListener;
};

/**
 * Discovers existing and HTMX-processed list owners before their load triggers are initialized.
 * Cleanup removes owner and control listeners; request-specific completion listeners remain independent.
 *
 * @param handlers - Callbacks for one pagination owner's lifecycle.
 * @example `bindPaginationLifecycle(handlers)` registers reset intent and request lifecycle hooks during startup.
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
export function bindPaginationLifecycle(handlers: LifecycleHandlers): void {
    const bound = new WeakSet<HTMLElement>();

    /** Binds one owner and its associated control once, and releases them on HTMX cleanup.
     * @author GPT-6.1 Terra
     * @author GPT-6.1 Sol
     */
    function bind(owner: HTMLElement): void {
        if(bound.has(owner)) {
            return;
        }

        bound.add(owner);

        const control = Array.from(document.querySelectorAll<HTMLElement>(`[${ CONTROL }]`))
            .find(element => element.getAttribute(CONTROL) === owner.id);

        /** Invalidates resets before HTMX's request lock can defer them; control retries keep their mode.
         * @author GPT-6.1 Sol
         */
        const confirm: EventListener = event => {
            const detail = (event as CustomEvent<{ verb?: string; triggeringEvent?: Event }>).detail;

            if(event.target === owner && detail.verb?.toLowerCase() === "get"
                && detail.triggeringEvent?.type !== LOAD_MORE_EVENT) {
                handlers.reset(owner);
            }
        };

        owner.addEventListener("htmx:confirm", confirm);
        owner.addEventListener("htmx:configRequest", handlers.configure);
        owner.addEventListener("htmx:beforeRequest", handlers.beforeRequest);
        owner.addEventListener("htmx:afterSettle", handlers.settle);
        control?.addEventListener("click", handlers.click);

        const cleanup: EventListener = event => {
            if(event.target !== owner) {
                return;
            }

            handlers.cleanup(event);
            owner.removeEventListener("htmx:confirm", confirm);
            owner.removeEventListener("htmx:configRequest", handlers.configure);
            owner.removeEventListener("htmx:beforeRequest", handlers.beforeRequest);
            owner.removeEventListener("htmx:afterSettle", handlers.settle);
            owner.removeEventListener("htmx:beforeCleanupElement", cleanup);
            control?.removeEventListener("click", handlers.click);
            bound.delete(owner);
        };

        owner.addEventListener("htmx:beforeCleanupElement", cleanup);
    }

    document.addEventListener("htmx:beforeProcessNode", event => {
        const target = event.target;

        if(target instanceof HTMLElement && target.hasAttribute(OWNER)) {
            bind(target);
        }
    });
    document.querySelectorAll<HTMLElement>(`[${ OWNER }]`).forEach(bind);
}
