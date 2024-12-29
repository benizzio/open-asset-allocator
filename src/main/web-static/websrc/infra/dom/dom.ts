import { bindExclusiveDisplayContainerInDescendants } from "./dom-exclusive-display";

export const domInfra = {
    bindDescendants: (element: HTMLElement) => {
        bindExclusiveDisplayContainerInDescendants(element);
    },
};