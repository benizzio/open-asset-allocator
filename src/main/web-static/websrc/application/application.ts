import { registerPortfolioAnalysisHandlebarsHelpers } from "./portfolio-analysis";

const application = {
    init() {
        registerPortfolioAnalysisHandlebarsHelpers();
    },
};

export default application;