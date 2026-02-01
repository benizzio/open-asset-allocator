import { logger, LogLevel } from "../logging";
import Format from "../format";
import { parseNonNegativeInt } from "../../utils/lang";

// =============================================================================
// FINANCIAL INPUT BINDING
// Automatically creates hidden raw value fields for financial inputs
// Visible input shows locale-formatted number, hidden input stores raw decimal
// Attribute value format: data-financial-input="fieldName"
// =============================================================================

const FINANCIAL_INPUT_ATTRIBUTE = "data-financial-input";
const FINANCIAL_INPUT_BOUND_FLAG = "data-financial-input-bound";
const FINANCIAL_INPUT_CONTAINER_ATTRIBUTE = "data-financial-input-container";
const FINANCIAL_INPUT_NULL_IF_EMPTY_ATTRIBUTE = "data-null-if-empty";
const FINANCIAL_INPUT_DECIMALS_ATTRIBUTE = "data-financial-input-decimals";
const FINANCIAL_INPUT_CONTAINER_DECIMALS_ATTRIBUTE = "data-financial-input-container-decimals";
const DEFAULT_DECIMAL_PLACES = 2;

/**
 * Binds financial input elements in descendants by creating hidden raw value fields
 * and setting up automatic synchronization between display and storage values.
 *
 * @param element The root element to search for financial input fields
 *
 * @author GitHub Copilot
 */
export function bindFinancialInputsInDescendants(element: HTMLElement): void {

    const financialInputs = element.querySelectorAll(
        `[${ FINANCIAL_INPUT_ATTRIBUTE }]:not([${ FINANCIAL_INPUT_BOUND_FLAG }])`,
    ) as NodeListOf<HTMLElement>;

    bindFinancialInputElements(financialInputs);
}

/**
 * Binds multiple financial input elements.
 *
 * @param financialInputs NodeList of input elements to bind
 *
 * @author GitHub Copilot
 */
function bindFinancialInputElements(financialInputs: NodeListOf<HTMLElement>): void {

    financialInputs.forEach((inputElement: HTMLInputElement) => {

        logger(LogLevel.INFO, "Binding financial input for element", inputElement);

        inputElement.setAttribute(FINANCIAL_INPUT_BOUND_FLAG, "binding");

        try {
            const fieldName = inputElement.getAttribute(FINANCIAL_INPUT_ATTRIBUTE);

            if(!fieldName) {
                logger(LogLevel.WARN, "Financial input element missing field name", inputElement);
                inputElement.removeAttribute(FINANCIAL_INPUT_BOUND_FLAG);
                return;
            }

            bindFinancialInput(inputElement, fieldName);
            inputElement.setAttribute(FINANCIAL_INPUT_BOUND_FLAG, "true");
        } catch(error) {
            inputElement.removeAttribute(FINANCIAL_INPUT_BOUND_FLAG);
            throw error;
        }
    });
}

/**
 * Binds a single financial input element by creating the hidden raw value field
 * and setting up the synchronization logic. Wraps both inputs in a container div
 * to allow scoped field referencing by name.
 *
 * @param financialInput The visible financial input element
 * @param fieldName The form field name for the hidden raw value input
 *
 * @author GitHub Copilot
 */
function bindFinancialInput(financialInput: HTMLInputElement, fieldName: string): void {

    // Create container div to wrap both inputs
    const container = document.createElement("div");
    container.setAttribute(FINANCIAL_INPUT_CONTAINER_ATTRIBUTE, fieldName);

    // Read and store custom decimal places if specified
    const decimalPlaces = getDecimalPlacesFromInput(financialInput);
    container.setAttribute(FINANCIAL_INPUT_CONTAINER_DECIMALS_ATTRIBUTE, decimalPlaces.toString());

    // Wrap the financial input in the container
    financialInput.parentNode.insertBefore(container, financialInput);
    container.appendChild(financialInput);

    // Create hidden raw value field (transfers oninput attribute)
    const hiddenRawInput = createHiddenRawValueField(financialInput, fieldName);

    // Insert hidden field inside the container
    container.appendChild(hiddenRawInput);

    // Initialize display value from raw if present
    initializeFinancialDisplay(financialInput, hiddenRawInput, decimalPlaces);

    // Set up synchronization on blur (when user finishes editing)
    financialInput.addEventListener("blur", () => {
        syncDisplayToRawInContainer(container, fieldName);
    });

    // Set up synchronization on input for the raw value (without formatting display)
    // After syncing, trigger the oninput event on the hidden field so handlers receive raw value
    financialInput.addEventListener("input", () => {
        syncInputToRawInContainer(container, fieldName);
        triggerInputEventOnHiddenField(container);
    });

    // Set up observer to detect programmatic changes in hidden raw field
    setupRawValueObserver(hiddenRawInput, financialInput, decimalPlaces);

    // Configure financial input attributes (removes oninput from visible input)
    configureFinancialInputAttributes(financialInput);
}

