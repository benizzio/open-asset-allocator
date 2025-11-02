import * as handlebars from "handlebars";

const PARTIALS_REGISTRY = new Map<string, HTMLElement>();

/**
 * Registers a Handlebars partial template to a container element.
 * The partial template is identified by its ID, and the main container is identified by its ID.
 * The partial will be unregistered when the main container is removed from the DOM.
 *
 * @param containerId - The ID of the main container element.
 * @param partialTemplateId - The ID of the element containing the Handlebars partial template to register.
 */
export function registerPartialToContainer(containerId: string, partialTemplateId: string) {

    const mainContainerElement = window[containerId];
    const partialTemplateElement = window[partialTemplateId];
    const partialTemplateHTML = partialTemplateElement.innerHTML;

    if(!PARTIALS_REGISTRY.has(partialTemplateId)) {

        handlebars.registerPartial(partialTemplateId, partialTemplateHTML);

        const allocationMapObserver = new MutationObserver((_, observer) => {
            if(!document.body.contains(mainContainerElement)) {
                console.info(`Main container removed, unregistering ${ partialTemplateId } partial and observer`);
                observer.disconnect();
                Handlebars.unregisterPartial(partialTemplateId);
                PARTIALS_REGISTRY.delete(partialTemplateId);
            }
        });
        allocationMapObserver.observe(document, { childList: true, subtree: true });

        PARTIALS_REGISTRY.set(partialTemplateId, partialTemplateElement);
    }

}