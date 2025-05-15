export type EventDetail = {
    routerPathData?: { [key: string]: unknown };
    [key: string]: unknown;
};

const configRequestEventListener = (event: CustomEvent) => {
    replaceRequestPathParams(event);
};

function replaceRequestPathParams(event: CustomEvent) {

    const requestPath = event.detail.path as string;

    if(requestPath.includes(":")) {
        replaceFromEventChain(event);
        replaceFromFormData(event);
    }
}

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
    document.addEventListener("htmx:configRequest", configRequestEventListener);
    document.body.addEventListener("htmx:afterSettle", domSettlingBehaviorEventHandler);
}

export const htmxInfra = {
    init(domSettlingBehaviorEventHandler: (event: CustomEvent) => void) {
        addEventListeners(domSettlingBehaviorEventHandler);
    },
};