import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";
import allocationPlanManagement from "../components/allocation-plan-management";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
const Application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        portfolioHistoryManagement.init();
        allocationPlanManagement.init();
        window["portfolioHistoryManagement"] = portfolioHistoryManagement;
    },
    chartContents: { toChartContent },
};

export default Application;