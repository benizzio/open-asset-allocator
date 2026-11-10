# Graph Report - web-static  (2026-11-10)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 935 nodes · 2136 edges · 55 communities (51 shown, 4 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 118 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `580e6e46`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan-management.ts
- portfolio-history-management.ts
- portfolio-history-management-observation-clone/index.ts
- handlebars-lang.ts
- asset.ts
- dom/index.ts
- portfolio-chart.ts
- asset-composed-columns-input/index.ts
- package.json
- pagination.ts
- asset-row-controller.ts
- htmx/index.ts
- AssetRowController
- AssetSearchAutocompleteController
- service/index.ts
- logging.ts
- allocation-plan-chart.ts
- Portfolio History Management
- MultiChartDataSource
- devDependencies
- chart-utils.ts
- asset-composed-columns-input/constants.ts
- chart-types.ts
- allocation.ts
- autocomplete-interactions.ts
- routing/index.ts
- dependencies
- fractal-allocation-plan-mapping.ts
- Portfolio Detail Page
- dom-utils.ts
- binding-htmx-trigger-on-route.ts
- handlebars/index.ts
- Portfolio History Pagination
- compilerOptions
- allocation-plan.ts
- Allocation Plan Pagination
- focusOnAcceptedRecord
- controls.ts
- chart.ts
- logger
- binding-htmx-transform-response.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-attribute-on-route.ts
- Asset Search Autocomplete
- .proxyrc.js
- application/portfolio-analysis.ts
- Allocation Plan Management Pagination
- lifecycle.ts
- scripts
- chart-contents.ts
- focus-tracking.ts
- Frontend Module Architecture
- alias
- Hierarchical Divergence Analysis
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 22 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `createPortfolioHistoryQuoteAction()` - 16 edges
5. `LogLevel` - 15 edges
6. `handlebars` - 15 edges
7. `AssetSearchAutocompleteController` - 14 edges
8. `bignumber.js` - 14 edges
9. `AutocompleteControllerActions` - 13 edges
10. `FractalPortfolioMultiChartDataSource` - 13 edges

## Surprising Connections (you probably didn't know these)
- `allocation-plan-management persisted record focus markers` --references--> `focusOnAcceptedRecord()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/paginated-list/controls.ts
- `portfolio-history persisted record focus markers` --conceptually_related_to--> `focusOnAcceptedRecord()`  [INFERRED]
  websrc/components/portfolio-history/portfolio-history.html → websrc/components/paginated-list/controls.ts
- `allocation-plan portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan.html → websrc/components/paginated-list/context.ts
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history/portfolio-history-management.html → websrc/components/paginated-list/context.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, websrc_components_portfolio_history_portfolio_history_portfolio_history_viewer, websrc_components_allocation_plan_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (55 total, 4 thin omitted)

### Community 0 - "allocation-plan-management.ts"
Cohesion: 0.05
Nodes (53): htmx.org, AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView() (+45 more)

### Community 1 - "portfolio-history-management.ts"
Cohesion: 0.06
Nodes (49): Application, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction (+41 more)

### Community 2 - "portfolio-history-management-observation-clone/index.ts"
Cohesion: 0.07
Nodes (52): captureAssetDraft(), input(), restoreAssetDraft(), isManagementMetadataEditorActive(), AllocationDraft, captureKeys(), captureObservationDraft(), captureRow() (+44 more)

### Community 3 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (42): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper() (+34 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (29): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+21 more)

### Community 5 - "dom/index.ts"
Cohesion: 0.11
Nodes (30): bootstrap, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning() (+22 more)

### Community 6 - "portfolio-chart.ts"
Cohesion: 0.14
Nodes (18): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+10 more)

### Community 7 - "asset-composed-columns-input/index.ts"
Cohesion: 0.13
Nodes (22): ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_INPUT_ATTRIBUTE, ASSET_SELECTION_ERROR_MESSAGE, ASSET_TICKER_INPUT_ATTRIBUTE, AssetSelectionState, EXISTING (+14 more)

### Community 8 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 9 - "pagination.ts"
Cohesion: 0.19
Nodes (21): getContextId(), stopTrackingFocus(), bindRequestToXhr(), commitSettledPage(), configureRequest(), finishRequest(), getOwner(), getState() (+13 more)

### Community 10 - "asset-row-controller.ts"
Cohesion: 0.13
Nodes (18): ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_NOT_FOUND_ERROR_MESSAGE, FAILED_TO_FETCH_ASSET_ERROR_PREFIX, INVALID_TICKER_ERROR_MESSAGE, NEW_ASSET_TICKER_MESSAGE_SELECTOR, REQUIRED_FOR_SEARCH_ERROR_MESSAGE, TICKER_CANDIDATE_PATTERN (+10 more)

