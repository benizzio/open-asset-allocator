import * as handlebars from "handlebars";
import format from "./format";

const hanldebarsFormatCurrency = (value: number, currency: unknown) => {
    const localCurrency = typeof currency === "string" ? currency : undefined;
    return format.formatCurrency(value, localCurrency);
};

const handlebarsFormat = {
    registerHandlebarsFormatHelper() {
        handlebars.registerHelper("formatCurrency", hanldebarsFormatCurrency);
    },
};

export default handlebarsFormat;