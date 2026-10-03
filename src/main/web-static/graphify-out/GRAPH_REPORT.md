# Graph Report - web-static  (2026-10-03)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 623 nodes · 1406 edges · 30 communities (26 shown, 4 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `99fdbd44`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan-chart.ts
- chart-utils.ts
- allocation-plan.ts
- asset-composed-columns-input.ts
- asset.ts
- handlebars-lang.ts
- package.json
- allocation-plan-management.ts
- binding-financial-input.ts
- devDependencies
- dom-utils.ts
- dependencies
- notifications.ts
- Portfolio Detail Page
- portfolio-history-management.ts
- infra.ts
- ChartDataSource
- compilerOptions
- dom/index.ts
- application/index.ts
- .proxyrc.js
- Asset Search Autocomplete
- scripts
- Portfolio Section Navigation
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `AssetComposedColumnInput` - 17 edges
3. `registerHandlebarsLangHelpers()` - 16 edges
4. `LogLevel` - 15 edges
5. `handlebars` - 14 edges
6. `FractalPortfolioMultiChartDataSource` - 13 edges
7. `bignumber.js` - 13 edges
8. `MultiChartDataSource` - 11 edges
9. `NotificationType` - 10 edges
10. `bindFinancialInput()` - 10 edges

## Surprising Connections (you probably didn't know these)
- `Asset Composed Columns Input` --references--> `AssetComposedColumnsInput`  [EXTRACTED]
  websrc/components/asset-composed-columns-input.html → websrc/components/asset-composed-columns-input.ts
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPlannedAllocationMultiChartDataSource` --references--> `FractalPlannedAllocation`  [EXTRACTED]
  websrc/application/allocation-plan-chart.ts → websrc/domain/allocation-plan.ts
- `registerTransformResponseFunction()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/htmx/binding-htmx-transform-response.ts → websrc/infra/logging.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (30 total, 4 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (82): navigo, api, APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes() (+74 more)

### Community 1 - "allocation-plan-chart.ts"
Cohesion: 0.07
Nodes (37): bignumber.js, chart.js, allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback() (+29 more)

### Community 2 - "chart-utils.ts"
Cohesion: 0.06
Nodes (47): chartjs-plugin-datalabels, chroma-js, handlebars, patternomaly, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis (+39 more)

### Community 3 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (42): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+34 more)

### Community 4 - "asset-composed-columns-input.ts"
Cohesion: 0.12
Nodes (21): ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, AssetSearchAutocompleteOption, AssetSearchAutocompleteState, autocompleteStates, clearAssetSearchOptionHighlight(), closeAssetSearchAutocomplete(), ensureAssetDatalistLifecycleListener() (+13 more)

### Community 5 - "asset.ts"
Cohesion: 0.10
Nodes (31): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetPage, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft (+23 more)

### Community 6 - "handlebars-lang.ts"
Cohesion: 0.10
Nodes (36): arrayHelper(), comparatorHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper(), ifNotEqualsHelper(), ifNotNullishHelper() (+28 more)

### Community 7 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 8 - "allocation-plan-management.ts"
Cohesion: 0.15
Nodes (8): htmx.org, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), Router, PortfolioPage

### Community 9 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 10 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 11 - "dom-utils.ts"
Cohesion: 0.19
Nodes (10): addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), addRemoveObserver(), contextDataCache, ensureSharedObserver() (+2 more)

### Community 12 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 13 - "notifications.ts"
Cohesion: 0.16
Nodes (10): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 14 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 15 - "portfolio-history-management.ts"
Cohesion: 0.17
Nodes (6): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), FormRowValueElements, getNextPortfolioHistoryManagementIndex(), AfterRequestEventDetail, toInt()

### Community 16 - "infra.ts"
Cohesion: 0.21
Nodes (9): notifications, handlebarsInfra, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra (+1 more)

### Community 17 - "ChartDataSource"
Cohesion: 0.15
Nodes (3): ChartDataSource, ChartDataSourceVisitor, SingleChartDataSource

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - "dom/index.ts"
Cohesion: 0.31
Nodes (8): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), maskNumberDecimalPlaces(), maskTagInput(), maskTickerInput(), DomInfra

### Community 20 - "application/index.ts"
Cohesion: 0.22
Nodes (8): Recursive Planned Allocation Rows, Asset Composed Columns Input, Application, allocationPlanManagement, AssetComposedColumnsInput, portfolioHistoryManagement, websrc_pages_index_assetpage, websrc_pages_index_portfoliopage

### Community 21 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 22 - "Asset Search Autocomplete"
Cohesion: 0.53
Nodes (6): Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 23 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 24 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **120 isolated node(s):** `APIError`, `EventDetail`, `Level`, `Levels`, `ReducedAllocation` (+115 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 165 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **4 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan-chart.ts` to `logger`, `chart-utils.ts`, `allocation-plan.ts`, `handlebars-lang.ts`, `package.json`, `allocation-plan-management.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.092) - this node is a cross-community bridge._
- **Why does `handlebars` connect `chart-utils.ts` to `asset.ts`, `handlebars-lang.ts`, `package.json`, `allocation-plan-management.ts`, `notifications.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `infra.ts`, `binding-financial-input.ts`, `dom-utils.ts`, `handlebars-lang.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `APIError`, `EventDetail`, `Level` to the rest of the system?**
  _120 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05307017543859649 - nodes in this community are weakly interconnected._
- **Should `allocation-plan-chart.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.0706605222734255 - nodes in this community are weakly interconnected._