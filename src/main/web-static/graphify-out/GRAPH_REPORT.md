# Graph Report - web-static  (2026-11-21)

## Corpus Check
- 125 files · ~56,013 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 1027 nodes · 2458 edges · 67 communities (63 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 127 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `cb8f4aee`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- handlebars-lang.ts
- rendering.ts
- asset-external.ts
- chart-utils.ts
- asset-row-controller.ts
- AssetSearchAutocompleteController
- htmx/index.ts
- allocation-plan-management.ts
- package.json
- AssetRowController
- bindObservationCloning
- dom/index.ts
- application/index.ts
- pagination.ts
- portfolio-history-management.ts
- portfolio-chart.ts
- allocation-plan-management-clone/index.ts
- asset-composed-columns-input/index.ts
- Portfolio History Management
- allocation-plan-chart.ts
- FractalPortfolioMultiChartDataSource
- MultiChartDataSource
- binding-percentage-input.ts
- binding-financial-input.ts
- devDependencies
- routing/index.ts
- allocation-plan.ts
- dom-utils.ts
- dependencies
- asset-composed-columns-input/constants.ts
- portfolio-history-management-observation-details/index.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- Portfolio History Pagination
- binding-htmx-transform-response.ts
- compilerOptions
- management-metadata-editor/index.ts
- autocomplete-interactions.ts
- binding-htmx-trigger-on-route.ts
- Allocation Plan Pagination
- Allocation Plan Management Pagination
- controls.ts
- chart-contents.ts
- binding-dom-attribute-on-route.ts
- logger
- Asset Search Autocomplete
- Allocation Plan Management
- alias
- portfolio-history-management-observation-clone/draft.ts
- getContextId
- binding-dom-display-on-route.ts
- .proxyrc.js
- binding-accordion-field-focus.ts
- portfolio-history-management-observation-clone/index.ts
- Portfolio History Management Pagination
- scripts
- notifications.ts
- focus-tracking.ts
- ObservationDetailsState
- Frontend Module Architecture
- asset.ts
- Hierarchical Divergence Analysis
- Toast Notification
- chart.ts
- infra.ts
- logging.ts
- binding-dom-bootstrap-form.ts

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetRowController` - 23 edges
3. `handlebars` - 18 edges
4. `registerHandlebarsLangHelpers()` - 18 edges
5. `createPortfolioHistoryQuoteAction()` - 17 edges
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

## Communities (67 total, 4 thin omitted)

### Community 0 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (51): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate() (+43 more)

### Community 1 - "rendering.ts"
Cohesion: 0.39
Nodes (8): PlanDraft, PlanRowDraft, buildRowHierarchy(), renderPlanMarkup(), renderPlanRows(), validatePlanMarkup(), hierarchyFieldName(), hierarchyFieldSuffix()

### Community 2 - "asset-external.ts"
Cohesion: 0.16
Nodes (22): Asset, ExternalAsset, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields(), handleExternalSearchAfterRequest() (+14 more)

### Community 3 - "chart-utils.ts"
Cohesion: 0.15
Nodes (16): chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions, UNIDIMENSIONAL_DATASET_SUM_FIELD (+8 more)

### Community 4 - "asset-row-controller.ts"
Cohesion: 0.14
Nodes (16): CommittedAssetControls, CommittedAssetSelection, renderCommittedAssetSelection(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_NOT_FOUND_ERROR_MESSAGE, AssetSelectionState, EXISTING, NEW (+8 more)

### Community 5 - "AssetSearchAutocompleteController"
Cohesion: 0.19
Nodes (7): AssetSearchAutocompleteController, loadDatalists(), bindOutsidePointerDismissal(), clearOutsidePointerDismissal(), clearPendingOutsidePointer(), AssetSearchAutocompleteState, AssetTickerSelectionHandler

### Community 6 - "htmx/index.ts"
Cohesion: 0.16
Nodes (16): api, APIError, APIErrorResponse, addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs() (+8 more)

### Community 7 - "allocation-plan-management.ts"
Cohesion: 0.13
Nodes (19): AllocationPlanningHierarchicalFormEntry, bindNativePlanRowSubmitGuard(), copyAssetTickersToHierarchicalIdFields(), FORM_FIELD_DEPENDENT_ATTRIBUTE, FORM_FIELD_LABEL_ATTRIBUTE, FORM_LAST_ROW_INDEX_INPUT_NAME, PLAN_CASH_RESERVE_FIELD_NAME, PLAN_DETAILS_FIELD_NAME (+11 more)

### Community 8 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 10 - "bindObservationCloning"
Cohesion: 0.23
Nodes (13): bindAllocationPlanCloning(), presentDraft(), CloneCoordinator, focusDraftDestination(), openPanel(), bindCloneRequestGuards(), CloneOperation, currentPortfolio() (+5 more)

### Community 11 - "dom/index.ts"
Cohesion: 0.20
Nodes (17): captureRow(), HIERARCHY_FIELD_PATTERN, HierarchyValue, input(), restorePlanRows(), captureAssetDraft(), input(), restoreAssetDraft() (+9 more)

### Community 12 - "application/index.ts"
Cohesion: 0.10
Nodes (23): htmx.org, AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView() (+15 more)

### Community 13 - "pagination.ts"
Cohesion: 0.17
Nodes (19): stopTrackingFocus(), bindPaginationLifecycle(), LifecycleHandlers, LOAD_MORE_EVENT, trackRequestCompletion(), bindRequestToXhr(), commitSettledPage(), finishRequest() (+11 more)

### Community 14 - "portfolio-history-management.ts"
Cohesion: 0.18
Nodes (17): notifications, PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR (+9 more)

### Community 15 - "portfolio-chart.ts"
Cohesion: 0.32
Nodes (9): chart.js, getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+1 more)

### Community 16 - "allocation-plan-management-clone/index.ts"
Cohesion: 0.18
Nodes (18): capturePlanDraft(), boundLists, copyDraft(), destinationControls(), allocationPlanNameValidationMessage(), bindAllocationPlanMetadataSubmitGuard(), boundContainers, captureAllocationPlanName() (+10 more)

### Community 17 - "asset-composed-columns-input/index.ts"
Cohesion: 0.11
Nodes (26): ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR, ASSET_SEARCH_INPUT_ATTRIBUTE, ASSET_SELECTION_ERROR_MESSAGE, ASSET_TICKER_INPUT_ATTRIBUTE, INVALID_INPUT_CLASS (+18 more)

### Community 18 - "Portfolio History Management"
Cohesion: 0.14
Nodes (15): Icon-Only Observation Clone Control, Dated Observation Header, Disabled Observation Clone Tooltip Wrapper, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor (+7 more)

### Community 19 - "allocation-plan-chart.ts"
Cohesion: 0.21
Nodes (12): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+4 more)

### Community 20 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.26
Nodes (4): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey()

### Community 21 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 22 - "binding-percentage-input.ts"
Cohesion: 0.42
Nodes (8): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer()

### Community 23 - "binding-financial-input.ts"
Cohesion: 0.17
Nodes (21): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+13 more)

### Community 24 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 25 - "routing/index.ts"
Cohesion: 0.31
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), buildParameterizedDestinationPathFromCurrentLocationContext(), hasPathParam() (+4 more)

### Community 26 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (42): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+34 more)

### Community 27 - "dom-utils.ts"
Cohesion: 0.28
Nodes (5): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements

### Community 28 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 29 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.23
Nodes (13): ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_AUTOCOMPLETE_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_OPTION_SELECTOR (+5 more)

### Community 30 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.29
Nodes (14): bindObservationDetailsLoader(), handleAfterRequest(), handleAfterSettle(), handleBeforeRequest(), handleBeforeSwap(), readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail (+6 more)

### Community 31 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 32 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 33 - "Portfolio History Pagination"
Cohesion: 0.17
Nodes (13): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+5 more)

### Community 34 - "binding-htmx-transform-response.ts"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 35 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 36 - "management-metadata-editor/index.ts"
Cohesion: 0.25
Nodes (10): activateManagementMetadataEditor(), bindReadOnlyManagementMetadataHeaderToggle(), boundHeaderToggleContainers, hideManagementMetadataEditor(), interactiveHeaderControlSelector, METADATA_EDITOR_ACTIVE_ATTRIBUTE, revealManagementMetadataEditor(), toggleReadOnlyManagementMetadataPanel() (+2 more)

### Community 37 - "autocomplete-interactions.ts"
Cohesion: 0.27
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AssetSearchAutocompleteOption, AutocompleteControllerActions

### Community 38 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.29
Nodes (11): addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute(), extractBindingData() (+3 more)

### Community 39 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 40 - "Allocation Plan Management Pagination"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, Allocation Plan Management Terminal State Without Completion Message, focusOnAcceptedRecord() (+2 more)

### Community 41 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 42 - "chart-contents.ts"
Cohesion: 0.14
Nodes (16): bignumber.js, toChartContent(), toUnidimensionalMultiChartContent(), getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), portfolioChart, ObservationTimestamp, PortfolioAllocation (+8 more)

### Community 43 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 44 - "logger"
Cohesion: 0.51
Nodes (9): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions() (+1 more)

### Community 45 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 46 - "Allocation Plan Management"
Cohesion: 0.31
Nodes (9): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Allocation Plan Clone Control, Dated Allocation Plan Header, activateAllocationPlanEditor() (+1 more)

### Community 48 - "portfolio-history-management-observation-clone/draft.ts"
Cohesion: 0.17
Nodes (18): OBSERVATION_ALLOCATION_FIELDS, OBSERVATION_ALLOCATION_SELECTORS, OBSERVATION_ASSET_QUANTITY_FIELD, OBSERVATION_ASSET_QUANTITY_SELECTOR, observationAllocationFieldSelector(), AllocationDraft, captureKeys(), captureRow() (+10 more)

### Community 49 - "getContextId"
Cohesion: 0.32
Nodes (7): allocation-plan-management portfolio context selector, getContextId(), configureRequest(), getOwner(), getState(), requestNextPage(), updateControls()

### Community 50 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 51 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 52 - "binding-accordion-field-focus.ts"
Cohesion: 0.52
Nodes (6): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains()

### Community 53 - "portfolio-history-management-observation-clone/index.ts"
Cohesion: 0.20
Nodes (18): NEW_OBSERVATION_FORM_ID, renderRows(), captureObservationDraft(), restoreObservationRows(), applyMetadata(), boundLists, copyDraft(), DestinationControls (+10 more)

### Community 54 - "Portfolio History Management Pagination"
Cohesion: 0.33
Nodes (6): Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management portfolio context selector, Portfolio History Management Terminal State Without Completion Message

### Community 55 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 56 - "notifications.ts"
Cohesion: 0.15
Nodes (11): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR (+3 more)

### Community 57 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 58 - "ObservationDetailsState"
Cohesion: 0.40
Nodes (5): ObservationDetailsState, FAILED, LOADING, READY, UNLOADED

### Community 59 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 60 - "asset.ts"
Cohesion: 0.15
Nodes (9): AssetBeforeSwapEvent, AssetRequestEvent, handleExternalSearchBeforeRequest(), hasValidExternalData(), isExternalAsset(), getAssetIdentifierFromLocation(), getCurrentAssetIdentifierForResponse(), normalizeAsset() (+1 more)

### Community 63 - "chart.ts"
Cohesion: 0.17
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+7 more)

### Community 64 - "infra.ts"
Cohesion: 0.21
Nodes (8): Application, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra, setupGlobalErrorHandler()

### Community 65 - "logging.ts"
Cohesion: 0.31
Nodes (8): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), Level, Levels, LogLevel

### Community 66 - "binding-dom-bootstrap-form.ts"
Cohesion: 0.70
Nodes (4): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants()

## Knowledge Gaps
- **188 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+183 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 228 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `application/index.ts` to `htmx/index.ts`, `allocation-plan-management.ts`, `package.json`, `binding-htmx-trigger-on-route.ts`, `pagination.ts`, `portfolio-history-management.ts`, `allocation-plan-management-clone/index.ts`, `asset-composed-columns-input/index.ts`, `portfolio-history-management-observation-clone/index.ts`, `asset.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.103) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `chart-contents.ts` to `handlebars-lang.ts`, `chart-utils.ts`, `allocation-plan-management.ts`, `package.json`, `portfolio-history-management.ts`, `binding-percentage-input.ts`, `allocation-plan.ts`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `rendering.ts`, `allocation-plan-management.ts`, `package.json`, `chart-contents.ts`, `portfolio-history-management.ts`, `allocation-plan-management-clone/index.ts`, `portfolio-history-management-observation-clone/draft.ts`, `portfolio-history-management-observation-clone/index.ts`, `notifications.ts`, `asset.ts`, `chart.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _188 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06954997077732321 - nodes in this community are weakly interconnected._
- **Should `asset-row-controller.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.13725490196078433 - nodes in this community are weakly interconnected._