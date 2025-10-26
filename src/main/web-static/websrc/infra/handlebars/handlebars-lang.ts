import * as handlebars from "handlebars";
import { HelperOptions } from "handlebars";

/**
 * Creates an object from named hash parameters.
 *
 * @this unknown
 * @param options - Handlebars helper options containing the `hash` key-value pairs.
 * @returns The provided `hash` object as-is.
 *
 *  @example
 * {{#with (object key1="value1" key2="value2")}}
 *   {{key1}} - {{key2}}
 * {{/with}}
 *
 * @author GitHub Copilot
 */
function objectHelper(this: unknown, options: { hash: Record<string, unknown> }): Record<string, unknown> {
    return options.hash;
}

/**
 * Converts arguments into an array (excluding the trailing Handlebars options object).
 *
 * @param args - Values to include in the array; the last argument (options) is ignored.
 * @returns A new array with the provided values.
 *
 *  @example
 * {{#each (array "a" "b" "c")}}
 *   {{this}}
 * {{/each}}
 *
 * {{!-- yields: a b c --}}
 *
 * @author GitHub Copilot
 */
function arrayHelper(...args: unknown[]): unknown[] {
    // The last argument is the Handlebars options object; exclude it.
    return args.slice(0, -1);
}

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
    return `\n            <script id="${ id }" type="application/json">${ JSON.stringify(object) }</script>\n        `;
}

/**
 * Repeats a text N times with optional prefix and suffix; returns empty for count <= 0.
 *
 * @param text - The value to repeat.
 * @param count - Number of repetitions; must be > 0.
 * @param prefix - Optional text to prepend to the repeated result.
 * @param suffix - Optional text to append to the repeated result.
 * @returns The repeated string, optionally wrapped with prefix/suffix.
 *
 *  @example
 * {{{repeater "&nbsp;" depth "" ""}}}
 * {{{repeater "★" rating "Rating: " "/5"}}}
 *
 * @author GitHub Copilot
 */
function repeaterHelper(text: string | number, count: number, prefix: string, suffix: string): string {

    if(count <= 0) {
        return "";
    }

    let result = String(text).repeat(count);
    result = prefix ? prefix + result : result;
    result = suffix ? result + suffix : result;

    return result;
}

/**
 * Converts a JavaScript value to a JSON string.
 *
 * @param object - The value to serialize.
 * @returns A JSON string representation of the value.
 *
 * @example
 * <pre>{{stringify data}}</pre>
 *
 * @author GitHub Copilot
 */
function stringifyHelper(object: unknown): string {
    return JSON.stringify(object);
}

/**
 * Concatenates multiple values into a single string.
 *
 * @param args - Values to concatenate (the trailing options object is ignored).
 * @returns The concatenated string.
 *
 * @example
 * {{concat protocol "://" host path}}
 *
 * @author GitHub Copilot
 */
function concatHelper(...args: unknown[]): string {
    const parts = args.slice(0, -1).map(v => String(v));
    return parts.join("");
}

/**
 * Iterates over an array from end to beginning, mirroring {{#each}} semantics.
 *
 * Exposes in @data:
 * - @index: 0..n-1 for the reversed order
 * - @first: true on the first item in reverse (original last)
 * - @last: true on the last item in reverse (original first)
 * - @key: original index in the source array
 *
 * @this unknown
 * @param collection - Array or function returning an array to iterate in reverse.
 * @param options - Handlebars block options providing `fn`/`inverse` and `data`.
 * @returns Rendered string for each item in reverse order (or the inverse block for empty/non-arrays).
 *
 * @example
 * {{#eachReverse items}}
 *   {{@index}} {{this}}
 * {{else}}
 *   No items.
 * {{/eachReverse}}
 *
 * @author GitHub Copilot
 */
function eachReverseHelper(this: unknown, collection: unknown, options: HelperOptions): string {
    const value = typeof collection === "function"
        ? (collection as (this: unknown) => unknown).call(this)
        : collection;

    if(!Array.isArray(value) || value.length === 0) {
        return options.inverse(this);
    }

    let out = "";

    for(let i = value.length - 1; i >= 0; i--) {
        const data = handlebars.createFrame(options.data || {});
        data.index = (value.length - 1) - i; // 0..n-1 in reverse order
        data.first = (i === value.length - 1);
        data.last = (i === 0);
        data.key = i; // original index

        out += options.fn(value[i], { data });
    }

    return out;
}

/**
 * Returns obj[key] or undefined.
 *
 * @param obj - Source object (nullable).
 * @param propertyKey - Key to read from the object.
 * @returns The property value or undefined when missing.
 *
 * @example
 * {{getProperty user "name"}}
 *
 * @author GitHub Copilot
 */
function getPropertyHelper(obj: Record<string, unknown> | undefined | null, propertyKey: string): unknown {
    return obj ? (obj as Record<string, unknown>)[propertyKey] : undefined;
}

/**
 * Renders the block when arg1 === arg2; otherwise renders the inverse block.
 *
 * @this unknown
 * @param arg1 - Left-hand value for strict equality comparison.
 * @param arg2 - Right-hand value for strict equality comparison.
 * @param options - Handlebars block options providing `fn`/`inverse`.
 * @returns The rendered block depending on the comparison result.
 *
 * @example
 * {{#ifEquals a b}} equal {{else}} not equal {{/ifEquals}}
 *
 * @author GitHub Copilot
 */
function ifEqualsHelper(this: unknown, arg1: unknown, arg2: unknown, options: HelperOptions): string {
    return arg1 === arg2 ? options.fn(this) : options.inverse(this);
}

/**
 * Registers custom Handlebars helpers that extend language functionality for template rendering.
 *
 * @author GitHub Copilot
 */
export function registerHandlebarsLangHelpers() {

    // Register all helpers with their names
    handlebars.registerHelper("object", objectHelper);
    handlebars.registerHelper("array", arrayHelper);
    handlebars.registerHelper("domJSON", domJSONHelper);
    handlebars.registerHelper("repeater", repeaterHelper);
    handlebars.registerHelper("stringify", stringifyHelper);
    handlebars.registerHelper("concat", concatHelper);
    handlebars.registerHelper("eachReverse", eachReverseHelper);
    handlebars.registerHelper("getProperty", getPropertyHelper);
    handlebars.registerHelper("ifEquals", ifEqualsHelper);
}