/**
 * Validates fresh allocation markup for an independently cloned observation draft.
 * @module components/portfolio-history/portfolio-history-management-observation-clone/destination
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */
import type { TemplateDelegate } from "handlebars";
import type { ObservationDraft } from "./draft";
import {
    OBSERVATION_ALLOCATION_FIELDS as fields,
    OBSERVATION_ASSET_QUANTITY_SELECTOR,
    observationAllocationFieldSelector as fieldSelector,
} from "../portfolio-history-management-form";
import {
    ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE,
    ASSET_SEARCH_INPUT_ATTRIBUTE,
    ASSET_TICKER_INPUT_ATTRIBUTE,
    ASSET_ACTION_BUTTON_SELECTOR,
    NEW_ASSET_TICKER_MESSAGE_SELECTOR,
    ASSET_NAME_INPUT_SELECTOR,
} from "../../asset-composed-columns-input";
import { FINANCIAL_INPUT_ATTRIBUTE } from "../../../infra/dom";
import { EXTERNAL_ASSET_KEYS_SELECTOR } from "../portfolio-history-quote";

/**
 * Builds and validates destination-specific row markup before the existing draft is touched.
 * @param template - Initialized allocation row template.
 * @param draft - Detached current source snapshot.
 * @returns Complete ID-0 row markup.
 * @example `const html = renderRows(template, draft);`
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
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
            `[${ ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE }]`, `[${ ASSET_SEARCH_INPUT_ATTRIBUTE }]`,
            `[${ ASSET_TICKER_INPUT_ATTRIBUTE }]`, ASSET_ACTION_BUTTON_SELECTOR, NEW_ASSET_TICKER_MESSAGE_SELECTOR,
            EXTERNAL_ASSET_KEYS_SELECTOR, ASSET_NAME_INPUT_SELECTOR, fieldSelector(fields.assetId, "name"),
            fieldSelector(fields.className, "name"), fieldSelector(fields.cashReserve, "name"),
            OBSERVATION_ASSET_QUANTITY_SELECTOR, fieldSelector(fields.marketPrice, FINANCIAL_INPUT_ATTRIBUTE),
            fieldSelector(fields.total, FINANCIAL_INPUT_ATTRIBUTE),
        ];

        if(requiredSelectors.some(selector => !row.querySelector(selector))) {
            throw new Error("Rendered observation draft row is incomplete.");
        }
    }
    return html;
}
