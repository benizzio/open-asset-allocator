/**
 * Owns the allocation-plan row form contract shared by row management and cloning.
 * Keeps serialized field paths independent of controller initialization.
 * @module components/allocation-plan/allocation-plan-management-form
 * @author GPT-6.1 Sol
 */
export const FORM_LAST_ROW_INDEX_INPUT_NAME = "last-planned-allocation-row-index";
export const FORM_FIELD_DEPENDENT_ATTRIBUTE = "data-bind-to-name";
export const FORM_FIELD_LABEL_ATTRIBUTE = "data-label-for-name";
export const PLAN_DETAILS_FIELD_NAME = "details";
export const PLAN_HIERARCHICAL_ID_FIELD_NAME = "hierarchicalId";
export const PLAN_CASH_RESERVE_FIELD_NAME = "cashReserve";

/** Builds the suffix identifying one hierarchy field within an allocation row.
 * @param level - Zero-based hierarchy level.
 * @returns Serialized hierarchy field suffix.
 * @example `hierarchyFieldSuffix(1)` returns `[hierarchicalId][1]`.
 * @author GPT-6.1 Sol
 */
export function hierarchyFieldSuffix(level: number): string {
    return `[${ PLAN_HIERARCHICAL_ID_FIELD_NAME }][${ level }]`;
}

/** Builds the canonical name of a row's hierarchy field.
 * @param rowIndex - Allocation row index in the form payload.
 * @param level - Zero-based hierarchy level.
 * @returns Complete serialized form field name.
 * @example `hierarchyFieldName(2, 1)` returns `details[2][hierarchicalId][1]`.
 * @author GPT-6.1 Sol
 */
export function hierarchyFieldName(rowIndex: number | string, level: number): string {
    return `${ PLAN_DETAILS_FIELD_NAME }[${ rowIndex }]${ hierarchyFieldSuffix(level) }`;
}
