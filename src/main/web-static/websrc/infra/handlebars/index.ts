import * as handlebars from "handlebars";
import { registerHandlebarsChartHelper } from "./handlebars-chart";
import { registerHandlebarsFormatHelper } from "./handlebars-format";
import { registerHandlebarsLangHelpers } from "./handlebars-lang";
import { registerPartialToContainer } from "./handlebars-partial";
import { registerHandlebarsDOMHelpers } from "./handlebars-dom";
import { registerHandlebarsUtilHelpers } from "./handlebars-util";

export const handlebarsInfra = {
    register: () => {
        registerHandlebarsChartHelper();
        registerHandlebarsFormatHelper();
        registerHandlebarsLangHelpers();
        registerHandlebarsDOMHelpers();
        registerHandlebarsUtilHelpers();
        return handlebars;
    },
    utils: { registerPartialToContainer: registerPartialToContainer },
};