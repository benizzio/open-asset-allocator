import BigNumber from "bignumber.js";
import { PortfolioSnapshot, PortfolioSnapshotDTO } from "../../portfolio-allocation";

export function mapToPortfolioSnapshot(
    dto: PortfolioSnapshotDTO,
): PortfolioSnapshot {
    return {
        ...dto,
        allocations: dto.allocations.map((allocationDTO) => ({
            ...allocationDTO,
            totalMarketValue: new BigNumber(allocationDTO.totalMarketValue),
        })),
    };
}