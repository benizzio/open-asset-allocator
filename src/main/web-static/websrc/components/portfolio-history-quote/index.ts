/**
 * Public API for row-scoped latest-closing-quote actions in portfolio history.
 *
 * @example
 * ```ts
 * const quoteAction = createPortfolioHistoryQuoteAction(recalculateAllocation);
 * quoteAction.bindRow(row);
 * quoteAction.updateButtonVisibility(row);
 * ```
 *
 * @author GPT-6 Luna
 */

export { createPortfolioHistoryQuoteAction } from "./quote-action";
export type {
    PortfolioHistoryQuoteAction,
    RecalculatePortfolioAllocation,
} from "./quote-action";
