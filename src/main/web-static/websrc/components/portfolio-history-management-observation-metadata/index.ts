/**
 * Prepares observation tag and timestamp metadata for the portfolio-history management POST form.
 *
 * Prepare the form before submission. It preserves exact persisted instants when the browser-local
 * calendar date is unchanged and converts only newly selected dates to UTC local-midnight timestamps.
 * Example: `prepareObservationMetadataForSubmit(form)` copies the current editor values into the form.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */

import { calendarDateToUtcTimestamp } from "../../utils/calendar-date";
import { isManagementMetadataEditorActive } from "../management-metadata-editor";
import { areObservationDetailsReady } from "../portfolio-history-management-observation-details";

const FORM_PREFIX = "portfolio-history-management-form-";
const NEW_DATE_INPUT_ID = "portfolio-history-management-observation-date";
const OBSERVATION_TAG_NAME = "observationTimestamp.timeTag";
const OBSERVATION_TIMESTAMP_NAME = "observationTimestamp.timestamp";

/** Validates and prepares canonical metadata inputs before native submission collects form values.
 *
 * @param form - The owning observation management form.
 * @returns Whether submission may proceed.
 * @example `if (prepareObservationMetadataForSubmit(form)) form.requestSubmit();`
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
export function prepareObservationMetadataForSubmit(form: HTMLFormElement): boolean {
    const formId = form.getAttribute("id");

    if(!formId?.startsWith(FORM_PREFIX) || !areObservationDetailsReady(form)) {
        return false;
    }

    const observationId = formId.slice(FORM_PREFIX.length);
    const isNewObservation = observationId === "0";
    const item = form.closest<HTMLElement>(".accordion-item");
    const fields = item?.querySelector<HTMLElement>("[data-metadata-editor-active]");
    const isEditing = fields instanceof HTMLElement && isManagementMetadataEditorActive(fields);
    const tagInput = item?.querySelector<HTMLInputElement>("[data-observation-editor-tag]");

    const dateInput = isNewObservation
        ? item?.querySelector<HTMLInputElement>(`#${ NEW_DATE_INPUT_ID }`)
        : item?.querySelector<HTMLInputElement>("[data-observation-editor-date]");
    const tagPayload = form.elements.namedItem(OBSERVATION_TAG_NAME);
    const timestampPayload = form.elements.namedItem(OBSERVATION_TIMESTAMP_NAME);

    if(!(tagPayload instanceof HTMLInputElement) || !(timestampPayload instanceof HTMLInputElement)) {
        return false;
    }

    if(isEditing && !copyEditedTag(tagInput, tagPayload)) {
        return false;
    }

    if((isNewObservation || isEditing)
        && !setObservationTimestamp(form, dateInput, timestampPayload, isNewObservation)) {
        return false;
    }

    if(!tagPayload.value) {
        return false;
    }

    if(!isNewObservation) {
        form.dataset.submittedMetadataChanged = String(
            tagPayload.value !== form.dataset.persistedTimeTag
                || timestampPayload.value !== form.dataset.persistedTimestamp,
        );
    }
    return true;
}

/** Copies and validates the active observation-tag editor value.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function copyEditedTag(
    tagInput: HTMLInputElement | null,
    tagPayload: HTMLInputElement,
): boolean {
    if(!tagInput?.checkValidity()) {
        tagInput?.reportValidity();
        return false;
    }

    tagPayload.value = tagInput.value;
    return true;
}

/** Sets an observation timestamp from the selected date or exact persisted instant.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
function setObservationTimestamp(
    form: HTMLFormElement,
    dateInput: HTMLElement | null,
    timestampPayload: HTMLInputElement,
    isNewObservation: boolean,
): boolean {
    if(!(dateInput instanceof HTMLInputElement) || !dateInput.value) {
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
