/**
 * Validates fresh allocation markup for an independently cloned observation draft.
 * @module components/portfolio-history/portfolio-history-management-observation-clone/destination
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */
import type { TemplateDelegate } from "handlebars";
import type { ObservationDraft } from "./draft";

/**
 * Builds and validates destination-specific row markup before the existing draft is touched.
 * @param template - Initialized allocation row template.
 * @param draft - Detached current source snapshot.
 * @returns Complete ID-0 row markup.
 * @example `const html = renderRows(template, draft);`
 * @author GPT-6.1 Sol
 */
export function renderRows(template: TemplateDelegate, draft: ObservationDraft): string {
    const html = draft.rows.map((_, allocationIndex) =>
        template({ allocationIndex, observationTimestampId: 0 })).join("");
    const documentFragment = document.createElement("template");
    documentFragment.innerHTML = `<table><tbody>${ html }</tbody></table>`;
    const rows = documentFragment.content.querySelectorAll("tbody > tr");

    if(rows.length !== draft.rows.length) {
        throw new Error("Could not render observation draft rows.");
    }

    for(const row of rows) {
        const requiredSelectors = [
            "[data-asset-search-autocomplete]", "[data-asset-search-input]", "[data-asset-ticker-input]",
            "[data-asset-action-button]", "[data-new-asset-ticker-message]", "[data-external-asset-keys]",
            "[aria-label='Asset name']", "[name$='[assetId]']", "[name$='[class]']", "[name$='[cashReserve]']",
            "[name$='[assetQuantity]']", "[data-financial-input$='[assetMarketPrice]']",
            "[data-financial-input$='[totalMarketValue]']",
        ];

        if(requiredSelectors.some(selector => !row.querySelector(selector))) {
            throw new Error("Rendered observation draft row is incomplete.");
        }
    }
    return html;
}
