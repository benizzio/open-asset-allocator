import { bindExclusiveDisplayContainerInDescendants } from "./dom-exclusive-display";
import { bindFormsInDescendants } from "./dom-form";
import { maskNumberDecimalPlaces, maskTagInput, maskTickerInput } from "./dom-form-input";
import { bindPercentageInputsInDescendants } from "./binding-percentage-input";
import DomUtils from "./dom-utils";

const DomInfra = {

    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
        bindFormsInDescendants(element);
        bindPercentageInputsInDescendants(element);
    },

    bindGlobalFunctions() {
        window["maskTagInput"] = maskTagInput;
        window["maskTickerInput"] = maskTickerInput;
        window["maskNumberDecimalPlaces"] = maskNumberDecimalPlaces;
    },

    DomUtils,
};

export default DomInfra;