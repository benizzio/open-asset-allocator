import { AllocationStructure, AllocationStructureDTO } from "./allocation";

export type Portfolio = {
    id: number;
    name: string;
    allocationStructure: AllocationStructure;
};

export type PortfolioAllocation = {
    assetName: string;
    assetTicker: string;
    class: string;
    cashReserve: boolean;
    totalMarketValue: number;
};

export type PortfolioAtTime = {
    timeFrameTag: string;
    allocations: PortfolioAllocation[];
};

export type PortfolioDTO = Omit<Portfolio, "allocationStructure"> & { allocationStructure: AllocationStructureDTO; };
