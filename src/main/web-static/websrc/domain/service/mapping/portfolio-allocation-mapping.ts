import BigNumber from "bignumber.js";
import { PortfolioSnapshot, PortfolioSnapshotDTO } from "../../portfolio-allocation";

export function mapToPortfolioSnapshot(
    dto: PortfolioSnapshotDTO,
): PortfolioSnapshot {

    return {

        ...dto,
        allocations: dto.allocations.map((allocationDTO) => {

            const totalMarketValue = new BigNumber(allocationDTO.totalMarketValue);

            if(totalMarketValue.isNaN()) {
                throw new Error(`Invalid totalMarketValue value: ${ allocationDTO.totalMarketValue }`);
            }

            return {
                ...allocationDTO,
                totalMarketValue: totalMarketValue,
            };
        }),
    };
}