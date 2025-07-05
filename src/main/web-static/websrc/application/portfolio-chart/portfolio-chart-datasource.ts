import { MultiChartDataSource } from "../../infra/chart/chart-types";
import { PortfolioSnapshot } from "../../domain/portfolio";
import {
    AllocationHierarchyLevel,
    AllocationStructure,
    LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX,
} from "../../domain/allocation";
import { allocationDomainService } from "../../domain/allocation-service";
import { AppliedAllocationHierarchyLevel, MappedChartData } from "./portfolio-chart-model";
import { mapChartData } from "./portfolio-chart-mapping";

/**
 * MultiChartDataSource that keeps track of an applied hierarchy level,
 * based on the Portfolio allocation hierarchy structure.
 * It stores the current level being used and
 * the applied hierarchy levels for filtering data for previously selected levels only.
 */
export class FractalPortfolioMultiChartDataSource extends MultiChartDataSource {

    private readonly appliedHierarchyLevels: AppliedAllocationHierarchyLevel[] = [];
    private currentHierarchyLevelIndex: number;

    constructor(
        initialChartData: MappedChartData,
        private readonly portfolioAtTime: PortfolioSnapshot,
        private readonly portfolioAllocationStructure: AllocationStructure,
    ) {

        const {
            topHierarchyLevel,
            topLevelIndex,
        } = allocationDomainService.getTopHierarchyLevelInfoFromAllocationStructure(portfolioAllocationStructure);

        const initialDataKey = generateDataKey(topHierarchyLevel.field);
        const chartDataMap = new Map([[initialDataKey, initialChartData]]);
        super(chartDataMap, initialDataKey);

        this.currentHierarchyLevelIndex = topLevelIndex;
    }

    toNextLevel(filteringValue: string | number): string {

        if(this.currentHierarchyLevelIndex === LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX) {
            return;
        }

        this.movePropertiesToNextLevel(filteringValue);

        return this.mapAndReferenceChartData();
    }

    private movePropertiesToNextLevel(filteringValue: string | number) {

        const currentHierarchyLevel = allocationDomainService.getHierarchyLevelFromAllocationStructure(
            this.portfolioAllocationStructure,
            this.currentHierarchyLevelIndex,
        );

        this.appliedHierarchyLevels.push({
            field: currentHierarchyLevel.field,
            value: filteringValue,
        });

        this.currentHierarchyLevelIndex--;
    }

    private mapAndReferenceChartData() {

        const dataKey = this.generateDataKey();

        if(!this.chartDataMap.has(dataKey)) {
            const chartData = mapChartData(
                this.portfolioAtTime,
                this.portfolioAllocationStructure,
                this.currentHierarchyLevelIndex,
                this.appliedHierarchyLevels,
            );
            this.chartDataMap.set(dataKey, chartData);
        }

        return dataKey;
    }

    toPreviousLevel() {

        const hierarchyTopLevelIndex =
            allocationDomainService.getTopLevelHierarchyIndexFromAllocationStructure(this.portfolioAllocationStructure);

        if(this.currentHierarchyLevelIndex === hierarchyTopLevelIndex) {
            return;
        }

        this.movePropertiesToPreviousLevel();

        return this.generateDataKey();
    }

    private movePropertiesToPreviousLevel() {
        this.appliedHierarchyLevels.pop();
        this.currentHierarchyLevelIndex++;
    }

    private generateDataKey() {
        const currentHierarchyLevel = this.getCurrentHierarchyLevel();
        return generateDataKey(currentHierarchyLevel.field, this.appliedHierarchyLevels);
    }

    public getCurrentHierarchyLevel(): AllocationHierarchyLevel {
        return allocationDomainService.getHierarchyLevelFromAllocationStructure(
            this.portfolioAllocationStructure,
            this.currentHierarchyLevelIndex,
        );
    }

    public getLastAppliedHierarchyLevel(): AppliedAllocationHierarchyLevel {
        return this.appliedHierarchyLevels[this.appliedHierarchyLevels.length - 1];
    }
}

/**
 * @return key format: "A(accumulationProperty)F(field1=value1,field2=value2,...)"
 */
function generateDataKey(
    accumulationProperty: string,
    hierarchicalFilteringProperties?: AppliedAllocationHierarchyLevel[],
) {

    let dataKey = "";

    if(hierarchicalFilteringProperties && hierarchicalFilteringProperties.length > 0) {

        dataKey += "F(";

        const filteringPropertiesAsString = hierarchicalFilteringProperties.map((filter) => {
            dataKey += `${ filter.field }=${ filter.value }`;
        });
        dataKey += filteringPropertiesAsString.join(",");

        dataKey += ")";
    }

    dataKey += `A(${ accumulationProperty })`;
    return dataKey;
}