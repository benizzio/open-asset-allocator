import DomUtils from "./dom-utils";

function bindSubmit(form: HTMLFormElement) {
    form.addEventListener("submit", event => {

        if(!form.checkValidity()) {
            event.preventDefault();
            event.stopPropagation();
        }

        form.classList.add("was-validated");

    }, false);
}

function bindExtract(form: HTMLFormElement) {
    form.addEventListener("reset", () => {
        form.classList.remove("was-validated");
    }, false);
}

export function bindFormsInDescendants(element: HTMLElement) {

    const forms = DomUtils.queryAllInDescendants(element, "form.needs-validation:not([novalidate])");

    forms.forEach((form: HTMLFormElement) => {
        bindSubmit(form);
        bindExtract(form);
        form.setAttribute("novalidate", "true");
    });
}