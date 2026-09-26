/**
 * Exposes Handlebars initialization and template utilities through the infrastructure module's public API.
 *
 * @module infra/handlebars
 * @author Igor Benicio de Mesquita
 * @author OpenCode
 */

import * as handlebars from "handlebars";
import { registerHandlebarsChartHelper } from "./handlebars-chart";
import { registerHandlebarsFormatHelper } from "./handlebars-format";
import { registerHandlebarsLangHelpers } from "./handlebars-lang";
import { registerPartialToContainer } from "./handlebars-partial";
import { registerHandlebarsDOMHelpers } from "./handlebars-dom";
import { registerHandlebarsUtilHelpers } from "./handlebars-util";
import { renderTemplateByElementId } from "./handlebars-render";

/**
 * Provides Handlebars helper registration and browser-facing template utilities.
 *
 * @example
 * ```ts
 * handlebarsInfra.register();
 * const html = handlebarsInfra.utils.renderTemplateByElementId("row-template", { name: "Example" });
 * ```
 *
 * @author Igor Benicio de Mesquita
 * @author OpenCode
 */
export const handlebarsInfra = {
    register: () => {
        registerHandlebarsChartHelper();
        registerHandlebarsFormatHelper();
        registerHandlebarsLangHelpers();
        registerHandlebarsDOMHelpers();
        registerHandlebarsUtilHelpers();
        return handlebars;
    },
    utils: { registerPartialToContainer, renderTemplateByElementId },
};
