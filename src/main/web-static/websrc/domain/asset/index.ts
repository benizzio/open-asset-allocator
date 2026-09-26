/**
 * Exposes the shared asset shape used across allocation and asset-management flows.
 *
 * @author OpenCode
 */

/**
 * Describes asset fields available before and after persistence.
 * ID and name are optional because allocation inputs may supply only a ticker. External data
 * stores an ordered list of provider identifiers; the first entry is preferred by portfolio views.
 *
 * @example
 * ```ts
 * const candidate: Asset = { ticker: "EXAMPLE" };
 * const loaded: Asset = { id: 1, ticker: "EXAMPLE", name: "Example", externalData: null };
 * ```
 *
 * @author Igor Benicio de Mesquita
 * @author OpenCode
 */
export type Asset = {
    id?: number;
    name?: string;
    ticker: string;
    externalData?: { data: ExternalAsset[]; } | null;
};

/**
 * Identifies a provider asset in the ordered external data list. Names are transient search
 * metadata; the backend persists only source, ticker and exchangeId.
 *
 * @example
 * ```ts
 * const external: ExternalAsset = { source: "YAHOO_FINANCE", ticker: "IAU", exchangeId: "PCX" };
 * ```
 *
 * @author OpenCode
 */
export type ExternalAsset = {
    source: string;
    ticker: string;
    exchangeId: string;
    name?: string;
    exchangeName?: string;
};
