# Graph Report - web-static  (2026-10-18)

## Corpus Check
- 88 files · ~38,504 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 742 nodes · 1689 edges · 47 communities (41 shown, 6 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 83 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1a409726`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- routing/index.ts
- allocation-plan.ts
- htmx/index.ts
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
- application/portfolio-analysis.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- autocomplete-interactions.ts
- compilerOptions
- FractalPortfolioMultiChartDataSource
- Asset Search Autocomplete
- allocation-plan-chart.ts
- .proxyrc.js
- scripts
- AssetSelectionState
- Portfolio Section Navigation
- Observation Editor
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- portfolio-history-management.ts
- infra.ts
- notifications.ts
- binding-htmx-trigger-on-route.ts
- logger
- binding-dom-attribute-on-route.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- binding-dom-exclusive-display.ts
- FormRowValueElements
- alias

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

## Communities (47 total, 6 thin omitted)

### Community 0 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.08
Nodes (45): FractalPlannedAllocationMultiChartDataSource, getSelectedDataKey(), interactionObserverCallback(), AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN (+37 more)

### Community 2 - "htmx/index.ts"
Cohesion: 0.12
Nodes (25): APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP (+17 more)

### Community 3 - "asset.ts"
Cohesion: 0.09
Nodes (32): api, APIError, Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch() (+24 more)

### Community 4 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (51): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper() (+43 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.13
Nodes (24): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes() (+16 more)

### Community 6 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+21 more)

### Community 7 - "allocation-plan-management.ts"
Cohesion: 0.11
Nodes (15): htmx.org, allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), AssetComposedColumnsInput (+7 more)

### Community 8 - "chart-utils.ts"
Cohesion: 0.15
Nodes (16): bignumber.js, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD (+8 more)

### Community 9 - "dom-utils.ts"
Cohesion: 0.28
Nodes (5): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements

### Community 10 - "package.json"
Cohesion: 0.11
Nodes (20): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+12 more)

### Community 12 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 13 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (15): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+7 more)

### Community 14 - "chart.ts"
Cohesion: 0.19
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+7 more)

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
Cohesion: 0.20
Nodes (15): chart.js, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel (+7 more)

### Community 19 - "application/portfolio-analysis.ts"
Cohesion: 0.38
Nodes (5): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

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
Cohesion: 0.27
Nodes (3): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource

### Community 25 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 26 - "allocation-plan-chart.ts"
Cohesion: 0.19
Nodes (14): allocationPlanChart, chartDataSelectionEventHandler(), getChartContent(), mapChildDatasets(), mapDataset(), toChartDataMap(), toChartContent(), toUnidimensionalMultiChartContent() (+6 more)

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

### Community 36 - "portfolio-history-management.ts"
Cohesion: 0.13
Nodes (21): handleAssetRowSelectionChange(), handleNewObservationFieldFocus(), handleObservationAccordionTransition(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, setNewObservationTimestampFromDate(), showFocusedObservationPanel() (+13 more)

### Community 37 - "infra.ts"
Cohesion: 0.17
Nodes (11): chartjs-plugin-datalabels, Application, notifications, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+3 more)

### Community 38 - "notifications.ts"
Cohesion: 0.16
Nodes (10): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 39 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.27
Nodes (11): navigo, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute() (+3 more)

### Community 40 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 41 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 42 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 43 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 44 - "binding-dom-exclusive-display.ts"
Cohesion: 0.60
Nodes (5): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings()

## Knowledge Gaps
- **130 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+125 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 176 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **6 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `chart-utils.ts` to `allocation-plan.ts`, `portfolio-history-management.ts`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `logger`, `package.json`, `portfolio-chart.ts`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.090) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `htmx/index.ts`, `asset.ts`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`, `package.json`, `asset-composed-columns-input/index.ts`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `portfolio-history-management.ts`, `notifications.ts`, `allocation-plan-management.ts`, `package.json`, `chart.ts`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.060) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _130 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07731694828469023 - nodes in this community are weakly interconnected._
- **Should `htmx/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.11724137931034483 - nodes in this community are weakly interconnected._