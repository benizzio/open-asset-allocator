import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { handlebarsInfra } from "./handlebars";
import { HtmxInfra } from "./htmx";
import router from "./routing/router";
import * as bootstrap from "bootstrap";
import chart from "./chart/chart";
import { CustomEventHandler } from "./infra-types";
import DomInfra from "./dom";

const ROUTER_BOOT_DELAY_MS = 500;
let routerBootTimeoutId: number | undefined;

function bootRouterDebouncing() {

    // Debounce router boot to avoid repeated immediate calls on rapid DOM settle events.
    if(routerBootTimeoutId) {
        clearTimeout(routerBootTimeoutId);
    }

    routerBootTimeoutId = window.setTimeout(() => {
        routerBootTimeoutId = undefined;
        router.boot();
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
    router.bindDescendants(target);
    DomInfra.bindDescendants(target);
    chart.loadDescendantCharts(target);

    bootRouterDebouncing();
};

/**
 * Component that controls the multiple external libraries and its components to the desired behaviour of the
 * application.
 */
export const Infra = {

    init: (afterRequestErrorHandler: CustomEventHandler) => {

        Chart.register(...registerables, ChartDataLabels);

        window.Handlebars = handlebarsInfra.register();
        window["HandlebarsUtils"] = handlebarsInfra.utils;

        DomInfra.bindGlobalFunctions();

        const onPageLoad = () => {
            router.init(window);
            HtmxInfra.init(DOM_SETTLING_BEHAVIOR_EVENT_HANDLER, afterRequestErrorHandler);
        };
        document.addEventListener("DOMContentLoaded", onPageLoad);

        return { bootstrap };
    },
};
