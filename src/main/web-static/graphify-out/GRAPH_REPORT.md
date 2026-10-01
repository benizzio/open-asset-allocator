# Graph Report - web-static  (2026-10-01)

## Corpus Check
- 77 files · ~29,368 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 598 nodes · 1343 edges · 32 communities (26 shown, 6 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `253655c4`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan.ts
- handlebars-lang.ts
- package.json
- asset.ts
- binding-financial-input.ts
- allocation-plan-chart.ts
- chart-utils.ts
- portfolio-chart.ts
- chart.ts
- dependencies
- dom/index.ts
- devDependencies
- FractalPortfolioMultiChartDataSource
- ChartDataSource
- Portfolio Detail Page
- AssetComposedColumnInput
- compilerOptions
- chart-contents.ts
- allocation-plan-management.ts
- .proxyrc.js
- htmx/index.ts
- eslint.config.mjs
- Portfolio Section Navigation
- scripts
- Frontend Module Architecture
- Recursive Planned Allocation Rows
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- alias

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `registerHandlebarsLangHelpers()` - 16 edges
3. `LogLevel` - 15 edges
4. `handlebars` - 14 edges
5. `bignumber.js` - 13 edges
6. `FractalPortfolioMultiChartDataSource` - 13 edges
7. `AssetComposedColumnInput` - 12 edges
8. `MultiChartDataSource` - 11 edges
9. `chart.js` - 10 edges
10. `compilerOptions` - 10 edges

## Surprising Connections (you probably didn't know these)
- `registerTransformResponseFunction()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/htmx/binding-htmx-transform-response.ts → websrc/infra/logging.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPortfolioMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/infra/chart/chart-types.ts
- `changeChartData()` --calls--> `changeChartDataOnDatasource()`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart.ts → websrc/infra/chart/chart-utils.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (32 total, 6 thin omitted)

### Community 0 - "logger"
Cohesion: 0.09
Nodes (52): navigo, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal() (+44 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (42): bignumber.js, AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+34 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (42): handlebars, handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper() (+34 more)

### Community 3 - "package.json"
Cohesion: 0.12
Nodes (15): bootstrap-icons, bootswatch, chroma-js, eslint, htmx-ext-client-side-templates, htmx-ext-form-json, parcel, @parcel/transformer-raw (+7 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (29): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+21 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.14
Nodes (22): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants() (+14 more)

### Community 6 - "allocation-plan-chart.ts"
Cohesion: 0.17
Nodes (12): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+4 more)

### Community 7 - "chart-utils.ts"
Cohesion: 0.16
Nodes (15): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD (+7 more)

### Community 8 - "portfolio-chart.ts"
Cohesion: 0.23
Nodes (13): chart.js, getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, MappedChartData, AllocationStructure (+5 more)

### Community 9 - "chart.ts"
Cohesion: 0.20
Nodes (14): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+6 more)

### Community 10 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 11 - "dom/index.ts"
Cohesion: 0.10
Nodes (23): BootstrapClasses, BootstrapIconClasses, bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay() (+15 more)

### Community 12 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 13 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.26
Nodes (5): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AppliedAllocationHierarchyLevel

### Community 14 - "ChartDataSource"
Cohesion: 0.15
Nodes (3): ChartDataSource, ChartDataSourceVisitor, SingleChartDataSource

### Community 15 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 16 - "AssetComposedColumnInput"
Cohesion: 0.26
Nodes (3): AssetComposedColumnInput, getAsset(), Asset

### Community 17 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 18 - "chart-contents.ts"
Cohesion: 0.24
Nodes (8): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, PortfolioDTO, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 19 - "allocation-plan-management.ts"
Cohesion: 0.05
Nodes (41): bootstrap, htmx.org, Application, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation() (+33 more)

### Community 20 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 21 - "htmx/index.ts"
Cohesion: 0.10
Nodes (30): api, APIError, APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse (+22 more)

### Community 22 - "eslint.config.mjs"
Cohesion: 0.33
Nodes (5): @eslint/js, eslint-plugin-sonarjs, globals, @stylistic/eslint-plugin, typescript-eslint

### Community 23 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 24 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **117 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+112 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 162 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **6 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan.ts` to `logger`, `handlebars-lang.ts`, `package.json`, `binding-financial-input.ts`, `chart-utils.ts`, `portfolio-chart.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.096) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `handlebars-lang.ts`, `binding-financial-input.ts`, `dom/index.ts`, `allocation-plan-management.ts`, `htmx/index.ts`?**
  _High betweenness centrality (0.067) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `package.json`, `asset.ts`, `binding-financial-input.ts`, `chart.ts`, `dom/index.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _117 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.08644067796610169 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.08888888888888889 - nodes in this community are weakly interconnected._