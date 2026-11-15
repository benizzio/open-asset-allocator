# Graph Report - web-static  (2026-11-15)

## Corpus Check
- 123 files · ~55,159 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 1011 nodes · 2380 edges · 63 communities (59 shown, 4 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 127 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `29742a2a`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- portfolio-history-management-observation-clone/index.ts
- asset.ts
- allocation-plan-management.ts
- allocation-plan.ts
- AssetRowController
- handlebars-lang.ts
- asset-composed-columns-input/index.ts
- package.json
- pagination.ts
- Portfolio History Management
- dom/index.ts
- chart-utils.ts
- dom-utils.ts
- htmx/index.ts
- allocation-plan-management-clone/index.ts
- allocation-plan-chart.ts
- Allocation Plan Pagination
- binding-htmx-trigger-on-route.ts
- binding-financial-input.ts
- devDependencies
- chart-contents.ts
- routing/index.ts
- asset-composed-columns-input/constants.ts
- dependencies
- portfolio-chart.ts
- getContextId
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- portfolio-history-management-observation-details/index.ts
- Portfolio History Pagination
- compilerOptions
- application/index.ts
- FractalPortfolioMultiChartDataSource
- logger
- Portfolio History Management Pagination
- autocomplete-controller.ts
- controls.ts
- binding-dom-attribute-on-route.ts
- binding-htmx-wait-for-ready.ts
- root.ts
- binding-dom-display-on-route.ts
- ObservationDetailsState
- bindObservationCloning
- Asset Search Autocomplete
- autocomplete-interactions.ts
- MultiChartDataSource
- .proxyrc.js
- Allocation Plan Management Pagination
- portfolio-history-management.ts
- portfolio-history-management-observation-clone/draft.ts
- scripts
- focus-tracking.ts
- allocation-plan-management-metadata/index.ts
- Frontend Module Architecture
- management-metadata-editor/index.ts
- Hierarchical Divergence Analysis
- Toast Notification
- application/portfolio-analysis.ts
- Allocation Plan Management
- binding-accordion-field-focus.ts
- response.ts
- api.ts
- alias

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
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan/allocation-plan-management.html → websrc/components/allocation-plan/allocation-plan-management-metadata/index.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, websrc_components_portfolio_history_portfolio_history_portfolio_history_viewer, websrc_components_allocation_plan_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (63 total, 4 thin omitted)

### Community 0 - "portfolio-history-management-observation-clone/index.ts"
Cohesion: 0.19
Nodes (19): isManagementMetadataEditorActive(), NEW_OBSERVATION_FORM_ID, renderRows(), captureObservationDraft(), ObservationDraft, applyMetadata(), boundLists, copyDraft() (+11 more)

### Community 1 - "asset.ts"
Cohesion: 0.07
Nodes (38): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, ExternalAsset, DomInfra, Notification, NotificationType, ERROR, INFO (+30 more)

### Community 2 - "allocation-plan-management.ts"
Cohesion: 0.18
Nodes (13): AllocationPlanningHierarchicalFormEntry, bindNativePlanRowSubmitGuard(), copyAssetTickersToHierarchicalIdFields(), FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), handleHierarchicalIdLevelChange(), removePlannedAllocationSubtree(), mapFormRowHierarchicalStructure() (+5 more)

### Community 3 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (41): htmx.org, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+33 more)

### Community 5 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (53): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate() (+45 more)

### Community 6 - "asset-composed-columns-input/index.ts"
Cohesion: 0.16
Nodes (16): UNABLE_TO_RESOLVE_ASSET_ROW_ERROR, bindAssetFormValidationGuards(), guardedAssetForms, invalidateSelectedAsset(), validateAssetRowForPost(), validateAssetRowsForPost(), autocompleteController, selectAssetTicker() (+8 more)

### Community 7 - "package.json"
Cohesion: 0.10
Nodes (21): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, eslint-plugin-sonarjs, globals, htmx-ext-client-side-templates (+13 more)

### Community 8 - "pagination.ts"
Cohesion: 0.17
Nodes (19): stopTrackingFocus(), bindPaginationLifecycle(), LifecycleHandlers, LOAD_MORE_EVENT, trackRequestCompletion(), bindRequestToXhr(), commitSettledPage(), finishRequest() (+11 more)

### Community 9 - "Portfolio History Management"
Cohesion: 0.14
Nodes (15): Icon-Only Observation Clone Control, Dated Observation Header, Disabled Observation Clone Tooltip Wrapper, External Asset Quote Action, FormRowValueElements, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor (+7 more)

### Community 10 - "dom/index.ts"
Cohesion: 0.22
Nodes (13): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), PERCENTAGE_INPUT_ATTRIBUTE, PERCENTAGE_INPUT_CONTAINER_ATTRIBUTE, maskNumberDecimalPlaces(), maskTagInput() (+5 more)

