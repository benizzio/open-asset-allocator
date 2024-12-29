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
}