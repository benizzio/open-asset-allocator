import * as handlebars from "handlebars";
import { HelperOptions } from "handlebars";
import { coerceToFiniteNumber } from "../../utils/lang";

/**
 * Internal store that keeps iterator maps per template render root using WeakMap.
 * The state is scoped to the lifetime of a single template rendering (data.root).
 *
 * Authored by: GitHub Copilot
 */
const RENDER_ITERATORS_STORE: WeakMap<object, Map<string, number>> = new WeakMap();

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
 * Narrow unknown to Handlebars HelperOptions.
 *
 * @param value - Unknown value to test.
 * @returns True if value looks like HelperOptions.
 *
 * Authored by: GitHub Copilot
 */
export function isHelperOptions(value: unknown): value is HelperOptions {

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


export function registerHandlebarsUtilHelpers() {

    // Register all helpers with their names
    handlebars.registerHelper("repeater", repeaterHelper);
    handlebars.registerHelper("iteratorInit", iteratorInitHelper);
    handlebars.registerHelper("iteratorNext", iteratorNextHelper);
}