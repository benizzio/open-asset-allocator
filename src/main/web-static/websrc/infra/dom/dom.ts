import { bindExclusiveDisplayContainerInDescendants } from "./dom-exclusive-display";
import { bindFormsInDescendants } from "./dom-form";
import { maskNumberDecimalPlaces, maskTagInput, maskTickerInput } from "./dom-form-input";

export const domInfra = {
    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
        bindFormsInDescendants(element);
    },
    bindGlobalFunctions() {
        window["maskTagInput"] = maskTagInput;
        window["maskTickerInput"] = maskTickerInput;
        window["maskNumberDecimalPlaces"] = maskNumberDecimalPlaces;
    },
};