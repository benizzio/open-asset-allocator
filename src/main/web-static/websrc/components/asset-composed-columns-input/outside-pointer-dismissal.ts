/** Coordinates outside-pointer dismissal without moving the eventual click target before activation.
 *
 * @author GPT-6 Luna
 */

import type { AssetSearchAutocompleteState } from "./types";

/** Registers deferred outside-pointer dismissal listeners for one open autocomplete.
 *
 * @param state - Open autocomplete state whose popup can affect document layout.
 * @param dismiss - Callback that closes the autocomplete after outside activation.
 * @example `bindOutsidePointerDismissal(state, () => controller.close(input));`
 * @author GPT-6 Luna
 */
export function bindOutsidePointerDismissal(
    state: AssetSearchAutocompleteState,
    dismiss: () => void,
): void {
    state.outsidePointerHandler = event => {
        if(state.outsidePointerPending) {
            clearPendingOutsidePointer(state);
        }

        const target = event.target;

        if(target instanceof Node && state.wrapper.contains(target)) {
            return;
        }

        if(event.button !== 0) {
            dismiss();
            return;
        }

        state.outsidePointerPending = true;
        state.outsidePointerId = event.pointerId;

        state.outsideClickHandler = clickEvent => {
            if(!state.outsidePointerPending) {
                return;
            }

            const clickTarget = clickEvent.target;
            const shouldDismiss = !(clickTarget instanceof Node) || !state.wrapper.contains(clickTarget);
            clearPendingOutsidePointer(state);

            if(shouldDismiss) {
                dismiss();
            }
        };

        state.outsidePointerUpHandler = pointerEvent => {
            if(pointerEvent.pointerId !== state.outsidePointerId) {
                return;
            }

            state.outsidePointerTimeout = window.setTimeout(() => {
                if(state.outsidePointerPending) {
                    clearPendingOutsidePointer(state);
                    dismiss();
                }
            }, 0);
        };

        state.outsidePointerCancelHandler = pointerEvent => {
            if(pointerEvent.pointerId !== state.outsidePointerId) {
                return;
            }

            clearPendingOutsidePointer(state);
            dismiss();
        };

        state.outsidePointerBlurHandler = () => {
            if(state.outsidePointerPending) {
                clearPendingOutsidePointer(state);
                dismiss();
            }
        };

        document.addEventListener("click", state.outsideClickHandler, true);
        document.addEventListener("pointerup", state.outsidePointerUpHandler, true);
        document.addEventListener("pointercancel", state.outsidePointerCancelHandler, true);
        window.addEventListener("blur", state.outsidePointerBlurHandler);
    };

    document.addEventListener("pointerdown", state.outsidePointerHandler, true);
}

/** Removes every listener and timeout for a pending outside-pointer interaction.
 *
 * @param state - Autocomplete state whose outside-pointer interaction is being cleared.
 * @example `clearOutsidePointerDismissal(state);`
 * @author GPT-6 Luna
 */
export function clearOutsidePointerDismissal(state: AssetSearchAutocompleteState): void {
    if(state.outsidePointerHandler) {
        document.removeEventListener("pointerdown", state.outsidePointerHandler, true);
    }

    state.outsidePointerHandler = undefined;
    clearPendingOutsidePointer(state);
}

/** Clears temporary completion listeners after an outside pointer sequence settles.
 *
 * @param state - Autocomplete state whose temporary listeners are being cleared.
 * @author GPT-6 Luna
 */
function clearPendingOutsidePointer(state: AssetSearchAutocompleteState): void {
    if(state.outsideClickHandler) {
        document.removeEventListener("click", state.outsideClickHandler, true);
    }

    if(state.outsidePointerUpHandler) {
        document.removeEventListener("pointerup", state.outsidePointerUpHandler, true);
    }

    if(state.outsidePointerCancelHandler) {
        document.removeEventListener("pointercancel", state.outsidePointerCancelHandler, true);
    }

    if(state.outsidePointerBlurHandler) {
        window.removeEventListener("blur", state.outsidePointerBlurHandler);
    }

    if(state.outsidePointerTimeout !== undefined) {
        window.clearTimeout(state.outsidePointerTimeout);
    }

    state.outsideClickHandler = undefined;
    state.outsidePointerUpHandler = undefined;
    state.outsidePointerCancelHandler = undefined;
    state.outsidePointerBlurHandler = undefined;
    state.outsidePointerTimeout = undefined;
    state.outsidePointerId = undefined;
    state.outsidePointerPending = false;
}
