# Graph Report - web-static  (2026-11-08)

## Corpus Check
- 110 files · ~50,926 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 935 nodes · 2136 edges · 55 communities (51 shown, 4 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 118 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6e8ba89d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset.ts
- handlebars-lang.ts
- portfolio-history-management-observation-clone/index.ts
- portfolio-chart.ts
- asset-composed-columns-input/index.ts
- asset-composed-columns-input/constants.ts
- service/index.ts
- allocation-plan-chart.ts
- allocation-plan-management.ts
- fractal-allocation-plan-mapping.ts
- package.json
- pagination.ts
- AssetRowController
- AssetSearchAutocompleteController
- portfolio-history-management.ts
- htmx/index.ts
- asset-row-controller.ts
- chart-types.ts
- chart-utils.ts
- dom/index.ts
- handlebars/index.ts
- Portfolio History Management
- MultiChartDataSource
- devDependencies
- allocation-plan.ts
- routing/index.ts
- dependencies
- chart.ts
- allocation.ts
- Portfolio Detail Page
- autocomplete-interactions.ts
- chart-contents.ts
- alias
- Portfolio History Pagination
- compilerOptions
- binding-htmx-trigger-on-route.ts
- application/portfolio-analysis.ts
- Allocation Plan Pagination
- focusOnAcceptedRecord
- controls.ts
- logger
- binding-htmx-transform-response.ts
- binding-dom-attribute-on-route.ts
- dom-utils.ts
- binding-htmx-wait-for-ready.ts
- Asset Search Autocomplete
- .proxyrc.js
- Allocation Plan Management Pagination
- lifecycle.ts
- logging.ts
- scripts
- focus-tracking.ts
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 22 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `createPortfolioHistoryQuoteAction()` - 16 edges
5. `handlebars` - 15 edges
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

## Communities (55 total, 4 thin omitted)

### Community 0 - "asset.ts"
Cohesion: 0.10
Nodes (29): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+21 more)

### Community 1 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (42): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper() (+34 more)

### Community 2 - "portfolio-history-management-observation-clone/index.ts"
Cohesion: 0.07
Nodes (52): captureAssetDraft(), input(), restoreAssetDraft(), isManagementMetadataEditorActive(), AllocationDraft, captureKeys(), captureObservationDraft(), captureRow() (+44 more)

### Community 3 - "portfolio-chart.ts"
Cohesion: 0.14
Nodes (18): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+10 more)

### Community 4 - "asset-composed-columns-input/index.ts"
Cohesion: 0.13
Nodes (22): ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_INPUT_ATTRIBUTE, ASSET_SELECTION_ERROR_MESSAGE, ASSET_TICKER_INPUT_ATTRIBUTE, AssetSelectionState, EXISTING (+14 more)

### Community 5 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 6 - "service/index.ts"
Cohesion: 0.26
Nodes (13): CompleteAllocationPlan, Portfolio, PortfolioDTO, mapAllocationStructure(), mapToAllocationPlan(), mapToCompleteAllocationPlan(), mapToCompleteAllocationPlans(), mapToSerializableCompleteAllocationPlan() (+5 more)

### Community 7 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 8 - "allocation-plan-management.ts"
Cohesion: 0.05
Nodes (53): htmx.org, AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView() (+45 more)

### Community 9 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.33
Nodes (12): getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure(), connectFractalStructure() (+4 more)

