import { BigNumber } from "bignumber.js";
import { PortfolioSnapshot, PortfolioSnapshotDTO } from "../../portfolio-allocation";

export function mapToPortfolioSnapshot(
    dto: PortfolioSnapshotDTO,
): PortfolioSnapshot {

    return {

        ...dto,
        allocations: dto.allocations.map((allocationDTO, index) => {

            try {
                const totalMarketValue = new BigNumber(allocationDTO.totalMarketValue);

                if(totalMarketValue.isNaN()) {
                    throw new Error("BigNumber resolved to NaN");
                }

                return {
                    ...allocationDTO,
                    totalMarketValue: totalMarketValue,
                };
            } catch(error) {
                throw new Error(
                    `Invalid totalMarketValue "${ allocationDTO.totalMarketValue }" for allocation` +
                    ` [${ index }] (ticker: ${ allocationDTO.assetTicker })`,
                    { cause: error },
                );
            }
        }),
    };
}