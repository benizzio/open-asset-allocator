import { navigoRouter } from "./routing-navigo";
import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";

// =============================================================================
// DISPLAY ON ROUTE
// binds elements to change display based on route
// =============================================================================

const DISPLAY_ON_ROUTE_ATTRIBUTE = "data-display-on-route";
const DISPLAY_ON_ROUTE_BOUND_FLAG = "display-on-route-bound";

export function bindDisplayOnRouteInDescendants(element: HTMLElement) {
    const displayOnRouteElements = DomUtils.queryAllInDescendants(element, `[${ DISPLAY_ON_ROUTE_ATTRIBUTE }]`);
    bindDisplayOnRouteElements(displayOnRouteElements);
}

function bindDisplayOnRouteElements(displayOnRouteElements: NodeListOf<HTMLElement>) {
    displayOnRouteElements.forEach((element) => {
        bindDisplayOnRoute(element);
    });
}

function bindDisplayOnRoute(element: HTMLElement) {

    if(!element.getAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG)) {//TODO replace this if for improving the element selector

        const route = element.getAttribute(DISPLAY_ON_ROUTE_ATTRIBUTE);

        logger(LogLevel.INFO, "Binding display on route hooks for element", element, route);

        const cleanupCallback = configDisplayOnRouteHooks(route, element);
        addDisplayOnRouteRemovalObserver(element, cleanupCallback);

        element.setAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG, "true");

        executeImmediatelyIfOnRoute(route, element);
    }
}

function configDisplayOnRouteHooks(route: string, element: HTMLElement): () => void {

    const hasRoute = !!navigoRouter.getRoute(route);

    const displayCallback = () => {
        changeElementDisplay(element, true);
    };

    let cleanupCallback: () => void;

    if(hasRoute) {
        cleanupCallback = navigoRouter.addAfterHook(route, displayCallback) as () => void;
    }
    else {
        navigoRouter.on(route, displayCallback);

        cleanupCallback = () => {
            navigoRouter.off(displayCallback);
        };
    }

    return cleanupCallback;
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


function executeImmediatelyIfOnRoute(route: string, element: HTMLElement) {
    changeElementDisplay(element, !!navigoRouter.matchLocation(route));
}
