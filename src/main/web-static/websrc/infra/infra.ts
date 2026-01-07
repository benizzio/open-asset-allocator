import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { handlebarsInfra } from "./handlebars";
import { HtmxInfra } from "./htmx";
import Router from "./routing";
import * as bootstrap from "bootstrap";
import chart from "./chart/chart";
import { CustomEventHandler } from "./infra-types";
import DomInfra from "./dom";
import { logger, LogLevel } from "./logging";

type GeneralErrorHandler = (error: Error) => void;

const ROUTER_BOOT_DELAY_MS = 500;
let routerBootTimeoutId: number | undefined;

function bootRouterDebouncing() {

    // Debounce router boot to avoid repeated immediate calls on rapid DOM settle events.
    if(routerBootTimeoutId) {
        clearTimeout(routerBootTimeoutId);
    }

    routerBootTimeoutId = window.setTimeout(() => {
        routerBootTimeoutId = undefined;
        Router.boot();
    }, ROUTER_BOOT_DELAY_MS);
}

/**
 * Ties multiple components of the application infrastructure to HTMX async DOM behaviour.
 * When the DOM "settles" after HTMX modifying it, this function controls and binds other components from different
 * libraries, when needed.
 *
 * @param event - The event that is triggered when the DOM settles.
 */
const DOM_SETTLING_BEHAVIOR_EVENT_HANDLER: CustomEventHandler = (event: CustomEvent) => {

    const target = event.target as HTMLElement;
    Router.bindDescendants(target);
    DomInfra.bindDescendants(target);
    chart.loadDescendantCharts(target);

    bootRouterDebouncing();
};

function handleError(
    errorMessage: string,
    error: Error,
    uncaughtErrorHandler: (error: Error) => void,
    message: Event | string,
) {

    logger(LogLevel.ERROR, errorMessage, error);

    if(error) {
        uncaughtErrorHandler(error);
    }
    else {
        uncaughtErrorHandler(new Error(String(message)));
    }
}

/**
 * Sets up global error handlers to capture and log any uncaught errors in the application.
 * Handles both synchronous errors via `window.onerror` and unhandled promise rejections
 * via `unhandledrejection` event.
 *
 * @param uncaughtErrorHandler - Custom handler to be called when an uncaught error occurs.
 *
 * @author GitHub Copilot
 */
function setupGlobalErrorHandler(uncaughtErrorHandler: GeneralErrorHandler): void {

    window.onerror = (message, source, lineno, colno, error) => {
        const errorMessage = `Uncaught error: ${ message } at ${ source }:${ lineno }:${ colno }`;
        handleError(errorMessage, error, uncaughtErrorHandler, message);
        return true;
    };

    window.addEventListener("unhandledrejection", (event: PromiseRejectionEvent) => {
        const reason = event.reason;
        const errorMessage = "Unhandled promise rejection:";
        handleError(errorMessage, reason instanceof Error ? reason : null, uncaughtErrorHandler, event);
    });
}

/**
 * Component that controls the multiple external libraries and its components to the desired behaviour of the
 * application.
 */
export const Infra = {

    init: (afterRequestErrorHandler: CustomEventHandler, generalUncaughtErrorHandler: GeneralErrorHandler) => {

        Chart.register(...registerables, ChartDataLabels);

        window.Handlebars = handlebarsInfra.register();
        window["HandlebarsUtils"] = handlebarsInfra.utils;

        DomInfra.bindGlobalFunctions();
        setupGlobalErrorHandler(generalUncaughtErrorHandler);

        const onPageLoad = () => {
            Router.init(window);
            HtmxInfra.init(DOM_SETTLING_BEHAVIOR_EVENT_HANDLER, afterRequestErrorHandler);
        };
        document.addEventListener("DOMContentLoaded", onPageLoad);

        return { bootstrap };
    },
};
