# Graph Report - web-static  (2026-11-04)

## Corpus Check
- 104 files · ~47,069 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 895 nodes · 2017 edges · 53 communities (50 shown, 3 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 117 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `072527f9`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan.ts
- portfolio-history-management.ts
- allocation-plan-management.ts
- handlebars-lang.ts
- asset-composed-columns-input/constants.ts
- dom-utils.ts
- portfolio-chart.ts
- package.json
- asset-external.ts
- AssetRowController
- AssetSearchAutocompleteController
- chart-utils.ts
- htmx/index.ts
- pagination.ts
- asset.ts
- asset-composed-columns-input/index.ts
- MultiChartDataSource
- binding-financial-input.ts
- devDependencies
- routing/index.ts
- dependencies
- notifications.ts
- FractalPortfolioMultiChartDataSource
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- autocomplete-interactions.ts
- chart-types.ts
- compilerOptions
- Allocation Plan Pagination
- Portfolio History Pagination
- binding-htmx-trigger-on-route.ts
- Portfolio History Management Pagination
- controls.ts
- chart.ts
- dom/index.ts
- logging.ts
- logger
- binding-htmx-transform-response.ts
- binding-dom-attribute-on-route.ts
- infra.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- getContextId
- .proxyrc.js
- Allocation Plan Management Pagination
- Asset Search Autocomplete
- binding-accordion-field-focus.ts
- lifecycle.ts
- scripts
- AssetSelectionState
- Frontend Module Architecture
- Hierarchical Divergence Analysis
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
- `portfolio-history portfolio context selector` --conceptually_related_to--> `getContextId()`  [INFERRED]
  websrc/components/portfolio-history.html → websrc/components/paginated-list/context.ts
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/paginated-list/context.ts
- `allocation-plan-management persisted record focus markers` --references--> `focusOnAcceptedRecord()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/paginated-list/controls.ts
- `portfolio-history persisted record focus markers` --conceptually_related_to--> `focusOnAcceptedRecord()`  [INFERRED]
  websrc/components/portfolio-history.html → websrc/components/paginated-list/controls.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (53 total, 3 thin omitted)

### Community 0 - "allocation-plan.ts"
Cohesion: 0.06
Nodes (57): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+49 more)

### Community 1 - "portfolio-history-management.ts"
Cohesion: 0.05
Nodes (53): bignumber.js, Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization (+45 more)

### Community 2 - "allocation-plan-management.ts"
Cohesion: 0.06
Nodes (52): htmx.org, Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, Recursive Planned Allocation Rows (+44 more)

### Community 3 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (40): addPlannedAllocationRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper() (+32 more)

### Community 4 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+21 more)

### Community 5 - "dom-utils.ts"
Cohesion: 0.11
Nodes (19): handlebars, addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements, Format (+11 more)

### Community 6 - "portfolio-chart.ts"
Cohesion: 0.16
Nodes (18): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation (+10 more)

### Community 7 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 8 - "asset-external.ts"
Cohesion: 0.17
Nodes (21): addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields(), handleExternalSearchAfterRequest(), handleExternalSearchBeforeRequest(), initializeExternalDraft() (+13 more)

### Community 10 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 11 - "chart-utils.ts"
Cohesion: 0.15
Nodes (16): chart.js, chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS (+8 more)

### Community 12 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+8 more)

### Community 13 - "pagination.ts"
Cohesion: 0.20
Nodes (18): FocusContext, onFocusChange(), startTrackingFocus(), stopTrackingFocus(), tracked, bindRequestToXhr(), commitSettledPage(), finishRequest() (+10 more)

### Community 14 - "asset.ts"
Cohesion: 0.13
Nodes (11): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetPage, AssetRequestEvent, hasValidExternalData(), isExternalAsset(), getAssetIdentifierFromLocation() (+3 more)

### Community 15 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (15): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+7 more)

### Community 16 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 17 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 18 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 19 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 20 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 21 - "notifications.ts"
Cohesion: 0.16
Nodes (10): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 22 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.24
Nodes (5): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AppliedAllocationHierarchyLevel

### Community 23 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 24 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 25 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 26 - "chart-types.ts"
Cohesion: 0.22
Nodes (11): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, LocalChartOptions, MeasuramentUnit, CURRENCY (+3 more)

### Community 27 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 28 - "Allocation Plan Pagination"
Cohesion: 0.22
Nodes (11): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Allocation Plan Wrapper-Owned Bottom Padding of Half the Bootstrap Horizontal Gutter, Independent of Hidden Pagination Controls, Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination Retry and Loading Feedback (+3 more)

### Community 29 - "Portfolio History Pagination"
Cohesion: 0.20
Nodes (11): Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Pagination Retry and Loading Feedback, portfolio-history portfolio context selector, Portfolio History Pagination, Portfolio History Viewer, Portfolio History Wrapper-Owned Bottom Padding of Half the Bootstrap Horizontal Gutter, Independent of Hidden Pagination Controls, Route-Controlled Portfolio History and Management Views (+3 more)

### Community 30 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.31
Nodes (10): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+2 more)

### Community 31 - "Portfolio History Management Pagination"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management persisted record focus markers, portfolio-history-management portfolio context selector, Portfolio History Management Pagination, Portfolio History Management Terminal State Without Completion Message (+2 more)

### Community 32 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 33 - "chart.ts"
Cohesion: 0.24
Nodes (6): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), convertUnidimensionalDatasetBackgroundToPattern(), getPieDoughnutChartColorScale()

### Community 34 - "dom/index.ts"
Cohesion: 0.36
Nodes (7): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), maskNumberDecimalPlaces(), maskTagInput(), maskTickerInput()

### Community 35 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 36 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 37 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 38 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 39 - "infra.ts"
Cohesion: 0.28
Nodes (8): DomInfra, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), setupGlobalErrorHandler(), Router

### Community 40 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 41 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 42 - "getContextId"
Cohesion: 0.32
Nodes (7): allocation-plan portfolio context selector, getContextId(), configureRequest(), getOwner(), getState(), requestNextPage(), updateControls()

### Community 43 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 44 - "Allocation Plan Management Pagination"
Cohesion: 0.29
Nodes (7): Allocation Plan Management Pagination, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message

### Community 45 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 46 - "binding-accordion-field-focus.ts"
Cohesion: 0.52
Nodes (6): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains()

### Community 47 - "lifecycle.ts"
Cohesion: 0.33
Nodes (4): bindPaginationLifecycle(), LifecycleHandlers, trackRequestCompletion(), init()

### Community 48 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 49 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 50 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **171 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+166 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 213 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `allocation-plan.ts`, `portfolio-history-management.ts`, `package.json`, `htmx/index.ts`, `pagination.ts`, `asset.ts`, `asset-composed-columns-input/index.ts`, `binding-htmx-trigger-on-route.ts`?**
  _High betweenness centrality (0.099) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `portfolio-history-management.ts` to `allocation-plan.ts`, `allocation-plan-management.ts`, `handlebars-lang.ts`, `logger`, `dom-utils.ts`, `portfolio-chart.ts`, `package.json`, `chart-utils.ts`?**
  _High betweenness centrality (0.076) - this node is a cross-community bridge._
- **Why does `handlebars` connect `dom-utils.ts` to `portfolio-history-management.ts`, `allocation-plan-management.ts`, `handlebars-lang.ts`, `package.json`, `asset.ts`, `notifications.ts`, `chart-types.ts`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _171 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06481481481481481 - nodes in this community are weakly interconnected._
- **Should `portfolio-history-management.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05217391304347826 - nodes in this community are weakly interconnected._