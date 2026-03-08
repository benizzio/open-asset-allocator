import htmx from "htmx.org";
import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";
import { HookCleanupFunction, navigoRouter } from "./routing-navigo";
import { RequestConfigEventDetail } from "../htmx";
import { Match } from "navigo";

// =============================================================================
// HTMX TRIGGER EVENT ON ROUTE
// triggers a custom HTMX event when a route is navigated to
// attribute value format is "route/:path-param!event"
// =============================================================================

const HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE = "data-hx-trigger-on-route";
const HTMX_CLEAN_ON_EXIT_ROUTE_ATTRIBUTE = "data-hx-trigger-on-route-clean-on-exit";
const HTMX_TRIGGER_ON_ROUTE_EVENT_SEPARATOR = "!";
const HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG = "data-hx-trigger-on-route-bound";
const HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_SETTLED_ATTRIBUTE = "data-hx-trigger-on-route-wait-for-settled";
const HTMX_TRIGGER_ON_ROUTE_SETTLED_FLAG = "data-hx-trigger-on-route-settled";
const HTMX_TRIGGER_ON_ROUTE_SETTLED_RESET_BOUND_FLAG = "data-hx-trigger-on-route-settled-reset-bound";

const HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_DEPENDENCIES_READY_ATTRIBUTE =
    "data-hx-trigger-on-route-wait-for-dependencies-ready";

const HTMX_TRIGGER_ON_ROUTE_DEPENDENCIES_READY_FLAG =
    "data-hx-trigger-on-route-dependencies-ready";

const ROUTE_HANDLER_MAP = new Map<HTMLElement, (match: Match) => void>();

const CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS = {
    "htmx:afterSettle": (event: Event) => {
        if(event.target === event.currentTarget) {
            const eventElement = event.target as HTMLElement;
            eventElement.setAttribute(HTMX_CLEAN_ON_EXIT_ROUTE_ATTRIBUTE, "false");
        }
    },
    "htmx:confirm": (event: Event) => {

        if(event.target !== event.currentTarget) {
            return;
        }

        const eventElement = event.target as HTMLElement;
        const cleanAttributeValue = eventElement.getAttribute(HTMX_CLEAN_ON_EXIT_ROUTE_ATTRIBUTE);
        const isClean = cleanAttributeValue !== "false";

        if(!isClean) {
            event.preventDefault();
        }
    },
};

export function bindHTMXTriggerOnRouteInDescendants(element: HTMLElement) {
    const htmxRoutedElements = element.querySelectorAll(
        `[${ HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE }]:not([${ HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG }])`,
    ) as NodeListOf<HTMLElement>;
    bindRouteToHTMXEventOnElements(htmxRoutedElements);
}

function bindRouteToHTMXEventOnElements(htmxRoutedElements: NodeListOf<HTMLElement>) {

    htmxRoutedElements.forEach((element) => {

        element.setAttribute(HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG, "binding");

        try {

            logger(LogLevel.INFO, "Binding HTMX event on route for element", element);

            const { route, event } = extractBindingData(element);

            bindRouteToHTMXTriggerOnElement(element, route, event);
            bindCleanOnExitRouteBehaviourOnElement(element, route);
            addDisableRouteRemovalObserver(element, route);

            element.setAttribute(HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG, "true");

            executeImmediatelyIfOnRoute(route, element, event);
        } catch(error) {
            element.removeAttribute(HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG);
            throw error;
        }
    });
}

function extractBindingData(element: HTMLElement) {
    const routeToEvent = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE);
    const [route, event] = routeToEvent.split(HTMX_TRIGGER_ON_ROUTE_EVENT_SEPARATOR);
    return { route, event };
}

function bindRouteToHTMXTriggerOnElement(element: HTMLElement, route: string, event: string) {

    const handler = ({ data }: Match) => {

        if(isWaitingForSettled(element)) {
            logger(
                LogLevel.DEBUG,
                `Route matched (${ route }), but waiting for settled element before triggering`,
                element,
            );
            return;
        }

        logger(LogLevel.DEBUG, `Route matched (${ route }), triggering HTMX event (${ event }) on element`, element);
        htmx.trigger(element, event, { routerPathData: data } as RequestConfigEventDetail);
    };

    navigoRouter.on(route, handler);
    ROUTE_HANDLER_MAP.set(element, handler);
}

function bindCleanOnExitRouteBehaviourOnElement(element: HTMLElement, route: string) {

    if(element.hasAttribute(HTMX_CLEAN_ON_EXIT_ROUTE_ATTRIBUTE)) {

        navigoRouter.addLeaveHook(route, (done: HookCleanupFunction) => {
            element.innerHTML = "";
            element.setAttribute(HTMX_CLEAN_ON_EXIT_ROUTE_ATTRIBUTE, "true");
            done();
        });

        for(const handler in CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS) {
            element.addEventListener(handler, CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS[handler]);
        }
    }
}

