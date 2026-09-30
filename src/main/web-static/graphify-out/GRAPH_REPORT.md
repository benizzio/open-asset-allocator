# Graph Report - web-static  (2026-09-30)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 617 nodes · 1393 edges · 29 communities (25 shown, 4 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `4e5619ea`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan-management.ts
- allocation-plan.ts
- binding-htmx-trigger-on-route.ts
- logger
- asset.ts
- handlebars-lang.ts
- asset-composed-columns-input.ts
- binding-financial-input.ts
- portfolio-chart.ts
- package.json
- chart-utils.ts
- chart.ts
- allocation-plan-chart.ts
- MultiChartDataSource
- devDependencies
- dependencies
- FractalPortfolioMultiChartDataSource
- Portfolio Detail Page
- compilerOptions
- .proxyrc.js
- Asset Ticker Autocomplete
- scripts
- Portfolio Section Navigation
- Frontend Module Architecture
- ChartDataType
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetComposedColumnInput` - 17 edges
3. `registerHandlebarsLangHelpers()` - 16 edges
4. `LogLevel` - 15 edges
5. `handlebars` - 14 edges
6. `FractalPortfolioMultiChartDataSource` - 13 edges
7. `bignumber.js` - 13 edges
8. `MultiChartDataSource` - 11 edges
9. `NotificationType` - 10 edges
10. `bindFinancialInput()` - 10 edges

## Surprising Connections (you probably didn't know these)
- `Asset Composed Columns Input` --references--> `AssetComposedColumnsInput`  [EXTRACTED]
  websrc/components/asset-composed-columns-input.html → websrc/components/asset-composed-columns-input.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPlannedAllocationMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/infra/chart/chart-types.ts
- `FractalPortfolioMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/infra/chart/chart-types.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (29 total, 4 thin omitted)

### Community 0 - "allocation-plan-management.ts"
Cohesion: 0.05
Nodes (37): htmx.org, Recursive Planned Allocation Rows, Asset Composed Columns Input, Application, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), addPlannedAllocationRow(), allocationPlanManagement (+29 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (43): bignumber.js, AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO (+35 more)

### Community 2 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.08
Nodes (45): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, RequestConfigEventDetail, addAttributes(), addRouterHooks() (+37 more)

### Community 3 - "logger"
Cohesion: 0.09
Nodes (45): APIError, APIErrorResponse, addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), bindPercentageInput() (+37 more)

### Community 4 - "asset.ts"
Cohesion: 0.07
Nodes (41): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, Asset, ExternalAsset, CustomEventHandler, Notification (+33 more)

### Community 5 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (44): handlebars, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+36 more)

### Community 6 - "asset-composed-columns-input.ts"
Cohesion: 0.12
Nodes (20): api, ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, AssetTickerAutocompleteOption, AssetTickerAutocompleteState, autocompleteStates, clearAssetTickerOptionHighlight(), closeAssetTickerAutocomplete() (+12 more)

### Community 7 - "binding-financial-input.ts"
Cohesion: 0.13
Nodes (24): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes() (+16 more)

### Community 8 - "portfolio-chart.ts"
Cohesion: 0.18
Nodes (17): chart.js, allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData() (+9 more)

### Community 9 - "package.json"
Cohesion: 0.10
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, globals, htmx-ext-client-side-templates (+12 more)

### Community 10 - "chart-utils.ts"
Cohesion: 0.14
Nodes (17): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+9 more)

### Community 11 - "chart.ts"
Cohesion: 0.19
Nodes (14): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, LocalChartOptions, MeasuramentUnit (+6 more)

### Community 12 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 13 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 14 - "devDependencies"
Cohesion: 0.13
Nodes (15): devDependencies, eslint, @eslint/js, globals, http-proxy-middleware, parcel, @parcel/transformer-raw, @parcel/transformer-sass (+7 more)

### Community 15 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 16 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.24
Nodes (5): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AppliedAllocationHierarchyLevel

### Community 17 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 20 - "Asset Ticker Autocomplete"
Cohesion: 0.53
Nodes (6): Unnamed Asset Search Field (filters the complete ticker-and-name label), Asset Ticker Autocomplete, Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 21 - "scripts"
Cohesion: 0.50
Nodes (4): scripts, build, clean, dev

### Community 22 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 23 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 24 - "ChartDataType"
Cohesion: 0.67
Nodes (3): ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

## Knowledge Gaps
- **118 isolated node(s):** `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure`, `DivergenceAnalysis`, `GeneralErrorHandler`, `PlannedAllocationDTO` (+113 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 163 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan.ts` to `allocation-plan-management.ts`, `logger`, `handlebars-lang.ts`, `portfolio-chart.ts`, `package.json`, `chart-utils.ts`?**
  _High betweenness centrality (0.091) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `allocation-plan-management.ts`, `binding-htmx-trigger-on-route.ts`, `handlebars-lang.ts`, `binding-financial-input.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `allocation-plan-management.ts`, `package.json`, `chart.ts`, `asset.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure`, `DivergenceAnalysis` to the rest of the system?**
  _118 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan-management.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05254237288135593 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09147869674185463 - nodes in this community are weakly interconnected._