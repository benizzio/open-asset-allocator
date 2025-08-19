import DomUtils from "../dom/dom-utils";
import { BeforeSwapEventDetail } from "./htmx";
import { logger, LogLevel } from "../logging";

// =============================================================================
// HTMX response transform binding
// modifies the response body of HTMX requests
// using a specified function on the beforeSwap event
// allows configuration of a validating RegExp to call function only when needed
// =============================================================================

const HTMX_TRANSFORM_RESPONSE_ATTRIBUTE = "data-hx-transform-response";
const HTMX_TRANSFORM_RESPONSE_ON_ROUTE_MATCHING_ATTRIBUTE = "data-hx-transform-response-on-route-matching";
const HTMX_TRANSFORM_RESPONSE_BOUND_FLAG = "data-hx-trigger-on-route-bound";

const TRANSFORM_RESPONSE_FUNCTION_MAP = new Map<string, (responseBody: string) => string>();

function extractRouteRegExpForTransform(element: HTMLElement) {
    const transformResponseRouteAttribute =
        element.getAttribute(HTMX_TRANSFORM_RESPONSE_ON_ROUTE_MATCHING_ATTRIBUTE);
    return transformResponseRouteAttribute
        ? new RegExp(transformResponseRouteAttribute)
        : null;
}

function transformResponse(eventDetail: BeforeSwapEventDetail, transformFunctionKey: string) {

    const originalServerResponseJSON = eventDetail.serverResponse;
    const transformFunction = TRANSFORM_RESPONSE_FUNCTION_MAP.get(transformFunctionKey);

    if(!transformFunction) {
        throw new Error(`No transform function found for key "${ transformFunctionKey }".`);
    }

    eventDetail.serverResponse = transformFunction(originalServerResponseJSON);
}

function bindHTMXTransformResponseElement(element: HTMLElement) {

    const transformResponseRegExp = extractRouteRegExpForTransform(element);

    element.addEventListener("htmx:beforeSwap", (event: CustomEvent) => {

        const eventDetail = event.detail as BeforeSwapEventDetail;
        const eventRequestPath = eventDetail.pathInfo.finalRequestPath;

        if(eventDetail.isError || (!transformResponseRegExp || !eventRequestPath.match(transformResponseRegExp))) {
            return;
        }

        const transformFunctionKey = element.getAttribute(HTMX_TRANSFORM_RESPONSE_ATTRIBUTE);

        try {
            transformResponse(eventDetail, transformFunctionKey);
        } catch(error) {
            logger(LogLevel.ERROR, "Error transforming response for element", element, error);
            return;
        }
    });
}

function bindHTMXTransformResponseElements(elementsToBind: NodeListOf<HTMLElement>) {
    elementsToBind.forEach((element) => {
        logger(LogLevel.INFO, "Binding HTMX transform response for element", element);
        bindHTMXTransformResponseElement(element);
        element.setAttribute(HTMX_TRANSFORM_RESPONSE_BOUND_FLAG, "true");
    });
}

export function bindHTMXTransformResponseInDescendants(element: HTMLElement) {
    const elementsToBind = DomUtils.queryAllInDescendants(
        element,
        `[${ HTMX_TRANSFORM_RESPONSE_ATTRIBUTE }]:not([${ HTMX_TRANSFORM_RESPONSE_BOUND_FLAG }])`,
    );
    bindHTMXTransformResponseElements(elementsToBind);
}

function registerTransformResponseFunction(
    functionName: string,
    transformFunction: (responseBody: string) => string,
) {

    if(TRANSFORM_RESPONSE_FUNCTION_MAP.has(functionName)) {
        logger(
            LogLevel.WARN,
            `Transform response function with name ${ functionName } already exists and will be replaced.`,
        );
    }

    TRANSFORM_RESPONSE_FUNCTION_MAP.set(functionName, transformFunction);
}

export const htmxTransformResponse = { registerTransformResponseFunction };