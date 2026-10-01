# Graph Report - web-static  (2026-10-02)

## Corpus Check
- 77 files · ~32,329 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 623 nodes · 1406 edges · 30 communities (26 shown, 4 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `c3ec3731`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- allocation-plan-management.ts
- allocation-plan.ts
- dom-utils.ts
- logger
- asset.ts
- handlebars-lang.ts
- asset-composed-columns-input.ts
- binding-financial-input.ts
- allocation-plan-chart.ts
- package.json
- chart-utils.ts
- notifications.ts
- portfolio-history-management.ts
- ChartDataSource
- devDependencies
- dependencies
- infra.ts
- Portfolio Detail Page
- compilerOptions
- .proxyrc.js
- Asset Ticker Autocomplete
- scripts
- Portfolio Section Navigation
- Frontend Module Architecture
- dom/index.ts
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- application/index.ts

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetComposedColumnInput` - 17 edges
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
- `FractalPlannedAllocationMultiChartDataSource` --references--> `FractalPlannedAllocation`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/domain/allocation-plan.ts
- `setHierarchicalIdFromParentRow()` --calls--> `toInt()`  [EXTRACTED]
  websrc/components/allocation-plan-management.ts → websrc/utils/lang.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (30 total, 4 thin omitted)

### Community 0 - "allocation-plan-management.ts"
Cohesion: 0.15
Nodes (8): htmx.org, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), Router, PortfolioPage

### Community 1 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (42): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+34 more)

### Community 2 - "dom-utils.ts"
Cohesion: 0.19
Nodes (10): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), addRemoveObserver(), contextDataCache, ensureSharedObserver() (+2 more)

### Community 3 - "logger"
Cohesion: 0.05
Nodes (82): navigo, api, APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes() (+74 more)

### Community 4 - "asset.ts"
Cohesion: 0.10
Nodes (31): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetPage, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft (+23 more)

### Community 5 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (36): arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper(), ifNotEqualsHelper(), ifNotNullishHelper() (+28 more)

### Community 6 - "asset-composed-columns-input.ts"
Cohesion: 0.12
Nodes (21): ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, AssetTickerAutocompleteOption, AssetTickerAutocompleteState, autocompleteStates, clearAssetTickerOptionHighlight(), closeAssetTickerAutocomplete(), ensureAssetDatalistLifecycleListener() (+13 more)

### Community 7 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 8 - "allocation-plan-chart.ts"
Cohesion: 0.07
Nodes (37): bignumber.js, chart.js, allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback() (+29 more)

### Community 9 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 10 - "chart-utils.ts"
Cohesion: 0.06
Nodes (47): chartjs-plugin-datalabels, chroma-js, handlebars, patternomaly, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis (+39 more)

### Community 11 - "notifications.ts"
Cohesion: 0.16
Nodes (10): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 12 - "portfolio-history-management.ts"
Cohesion: 0.17
Nodes (6): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), FormRowValueElements, getNextPortfolioHistoryManagementIndex(), AfterRequestEventDetail, toInt()

### Community 13 - "ChartDataSource"
Cohesion: 0.15
Nodes (3): ChartDataSource, ChartDataSourceVisitor, SingleChartDataSource

### Community 14 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 15 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 16 - "infra.ts"
Cohesion: 0.21
Nodes (9): notifications, handlebarsInfra, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra (+1 more)

### Community 17 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 20 - "Asset Ticker Autocomplete"
Cohesion: 0.53
Nodes (6): Unnamed Asset Search Field (filters the complete ticker-and-name label), Asset Ticker Autocomplete, Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 21 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 22 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 23 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 24 - "dom/index.ts"
Cohesion: 0.31
Nodes (8): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), maskNumberDecimalPlaces(), maskTagInput(), maskTickerInput(), DomInfra

### Community 29 - "application/index.ts"
Cohesion: 0.22
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Application, allocationPlanManagement, AssetComposedColumnsInput, portfolioHistoryManagement, websrc_pages_index_assetpage, websrc_pages_index_portfoliopage

## Knowledge Gaps
- **120 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+115 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 165 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan-chart.ts` to `allocation-plan-management.ts`, `allocation-plan.ts`, `logger`, `handlebars-lang.ts`, `package.json`, `chart-utils.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.092) - this node is a cross-community bridge._
- **Why does `handlebars` connect `chart-utils.ts` to `allocation-plan-management.ts`, `asset.ts`, `handlebars-lang.ts`, `package.json`, `notifications.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `infra.ts`, `dom-utils.ts`, `handlebars-lang.ts`, `binding-financial-input.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _120 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `allocation-plan.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.09090909090909091 - nodes in this community are weakly interconnected._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05307017543859649 - nodes in this community are weakly interconnected._