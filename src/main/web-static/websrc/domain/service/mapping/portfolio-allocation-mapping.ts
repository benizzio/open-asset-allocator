import BigNumber from "bignumber.js";
import { PortfolioSnapshot, PortfolioSnapshotDTO } from "../../portfolio-allocation";

export function mapToPortfolioSnapshot(
    dto: PortfolioSnapshotDTO,
): PortfolioSnapshot {

    return {

        ...dto,
        allocations: dto.allocations.map((allocationDTO) => {

            try {
                const totalMarketValue = new BigNumber(allocationDTO.totalMarketValue);

                return {
                    ...allocationDTO,
                    totalMarketValue: totalMarketValue,
                };
            } catch(error) {
                throw new Error(
                    `Invalid totalMarketValue value: ${ allocationDTO.totalMarketValue }`,
                    { cause: error },
                );
            }
        }),
    };
}