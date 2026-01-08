import BigNumber from "bignumber.js";

export type ObservationTimestamp = {
    id: number;
    timeTag: string;
    timestamp?: Date;
};

export type PortfolioAllocation = {
    assetName: string;
    assetTicker: string;
    class: string;
    cashReserve: boolean;
    totalMarketValue: BigNumber;
};

export type PortfolioSnapshot = {
    observationTimestamp: ObservationTimestamp;
    allocations: PortfolioAllocation[];
};

export type PortfolioAllocationDTO = Omit<PortfolioAllocation, "totalMarketValue"> & { totalMarketValue: string, };

export type PortfolioSnapshotDTO = Omit<PortfolioSnapshot, "allocations">
    & { allocations: PortfolioAllocationDTO[]; };