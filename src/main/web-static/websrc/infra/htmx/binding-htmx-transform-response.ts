import DomUtils from "../dom/dom-utils";
import { BeforeSwapEventDetail } from "./htmx";
import { logger, LogLevel } from "../logging";

// =============================================================================
// HTMX response transform binding
// modifyies the response body of HTMX requests
// using a specified function on the beforeSwap event
// allows configuration of a validating RegExp to call function only when needed
// =============================================================================

const HTMX_TRANSFORM_RESPONSE_ATTRIBUTE = "data-hx-transform-response";
const HTMX_TRANSFORM_RESPONSE_ON_ROUTE_MATCHING_ATTRIBUTE = "data-hx-transform-response-on-route-matching";
const HTMX_TRANSFORM_RESPONSE_BOUND_FLAG = "data-hx-trigger-on-route-bound";

const TRANSFORM_RESPONSE_FUNCTION_MAP = new Map<string, (responseBody: string) => string>();

//TODO clean
export function bindHTMXTransformResponse(element: HTMLElement) {

    const elemementsToBind = DomUtils.queryAllInDescendants(
        element,
        `[${ HTMX_TRANSFORM_RESPONSE_ATTRIBUTE }]:not([${ HTMX_TRANSFORM_RESPONSE_BOUND_FLAG }])`,
    );

    elemementsToBind.forEach((element) => {

        logger(LogLevel.INFO, "Binding HTMX transform response for element", element);

        const transformResponseRouteAttribute =
            element.getAttribute(HTMX_TRANSFORM_RESPONSE_ON_ROUTE_MATCHING_ATTRIBUTE);

        const transformResponseRegExp = transformResponseRouteAttribute
            ? new RegExp(transformResponseRouteAttribute)
            : null;

        element.addEventListener("htmx:beforeSwap", (event: CustomEvent) => {

            const eventDetail = event.detail as BeforeSwapEventDetail;
            const eventRequestPath = eventDetail.pathInfo.finalRequestPath;

            if(eventDetail.isError || (!transformResponseRegExp || !eventRequestPath.match(transformResponseRegExp))) {
                return;
            }


            const originalServerResponseJSON = eventDetail.serverResponse;
            const transformFunctionKey = element.getAttribute(HTMX_TRANSFORM_RESPONSE_ATTRIBUTE);
            const transformFunction = TRANSFORM_RESPONSE_FUNCTION_MAP.get(transformFunctionKey);

            if(!transformFunction) {
                logger(LogLevel.WARN, "No transform function found for element", element);
                return;
            }

            eventDetail.serverResponse = transformFunction(originalServerResponseJSON);
        });

        element.setAttribute(HTMX_TRANSFORM_RESPONSE_BOUND_FLAG, "true");
    });
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