import router from "./routing/router";
import { domInfra } from "./dom/dom";
import chart from "./chart/chart";

export type EventDetail = {
    routerPathData?: { [key: string]: unknown };
    [key: string]: unknown;
};

const configRequestEventListener = (event: CustomEvent) => {
    replaceRequestPathParams(event);
};

const afterSettleEventListener = (event: CustomEvent) => {
    const target = event.target as HTMLElement;
    router.bindDescendants(target);
    domInfra.bindDescendants(target);
    chart.loadDescendantCharts(target);
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

function addEventListeners() {
    document.addEventListener("htmx:configRequest", configRequestEventListener);
    document.body.addEventListener("htmx:afterSettle", afterSettleEventListener);
}

export const htmxInfra = {
    init() {
        addEventListeners();
    },
};