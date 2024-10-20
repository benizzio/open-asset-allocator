export type PortfolioSlice = {
    assetName: string;
    assetTicker: string;
    class: string;
    cashReserve: boolean;
    totalMarketValue: number;
};

export type PortfolioAtTime = {
    timeFrameTag: string;
    slices: PortfolioSlice[];
};
