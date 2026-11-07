# Graph Report - web-static  (2026-11-04)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 897 nodes · 2019 edges · 52 communities (48 shown, 4 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 117 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `68d315cf`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset.ts
- chart-utils.ts
- handlebars-lang.ts
- asset-composed-columns-input/constants.ts
- dom/index.ts
- htmx/index.ts
- allocation-plan.ts
- binding-htmx-trigger-on-route.ts
- portfolio-chart.ts
- management-metadata-editor/index.ts
- package.json
- pagination.ts
- application/index.ts
- portfolio-history-management.ts
- allocation-plan-chart.ts
- AssetRowController
- AssetSearchAutocompleteController
- allocation-plan-management.ts
- asset-composed-columns-input/index.ts
- binding-financial-input.ts
- devDependencies
- Portfolio History Management
- fractal-allocation-plan-mapping.ts
- dependencies
- FractalPortfolioMultiChartDataSource
- portfolio-history-management-observation-details/index.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- response.ts
- autocomplete-interactions.ts
- Portfolio History Pagination
- compilerOptions
- allocation.ts
- Allocation Plan Pagination
- focusOnAcceptedRecord
- controls.ts
- logger
- binding-htmx-wait-for-ready.ts
- binding-dom-attribute-on-route.ts
- binding-dom-display-on-route.ts
- .proxyrc.js
- Asset Search Autocomplete
- Allocation Plan Management Pagination
- portfolio-history-management-observation-metadata/index.ts
- lifecycle.ts
- scripts
- AssetSelectionState
- focus-tracking.ts
- Frontend Module Architecture
- alias
- Hierarchical Divergence Analysis
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 21 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `createPortfolioHistoryQuoteAction()` - 16 edges
5. `LogLevel` - 15 edges
6. `AssetSearchAutocompleteController` - 14 edges
7. `bignumber.js` - 14 edges
8. `handlebars` - 14 edges
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

## Communities (52 total, 4 thin omitted)

### Community 0 - "asset.ts"
Cohesion: 0.05
Nodes (51): api, APIError, Application, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, Asset, ExternalAsset, DomInfra (+43 more)

### Community 1 - "chart-utils.ts"
Cohesion: 0.05
Nodes (43): bignumber.js, chart.js, chartjs-plugin-datalabels, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence (+35 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (46): handlebars, setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), andHelper(), arrayHelper(), comparatorHelper() (+38 more)

### Community 3 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+21 more)

### Community 4 - "dom/index.ts"
Cohesion: 0.10
Nodes (24): bootstrap, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning() (+16 more)

### Community 5 - "htmx/index.ts"
Cohesion: 0.11
Nodes (26): APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP (+18 more)

### Community 6 - "allocation-plan.ts"
Cohesion: 0.17
Nodes (20): AllocationPlan, AllocationPlanDTO, CompleteAllocationPlan, FractalHierarchicalAllocationPlan, PlannedAllocationDTO, SerializableCompleteAllocationPlan, SerializableFractalHierarchicalAllocationPlan, SerializableFractalPlannedAllocation (+12 more)

### Community 7 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.17
Nodes (23): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement() (+15 more)

### Community 8 - "portfolio-chart.ts"
Cohesion: 0.15
Nodes (18): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation (+10 more)

### Community 9 - "management-metadata-editor/index.ts"
Cohesion: 0.13
Nodes (23): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, activateAllocationPlanEditor(), bindAllocationPlanMetadataSubmitGuard() (+15 more)

### Community 10 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 11 - "pagination.ts"
Cohesion: 0.19
Nodes (21): getContextId(), stopTrackingFocus(), bindRequestToXhr(), commitSettledPage(), configureRequest(), finishRequest(), getOwner(), getState() (+13 more)

### Community 12 - "application/index.ts"
Cohesion: 0.14
Nodes (16): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView(), allocationPlanManagement (+8 more)

### Community 13 - "portfolio-history-management.ts"
Cohesion: 0.18
Nodes (17): notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR (+9 more)

### Community 14 - "allocation-plan-chart.ts"
Cohesion: 0.16
Nodes (12): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+4 more)

### Community 16 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 17 - "allocation-plan-management.ts"
Cohesion: 0.14
Nodes (12): htmx.org, addPlannedAllocationRow(), AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow(), mapFormRowHierarchicalStructure() (+4 more)

### Community 18 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (15): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+7 more)

### Community 19 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 20 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 21 - "Portfolio History Management"
Cohesion: 0.16
Nodes (13): Dated Observation Header, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization (+5 more)

### Community 22 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.30
Nodes (13): PlannedAllocation, getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure() (+5 more)

### Community 23 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 24 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.24
Nodes (5): chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AppliedAllocationHierarchyLevel, AllocationHierarchyLevel

### Community 25 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.24
Nodes (13): bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, FAILED, LOADING, READY, UNLOADED (+5 more)

### Community 26 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 27 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 28 - "response.ts"
Cohesion: 0.28
Nodes (9): mapForAllocationPlanManagement(), mapPlanPage(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), websrc_pages_index_portfoliopage, formatCalendarDate(), getBrowserLocalCalendarDate() (+1 more)

### Community 29 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 30 - "Portfolio History Pagination"
Cohesion: 0.18
Nodes (12): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+4 more)

### Community 31 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 32 - "allocation.ts"
Cohesion: 0.25
Nodes (8): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, AllocationDomainService

### Community 33 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 34 - "focusOnAcceptedRecord"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, focusOnAcceptedRecord(), Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management persisted record focus markers, portfolio-history-management portfolio context selector (+2 more)

### Community 35 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 36 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 37 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 38 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 39 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 40 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 41 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 42 - "Allocation Plan Management Pagination"
Cohesion: 0.29
Nodes (7): Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message

### Community 43 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.52
Nodes (6): areObservationDetailsReady(), copyEditedTag(), getObservationTimestamp(), prepareObservationMetadataForSubmit(), setObservationTimestamp(), calendarDateToUtcTimestamp()

### Community 44 - "lifecycle.ts"
Cohesion: 0.33
Nodes (4): bindPaginationLifecycle(), LifecycleHandlers, trackRequestCompletion(), init()

### Community 45 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 46 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 47 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 48 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **171 isolated node(s):** `APIError`, `Draft`, `SearchEvent`, `AssetBeforeSwapEvent`, `AssetRequestEvent` (+166 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 213 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset.ts`, `htmx/index.ts`, `binding-htmx-trigger-on-route.ts`, `package.json`, `pagination.ts`, `portfolio-history-management.ts`, `asset-composed-columns-input/index.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.099) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `chart-utils.ts` to `handlebars-lang.ts`, `logger`, `allocation-plan.ts`, `portfolio-chart.ts`, `package.json`, `portfolio-history-management.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.077) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `chart-utils.ts`, `package.json`, `portfolio-history-management.ts`, `allocation-plan-management.ts`, `response.ts`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `APIError`, `Draft`, `SearchEvent` to the rest of the system?**
  _171 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05223880597014925 - nodes in this community are weakly interconnected._
- **Should `chart-utils.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05384615384615385 - nodes in this community are weakly interconnected._