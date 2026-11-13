/**
 * Shares committed existing/new asset presentation between lookup transitions and draft restoration.
 * Rendering does not complete lookups, clear validation, move focus, or publish selection events.
 * @module components/asset-composed-columns-input/committed-selection-renderer
 * @author GPT-6.1 Sol
 */
import { ASSET_ACTION_BUTTON_IDENTITIES, AssetSelectionState } from "./constants";

/** Explicit committed values, including an incomplete new asset's captured name.
 * @author GPT-6.1 Sol
 */
export type CommittedAssetSelection = {
    state: AssetSelectionState.EXISTING | AssetSelectionState.NEW;
    id: string;
    ticker: string;
    name: string;
};

/** Resolved controls whose committed presentation is owned by this internal renderer.
 * @author GPT-6.1 Sol
 */
type CommittedAssetControls = {
    wrapper: HTMLElement;
    button: HTMLButtonElement;
    message: HTMLElement;
    search: HTMLInputElement;
    ticker: HTMLInputElement;
    name: HTMLInputElement;
    id: HTMLInputElement;
};

/** Renders a committed selection without invoking the owning controller's transition side effects.
 * The supplied name is preserved verbatim, even when a new asset is incomplete.
 * @param controls - Already resolved controls from the same composed asset row.
 * @param selection - Explicit state and asset values to display; new selections always clear the ID.
 * @example `renderCommittedAssetSelection(controls, { state: AssetSelectionState.NEW, id: "", ticker, name });`
 * @author GPT-6.1 Sol
 */
export function renderCommittedAssetSelection(
    controls: CommittedAssetControls,
    selection: CommittedAssetSelection,
): void {
    const existing = selection.state === AssetSelectionState.EXISTING;
    controls.wrapper.dataset.assetSelectionState = selection.state;
    controls.search.disabled = true;
    controls.search.hidden = true;
    controls.ticker.disabled = false;
    controls.ticker.hidden = false;
    controls.ticker.readOnly = true;
    controls.ticker.value = selection.ticker;
    controls.name.style.display = "";
    controls.name.value = selection.name;
    controls.name.readOnly = existing;
    controls.name.required = !existing;
    controls.id.value = existing ? selection.id : "";
    controls.message.style.display = existing ? "none" : "";
    controls.button.className = ASSET_ACTION_BUTTON_IDENTITIES.reset.classes;
    controls.button.innerHTML = `<span class="${ ASSET_ACTION_BUTTON_IDENTITIES.reset.iconClasses }"></span>`;
}
