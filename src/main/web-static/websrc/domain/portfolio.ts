import { AllocationStructure } from "./allocation";

export type PortfolioAllocation = {
    assetName: string;
    assetTicker: string;
    class: string;
    cashReserve: boolean;
    totalMarketValue: number;
};

export type PortfolioAtTime = {
    timeFrameTag: string;
    //TODO change property name to allocations with back-end
    slices: PortfolioAllocation[];
    structure?: AllocationStructure
};
