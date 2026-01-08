import { navigoRouter } from "./routing-navigo";
import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";

// =============================================================================
// DISPLAY ON ROUTE
// binds elements to change display based on route
// =============================================================================

const DISPLAY_ON_ROUTE_ATTRIBUTE = "data-display-on-route";
const DISPLAY_ON_REGULAR_EXPRESSION_ROUTE_ATTRIBUTE = "data-display-on-regexp-route";
const DISPLAY_ON_ROUTE_BOUND_FLAG = "display-on-route-bound";

export function bindDisplayOnRouteInDescendants(element: HTMLElement) {
    const displayOnRouteElements = element.querySelectorAll(
        `[${ DISPLAY_ON_ROUTE_ATTRIBUTE }]:not([${ DISPLAY_ON_ROUTE_BOUND_FLAG }])`
        + `, [${ DISPLAY_ON_REGULAR_EXPRESSION_ROUTE_ATTRIBUTE }]:not([${ DISPLAY_ON_ROUTE_BOUND_FLAG }])`,
    ) as NodeListOf<HTMLElement>;
    bindDisplayOnRouteElements(displayOnRouteElements);
}

function bindDisplayOnRouteElements(displayOnRouteElements: NodeListOf<HTMLElement>) {

    displayOnRouteElements.forEach((element) => {

        element.setAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG, "binding");

        try {
            const isBound = bindDisplayOnRoute(element);

            if(!isBound) {
                element.removeAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG);
                return;
            }

            element.setAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG, "true");
        } catch(error) {
            element.removeAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG);
            throw error;
        }
    });
}

function bindDisplayOnRoute(element: HTMLElement): boolean {

    const isRegularExpressionRoute = element.hasAttribute(DISPLAY_ON_REGULAR_EXPRESSION_ROUTE_ATTRIBUTE);

    let route: string | RegExp = isRegularExpressionRoute
        ? element.getAttribute(DISPLAY_ON_REGULAR_EXPRESSION_ROUTE_ATTRIBUTE)
        : element.getAttribute(DISPLAY_ON_ROUTE_ATTRIBUTE);

    if(!route) {
        return false;
    }

    route = isRegularExpressionRoute ? new RegExp(route, "g") : route;

    logger(LogLevel.INFO, "Binding display on route hooks for element", element, route);

    const cleanupCallback = configDisplayOnRouteHooks(route, element);
    addDisplayOnRouteRemovalObserver(element, cleanupCallback);

    executeImmediatelyIfOnRoute(route, element);

    return true;
}

function configDisplayOnRouteHooks(route: string | RegExp, element: HTMLElement): () => void {

    const displayCallback = () => {
        changeElementDisplay(element, true);
    };

    navigoRouter.on(route, displayCallback);

    return () => {
        navigoRouter.off(displayCallback);
    };
}

function changeElementDisplay(element: HTMLElement, display: boolean) {
    element.style.display = display ? null : "none";
}

function addDisplayOnRouteRemovalObserver(
    element: HTMLElement,
    cleanupCallback: () => void,
) {
    const observer = new MutationObserver((_, observer) => {
        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Element removed, removing display on route action", element);
            observer.disconnect();
            cleanupCallback();
        }
    });
    observer.observe(document, { childList: true, subtree: true });
}


function executeImmediatelyIfOnRoute(route: string | RegExp, element: HTMLElement) {
    changeElementDisplay(element, !!navigoRouter.matchLocation(route));
}
