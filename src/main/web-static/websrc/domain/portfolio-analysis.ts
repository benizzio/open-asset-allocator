import { ObservationTimestamp } from "./portfolio";

export type PotentialDivergence = {
    hierarchyLevelKey: string,
    hierarchicalId: string,
    totalMarketValue: number,
    totalMarketValueDivergence: number,
    depth: number,
    internalDivergences?: PotentialDivergence[],
};

export type DivergenceAnalysis = {
    portfolioId: number;
    observationTimestamp: ObservationTimestamp;
    allocationPlanId: number;
    portfolioTotalMarketValue: number;
    root: PotentialDivergence[];
};