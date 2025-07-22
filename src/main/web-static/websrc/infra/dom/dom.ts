import { bindExclusiveDisplayContainerInDescendants } from "./dom-exclusive-display";
import { bindFormsInDescendants } from "./dom-form";
import { maskTagInput, maskTickerInput } from "./dom-form-input";

export const domInfra = {
    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
        bindFormsInDescendants(element);
    },
    bindGlobalFunctions() {
        window["maskTagInput"] = maskTagInput;
        window["maskTickerInput"] = maskTickerInput;
    },
};