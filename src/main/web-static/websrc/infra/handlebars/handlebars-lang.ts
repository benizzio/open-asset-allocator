import * as handlebars from "handlebars";
import { HelperOptions } from "handlebars";

// TODO generalize and componentize generic utility functions from here

/**
 * Internal store that keeps iterator maps per template render root using WeakMap.
 * The state is scoped to the lifetime of a single template rendering (data.root).
 *
 * Authored by: GitHub Copilot
 */
const RENDER_ITERATORS_STORE: WeakMap<object, Map<string, number>> = new WeakMap();

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
 * Narrow unknown to Handlebars HelperOptions.
 *
 * @param value - Unknown value to test.
 * @returns True if value looks like HelperOptions.
 *
 * Authored by: GitHub Copilot
 */
function isHelperOptions(value: unknown): value is HelperOptions {

    if(typeof value !== "object" || value === null) {
        return false;
    }

    const record = value as Record<string, unknown>;
    return (
        Object.prototype.hasOwnProperty.call(record, "data") &&
        Object.prototype.hasOwnProperty.call(record, "hash")
    );
}

/**
 * Gets (or creates) the iterator map for the current template rendering.
 *
 * @param options - Handlebars helper options (must contain data.root).
 * @returns The iterator map bound to the current render root.
 *
 * Authored by: GitHub Copilot
 */
function getRenderIteratorMap(options?: HelperOptions): Map<string, number> {

    const dataUnknown = options?.data as unknown;

    let root: object | undefined = undefined;

    if(typeof dataUnknown === "object" && dataUnknown !== null) {
        const candidate = (dataUnknown as Record<string, unknown>)["root"];

        if(typeof candidate === "object" && candidate !== null) {
            root = candidate as object;
        }
    }

    if(!root) {
        // As a last resort, use a unique per-call object to avoid leaking across renders.
        // Note: without data.root, iterators won't share state between helper calls.
        const isolatedRoot = {} as object;
        const isolated = new Map<string, number>();
        RENDER_ITERATORS_STORE.set(isolatedRoot, isolated);
        return isolated;
    }

    let map = RENDER_ITERATORS_STORE.get(root);

    if(!map) {
        map = new Map<string, number>();
        RENDER_ITERATORS_STORE.set(root, map);
    }

    return map;
}

/**
 * Initializes (or resets) a named iterator for the current template rendering.
 * The iterator starts at the provided initial value; the first {{iteratorNext id}} call
 * will yield that initial value, then increment by 1 for subsequent calls.
 * If no initial value is supplied, it starts at 0.
 *
 * @param id - Unique iterator id within this template rendering.
 * @param startOrOptions - Initial value (number-like) or the Handlebars options when omitted.
 * @param maybeOptions - The Handlebars options object when start is provided.
 * @returns An empty string (no output); use {{iteratorNext id}} to consume values.
 *
 * @example
 * {{iteratorInit "row"}}        {{!-- starts at 0 --}}
 * {{iteratorInit "row" 10}}     {{!-- starts at 10 --}}
 * {{iteratorNext "row"}}        {{!-- 10 --}}
 *
 * @author GitHub Copilot
 */
function iteratorInitHelper(
    this: unknown,
    id: unknown,
    startOrOptions?: unknown,
    maybeOptions?: HelperOptions,
): string {

    // Determine options and start value allowing {{iteratorInit id}} and {{iteratorInit id start}}
    let options: HelperOptions | undefined = undefined;
    let startValue = 0;

    const maybeOptionsIsHelper = isHelperOptions(maybeOptions);
    const startOrOptionsIsHelper = isHelperOptions(startOrOptions);

    if(maybeOptionsIsHelper) {

        options = maybeOptions as HelperOptions;

        if(!startOrOptionsIsHelper && typeof startOrOptions !== "undefined") {
            startValue = coerceToFiniteNumber(startOrOptions);
        }
    }
    else if(startOrOptionsIsHelper) {
        options = startOrOptions as HelperOptions;
        startValue = 0;
    }
    else {
        // Fallback: no options detected (should not happen in normal Handlebars usage)
        startValue = typeof startOrOptions !== "undefined" ? coerceToFiniteNumber(startOrOptions) : 0;
    }

    const map = getRenderIteratorMap(options);
    const key = String(id);

    map.set(key, startValue);

    return "";
}

