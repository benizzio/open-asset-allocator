# Graph Report - web-static  (2026-11-13)

## Corpus Check
- 114 files · ~51,643 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 951 nodes · 2183 edges · 52 communities (49 shown, 3 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 118 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f61d4fed`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- portfolio-history-management-observation-clone/index.ts
- portfolio-history-management.ts
- handlebars-lang.ts
- asset-row-controller.ts
- asset.ts
- allocation-plan-chart.ts
- allocation-plan.ts
- binding-dom-attribute-on-route.ts
- package.json
- pagination.ts
- AssetSearchAutocompleteController
- AssetRowController
- htmx/index.ts
- portfolio-chart.ts
- dom/index.ts
- devDependencies
- chart-utils.ts
- application/index.ts
- allocation-plan-management.ts
- chart-types.ts
- asset-composed-columns-input/index.ts
- fractal-allocation-plan-mapping.ts
- dependencies
- portfolio-history-management-observation-details/index.ts
- infra.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- asset-composed-columns-input/constants.ts
- autocomplete-interactions.ts
- AllocationPlanType
- MultiChartDataSource
- dom-utils.ts
- handlebars/index.ts
- binding-dom-navigate-to.ts
- Portfolio History Management
- compilerOptions
- management-metadata-editor/index.ts
- binding-htmx-trigger-on-route.ts
- application/portfolio-analysis.ts
- chart-contents.ts
- binding-dom-exclusive-display.ts
- logger
- binding-htmx-transform-response.ts
- chart.ts
- binding-htmx-wait-for-ready.ts
- Asset Search Autocomplete
- binding-dom-display-on-route.ts
- .proxyrc.js
- scripts
- Frontend Module Architecture
- Hierarchical Divergence Analysis
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 23 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `handlebars` - 16 edges
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
- `portfolio-history persisted record focus markers` --conceptually_related_to--> `focusOnAcceptedRecord()`  [INFERRED]
  websrc/components/portfolio-history/portfolio-history.html → websrc/components/paginated-list/controls.ts
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history/portfolio-history-management.html → websrc/components/paginated-list/context.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, websrc_components_portfolio_history_portfolio_history_portfolio_history_viewer, websrc_components_allocation_plan_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (52 total, 3 thin omitted)

### Community 0 - "portfolio-history-management-observation-clone/index.ts"
Cohesion: 0.08
Nodes (45): captureAssetDraft(), input(), restoreAssetDraft(), isManagementMetadataEditorActive(), NEW_OBSERVATION_FORM_ID, focusDestination(), openPanel(), renderRows() (+37 more)

### Community 1 - "portfolio-history-management.ts"
Cohesion: 0.16
Nodes (17): AssetComposedColumnsInput, hideManagementMetadataEditor(), PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, propagateRefreshDataAfterPost() (+9 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (38): andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper(), ifNotEqualsHelper() (+30 more)

### Community 3 - "asset-row-controller.ts"
Cohesion: 0.10
Nodes (23): CommittedAssetControls, CommittedAssetSelection, renderCommittedAssetSelection(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, ASSET_NOT_FOUND_ERROR_MESSAGE (+15 more)

### Community 4 - "asset.ts"
Cohesion: 0.09
Nodes (32): htmx.org, Asset, ExternalAsset, handlebarsInfra, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch() (+24 more)

### Community 5 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 6 - "allocation-plan.ts"
Cohesion: 0.15
Nodes (23): bignumber.js, AllocationPlan, AllocationPlanDTO, CompleteAllocationPlan, PlannedAllocation, PlannedAllocationDTO, SerializableCompleteAllocationPlan, SerializableFractalHierarchicalAllocationPlan (+15 more)

### Community 7 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 8 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 9 - "pagination.ts"
Cohesion: 0.06
Nodes (54): allocation-plan persisted record focus markers, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message (+46 more)

### Community 10 - "AssetSearchAutocompleteController"
Cohesion: 0.19
Nodes (7): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 12 - "htmx/index.ts"
Cohesion: 0.14
Nodes (18): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+10 more)

### Community 13 - "portfolio-chart.ts"
Cohesion: 0.13
Nodes (18): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+10 more)

### Community 14 - "dom/index.ts"
Cohesion: 0.11
Nodes (29): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit() (+21 more)

### Community 15 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 16 - "chart-utils.ts"
Cohesion: 0.20
Nodes (12): chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, getDatasetSum(), LABEL_CALLBACKS (+4 more)

### Community 17 - "application/index.ts"
Cohesion: 0.12
Nodes (19): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), Application, HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView() (+11 more)

### Community 18 - "allocation-plan-management.ts"
Cohesion: 0.12
Nodes (16): addPlannedAllocationRow(), AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey() (+8 more)

### Community 19 - "chart-types.ts"
Cohesion: 0.17
Nodes (14): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, ChartInteraction, ChartInteractions, LocalChartOptions (+6 more)

### Community 20 - "asset-composed-columns-input/index.ts"
Cohesion: 0.14
Nodes (19): ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_SELECTION_ERROR_MESSAGE, INVALID_INPUT_CLASS, TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), guardedAssetForms, invalidateSelectedAsset() (+11 more)

### Community 21 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.20
Nodes (18): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX, getAllocationHierarchySize(), getHierarchicalIdAsString() (+10 more)

### Community 22 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 23 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.20
Nodes (18): bindObservationDetailsLoader(), handleAfterRequest(), handleAfterSettle(), handleBeforeRequest(), handleBeforeSwap(), ObservationDetailsState, FAILED, LOADING (+10 more)

### Community 24 - "infra.ts"
Cohesion: 0.11
Nodes (18): chartjs-plugin-datalabels, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError() (+10 more)

### Community 25 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 26 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 27 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 28 - "autocomplete-interactions.ts"
Cohesion: 0.27
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AssetSearchAutocompleteOption, AutocompleteControllerActions

### Community 29 - "AllocationPlanType"
Cohesion: 0.67
Nodes (3): AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN

### Community 30 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 31 - "dom-utils.ts"
Cohesion: 0.28
Nodes (5): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements

### Community 32 - "handlebars/index.ts"
Cohesion: 0.31
Nodes (7): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), PARTIALS_REGISTRY, registerPartialToContainer(), renderTemplateByElementId()

### Community 33 - "binding-dom-navigate-to.ts"
Cohesion: 0.57
Nodes (6): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate()

### Community 34 - "Portfolio History Management"
Cohesion: 0.06
Nodes (38): Allocation Map, Portfolio Section Navigation, Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector (+30 more)

### Community 35 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 36 - "management-metadata-editor/index.ts"
Cohesion: 0.13
Nodes (21): bootstrap, Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, activateAllocationPlanEditor() (+13 more)

### Community 37 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.18
Nodes (20): navigo, RequestConfigEventDetail, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS (+12 more)

### Community 38 - "application/portfolio-analysis.ts"
Cohesion: 0.38
Nodes (5): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

### Community 39 - "chart-contents.ts"
Cohesion: 0.40
Nodes (5): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartContent

### Community 42 - "binding-dom-exclusive-display.ts"
Cohesion: 0.60
Nodes (5): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings()

### Community 43 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 44 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

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

### Community 53 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 56 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **178 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+173 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 220 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `asset.ts` to `portfolio-history-management-observation-clone/index.ts`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`, `allocation-plan.ts`, `package.json`, `pagination.ts`, `htmx/index.ts`, `allocation-plan-management.ts`, `asset-composed-columns-input/index.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.105) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `allocation-plan.ts` to `portfolio-history-management.ts`, `handlebars-lang.ts`, `application/portfolio-analysis.ts`, `package.json`, `logger`, `portfolio-chart.ts`, `chart-utils.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.071) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars/index.ts` to `portfolio-history-management-observation-clone/index.ts`, `portfolio-history-management.ts`, `handlebars-lang.ts`, `asset.ts`, `application/portfolio-analysis.ts`, `package.json`, `allocation-plan-management.ts`, `chart-types.ts`, `infra.ts`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _178 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `portfolio-history-management-observation-clone/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.08315863032844165 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09878048780487805 - nodes in this community are weakly interconnected._