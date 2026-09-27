# Graph Report - web-static  (2026-09-26)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 594 nodes · 1335 edges · 31 communities (24 shown, 7 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 57 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `da880bc5`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Community 0
- Community 1
- Community 2
- Community 3
- Community 4
- Community 5
- Community 6
- Community 7
- Community 8
- Community 9
- Community 10
- Community 11
- Community 12
- Community 13
- Community 14
- Community 15
- Community 16
- Community 17
- Community 18
- Community 19
- Community 20
- Community 21
- Community 22
- Community 23
- Community 24
- Community 25
- Community 26
- Community 27
- Community 28
- Community 29
- Community 30

## God Nodes (most connected - your core abstractions)
1. `logger()` - 43 edges
2. `registerHandlebarsLangHelpers()` - 16 edges
3. `LogLevel` - 15 edges
4. `handlebars` - 14 edges
5. `FractalPortfolioMultiChartDataSource` - 13 edges
6. `bignumber.js` - 13 edges
7. `AssetComposedColumnInput` - 12 edges
8. `MultiChartDataSource` - 11 edges
9. `NotificationType` - 10 edges
10. `bindFinancialInput()` - 10 edges

## Surprising Connections (you probably didn't know these)
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPortfolioMultiChartDataSource` --inherits--> `MultiChartDataSource`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/infra/chart/chart-types.ts
- `bindExclusiveDisplayInDescendants()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/dom/binding-dom-exclusive-display.ts → websrc/infra/logging.ts
- `bindFinancialInputElements()` --calls--> `logger()`  [EXTRACTED]
  websrc/infra/dom/binding-financial-input.ts → websrc/infra/logging.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (31 total, 7 thin omitted)

### Community 0 - "Community 0"
Cohesion: 0.05
Nodes (80): APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay() (+72 more)

### Community 1 - "Community 1"
Cohesion: 0.09
Nodes (42): bignumber.js, AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO (+34 more)

### Community 2 - "Community 2"
Cohesion: 0.09
Nodes (41): handlebars, DomUtils, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper() (+33 more)

### Community 3 - "Community 3"
Cohesion: 0.05
Nodes (44): alias, bignumber.js, dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js (+36 more)

### Community 4 - "Community 4"
Cohesion: 0.10
Nodes (31): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+23 more)

### Community 5 - "Community 5"
Cohesion: 0.09
Nodes (16): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+8 more)

### Community 6 - "Community 6"
Cohesion: 0.13
Nodes (17): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+9 more)

### Community 7 - "Community 7"
Cohesion: 0.13
Nodes (18): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+10 more)

### Community 8 - "Community 8"
Cohesion: 0.20
Nodes (9): chart.js, FractalPortfolioMultiChartDataSource, generateDataKey(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+1 more)

### Community 9 - "Community 9"
Cohesion: 0.14
Nodes (10): htmx.org, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure(), mapPlannedAllocationFormEntriesPerHierarchicalKey(), SerializableCompleteAllocationPlan, Portfolio (+2 more)

### Community 10 - "Community 10"
Cohesion: 0.21
Nodes (13): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, LocalChartOptions (+5 more)

### Community 11 - "Community 11"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 12 - "Community 12"
Cohesion: 0.13
Nodes (15): devDependencies, eslint, @eslint/js, globals, http-proxy-middleware, parcel, @parcel/transformer-raw, @parcel/transformer-sass (+7 more)

### Community 13 - "Community 13"
Cohesion: 0.15
Nodes (11): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, CustomEventHandler, Notification, NotificationType, ERROR (+3 more)

### Community 14 - "Community 14"
Cohesion: 0.20
Nodes (12): toChartContent(), toUnidimensionalMultiChartContent(), changeChartData(), chartDataSelectionEventHandler(), getChartContent(), interactionObserverCallback(), portfolioChart, DomainService (+4 more)

### Community 15 - "Community 15"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 17 - "Community 17"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 18 - "Community 18"
Cohesion: 0.23
Nodes (8): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO, DivergenceAnalysis, PotentialDivergence

### Community 19 - "Community 19"
Cohesion: 0.22
Nodes (7): Application, allocationPlanManagement, notifications, AfterRequestEventDetail, AssetPage, websrc_pages_index_assetpage, websrc_pages_index_portfoliopage

### Community 20 - "Community 20"
Cohesion: 0.22
Nodes (7): addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), portfolioHistoryManagement, convertNumberForToInt(), reportToIntCoercion(), toInt()

### Community 21 - "Community 21"
Cohesion: 0.28
Nodes (5): api, ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnsInput, BootstrapClasses, BootstrapIconClasses

### Community 22 - "Community 22"
Cohesion: 0.28
Nodes (8): handlebarsInfra, HtmxInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler, handleError(), Infra, setupGlobalErrorHandler()

### Community 23 - "Community 23"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 25 - "Community 25"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **115 isolated node(s):** `Level`, `Levels`, `APIError`, `EventDetail`, `AllocationHierarchyLevelDTO` (+110 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 161 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **7 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `Community 1` to `Community 0`, `Community 2`, `Community 3`, `Community 7`, `Community 9`, `Community 18`, `Community 20`?**
  _High betweenness centrality (0.094) - this node is a cross-community bridge._
- **Why does `logger()` connect `Community 0` to `Community 2`, `Community 11`, `Community 6`, `Community 22`?**
  _High betweenness centrality (0.068) - this node is a cross-community bridge._
- **Why does `handlebars` connect `Community 2` to `Community 3`, `Community 4`, `Community 9`, `Community 10`, `Community 13`, `Community 18`, `Community 20`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Are the 14 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `arrayHelper()` and `comparatorHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 14 INFERRED edges - model-reasoned connections that need verification._
- **What connects `Level`, `Levels`, `APIError` to the rest of the system?**
  _115 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Community 0` be split into smaller, more focused modules?**
  _Cohesion score 0.05467856325783574 - nodes in this community are weakly interconnected._
- **Should `Community 1` be split into smaller, more focused modules?**
  _Cohesion score 0.09350649350649351 - nodes in this community are weakly interconnected._