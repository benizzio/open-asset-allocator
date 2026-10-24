# Graph Report - web-static  (2026-10-24)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 791 nodes · 1815 edges · 30 communities (28 shown, 2 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 110 edges (avg confidence: 0.86)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `2272351b`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- htmx/index.ts
- asset-composed-columns-input/constants.ts
- handlebars-lang.ts
- portfolio-history-management.ts
- allocation-plan.ts
- portfolio-history-management-observation-metadata/index.ts
- asset.ts
- allocation-plan-management.ts
- dom/index.ts
- AssetRowController
- chart-utils.ts
- package.json
- portfolio-chart.ts
- chart.ts
- allocation-plan-chart.ts
- MultiChartDataSource
- binding-financial-input.ts
- devDependencies
- dependencies
- Portfolio Detail Page
- FractalPortfolioMultiChartDataSource
- compilerOptions
- chart-contents.ts
- .proxyrc.js
- Asset Search Autocomplete
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
- `Allocation Plan Pencil Metadata Editor` --references--> `activateAllocationPlanEditor()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management-metadata/index.ts
- `selectAssetTicker()` --calls--> `AssetRowController`  [EXTRACTED]
  websrc/components/asset-composed-columns-input/index.ts → websrc/components/asset-composed-columns-input/asset-row-controller.ts
- `Observation Detail Lazy-Load and Retry State` --references--> `retryObservationDetails()`  [EXTRACTED]
  websrc/components/portfolio-history-management.html → websrc/components/portfolio-history-management-observation-details/index.ts
- `Recursive Planned Allocation Rows` --references--> `handleHierarchicalIdLevelChange()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts
- `Recursive Planned Allocation Rows` --references--> `handleRemovePlannedAllocationRow()`  [EXTRACTED]
  websrc/components/allocation-plan-management.html → websrc/components/allocation-plan-management.ts

