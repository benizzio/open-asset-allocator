import * as handlebars from "handlebars";
import Format from "../format";

const hanldebarsFormatCurrency = (value: number, currency: unknown) => {
    const localCurrency = typeof currency === "string" ? currency : undefined;
    return Format.formatCurrency(value, localCurrency);
};

export function registerHandlebarsFormatHelper() {
    handlebars.registerHelper("formatCurrency", hanldebarsFormatCurrency);
}

