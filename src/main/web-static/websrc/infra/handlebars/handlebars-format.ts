import * as handlebars from "handlebars";
import Format from "../format";

const handlebarsFormatCurrency = (value: number | string, currency: unknown) => {
    const localCurrency = typeof currency === "string" ? currency : undefined;
    const numericValue = typeof value === "number" ? value : parseFloat(value);
    return Format.formatCurrency(numericValue, localCurrency);
};

export function registerHandlebarsFormatHelper() {
    handlebars.registerHelper("formatCurrency", handlebarsFormatCurrency);
}

