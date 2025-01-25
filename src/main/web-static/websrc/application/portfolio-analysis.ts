import * as handlebars from "handlebars";
import { PotentialDivergence } from "../domain/portfolio-analysis";
import BigNumber from "bignumber.js";
import format from "../infra/format";

export function registerPortfolioAnalysisHandlebarsHelpers() {

    handlebars.registerHelper(
        "divergenceVisualization",
        function(totalMarketValue: number, potentialDivergence: PotentialDivergence, field: string) {

            const localMarketValue = potentialDivergence.totalMarketValue;
            const localDivergence = potentialDivergence.totalMarketValueDivergence;
            const plannedMarketValue = localMarketValue - localDivergence;

            switch(field) {
                case "totalMarketValue": {
                    return getDivergenceValueLabel(localMarketValue, totalMarketValue);
                }

                case "plannedMarketValue": {

                    return getDivergenceValueLabel(plannedMarketValue, totalMarketValue);
                }

                case "divergence": {
                    return getDivergenceValueLabel(localDivergence, totalMarketValue);
                }

                case "divergenceBar": {

                    const divergenceOnTotal =
                        new BigNumber(localDivergence).div(totalMarketValue).times(200).toNumber();
                    const barStyle = divergenceOnTotal > 0 ? "bg-danger" : "bg-success";

                    return `<div class="progress"
                             role="progressbar"
                        >
                            <div 
                                class="progress-bar progress-bar-striped ${ barStyle }" 
                                style="width: ${ Math.abs(divergenceOnTotal) }%">
                            </div>
                        </div>
                    `;
                }
            }
        },
    );
}

function getDivergenceValueLabel(value: number, total: number) {
    return format.formatCurrency(value) + " (" + format.formatPercent(new BigNumber(value).div(total).toNumber()) + ")";
}
