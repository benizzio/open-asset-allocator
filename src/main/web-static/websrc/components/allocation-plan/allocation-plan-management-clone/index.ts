/**
 * Copies a persisted plan's current editable state into the existing new-plan form.
 * Capture and replacement are synchronous; asynchronous accordion presentation
 * is separately guarded against context changes, edits, and newer operations.
 * @module components/allocation-plan/allocation-plan-management-clone
 * @author GPT-6.1 Terra
 */
import htmx from "htmx.org";
import type { TemplateDelegate } from "handlebars";
import { PortfolioPage } from "../../../pages";
import { getBrowserLocalCalendarDate } from "../../../utils/calendar-date";
import { bindCloneRequestGuards, focusDraftDestination, trackCloneOpening, type CloneOperation } from
    "../../management-draft-presentation";
import { allocationPlanNameValidationMessage } from "../allocation-plan-management-metadata";
import { capturePlanDraft, restorePlanRows } from "./draft";
import { renderPlanRows } from "./rendering";

const FORM_ID = "allocation-plan-management-form-0";
const ROUTE = "/portfolio/:portfolioId/allocation/manage";
const boundLists = new WeakSet<HTMLElement>();

/** Finds and validates the destination controls before changing the existing draft.
 * @author GPT-6.1 Terra
 */
function destinationControls(item: HTMLElement) {
    const form = item.querySelector<HTMLFormElement>(`#${ FORM_ID }`);
    const tbody = form?.querySelector<HTMLElement>("tbody");
    const name = item.querySelector<HTMLInputElement>(".allocation-plan-management-new-name");
    const date = item.querySelector<HTMLInputElement>(".allocation-plan-management-new-date");
    const payload = form?.elements.namedItem("name");
    const counter = form?.elements.namedItem("last-planned-allocation-row-index");

    if(!form || !tbody || !name || !date) {
        throw new Error("Allocation plan destination controls are unavailable.");
    }

    if(!(payload instanceof HTMLInputElement) || !(counter instanceof HTMLInputElement)) {
        throw new Error("Allocation plan destination controls are unavailable.");
    }
    return { form, tbody, name, date, payload, counter };
}

/** Replaces the destination after complete source capture and markup preflight.
 * @author GPT-6.1 Terra
 */
function copyDraft(source: HTMLFormElement, destination: HTMLElement, template: TemplateDelegate): HTMLInputElement {
    const date = getBrowserLocalCalendarDate();
    const controls = destinationControls(destination);
    const hierarchy = PortfolioPage.getContextPortfolio().allocationStructure.hierarchy;
    const draft = capturePlanDraft(source, date, hierarchy.length);
    const markup = renderPlanRows(template, hierarchy, draft);

    htmx.swap(controls.tbody, markup, { swapStyle: "innerHTML", swapDelay: 0, settleDelay: 0 },
        { afterSettleCallback: () => restorePlanRows(controls.tbody, draft) });
    controls.name.value = draft.name;
    controls.name.setCustomValidity(allocationPlanNameValidationMessage(draft.name));
    controls.payload.value = draft.name;
    controls.date.value = draft.date;
    controls.date.setCustomValidity("");
    controls.counter.value = String(draft.rows.length - 1);
    controls.form.classList.remove("was-validated");
    controls.name.classList.remove("is-invalid", "is-valid");
    controls.date.classList.remove("is-invalid", "is-valid");
    return controls.name;
}

/** Opens the replaced draft and releases this operation's listeners and busy state.
 * @author GPT-6.1 Terra
 */
async function presentDraft(
    destination: HTMLElement, name: HTMLInputElement, button: HTMLButtonElement, status: HTMLElement,
    operation: CloneOperation, finished: () => void,
): Promise<void> {
    button.setAttribute("aria-busy", "true");

    try {
        await focusDraftDestination(destination, name, operation);

        if(operation.owns()) {
            status.textContent = "Allocation plan copied to the new draft.";
        }
    } catch(error) {
        if(operation.owns()) {
            status.textContent = "Unable to open the new allocation plan. Please try again.";
            console.error(error);
        }
    } finally {
        operation.dispose();
        finished();
    }
}

/** Binds delegated cloning for initial and subsequently appended persisted plans.
 * @param list - Stable paginated list owner.
 * @param template - Returns the initialized row template.
 * @example `bindAllocationPlanCloning(list, () => rowTemplate);`
 * @author GPT-6.1 Terra
 */
export function bindAllocationPlanCloning(list: HTMLElement, template: () => TemplateDelegate): void {
    if(boundLists.has(list)) {
        return;
    }
    boundLists.add(list);
    let active: CloneOperation | undefined;
    let pendingSource: HTMLElement | undefined;
    const owners = new WeakMap<HTMLButtonElement, CloneOperation>();
    const cancel = () => active?.controller.abort();
    const isSaving = bindCloneRequestGuards(list, cancel, FORM_ID);

    list.addEventListener("click", event => {
        const button = (event.target as Element)?.closest<HTMLButtonElement>("[data-allocation-plan-clone]");

        if(!button || !list.contains(button)) {
            return;
        }
        event.preventDefault();
        event.stopPropagation();
        const item = button.closest<HTMLElement>(".accordion-item");
        const source = item?.querySelector<HTMLFormElement>("form");

        const destination = list.querySelector<HTMLFormElement>(`#${ FORM_ID }`)
            ?.closest<HTMLElement>(".accordion-item");

        const status = item?.querySelector<HTMLElement>("[data-allocation-plan-clone-status]");

        if(!item || !source || !destination || !status) {
            return;
        }

        if(item === destination) {
            return;
        }

        if(isSaving()) {
            status.textContent = "Wait for the current save to finish, then try again.";
            return;
        }

        if(active && !active.controller.signal.aborted && pendingSource === item) {
            return;
        }

        cancel();

        try {
            const name = copyDraft(source, destination, template());
            const operation = trackCloneOpening(list, item, destination, FORM_ID, ROUTE);
            active = operation;
            pendingSource = item;
            owners.set(button, operation);

            void presentDraft(destination, name, button, status, operation, () => {
                if(owners.get(button) === operation) {
                    button.removeAttribute("aria-busy");
                    owners.delete(button);
                }

                if(active === operation) {
                    active = undefined;
                    pendingSource = undefined;
                }
            });
        } catch(error) {
            status.textContent = "Unable to clone allocation plan. Please try again.";
            console.error(error);
        }
    });
}
