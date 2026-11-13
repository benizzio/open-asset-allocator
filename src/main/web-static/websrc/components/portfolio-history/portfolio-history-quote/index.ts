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
 * @author GPT-6 Sol
 * @author GPT-6.1 Sol
 */

export {
    createPortfolioHistoryQuoteAction,
    readExternalAssetKeys,
    EXTERNAL_ASSET_KEYS_SELECTOR,
    EXTERNAL_ASSET_SOURCE_SELECTOR,
    EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR,
    EXTERNAL_ASSET_TICKER_SELECTOR,
} from "./quote-action";
export type {
    PortfolioHistoryQuoteAction,
    RecalculatePortfolioAllocation,
} from "./quote-action";
