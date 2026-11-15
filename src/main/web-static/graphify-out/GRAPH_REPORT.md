# Graph Report - web-static  (2026-11-15)

## Corpus Check
- 122 files · ~54,549 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 989 nodes · 2315 edges · 47 communities (43 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 127 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e8d4a211`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- dom/index.ts
- asset.ts
- allocation-plan-management.ts
- allocation-plan.ts
- AssetRowController
- handlebars-lang.ts
- asset-composed-columns-input/index.ts
- package.json
- pagination.ts
- portfolio-history-management-observation-details/index.ts
- bignumber.js
- chart-utils.ts
- chart.ts
- htmx/index.ts
- asset-composed-columns-input/constants.ts
- allocation-plan-chart.ts
- Allocation Plan Pagination
- MultiChartDataSource
- binding-financial-input.ts
- devDependencies
- chart-contents.ts
- logger
- asset-row-controller.ts
- dependencies
- portfolio-chart.ts
- getContextId
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- Portfolio History Pagination
- compilerOptions
- application/index.ts
- Portfolio History Management Pagination
- AssetSearchAutocompleteController
- controls.ts
- notifications.ts
- Asset Search Autocomplete
- autocomplete-interactions.ts
- .proxyrc.js
- Allocation Plan Management Pagination
- portfolio-history-management.ts
- scripts
- focus-tracking.ts
- Frontend Module Architecture
- alias
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

## Communities (47 total, 4 thin omitted)

### Community 0 - "dom/index.ts"
Cohesion: 0.05
Nodes (71): bootstrap, capturePlanDraft(), captureRow(), HierarchyValue, input(), PlanDraft, PlanRowDraft, restorePlanRows() (+63 more)

### Community 1 - "asset.ts"
Cohesion: 0.10
Nodes (30): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+22 more)

### Community 2 - "allocation-plan-management.ts"
Cohesion: 0.08
Nodes (39): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, AllocationPlanningHierarchicalFormEntry, bindNativePlanRowSubmitGuard(), Allocation Plan Clone Control (+31 more)

### Community 3 - "allocation-plan.ts"
Cohesion: 0.08
Nodes (49): htmx.org, mapForAllocationPlanManagement(), mapPlanPage(), AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN (+41 more)

### Community 5 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (53): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate() (+45 more)

### Community 6 - "asset-composed-columns-input/index.ts"
Cohesion: 0.14
Nodes (19): ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_SELECTION_ERROR_MESSAGE, INVALID_INPUT_CLASS, TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), guardedAssetForms, invalidateSelectedAsset() (+11 more)

### Community 7 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 8 - "pagination.ts"
Cohesion: 0.17
Nodes (19): stopTrackingFocus(), bindPaginationLifecycle(), LifecycleHandlers, LOAD_MORE_EVENT, trackRequestCompletion(), bindRequestToXhr(), commitSettledPage(), finishRequest() (+11 more)

### Community 9 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.08
Nodes (33): Icon-Only Observation Clone Control, Dated Observation Header, Disabled Observation Clone Tooltip Wrapper, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, bindObservationDetailsLoader() (+25 more)

### Community 10 - "bignumber.js"
Cohesion: 0.19
Nodes (11): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO, DivergenceAnalysis (+3 more)

### Community 11 - "chart-utils.ts"
Cohesion: 0.16
Nodes (15): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD (+7 more)

### Community 12 - "chart.ts"
Cohesion: 0.17
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+7 more)

### Community 13 - "htmx/index.ts"
Cohesion: 0.22
Nodes (12): addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs(), prepareFormData(), replaceRequestPathParams(), replaceRequestPathParamsFromCurrentRoute() (+4 more)

### Community 14 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 15 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (10): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+2 more)

### Community 16 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 17 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 18 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 19 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 20 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 21 - "logger"
Cohesion: 0.05
Nodes (79): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants() (+71 more)

### Community 22 - "asset-row-controller.ts"
Cohesion: 0.10
Nodes (23): CommittedAssetControls, CommittedAssetSelection, renderCommittedAssetSelection(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, ASSET_NOT_FOUND_ERROR_MESSAGE (+15 more)

### Community 23 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 24 - "portfolio-chart.ts"
Cohesion: 0.16
Nodes (14): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+6 more)

### Community 25 - "getContextId"
Cohesion: 0.32
Nodes (7): allocation-plan-management portfolio context selector, getContextId(), configureRequest(), getOwner(), getState(), requestNextPage(), updateControls()

### Community 26 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 27 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 29 - "Portfolio History Pagination"
Cohesion: 0.17
Nodes (13): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+5 more)

### Community 30 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 31 - "application/index.ts"
Cohesion: 0.13
Nodes (18): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView(), allocationPlanManagement (+10 more)

### Community 34 - "Portfolio History Management Pagination"
Cohesion: 0.33
Nodes (6): Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management portfolio context selector, Portfolio History Management Terminal State Without Completion Message

### Community 35 - "AssetSearchAutocompleteController"
Cohesion: 0.19
Nodes (7): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 36 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 37 - "notifications.ts"
Cohesion: 0.15
Nodes (11): api, APIError, APIErrorResponse, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, Notification, NotificationType, ERROR (+3 more)

### Community 43 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 44 - "autocomplete-interactions.ts"
Cohesion: 0.27
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AssetSearchAutocompleteOption, AutocompleteControllerActions

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
Cohesion: 0.15
Nodes (12): Application, notifications, DomInfra, handlebarsInfra, AfterRequestEventDetail, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+4 more)

## Knowledge Gaps
- **182 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+177 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 222 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan.ts` to `dom/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `asset-composed-columns-input/index.ts`, `package.json`, `pagination.ts`, `portfolio-history-management-observation-details/index.ts`, `htmx/index.ts`, `portfolio-history-management.ts`, `logger`?**
  _High betweenness centrality (0.106) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `bignumber.js` to `allocation-plan-management.ts`, `allocation-plan.ts`, `handlebars-lang.ts`, `package.json`, `chart-utils.ts`, `portfolio-history-management.ts`, `logger`?**
  _High betweenness centrality (0.067) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `dom/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `notifications.ts`, `package.json`, `bignumber.js`, `chart.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _182 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `dom/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05303392259913999 - nodes in this community are weakly interconnected._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.10241820768136557 - nodes in this community are weakly interconnected._