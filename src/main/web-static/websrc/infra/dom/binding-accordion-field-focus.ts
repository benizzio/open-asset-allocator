/**
 * Binds focus-driven opening for accordion panels whose header contains editable fields.
 *
 * @example `bindAccordionFieldFocusInDescendants(document.body)` binds marked field groups.
 * @author GPT-6 Luna
 */

import { Collapse } from "bootstrap";

const FIELD_GROUP_SELECTOR = "[data-accordion-field-focus]";
const BOUND_ATTRIBUTE = "data-accordion-field-focus-bound";
const ACCORDION_ITEM_SELECTOR = ".accordion-item";
const COLLAPSE_PANEL_SELECTOR = ".accordion-collapse";
const SHOWN_EVENT = "shown.bs.collapse";
const HIDDEN_EVENT = "hidden.bs.collapse";

/** Tracks whether a focused field still requires its associated panel to open.
 * @author GPT-6 Luna
 */
type AccordionFieldFocusState = { pendingOpen: boolean };

/**
 * Binds marked accordion header field groups in an element and its descendants.
 * A focused field opens its sibling panel without changing focus; pending opening is retried after a transition.
 *
 * @param element - Root whose marked field groups should be bound.
 * @example `bindAccordionFieldFocusInDescendants(container)` binds any marked groups below the container.
 * @author GPT-6 Luna
 */
export function bindAccordionFieldFocusInDescendants(element: HTMLElement): void {
    const fieldGroups = Array.from(element.querySelectorAll<HTMLElement>(FIELD_GROUP_SELECTOR));

    if(element.matches(FIELD_GROUP_SELECTOR)) {
        fieldGroups.unshift(element);
    }

    fieldGroups.forEach(bindAccordionFieldFocus);
}

/** Binds focus and transition listeners to one marked field group when it has a matching panel.
 * @author GPT-6 Luna
 */
function bindAccordionFieldFocus(fields: HTMLElement): void {
    if(fields.hasAttribute(BOUND_ATTRIBUTE)) {
        return;
    }

    const accordionItem = fields.closest(ACCORDION_ITEM_SELECTOR);
    const panel = accordionItem?.querySelector<HTMLElement>(COLLAPSE_PANEL_SELECTOR);

    if(!panel) {
        return;
    }

    const state: AccordionFieldFocusState = { pendingOpen: false };
    fields.addEventListener("focusin", () => handleFieldFocus(fields, panel, state));
    fields.addEventListener("focusout", event => handleFieldBlur(event, fields, state));
    panel.addEventListener(SHOWN_EVENT, () => showPanelIfFocusRemains(fields, panel, state));
    panel.addEventListener(HIDDEN_EVENT, () => showPanelIfFocusRemains(fields, panel, state));
    fields.setAttribute(BOUND_ATTRIBUTE, "true");
}

/** Requests opening when a field in the header receives focus.
 * @author GPT-6 Luna
 */
function handleFieldFocus(
    fields: HTMLElement,
    panel: HTMLElement,
    state: AccordionFieldFocusState,
): void {
    state.pendingOpen = true;
    showPanelIfFocusRemains(fields, panel, state);
}

/** Cancels a deferred open when focus leaves the marked fields.
 * @author GPT-6 Luna
 */
function handleFieldBlur(
    event: FocusEvent,
    fields: HTMLElement,
    state: AccordionFieldFocusState,
): void {
    if(!(event.relatedTarget instanceof Node) || !fields.contains(event.relatedTarget)) {
        state.pendingOpen = false;
    }
}

/** Shows the panel if it remains connected and focus is still within the associated field group.
 * @author GPT-6 Luna
 */
function showPanelIfFocusRemains(
    fields: HTMLElement,
    panel: HTMLElement,
    state: AccordionFieldFocusState,
): void {
    if(!state.pendingOpen) {
        return;
    }

    if(!fields.isConnected || !panel.isConnected || !fields.contains(document.activeElement)) {
        state.pendingOpen = false;
        return;
    }

    if(panel.classList.contains("show")) {
        state.pendingOpen = false;
        return;
    }

    Collapse.getOrCreateInstance(panel, { toggle: false }).show();
}
