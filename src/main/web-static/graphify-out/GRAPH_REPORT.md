# Graph Report - web-static  (2026-10-08)

## Corpus Check
- 86 files · ~37,092 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 721 nodes · 1640 edges · 34 communities (31 shown, 3 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 82 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `039bbce1`
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
- allocation-plan-chart.ts
- portfolio-history-management.ts
- MultiChartDataSource
- devDependencies
- portfolio-chart.ts
- dependencies
- infra.ts
- chart-contents.ts
- notifications.ts
- chart.ts
- Portfolio Detail Page
- createPortfolioHistoryQuoteAction
- compilerOptions
- dom/index.ts
- Asset Search Autocomplete
- .proxyrc.js
- quote-action.ts
- scripts
- Portfolio Section Navigation
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

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
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `getNextPortfolioHistoryManagementIndex()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/portfolio-history-management.ts → websrc/utils/lang.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (34 total, 3 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (83): api, APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField() (+75 more)

### Community 1 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.06
Nodes (55): AssetSearchAutocompleteController, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR (+47 more)

### Community 2 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (44): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+36 more)

### Community 3 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (47): handlebars, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), andHelper(), arrayHelper(), comparatorHelper() (+39 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+22 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.14
Nodes (22): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants() (+14 more)

### Community 6 - "allocation-plan-management.ts"
Cohesion: 0.11
Nodes (14): addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), setHierarchicalIdFromParentRow() (+6 more)

### Community 7 - "chart-utils.ts"
Cohesion: 0.14
Nodes (17): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+9 more)

### Community 8 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 9 - "AssetRowController"
Cohesion: 0.15
Nodes (7): AssetRowController, handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), selectAssetTicker(), AutocompleteControllerActions

### Community 10 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (10): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+2 more)

### Community 11 - "portfolio-history-management.ts"
Cohesion: 0.15
Nodes (10): htmx.org, FormRowValueElements, getNextPortfolioHistoryManagementIndex(), getTrimmedExternalAssetKeys(), handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction (+2 more)

### Community 12 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 13 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 14 - "portfolio-chart.ts"
Cohesion: 0.14
Nodes (18): bignumber.js, chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback() (+10 more)

### Community 15 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 16 - "infra.ts"
Cohesion: 0.28
Nodes (8): bootstrap, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), setupGlobalErrorHandler(), CustomEventHandler, Router

### Community 17 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 18 - "notifications.ts"
Cohesion: 0.20
Nodes (8): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, Notification, NotificationType, ERROR, INFO, SUCCESS, WARNING

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

### Community 23 - "dom/index.ts"
Cohesion: 0.12
Nodes (18): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+10 more)

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 25 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 26 - "quote-action.ts"
Cohesion: 0.21
Nodes (8): Application, notifications, PortfolioHistoryQuoteAction, RecalculatePortfolioAllocation, AfterRequestEventDetail, HtmxInfra, RequestConfigEventDetail, Infra

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

## Knowledge Gaps
- **128 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+123 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 172 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `portfolio-chart.ts` to `logger`, `allocation-plan.ts`, `handlebars-lang.ts`, `binding-financial-input.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `portfolio-history-management.ts`, `quote-action.ts`?**
  _High betweenness centrality (0.089) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `portfolio-history-management.ts` to `logger`, `asset-composed-columns-input/constants.ts`, `allocation-plan.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `binding-financial-input.ts`, `allocation-plan-management.ts`, `package.json`, `portfolio-history-management.ts`, `notifications.ts`, `chart.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _128 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05197594501718213 - nodes in this community are weakly interconnected._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05754385964912281 - nodes in this community are weakly interconnected._