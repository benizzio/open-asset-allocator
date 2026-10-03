# Graph Report - web-static  (2026-10-03)

## Corpus Check
- 77 files · ~32,878 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 3 file(s) not represented in the graph (top: (none) 2, .scss 1)

## Summary
- 631 nodes · 1418 edges · 36 communities (30 shown, 6 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3e518790`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- logger
- allocation-plan-chart.ts
- chart-utils.ts
- portfolio-chart.ts
- asset-composed-columns-input.ts
- asset.ts
- handlebars-lang.ts
- package.json
- allocation-plan-management.ts
- chart.ts
- devDependencies
- binding-financial-input.ts
- allocation-plan.ts
- notifications.ts
- Portfolio Detail Page
- portfolio-allocation.ts
- chart-contents.ts
- dependencies
- compilerOptions
- ChartDataSource
- portfolio-history-management.ts
- infra.ts
- Asset Search Autocomplete
- .proxyrc.js
- Portfolio Section Navigation
- Frontend Module Architecture
- Observation Editor
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- application/index.ts
- eslint.config.mjs
- scripts
- AssetSelectionState
- FormRowValueElements
- alias

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
- `FractalPortfolioMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/infra/chart/chart-types.ts
- `changeChartData()` --calls--> `changeChartDataOnDatasource()`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart.ts → websrc/infra/chart/chart-utils.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (36 total, 6 thin omitted)

### Community 0 - "logger"
Cohesion: 0.05
Nodes (82): navigo, api, APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes() (+74 more)

### Community 1 - "allocation-plan-chart.ts"
Cohesion: 0.17
Nodes (12): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+4 more)

### Community 2 - "chart-utils.ts"
Cohesion: 0.15
Nodes (17): bignumber.js, chartjs-plugin-datalabels, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions (+9 more)

### Community 3 - "portfolio-chart.ts"
Cohesion: 0.17
Nodes (14): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+6 more)

### Community 4 - "asset-composed-columns-input.ts"
Cohesion: 0.11
Nodes (24): ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, AssetSearchAutocompleteOption, AssetSearchAutocompleteState, autocompleteStates, clearAssetSearchOptionHighlight(), closeAssetSearchAutocomplete(), createAutocompleteState() (+16 more)

### Community 5 - "asset.ts"
Cohesion: 0.10
Nodes (31): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetPage, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft (+23 more)

### Community 6 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (41): handlebars, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+33 more)

### Community 7 - "package.json"
Cohesion: 0.12
Nodes (15): bootstrap-icons, bootswatch, chroma-js, eslint, htmx-ext-client-side-templates, htmx-ext-form-json, parcel, @parcel/transformer-raw (+7 more)

### Community 8 - "allocation-plan-management.ts"
Cohesion: 0.14
Nodes (9): htmx.org, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), SerializableCompleteAllocationPlan, Router (+1 more)

### Community 9 - "chart.ts"
Cohesion: 0.19
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+7 more)

### Community 10 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 11 - "binding-financial-input.ts"
Cohesion: 0.08
Nodes (34): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+26 more)

### Community 12 - "allocation-plan.ts"
Cohesion: 0.10
Nodes (41): AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL (+33 more)

### Community 13 - "notifications.ts"
Cohesion: 0.16
Nodes (10): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, CustomEventHandler, Notification, NotificationType, ERROR, INFO (+2 more)

### Community 14 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 15 - "portfolio-allocation.ts"
Cohesion: 0.21
Nodes (9): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO, DivergenceAnalysis, PotentialDivergence (+1 more)

### Community 16 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 17 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 18 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 19 - "ChartDataSource"
Cohesion: 0.15
Nodes (3): ChartDataSource, ChartDataSourceVisitor, SingleChartDataSource

### Community 20 - "portfolio-history-management.ts"
Cohesion: 0.24
Nodes (7): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), convertBooleanForToInt(), convertNumberForToInt(), reportToIntCoercion(), toInt()

### Community 21 - "infra.ts"
Cohesion: 0.17
Nodes (11): bootstrap, Application, handlebarsInfra, AfterRequestEventDetail, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+3 more)

### Community 22 - "Asset Search Autocomplete"
Cohesion: 0.53
Nodes (6): Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 23 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 24 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 30 - "application/index.ts"
Cohesion: 0.25
Nodes (7): Recursive Planned Allocation Rows, Asset Composed Columns Input, allocationPlanManagement, AssetComposedColumnsInput, portfolioHistoryManagement, websrc_pages_index_assetpage, websrc_pages_index_portfoliopage

### Community 31 - "eslint.config.mjs"
Cohesion: 0.33
Nodes (5): @eslint/js, eslint-plugin-sonarjs, globals, @stylistic/eslint-plugin, typescript-eslint

### Community 32 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 33 - "AssetSelectionState"
Cohesion: 0.40
Nodes (5): AssetSelectionState, EXISTING, NEW, PENDING, SEARCH

## Knowledge Gaps
- **124 isolated node(s):** `{ createProxyMiddleware }`, `fs`, `path`, `@eslint/js`, `@parcel/transformer-raw` (+119 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 169 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **6 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `chart-utils.ts` to `logger`, `handlebars-lang.ts`, `package.json`, `allocation-plan-management.ts`, `allocation-plan.ts`, `portfolio-allocation.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.091) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `package.json`, `allocation-plan-management.ts`, `chart.ts`, `notifications.ts`, `portfolio-allocation.ts`, `portfolio-history-management.ts`?**
  _High betweenness centrality (0.064) - this node is a cross-community bridge._
- **Why does `logger()` connect `logger` to `binding-financial-input.ts`, `infra.ts`, `handlebars-lang.ts`?**
  _High betweenness centrality (0.062) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `{ createProxyMiddleware }`, `fs`, `path` to the rest of the system?**
  _124 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `logger` be split into smaller, more focused modules?**
  _Cohesion score 0.05307017543859649 - nodes in this community are weakly interconnected._
- **Should `chart-utils.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.1471861471861472 - nodes in this community are weakly interconnected._