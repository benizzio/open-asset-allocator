# Graph Report - web-static  (2026-10-18)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 743 nodes · 1693 edges · 40 communities (34 shown, 6 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 83 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `62a3fb13`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan.ts
- handlebars-lang.ts
- asset.ts
- asset-composed-columns-input/constants.ts
- allocation-plan-management.ts
- dom/index.ts
- package.json
- infra.ts
- chart.ts
- chart-utils.ts
- AssetRowController
- AssetSearchAutocompleteController
- portfolio-history-management.ts
- bignumber.js
- portfolio-chart.ts
- asset-composed-columns-input/index.ts
- MultiChartDataSource
- binding-financial-input.ts
- devDependencies
- dependencies
- createPortfolioHistoryQuoteAction
- htmx/index.ts
- Portfolio Detail Page
- autocomplete-interactions.ts
- compilerOptions
- FractalPortfolioMultiChartDataSource
- quote-action.ts
- Asset Search Autocomplete
- .proxyrc.js
- scripts
- AssetSelectionState
- FormRowValueElements
- NotificationType
- Portfolio Section Navigation
- Observation Editor
- Frontend Module Architecture
- Allocation Plan Date Picker
- Hierarchical Divergence Analysis
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
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
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

## Communities (40 total, 6 thin omitted)

### Community 0 - "logger"
Cohesion: 0.06
Nodes (70): APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal() (+62 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.07
Nodes (55): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+47 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (48): handlebars, addPlannedAllocationRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), andHelper() (+40 more)

### Community 3 - "asset.ts"
Cohesion: 0.09
Nodes (33): api, APIError, Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch() (+25 more)

### Community 4 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+21 more)

### Community 5 - "allocation-plan-management.ts"
Cohesion: 0.11
Nodes (16): htmx.org, Application, allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey() (+8 more)

### Community 6 - "dom/index.ts"
Cohesion: 0.12
Nodes (18): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+10 more)

### Community 7 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 8 - "infra.ts"
Cohesion: 0.12
Nodes (15): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, handlebarsInfra, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing() (+7 more)

### Community 9 - "chart.ts"
Cohesion: 0.15
Nodes (18): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+10 more)

### Community 10 - "chart-utils.ts"
Cohesion: 0.14
Nodes (17): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+9 more)

### Community 12 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 13 - "portfolio-history-management.ts"
Cohesion: 0.16
Nodes (12): mapForAllocationPlanManagement(), handleAssetRowSelectionChange(), handleNewObservationFieldFocus(), handleObservationAccordionTransition(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, setNewObservationTimestampFromDate() (+4 more)

### Community 14 - "bignumber.js"
Cohesion: 0.18
Nodes (12): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshot, PortfolioSnapshotDTO (+4 more)

### Community 15 - "portfolio-chart.ts"
Cohesion: 0.24
Nodes (13): chart.js, changeChartData(), chartDataSelectionEventHandler(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation (+5 more)

### Community 16 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (15): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+7 more)

### Community 17 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 18 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 19 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 20 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 21 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 22 - "htmx/index.ts"
Cohesion: 0.22
Nodes (12): addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs(), prepareFormData(), replaceRequestPathParams(), replaceRequestPathParamsFromCurrentRoute() (+4 more)

### Community 23 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 24 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 25 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 27 - "quote-action.ts"
Cohesion: 0.33
Nodes (9): EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR, EXTERNAL_ASSET_TICKER_SELECTOR, HISTORY_ROW_ID_PATTERN, PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, PortfolioHistoryQuoteAction, readExternalAssetKeys() (+1 more)

### Community 28 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 29 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 30 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 31 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 33 - "NotificationType"
Cohesion: 0.40
Nodes (5): NotificationType, ERROR, INFO, SUCCESS, WARNING

### Community 34 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 35 - "Observation Editor"
Cohesion: 0.50
Nodes (4): External Asset Quote Action, Observation Date Picker, Observation Editor, Portfolio History Management

### Community 36 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **131 isolated node(s):** `Level`, `Levels`, `AllocationHierarchyLevelDTO`, `PlannedAllocationDTO`, `AssetSearchAutocompleteOption` (+126 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 176 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **6 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `bignumber.js` to `logger`, `allocation-plan.ts`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `package.json`, `chart-utils.ts`, `portfolio-history-management.ts`, `quote-action.ts`?**
  _High betweenness centrality (0.089) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `logger`, `asset.ts`, `package.json`, `portfolio-history-management.ts`, `asset-composed-columns-input/index.ts`, `htmx/index.ts`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `allocation-plan-management.ts`, `package.json`, `infra.ts`, `chart.ts`, `portfolio-history-management.ts`, `bignumber.js`?**
  _High betweenness centrality (0.060) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `Level`, `Levels`, `AllocationHierarchyLevelDTO` to the rest of the system?**
  _131 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.06172839506172839 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.0687719298245614 - nodes in this community are weakly interconnected._