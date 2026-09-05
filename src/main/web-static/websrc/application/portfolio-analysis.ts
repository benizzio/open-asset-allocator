import * as handlebars from "handlebars";
import { PotentialDivergence } from "../domain/portfolio-analysis";
import { BigNumber } from "bignumber.js";
import Format from "../infra/format";

/**
 * Registers the Handlebars helpers used to render portfolio divergence analysis values and bars.
 *
 * Example: call `registerPortfolioAnalysisHandlebarsHelpers()` during frontend infrastructure
 * initialization before compiling `template-divergence-analysis`.
 *
 * Co-authored by: OpenCode and Benizzio
 */
export function registerPortfolioAnalysisHandlebarsHelpers() {

    handlebars.registerHelper(
        "divergenceVisualization",
        function(totalMarketValue: number, potentialDivergence: PotentialDivergence, field: string) {

            const localMarketValue = potentialDivergence.totalMarketValue;
            const localDivergence = potentialDivergence.totalMarketValueDivergence;
            const plannedMarketValue = localMarketValue - localDivergence;

            switch(field) {
                case "totalMarketValue": {
                    return getValueLabel(localMarketValue, totalMarketValue);
                }

                case "plannedMarketValue": {
                    return getValueLabel(plannedMarketValue, totalMarketValue);
                }

                case "divergence": {
                    return getValueLabel(localDivergence, totalMarketValue);
                }

                case "divergenceBar": {

                    // A zero-valued parent has no meaningful percentage scale, so its zero bar stays empty.
                    const divergenceOnTotal = totalMarketValue === 0
                        ? 0
                        : new BigNumber(localDivergence).div(totalMarketValue).times(200).toNumber();
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

function getValueLabel(value: number, total: number) {

    const formattedValue = Format.formatCurrency(value);

    const formattedPercent = total === 0
        ? ""
        : " (" + Format.formatPercent(new BigNumber(value).div(total).toNumber()) + ")";

    return `${ formattedValue } ${ formattedPercent }`;
}
