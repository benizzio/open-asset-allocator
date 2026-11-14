/**
 * Preserves the independently editable display and submitted fraction of a bound
 * percentage input without invoking its conversion or validation handlers.
 * @module infra/dom/percentage-input-snapshot
 * @author GPT-6.1 Terra
 */

/** Browser-exposed values of a bound percentage input.
 * @author GPT-6.1 Terra
 */
export type PercentageInputSnapshot = {
    display: string;
    fraction: string;
};

/** Finds and validates the hidden fraction associated with a visible input.
 * @author GPT-6.1 Terra
 */
function fractionInputFor(display: HTMLInputElement): HTMLInputElement {
    const container = display.parentElement;
    const fraction = container?.children[1];

    if(!container?.hasAttribute("data-percentage-input-container")
        || container.children[0] !== display
        || !(fraction instanceof HTMLInputElement)
        || fraction.type !== "hidden") {
        throw new Error("Percentage input is not bound.");
    }

    if(fraction.name !== display.getAttribute("data-percentage-input")
        || container.getAttribute("data-percentage-input-container") !== fraction.name) {
        throw new Error("Percentage input is not bound.");
    }

    return fraction;
}

/** Captures both values without parsing or changing the input.
 *
 * @param display - Visible percentage input after binding.
 * @returns Independent display and submitted fraction strings.
 * @example `const snapshot = capturePercentageInput(percentageInput);`
 * @author GPT-6.1 Terra
 */
export function capturePercentageInput(display: HTMLInputElement): PercentageInputSnapshot {
    return { display: display.value, fraction: fractionInputFor(display).value };
}

/** Restores both values without conversion, rounding, clamping, or input events.
 *
 * @param display - Fresh visible percentage input after binding.
 * @param snapshot - Previously captured display and fraction strings.
 * @example `restorePercentageInput(newPercentageInput, snapshot);`
 * @author GPT-6.1 Terra
 */
export function restorePercentageInput(display: HTMLInputElement, snapshot: PercentageInputSnapshot): void {
    const fraction = fractionInputFor(display);
    display.value = snapshot.display;
    fraction.value = snapshot.fraction;
}
