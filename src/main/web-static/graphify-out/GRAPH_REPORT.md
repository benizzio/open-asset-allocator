# Graph Report - web-static  (2026-10-20)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 755 nodes · 1720 edges · 31 communities (28 shown, 3 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 96 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f4578681`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- asset-composed-columns-input/constants.ts
- handlebars-lang.ts
- allocation-plan.ts
- chart-utils.ts
- asset.ts
- binding-dom-bootstrap-form.ts
- AssetRowController
- logger
- allocation-plan-chart.ts
- createPortfolioHistoryQuoteAction
- package.json
- portfolio-chart.ts
- Observation Editor
- MultiChartDataSource
- devDependencies
- allocation-plan-management.ts
- binding-htmx-trigger-on-route.ts
- dependencies
- chart.ts
- Portfolio Detail Page
- FractalPortfolioMultiChartDataSource
- Asset Search Autocomplete
- compilerOptions
- bignumber.js
- chart-contents.ts
- .proxyrc.js
- scripts
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
- `Portfolio Route Container` --references--> `Portfolio Detail Page`  [EXTRACTED]
  root.html → websrc/pages/portfolio.html
- `Portfolios Route Container` --references--> `Portfolio List Page`  [EXTRACTED]
  root.html → websrc/pages/portfolios.html
- `FractalPortfolioMultiChartDataSource` --references--> `AppliedAllocationHierarchyLevel`  [EXTRACTED]
  websrc/application/portfolio-chart/portfolio-chart-datasource.ts → websrc/application/portfolio-chart/portfolio-chart-model.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (31 total, 3 thin omitted)

### Community 0 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.05
Nodes (62): AssetSearchAutocompleteController, handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE (+54 more)

### Community 1 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (42): handlebars, handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate(), registerHandlebarsFormatHelper(), andHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+34 more)

### Community 2 - "allocation-plan.ts"
Cohesion: 0.10
Nodes (41): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+33 more)

### Community 3 - "chart-utils.ts"
Cohesion: 0.14
Nodes (16): chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction, ChartInteractions (+8 more)

### Community 4 - "asset.ts"
Cohesion: 0.05
Nodes (56): bootstrap, chartjs-plugin-datalabels, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, ExternalAsset, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants() (+48 more)

### Community 5 - "binding-dom-bootstrap-form.ts"
Cohesion: 0.39
Nodes (6): BootstrapClasses, BootstrapIconClasses, bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants()

### Community 7 - "logger"
Cohesion: 0.07
Nodes (60): api, APIError, APIErrorResponse, bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField() (+52 more)

### Community 8 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 9 - "createPortfolioHistoryQuoteAction"
Cohesion: 0.33
Nodes (14): createPortfolioHistoryQuoteAction(), applyLatestClosingQuote(), bindRow(), getQuoteRequestSnapshot(), handleAfterRequest(), handleBeforeSend(), handleRequestConfiguration(), isOwnedAfterRequest() (+6 more)

### Community 10 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 11 - "portfolio-chart.ts"
Cohesion: 0.21
Nodes (14): chart.js, getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+6 more)

### Community 12 - "Observation Editor"
Cohesion: 0.33
Nodes (7): Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Editor, Observation Header Field Focus Opens Accordion, Portfolio History Management, Portfolio History New Observation Header Field Group

### Community 13 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 14 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 16 - "allocation-plan-management.ts"
Cohesion: 0.05
Nodes (49): htmx.org, Application, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapForAllocationPlanManagement() (+41 more)

### Community 17 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.07
Nodes (53): navigo, addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants(), hideAllSiblings(), addRemoveObserver(), contextDataCache (+45 more)

### Community 18 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 20 - "chart.ts"
Cohesion: 0.20
Nodes (14): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+6 more)

### Community 21 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 22 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.24
Nodes (5): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), AllocationHierarchyLevel

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.21
Nodes (13): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Dated Allocation Plan Header, Recursive Planned Allocation Rows, Asset Composed Columns Input, Asset Search Autocomplete (+5 more)

### Community 25 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 28 - "bignumber.js"
Cohesion: 0.29
Nodes (7): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, DivergenceAnalysis, PotentialDivergence, Format

### Community 31 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 32 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

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
- **130 isolated node(s):** `AssignAtPathOptions`, `ToIntOptions`, `ReducedAllocation`, `PortfolioAllocationDTO`, `GeneralErrorHandler` (+125 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 174 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **3 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `bignumber.js` to `handlebars-lang.ts`, `allocation-plan.ts`, `chart-utils.ts`, `logger`, `package.json`, `portfolio-chart.ts`, `allocation-plan-management.ts`?**
  _High betweenness centrality (0.087) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `allocation-plan-management.ts` to `asset-composed-columns-input/constants.ts`, `asset.ts`, `logger`, `package.json`, `binding-htmx-trigger-on-route.ts`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `asset.ts`, `package.json`, `allocation-plan-management.ts`, `binding-htmx-trigger-on-route.ts`, `chart.ts`, `bignumber.js`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssignAtPathOptions`, `ToIntOptions`, `ReducedAllocation` to the rest of the system?**
  _130 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05198537095088819 - nodes in this community are weakly interconnected._
- **Should `handlebars-lang.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.08503401360544217 - nodes in this community are weakly interconnected._