/**
 * DOM selectors for the row-scoped portfolio-history quote controls.
 * @module components/portfolio-history/portfolio-history-quote/selectors
 * @author GPT-6.1 Terra
 */
import { FINANCIAL_INPUT_ATTRIBUTE, FINANCIAL_INPUT_DECIMALS_ATTRIBUTE } from "../../../infra/dom";
import { OBSERVATION_ALLOCATION_SELECTORS } from "../portfolio-history-management-form";

/**
 * Identifies the quote controls and submitted financial fields in an observation row.
 * @example `row.querySelector(quoteSelectors.action)` finds the quote button.
 * @author GPT-6.1 Terra
 */
export const quoteSelectors = {
    group: ".portfolio-history-market-price",
    action: "[data-quote-action]",
    icon: "[data-quote-icon]",
    spinner: "[data-quote-spinner]",
    priceDisplay: `[${ FINANCIAL_INPUT_ATTRIBUTE }][${ FINANCIAL_INPUT_DECIMALS_ATTRIBUTE }]`,
    priceRaw: `input[type="hidden"]${ OBSERVATION_ALLOCATION_SELECTORS.marketPrice }`,
    totalRaw: `input[type="hidden"]${ OBSERVATION_ALLOCATION_SELECTORS.total }`,
} as const;
