import BigNumber from "bignumber.js";
import { PortfolioSnapshot, PortfolioSnapshotDTO } from "../../portfolio-allocation";

export function mapToPortfolioSnapshot(
    dto: PortfolioSnapshotDTO,
): PortfolioSnapshot {

    return {

        ...dto,
        allocations: dto.allocations.map((allocationDTO) => {

            let totalMarketValue: BigNumber;

            try {
                totalMarketValue = new BigNumber(allocationDTO.totalMarketValue);
            } catch(error) {
                throw new Error(
                    `Invalid totalMarketValue value: ${ allocationDTO.totalMarketValue }`,
                    { cause: error },
                );
            }

            return {
                ...allocationDTO,
                totalMarketValue: totalMarketValue,
            };
        }),
    };
}