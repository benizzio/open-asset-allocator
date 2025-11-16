import * as handlebars from "handlebars";
import { HelperOptions } from "handlebars";
import { logger, LogLevel } from "../logging";
import {
    assignValueAtPath,
    coerceToFiniteNumber,
    toCompararableString,
    toPropertyPathSegments,
    tryCoerceToFiniteNumber,
} from "../../utils/lang";


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
 * Sets (mutates) a property on a target object, supporting dot-separated nested paths.
 * Intermediate objects are created when missing. Non-object targets are ignored.
 * Returns the assigned value to allow inline usage inside expressions.
 *
 * Usage patterns:
 *   {{setProperty obj "name" value}}            -> sets obj.name = value
 *   {{setProperty obj "stats.count" 10}}       -> creates stats if missing and sets count
 *   {{setProperty obj dynamicPath value}}       -> path can be computed
 *
 * Edge cases & safety:
 * - When target is null/undefined or not an object, no mutation occurs (returns value).
 * - Empty or invalid path produces no mutation (returns value).
 * - Existing non-object intermediate segment is replaced with a plain object to continue path creation.
 * - Does not attempt array index semantics; numeric segments become plain object keys.
 *
 * @param target - Target object to mutate.
 * @param path - Property key or dot-separated path string.
 * @param value - Value to assign at the path.
 * @param _options - Handlebars helper options (unused but required by helper signature).
 * @returns The empty string (no direct output) to avoid interfering with template rendering.
 *
 * @example
 * {{setProperty user "firstName" "Alice"}}
 * {{setProperty user "profile.details.age" 42}}
 * <pre>{{stringify user}}</pre>
 *
 * @author GitHub Copilot
 */
function setPropertyHelper(
    target: unknown,
    path: unknown,
    value: unknown,
    _options: HelperOptions, // underscore to indicate intentionally unused per Handlebars signature
): string { // returns empty string to avoid interfering with template output

    // Explicit no-op usage to satisfy eslint unused variable rule while documenting intention.
    void _options;

    if(typeof target !== "object" || target === null) {
        logger(
            LogLevel.ERROR,
            "[handlebars:setProperty] Target is not an object. Skipping assignment.",
            { targetType: typeof target, path, value },
        );
        return "";
    }

    const segments = toPropertyPathSegments(path);

    if(segments.length === 0) {
        logger(
            LogLevel.ERROR,
            "[handlebars:setProperty] Empty or invalid path. Skipping assignment.",
            { path, value },
        );
        return ""; // nothing to assign
    }

    // Delegate the traversal and assignment to reduce cognitive complexity
    assignValueAtPath(
        target as Record<string, unknown>,
        segments,
        value,
        {
            onWarn: (message, details) =>
                logger(LogLevel.WARN, `[handlebars:setProperty] ${ message }`, details ?? {}),
            onError: (message, details) =>
                logger(LogLevel.ERROR, `[handlebars:setProperty] ${ message }`, details ?? {}),
        },
    );

    return ""; // always empty output
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
 * Renders the block when the provided value is null or undefined; otherwise renders the inverse block.
 *
 * This helper matches JavaScript "nullish" semantics (value === null || value === undefined).
 *
 * @this unknown
 * @param value - Value to test for nullishness.
 * @param options - Handlebars block options providing `fn`/`inverse`.
 * @returns The rendered block depending on whether the value is nullish.
 *
 * @example
 * {{#ifNullish user.middleName}}
 *   <span class="muted">N/A</span>
 * {{else}}
 *   {{user.middleName}}
 * {{/ifNullish}}
 *
 * @author GitHub Copilot
 */
function ifNullishHelper(this: unknown, value: unknown, options: HelperOptions): string {
    return (value === null || typeof value === "undefined")
        ? options.fn(this)
        : options.inverse(this);
}

/**
 * Renders the block when the provided value is NOT null or undefined; otherwise renders the inverse block.
 *
 * This helper matches the inverted JavaScript nullish check: value !== null && value !== undefined.
 *
 * @this unknown
 * @param value - Value to test for not-nullishness.
 * @param options - Handlebars block options providing `fn`/`inverse`.
 * @returns The rendered block depending on whether the value is not nullish.
 *
 * @example
 * {{#ifNotNullish user.middleName}}
 *   {{user.middleName}}
 * {{else}}
 *   <span class="muted">N/A</span>
 * {{/notNullish}}
 *
 * @author GitHub Copilot
 */
function ifNotNullishHelper(this: unknown, value: unknown, options: HelperOptions): string {
    return (value !== null && typeof value !== "undefined")
        ? options.fn(this)
        : options.inverse(this);
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

    const leftComparableKey = toCompararableString(a);
    const rightComparableKey = toCompararableString(b);

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
    handlebars.registerHelper("setProperty", setPropertyHelper); // newly added helper
    handlebars.registerHelper("ifEquals", ifEqualsHelper);
    handlebars.registerHelper("ifNotEquals", ifNotEqualsHelper);
    handlebars.registerHelper("ifNullish", ifNullishHelper);
    handlebars.registerHelper("ifNotNullish", ifNotNullishHelper);
    handlebars.registerHelper("math", mathHelper);
    handlebars.registerHelper("iteratorInit", iteratorInitHelper);
    handlebars.registerHelper("iteratorNext", iteratorNextHelper);
    handlebars.registerHelper("split", splitHelper);
    handlebars.registerHelper("comparator", comparatorHelper);
}