# Graph Report - web-static  (2026-10-16)

## Corpus Check
- 87 files · ~37,851 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 733 nodes · 1673 edges · 51 communities (46 shown, 5 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 83 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e42df97c`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- htmx/index.ts
- asset-composed-columns-input/constants.ts
- fractal-allocation-plan-mapping.ts
- handlebars-lang.ts
- asset.ts
- binding-financial-input.ts
- allocation-plan-management.ts
- chart-utils.ts
- package.json
- AssetRowController
- logging.ts
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
- service/index.ts
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
- autocomplete-controller.ts
- routing/index.ts
- FormRowValueElements
- handlebars/index.ts
- AssetSearchAutocompleteController
- binding-htmx-trigger-on-route.ts
- autocomplete-interactions.ts
- allocation-plan.ts
- logger
- binding-htmx-transform-response.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-attribute-on-route.ts
- bignumber.js
- binding-dom-display-on-route.ts
- AssetSelectionState
- ChartDataType

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
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `registerTransformResponseFunction()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/htmx/binding-htmx-transform-response.ts → websrc/infra/logging.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (51 total, 5 thin omitted)

### Community 0 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, bindHTMXTransformResponseInDescendants(), addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail (+8 more)

### Community 1 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.10
Nodes (37): ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+29 more)

### Community 2 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.19
Nodes (19): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX, PlannedAllocation, getAllocationHierarchySize() (+11 more)

### Community 3 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (39): getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper() (+31 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+22 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.13
Nodes (24): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes() (+16 more)

### Community 6 - "allocation-plan-management.ts"
Cohesion: 0.11
Nodes (14): addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), setHierarchicalIdFromParentRow() (+6 more)

### Community 7 - "chart-utils.ts"
Cohesion: 0.15
Nodes (16): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+8 more)

### Community 8 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 10 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 11 - "portfolio-history-management.ts"
Cohesion: 0.19
Nodes (14): handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR (+6 more)

### Community 12 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 13 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 14 - "portfolio-chart.ts"
Cohesion: 0.16
Nodes (20): chart.js, allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData() (+12 more)

### Community 15 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 16 - "infra.ts"
Cohesion: 0.15
Nodes (13): bootstrap, Application, notifications, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+5 more)

### Community 17 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (10): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+2 more)

### Community 18 - "notifications.ts"
Cohesion: 0.20
Nodes (8): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, Notification, NotificationType, ERROR, INFO, SUCCESS, WARNING

### Community 19 - "chart.ts"
Cohesion: 0.18
Nodes (14): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, LocalChartOptions, MeasuramentUnit (+6 more)

### Community 20 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 21 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 22 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 23 - "service/index.ts"
Cohesion: 0.22
Nodes (15): htmx.org, AllocationPlanDTO, CompleteAllocationPlan, Portfolio, PortfolioDTO, AllocationDomainService, mapAllocationStructure(), mapToAllocationPlan() (+7 more)

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 25 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 26 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.26
Nodes (4): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey()

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

### Community 35 - "autocomplete-controller.ts"
Cohesion: 0.15
Nodes (14): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_STATUS_ATTRIBUTE, ASSET_SEARCH_SUGGESTIONS_ATTRIBUTE (+6 more)

### Community 36 - "routing/index.ts"
Cohesion: 0.28
Nodes (13): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), buildParameterizedDestinationPathFromCurrentLocationContext(), hasPathParam() (+5 more)

### Community 38 - "handlebars/index.ts"
Cohesion: 0.23
Nodes (10): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY, registerPartialToContainer() (+2 more)

### Community 39 - "AssetSearchAutocompleteController"
Cohesion: 0.28
Nodes (3): AssetSearchAutocompleteController, loadDatalists(), AssetSearchAutocompleteState

### Community 40 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.26
Nodes (12): RequestConfigEventDetail, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute() (+4 more)

### Community 41 - "autocomplete-interactions.ts"
Cohesion: 0.32
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_SEARCH_OPTION_SELECTOR, AutocompleteControllerActions

### Community 42 - "allocation-plan.ts"
Cohesion: 0.21
Nodes (10): AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationPlan, FractalHierarchicalAllocationPlan, PlannedAllocationDTO, SerializableCompleteAllocationPlan, SerializableFractalHierarchicalAllocationPlan (+2 more)

### Community 43 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 44 - "binding-htmx-transform-response.ts"
Cohesion: 0.31
Nodes (8): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse(), BeforeSwapEventDetail

### Community 45 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 46 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 47 - "bignumber.js"
Cohesion: 0.32
Nodes (6): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

### Community 48 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 49 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 50 - "ChartDataType"
Cohesion: 0.67
Nodes (3): ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

## Knowledge Gaps
- **130 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+125 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 175 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **5 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `bignumber.js` to `handlebars-lang.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `allocation-plan.ts`, `portfolio-history-management.ts`, `logger`, `portfolio-chart.ts`, `service/index.ts`?**
  _High betweenness centrality (0.090) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `service/index.ts` to `htmx/index.ts`, `asset-composed-columns-input/constants.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `binding-htmx-trigger-on-route.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `handlebars-lang.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `portfolio-history-management.ts`, `bignumber.js`, `notifications.ts`, `chart.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _130 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09830866807610994 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09872241579558652 - nodes in this community are weakly interconnected._