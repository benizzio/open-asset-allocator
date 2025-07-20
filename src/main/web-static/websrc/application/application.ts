import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
const application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        window["portfolioHistoryManagement"] = portfolioHistoryManagement;
    },
    chartContents: { toChartContent },
};

export default application;