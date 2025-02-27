import * as handlebars from "handlebars";

export function registerHandlebarsLangHelpers() {

    handlebars.registerHelper("object", ({ hash }) => {
        return hash;
    });

    handlebars.registerHelper("array", function(...args) {
        return Array.from(args).slice(0, arguments.length - 1);
    });

    handlebars.registerHelper("domJSON", function(id: string, object: object) {
        return `
            <script id="${ id }" type="application/json">${ JSON.stringify(object) }</script>
        `;
    });

    handlebars.registerHelper("repeater", function(text: unknown, count: number, prefix: string, suffix: string) {
        if(count <= 0) {
            return "";
        }
        let result = String(text).repeat(count);
        result = prefix ? prefix + result : result;
        result = suffix ? result + suffix : result;
        return result;
    });

    handlebars.registerHelper("stringify", function(object: object) {
        return JSON.stringify(object);
    });

    handlebars.registerHelper("concat", function(...args) {
        return Array.from(args).slice(0, arguments.length - 1).join("");
    });
}