import * as handlebars from "handlebars";
import Format from "../format";
import { timestampToBrowserLocalCalendarDate } from "../../utils/calendar-date";

/** Registers Handlebars formatting helpers used by application templates.
 *
 * @module infra/handlebars/handlebars-format
 * @author GPT-6 Luna
 * @author benizzio
 */

const handlebarsFormatCurrency = (value: number | string, currency: unknown) => {
    const localCurrency = typeof currency === "string" ? currency : undefined;
    const numericValue = typeof value === "number" ? value : parseFloat(value);
    return Format.formatCurrency(numericValue, undefined, localCurrency);
};

/** Formats a persisted timestamp using the browser's local calendar date.
 *
 * @param value - An ISO-compatible timestamp.
 * @returns The timestamp's local date in `YYYY-MM-DD` format.
 * @example `{{formatTimestampLocalDate timestamp}}` renders a timestamp as the viewer's local date.
 * @author GPT-6 Luna
 */
const handlebarsFormatTimestampLocalDate = (value: string) => timestampToBrowserLocalCalendarDate(value);

/** Registers currency and local timestamp date formatting helpers with Handlebars.
 *
 * @example `registerHandlebarsFormatHelper()` makes `formatTimestampLocalDate` available to templates.
 * @author GPT-6 Luna
 * @author benizzio
 */
export function registerHandlebarsFormatHelper() {
    handlebars.registerHelper("formatCurrency", handlebarsFormatCurrency);
    handlebars.registerHelper("formatTimestampLocalDate", handlebarsFormatTimestampLocalDate);
}
