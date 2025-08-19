import { Portfolio, PortfolioDTO } from "../portfolio";
import { mapAllocationStructure } from "./allocation-mapping";

export function mapToPortfolio(portfolioDTO: PortfolioDTO): Portfolio {
    return {
        id: portfolioDTO.id,
        name: portfolioDTO.name,
        allocationStructure: mapAllocationStructure(portfolioDTO.allocationStructure),
    };
}