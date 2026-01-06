import DomUtils from "./dom-utils";
import { BootstrapClasses } from "../bootstrap/constants";

const ATTRIBUTE_BOOTSRAP_VALIDATION_BOUND_FLAG = "data-bootstrap-validation-bound";

function bindBootstrapValidationOnSubmit(form: HTMLFormElement) {

    form.addEventListener("submit", event => {

        if(!form.checkValidity()) {
            event.preventDefault();
            event.stopPropagation();
        }

        form.classList.add(BootstrapClasses.WAS_VALIDATED);

    }, false);
}

function bindBootstrapValidationCleaning(form: HTMLFormElement) {
    form.addEventListener("reset", () => {
        form.classList.remove(BootstrapClasses.WAS_VALIDATED);
    }, false);
}

function bindBootstrapValidationToDefaultForm(form: HTMLFormElement) {
    form.addEventListener("invalid", () => {
        form.classList.add(BootstrapClasses.WAS_VALIDATED);
    }, true);
}

export function bindFormsInDescendants(element: HTMLElement) {

    const forms = DomUtils.queryAllInDescendants(
        element,
        `form.${ BootstrapClasses.NEEDS_VALIDATION }:not([${ ATTRIBUTE_BOOTSRAP_VALIDATION_BOUND_FLAG }])`,
    );

    forms.forEach((form: HTMLFormElement) => {

        form.setAttribute(ATTRIBUTE_BOOTSRAP_VALIDATION_BOUND_FLAG, "binding");

        try {
            if(form.noValidate) {
                bindBootstrapValidationOnSubmit(form);
            }
            else {
                bindBootstrapValidationToDefaultForm(form);
            }

            bindBootstrapValidationCleaning(form);

            form.setAttribute(ATTRIBUTE_BOOTSRAP_VALIDATION_BOUND_FLAG, "true");
        } catch(error) {
            form.removeAttribute(ATTRIBUTE_BOOTSRAP_VALIDATION_BOUND_FLAG);
            throw error;
        }
    });
}