/**
 * Captures and restores allocation-plan rows as independent editable snapshots.
 * Structural identities do not depend on hierarchy values, which may be blank,
 * duplicated, or temporarily out of sync with their inherited labels.
 * @module components/allocation-plan/allocation-plan-management-clone/draft
 * @author GPT-6.1 Terra
 */
import { capturePercentageInput, restorePercentageInput, type PercentageInputSnapshot } from "../../../infra/dom";
import { captureAssetDraft, restoreAssetDraft, type AssetDraft } from "../../asset-composed-columns-input";
import { captureAllocationPlanName } from "../allocation-plan-management-metadata";

/** One hierarchy value and its separately displayed inherited label.
 * @author GPT-6.1 Terra
 */
type HierarchyValue = { level: number; value: string; label?: string };

/** An independent row snapshot with structural parent identity.
 * @author GPT-6.1 Terra
 */
export type PlanRowDraft = {
    id: string;
    parent?: string;
    level: number;
    hierarchy: HierarchyValue[];
    cash: boolean;
    percentage: PercentageInputSnapshot;
    asset?: AssetDraft;
};

/** Complete editable allocation-plan snapshot.
 * @author GPT-6.1 Terra
 */
export type PlanDraft = { name: string; date: string; rows: PlanRowDraft[] };

/** Finds a required row input.
 * @author GPT-6.1 Terra
 */
function input(row: HTMLElement, selector: string): HTMLInputElement {
    const element = row.querySelector<HTMLInputElement>(selector);

    if(!element) {
        throw new Error(`Missing plan draft input: ${ selector }`);
    }
    return element;
}

/** Captures one row and its current, potentially unsynchronized values.
 * @author GPT-6.1 Terra
 */
function captureRow(row: HTMLTableRowElement): PlanRowDraft {
    const level = Number(row.dataset.hierarchyLevel);

    if(!row.id || row.dataset.hierarchyLevel === undefined || !Number.isInteger(level) || level < 0) {
        throw new Error("Plan row has an invalid identity or hierarchy level.");
    }

    const hierarchy = Array.from(row.querySelectorAll<HTMLInputElement>("input[name*='[hierarchicalId]']"), field => {
        const match = field.name.match(/\[hierarchicalId\]\[(\d+)\]$/);

        if(!match) {
            throw new Error("Plan hierarchy field has an invalid name.");
        }
        return {
            level: Number(match[1]),
            value: field.value,
            label: field.parentElement?.querySelector<HTMLElement>("[data-label-for-name]")?.textContent ?? undefined,
        };
    });
    return {
        id: row.id,
        parent: row.dataset.parentRowId,
        level,
        hierarchy,
        cash: input(row, "input[name$='[cashReserve]']").checked,
        percentage: capturePercentageInput(input(row, "[data-percentage-input]")),
        asset: level === 0 ? captureAssetDraft(row) : undefined,
    };
}

/** Captures a source plan and validates its explicit parent relationships.
 * @param form - Persisted source form.
 * @param date - Browser-local date captured at activation.
 * @param hierarchySize - Number of levels in the current portfolio hierarchy.
 * @returns Independent ordered draft values.
 * @example `const draft = capturePlanDraft(source, today, 2);`
 * @author GPT-6.1 Terra
 */
export function capturePlanDraft(form: HTMLFormElement, date: string, hierarchySize: number): PlanDraft {
    const rows = Array.from(form.querySelectorAll<HTMLTableRowElement>("tbody > tr"), captureRow);
    const seen = new Map<string, PlanRowDraft>();

    for(const row of rows) {
        const parent = row.parent ? seen.get(row.parent) : undefined;

        if(seen.has(row.id) || (row.parent && parent?.level !== row.level + 1)) {
            throw new Error("Plan hierarchy has an invalid parent relationship.");
        }
        const levels = new Set(row.hierarchy.map(value => value.level));

        if(levels.size !== hierarchySize - row.level || row.hierarchy.length !== levels.size) {
            throw new Error("Plan row has incomplete hierarchy controls.");
        }

        for(let level = row.level; level < hierarchySize; level++) {
            if(!levels.has(level)) {
                throw new Error("Plan row has incomplete hierarchy controls.");
            }
        }
        seen.set(row.id, row);
    }
    return { name: `${ captureAllocationPlanName(form) } (CLONE)`, date, rows };
}

/** Restores exact values into freshly rendered and bound rows without input events.
 * @param tbody - Destination table body after bindings are initialized.
 * @param draft - Captured source draft.
 * @example `restorePlanRows(tbody, draft);`
 * @author GPT-6.1 Terra
 */
export function restorePlanRows(tbody: HTMLElement, draft: PlanDraft): void {
    const rows = Array.from(tbody.querySelectorAll<HTMLTableRowElement>(":scope > tr"));

    if(rows.length !== draft.rows.length) {
        throw new Error("Rendered plan row count does not match the draft.");
    }

    rows.forEach((row, index) => {
        const snapshot = draft.rows[index];
        input(row, "input[name$='[cashReserve]']").checked = snapshot.cash;
        restorePercentageInput(input(row, "[data-percentage-input]"), snapshot.percentage);

        if(snapshot.asset) {
            restoreAssetDraft(row, snapshot.asset);
        }

        for(const value of snapshot.hierarchy) {
            const field = input(row, `input[name$='[hierarchicalId][${ value.level }]']`);
            field.value = value.value;
            const label = field.parentElement?.querySelector<HTMLElement>("[data-label-for-name]");

            if(label && value.label !== undefined) {
                label.textContent = value.label;
            }
        }
    });
}
