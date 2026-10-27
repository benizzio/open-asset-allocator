/**
 * Renders owner-scoped pagination feedback and provides focus targets after a terminal append.
 * @module components/paginated-list
 * @author GPT-6 Luna
 */

const LOAD_MORE_ATTRIBUTE = "data-paginated-list-more";
const STATUS_ATTRIBUTE = "data-paginated-list-status";
const CONTROLS_ATTRIBUTE = "data-paginated-list-controls";

/** Describes the request state needed to render one list's pagination controls. */
export type PaginationControlState = {
    initialized: boolean;
    cursor: string | null;
    loading: boolean;
    retryAvailable: boolean;
    controlRequestPending: boolean;
};

/**
 * Updates only the continuation controls associated with the supplied list owner.
 *
 * @param owner - The paginated accordion whose controls should be rendered.
 * @param state - The minimal cursor/request state needed for presentation.
 * @param message - Optional owner-scoped live-region feedback.
 * @param isError - Whether the message is an actionable failure alert.
 * @example `renderControls(owner, state, "Could not load records. Try again.", true)` renders a retry state.
 * @author GPT-6 Luna
 */
export function renderControls(
    owner: HTMLElement,
    state: PaginationControlState,
    message = "",
    isError = false,
): void {
    const buttonVisible = shouldShowButton(state);

    updateButton(findOwnedElement<HTMLButtonElement>(LOAD_MORE_ATTRIBUTE, owner.id), state, buttonVisible);
    updateStatus(findOwnedElement<HTMLElement>(STATUS_ATTRIBUTE, owner.id), message, isError);
    updateControlsGroup(findOwnedElement<HTMLElement>(CONTROLS_ATTRIBUTE, owner.id), buttonVisible, message);
}

/**
 * Moves focus to a visible persisted record toggle after the final continuation page commits.
 *
 * @param owner - The list whose accepted page has just settled.
 * @param acceptedIds - Persisted IDs accepted from the settled response, in response order.
 * @example `focusOnAcceptedRecord(owner, request.acceptedIds)` focuses the first newly appended record.
 * @author GPT-6 Luna
 */
export function focusOnAcceptedRecord(owner: HTMLElement, acceptedIds: readonly string[]): void {
    if(owner.getClientRects().length === 0) {
        return;
    }

    const records = Array.from(owner.querySelectorAll<HTMLElement>("[data-paginated-record-id]"));

    const acceptedRecord = acceptedIds
        .map(id => records.find(record => record.getAttribute("data-paginated-record-id") === id))
        .find((record): record is HTMLElement => record !== undefined);

    const lastPersistedRecord = records
        .filter(record => record.getAttribute("data-paginated-record-id") !== "0")
        .at(-1);

    const toggle = (acceptedRecord ?? lastPersistedRecord)
        ?.querySelector<HTMLButtonElement>(".accordion-button");

    if(toggle && toggle.getClientRects().length > 0) {
        toggle.focus();
    }
}

/** Returns whether the list currently has an action or a continuation cursor. */
function shouldShowButton(state: PaginationControlState): boolean {
    const hasContinuation = state.initialized && state.cursor !== null;
    return state.retryAvailable || hasContinuation || state.controlRequestPending;
}

/** Updates the native button's visibility, busy state, accessible name, and icon state. */
function updateButton(
    button: HTMLButtonElement | undefined,
    state: PaginationControlState,
    buttonVisible: boolean,
): void {
    if(!button) {
        return;
    }

    const loading = state.loading || state.controlRequestPending;
    const retry = state.retryAvailable && !loading;
    const label = getButtonLabel(loading, retry);
    button.hidden = !buttonVisible;
    button.disabled = loading;
    button.setAttribute("aria-label", label);
    button.title = label;
    button.setAttribute("aria-busy", String(loading));

    setChildHidden(button, "[data-paginated-list-more-icon]", loading || retry);
    setChildHidden(button, "[data-paginated-list-retry-icon]", loading || !retry);
    setChildHidden(button, "[data-paginated-list-loading-indicator]", !loading);
}

/** Chooses the accessible name for the currently rendered pagination action. */
function getButtonLabel(loading: boolean, retry: boolean): string {
    if(loading) {
        return "Loading records";
    }

    if(retry) {
        return "Retry loading records";
    }

    return "Show more";
}

/** Hides one decorative icon or loading indicator when it is not the active state. */
function setChildHidden(button: HTMLButtonElement, selector: string, hidden: boolean): void {
    const child = button.querySelector<HTMLElement>(selector);

    if(child) {
        child.hidden = hidden;
    }
}

/** Updates the owner-specific live region without changing its DOM identity. */
function updateStatus(status: HTMLElement | undefined, message: string, isError: boolean): void {
    if(!status) {
        return;
    }

    status.textContent = message;
    status.setAttribute("role", isError ? "alert" : "status");
}

/** Removes the controls margin block when it contains neither a button nor feedback. */
function updateControlsGroup(
    controls: HTMLElement | undefined,
    buttonVisible: boolean,
    message: string,
): void {
    if(!controls) {
        return;
    }

    controls.hidden = !buttonVisible && message.length === 0;
}

/** Finds one pagination element by its exact owner ID. */
function findOwnedElement<T extends HTMLElement>(attribute: string, ownerId: string): T | undefined {
    return Array.from(document.querySelectorAll<T>(`[${ attribute }]`))
        .find(element => element.getAttribute(attribute) === ownerId);
}
