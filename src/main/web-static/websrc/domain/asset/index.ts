/**
 * Exposes the shared asset shape used across allocation and asset-management flows.
 *
 * Co-authored by: OpenCode and Igor Benicio de Mesquita
 */

/**
 * Describes asset fields available before and after persistence.
 * ID and name are optional because allocation inputs may supply only a ticker. External data
 * stores an ordered list of provider identifiers; the first entry is preferred by portfolio views.
 *
 * Example: `const candidate: Asset = { ticker: "EXAMPLE" };`
 * Example: `const loaded: Asset = { id: 1, ticker: "EXAMPLE", name: "Example", externalData: null };`
 *
 * Co-authored by: OpenCode and Igor Benicio de Mesquita
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
 * Example: `const external: ExternalAsset = { source: "YAHOO_FINANCE", ticker: "IAU", exchangeId: "PCX" };`
 *
 * Authored by: OpenCode
 */
export type ExternalAsset = {
    source: string;
    ticker: string;
    exchangeId: string;
    name?: string;
    exchangeName?: string;
};
