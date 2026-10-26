# Graph Report - web-static  (2026-10-27)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 835 nodes · 1914 edges · 46 communities (42 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 101 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e0e005ff`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan.ts
- handlebars-lang.ts
- asset.ts
- binding-financial-input.ts
- htmx/index.ts
- portfolio-chart.ts
- pagination.ts
- package.json
- Portfolio History Management
- application/index.ts
- asset-composed-columns-input/index.ts
- asset-row-controller.ts
- AssetRowController
- AssetSearchAutocompleteController
- portfolio-history-management.ts
- chart.ts
- binding-htmx-trigger-on-route.ts
- allocation-plan-management.ts
- management-metadata-editor/index.ts
- chart-utils.ts
- MultiChartDataSource
- devDependencies
- asset-composed-columns-input/constants.ts
- routing/index.ts
- dependencies
- createPortfolioHistoryQuoteAction
- bignumber.js
- Portfolio Detail Page
- autocomplete-interactions.ts
- compilerOptions
- portfolio-history-management-observation-metadata/index.ts
- infra.ts
- portfolio-history-management-observation-details/index.ts
- logger
- binding-dom-attribute-on-route.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- Allocation Plan Management
- .proxyrc.js
- Asset Search Autocomplete
- binding-dom-exclusive-display.ts
- scripts
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
7. `handlebars` - 14 edges
8. `bignumber.js` - 14 edges
9. `AutocompleteControllerActions` - 13 edges
10. `FractalPortfolioMultiChartDataSource` - 13 edges

## Surprising Connections (you probably didn't know these)
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (46 total, 4 thin omitted)

### Community 0 - "allocation-plan.ts"
Cohesion: 0.06
Nodes (60): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+52 more)

### Community 1 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (52): handlebars, getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), andHelper() (+44 more)

### Community 2 - "asset.ts"
Cohesion: 0.10
Nodes (31): ExternalAsset, Router, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+23 more)

### Community 3 - "binding-financial-input.ts"
Cohesion: 0.11
Nodes (30): bootstrap, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning() (+22 more)

### Community 4 - "htmx/index.ts"
Cohesion: 0.11
Nodes (26): APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP (+18 more)

### Community 5 - "portfolio-chart.ts"
Cohesion: 0.14
Nodes (16): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData() (+8 more)

### Community 6 - "pagination.ts"
Cohesion: 0.14
Nodes (27): htmx.org, AllocationPlanPageItem, bindRequestToXhr(), commitSettledPage(), configureRequest(), finishRequest(), getOwner(), getPortfolioId() (+19 more)

### Community 7 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 8 - "Portfolio History Management"
Cohesion: 0.10
Nodes (20): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor (+12 more)

### Community 9 - "application/index.ts"
Cohesion: 0.11
Nodes (15): Application, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, paginatedList, Infra, CustomEventHandler, Notification (+7 more)

### Community 10 - "asset-composed-columns-input/index.ts"
Cohesion: 0.12
Nodes (21): ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_INPUT_ATTRIBUTE, ASSET_SELECTION_ERROR_MESSAGE, ASSET_TICKER_INPUT_ATTRIBUTE, AssetSelectionState, EXISTING (+13 more)

### Community 11 - "asset-row-controller.ts"
Cohesion: 0.13
Nodes (18): api, APIError, ASSET_NOT_FOUND_ERROR_MESSAGE, FAILED_TO_FETCH_ASSET_ERROR_PREFIX, INVALID_TICKER_ERROR_MESSAGE, NEW_ASSET_TICKER_MESSAGE_SELECTOR, REQUIRED_FOR_SEARCH_ERROR_MESSAGE, TICKER_CANDIDATE_PATTERN (+10 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "portfolio-history-management.ts"
Cohesion: 0.19
Nodes (16): handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR (+8 more)

### Community 15 - "chart.ts"
Cohesion: 0.15
Nodes (16): chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D (+8 more)

### Community 16 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.15
Nodes (15): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants() (+7 more)

### Community 17 - "allocation-plan-management.ts"
Cohesion: 0.12
Nodes (14): Recursive Planned Allocation Rows, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow() (+6 more)

### Community 18 - "management-metadata-editor/index.ts"
Cohesion: 0.18
Nodes (16): activateAllocationPlanEditor(), bindAllocationPlanMetadataRequest(), boundContainers, copyEditedPlanMetadata(), PlanRequest, prepareAllocationPlanMetadataForPost(), activateManagementMetadataEditor(), bindReadOnlyManagementMetadataHeaderToggle() (+8 more)

### Community 19 - "chart-utils.ts"
Cohesion: 0.16
Nodes (14): buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD, convertUnidimensionalDatasetBackgroundToPattern() (+6 more)

### Community 20 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 21 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 22 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 23 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 24 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 25 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 26 - "bignumber.js"
Cohesion: 0.21
Nodes (9): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioSnapshotDTO, DivergenceAnalysis, PotentialDivergence, mapToPortfolioSnapshot() (+1 more)

### Community 27 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 28 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 29 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 30 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.27
Nodes (11): areObservationDetailsReady(), bindObservationMetadataRequest(), boundContainers, copyEditedTag(), getObservationForm(), getObservationTimestamp(), ObservationRequest, prepareObservationMetadataForPost() (+3 more)

### Community 31 - "infra.ts"
Cohesion: 0.24
Nodes (8): chart.js, chartjs-plugin-datalabels, chart, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), setupGlobalErrorHandler()

### Community 32 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.40
Nodes (9): bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail, retryObservationDetails() (+1 more)

### Community 33 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 34 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 35 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 36 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 37 - "Allocation Plan Management"
Cohesion: 0.36
Nodes (8): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan Management Pagination, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, handleInputAllocationPlanDate()

### Community 38 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 39 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 40 - "binding-dom-exclusive-display.ts"
Cohesion: 0.60
Nodes (5): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings()

### Community 41 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 42 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **151 isolated node(s):** `AllocationHierarchyLevelDTO`, `PlannedAllocationDTO`, `AssignAtPathOptions`, `ToIntOptions`, `APIError` (+146 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 192 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `bignumber.js` to `allocation-plan.ts`, `logger`, `handlebars-lang.ts`, `portfolio-chart.ts`, `package.json`, `portfolio-history-management.ts`, `allocation-plan-management.ts`, `chart-utils.ts`?**
  _High betweenness centrality (0.084) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `pagination.ts` to `portfolio-history-management-observation-details/index.ts`, `allocation-plan.ts`, `asset.ts`, `htmx/index.ts`, `package.json`, `asset-composed-columns-input/index.ts`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.078) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `package.json`, `application/index.ts`, `portfolio-history-management.ts`, `chart.ts`, `allocation-plan-management.ts`, `bignumber.js`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AllocationHierarchyLevelDTO`, `PlannedAllocationDTO`, `AssignAtPathOptions` to the rest of the system?**
  _151 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06190476190476191 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06775956284153005 - nodes in this community are weakly interconnected._