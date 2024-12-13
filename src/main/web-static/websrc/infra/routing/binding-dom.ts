import { HookCleanupFunction, navigoRouter } from "./routing-navigo";
import DomUtils from "../dom-utils";
import { logger, LogLevel } from "../logging";

// =============================================================================
// DISPLAY ON ROUTE
// binds elements to change display based on route
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
    afterHookCleanup: () => void,
    leaveHookCleanup: () => void,
) {
    const observer = new MutationObserver((_, observer) => {
        if (DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Element removed, removing display on route action", element);
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
        addDisplayOnRouteRemovalObserver(element, afterHookCleanup, leaveHookCleanup);

        element.setAttribute(DISPLAY_ON_ROUTE_BOUND_FLAG, "true");
    }
}

function bindDisplayOnRouteElements(displayOnRouteElements: NodeListOf<HTMLElement>) {
    displayOnRouteElements.forEach((element) => {
        bindDisplayOnRoute(element);
    });
}

export function bindDisplayOnRouteInDescendants(element: HTMLElement) {
    const displayOnRouteElements = DomUtils.queryAllInDescendants(element, `[${DISPLAY_ON_ROUTE_ATTRIBUTE}]`);
    bindDisplayOnRouteElements(displayOnRouteElements);
}

// =============================================================================
// ATTRIBUTE ON ROUTE
// binds elements to contain an attribute based on route
// attribute value format is "route:attribute=value"
// =============================================================================

const ATTRIBUTE_ON_ROUTE_ATTRIBUTE = "data-attribute-on-route";
const ATTRIBUTE_ON_ROUTE_BOUND_FLAG = "attribute-on-route-bound";

const ADD_ATTRIBUTE_FUNCTION =
    (element: HTMLElement, attributeName: string, attributeValue: string) =>
        element.setAttribute(attributeName, attributeValue || "");

const REMOVE_ATTRIBUTE_FUNCTION =
    (element: HTMLElement, attributeName: string, done?: HookCleanupFunction) => {
        element.removeAttribute(attributeName);
        done();
    };

function bindAttributeOnRoute(element: HTMLElement) {

    if(!element.getAttribute(ATTRIBUTE_ON_ROUTE_BOUND_FLAG)) {

        const attributeOnRouteValue = element.getAttribute(ATTRIBUTE_ON_ROUTE_ATTRIBUTE);
        const [route, attribute] = attributeOnRouteValue.split(":");
        const [attributeName, attributeValue] = attribute.split("=");
        logger(LogLevel.INFO, "Binding attribute on route for element", element, route, attribute);

        const afterHookCleanup = navigoRouter.addAfterHook(route, () => {
            ADD_ATTRIBUTE_FUNCTION(element, attributeName, attributeValue);
        }) as () => void;

        const leaveHookCleanup = navigoRouter.addLeaveHook(route, (done: HookCleanupFunction) => {
            REMOVE_ATTRIBUTE_FUNCTION(element, attributeName, done);
        }) as () => void;

        const observer = new MutationObserver((_, observer) => {
            if (DomUtils.wasElementRemoved(element)) {
                logger(LogLevel.INFO, "Element removed, removing attribute on route action", element);
                observer.disconnect();
                afterHookCleanup();
                leaveHookCleanup();
            }
        });
        observer.observe(document, { childList: true, subtree: true });

        if(navigoRouter.matchLocation(route)) {
            ADD_ATTRIBUTE_FUNCTION(element, attributeName, attributeValue);
        }
        else {
            REMOVE_ATTRIBUTE_FUNCTION(element, attributeName, () => {});
        }

        element.setAttribute(ATTRIBUTE_ON_ROUTE_BOUND_FLAG, "true");
    }
}

function bindAttributeOnRouteElements(attributeOnRouteElements: NodeListOf<HTMLElement>) {
    attributeOnRouteElements.forEach((element) => {
        bindAttributeOnRoute(element);
    });
}

export function bindAttributeOnRouteInDescendants(element: HTMLElement) {
    const attributeOnRouteElements = DomUtils.queryAllInDescendants(element, `[${ATTRIBUTE_ON_ROUTE_ATTRIBUTE}]`);
    bindAttributeOnRouteElements(attributeOnRouteElements);
}