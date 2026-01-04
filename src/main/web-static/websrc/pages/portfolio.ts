import { Portfolio, PortfolioDTO } from "../domain/portfolio";
import { DomainService } from "../domain/service";

const PortfolioPage = {

    getContextPortfolio(): Portfolio {
        const portfolioElement = window["portfolio"] as HTMLScriptElement;
        const portfolioJSON = portfolioElement.textContent;
        const portfolioDTO = JSON.parse(portfolioJSON) as PortfolioDTO;
        return DomainService.mapping.mapToPortfolio(portfolioDTO);
    },
};

export default PortfolioPage;