### Community 10 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 11 - "pagination.ts"
Cohesion: 0.19
Nodes (21): getContextId(), stopTrackingFocus(), bindRequestToXhr(), commitSettledPage(), configureRequest(), finishRequest(), getOwner(), getState() (+13 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.19
Nodes (7): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "portfolio-history-management.ts"
Cohesion: 0.06
Nodes (49): Application, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction (+41 more)

### Community 15 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+8 more)

### Community 16 - "asset-row-controller.ts"
Cohesion: 0.13
Nodes (18): ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_NOT_FOUND_ERROR_MESSAGE, FAILED_TO_FETCH_ASSET_ERROR_PREFIX, INVALID_TICKER_ERROR_MESSAGE, NEW_ASSET_TICKER_MESSAGE_SELECTOR, REQUIRED_FOR_SEARCH_ERROR_MESSAGE, TICKER_CANDIDATE_PATTERN (+10 more)

### Community 17 - "chart-types.ts"
Cohesion: 0.17
Nodes (14): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, ChartInteraction, ChartInteractions, LocalChartOptions (+6 more)

### Community 18 - "chart-utils.ts"
Cohesion: 0.23
Nodes (11): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, getDatasetSum(), LABEL_CALLBACKS, valueAsPercentageOfDataset() (+3 more)

### Community 19 - "dom/index.ts"
Cohesion: 0.11
Nodes (30): bootstrap, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning() (+22 more)

### Community 20 - "handlebars/index.ts"
Cohesion: 0.30
Nodes (8): handlebars, handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY, registerPartialToContainer(), renderTemplateByElementId(), timestampToBrowserLocalCalendarDate()

### Community 21 - "Portfolio History Management"
Cohesion: 0.15
Nodes (14): Observation Clone Control, Dated Observation Header, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion (+6 more)

### Community 22 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 23 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 24 - "allocation-plan.ts"
Cohesion: 0.29
Nodes (8): bignumber.js, AllocationPlan, AllocationPlanDTO, PlannedAllocation, PlannedAllocationDTO, SerializableFractalHierarchicalAllocationPlan, SerializableFractalPlannedAllocation, SerializablePortfolioCompleteAllocationPlanSet

### Community 25 - "routing/index.ts"
Cohesion: 0.31
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), buildParameterizedDestinationPathFromCurrentLocationContext(), hasPathParam() (+4 more)

### Community 26 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 27 - "chart.ts"
Cohesion: 0.24
Nodes (7): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), ChartContent, convertUnidimensionalDatasetBackgroundToPattern(), getPieDoughnutChartColorScale()

### Community 28 - "allocation.ts"
Cohesion: 0.18
Nodes (10): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+2 more)

### Community 29 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 30 - "autocomplete-interactions.ts"
Cohesion: 0.27
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AssetSearchAutocompleteOption, AutocompleteControllerActions

### Community 31 - "chart-contents.ts"
Cohesion: 0.50
Nodes (4): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart

### Community 33 - "Portfolio History Pagination"
Cohesion: 0.18
Nodes (12): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+4 more)

### Community 34 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 35 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.26
Nodes (12): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+4 more)

### Community 36 - "application/portfolio-analysis.ts"
Cohesion: 0.38
Nodes (5): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

### Community 37 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 38 - "focusOnAcceptedRecord"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, focusOnAcceptedRecord(), Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management persisted record focus markers, portfolio-history-management portfolio context selector (+2 more)

### Community 39 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 40 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 41 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 42 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 43 - "dom-utils.ts"
Cohesion: 0.21
Nodes (8): addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements, domJSONHelper(), registerHandlebarsDOMHelpers()

### Community 44 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 45 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 47 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 48 - "Allocation Plan Management Pagination"
Cohesion: 0.29
Nodes (7): Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message

### Community 49 - "lifecycle.ts"
Cohesion: 0.33
Nodes (4): bindPaginationLifecycle(), LifecycleHandlers, trackRequestCompletion(), init()

### Community 50 - "logging.ts"
Cohesion: 0.20
Nodes (15): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel (+7 more)

### Community 51 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 53 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 54 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **175 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+170 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 217 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset.ts`, `portfolio-history-management-observation-clone/index.ts`, `binding-htmx-trigger-on-route.ts`, `asset-composed-columns-input/index.ts`, `package.json`, `pagination.ts`, `portfolio-history-management.ts`, `htmx/index.ts`?**
  _High betweenness centrality (0.104) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `allocation-plan.ts` to `handlebars-lang.ts`, `portfolio-chart.ts`, `application/portfolio-analysis.ts`, `allocation-plan-management.ts`, `logger`, `package.json`, `portfolio-history-management.ts`, `chart-utils.ts`?**
  _High betweenness centrality (0.072) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `asset.ts`, `handlebars-lang.ts`, `portfolio-history-management-observation-clone/index.ts`, `application/portfolio-analysis.ts`, `allocation-plan-management.ts`, `package.json`, `dom-utils.ts`, `portfolio-history-management.ts`, `chart-types.ts`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _175 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.10476190476190476 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.08888888888888889 - nodes in this community are weakly interconnected._