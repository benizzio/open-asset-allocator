/**
 * Loads, tracks, and disposes Chart.js instances attached to HTMX-managed DOM subtrees.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author benizzio
 * @author GPT-6.1 Terra
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

/** Creates a chart once for a canvas that has registered chart content.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author benizzio
 * @author GPT-6.1 Terra
 */
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

/** Destroys charts below a node before HTMX removes their canvases.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 */
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

/**
 * Provides chart content registration, lookup, lifecycle management, and color helpers.
 * Register content by canvas ID before loading charts in a DOM subtree. Destroy charts
 * before removing their subtree; registered content is retained if a matching canvas
 * replaces the removed canvas before the deferred cleanup runs.
 *
 * @example
 * chart.saveChartContent(canvas.id, content);
 * chart.loadDescendantCharts(container);
 * chart.destroyDescendantCharts(container);
 * @author benizzio
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author GPT-6.1 Terra
 */
const chart = {
    saveChartContent,
    getChartContent,
    getChartContentFromChart,
    destroyDescendantCharts,
    /**
     * Loads registered charts in matching descendant canvases that have no chart yet.
     * The element itself is not included, even if it is a matching canvas.
     * @param element - Root whose descendant canvases should be loaded.
     * @example
     * chart.loadDescendantCharts(container);
     * @author benizzio
     * @author GPT-6 Luna
     * @author GPT-6.1 Terra
     */
    loadDescendantCharts(element: HTMLElement) {
        element.querySelectorAll(`canvas[${ CHART_ATTRIBUTE }]`).forEach((canvas: HTMLCanvasElement) => {
            loadChart(canvas);
        });
    },
    getPieDoughnutChartColorScale,
    convertUnidimensionalDatasetBackgroundToPattern,
};

export default chart;
