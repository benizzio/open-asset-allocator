/**
 * Provides browser-local calendar-date formatting and UTC timestamp conversion.
 * Use these helpers to serialize a selected calendar date's local midnight as a UTC instant.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author benizzio
 */

/** Formats a local year, month, and day using the date-input calendar format.
 *
 * @param year - Local calendar year.
 * @param month - One-based local calendar month.
 * @param day - Local calendar day of the month.
 * @returns The local calendar date as `YYYY-MM-DD`.
 * @author GPT-6 Luna
 */
function formatCalendarDate(year: number, month: number, day: number): string {
    return `${ String(year).padStart(4, "0") }-${ String(month).padStart(2, "0") }-${ String(day).padStart(2, "0") }`;
}

/**
 * Returns the browser's current local calendar date in the native date-input format.
 * @example `getBrowserLocalCalendarDate()` returns a value such as `2026-10-18`.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
export function getBrowserLocalCalendarDate(): string {
    const today = new Date();

    return formatCalendarDate(today.getFullYear(), today.getMonth() + 1, today.getDate());
}

/** Formats a stored timestamp as a calendar date in the viewing browser's local timezone.
 *
 * @param timestamp - An ISO-compatible timestamp representing a persisted instant.
 * @returns The instant's browser-local date as `YYYY-MM-DD`.
 * @throws {RangeError} If the timestamp cannot be parsed as a valid date.
 * @example `timestampToBrowserLocalCalendarDate("2026-10-12T15:30:00Z")` returns `"2026-10-13"` in Tokyo.
 * @author GPT-6 Luna
 */
export function timestampToBrowserLocalCalendarDate(timestamp: string): string {
    const instant = new Date(timestamp);

    if(Number.isNaN(instant.getTime())) {
        throw new RangeError("Timestamp must represent a valid date and time.");
    }

    return formatCalendarDate(instant.getFullYear(), instant.getMonth() + 1, instant.getDate());
}

/**
 * Converts a selected YYYY-MM-DD calendar date to the UTC instant at browser-local midnight.
 * The timezone offset is resolved for the selected date, so daylight-saving changes are handled by the browser.
 * @param calendarDate - A native date-input value such as `2026-10-18`.
 * @example `calendarDateToUtcTimestamp("2026-10-18")` returns that date's local midnight as an ISO UTC string.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
export function calendarDateToUtcTimestamp(calendarDate: string): string {
    const dateParts = /^(\d{4})-(\d{2})-(\d{2})$/.exec(calendarDate);

    if(!dateParts) {
        throw new RangeError("Calendar date must use the YYYY-MM-DD format.");
    }

    const year = Number(dateParts[1]);
    const month = Number(dateParts[2]);
    const day = Number(dateParts[3]);
    const localMidnight = new Date(`${ calendarDate }T00:00:00`);

    if(Number.isNaN(localMidnight.getTime())
        || localMidnight.getFullYear() !== year
        || localMidnight.getMonth() !== month - 1
        || localMidnight.getDate() !== day) {
        throw new RangeError("Invalid calendar date.");
    }

    return localMidnight.toISOString();
}
