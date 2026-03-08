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

const HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_READY_ATTRIBUTE = "data-hx-trigger-on-route-wait-for-ready";
const HTMX_TRIGGER_ON_ROUTE_READY_FLAG = "data-hx-trigger-on-route-ready";

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

        if(isWaitingForReady(element)) {
            logger(
                LogLevel.DEBUG,
                `Route matched (${ route }), but waiting for ready conditions before triggering`,
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
 * Checks whether the element must wait for its conditions to be ready before triggering.
 *
 * Reads the wait-for-ready attribute for required conditions and checks whether
 * all conditions are present on the element's own ready flag.
 *
 * @param element - The HTML element to check.
 * @returns true if the element must wait (conditions not yet met), false otherwise.
 *
 * @author GitHub Copilot
 */
function isWaitingForReady(element: HTMLElement): boolean {

    const requiredConditions = parseWaitForReadyConditions(element);

    if(requiredConditions.length === 0) {
        return false;
    }

    return !areAllConditionsReady(element, requiredConditions);
}

/**
 * Parses the wait-for-ready attribute value into a list of required conditions.
 *
 * The expected format is "condition1,condition2,...".
 *
 * @param element - The HTML element to read the attribute from.
 * @returns The list of required conditions, or an empty array if the attribute is not set.
 *
 * @author GitHub Copilot
 */
function parseWaitForReadyConditions(element: HTMLElement): string[] {

    const attributeValue = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_READY_ATTRIBUTE);

    if(!attributeValue) {
        return [];
    }

    return attributeValue.split(",").map(c => c.trim()).filter(Boolean);
}

/**
 * Checks whether all required conditions are present in the element's ready flag.
 *
 * @param element - The element to check.
 * @param requiredConditions - The conditions that must all be present.
 * @returns true if all conditions are met, false otherwise.
 *
 * @author GitHub Copilot
 */
function areAllConditionsReady(element: HTMLElement, requiredConditions: string[]): boolean {

    const readyValue = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_READY_FLAG) || "";
    const fulfilledConditions = readyValue.split(",").map(c => c.trim()).filter(Boolean);

    return requiredConditions.every(condition => fulfilledConditions.includes(condition));
}

/**
 * Defers the htmx trigger until all required conditions are met on the element itself.
 *
 * Reads the wait-for-ready attribute to get the required conditions.
 * If all conditions are already met, triggers immediately. Otherwise, sets up a
 * MutationObserver on the element's ready flag attribute and waits for all
 * conditions to be fulfilled before triggering.
 *
 * @param element - The HTML element to trigger the event on.
 * @param event - The event name to trigger.
 * @param routerMatch - The matched route data from the router.
 *
 * @author GitHub Copilot
 */
function executeAfterElementReady(
    element: HTMLElement,
    event: string,
    routerMatch: Match,
) {

    const requiredConditions = parseWaitForReadyConditions(element);

    if(requiredConditions.length === 0) {
        return;
    }

    if(areAllConditionsReady(element, requiredConditions)) {
        htmx.trigger(element, event, { routerPathData: routerMatch.data } as RequestConfigEventDetail);
        return;
    }

    logger(
        LogLevel.DEBUG,
        "Conditions not yet met, waiting for ready conditions on element",
        element,
        requiredConditions,
    );

    const readyObserver = new MutationObserver(() => {

        if(areAllConditionsReady(element, requiredConditions)) {
            readyObserver.disconnect();
            removalObserver.disconnect();
            htmx.trigger(element, event, { routerPathData: routerMatch.data } as RequestConfigEventDetail);
        }
    });

    readyObserver.observe(element, {
        attributes: true,
        attributeFilter: [HTMX_TRIGGER_ON_ROUTE_READY_FLAG],
    });

    const removalObserver = new MutationObserver(() => {

        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Waiting element removed, cleaning up ready observer", element);
            readyObserver.disconnect();
            removalObserver.disconnect();
        }
    });

    removalObserver.observe(document, { childList: true, subtree: true });
}

/**
 * Adds a condition value to the ready flag of all elements matching the selector.
 *
 * Appends the condition to the comma-separated ready flag list if not already present.
 * The MutationObserver on waiting elements will detect this change and check
 * whether all required conditions are now fulfilled.
 *
 * @param selector - CSS selector to find the elements to add the condition to.
 * @param condition - The condition value to add.
 *
 * @example
 * addRouteReadyCondition("[data-hx-trigger-on-route-wait-for-ready]", "settled");
 * addRouteReadyCondition("[data-hx-trigger-on-route-wait-for-ready]", "partials-registered");
 *
 * @author GitHub Copilot
 */
export function addRouteReadyCondition(selector: string, condition: string) {

    const elements = document.querySelectorAll<HTMLElement>(selector);

    elements.forEach(element => {

        const currentValue = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_READY_FLAG) || "";

        const currentConditions = currentValue
            ? currentValue.split(",").map(c => c.trim()).filter(Boolean)
            : [];

        if(!currentConditions.includes(condition)) {
            currentConditions.push(condition);
            element.setAttribute(HTMX_TRIGGER_ON_ROUTE_READY_FLAG, currentConditions.join(","));
        }
    });
}

/**
 * Executes the htmx trigger immediately if the current route matches the provided route.
 *
 * When the element has a wait-for-ready attribute, defers execution until all required
 * conditions are met on the element itself. Otherwise, uses setTimeout to defer the trigger,
 * allowing htmx to fully process the element's trigger setup before the event is dispatched.
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

        const waitForReadyValue = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_WAIT_FOR_READY_ATTRIBUTE);

        if(waitForReadyValue) {
            executeAfterElementReady(element, event, routerMatch);
        }
        else {
            window.setTimeout(() => {
                htmx.trigger(element, event, { routerPathData: routerMatch.data } as RequestConfigEventDetail);
            }, 500);
        }
    }
}
