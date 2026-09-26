/**
 * Provides shared Handlebars template rendering for templates stored in DOM elements.
 *
 * @module infra/handlebars/handlebars-render
 * @author OpenCode
 */

import * as handlebars from "handlebars";

/**
 * Compiles a Handlebars template from a DOM element and renders it with the supplied data.
 * Standard Handlebars interpolation, such as `{{name}}`, remains HTML-escaped. When the template element is absent,
 * the function returns an empty string so callers can safely assign the result to an element's `innerHTML`.
 *
 * @param templateElementId - ID of the DOM element whose inner HTML contains the Handlebars template.
 * @param data - Values made available to the Handlebars template during rendering.
 * @returns Rendered HTML, or an empty string when the template element does not exist.
 *
 * @example
 * ```ts
 * const html = renderTemplateByElementId("asset-row-template", { name: "Stocks & Bonds" });
 * // Normal {{name}} interpolation renders "Stocks &amp; Bonds".
 * container.innerHTML = html;
 * ```
 *
 * @author OpenCode
 */
export function renderTemplateByElementId(templateElementId: string, data: unknown): string {
    const templateElement = document.getElementById(templateElementId);

    return templateElement ? handlebars.compile(templateElement.innerHTML)(data) : "";
}
