# Graph Report - web-static  (2026-11-13)

## Corpus Check
- 113 files · ~51,408 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 946 nodes · 2169 edges · 59 communities (56 shown, 3 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 118 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ac708370`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- portfolio-history-management-observation-clone/index.ts
- portfolio-history-management.ts
- handlebars-lang.ts
- asset-composed-columns-input/constants.ts
- asset.ts
- allocation-plan-chart.ts
- allocation-plan.ts
- routing/index.ts
- package.json
- pagination.ts
- AssetSearchAutocompleteController
- AssetRowController
- htmx/index.ts
- portfolio-chart.ts
- binding-financial-input.ts
- devDependencies
- chart-utils.ts
- allocation-plan-pagination.ts
- allocation-plan-management.ts
- chart-types.ts
- asset-composed-columns-input/index.ts
- fractal-allocation-plan-mapping.ts
- dependencies
- application/index.ts
- notifications.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- FractalPortfolioMultiChartDataSource
- autocomplete-interactions.ts
- allocation.ts
- ChartDataSource
- dom-utils.ts
- handlebars/index.ts
- response.ts
- Portfolio History Pagination
- compilerOptions
- Allocation Plan Management
- binding-htmx-trigger-on-route.ts
- bignumber.js
- Allocation Plan Pagination
- focusOnAcceptedRecord
- controls.ts
- logging.ts
- logger
- binding-htmx-transform-response.ts
- infra.ts
- chart.ts
- binding-htmx-wait-for-ready.ts
- Asset Search Autocomplete
- binding-dom-display-on-route.ts
- .proxyrc.js
- Allocation Plan Management Pagination
- lifecycle.ts
- scripts
- AssetSelectionState
- focus-tracking.ts
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 22 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `handlebars` - 16 edges
5. `createPortfolioHistoryQuoteAction()` - 16 edges
6. `LogLevel` - 15 edges
7. `bignumber.js` - 14 edges
8. `AssetSearchAutocompleteController` - 14 edges
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `allocation-plan portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan.html → websrc/components/paginated-list/context.ts
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history/portfolio-history-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history portfolio context selector` --conceptually_related_to--> `getContextId()`  [INFERRED]
  websrc/components/portfolio-history/portfolio-history.html → websrc/components/paginated-list/context.ts
- `allocation-plan-management persisted record focus markers` --references--> `focusOnAcceptedRecord()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/paginated-list/controls.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, websrc_components_portfolio_history_portfolio_history_portfolio_history_viewer, websrc_components_allocation_plan_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (59 total, 3 thin omitted)

### Community 0 - "portfolio-history-management-observation-clone/index.ts"
Cohesion: 0.06
Nodes (56): bootstrap, captureAssetDraft(), input(), restoreAssetDraft(), isManagementMetadataEditorActive(), NEW_OBSERVATION_FORM_ID, PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, focusDestination() (+48 more)

### Community 1 - "portfolio-history-management.ts"
Cohesion: 0.05
Nodes (56): activateManagementMetadataEditor(), bindReadOnlyManagementMetadataHeaderToggle(), boundHeaderToggleContainers, hideManagementMetadataEditor(), interactiveHeaderControlSelector, revealManagementMetadataEditor(), toggleReadOnlyManagementMetadataPanel(), updateManagementMetadataHeader() (+48 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (41): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper() (+33 more)

### Community 3 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.10
Nodes (34): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE (+26 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (31): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+23 more)

### Community 5 - "allocation-plan-chart.ts"
Cohesion: 0.12
Nodes (18): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+10 more)

### Community 6 - "allocation-plan.ts"
Cohesion: 0.17
Nodes (19): AllocationPlan, AllocationPlanDTO, CompleteAllocationPlan, FractalHierarchicalAllocationPlan, PlannedAllocationDTO, SerializableFractalHierarchicalAllocationPlan, SerializableFractalPlannedAllocation, SerializablePortfolioCompleteAllocationPlanSet (+11 more)

### Community 7 - "routing/index.ts"
Cohesion: 0.18
Nodes (22): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+14 more)

### Community 8 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 9 - "pagination.ts"
Cohesion: 0.19
Nodes (21): getContextId(), stopTrackingFocus(), bindRequestToXhr(), commitSettledPage(), configureRequest(), finishRequest(), getOwner(), getState() (+13 more)

### Community 10 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 12 - "htmx/index.ts"
Cohesion: 0.17
Nodes (15): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+7 more)

### Community 13 - "portfolio-chart.ts"
Cohesion: 0.25
Nodes (11): getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData, PortfolioAllocation (+3 more)

### Community 14 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 15 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 16 - "chart-utils.ts"
Cohesion: 0.20
Nodes (12): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, getDatasetSum() (+4 more)

