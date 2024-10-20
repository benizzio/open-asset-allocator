const DomUtils = {
    queryFirst(selector: string): HTMLElement {
        return document.querySelector(selector);
    },
    queryAll(selector: string): NodeListOf<HTMLElement> {
        return document.querySelectorAll(selector);
    },
};

export default DomUtils;