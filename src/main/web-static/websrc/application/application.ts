import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
//TODO restrict folder imports to this module only
const application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        portfolioHistoryManagement.init();
        window["portfolioHistoryManagement"] = portfolioHistoryManagement;
    },
    chartContents: { toChartContent },
};

export default application;