export type EventDetail = {
    routerPathData?: { [key: string]: unknown };
    [key: string]: unknown;
};

const configEnhancedRequestEventListener = (event: CustomEvent) => {
    replaceRequestPathParams(event);
};

/**
 * Enhances the htmx request by replacing path parameters from alternative sources.
 *
 * @param event - The htmx event
 */
function replaceRequestPathParams(event: CustomEvent) {

    const requestPath = event.detail.path as string;

    if(requestPath.includes(":")) {
        replaceFromEventChain(event);
        replaceFromFormData(event);
    }
}

/**
 * Enhances the htmx request by replacing path parameters from the event chain
 * (e.g., from the event that triggered the current one).
 *
 * @param event - The htmx event, where the triggering event may contain
 * { detail: { routerPathData: { [key: string]: unknown } } }
 */
function replaceFromEventChain(event: CustomEvent) {

    const triggeringEvent = event.detail?.triggeringEvent as CustomEvent;
    const detail = triggeringEvent?.detail as EventDetail;

    if(detail?.routerPathData) {

        let path = event.detail.path as string;

        for(const key in detail.routerPathData) {
            path = path.replace(`:${ key }`, detail.routerPathData[key] as string);
        }
        event.detail.path = path;
    }
}

/**
 * Enhances the htmx request by replacing path parameters from form data values.
 *
 * @param event - The htmx event
 */
function replaceFromFormData(event: CustomEvent) {

    const formData = event.detail.formData as FormData;
    const requestPath = event.detail.path as string;

    const splittedRequestPath = requestPath.split("/");

    const resolvedSplittedRequestPath = splittedRequestPath.map((pathPart) => {

        let processedPart = pathPart;

        if(processedPart.startsWith(":")) {

            const paramName = processedPart.substring(1);

            if(formData.has(paramName)) {
                const paramValue = formData.get(paramName);
                processedPart = typeof paramValue === "string" ? paramValue : processedPart;
                formData.delete(paramName);
            }
        }

        return processedPart;
    });

    event.detail.path = resolvedSplittedRequestPath.join("/");
}

function addEventListeners(domSettlingBehaviorEventHandler: (event: CustomEvent) => void) {
    document.addEventListener("htmx:configRequest", configEnhancedRequestEventListener);
    document.body.addEventListener("htmx:afterSettle", domSettlingBehaviorEventHandler);
}

export const htmxInfra = {
    /**
     * Initializes the htmx infrastructure of the application.
     *
     * @param domSettlingBehaviorEventHandler - The handler for the default DOM settling behavior event.
     * Will be applied to the body and be triggered in after settiling of any child element.
     */
    init(domSettlingBehaviorEventHandler: (event: CustomEvent) => void) {
        addEventListeners(domSettlingBehaviorEventHandler);
    },
};