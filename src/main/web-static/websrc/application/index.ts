import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";
import allocationPlanManagement from "../components/allocation-plan-management";
import { AssetPage, PortfolioPage } from "../pages";
import notifications from "../components/notifications";
import AssetComposedColumnsInput from "../components/asset-composed-columns-input";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 * Call `Application.init()` once after loading the application's frontend modules to register the
 * global bindings used by templates and page controllers.
 *
 * @author Igor Benicio de Mesquita
 * @author OpenCode
 */
const Application = {

    /**
     * Registers the application services and page controllers on `globalThis` for HTML bindings.
     *
     * @example
     * Application.init();
     *
     * @author OpenCode
     */
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        globalThis["portfolioHistoryManagement"] = portfolioHistoryManagement;
        globalThis["allocationPlanManagement"] = allocationPlanManagement;
        globalThis["portfolioPage"] = PortfolioPage;
        globalThis["assetPage"] = AssetPage;
        globalThis["notifications"] = notifications;
        globalThis["AssetComposedColumnsInput"] = AssetComposedColumnsInput;
    },

    chartContents: { toChartContent },
};

export default Application;
