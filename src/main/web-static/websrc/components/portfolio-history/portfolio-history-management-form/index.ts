/**
 * Owns identifiers shared by portfolio-history observation forms and their row controllers.
 * This dependency-free boundary avoids importing the management controller from its child modules.
 * @module components/portfolio-history/portfolio-history-management-form
 * @author GPT-6.1 Sol
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
