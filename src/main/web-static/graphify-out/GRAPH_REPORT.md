# Graph Report - web-static  (2026-09-27)

## Corpus Check
- 76 files · ~30,844 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 607 nodes · 1371 edges · 30 communities (25 shown, 5 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `6f320cdc`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- chart-utils.ts
- handlebars-lang.ts
- package.json
- asset.ts
- asset-composed-columns-input.ts
- portfolio-chart.ts
- dom/index.ts
- fractal-allocation-plan-mapping.ts
- allocation-plan-chart.ts
- service/index.ts
- binding-financial-input.ts
- devDependencies
- notifications.ts
- bignumber.js
- allocation-plan.ts
- Portfolio Detail Page
- application/index.ts
- compilerOptions
- allocation-plan-management.ts
- portfolio-history-management.ts
- infra.ts
- root.ts
- Portfolio Section Navigation
- FormRowValueElements
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetComposedColumnInput` - 16 edges
3. `registerHandlebarsLangHelpers()` - 16 edges
4. `LogLevel` - 15 edges
5. `handlebars` - 14 edges
6. `bignumber.js` - 13 edges
7. `FractalPortfolioMultiChartDataSource` - 13 edges
8. `MultiChartDataSource` - 11 edges
9. `chart.js` - 10 edges
10. `compilerOptions` - 10 edges

## Surprising Connections (you probably didn't know these)
- `Asset Composed Columns Input` --references--> `AssetComposedColumnsInput`  [EXTRACTED]
  websrc/components/asset-composed-columns-input.html → websrc/components/asset-composed-columns-input.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPlannedAllocationMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/infra/chart/chart-types.ts
- `FractalPortfolioMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/infra/chart/chart-types.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (30 total, 5 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (80): APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay() (+72 more)

### Community 1 - "chart-utils.ts"
Cohesion: 0.06
Nodes (38): chartjs-plugin-datalabels, chroma-js, patternomaly, Application, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart() (+30 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (44): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper() (+36 more)

### Community 3 - "package.json"
Cohesion: 0.05
Nodes (44): alias, bignumber.js, dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js (+36 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+22 more)

### Community 5 - "asset-composed-columns-input.ts"
Cohesion: 0.13
Nodes (15): api, ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, AssetTickerAutocompleteState, autocompleteStates, closeAssetTickerAutocomplete(), ensureAssetDatalistLifecycleListener(), getAsset() (+7 more)

### Community 6 - "portfolio-chart.ts"
Cohesion: 0.16
Nodes (15): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+7 more)

### Community 7 - "dom/index.ts"
Cohesion: 0.13
Nodes (17): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+9 more)

### Community 8 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.19
Nodes (19): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX, PlannedAllocation, getAllocationHierarchySize() (+11 more)

### Community 9 - "allocation-plan-chart.ts"
Cohesion: 0.16
Nodes (15): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+7 more)

### Community 10 - "service/index.ts"
Cohesion: 0.26
Nodes (13): CompleteAllocationPlan, Portfolio, PortfolioDTO, AllocationDomainService, mapAllocationStructure(), mapToAllocationPlan(), mapToCompleteAllocationPlan(), mapToCompleteAllocationPlans() (+5 more)

### Community 11 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 12 - "devDependencies"
Cohesion: 0.13
Nodes (15): devDependencies, eslint, @eslint/js, globals, http-proxy-middleware, parcel, @parcel/transformer-raw, @parcel/transformer-sass (+7 more)

### Community 13 - "notifications.ts"
Cohesion: 0.15
Nodes (11): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR (+3 more)

### Community 14 - "bignumber.js"
Cohesion: 0.21
Nodes (10): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO, DivergenceAnalysis (+2 more)

### Community 15 - "allocation-plan.ts"
Cohesion: 0.20
Nodes (11): AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationPlan, AllocationPlanDTO, FractalHierarchicalAllocationPlan, PlannedAllocationDTO, SerializableCompleteAllocationPlan (+3 more)

### Community 16 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 17 - "application/index.ts"
Cohesion: 0.17
Nodes (10): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Ticker Autocomplete, allocationPlanManagement, AssetComposedColumnsInput, portfolioHistoryManagement, AssetPage, websrc_pages_index_assetpage (+2 more)

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - "allocation-plan-management.ts"
Cohesion: 0.20
Nodes (5): AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey()

### Community 20 - "portfolio-history-management.ts"
Cohesion: 0.25
Nodes (5): htmx.org, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), toInt()

### Community 21 - "infra.ts"
Cohesion: 0.28
Nodes (8): handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra, setupGlobalErrorHandler(), Router

### Community 22 - "root.ts"
Cohesion: 0.33
Nodes (3): notifications, AfterRequestEventDetail, HtmxInfra

### Community 23 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **117 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+112 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 163 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **5 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `bignumber.js` to `logger`, `chart-utils.ts`, `handlebars-lang.ts`, `package.json`, `allocation-plan.ts`, `allocation-plan-management.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.092) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `handlebars-lang.ts`, `binding-financial-input.ts`, `infra.ts`, `dom/index.ts`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `chart-utils.ts`, `package.json`, `asset.ts`, `notifications.ts`, `bignumber.js`, `allocation-plan-management.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _117 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05467856325783574 - nodes in this community are weakly interconnected._
- **Should `chart-utils.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05593220338983051 - nodes in this community are weakly interconnected._