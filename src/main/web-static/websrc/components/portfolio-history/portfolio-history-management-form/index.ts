/**
 * Owns identifiers shared by portfolio-history observation forms and their row controllers.
 * This dependency-free boundary avoids importing the management controller from its child modules.
 * @module components/portfolio-history/portfolio-history-management-form
 * @author GPT-6.1 Sol
 * @author GPT-6.1 Terra
 */

/**
 * Prefix for observation forms and allocation row IDs.
 * @example `const formId = PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX + observationId;`
 * @author GPT-6.1 Sol
 */
export const PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX = "portfolio-history-management-form-";

/**
 * Identifies the single unsaved observation destination.
 * @example `document.getElementById(NEW_OBSERVATION_FORM_ID)` finds the editable draft form.
 * @author GPT-6.1 Sol
 */
export const NEW_OBSERVATION_FORM_ID = `${ PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX }0`;

/**
 * Field suffix for an observation allocation's asset quantity.
 * @example `field(row, OBSERVATION_ASSET_QUANTITY_FIELD)` resolves the quantity in a draft row.
 * @author GPT-6.1 Terra
 */
export const OBSERVATION_ASSET_QUANTITY_FIELD = "assetQuantity";

/**
 * Selector for the named asset quantity control in an observation allocation row.
 * @example `row.querySelector(OBSERVATION_ASSET_QUANTITY_SELECTOR)` finds its quantity control.
 * @author GPT-6.1 Terra
 */
export const OBSERVATION_ASSET_QUANTITY_SELECTOR = `[name$='[${ OBSERVATION_ASSET_QUANTITY_FIELD }]']`;
