/**
 * Registers pointer, focus, and keyboard behaviors for asset-search autocomplete controls.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */

import {
    ASSET_ACTION_BUTTON_IDENTITIES,
    ASSET_ACTION_BUTTON_SELECTOR,
    ASSET_SEARCH_OPTION_SELECTOR,
} from "./constants";
import type { AssetSearchAutocompleteState, AutocompleteControllerActions } from "./types";

/** Attaches pointer, focus, and accessible-click handlers to one autocomplete widget.
 *
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
export function registerAutocompleteInteractions(
    input: HTMLInputElement,
    state: AssetSearchAutocompleteState,
    controller: AutocompleteControllerActions,
): void {
    const { listbox, wrapper } = state;

    listbox.addEventListener("pointerover", event => {
        const pointerEvent = event as PointerEvent;
        const option = (pointerEvent.target as HTMLElement).closest<HTMLElement>(ASSET_SEARCH_OPTION_SELECTOR);

        if(pointerEvent.pointerType === "mouse" && option) {
            controller.highlight(input, Number(option.dataset.optionIndex), "pointer", false);
        }
    });

    listbox.addEventListener("pointerout", event => {
        const pointerEvent = event as PointerEvent;

        if(pointerEvent.pointerType !== "mouse"
            || (pointerEvent.relatedTarget instanceof Node && listbox.contains(pointerEvent.relatedTarget))) {
            return;
        }

        if(state.activeOptionSource === "pointer" && !state.pointerSelection) {
            controller.clearHighlight(input);
        }
    });

    listbox.addEventListener("pointerdown", event => {
        const pointerEvent = event as PointerEvent;
        const option = (pointerEvent.target as HTMLElement).closest<HTMLElement>(ASSET_SEARCH_OPTION_SELECTOR);

        if(!option || pointerEvent.button !== 0) {
            state.pointerSelection = undefined;
            return;
        }

        pointerEvent.preventDefault();

        state.pointerSelection = {
            optionIndex: Number(option.dataset.optionIndex),
            pointerId: pointerEvent.pointerId,
            x: pointerEvent.clientX,
            y: pointerEvent.clientY,
            isCancelled: false,
        };
        controller.highlight(input, state.pointerSelection.optionIndex, "pointer", false);
    });

    listbox.addEventListener("pointermove", event => {
        const pointerEvent = event as PointerEvent;
        const pointerSelection = state.pointerSelection;

        if(pointerSelection?.pointerId !== pointerEvent.pointerId || pointerSelection?.isCancelled !== false) {
            return;
        }

        const movedX = pointerEvent.clientX - pointerSelection.x;
        const movedY = pointerEvent.clientY - pointerSelection.y;

        if(Math.hypot(movedX, movedY) > 8) {
            pointerSelection.isCancelled = true;
            controller.clearHighlight(input);
        }
    });

    listbox.addEventListener("pointerup", event => {
        const pointerEvent = event as PointerEvent;
        const pointerSelection = state.pointerSelection;
        state.pointerSelection = undefined;

        if(pointerSelection?.pointerId !== pointerEvent.pointerId || pointerSelection?.isCancelled !== false) {
            return;
        }

        const pointerTarget = document.elementFromPoint(pointerEvent.clientX, pointerEvent.clientY);
        const option = pointerTarget?.closest<HTMLElement>(ASSET_SEARCH_OPTION_SELECTOR);
        const optionIndex = option ? Number(option.dataset.optionIndex) : -1;
        const movedX = pointerEvent.clientX - pointerSelection.x;
        const movedY = pointerEvent.clientY - pointerSelection.y;

        if(optionIndex !== pointerSelection.optionIndex || Math.hypot(movedX, movedY) > 8) {
            return;
        }

        const selectedOption = state.options[optionIndex];

        if(selectedOption) {
            controller.selectOption(input, selectedOption.ticker);
        }
    });

    listbox.addEventListener("pointercancel", () => {
        state.pointerSelection = undefined;
        controller.clearHighlight(input);
    });

    listbox.addEventListener("click", event => {
        const clickEvent = event as MouseEvent;

        if(clickEvent.detail !== 0) {
            return;
        }

        const option = (clickEvent.target as HTMLElement).closest<HTMLElement>(ASSET_SEARCH_OPTION_SELECTOR);
        const selectedOption = option ? state.options[Number(option.dataset.optionIndex)] : undefined;

        if(selectedOption) {
            controller.selectOption(input, selectedOption.ticker);
        }
    });

    wrapper.addEventListener("focusout", event => {
        const focusEvent = event as FocusEvent;

        if(focusEvent.relatedTarget && wrapper.contains(focusEvent.relatedTarget as Node)) {
            return;
        }

        controller.close(input);
    });
}

/** Handles autocomplete navigation, selection, literal lookup fallback, Escape, and Tab.
 *
 * @author GPT-6.1 Sol
 * @author GPT-6 Sol
 */
export function handleAssetSearchKeydown(
    input: HTMLInputElement,
    event: KeyboardEvent,
    controller: AutocompleteControllerActions,
): void {
    if(event.isComposing) {
        return;
    }

    const state = controller.ensureState(input);

    if(event.key === "ArrowDown" || event.key === "ArrowUp") {
        navigateAssetSearchOptions(input, event, controller);
        return;
    }

    if(event.key === "Escape") {
        if(state?.isOpen) {
            event.preventDefault();
            controller.close(input);
        }
        return;
    }

    if(event.key === "Tab") {
        controller.close(input);
        return;
    }

    if(event.key === "Enter") {
        submitAssetSearchLookup(input, state, event, controller);
    }
}

/** Opens editable suggestions and moves the active option in the arrow-key direction.
 *
 * @author GPT-6.1 Sol
 * @author GPT-6 Sol
 */
function navigateAssetSearchOptions(
    input: HTMLInputElement,
    event: KeyboardEvent,
    controller: AutocompleteControllerActions,
): void {
    if(input.readOnly) {
        return;
    }

    controller.open(input);
    const state = controller.ensureState(input);

    if(!state?.options.length) {
        return;
    }

    event.preventDefault();
    const direction = event.key === "ArrowDown" ? 1 : -1;
    let nextIndex = state.activeOptionIndex + direction;

    if(state.activeOptionIndex < 0) {
        nextIndex = direction === 1 ? 0 : state.options.length - 1;
    }

    controller.highlight(input, nextIndex, "keyboard", true);
}

/** Selects the active suggestion or clicks the row's search action for a literal query.
 *
 * @author GPT-6.1 Sol
 * @author GPT-6 Sol
 */
function submitAssetSearchLookup(
    input: HTMLInputElement,
    state: AssetSearchAutocompleteState | null,
    event: KeyboardEvent,
    controller: AutocompleteControllerActions,
): void {
    if(state && state.activeOptionIndex >= 0) {
        event.preventDefault();
        const selectedOption = state.options[state.activeOptionIndex];

        if(selectedOption) {
            controller.selectOption(input, selectedOption.ticker);
        }
        return;
    }

    const actionButton = input
        .closest(".input-group")
        ?.querySelector<HTMLButtonElement>(ASSET_ACTION_BUTTON_SELECTOR);

    if(actionButton?.className === ASSET_ACTION_BUTTON_IDENTITIES.search.classes) {
        event.preventDefault();
        actionButton.click();
    }
}
