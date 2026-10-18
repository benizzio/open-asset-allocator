# Graph Report - web-static  (2026-10-18)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 750 nodes · 1709 edges · 40 communities (37 shown, 3 thin omitted)
- Extraction: 95% EXTRACTED · 5% INFERRED · 0% AMBIGUOUS · INFERRED: 93 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `fdadb686`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset-composed-columns-input/constants.ts
- handlebars-lang.ts
- allocation-plan.ts
- chart-utils.ts
- asset.ts
- dom/index.ts
- AssetRowController
- htmx/index.ts
- allocation-plan-management.ts
- binding-financial-input.ts
- package.json
- portfolio-chart.ts
- allocation-plan-chart.ts
- MultiChartDataSource
- devDependencies
- infra.ts
- portfolio-history-management.ts
- routing/index.ts
- dependencies
- notifications.ts
- createPortfolioHistoryQuoteAction
- Portfolio Detail Page
- FractalPortfolioMultiChartDataSource
- binding-htmx-trigger-on-route.ts
- Asset Search Autocomplete
- compilerOptions
- quote-action.ts
- logger
- binding-dom-attribute-on-route.ts
- binding-htmx-wait-for-ready.ts
- binding-dom-display-on-route.ts
- chart-contents.ts
- .proxyrc.js
- binding-dom-exclusive-display.ts
- scripts
- FormRowValueElements
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
9. `FractalPortfolioMultiChartDataSource` - 13 edges
10. `AutocompleteControllerActions` - 13 edges

## Surprising Connections (you probably didn't know these)
- `Allocation Plan Header Field Focus Opens Accordion` --references--> `bindAccordionFieldFocusInDescendants()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/infra/dom/binding-accordion-field-focus.ts
- `Observation Header Field Focus Opens Accordion` --references--> `bindAccordionFieldFocusInDescendants()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/infra/dom/binding-accordion-field-focus.ts
- `addPlannedAllocationRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts
- `registerTransformResponseFunction()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/htmx/binding-htmx-transform-response.ts → websrc/infra/logging.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (40 total, 3 thin omitted)

### Community 0 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.05
Nodes (61): api, APIError, AssetSearchAutocompleteController, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_ACTION_BUTTON_IDENTITIES (+53 more)

### Community 1 - "handlebars-lang.ts"
Cohesion: 0.07
Nodes (49): handlebars, setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper() (+41 more)

### Community 2 - "allocation-plan.ts"
Cohesion: 0.10
Nodes (41): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+33 more)

### Community 3 - "chart-utils.ts"
Cohesion: 0.10
Nodes (31): chart.js, chroma-js, patternomaly, chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart() (+23 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts, fillEmptyAssetFields() (+22 more)

### Community 5 - "dom/index.ts"
Cohesion: 0.09
Nodes (25): bootstrap, External Asset Quote Action, Observation Date Picker, Observation Editor, Observation Header Field Focus Opens Accordion, Portfolio History Management, Portfolio History New Observation Header Field Group, AccordionFieldFocusState (+17 more)

### Community 6 - "AssetRowController"
Cohesion: 0.15
Nodes (7): AssetRowController, handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), selectAssetTicker(), AutocompleteControllerActions

### Community 7 - "htmx/index.ts"
Cohesion: 0.12
Nodes (25): APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP (+17 more)

### Community 8 - "allocation-plan-management.ts"
Cohesion: 0.11
Nodes (16): htmx.org, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey() (+8 more)

### Community 9 - "binding-financial-input.ts"
Cohesion: 0.14
Nodes (22): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants() (+14 more)

### Community 10 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 11 - "portfolio-chart.ts"
Cohesion: 0.21
Nodes (14): bignumber.js, getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+6 more)

### Community 12 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 13 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 14 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 15 - "infra.ts"
Cohesion: 0.16
Nodes (12): chartjs-plugin-datalabels, Application, DomInfra, handlebarsInfra, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER() (+4 more)

### Community 16 - "portfolio-history-management.ts"
Cohesion: 0.19
Nodes (10): mapForAllocationPlanManagement(), handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, setNewObservationTimestampFromDate(), synchronizeExternalAssetKeyGroup(), RequestConfigEventDetail (+2 more)

### Community 17 - "routing/index.ts"
Cohesion: 0.32
Nodes (12): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), bootNavigoRouter(), buildParameterizedDestinationPathFromCurrentLocationContext() (+4 more)

### Community 18 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 19 - "notifications.ts"
Cohesion: 0.16
Nodes (10): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 20 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 21 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 22 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.26
Nodes (4): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey()

### Community 23 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.27
Nodes (11): navigo, addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement(), bindHTMXTriggerOnRouteInDescendants(), bindRouteToHTMXEventOnElements(), bindRouteToHTMXTriggerOnElement(), CLEAN_ON_EXIT_HTMX_EVENT_HANDLERS, executeImmediatelyIfOnRoute() (+3 more)

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.23
Nodes (12): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label) (+4 more)

### Community 25 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 26 - "quote-action.ts"
Cohesion: 0.33
Nodes (9): EXTERNAL_ASSET_EXCHANGE_ID_SELECTOR, EXTERNAL_ASSET_KEYS_SELECTOR, EXTERNAL_ASSET_SOURCE_SELECTOR, EXTERNAL_ASSET_TICKER_SELECTOR, HISTORY_ROW_ID_PATTERN, PORTFOLIO_ALLOCATION_MANAGEMENT_FORM_PREFIX, PortfolioHistoryQuoteAction, readExternalAssetKeys() (+1 more)

### Community 27 - "logger"
Cohesion: 0.44
Nodes (9): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+1 more)

### Community 28 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.40
Nodes (9): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes() (+1 more)

### Community 29 - "binding-htmx-wait-for-ready.ts"
Cohesion: 0.50
Nodes (8): addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady(), bindConfirmGateOnElement(), bindHTMXWaitForReadyInDescendants(), bindWaitForReadyOnElements(), parseCommaSeparatedSet(), parseRequiredConditions()

### Community 30 - "binding-dom-display-on-route.ts"
Cohesion: 0.42
Nodes (8): addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements(), bindDisplayOnRouteInDescendants(), changeElementDisplay(), configDisplayOnRouteHooks(), executeImmediatelyIfOnRoute(), navigoRouter

### Community 31 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 32 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 33 - "binding-dom-exclusive-display.ts"
Cohesion: 0.60
Nodes (5): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings()

### Community 34 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 36 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 37 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **128 isolated node(s):** `AssetSearchAutocompleteOption`, `APIError`, `AssignAtPathOptions`, `ToIntOptions`, `PortfolioAllocationDTO` (+123 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 172 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `portfolio-chart.ts` to `handlebars-lang.ts`, `allocation-plan.ts`, `chart-utils.ts`, `allocation-plan-management.ts`, `binding-financial-input.ts`, `package.json`, `portfolio-history-management.ts`, `quote-action.ts`, `logger`?**
  _High betweenness centrality (0.088) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset-composed-columns-input/constants.ts`, `asset.ts`, `htmx/index.ts`, `package.json`, `portfolio-history-management.ts`, `binding-htmx-trigger-on-route.ts`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `chart-utils.ts`, `asset.ts`, `allocation-plan-management.ts`, `binding-financial-input.ts`, `package.json`, `portfolio-history-management.ts`, `notifications.ts`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchAutocompleteOption`, `APIError`, `AssignAtPathOptions` to the rest of the system?**
  _128 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05249569707401033 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.07205513784461152 - nodes in this community are weakly interconnected._