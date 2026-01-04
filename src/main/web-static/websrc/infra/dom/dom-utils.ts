const contextDataCache = new WeakMap<HTMLElement, unknown>();

const observedElements = new Set<HTMLElement>();
let sharedObserver: MutationObserver | null = null;

function getCacheableContextData(
    contextDataElement: HTMLElement,
): unknown {

    if(!contextDataElement) {
        return null;
    }

    let elementData: unknown = null;

    if(contextDataCache.has(contextDataElement)) {
        elementData = contextDataCache.get(contextDataElement);
    }
    else if(contextDataElement.textContent.trim()) {
        elementData = JSON.parse(contextDataElement.textContent);
        contextDataCache.set(contextDataElement, elementData);
        addRemoveObserver(contextDataElement);
    }

    return elementData;
}

function ensureSharedObserver() {

    if(!sharedObserver) {
        sharedObserver = new MutationObserver(() => {
            observedElements.forEach(element => {
                if(DomUtils.wasElementRemoved(element)) {
                    contextDataCache.delete(element);
                    observedElements.delete(element);
                }
            });

            if(observedElements.size === 0) {
                sharedObserver?.disconnect();
                sharedObserver = null;
            }
        });
        sharedObserver.observe(document.body, { childList: true, subtree: true });
    }
}

function addRemoveObserver(contextDataElement: HTMLElement) {
    observedElements.add(contextDataElement);
    ensureSharedObserver();
}

function queryFirst(selector: string): HTMLElement {
    return document.querySelector(selector);
}

function escapeHtmlValue(value: string): string {
    return (value ?? "")
        .replace(/&/g, "&amp;")
        .replace(/</g, "&lt;")
        .replace(/>/g, "&gt;")
        .replace(/"/g, "&quot;")
        .replace(/'/g, "&#39;");
}

const DomUtils = {

    //TODO remove simple queries from here and use directly
    queryFirst,

    queryAll(selector: string): NodeListOf<HTMLElement> {
        return document.querySelectorAll(selector);
    },

    queryFirstInDescendants(element: HTMLElement, selector: string): HTMLElement | null {
        return element.querySelector(selector);
    },

    queryAllInDescendants(element: HTMLElement, selector: string): NodeListOf<HTMLElement> {
        return element.querySelectorAll(selector);
    },

    queryDirectDescendants(element: HTMLElement, selector: string): NodeListOf<HTMLElement> {
        return element.querySelectorAll(`:scope > ${ selector }`);
    },

    wasElementRemoved(element: HTMLElement) {
        return !document.body.contains(element);
    },

    /**
     * Gets context data from root document element matching the given selector.
     * Caches the parsed data for future retrievals.
     *
     * Context data elements are expected to be script tags on the document with type "application/json" and content
     * relevant to the current structure of pages.
     *
     * @param contextDataSelector CSS selector to find the context data element in the root document.
     * @returns The parsed context data object, or null if the element is not found.
     */
    getContextDataFromRoot(contextDataSelector: string): unknown {
        const dataElement = queryFirst(contextDataSelector);
        return getCacheableContextData(dataElement);
    },

    /**
     * Escapes HTML special characters in untrusted text
     * to prevent DOM injection when rendering server-provided content.
     *
     * Usage Guidance:
     * - Call before inserting dynamic strings into innerHTML or template output.
     * - This is a conservative encoder; it targets &, <, >, " and '.
     * - Keeps input null-safe by treating null/undefined as an empty string.
     *
     * @param rawValue Raw text value that may contain unsafe characters.
     * @returns Safe, escaped string suitable for HTML contexts.
     *
     * @example
     * // "&lt;script&gt;alert(&quot;x&quot;)&lt;/script&gt;"
     * const safe = DomUtils.escapeHtml('<script>alert("x")</script>');
     *
     * @author GitHub Copilot
     */
    escapeHtml(rawValue: string): string {
        return escapeHtmlValue(rawValue);
    },
};

export default DomUtils;