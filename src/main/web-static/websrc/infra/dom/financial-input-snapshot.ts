/**
 * Captures and restores the two independent values of a bound financial input.
 * Restoration bypasses the raw input's formatting observer so incomplete display
 * text and the exact submitted value survive draft reconstruction unchanged.
 * @module infra/dom/financial-input-snapshot
 * @author GPT-6.1 Terra
 */

/** The browser-exposed display and submitted raw values of a financial input.
 * @author GPT-6.1 Terra
 */
export type FinancialInputSnapshot = {
    display: string;
    raw: string;
};

/** Finds the hidden raw input created for a bound financial input.
 * @author GPT-6.1 Terra
 */
function rawInputFor(display: HTMLInputElement): HTMLInputElement {
    const container = display.parentElement;
    const raw = container?.children[1];

    if(!container?.hasAttribute("data-financial-input-container")
        || container.children[0] !== display
        || !(raw instanceof HTMLInputElement)
        || raw.type !== "hidden") {
        throw new Error("Financial input is not bound.");
    }

    return raw;
}

/** Captures a bound financial input without parsing or changing either value.
 *
 * @param display - The visible input after financial binding.
 * @returns Independent display and raw strings.
 * @example `const snapshot = captureFinancialInput(displayInput);`
 * @author GPT-6.1 Terra
 */
export function captureFinancialInput(display: HTMLInputElement): FinancialInputSnapshot {
    return { display: display.value, raw: rawInputFor(display).value };
}

/** Restores a bound financial input without formatting, recalculation, or input events.
 *
 * @param display - The visible input after financial binding.
 * @param snapshot - Previously captured display and raw strings.
 * @example `restoreFinancialInput(newDisplayInput, snapshot);`
 * @author GPT-6.1 Terra
 */
export function restoreFinancialInput(display: HTMLInputElement, snapshot: FinancialInputSnapshot): void {
    const raw = rawInputFor(display);
    const nativeSetter = Object.getOwnPropertyDescriptor(HTMLInputElement.prototype, "value")?.set;

    if(!nativeSetter) {
        throw new Error("Native input value setter is unavailable.");
    }

    nativeSetter.call(raw, snapshot.raw);
    display.value = snapshot.display;
}
