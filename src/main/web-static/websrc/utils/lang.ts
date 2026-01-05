import BigNumber from "bignumber.js";

/**
 * Generic language utilities for internal use across the websrc codebase.
 *
 * These helpers are framework-agnostic and safe to reuse anywhere.
 *
 * Public Usage Guidance:
 * - All exported functions are pure (no side effects) except assignValueAtPath (mutates target object)
 * - Defensive coercion helpers return neutral defaults (0 or empty arrays) instead of throwing
 * - Prefer tryCoerceToFiniteNumber when you need success/failure signalling over silent coercion
 * - toInt defaults to coercions disabled; pass { allowCoercions: true } to enable explicit coercions
 *   (useful for tolerant parsing in UI forms), and keep the default for strict validation layers.
 *
 * Authoring Guidance:
 * - When adding new utilities keep them single-purpose and side-effect free
 * - Provide at least one concrete @example covering success and an edge case
 * - If a function performs mutation clearly document it and highlight expected invariants
 *
 * Authored by: GitHub Copilot
 */

/**
 * Converts an arbitrary property path specification into string segments.
 *
 * Rules:
 * - string: split on '.' and filter out empty segments
 * - number, boolean, bigint, symbol: coerced to string as a single segment
 * - others (null/undefined/object/function): yields an empty array (ignored)
 *
 * Edge Cases:
 * - Empty string => []
 * - Numeric values preserve exact string form (e.g. 0 => ["0"], -1 => ["-1"])
 * - Symbols convert via String(symbol)
 *
 * @param path - Raw path value provided by template or runtime usage.
 * @returns Array of path segments (possibly empty).
 *
 * @example
 * toPropertyPathSegments("a.b.c") //=> ["a","b","c"]
 * toPropertyPathSegments(10) //=> ["10"]
 * toPropertyPathSegments({}) //=> []
 *
 * @author GitHub Copilot
 */
export function toPropertyPathSegments(path: unknown): string[] {

    if(typeof path === "string") {
        if(path.trim().length === 0) {
            return [];
        }
        return path.split(".").filter(p => p.length > 0);
    }

    if(typeof path === "number" || typeof path === "boolean" || typeof path === "bigint" || typeof path === "symbol") {
        return [String(path)];
    }

    return [];
}

/**
 * Options for assignValueAtPath to provide decoupled logging or introspection.
 *
 * All callbacks are optional and will only be called when relevant.
 *
 * @property onWarn - Called when overwriting a non-object intermediate segment.
 * @property onError - Called when an unexpected mutation failure occurs.
 *
 * @author GitHub Copilot
 */
export type AssignAtPathOptions = {
    onWarn?: (message: string, details?: Record<string, unknown>) => void;
    onError?: (message: string, details?: Record<string, unknown>) => void;
};

/**
 * Assigns a value into a target object at the provided path segments, creating
 * intermediate plain objects when necessary.
 *
 * The function can report noteworthy events via the provided callbacks.
 *
 * Contract:
 * - Inputs: target (object), segments (non-empty array), value (any)
 * - Side effects: mutates target
 * - Returns: true on success; false when an error occurs during assignment
 * - Empty segments array: returns false and triggers onError (guard added)
 *
 * Mutation Invariants:
 * - Existing non-object intermediate values are replaced with new objects (warning emitted)
 * - Target is only mutated along the specified segment chain
 *
 * @param target - Target object to mutate.
 * @param segments - Non-empty array of property path segments.
 * @param value - Value to assign at the final segment.
 * @param options - Optional callbacks for warnings and errors.
 * @returns True when assignment was successful; otherwise false.
 *
 * @example
 * const obj: Record<string, unknown> = {};
 * assignValueAtPath(obj, ["a", "b", "c"], 10); // obj => { a: { b: { c: 10 } } }
 * const obj2: Record<string, unknown> = { a: 1 };
 * assignValueAtPath(obj2, ["a", "b"], 5, { onWarn: console.warn }); // warns overwrite
 * assignValueAtPath({}, [], 5) // => false (empty path)
 *
 * @author GitHub Copilot
 */