### Community 11 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+8 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.19
Nodes (7): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "service/index.ts"
Cohesion: 0.26
Nodes (13): CompleteAllocationPlan, Portfolio, PortfolioDTO, mapAllocationStructure(), mapToAllocationPlan(), mapToCompleteAllocationPlan(), mapToCompleteAllocationPlans(), mapToSerializableCompleteAllocationPlan() (+5 more)

### Community 15 - "logging.ts"
Cohesion: 0.20
Nodes (15): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel (+7 more)

### Community 16 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 17 - "Portfolio History Management"
Cohesion: 0.15
Nodes (14): Icon-Only Observation Clone Control, Dated Observation Header, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion (+6 more)

### Community 18 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 19 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 20 - "chart-utils.ts"
Cohesion: 0.23
Nodes (11): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, getDatasetSum(), LABEL_CALLBACKS, valueAsPercentageOfDataset() (+3 more)

### Community 21 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 22 - "chart-types.ts"
Cohesion: 0.17
Nodes (14): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, ChartInteraction, ChartInteractions, LocalChartOptions (+6 more)

### Community 23 - "allocation.ts"
Cohesion: 0.18
Nodes (10): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+2 more)

### Community 24 - "autocomplete-interactions.ts"
Cohesion: 0.27
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AssetSearchAutocompleteOption, AutocompleteControllerActions

### Community 25 - "routing/index.ts"
Cohesion: 0.31
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), buildParameterizedDestinationPathFromCurrentLocationContext(), hasPathParam() (+4 more)

### Community 26 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 27 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.33
Nodes (12): getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure(), connectFractalStructure() (+4 more)

### Community 28 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 29 - "dom-utils.ts"
Cohesion: 0.21
Nodes (8): addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements, domJSONHelper(), registerHandlebarsDOMHelpers()

### Community 30 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.26
Nodes (12): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+4 more)

### Community 31 - "handlebars/index.ts"
Cohesion: 0.30
Nodes (8): handlebars, handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY, registerPartialToContainer(), renderTemplateByElementId(), timestampToBrowserLocalCalendarDate()

### Community 32 - "Portfolio History Pagination"
Cohesion: 0.18
Nodes (12): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+4 more)

### Community 33 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 34 - "allocation-plan.ts"
Cohesion: 0.29
Nodes (8): bignumber.js, AllocationPlan, AllocationPlanDTO, PlannedAllocation, PlannedAllocationDTO, SerializableFractalHierarchicalAllocationPlan, SerializableFractalPlannedAllocation, SerializablePortfolioCompleteAllocationPlanSet

### Community 35 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 36 - "focusOnAcceptedRecord"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, focusOnAcceptedRecord(), Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management persisted record focus markers, portfolio-history-management portfolio context selector (+2 more)

### Community 37 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 38 - "chart.ts"
Cohesion: 0.24
Nodes (7): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), ChartContent, convertUnidimensionalDatasetBackgroundToPattern(), getPieDoughnutChartColorScale()

### Community 39 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 40 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 41 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 42 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 43 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 44 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 45 - "application/portfolio-analysis.ts"
Cohesion: 0.38
Nodes (5): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

### Community 46 - "Allocation Plan Management Pagination"
Cohesion: 0.29
Nodes (7): Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message

### Community 47 - "lifecycle.ts"
Cohesion: 0.33
Nodes (4): bindPaginationLifecycle(), LifecycleHandlers, trackRequestCompletion(), init()

### Community 48 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 49 - "chart-contents.ts"
Cohesion: 0.50
Nodes (4): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart

### Community 50 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 51 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **175 isolated node(s):** `AllocationPlanPageItem`, `AllocationPlanViewPage`, `HistoryPageItem`, `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure` (+170 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 217 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `portfolio-history-management.ts`, `portfolio-history-management-observation-clone/index.ts`, `asset.ts`, `asset-composed-columns-input/index.ts`, `package.json`, `pagination.ts`, `htmx/index.ts`, `binding-htmx-trigger-on-route.ts`?**
  _High betweenness centrality (0.104) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `allocation-plan.ts` to `allocation-plan-management.ts`, `portfolio-history-management.ts`, `handlebars-lang.ts`, `portfolio-chart.ts`, `logger`, `package.json`, `application/portfolio-analysis.ts`, `chart-utils.ts`?**
  _High betweenness centrality (0.072) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `allocation-plan-management.ts`, `portfolio-history-management.ts`, `portfolio-history-management-observation-clone/index.ts`, `handlebars-lang.ts`, `asset.ts`, `package.json`, `application/portfolio-analysis.ts`, `chart-types.ts`, `dom-utils.ts`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AllocationPlanPageItem`, `AllocationPlanViewPage`, `HistoryPageItem` to the rest of the system?**
  _175 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan-management.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.0505175983436853 - nodes in this community are weakly interconnected._
- **Should `portfolio-history-management.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06240084611316764 - nodes in this community are weakly interconnected._