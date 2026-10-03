# Graph Report - web-static  (2026-10-03)

## Corpus Check
- 83 files · ~33,656 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 679 nodes · 1542 edges · 37 communities (32 shown, 5 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 78 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `9fd8de78`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan-chart.ts
- chart-utils.ts
- portfolio-chart.ts
- asset-composed-columns-input/constants.ts
- asset.ts
- handlebars-lang.ts
- package.json
- allocation-plan-management.ts
- service/index.ts
- devDependencies
- binding-financial-input.ts
- fractal-allocation-plan-mapping.ts
- notifications.ts
- Portfolio Detail Page
- application/index.ts
- chart-contents.ts
- dependencies
- compilerOptions
- asset-composed-columns-input/index.ts
- portfolio-history-management.ts
- infra.ts
- Asset Search Autocomplete
- .proxyrc.js
- Portfolio Section Navigation
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- AssetRowController
- AssetSearchAutocompleteController
- scripts
- dom-utils.ts
- autocomplete-interactions.ts
- allocation-plan.ts
- dom/index.ts

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 19 edges
3. `registerHandlebarsLangHelpers()` - 16 edges
4. `LogLevel` - 15 edges
5. `handlebars` - 14 edges
6. `AssetSearchAutocompleteController` - 14 edges
7. `bignumber.js` - 13 edges
8. `FractalPortfolioMultiChartDataSource` - 13 edges
9. `AutocompleteControllerActions` - 13 edges
10. `MultiChartDataSource` - 11 edges

## Surprising Connections (you probably didn't know these)
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `getNextPortfolioHistoryManagementIndex()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/portfolio-history-management.ts → websrc/utils/lang.ts
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

## Communities (37 total, 5 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (82): navigo, APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField() (+74 more)

### Community 1 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 2 - "chart-utils.ts"
Cohesion: 0.06
Nodes (37): chartjs-plugin-datalabels, chroma-js, patternomaly, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions() (+29 more)

### Community 3 - "portfolio-chart.ts"
Cohesion: 0.14
Nodes (18): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+10 more)

### Community 4 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.14
Nodes (23): api, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_INPUT_ATTRIBUTE (+15 more)

### Community 5 - "asset.ts"
Cohesion: 0.09
Nodes (32): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetPage, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft (+24 more)

### Community 6 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (44): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper() (+36 more)

### Community 7 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 8 - "allocation-plan-management.ts"
Cohesion: 0.17
Nodes (8): addPlannedAllocationRow(), AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), setHierarchicalIdFromParentRow(), toInt()

### Community 9 - "service/index.ts"
Cohesion: 0.18
Nodes (17): CompleteAllocationPlan, SerializableCompleteAllocationPlan, SerializablePortfolioCompleteAllocationPlanSet, Portfolio, PortfolioDTO, AllocationDomainService, mapAllocationStructure(), mapToAllocationPlan() (+9 more)

### Community 10 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 11 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 12 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.20
Nodes (18): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX, getAllocationHierarchySize(), getHierarchicalIdAsString() (+10 more)

### Community 13 - "notifications.ts"
Cohesion: 0.15
Nodes (11): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, CustomEventHandler, Notification, NotificationType, ERROR (+3 more)

### Community 14 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 15 - "application/index.ts"
Cohesion: 0.19
Nodes (10): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), allocationPlanManagement, AssetComposedColumnsInput, portfolioHistoryManagement, ObservationTimestamp, DivergenceAnalysis, PotentialDivergence (+2 more)

### Community 16 - "chart-contents.ts"
Cohesion: 0.40
Nodes (5): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartContent

### Community 17 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (18): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_TICKER_INPUT_ATTRIBUTE, AssetSelectionState, EXISTING, NEW, PENDING (+10 more)

### Community 20 - "portfolio-history-management.ts"
Cohesion: 0.20
Nodes (3): htmx.org, FormRowValueElements, getNextPortfolioHistoryManagementIndex()

### Community 21 - "infra.ts"
Cohesion: 0.17
Nodes (11): Application, DomInfra, handlebarsInfra, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+3 more)

### Community 22 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 23 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 24 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 31 - "AssetSearchAutocompleteController"
Cohesion: 0.20
Nodes (5): AssetSearchAutocompleteController, loadDatalists(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 32 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 33 - "dom-utils.ts"
Cohesion: 0.19
Nodes (10): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), addRemoveObserver(), contextDataCache, ensureSharedObserver() (+2 more)

### Community 34 - "autocomplete-interactions.ts"
Cohesion: 0.25
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 35 - "allocation-plan.ts"
Cohesion: 0.23
Nodes (10): bignumber.js, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationPlan, AllocationPlanDTO, PlannedAllocation, PlannedAllocationDTO (+2 more)

### Community 36 - "dom/index.ts"
Cohesion: 0.36
Nodes (7): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), maskNumberDecimalPlaces(), maskTagInput(), maskTickerInput()

## Knowledge Gaps
- **122 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+117 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 167 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **5 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan.ts` to `logger`, `chart-utils.ts`, `portfolio-chart.ts`, `handlebars-lang.ts`, `package.json`, `allocation-plan-management.ts`, `application/index.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.081) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `portfolio-history-management.ts` to `logger`, `asset.ts`, `package.json`, `allocation-plan-management.ts`, `service/index.ts`, `asset-composed-columns-input/index.ts`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `chart-utils.ts`, `asset.ts`, `package.json`, `allocation-plan-management.ts`, `notifications.ts`, `application/index.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _122 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.053289473684210525 - nodes in this community are weakly interconnected._
- **Should `chart-utils.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05727644652250146 - nodes in this community are weakly interconnected._