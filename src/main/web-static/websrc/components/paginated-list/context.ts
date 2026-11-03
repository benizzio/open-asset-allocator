/**
 * Resolves a list's consumer-configured context identity for cursor lifecycle checks.
 * @module components/paginated-list/context
 * @author GPT-6.1 Terra
 */

const CONTEXT_SELECTOR_ATTRIBUTE = "data-paginated-list-context-selector";

/**
 * Reads the input selected by the owner's document-wide context selector. An omitted selector or
 * missing input yields the empty identity. Pagination rejects responses when this value changes
 * and resets the cursor on the next request.
 *
 * @param owner - The list with an optional context selector attribute.
 * @returns The selected input's current value, or an empty string.
 * @example `getContextId(owner)` reads the input selected by `[name='accountId']` when that
 * selector is configured as `data-paginated-list-context-selector` on the owner.
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
export function getContextId(owner: HTMLElement): string {
    const selector = owner.getAttribute(CONTEXT_SELECTOR_ATTRIBUTE);
    return selector ? document.querySelector<HTMLInputElement>(selector)?.value ?? "" : "";
}
