# Graph Report - web-static  (2026-10-10)

## Corpus Check
- 86 files · ~37,490 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 729 nodes · 1661 edges · 36 communities (31 shown, 5 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 83 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `c6acf194`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- asset-composed-columns-input/constants.ts
- allocation-plan.ts
- handlebars-lang.ts
- asset.ts
- binding-financial-input.ts
- allocation-plan-management.ts
- chart-utils.ts
- package.json
- AssetRowController
- dom/index.ts
- portfolio-history-management.ts
- MultiChartDataSource
- devDependencies
- portfolio-chart.ts
- dependencies
- infra.ts
- allocation-plan-chart.ts
- notifications.ts
- chart.ts
- Portfolio Detail Page
- createPortfolioHistoryQuoteAction
- compilerOptions
- application/index.ts
- Asset Search Autocomplete
- .proxyrc.js
- FractalPortfolioMultiChartDataSource
- scripts
- Portfolio Section Navigation
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- dom-utils.ts
- FormRowValueElements

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 21 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `createPortfolioHistoryQuoteAction()` - 16 edges
5. `LogLevel` - 15 edges
6. `bignumber.js` - 14 edges
7. `handlebars` - 14 edges
8. `AssetSearchAutocompleteController` - 14 edges
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPlannedAllocationMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/infra/chart/chart-types.ts
- `FractalPortfolioMultiChartDataSource` --references--> `AppliedAllocationHierarchyLevel`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/application/portfolio-chart/portfolio-chart-model.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (36 total, 5 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (82): APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal() (+74 more)

### Community 1 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.05
Nodes (61): api, APIError, AssetSearchAutocompleteController, handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ARIA_ACTIVE_DESCENDANT_ATTRIBUTE (+53 more)

### Community 2 - "allocation-plan.ts"
Cohesion: 0.07
Nodes (45): FractalPlannedAllocationMultiChartDataSource, getSelectedDataKey(), interactionObserverCallback(), AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN (+37 more)

### Community 3 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (49): handlebars, setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), andHelper() (+41 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+22 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.14
Nodes (22): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants() (+14 more)

### Community 6 - "allocation-plan-management.ts"
Cohesion: 0.13
Nodes (9): htmx.org, addPlannedAllocationRow(), AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), Router (+1 more)

### Community 7 - "chart-utils.ts"
Cohesion: 0.14
Nodes (17): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+9 more)

### Community 8 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 10 - "dom/index.ts"
Cohesion: 0.18
Nodes (14): BootstrapClasses, BootstrapIconClasses, bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay() (+6 more)

### Community 11 - "portfolio-history-management.ts"
Cohesion: 0.19
Nodes (15): bignumber.js, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR (+7 more)

### Community 12 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 13 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 14 - "portfolio-chart.ts"
Cohesion: 0.20
Nodes (16): chart.js, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel (+8 more)

### Community 15 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 16 - "infra.ts"
Cohesion: 0.32
Nodes (7): HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra, setupGlobalErrorHandler()

### Community 17 - "allocation-plan-chart.ts"
Cohesion: 0.19
Nodes (14): allocationPlanChart, chartDataSelectionEventHandler(), getChartContent(), mapChildDatasets(), mapDataset(), toChartDataMap(), toChartContent(), toUnidimensionalMultiChartContent() (+6 more)

### Community 18 - "notifications.ts"
Cohesion: 0.15
Nodes (11): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR (+3 more)

### Community 19 - "chart.ts"
Cohesion: 0.19
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+7 more)

### Community 20 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 21 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 22 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 23 - "application/index.ts"
Cohesion: 0.18
Nodes (9): Application, allocationPlanManagement, AssetComposedColumnsInput, notifications, portfolioHistoryManagement, AfterRequestEventDetail, AssetPage, websrc_pages_index_assetpage (+1 more)

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 25 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 26 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.27
Nodes (3): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource

### Community 27 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 28 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 29 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 30 - "Observation Editor"
Cohesion: 0.67
Nodes (3): External Asset Quote Action, Observation Editor, Portfolio History Management

### Community 34 - "dom-utils.ts"
Cohesion: 0.28
Nodes (5): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements

## Knowledge Gaps
- **130 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+125 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 175 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **5 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `portfolio-history-management.ts` to `logger`, `allocation-plan.ts`, `handlebars-lang.ts`, `binding-financial-input.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `portfolio-chart.ts`?**
  _High betweenness centrality (0.091) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `logger`, `asset-composed-columns-input/constants.ts`, `asset.ts`, `package.json`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `binding-financial-input.ts`, `allocation-plan-management.ts`, `package.json`, `portfolio-history-management.ts`, `notifications.ts`, `chart.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _130 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05352743561030235 - nodes in this community are weakly interconnected._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.051590483827853514 - nodes in this community are weakly interconnected._