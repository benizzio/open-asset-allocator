/**
 * Rejects invalid user edits before observation numeric inputs update their dependent fields.
 * @module infra/dom/numeric-edit-guard
 * @author GPT-6.1 Terra
 */

const guardedInputs = new WeakSet<HTMLInputElement>();

/** Identifies editable non-negative financial values without enforcing grouping positions.
 * @author GPT-6.1 Terra
 */
function isFinancialDraft(value: string): boolean {
    return /^\d*(?:,\d*)*(?:\.\d*)?$/.test(value);
}

/** Rejects incomplete grouping in pasted financial values.
 * @author GPT-6.1 Terra
 */
function isFinancialPaste(value: string): boolean {
    return value.includes(",")
        ? /^\d{1,3}(?:,\d{3})+(?:\.\d*)?$/.test(value)
        : /^\d*(?:\.\d*)?$/.test(value);
}

/** Builds the value resulting from a text input's current selection and inserted text.
 * @author GPT-6.1 Terra
 */
function replacement(input: HTMLInputElement, text: string): string {
    const start = input.selectionStart ?? input.value.length;
    const end = input.selectionEnd ?? start;
    return input.value.slice(0, start) + text + input.value.slice(end);
}

/**
 * Guards one observation numeric input against invalid user edits without intercepting programmatic updates.
 * Financial text fields retain their submitted raw-value synchronization; native number inputs retain their spinbutton.
 * @param input - Visible numeric input.
 * @param financial - Whether the input accepts financial grouping separators.
 * @example `guardNumericEdits(priceInput, true);`
 * @author GPT-6.1 Terra
 */
export function guardNumericEdits(input: HTMLInputElement, financial: boolean): void {
    if(guardedInputs.has(input)) {
        return;
    }

    guardedInputs.add(input);
    let accepted = input.value;

    input.addEventListener("focus", () => {
        accepted = input.value;
    });

    input.addEventListener("beforeinput", event => {
        const edit = event as InputEvent;

        if(!edit.inputType.startsWith("insert") || edit.data === null) {
            return;
        }

        const valid = financial
            ? isFinancialDraft(replacement(input, edit.data))
            : /^[0-9.]*$/.test(edit.data);

        if(!valid && edit.cancelable) {
            edit.preventDefault();
        }
    });

    input.addEventListener("paste", event => {
        const text = event.clipboardData?.getData("text/plain");

        if(text === undefined) {
            return;
        }

        const valid = financial
            ? isFinancialPaste(replacement(input, text))
            : /^\d*(?:\.\d*)?$/.test(text);

        if(!valid) {
            event.preventDefault();
        }
    });

    input.addEventListener("keydown", event => {
        const modifier = event.ctrlKey || event.metaKey || event.altKey;

        if(!financial && !modifier && /^[a-z]$/i.test(event.key)) {
            event.preventDefault();
        }
    });

    input.addEventListener("input", event => {
        const valid = financial ? isFinancialDraft(input.value) : !input.validity.badInput;

        if(!valid) {
            input.value = accepted;
            event.stopImmediatePropagation();
            return;
        }

        accepted = input.value;
    }, true);
}
