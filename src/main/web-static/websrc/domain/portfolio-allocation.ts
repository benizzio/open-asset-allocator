import { BigNumber } from "bignumber.js";

/**
 * Observation timestamp returned by the portfolio history API and submitted unchanged for existing records.
 * @example `{ id: 12, timeTag: "Monthly", timestamp: "2026-10-18T00:00:00.000Z" }`
 * @author GPT-6 Luna
 * @author benizzio
 */
export type ObservationTimestamp = {
    id: number;
    timeTag: string;
    timestamp?: string;
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
