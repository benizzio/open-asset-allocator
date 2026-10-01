# Graph Report - web-static  (2026-10-01)

## Corpus Check
- 77 files · ~29,310 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 597 nodes · 1339 edges · 26 communities (19 shown, 7 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6f320cdc`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan.ts
- handlebars-lang.ts
- package.json
- asset.ts
- dom-utils.ts
- chart-utils.ts
- portfolio-chart.ts
- allocation-plan-management.ts
- binding-financial-input.ts
- devDependencies
- notifications.ts
- Portfolio Detail Page
- AssetComposedColumnInput
- compilerOptions
- portfolio-history-management.ts
- htmx/index.ts
- infra.ts
- Portfolio Section Navigation
- FormRowValueElements
- Frontend Module Architecture
- Recursive Planned Allocation Rows
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `registerHandlebarsLangHelpers()` - 16 edges
3. `LogLevel` - 15 edges
4. `handlebars` - 14 edges
5. `bignumber.js` - 13 edges
6. `FractalPortfolioMultiChartDataSource` - 13 edges
7. `AssetComposedColumnInput` - 12 edges
8. `MultiChartDataSource` - 11 edges
9. `chart.js` - 10 edges
10. `compilerOptions` - 10 edges

## Surprising Connections (you probably didn't know these)
- `getNextPortfolioHistoryManagementIndex()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/portfolio-history-management.ts → websrc/utils/lang.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPlannedAllocationMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/infra/chart/chart-types.ts
- `chartDataSelectionEventHandler()` --calls--> `changeChartDataOnDatasource()`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/infra/chart/chart-utils.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (26 total, 7 thin omitted)

### Community 0 - "logger"
Cohesion: 0.07
Nodes (61): navigo, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal() (+53 more)

### Community 1 - "allocation-plan.ts"
Cohesion: 0.06
Nodes (58): bignumber.js, allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets() (+50 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (45): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper() (+37 more)

### Community 3 - "package.json"
Cohesion: 0.05
Nodes (45): alias, bignumber.js, dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js (+37 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+22 more)

### Community 6 - "dom-utils.ts"
Cohesion: 0.19
Nodes (13): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, addReadyConditionToWaitingElement(), addReadyFlagObserverOnElement(), areAllConditionsReady() (+5 more)

### Community 7 - "chart-utils.ts"
Cohesion: 0.06
Nodes (38): chartjs-plugin-datalabels, chroma-js, patternomaly, Application, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart() (+30 more)

### Community 8 - "portfolio-chart.ts"
Cohesion: 0.10
Nodes (25): chart.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent() (+17 more)

### Community 9 - "allocation-plan-management.ts"
Cohesion: 0.17
Nodes (8): addPlannedAllocationRow(), AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), setHierarchicalIdFromParentRow(), toInt()

### Community 11 - "binding-financial-input.ts"
Cohesion: 0.10
Nodes (31): BootstrapClasses, BootstrapIconClasses, bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay() (+23 more)

### Community 12 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 13 - "notifications.ts"
Cohesion: 0.16
Nodes (10): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 15 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 17 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - "portfolio-history-management.ts"
Cohesion: 0.14
Nodes (11): htmx.org, allocationPlanManagement, ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnsInput, notifications, getNextPortfolioHistoryManagementIndex(), portfolioHistoryManagement, AssetPage (+3 more)

### Community 21 - "htmx/index.ts"
Cohesion: 0.19
Nodes (13): api, APIError, APIErrorResponse, addEventListeners(), configEnhancedRequestEventListener(), EventDetail, prepareFormData(), replaceRequestPathParams() (+5 more)

### Community 22 - "infra.ts"
Cohesion: 0.21
Nodes (9): AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra, setupGlobalErrorHandler() (+1 more)

### Community 23 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **117 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+112 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 163 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **7 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan.ts` to `logger`, `handlebars-lang.ts`, `package.json`, `chart-utils.ts`, `portfolio-chart.ts`, `allocation-plan-management.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.095) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `handlebars-lang.ts`, `dom-utils.ts`, `binding-financial-input.ts`, `htmx/index.ts`, `infra.ts`?**
  _High betweenness centrality (0.067) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `package.json`, `asset.ts`, `chart-utils.ts`, `portfolio-chart.ts`, `allocation-plan-management.ts`, `notifications.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _117 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.07122153209109731 - nodes in this community are weakly interconnected._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06265432098765432 - nodes in this community are weakly interconnected._