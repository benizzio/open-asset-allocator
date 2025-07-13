import * as handlebars from "handlebars";

/**
 * Registers custom Handlebars helpers that extend language functionality for template rendering.
 *
 * Usage:
 *   Call this function during application initialization to register all custom helpers.
 *   Example: registerHandlebarsLangHelpers();
 *
 * Documentation authored by: GitHub Copilot
 */
export function registerHandlebarsLangHelpers() {

    /**
     * Creates an object from the provided hash parameters.
     *
     * This helper allows creating objects dynamically within templates using key-value pairs.
     *
     * Usage in templates:
     *   {{#with (object key1="value1" key2="value2")}}
     *     {{key1}} - {{key2}}
     *   {{/with}}
     *
     * Parameters:
     *   - hash: Object containing key-value pairs passed as named parameters
     *
     * Returns:
     *   - Object with the provided key-value pairs
     *
     * Documentation authored by: GitHub Copilot
     */
    handlebars.registerHelper("object", ({ hash }) => {
        return hash;
    });

    /**
     * Creates an array from the provided arguments.
     *
     * This helper converts multiple arguments into a JavaScript array, excluding the Handlebars
     * options object that is automatically appended as the last argument.
     *
     * Usage in templates:
     *   {{#each (array "item1" "item2" "item3")}}
     *     <li>{{this}}</li>
     *   {{/each}}
     *
     * Parameters:
     *   - ...args: Variable number of arguments to be converted into an array
     *
     * Returns:
     *   - Array containing all provided arguments (excluding the Handlebars options object)
     *
     * Documentation authored by: GitHub Copilot
     */
    handlebars.registerHelper("array", function(...args) {
        return Array.from(args).slice(0, arguments.length - 1);
    });

    /**
     * Generates a script tag containing JSON data for DOM manipulation.
     *
     * This helper creates a script element with type "application/json" containing the serialized
     * object data. This is commonly used for passing server-side data to client-side JavaScript
     * in a safe and structured way.
     *
     * Usage in templates:
     *   {{{ domJSON "portfolio" this }}}
     *
     * Parameters:
     *   - id: Unique identifier for the script element
     *   - object: JavaScript object to be serialized as JSON
     *
     * Returns:
     *   - HTML string containing a script element with the JSON data
     *
     * Documentation authored by: GitHub Copilot
     */
    handlebars.registerHelper("domJSON", function(id: string, object: object) {
        return `
            <script id="${ id }" type="application/json">${ JSON.stringify(object) }</script>
        `;
    });

    /**
     * Repeats a text string a specified number of times with optional prefix and suffix.
     *
     * This helper is useful for creating indentation, spacing, or repetitive content in templates.
     * Returns an empty string if count is zero or negative.
     *
     * Usage in templates:
     *   {{{repeater "&nbsp;&nbsp;&nbsp;&nbsp;" depth "" ""}}}
     *   {{{repeater "★" rating "Rating: " "/5"}}}
     *
     * Parameters:
     *   - text: The string to be repeated
     *   - count: Number of times to repeat the text (must be > 0)
     *   - prefix: Optional string to prepend to the result
     *   - suffix: Optional string to append to the result
     *
     * Returns:
     *   - String with the repeated text, optionally wrapped with prefix and suffix
     *   - Empty string if count <= 0
     *
     * Documentation authored by: GitHub Copilot
     */
    handlebars.registerHelper(
        "repeater",
        function(text: string | number, count: number, prefix: string, suffix: string) {

            if(count <= 0) {
                return "";
            }

            let result = String(text).repeat(count);
            result = prefix ? prefix + result : result;
            result = suffix ? result + suffix : result;

            return result;
        },
    );

    /**
     * Converts a JavaScript object to its JSON string representation.
     *
     * This helper provides a simple way to serialize objects as JSON strings within templates.
     * The output is safe for use in HTML attributes and content.
     *
     * Usage in templates:
     *   <pre>{{stringify userData}}</pre>
     *   <input type="hidden" value="{{stringify formData}}">
     *
     * Parameters:
     *   - object: JavaScript object to be converted to JSON string
     *
     * Returns:
     *   - JSON string representation of the object
     *
     * Documentation authored by: GitHub Copilot
     */
    handlebars.registerHelper("stringify", function(object: object) {
        return JSON.stringify(object);
    });

    /**
     * Concatenates multiple string arguments into a single string.
     *
     * Usage in templates:
     *   <div class="{{concat baseClass "-" modifier}}">
     *   <a href="{{concat protocol "://" domain path}}">
     *
     * Parameters:
     *   - ...args: Variable number of arguments to be concatenated
     *
     * Returns:
     *   - Single string containing all arguments joined together
     *
     * Documentation authored by: GitHub Copilot
     */
    handlebars.registerHelper("concat", function(...args) {
        return Array.from(args).slice(0, arguments.length - 1).join("");
    });

    /**
     * Retrieves a property value from an object by its key.
     *
     * Usage in templates:
     *  {{getProperty user "name"}}
     *  {{getProperty settings "theme"}}
     *
     *  Parameters:
     *  - obj: The object from which to retrieve the property
     *  - key: The key of the property to retrieve
     *
     *  Returns:
     *  - The value of the specified property, or undefined if the property does not exist
     *
     *  Authored by: GitHub Copilot
     */
    handlebars.registerHelper("getProperty", (obj: object, propertyKey: string): unknown => {
        return obj && obj[propertyKey];
    });
}