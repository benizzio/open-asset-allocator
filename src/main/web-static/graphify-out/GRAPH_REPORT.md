# Graph Report - web-static  (2026-10-25)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 833 nodes · 1909 edges · 51 communities (47 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 101 edges (avg confidence: 0.84)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `4154ef0f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan-management.ts
- asset-composed-columns-input/constants.ts
- handlebars-lang.ts
- chart-utils.ts
- dom-utils.ts
- binding-financial-input.ts
- pagination.ts
- Portfolio History Management
- package.json
- portfolio-history-management.ts
- portfolio-chart-datasource.ts
- AssetRowController
- allocation-plan-chart.ts
- asset-external.ts
- htmx/index.ts
- routing/index.ts
- portfolio-chart.ts
- autocomplete-controller.ts
- MultiChartDataSource
- asset.ts
- devDependencies
- dependencies
- notifications.ts
- infra.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- AssetSearchAutocompleteController
- compilerOptions
- autocomplete-interactions.ts
- portfolio-history-management-observation-metadata/index.ts
- binding-htmx-transform-response.ts
- binding-htmx-trigger-on-route.ts
- portfolio-history-management-observation-details/index.ts
- dom/index.ts
- logging.ts
- logger
- allocation-plan-management-metadata/index.ts
- management-metadata-editor/index.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-attribute-on-route.ts
- Allocation Plan Management
- binding-dom-display-on-route.ts
- .proxyrc.js
- Asset Search Autocomplete
- binding-accordion-field-focus.ts
- scripts
- AssetSelectionState
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
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `Recursive Planned Allocation Rows` --references--> `handleHierarchicalIdLevelChange()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (51 total, 4 thin omitted)

### Community 0 - "allocation-plan-management.ts"
Cohesion: 0.06
Nodes (63): htmx.org, Recursive Planned Allocation Rows, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange() (+55 more)

### Community 1 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.10
Nodes (37): ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_NOT_FOUND_ERROR_MESSAGE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_INPUT_ATTRIBUTE (+29 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (39): getNextPortfolioHistoryManagementIndex(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper() (+31 more)

### Community 3 - "chart-utils.ts"
Cohesion: 0.11
Nodes (28): chart.js, chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions(), buildChartOptions() (+20 more)

### Community 4 - "dom-utils.ts"
Cohesion: 0.12
Nodes (18): handlebars, addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements, domJSONHelper() (+10 more)

### Community 5 - "binding-financial-input.ts"
Cohesion: 0.13
Nodes (23): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, bindFinancialInput(), bindFinancialInputElements() (+15 more)

### Community 6 - "pagination.ts"
Cohesion: 0.17
Nodes (24): AllocationPlanPageItem, bindRequestToXhr(), commitSettledPage(), configureRequest(), finishRequest(), getOwner(), getPortfolioId(), getState() (+16 more)

### Community 7 - "Portfolio History Management"
Cohesion: 0.10
Nodes (20): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor (+12 more)

### Community 8 - "package.json"
Cohesion: 0.11
Nodes (20): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+12 more)

### Community 9 - "portfolio-history-management.ts"
Cohesion: 0.18
Nodes (17): notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR (+9 more)

### Community 10 - "portfolio-chart-datasource.ts"
Cohesion: 0.19
Nodes (8): chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData

### Community 12 - "allocation-plan-chart.ts"
Cohesion: 0.18
Nodes (14): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+6 more)

### Community 13 - "asset-external.ts"
Cohesion: 0.18
Nodes (19): ExternalAsset, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields(), handleExternalSearchAfterRequest(), moveExternalAsset() (+11 more)

### Community 14 - "htmx/index.ts"
Cohesion: 0.17
Nodes (15): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+7 more)

### Community 15 - "routing/index.ts"
Cohesion: 0.24
Nodes (15): navigo, bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter() (+7 more)

### Community 16 - "portfolio-chart.ts"
Cohesion: 0.17
Nodes (14): toChartContent(), toUnidimensionalMultiChartContent(), getChartContent(), interactionObserverCallback(), portfolioChart, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshot (+6 more)

### Community 17 - "autocomplete-controller.ts"
Cohesion: 0.15
Nodes (14): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_STATUS_ATTRIBUTE, ASSET_SEARCH_SUGGESTIONS_ATTRIBUTE (+6 more)

### Community 18 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 19 - "asset.ts"
Cohesion: 0.14
Nodes (11): AssetBeforeSwapEvent, AssetRequestEvent, handleExternalSearchBeforeRequest(), hasValidExternalData(), initializeExternalDraft(), isExternalAsset(), getAssetIdentifierFromLocation(), getCurrentAssetIdentifierForResponse() (+3 more)

### Community 20 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 21 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 22 - "notifications.ts"
Cohesion: 0.16
Nodes (10): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 23 - "infra.ts"
Cohesion: 0.18
Nodes (10): chartjs-plugin-datalabels, Application, DomInfra, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError() (+2 more)

### Community 24 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 25 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 26 - "AssetSearchAutocompleteController"
Cohesion: 0.28
Nodes (3): AssetSearchAutocompleteController, loadDatalists(), AssetSearchAutocompleteState

### Community 27 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 28 - "autocomplete-interactions.ts"
Cohesion: 0.32
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_SEARCH_OPTION_SELECTOR, AutocompleteControllerActions

### Community 29 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.29
Nodes (11): isManagementMetadataEditorActive(), areObservationDetailsReady(), bindObservationMetadataRequest(), boundContainers, copyEditedTag(), getObservationForm(), getObservationTimestamp(), ObservationRequest (+3 more)

### Community 30 - "binding-htmx-transform-response.ts"
Cohesion: 0.25
Nodes (10): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+2 more)

### Community 31 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.31
Nodes (10): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+2 more)

### Community 32 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.40
Nodes (9): bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail, retryObservationDetails() (+1 more)

### Community 33 - "dom/index.ts"
Cohesion: 0.36
Nodes (7): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), maskNumberDecimalPlaces(), maskTagInput(), maskTickerInput()

### Community 34 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 35 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 36 - "allocation-plan-management-metadata/index.ts"
Cohesion: 0.31
Nodes (8): activateAllocationPlanEditor(), bindAllocationPlanMetadataRequest(), boundContainers, copyEditedPlanMetadata(), PlanRequest, prepareAllocationPlanMetadataForPost(), activateManagementMetadataEditor(), RequestConfigEventDetail

### Community 37 - "management-metadata-editor/index.ts"
Cohesion: 0.31
Nodes (8): bindReadOnlyManagementMetadataHeaderToggle(), boundHeaderToggleContainers, hideManagementMetadataEditor(), interactiveHeaderControlSelector, revealManagementMetadataEditor(), toggleReadOnlyManagementMetadataPanel(), updateManagementMetadataHeader(), propagateRefreshDataAfterPost()

### Community 38 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 39 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 40 - "Allocation Plan Management"
Cohesion: 0.36
Nodes (8): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan Management Pagination, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, handleInputAllocationPlanDate()

### Community 41 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 42 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 43 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 44 - "binding-accordion-field-focus.ts"
Cohesion: 0.52
Nodes (6): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains()

### Community 45 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 46 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 47 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **150 isolated node(s):** `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure`, `AllocationHierarchyLevelDTO`, `PlannedAllocationDTO`, `ReducedAllocation` (+145 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 191 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `binding-financial-input.ts` to `allocation-plan-management.ts`, `handlebars-lang.ts`, `chart-utils.ts`, `logger`, `package.json`, `portfolio-history-management.ts`, `portfolio-chart.ts`?**
  _High betweenness centrality (0.084) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `portfolio-history-management-observation-details/index.ts`, `asset-composed-columns-input/constants.ts`, `pagination.ts`, `package.json`, `portfolio-history-management.ts`, `htmx/index.ts`, `asset.ts`, `binding-htmx-trigger-on-route.ts`?**
  _High betweenness centrality (0.079) - this node is a cross-community bridge._
- **Why does `handlebars` connect `dom-utils.ts` to `allocation-plan-management.ts`, `handlebars-lang.ts`, `chart-utils.ts`, `binding-financial-input.ts`, `package.json`, `portfolio-history-management.ts`, `asset.ts`, `notifications.ts`?**
  _High betweenness centrality (0.056) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AllocationPlanningHierarchicalFormEntry`, `FormRowHierarchicalStructure`, `AllocationHierarchyLevelDTO` to the rest of the system?**
  _150 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan-management.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.055381400208986416 - nodes in this community are weakly interconnected._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09830866807610994 - nodes in this community are weakly interconnected._