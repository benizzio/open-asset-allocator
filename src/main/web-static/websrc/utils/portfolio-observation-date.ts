/**
 * Provides browser-local calendar-date formatting and observation timestamp conversion.
 * Use these helpers when an observation is selected by date: serialize its local midnight as a UTC instant.
 * @author GPT-6 Luna
 */

/**
 * Returns the browser's current local calendar date in the native date-input format.
 * @example `getBrowserLocalCalendarDate()` returns a value such as `2026-10-18`.
 * @author GPT-6 Luna
 */
export function getBrowserLocalCalendarDate(): string {
    const today = new Date();
    const month = String(today.getMonth() + 1).padStart(2, "0");
    const day = String(today.getDate()).padStart(2, "0");

    return `${ today.getFullYear() }-${ month }-${ day }`;
}

/**
 * Converts a selected YYYY-MM-DD calendar date to the UTC instant at browser-local midnight.
 * The timezone offset is resolved for the selected date, so daylight-saving changes are handled by the browser.
 * @param calendarDate - A native date-input value such as `2026-10-18`.
 * @example `calendarDateToUtcTimestamp("2026-10-18")` returns that date's local midnight as an ISO UTC string.
 * @author GPT-6 Luna
 */
export function calendarDateToUtcTimestamp(calendarDate: string): string {
    const dateParts = /^(\d{4})-(\d{2})-(\d{2})$/.exec(calendarDate);

    if(!dateParts) {
        throw new RangeError("Observation date must use the YYYY-MM-DD format.");
    }

    const year = Number(dateParts[1]);
    const month = Number(dateParts[2]);
    const day = Number(dateParts[3]);
    const localMidnight = new Date(`${ calendarDate }T00:00:00`);

    if(Number.isNaN(localMidnight.getTime())
        || localMidnight.getFullYear() !== year
        || localMidnight.getMonth() !== month - 1
        || localMidnight.getDate() !== day) {
        throw new RangeError("Observation date is not a valid calendar date.");
    }

    return localMidnight.toISOString();
}
