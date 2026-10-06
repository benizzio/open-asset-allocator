# Graph Report - go  (2026-10-06)

## Corpus Check
- 119 files · ~57,801 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1068 nodes · 3155 edges · 46 communities (36 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 233 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `b72edb38`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- PortfolioAllocation
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- allocation_plan_domain_service.go
- ExternalAsset
- reflect.Type
- context.Context
- deferCloseResponseBody
- GinServer
- Asset
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- BuildCleanupFunctionBuilder
- golang_ext_util_concurrent.go
- NewOrderedMapIterator
- testing.T
- Adapter
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- rdbms_adapter.go
- BuildQuery
- T
- go_pkg_context
- assert_json_extension.go
- .buildAppComponents
- http_client.go
- PropagateAsAppErrorWithNewMessage
- HierarchicalId
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- go_pkg_net_http
- App
- Portfolio
- Configuration
- hierarchical_id_test.go
- go_pkg_fmt

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 76 edges
2. `App` - 35 edges
3. `PortfolioAllocation` - 32 edges
4. `PropagateAsAppErrorWithNewMessage()` - 31 edges
5. `Asset` - 30 edges
6. `BuildCleanupFunctionBuilder()` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 25 edges
9. `PlannedAllocation` - 23 edges
10. `AllocationPlan` - 23 edges

## Surprising Connections (you probably didn't know these)
- `AssetRESTController` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `BuildAssetRESTController()` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `DivergenceAnalysisRESTController` --references--> `PortfolioDivergenceAnalysisAppService`  [EXTRACTED]
  api/rest/divergence_analysis_controller.go → application/portfolio_divergence_analysis_service.go
- `BuildDivergenceAnalysisRESTController()` --references--> `PortfolioDivergenceAnalysisAppService`  [EXTRACTED]
  api/rest/divergence_analysis_controller.go → application/portfolio_divergence_analysis_service.go
- `mapToAllocationStructureDTS()` --references--> `AllocationStructure`  [EXTRACTED]
  api/rest/model/allocation_rest_model.go → domain/allocation.go

## Import Cycles
- None detected.

## Communities (46 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (46): BuildDivergenceAnalysisRESTController(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+38 more)

### Community 1 - "PortfolioAllocation"
Cohesion: 0.06
Nodes (54): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS() (+46 more)

### Community 2 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.27
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation (+1 more)

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.06
Nodes (41): BuildAllocationPlanRESTController(), cleanNilAllocations(), AllocationPlanManagementAppService, AllocationHierarchy, AllocationStructure, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService() (+33 more)

### Community 4 - "ExternalAsset"
Cohesion: 0.11
Nodes (24): YahooFinanceAssetIntegrationService, ExternalAsset, ExternalAssetQuote, AssetExternalSource, BuildYahooFinanceAssetIntegrationService(), extractLastClose(), mapToExternalAsset(), mapToExternalAssetQuote() (+16 more)

### Community 5 - "reflect.Type"
Cohesion: 0.05
Nodes (46): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_github_com_go_playground_validator_v10, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type (+38 more)

### Community 6 - "context.Context"
Cohesion: 0.10
Nodes (33): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+25 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.10
Nodes (36): go_pkg_io, TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches(), TestPostAllocationPlanValidation_MissingDetails() (+28 more)

### Community 8 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 9 - "Asset"
Cohesion: 0.08
Nodes (36): MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS(), mapToPlannedAllocationDTSs() (+28 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.14
Nodes (21): go_pkg_net_url, net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets() (+13 more)

### Community 11 - "infra.go"
Cohesion: 0.06
Nodes (38): go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os, go_pkg_path_filepath (+30 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.14
Nodes (16): TestParseAssetTextSearch(), go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_strconv, go_pkg_testing, github.com/go-ozzo/ozzo-dbx.Params (+8 more)

### Community 14 - "BuildCleanupFunctionBuilder"
Cohesion: 0.25
Nodes (26): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAsset(), assertPersistedAssetWithExternalData(), TestPostPortfolioAllocationHistoryFullMerge() (+18 more)

### Community 15 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "testing.T"
Cohesion: 0.13
Nodes (30): testing.T, postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios(), TestPostPortfolio() (+22 more)

### Community 18 - "Adapter"
Cohesion: 0.13
Nodes (9): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey (+1 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.27
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 21 - "rdbms_adapter.go"
Cohesion: 0.15
Nodes (12): assetRowScanner(), go_pkg_database_sql, go_pkg_github_com_lib_pq, database/sql.Row, database/sql.Rows, database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement() (+4 more)

### Community 22 - "BuildQuery"
Cohesion: 0.14
Nodes (10): BuildAssetRDBMSRepository(), BuildUniqueConstraintViolationError(), BuildQuery(), BuildQueryInTransaction(), T, IsUniqueConstraintViolation(), BuildILikeSubstringPattern(), ToPointerSlice() (+2 more)

### Community 23 - "T"
Cohesion: 0.10
Nodes (17): database/sql.Tx, github.com/go-ozzo/ozzo-dbx.Query, T, withParams(), processParamsForPostgreSQL(), processSQL(), T, QueryBuilder (+9 more)

### Community 24 - "go_pkg_context"
Cohesion: 0.13
Nodes (13): AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), go_pkg_context, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration (+5 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - ".buildAppComponents"
Cohesion: 0.21
Nodes (9): BuildAssetRESTController(), BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), github.com/go-ozzo/ozzo-dbx.Rows, RepositoryRDBMSAdapter, AllocationRDBMSRepository (+1 more)

### Community 27 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 28 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.25
Nodes (5): BuildAppError(), PropagateAsAppErrorWithNewMessage(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository, PortfolioAllocationRDBMSRepository

### Community 29 - "HierarchicalId"
Cohesion: 0.16
Nodes (6): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn(), IsNilPointer()

### Community 40 - "go_pkg_net_http"
Cohesion: 0.20
Nodes (8): malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_errors, go_pkg_github_com_gin_gonic_gin, go_pkg_github_com_golang_glog, go_pkg_net, go_pkg_net_http, go_pkg_net_http_httptest

### Community 41 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 42 - "Portfolio"
Cohesion: 0.21
Nodes (8): BuildPortfolioRESTController(), AllocationRepository, Portfolio, PortfolioRepository, BuildAllocationDomService(), AllocationDomService, BuildPortfolioDomService(), PortfolioDomService

### Community 43 - "Configuration"
Cohesion: 0.20
Nodes (12): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), Configuration, RDBMSConfiguration, YahooFinanceConfiguration, ReadConfig(), TestReadConfigAPIOnly() (+4 more)

### Community 44 - "hierarchical_id_test.go"
Cohesion: 0.24
Nodes (9): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, database/sql.NullString, StringPointerToNullString(), StringToNullString() (+1 more)

### Community 45 - "go_pkg_fmt"
Cohesion: 0.27
Nodes (5): go_pkg_database_sql_driver, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_golang_org_x_text_currency

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `go_pkg_net_http`, `GinServer`, `infra.go`, `Configuration`, `Adapter`, `go_pkg_context`, `.buildAppComponents`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `github.com/gin-gonic/gin.Context`, `PortfolioAllocation`, `ExecuteDBQuery`, `rdbms_adapter.go`, `BuildQuery`, `PropagateAsAppErrorWithNewMessage`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **Why does `GinServer` connect `GinServer` to `go_pkg_net_http`, `App`, `Configuration`?**
  _High betweenness centrality (0.028) - this node is a cross-community bridge._
- **Are the 74 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestGetAllocationPlans()` and `TestPostAllocationPlanForInsertion()`) actually correct?**
  _`deferCloseResponseBody()` has 74 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.06486486486486487 - nodes in this community are weakly interconnected._
- **Should `PortfolioAllocation` be split into smaller, more focused modules?**
  _Cohesion score 0.06368011847463902 - nodes in this community are weakly interconnected._