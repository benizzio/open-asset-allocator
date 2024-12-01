import * as handlebarsModule from "handlebars";
import { Chart, registerables } from "chart.js";
import ChartDataLabels from "chartjs-plugin-datalabels";
import { registerHandlebarsChartHelper } from "./handlebars-chart";
import { registerHandlebarsFormatHelper } from "./handlebars-format";
import { registerHandlebarsObjectHelpers } from "./handlebars-object";

export const handlebars ={
    register: () => {
        Chart.register(...registerables, ChartDataLabels);
        registerHandlebarsChartHelper();
        registerHandlebarsFormatHelper();
        registerHandlebarsObjectHelpers();
        return handlebarsModule;
    },
};