/**
 * Gets the decimal places value from the input element's data attribute.
 * Returns the default value if the attribute is not present or invalid.
 *
 * @param financialInput The financial input element
 * @returns The number of decimal places to use
 *
 * @author GitHub Copilot
 */
function getDecimalPlacesFromInput(financialInput: HTMLInputElement): number {

    const decimalsAttr = financialInput.getAttribute(FINANCIAL_INPUT_DECIMALS_ATTRIBUTE);

    return parseNonNegativeInt(decimalsAttr, DEFAULT_DECIMAL_PLACES);
}

/**
 * Gets the decimal places value from the container element's data attribute.
 * Returns the default value if the attribute is not present or invalid.
 *
 * @param container The container element
 * @returns The number of decimal places to use
 *
 * @author GitHub Copilot
 */
function getDecimalPlacesFromContainer(container: HTMLElement): number {

    const decimalsAttr = container.getAttribute(FINANCIAL_INPUT_CONTAINER_DECIMALS_ATTRIBUTE);

    return parseNonNegativeInt(decimalsAttr, DEFAULT_DECIMAL_PLACES);
}

/**
 * Creates the hidden input field that stores the raw numeric value.
 * Transfers the oninput attribute from the visible input so handlers receive raw values.
 *
 * @param financialInput The visible financial input element
 * @param fieldName The form field name for the hidden input
 * @returns The created hidden input element
 *
 * @author GitHub Copilot
 */
function createHiddenRawValueField(financialInput: HTMLInputElement, fieldName: string): HTMLInputElement {

    const hiddenInput = document.createElement("input");
    hiddenInput.type = "hidden";
    hiddenInput.name = fieldName;

    // Transfer data-null-if-empty attribute if present
    if(financialInput.hasAttribute(FINANCIAL_INPUT_NULL_IF_EMPTY_ATTRIBUTE)) {
        hiddenInput.setAttribute(FINANCIAL_INPUT_NULL_IF_EMPTY_ATTRIBUTE, "");
    }

    // Transfer oninput attribute if present (so handlers receive raw numeric value)
    const onInputAttr = financialInput.getAttribute("oninput");

    if(onInputAttr) {
        hiddenInput.setAttribute("oninput", onInputAttr);
    }

    return hiddenInput;
}

/**
 * Configures the financial input attributes for proper display behavior.
 * Removes attributes that were transferred to the hidden input.
 *
 * @param financialInput The visible financial input element
 *
 * @author GitHub Copilot
 */
function configureFinancialInputAttributes(financialInput: HTMLInputElement): void {

    // Use text type for formatted display (number type doesn't allow formatting characters)
    financialInput.type = "text";

    // Remove name attribute so only hidden field is submitted
    financialInput.removeAttribute("name");

    // Remove oninput attribute (transferred to hidden field)
    financialInput.removeAttribute("oninput");

    // Set inputmode for numeric keyboard on mobile devices
    financialInput.setAttribute("inputmode", "decimal");
}

/**
 * Initializes the financial display by converting existing raw value to formatted display.
 * Called once during binding to properly display server-provided values.
 *
 * @param financialInput The visible financial input element
 * @param rawInput The hidden raw value input element
 * @param decimalPlaces The number of decimal places to use for formatting
 *
 * @author GitHub Copilot
 */
