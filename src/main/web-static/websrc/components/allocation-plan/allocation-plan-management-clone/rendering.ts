/**
 * Adapts structural plan snapshots to the existing allocation-row template with
 * independent hierarchy contexts, new row indexes, and ancestor field bindings.
 * @module components/allocation-plan/allocation-plan-management-clone/rendering
 * @author GPT-6.1 Terra
 */
import type { TemplateDelegate } from "handlebars";
import type { AllocationHierarchyLevel } from "../../../domain/allocation";
import type { PlanDraft, PlanRowDraft } from "./draft";

/** Builds fresh row markup and verifies its structural prerequisites.
 * @param template - Compiled existing allocation-row template.
 * @param hierarchy - Current portfolio hierarchy, ordered from asset level upward.
 * @param draft - Detached source snapshot.
 * @returns Markup for the destination body.
 * @example `const html = renderPlanRows(rowTemplate, hierarchy, draft);`
 * @author GPT-6.1 Terra
 */
export function renderPlanRows(
    template: TemplateDelegate, hierarchy: AllocationHierarchyLevel[], draft: PlanDraft,
): string {
    const indices = new Map<string, number>();
    const rows = new Map<string, PlanRowDraft>();

    const markup = draft.rows.map((row, index) => {
        if(row.level >= hierarchy.length || (!row.parent && row.level !== hierarchy.length - 1)) {
            throw new Error("Plan row is outside the portfolio hierarchy.");
        }
        const levels = hierarchy.map(level => ({ ...level, currentManagingFieldName: "" }));
        let parent = row.parent;

        while (parent) {
            const ancestor = rows.get(parent);
            const ancestorIndex = indices.get(parent);

            if(!ancestor || ancestorIndex === undefined) {
                throw new Error("Plan row parent is missing.");
            }

            levels[ancestor.level].currentManagingFieldName =
                `details[${ ancestorIndex }][hierarchicalId][${ ancestor.level }]`;
            parent = ancestor.parent;
        }

        const level = levels[row.level];

        const html = template({
            allocationPlanId: 0,
            allocationIndex: index,
            parentRowIndex: row.parent ? indices.get(row.parent) : undefined,
            hierarchy: levels,
            targetLevelKey: "",
            fractalPlannedAllocation: {
                level,
                subLevel: row.level > 0 ? levels[row.level - 1] : undefined,
                allocation: { hierarchicalId: [], cashReserve: false, sliceSizePercentage: 0, asset: { ticker: "" } },
            },
        });
        indices.set(row.id, index);
        rows.set(row.id, row);
        return html;
    }).join("");
    const fragment = document.createElement("template");
    fragment.innerHTML = `<table><tbody>${ markup }</tbody></table>`;
    const rendered = fragment.content.querySelectorAll<HTMLTableRowElement>("tbody > tr");

    if(rendered.length !== draft.rows.length) {
        throw new Error("Could not render the complete plan draft.");
    }

    rendered.forEach((row, index) => {
        const required = ["[data-percentage-input]", "input[name$='[cashReserve]']"];

        if(draft.rows[index].asset) {
            required.push("[data-asset-search-autocomplete]", "[data-asset-search-input]");
        }

        for(const value of draft.rows[index].hierarchy) {
            required.push(`input[name$='[hierarchicalId][${ value.level }]']`);
        }

        if(required.some(selector => !row.querySelector(selector))) {
            throw new Error("Rendered plan draft row is incomplete.");
        }
    });
    return markup;
}