### Community 11 - "chart-utils.ts"
Cohesion: 0.10
Nodes (31): chart.js, chartjs-plugin-datalabels, chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions() (+23 more)

### Community 12 - "dom-utils.ts"
Cohesion: 0.19
Nodes (10): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), addRemoveObserver(), contextDataCache, ensureSharedObserver() (+2 more)

### Community 13 - "htmx/index.ts"
Cohesion: 0.11
Nodes (24): APIErrorResponse, bindHTMXTransformResponseInDescendants(), addEventListeners(), afterRequestErrorEventHandlers, configEnhancedRequestEventListener(), EventDetail, getAfterRequestErrorNotificationHandlerSkippingSuppressedXhrs(), prepareFormData() (+16 more)

### Community 14 - "allocation-plan-management-clone/index.ts"
Cohesion: 0.15
Nodes (24): capturePlanDraft(), captureRow(), HIERARCHY_FIELD_PATTERN, HierarchyValue, input(), PlanDraft, PlanRowDraft, restorePlanRows() (+16 more)

### Community 15 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 16 - "Allocation Plan Pagination"
Cohesion: 0.24
Nodes (10): Allocation Plan Consistent List-to-Control Gap, Allocation Plan Icon-Only Continuation Controls, Allocation Plan Management Navigation Outside Swapped List, Allocation Plan Pagination, Allocation Plan Pagination Retry and Loading Feedback, allocation-plan portfolio context selector, Allocation Plan Terminal State Without Completion Message, Allocation Plan Viewer (+2 more)

### Community 17 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.27
Nodes (11): RequestConfigEventDetail, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute() (+3 more)

### Community 18 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 19 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 20 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 21 - "routing/index.ts"
Cohesion: 0.30
Nodes (13): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+5 more)

### Community 22 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.12
Nodes (29): CommittedAssetControls, CommittedAssetSelection, renderCommittedAssetSelection(), ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR, ASSET_ID_INPUT_SELECTOR, ASSET_LOOKUP_PENDING_ERROR_MESSAGE, ASSET_NAME_INPUT_SELECTOR (+21 more)

### Community 23 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 24 - "portfolio-chart.ts"
Cohesion: 0.20
Nodes (13): getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, MappedChartData, AllocationStructure, PortfolioAllocation (+5 more)

### Community 25 - "getContextId"
Cohesion: 0.32
Nodes (7): allocation-plan-management portfolio context selector, getContextId(), configureRequest(), getOwner(), getState(), requestNextPage(), updateControls()

### Community 26 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 27 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, htmx.trigger, navigateTo, New Portfolio Form, Portfolio List Page (+5 more)

### Community 28 - "portfolio-history-management-observation-details/index.ts"
Cohesion: 0.32
Nodes (13): bindObservationDetailsLoader(), handleAfterRequest(), handleAfterSettle(), handleBeforeRequest(), handleBeforeSwap(), readObservationDetailsState(), reloadObservationDetails(), RequestEventDetail (+5 more)

### Community 29 - "Portfolio History Pagination"
Cohesion: 0.17
Nodes (13): Allocation Map, Portfolio Section Navigation, History Management Navigation Outside Swapped List Revealed After List Swap, Portfolio History Consistent List-to-Control Gap, Portfolio History Icon-Only Continuation Controls, Portfolio History Item-Only Template, Portfolio History Pagination, Portfolio History Pagination Retry and Loading Feedback (+5 more)

### Community 30 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 31 - "application/index.ts"
Cohesion: 0.13
Nodes (15): AllocationPlanPageItem, AllocationPlanViewPage, registerAllocationPlanPageTransform(), transformAllocationPlanView(), HistoryPageItem, registerPortfolioHistoryPageTransform(), transformPortfolioHistoryView(), allocationPlanManagement (+7 more)

### Community 32 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.23
Nodes (6): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AppliedAllocationHierarchyLevel, AllocationHierarchyLevel

### Community 33 - "logger"
Cohesion: 0.19
Nodes (18): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+10 more)

### Community 34 - "Portfolio History Management Pagination"
Cohesion: 0.33
Nodes (6): Portfolio History Management Consistent List-to-Control Gap, Portfolio History Management Icon-Only Continuation Controls, Portfolio History Management Pagination, Portfolio History Management Pagination Retry and Loading Feedback, portfolio-history-management portfolio context selector, Portfolio History Management Terminal State Without Completion Message

### Community 35 - "autocomplete-controller.ts"
Cohesion: 0.11
Nodes (17): AssetSearchAutocompleteController, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_SEARCH_LISTBOX_ATTRIBUTE, ASSET_SEARCH_OPTION_CLASS, ASSET_SEARCH_STATUS_ATTRIBUTE (+9 more)

### Community 36 - "controls.ts"
Cohesion: 0.36
Nodes (9): findOwnedElement(), getButtonLabel(), PaginationControlState, renderControls(), setChildHidden(), shouldShowButton(), updateButton(), updateControlsGroup() (+1 more)

