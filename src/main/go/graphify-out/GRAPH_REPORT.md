# Graph Report - go  (2026-10-20)

## Corpus Check
- 120 files · ~64,082 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1110 nodes · 3377 edges · 46 communities (35 shown, 11 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 243 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `d462dbce`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- .buildAppComponents
- Adapter
- allocation_plan_domain_service.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- reflect.Type
- context.Context
- testing.T
- PlannedAllocation
- BuildAppError
- asset_integration_test.go
- infra_db.go
- golang_ext_tree_test.go
- go_pkg_testing
- BuildCleanupFunctionBuilder
- App
- NewOrderedMapIterator
- go_pkg_strings
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- Asset
- go_pkg_fmt
- T
- portfolio_observation_ownership_migration_integration_test.go
- assert_json_extension.go
- error.go
- T
- portfolio_rest_mapping.go
- HierarchicalId
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- manage_db.go
- GinServer
- PropagateAsAppErrorWithNewMessage
- portfolio_allocation_controller.go
- portfolio_allocation_integration_test.go
- rdbms_adapter.go

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 85 edges
2. `BuildCleanupFunctionBuilder()` - 37 edges
3. `App` - 35 edges
4. `PropagateAsAppErrorWithNewMessage()` - 33 edges
5. `PortfolioAllocation` - 32 edges
6. `Asset` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 26 edges
9. `PlannedAllocation` - 25 edges
10. `ExecuteDBQuery()` - 25 edges

## Surprising Connections (you probably didn't know these)
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `AssetRESTController` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `BuildAssetRESTController()` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go

## Import Cycles
- None detected.

## Communities (46 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (42): BuildAllocationPlanRESTController(), validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), bindAndValidatePortfolioSnapshot(), BuildPortfolioRESTController(), AllocationPlanManagementAppService (+34 more)

### Community 1 - ".buildAppComponents"
Cohesion: 0.20
Nodes (12): BuildAssetRESTController(), BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildPortfolioRepository() (+4 more)

### Community 2 - "Adapter"
Cohesion: 0.13
Nodes (8): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, Adapter, SQLTransactionalContext, withTransaction(), contextKey, TransactionalContext

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.07
Nodes (43): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+35 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.27
Nodes (9): AssetIntegrationService, collectIntegrationServices(), go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext (+1 more)

### Community 5 - "reflect.Type"
Cohesion: 0.05
Nodes (48): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_github_com_go_playground_validator_v10, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type (+40 more)

### Community 6 - "context.Context"
Cohesion: 0.06
Nodes (66): BuildDivergenceAnalysisRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext(), buildDivergenceAnalysisContext() (+58 more)

### Community 7 - "testing.T"
Cohesion: 0.10
Nodes (44): testing.T, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+36 more)

### Community 8 - "PlannedAllocation"
Cohesion: 0.26
Nodes (9): GetPlanType(), PlanType, AllocationPlan, PlannedAllocation, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow(), mapPlannedAllocationRow(), mapPlannedAllocationRows() (+1 more)

### Community 9 - "BuildAppError"
Cohesion: 0.25
Nodes (7): allocationPlanOwnershipValidationError(), plannedAllocationValidationError(), rejectPersistedDetailsOnNewPlan(), validatePlannedAllocationOwnership(), BuildAppError(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 10 - "asset_integration_test.go"
Cohesion: 0.11
Nodes (25): go_pkg_net_url, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+17 more)

### Community 11 - "infra_db.go"
Cohesion: 0.07
Nodes (27): go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M (+19 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.09
Nodes (33): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx (+25 more)

### Community 14 - "BuildCleanupFunctionBuilder"
Cohesion: 0.19
Nodes (33): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+25 more)

### Community 15 - "App"
Cohesion: 0.21
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), App

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "go_pkg_strings"
Cohesion: 0.22
Nodes (7): go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, go_pkg_strings, go_pkg_time

### Community 18 - "PortfolioAllocation"
Cohesion: 0.17
Nodes (8): BuildPortfolioAllocationRepository(), preparePortfolioAllocationTempInserts(), PortfolioAllocation, PortfolioAllocationRepository, PortfolioObservationTimestamp, BuildPortfolioAllocationDomService(), PortfolioAllocationDomService, PortfolioAllocationRDBMSRepository

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.26
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 21 - "Asset"
Cohesion: 0.07
Nodes (41): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+33 more)

### Community 22 - "go_pkg_fmt"
Cohesion: 0.18
Nodes (8): go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_net, go_pkg_path_filepath, ConfigLogger(), main()

### Community 23 - "T"
Cohesion: 0.11
Nodes (16): github.com/go-ozzo/ozzo-dbx.Query, T, withParams(), processParamsForPostgreSQL(), processSQL(), T, QueryBuilder, QueryBuilder[T] (+8 more)

### Community 24 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.53
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 25 - "assert_json_extension.go"
Cohesion: 0.12
Nodes (21): go_pkg_encoding_json, go_pkg_regexp, net/http.Response, RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON() (+13 more)

### Community 26 - "error.go"
Cohesion: 0.23
Nodes (8): mapToExternalAssetQuote(), classifyObservationWriteError(), BuildAppErrorFormatted(), BuildUniqueConstraintViolationError(), logError(), newAppError(), IsUniqueConstraintViolation(), UniqueConstraintViolationError

### Community 28 - "portfolio_rest_mapping.go"
Cohesion: 0.09
Nodes (47): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+39 more)

### Community 29 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 40 - "manage_db.go"
Cohesion: 0.48
Nodes (4): github.com/go-ozzo/ozzo-dbx.Params, createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 43 - "GinServer"
Cohesion: 0.06
Nodes (33): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), BuildYahooFinanceAssetIntegrationClient(), YahooFinanceAssetIntegrationClient (+25 more)

### Community 47 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.18
Nodes (9): BuildAssetRDBMSRepository(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, ToPointerSlice(), AssetRDBMSRepository, PortfolioRDBMSRepository (+1 more)

### Community 49 - "portfolio_allocation_controller.go"
Cohesion: 0.24
Nodes (11): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin (+3 more)

### Community 50 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (17): portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest(), registerObservationFixtureRestore(), TestGetAvailableHistoryObservations(), TestGetPortfolioAllocationHistory(), TestGetPortfolioAllocationHistoryForObservationTimestamp(), TestGetPortfolioAllocationHistoryForObservationTimestampNoneFound() (+9 more)

### Community 52 - "rdbms_adapter.go"
Cohesion: 0.13
Nodes (13): assetRowScanner(), go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_github_com_lib_pq, database/sql.Row, database/sql.Rows, database/sql.Stmt, createBulkInsertPreparedStatement() (+5 more)

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `.buildAppComponents`, `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `context.Context`, `infra_db.go`, `GinServer`, `portfolio_allocation_controller.go`, `go_pkg_strings`, `go_pkg_fmt`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `github.com/gin-gonic/gin.Context`, `PlannedAllocation`, `BuildAppError`, `PropagateAsAppErrorWithNewMessage`, `PortfolioAllocation`, `rdbms_adapter.go`, `ExecuteDBQuery`, `error.go`, `portfolio_rest_mapping.go`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `HierarchicalId` to `PlannedAllocation`, `go_pkg_strings`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 83 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 83 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.06220095693779904 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.12666666666666668 - nodes in this community are weakly interconnected._