/**
 * Coordinates editable and read-only header states for persisted management metadata.
 *
 * Record-specific management controllers use these functions to reveal prefilled metadata fields and restore the
 * read-only label and edit control after a successful save. Example: `revealManagementMetadataEditor(fields, nameInput,
 * dateInput, editButton)` activates a persisted record's complete header editor.
 * @author GPT-6 Luna
 * @author GPT-6.1 Sol
 */

import { Collapse } from "bootstrap";

const boundHeaderToggleContainers = new WeakSet<HTMLElement>();

const interactiveHeaderControlSelector = [
    "button",
    "a",
    "input",
    "select",
    "textarea",
    "label",
    "summary",
    "[role='button']",
    "[role='link']",
    "[contenteditable]:not([contenteditable='false'])",
].join(", ");

/** Binds read-only management headers so clicks outside interactive controls toggle their accordion panel.
 *
 * Bind once to a management container. Delegation also covers header markup inserted or replaced by HTMX. For example,
 * `bindReadOnlyManagementMetadataHeaderToggle(container)` enables label and background clicks without intercepting the
 * pencil, accordion button, or editable controls.
 * @param container - Management-list container owning persisted metadata headers.
 * @author GPT-6 Luna
 */
export function bindReadOnlyManagementMetadataHeaderToggle(container: HTMLElement): void {
    if(boundHeaderToggleContainers.has(container)) {
        return;
    }

    boundHeaderToggleContainers.add(container);
    container.addEventListener("click", toggleReadOnlyManagementMetadataPanel);
}

/** Toggles the panel for a read-only header click that did not originate from an interactive control.
 * @author GPT-6 Luna
 */
function toggleReadOnlyManagementMetadataPanel(event: MouseEvent): void {
    const target = event.target;

    if(!(target instanceof Element)) {
        return;
    }

    const header = target.closest<HTMLElement>("[data-management-metadata-header]");

    if(header?.dataset.metadataEditorState !== "read-only" || target.closest(interactiveHeaderControlSelector)) {
        return;
    }

    const item = header.closest<HTMLElement>(".accordion-item");
    const panel = item?.querySelector<HTMLElement>(":scope > .accordion-collapse");

    if(panel) {
        Collapse.getOrCreateInstance(panel, { toggle: false }).toggle();
    }
}

/** Resolves one record's metadata controls and reveals its editor when every required control exists.
 *
 * Pass the owning record element and its text/date input selectors. A missing record, editor field group, or input
 * leaves the editor unchanged. Complete records use the standard reveal behavior, including opening the panel and
 * focusing the text input.
 * @param item - Record element containing the metadata controls, or null when the record is absent.
 * @param textInputSelector - Selector for the record's text input, scoped to the record element.
 * @param dateInputSelector - Selector for the record's date input, scoped to the record element.
 * @param editButton - Pencil control whose visibility and expanded state follow the editor.
 * @example `activateManagementMetadataEditor(item, "[data-editor-text]", "[data-editor-date]", pencil)` activates the
 * record's editor when its controls exist.
 * @author GPT-6.1 Sol
 */
export function activateManagementMetadataEditor(
    item: HTMLElement | null,
    textInputSelector: string,
    dateInputSelector: string,
    editButton: HTMLButtonElement,
): void {
    const fields = item?.querySelector<HTMLElement>("[data-metadata-editor-active]");
    const textInput = item?.querySelector<HTMLInputElement>(textInputSelector);
    const dateInput = item?.querySelector<HTMLInputElement>(dateInputSelector);

    if(fields && textInput && dateInput) {
        revealManagementMetadataEditor(fields, textInput, dateInput, editButton);
    }
}

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
