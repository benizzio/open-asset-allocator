import * as handlebars from "handlebars";
import DomUtils from "../dom/dom-utils";

/**
 * Serializes an object into a <script type="application/json"> tag.
 *
 * @param id - Unique element id for the script tag.
 * @param object - The value to serialize as JSON.
 * @returns An HTML string containing the JSON payload.
 *
 * @example
 * {{{ domJSON "portfolio" this }}}
 *
 * @author GitHub Copilot
 */
function domJSONHelper(id: string, object: object): string {
    const safeId = DomUtils.escapeHtml(id ?? "");
    const safeJson = DomUtils.escapeHtmlPreserveQuotes(JSON.stringify(object));

    return `\n <script id="${ safeId }" type="application/json">${ safeJson }</script> \n`;
}

export function registerHandlebarsDOMHelpers() {

    // Register all helpers with their names
    handlebars.registerHelper("domJSON", domJSONHelper);
}