function addDisableRouteRemovalObserver(element: HTMLElement, route: string) {

    const observer = new MutationObserver((_, observer) => {
        if(DomUtils.wasElementRemoved(element)) {

            logger(LogLevel.INFO, "Element removed, removing related handler for route", element, route);

            observer.disconnect();

            const handler = ROUTE_HANDLER_MAP.get(element);

            if(handler) {
                navigoRouter.off(handler);
                ROUTE_HANDLER_MAP.delete(element);
            }
        }
    });

    observer.observe(document, { childList: true, subtree: true });
}

/**
 * Checks whether the element must wait for another element to settle before triggering.
 *
 * Reads the wait-for-settled attribute and checks whether the referenced element
 * has already been marked as settled.
 *
 * @param element - The HTML element to check.
 * @returns true if the element must wait (target exists and has not settled yet), false otherwise.
 *
 * @author GitHub Copilot
 */
function isWaitingForSettled(element: HTMLElement): boolean {

    const waitForSelector = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_SETTLED_ATTRIBUTE);

    if(!waitForSelector) {
        return false;
    }

    const targetElement = document.querySelector(waitForSelector) as HTMLElement;

    return !!targetElement && !targetElement.hasAttribute(HTMX_TRIGGER_ON_ROUTE_SETTLED_FLAG);
}

/**
 * Defers the htmx trigger until the element referenced by waitForSelector has settled.
 *
 * Listens for the htmx:afterSettle event on the target element. When the event fires
 * (and the event target matches the waited element), the settled flag is set on the target
 * and the htmx event is triggered on the route element. If the target has already settled,
 * the trigger fires immediately.
 *
 * @param waitForSelector - CSS selector for the element to wait for.
 * @param element - The HTML element to trigger the event on.
 * @param event - The event name to trigger.
 * @param routerMatch - The matched route data from the router.
 *
 * @author GitHub Copilot
 */
function executeAfterElementSettled(
    waitForSelector: string,
    element: HTMLElement,
    event: string,
    routerMatch: Match,
) {

    const targetElement = document.querySelector(waitForSelector) as HTMLElement;

    if(!targetElement) {
        logger(LogLevel.WARN, "Wait-for-settled target element not found", waitForSelector);
        return;
    }

    bindSettledFlagCleanupOnNewRequest(targetElement);

    if(targetElement.hasAttribute(HTMX_TRIGGER_ON_ROUTE_SETTLED_FLAG)) {
        htmx.trigger(element, event, { routerPathData: routerMatch.data } as RequestConfigEventDetail);
        return;
    }

    const settleCleanupRef: { observer?: MutationObserver } = {};

    const settleHandler = (settleEvent: Event) => {

        if(settleEvent.target === targetElement) {
            cleanupSettleWait(targetElement, settleHandler, settleCleanupRef.observer);

            if(areDependenciesReady(element, targetElement)) {
                markSettledAndTrigger(targetElement, element, event, routerMatch);
            }
            else {
                waitForDependenciesReadyAndTrigger(element, targetElement, event, routerMatch);
            }
        }
    };

    targetElement.addEventListener("htmx:afterSettle", settleHandler);

    settleCleanupRef.observer = addSettleListenerRemovalObserver(element, targetElement, settleHandler);
}

/**
 * Binds a listener that clears the settled flag when the target element starts a new HTMX request.
 *
 * Uses a data attribute to ensure only one cleanup listener is bound per target element.
 *
 * @param targetElement - The element to observe for new HTMX requests.
 *
 * @author GitHub Copilot
 */
function bindSettledFlagCleanupOnNewRequest(targetElement: HTMLElement) {

    if(targetElement.hasAttribute(HTMX_TRIGGER_ON_ROUTE_SETTLED_RESET_BOUND_FLAG)) {
        return;
    }

    targetElement.addEventListener("htmx:beforeRequest", (event: Event) => {

        if(event.target === targetElement) {
            targetElement.removeAttribute(HTMX_TRIGGER_ON_ROUTE_SETTLED_FLAG);
        }
    });

    targetElement.setAttribute(HTMX_TRIGGER_ON_ROUTE_SETTLED_RESET_BOUND_FLAG, "true");
}

/**
 * Removes the settle event listener from the target element and disconnects the removal observer.
 *
 * @param targetElement - The element the settle listener is attached to.
 * @param settleHandler - The settle event handler to remove.
 * @param observer - The MutationObserver to disconnect, if present.
 *
 * @author GitHub Copilot
 */
function cleanupSettleWait(
    targetElement: HTMLElement,
    settleHandler: (settleEvent: Event) => void,
    observer?: MutationObserver,
) {

    targetElement.removeEventListener("htmx:afterSettle", settleHandler);
    observer?.disconnect();
}

/**
 * Sets the settled flag on the target element and triggers the htmx event on the waiting element.
 *
 * @param targetElement - The element to mark as settled.
 * @param element - The waiting element to trigger the event on.
 * @param event - The event name to trigger.
 * @param routerMatch - The matched route data from the router.
 *
 * @author GitHub Copilot
 */
