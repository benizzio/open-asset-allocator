/**
 * Prepares allocation-plan name and planned-date values for management-form POST requests.
 *
 * Bind once to the allocation-plan management container. Existing editor fields are copied into the hidden canonical
 * form fields immediately before HTMX serializes the request, while untouched and new-plan forms keep their values.
 * Example: `bindAllocationPlanMetadataRequest(container)` installs the idempotent request hook.
 * @author GPT-6 Luna
 */

import { RequestConfigEventDetail } from "../../infra/htmx";
import {
    isManagementMetadataEditorActive,
    revealManagementMetadataEditor,
} from "../management-metadata-editor";

const FORM_PREFIX = "allocation-plan-management-form-";
const boundContainers = new WeakSet<HTMLElement>();

type PlanRequest = RequestConfigEventDetail & { formData: FormData; };

/** Binds allocation-plan metadata serialization to one management-list container, at most once.
 *
 * @param container - Allocation-plan management root receiving bubbling HTMX request events.
 * @example `bindAllocationPlanMetadataRequest(managementContainer)` serializes active plan editor values.
 * @author GPT-6 Luna
 */
export function bindAllocationPlanMetadataRequest(container: HTMLElement): void {
    if(boundContainers.has(container)) {
        return;
    }

    boundContainers.add(container);
    container.addEventListener("htmx:configRequest", prepareAllocationPlanMetadataForPost);
}

/** Reveals one plan's persisted name/date editor and focuses its name field.
 *
 * @param button - Pencil control activated by pointer or keyboard.
 * @param allocationPlanId - Persisted plan whose metadata should be edited.
 * @example `activateAllocationPlanEditor(button, 12)` opens plan 12's editor.
 * @author GPT-6 Luna
 */
export function activateAllocationPlanEditor(button: HTMLButtonElement, allocationPlanId: number): void {
    const item = document.getElementById(`allocation-plan-management-container-${ allocationPlanId }`);
    const fields = item?.querySelector<HTMLElement>("[data-metadata-editor-active]");
    const nameInput = item?.querySelector<HTMLInputElement>("[data-allocation-plan-editor-name]");
    const dateInput = item?.querySelector<HTMLInputElement>("[data-allocation-plan-editor-date]");

    if(!fields || !nameInput || !dateInput) {
        return;
    }

    revealManagementMetadataEditor(fields, nameInput, dateInput, button);
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

/** Copies current allocation-plan values into the owning form's request payload.
 * @author GPT-6 Luna
 */
function prepareAllocationPlanMetadataForPost(event: Event): void {
    const request = (event as CustomEvent<PlanRequest>).detail;

    if(request.verb.toLowerCase() !== "post") {
        return;
    }

    const requestElement = request.elt;
    const form = requestElement instanceof HTMLFormElement ? requestElement : requestElement.closest("form");
    const formId = form instanceof HTMLFormElement ? form.getAttribute("id") : null;

    if(!(form instanceof HTMLFormElement) || !formId?.startsWith(FORM_PREFIX)) {
        return;
    }

    const item = form.closest<HTMLElement>(".accordion-item");
    const fields = item?.querySelector<HTMLElement>("[data-metadata-editor-active]");
    const nameEditor = item?.querySelector<HTMLInputElement>("[data-allocation-plan-editor-name]");
    const dateEditor = item?.querySelector<HTMLInputElement>("[data-allocation-plan-editor-date]");
    const namePayload = form.elements.namedItem("name");
    const datePayload = form.elements.namedItem("plannedDate");

    if(!(namePayload instanceof HTMLInputElement) || !(datePayload instanceof HTMLInputElement)) {
        event.preventDefault();
        return;
    }

    if(fields && isManagementMetadataEditorActive(fields)
        && !copyEditedPlanMetadata(nameEditor, dateEditor, namePayload, datePayload, event)) {
        return;
    }

    request.formData.set(namePayload.name, namePayload.value);
    request.formData.set(datePayload.name, datePayload.value);

    if(form.dataset.persistedName !== undefined) {
        form.dataset.submittedMetadataChanged = String(
            namePayload.value !== form.dataset.persistedName || datePayload.value !== form.dataset.persistedDate,
        );
    }
}

/** Validates and copies active name/date controls into the form's canonical payload fields.
 * @author GPT-6 Luna
 */
function copyEditedPlanMetadata(
    nameEditor: HTMLInputElement | null,
    dateEditor: HTMLInputElement | null,
    namePayload: HTMLInputElement,
    datePayload: HTMLInputElement,
    event: Event,
): boolean {
    if(!nameEditor || !dateEditor || !nameEditor.checkValidity() || !dateEditor.checkValidity()) {
        event.preventDefault();
        nameEditor?.reportValidity();
        dateEditor?.reportValidity();
        return false;
    }

    namePayload.value = nameEditor.value;
    datePayload.value = dateEditor.value;
    return true;
}
