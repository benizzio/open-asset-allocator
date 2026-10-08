/**
 * Public event contract for committed asset-row selection transitions.
 *
 * @author GPT-6 Luna
 */

import type { Asset } from "../../domain/asset";

/**
 * Event name emitted by a composed asset row after an existing, new, or cleared selection transition.
 * Listen through the public `components/asset-composed-columns-input` API. The event bubbles from the row and
 * its detail generation matches the row's `data-asset-selection-generation` attribute.
 *
 * @example
 * ```ts
 * document.addEventListener(ASSET_ROW_SELECTION_CHANGE_EVENT, event => {
 *     const selection = (event as AssetRowSelectionChangeEvent).detail;
 *     console.log(selection.state, selection.generation);
 * });
 * ```
 *
 * @author GPT-6 Luna
 */
export const ASSET_ROW_SELECTION_CHANGE_EVENT = "asset-row-selection-change";

/**
 * Describes one asset-row transition. Existing selections include the resolved asset; new and search states
 * intentionally carry no asset because consumers should clear any association tied to the previous selection.
 * Generation increases on every committed existing/new selection and every reset to search mode.
 *
 * @example
 * ```ts
 * const detail: AssetRowSelectionChangeDetail = {
 *     state: "existing",
 *     asset: { id: 1, ticker: "EXAMPLE" },
 *     generation: 1,
 * };
 * ```
 *
 * @author GPT-6 Luna
 */
export type AssetRowSelectionChangeDetail =
    | { state: "existing"; asset: Asset; generation: number }
    | { state: "new" | "search"; generation: number };

/**
 * Typed custom event emitted by a composed asset row. It bubbles from the owning row after its DOM state changes.
 *
 * @example
 * ```ts
 * row.addEventListener(ASSET_ROW_SELECTION_CHANGE_EVENT, event => {
 *     const detail = (event as AssetRowSelectionChangeEvent).detail;
 *     if(detail.state === "existing") console.log(detail.asset.ticker);
 * });
 * ```
 *
 * @author GPT-6 Luna
 */
export type AssetRowSelectionChangeEvent = CustomEvent<AssetRowSelectionChangeDetail>;

/**
 * Payload accepted by the internal transition publisher before it adds the row generation.
 *
 * @author GPT-6 Luna
 */
export type AssetRowSelectionChangePayload =
    | { state: "existing"; asset: Asset }
    | { state: "new" | "search" };
