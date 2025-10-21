import * as handlebars from "handlebars";
import chart from "../chart/chart";
import {
    CHART_ATTRIBUTE,
    CHART_OPTIONS_JSON_ELEMENT_ID,
    LocalChartOptions,
    MeasuramentUnit,
} from "../chart/chart-types";
import DomUtils from "../dom/dom-utils";
import Application from "../../application";

const handlebarsChartHelper = (
    source: object,
    chartDataType: string,
    optionsJSonElementId: string,
    idPrefix: string,
    idSuffix: string,
    contextDataSelector: string,
) => {

    const contextData = DomUtils.getContextDataFromRoot(contextDataSelector);
    const chartContent = Application.chartContents.toChartContent(chartDataType, source, contextData);

    const options = DomUtils.getContextDataFromRoot("#" + optionsJSonElementId) as LocalChartOptions;
    options.measuramentUnit = options.measuramentUnit || MeasuramentUnit.CURRENCY;

    const id = `${ idPrefix }-${ idSuffix }`;

    chart.saveChartContent(id, chartContent);

    return `<canvas id="${ id }" ${ CHART_OPTIONS_JSON_ELEMENT_ID }="${ optionsJSonElementId }" ${ CHART_ATTRIBUTE }>
            </canvas>`;
};

export function registerHandlebarsChartHelper() {
    handlebars.registerHelper("chart", handlebarsChartHelper);
}
