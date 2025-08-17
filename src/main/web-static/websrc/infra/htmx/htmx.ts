import { HtmxBeforeSwapDetails, HtmxRequestConfig, HtmxResponseInfo } from "htmx.org";
import { bindHTMXTransformResponse, htmxTransformResponse } from "./binding-htmx-transform-response";
import { CustomEventHandler } from "../infra-types";

const NULL_IF_EMPTY_ATTRIBUTE = "data-null-if-empty";

type EventDetail = { [key: string]: unknown; };

export type RequestConfigEventDetail =
    { routerPathData?: { [key: string]: unknown }; }
    & EventDetail
    & HtmxRequestConfig;

export type AfterRequestEventDetail = EventDetail & RequestConfigEventDetail & HtmxResponseInfo;

export type BeforeSwapEventDetail = EventDetail & HtmxBeforeSwapDetails;

const configEnhancedRequestEventListener = (event: CustomEvent) => {
    replaceRequestPathParams(event);
    prepareFormData(event);
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
    const detail = triggeringEvent?.detail as RequestConfigEventDetail;

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

function prepareFormData(event: CustomEvent) {

    const eventDetail = event.detail as EventDetail;

    if(eventDetail.verb === "post") {

        const formElement = event.target as HTMLFormElement;
        const formData = event.detail.formData as FormData;

        for(const element of Array.from(formElement.elements) as HTMLInputElement[]) {

            if(element.hasAttribute(NULL_IF_EMPTY_ATTRIBUTE) && !element.value) {
                formData.delete(element.name);
            }
        }
    }

}

function addEventListeners(domSettlingBehaviorEventHandler: CustomEventHandler) {

    document.addEventListener("htmx:configRequest", configEnhancedRequestEventListener);

    const afterSettleCustomEventHandler = (event: CustomEvent) => {
        domSettlingBehaviorEventHandler(event);
        const eventTarget = event.target as HTMLElement;
        bindHTMXTransformResponse(eventTarget);
    };
    document.body.addEventListener("htmx:afterSettle", afterSettleCustomEventHandler);
}

export const htmxInfra = {
    /**
     * Initializes the htmx infrastructure of the application.
     *
     * @param domSettlingBehaviorEventHandler - The handler for the default DOM settling behavior event.
     * Will be applied to the body and be triggered in after settling of any child element.
     */
    init(domSettlingBehaviorEventHandler: CustomEventHandler) {
        addEventListeners(domSettlingBehaviorEventHandler);
    },
    htmxTransformResponse,
};