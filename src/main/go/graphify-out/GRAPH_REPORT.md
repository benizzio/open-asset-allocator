# Graph Report - go  (2026-10-22)

## Corpus Check
- 123 files · ~65,563 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1126 nodes · 3449 edges · 44 communities (34 shown, 10 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 261 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ee116a9c`
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
- deferCloseResponseBody
- testing.T
- go_pkg_fmt
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- asset_integration_client_yahoo.go
- NewOrderedMapIterator
- PropagateAsAppErrorWithNewMessage
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- Asset
- AllocationDomService
- portfolio_observation_ownership_migration_integration_test.go
- assert_json_extension.go
- portfolio_rest_mapping.go
- AllocationStructure
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- App
- Portfolio
- AssertDBWithQueryMultipleRows
- BuildAppError
- portfolio_allocation_controller.go
- http_client.go
- rdbms_error.go

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 89 edges
2. `BuildCleanupFunctionBuilder()` - 40 edges
3. `App` - 35 edges
4. `PropagateAsAppErrorWithNewMessage()` - 34 edges
5. `PortfolioAllocation` - 32 edges
6. `Asset` - 30 edges
7. `IsZeroValue()` - 30 edges
8. `NewMapTree()` - 27 edges
9. `ExecuteDBQuery()` - 26 edges
10. `AssertDBWithQueryMultipleRows()` - 26 edges

## Surprising Connections (you probably didn't know these)
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `AllocationPlanRESTController` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go

## Import Cycles
- None detected.

## Communities (44 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (36): BuildAllocationPlanRESTController(), validateAllocationPlanPlannedDate(), bindAndValidatePortfolioSnapshot(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.Engine, github.com/gin-gonic/gin.HandlersChain, net/http.Server (+28 more)

### Community 1 - ".buildAppComponents"
Cohesion: 0.19
Nodes (12): BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository() (+4 more)

### Community 2 - "Adapter"
Cohesion: 0.12
Nodes (10): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, BuildQueryInTransaction(), Adapter, T, SQLTransactionalContext, withTransaction() (+2 more)

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.06
Nodes (41): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+33 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.29
Nodes (9): AssetIntegrationService, collectIntegrationServices(), go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext (+1 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (43): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+35 more)

### Community 6 - "context.Context"
Cohesion: 0.06
Nodes (62): BuildDivergenceAnalysisRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext(), buildDivergenceAnalysisContext() (+54 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.15
Nodes (26): assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds(), TestPostAllocationPlanUpdatesOwnedMetadata() (+18 more)

### Community 8 - "testing.T"
Cohesion: 0.07
Nodes (45): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), testing.T, BuildILikeSubstringPattern(), TestBuildILikeSubstringPattern(), ValueToString() (+37 more)

### Community 9 - "go_pkg_fmt"
Cohesion: 0.13
Nodes (15): assetRowScanner(), go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_golang_glog, go_pkg_github_com_lib_pq, go_pkg_net (+7 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.17
Nodes (15): go_pkg_net_url, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+7 more)

### Community 11 - "infra.go"
Cohesion: 0.06
Nodes (37): go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_path_filepath, go_pkg_sync (+29 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.18
Nodes (16): malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require (+8 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.18
Nodes (21): TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRejectsWhitespaceTags(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest(), registerObservationFixtureRestore() (+13 more)

### Community 15 - "asset_integration_client_yahoo.go"
Cohesion: 0.15
Nodes (14): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), BuildYahooFinanceAssetIntegrationClient(), YahooFinanceAssetIntegrationClient, TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU() (+6 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.24
Nodes (8): classifyObservationWriteError(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), IsUniqueConstraintViolation(), ToPointerSlice(), AssetRDBMSRepository, S

### Community 18 - "PortfolioAllocation"
Cohesion: 0.14
Nodes (12): BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AllocationHierarchyLevel, AssetsPerTicker, preparePortfolioAllocationTempInserts(), PortfolioAllocation (+4 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.15
Nodes (32): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), TestPostAssetWithoutExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted() (+24 more)

### Community 21 - "Asset"
Cohesion: 0.07
Nodes (49): BuildAssetRESTController(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+41 more)

### Community 22 - "AllocationDomService"
Cohesion: 0.38
Nodes (4): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 23 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.08
Nodes (26): database/sql.Row, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.Query, T, withParams(), processParamsForPostgreSQL(), processSQL(), ReturningIntIdSingleRowScanner() (+18 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 28 - "portfolio_rest_mapping.go"
Cohesion: 0.06
Nodes (63): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+55 more)

### Community 29 - "AllocationStructure"
Cohesion: 0.10
Nodes (15): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationStructure, HierarchicalId, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, database/sql/driver.Value (+7 more)

### Community 44 - "App"
Cohesion: 0.13
Nodes (16): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os, go_pkg_os_signal, go_pkg_syscall, os.Signal, Configuration (+8 more)

### Community 45 - "Portfolio"
Cohesion: 0.24
Nodes (6): BuildPortfolioDivergenceAnalysisAppService(), Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioRDBMSRepository

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.19
Nodes (30): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+22 more)

### Community 47 - "BuildAppError"
Cohesion: 0.25
Nodes (4): BuildAppError(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository, PortfolioAllocationRDBMSRepository

### Community 49 - "portfolio_allocation_controller.go"
Cohesion: 0.26
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_time

### Community 50 - "http_client.go"
Cohesion: 0.42
Nodes (8): net/http.Response, RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 53 - "rdbms_error.go"
Cohesion: 0.25
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 94 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `github.com/gin-gonic/gin.Context`, `.buildAppComponents`, `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `context.Context`, `go_pkg_fmt`, `infra.go`, `asset_integration_client_yahoo.go`, `portfolio_allocation_controller.go`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `allocation_plan_domain_service.go`, `go_pkg_fmt`, `BuildAppError`, `PropagateAsAppErrorWithNewMessage`, `PortfolioAllocation`, `BuildCleanupFunctionBuilder`, `portfolio_rest_mapping.go`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `AllocationStructure` to `portfolio_rest_mapping.go`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 87 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 87 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.056790123456790124 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.12 - nodes in this community are weakly interconnected._