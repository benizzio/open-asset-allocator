import { buildParameterizedDestinationPath, navigoRouter } from "./routing-navigo";
import { logger, LogLevel } from "../logging";
import DomUtils from "../dom/dom-utils";

// =============================================================================
// NAVIGATE TO
// triggers navigation to a route on click and enter key press
// =============================================================================

const NAVIGATE_TO_ATTRIBUTE = "data-navigate-to";
const NAVIGATE_TO_BOUND_FLAG = "navigate-to-bound";

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
    const destinationPath = element.getAttribute(NAVIGATE_TO_ATTRIBUTE);
    return buildParameterizedDestinationPath(destinationPath);
}

function bindNavigateToElements(navigationElement: NodeListOf<HTMLElement>) {
    navigationElement.forEach((element) => {
        if(!element.getAttribute(NAVIGATE_TO_BOUND_FLAG)) {//TODO replace this if for improving the element selector
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