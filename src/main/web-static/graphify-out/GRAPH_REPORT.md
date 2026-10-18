# Graph Report - web-static  (2026-10-18)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 738 nodes · 1682 edges · 36 communities (32 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 83 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `aa4d0dc1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan.ts
- portfolio-history-management.ts
- asset.ts
- handlebars-lang.ts
- binding-financial-input.ts
- asset-composed-columns-input/constants.ts
- allocation-plan-management.ts
- chart-utils.ts
- dom-utils.ts
- package.json
- AssetRowController
- AssetSearchAutocompleteController
- asset-composed-columns-input/index.ts
- chart.ts
- MultiChartDataSource
- devDependencies
- dependencies
- portfolio-chart.ts
- portfolio-allocation.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- autocomplete-interactions.ts
- compilerOptions
- FractalPortfolioMultiChartDataSource
- Asset Search Autocomplete
- chart-contents.ts
- .proxyrc.js
- scripts
- AssetSelectionState
- Portfolio Section Navigation
- Observation Editor
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 21 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `createPortfolioHistoryQuoteAction()` - 16 edges
5. `LogLevel` - 15 edges
6. `AssetSearchAutocompleteController` - 14 edges
7. `bignumber.js` - 14 edges
8. `handlebars` - 14 edges
9. `AutocompleteControllerActions` - 13 edges
10. `FractalPortfolioMultiChartDataSource` - 13 edges

## Surprising Connections (you probably didn't know these)
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `registerTransformResponseFunction()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/htmx/binding-htmx-transform-response.ts → websrc/infra/logging.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPlannedAllocationMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/infra/chart/chart-types.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (36 total, 4 thin omitted)

### Community 0 - "logger"
Cohesion: 0.07
Nodes (67): navigo, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal() (+59 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.07
Nodes (51): htmx.org, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+43 more)

### Community 2 - "portfolio-history-management.ts"
Cohesion: 0.05
Nodes (49): api, APIError, APIErrorResponse, Application, notifications, FormRowValueElements, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow() (+41 more)

### Community 3 - "asset.ts"
Cohesion: 0.07
Nodes (40): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR (+32 more)

### Community 4 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (42): bignumber.js, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+34 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.10
Nodes (30): BootstrapClasses, BootstrapIconClasses, bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay() (+22 more)

### Community 6 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.13
Nodes (27): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+19 more)

### Community 7 - "allocation-plan-management.ts"
Cohesion: 0.10
Nodes (15): allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), AssetComposedColumnsInput, portfolioHistoryManagement (+7 more)

### Community 8 - "chart-utils.ts"
Cohesion: 0.13
Nodes (18): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+10 more)

### Community 9 - "dom-utils.ts"
Cohesion: 0.14
Nodes (14): handlebars, addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements, domJSONHelper() (+6 more)

### Community 10 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 12 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 13 - "asset-composed-columns-input/index.ts"
Cohesion: 0.13
Nodes (17): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+9 more)

### Community 14 - "chart.ts"
Cohesion: 0.21
Nodes (13): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent, LocalChartOptions (+5 more)

### Community 15 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 16 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 17 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 18 - "portfolio-chart.ts"
Cohesion: 0.26
Nodes (11): chart.js, changeChartData(), chartDataSelectionEventHandler(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation (+3 more)

### Community 19 - "portfolio-allocation.ts"
Cohesion: 0.20
Nodes (10): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshot, PortfolioSnapshotDTO, DivergenceAnalysis (+2 more)

### Community 20 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 21 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 22 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 23 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 24 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.29
Nodes (3): FractalPortfolioMultiChartDataSource, generateDataKey(), AppliedAllocationHierarchyLevel

### Community 25 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 26 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 27 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 28 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 29 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 30 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 31 - "Observation Editor"
Cohesion: 0.50
Nodes (4): External Asset Quote Action, Observation Date Picker, Observation Editor, Portfolio History Management

### Community 32 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **130 isolated node(s):** `GeneralErrorHandler`, `Level`, `Levels`, `AllocationHierarchyLevelDTO`, `PlannedAllocationDTO` (+125 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 175 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `handlebars-lang.ts` to `logger`, `allocation-plan.ts`, `portfolio-history-management.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `portfolio-allocation.ts`?**
  _High betweenness centrality (0.090) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan.ts` to `logger`, `portfolio-history-management.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `asset-composed-columns-input/index.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `handlebars` connect `dom-utils.ts` to `portfolio-history-management.ts`, `asset.ts`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `package.json`, `chart.ts`, `portfolio-allocation.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `GeneralErrorHandler`, `Level`, `Levels` to the rest of the system?**
  _130 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.06596035543403965 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.0745814307458143 - nodes in this community are weakly interconnected._