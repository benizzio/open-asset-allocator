import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { handlebarsInfra } from "./handlebars";
import { htmxInfra } from "./htmx/htmx";
import router from "./routing/router";
import * as bootstrap from "bootstrap";
import chart from "./chart/chart";
import { CustomEventHandler } from "./infra-types";
import DomInfra from "./dom";

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
    router.boot();
    DomInfra.bindDescendants(target);
    chart.loadDescendantCharts(target);
};

/**
 * Component that controls the multiple external libraries and its components to the desired behaviour of the
 * application.
 */
export const infra = {

    init: () => {

        Chart.register(...registerables, ChartDataLabels);

        window.Handlebars = handlebarsInfra.register();
        window["HandlebarsUtils"] = handlebarsInfra.utils;

        DomInfra.bindGlobalFunctions();

        htmxInfra.init(DOM_SETTLING_BEHAVIOR_EVENT_HANDLER);

        const onPageLoad = () => {
            router.init(window);
        };
        document.addEventListener("DOMContentLoaded", onPageLoad);

        return { bootstrap };
    },
};
