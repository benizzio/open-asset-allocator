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
 * Authored by: GitHub Copilot
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
 * Authored by: GitHub Copilot
 */
function getPropertyHelper(obj: Record<string, unknown> | undefined | null, propertyKey: string): unknown {
    return obj ? (obj as Record<string, unknown>)[propertyKey] : undefined;
}

/**
 * Safely coerces a value to a finite number.
 *
 * Supported coercions:
 * - number: returned if finite; otherwise 0
 * - string: Number(trimmed); non-finite becomes 0
 * - boolean: true => 1, false => 0
 * - bigint: converted via Number()
 * - others (null/undefined/object/symbol): 0
 *
 * Authored by: GitHub Copilot
 */
function coerceToFiniteNumber(value: unknown): number {
    if(typeof value === "number") {
        return Number.isFinite(value) ? value : 0;
    }

    if(typeof value === "string") {
        const n = Number(value.trim());
        return Number.isFinite(n) ? n : 0;
    }

    if(typeof value === "boolean") {
        return value ? 1 : 0;
    }

    if(typeof value === "bigint") {
        return Number(value);
    }

    return 0;
}

/**
 * Performs a basic arithmetic operation between two values.
 *
 * Allowed operations (case-insensitive):
 * - Addition: "+", "add", "plus", "sum"
 * - Subtraction: "-", "sub", "minus"
 * - Multiplication: "*", "x", "mul", "times"
 * - Division: "/", "div" (division by zero yields 0)
 * - Modulo: "%", "mod", "rem" (modulo by zero yields 0)
 *
 * @param a - Left operand (number-like).
 * @param op - Operation symbol or keyword.
 * @param b - Right operand (number-like).
 * @returns The finite numeric result of the operation; returns 0 for invalid ops or non-finite results.
 *
 * @example
 * {{math 2 "+" 3}}         {{!-- 5 --}}
 * {{math "10" "/" x}}      {{!-- coerces inputs; division by zero => 0 --}}
 * {{math price "*" qty}}    {{!-- multiplication --}}
 *
 * Authored by: GitHub Copilot
 */
function mathHelper(a: unknown, op: unknown, b: unknown): number {

    const left = coerceToFiniteNumber(a);
    const right = coerceToFiniteNumber(b);

    const opStr = typeof op === "string" ? op.trim().toLowerCase() : String(op);

    let result: number;

    switch(opStr) {
        case "+":
        case "add":
        case "plus":
        case "sum":
            result = left + right;
            break;

        case "-":
        case "sub":
        case "minus":
            result = left - right;
            break;

        case "*":
        case "x":
        case "mul":
        case "times":
            result = left * right;
            break;

        case "/":
        case "div":
            result = right === 0 ? 0 : left / right;
            break;

        case "%":
        case "mod":
        case "rem":
            result = right === 0 ? 0 : left % right;
            break;

        default:
            result = 0;
            break;
    }

    return Number.isFinite(result) ? result : 0;
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
 * Renders the block when arg1 !== arg2; otherwise renders the inverse block.
 *
 * @this unknown
 * @param arg1 - Left-hand value for strict inequality comparison.
 * @param arg2 - Right-hand value for strict inequality comparison.
 * @param options - Handlebars block options providing `fn`/`inverse`.
 * @returns The rendered block depending on the comparison result.
 *
 * @example
 * {{#ifNotEquals a b}} not equal {{else}} equal {{/ifNotEquals}}
 *
 * @author GitHub Copilot
 */
function ifNotEqualsHelper(this: unknown, arg1: unknown, arg2: unknown, options: HelperOptions): string {
    return arg1 !== arg2 ? options.fn(this) : options.inverse(this);
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
    handlebars.registerHelper("stringify", stringifyHelper);
    handlebars.registerHelper("concat", concatHelper);
    handlebars.registerHelper("eachReverse", eachReverseHelper);
    handlebars.registerHelper("getProperty", getPropertyHelper);
    handlebars.registerHelper("ifEquals", ifEqualsHelper);
    handlebars.registerHelper("ifNotEquals", ifNotEqualsHelper);
    handlebars.registerHelper("math", mathHelper);
}