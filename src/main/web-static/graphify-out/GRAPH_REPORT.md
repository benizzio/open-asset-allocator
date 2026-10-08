# Graph Report - web-static  (2026-10-08)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 711 nodes · 1618 edges · 37 communities (33 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 80 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `4256a2f9`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan.ts
- portfolio-chart.ts
- handlebars-lang.ts
- binding-financial-input.ts
- allocation-plan-management.ts
- asset-composed-columns-input/index.ts
- infra.ts
- asset-composed-columns-input/constants.ts
- package.json
- asset-external.ts
- AssetRowController
- createPortfolioHistoryQuoteAction
- AssetSearchAutocompleteController
- asset.ts
- MultiChartDataSource
- devDependencies
- chart-types.ts
- portfolio-history-management.ts
- chart-utils.ts
- autocomplete-interactions.ts
- dependencies
- handlebars/index.ts
- Portfolio Detail Page
- compilerOptions
- chart.ts
- dom-utils.ts
- Asset Search Autocomplete
- .proxyrc.js
- scripts
- NotificationType
- Portfolio Section Navigation
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 21 edges
3. `registerHandlebarsLangHelpers()` - 16 edges
4. `createPortfolioHistoryQuoteAction()` - 15 edges
5. `LogLevel` - 15 edges
6. `AssetSearchAutocompleteController` - 14 edges
7. `bignumber.js` - 14 edges
8. `handlebars` - 14 edges
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `getNextPortfolioHistoryManagementIndex()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/portfolio-history-management.ts → websrc/utils/lang.ts
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (37 total, 4 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (84): APIError, APIErrorResponse, addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), bindPercentageInput() (+76 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.07
Nodes (55): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+47 more)

### Community 2 - "portfolio-chart.ts"
Cohesion: 0.09
Nodes (28): bignumber.js, chart.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey() (+20 more)

### Community 3 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (37): addPlannedAllocationRow(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper(), ifNotEqualsHelper() (+29 more)

### Community 4 - "binding-financial-input.ts"
Cohesion: 0.12
Nodes (25): BootstrapClasses, BootstrapIconClasses, bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), bindFinancialInput(), bindFinancialInputElements() (+17 more)

### Community 5 - "allocation-plan-management.ts"
Cohesion: 0.11
Nodes (15): htmx.org, allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), setHierarchicalIdFromParentRow() (+7 more)

### Community 6 - "asset-composed-columns-input/index.ts"
Cohesion: 0.12
Nodes (22): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_TICKER_INPUT_ATTRIBUTE, AssetSelectionState, EXISTING, NEW, PENDING (+14 more)

### Community 7 - "infra.ts"
Cohesion: 0.12
Nodes (16): bootstrap, chartjs-plugin-datalabels, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, DomInfra, handlebarsInfra, HtmxInfra (+8 more)

### Community 8 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.16
Nodes (21): api, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_INPUT_ATTRIBUTE (+13 more)

### Community 9 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 10 - "asset-external.ts"
Cohesion: 0.17
Nodes (21): addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields(), handleExternalSearchAfterRequest(), handleExternalSearchBeforeRequest(), initializeExternalDraft() (+13 more)

### Community 12 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.19
Nodes (16): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeRequest(), handleRequestConfiguration(), isSnapshotCurrent(), parseLatestClosingQuote() (+8 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.20
Nodes (5): AssetSearchAutocompleteController, loadDatalists(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "asset.ts"
Cohesion: 0.14
Nodes (10): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, hasValidExternalData(), isExternalAsset(), getAssetIdentifierFromLocation(), getCurrentAssetIdentifierForResponse() (+2 more)

### Community 15 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 16 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 17 - "chart-types.ts"
Cohesion: 0.17
Nodes (14): Application, CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, ChartInteraction, ChartInteractions (+6 more)

### Community 18 - "portfolio-history-management.ts"
Cohesion: 0.17
Nodes (9): FormRowValueElements, getNextPortfolioHistoryManagementIndex(), getTrimmedExternalAssetKeys(), handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, recalculatePortfolioHistoryAllocation() (+1 more)

### Community 19 - "chart-utils.ts"
Cohesion: 0.22
Nodes (10): chroma-js, patternomaly, PIE_DOUGHNUT_CHART_OPTIONS, UNIDIMENSIONAL_DATASET_SUM_FIELD, getDatasetSum(), LABEL_CALLBACKS, valueAsPercentageOfDataset(), BOOTSTRAP_BODY_BACKGROUND_COLOR (+2 more)

### Community 20 - "autocomplete-interactions.ts"
Cohesion: 0.26
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 21 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 22 - "handlebars/index.ts"
Cohesion: 0.25
Nodes (9): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY, registerPartialToContainer() (+1 more)

### Community 23 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 24 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 25 - "chart.ts"
Cohesion: 0.24
Nodes (9): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), convertUnidimensionalDatasetBackgroundToPattern() (+1 more)

### Community 26 - "dom-utils.ts"
Cohesion: 0.28
Nodes (5): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements

### Community 27 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 28 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 29 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 30 - "NotificationType"
Cohesion: 0.40
Nodes (5): NotificationType, ERROR, INFO, SUCCESS, WARNING

### Community 31 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 32 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 33 - "Observation Editor"
Cohesion: 0.67
Nodes (3): External Asset Quote Action, Observation Editor, Portfolio History Management

## Knowledge Gaps
- **123 isolated node(s):** `APIError`, `EventDetail`, `Level`, `Levels`, `AllocationHierarchyLevelDTO` (+118 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 168 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `portfolio-chart.ts` to `logger`, `allocation-plan.ts`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `package.json`, `createPortfolioHistoryQuoteAction`, `portfolio-history-management.ts`, `chart-utils.ts`?**
  _High betweenness centrality (0.088) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `logger`, `asset-composed-columns-input/index.ts`, `package.json`, `asset.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `portfolio-chart.ts`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `infra.ts`, `package.json`, `asset.ts`, `chart-types.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `APIError`, `EventDetail`, `Level` to the rest of the system?**
  _123 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05153576582148011 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06631578947368422 - nodes in this community are weakly interconnected._