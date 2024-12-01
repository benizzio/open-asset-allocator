import * as handlebars from "handlebars";

export function registerHandlebarsObjectHelpers() {

    handlebars.registerHelper("object", ({ hash }) => {
        return hash;
    });

    handlebars.registerHelper("array", function(...args) {
        return Array.from(args).slice(0, arguments.length-1);
    });
}