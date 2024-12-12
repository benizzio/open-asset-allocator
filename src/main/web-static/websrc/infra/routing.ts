import Navigo from "navigo";
import { LogLevel, logger } from "./logging";
import DomUtils from "./dom-utils";
import htmx from "htmx.org";

type HookCleanupFunction = (success?: boolean) => void;

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

// =============================================================================
// NAVIGATE TO
// triggers navigation to a route on click and enter key press
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

// =============================================================================
// SHOW ON ROUTE
// binds elements to trigger navigation to a route
// =============================================================================

const DISPLAY_ON_ROUTE_ATTRIBUTE = "data-display-on-route";
const DISPLAY_ON_ROUTE_BOUND_FLAG = "display-on-route-bound";

function configDisplayOnRouteHooks(route: string, element: HTMLElement) {

    const afterHookCleanup = navigoRouter.addAfterHook(route, () => {
        element.style.display = "inline";
    }) as () => void;

    const leaveHookCleanup = navigoRouter.addLeaveHook(route, (done: HookCleanupFunction) => {
        element.style.display = "none";
        done();
    }) as () => void;

    return { afterHookCleanup, leaveHookCleanup };
}

function addDisplayOnRouteRemovalObserver(
    element: HTMLElement,
    route: string,
    afterHookCleanup: () => void,
    leaveHookCleanup: () => void,
) {
    const observer = new MutationObserver((_, observer) => {
        if (DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Element removed, removing display on route action", element, route);
            observer.disconnect();
            afterHookCleanup();
            leaveHookCleanup();
        }
    });
    observer.observe(document, { childList: true, subtree: true });
}

function bindDisplayOnRoute(element: HTMLElement) {

    if(!element.getAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG)) {

        const route = element.getAttribute(DISPLAY_ON_ROUTE_ATTRIBUTE);
        logger(LogLevel.INFO, "Binding display on route hooks for element", element, route);

        const { afterHookCleanup, leaveHookCleanup } = configDisplayOnRouteHooks(route, element);
        addDisplayOnRouteRemovalObserver(element, route, afterHookCleanup, leaveHookCleanup);

        element.setAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG, "true");
    }
}

function bindDisplayOnRouteElements(displayOnRouteElements: NodeListOf<HTMLElement>) {
    displayOnRouteElements.forEach((element) => {
        bindDisplayOnRoute(element);
    });
}

// =============================================================================
// Module object
// =============================================================================

const router = {
    resolveBrowserRoute() {
        navigoRouter.resolve();
    },
    bindHTMXRouting() {
        const htmxRoutedElements = DomUtils.queryAll(`[${HTMX_EVENT_ON_ROUTE_ATTRIBUTE}]`);
        bindRouteToHTMXEventOnElements(htmxRoutedElements);
    },
    bindNavigateTo() {
        const navigationElements = DomUtils.queryAll(`[${NAVIGATE_TO_ATTRIBUTE}]`);
        bindNavigateToElements(navigationElements);
    },
    bindNavigateToInDescendants(element: HTMLElement) {
        const navigationElements = DomUtils.queryAllInDescendants(element, `[${NAVIGATE_TO_ATTRIBUTE}]`);
        bindNavigateToElements(navigationElements);
    },
    bindDisplayOnRoute() {
        const displayOnRouteElements = DomUtils.queryAll(`[${DISPLAY_ON_ROUTE_ATTRIBUTE}]`);
        bindDisplayOnRouteElements(displayOnRouteElements);
    },
};

export default router;