function markSettledAndTrigger(
    targetElement: HTMLElement,
    element: HTMLElement,
    event: string,
    routerMatch: Match,
) {

    targetElement.setAttribute(HTMX_TRIGGER_ON_ROUTE_SETTLED_FLAG, "true");
    htmx.trigger(element, event, { routerPathData: routerMatch.data } as RequestConfigEventDetail);
}

/**
 * Checks whether the target element's dependencies are ready.
 *
 * Returns true if the waiting element does not require dependencies, or if the target
 * element has the dependencies-ready flag set.
 *
 * @param element - The waiting element that may require dependencies.
 * @param targetElement - The target element to check for the dependencies-ready flag.
 * @returns true if dependencies are satisfied, false otherwise.
 *
 * @author GitHub Copilot
 */
function areDependenciesReady(element: HTMLElement, targetElement: HTMLElement): boolean {

    if(!element.hasAttribute(HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_DEPENDENCIES_READY_ATTRIBUTE)) {
        return true;
    }

    return targetElement.hasAttribute(HTMX_TRIGGER_ON_ROUTE_DEPENDENCIES_READY_FLAG);
}

/**
 * Waits for the dependencies-ready flag on the target element, then marks it as settled and
 * triggers the htmx event on the waiting element.
 *
 * Uses a MutationObserver on the target element's attributes and includes a removal observer
 * for cleanup if the waiting element is removed during the wait.
 *
 * @param element - The waiting element to trigger the event on.
 * @param targetElement - The target element to observe for the dependencies-ready flag.
 * @param event - The event name to trigger.
 * @param routerMatch - The matched route data from the router.
 *
 * @author GitHub Copilot
 */
function waitForDependenciesReadyAndTrigger(
    element: HTMLElement,
    targetElement: HTMLElement,
    event: string,
    routerMatch: Match,
) {

    logger(
        LogLevel.DEBUG,
        "Target element settled but dependencies not ready, waiting for dependencies",
        element,
    );

    const dependenciesObserver = new MutationObserver(() => {

        if(targetElement.hasAttribute(HTMX_TRIGGER_ON_ROUTE_DEPENDENCIES_READY_FLAG)) {
            dependenciesObserver.disconnect();
            removalObserver.disconnect();
            markSettledAndTrigger(targetElement, element, event, routerMatch);
        }
    });

    dependenciesObserver.observe(targetElement, {
        attributes: true,
        attributeFilter: [HTMX_TRIGGER_ON_ROUTE_DEPENDENCIES_READY_FLAG],
    });

    const removalObserver = new MutationObserver(() => {

        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Waiting element removed, cleaning up dependencies observer", element);
            dependenciesObserver.disconnect();
            removalObserver.disconnect();
        }
    });

    removalObserver.observe(document, { childList: true, subtree: true });
}

/**
 * Observes the DOM for the removal of the waiting element and cleans up the settle listener
 *
 * on the target element to prevent event listener leaks.
 *
 * @param element - The waiting element to observe for removal.
 * @param targetElement - The element the settle listener is attached to.
 * @param settleHandler - The settle event handler to remove on cleanup.
 * @returns The MutationObserver instance, so it can be disconnected externally when no longer needed.
 *
 * @author GitHub Copilot
 */
function addSettleListenerRemovalObserver(
    element: HTMLElement,
    targetElement: HTMLElement,
    settleHandler: (settleEvent: Event) => void,
): MutationObserver {

    const observer = new MutationObserver((_, observer) => {

        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Waiting element removed, cleaning up settle listener", element);
            cleanupSettleWait(targetElement, settleHandler, observer);
        }
    });

    observer.observe(document, { childList: true, subtree: true });

    return observer;
}

/**
 * Executes the htmx trigger immediately if the current route matches the provided route.
 *
 * When the element has a wait-for-settled attribute, defers execution until the referenced
 * element has settled its htmx request, preventing race conditions on dependent lazy-loaded components.
 * Otherwise, uses setTimeout to defer the trigger, allowing htmx to fully process the element's trigger
 * setup before the event is dispatched.
 *
 * @param route - The route pattern to match against the current location.
 * @param element - The HTML element to trigger the event on.
 * @param event - The event name to trigger.
 *
 * @author GitHub Copilot
 */
function executeImmediatelyIfOnRoute(route: string, element: HTMLElement, event: string) {

    const routerMatch = navigoRouter.matchLocation(route);

    if(routerMatch) {

        const waitForSelector = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_SETTLED_ATTRIBUTE);

        if(waitForSelector) {
            executeAfterElementSettled(waitForSelector, element, event, routerMatch);
        }
        else {
            window.setTimeout(() => {
                htmx.trigger(element, event, { routerPathData: routerMatch.data } as RequestConfigEventDetail);
            }, 500);
        }
    }
}
