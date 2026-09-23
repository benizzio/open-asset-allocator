# Graph Report - open-asset-allocator  (2026-09-23)

## Corpus Check
- 286 files · ~299,806 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 24 file(s) not represented in the graph (top: (none) 16, .csv 2, .dbm 1)

## Summary
- 2270 nodes · 5262 edges · 155 communities (109 shown, 46 thin omitted)
- Extraction: 94% EXTRACTED · 6% INFERRED · 0% AMBIGUOUS · INFERRED: 313 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `990ee51f`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- AllocationPlanDomService
- T
- generate_report.py
- logger
- portfolio-allocation-plan-management.e2e.spec.ts
- Comparative E2E Framework Research
- allocation-plan-management.ts
- portfolio-allocation-history-management.e2e.spec.ts
- insertTestPortfolio
- BuildCleanupFunctionBuilder
- golang_ext_tree_test.go
- createBulkInsertPreparedStatement
- github.com/gin-gonic/gin.Context
- chart-utils.ts
- AssetDomService
- IsZeroValue
- portfolio-chart.ts
- portfolio-allocation-map.e2e.spec.ts
- PropagateAsAppErrorWithNewMessage
- Adapter
- go_pkg_testing
- PortfolioAllocation
- logFanOutConsumer
- allocation-plan.ts
- context.Context
- asset_integration_test.go
- binding-htmx-trigger-on-route.ts
- AssetComposedColumnInput
- allocation-plan-chart.ts
- fractal-allocation-plan-mapping.ts
- portfolio_rest_mapping.go
- NewOrderedMapIterator
- infra.ts
- AllocationPlan
- web-static/package.json
- e2e/package.json
- doughnut-chart.ts
- Allocation Map View
- e2e.sh
- validate-playwright-version.mjs
- portfolio-visualization.e2e.spec.ts
- Portfolio Holdings Editor
- asset_rest_model.go
- app.go
- go_pkg_fmt
- allocation_plan_rest_model.go
- handlebars-lang.ts
- portfolio-editing.e2e.spec.ts
- golang_ext_util_struct_unify_pointers_test.go
- fixtures.ts
- MultiChartDataSource
- binding-financial-input.ts
- asset-management.e2e.spec.ts
- GinServer
- dom/index.ts
- binding-dom-attribute-on-route.ts
- BONDS (60% slice)
- CustomFieldError
- devDependencies
- Asset Allocation Donut Chart
- testing.T
- Portfolio
- dependencies
- Portfolio Detail Page
- compilerOptions
- reflect.Type
- compilerOptions
- Asset Allocation Donut Chart
- Portfolio History Detail View
- custom_deep_validation.go
- database/sql/driver.Value
- Base Application Service
- Configuration
- Graphify Pipeline
- Portfolio Card Grid
- assert_json_extension.go
- allocation_plan_domain_service.go
- golang_ext_util_struct.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- htmx/index.ts
- asset_integration_client_yahoo_model.go
- AllocationStructure
- custom_validation_error_handling.go
- CustomValidationErrorsBuilder
- .query
- golang_ext_custom_slice.go
- allocation_rest_model.go
- T
- golang_ext_util_slice.go
- chart-contents.ts
- Verify Assets Against Planned Allocation Percentages
- Asset Class Allocation
- RESTRoute
- ExecuteDBQuery
- ExternalAsset
- .proxyrc.js
- HierarchicalId
- parseAssetTextSearch
- Incremental Graph Re-Extraction
- destroy.sh
- dom-utils.ts
- macos-provisioning.sh
- scripts
- bignumber.js
- Portfolio Section Navigation
- start-debug.sh
- stop.sh
- validate-node-version.sh
- Graph Query Traversal
- Semantic Extraction Contract
- Go External Integration Tests
- build-dev.sh
- dev.sh
- duckdb-cli.sh
- arch-provisioning.sh
- Frontend Module Architecture
- start.sh
- Property Graph Exports
- E2E Diagnostics Artifact
- build.sh
- go-dependency-graph.sh
- CustomSliceTable[T]
- migrate.sh
- ErrorResponse
- duckdb-cli-build.sh
- External Integration Tests
- Go Coding Standards
- init-db.sh
- Recursive Planned Allocation Rows
- Observation Editor
- test.sh
- Background Folder Watcher
- Post-Commit Graph Rebuild
- Saved Query Result Memory
- Media Transcription Pipeline
- CodeRabbit Review Configuration
- Go Lint Workflow
- Problem-Domain Literacy
- Weighted Project Fit Score
- github.com/benizzio/open-asset-allocator
- Split Development Stack
- Local Split-Application E2E Overlay
- DuckDB CLI Service
- Flyway Migration Service
- Flyway Migration Standards
- Go Lint Configuration
- Hierarchical Divergence Analysis
- Allocation Plan Management
- Toast Notification
- alias
- .buildAppComponents
- src_main_web_static_websrc_pages_index_assetspage

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
- `Allocation Plan` --semantically_similar_to--> `Allocation Planning`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Convergence Planning` --semantically_similar_to--> `Convergence Analysis and Planning`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Divergence Analysis` --semantically_similar_to--> `Divergence Analysis`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Fractal Portfolio Hierarchy` --semantically_similar_to--> `Customizable Fractal Hierarchy`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md
- `Portfolio History` --semantically_similar_to--> `Historical Portfolio Snapshots`  [INFERRED] [semantically similar]
  README.md → docs/readme-features.md

## Import Cycles
- 3-file cycle: `src/main/web-static/websrc/infra/htmx/index.ts -> src/main/web-static/websrc/infra/routing/index.ts -> src/main/web-static/websrc/infra/routing/binding-htmx-trigger-on-route.ts -> src/main/web-static/websrc/infra/htmx/index.ts`

## Hyperedges (group relationships)
- **Asset Allocation Monitoring and Rebalancing Cycle** — docs_images_asset_allocation_flow_allocation_plan_creation_or_modification, docs_images_asset_allocation_flow_out_of_balance_asset_verification, docs_images_asset_allocation_flow_resource_reallocation, docs_images_asset_allocation_flow_planned_interval_wait, docs_images_asset_allocation_flow_market_fluctuation_check, docs_images_asset_allocation_flow_scenario_change_check [EXTRACTED 1.00]
- **60/40 Asset Allocation** — docs_images_allocation_plan_detail_60_40_portfolio_classic, docs_images_allocation_plan_detail_asset_allocation_donut_chart, docs_images_allocation_plan_detail_bonds_60_percent_allocation, docs_images_allocation_plan_detail_stocks_40_percent_allocation [EXTRACTED 1.00]
- **Portfolio Detail Navigation** — docs_images_allocation_plan_detail_portfolio_tab, docs_images_allocation_plan_detail_allocation_plan_tab, docs_images_allocation_plan_detail_allocation_map_tab [EXTRACTED 1.00]
- **BONDS Slice Composition** — docs_images_allocation_plan_management_bonds, docs_images_allocation_plan_management_arca_bil, docs_images_allocation_plan_management_nasdaqgm_ief, docs_images_allocation_plan_management_nasdaqgm_tlt, docs_images_allocation_plan_management_arca_stip [EXTRACTED 1.00]
- **STOCKS Slice Composition** — docs_images_allocation_plan_management_stocks, docs_images_allocation_plan_management_nasdaqgm_shv, docs_images_allocation_plan_management_arca_spy, docs_images_allocation_plan_management_arca_ewz [EXTRACTED 1.00]
- **Top-Level 60/40 Allocation** — docs_images_allocation_plan_management_60_40_portfolio_classic_example_20260210_230012, docs_images_allocation_plan_management_bonds, docs_images_allocation_plan_management_stocks [EXTRACTED 1.00]
- **Hierarchical Allocation Divergence Rows** — docs_images_divergence_analysis_bonds_underweight_9_000_15, docs_images_divergence_analysis_stocks_overweight_9_000_15, docs_images_divergence_analysis_bil_underweight_800_2_96, docs_images_divergence_analysis_stip_overweight_5_300_19_63, docs_images_divergence_analysis_ief_underweight_2_100_7_78, docs_images_divergence_analysis_tlt_underweight_2_400_8_89, docs_images_divergence_analysis_shv_underweight_7_500_22_73, docs_images_divergence_analysis_spy_overweight_7_650_23_18, docs_images_divergence_analysis_ewz_underweight_150_0_45 [EXTRACTED 1.00]
- **Portfolio Allocation Map Interface Composition** — docs_images_divergence_analysis_portfolio_workflow_tabs, docs_images_divergence_analysis_monthly_snapshot_accordion, docs_images_divergence_analysis_allocation_plan_selector, docs_images_divergence_analysis_actual_vs_planned_market_value_comparison, docs_images_divergence_analysis_hierarchical_asset_allocation_breakdown, docs_images_divergence_analysis_directional_divergence_bars [EXTRACTED 1.00]
- **Asset Allocation Composition** — docs_images_portfolio_history_detail2_asset_allocation_donut_chart, docs_images_portfolio_history_detail2_nasdaqgm_shv_cash_reserve_allocation, docs_images_portfolio_history_detail2_arca_spy_allocation, docs_images_portfolio_history_detail2_arca_ewz_allocation [EXTRACTED 1.00]
- **Portfolio Detail Navigation Tabs** — docs_images_portfolio_history_detail2_portfolio_tab, docs_images_portfolio_history_detail2_allocation_plan_tab, docs_images_portfolio_history_detail2_allocation_map_tab [EXTRACTED 1.00]
- **Portfolio Choice Grid** — docs_images_portfolio_selection_global_all_assets_portfolio, docs_images_portfolio_selection_my_portfolio_example, docs_images_portfolio_selection_my_portfolio_example_20260210_230012, docs_images_portfolio_selection_new_portfolio_action [EXTRACTED 1.00]
- **Fourteen Evaluated E2E Candidates** — docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_playwright_test, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_cypress, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_webdriverio, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_selenium_webdriver_javascript, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_nightwatchjs, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_testcafe, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_puppeteer_vitest_jest, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_rod_go_testing, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_chromedp_go_testing, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_codeceptjs, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_testplane, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_cucumber_playwright_bdd, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_robot_framework_browser, docs_research_best_maintained_and_community_adopted_e2e_testing_tool_for_the_open_asset_allocator_project_report_grafana_k6_browser [EXTRACTED 1.00]
- **February 2026 Class Allocation Summary** — docs_images_portfolio_history_detail1_expanded_202602_history_panel, docs_images_portfolio_history_detail1_total_market_value_60000, docs_images_portfolio_history_detail1_class_level_asset_allocation_chart, docs_images_portfolio_history_detail1_bonds_45_percent, docs_images_portfolio_history_detail1_stocks_55_percent [EXTRACTED 1.00]
- **Graphify Semantic Extraction Integrity Contract** — _agents_skills_graphify_references_extraction_spec_semantic_extraction_contract, _agents_skills_graphify_references_extraction_spec_confidence_rubric, _agents_skills_graphify_references_extraction_spec_deterministic_node_identity [EXTRACTED 1.00]
- **Monthly Portfolio Valuation Breakdown** — docs_images_portfolio_history_monthly_portfolio_snapshot_202602, docs_images_portfolio_history_total_market_value_60000, docs_images_portfolio_history_asset_class_allocation, docs_images_portfolio_history_bonds_allocation_45_percent, docs_images_portfolio_history_stocks_allocation_55_percent [EXTRACTED 1.00]
- **Portfolio Detail Navigation** — docs_images_portfolio_history_detail1_active_portfolio_tab, docs_images_portfolio_history_detail1_allocation_plan_tab, docs_images_portfolio_history_detail1_allocation_map_tab [EXTRACTED 1.00]
- **Portfolio HTMX Route Flow** — src_main_web_static_root_htmx_lazy_route_loading, src_main_web_static_websrc_pages_portfolios_portfolio_list_page, src_main_web_static_websrc_pages_portfolio_portfolio_detail_page [EXTRACTED 1.00]
- **Portfolio Observation Holding Fields** — docs_images_portfolio_history_management_asset_identity, docs_images_portfolio_history_management_asset_classification, docs_images_portfolio_history_management_cash_reserve_designation, docs_images_portfolio_history_management_position_quantity, docs_images_portfolio_history_management_market_price, docs_images_portfolio_history_management_total_market_value [EXTRACTED 1.00]
- **Portfolio Allocation User Interface Flow** — src_main_web_static_websrc_components_portfolio_navigation_portfolio_section_navigation, src_main_web_static_websrc_components_portfolio_history_portfolio_history_viewer, src_main_web_static_websrc_components_allocation_plan_allocation_plan_viewer, src_main_web_static_websrc_components_allocation_map_allocation_map, src_main_web_static_websrc_components_asset_composed_columns_input_asset_composed_columns_input [INFERRED 0.85]
- **Containerized E2E Execution Topology** — src_main_docker_docker_compose_e2e_e2e_database, src_main_docker_docker_compose_e2e_e2e_migration_engine, src_main_docker_docker_compose_e2e_playwright_runner, src_main_docker_docker_compose_e2e_ci_immutable_monolith, _github_workflows_e2e_e2e_tests_workflow [INFERRED 0.95]

## Communities (155 total, 46 thin omitted)

### Community 0 - "AllocationPlanDomService"
Cohesion: 0.17
Nodes (10): DomainValidationError, UniqueConstraintViolationError, AllocationPlanRepository, BuildAllocationPlanDomService(), AllocationPlanDomService, BuildAppErrorFormattedUnconverted(), BuildDomainValidationError(), AppError (+2 more)

### Community 1 - "T"
Cohesion: 0.11
Nodes (16): github.com/go-ozzo/ozzo-dbx.Query, QueryBuilder, QueryBuilder[T], QueryExecutor, QueryExecutor[T], RowScanner, SingleRowScanner, SQLTransactionalQueryBuilder (+8 more)

### Community 2 - "generate_report.py"
Cohesion: 0.06
Nodes (57): Any, _anchor(), _base_result_name(), _build_output_mapping(), _category_aliases(), _collect_extra_fields(), visit(), _compact_summary() (+49 more)

### Community 3 - "logger"
Cohesion: 0.25
Nodes (17): bindPercentageInput(), bindPercentageInputElements(), bindPercentageInputsInDescendants(), configurePercentageInputAttributes(), createHiddenDecimalField(), initializePercentageDisplay(), syncPercentageToDecimal(), syncPercentageToDecimalInContainer() (+9 more)

### Community 4 - "portfolio-allocation-plan-management.e2e.spec.ts"
Cohesion: 0.06
Nodes (43): addAssetAllocationRow(), addClassAllocationRow(), DEFAULT_ALLOCATION_STRUCTURE, expectAllocationPlan(), expectAllocationPlanManagement(), expectAllocationPlanManagementForm(), expectDraftAllocationRow(), expectDraftAllocationRows() (+35 more)

### Community 5 - "Comparative E2E Framework Research"
Cohesion: 0.05
Nodes (49): Root AGENTS Instructions for Copilot, Dependabot Dependency Updates, Renovate Runtime Coordination, Atomic Persistence-Verified E2E Testing, Open Asset Allocator Architecture, Runtime Version Coordination, Portfolio History Form TODOs, Allocation Planning (+41 more)

### Community 6 - "allocation-plan-management.ts"
Cohesion: 0.05
Nodes (33): htmx.org, api, APIError, APIErrorResponse, Application, addPlannedAllocationRow(), allocationPlanManagement, AllocationPlanningHierarchicalFormEntry (+25 more)

### Community 7 - "portfolio-allocation-history-management.e2e.spec.ts"
Cohesion: 0.06
Nodes (34): installCanvasTextRecorder(), DEFAULT_ALLOCATION_STRUCTURE, expectEditableObservationRows(), ExpectedObservationRow, expectExistingAsset(), expectNewObservationRows(), expectPersistedAllocations(), expectPersistedModifiedPortfolioHistory() (+26 more)

### Community 8 - "insertTestPortfolio"
Cohesion: 0.19
Nodes (15): database/sql.NullString, StringPointerToNullString(), StringToNullString(), putForValidationFailure(), TestPostPortfolio(), TestPostPortfolioWithAllocationStructure(), TestPutPortfolio(), TestPutPortfolioFailureWithNameExceedingMaxLength() (+7 more)

### Community 9 - "BuildCleanupFunctionBuilder"
Cohesion: 0.25
Nodes (26): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAsset(), assertPersistedAssetWithExternalData(), TestPostPortfolioAllocationHistoryFullMerge() (+18 more)

### Community 10 - "golang_ext_tree_test.go"
Cohesion: 0.15
Nodes (25): NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch(), TestAddBranchBreakingOnZeroValues_IntZeroValue(), TestAddBranchBreakingOnZeroValues_StopsAtZeroValue() (+17 more)

### Community 11 - "createBulkInsertPreparedStatement"
Cohesion: 0.29
Nodes (6): database/sql.Stmt, database/sql.Tx, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), prepareBulkInsertValues(), quoteIdentifiers()

### Community 12 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.14
Nodes (19): go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, PortfolioDTS, AssetRESTController, PortfolioRESTController, MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs() (+11 more)

### Community 13 - "chart-utils.ts"
Cohesion: 0.11
Nodes (31): chartjs-plugin-datalabels, chart, chartContentRepo, getChartContent(), getChartContentFromChart(), loadChart(), buildChartInteractions(), buildChartOptions() (+23 more)

### Community 14 - "AssetDomService"
Cohesion: 0.12
Nodes (15): AllocationPlanRESTController, AssetIntegrationServicesPerSource, BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), AssetIntegrationService (+7 more)

### Community 15 - "IsZeroValue"
Cohesion: 0.27
Nodes (4): ParseableInt, PortfolioAllocationRESTController, T, IsZeroValue()

### Community 16 - "portfolio-chart.ts"
Cohesion: 0.16
Nodes (15): chart.js, changeChartData(), chartDataSelectionEventHandler(), FractalPortfolioMultiChartDataSource, generateDataKey(), getChartContent(), interactionObserverCallback(), getAccumulatedAllocationsPerProperty() (+7 more)

### Community 17 - "portfolio-allocation-map.e2e.spec.ts"
Cohesion: 0.08
Nodes (33): ASSET_DATA, DatabaseSnapshot, DEFAULT_ALLOCATION_STRUCTURE, expandRootAndAssertChildren(), expectAllocationMapShell(), expectAnalysisTable(), ExpectedDivergenceNode, expectNodeRow() (+25 more)

### Community 18 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.13
Nodes (18): database/sql.Rows, AllocationPlanRDBMSRepository, AssetRDBMSRepository, S, Asset, assetRowScanner(), buildAssetInsertValue(), buildAssetInsertValues() (+10 more)

### Community 19 - "Adapter"
Cohesion: 0.14
Nodes (8): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, contextKey, TransactionalContext, Adapter, SQLTransactionalContext, withTransaction()

### Community 20 - "go_pkg_testing"
Cohesion: 0.06
Nodes (37): malformedQueryBindingTarget, go_pkg_context, go_pkg_database_sql, go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_lib_pq (+29 more)

### Community 21 - "PortfolioAllocation"
Cohesion: 0.12
Nodes (18): go_pkg_github_com_shopspring_decimal, portfolioAllocationJoinedRowDTS, PortfolioAllocationRDBMSRepository, MapToPortfolioObservationTimestamp(), BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations() (+10 more)

### Community 22 - "logFanOutConsumer"
Cohesion: 0.10
Nodes (20): go_pkg_github_com_nhatthm_httpmock, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry() (+12 more)

### Community 23 - "allocation-plan.ts"
Cohesion: 0.11
Nodes (30): AllocationHierarchyLevelDTO, AllocationPlanType, ASSET_ALLOCATION_PLAN, BALANCING_EXECUTION_PLAN, AllocationStructure, AllocationStructureDTO, LOWEST_AVAILABLE_HIERARCHY_LEVEL, LOWEST_AVAILABLE_HIERARCHY_LEVEL_INDEX (+22 more)

### Community 24 - "context.Context"
Cohesion: 0.06
Nodes (61): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, potentialDivergencesPerHierarchicalId, context.CancelFunc, context.Context, sync.WaitGroup, I (+53 more)

### Community 25 - "asset_integration_test.go"
Cohesion: 0.13
Nodes (31): go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_io, go_pkg_net_http, go_pkg_strconv, go_pkg_strings, TestGetExternalAssetsFailure() (+23 more)

### Community 26 - "binding-htmx-trigger-on-route.ts"
Cohesion: 0.16
Nodes (24): bindCLickNavigation(), bindKeypressNavigation(), bindNavigateToElements(), bindNavigateToInDescendants(), buildDestinationPath(), navigate(), addDisableRouteRemovalObserver(), bindCleanOnExitRouteBehaviourOnElement() (+16 more)

### Community 28 - "allocation-plan-chart.ts"
Cohesion: 0.24
Nodes (9): chartDataSelectionEventHandler(), FractalPlannedAllocationMultiChartDataSource, getChartContent(), getSelectedDataKey(), interactionObserverCallback(), mapChildDatasets(), mapDataset(), toChartDataMap() (+1 more)

### Community 29 - "fractal-allocation-plan-mapping.ts"
Cohesion: 0.30
Nodes (13): PlannedAllocation, getAllocationHierarchySize(), getHierarchicalIdAsString(), getHierarchyLevelIndex(), getPlannedAllocationHierarchicalIdAsString(), getTopLevelHierarchyIndexFromAllocationStructure(), getTopLevelHierarchyIndexFromPlannedAllocation(), connectAllocationsToFractalStructure() (+5 more)

### Community 30 - "portfolio_rest_mapping.go"
Cohesion: 0.16
Nodes (27): github.com/shopspring/decimal.Decimal, AllocationPlanIdentifierDTS, AnalysisOptionsDTS, DivergenceAnalysisDTS, portfolioAllocationsPerObservationTimestamp, PortfolioObservationTimestampDTS, PotentialDivergenceDTS, AggregateAndMapToPortfolioHistoryDTSs() (+19 more)

### Community 31 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, KeyValue, MapIterator, OrderedMapIterator, OrderedMapIterator[K, V], NewOrderedMapIterator() (+9 more)

### Community 32 - "infra.ts"
Cohesion: 0.12
Nodes (17): bootstrap, BootstrapNotification, NOTIFICATION_TYPE_BOOTSTRAP_CLASSES, DomInfra, handlebarsInfra, bootRouterDebouncing(), DOM_SETTLING_BEHAVIOR_EVENT_HANDLER(), GeneralErrorHandler (+9 more)

### Community 33 - "AllocationPlan"
Cohesion: 0.16
Nodes (13): plannedAllocationJoinedRowDTS, mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), GetPlanType(), PlanType, Asset, AllocationPlan, PlannedAllocation (+5 more)

### Community 34 - "web-static/package.json"
Cohesion: 0.10
Nodes (20): bootstrap-icons, bootswatch, chroma-js, eslint, @eslint/js, globals, htmx-ext-client-side-templates, htmx-ext-form-json (+12 more)

### Community 35 - "e2e/package.json"
Cohesion: 0.09
Nodes (21): pg, @types/node, @types/pg, author, devDependencies, pg, @playwright/test, @types/node (+13 more)

### Community 36 - "doughnut-chart.ts"
Cohesion: 0.32
Nodes (20): CanvasPoint, CanvasTextRecorder, clickCanvasPoint(), expectChartTooltip(), expectLatestCanvasPatternState(), expectLatestCanvasTextContains(), expectLatestCanvasTextSet(), findDoughnutSlicePointByTooltip() (+12 more)

### Community 37 - "Allocation Map View"
Cohesion: 0.12
Nodes (20): $60,000 Total Market Value, 60/40 Portfolio Classic - Example - 20260210-230012, Actual vs Planned Market Value Comparison, Allocation Map View, Allocation Plan Selector, BIL Underweight: -$800 (-2.96%), Bonds Underweight: -$9,000 (-15%), Directional Divergence Bars (+12 more)

### Community 38 - "e2e.sh"
Cohesion: 0.24
Nodes (17): build_images(), capture_logs(), cleanup_stack(), compose(), compose_all(), compose_debug(), fail(), main() (+9 more)

### Community 39 - "validate-playwright-version.mjs"
Cohesion: 0.11
Nodes (18): ref_node_fs, ref_node_path, ref_node_url, @playwright/test, defaultArtifactsDirectory, packageDirectory, argumentsByName, assertEqual() (+10 more)

### Community 40 - "portfolio-visualization.e2e.spec.ts"
Cohesion: 0.20
Nodes (19): DEFAULT_ALLOCATION_STRUCTURE, EMPTY_PORTFOLIO_NAMES, expectAllocationMap(), expectAllocationPlan(), expectPortfolioHistory(), expectPortfolioList(), expectPortfolioNavigation(), expectPortfolioShell() (+11 more)

### Community 41 - "Portfolio Holdings Editor"
Cohesion: 0.12
Nodes (19): Manage Portfolio Allocation Data Panel, Allocation Map Tab, Allocation Plan Tab, Asset Classification, Asset Identity and Description, Bonds Asset Class, Cash Reserve Designation, Add and Remove Holding Controls (+11 more)

### Community 42 - "asset_rest_model.go"
Cohesion: 0.20
Nodes (18): AssetDTS, AssetSearchQueryDTS, ExternalAssetDataDTS, ExternalAssetDTS, ExternalAssetSearchQueryDTS, MapToAsset(), MapToAssetDTS(), MapToAssetDTSs() (+10 more)

### Community 43 - "app.go"
Cohesion: 0.16
Nodes (11): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os, go_pkg_os_signal, go_pkg_syscall, os.Signal (+3 more)

### Community 44 - "go_pkg_fmt"
Cohesion: 0.10
Nodes (19): go_pkg_encoding_json, go_pkg_errors, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_go_playground_universal_translator, go_pkg_net, go_pkg_net_url, go_pkg_path_filepath (+11 more)

### Community 45 - "allocation_plan_rest_model.go"
Cohesion: 0.27
Nodes (15): time.Time, AllocationPlanAssetDTS, AllocationPlanDTS, PlannedAllocationDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs() (+7 more)

### Community 46 - "handlebars-lang.ts"
Cohesion: 0.09
Nodes (41): handlebars, domJSONHelper(), registerHandlebarsDOMHelpers(), handlebarsFormatCurrency(), registerHandlebarsFormatHelper(), arrayHelper(), comparatorHelper(), concatHelper() (+33 more)

### Community 47 - "portfolio-editing.e2e.spec.ts"
Cohesion: 0.16
Nodes (15): DEFAULT_ALLOCATION_STRUCTURE, expectEditPortfolio(), ExpectedPortfolio, expectPortfolioList(), expectPortfolioNavigation(), expectPortfolioShell(), expectRootShell(), expectRoute() (+7 more)

### Community 48 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 49 - "fixtures.ts"
Cohesion: 0.12
Nodes (17): createE2eDatabase(), E2eDatabase, FLYWAY_HISTORY_TABLES, parsePort(), quoteIdentifier(), requiredEnvironment(), src_test_e2e_support_fixtures_expect, test (+9 more)

### Community 50 - "MultiChartDataSource"
Cohesion: 0.12
Nodes (4): ChartDataSource, ChartDataSourceVisitor, MultiChartDataSource, SingleChartDataSource

### Community 51 - "binding-financial-input.ts"
Cohesion: 0.24
Nodes (16): bindFinancialInput(), bindFinancialInputElements(), bindFinancialInputsInDescendants(), configureFinancialInputAttributes(), createHiddenRawValueField(), getDecimalPlacesFromContainer(), getDecimalPlacesFromInput(), initializeFinancialDisplay() (+8 more)

### Community 52 - "asset-management.e2e.spec.ts"
Cohesion: 0.14
Nodes (3): Asset, AssetRow, PERSISTED_EXTERNAL_DATA

### Community 53 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 54 - "dom/index.ts"
Cohesion: 0.23
Nodes (12): bindBootstrapValidationCleaning(), bindBootstrapValidationOnSubmit(), bindBootstrapValidationToDefaultForm(), bindFormsInDescendants(), addDisplayObserver(), bindExclusiveDisplay(), bindExclusiveDisplayContainerInDescendants(), bindExclusiveDisplayInDescendants() (+4 more)

### Community 55 - "binding-dom-attribute-on-route.ts"
Cohesion: 0.47
Nodes (8): addAttributes(), addRouterHooks(), bindAttributeOnRoute(), bindAttributeOnRouteElements(), bindAttributeOnRouteInDescendants(), executeImmediatelyIfOnRoute(), extractBindingData(), removeAttributes()

### Community 56 - "BONDS (60% slice)"
Cohesion: 0.14
Nodes (15): 60/40 Portfolio Classic - Example - 20260210-230012, Allocation Plan Management Interface, ARCA:BIL - SPDR Bloomberg 1-3 Month T-Bill ETF (40%), ARCA:EWZ - iShares MSCI Brazil ETF (5%), ARCA:SPY - SPDR S&P 500 ETF Trust (45%), ARCA:STIP - iShares 0-5 Year TIPS Bond ETF (10%), BONDS (60% slice), Cash Reserve Designation (+7 more)

### Community 57 - "CustomFieldError"
Cohesion: 0.13
Nodes (3): github.com/go-playground/universal-translator.Translator, reflect.Kind, CustomFieldError

### Community 58 - "devDependencies"
Cohesion: 0.13
Nodes (15): devDependencies, eslint, @eslint/js, globals, http-proxy-middleware, parcel, @parcel/transformer-raw, @parcel/transformer-sass (+7 more)

### Community 59 - "Asset Allocation Donut Chart"
Cohesion: 0.15
Nodes (14): Allocation Map Tab, Allocation Plan Tab, ARCA:EWZ Allocation 4.55%, ARCA:SPY Allocation 68.18%, Asset Allocation Donut Chart, Assets for STOCKS Level, February 2026 Portfolio Snapshot, My Portfolio Example (+6 more)

### Community 60 - "testing.T"
Cohesion: 0.08
Nodes (54): testing.T, demoStringer, TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches() (+46 more)

### Community 61 - "Portfolio"
Cohesion: 0.21
Nodes (6): AllocationStructure, PortfolioRDBMSRepository, Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService

### Community 62 - "dependencies"
Cohesion: 0.14
Nodes (14): dependencies, bignumber.js, bootstrap, bootstrap-icons, bootswatch, chart.js, chartjs-plugin-datalabels, chroma-js (+6 more)

### Community 63 - "Portfolio Detail Page"
Cohesion: 0.15
Nodes (13): HTMX Lazy Route Loading, Open Asset Allocator Shell, Portfolio Route Container, Portfolios Route Container, Edit Portfolio Form, Portfolio Context API Loading, Portfolio Detail Page, Portfolio Route Components (+5 more)

### Community 64 - "compilerOptions"
Cohesion: 0.15
Nodes (12): compilerOptions, allowImportingTsExtensions, esModuleInterop, forceConsistentCasingInFileNames, lib, module, moduleResolution, noEmit (+4 more)

### Community 65 - "reflect.Type"
Cohesion: 0.37
Nodes (12): reflect.StructField, reflect.Type, buildJSONFieldPath(), GetJSONFieldName(), mapIndexedNamespacePart(), mapNamespacePart(), parseNamespace(), buildFieldPath() (+4 more)

### Community 66 - "compilerOptions"
Cohesion: 0.17
Nodes (11): compilerOptions, esModuleInterop, isolatedModules, lib, module, moduleResolution, noEmit, skipLibCheck (+3 more)

### Community 67 - "Asset Allocation Donut Chart"
Cohesion: 0.22
Nodes (11): 60/40 Portfolio Classic - Example - 20260210-230012, Allocation Map Tab, Allocation Plan Tab, Asset Allocation Donut Chart, Asset Classes Level, Bonds 60 Percent Allocation, Edit Portfolio Action, My Portfolio Example - 20260210-230012 (+3 more)

### Community 68 - "Portfolio History Detail View"
Cohesion: 0.22
Nodes (11): Active Portfolio Tab, Allocation Map Tab, Allocation Plan Tab, Bonds: 45%, Class-Level Asset Allocation Donut Chart, Edit Portfolio Control, Expanded 202602 History Panel, My Portfolio Example - 20260210-230012 (+3 more)

### Community 69 - "custom_deep_validation.go"
Cohesion: 0.42
Nodes (9): reflect.Value, buildValidationErrorPathWithJSON(), validator.FieldError, isDirectValidationErrorForStruct(), validateSliceDeep(), validateStructDeep(), validateStructWithValidator(), validateValueDeep() (+1 more)

### Community 70 - "database/sql/driver.Value"
Cohesion: 0.21
Nodes (5): database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn(), IsNilPointer()

### Community 71 - "Base Application Service"
Cohesion: 0.27
Nodes (11): Base Application Service, Base Migration Engine, Base PostgreSQL Service, CI E2E Overlay, Immutable E2E Monolith, E2E PostgreSQL Service, E2E Migration Engine, Containerized Playwright Runner (+3 more)

### Community 72 - "Configuration"
Cohesion: 0.22
Nodes (12): GinServerConfiguration, IntegrationConfiguration, BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), Configuration, RDBMSConfiguration, YahooFinanceConfiguration (+4 more)

### Community 73 - "Graphify Pipeline"
Cohesion: 0.20
Nodes (10): Existing Graph Fast Path, Graph Health Diagnostics, Graphify Pipeline, Graphify Honesty Rules, Graph Query, Path, and Explain, Semantic Subagent Extraction, Structural AST Extraction, Exact Graphify Argument Forwarding (+2 more)

### Community 74 - "Portfolio Card Grid"
Cohesion: 0.22
Nodes (10): Dark Theme, Global All Assets Portfolio, My Portfolio Example, My Portfolio Example - 20260210-230012, New Portfolio Action, New Portfolio Focus State, Open Asset Allocator Brand, Portfolio Card Grid (+2 more)

### Community 75 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 76 - "allocation_plan_domain_service.go"
Cohesion: 0.36
Nodes (11): allocationPlanValidationData, levelSliceSizeValidationData, AllocationHierarchy, appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData() (+3 more)

### Community 77 - "golang_ext_util_struct.go"
Cohesion: 0.29
Nodes (7): deepCompleteReflective(), DeepCompleteStruct(), GetStructName(), GetStructNamespaceDescription(), GetStructType(), T, StructString()

### Community 78 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.19
Nodes (15): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms (+7 more)

### Community 79 - "htmx/index.ts"
Cohesion: 0.13
Nodes (22): bindHTMXTransformResponseElement(), bindHTMXTransformResponseElements(), bindHTMXTransformResponseInDescendants(), extractPathRegExpForTransform(), htmxTransformResponse, registerTransformResponseFunction(), TRANSFORM_RESPONSE_FUNCTION_MAP, transformResponse() (+14 more)

### Community 80 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.23
Nodes (10): YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS, YahooFinanceSearchResponseDTS, buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceChartResponseDTS (+2 more)

### Community 81 - "AllocationStructure"
Cohesion: 0.25
Nodes (5): go_pkg_database_sql_driver, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, AllocationStructure, T, ScanJsonColumn()

### Community 82 - "custom_validation_error_handling.go"
Cohesion: 0.39
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, asValidationErrors(), formatErrorMessage(), formatValidationError(), FormatValidationErrorMessages(), validator.FieldError, validator.ValidationErrors, MapValidationErrorsToMessages()

### Community 83 - "CustomValidationErrorsBuilder"
Cohesion: 0.33
Nodes (5): buildCustomValidationError(), BuildCustomValidationErrorsBuilder(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 84 - ".query"
Cohesion: 0.24
Nodes (9): seedPortfolioHistoryModificationData(), readDatabaseSnapshot(), seedAllocationMapData(), expectedPlannedAllocation(), expectPersistedAllocationPlan(), expectPersistedAllocationPlanManagement(), plannedAllocationId(), queryPlannedAllocations() (+1 more)

### Community 85 - "golang_ext_custom_slice.go"
Cohesion: 0.25
Nodes (6): CustomSlice[T], stripRootFromBranches(), CustomSlice, CustomSliceTable, T, joinAny()

### Community 86 - "allocation_rest_model.go"
Cohesion: 0.54
Nodes (7): AllocationHierarchyLevelDTS, AllocationStructureDTS, mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel

### Community 87 - "T"
Cohesion: 0.47
Nodes (3): MapTreeNode[T], MapTreeNode, T

### Community 88 - "golang_ext_util_slice.go"
Cohesion: 0.36
Nodes (7): cleanNilAllocations(), CleanNilPointersInSlice(), DereferenceSliceContent(), T, IsSlice(), ReverseSlice(), SliceContainsZeroValue()

### Community 89 - "chart-contents.ts"
Cohesion: 0.29
Nodes (7): allocationPlanChart, toChartContent(), toUnidimensionalMultiChartContent(), portfolioChart, ChartDataType, ASSET_ALLOCATION_PLAN_1D, PORTFOLIO_HISTORY_1D

### Community 90 - "Verify Assets Against Planned Allocation Percentages"
Cohesion: 0.43
Nodes (7): Create or Modify Allocation Plan, External Cash Inflow, Big or Disruptive Market Fluctuations?, Verify Assets Against Planned Allocation Percentages, Wait for Planned Interval, Move Resources to Fit the Allocation Plan, Scenario Changed?

### Community 91 - "Asset Class Allocation"
Cohesion: 0.33
Nodes (7): Asset Class Allocation, Bonds Allocation 45 Percent, Monthly Portfolio Snapshot 202602, Portfolio Detail Screen, Portfolio Workspace Navigation, Stocks Allocation 55 Percent, Total Market Value 60000

### Community 93 - "ExecuteDBQuery"
Cohesion: 0.27
Nodes (11): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), addAssetExternalDataRestoreCleanup(), capturePersistedAssetExternalData(), ExecuteDBQuery(), TestGetPortfolioAllocationHistoryOmitsExternalAssetWhenNull(), TestGetPortfolioAllocationHistoryWithProjectedExternalAsset(), TestPostPortfolioAllocationHistoryDoesNotOverwriteExistingAssetExternalData() (+3 more)

### Community 94 - "ExternalAsset"
Cohesion: 0.19
Nodes (14): YahooFinanceAssetIntegrationService, AssetExternalSource, go_pkg_golang_org_x_text_currency, golang.org/x/text/currency.Unit, ExternalAsset, ExternalAssetQuote, BuildYahooFinanceAssetIntegrationService(), extractLastClose() (+6 more)

### Community 95 - ".proxyrc.js"
Cohesion: 0.29
Nodes (6): ref_fs, http-proxy-middleware, ref_path, { createProxyMiddleware }, fs, path

### Community 97 - "parseAssetTextSearch"
Cohesion: 0.60
Nodes (4): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch(), TestParseAssetTextSearch()

### Community 99 - "Incremental Graph Re-Extraction"
Cohesion: 0.50
Nodes (4): Graphify URL Ingestion, Cluster-Only Graph Refresh, Incremental Graph Re-Extraction, Replace-on-Re-Extract Graph Merge

### Community 100 - "destroy.sh"
Cohesion: 0.50
Nodes (3): POSTGRES_DATA_DIR, POSTGRES_DEV_DATA_DIR, destroy.sh script

### Community 101 - "dom-utils.ts"
Cohesion: 0.18
Nodes (12): addRemoveObserver(), contextDataCache, ensureSharedObserver(), getCacheableContextData(), observedElements, addDisplayOnRouteRemovalObserver(), bindDisplayOnRoute(), bindDisplayOnRouteElements() (+4 more)

### Community 102 - "macos-provisioning.sh"
Cohesion: 0.50
Nodes (3): NVM_DIR, PATH, macos-provisioning.sh script

### Community 103 - "scripts"
Cohesion: 0.50
Nodes (4): scripts, build, clean, dev

### Community 104 - "bignumber.js"
Cohesion: 0.19
Nodes (11): bignumber.js, getValueLabel(), registerPortfolioAnalysisHandlebarsHelpers(), ObservationTimestamp, PortfolioAllocation, PortfolioAllocationDTO, PortfolioSnapshotDTO, DivergenceAnalysis (+3 more)

### Community 105 - "Portfolio Section Navigation"
Cohesion: 0.50
Nodes (4): Allocation Map, Allocation Plan Viewer, Portfolio History Viewer, Portfolio Section Navigation

### Community 107 - "stop.sh"
Cohesion: 0.50
Nodes (3): POSTGRES_DATA_DIR, POSTGRES_DEV_DATA_DIR, stop.sh script

### Community 108 - "validate-node-version.sh"
Cohesion: 0.83
Nodes (3): fail(), read_node_image(), validate-node-version.sh script

### Community 109 - "Graph Query Traversal"
Cohesion: 0.67
Nodes (3): Graphify MCP Graph Server, Constrained Graph Vocabulary Expansion, Graph Query Traversal

### Community 110 - "Semantic Extraction Contract"
Cohesion: 0.67
Nodes (3): Extraction Confidence Rubric, Deterministic Full-Path Node Identity, Semantic Extraction Contract

### Community 111 - "Go External Integration Tests"
Cohesion: 1.00
Nodes (3): Go External Integration Tests, Go Test Workflow, Go Unit Tests

### Community 116 - "Frontend Module Architecture"
Cohesion: 0.67
Nodes (3): Frontend Module Architecture, HTMX-First API Calls, index.ts Module API Boundaries

### Community 157 - ".buildAppComponents"
Cohesion: 0.16
Nodes (12): github.com/go-ozzo/ozzo-dbx.Rows, AllocationRDBMSRepository, BuildPortfolioRESTController(), AllocationRepository, BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository() (+4 more)

## Knowledge Gaps
- **341 isolated node(s):** `build-dev.sh script`, `POSTGRES_DEV_DATA_DIR`, `build.sh script`, `destroy.sh script`, `POSTGRES_DEV_DATA_DIR` (+336 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 564 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **46 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `PlannedAllocation` connect `AllocationPlan` to `HierarchicalId`, `allocation_plan_domain_service.go`, `allocation_plan_rest_model.go`, `PropagateAsAppErrorWithNewMessage`, `go_pkg_testing`, `context.Context`, `portfolio_rest_mapping.go`?**
  _High betweenness centrality (0.011) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `Configuration`, `createBulkInsertPreparedStatement`, `app.go`, `go_pkg_testing`, `context.Context`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `AllocationPlanDomService`, `AllocationPlan`, `asset_rest_model.go`, `allocation_plan_rest_model.go`, `AssetDomService`, `PortfolioAllocation`, `asset_integration_test.go`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **Are the 73 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestGetAllocationPlans()` and `TestPostAllocationPlanForInsertion()`) actually correct?**
  _`deferCloseResponseBody()` has 73 INFERRED edges - model-reasoned connections that need verification._
- **What connects `build-dev.sh script`, `POSTGRES_DEV_DATA_DIR`, `build.sh script` to the rest of the system?**
  _341 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `T` be split into smaller, more focused modules?**
  _Cohesion score 0.11491935483870967 - nodes in this community are weakly interconnected._
- **Should `generate_report.py` be split into smaller, more focused modules?**
  _Cohesion score 0.06170598911070781 - nodes in this community are weakly interconnected._