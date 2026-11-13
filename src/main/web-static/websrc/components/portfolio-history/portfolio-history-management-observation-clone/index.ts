/**
 * Copies ready observations into the independent ID-0 draft without fetching or persisting data.
 * Capturing and applying editable values is synchronous; only destination presentation is asynchronous.
 * @module components/portfolio-history/portfolio-history-management-observation-clone
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
import htmx from "htmx.org";
import type { TemplateDelegate } from "handlebars";
import { calendarDateToUtcTimestamp, getBrowserLocalCalendarDate } from "../../../utils/calendar-date";
import { areObservationDetailsReady } from "../portfolio-history-management-observation-details";
import { NEW_OBSERVATION_FORM_ID } from "../portfolio-history-management-form";
import { captureObservationDraft, restoreObservationRows, type ObservationDraft } from "./draft";
import { focusDestination, renderRows } from "./destination";
import { bindCloneRequestGuards, trackCloneOpening, type CloneOperation } from "./operation";

const boundLists = new WeakSet<HTMLElement>();

/** Validated destination controls needed to replace an observation draft.
 * @author GPT-6.1 Sol
 */
type DestinationControls = {
    form: HTMLFormElement;
    tbody: HTMLElement;
    tag: HTMLInputElement;
    dateInput: HTMLInputElement;
};

/** Resolves and validates destination controls before any draft values are changed.
 * @author GPT-6.1 Sol
 */
function destinationControls(destination: HTMLElement): DestinationControls {
    const form = destination.querySelector<HTMLFormElement>(`#${ NEW_OBSERVATION_FORM_ID }`);
    const tbody = form?.querySelector<HTMLElement>("tbody");
    const tag = destination.querySelector<HTMLInputElement>("[data-observation-editor-tag]");
    const dateInput = destination.querySelector<HTMLInputElement>("#portfolio-history-management-observation-date");

    if(!form || !tbody || !tag || !dateInput) {
        throw new Error("Observation clone controls are unavailable.");
    }

    if(!(form.elements.namedItem("observationTimestamp.timeTag") instanceof HTMLInputElement)
        || !(form.elements.namedItem("observationTimestamp.timestamp") instanceof HTMLInputElement)) {
        throw new Error("Observation clone metadata payload is unavailable.");
    }
    return { form, tbody, tag, dateInput };
}

/** Writes the new-record metadata after the fresh rows have settled.
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
function applyMetadata(
    form: HTMLFormElement, tag: HTMLInputElement, dateInput: HTMLInputElement, draft: ObservationDraft,
): void {
    tag.value = draft.tag;
    tag.setCustomValidity(Array.from(draft.tag).length > 100 ? "Observation tag must not exceed 100 characters." : "");
    tag.classList.remove("is-invalid", "is-valid");
    const payload = form.elements.namedItem("observationTimestamp.timeTag");

    if(!(payload instanceof HTMLInputElement)) {
        throw new Error("Observation tag payload is missing.");
    }
    payload.value = draft.tag;
    dateInput.value = draft.date;
    dateInput.setCustomValidity("");
    dateInput.classList.remove("is-invalid", "is-valid");
    const timestamp = form.elements.namedItem("observationTimestamp.timestamp");

    if(!(timestamp instanceof HTMLInputElement)) {
        throw new Error("Observation timestamp payload is missing.");
    }
    timestamp.value = calendarDateToUtcTimestamp(draft.date);
    form.classList.remove("was-validated");
}

/** Captures and applies a complete ready source before waiting for any accordion transitions.
 * @author GPT-6.1 Sol
 */
function copyDraft(
    source: HTMLFormElement, destination: HTMLElement, template: TemplateDelegate,
    normalize: (row: HTMLTableRowElement) => void,
): HTMLInputElement {
    const controls = destinationControls(destination);
    const draft = captureObservationDraft(source, getBrowserLocalCalendarDate());
    const html = renderRows(template, draft);

    htmx.swap(controls.tbody, html, { swapStyle: "innerHTML", swapDelay: 0, settleDelay: 0 }, {
        afterSettleCallback: () => {
            restoreObservationRows(controls.tbody, draft);
            controls.tbody.querySelectorAll<HTMLTableRowElement>("tr").forEach(normalize);
        },
    });
    applyMetadata(controls.form, controls.tag, controls.dateInput, draft);
    return controls.tag;
}

/** Presents an applied draft and releases only the operation that still owns its source button.
 * @author GPT-6.1 Sol
 */
async function presentDraft(
    destination: HTMLElement, tag: HTMLInputElement, button: HTMLButtonElement,
    status: HTMLElement, operation: CloneOperation, finished: () => void,
): Promise<void> {
    button.setAttribute("aria-busy", "true");

    try {
        await focusDestination(destination, tag, operation);

        if(operation.owns()) {
            status.textContent = "Observation copied to the new draft.";
        }
    } catch(error) {
        if(operation.owns()) {
            status.textContent = "Unable to clone observation. Please try again.";
            console.error(error);
        }
    } finally {
        operation.dispose();
        finished();
    }
}

/**
 * Binds one delegated clone controller. Only ready source forms may be copied.
 * @param list - Stable paginated list owner.
 * @param rowTemplate - Returns the initialized allocation-row template.
 * @param normalize - Rebinds quote state after restoration.
 * @example `bindObservationCloning(list, () => rowTemplate, normalizeRow);`
 * @author GPT-6.1 Terra
 * @author GPT-6.1 Sol
 */
export function bindObservationCloning(
    list: HTMLElement, rowTemplate: () => TemplateDelegate, normalize: (row: HTMLTableRowElement) => void,
): void {
    if(boundLists.has(list)) {
        return;
    }
    boundLists.add(list);
    let active: CloneOperation | undefined;
    let pendingSource: HTMLElement | undefined;
    const buttonOwners = new WeakMap<HTMLButtonElement, CloneOperation>();
    const cancel = () => active?.controller.abort();
    const isSaving = bindCloneRequestGuards(list, cancel);

    list.addEventListener("click", event => {
        const button = (event.target as Element)?.closest<HTMLButtonElement>("[data-observation-clone]");

        if(!button || button.disabled || !list.contains(button)) {
            return;
        }
        event.preventDefault();
        event.stopPropagation();
        const item = button.closest<HTMLElement>(".accordion-item");
        const source = item?.querySelector<HTMLFormElement>("form");

        const destination = list.querySelector<HTMLFormElement>(`#${ NEW_OBSERVATION_FORM_ID }`)
            ?.closest<HTMLElement>(".accordion-item");
        const status = item?.querySelector<HTMLElement>("[data-observation-clone-status]");

        if(!item || !source || !destination || !status) {
            return;
        }

        if(item === destination || !areObservationDetailsReady(source)) {
            return;
        }

        if(isSaving()) {
            status.textContent = "Wait for the current save to finish, then try again.";
            return;
        }

        if(active && !active.controller.signal.aborted && pendingSource === item) {
            return;
        }
        cancel();

        try {
            const tag = copyDraft(source, destination, rowTemplate(), normalize);
            const operation = trackCloneOpening(list, item, destination);
            active = operation;
            pendingSource = item;
            buttonOwners.set(button, operation);

            void presentDraft(destination, tag, button, status, operation, () => {
                if(buttonOwners.get(button) === operation) {
                    button.removeAttribute("aria-busy");
                    buttonOwners.delete(button);
                }

                if(active === operation) {
                    active = undefined;
                    pendingSource = undefined;
                }
            });
        } catch(error) {
            status.textContent = "Unable to clone observation. Please try again.";
            console.error(error);
        }
    });
}
