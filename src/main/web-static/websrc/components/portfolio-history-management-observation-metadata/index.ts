/**
 * Prepares observation tag and timestamp metadata for the portfolio-history management POST form.
 *
 * Bind this module once to the management-list container. It preserves exact persisted instants when the browser-local
 * calendar date is unchanged and converts only newly selected dates to UTC local-midnight timestamps.
 * Example: `bindObservationMetadataRequest(container)` installs the idempotent HTMX request hook.
 * @author GPT-6 Luna
 */

import { RequestConfigEventDetail } from "../../infra/htmx";
import { calendarDateToUtcTimestamp } from "../../utils/calendar-date";
import { isManagementMetadataEditorActive } from "../management-metadata-editor";
import { areObservationDetailsReady } from "../portfolio-history-management-observation-details";

const FORM_PREFIX = "portfolio-history-management-form-";
const NEW_DATE_INPUT_ID = "portfolio-history-management-observation-date";
const OBSERVATION_TAG_NAME = "observationTimestamp.timeTag";
const OBSERVATION_TIMESTAMP_NAME = "observationTimestamp.timestamp";
const boundContainers = new WeakSet<HTMLElement>();

type ObservationRequest = RequestConfigEventDetail & { formData: FormData; };

/** Binds observation metadata serialization to one management-list container, at most once.
 *
 * @param container - Portfolio-history management root receiving bubbling HTMX request events.
 * @example `bindObservationMetadataRequest(managementContainer)` serializes active observation editor values.
 * @author GPT-6 Luna
 */
export function bindObservationMetadataRequest(container: HTMLElement): void {
    if(boundContainers.has(container)) {
        return;
    }

    boundContainers.add(container);
    container.addEventListener("htmx:configRequest", prepareObservationMetadataForPost);
}

/** Serializes metadata only for POSTs owned by one observation management form.
 * @author GPT-6 Luna
 */
function prepareObservationMetadataForPost(event: Event): void {
    const request = (event as CustomEvent<ObservationRequest>).detail;

    if(request.verb.toLowerCase() !== "post") {
        return;
    }

    const form = getObservationForm(request);

    if(!form) {
        return;
    }

    if(!areObservationDetailsReady(form)) {
        event.preventDefault();
        return;
    }

    const observationId = form.getAttribute("id")?.slice(FORM_PREFIX.length) ?? "";
    const isNewObservation = observationId === "0";
    const item = form.closest<HTMLElement>(".accordion-item");
    const fields = item?.querySelector<HTMLElement>("[data-metadata-editor-active]");
    const isEditing = fields instanceof HTMLElement && isManagementMetadataEditorActive(fields);
    const tagInput = item?.querySelector<HTMLInputElement>("[data-observation-editor-tag]");

    const dateInput = isNewObservation
        ? document.getElementById(NEW_DATE_INPUT_ID)
        : item?.querySelector<HTMLInputElement>("[data-observation-editor-date]");
    const tagPayload = form.elements.namedItem(OBSERVATION_TAG_NAME);
    const timestampPayload = form.elements.namedItem(OBSERVATION_TIMESTAMP_NAME);

    if(!(tagPayload instanceof HTMLInputElement) || !(timestampPayload instanceof HTMLInputElement)) {
        event.preventDefault();
        return;
    }

    if(isEditing && !copyEditedTag(tagInput, tagPayload, event)) {
        return;
    }

    if((isNewObservation || isEditing)
        && !setObservationTimestamp(form, dateInput, timestampPayload, isNewObservation, event)) {
        return;
    }

    if(!tagPayload.value) {
        event.preventDefault();
        return;
    }

    request.formData.set(tagPayload.name, tagPayload.value);
    request.formData.set(timestampPayload.name, timestampPayload.value);

    if(!isNewObservation) {
        form.dataset.submittedMetadataChanged = String(
            tagPayload.value !== form.dataset.persistedTimeTag
                || timestampPayload.value !== form.dataset.persistedTimestamp,
        );
    }
}

/** Finds the observation form that owns the current HTMX POST.
 * @author GPT-6 Luna
 */
function getObservationForm(request: ObservationRequest): HTMLFormElement | null {
    const requestElement = request.elt;
    const form = requestElement instanceof HTMLFormElement ? requestElement : requestElement.closest("form");
    const formId = form instanceof HTMLFormElement ? form.getAttribute("id") : null;

    return form instanceof HTMLFormElement && formId?.startsWith(FORM_PREFIX) ? form : null;
}

/** Copies and validates the active observation-tag editor value.
 * @author GPT-6 Luna
 */
function copyEditedTag(
    tagInput: HTMLInputElement | null,
    tagPayload: HTMLInputElement,
    event: Event,
): boolean {
    if(!tagInput?.checkValidity()) {
        event.preventDefault();
        tagInput?.reportValidity();
        return false;
    }

    tagPayload.value = tagInput.value;
    return true;
}

/** Sets an observation timestamp from the selected date or exact persisted instant.
 * @author GPT-6 Luna
 */
function setObservationTimestamp(
    form: HTMLFormElement,
    dateInput: HTMLElement | null,
    timestampPayload: HTMLInputElement,
    isNewObservation: boolean,
    event: Event,
): boolean {
    if(!(dateInput instanceof HTMLInputElement) || !dateInput.value) {
        event.preventDefault();

        if(dateInput instanceof HTMLInputElement) {
            dateInput.reportValidity();
        }
        return false;
    }

    try {
        timestampPayload.value = getObservationTimestamp(form, dateInput.value, isNewObservation);
        return true;
    } catch(error) {
        if(!(error instanceof RangeError)) {
            throw error;
        }

        event.preventDefault();
        dateInput.setCustomValidity("Observation date is not valid.");
        dateInput.reportValidity();
        dateInput.setCustomValidity("");
        return false;
    }
}

/** Resolves the timestamp while retaining the persisted string when the local date is unchanged.
 * @author GPT-6 Luna
 */
function getObservationTimestamp(form: HTMLFormElement, selectedDate: string, isNewObservation: boolean): string {
    const persistedDate = form.dataset.persistedLocalDate;
    const persistedTimestamp = form.dataset.persistedTimestamp;

    return !isNewObservation && selectedDate === persistedDate && persistedTimestamp
        ? persistedTimestamp
        : calendarDateToUtcTimestamp(selectedDate);
}
