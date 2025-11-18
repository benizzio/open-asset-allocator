import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";
import allocationPlanManagement from "../components/allocation-plan-management";
import PortfolioPage from "../pages/portfolio";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
const Application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        portfolioHistoryManagement.init();
        allocationPlanManagement.init();
        window["portfolioHistoryManagement"] = portfolioHistoryManagement;
        window["allocationPlanManagement"] = allocationPlanManagement;
        window["portfolioPage"] = PortfolioPage;
    },
    chartContents: { toChartContent },
};

export default Application;