### Community 17 - "allocation-plan-pagination.ts"
Cohesion: 0.18
Nodes (13): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView(), paginatedList (+5 more)

### Community 18 - "allocation-plan-management.ts"
Cohesion: 0.18
Nodes (10): AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), bindAllocationPlanMetadataSubmitGuard(), boundContainers, copyEditedPlanMetadata() (+2 more)

### Community 19 - "chart-types.ts"
Cohesion: 0.17
Nodes (14): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, ChartInteraction, ChartInteractions, LocalChartOptions (+6 more)

### Community 20 - "asset-composed-columns-input/index.ts"
Cohesion: 0.18
Nodes (13): UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker(), ASSET_ROW_SELECTION_CHANGE_EVENT, AssetRowSelectionChangeDetail (+5 more)

### Community 21 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.30
Nodes (13): PlannedAllocation, getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure() (+5 more)

### Community 22 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 23 - "application/index.ts"
Cohesion: 0.16
Nodes (8): Application, allocationPlanManagement, AssetComposedColumnsInput, notifications, AfterRequestEventDetail, Infra, AssetPage, websrc_pages_index_assetpage

### Community 24 - "notifications.ts"
Cohesion: 0.16
Nodes (10): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 25 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 26 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 27 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.26
Nodes (4): chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AllocationHierarchyLevel

### Community 28 - "autocomplete-interactions.ts"
Cohesion: 0.32
Nodes (5): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AutocompleteControllerActions

### Community 29 - "allocation.ts"
Cohesion: 0.23
Nodes (9): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+1 more)

### Community 30 - "ChartDataSource"
Cohesion: 0.15
Nodes (3): ChartDataSource, ChartDataSourceVisitor, SingleChartDataSource

### Community 31 - "dom-utils.ts"
Cohesion: 0.21
Nodes (8): addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements, domJSONHelper(), registerHandlebarsDOMHelpers()

### Community 32 - "handlebars/index.ts"
Cohesion: 0.30
Nodes (8): handlebars, handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY, registerPartialToContainer(), renderTemplateByElementId(), timestampToBrowserLocalCalendarDate()

### Community 33 - "response.ts"
Cohesion: 0.23
Nodes (8): htmx.org, mapForAllocationPlanManagement(), mapPlanPage(), SerializableCompleteAllocationPlan, DomainService, Router, websrc_pages_index_portfoliopage, PortfolioPage

### Community 34 - "Portfolio History Pagination"
Cohesion: 0.18
Nodes (12): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+4 more)

### Community 35 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 36 - "Allocation Plan Management"
Cohesion: 0.24
Nodes (11): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow() (+3 more)

### Community 37 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.31
Nodes (10): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+2 more)

### Community 38 - "bignumber.js"
Cohesion: 0.29
Nodes (7): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, Format

### Community 39 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 40 - "focusOnAcceptedRecord"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, focusOnAcceptedRecord(), Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management persisted record focus markers, portfolio-history-management portfolio context selector (+2 more)

### Community 41 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 42 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 43 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 44 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 45 - "infra.ts"
Cohesion: 0.28
Nodes (8): chart.js, chart, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), setupGlobalErrorHandler()

### Community 46 - "chart.ts"
Cohesion: 0.28
Nodes (6): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), convertUnidimensionalDatasetBackgroundToPattern(), getPieDoughnutChartColorScale()

### Community 47 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 48 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 49 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 50 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 51 - "Allocation Plan Management Pagination"
Cohesion: 0.29
Nodes (7): Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message

### Community 52 - "lifecycle.ts"
Cohesion: 0.33
Nodes (4): bindPaginationLifecycle(), LifecycleHandlers, trackRequestCompletion(), init()

### Community 53 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 54 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 55 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 56 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **177 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+172 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 219 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `response.ts` to `portfolio-history-management-observation-clone/index.ts`, `portfolio-history-management.ts`, `asset.ts`, `binding-htmx-trigger-on-route.ts`, `package.json`, `pagination.ts`, `htmx/index.ts`, `allocation-plan-management.ts`, `asset-composed-columns-input/index.ts`?**
  _High betweenness centrality (0.105) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `bignumber.js` to `portfolio-history-management.ts`, `handlebars-lang.ts`, `allocation-plan.ts`, `package.json`, `logger`, `portfolio-chart.ts`, `chart-utils.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.071) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `portfolio-history-management-observation-clone/index.ts`, `portfolio-history-management.ts`, `handlebars-lang.ts`, `asset.ts`, `bignumber.js`, `package.json`, `allocation-plan-management.ts`, `chart-types.ts`, `notifications.ts`, `dom-utils.ts`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _177 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `portfolio-history-management-observation-clone/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06453028972783142 - nodes in this community are weakly interconnected._
- **Should `portfolio-history-management.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05454545454545454 - nodes in this community are weakly interconnected._