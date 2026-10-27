# Graph Report - web-static  (2026-10-27)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 867 nodes · 1960 edges · 52 communities (48 shown, 4 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 120 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `05793769`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset-composed-columns-input/constants.ts
- handlebars-lang.ts
- allocation-plan-management.ts
- pagination.ts
- portfolio-chart.ts
- Allocation Plan Management
- allocation-plan.ts
- htmx/index.ts
- allocation-plan-chart.ts
- package.json
- infra.ts
- AssetRowController
- asset.ts
- portfolio-history-management.ts
- asset-external.ts
- chart-utils.ts
- MultiChartDataSource
- binding-financial-input.ts
- devDependencies
- Portfolio History Management
- logging.ts
- allocation.ts
- dependencies
- createPortfolioHistoryQuoteAction
- fractal-allocation-plan-mapping.ts
- binding-htmx-trigger-on-route.ts
- Portfolio Detail Page
- chart-types.ts
- routing/index.ts
- compilerOptions
- portfolio-history-management-observation-metadata/index.ts
- bignumber.js
- management-metadata-editor/index.ts
- portfolio-history-management-observation-details/index.ts
- dom/index.ts
- logger
- dom-utils.ts
- chart.ts
- binding-htmx-wait-for-ready.ts
- allocation-plan-management-metadata/index.ts
- binding-dom-display-on-route.ts
- .proxyrc.js
- Asset Search Autocomplete
- binding-accordion-field-focus.ts
- binding-dom-exclusive-display.ts
- scripts
- Portfolio History Management Pagination
- NotificationType
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
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `Recursive Planned Allocation Rows` --references--> `handleHierarchicalIdLevelChange()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts
- `Recursive Planned Allocation Rows` --references--> `handleRemovePlannedAllocationRow()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (52 total, 4 thin omitted)

### Community 0 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.05
Nodes (63): AssetSearchAutocompleteController, handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE (+55 more)

### Community 1 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (45): handlebars, getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+37 more)

### Community 2 - "allocation-plan-management.ts"
Cohesion: 0.07
Nodes (32): htmx.org, Recursive Planned Allocation Rows, Application, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation() (+24 more)

### Community 3 - "pagination.ts"
Cohesion: 0.11
Nodes (36): findOwnedElement(), focusOnAcceptedRecord(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton() (+28 more)

### Community 4 - "portfolio-chart.ts"
Cohesion: 0.15
Nodes (17): chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation (+9 more)

### Community 5 - "Allocation Plan Management"
Cohesion: 0.08
Nodes (29): Allocation Map, Allocation Plan Pagination, Allocation Plan Viewer, Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management (+21 more)

### Community 6 - "allocation-plan.ts"
Cohesion: 0.18
Nodes (19): AllocationPlan, AllocationPlanDTO, CompleteAllocationPlan, PlannedAllocation, PlannedAllocationDTO, SerializableFractalHierarchicalAllocationPlan, SerializableFractalPlannedAllocation, SerializablePortfolioCompleteAllocationPlanSet (+11 more)

### Community 7 - "htmx/index.ts"
Cohesion: 0.12
Nodes (23): APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP (+15 more)

### Community 8 - "allocation-plan-chart.ts"
Cohesion: 0.15
Nodes (16): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+8 more)

### Community 9 - "package.json"
Cohesion: 0.11
Nodes (20): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+12 more)

### Community 10 - "infra.ts"
Cohesion: 0.13
Nodes (14): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, DomInfra, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER() (+6 more)

### Community 12 - "asset.ts"
Cohesion: 0.11
Nodes (13): api, APIError, Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, handleExternalSearchBeforeRequest(), hasValidExternalData() (+5 more)

### Community 13 - "portfolio-history-management.ts"
Cohesion: 0.19
Nodes (16): handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR (+8 more)

### Community 14 - "asset-external.ts"
Cohesion: 0.18
Nodes (20): addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields(), handleExternalSearchAfterRequest(), initializeExternalDraft(), moveExternalAsset() (+12 more)

### Community 15 - "chart-utils.ts"
Cohesion: 0.16
Nodes (15): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD (+7 more)

### Community 16 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 17 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 18 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 19 - "Portfolio History Management"
Cohesion: 0.16
Nodes (13): Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization, Portfolio History Management (+5 more)

