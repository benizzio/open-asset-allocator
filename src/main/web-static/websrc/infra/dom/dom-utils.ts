const contextDataCache = new WeakMap<HTMLElement, unknown>();

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

function addRemoveObserver(contextDataElement: HTMLElement) {
    const observer = new MutationObserver(() => {
        if(DomUtils.wasElementRemoved(contextDataElement)) {
            contextDataCache.delete(contextDataElement);
            observer.disconnect();
        }
    });
    observer.observe(document.body, { childList: true, subtree: true });
}

function queryFirst(selector: string): HTMLElement {
    return document.querySelector(selector);
}

const DomUtils = {
    queryFirst,
    queryAll(selector: string): NodeListOf<HTMLElement> {
        return document.querySelectorAll(selector);
    },
    queryFirstInDescendants(element: HTMLElement, selector: string): HTMLElement {
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
};

export default DomUtils;