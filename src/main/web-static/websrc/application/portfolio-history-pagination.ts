/**
 * Adapts portfolio-history page responses to the history view template.
 * @module application/portfolio-history-pagination
 * @author GPT-6.1 Terra
 */
import paginatedList, { PaginationMode } from "../components/paginated-list";
import { HtmxInfra, type BeforeSwapEventDetail } from "../infra/htmx";

/**
 * Identifies the history snapshot within a page.
 * @author GPT-6.1 Terra
 */
type HistoryPageItem = { observationTimestamp?: { id?: number } };

/**
 * Prepares history snapshots for the view template.
 * @author GPT-6 Luna
 * @author GPT-6 Sol
 * @author GPT-6.1 Terra
 */
function transformPortfolioHistoryView(responseBody: string, eventDetail: BeforeSwapEventDetail): string {
    return paginatedList.transformPageResponse<HistoryPageItem, { items: HistoryPageItem[]; initial: boolean }>(
        responseBody,
        eventDetail,
        item => item.observationTimestamp?.id,
        (items, mode) => ({ items, initial: mode === PaginationMode.Initial }),
    );
}

/**
 * Registers the history view's named HTMX response adapter.
 * @example `registerPortfolioHistoryPageTransform()` registers `paginatePortfolioHistoryView`.
 * @author GPT-6.1 Terra
 */
export function registerPortfolioHistoryPageTransform(): void {
    HtmxInfra.htmxTransformResponse.registerTransformResponseFunction(
        "paginatePortfolioHistoryView", transformPortfolioHistoryView,
    );
}
