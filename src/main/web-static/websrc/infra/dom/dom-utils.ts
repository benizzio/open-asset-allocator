const contextDataCache = new WeakMap<HTMLElement, unknown>();

function getCacheableContextData(
    contextDataElement: HTMLElement,
): unknown {

    if(!contextDataElement) {
        return null;
    }

    let elementData: unknown;

    if(contextDataCache.has(contextDataElement)) {
        elementData = contextDataCache.get(contextDataElement);
    }
    else {
        elementData = JSON.parse(contextDataElement.textContent);
        contextDataCache.set(contextDataElement, elementData);
    }

    return elementData;
}

function queryFirst(selector: string): HTMLElement {
    return document.querySelector(selector);
}

const DomUtils = {
    queryFirst,
    queryAll(selector: string): NodeListOf<HTMLElement> {
        return document.querySelectorAll(selector);
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
    getContextDataFromRoot(contextDataSelector: string): unknown {
        const dataElement = queryFirst(contextDataSelector);
        return getCacheableContextData(dataElement);
    },
};

export default DomUtils;