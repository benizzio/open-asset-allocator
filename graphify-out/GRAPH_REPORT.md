# Graph Report - open-asset-allocator  (2026-09-20)

## Corpus Check
- 319 files · ~296,278 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 24 file(s) not represented in the graph (top: (none) 16, .csv 2, .dbm 1)

## Summary
- 2240 nodes · 5213 edges · 153 communities (110 shown, 43 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 310 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Community Hubs (Navigation)
- Allocation Plan Domain
- Database Query Infrastructure
- Research Report Generator
- Divergence Analysis Service
- Allocation Plan E2E
- Project Research Documentation
- Frontend Application Controllers
- Portfolio History E2E
- Go Package Architecture
- Allocation Plan Integration
- Go Utility Tests
- Database Repository Adapters
- Gin Validation Errors
- Chart Utilities
- Application Component Assembly
- Portfolio Allocation Integration
- Portfolio Charting
- Allocation Map E2E
- Asset Repository
- Portfolio Integration Tests
- Domain Unit Tests
- Portfolio Allocation Domain
- Yahoo Finance Mocking
- Allocation Plan Models
- Concurrent Utilities
- Asset Integration Tests
- Frontend Routing
- Asset Input Component
- Allocation Plan Charts
- Allocation Mapping
- Portfolio REST Mapping
- Ordered Map Iterator
- Frontend Infrastructure
- Allocation Plan Repository
- Frontend Package Configuration
- E2E Package Configuration
- Chart Interaction Tests
- Allocation Divergence UI
- E2E Shell Automation
- Playwright Version Validation
- Portfolio Visualization E2E
- Portfolio Holdings UI
- Asset REST Mapping
- Frontend Event Bindings
- HTTP Client Utilities
- DOM Routing Utilities
- Handlebars Language Helpers
- Portfolio Editing E2E
- Struct Pointer Utilities
- Playwright Test Fixtures
- Chart Data Sources
- Financial Input Binding
- Yahoo Finance Client
- Gin Server
- DOM Form Bindings
- Frontend Logging Bindings
- Allocation Plan UI
- Validation Field Errors
- Frontend Dependencies
- Portfolio Snapshot UI
- Application Configuration
- Allocation Plan REST
- Frontend Runtime Dependencies
- Portfolio Page Routing
- E2E TypeScript Configuration
- JSON Reflection Utilities
- Frontend TypeScript Configuration
- Allocation Plan Detail
- Portfolio History Detail
- Deep Validation
- Handlebars Core Helpers
- Docker Service Topology
- E2E Database Utilities
- Graphify Command Pipeline
- Portfolio Selection UI
- JSON Assertion Utilities
- Application Lifecycle
- Struct Completion Utilities
- Language Comparison Helpers
- HTMX Response Transformation
- Allocation Plan Persistence
- Divergence Analysis Options
- Validation Error Formatting
- Validation Error Builder
- Generic Map Tree
- Handlebars Utility Helpers
- Portfolio REST API
- Asset Persistence
- Slice Iterator
- External Asset Quotes
- Rebalancing Workflow
- Portfolio History Snapshot
- REST Route Building
- Database Cleanup Utilities
- Yahoo Finance Models
- Parcel Proxy Configuration
- Reflection Validation Helpers
- Custom Slice Formatting
- Portfolio Analysis Services
- Incremental Graph Updates
- Database Destruction Scripts
- Nullable SQL Strings
- macOS Provisioning
- Frontend Build Scripts
- Portfolio Analysis Helpers
- Portfolio Section Navigation
- Debug Startup Script
- Database Stop Script
- Node Version Validation
- Graph Query System
- Semantic Extraction Rules
- Go Test Workflows
- Development Build Script
- Development Startup Script
- DuckDB CLI Script
- Arch Linux Provisioning
- Frontend Module Architecture
- Database Start Script
- Cross Repository Exports
- E2E Workflow Diagnostics
- Build Shell Script
- Go Dependency Graph
- Custom Slice Table
- Database Migration Script
- REST Error Response
- DuckDB Build Script
- External Integration CI
- Go Testing Standards
- Database Initialization Script
- Recursive Allocation Inputs
- Portfolio Observation Editor
- Test Shell Script
- Folder Watcher
- Post Commit Graph
- Query Result Memory
- Media Transcription
- CodeRabbit Configuration
- Go Lint Workflow
- Domain Literacy
- Project Fit Scoring
- Go Module Root
- Development Docker Stack
- Local E2E Stack
- DuckDB CLI Service
- Flyway Service
- Flyway Standards
- Go Lint Configuration
- Hierarchical Divergence
- Allocation Plan Management
- Toast Notifications

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 75 edges
2. `logger()` - 43 edges
3. `PortfolioAllocation` - 32 edges
4. `PropagateAsAppErrorWithNewMessage()` - 31 edges
5. `NewMapTree()` - 27 edges
6. `Asset` - 26 edges
7. `BuildCleanupFunctionBuilder()` - 26 edges
8. `IsZeroValue()` - 24 edges
9. `AllocationPlan` - 23 edges
10. `PlannedAllocation` - 22 edges

## Surprising Connections (you probably didn't know these)
- `Fractal Portfolio Hierarchy` --semantically_similar_to--> `Customizable Fractal Hierarchy`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Portfolio History` --semantically_similar_to--> `Historical Portfolio Snapshots`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Allocation Plan` --semantically_similar_to--> `Allocation Planning`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Divergence Analysis` --semantically_similar_to--> `Divergence Analysis`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Convergence Planning` --semantically_similar_to--> `Convergence Analysis and Planning`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md

## Import Cycles
- 3-file cycle: `src/main/web-static/websrc/infra/htmx/index.ts -> src/main/web-static/websrc/infra/routing/index.ts -> src/main/web-static/websrc/infra/routing/binding-htmx-trigger-on-route.ts -> src/main/web-static/websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Graphify Semantic Extraction Integrity Contract** — _agents_skills_graphify_references_extraction_spec_semantic_extraction_contract, _agents_skills_graphify_references_extraction_spec_confidence_rubric, _agents_skills_graphify_references_extraction_spec_deterministic_node_identity [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]
- **Containerized E2E Execution Topology** — src_main_docker_docker_compose_e2e_e2e_database, src_main_docker_docker_compose_e2e_e2e_migration_engine, src_main_docker_docker_compose_e2e_playwright_runner, src_main_docker_docker_compose_e2e_ci_immutable_monolith, _github_workflows_e2e_e2e_tests_workflow [INFERRED 0.95]
- **Fourteen Evaluated E2E Candidates** — docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_playwright_test, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_cypress, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_webdriverio, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_selenium_webdriver_javascript, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_nightwatchjs, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_testcafe, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_puppeteer_vitest_jest, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_rod_go_testing, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_chromedp_go_testing, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_codeceptjs, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_testplane, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_cucumber_playwright_bdd, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_robot_framework_browser, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_grafana_k6_browser [EXTRACTED 1.00]
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Detail Navigation** — docs_images_allocation_plan_detail_portfolio_tab, docs_images_allocation_plan_detail_allocation_plan_tab, docs_images_allocation_plan_detail_allocation_map_tab [EXTRACTED 1.00]
- **60/40 Asset Allocation** — docs_images_allocation_plan_detail_60_40_portfolio_classic, docs_images_allocation_plan_detail_asset_allocation_donut_chart, docs_images_allocation_plan_detail_bonds_60_percent_allocation, docs_images_allocation_plan_detail_stocks_40_percent_allocation [EXTRACTED 1.00]
- **Top-Level 60/40 Allocation** — docs_images_allocation_plan_management_60_40_portfolio_classic_example_20260210_230012, docs_images_allocation_plan_management_bonds, docs_images_allocation_plan_management_stocks [EXTRACTED 1.00]
- **BONDS Slice Composition** — docs_images_allocation_plan_management_bonds, docs_images_allocation_plan_management_arca_bil, docs_images_allocation_plan_management_nasdaqgm_ief, docs_images_allocation_plan_management_nasdaqgm_tlt, docs_images_allocation_plan_management_arca_stip [EXTRACTED 1.00]
- **STOCKS Slice Composition** — docs_images_allocation_plan_management_stocks, docs_images_allocation_plan_management_nasdaqgm_shv, docs_images_allocation_plan_management_arca_spy, docs_images_allocation_plan_management_arca_ewz [EXTRACTED 1.00]
- **Asset Allocation Monitoring and Rebalancing Cycle** — docs_images_asset_allocation_flow_allocation_plan_creation_or_modification, docs_images_asset_allocation_flow_out_of_balance_asset_verification, docs_images_asset_allocation_flow_resource_reallocation, docs_images_asset_allocation_flow_planned_interval_wait, docs_images_asset_allocation_flow_market_fluctuation_check, docs_images_asset_allocation_flow_scenario_change_check [EXTRACTED 1.00]
- **Portfolio Allocation Map Interface Composition** — docs_images_divergence_analysis_portfolio_workflow_tabs, docs_images_divergence_analysis_monthly_snapshot_accordion, docs_images_divergence_analysis_allocation_plan_selector, docs_images_divergence_analysis_actual_vs_planned_market_value_comparison, docs_images_divergence_analysis_hierarchical_asset_allocation_breakdown, docs_images_divergence_analysis_directional_divergence_bars [EXTRACTED 1.00]
- **Hierarchical Allocation Divergence Rows** — docs_images_divergence_analysis_bonds_underweight_9_000_15, docs_images_divergence_analysis_stocks_overweight_9_000_15, docs_images_divergence_analysis_bil_underweight_800_2_96, docs_images_divergence_analysis_stip_overweight_5_300_19_63, docs_images_divergence_analysis_ief_underweight_2_100_7_78, docs_images_divergence_analysis_tlt_underweight_2_400_8_89, docs_images_divergence_analysis_shv_underweight_7_500_22_73, docs_images_divergence_analysis_spy_overweight_7_650_23_18, docs_images_divergence_analysis_ewz_underweight_150_0_45 [EXTRACTED 1.00]
- **Portfolio Detail Navigation** — docs_images_portfolio_history_detail1_active_portfolio_tab, docs_images_portfolio_history_detail1_allocation_plan_tab, docs_images_portfolio_history_detail1_allocation_map_tab [EXTRACTED 1.00]
- **February 2026 Class Allocation Summary** — docs_images_portfolio_history_detail1_expanded_202602_history_panel, docs_images_portfolio_history_detail1_total_market_value_60000, docs_images_portfolio_history_detail1_class_level_asset_allocation_chart, docs_images_portfolio_history_detail1_bonds_45_percent, docs_images_portfolio_history_detail1_stocks_55_percent [EXTRACTED 1.00]
- **Portfolio Detail Navigation Tabs** — docs_images_portfolio_history_detail2_portfolio_tab, docs_images_portfolio_history_detail2_allocation_plan_tab, docs_images_portfolio_history_detail2_allocation_map_tab [EXTRACTED 1.00]
- **Asset Allocation Composition** — docs_images_portfolio_history_detail2_asset_allocation_donut_chart, docs_images_portfolio_history_detail2_nasdaqgm_shv_cash_reserve_allocation, docs_images_portfolio_history_detail2_arca_spy_allocation, docs_images_portfolio_history_detail2_arca_ewz_allocation [EXTRACTED 1.00]
- **Portfolio Observation Holding Fields** — docs_images_portfolio_history_management_asset_identity, docs_images_portfolio_history_management_asset_classification, docs_images_portfolio_history_management_cash_reserve_designation, docs_images_portfolio_history_management_position_quantity, docs_images_portfolio_history_management_market_price, docs_images_portfolio_history_management_total_market_value [EXTRACTED 1.00]
- **Monthly Portfolio Valuation Breakdown** — docs_images_portfolio_history_monthly_portfolio_snapshot_202602, docs_images_portfolio_history_total_market_value_60000, docs_images_portfolio_history_asset_class_allocation, docs_images_portfolio_history_bonds_allocation_45_percent, docs_images_portfolio_history_stocks_allocation_55_percent [EXTRACTED 1.00]
- **Portfolio Choice Grid** — docs_images_portfolio_selection_global_all_assets_portfolio, docs_images_portfolio_selection_my_portfolio_example, docs_images_portfolio_selection_my_portfolio_example_20260210_230012, docs_images_portfolio_selection_new_portfolio_action [EXTRACTED 1.00]

## Communities (153 total, 43 thin omitted)

### Community 0 - "Allocation Plan Domain"
Cohesion: 0.06
Nodes (42): DomainValidationError, UniqueConstraintViolationError, AllocationHierarchyLevelDTS, AllocationStructureDTS, allocationPlanValidationData, levelSliceSizeValidationData, mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels() (+34 more)

### Community 1 - "Database Query Infrastructure"
Cohesion: 0.05
Nodes (30): HierarchicalId, go_pkg_github_com_lib_pq, database/sql/driver.Value, database/sql.Row, database/sql.Rows, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows (+22 more)

### Community 2 - "Research Report Generator"
Cohesion: 0.06
Nodes (57): Any, _anchor(), _base_result_name(), _build_output_mapping(), _category_aliases(), _collect_extra_fields(), visit(), _compact_summary() (+49 more)

### Community 3 - "Divergence Analysis Service"
Cohesion: 0.09
Nodes (38): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, potentialDivergencesPerHierarchicalId, context.Context, SliceIterator, PlannedAllocation, buildAllocationIterationContext() (+30 more)

### Community 4 - "Allocation Plan E2E"
Cohesion: 0.06
Nodes (43): addAssetAllocationRow(), addClassAllocationRow(), DEFAULT_ALLOCATION_STRUCTURE, expectAllocationPlan(), expectAllocationPlanManagement(), expectAllocationPlanManagementForm(), expectDraftAllocationRow(), expectDraftAllocationRows() (+35 more)

### Community 5 - "Project Research Documentation"
Cohesion: 0.05
Nodes (49): Root AGENTS Instructions for Copilot, Dependabot Dependency Updates, Renovate Runtime Coordination, Atomic Persistence-Verified E2E Testing, Open Asset Allocator Architecture, Runtime Version Coordination, Portfolio History Form TODOs, Allocation Planning (+41 more)

### Community 6 - "Frontend Application Controllers"
Cohesion: 0.07
Nodes (29): htmx.org, Application, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry, FormRowHierarchicalStructure, getHierarchicalFieldForValidation(), mapFormRowHierarchicalStructure() (+21 more)

### Community 7 - "Portfolio History E2E"
Cohesion: 0.06
Nodes (34): installCanvasTextRecorder(), DEFAULT_ALLOCATION_STRUCTURE, expectEditableObservationRows(), ExpectedObservationRow, expectExistingAsset(), expectNewObservationRows(), expectPersistedAllocations(), expectPersistedModifiedPortfolioHistory() (+26 more)

### Community 8 - "Go Package Architecture"
Cohesion: 0.12
Nodes (24): go_pkg_context, go_pkg_database_sql_driver, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation (+16 more)

### Community 9 - "Allocation Plan Integration"
Cohesion: 0.13
Nodes (41): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestGetAllocationPlans(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds() (+33 more)

### Community 10 - "Go Utility Tests"
Cohesion: 0.11
Nodes (39): testing.T, demoStringer, TestGetDivergenceAnalysisOptions(), TestGetDivergenceAnalysisV2(), TestGetDivergenceAnalysisV2FullDivergence(), TestGetDivergenceAnalysisV2WhenHistoryHighestLevelHasExtraAllocation(), TestGetDivergenceAnalysisV2WhenPlanHighestLevelHasExtraAllocation(), TestGetDivergenceAnalysisV2WhenPlanLowestLevelHasDifferentRanges() (+31 more)

### Community 11 - "Database Repository Adapters"
Cohesion: 0.09
Nodes (17): database/sql.DB, database/sql.Result, database/sql.Stmt, github.com/go-ozzo/ozzo-dbx.DB, contextKey, TransactionalContext, AllocationPlanRDBMSRepository, BuildAppError() (+9 more)

### Community 12 - "Gin Validation Errors"
Cohesion: 0.13
Nodes (18): github.com/gin-gonic/gin.Context, AssetRESTController, PortfolioAllocationRESTController, PortfolioRESTController, HandleAPIError(), handleDomainError(), handleInfrastructureError(), SendDataNotFoundResponse() (+10 more)

### Community 13 - "Chart Utilities"
Cohesion: 0.10
Nodes (32): chartjs-plugin-datalabels, chroma-js, patternomaly, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions() (+24 more)

### Community 14 - "Application Component Assembly"
Cohesion: 0.08
Nodes (25): AllocationPlanRESTController, AssetIntegrationServicesPerSource, BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildPortfolioAllocationRESTController(), BuildPortfolioRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService (+17 more)

### Community 15 - "Portfolio Allocation Integration"
Cohesion: 0.08
Nodes (30): go_pkg_database_sql, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go (+22 more)

### Community 16 - "Portfolio Charting"
Cohesion: 0.11
Nodes (23): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+15 more)

### Community 17 - "Allocation Map E2E"
Cohesion: 0.08
Nodes (33): ASSET_DATA, DatabaseSnapshot, DEFAULT_ALLOCATION_STRUCTURE, expandRootAndAssertChildren(), expectAllocationMapShell(), expectAnalysisTable(), ExpectedDivergenceNode, expectNodeRow() (+25 more)

### Community 18 - "Asset Repository"
Cohesion: 0.12
Nodes (19): AllocationRDBMSRepository, AssetRDBMSRepository, PortfolioRDBMSRepository, S, Asset, BuildAllocationPlanRepository(), BuildAllocationRepository(), buildAssetInsertValue() (+11 more)

### Community 19 - "Portfolio Integration Tests"
Cohesion: 0.12
Nodes (30): go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, postAssetForValidationFailure(), TestPostAsset(), TestPostAssetFailureWithMissingRequiredFields(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), TestPostAssetWithoutExternalData(), postAsset() (+22 more)

### Community 20 - "Domain Unit Tests"
Cohesion: 0.09
Nodes (22): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_net_http_httptest, go_pkg_testing, TestHierarchicalIdIsTopLevel_Empty() (+14 more)

### Community 21 - "Portfolio Allocation Domain"
Cohesion: 0.15
Nodes (13): portfolioAllocationJoinedRowDTS, PortfolioAllocationRDBMSRepository, buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), BuildPortfolioAllocationRepository(), preparePortfolioAllocationTempInserts(), Asset (+5 more)

### Community 22 - "Yahoo Finance Mocking"
Cohesion: 0.10
Nodes (20): go_pkg_github_com_nhatthm_httpmock, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry() (+12 more)

### Community 23 - "Allocation Plan Models"
Cohesion: 0.18
Nodes (19): bignumber.js, AllocationPlan, AllocationPlanDTO, CompleteAllocationPlan, PlannedAllocation, PlannedAllocationDTO, SerializableFractalHierarchicalAllocationPlan, SerializableFractalPlannedAllocation (+11 more)

### Community 24 - "Concurrent Utilities"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, SliceFromItem, SliceFromItemWithContext (+17 more)

### Community 25 - "Asset Integration Tests"
Cohesion: 0.15
Nodes (26): getAssetsByTextSearch(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsRejectsInvalidTextSearch(), TestGetKnownAssetsWithTextSearch() (+18 more)

### Community 26 - "Frontend Routing"
Cohesion: 0.17
Nodes (23): RequestConfigEventDetail, bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), addDisableRouteRemovalObserver() (+15 more)

### Community 27 - "Asset Input Component"
Cohesion: 0.14
Nodes (9): api, APIError, APIErrorResponse, ASSET_ACTION_BUTTON_IDENTITIES, AssetComposedColumnInput, getAsset(), Asset, BootstrapClasses (+1 more)

### Community 28 - "Allocation Plan Charts"
Cohesion: 0.13
Nodes (19): allocationPlanChart, chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset() (+11 more)

### Community 29 - "Allocation Mapping"
Cohesion: 0.15
Nodes (21): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX, AllocationDomainService (+13 more)

### Community 30 - "Portfolio REST Mapping"
Cohesion: 0.19
Nodes (21): github.com/shopspring/decimal.Decimal, DivergenceAnalysisDTS, portfolioAllocationsPerObservationTimestamp, PortfolioObservationTimestampDTS, PotentialDivergenceDTS, AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS() (+13 more)

### Community 31 - "Ordered Map Iterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, KeyValue, MapIterator, OrderedMapIterator, OrderedMapIterator[K, V], NewOrderedMapIterator() (+9 more)

### Community 32 - "Frontend Infrastructure"
Cohesion: 0.11
Nodes (18): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+10 more)

### Community 33 - "Allocation Plan Repository"
Cohesion: 0.17
Nodes (12): time.Time, plannedAllocationJoinedRowDTS, GetPlanType(), PlanType, Asset, AllocationPlan, PlannedAllocation, buildAllocationPlanFromRow() (+4 more)

### Community 34 - "Frontend Package Configuration"
Cohesion: 0.10
Nodes (20): bootstrap-icons, bootswatch, eslint, @eslint/js, globals, htmx-ext-client-side-templates, htmx-ext-form-json, navigo (+12 more)

### Community 35 - "E2E Package Configuration"
Cohesion: 0.09
Nodes (21): pg, @types/node, @types/pg, author, devDependencies, pg, @playwright/test, @types/node (+13 more)

### Community 36 - "Chart Interaction Tests"
Cohesion: 0.32
Nodes (20): CanvasPoint, CanvasTextRecorder, clickCanvasPoint(), expectChartTooltip(), expectLatestCanvasPatternState(), expectLatestCanvasTextContains(), expectLatestCanvasTextSet(), findDoughnutSlicePointByTooltip() (+12 more)

### Community 37 - "Allocation Divergence UI"
Cohesion: 0.12
Nodes (20): $60,000 Total Market Value, 60/40 Portfolio Classic - Example - 20260210-230012, Actual vs Planned Market Value Comparison, Allocation Map View, Allocation Plan Selector, BIL Underweight: -$800 (-2.96%), Bonds Underweight: -$9,000 (-15%), Directional Divergence Bars (+12 more)

### Community 38 - "E2E Shell Automation"
Cohesion: 0.24
Nodes (17): build_images(), capture_logs(), cleanup_stack(), compose(), compose_all(), compose_debug(), fail(), main() (+9 more)

### Community 39 - "Playwright Version Validation"
Cohesion: 0.12
Nodes (17): ref_node_fs, ref_node_path, ref_node_url, defaultArtifactsDirectory, packageDirectory, argumentsByName, assertEqual(), assertNodeMajor() (+9 more)

### Community 40 - "Portfolio Visualization E2E"
Cohesion: 0.20
Nodes (19): DEFAULT_ALLOCATION_STRUCTURE, EMPTY_PORTFOLIO_NAMES, expectAllocationMap(), expectAllocationPlan(), expectPortfolioHistory(), expectPortfolioList(), expectPortfolioNavigation(), expectPortfolioShell() (+11 more)

### Community 41 - "Portfolio Holdings UI"
Cohesion: 0.12
Nodes (19): Manage Portfolio Allocation Data Panel, Allocation Map Tab, Allocation Plan Tab, Asset Classification, Asset Identity and Description, Bonds Asset Class, Cash Reserve Designation, Add and Remove Holding Controls (+11 more)

### Community 42 - "Asset REST Mapping"
Cohesion: 0.25
Nodes (18): AssetExternalSource, AssetDTS, AssetSearchQueryDTS, ExternalAssetDataDTS, ExternalAssetDTS, ExternalAssetSearchQueryDTS, MapToAsset(), MapToAssetDTS() (+10 more)

### Community 43 - "Frontend Event Bindings"
Cohesion: 0.25
Nodes (17): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+9 more)

### Community 44 - "HTTP Client Utilities"
Cohesion: 0.16
Nodes (13): go_pkg_encoding_json, go_pkg_strconv, net/http.Response, RequestOption, ParseableInt, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet() (+5 more)

### Community 45 - "DOM Routing Utilities"
Cohesion: 0.17
Nodes (13): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements() (+5 more)

### Community 46 - "Handlebars Language Helpers"
Cohesion: 0.22
Nodes (17): arrayHelper(), concatHelper(), eachReverseHelper(), getPropertyHelper(), ifEqualsHelper(), ifNotEqualsHelper(), ifNotNullishHelper(), ifNullishHelper() (+9 more)

### Community 47 - "Portfolio Editing E2E"
Cohesion: 0.16
Nodes (15): DEFAULT_ALLOCATION_STRUCTURE, expectEditPortfolio(), ExpectedPortfolio, expectPortfolioList(), expectPortfolioNavigation(), expectPortfolioShell(), expectRootShell(), expectRoute() (+7 more)

### Community 48 - "Struct Pointer Utilities"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 49 - "Playwright Test Fixtures"
Cohesion: 0.17
Nodes (12): @playwright/test, src_test_e2e_support_fixtures_expect, test, TestFixtures, WorkerFixtures, delay(), hasSuccessfulResponse(), waitFor() (+4 more)

### Community 50 - "Chart Data Sources"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 51 - "Financial Input Binding"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 52 - "Yahoo Finance Client"
Cohesion: 0.17
Nodes (12): YahooFinanceAssetIntegrationService, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_net_url, YahooFinanceSearchResponseDTS, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL() (+4 more)

### Community 53 - "Gin Server"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 54 - "DOM Form Bindings"
Cohesion: 0.23
Nodes (12): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+4 more)

### Community 55 - "Frontend Logging Bindings"
Cohesion: 0.22
Nodes (13): InfraTypesUtils, Level, Levels, LogLevel, addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements() (+5 more)

### Community 56 - "Allocation Plan UI"
Cohesion: 0.14
Nodes (15): 60/40 Portfolio Classic - Example - 20260210-230012, Allocation Plan Management Interface, ARCA:BIL - SPDR Bloomberg 1-3 Month T-Bill ETF (40%), ARCA:EWZ - iShares MSCI Brazil ETF (5%), ARCA:SPY - SPDR S&P 500 ETF Trust (45%), ARCA:STIP - iShares 0-5 Year TIPS Bond ETF (10%), BONDS (60% slice), Cash Reserve Designation (+7 more)

### Community 57 - "Validation Field Errors"
Cohesion: 0.13
Nodes (3): github.com/go-playground/universal-translator.Translator, reflect.Kind, CustomFieldError

### Community 58 - "Frontend Dependencies"
Cohesion: 0.13
Nodes (15): devDependencies, eslint, @eslint/js, globals, http-proxy-middleware, parcel, @parcel/transformer-raw, @parcel/transformer-sass (+7 more)

### Community 59 - "Portfolio Snapshot UI"
Cohesion: 0.15
Nodes (14): Allocation Map Tab, Allocation Plan Tab, ARCA:EWZ Allocation 4.55%, ARCA:SPY Allocation 68.18%, Asset Allocation Donut Chart, Assets for STOCKS Level, February 2026 Portfolio Snapshot, My Portfolio Example (+6 more)

### Community 60 - "Application Configuration"
Cohesion: 0.24
Nodes (12): go_pkg_os, GinServerConfiguration, IntegrationConfiguration, BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), Configuration, RDBMSConfiguration (+4 more)

### Community 61 - "Allocation Plan REST"
Cohesion: 0.37
Nodes (13): AllocationPlanAssetDTS, AllocationPlanDTS, PlannedAllocationDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan() (+5 more)

### Community 62 - "Frontend Runtime Dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 63 - "Portfolio Page Routing"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 64 - "E2E TypeScript Configuration"
Cohesion: 0.15
Nodes (12): compilerOptions, allowImportingTsExtensions, esModuleInterop, forceConsistentCasingInFileNames, lib, module, moduleResolution, noEmit (+4 more)

### Community 65 - "JSON Reflection Utilities"
Cohesion: 0.42
Nodes (11): reflect.StructField, reflect.Type, buildJSONFieldPath(), GetJSONFieldName(), mapIndexedNamespacePart(), mapNamespacePart(), parseNamespace(), buildFieldPath() (+3 more)

### Community 66 - "Frontend TypeScript Configuration"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 67 - "Allocation Plan Detail"
Cohesion: 0.22
Nodes (11): 60/40 Portfolio Classic - Example - 20260210-230012, Allocation Map Tab, Allocation Plan Tab, Asset Allocation Donut Chart, Asset Classes Level, Bonds 60 Percent Allocation, Edit Portfolio Action, My Portfolio Example - 20260210-230012 (+3 more)

### Community 68 - "Portfolio History Detail"
Cohesion: 0.22
Nodes (11): Active Portfolio Tab, Allocation Map Tab, Allocation Plan Tab, Bonds: 45%, Class-Level Asset Allocation Donut Chart, Edit Portfolio Control, Expanded 202602 History Panel, My Portfolio Example - 20260210-230012 (+3 more)

### Community 69 - "Deep Validation"
Cohesion: 0.38
Nodes (10): reflect.Value, buildValidationErrorPathWithJSON(), DeepValidate(), validator.FieldError, isDirectValidationErrorForStruct(), validateSliceDeep(), validateStructDeep(), validateStructWithValidator() (+2 more)

### Community 70 - "Handlebars Core Helpers"
Cohesion: 0.33
Nodes (7): handlebars, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), PARTIALS_REGISTRY, registerPartialToContainer()

### Community 71 - "Docker Service Topology"
Cohesion: 0.27
Nodes (11): Base Application Service, Base Migration Engine, Base PostgreSQL Service, CI E2E Overlay, Immutable E2E Monolith, E2E PostgreSQL Service, E2E Migration Engine, Containerized Playwright Runner (+3 more)

### Community 72 - "E2E Database Utilities"
Cohesion: 0.27
Nodes (6): createE2eDatabase(), E2eDatabase, FLYWAY_HISTORY_TABLES, parsePort(), quoteIdentifier(), requiredEnvironment()

### Community 73 - "Graphify Command Pipeline"
Cohesion: 0.20
Nodes (10): Existing Graph Fast Path, Graph Health Diagnostics, Graphify Pipeline, Graphify Honesty Rules, Graph Query, Path, and Explain, Semantic Subagent Extraction, Structural AST Extraction, Exact Graphify Argument Forwarding (+2 more)

### Community 74 - "Portfolio Selection UI"
Cohesion: 0.22
Nodes (10): Dark Theme, Global All Assets Portfolio, My Portfolio Example, My Portfolio Example - 20260210-230012, New Portfolio Action, New Portfolio Focus State, Open Asset Allocator Brand, Portfolio Card Grid (+2 more)

### Community 75 - "JSON Assertion Utilities"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 76 - "Application Lifecycle"
Cohesion: 0.31
Nodes (4): os.Signal, buildStopChannel(), buildStopContext(), App

### Community 77 - "Struct Completion Utilities"
Cohesion: 0.29
Nodes (7): deepCompleteReflective(), DeepCompleteStruct(), GetStructName(), GetStructNamespaceDescription(), GetStructType(), T, StructString()

### Community 78 - "Language Comparison Helpers"
Cohesion: 0.29
Nodes (9): comparatorHelper(), AssignAtPathOptions, convertNumberForToInt(), isFiniteNumberValue(), isNullish(), reportToIntCoercion(), toComparableString(), ToIntOptions (+1 more)

### Community 79 - "HTMX Response Transformation"
Cohesion: 0.29
Nodes (9): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+1 more)

### Community 80 - "Allocation Plan Persistence"
Cohesion: 0.24
Nodes (9): seedPortfolioHistoryModificationData(), readDatabaseSnapshot(), seedAllocationMapData(), expectedPlannedAllocation(), expectPersistedAllocationPlan(), expectPersistedAllocationPlanManagement(), plannedAllocationId(), queryPlannedAllocations() (+1 more)

### Community 81 - "Divergence Analysis Options"
Cohesion: 0.33
Nodes (7): AllocationPlanIdentifierDTS, AnalysisOptionsDTS, mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), AllocationPlanIdentifier, AnalysisOptions

### Community 82 - "Validation Error Formatting"
Cohesion: 0.39
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, asValidationErrors(), formatErrorMessage(), formatValidationError(), FormatValidationErrorMessages(), validator.FieldError, validator.ValidationErrors, MapValidationErrorsToMessages()

### Community 83 - "Validation Error Builder"
Cohesion: 0.31
Nodes (5): go_pkg_github_com_go_playground_validator_v10, buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 84 - "Generic Map Tree"
Cohesion: 0.47
Nodes (3): MapTreeNode[T], MapTreeNode, T

### Community 85 - "Handlebars Utility Helpers"
Cohesion: 0.42
Nodes (8): getRenderIteratorMap(), isHelperOptions(), iteratorInitHelper(), iteratorNextHelper(), registerHandlebarsUtilHelpers(), RENDER_ITERATORS_STORE, repeaterHelper(), coerceToFiniteNumber()

### Community 86 - "Portfolio REST API"
Cohesion: 0.36
Nodes (6): AllocationStructure, PortfolioDTS, MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), Portfolio

### Community 87 - "Asset Persistence"
Cohesion: 0.32
Nodes (5): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker

### Community 89 - "External Asset Quotes"
Cohesion: 0.38
Nodes (5): golang.org/x/text/currency.Unit, ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), BuildAppErrorFormatted()

### Community 90 - "Rebalancing Workflow"
Cohesion: 0.43
Nodes (7): Create or Modify Allocation Plan, External Cash Inflow, Big or Disruptive Market Fluctuations?, Verify Assets Against Planned Allocation Percentages, Wait for Planned Interval, Move Resources to Fit the Allocation Plan, Scenario Changed?

### Community 91 - "Portfolio History Snapshot"
Cohesion: 0.33
Nodes (7): Asset Class Allocation, Bonds Allocation 45 Percent, Monthly Portfolio Snapshot 202602, Portfolio Detail Screen, Portfolio Workspace Navigation, Stocks Allocation 55 Percent, Total Market Value 60000

### Community 93 - "Database Cleanup Utilities"
Cohesion: 0.48
Nodes (4): github.com/go-ozzo/ozzo-dbx.Params, createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 94 - "Yahoo Finance Models"
Cohesion: 0.52
Nodes (6): YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS, YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS

### Community 95 - "Parcel Proxy Configuration"
Cohesion: 0.29
Nodes (6): ref_fs, http-proxy-middleware, ref_path, { createProxyMiddleware }, fs, path

### Community 96 - "Reflection Validation Helpers"
Cohesion: 0.40
Nodes (3): go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, IsNilPointer()

### Community 97 - "Custom Slice Formatting"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

### Community 98 - "Portfolio Analysis Services"
Cohesion: 0.60
Nodes (4): DivergenceAnalysisRESTController, BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService

### Community 99 - "Incremental Graph Updates"
Cohesion: 0.50
Nodes (4): Graphify URL Ingestion, Cluster-Only Graph Refresh, Incremental Graph Re-Extraction, Replace-on-Re-Extract Graph Merge

### Community 100 - "Database Destruction Scripts"
Cohesion: 0.50
Nodes (3): POSTGRES_DATA_DIR, POSTGRES_DEV_DATA_DIR, destroy.sh script

### Community 101 - "Nullable SQL Strings"
Cohesion: 0.83
Nodes (3): database/sql.NullString, StringPointerToNullString(), StringToNullString()

### Community 102 - "macOS Provisioning"
Cohesion: 0.50
Nodes (3): NVM_DIR, PATH, macos-provisioning.sh script

### Community 103 - "Frontend Build Scripts"
Cohesion: 0.50
Nodes (4): scripts, build, clean, dev

### Community 104 - "Portfolio Analysis Helpers"
Cohesion: 0.67
Nodes (3): getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), PotentialDivergence

### Community 105 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 107 - "Database Stop Script"
Cohesion: 0.50
Nodes (3): POSTGRES_DATA_DIR, POSTGRES_DEV_DATA_DIR, stop.sh script

### Community 108 - "Node Version Validation"
Cohesion: 0.83
Nodes (3): fail(), read_node_image(), validate-node-version.sh script

### Community 109 - "Graph Query System"
Cohesion: 0.67
Nodes (3): Graphify MCP Graph Server, Constrained Graph Vocabulary Expansion, Graph Query Traversal

### Community 110 - "Semantic Extraction Rules"
Cohesion: 0.67
Nodes (3): Extraction Confidence Rubric, Deterministic Full-Path Node Identity, Semantic Extraction Contract

### Community 111 - "Go Test Workflows"
Cohesion: 1.00
Nodes (3): Go External Integration Tests, Go Test Workflow, Go Unit Tests

### Community 116 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

## Knowledge Gaps
- **336 isolated node(s):** `build-dev.sh script`, `POSTGRES_DEV_DATA_DIR`, `build.sh script`, `destroy.sh script`, `POSTGRES_DEV_DATA_DIR` (+331 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 538 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **43 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `PlannedAllocation` connect `Allocation Plan Repository` to `Allocation Plan Domain`, `Database Query Infrastructure`, `Divergence Analysis Service`, `Go Package Architecture`, `Database Repository Adapters`, `Asset Persistence`, `Allocation Plan REST`, `Portfolio REST Mapping`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Database Repository Adapters` to `Application Configuration`, `Application Lifecycle`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset Repository` to `Allocation Plan Domain`, `Allocation Plan Repository`, `Database Query Infrastructure`, `Asset REST Mapping`, `Application Component Assembly`, `Portfolio Allocation Domain`, `Asset Persistence`, `Asset Integration Tests`, `Allocation Plan REST`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Are the 73 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestGetAllocationPlans()` and `TestPostAllocationPlanForInsertion()`) actually correct?**
  _`deferCloseResponseBody()` has 73 INFERRED edges - model-reasoned connections that need verification._
- **What connects `build-dev.sh script`, `POSTGRES_DEV_DATA_DIR`, `build.sh script` to the rest of the system?**
  _336 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Allocation Plan Domain` be split into smaller, more focused modules?**
  _Cohesion score 0.06346153846153846 - nodes in this community are weakly interconnected._
- **Should `Database Query Infrastructure` be split into smaller, more focused modules?**
  _Cohesion score 0.05367231638418079 - nodes in this community are weakly interconnected._