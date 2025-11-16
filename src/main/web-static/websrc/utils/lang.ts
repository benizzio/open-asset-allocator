/**
 * Generic language utilities for internal use across the websrc codebase.
 *
 * These helpers are framework-agnostic and safe to reuse anywhere.
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
 * @param path - Raw path value provided by template or runtime usage.
 * @returns Array of path segments (possibly empty).
 *
 * @example
 * toPropertyPathSegments("a.b.c") //=> ["a","b","c"]
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
 *
 * @param target - Target object to mutate.
 * @param segments - Non-empty array of property path segments.
 * @param value - Value to assign at the final segment.
 * @param options - Optional callbacks for warnings and errors.
 * @returns True when assignment was successful; otherwise false.
 *
 * @example
 * const obj: Record<string, unknown> = {};
 * assignValueAtPath(obj, ["a", "b", "c"], 10);
 * // obj is now: { a: { b: { c: 10 } } }
 *
 * @author GitHub Copilot
 */
export function assignValueAtPath(
    target: Record<string, unknown>,
    segments: string[],
    value: unknown,
    options?: AssignAtPathOptions,
): boolean {

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
 * @returns { ok: true, value: number } on success; { ok: false } on failure.
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
 * @author GitHub Copilot
 */
export function toCompararableString(value: unknown): string {

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
