/**
 * Removes allocation-plan hierarchy rows together with their complete descendants.
 * Relationships are resolved from row identities within the owning form, independently
 * of editable hierarchy values.
 * @module components/allocation-plan/allocation-plan-management-hierarchy
 * @author GPT-6.1 Terra
 */

/** Removes the selected row and every descendant belonging to the same form.
 *
 * @param target - Control inside the row to remove.
 * @example `removePlannedAllocationSubtree(removeButton);`
 * @author GPT-6.1 Terra
 */
export function removePlannedAllocationSubtree(target: HTMLElement): void {
    const root = target.closest<HTMLTableRowElement>("tr");
    const form = root?.closest<HTMLFormElement>("form");

    if(!root || !form || !root.id) {
        return;
    }

    const rows = Array.from(form.querySelectorAll<HTMLTableRowElement>("tbody > tr"));
    const removedIds = new Set([root.id]);
    let previousSize: number;

    do {
        previousSize = removedIds.size;

        for(const row of rows) {
            if(row.id && row.dataset.parentRowId && removedIds.has(row.dataset.parentRowId)) {
                removedIds.add(row.id);
            }
        }
    } while (removedIds.size !== previousSize);

    for(const row of rows) {
        if(removedIds.has(row.id)) {
            row.remove();
        }
    }
}
