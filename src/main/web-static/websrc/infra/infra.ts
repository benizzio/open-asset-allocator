import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { handlebarsInfra } from "./handlebars/handlebars";
import { htmxInfra } from "./htmx";
import router from "./routing/router";
import * as bootstrap from "bootstrap";

//eslint-disable-next-line
const browserGlobal = (window as any);

export const infra = {

    init: () => {

        Chart.register(...registerables, ChartDataLabels);

        browserGlobal.Handlebars = handlebarsInfra.register();

        htmxInfra.init();

        const onPageLoad = () => {
            router.init();
        };
        document.addEventListener("DOMContentLoaded", onPageLoad);

        document.body.addEventListener("htmx:afterSettle", function() {
            router.boot();
        });

        return { bootstrap };
    },
};
