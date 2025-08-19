import { AllocationStructure, AllocationStructureDTO } from "./allocation";

export type Portfolio = {
    id: number;
    name: string;
    allocationStructure: AllocationStructure;
};

export type PortfolioDTO = Omit<Portfolio, "allocationStructure"> & { allocationStructure: AllocationStructureDTO; };
