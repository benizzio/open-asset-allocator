import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";
import allocationPlanManagement from "../components/allocation-plan-management";
import { AssetsPage, PortfolioPage } from "../pages";
import notifications from "../components/notifications";
import AssetComposedColumnsInput from "../components/asset-composed-columns-input";

// Co-authored by: OpenCode

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
const Application = {

    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        globalThis["portfolioHistoryManagement"] = portfolioHistoryManagement;
        globalThis["allocationPlanManagement"] = allocationPlanManagement;
        globalThis["portfolioPage"] = PortfolioPage;
        globalThis["assetsPage"] = AssetsPage;
        globalThis["notifications"] = notifications;
        globalThis["AssetComposedColumnsInput"] = AssetComposedColumnsInput;
    },

    chartContents: { toChartContent },
};

export default Application;
