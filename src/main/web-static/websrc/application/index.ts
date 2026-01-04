import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";
import allocationPlanManagement from "../components/allocation-plan-management";
import PortfolioPage from "../pages/portfolio";
import notifications from "../components/notifications";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
const Application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        allocationPlanManagement.init();
        window["portfolioHistoryManagement"] = portfolioHistoryManagement;
        window["allocationPlanManagement"] = allocationPlanManagement;
        window["portfolioPage"] = PortfolioPage;
        window["notifications"] = notifications;
    },
    chartContents: { toChartContent },
};

export default Application;