import htmx from "htmx.org";
import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";
import { navigoRouter } from "./routing-navigo";
import { EventDetail } from "../htmx";

// =============================================================================
// HTMX TRIGGER EVENT ON ROUTE
// triggers a custom HTMX event when a route is navigated to
// attribute value format is "route/:path-param!event"
// =============================================================================

const HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE = "data-hx-trigger-on-route";
const HTMX_TRIGGER_ON_ROUTE_EVENT_SEPARATOR = "!";
const HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG = "hx-trigger-on-route-bound";

export function bindHTMXTriggerOnRouteInDescendants(element: HTMLElement) {
    const htmxRoutedElements = DomUtils.queryAllInDescendants(element, `[${ HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE }]`);
    bindRouteToHTMXEventOnElements(htmxRoutedElements);
}

function bindRouteToHTMXEventOnElements(htmxRoutedElements: NodeListOf<HTMLElement>) {

    htmxRoutedElements.forEach((element) => {

        if(!element.getAttribute(HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG)) {

            logger(LogLevel.INFO, "Binding HTMX event on route for element", element);

            const { route, event } = extractBindingData(element);

            bindRouteToHTMXTriggerOnElement(element, route, event);
            addDisableRouteRemovalObserver(element, route);

            element.setAttribute(HTMX_TRIGGER_ON_ROUTE_BOUND_FLAG, "true");

            executeImmediatelyIfOnRoute(route, element, event);
        }
    });
}

function extractBindingData(element: HTMLElement) {
    const routeToEvent = element.getAttribute(HTMX_TRIGGER_ON_ROUTE_ATTRIBUTE);
    const [route, event] = routeToEvent.split(HTMX_TRIGGER_ON_ROUTE_EVENT_SEPARATOR);
    return { route, event };
}

function bindRouteToHTMXTriggerOnElement(element: HTMLElement, route: string, event: string) {
    navigoRouter.on(route, ({ data }) => {
        htmx.trigger(element, event, { routerPathData: data } as EventDetail);
    });
    return route;
}

function addDisableRouteRemovalObserver(element: HTMLElement, route: string) {
    const observer = new MutationObserver((_, observer) => {
        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Element removed, removing route", element, route);
            observer.disconnect();
            navigoRouter.off(route);
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