export function assignValueAtPath(
    target: Record<string, unknown>,
    segments: string[],
    value: unknown,
    options?: AssignAtPathOptions,
): boolean {

    if(segments.length === 0) {
        options?.onError?.("Empty path segments array.", { value });
        return false;
    }

    let cursor: Record<string, unknown> = target;

    for(let i = 0; i < segments.length - 1; i++) {
        const segment = segments[i];
        const existing = cursor[segment];

        if(typeof existing === "object" && existing !== null) {
            cursor = existing as Record<string, unknown>;
            continue;
        }

        if(typeof existing !== "undefined") {
            options?.onWarn?.(
                "Overwriting non-object intermediate path segment with a new object.",
                {
                    segment,
                    previousType: typeof existing,
                    path: segments.join("."),
                    value,
                },
            );
        }

        const next: Record<string, unknown> = {};

        try {
            cursor[segment] = next;
        } catch(err) {
            options?.onError?.(
                "Failed to create intermediate object during assignment.",
                { segment, error: err as unknown },
            );
            return false;
        }

        cursor = next;
    }

    const finalKey = segments[segments.length - 1];

    try {
        cursor[finalKey] = value;
        return true;
    } catch(err) {
        options?.onError?.(
            "Failed to assign value.",
            { key: finalKey, error: err as unknown, value },
        );
        return false;
    }
}

/**
 * Determines whether a value is a finite number primitive.
 *
 * @param value - Value to test.
 * @returns True when value is a number and Number.isFinite(value) is true.
 *
 * @example
 * isFiniteNumberValue(10) // => true
 * isFiniteNumberValue(Infinity) // => false
 * isFiniteNumberValue("10") // => false
 *
 * @author GitHub Copilot
 */
