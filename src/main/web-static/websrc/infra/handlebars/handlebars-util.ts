import * as handlebars from "handlebars";

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

export function registerHandlebarsUtilHelpers() {

    // Register all helpers with their names
    handlebars.registerHelper("repeater", repeaterHelper);
}