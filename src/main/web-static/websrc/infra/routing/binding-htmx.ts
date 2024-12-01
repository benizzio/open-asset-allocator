// =============================================================================
// HTMX EVENT ON ROUTE
// triggers a custom HTMX event when a route is navigated to
// attribute value format is "route:event"
// =============================================================================

import htmx from "htmx.org";
import DomUtils from "../dom-utils";
import { logger, LogLevel } from "../logging";
import { navigoRouter } from "./routing-navigo";

const HTMX_EVENT_ON_ROUTE_ATTRIBUTE = "data-htmx-event-on-route";
const HTMX_EVENT_ON_ROUTE_BOUND_FLAG = "htmx-event-on-route-bound";

function bindRouteToHTMXEventOnElement(element: HTMLElement, route: string, event: string) {
    navigoRouter.on(route, () => {
        htmx.trigger(element, event);
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

function bindRouteToHTMXEventOnElements(htmxRoutedElements: NodeListOf<HTMLElement>) {

    htmxRoutedElements.forEach((element) => {

        if(!element.getAttribute(HTMX_EVENT_ON_ROUTE_BOUND_FLAG)) {

            logger(LogLevel.INFO, "Binding HTMX event on route for element", element);
            const routeToEvent = element.getAttribute(HTMX_EVENT_ON_ROUTE_ATTRIBUTE);
            const [route, event] = routeToEvent.split(":");
            bindRouteToHTMXEventOnElement(element, route, event);
            addDisableRouteRemovalObserver(element, route);

            element.setAttribute(HTMX_EVENT_ON_ROUTE_BOUND_FLAG, "true");
        }
    });
}

export function bindHTMXEventOnRouteInDescendants(element: HTMLElement) {
    const htmxRoutedElements = DomUtils.queryAllInDescendants(element, `[${HTMX_EVENT_ON_ROUTE_ATTRIBUTE}]`);
    bindRouteToHTMXEventOnElements(htmxRoutedElements);
}