import { buildParameterizedDestinationPathFromCurrentLocationContext, navigoRouter } from "./routing-navigo";
import { logger, LogLevel } from "../logging";

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
    return buildParameterizedDestinationPathFromCurrentLocationContext(destinationPath);
}

function bindNavigateToElements(navigationElement: NodeListOf<HTMLElement>) {

    navigationElement.forEach((element) => {

        element.setAttribute(NAVIGATE_TO_BOUND_FLAG, "binding");

        try {
            logger(LogLevel.INFO, "Binding navigation element", element);
            bindKeypressNavigation(element);
            bindCLickNavigation(element);
            element.setAttribute(NAVIGATE_TO_BOUND_FLAG, "true");
        } catch(error) {
            element.removeAttribute(NAVIGATE_TO_BOUND_FLAG);
            throw error;
        }
    });
}

export function bindNavigateToInDescendants(element: HTMLElement) {
    const navigationElements = element.querySelectorAll(
        `[${ NAVIGATE_TO_ATTRIBUTE }]:not([${ NAVIGATE_TO_BOUND_FLAG }])`,
    ) as NodeListOf<HTMLElement>;
    bindNavigateToElements(navigationElements);
}