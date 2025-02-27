import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";
import { toChartContent } from "./chart-contents";

const application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
    },
    chartContents: { toChartContent },
};

export default application;