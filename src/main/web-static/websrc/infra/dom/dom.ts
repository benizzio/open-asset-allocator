import { bindExclusiveDisplayContainerInDescendants } from "./dom-exclusive-display";
import { bindFormsInDescendants } from "./form";

export const domInfra = {
    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
        bindFormsInDescendants(element);
    },
};