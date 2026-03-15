import DomUtils from "../dom/dom-utils";
import { logger, LogLevel } from "../logging";

// =============================================================================
// HTMX WAIT FOR READY
// prevents HTMX triggers from firing until all required conditions are met
// attribute value format is "condition1,condition2,..."
// conditions are fulfilled externally via addReadyConditionToWaitingElement(selector, condition)
// =============================================================================

const HTMX_WAIT_FOR_READY_ATTRIBUTE = "data-hx-wait-for-ready";
const HTMX_WAIT_FOR_READY_FLAG = "data-hx-wait-for-ready-fulfilled";
const HTMX_WAIT_FOR_READY_BOUND_FLAG = "data-hx-wait-for-ready-bound";

export function bindHTMXWaitForReadyInDescendants(element: HTMLElement) {

    const elementsToBind = element.querySelectorAll(
        `[${ HTMX_WAIT_FOR_READY_ATTRIBUTE }]:not([${ HTMX_WAIT_FOR_READY_BOUND_FLAG }])`,
    ) as NodeListOf<HTMLElement>;

    bindWaitForReadyOnElements(elementsToBind);
}

function bindWaitForReadyOnElements(elements: NodeListOf<HTMLElement>) {

    elements.forEach((element) => {

        element.setAttribute(HTMX_WAIT_FOR_READY_BOUND_FLAG, "binding");

        try {

            logger(LogLevel.INFO, "Binding HTMX wait for ready on element", element);

            bindConfirmGateOnElement(element);

            element.setAttribute(HTMX_WAIT_FOR_READY_BOUND_FLAG, "true");
        } catch(error) {
            element.removeAttribute(HTMX_WAIT_FOR_READY_BOUND_FLAG);
            throw error;
        }
    });
}

/**
 * Parses a comma-separated string value into a set of trimmed, non-empty values.
 *
 * @author GitHub Copilot
 */
function parseCommaSeparatedSet(value: string): Set<string> {

    return new Set(
        value.split(",").map(c => c.trim()).filter(Boolean),
    );
}

/**
 * Parses the wait-for-ready attribute value into a set of required conditions.
 *
 * @author GitHub Copilot
 */
function parseRequiredConditions(element: HTMLElement): Set<string> {

    const attributeValue = element.getAttribute(HTMX_WAIT_FOR_READY_ATTRIBUTE);

    if(!attributeValue) {
        return new Set();
    }

    return parseCommaSeparatedSet(attributeValue);
}

/**
 * Checks whether all required conditions are present in the element's fulfilled flag.
 *
 * @author GitHub Copilot
 */
function areAllConditionsReady(element: HTMLElement, requiredConditions: Set<string>): boolean {

    const readyValue = element.getAttribute(HTMX_WAIT_FOR_READY_FLAG) || "";
    const fulfilledConditions = parseCommaSeparatedSet(readyValue);

    return requiredConditions.difference(fulfilledConditions).size === 0;
}

/**
 * Binds an htmx:confirm gate on the element that prevents HTMX triggers from firing
 * until all required conditions are met.
 *
 * When conditions are not yet fulfilled, the event is prevented and the issueRequest
 * function is stored. A MutationObserver watches the fulfilled flag attribute and calls
 * issueRequest() when all conditions are met, allowing the original HTMX request to proceed.
 *
 * @author GitHub Copilot
 */
function bindConfirmGateOnElement(element: HTMLElement) {

    let pendingIssueRequest: (() => void) | null = null;

    element.addEventListener("htmx:confirm", (event: CustomEvent) => {

        if(event.target !== element) {
            return;
        }

        const requiredConditions = parseRequiredConditions(element);

        if(requiredConditions.size === 0) {
            return;
        }

        if(areAllConditionsReady(element, requiredConditions)) {
            return;
        }

        event.preventDefault();
        pendingIssueRequest = event.detail.issueRequest;

        logger(
            LogLevel.DEBUG,
            "HTMX trigger blocked, waiting for ready conditions",
            element,
            requiredConditions,
        );
    });

    addReadyFlagObserverOnElement(element, () => {

        if(pendingIssueRequest) {

            logger(
                LogLevel.DEBUG,
                "All ready conditions met, issuing pending HTMX request on element",
                element,
            );

            const issueRequest = pendingIssueRequest;
            pendingIssueRequest = null;
            issueRequest();
        }
    });
}

/**
 * Adds a MutationObserver on the element's fulfilled flag attribute. When all conditions
 * are met, calls the provided callback to issue the pending HTMX request.
 *
 * Also observes the document for element removal to clean up the observer if needed.
 *
 * @author GitHub Copilot
 */
function addReadyFlagObserverOnElement(element: HTMLElement, onReady: () => void) {

    const readyObserver = new MutationObserver(() => {

        const requiredConditions = parseRequiredConditions(element);
        const hasConditions = requiredConditions.size > 0;
        const allReady = hasConditions && areAllConditionsReady(element, requiredConditions);

        if(allReady) {
            readyObserver.disconnect();
            removalObserver.disconnect();
            onReady();
        }
    });

    readyObserver.observe(element, {
        attributes: true,
        attributeFilter: [HTMX_WAIT_FOR_READY_FLAG],
    });

    const removalObserver = new MutationObserver(() => {

        if(DomUtils.wasElementRemoved(element)) {
            logger(LogLevel.INFO, "Waiting element removed, cleaning up ready observer", element);
            readyObserver.disconnect();
            removalObserver.disconnect();
        }
    });

    removalObserver.observe(document, { childList: true, subtree: true });
}

/**
 * Adds a condition value to the fulfilled flag of all elements matching the selector.
 *
 * Appends the condition to the comma-separated fulfilled flag list if not already present.
 * The MutationObserver on waiting elements will detect this change and check
 * whether all required conditions are now fulfilled.
 *
 * @param selector - CSS selector to find the elements to add the condition to.
 * @param condition - The condition value to add.
 *
 * @example
 * addReadyConditionToWaitingElement("[data-hx-wait-for-ready]", "settled");
 * addReadyConditionToWaitingElement("[data-hx-wait-for-ready]", "partials-registered");
 *
 * @author GitHub Copilot
 */
export function addReadyConditionToWaitingElement(selector: string, condition: string) {

    const elements = document.querySelectorAll<HTMLElement>(selector);

    if(elements.length === 0) {
        logger(LogLevel.WARN, "addReadyConditionToWaitingElement found no elements for selector", selector, condition);
        return;
    }

    elements.forEach(element => {

        const currentValue = element.getAttribute(HTMX_WAIT_FOR_READY_FLAG) || "";
        const currentConditions = parseCommaSeparatedSet(currentValue);

        if(!currentConditions.has(condition)) {
            currentConditions.add(condition);
            element.setAttribute(HTMX_WAIT_FOR_READY_FLAG, [...currentConditions].join(","));
        }
    });
}
