import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";
import portfolioHistoryManagement from "../components/portfolio-history-management";
import allocationPlanManagement from "../components/allocation-plan-management";
import { AssetPage, PortfolioPage } from "../pages";
import notifications from "../components/notifications";
import AssetComposedColumnsInput from "../components/asset-composed-columns-input";
import paginatedList from "../components/paginated-list";
import { registerPortfolioHistoryPageTransform } from "./portfolio-history-pagination";
import { registerAllocationPlanPageTransform } from "./allocation-plan-pagination";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 * Call `Application.init()` once after loading the application's frontend modules to register the
 * global bindings used by templates and page controllers.
 *
 * @author Igor Benicio de Mesquita
 * @author OpenCode
 * @author GPT-6 Luna
 * @author GPT-6.1 Terra
 */
const Application = {

    /**
     * Registers the application services and page controllers on `globalThis` for HTML bindings.
     *
     * @example
     * Application.init();
     *
     * @author OpenCode
     * @author GPT-6.1 Terra
     */
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
        paginatedList.init();
        registerPortfolioHistoryPageTransform();
        registerAllocationPlanPageTransform();
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
