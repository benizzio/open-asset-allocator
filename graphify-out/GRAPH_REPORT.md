# Graph Report - open-asset-allocator  (2026-09-26)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 818 nodes · 1159 edges · 91 communities (46 shown, 45 thin omitted)
- Extraction: 96% EXTRACTED · 4% INFERRED · 0% AMBIGUOUS · INFERRED: 50 edges (avg confidence: 0.88)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `1a2fda2d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- portfolio-allocation-history-management.e2e.spec.ts
- portfolio-allocation-plan-management.e2e.spec.ts
- generate_report.py
- Comparative E2E Framework Research
- portfolio-allocation-map.e2e.spec.ts
- Three scoped Graphify graphs
- package.json
- Allocation Map View
- e2e.sh
- portfolio-visualization.e2e.spec.ts
- Three-Scope Graph Selection
- Portfolio Holdings Editor
- asset
- portfolio-editing.e2e.spec.ts
- prune-sql-stubs.py
- asset-management.e2e.spec.ts
- BONDS (60% slice)
- validate-playwright-version.mjs
- Asset Allocation Donut Chart
- Open Asset Allocator Architecture
- test
- compilerOptions
- Asset Allocation Donut Chart
- Portfolio History Detail View
- ingest-ghostfolio-activity.sql
- Base Application Service
- database.ts
- Portfolio Card Grid
- migration-[5]-allocation_plan_improvements.sql
- fixtures.ts
- Graphify Pipeline
- Verify Assets Against Planned Allocation Percentages
- Asset Class Allocation
- portfolio_allocation_fact_insertion
- portfolio_allocation_fact_insertion
- Repository Scope Routing
- @playwright/test
- migration-[2]-initial_fixes.sql
- Incremental Graph Re-Extraction
- destroy.sh
- macos-provisioning.sh
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
- migration-[11]-portfolio_allocation_fact_time_frame_tag_substitution.sql
- portfolio_allocation_obs_time
- migration-[9]-portfolio_allocation_timestamp_dimension_adjustments.sql
- start.sh
- Property Graph Exports
- E2E Diagnostics Artifact
- build.sh
- go-dependency-graph.sh
- migrate.sh
- delete-portfolio.sql
- move-portfolio-id.sql
- duckdb-cli-build.sh
- migration-[10]-portfolio_allocation_obs_time_tag_unique.sql
- migration-[12]-planned_allocation_structural_id_renaming.sql
- migration-[14]-asset_external_data.sql
- init-db.sh
- graphify-integration.sh
- test.sh
- Background Folder Watcher
- Post-Commit Graph Rebuild
- Saved Query Result Memory
- Media Transcription Pipeline
- Automatic Non-Draft CodeRabbit Review Policy
- Go Lint Workflow
- AI-Generated Code Documentation and Authorship
- Atomic Persistence-Verified E2E Testing
- Problem-Domain Literacy
- Weighted Project Fit Score
- Split Development Stack
- Local Split-Application E2E Overlay
- DuckDB CLI Service
- Flyway Migration Service
- Flyway Migration Standards
- External Integration CI Gate

## God Nodes (most connected - your core abstractions)
1. `E2eDatabase` - 16 edges
2. `Comparative E2E Framework Research` - 15 edges
3. `main()` - 14 edges
4. `getDoughnutCenterPoint()` - 13 edges
5. `getDoughnutSlicePoint()` - 12 edges
6. `@playwright/test` - 12 edges
7. `expectLatestCanvasTextSet()` - 11 edges
8. `test` - 11 edges
9. `compilerOptions` - 11 edges
10. `clickCanvasPoint()` - 10 edges

## Surprising Connections (you probably didn't know these)
- `Allocation Planning` --semantically_similar_to--> `Allocation Plan`  [INFERRED] [semantically similar]
  docs/readme-features.md → README.md
- `Convergence Analysis and Planning` --semantically_similar_to--> `Convergence Planning`  [INFERRED] [semantically similar]
  docs/readme-features.md → README.md
- `Divergence Analysis` --semantically_similar_to--> `Divergence Analysis`  [INFERRED] [semantically similar]
  docs/readme-features.md → README.md
- `Customizable Fractal Hierarchy` --semantically_similar_to--> `Fractal Portfolio Hierarchy`  [INFERRED] [semantically similar]
  docs/readme-features.md → README.md
- `Historical Portfolio Snapshots` --semantically_similar_to--> `Portfolio History`  [INFERRED] [semantically similar]
  docs/readme-features.md → README.md

## Import Cycles
- None detected.

## Hyperedges (group relationships)
- **Remainder host-agent semantic refresh safeguards** — docs_graphify_scopes_opencode_host_agent_refresh, docs_graphify_scopes_root_candidate_detection, docs_graphify_scopes_semantic_provenance_normalization, docs_graphify_scopes_scoped_build_configurations [EXTRACTED 1.00]
- **Independent frontend, backend, and remainder graph partition** — docs_graphify_scopes_frontend_graph, docs_graphify_scopes_backend_graph, docs_graphify_scopes_remainder_graph [EXTRACTED 1.00]
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
- **Portfolio Observation Holding Fields** — docs_images_portfolio_history_management_asset_identity, docs_images_portfolio_history_management_asset_classification, docs_images_portfolio_history_management_cash_reserve_designation, docs_images_portfolio_history_management_position_quantity, docs_images_portfolio_history_management_market_price, docs_images_portfolio_history_management_total_market_value [EXTRACTED 1.00]
- **Containerized E2E Execution Topology** — src_main_docker_docker_compose_e2e_e2e_database, src_main_docker_docker_compose_e2e_e2e_migration_engine, src_main_docker_docker_compose_e2e_playwright_runner, src_main_docker_docker_compose_e2e_ci_immutable_monolith, _github_workflows_e2e_e2e_tests_workflow [INFERRED 0.95]

## Communities (91 total, 45 thin omitted)

### Community 0 - "portfolio-allocation-history-management.e2e.spec.ts"
Cohesion: 0.06
Nodes (54): CanvasPoint, CanvasTextRecorder, clickCanvasPoint(), expectChartTooltip(), expectLatestCanvasPatternState(), expectLatestCanvasTextContains(), expectLatestCanvasTextSet(), findDoughnutSlicePointByTooltip() (+46 more)

### Community 1 - "portfolio-allocation-plan-management.e2e.spec.ts"
Cohesion: 0.05
Nodes (52): seedPortfolioHistoryModificationData(), readDatabaseSnapshot(), seedAllocationMapData(), addAssetAllocationRow(), addClassAllocationRow(), DEFAULT_ALLOCATION_STRUCTURE, expectAllocationPlan(), expectAllocationPlanManagement() (+44 more)

### Community 2 - "generate_report.py"
Cohesion: 0.06
Nodes (55): Any, _anchor(), _base_result_name(), _build_output_mapping(), _category_aliases(), _collect_extra_fields(), visit(), _compact_summary() (+47 more)

### Community 3 - "Comparative E2E Framework Research"
Cohesion: 0.05
Nodes (49): Dependabot Dependency Updates, Renovate Runtime Coordination, Go and Air Image Version Coordination, Application Node.js LTS Runtime Coordination, Playwright and E2E Node Type Major Alignment, Runtime Version Coordination, Portfolio History Form TODOs, Allocation Planning (+41 more)

### Community 4 - "portfolio-allocation-map.e2e.spec.ts"
Cohesion: 0.08
Nodes (33): ASSET_DATA, DatabaseSnapshot, DEFAULT_ALLOCATION_STRUCTURE, expandRootAndAssertChildren(), expectAllocationMapShell(), expectAnalysisTable(), ExpectedDivergenceNode, expectNodeRow() (+25 more)

### Community 5 - "Three scoped Graphify graphs"
Cohesion: 0.10
Nodes (29): Absolute module scan paths, asset.external_data migration retrieval evidence, Backend Graphify graph, Build configuration file retrieval miss, Code-only graph bootstrap, Committed graph integrity check, Operator-configured Graphify semantic backend, Disposable Graphify integration fixture (+21 more)

### Community 6 - "package.json"
Cohesion: 0.09
Nodes (21): pg, @types/node, @types/pg, typescript, author, devDependencies, pg, @playwright/test (+13 more)

### Community 7 - "Allocation Map View"
Cohesion: 0.12
Nodes (20): $60,000 Total Market Value, 60/40 Portfolio Classic - Example - 20260210-230012, Actual vs Planned Market Value Comparison, Allocation Map View, Allocation Plan Selector, BIL Underweight: -$800 (-2.96%), Bonds Underweight: -$9,000 (-15%), Directional Divergence Bars (+12 more)

### Community 8 - "e2e.sh"
Cohesion: 0.24
Nodes (17): build_images(), capture_logs(), cleanup_stack(), compose(), compose_all(), compose_debug(), fail(), main() (+9 more)

### Community 9 - "portfolio-visualization.e2e.spec.ts"
Cohesion: 0.20
Nodes (19): DEFAULT_ALLOCATION_STRUCTURE, EMPTY_PORTFOLIO_NAMES, expectAllocationMap(), expectAllocationPlan(), expectPortfolioHistory(), expectPortfolioList(), expectPortfolioNavigation(), expectPortfolioShell() (+11 more)

### Community 10 - "Three-Scope Graph Selection"
Cohesion: 0.12
Nodes (19): CodeRabbit Review Path Filters, Graphify Build Manifest Inclusion Patterns, Graphify Output Directory Exclusion Patterns, Vendored Graphify Skill Exclusion Pattern, Backend Graph Corpus, Disconnected Flyway SQL Stubs after Root Update, Frontend Graph Corpus, Graphify-First Codebase Navigation (+11 more)

### Community 11 - "Portfolio Holdings Editor"
Cohesion: 0.12
Nodes (19): Manage Portfolio Allocation Data Panel, Allocation Map Tab, Allocation Plan Tab, Asset Classification, Asset Identity and Description, Bonds Asset Class, Cash Reserve Designation, Add and Remove Holding Controls (+11 more)

### Community 12 - "asset"
Cohesion: 0.21
Nodes (13): asset_price_last_market_data, asset_ticker_market_data_source, allocation_plan, allocation_plan_unit, asset, asset_value_fact, asset_market_data_source, asset_price_last_market_data (+5 more)

### Community 13 - "portfolio-editing.e2e.spec.ts"
Cohesion: 0.16
Nodes (15): DEFAULT_ALLOCATION_STRUCTURE, expectEditPortfolio(), ExpectedPortfolio, expectPortfolioList(), expectPortfolioNavigation(), expectPortfolioShell(), expectRootShell(), expectRoute() (+7 more)

### Community 14 - "prune-sql-stubs.py"
Cohesion: 0.13
Nodes (15): graphify_cluster, graphify_export, json, networkx_readwrite, pathlib, re, _main(), Remove disconnected Flyway SQL parser stubs from the remainder graph. Graphify… (+7 more)

### Community 15 - "asset-management.e2e.spec.ts"
Cohesion: 0.16
Nodes (8): Asset, AssetRow, expectAssetEditor(), expectIconOnlyButton(), expectNewAssetForm(), expectSupersededAssetResponseIgnored(), navigateWithinApp(), PERSISTED_EXTERNAL_DATA

### Community 16 - "BONDS (60% slice)"
Cohesion: 0.14
Nodes (15): 60/40 Portfolio Classic - Example - 20260210-230012, Allocation Plan Management Interface, ARCA:BIL - SPDR Bloomberg 1-3 Month T-Bill ETF (40%), ARCA:EWZ - iShares MSCI Brazil ETF (5%), ARCA:SPY - SPDR S&P 500 ETF Trust (45%), ARCA:STIP - iShares 0-5 Year TIPS Bond ETF (10%), BONDS (60% slice), Cash Reserve Designation (+7 more)

### Community 17 - "validate-playwright-version.mjs"
Cohesion: 0.16
Nodes (13): ref_node_fs, argumentsByName, assertEqual(), assertNodeMajor(), defaultDockerfile, fail(), nodeTypeVersions, packageDirectory (+5 more)

### Community 18 - "Asset Allocation Donut Chart"
Cohesion: 0.15
Nodes (14): Allocation Map Tab, Allocation Plan Tab, ARCA:EWZ Allocation 4.55%, ARCA:SPY Allocation 68.18%, Asset Allocation Donut Chart, Assets for STOCKS Level, February 2026 Portfolio Snapshot, My Portfolio Example (+6 more)

### Community 19 - "Open Asset Allocator Architecture"
Cohesion: 0.17
Nodes (12): Root AGENTS Instructions for Copilot, Fractal Long-Term Asset Allocation Strategies, Backend Serves Frontend Static Assets in Production, Go Gin Backend, Hybrid HTMX Lazy-Loading SPA Frontend, Open Asset Allocator Architecture, PostgreSQL Flyway and DuckDB Data Layer, Graphify scopes (+4 more)

### Community 20 - "test"
Cohesion: 0.17
Nodes (7): src_test_e2e_support_fixtures_expect, test, FIRST, SECOND, DEFAULT_ALLOCATION_STRUCTURE, Portfolio, PortfolioRow

### Community 21 - "compilerOptions"
Cohesion: 0.15
Nodes (12): compilerOptions, allowImportingTsExtensions, esModuleInterop, forceConsistentCasingInFileNames, lib, module, moduleResolution, noEmit (+4 more)

### Community 22 - "Asset Allocation Donut Chart"
Cohesion: 0.22
Nodes (11): 60/40 Portfolio Classic - Example - 20260210-230012, Allocation Map Tab, Allocation Plan Tab, Asset Allocation Donut Chart, Asset Classes Level, Bonds 60 Percent Allocation, Edit Portfolio Action, My Portfolio Example - 20260210-230012 (+3 more)

### Community 23 - "Portfolio History Detail View"
Cohesion: 0.22
Nodes (11): Active Portfolio Tab, Allocation Map Tab, Allocation Plan Tab, Bonds: 45%, Class-Level Asset Allocation Donut Chart, Edit Portfolio Control, Expanded 202602 History Panel, My Portfolio Example - 20260210-230012 (+3 more)

### Community 24 - "ingest-ghostfolio-activity.sql"
Cohesion: 0.31
Nodes (10): pgsql.asset_market_data_source, asset_dimension_mapping, asset_insertion, asset_market_data_source_insertion, ghostf_activity, ghostf_symbol_aggegation, pgsql.asset, yahoo_asset_list (+2 more)

### Community 25 - "Base Application Service"
Cohesion: 0.27
Nodes (11): Base Application Service, Base Migration Engine, Base PostgreSQL Service, CI E2E Overlay, Immutable E2E Monolith, E2E PostgreSQL Service, E2E Migration Engine, Containerized Playwright Runner (+3 more)

### Community 26 - "database.ts"
Cohesion: 0.27
Nodes (6): createE2eDatabase(), E2eDatabase, FLYWAY_HISTORY_TABLES, parsePort(), quoteIdentifier(), requiredEnvironment()

### Community 27 - "Portfolio Card Grid"
Cohesion: 0.22
Nodes (10): Dark Theme, Global All Assets Portfolio, My Portfolio Example, My Portfolio Example - 20260210-230012, New Portfolio Action, New Portfolio Focus State, Open Asset Allocator Brand, Portfolio Card Grid (+2 more)

### Community 28 - "migration-[5]-allocation_plan_improvements.sql"
Cohesion: 0.28
Nodes (7): allocation_plan.create_timestamp TIMESTAMP DEFAULT now() IF NOT EXISTS, allocation_plan.structure JSONB USING structure::jsonb, allocation_plan_unit.slice renamed slice_size_percentage NUMERIC(5,2), allocation_plan_unit.structural_id text[] USING structural_id::text[], slice_percentage_ck CHECK slice_size_percentage between 0 and 100, allocation_plan_unit.slice_size_percentage NUMERIC(10,5) then divide by 100 then NUMERIC(6,5), slice_percentage_ck replaced with CHECK slice_size_percentage between 0 and 1

### Community 29 - "fixtures.ts"
Cohesion: 0.39
Nodes (6): TestFixtures, WorkerFixtures, delay(), hasSuccessfulResponse(), waitFor(), waitForE2eReadiness()

### Community 30 - "Graphify Pipeline"
Cohesion: 0.29
Nodes (7): Existing Graph Fast Path, Graph Health Diagnostics, Graphify Pipeline, Graphify Honesty Rules, Graph Query, Path, and Explain, Semantic Subagent Extraction, Structural AST Extraction

### Community 31 - "Verify Assets Against Planned Allocation Percentages"
Cohesion: 0.43
Nodes (7): Create or Modify Allocation Plan, External Cash Inflow, Big or Disruptive Market Fluctuations?, Verify Assets Against Planned Allocation Percentages, Wait for Planned Interval, Move Resources to Fit the Allocation Plan, Scenario Changed?

### Community 32 - "Asset Class Allocation"
Cohesion: 0.33
Nodes (7): Asset Class Allocation, Bonds Allocation 45 Percent, Monthly Portfolio Snapshot 202602, Portfolio Detail Screen, Portfolio Workspace Navigation, Stocks Allocation 55 Percent, Total Market Value 60000

### Community 33 - "portfolio_allocation_fact_insertion"
Cohesion: 0.48
Nodes (6): asset_dimension_mapping, asset_insertion, portfolio_allocation_fact_insertion, pgsql.asset, pgsql.portfolio_allocation_obs_time, sws_summary

### Community 34 - "portfolio_allocation_fact_insertion"
Cohesion: 0.48
Nodes (6): asset_dimension_mapping, asset_insertion, portfolio_allocation_fact_insertion, pgsql.asset, pgsql.portfolio_allocation_obs_time, sws_summary

### Community 35 - "Repository Scope Routing"
Cohesion: 0.33
Nodes (6): Explicit Module Graph Query, Graphify Skill Invocation, Mixed Scope Boundary Verification, Native Per-Scope Graph Maintenance, Repository Scope Routing, Unknown Scope Write Pause

### Community 36 - "@playwright/test"
Cohesion: 0.33
Nodes (5): ref_node_path, ref_node_url, @playwright/test, defaultArtifactsDirectory, packageDirectory

### Community 37 - "migration-[2]-initial_fixes.sql"
Cohesion: 0.40
Nodes (4): asset.name DROP NOT NULL, asset_ticker_uk UNIQUE (asset.ticker), asset_value_fact_pk PRIMARY KEY (class, cash_reserve, asset_id, time_frame_tag), asset_value_fact.time_frame_tag text NOT NULL snapshot classifier

### Community 38 - "Incremental Graph Re-Extraction"
Cohesion: 0.50
Nodes (4): Graphify URL Ingestion, Cluster-Only Graph Refresh, Incremental Graph Re-Extraction, Replace-on-Re-Extract Graph Merge

### Community 39 - "destroy.sh"
Cohesion: 0.50
Nodes (3): POSTGRES_DATA_DIR, POSTGRES_DEV_DATA_DIR, destroy.sh script

### Community 40 - "macos-provisioning.sh"
Cohesion: 0.50
Nodes (3): NVM_DIR, PATH, macos-provisioning.sh script

### Community 42 - "stop.sh"
Cohesion: 0.50
Nodes (3): POSTGRES_DATA_DIR, POSTGRES_DEV_DATA_DIR, stop.sh script

### Community 43 - "validate-node-version.sh"
Cohesion: 0.83
Nodes (3): fail(), read_node_image(), validate-node-version.sh script

### Community 44 - "Graph Query Traversal"
Cohesion: 0.67
Nodes (3): Graphify MCP Graph Server, Constrained Graph Vocabulary Expansion, Graph Query Traversal

### Community 45 - "Semantic Extraction Contract"
Cohesion: 0.67
Nodes (3): Extraction Confidence Rubric, Deterministic Full-Path Node Identity, Semantic Extraction Contract

### Community 46 - "Go External Integration Tests"
Cohesion: 1.00
Nodes (3): Go External Integration Tests, Go Test Workflow, Go Unit Tests

## Knowledge Gaps
- **252 isolated node(s):** `CanvasPoint`, `CanvasTextRecorder`, `ExpectedObservationRow`, `NavigationOptionName`, `PersistedAllocation` (+247 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 381 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **45 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `@playwright/test` connect `@playwright/test` to `portfolio-allocation-history-management.e2e.spec.ts`, `portfolio-allocation-plan-management.e2e.spec.ts`, `portfolio-allocation-map.e2e.spec.ts`, `package.json`, `portfolio-visualization.e2e.spec.ts`, `portfolio-editing.e2e.spec.ts`, `asset-management.e2e.spec.ts`, `test`, `fixtures.ts`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Why does `E2eDatabase` connect `database.ts` to `portfolio-allocation-history-management.e2e.spec.ts`, `portfolio-allocation-plan-management.e2e.spec.ts`, `portfolio-allocation-map.e2e.spec.ts`, `portfolio-visualization.e2e.spec.ts`, `portfolio-editing.e2e.spec.ts`, `asset-management.e2e.spec.ts`, `test`, `fixtures.ts`?**
  _High betweenness centrality (0.013) - this node is a cross-community bridge._
- **Why does `test` connect `test` to `portfolio-allocation-history-management.e2e.spec.ts`, `portfolio-allocation-plan-management.e2e.spec.ts`, `portfolio-allocation-map.e2e.spec.ts`, `portfolio-visualization.e2e.spec.ts`, `portfolio-editing.e2e.spec.ts`, `asset-management.e2e.spec.ts`, `database.ts`, `fixtures.ts`?**
  _High betweenness centrality (0.010) - this node is a cross-community bridge._
- **What connects `CanvasPoint`, `CanvasTextRecorder`, `ExpectedObservationRow` to the rest of the system?**
  _252 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `portfolio-allocation-history-management.e2e.spec.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.06459627329192547 - nodes in this community are weakly interconnected._
- **Should `portfolio-allocation-plan-management.e2e.spec.ts` be split into smaller, more focused modules?**
  _Cohesion score 0.051203277009728626 - nodes in this community are weakly interconnected._
- **Should `generate_report.py` be split into smaller, more focused modules?**
  _Cohesion score 0.06493506493506493 - nodes in this community are weakly interconnected._