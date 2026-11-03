/**
 * Public API for independent HTMX cursor-paginated lists.
 * Set `data-paginated-list-context-selector` on an owner to a document-wide CSS selector
 * matching an input whose value identifies its current context. A changed value rejects
 * stale responses and resets the cursor on the next request. Without the attribute, the
 * owner has a constant empty context. For example, use `[name='accountId']` for an account list.
 * @module components/paginated-list
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author GPT-6.1 Terra
 */

export { default } from "./pagination";
export { PaginationMode } from "./pagination";