export function isFiniteNumberValue(value: unknown): value is number {
    return typeof value === "number" && Number.isFinite(value);
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
 * @param value - Arbitrary input value to coerce.
 * @returns Finite number result; 0 when the value cannot be coerced to a finite number.
 *
 * @example
 * coerceToFiniteNumber(10) // => 10
 * coerceToFiniteNumber(Infinity) // => 0
 * coerceToFiniteNumber(" 42 ") // => 42
 * coerceToFiniteNumber("oops") // => 0
 *
 * @author GitHub Copilot
 */
export function coerceToFiniteNumber(value: unknown): number {

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
 * @returns An object representing the coercion outcome:
 *   - { ok: true, value: number } when coercion succeeds
 *   - { ok: false } when coercion fails
 *
 * @example
 * tryCoerceToFiniteNumber(15) // => { ok: true, value: 15 }
 * tryCoerceToFiniteNumber("15") // => { ok: true, value: 15 }
 * tryCoerceToFiniteNumber("not") // => { ok: false }
 * tryCoerceToFiniteNumber("   ") // => { ok: false }
 *
 * @author GitHub Copilot
 */
export function tryCoerceToFiniteNumber(value: unknown): { ok: true; value: number } | { ok: false } {

    if(isFiniteNumberValue(value)) {
        return { ok: true, value: value } as const;
    }

    if(typeof value === "string") {
        const trimmed = value.trim();

        if(trimmed.length === 0) {
            return { ok: false } as const;
        }

        const n = Number(trimmed);
        return Number.isFinite(n) ? ({ ok: true, value: n } as const) : ({ ok: false } as const);
    }

    if(value instanceof Date) {
        const n = value.getTime();
        return Number.isFinite(n) ? ({ ok: true, value: n } as const) : ({ ok: false } as const);
    }

    try {
        const n = Number(value as never);
        return Number.isFinite(n) ? ({ ok: true, value: n } as const) : ({ ok: false } as const);
    } catch {
        return { ok: false } as const;
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
 * @example
 * toComparableString({ b: 2, a: 1 }) // => '{"a":1,"b":2}'
 * toComparableString(Symbol("x")) // => 'Symbol(x)'
 * toComparableString(undefined) // => 'undefined'
 *
 * @author GitHub Copilot
 */
export function toComparableString(value: unknown): string {

    if(isNullish(value)) {
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
 * Options controlling integer conversion behavior for toInt.
 *
 * @property {boolean} [allowCoercions] When true, apply all documented coercion rules; when false, reject coercions.
 * @property {(message: string, details?: Record<string, unknown>) => void} [onCoercion]
 *   Callback invoked when a coercion is blocked (non-coercion mode) or needs reporting.
 *
 * Defaults:
 * - allowCoercions: false (non-coercion/strict mode by default)
 *
 * Usage Examples:
 * @example
 * // Default (coercions disabled): non-number values are rejected with undefined
 * toInt("42") // => undefined
 * toInt(true) // => undefined
 *
 * // Enable coercions explicitly when intended
 * toInt("42", { allowCoercions: true }) // => 42
 * toInt(42.5, { allowCoercions: false, onCoercion: console.error }) // => undefined
 * toInt("x", { allowCoercions: false, onCoercion: (m, d) => console.log(m, d) }) // => undefined
 *
 * @author GitHub Copilot
 */
export type ToIntOptions = {
    allowCoercions?: boolean;
    onCoercion?: (message: string, details?: Record<string, unknown>) => void;
};

/**
 * Internal helper that reports a blocked coercion attempt for toInt.
 *
 * Logs using the provided onCoercion callback or falls back to console.error.
 * Returns undefined consistently to streamline caller usage.
 *
 * @param options - ToIntOptions passed to toInt (can be undefined).
 * @param reason - Human readable reason why coercion was blocked.
 * @param original - Original value provided to toInt.
 * @returns Always undefined.
 *
 * @author GitHub Copilot
 */
function reportToIntCoercion(
    options: ToIntOptions | undefined,
    reason: string,
    original: unknown,
): undefined {

    const message = `toInt coercion blocked: ${ reason }`;
    const details = { value: original };

    if(options?.onCoercion) {
        options.onCoercion(message, details);
    }
    else {
        // Fallback: use console.error to highlight coercion problems that could indicate data quality issues.
        // Authored by: GitHub Copilot
        console.error(message, details);
    }

    return undefined;
}

/**
 * Handles number conversion logic for toInt, applying coercion rules or logging.
 *
 * @param value - Number input.
 * @param options - Original toInt options for logging callback.
 * @returns Converted integer number or undefined when coercion blocked.
 *
 * @example
 * convertNumberForToInt(10, { allowCoercions: false }) // => 10
 * convertNumberForToInt(10.9, { allowCoercions: true }) // => 10
 * convertNumberForToInt(10.9, { allowCoercions: false }) // => undefined
 * convertNumberForToInt(Infinity, { allowCoercions: true }) // => 0
 *
 * @author GitHub Copilot
 */
function convertNumberForToInt(
    value: number,
    options?: ToIntOptions,
): number | undefined {

    const allowCoercions = options?.allowCoercions === true;

    if(!Number.isFinite(value)) {
        return allowCoercions ? 0 : reportToIntCoercion(options, "non-finite number", value);
    }

    if(Number.isInteger(value)) {
        return value;
    }

    return allowCoercions
        ? Math.trunc(value)
        : reportToIntCoercion(options, "non-integer finite number requires truncation", value);
}

/**
 * Converts an arbitrary value to an integer number with optional coercion control.
 *
 * Coercion Rules (when allowCoercions === true; default is false):
 * - number (finite): truncated to integer via Math.trunc
 * - number (NaN, +Infinity, -Infinity): coerced to 0 for safety and consistency
 * - string: parsed as base-10 integer; non-numeric yields 0
 * - bigint: converted to number (may lose precision)
 * - boolean: true => 1, false => 0
 * - others (null/undefined/object/symbol): 0
 *
 * Non-Coercion Mode (allowCoercions !== true) [DEFAULT]:
 * - Accepts only finite integer number values (e.g. 5, 0, -12)
 * - Returns undefined for:
 *   - Non-integer finite numbers (would require truncation)
 *   - Non-finite numbers (NaN / ±Infinity)
 *   - Any non-number type (string, bigint, boolean, object, etc.)
 *   - Each rejected coercion attempt triggers logging via onCoercion or console.error
 *
 * @param value - The input value to convert.
 * @param options - Optional ToIntOptions controlling coercions and logging.
 * @returns Integer number or undefined when coercions are disallowed.
 *
 * @example
 * // Defaults to strict (non-coercion) mode
 * toInt(42) // => 42
 * toInt(42.9) // => undefined
 * toInt("42") // => 42
 * toInt(Infinity) // => undefined
 * toInt(true) // => undefined
 *
 * // Coercion-enabled examples
 * toInt(42.9, { allowCoercions: true }) // => 42
 * toInt(Infinity, { allowCoercions: true }) // => 0
 * toInt("aaa", { allowCoercions: true }) // => 0
 * toInt(true, { allowCoercions: true }) // => 1
 *
 * @author GitHub Copilot
 * @author benizzio
 */
export function toInt(value: unknown, options?: ToIntOptions): number | undefined {

    const allowCoercions = options?.allowCoercions === true;

    if(typeof value === "number") {
        return convertNumberForToInt(value, options);
    }

    switch(typeof value) {
        case "string": {
            const n = parseInt(value, 10);
            const isNotANumber = Number.isNaN(n);

            if(isNotANumber && !allowCoercions) {
                return reportToIntCoercion(options, "string is NaN", value);
            }

            return isNotANumber ? 0 : n;
        }

        case "bigint": {
            return allowCoercions
                ? Number(value)
                : reportToIntCoercion(options, "bigint conversion disallowed", value);
        }

        case "boolean": {
            return allowCoercions
                ? (value ? 1 : 0)
                : reportToIntCoercion(options, "boolean conversion disallowed", value);
        }

        default: {
            return allowCoercions
                ? 0
                : reportToIntCoercion(options, "unsupported type conversion disallowed", value);
        }
    }
}

/**
 * Determines whether a value is null or undefined (JavaScript "nullish").
 *
 * @param value - Arbitrary value to test.
 * @returns True when the value is exactly null or undefined.
 *
 * @example
 * isNullish(null) // => true
 * isNullish(undefined) // => true
 * isNullish(0) // => false
 * isNullish("") // => false
 * isNullish(false) // => false
 *
 * @example
 * // Safe defaulting pattern without touching other falsy values:
 * const height = maybeHeightValue;
 * const resolved = isNullish(height) ? 100 : height;
 *
 * @author GitHub Copilot
 */
export function isNullish(value: unknown): value is null | undefined {
    // Loose equality is intentional: only matches null or undefined.
    return value == null;
}

/**
 * Safely coerces a value to a BigNumber instance.
 *
 * This function handles all edge cases including null, undefined, invalid types,
 * and values that cannot be converted to valid numbers. It uses try-catch to
 * handle BigNumber constructor errors and checks for NaN results.
 *
 * Coercion Rules:
 * - Valid numeric values (numbers, numeric strings, etc.): converted to BigNumber
 * - null or undefined: coalesced to 0, then converted to BigNumber(0)
 * - Invalid inputs (objects, functions, NaN results): returns BigNumber(0)
 * - BigNumber constructor errors: caught and returns BigNumber(0)
 *
 * @param value - Arbitrary input value to coerce.
 * @returns A valid BigNumber instance; BigNumber(0) when the value cannot be coerced.
 *
 * @example
 * coerceToBigNumber(10) // => BigNumber(10)
 * coerceToBigNumber("0.00111") // => BigNumber(0.00111)
 * coerceToBigNumber(null) // => BigNumber(0)
 * coerceToBigNumber(undefined) // => BigNumber(0)
 * coerceToBigNumber({}) // => BigNumber(0)
 * coerceToBigNumber("invalid") // => BigNumber(0)
 *
 * @author GitHub Copilot
 */
export function coerceToBigNumber(value: unknown): BigNumber {

    try {
        // Convert unknown value to BigNumber.Value type before passing to constructor
        const coercedValue: BigNumber.Value = (typeof value === "string" ||
            typeof value === "number" ||
            typeof value === "bigint" ||
            value instanceof BigNumber)
            ? value as BigNumber.Value
            : 0;

        const bn = new BigNumber(coercedValue);

        if (bn.isNaN()) {
            return new BigNumber(0);
        }

        return bn;
    } catch {
        return new BigNumber(0);
    }
}