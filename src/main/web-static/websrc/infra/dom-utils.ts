const DomUtils = {
    queryFirst(selector: string): HTMLElement {
        return document.querySelector(selector);
    },
    queryAll(selector: string): NodeListOf<HTMLElement> {
        return document.querySelectorAll(selector);
    },
    queryAllInDescendants(element: HTMLElement, selector: string): NodeListOf<HTMLElement> {
        return element.querySelectorAll(selector);
    },
    wasElementRemoved(element: HTMLElement) {
        return !document.body.contains(element);
    },
};

export default DomUtils;