/**
 * Captures observation allocation rows as independent editable drafts and restores them
 * into freshly rendered ID-0 rows without triggering calculations or network requests.
 * @module components/portfolio-history/portfolio-history-management-observation-clone/draft
 * @author GPT-6.1 Terra
 */
import { captureFinancialInput, restoreFinancialInput, FINANCIAL_INPUT_ATTRIBUTE } from "../../../infra/dom";
import type { FinancialInputSnapshot } from "../../../infra/dom";
import { AssetSelectionState, captureAssetDraft, restoreAssetDraft } from "../../asset-composed-columns-input";
import type { AssetDraft } from "../../asset-composed-columns-input";
import { captureObservationTag } from "../portfolio-history-management-observation-metadata";
import {
    OBSERVATION_ALLOCATION_FIELDS as fields,
    observationAllocationFieldSelector,
} from "../portfolio-history-management-form";
import { EXTERNAL_ASSET_KEYS_SELECTOR } from "../portfolio-history-quote";

/** Independent allocation fields captured from one editable observation row.
 * @author GPT-6.1 Terra
 */
type AllocationDraft = {
    asset: AssetDraft;
    className: string;
    cashReserve: boolean;
    quantity: string;
    price: FinancialInputSnapshot;
    total: FinancialInputSnapshot;
    keys?: { source: string; exchangeId: string; ticker: string };
};

/** Complete editable observation snapshot.
 * @author GPT-6.1 Terra
 */
export type ObservationDraft = { tag: string; date: string; rows: AllocationDraft[] };

/** Resolves a required row input by its field suffix or financial binding name.
 * @author GPT-6.1 Terra
 */
function field(row: HTMLElement, suffix: string, financial = false): HTMLInputElement {
    const attr = financial ? FINANCIAL_INPUT_ATTRIBUTE : "name";
    const element = row.querySelector<HTMLInputElement>(`input${ observationAllocationFieldSelector(suffix, attr) }`);

    if(!element) {
        throw new Error(`Missing observation field ${ suffix }.`);
    }
    return element;
}

/** Captures a complete enabled provider key triplet for an existing selected asset.
 * @author GPT-6.1 Terra
 */
function captureKeys(row: HTMLElement, asset: AssetDraft): AllocationDraft["keys"] {
    const group = row.querySelector<HTMLFieldSetElement>(EXTERNAL_ASSET_KEYS_SELECTOR);

    if(asset.state !== AssetSelectionState.EXISTING || !group || group.disabled) {
        return undefined;
    }
    const value = (key: string) => group.querySelector<HTMLInputElement>(`[data-external-asset-key='${ key }']`)?.value;
    const source = value("source");
    const exchangeId = value("exchangeId");
    const ticker = value("ticker");
    return source && exchangeId && ticker ? { source, exchangeId, ticker } : undefined;
}

/** Captures one source row without mutating its fields.
 * @author GPT-6.1 Terra
 */
function captureRow(row: HTMLTableRowElement): AllocationDraft {
    const asset = captureAssetDraft(row);
    return {
        asset,
        className: field(row, fields.className).value,
        cashReserve: field(row, fields.cashReserve).checked,
        quantity: field(row, fields.quantity).value,
        price: captureFinancialInput(field(row, fields.marketPrice, true)),
        total: captureFinancialInput(field(row, fields.total, true)),
        keys: captureKeys(row, asset),
    };
}

/** Captures the current source tag and rows after detail initialization.
 * @param form - The loaded persisted observation form.
 * @param date - Browser-local date recorded at clone activation.
 * @returns A detached observation draft.
 * @example `const draft = captureObservationDraft(sourceForm, localToday);`
 * @author GPT-6.1 Terra
 */
export function captureObservationDraft(form: HTMLFormElement, date: string): ObservationDraft {
    return {
        tag: `${ captureObservationTag(form) } (CLONE)`,
        date,
        rows: Array.from(form.querySelectorAll<HTMLTableRowElement>("tbody > tr"), captureRow),
    };
}

/** Restores one row after standard HTMX and financial bindings have settled.
 * @author GPT-6.1 Terra
 */
function restoreRow(row: HTMLTableRowElement, draft: AllocationDraft): void {
    restoreAssetDraft(row, draft.asset);
    field(row, fields.className).value = draft.className;
    field(row, fields.cashReserve).checked = draft.cashReserve;
    field(row, fields.quantity).value = draft.quantity;
    restoreFinancialInput(field(row, fields.marketPrice, true), draft.price);
    restoreFinancialInput(field(row, fields.total, true), draft.total);
    const group = row.querySelector<HTMLFieldSetElement>(EXTERNAL_ASSET_KEYS_SELECTOR);

    if(!group) {
        throw new Error("Missing external asset keys.");
    }
    group.disabled = !draft.keys;

    if(draft.keys) {
        for(const key of ["source", "exchangeId", "ticker"] as const) {
            const target = group.querySelector<HTMLInputElement>(`[data-external-asset-key='${ key }']`);

            if(!target) {
                throw new Error("Missing external asset key.");
            }
            target.value = draft.keys[key];
        }
    }
}

/** Restores allocation values into fresh destination rows in their original order.
 * @param tbody - Settled destination table body.
 * @param draft - Independent source snapshot.
 * @example `restoreObservationRows(tbody, draft);`
 * @author GPT-6.1 Terra
 */
export function restoreObservationRows(tbody: HTMLElement, draft: ObservationDraft): void {
    const rows = Array.from(tbody.querySelectorAll<HTMLTableRowElement>(":scope > tr"));

    if(rows.length !== draft.rows.length) {
        throw new Error("Rendered observation row count does not match the draft.");
    }
    rows.forEach((row, index) => restoreRow(row, draft.rows[index]));
}