### Community 37 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 38 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 39 - "root.ts"
Cohesion: 0.25
Nodes (5): Application, notifications, AfterRequestEventDetail, HtmxInfra, Infra

### Community 40 - "binding-dom-display-on-route.ts"
Cohesion: 0.50
Nodes (7): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute()

### Community 41 - "ObservationDetailsState"
Cohesion: 0.40
Nodes (5): ObservationDetailsState, FAILED, LOADING, READY, UNLOADED

### Community 42 - "bindObservationCloning"
Cohesion: 0.22
Nodes (14): bootstrap, bindAllocationPlanCloning(), presentDraft(), CloneCoordinator, focusDraftDestination(), openPanel(), bindCloneRequestGuards(), CloneOperation (+6 more)

### Community 43 - "Asset Search Autocomplete"
Cohesion: 0.36
Nodes (8): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Static Asset Ticker and Name Labels, Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 44 - "autocomplete-interactions.ts"
Cohesion: 0.32
Nodes (6): handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ASSET_SEARCH_OPTION_SELECTOR, AutocompleteControllerActions

### Community 45 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 46 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 47 - "Allocation Plan Management Pagination"
Cohesion: 0.20
Nodes (10): allocation-plan persisted record focus markers, Allocation Plan Management Consistent List-to-Control Gap, Allocation Plan Management Icon-Only Continuation Controls, Allocation Plan Management Pagination, Allocation Plan Management Pagination Retry and Loading Feedback, allocation-plan-management persisted record focus markers, Allocation Plan Management Terminal State Without Completion Message, focusOnAcceptedRecord() (+2 more)

### Community 48 - "portfolio-history-management.ts"
Cohesion: 0.19
Nodes (15): bignumber.js, PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup(), EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR (+7 more)

### Community 49 - "portfolio-history-management-observation-clone/draft.ts"
Cohesion: 0.26
Nodes (13): captureAssetDraft(), input(), restoreAssetDraft(), AllocationDraft, captureKeys(), captureRow(), field(), restoreObservationRows() (+5 more)

### Community 50 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 51 - "focus-tracking.ts"
Cohesion: 0.50
Nodes (4): FocusContext, onFocusChange(), startTrackingFocus(), tracked

### Community 52 - "allocation-plan-management-metadata/index.ts"
Cohesion: 0.22
Nodes (12): allocationPlanNameValidationMessage(), bindAllocationPlanMetadataSubmitGuard(), boundContainers, copyEditedPlanMetadata(), handleInputAllocationPlanName(), NEW_PLAN_DATE_SELECTOR, NEW_PLAN_NAME_SELECTOR, PLAN_DATE_EDITOR_ATTRIBUTE (+4 more)

### Community 53 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 54 - "management-metadata-editor/index.ts"
Cohesion: 0.23
Nodes (11): activateAllocationPlanEditor(), activateManagementMetadataEditor(), bindReadOnlyManagementMetadataHeaderToggle(), boundHeaderToggleContainers, hideManagementMetadataEditor(), interactiveHeaderControlSelector, METADATA_EDITOR_ACTIVE_ATTRIBUTE, revealManagementMetadataEditor() (+3 more)

### Community 57 - "application/portfolio-analysis.ts"
Cohesion: 0.31
Nodes (6): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, Format

### Community 58 - "Allocation Plan Management"
Cohesion: 0.36
Nodes (8): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Allocation Plan Clone Control, Dated Allocation Plan Header, handleInputAllocationPlanDate()

### Community 59 - "binding-accordion-field-focus.ts"
Cohesion: 0.52
Nodes (6): AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains()

### Community 60 - "response.ts"
Cohesion: 0.47
Nodes (4): mapForAllocationPlanManagement(), mapPlanPage(), websrc_pages_index_portfoliopage, PortfolioPage

### Community 61 - "api.ts"
Cohesion: 0.50
Nodes (3): api, APIError, Asset

## Knowledge Gaps
- **186 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+181 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 226 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `htmx.org` connect `allocation-plan.ts` to `portfolio-history-management-observation-clone/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `asset-composed-columns-input/index.ts`, `package.json`, `pagination.ts`, `htmx/index.ts`, `allocation-plan-management-clone/index.ts`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`, `portfolio-history-management-observation-details/index.ts`?**
  _High betweenness centrality (0.105) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `portfolio-history-management-observation-clone/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `chart-utils.ts`, `allocation-plan-management-clone/index.ts`, `portfolio-history-management.ts`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.068) - this node is a cross-community bridge._
- **Why does `bignumber.js` connect `portfolio-history-management.ts` to `logger`, `allocation-plan-management.ts`, `allocation-plan.ts`, `handlebars-lang.ts`, `package.json`, `chart-utils.ts`, `portfolio-chart.ts`, `application/portfolio-analysis.ts`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _186 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.0726950354609929 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09155844155844156 - nodes in this community are weakly interconnected._