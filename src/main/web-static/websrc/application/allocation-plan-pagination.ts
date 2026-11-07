/**
 * Adapts allocation-plan page responses to the allocation-plan view template.
 * @module application/allocation-plan-pagination
 * @author GPT-6.1 Terra
 */
import paginatedList, { PaginationMode } from "../components/paginated-list";
import { HtmxInfra, type BeforeSwapEventDetail } from "../infra/htmx";

/**
 * Identifies an allocation plan within a page.
 * @author GPT-6.1 Terra
 */
type AllocationPlanPageItem = { id?: number };

/** Describes the template input for one accepted plan page.
 * @author GPT-6.1 Terra
 */
type AllocationPlanViewPage = { items: AllocationPlanPageItem[]; initial: boolean };

/**
 * Prepares plans for the view template.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author GPT-6.1 Terra
 */
function transformAllocationPlanView(responseBody: string, eventDetail: BeforeSwapEventDetail): string {
    return paginatedList.transformPageResponse<AllocationPlanPageItem, AllocationPlanViewPage>(
        responseBody,
        eventDetail,
        item => item.id,
        (items, mode) => ({ items, initial: mode === PaginationMode.Initial }),
    );
}

/**
 * Registers the allocation-plan view's named HTMX response adapter.
 * @example `registerAllocationPlanPageTransform()` registers `paginateAllocationPlanView`.
 * @author GPT-6.1 Terra
 */
export function registerAllocationPlanPageTransform(): void {
    HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
        "paginateAllocationPlanView", transformAllocationPlanView,
    );
}