function initializeFinancialDisplay(
    financialInput: HTMLInputElement,
    rawInput: HTMLInputElement,
    decimalPlaces: number,
): void {

    // Check if financial input already has a value (from template)
    const existingDisplayValue = financialInput.value;

    if(existingDisplayValue && existingDisplayValue !== "") {

        const numericValue = parseFloat(existingDisplayValue);

        if(!isNaN(numericValue)) {
            // Value is a raw number, format it for display and store raw in hidden field
            rawInput.value = numericValue.toString();
            financialInput.value = Format.formatFinancialNumber(numericValue, undefined, decimalPlaces);
            return;
        }
    }

    // No display value, check if there's a raw value to initialize from
    const rawValue = rawInput.value;

    if(!rawValue || rawValue === "") {
        return;
    }

    const numericRaw = parseFloat(rawValue);

    if(isNaN(numericRaw)) {
        return;
    }

    // Format the raw value for display
    financialInput.value = Format.formatFinancialNumber(numericRaw, undefined, decimalPlaces);
}

/**
 * Synchronizes the visible formatted input with the hidden raw input on blur.
 * Re-formats the display value when user finishes editing.
 *
 * @param container The container div wrapping both inputs
 * @param fieldName The form field name (used only for logging)
 *
 * @author GitHub Copilot
 */
function syncDisplayToRawInContainer(container: HTMLElement, fieldName: string): void {

    // Access inputs by position: first child is financial, second is hidden raw
    const financialInput = container.children[0] as HTMLInputElement;
    const rawInput = container.children[1] as HTMLInputElement;

    if(!financialInput || !rawInput) {
        logger(LogLevel.WARN, "Container missing expected child inputs", container, fieldName);
        return;
    }

    const decimalPlaces = getDecimalPlacesFromContainer(container);
    syncDisplayToRaw(financialInput, rawInput, decimalPlaces);
}

/**
 * Synchronizes the visible input with the hidden raw input during typing.
 * Updates raw value without reformatting display (allows user to type freely).
 *
 * @param container The container div wrapping both inputs
 * @param fieldName The form field name (used only for logging)
 *
 * @author GitHub Copilot
 */
function syncInputToRawInContainer(container: HTMLElement, fieldName: string): void {

    // Access inputs by position: first child is financial, second is hidden raw
    const financialInput = container.children[0] as HTMLInputElement;
    const rawInput = container.children[1] as HTMLInputElement;

    if(!financialInput || !rawInput) {
        logger(LogLevel.WARN, "Container missing expected child inputs", container, fieldName);
        return;
    }

    syncInputToRaw(financialInput, rawInput);
}

/**
 * Synchronizes the visible formatted input with the hidden raw input.
 * Called on blur to reformat display and update raw value.
 * Re-parses the formatted value to ensure raw matches displayed (handles decimal truncation).
 *
 * @param financialInput The visible financial input element
 * @param rawInput The hidden raw value input element
 * @param decimalPlaces The number of decimal places to use for formatting
 *
 * @author GitHub Copilot
 */
function syncDisplayToRaw(
    financialInput: HTMLInputElement,
    rawInput: HTMLInputElement,
    decimalPlaces: number,
): void {

    const displayValue = financialInput.value;

    if(!displayValue || displayValue.trim() === "") {
        rawInput.value = "";
        financialInput.value = "";
        return;
    }

    // Parse the formatted value to get raw number
    const numericValue = Format.parseFinancialNumber(displayValue);

    if(isNaN(numericValue)) {
        rawInput.value = "";
        return;
    }

    // Ensure non-negative values
    const finalValue = Math.max(0, numericValue);

    // Reformat display (may round/truncate decimal places)
    const formattedValue = Format.formatFinancialNumber(finalValue, undefined, decimalPlaces);
    financialInput.value = formattedValue;

    // Re-parse the formatted value to ensure raw value matches displayed value
    // This handles cases where formatting truncates decimal places
    const correctedRawValue = Format.parseFinancialNumber(formattedValue);
    rawInput.value = correctedRawValue.toString();
}

