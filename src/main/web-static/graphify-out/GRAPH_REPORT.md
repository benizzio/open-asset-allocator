# Graph Report - web-static  (2026-11-04)

## Corpus Check
- 104 files · ~46,886 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 892 nodes · 2011 edges · 48 communities (44 shown, 4 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 115 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `fc84e195`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset.ts
- MultiChartDataSource
- handlebars-lang.ts
- bignumber.js
- Allocation Plan Pagination
- pagination.ts
- asset-composed-columns-input/constants.ts
- binding-financial-input.ts
- allocation-plan-chart.ts
- package.json
- controls.ts
- Portfolio History Pagination
- AssetRowController
- AssetSearchAutocompleteController
- htmx/index.ts
- Portfolio History Management Pagination
- allocation-plan.ts
- portfolio-chart-datasource.ts
- asset-composed-columns-input/index.ts
- AssetSelectionState
- devDependencies
- routing/index.ts
- dependencies
- alias
- Portfolio Detail Page
- autocomplete-interactions.ts
- binding-htmx-trigger-on-route.ts
- portfolio-history-management-observation-details/index.ts
- compilerOptions
- chart-utils.ts
- getContextId
- focus-tracking.ts
- logger
- binding-htmx-transform-response.ts
- binding-dom-attribute-on-route.ts
- logging.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- .proxyrc.js
- Asset Search Autocomplete
- scripts
- allocation-plan-management.ts
- Frontend Module Architecture
- chart.ts
- Hierarchical Divergence Analysis
- Toast Notification
- portfolio-history-management.ts
- Allocation Plan Management Pagination

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
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history.html → websrc/components/paginated-list/context.ts
- `allocation-plan persisted record focus markers` --references--> `focusOnAcceptedRecord()`  [EXTRACTED]
  websrc/components/allocation-plan.html → websrc/components/paginated-list/controls.ts
- `allocation-plan portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan.html → websrc/components/paginated-list/context.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (48 total, 4 thin omitted)

### Community 0 - "asset.ts"
Cohesion: 0.06
Nodes (47): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, ExternalAsset, DomInfra, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+39 more)

### Community 1 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (47): handlebars, setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), andHelper(), arrayHelper() (+39 more)

### Community 3 - "bignumber.js"
Cohesion: 0.24
Nodes (8): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, DivergenceAnalysis, PotentialDivergence

### Community 4 - "Allocation Plan Pagination"
Cohesion: 0.18
Nodes (11): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Allocation Plan Wrapper-Owned Bottom Padding of Half the Bootstrap Horizontal Gutter, Independent of Hidden Pagination Controls, Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan persisted record focus markers (+3 more)

### Community 5 - "pagination.ts"
Cohesion: 0.18
Nodes (18): stopTrackingFocus(), bindPaginationLifecycle(), LifecycleHandlers, trackRequestCompletion(), bindRequestToXhr(), commitSettledPage(), finishRequest(), handleRequestCompletion() (+10 more)

### Community 6 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.13
Nodes (28): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE (+20 more)

### Community 7 - "binding-financial-input.ts"
Cohesion: 0.10
Nodes (31): BootstrapClasses, BootstrapIconClasses, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains() (+23 more)

### Community 8 - "allocation-plan-chart.ts"
Cohesion: 0.15
Nodes (21): allocationPlanChart, chartDataSelectionEventHandler(), getChartContent(), mapChildDatasets(), mapDataset(), toChartDataMap(), toChartContent(), toUnidimensionalMultiChartContent() (+13 more)

### Community 9 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 10 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 11 - "Portfolio History Pagination"
Cohesion: 0.22
Nodes (9): Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Pagination Retry and Loading Feedback, portfolio-history portfolio context selector, Portfolio History Pagination, Portfolio History Viewer, Portfolio History Wrapper-Owned Bottom Padding of Half the Bootstrap Horizontal Gutter, Independent of Hidden Pagination Controls, Route-Controlled Portfolio History and Management Views (+1 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.20
Nodes (6): AssetSearchAutocompleteController, bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+8 more)

### Community 15 - "Portfolio History Management Pagination"
Cohesion: 0.33
Nodes (6): Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management portfolio context selector, Portfolio History Management Pagination, Portfolio History Management Terminal State Without Completion Message

### Community 16 - "allocation-plan.ts"
Cohesion: 0.07
Nodes (47): FractalPlannedAllocationMultiChartDataSource, getSelectedDataKey(), interactionObserverCallback(), mapForAllocationPlanManagement(), mapPlanPage(), AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN (+39 more)

### Community 17 - "portfolio-chart-datasource.ts"
Cohesion: 0.19
Nodes (11): chart.js, FractalPortfolioMultiChartDataSource, generateDataKey(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+3 more)

### Community 18 - "asset-composed-columns-input/index.ts"
Cohesion: 0.12
Nodes (18): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, loadDatalists() (+10 more)

### Community 19 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 20 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 21 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 22 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 24 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 25 - "autocomplete-interactions.ts"
Cohesion: 0.27
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AssetSearchAutocompleteOption, AutocompleteControllerActions

### Community 26 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.15
Nodes (15): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants() (+7 more)

### Community 27 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.12
Nodes (25): areObservationDetailsReady(), bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, FAILED, LOADING, READY (+17 more)

### Community 28 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 29 - "chart-utils.ts"
Cohesion: 0.18
Nodes (14): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD (+6 more)

### Community 30 - "getContextId"
Cohesion: 0.32
Nodes (7): allocation-plan portfolio context selector, getContextId(), configureRequest(), getOwner(), getState(), requestNextPage(), updateControls()

### Community 31 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 33 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 34 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 35 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 36 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 37 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 38 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 39 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 40 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 41 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 42 - "allocation-plan-management.ts"
Cohesion: 0.07
Nodes (36): bootstrap, htmx.org, Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header (+28 more)

### Community 43 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 44 - "chart.ts"
Cohesion: 0.15
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, LocalChartOptions (+7 more)

### Community 47 - "portfolio-history-management.ts"
Cohesion: 0.05
Nodes (62): Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization, Portfolio History Management (+54 more)

### Community 51 - "Allocation Plan Management Pagination"
Cohesion: 0.20
Nodes (10): Allocation Plan Management Pagination, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message, portfolio-history-management persisted record focus markers (+2 more)

## Knowledge Gaps
- **172 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+167 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 214 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset.ts`, `pagination.ts`, `package.json`, `htmx/index.ts`, `portfolio-history-management.ts`, `asset-composed-columns-input/index.ts`, `binding-htmx-trigger-on-route.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.099) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `bignumber.js` to `logger`, `handlebars-lang.ts`, `package.json`, `allocation-plan-management.ts`, `portfolio-history-management.ts`, `allocation-plan.ts`, `chart-utils.ts`?**
  _High betweenness centrality (0.076) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `bignumber.js`, `package.json`, `allocation-plan-management.ts`, `chart.ts`, `portfolio-history-management.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _172 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06019871420222092 - nodes in this community are weakly interconnected._
- **Should `MultiChartDataSource` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._