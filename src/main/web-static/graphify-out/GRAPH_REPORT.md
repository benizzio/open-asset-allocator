# Graph Report - web-static  (2026-10-27)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 869 nodes · 1962 edges · 47 communities (44 shown, 3 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 120 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f8822e63`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset.ts
- allocation-plan.ts
- handlebars-lang.ts
- allocation-plan-management.ts
- Portfolio History Management
- pagination.ts
- asset-composed-columns-input/constants.ts
- binding-financial-input.ts
- chart-utils.ts
- package.json
- portfolio-chart-datasource.ts
- portfolio-history-management.ts
- AssetRowController
- AssetSearchAutocompleteController
- htmx/index.ts
- chart.ts
- management-metadata-editor/index.ts
- portfolio-chart.ts
- asset-composed-columns-input/index.ts
- MultiChartDataSource
- devDependencies
- routing/index.ts
- dependencies
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- autocomplete-interactions.ts
- binding-htmx-trigger-on-route.ts
- Allocation Plan Management
- compilerOptions
- portfolio-history-management-observation-metadata/index.ts
- allocation-plan-chart.ts
- portfolio-history-management-observation-details/index.ts
- logging.ts
- logger
- binding-htmx-transform-response.ts
- binding-dom-attribute-on-route.ts
- dom-utils.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- .proxyrc.js
- Asset Search Autocomplete
- scripts
- AssetSelectionState
- Frontend Module Architecture
- ChartDataType
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
9. `AutocompleteControllerActions` - 13 edges
10. `FractalPortfolioMultiChartDataSource` - 13 edges

## Surprising Connections (you probably didn't know these)
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `Allocation Plan Pencil Metadata Editor` --references--> `handleInputAllocationPlanDate()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (47 total, 3 thin omitted)

### Community 0 - "asset.ts"
Cohesion: 0.05
Nodes (49): Application, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, Asset, ExternalAsset, DomInfra, handlebarsInfra, bootRouterDebouncing() (+41 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.08
Nodes (44): bignumber.js, FractalPlannedAllocationMultiChartDataSource, getSelectedDataKey(), AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN (+36 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (45): handlebars, getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+37 more)

### Community 3 - "allocation-plan-management.ts"
Cohesion: 0.06
Nodes (37): htmx.org, Recursive Planned Allocation Rows, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure (+29 more)

### Community 4 - "Portfolio History Management"
Cohesion: 0.06
Nodes (36): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Allocation Plan Wrapper-Owned Bottom Padding of Half the Bootstrap Horizontal Gutter, Independent of Hidden Pagination Controls, Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Pagination Retry and Loading Feedback, Route-Controlled Allocation Plan and Management Views (+28 more)

### Community 5 - "pagination.ts"
Cohesion: 0.11
Nodes (36): findOwnedElement(), focusOnAcceptedRecord(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton() (+28 more)

### Community 6 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+21 more)

### Community 7 - "binding-financial-input.ts"
Cohesion: 0.11
Nodes (29): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit() (+21 more)

### Community 8 - "chart-utils.ts"
Cohesion: 0.13
Nodes (18): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+10 more)

### Community 9 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 10 - "portfolio-chart-datasource.ts"
Cohesion: 0.19
Nodes (10): chart.js, FractalPortfolioMultiChartDataSource, generateDataKey(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+2 more)

### Community 11 - "portfolio-history-management.ts"
Cohesion: 0.18
Nodes (17): notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR (+9 more)

### Community 13 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 14 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, bindHTMXTransformResponseInDescendants(), addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail (+8 more)

### Community 15 - "chart.ts"
Cohesion: 0.17
Nodes (14): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, LocalChartOptions (+6 more)

### Community 16 - "management-metadata-editor/index.ts"
Cohesion: 0.15
Nodes (17): bootstrap, activateAllocationPlanEditor(), bindAllocationPlanMetadataRequest(), boundContainers, copyEditedPlanMetadata(), PlanRequest, prepareAllocationPlanMetadataForPost(), activateManagementMetadataEditor() (+9 more)

### Community 17 - "portfolio-chart.ts"
Cohesion: 0.21
Nodes (12): toChartContent(), toUnidimensionalMultiChartContent(), changeChartData(), chartDataSelectionEventHandler(), getChartContent(), interactionObserverCallback(), portfolioChart, PortfolioAllocation (+4 more)

### Community 18 - "asset-composed-columns-input/index.ts"
Cohesion: 0.15
Nodes (15): ASSET_ID_INPUT_SELECTOR, ASSET_NAME_INPUT_SELECTOR, UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+7 more)

### Community 19 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 20 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 21 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

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
Cohesion: 0.27
Nodes (11): navigo, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute() (+3 more)

### Community 27 - "Allocation Plan Management"
Cohesion: 0.21
Nodes (12): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan Management Pagination, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Allocation Plan Management Consistent List-to-Control Gap, Dated Allocation Plan Header (+4 more)

### Community 28 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 29 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.29
Nodes (11): isManagementMetadataEditorActive(), areObservationDetailsReady(), bindObservationMetadataRequest(), boundContainers, copyEditedTag(), getObservationForm(), getObservationTimestamp(), ObservationRequest (+3 more)

### Community 30 - "allocation-plan-chart.ts"
Cohesion: 0.33
Nodes (9): allocationPlanChart, chartDataSelectionEventHandler(), getChartContent(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap(), DomainService (+1 more)

### Community 31 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.40
Nodes (9): bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail, retryObservationDetails() (+1 more)

### Community 32 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 33 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 34 - "binding-htmx-transform-response.ts"
Cohesion: 0.27
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse(), TransformResponseFunction (+1 more)

### Community 35 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 36 - "dom-utils.ts"
Cohesion: 0.28
Nodes (5): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements

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

### Community 42 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 43 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 44 - "ChartDataType"
Cohesion: 0.67
Nodes (3): ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

## Knowledge Gaps
- **169 isolated node(s):** `PlanRequest`, `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure`, `PlannedAllocationDTO`, `AllocationHierarchyLevelDTO` (+164 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 210 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset.ts`, `pagination.ts`, `package.json`, `portfolio-history-management.ts`, `htmx/index.ts`, `asset-composed-columns-input/index.ts`, `binding-htmx-trigger-on-route.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.082) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `allocation-plan.ts` to `logger`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `portfolio-history-management.ts`, `portfolio-chart.ts`?**
  _High betweenness centrality (0.080) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `allocation-plan-management.ts`, `package.json`, `portfolio-history-management.ts`, `chart.ts`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `PlanRequest`, `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure` to the rest of the system?**
  _169 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05480769230769231 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07667900581702802 - nodes in this community are weakly interconnected._