/**
 * Prepares allocation-plan name and planned-date values in the management form before submission.
 *
 * Active editor values are validated and copied to canonical form fields before HTMX collects parameters.
 * Example: `prepareAllocationPlanMetadataForSubmit(form)` returns whether submission may proceed.
 * @author GPT-6 Luna
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */

import {
    activateManagementMetadataEditor,
    isManagementMetadataEditorActive,
} from "../management-metadata-editor";

const FORM_PREFIX = "allocation-plan-management-form-";
const boundContainers = new WeakSet<HTMLElement>();

/** Installs an idempotent native-submit guard before HTMX collects form values.
 * @param container - Management list containing the plan forms.
 * @example `bindAllocationPlanMetadataSubmitGuard(container)` prepares native form submissions.
 * @author GPT-6.1 Terra
 */
export function bindAllocationPlanMetadataSubmitGuard(container: HTMLElement): void {
    if(boundContainers.has(container)) {
        return;
    }

    boundContainers.add(container);

    container.addEventListener("submit", event => {
        const form = event.target;

        if(form instanceof HTMLFormElement && form.getAttribute("id")?.startsWith(FORM_PREFIX)
            && !prepareAllocationPlanMetadataForSubmit(form)) {
            event.preventDefault();
            event.stopImmediatePropagation();
        }
    }, true);
}

/** Synchronizes the visible plan name to its canonical form field during editing.
 * @param nameInput - Visible plan name field.
 * @param allocationPlanId - ID of the owning plan.
 * @example `handleInputAllocationPlanName(input, 12)` updates plan 12's form.
 * @author benizzio
 * @author GPT-6.1 Terra
 */
export function handleInputAllocationPlanName(nameInput: HTMLInputElement, allocationPlanId: number): void {
    const form = document.getElementById(`${ FORM_PREFIX }${ allocationPlanId }`) as HTMLFormElement;
    const name = form.elements.namedItem("name") as HTMLInputElement;
    name.value = nameInput.value;
}

/** Reveals one plan's persisted name/date editor and focuses its name field.
 *
 * @param button - Pencil control activated by pointer or keyboard.
 * @param allocationPlanId - Persisted plan whose metadata should be edited.
 * @example `activateAllocationPlanEditor(button, 12)` opens plan 12's editor.
 * @author GPT-6 Luna
 * @author GPT-6.1 Sol
 */
export function activateAllocationPlanEditor(button: HTMLButtonElement, allocationPlanId: number): void {
    activateManagementMetadataEditor(
        document.getElementById(`allocation-plan-management-container-${ allocationPlanId }`),
        "[data-allocation-plan-editor-name]",
        "[data-allocation-plan-editor-date]",
        button,
    );
}

/** Copies one plan editor date into the hidden date field serialized by the existing form.
 *
 * @param dateInput - Visible date editor associated with the persisted plan.
 * @param allocationPlanId - Persisted plan whose hidden payload value should be updated.
 * @example `handleInputAllocationPlanDate(dateInput, 12)` updates plan 12's payload.
 * @author GPT-6 Luna
 */
export function handleInputAllocationPlanDate(dateInput: HTMLInputElement, allocationPlanId: number): void {
    const form = document.getElementById(`${ FORM_PREFIX }${ allocationPlanId }`);

    const plannedDateInput = form instanceof HTMLFormElement
        ? form.elements.namedItem("plannedDate")
        : null;

    if(plannedDateInput instanceof HTMLInputElement) {
        plannedDateInput.value = dateInput.value;
    }
}

/** Validates and copies current metadata into the owning form before submission.
 *
 * Returns false for missing canonical fields or invalid active editors, reporting invalid controls to the browser.
 * @param form - The allocation-plan management form to prepare.
 * @example `if (prepareAllocationPlanMetadataForSubmit(form)) form.requestSubmit();`
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
export function prepareAllocationPlanMetadataForSubmit(form: HTMLFormElement): boolean {
    const item = form.closest<HTMLElement>(".accordion-item");
    const fields = item?.querySelector<HTMLElement>("[data-metadata-editor-active]");
    const nameEditor = item?.querySelector<HTMLInputElement>("[data-allocation-plan-editor-name]");
    const dateEditor = item?.querySelector<HTMLInputElement>("[data-allocation-plan-editor-date]");
    const namePayload = form.elements.namedItem("name");
    const datePayload = form.elements.namedItem("plannedDate");

    if(!item || !(namePayload instanceof HTMLInputElement) || !(datePayload instanceof HTMLInputElement)) {
        return false;
    }

    if(fields && isManagementMetadataEditorActive(fields)
        && !copyEditedPlanMetadata(nameEditor, dateEditor, namePayload, datePayload)) {
        return false;
    }

    if(form.dataset.persistedName === undefined) {
        const newName = item.querySelector<HTMLInputElement>(".allocation-plan-management-new-name");

        if(!newName) {
            return false;
        }

        namePayload.value = newName.value;
    }

    return true;
}

/** Validates and copies active name/date controls into the form's canonical payload fields.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function copyEditedPlanMetadata(
    nameEditor: HTMLInputElement | null,
    dateEditor: HTMLInputElement | null,
    namePayload: HTMLInputElement,
    datePayload: HTMLInputElement,
): boolean {
    if(!nameEditor || !dateEditor || !nameEditor.checkValidity() || !dateEditor.checkValidity()) {
        nameEditor?.reportValidity();
        dateEditor?.reportValidity();
        return false;
    }

    namePayload.value = nameEditor.value;
    datePayload.value = dateEditor.value;
    return true;
}