/**
 * Synchronizes the visible input with the hidden raw value during typing.
 * Parses input to update raw value without reformatting display.
 *
 * @param financialInput The visible financial input element
 * @param rawInput The hidden raw value input element
 *
 * @author GitHub Copilot
 */
function syncInputToRaw(financialInput: HTMLInputElement, rawInput: HTMLInputElement): void {

    const displayValue = financialInput.value;

    if(!displayValue || displayValue.trim() === "") {
        rawInput.value = "";
        return;
    }

    // Parse the value (may be partially formatted or raw during typing)
    const numericValue = Format.parseFinancialNumber(displayValue);

    if(isNaN(numericValue)) {
        // During typing, user may enter invalid intermediate values
        // Keep raw value empty until valid
        rawInput.value = "";
        return;
    }

    // Update raw value
    rawInput.value = Math.max(0, numericValue).toString();
}

/**
 * Triggers an input event on the hidden raw value field.
 * This allows oninput handlers defined on the original visible input to be executed
 * with access to the raw numeric value.
 *
 * @param container The container div wrapping both inputs
 *
 * @author GitHub Copilot
 */
function triggerInputEventOnHiddenField(container: HTMLElement): void {

    const hiddenInput = container.children[1] as HTMLInputElement;

    if(!hiddenInput) {
        return;
    }

    // Dispatch an input event on the hidden field to trigger any oninput handlers
    const inputEvent = new Event("input", { bubbles: true, cancelable: true });
    hiddenInput.dispatchEvent(inputEvent);
}

/**
 * Sets up an observer to detect programmatic changes in the hidden raw value field
 * and reflect them in the visible display field. Uses property descriptor override
 * to intercept value setter since MutationObserver doesn't detect .value changes.
 *
 * @param rawInput The hidden raw value input element
 * @param financialInput The visible financial input element to update
 * @param decimalPlaces The number of decimal places to use for formatting
 *
 * @author GitHub Copilot
 */
function setupRawValueObserver(
    rawInput: HTMLInputElement,
    financialInput: HTMLInputElement,
    decimalPlaces: number,
): void {

    // Get the original value descriptor from HTMLInputElement prototype
    const originalDescriptor = Object.getOwnPropertyDescriptor(HTMLInputElement.prototype, "value");

    if(!originalDescriptor) {
        logger(LogLevel.WARN, "Could not get value descriptor for raw value observer");
        return;
    }

    // Track whether we're in a sync operation to prevent infinite loops
    let syncing = false;

    // Override the value property to intercept programmatic changes
    Object.defineProperty(rawInput, "value", {
        get: function() {
            return originalDescriptor.get.call(this);
        },

        set: function(newValue: string) {

            // Call original setter first
            originalDescriptor.set.call(this, newValue);

            // Skip if we're already syncing (prevents infinite loops)
            if(syncing) {
                return;
            }

            // Skip if financial input is focused (user is typing)
            if(document.activeElement === financialInput) {
                return;
            }

            // Sync the display value from the new raw value
            syncing = true;
            syncRawToDisplay(rawInput, financialInput, decimalPlaces);
            syncing = false;
        },

        configurable: true,
        enumerable: true,
    });
}

/**
 * Synchronizes the hidden raw value to the visible formatted display.
 * Called when the raw value is programmatically changed.
 *
 * @param rawInput The hidden raw value input element
 * @param financialInput The visible financial input element to update
 * @param decimalPlaces The number of decimal places to use for formatting
 *
 * @author GitHub Copilot
 */
function syncRawToDisplay(
    rawInput: HTMLInputElement,
    financialInput: HTMLInputElement,
    decimalPlaces: number,
): void {

    const rawValue = rawInput.value;

    if(!rawValue || rawValue.trim() === "") {
        financialInput.value = "";
        return;
    }

    const numericValue = parseFloat(rawValue);

    if(isNaN(numericValue)) {
        financialInput.value = "";
        return;
    }

    // Format the raw value for display
    financialInput.value = Format.formatFinancialNumber(numericValue, undefined, decimalPlaces);
}
