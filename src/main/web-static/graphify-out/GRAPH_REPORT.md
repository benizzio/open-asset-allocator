# Graph Report - web-static  (2026-11-15)

## Corpus Check
- 123 files · ~55,050 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 1008 nodes · 2374 edges · 53 communities (49 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 127 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `762a622d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan-management-clone/index.ts
- asset.ts
- allocation-plan-management.ts
- service/index.ts
- AssetRowController
- handlebars-lang.ts
- asset-composed-columns-input/index.ts
- package.json
- pagination.ts
- Portfolio History Management
- dom/index.ts
- chart-utils.ts
- dom-utils.ts
- htmx/index.ts
- fractal-allocation-plan-mapping.ts
- allocation-plan.ts
- Allocation Plan Pagination
- binding-htmx-trigger-on-route.ts
- binding-financial-input.ts
- devDependencies
- allocation-plan-chart.ts
- routing/index.ts
- asset-composed-columns-input/constants.ts
- dependencies
- portfolio-chart-datasource.ts
- getContextId
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- portfolio-history-management-observation-details/index.ts
- Portfolio History Pagination
- compilerOptions
- application/index.ts
- FractalPortfolioMultiChartDataSource
- binding-htmx-transform-response.ts
- Portfolio History Management Pagination
- autocomplete-controller.ts
- controls.ts
- binding-dom-attribute-on-route.ts
- logger
- root.ts
- binding-dom-display-on-route.ts
- ObservationDetailsState
- Asset Search Autocomplete
- autocomplete-interactions.ts
- .proxyrc.js
- Allocation Plan Management Pagination
- portfolio-history-management.ts
- scripts
- focus-tracking.ts
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Toast Notification
- infra.ts

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 23 edges
3. `handlebars` - 18 edges
4. `registerHandlebarsLangHelpers()` - 18 edges
5. `createPortfolioHistoryQuoteAction()` - 16 edges
6. `LogLevel` - 15 edges
7. `bignumber.js` - 14 edges
8. `AssetSearchAutocompleteController` - 14 edges
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `portfolio-history portfolio context selector` --conceptually_related_to--> `getContextId()`  [INFERRED]
  websrc/components/portfolio-history/portfolio-history.html → websrc/components/paginated-list/context.ts
- `allocation-plan portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history/portfolio-history-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history persisted record focus markers` --conceptually_related_to--> `focusOnAcceptedRecord()`  [INFERRED]
  websrc/components/portfolio-history/portfolio-history.html → websrc/components/paginated-list/controls.ts
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/paginated-list/context.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, websrc_components_portfolio_history_portfolio_history_portfolio_history_viewer, websrc_components_allocation_plan_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (53 total, 4 thin omitted)

### Community 0 - "allocation-plan-management-clone/index.ts"
Cohesion: 0.08
Nodes (49): bootstrap, bindAllocationPlanCloning(), boundLists, copyDraft(), destinationControls(), presentDraft(), NEW_PLAN_DATE_SELECTOR, NEW_PLAN_NAME_SELECTOR (+41 more)

### Community 1 - "asset.ts"
Cohesion: 0.07
Nodes (39): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, Asset, ExternalAsset, DomInfra, Notification, NotificationType, ERROR (+31 more)

### Community 2 - "allocation-plan-management.ts"
Cohesion: 0.06
Nodes (49): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, AllocationPlanningHierarchicalFormEntry, bindNativePlanRowSubmitGuard(), Allocation Plan Clone Control (+41 more)

### Community 3 - "service/index.ts"
Cohesion: 0.17
Nodes (17): AllocationStructureDTO, AllocationPlan, AllocationPlanDTO, CompleteAllocationPlan, SerializableCompleteAllocationPlan, SerializablePortfolioCompleteAllocationPlanSet, AllocationDomainService, mapAllocationStructure() (+9 more)

### Community 5 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (53): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate() (+45 more)

### Community 6 - "asset-composed-columns-input/index.ts"
Cohesion: 0.12
Nodes (22): renderCommittedAssetSelection(), ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, AssetDraft, captureAssetDraft(), input(), restoreAssetDraft() (+14 more)

### Community 7 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 8 - "pagination.ts"
Cohesion: 0.17
Nodes (19): stopTrackingFocus(), bindPaginationLifecycle(), LifecycleHandlers, LOAD_MORE_EVENT, trackRequestCompletion(), bindRequestToXhr(), commitSettledPage(), finishRequest() (+11 more)

### Community 9 - "Portfolio History Management"
Cohesion: 0.14
Nodes (15): Icon-Only Observation Clone Control, Dated Observation Header, Disabled Observation Clone Tooltip Wrapper, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor (+7 more)

### Community 10 - "dom/index.ts"
Cohesion: 0.09
Nodes (37): capturePlanDraft(), captureRow(), HIERARCHY_FIELD_PATTERN, HierarchyValue, input(), PlanDraft, PlanRowDraft, restorePlanRows() (+29 more)

### Community 11 - "chart-utils.ts"
Cohesion: 0.06
Nodes (39): chart.js, chartjs-plugin-datalabels, chroma-js, patternomaly, chart, chartContentRepo, getChartContent(), getChartContentFromChart() (+31 more)

### Community 12 - "dom-utils.ts"
Cohesion: 0.18
Nodes (11): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), addRemoveObserver(), contextDataCache, DomUtils (+3 more)

### Community 13 - "htmx/index.ts"
Cohesion: 0.16
Nodes (17): APIError, APIErrorResponse, bindHTMXTransformResponseInDescendants(), bindHTMXWaitForReadyInDescendants(), addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail (+9 more)

### Community 14 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.30
Nodes (13): PlannedAllocation, getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure() (+5 more)

### Community 15 - "allocation-plan.ts"
Cohesion: 0.16
Nodes (12): FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, FractalHierarchicalAllocationPlan (+4 more)

### Community 16 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 17 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.22
Nodes (13): htmx.org, navigo, RequestConfigEventDetail, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement() (+5 more)

### Community 18 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 19 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 20 - "allocation-plan-chart.ts"
Cohesion: 0.19
Nodes (17): allocationPlanChart, chartDataSelectionEventHandler(), mapChildDatasets(), mapDataset(), toChartDataMap(), toChartContent(), toUnidimensionalMultiChartContent(), changeChartData() (+9 more)

### Community 21 - "routing/index.ts"
Cohesion: 0.30
Nodes (13): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+5 more)

### Community 22 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (27): api, CommittedAssetControls, CommittedAssetSelection, ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE (+19 more)

### Community 23 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 24 - "portfolio-chart-datasource.ts"
Cohesion: 0.16
Nodes (15): bignumber.js, getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData, AllocationHierarchyLevel, AllocationHierarchyLevelDTO (+7 more)

### Community 25 - "getContextId"
Cohesion: 0.32
Nodes (7): allocation-plan-management portfolio context selector, getContextId(), configureRequest(), getOwner(), getState(), requestNextPage(), updateControls()

### Community 26 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 27 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 28 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.32
Nodes (13): bindObservationDetailsLoader(), handleAfterRequest(), handleAfterSettle(), handleBeforeRequest(), handleBeforeSwap(), readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail (+5 more)

### Community 29 - "Portfolio History Pagination"
Cohesion: 0.17
Nodes (13): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+5 more)

### Community 30 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 31 - "application/index.ts"
Cohesion: 0.09
Nodes (24): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), HistoryPageItem, registerPortfolioHistoryPageTransform() (+16 more)

### Community 33 - "binding-htmx-transform-response.ts"
Cohesion: 0.27
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse(), TransformResponseFunction (+1 more)

### Community 34 - "Portfolio History Management Pagination"
Cohesion: 0.33
Nodes (6): Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management portfolio context selector, Portfolio History Management Terminal State Without Completion Message

### Community 35 - "autocomplete-controller.ts"
Cohesion: 0.11
Nodes (17): AssetSearchAutocompleteController, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_STATUS_ATTRIBUTE (+9 more)

### Community 36 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 37 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 38 - "logger"
Cohesion: 0.58
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions(), logger()

### Community 39 - "root.ts"
Cohesion: 0.25
Nodes (5): Application, notifications, AfterRequestEventDetail, HtmxInfra, Infra

### Community 40 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 41 - "ObservationDetailsState"
Cohesion: 0.40
Nodes (5): ObservationDetailsState, FAILED, LOADING, READY, UNLOADED

### Community 43 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 44 - "autocomplete-interactions.ts"
Cohesion: 0.31
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_SEARCH_OPTION_SELECTOR, AutocompleteControllerActions

### Community 46 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 47 - "Allocation Plan Management Pagination"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, Allocation Plan Management Terminal State Without Completion Message, focusOnAcceptedRecord() (+2 more)

### Community 48 - "portfolio-history-management.ts"
Cohesion: 0.20
Nodes (14): PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR (+6 more)

### Community 50 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 51 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 53 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 57 - "infra.ts"
Cohesion: 0.22
Nodes (9): bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), setupGlobalErrorHandler(), CustomEventHandler, Level, Levels (+1 more)

## Knowledge Gaps
- **186 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+181 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 226 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `binding-htmx-trigger-on-route.ts` to `allocation-plan-management-clone/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `asset-composed-columns-input/index.ts`, `package.json`, `pagination.ts`, `htmx/index.ts`, `portfolio-history-management.ts`, `allocation-plan-chart.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.110) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `portfolio-chart-datasource.ts` to `allocation-plan-management.ts`, `service/index.ts`, `handlebars-lang.ts`, `package.json`, `dom/index.ts`, `chart-utils.ts`, `allocation-plan.ts`, `portfolio-history-management.ts`, `application/index.ts`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `allocation-plan-management-clone/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `dom/index.ts`, `chart-utils.ts`, `portfolio-history-management.ts`, `application/index.ts`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _186 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan-management-clone/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07738095238095238 - nodes in this community are weakly interconnected._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07294117647058823 - nodes in this community are weakly interconnected._