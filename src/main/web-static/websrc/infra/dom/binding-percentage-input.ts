import DomUtils from "./dom-utils";
import { logger, LogLevel } from "../logging";
import BigNumber from "bignumber.js";

// =============================================================================
// PERCENTAGE INPUT BINDING
// Automatically creates hidden decimal fields for percentage inputs
// Visible input shows percentage (0-100), hidden input stores decimal (0-1)
// Attribute value format: data-percentage-input="fieldName"
// =============================================================================

const PERCENTAGE_INPUT_ATTRIBUTE = "data-percentage-input";
const PERCENTAGE_INPUT_BOUND_FLAG = "data-percentage-input-bound";
const PERCENTAGE_INPUT_CONTAINER_ATTRIBUTE = "data-percentage-input-container";
const PERCENTAGE_INPUT_NULL_IF_EMPTY_ATTRIBUTE = "data-null-if-empty";

/**
 * Binds percentage input elements in descendants by creating hidden decimal fields
 * and setting up automatic synchronization between display and storage values.
 *
 * @param element The root element to search for percentage input fields
 *
 * @author GitHub Copilot
 */
export function bindPercentageInputsInDescendants(element: HTMLElement): void {

    const percentageInputs = DomUtils.queryAllInDescendants(
        element,
        `[${ PERCENTAGE_INPUT_ATTRIBUTE }]:not([${ PERCENTAGE_INPUT_BOUND_FLAG }])`,
    );

    bindPercentageInputElements(percentageInputs);
}

/**
 * Binds multiple percentage input elements.
 *
 * @param percentageInputs NodeList of input elements to bind
 *
 * @author GitHub Copilot
 */
function bindPercentageInputElements(percentageInputs: NodeListOf<HTMLElement>): void {

    percentageInputs.forEach((inputElement: HTMLInputElement) => {

        logger(LogLevel.INFO, "Binding percentage input for element", inputElement);

        inputElement.setAttribute(PERCENTAGE_INPUT_BOUND_FLAG, "binding");

        try {
            const fieldName = inputElement.getAttribute(PERCENTAGE_INPUT_ATTRIBUTE);

            if(!fieldName) {
                logger(LogLevel.WARN, "Percentage input element missing field name", inputElement);
                inputElement.removeAttribute(PERCENTAGE_INPUT_BOUND_FLAG);
                return;
            }

            bindPercentageInput(inputElement, fieldName);
            inputElement.setAttribute(PERCENTAGE_INPUT_BOUND_FLAG, "true");
        } catch(error) {
            inputElement.removeAttribute(PERCENTAGE_INPUT_BOUND_FLAG);
            throw error;
        }
    });
}

/**
 * Binds a single percentage input element by creating the hidden decimal field
 * and setting up the synchronization logic. Wraps both inputs in a container div
 * to allow scoped field referencing by name.
 *
 * @param percentageInput The visible percentage input element
 * @param fieldName The form field name for the hidden decimal input
 *
 * @author GitHub Copilot
 */
function bindPercentageInput(percentageInput: HTMLInputElement, fieldName: string): void {

    // Create container div to wrap both inputs
    const container = document.createElement("div");
    container.setAttribute(PERCENTAGE_INPUT_CONTAINER_ATTRIBUTE, fieldName);

    // Wrap the percentage input in the container
    percentageInput.parentNode.insertBefore(container, percentageInput);
    container.appendChild(percentageInput);


    // Create hidden decimal field
    const hiddenDecimalInput = createHiddenDecimalField(percentageInput, fieldName);

    // Insert hidden field inside the container
    container.appendChild(hiddenDecimalInput);

    // Initialize display value from decimal if present
    initializePercentageDisplay(percentageInput, hiddenDecimalInput);

    // Set up synchronization on input
    percentageInput.addEventListener("input", () => {
        syncPercentageToDecimalInContainer(container, fieldName);
    });

    // Configure percentage input attributes
    configurePercentageInputAttributes(percentageInput);
}

/**
 * Creates the hidden input field that stores the decimal value (0-1).
 *
 * @param percentageInput The visible percentage input element
 * @param fieldName The form field name for the hidden input
 * @returns The created hidden input element
 *
 * @author GitHub Copilot
 */
function createHiddenDecimalField(percentageInput: HTMLInputElement, fieldName: string): HTMLInputElement {

    const hiddenInput = document.createElement("input");
    hiddenInput.type = "hidden";
    hiddenInput.name = fieldName;

    // Transfer data-null-if-empty attribute if present
    if(percentageInput.hasAttribute(PERCENTAGE_INPUT_NULL_IF_EMPTY_ATTRIBUTE)) {
        hiddenInput.setAttribute(PERCENTAGE_INPUT_NULL_IF_EMPTY_ATTRIBUTE, "");
    }

    return hiddenInput;
}


/**
 * Configures the percentage input attributes for proper display behavior.
 *
 * @param percentageInput The visible percentage input element
 *
 * @author GitHub Copilot
 */
