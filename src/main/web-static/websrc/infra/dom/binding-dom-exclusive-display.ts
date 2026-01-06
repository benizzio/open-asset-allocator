import DomUtils from "./dom-utils";
import { logger, LogLevel } from "../logging";

const EXCLUSIVE_DISPLAY_ATTRIBUTE = "data-exclusive-display-container";
const EXCLUSIVE_DISPLAY_BOUND_FLAG = "data-exclusive-display-bound";

export function bindExclusiveDisplayContainerInDescendants(element: HTMLElement) {
    const exclusiveDisplayElements = DomUtils.queryAllInDescendants(element, `[${ EXCLUSIVE_DISPLAY_ATTRIBUTE }]`);
    bindExclusiveDisplayInDescendants(exclusiveDisplayElements);
}

function bindExclusiveDisplayInDescendants(exclusiveDisplayElements: NodeListOf<HTMLElement>) {

    exclusiveDisplayElements.forEach((element) => {

        if(!element.getAttribute(EXCLUSIVE_DISPLAY_BOUND_FLAG)) {

            element.setAttribute(EXCLUSIVE_DISPLAY_BOUND_FLAG, "binding");

            try {

                logger(LogLevel.INFO, "Binding exclusive display for element", element);

                const exclusiveDisplayElements = DomUtils.queryDirectDescendants(
                    element,
                    "*:not(script):not(style):not(link)",
                );
                bindExclusiveDisplay(exclusiveDisplayElements);
                element.setAttribute(EXCLUSIVE_DISPLAY_BOUND_FLAG, "true");
            } catch(error) {
                element.removeAttribute(EXCLUSIVE_DISPLAY_BOUND_FLAG);
                throw error;
            }
        }
    });
}

function bindExclusiveDisplay(exclusiveDisplayElements: NodeListOf<HTMLElement>) {
    exclusiveDisplayElements.forEach((element) => {
        addDisplayObserver(element, exclusiveDisplayElements);
    });
}

function addDisplayObserver(element: HTMLElement, exclusiveDisplayElements: NodeListOf<HTMLElement>) {

    const observer = new MutationObserver(() => {

        const display = element.style.display;

        if(display !== "none") {
            hideAllSiblings(element, exclusiveDisplayElements);
        }
    });

    observer.observe(element, { attributes: true, attributeFilter: ["style"] });
}

function hideAllSiblings(mutatedElement: HTMLElement, exclusiveDisplayElements: NodeListOf<HTMLElement>) {
    exclusiveDisplayElements.forEach((element) => {
        if(element !== mutatedElement) {
            element.style.display = "none";
        }
    });
}
