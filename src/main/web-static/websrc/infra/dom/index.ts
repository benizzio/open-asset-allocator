import { bindExclusiveDisplayContainerInDescendants } from "./binding-dom-exclusive-display";
import { bindFormsInDescendants } from "./binding-dom-bootstrap-form";
import { maskNumberDecimalPlaces, maskTagInput, maskTickerInput } from "./dom-form-input";
import { bindPercentageInputsInDescendants } from "./binding-percentage-input";
import { bindFinancialInputsInDescendants } from "./binding-financial-input";
import DomUtils from "./dom-utils";

const DomInfra = {

    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
        bindFormsInDescendants(element);
        bindPercentageInputsInDescendants(element);
        bindFinancialInputsInDescendants(element);
    },

    bindGlobalFunctions() {
        window["maskTagInput"] = maskTagInput;
        window["maskTickerInput"] = maskTickerInput;
        window["maskNumberDecimalPlaces"] = maskNumberDecimalPlaces;
    },

    DomUtils,
};

export default DomInfra;