/**
 * Coordinates editable and read-only header states for persisted management metadata.
 *
 * Record-specific management controllers use these functions to reveal prefilled metadata fields and restore the
 * read-only label and edit control after a successful save. Example: `revealManagementMetadataEditor(fields, nameInput,
 * dateInput, editButton)` activates a persisted record's complete header editor.
 * @author GPT-6 Luna
 */

import { Collapse } from "bootstrap";

/** Reveals and enables one record's metadata fields, then focuses its text input.
 *
 * @param fields - The initially hidden group containing the editable metadata controls.
 * @param textInput - Required text field to receive focus after activation.
 * @param dateInput - Required date field associated with the same record.
 * @param editButton - Pencil control whose visibility and expanded state follow the editor.
 * @example `revealManagementMetadataEditor(editorFields, tagInput, dateInput, pencil)` starts editing that record.
 * @author GPT-6 Luna
 */
export function revealManagementMetadataEditor(
    fields: HTMLElement,
    textInput: HTMLInputElement,
    dateInput: HTMLInputElement,
    editButton: HTMLButtonElement,
): void {
    fields.hidden = false;
    fields.dataset.metadataEditorActive = "true";
    fields.setAttribute("aria-hidden", "false");
    textInput.disabled = false;
    dateInput.disabled = false;
    updateManagementMetadataHeader(fields, true, editButton);
    const panel = fields.closest(".accordion-item")?.querySelector<HTMLElement>(".accordion-collapse");

    if(panel) {
        Collapse.getOrCreateInstance(panel, { toggle: false }).show();
    }
    textInput.focus();
}

/** Hides and disables one record's metadata fields and restores its read-only header after a successful save.
 *
 * @param fields - The metadata field group to return to its read-only state.
 * @param textInput - Text field to disable while the editor is hidden.
 * @param dateInput - Date field to disable while the editor is hidden.
 * @example `hideManagementMetadataEditor(editorFields, tagInput, dateInput)` restores the read-only header.
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
    updateManagementMetadataHeader(fields, false);
}

/** Updates the label, pencil, and editing state of the field group's owning accordion header.
 * @author GPT-6 Luna
 */
function updateManagementMetadataHeader(
    fields: HTMLElement,
    editing: boolean,
    editButton?: HTMLButtonElement,
): void {
    const header = fields.closest<HTMLElement>("[data-management-metadata-header]");

    if(!header) {
        return;
    }

    header.dataset.metadataEditorState = editing ? "editing" : "read-only";
    const label = header.querySelector<HTMLElement>("[data-metadata-editor-label]");
    const trigger = editButton ?? header.querySelector<HTMLButtonElement>("[data-metadata-editor-trigger]");

    if(label) {
        label.hidden = editing;
    }

    if(trigger) {
        trigger.hidden = editing;
        trigger.setAttribute("aria-expanded", String(editing));
    }
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
