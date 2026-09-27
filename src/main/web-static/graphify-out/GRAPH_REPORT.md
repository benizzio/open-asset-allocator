# Graph Report - web-static  (2026-09-27)

## Corpus Check
- 76 files · ~31,174 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 609 nodes · 1375 edges · 24 communities (20 shown, 4 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `cc084ba1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- chart-utils.ts
- handlebars-lang.ts
- package.json
- asset.ts
- asset-composed-columns-input.ts
- allocation-plan-chart.ts
- htmx/index.ts
- allocation-plan.ts
- binding-financial-input.ts
- devDependencies
- notifications.ts
- Portfolio Detail Page
- application/index.ts
- compilerOptions
- allocation-plan-management.ts
- portfolio-history-management.ts
- infra.ts
- Portfolio Section Navigation
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
- `registerTransformResponseFunction()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/htmx/binding-htmx-transform-response.ts → websrc/infra/logging.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (24 total, 4 thin omitted)

### Community 0 - "logger"
Cohesion: 0.09
Nodes (51): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+43 more)

### Community 1 - "chart-utils.ts"
Cohesion: 0.05
Nodes (40): chartjs-plugin-datalabels, chroma-js, patternomaly, chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart() (+32 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.08
Nodes (43): handlebars, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+35 more)

### Community 3 - "package.json"
Cohesion: 0.05
Nodes (44): alias, bignumber.js, dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js (+36 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (30): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+22 more)

### Community 5 - "asset-composed-columns-input.ts"
Cohesion: 0.13
Nodes (17): ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, AssetTickerAutocompleteOption, AssetTickerAutocompleteState, autocompleteStates, clearAssetTickerOptionHighlight(), closeAssetTickerAutocomplete(), ensureAssetDatalistLifecycleListener() (+9 more)

### Community 6 - "allocation-plan-chart.ts"
Cohesion: 0.06
Nodes (44): bignumber.js, chart.js, allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback() (+36 more)

### Community 7 - "htmx/index.ts"
Cohesion: 0.10
Nodes (30): api, APIError, APIErrorResponse, bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse (+22 more)

### Community 8 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (40): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+32 more)

### Community 11 - "binding-financial-input.ts"
Cohesion: 0.12
Nodes (28): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+20 more)

### Community 12 - "devDependencies"
Cohesion: 0.13
Nodes (15): devDependencies, eslint, @eslint/js, globals, http-proxy-middleware, parcel, @parcel/transformer-raw, @parcel/transformer-sass (+7 more)

### Community 13 - "notifications.ts"
Cohesion: 0.16
Nodes (10): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 16 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 17 - "application/index.ts"
Cohesion: 0.18
Nodes (11): Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Ticker Autocomplete, Application, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), allocationPlanManagement, AssetComposedColumnsInput (+3 more)

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - "allocation-plan-management.ts"
Cohesion: 0.15
Nodes (8): htmx.org, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), Router, PortfolioPage

### Community 20 - "portfolio-history-management.ts"
Cohesion: 0.17
Nodes (6): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), FormRowValueElements, getNextPortfolioHistoryManagementIndex(), portfolioHistoryManagement, toInt()

### Community 21 - "infra.ts"
Cohesion: 0.17
Nodes (11): bootstrap, notifications, handlebarsInfra, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+3 more)

### Community 23 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **118 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+113 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 163 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan-chart.ts` to `logger`, `chart-utils.ts`, `handlebars-lang.ts`, `package.json`, `allocation-plan.ts`, `application/index.ts`, `allocation-plan-management.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.092) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `handlebars-lang.ts`, `binding-financial-input.ts`, `infra.ts`, `htmx/index.ts`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `chart-utils.ts`, `package.json`, `asset.ts`, `notifications.ts`, `application/index.ts`, `allocation-plan-management.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _118 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.08825248392752776 - nodes in this community are weakly interconnected._
- **Should `chart-utils.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.053763440860215055 - nodes in this community are weakly interconnected._