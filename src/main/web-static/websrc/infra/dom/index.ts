/**
 * Public DOM infrastructure for binding form controls and preserving financial-input state.
 * @module infra/dom
 * @author benizzio
 * @author GPT-6.1 Terra
 */
import { bindExclusiveDisplayContainerInDescendants } from "./binding-dom-exclusive-display";
import { bindFormsInDescendants } from "./binding-dom-bootstrap-form";
import { maskNumberDecimalPlaces, maskTagInput, maskTickerInput } from "./dom-form-input";
import { bindPercentageInputsInDescendants } from "./binding-percentage-input";
import { bindFinancialInputsInDescendants } from "./binding-financial-input";
import { bindAccordionFieldFocusInDescendants } from "./binding-accordion-field-focus";
import DomUtils from "./dom-utils";

export { captureFinancialInput, restoreFinancialInput } from "./financial-input-snapshot";
export type { FinancialInputSnapshot } from "./financial-input-snapshot";
export { capturePercentageInput, restorePercentageInput } from "./percentage-input-snapshot";
export type { PercentageInputSnapshot } from "./percentage-input-snapshot";

const DomInfra = {

    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
        bindFormsInDescendants(element);
        bindPercentageInputsInDescendants(element);
        bindFinancialInputsInDescendants(element);
        bindAccordionFieldFocusInDescendants(element);
    },

    bindGlobalFunctions() {
        window["maskTagInput"] = maskTagInput;
        window["maskTickerInput"] = maskTickerInput;
        window["maskNumberDecimalPlaces"] = maskNumberDecimalPlaces;
    },

    DomUtils,
};

export default DomInfra;
