import * as clientSideTemplates from "htmx-ext-client-side-templates";
import * as handlebars from "handlebars";
import * as bootstrap from "bootstrap";
import router from "./infra/routing";
import { Chart, registerables } from "chart.js";
import chart from "./infra/chart";
import handlebarsChart from "./infra/handlebars-chart";
import ChartDataLabels from "chartjs-plugin-datalabels";
import handlebarsFormat from "./infra/handlebars-format";

Chart.register(...registerables, ChartDataLabels);
handlebarsChart.registerHandlebarsChartHelper();
handlebarsFormat.registerHandlebarsFormatHelper();

//eslint-disable-next-line
const browserGlobal = (window as any);

browserGlobal.router = router;
browserGlobal.Handlebars = handlebars;

browserGlobal.loadCharts = (element: HTMLElement) => {
    element.querySelectorAll("canvas").forEach((canvas) => {
        chart.loadChart(canvas);
    });
};

const onPageLoad = () => {
    router.bindHTMXRouting();
    router.resolveBrowserRoute();
};
document.addEventListener("DOMContentLoaded", onPageLoad);

export default { clientSideTemplates, handlebars, bootstrap };
