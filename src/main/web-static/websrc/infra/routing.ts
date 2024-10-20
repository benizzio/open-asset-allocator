import Navigo from "navigo";
import { LogLevel, logger } from "./logging";
import DomUtils from "./dom-utils";
import htmx from "htmx.org";

const navigoRouter = new Navigo("/");

// =============================================================================
// HTMX EVENT ON ROUTE
// triggers a custom HTMX event when a route is navigated to
// =============================================================================

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
        if (!document.body.contains(element)) {
            logger(LogLevel.INFO, "Element removed, removing route", element, route);
            observer.disconnect();
            navigoRouter.off(route);
        }
    });
    observer.observe(element.parentElement, { childList: true });
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

// =============================================================================
// NAVIGATE TO
// binds elements to trigger navigation to a route
// =============================================================================

const NAVIGATE_TO_ATTRIBUTE = "data-navigate-to";
const NAVIGATE_TO_BOUND_FLAG = "navigate-to-bound";

function bindKeypressNavigation(element: HTMLElement) {
    element.addEventListener("keypress", (event: KeyboardEvent) => {
        if (event.key === "Enter" || event.key === " ") {
            navigoRouter.navigate(element.getAttribute(NAVIGATE_TO_ATTRIBUTE));
        }
    });
}

function bindCLickNavigation(element: HTMLElement) {
    element.addEventListener("click", () => {
        navigoRouter.navigate(element.getAttribute(NAVIGATE_TO_ATTRIBUTE));
    });
}

function bindNavigateToElements(navigationElement: NodeListOf<HTMLElement>) {
    navigationElement.forEach((element) => {
        if (!element.getAttribute(NAVIGATE_TO_BOUND_FLAG)) {
            logger(LogLevel.INFO, "Binding navigation element", element);
            bindKeypressNavigation(element);
            bindCLickNavigation(element);
            element.setAttribute(NAVIGATE_TO_BOUND_FLAG, "true");
        }
    });
}

const router = {
    resolveBrowserRoute() {
        navigoRouter.resolve();
    },
    bindHTMXRouting() {
        const htmxRoutedElements = DomUtils.queryAll(`[${HTMX_EVENT_ON_ROUTE_ATTRIBUTE}]`);
        bindRouteToHTMXEventOnElements(htmxRoutedElements);
    },
    bindNavigateTo() {
        const navigationElement = DomUtils.queryAll(`[${NAVIGATE_TO_ATTRIBUTE}]`);
        bindNavigateToElements(navigationElement);
    },
};

export default router;