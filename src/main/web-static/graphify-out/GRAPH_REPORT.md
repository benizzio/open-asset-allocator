# Graph Report - web-static  (2026-11-04)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 893 nodes · 2014 edges · 45 communities (42 shown, 3 thin omitted)
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
- application/portfolio-analysis.ts
- Allocation Plan Pagination
- pagination.ts
- asset-composed-columns-input/constants.ts
- infra.ts
- allocation-plan-chart.ts
- package.json
- portfolio-history-management.ts
- management-metadata-editor/index.ts
- AssetRowController
- AssetSearchAutocompleteController
- htmx/index.ts
- binding-financial-input.ts
- allocation-plan.ts
- portfolio-chart.ts
- asset-composed-columns-input/index.ts
- AssetSelectionState
- devDependencies
- logging.ts
- dependencies
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- autocomplete-interactions.ts
- binding-htmx-trigger-on-route.ts
- portfolio-history-management-observation-details/index.ts
- compilerOptions
- notifications.ts
- Allocation Plan Management
- logger
- binding-htmx-transform-response.ts
- binding-dom-attribute-on-route.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- .proxyrc.js
- Asset Search Autocomplete
- scripts
- allocation-plan-management.ts
- Frontend Module Architecture
- chart-utils.ts
- Hierarchical Divergence Analysis
- Toast Notification
- application/index.ts

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 21 edges
3. `registerHandlebarsLangHelpers()` - 18 edges
4. `createPortfolioHistoryQuoteAction()` - 16 edges
5. `LogLevel` - 15 edges
6. `AssetSearchAutocompleteController` - 14 edges
7. `handlebars` - 14 edges
8. `bignumber.js` - 14 edges
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `allocation-plan portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan.html → websrc/components/paginated-list/context.ts
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `allocation-plan-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history-management portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/paginated-list/context.ts
- `portfolio-history portfolio context selector` --references--> `getContextId()`  [EXTRACTED]
  websrc/components/portfolio-history.html → websrc/components/paginated-list/context.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (45 total, 3 thin omitted)

### Community 0 - "asset.ts"
Cohesion: 0.09
Nodes (32): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetPage, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft (+24 more)

### Community 1 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (45): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+37 more)

### Community 3 - "application/portfolio-analysis.ts"
Cohesion: 0.31
Nodes (6): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, Format

### Community 4 - "Allocation Plan Pagination"
Cohesion: 0.08
Nodes (28): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Allocation Plan Wrapper-Owned Bottom Padding of Half the Bootstrap Horizontal Gutter, Independent of Hidden Pagination Controls, Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination Retry and Loading Feedback (+20 more)

### Community 5 - "pagination.ts"
Cohesion: 0.05
Nodes (60): Allocation Plan Management Pagination, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, allocation-plan-management portfolio context selector, Allocation Plan Management Terminal State Without Completion Message, allocation-plan persisted record focus markers (+52 more)

### Community 6 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+21 more)

### Community 7 - "infra.ts"
Cohesion: 0.07
Nodes (35): bootstrap, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning() (+27 more)

### Community 8 - "allocation-plan-chart.ts"
Cohesion: 0.13
Nodes (20): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+12 more)

### Community 9 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 10 - "portfolio-history-management.ts"
Cohesion: 0.20
Nodes (15): handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR (+7 more)

### Community 11 - "management-metadata-editor/index.ts"
Cohesion: 0.18
Nodes (15): bindAllocationPlanMetadataSubmitGuard(), boundContainers, copyEditedPlanMetadata(), handleInputAllocationPlanName(), prepareAllocationPlanMetadataForSubmit(), activateManagementMetadataEditor(), bindReadOnlyManagementMetadataHeaderToggle(), boundHeaderToggleContainers (+7 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+8 more)

### Community 15 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 16 - "allocation-plan.ts"
Cohesion: 0.08
Nodes (45): bignumber.js, mapForAllocationPlanManagement(), mapPlanPage(), AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN (+37 more)

### Community 17 - "portfolio-chart.ts"
Cohesion: 0.13
Nodes (19): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+11 more)

### Community 18 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (15): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+7 more)

### Community 19 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 20 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 21 - "logging.ts"
Cohesion: 0.31
Nodes (9): Level, Levels, LogLevel, bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath() (+1 more)

### Community 22 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 23 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 24 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 25 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 26 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.20
Nodes (18): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+10 more)

### Community 27 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.12
Nodes (25): areObservationDetailsReady(), bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, FAILED, LOADING, READY (+17 more)

### Community 28 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 29 - "notifications.ts"
Cohesion: 0.18
Nodes (9): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO, SUCCESS (+1 more)

### Community 30 - "Allocation Plan Management"
Cohesion: 0.36
Nodes (8): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, activateAllocationPlanEditor(), handleInputAllocationPlanDate()

### Community 33 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 34 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 35 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 37 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 38 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

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
Cohesion: 0.15
Nodes (11): htmx.org, Recursive Planned Allocation Rows, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow(), mapFormRowHierarchicalStructure() (+3 more)

### Community 43 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 44 - "chart-utils.ts"
Cohesion: 0.10
Nodes (29): chartjs-plugin-datalabels, chroma-js, patternomaly, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions() (+21 more)

### Community 47 - "application/index.ts"
Cohesion: 0.13
Nodes (18): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), Application, HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView() (+10 more)

## Knowledge Gaps
- **171 isolated node(s):** `BootstrapNotification`, `GeneralErrorHandler`, `AssetBeforeSwapEvent`, `AssetRequestEvent`, `Draft` (+166 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 213 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset.ts`, `pagination.ts`, `package.json`, `portfolio-history-management.ts`, `htmx/index.ts`, `asset-composed-columns-input/index.ts`, `binding-htmx-trigger-on-route.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.099) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `allocation-plan.ts` to `logger`, `handlebars-lang.ts`, `application/portfolio-analysis.ts`, `package.json`, `portfolio-history-management.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `portfolio-chart.ts`?**
  _High betweenness centrality (0.076) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `application/portfolio-analysis.ts`, `infra.ts`, `package.json`, `portfolio-history-management.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `portfolio-history-management-observation-details/index.ts`, `notifications.ts`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `BootstrapNotification`, `GeneralErrorHandler`, `AssetBeforeSwapEvent` to the rest of the system?**
  _171 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09487179487179487 - nodes in this community are weakly interconnected._
- **Should `MultiChartDataSource` be split into smaller, more focused modules?**
  _Cohesion score 0.125 - nodes in this community are weakly interconnected._