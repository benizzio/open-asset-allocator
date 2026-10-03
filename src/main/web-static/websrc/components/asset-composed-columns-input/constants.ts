/**
 * Shared DOM identities and messages used by the composed asset-column controllers.
 *
 * @author GPT-6 Luna
 */

import { BootstrapClasses, BootstrapIconClasses } from "../../infra/bootstrap/constants";

export const TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE = "data-asset-ticker-extra-error-message";
export const ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE = "data-asset-search-autocomplete";
export const ASSET_SEARCH_SUGGESTIONS_ATTRIBUTE = "data-asset-search-suggestions";
export const ASSET_SEARCH_LISTBOX_ATTRIBUTE = "data-asset-search-listbox";
export const ASSET_SEARCH_STATUS_ATTRIBUTE = "data-asset-search-status";
export const ASSET_SEARCH_INPUT_ATTRIBUTE = "data-asset-search-input";
export const ASSET_TICKER_INPUT_ATTRIBUTE = "data-asset-ticker-input";
export const ASSET_SEARCH_OPTION_SELECTOR = "[data-asset-search-option]";
export const ASSET_ACTION_BUTTON_SELECTOR = "[data-asset-action-button]";
export const NEW_ASSET_TICKER_MESSAGE_SELECTOR = "[data-new-asset-ticker-message]";
export const ASSET_NAME_INPUT_SELECTOR = "input[aria-label='Asset name']";
export const ASSET_ID_INPUT_SELECTOR = "input[type='hidden'][data-null-if-empty]";

export const ARIA_CONTROLS_ATTRIBUTE = "aria-controls";
export const ARIA_EXPANDED_ATTRIBUTE = "aria-expanded";
export const ARIA_ACTIVE_DESCENDANT_ATTRIBUTE = "aria-activedescendant";
export const ARIA_SELECTED_ATTRIBUTE = "aria-selected";

export const INVALID_INPUT_CLASS = "is-invalid";
export const VISUALLY_HIDDEN_CLASS = "visually-hidden";
export const ASSET_SEARCH_OPTION_CLASS = "asset-search-autocomplete__option";

export const ASSET_SELECTION_ERROR_MESSAGE = "Reference an existing asset or create a new one";
export const ASSET_LOOKUP_PENDING_ERROR_MESSAGE = "Asset lookup is still in progress";
export const REQUIRED_FOR_SEARCH_ERROR_MESSAGE = "Required for search";
export const TICKER_TOO_LONG_ERROR_MESSAGE = "Asset ticker cannot exceed 40 characters";
export const INVALID_TICKER_ERROR_MESSAGE = "Ticker may contain only letters, numbers, '-', '_', ':', and '.'";
export const ASSET_NOT_FOUND_ERROR_MESSAGE = "Data not found";
export const FAILED_TO_FETCH_ASSET_ERROR_PREFIX = "Failed to fetch asset data: ";
export const UNABLE_TO_RESOLVE_ASSET_ROW_ERROR = "Unable to resolve the asset row for the selected ticker.";

export const TICKER_CANDIDATE_PATTERN = /^[A-Za-z0-9_:.-]+$/;

/**
 * Identifies the search, in-flight lookup, or committed asset selection stored on an autocomplete wrapper.
 *
 * The HTML template initializes `data-asset-selection-state` to `search`. Read or update subsequent
 * states through `element.dataset.assetSelectionState` and these members, for example:
 * `wrapper.dataset.assetSelectionState = AssetSelectionState.PENDING`.
 *
 * @author GPT-6 Sol
 */
export enum AssetSelectionState {
    SEARCH = "search",
    PENDING = "pending",
    EXISTING = "existing",
    NEW = "new",
}

export const ASSET_ACTION_BUTTON_IDENTITIES = {
    search: {
        classes: `${ BootstrapClasses.BUTTON_PRIMARY } btn-xs`,
        iconClasses: `${ BootstrapIconClasses.SEARCH }`,
    },
    reset: {
        classes: `${ BootstrapClasses.BUTTON_DANGER } btn-xs`,
        iconClasses: `${ BootstrapIconClasses.RESET }`,
    },
};
