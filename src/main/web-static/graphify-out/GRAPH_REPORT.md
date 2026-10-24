# Graph Report - web-static  (2026-10-24)

## Corpus Check
- 93 files · ~42,096 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 795 nodes · 1823 edges · 51 communities (48 shown, 3 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 111 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `20d52e55`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- htmx/index.ts
- AssetSearchAutocompleteController
- handlebars-lang.ts
- portfolio-history-management.ts
- fractal-allocation-plan-mapping.ts
- portfolio-history-management-observation-metadata/index.ts
- asset.ts
- allocation-plan-management.ts
- binding-financial-input.ts
- AssetRowController
- chart-utils.ts
- package.json
- portfolio-chart.ts
- chart-types.ts
- allocation-plan-chart.ts
- ChartDataSource
- service/index.ts
- devDependencies
- dependencies
- Portfolio Detail Page
- handlebars/index.ts
- compilerOptions
- chart-contents.ts
- .proxyrc.js
- Asset Search Autocomplete
- scripts
- Portfolio Section Navigation
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Toast Notification
- asset-composed-columns-input/index.ts
- asset-row-controller.ts
- Portfolio History Management
- asset-composed-columns-input/constants.ts
- infra.ts
- autocomplete-interactions.ts
- routing/index.ts
- notifications.ts
- createPortfolioHistoryQuoteAction
- allocation-plan.ts
- binding-htmx-trigger-on-route.ts
- logging.ts
- logger
- dom-utils.ts
- binding-dom-attribute-on-route.ts
- binding-htmx-transform-response.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- chart.ts
- application/portfolio-analysis.ts
- AssetSelectionState

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
- `Allocation Plan Management` --references--> `mapForAllocationPlanManagement()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `Allocation Plan Pencil Metadata Editor` --references--> `handleInputAllocationPlanDate()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `Recursive Planned Allocation Rows` --references--> `handleHierarchicalIdLevelChange()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (51 total, 3 thin omitted)

### Community 0 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, bindHTMXTransformResponseInDescendants(), addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail (+8 more)

### Community 1 - "AssetSearchAutocompleteController"
Cohesion: 0.18
Nodes (7): AssetSearchAutocompleteController, bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (42): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper() (+34 more)

### Community 3 - "portfolio-history-management.ts"
Cohesion: 0.20
Nodes (15): handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR (+7 more)

### Community 4 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.20
Nodes (18): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX, getAllocationHierarchySize(), getHierarchicalIdAsString() (+10 more)

### Community 5 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.07
Nodes (43): bootstrap, Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, activateAllocationPlanEditor() (+35 more)

### Community 6 - "asset.ts"
Cohesion: 0.10
Nodes (30): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+22 more)

### Community 7 - "allocation-plan-management.ts"
Cohesion: 0.12
Nodes (16): htmx.org, Recursive Planned Allocation Rows, allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow() (+8 more)

### Community 8 - "binding-financial-input.ts"
Cohesion: 0.11
Nodes (29): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit() (+21 more)

### Community 10 - "chart-utils.ts"
Cohesion: 0.20
Nodes (12): chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, getDatasetSum(), LABEL_CALLBACKS (+4 more)

### Community 11 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 12 - "portfolio-chart.ts"
Cohesion: 0.16
Nodes (15): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+7 more)

### Community 13 - "chart-types.ts"
Cohesion: 0.17
Nodes (14): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, ChartInteraction, ChartInteractions, LocalChartOptions (+6 more)

### Community 14 - "allocation-plan-chart.ts"
Cohesion: 0.17
Nodes (12): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+4 more)

### Community 15 - "ChartDataSource"
Cohesion: 0.15
Nodes (3): ChartDataSource, ChartDataSourceVisitor, SingleChartDataSource

### Community 16 - "service/index.ts"
Cohesion: 0.22
Nodes (15): CompleteAllocationPlan, SerializableCompleteAllocationPlan, SerializablePortfolioCompleteAllocationPlanSet, Portfolio, PortfolioDTO, AllocationDomainService, mapAllocationStructure(), mapToAllocationPlan() (+7 more)

### Community 17 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 18 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 19 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 20 - "handlebars/index.ts"
Cohesion: 0.19
Nodes (13): handlebars, mapForAllocationPlanManagement(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY (+5 more)

### Community 21 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 22 - "chart-contents.ts"
Cohesion: 0.21
Nodes (10): bignumber.js, allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO (+2 more)

### Community 23 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 25 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 26 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 27 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 30 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (17): ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_INPUT_ATTRIBUTE, ASSET_SELECTION_ERROR_MESSAGE, ASSET_TICKER_INPUT_ATTRIBUTE, INVALID_INPUT_CLASS, TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE (+9 more)

### Community 31 - "asset-row-controller.ts"
Cohesion: 0.15
Nodes (15): ASSET_NOT_FOUND_ERROR_MESSAGE, FAILED_TO_FETCH_ASSET_ERROR_PREFIX, INVALID_TICKER_ERROR_MESSAGE, NEW_ASSET_TICKER_MESSAGE_SELECTOR, REQUIRED_FOR_SEARCH_ERROR_MESSAGE, TICKER_CANDIDATE_PATTERN, TICKER_TOO_LONG_ERROR_MESSAGE, ASSET_ROW_SELECTION_CHANGE_EVENT (+7 more)

### Community 32 - "Portfolio History Management"
Cohesion: 0.16
Nodes (13): Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization, Portfolio History Management (+5 more)

### Community 33 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 34 - "infra.ts"
Cohesion: 0.17
Nodes (11): chartjs-plugin-datalabels, Application, notifications, handlebarsInfra, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+3 more)

### Community 35 - "autocomplete-interactions.ts"
Cohesion: 0.26
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 36 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 37 - "notifications.ts"
Cohesion: 0.16
Nodes (10): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 38 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 39 - "allocation-plan.ts"
Cohesion: 0.24
Nodes (9): AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationPlan, AllocationPlanDTO, PlannedAllocation, PlannedAllocationDTO, SerializableFractalHierarchicalAllocationPlan (+1 more)

### Community 40 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.31
Nodes (10): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+2 more)

### Community 41 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 42 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 43 - "dom-utils.ts"
Cohesion: 0.24
Nodes (6): addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements

### Community 44 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 45 - "binding-htmx-transform-response.ts"
Cohesion: 0.31
Nodes (8): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse(), BeforeSwapEventDetail

### Community 46 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 47 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 48 - "chart.ts"
Cohesion: 0.32
Nodes (6): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), convertUnidimensionalDatasetBackgroundToPattern(), getPieDoughnutChartColorScale()

### Community 49 - "application/portfolio-analysis.ts"
Cohesion: 0.38
Nodes (5): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

### Community 50 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

## Knowledge Gaps
- **139 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+134 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 179 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `chart-contents.ts` to `handlebars-lang.ts`, `portfolio-history-management.ts`, `allocation-plan.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `logger`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.086) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `htmx/index.ts`, `portfolio-history-management.ts`, `portfolio-history-management-observation-metadata/index.ts`, `asset.ts`, `binding-htmx-trigger-on-route.ts`, `package.json`, `asset-composed-columns-input/index.ts`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `handlebars-lang.ts`, `portfolio-history-management.ts`, `notifications.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `chart-types.ts`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _139 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.08888888888888889 - nodes in this community are weakly interconnected._
- **Should `portfolio-history-management-observation-metadata/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07493061979648474 - nodes in this community are weakly interconnected._