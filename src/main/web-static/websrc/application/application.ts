import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";

/**
 * Component that represents the central application code tied to the underlying infrastructure.
 */
const application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
    },
    chartContents: { toChartContent },
};

export default application;