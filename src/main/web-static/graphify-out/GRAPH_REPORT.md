# Graph Report - web-static  (2026-10-22)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 790 nodes · 1806 edges · 48 communities (45 shown, 3 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 105 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `401db7b3`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan.ts
- handlebars-lang.ts
- asset.ts
- infra.ts
- allocation-plan-management.ts
- package.json
- chart-utils.ts
- portfolio-history-management.ts
- chart.ts
- portfolio-chart-datasource.ts
- AssetRowController
- AssetSearchAutocompleteController
- binding-htmx-trigger-on-route.ts
- htmx/index.ts
- asset-composed-columns-input/index.ts
- asset-row-controller.ts
- MultiChartDataSource
- binding-financial-input.ts
- devDependencies
- asset-composed-columns-input/constants.ts
- Portfolio History Management
- routing/index.ts
- dependencies
- portfolio-chart.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- autocomplete-interactions.ts
- allocation-plan-management-metadata/index.ts
- compilerOptions
- portfolio-history-management-observation-details/index.ts
- portfolio-history-management-observation-metadata/index.ts
- logging.ts
- logger
- binding-htmx-transform-response.ts
- binding-dom-attribute-on-route.ts
- Allocation Plan Management
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- handlebars-format.ts
- .proxyrc.js
- Asset Search Autocomplete
- application/portfolio-analysis.ts
- scripts
- AssetSelectionState
- Portfolio Section Navigation
- Frontend Module Architecture
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
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `Allocation Plan Pencil Metadata Editor` --references--> `handleInputAllocationPlanDate()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (48 total, 3 thin omitted)

### Community 0 - "allocation-plan.ts"
Cohesion: 0.07
Nodes (55): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+47 more)

### Community 1 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (48): handlebars, setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), andHelper(), arrayHelper() (+40 more)

### Community 2 - "asset.ts"
Cohesion: 0.08
Nodes (36): Asset, ExternalAsset, NotificationType, ERROR, INFO, SUCCESS, WARNING, AssetBeforeSwapEvent (+28 more)

### Community 3 - "infra.ts"
Cohesion: 0.09
Nodes (26): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus() (+18 more)

### Community 4 - "allocation-plan-management.ts"
Cohesion: 0.12
Nodes (16): htmx.org, Application, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure() (+8 more)

### Community 5 - "package.json"
Cohesion: 0.10
Nodes (21): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+13 more)

### Community 6 - "chart-utils.ts"
Cohesion: 0.13
Nodes (18): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+10 more)

### Community 7 - "portfolio-history-management.ts"
Cohesion: 0.18
Nodes (17): bignumber.js, notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR (+9 more)

### Community 8 - "chart.ts"
Cohesion: 0.16
Nodes (17): chart.js, chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID (+9 more)

### Community 9 - "portfolio-chart-datasource.ts"
Cohesion: 0.20
Nodes (9): FractalPortfolioMultiChartDataSource, generateDataKey(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData, AllocationStructure (+1 more)

### Community 11 - "AssetSearchAutocompleteController"
Cohesion: 0.17
Nodes (8): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteOption, AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 12 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.15
Nodes (15): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants() (+7 more)

### Community 13 - "htmx/index.ts"
Cohesion: 0.17
Nodes (15): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+7 more)

### Community 14 - "asset-composed-columns-input/index.ts"
Cohesion: 0.16
Nodes (16): ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_INPUT_ATTRIBUTE, ASSET_SELECTION_ERROR_MESSAGE, ASSET_TICKER_INPUT_ATTRIBUTE, INVALID_INPUT_CLASS, TICKER_EXTRA_ERROR_MESSAGE_ATTRIBUTE (+8 more)

### Community 15 - "asset-row-controller.ts"
Cohesion: 0.15
Nodes (15): ASSET_NOT_FOUND_ERROR_MESSAGE, FAILED_TO_FETCH_ASSET_ERROR_PREFIX, INVALID_TICKER_ERROR_MESSAGE, NEW_ASSET_TICKER_MESSAGE_SELECTOR, REQUIRED_FOR_SEARCH_ERROR_MESSAGE, TICKER_CANDIDATE_PATTERN, TICKER_TOO_LONG_ERROR_MESSAGE, ASSET_ROW_SELECTION_CHANGE_EVENT (+7 more)

### Community 16 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 17 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 18 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 19 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 20 - "Portfolio History Management"
Cohesion: 0.16
Nodes (12): Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization, Portfolio History Management (+4 more)

### Community 21 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 22 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 23 - "portfolio-chart.ts"
Cohesion: 0.20
Nodes (11): changeChartData(), chartDataSelectionEventHandler(), getChartContent(), interactionObserverCallback(), portfolioChart, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO (+3 more)

### Community 24 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 25 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 26 - "autocomplete-interactions.ts"
Cohesion: 0.28
Nodes (7): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, AutocompleteControllerActions

### Community 27 - "allocation-plan-management-metadata/index.ts"
Cohesion: 0.27
Nodes (10): Allocation Plan Pencil Metadata Editor, activateAllocationPlanEditor(), bindAllocationPlanMetadataRequest(), boundContainers, copyEditedPlanMetadata(), handleInputAllocationPlanDate(), PlanRequest, prepareAllocationPlanMetadataForPost() (+2 more)

### Community 28 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 29 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.30
Nodes (11): hideManagementMetadataEditor(), bindObservationDetailsLoader(), handleAfterRequest(), handleBeforeRequest(), ObservationDetailsState, readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail (+3 more)

### Community 30 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.27
Nodes (11): areObservationDetailsReady(), bindObservationMetadataRequest(), boundContainers, copyEditedTag(), getObservationForm(), getObservationTimestamp(), ObservationRequest, prepareObservationMetadataForPost() (+3 more)

### Community 31 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 32 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 33 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 34 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 35 - "Allocation Plan Management"
Cohesion: 0.25
Nodes (9): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Dated Allocation Plan Header, Recursive Planned Allocation Rows, handleHierarchicalIdLevelChange(), handleRemovePlannedAllocationRow() (+1 more)

### Community 36 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 37 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 38 - "handlebars-format.ts"
Cohesion: 0.46
Nodes (6): handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), formatCalendarDate(), getBrowserLocalCalendarDate(), timestampToBrowserLocalCalendarDate()

### Community 39 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 40 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 41 - "application/portfolio-analysis.ts"
Cohesion: 0.38
Nodes (5): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence

### Community 42 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 43 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

### Community 44 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 45 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **139 isolated node(s):** `PlannedAllocationDTO`, `AllocationHierarchyLevelDTO`, `AssignAtPathOptions`, `ToIntOptions`, `AssetSearchAutocompleteOption` (+134 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 179 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `portfolio-history-management.ts` to `allocation-plan.ts`, `logger`, `handlebars-lang.ts`, `allocation-plan-management.ts`, `package.json`, `chart-utils.ts`, `application/portfolio-analysis.ts`, `portfolio-chart.ts`?**
  _High betweenness centrality (0.089) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset.ts`, `package.json`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`, `htmx/index.ts`, `asset-composed-columns-input/index.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.067) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `infra.ts`, `allocation-plan-management.ts`, `package.json`, `handlebars-format.ts`, `portfolio-history-management.ts`, `chart.ts`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.060) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `PlannedAllocationDTO`, `AllocationHierarchyLevelDTO`, `AssignAtPathOptions` to the rest of the system?**
  _139 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06736842105263158 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07407407407407407 - nodes in this community are weakly interconnected._