## Import Cycles
- 3-file cycle: `websrc/infra/htmx/index.ts -> websrc/infra/routing/index.ts -> websrc/infra/routing/binding-htmx-trigger-on-route.ts -> websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]

## Communities (30 total, 2 thin omitted)

### Community 0 - "htmx/index.ts"
Cohesion: 0.05
Nodes (85): navigo, api, APIError, APIErrorResponse, bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes() (+77 more)

### Community 1 - "asset-composed-columns-input/constants.ts"
Cohesion: 0.05
Nodes (60): htmx.org, AssetSearchAutocompleteController, ARIA_ACTIVE_DESCENDANT_ATTRIBUTE, ARIA_CONTROLS_ATTRIBUTE, ARIA_EXPANDED_ATTRIBUTE, ARIA_SELECTED_ATTRIBUTE, ASSET_ACTION_BUTTON_IDENTITIES, ASSET_ACTION_BUTTON_SELECTOR (+52 more)

### Community 2 - "handlebars-lang.ts"
Cohesion: 0.06
Nodes (54): handlebars, addPlannedAllocationRow(), setHierarchicalIdFromParentRow(), getNextPortfolioHistoryManagementIndex(), domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), handlebarsFormatTimestampLocalDate() (+46 more)

### Community 3 - "portfolio-history-management.ts"
Cohesion: 0.07
Nodes (47): BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, notifications, handleAssetRowSelectionChange(), normalizePortfolioHistoryRow(), normalizePortfolioHistoryRowsWithin(), portfolioHistoryQuoteAction, synchronizeExternalAssetKeyGroup() (+39 more)

### Community 4 - "allocation-plan.ts"
Cohesion: 0.09
Nodes (43): bignumber.js, AllocationHierarchyLevel, AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO (+35 more)

### Community 5 - "portfolio-history-management-observation-metadata/index.ts"
Cohesion: 0.07
Nodes (44): Dated Observation Header, External Asset Quote Action, Observation Date Picker, Observation Detail Lazy-Load and Retry State, Observation Editor, Observation Header Field Focus Opens Accordion, Observation Timestamp and Tag Serialization, Portfolio History Management (+36 more)

### Community 6 - "asset.ts"
Cohesion: 0.10
Nodes (31): Asset, ExternalAsset, AssetBeforeSwapEvent, AssetRequestEvent, addExternalAsset(), clearSearch(), Draft, drafts (+23 more)

### Community 7 - "allocation-plan-management.ts"
Cohesion: 0.08
Nodes (29): Allocation Plan Date Picker, Allocation Plan Header Field Focus Opens Accordion, Allocation Plan Management, Allocation Plan New Entry Header Field Group, Allocation Plan Pencil Metadata Editor, Dated Allocation Plan Header, Recursive Planned Allocation Rows, Application (+21 more)

### Community 8 - "dom/index.ts"
Cohesion: 0.10
Nodes (24): bootstrap, AccordionFieldFocusState, bindAccordionFieldFocus(), bindAccordionFieldFocusInDescendants(), handleFieldBlur(), handleFieldFocus(), showPanelIfFocusRemains(), bindBootstrapValidationCleaning() (+16 more)

### Community 9 - "AssetRowController"
Cohesion: 0.15
Nodes (6): AssetRowController, handleAssetSearchKeydown(), navigateAssetSearchOptions(), registerAutocompleteInteractions(), submitAssetSearchLookup(), AutocompleteControllerActions

### Community 10 - "chart-utils.ts"
Cohesion: 0.13
Nodes (18): chartjs-plugin-datalabels, chroma-js, patternomaly, buildChartInteractions(), buildChartOptions(), getPieDoughnutChartOptions(), PIE_DOUGHNUT_CHART_OPTIONS, ChartInteraction (+10 more)

### Community 11 - "package.json"
Cohesion: 0.11
Nodes (20): alias, bignumber.js, bootstrap-icons, bootswatch, eslint, @eslint/js, eslint-plugin-sonarjs, globals (+12 more)

### Community 12 - "portfolio-chart.ts"
Cohesion: 0.24
Nodes (13): chart.js, getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty(), mapChartData(), ReducedAllocation, AppliedAllocationHierarchyLevel, MappedChartData (+5 more)

### Community 13 - "chart.ts"
Cohesion: 0.19
Nodes (15): chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), CHART_ATTRIBUTE, CHART_OPTIONS_JSON_ELEMENT_ID, ChartContent (+7 more)

### Community 14 - "allocation-plan-chart.ts"
Cohesion: 0.23
Nodes (11): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+3 more)

### Community 15 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 16 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 17 - "devDependencies"
Cohesion: 0.12
Nodes (16): devDependencies, eslint, @eslint/js, eslint-plugin-sonarjs, globals, http-proxy-middleware, parcel, @parcel/transformer-raw (+8 more)

### Community 18 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 19 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 20 - "FractalPortfolioMultiChartDataSource"
Cohesion: 0.26
Nodes (4): changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey()

### Community 21 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 22 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 23 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): { createProxyMiddleware }, fs, path, ref_fs, http-proxy-middleware, ref_path

### Community 24 - "Asset Search Autocomplete"
Cohesion: 0.43
Nodes (7): Asset Composed Columns Input, Asset Search Autocomplete, Unnamed Asset Search Field (filters the complete ticker-and-name label), Named Committed Ticker Control (submits only the canonical ticker), Asset Search Datalist, Asset Search Option Label (ticker - asset name), Canonical Asset Ticker (datalist option value)

### Community 25 - "scripts"
Cohesion: 0.40
Nodes (5): scripts, build, clean, dev, lint

### Community 26 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 27 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **137 isolated node(s):** `APIError`, `EventDetail`, `Level`, `Levels`, `AssetSearchAutocompleteOption` (+132 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 177 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **2 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `bignumber.js` connect `allocation-plan.ts` to `htmx/index.ts`, `handlebars-lang.ts`, `portfolio-history-management.ts`, `allocation-plan-management.ts`, `chart-utils.ts`, `package.json`, `portfolio-chart.ts`?**
  _High betweenness centrality (0.087) - this node is a cross-community bridge._
- **Why does `htmx.org` connect `asset-composed-columns-input/constants.ts` to `htmx/index.ts`, `portfolio-history-management.ts`, `allocation-plan.ts`, `portfolio-history-management-observation-metadata/index.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`?**
  _High betweenness centrality (0.065) - this node is a cross-community bridge._
- **Why does `handlebars` connect `handlebars-lang.ts` to `portfolio-history-management.ts`, `asset.ts`, `allocation-plan-management.ts`, `package.json`, `chart.ts`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **Are the 16 inferred relationships involving `registerHandlebarsLangHelpers()` (e.g. with `andHelper()` and `arrayHelper()`) actually correct?**
  _`registerHandlebarsLangHelpers()` has 16 INFERRED edges - model-reasoned connections that need verification._
- **What connects `APIError`, `EventDetail`, `Level` to the rest of the system?**
  _137 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `htmx/index.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.05071119356833643 - nodes in this community are weakly interconnected._
- **Should `asset-composed-columns-input/constants.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.053703703703703705 - nodes in this community are weakly interconnected._