/**
 * Returns the current value for the named iterator and advances it by 1.
 * If the iterator doesn't exist yet, it is implicitly initialized at 0.
 *
 * @param id - Iterator id.
 * @param options - Handlebars helper options providing data.root for render scoping.
 * @returns The iterator's current value (number).
 *
 * @example
 * {{iteratorInit "seq" 3}}
 * {{iteratorNext "seq"}}  {{!-- 3 --}}
 * {{iteratorNext "seq"}}  {{!-- 4 --}}
 *
 * @author GitHub Copilot
 */
function iteratorNextHelper(this: unknown, id: unknown, options: HelperOptions): number {

    const map = getRenderIteratorMap(options);
    const key = String(id);

    const current = map.has(key) ? (map.get(key) as number) : 0;

    map.set(key, current + 1);

    return current;
}

/**
 * Splits a string value into an array using a separator.
 *
 * Non-string values are coerced via String(). If the separator is omitted it defaults to ','.
 * When the provided value is null/undefined or an empty string, an empty array is returned.
 *
 * This helper gracefully handles the Handlebars options object being passed as the second
 * argument when the separator is omitted (e.g. {{split value}}).
 *
 * @param value - Source value to split (coerced to string when not null/undefined).
 * @param separatorOrOptions - Separator string (optional) or implicitly the Handlebars options object when omitted.
 * @param maybeOptions - Handlebars options object when both value and separator are provided.
 * @returns Array of substrings produced by String(value).split(separator).
 *
 * @example
 * {{#each (split "a,b,c" ",")}}
 *   {{this}}
 * {{/each}}            {{!-- yields: a b c --}}
 *
 * @example
 * {{#each (split csv)}}            {{!-- defaults to comma separator --}}
 *   {{this}}
 * {{/each}}
 *
 * @author GitHub Copilot
 */
function splitHelper(
    value: unknown,
    separatorOrOptions?: unknown,
    maybeOptions?: HelperOptions,
): string[] {

    if(value === null || typeof value === "undefined") {
        return [];
    }

    const str = String(value);

    if(str.length === 0) {
        return [];
    }

    let separator = ","; // default separator

    const maybeOptionsIsHelper = isHelperOptions(maybeOptions);
    const separatorOrOptionsIsHelper = isHelperOptions(separatorOrOptions);

    // Determine if the second argument is actually the options object.
    if(maybeOptionsIsHelper) {
        // Both value and separator provided.
        if(!separatorOrOptionsIsHelper && typeof separatorOrOptions !== "undefined") {
            separator = String(separatorOrOptions);
        }
    }
    else if(separatorOrOptionsIsHelper) {
        // Only value provided; keep default separator.
    }
    else if(typeof separatorOrOptions !== "undefined") {
        // Separator provided without options (unlikely in normal Handlebars usage, but supported).
        separator = String(separatorOrOptions);
    }

    return str.split(separator);
}

/**
 * Determines whether a value is a finite number primitive.
 *
 * @param value - Value to test.
 * @returns True when value is a number and Number.isFinite(value) is true.
 *
 * @author GitHub Copilot
 */
function isFiniteNumberValue(value: unknown): value is number {
    return typeof value === "number" && Number.isFinite(value);
}

