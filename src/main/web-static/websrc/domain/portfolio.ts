import { AllocationStructure, AllocationStructureDTO } from "./allocation";

export type Portfolio = {
    id: number;
    name: string;
    allocationStructure: AllocationStructure;
};

export type ObservationTimestamp = {
    id: number;
    timeTag: string;
    timestamp: Date;
};

export type PortfolioAllocation = {
    assetName: string;
    assetTicker: string;
    class: string;
    cashReserve: boolean;
    totalMarketValue: number;
};

export type PortfolioSnapshot = {
    /**
     * @deprecated use `observationTimestamp` instead
     */
    timeFrameTag: string;
    observationTimestamp: ObservationTimestamp;
    allocations: PortfolioAllocation[];
};

export type PortfolioDTO = Omit<Portfolio, "allocationStructure"> & { allocationStructure: AllocationStructureDTO; };
