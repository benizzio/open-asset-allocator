/**
 * Loads, tracks, and disposes Chart.js instances attached to HTMX-managed DOM subtrees.
 * @author GPT-6 Luna
 */

import { Chart, ChartType } from "chart.js";
import { buildChartOptions } from "./chart-options";
import { CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent, LocalChartOptions } from "./chart-types";
import DomUtils from "../dom/dom-utils";
import { convertUnidimensionalDatasetBackgroundToPattern, getPieDoughnutChartColorScale } from "./chart-utils";

const chartContentRepo = new Map<string, ChartContent>;

function saveChartContent(chartId: string, content: ChartContent): void {
    chartContentRepo.set(chartId, content);
}

function getChartContent(chartId: string): ChartContent {
    return chartContentRepo.get(chartId);
}

function getChartContentFromChart(chart: Chart) {
    const chartId = chart.canvas.id;
    const chartContent = getChartContent(chartId);
    return { chartId, chartContent };
}

/** Creates a chart once for a canvas that has registered chart content. */
function loadChart(canvas: HTMLCanvasElement): void {

    const id = canvas.id;
    const content = getChartContent(id);

    if(!content || Chart.getChart(canvas)) {
        return;
    }

    const interactionObserverCallback = content.interactionObserverCallback;
    const chartInteractions = content.chartInteractions;

    const localOptionsElementId = canvas.getAttribute(CHART_OPTIONS_JSON_ELEMENT_ID);
    const localOptions = DomUtils.getContextDataFromRoot("#" + localOptionsElementId) as LocalChartOptions;
    const chartType = localOptions.type;


    const options = buildChartOptions(
        localOptions,
        chartInteractions,
        interactionObserverCallback,
    );

    new Chart(canvas, {
        type: chartType as ChartType,
        data: content.chartDataSource.getChartData(),
        options: options,
    });
}

/** Destroys charts below a node before HTMX removes their canvases. */
function destroyDescendantCharts(element: HTMLElement): void {
    const canvases = Array.from(element.querySelectorAll<HTMLCanvasElement>(`canvas[${ CHART_ATTRIBUTE }]`));

    if(element.matches(`canvas[${ CHART_ATTRIBUTE }]`)) {
        canvases.unshift(element as HTMLCanvasElement);
    }

    canvases.forEach(canvas => {
        const chartId = canvas.id;
        Chart.getChart(canvas)?.destroy();

        globalThis.setTimeout(() => {
            const replacement = document.getElementById(chartId);

            if(!(replacement instanceof HTMLCanvasElement && replacement.hasAttribute(CHART_ATTRIBUTE))) {
                chartContentRepo.delete(chartId);
            }
        }, 0);
    });
}

const chart = {
    saveChartContent,
    getChartContent,
    getChartContentFromChart,
    destroyDescendantCharts,
    loadDescendantCharts(element: HTMLElement) {
        element.querySelectorAll(`canvas[${ CHART_ATTRIBUTE }]`).forEach((canvas: HTMLCanvasElement) => {
            loadChart(canvas);
        });
    },
    getPieDoughnutChartColorScale,
    convertUnidimensionalDatasetBackgroundToPattern,
};

export default chart;
