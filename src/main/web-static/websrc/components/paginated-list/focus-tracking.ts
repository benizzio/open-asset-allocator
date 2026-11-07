/**
 * Tracks focus changes only while keyboard-initiated pagination requests are active.
 * @module components/paginated-list
 * @author GPT-6.1 Terra
 */

/** Describes the focus information needed for one keyboard pagination request.
 * @example `startTrackingFocus({ focusButton: button, focusChanged: false })` tracks a request.
 * @author GPT-6.1 Terra
 */
type FocusContext = { focusButton?: HTMLButtonElement; focusChanged: boolean };

const tracked = new Set<FocusContext>();

/** Marks requests abandoned if focus moves to another element.
 * @author GPT-6.1 Terra
 */
function onFocusChange(event: FocusEvent): void {
    for(const context of tracked) {
        if(event.target !== context.focusButton && event.target !== document.body) {
            context.focusChanged = true;
        }
    }
}

/**
 * Starts tracking a keyboard request, installing the document listener only while needed.
 * @param context - Mutable focus state of the request.
 * @example `startTrackingFocus(context)` begins tracking before sending.
 * @author GPT-6.1 Terra
 */
export function startTrackingFocus(context: FocusContext): void {
    if(!context.focusButton) {
        return;
    }

    if(tracked.size === 0) {
        document.addEventListener("focusin", onFocusChange);
    }

    tracked.add(context);
}

/**
 * Stops tracking a request and releases the document listener after the last request.
 * @param context - The completed or invalidated request.
 * @example `stopTrackingFocus(context)` releases tracking after completion.
 * @author GPT-6.1 Terra
 */
export function stopTrackingFocus(context: FocusContext): void {
    tracked.delete(context);

    if(tracked.size === 0) {
        document.removeEventListener("focusin", onFocusChange);
    }
}
