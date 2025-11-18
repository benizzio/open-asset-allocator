import { Portfolio } from "../domain/portfolio";

const PortfolioPage = {

    getContextPortfolio(): Portfolio {
        const portfolioElement = window["portfolio"] as HTMLScriptElement;
        const portfolioJSON = portfolioElement.textContent;
        return JSON.parse(portfolioJSON) as Portfolio;
    },
};

export default PortfolioPage;