/**
 * Attempts to coerce an arbitrary value into a finite number.
 * Returns an object indicating success with the coerced number or failure.
 *
 * Rules:
 * - Numbers: accepted if finite.
 * - Strings: trimmed; empty strings are rejected; numeric strings must coerce to finite numbers.
 * - Dates: coerced using getTime().
 * - Other values: coerced via Number(); symbols will throw and be treated as failure.
 *
 * @param value - The input value to try to coerce to a finite number.
 * @returns { ok: true, value: number } on success; { ok: false } on failure.
 *
 * @author GitHub Copilot
 */
function tryCoerceToFiniteNumber(value: unknown): { ok: true; value: number } | { ok: false } {

    if(isFiniteNumberValue(value)) {
        return { ok: true, value: value };
    }

    if(typeof value === "string") {
        const trimmed = value.trim();

        if(trimmed.length === 0) {
            return { ok: false };
        }

        const n = Number(trimmed);
        return Number.isFinite(n) ? { ok: true, value: n } : { ok: false };
    }

    if(value instanceof Date) {
        const n = value.getTime();
        return Number.isFinite(n) ? { ok: true, value: n } : { ok: false };
    }

    try {
        const n = Number(value as never);
        return Number.isFinite(n) ? { ok: true, value: n } : { ok: false };
    } catch {
        return { ok: false };
    }
}

/**
 * Builds a deterministic string representation for a value used in fallback comparison.
 *
 * Rules:
 * - null/undefined: String(value)
 * - string: returned as-is
 * - symbol: symbol.toString()
 * - objects/arrays: JSON with sorted keys when possible; falls back to String(value) on failure
 * - others: String(value)
 *
 * @param value - Value to convert to a stable comparable string.
 * @returns Deterministic string representation used for lexicographic comparison.
 *
 * @author GitHub Copilot
 */
function toComparatorString(value: unknown): string {

    if(value === null || typeof value === "undefined") {
        return String(value);
    }

    if(typeof value === "string") {
        return value;
    }

    if(typeof value === "symbol") {
        return value.toString();
    }

    if(typeof value === "object") {
        try {
            const obj = value as Record<string, unknown>;
            const keys = Object.keys(obj).sort();
            const json = JSON.stringify(obj, keys);

            if(json !== undefined) {
                return json as string;
            }
        } catch {
            // ignore and fallback
        }
    }

    return String(value);
}

/**
 * Compares two values prioritizing numeric comparison when both values are number-like.
 * Internal helper logic extracted to reduce cognitive complexity.
 *
 * @param a - Left value for comparison.
 * @param b - Right value for comparison.
 * @returns -1 if a < b, 1 if a > b, 0 if equal under the comparator rules.
 *
 * @example
 * <pre>{{comparator 1 2}}</pre>                <!-- -1 (numeric) -->
 * <pre>{{comparator "10" "2"}}</pre>         <!-- 1 when numeric (10 > 2) -->
 * <pre>{{comparator "x10" "x2"}}</pre>       <!-- string compare when not numeric -->
 *
 * @author GitHub Copilot
 */
function comparatorHelper(a: unknown, b: unknown): number {

    const leftNumberCoercion = tryCoerceToFiniteNumber(a);
    const rightNumberCoercion = tryCoerceToFiniteNumber(b);

    if(leftNumberCoercion.ok && rightNumberCoercion.ok) {
        if(leftNumberCoercion.value === rightNumberCoercion.value) {
            return 0;
        }
        return leftNumberCoercion.value < rightNumberCoercion.value ? -1 : 1;
    }

    const leftComparableKey = toComparatorString(a);
    const rightComparableKey = toComparatorString(b);

    if(leftComparableKey === rightComparableKey) {
        return 0;
    }
    return leftComparableKey < rightComparableKey ? -1 : 1;
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
    handlebars.registerHelper("iteratorInit", iteratorInitHelper);
    handlebars.registerHelper("iteratorNext", iteratorNextHelper);
    handlebars.registerHelper("split", splitHelper);
    handlebars.registerHelper("comparator", comparatorHelper);
}