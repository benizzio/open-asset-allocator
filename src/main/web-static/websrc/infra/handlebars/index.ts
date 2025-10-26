import * as handlebars from "handlebars";
import { registerHandlebarsChartHelper } from "./handlebars-chart";
import { registerHandlebarsFormatHelper } from "./handlebars-format";
import { registerHandlebarsLangHelpers } from "./handlebars-lang";
import { registerPartialToContainer } from "./handlebars-partial";

export const handlebarsInfra = {
    register: () => {
        registerHandlebarsChartHelper();
        registerHandlebarsFormatHelper();
        registerHandlebarsLangHelpers();
        return handlebars;
    },
    utils: { registerPartialToContainer: registerPartialToContainer },
};