### Community 20 - "logging.ts"
Cohesion: 0.24
Nodes (13): Level, Levels, LogLevel, addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants() (+5 more)

### Community 21 - "allocation.ts"
Cohesion: 0.16
Nodes (10): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+2 more)

### Community 22 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 23 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 24 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.33
Nodes (12): getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure(), connectFractalStructure() (+4 more)

### Community 25 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.26
Nodes (12): navigo, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute() (+4 more)

### Community 26 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 27 - "chart-types.ts"
Cohesion: 0.22
Nodes (11): CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D, LocalChartOptions, MeasuramentUnit, CURRENCY (+3 more)

### Community 28 - "routing/index.ts"
Cohesion: 0.31
Nodes (11): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), buildParameterizedDestinationPathFromCurrentLocationContext(), hasPathParam() (+3 more)

### Community 29 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 30 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.31
Nodes (10): areObservationDetailsReady(), bindObservationMetadataRequest(), boundContainers, copyEditedTag(), getObservationForm(), getObservationTimestamp(), ObservationRequest, prepareObservationMetadataForPost() (+2 more)

### Community 31 - "bignumber.js"
Cohesion: 0.29
Nodes (7): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, Format

### Community 32 - "management-metadata-editor/index.ts"
Cohesion: 0.29
Nodes (9): activateManagementMetadataEditor(), bindReadOnlyManagementMetadataHeaderToggle(), boundHeaderToggleContainers, hideManagementMetadataEditor(), interactiveHeaderControlSelector, revealManagementMetadataEditor(), toggleReadOnlyManagementMetadataPanel(), updateManagementMetadataHeader() (+1 more)

### Community 33 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.40
Nodes (9): bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail, retryObservationDetails() (+1 more)

### Community 34 - "dom/index.ts"
Cohesion: 0.36
Nodes (7): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), maskNumberDecimalPlaces(), maskTagInput(), maskTickerInput()

### Community 35 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 36 - "dom-utils.ts"
Cohesion: 0.24
Nodes (6): addRemoveObserver(), contextDataCache, DomUtils, ensureSharedObserver(), getCacheableContextData(), observedElements

### Community 37 - "chart.ts"
Cohesion: 0.28
Nodes (6): chart.js, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), ChartContent

### Community 38 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 39 - "allocation-plan-management-metadata/index.ts"
Cohesion: 0.36
Nodes (7): bindAllocationPlanMetadataRequest(), boundContainers, copyEditedPlanMetadata(), PlanRequest, prepareAllocationPlanMetadataForPost(), isManagementMetadataEditorActive(), RequestConfigEventDetail

### Community 40 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 41 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 42 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 43 - "binding-accordion-field-focus.ts"
Cohesion: 0.52
Nodes (6): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains()

### Community 44 - "binding-dom-exclusive-display.ts"
Cohesion: 0.60
Nodes (5): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings()

### Community 45 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 46 - "Portfolio History Management Pagination"
Cohesion: 0.40
Nodes (5): Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination Retry and Loading Feedback, Portfolio History Management Pagination, Portfolio History Management Terminal State Without Completion Message

### Community 47 - "NotificationType"
Cohesion: 0.40
Nodes (5): NotificationType, ERROR, INFO, SUCCESS, WARNING

### Community 48 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **167 isolated node(s):** `AssetSearchAutocompleteOption`, `AssignAtPathOptions`, `ToIntOptions`, `GeneralErrorHandler`, `BootstrapNotification` (+162 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 208 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset-composed-columns-input/constants.ts`, `portfolio-history-management-observation-details/index.ts`, `pagination.ts`, `htmx/index.ts`, `package.json`, `asset.ts`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`?**
  _High betweenness centrality (0.082) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `bignumber.js` to `handlebars-lang.ts`, `allocation-plan-management.ts`, `logger`, `portfolio-chart.ts`, `allocation-plan.ts`, `package.json`, `portfolio-history-management.ts`, `chart-utils.ts`?**
  _High betweenness centrality (0.080) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `allocation-plan-management.ts`, `package.json`, `infra.ts`, `asset.ts`, `portfolio-history-management.ts`, `chart-types.ts`, `bignumber.js`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchAutocompleteOption`, `AssignAtPathOptions`, `ToIntOptions` to the rest of the system?**
  _167 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.050817160367722165 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.08069381598793364 - nodes in this community are weakly interconnected._