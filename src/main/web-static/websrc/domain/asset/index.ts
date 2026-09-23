/**
 * Exposes the shared asset shape used across allocation and asset-management flows.
 *
 * Co-authored by: OpenCode and Igor Benicio de Mesquita
 */

/**
 * Describes asset fields available before and after persistence.
 * ID and name are optional because allocation inputs may supply only a ticker. When updating a
 * loaded asset, keep its externalData unchanged until the frontend supports editing that data.
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
    externalData?: Record<string, unknown> | null;
};
