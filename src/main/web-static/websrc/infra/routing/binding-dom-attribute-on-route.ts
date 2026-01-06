import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";
import { HookCleanupFunction, navigoRouter } from "./routing-navigo";

// =============================================================================
// ATTRIBUTE ON ROUTE
// binds elements to contain an attribute based on route
// attribute value format is "route[attribute=value,attribute]"
// =============================================================================

const ATTRIBUTE_ON_ROUTE_ATTRIBUTE = "data-attribute-on-route";
const ATTRIBUTE_ON_ROUTE_BOUND_FLAG = "attribute-on-route-bound";
const ATTRIBUTE_ON_ROUTE_ATTRIBUTE_SEPARATOR = ",";
const ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_SEPARATOR = "=";
const ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_GROUP_PREFIX = "[";
const ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_GROUP_SUFFIX = "]";

export function bindAttributeOnRouteInDescendants(element: HTMLElement) {
    const attributeOnRouteElements = DomUtils.queryAllInDescendants(
        element,
        `[${ ATTRIBUTE_ON_ROUTE_ATTRIBUTE }]:not([${ ATTRIBUTE_ON_ROUTE_BOUND_FLAG }])`,
    );
    bindAttributeOnRouteElements(attributeOnRouteElements);
}

function bindAttributeOnRouteElements(attributeOnRouteElements: NodeListOf<HTMLElement>) {

    attributeOnRouteElements.forEach((element) => {

        element.setAttribute(ATTRIBUTE_ON_ROUTE_BOUND_FLAG, "binding");

        try {
            const isBound = bindAttributeOnRoute(element);

            if(!isBound) {
                element.removeAttribute(ATTRIBUTE_ON_ROUTE_BOUND_FLAG);
                return;
            }

            element.setAttribute(ATTRIBUTE_ON_ROUTE_BOUND_FLAG, "true");
        } catch(error) {
            element.removeAttribute(ATTRIBUTE_ON_ROUTE_BOUND_FLAG);
            throw error;
        }
    });
}

function bindAttributeOnRoute(element: HTMLElement): boolean {

    const { route, attributesStrings } = extractBindingData(element);

    logger(LogLevel.INFO, "Binding attributes on route for element", element, route, attributesStrings);

    if(attributesStrings.length === 0) {
        return false;
    }

    addRouterHooks(route, attributesStrings, element);

    executeImmediatelyIfOnRoute(route, attributesStrings, element);

    return true;
}

function extractBindingData(element: HTMLElement) {

    const attributeOnRouteValue = element.getAttribute(ATTRIBUTE_ON_ROUTE_ATTRIBUTE);

    const route = attributeOnRouteValue.substring(
        0,
        attributeOnRouteValue.indexOf(ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_GROUP_PREFIX),
    );

    //"attribute1=value1,attribute"
    const attributesString = attributeOnRouteValue.substring(
        attributeOnRouteValue.indexOf(ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_GROUP_PREFIX) + 1,
        attributeOnRouteValue.indexOf(ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_GROUP_SUFFIX),
    );

    //["attribute1=value1","attribute2=value2"]
    const attributesStrings = attributesString.split(ATTRIBUTE_ON_ROUTE_ATTRIBUTE_SEPARATOR);
    return { route, attributesStrings };
}

function addRouterHooks(route: string, attributesStrings: string[], element: HTMLElement) {

    const afterHookCleanup = navigoRouter.addAfterHook(route, () => {
        addAttributes(attributesStrings, element);
    }) as () => void;

    const leaveHookCleanup = navigoRouter.addLeaveHook(route, (done: HookCleanupFunction) => {
        removeAttributes(attributesStrings, element);
        done();
    }) as () => void;

    const observer = new MutationObserver((_, observer) => {
        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Element removed, removing attribute on route action", element);
            observer.disconnect();
            afterHookCleanup();
            leaveHookCleanup();
        }
    });

    observer.observe(document, { childList: true, subtree: true });
}

function addAttributes(attributesStrings: string[], element: HTMLElement) {
    attributesStrings.forEach((attributeString) => {
        const [attributeName, attributeValue] = attributeString.split(
            ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_SEPARATOR);
        element.setAttribute(attributeName, attributeValue || "");
    });
}

function removeAttributes(attributesStrings: string[], element: HTMLElement) {
    attributesStrings.forEach((attributeString) => {
        const [attributeName] = attributeString.split(ATTRIBUTE_ON_ROUTE_ATTRIBUTE_NAME_VALUE_SEPARATOR);
        element.removeAttribute(attributeName);
    });
}

function executeImmediatelyIfOnRoute(route: string, attributesStrings: string[], element: HTMLElement) {
    if(navigoRouter.matchLocation(route)) {
        addAttributes(attributesStrings, element);
    }
    else {
        removeAttributes(attributesStrings, element);
    }
}
