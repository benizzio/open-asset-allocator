import { HtmxBeforeSwapDetails, HtmxRequestConfig, HtmxResponseInfo } from "htmx.org";
import { bindHTMXTransformResponseInDescendants, htmxTransformResponse } from "./binding-htmx-transform-response";
import { addReadyConditionToWaitingElement, bindHTMXWaitForReadyInDescendants } from "./binding-htmx-wait-for-ready";
import { CustomEventHandler } from "../infra-types";
import InfraTypesUtils from "../infra-types-utils";
import Router from "../routing";
import { logger, LogLevel } from "../logging";
import { APIErrorResponse } from "../../api/api";

const NULL_IF_EMPTY_ATTRIBUTE = "data-null-if-empty";
const suppressedAfterRequestErrorNotifications = new WeakSet<XMLHttpRequest>();
const afterRequestErrorEventHandlers = new WeakMap<CustomEventHandler, CustomEventHandler>();

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

    replaceRequestPathParamsFromEventChain(event);
    replaceRequestPathParamsFromFormData(event);
    replaceRequestPathParamsFromCurrentRoute(event);

    const requestPath = event.detail.path as string;

    if(requestPath.includes(Router.NAVIGO_PATH_PARAM_PREFIX)) {
        logger(LogLevel.WARN, "Could not resolve all path parameters for htmx request", event.detail.path);
    }
}

/**
 * Enhances the htmx request by replacing path parameters from the event chain
 * (e.g., from the event that triggered the current one).
 *
 * @param event - The htmx event, where the triggering event may contain
 * { detail: { routerPathData: { [key: string]: unknown } } }
 */
function replaceRequestPathParamsFromEventChain(event: CustomEvent) {

    const requestPath = event.detail.path as string;

    if(!requestPath.includes(Router.NAVIGO_PATH_PARAM_PREFIX)) {
        return;
    }

    const triggeringEvent = event.detail?.triggeringEvent as CustomEvent;
    const detail = triggeringEvent?.detail as RequestConfigEventDetail;

    if(detail?.routerPathData) {

        let path = requestPath;

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
function replaceRequestPathParamsFromFormData(event: CustomEvent) {

    const requestPath = event.detail.path as string;

    if(!requestPath.includes(Router.NAVIGO_PATH_PARAM_PREFIX)) {
        return;
    }

    const formData = event.detail.formData as FormData;

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

function replaceRequestPathParamsFromCurrentRoute(event: CustomEvent) {

    const requestPath = event.detail.path as string;

    if(!requestPath.includes(Router.NAVIGO_PATH_PARAM_PREFIX)) {
        return;
    }

    let path = requestPath;
    path = Router.buildParameterizedDestinationPathFromCurrentLocationContext(path);
    event.detail.path = path;
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

/** Returns an idempotent afterRequest wrapper that skips only explicitly suppressed request notifications.
 *
 * @author GPT-6 Luna
 */
function getAfterRequestErrorEventHandler(handler: CustomEventHandler): CustomEventHandler {
    const cachedHandler = afterRequestErrorEventHandlers.get(handler);

    if(cachedHandler) {
        return cachedHandler;
    }

    const wrappedHandler: CustomEventHandler = event => {
        const detail = (event as CustomEvent<AfterRequestEventDetail>).detail;

        if(detail?.xhr && suppressedAfterRequestErrorNotifications.has(detail.xhr)) {
            return;
        }

        handler(event);
    };

    afterRequestErrorEventHandlers.set(handler, wrappedHandler);
    return wrappedHandler;
}

function addEventListeners(
    domSettlingBehaviorEventHandler: CustomEventHandler,
    afterRequestErrorHandler: CustomEventHandler,
) {

    document.addEventListener("htmx:configRequest", configEnhancedRequestEventListener);

    document.body.addEventListener("htmx:afterRequest", getAfterRequestErrorEventHandler(afterRequestErrorHandler));

    // Add settling behaviour needed for HTMX own bindings
    const afterSettleCustomEventHandler = (event: CustomEvent) => {

        const eventTarget = event.target as HTMLElement;

        // Wait-for-ready gates must be bound before any other binding that may trigger HTMX requests
        bindHTMXWaitForReadyInDescendants(eventTarget);

        domSettlingBehaviorEventHandler(event);
        bindHTMXTransformResponseInDescendants(eventTarget);
    };

    document.body.addEventListener("htmx:afterSettle", afterSettleCustomEventHandler);
}

function toErrorResponse(eventDetail: AfterRequestEventDetail): APIErrorResponse | undefined {

    const contentType = eventDetail.xhr.getResponseHeader("content-type");

    if(contentType?.includes("application/json")) {
        return InfraTypesUtils.toErrorResponse(eventDetail.xhr.response);
    }

    return undefined;
}

export const HtmxInfra = {

    /** Suppresses only the global error notification for one stale HTMX request, without stopping event propagation.
     *
     * Call synchronously from that request's `htmx:afterRequest` handler when its owning UI has become stale.
     * Other listeners still receive the event; only the shared error-notification callback skips this XHR.
     *
     * @param xhr - The failed request whose error should not be shown because its owning UI is stale.
     * @example
     * ```ts
     * if(isStaleRequest) {
     *     HtmxInfra.suppressAfterRequestErrorNotification(event.detail.xhr);
     * }
     * ```
     * @author GPT-6 Luna
     */
    suppressAfterRequestErrorNotification(xhr: XMLHttpRequest): void {
        suppressedAfterRequestErrorNotifications.add(xhr);
    },

    /**
     * Initializes the htmx infrastructure of the application.
     * All handlers will be applied to the body and be triggered in after the events of any child element.
     *
     * @param domSettlingBehaviorEventHandler - The handler for the default DOM settling behavior event.
     * @param afterRequestErrorHandler - The handler for after request error events.
     */
    init(domSettlingBehaviorEventHandler: CustomEventHandler, afterRequestErrorHandler: CustomEventHandler) {
        addEventListeners(domSettlingBehaviorEventHandler, afterRequestErrorHandler);
    },

    htmxTransformResponse,
    toErrorResponse,
    addReadyConditionToWaitingElement,
};
