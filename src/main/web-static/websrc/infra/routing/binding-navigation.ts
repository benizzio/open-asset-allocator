// =============================================================================
// NAVIGATE TO
// triggers navigation to a route on click and enter key press
// =============================================================================

import { navigoRouter } from "./routing-navigo";
import { logger, LogLevel } from "../logging";
import DomUtils from "../dom-utils";

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

export function bindNavigateToInDescendants(element: HTMLElement) {
    const navigationElements = DomUtils.queryAllInDescendants(element, `[${NAVIGATE_TO_ATTRIBUTE}]`);
    bindNavigateToElements(navigationElements);
}