import * as handlebars from "handlebars";
import { registerHandlebarsChartHelper } from "./handlebars-chart";
import { registerHandlebarsFormatHelper } from "./handlebars-format";
import { registerHandlebarsLangHelpers } from "./handlebars-lang";

export const handlebarsInfra = {
    register: () => {
        registerHandlebarsChartHelper();
        registerHandlebarsFormatHelper();
        registerHandlebarsLangHelpers();
        return handlebars;
    },
};