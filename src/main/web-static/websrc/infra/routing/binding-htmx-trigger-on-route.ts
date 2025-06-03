import htmx from "htmx.org";
import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";
import { HookCleanupFunction, navigoRouter } from "./routing-navigo";
import { EventDetail } from "../htmx";
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
    const htmxRoutedElements = DomUtils.queryAllInDescendants(
        element,
        `[${ HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE }]:not([${ HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG }])`,
    );
    bindRouteToHTMXEventOnElements(htmxRoutedElements);
}

function bindRouteToHTMXEventOnElements(htmxRoutedElements: NodeListOf<HTMLElement>) {

    htmxRoutedElements.forEach((element) => {

        logger(LogLevel.INFO, "Binding HTMX event on route for element", element);

        const { route, event } = extractBindingData(element);

        bindRouteToHTMXTriggerOnElement(element, route, event);
        bindCleanOnExitRouteBehaviourOnElement(element, route);
        addDisableRouteRemovalObserver(element, route);

        element.setAttribute(HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG, "true");

        executeImmediatelyIfOnRoute(route, element, event);
    });
}

function extractBindingData(element: HTMLElement) {
    const routeToEvent = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE);
    const [route, event] = routeToEvent.split(HTMX_TRIGGER_ON_ROUTE_EVENT_SEPARATOR);
    return { route, event };
}

function bindRouteToHTMXTriggerOnElement(element: HTMLElement, route: string, event: string) {

    const handler = ({ data }: Match) => {
        htmx.trigger(element, event, { routerPathData: data } as EventDetail);
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

function executeImmediatelyIfOnRoute(route: string, element: HTMLElement, event: string) {

    const routerMatch = navigoRouter.matchLocation(route);

    if(routerMatch) {
        htmx.trigger(element, event, { routerPathData: routerMatch.data } as EventDetail);
    }
}