function configurePercentageInputAttributes(percentageInput: HTMLInputElement): void {

    // Ensure input type is number
    if(percentageInput.type !== "number") {
        percentageInput.type = "number";
    }

    // Set min/max if not already set
    if(!percentageInput.hasAttribute("min")) {
        percentageInput.min = "0";
    }

    if(!percentageInput.hasAttribute("max")) {
        percentageInput.max = "100";
    }

    // Set step if not already set
    if(!percentageInput.hasAttribute("step")) {
        percentageInput.step = "any";
    }

    // Remove name attribute so only hidden field is submitted
    percentageInput.removeAttribute("name");
}

/**
 * Initializes the percentage display by converting existing decimal value to percentage.
 * Called once during binding to properly display server-provided decimal values.
 * Uses BigNumber for precise decimal arithmetic.
 *
 * @param percentageInput The visible percentage input element
 * @param decimalInput The hidden decimal input element
 *
 * @author GitHub Copilot
 */
function initializePercentageDisplay(percentageInput: HTMLInputElement, decimalInput: HTMLInputElement): void {

    // Check if percentage input already has a value (from template)
    const existingDisplayValue = percentageInput.value;

    if(existingDisplayValue && existingDisplayValue !== "") {

        try {
            const numericValue = new BigNumber(existingDisplayValue);

            if(!numericValue.isNaN()) {
                // Value is in percentage format, convert to decimal for hidden field
                // percentage / 100 = decimal
                const decimalValue = numericValue.dividedBy(100);
                decimalInput.value = decimalValue.toString();
                return;
            }
        } catch(error) {
            logger(LogLevel.WARN, "Invalid percentage value during initialization", existingDisplayValue, error);
        }
    }

    // No display value, check if there's a decimal value to initialize from
    const decimalValue = decimalInput.value;

    if(!decimalValue || decimalValue === "") {
        return;
    }

    try {
        const numericDecimal = new BigNumber(decimalValue);

        if(numericDecimal.isNaN()) {
            return;
        }

        // Convert decimal (0-1) to percentage (0-100) for display
        // decimal * 100 = percentage
        const percentageValue = numericDecimal.multipliedBy(100);
        percentageInput.value = percentageValue.toString();
    } catch(error) {
        logger(LogLevel.WARN, "Invalid decimal value during initialization", decimalValue, error);
    }
}

/**
 * Synchronizes the visible percentage input (0-100) with the hidden decimal input (0-1)
 * by accessing fields by position within the container.
 * Called on every input event to keep values in sync.
 *
 * The container structure is always:
 * - First child: percentage input (visible)
 * - Second child: decimal input (hidden)
 *
 * @param container The container div wrapping both inputs
 * @param fieldName The form field name (used only for logging)
 *
 * @author GitHub Copilot
 */
function syncPercentageToDecimalInContainer(container: HTMLElement, fieldName: string): void {

    // Access inputs by position: first child is percentage, second is hidden decimal
    const percentageInput = container.children[0] as HTMLInputElement;
    const decimalInput = container.children[1] as HTMLInputElement;

    if(!percentageInput || !decimalInput) {
        logger(LogLevel.WARN, "Container missing expected child inputs", container, fieldName);
        return;
    }

    syncPercentageToDecimal(percentageInput, decimalInput);
}

/**
 * Synchronizes the visible percentage input (0-100) with the hidden decimal input (0-1).
 * Called on every input event to keep values in sync.
 * Uses BigNumber for precise decimal arithmetic.
 * Limits percentage to 3 decimal places to ensure decimal field has maximum 5 decimal places.
 *
 * @param percentageInput The visible percentage input element
 * @param decimalInput The hidden decimal input element
 *
 * @author GitHub Copilot
 */
function syncPercentageToDecimal(percentageInput: HTMLInputElement, decimalInput: HTMLInputElement): void {

    const percentageValue = percentageInput.value;

    if(!percentageValue || percentageValue === "") {
        decimalInput.value = "";
        return;
    }

    try {
        let numericPercentage = new BigNumber(percentageValue);

        if(numericPercentage.isNaN()) {
            decimalInput.value = "";
            return;
        }

        // Limit percentage to 3 decimal places
        // This ensures the decimal value will have at most 5 decimal places (3 + 2 from /100)
        numericPercentage = numericPercentage.decimalPlaces(3, BigNumber.ROUND_DOWN);

        // Update the input field if value was truncated
        if(percentageInput.value !== numericPercentage.toString()) {
            percentageInput.value = numericPercentage.toString();
        }

        // Cap values at 0-100 range
        let cappedPercentage = numericPercentage;

        if(numericPercentage.isGreaterThan(100)) {
            cappedPercentage = new BigNumber(100);
            percentageInput.value = "100";
        }
        else if(numericPercentage.isLessThan(0)) {
            cappedPercentage = new BigNumber(0);
            percentageInput.value = "0";
        }

        // Convert percentage (0-100) to decimal (0-1)
        // percentage / 100 = decimal
        // Result will have at most 5 decimal places
        const decimalValue = cappedPercentage.dividedBy(100);
        decimalInput.value = decimalValue.toString();
    } catch(error) {
        logger(LogLevel.WARN, "Invalid percentage value during synchronization", percentageValue, error);
        decimalInput.value = "";
    }
}
