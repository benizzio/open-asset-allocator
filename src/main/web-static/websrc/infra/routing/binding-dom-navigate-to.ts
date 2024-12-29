import { navigoRouter } from "./routing-navigo";
import { logger, LogLevel } from "../logging";
import DomUtils from "../dom/dom-utils";

// =============================================================================
// NAVIGATE TO
// triggers navigation to a route on click and enter key press
// =============================================================================

const NAVIGATE_TO_ATTRIBUTE = "data-navigate-to";
const NAVIGATE_TO_BOUND_FLAG = "navigate-to-bound";
const NAVIGO_PATH_PARAM_PREFIX = ":";

function hasPathParam(path: string) {
    return path.includes(NAVIGO_PATH_PARAM_PREFIX);
}

function bindKeypressNavigation(element: HTMLElement) {
    element.addEventListener("keypress", (event: KeyboardEvent) => {
        if(event.key === "Enter" || event.key === " ") {
            navigate(element);
        }
    });
}

function bindCLickNavigation(element: HTMLElement) {
    element.addEventListener("click", () => {
        navigate(element);
    });
}

function navigate(element: HTMLElement) {
    const destinationPath = buildDestinationPath(element);
    navigoRouter.navigate(destinationPath);
}

function buildDestinationPath(element: HTMLElement) {

    let destinationPath = element.getAttribute(NAVIGATE_TO_ATTRIBUTE);

    if(hasPathParam(destinationPath)) {
        destinationPath = resolvePathParamsFromCurrentLocationContext(destinationPath);
    }

    return destinationPath;
}

function resolvePathParamsFromCurrentLocationContext(path: string) {

    let resolvedPath = path;
    const currentLocation = navigoRouter.getCurrentLocation();
    const currentMatch = navigoRouter.match(currentLocation.url)[0];

    if(currentMatch) {
        for(const key in currentMatch.data) {
            resolvedPath = resolvedPath.replace(`${ NAVIGO_PATH_PARAM_PREFIX }${ key }`, currentMatch.data[key]);
        }
    }

    if(hasPathParam(resolvedPath)) {
        const errorMessage = "Could not find path param in current route navigate to";
        logger(LogLevel.ERROR, errorMessage, resolvedPath);
        throw new Error(errorMessage);
    }

    return resolvedPath;
}

function bindNavigateToElements(navigationElement: NodeListOf<HTMLElement>) {
    navigationElement.forEach((element) => {
        if(!element.getAttribute(NAVIGATE_TO_BOUND_FLAG)) {
            logger(LogLevel.INFO, "Binding navigation element", element);
            bindKeypressNavigation(element);
            bindCLickNavigation(element);
            element.setAttribute(NAVIGATE_TO_BOUND_FLAG, "true");
        }
    });
}

export function bindNavigateToInDescendants(element: HTMLElement) {
    const navigationElements = DomUtils.queryAllInDescendants(element, `[${ NAVIGATE_TO_ATTRIBUTE }]`);
    bindNavigateToElements(navigationElements);
}