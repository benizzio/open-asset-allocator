/**
 * Shared behavior for revealing persisted management metadata fields.
 *
 * Use these functions from a record-specific management controller when a user activates its edit control.
 * Example: `revealManagementMetadataEditor(fields, nameInput, dateInput)` reveals the inputs and focuses the name.
 * @author GPT-6 Luna
 */

import { Collapse } from "bootstrap";

/** Reveals and enables one record's metadata fields, then focuses its text input.
 *
 * @param fields - The initially hidden group containing the editable metadata controls.
 * @param textInput - Required text field to receive focus after activation.
 * @param dateInput - Required date field associated with the same record.
 * @example `revealManagementMetadataEditor(editorFields, tagInput, dateInput)` starts editing that record.
 * @author GPT-6 Luna
 */
export function revealManagementMetadataEditor(
    fields: HTMLElement,
    textInput: HTMLInputElement,
    dateInput: HTMLInputElement,
): void {
    fields.hidden = false;
    fields.dataset.metadataEditorActive = "true";
    fields.setAttribute("aria-hidden", "false");
    textInput.disabled = false;
    dateInput.disabled = false;
    const panel = fields.closest(".accordion-item")?.querySelector<HTMLElement>(".accordion-collapse");

    if(panel) {
        Collapse.getOrCreateInstance(panel, { toggle: false }).show();
    }
    textInput.focus();
}

/** Hides and disables one record's metadata fields after a successful unchanged-metadata save.
 *
 * @param fields - The metadata field group to return to its read-only state.
 * @param textInput - Text field to disable while the editor is hidden.
 * @param dateInput - Date field to disable while the editor is hidden.
 * @example `hideManagementMetadataEditor(editorFields, tagInput, dateInput)` closes the editor.
 * @author GPT-6 Luna
 */
export function hideManagementMetadataEditor(
    fields: HTMLElement,
    textInput: HTMLInputElement,
    dateInput: HTMLInputElement,
): void {
    fields.hidden = true;
    fields.dataset.metadataEditorActive = "false";
    fields.setAttribute("aria-hidden", "true");
    textInput.disabled = true;
    dateInput.disabled = true;
}

/** Reports whether a metadata field group is currently being edited.
 *
 * @param fields - The record's editor field group.
 * @returns `true` when the editor has been activated and its draft should be submitted.
 * @example `isManagementMetadataEditorActive(editorFields)` selects draft or persisted values.
 * @author GPT-6 Luna
 */
export function isManagementMetadataEditorActive(fields: HTMLElement): boolean {
    return fields.dataset.metadataEditorActive === "true";
}
