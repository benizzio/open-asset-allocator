# Graph Report - go  (2026-10-06)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 1067 nodes · 3150 edges · 40 communities (30 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 233 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `557288c1`
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
- testing.T
- App
- AllocationPlan
- asset_integration_test.go
- infra_db.go
- golang_ext_tree_test.go
- go_pkg_testing
- assert_db.go
- golang_ext_util_concurrent.go
- NewOrderedMapIterator
- portfolio_integration_test.go
- Adapter
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- rdbms_adapter.go
- rdbms_query.go
- T
- T
- assert_json_extension.go
- QueryExecutor[T]
- http_client.go
- go_pkg_database_sql
- manage_db.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 76 edges
2. `App` - 35 edges
3. `PortfolioAllocation` - 32 edges
4. `PropagateAsAppErrorWithNewMessage()` - 31 edges
5. `Asset` - 30 edges
6. `BuildCleanupFunctionBuilder()` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 24 edges
9. `AllocationPlan` - 23 edges
10. `PlannedAllocation` - 23 edges

## Surprising Connections (you probably didn't know these)
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `MapToAsset()` --references--> `Asset`  [EXTRACTED]
  api/rest/model/asset_rest_model.go → domain/asset.go
- `MapToAssetDTS()` --references--> `Asset`  [EXTRACTED]
  api/rest/model/asset_rest_model.go → domain/asset.go
- `MapToAssetDTS()` --calls--> `ParseableInt64`  [EXTRACTED]
  api/rest/model/asset_rest_model.go → langext/golang_ext_parseable_int64.go
- `MapToAssetDTSs()` --references--> `Asset`  [EXTRACTED]
  api/rest/model/asset_rest_model.go → domain/asset.go

## Import Cycles
- None detected.

## Communities (40 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.05
Nodes (72): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+64 more)

### Community 1 - "PortfolioAllocation"
Cohesion: 0.05
Nodes (45): MapToPortfolioObservationTimestamp(), BuildPortfolioAllocationRESTController(), BuildPortfolioRESTController(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AllocationRepository (+37 more)

### Community 2 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.06
Nodes (42): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildDivergenceAnalysisRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, AllocationPlanRepository (+34 more)

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.05
Nodes (42): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+34 more)

### Community 4 - "ExternalAsset"
Cohesion: 0.06
Nodes (40): YahooFinanceAssetIntegrationService, ExternalAsset, ExternalAssetQuote, AssetExternalSource, BuildYahooFinanceAssetIntegrationService(), extractLastClose(), mapToExternalAsset(), mapToExternalAssetQuote() (+32 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.08
Nodes (37): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+29 more)

### Community 7 - "testing.T"
Cohesion: 0.10
Nodes (44): go_pkg_net_http, testing.T, TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches() (+36 more)

### Community 8 - "App"
Cohesion: 0.07
Nodes (27): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os (+19 more)

### Community 9 - "AllocationPlan"
Cohesion: 0.10
Nodes (30): MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS(), mapToPlannedAllocationDTSs() (+22 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.11
Nodes (35): go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_moby_moby_api_types_container, go_pkg_io, go_pkg_net_url, net/http.Response, contextKey, getAssetsByTextSearch(), postAssetForValidationFailure() (+27 more)

### Community 11 - "infra_db.go"
Cohesion: 0.07
Nodes (27): go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M (+19 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.10
Nodes (22): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), malformedQueryBindingTarget (+14 more)

### Community 14 - "assert_db.go"
Cohesion: 0.21
Nodes (27): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+19 more)

### Community 15 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "portfolio_integration_test.go"
Cohesion: 0.17
Nodes (19): go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_strconv, postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios() (+11 more)

### Community 18 - "Adapter"
Cohesion: 0.17
Nodes (6): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, Adapter, SQLTransactionalContext, withTransaction()

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.38
Nodes (14): TestGetKnownAssetsIncludesPersistedExternalData(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange(), TestGetExternalAssetQuoteRejectsMismatchedProviderTicker(), TestGetExternalAssetQuoteSuccess(), addAssetExternalDataRestoreCleanup() (+6 more)

### Community 21 - "rdbms_adapter.go"
Cohesion: 0.21
Nodes (9): database/sql.Stmt, database/sql.Tx, buildPingContext(), BuildQueryInTransaction(), createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), T, prepareBulkInsertValues() (+1 more)

### Community 22 - "rdbms_query.go"
Cohesion: 0.18
Nodes (8): assetRowScanner(), database/sql.Row, database/sql.Rows, BuildILikeSubstringPattern(), ReturningIntIdRowScanner(), ReturningIntIdSingleRowScanner(), SingleRowScanner, SQLTransactionalQueryExecutor[T]

### Community 23 - "T"
Cohesion: 0.38
Nodes (6): github.com/go-ozzo/ozzo-dbx.Query, T, withParams(), QueryBuilder, QueryBuilder[T], QueryExecutor

### Community 24 - "T"
Cohesion: 0.35
Nodes (6): processParamsForPostgreSQL(), processSQL(), T, SQLTransactionalQueryBuilder, SQLTransactionalQueryBuilder[T], SQLTransactionalQueryExecutor

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "QueryExecutor[T]"
Cohesion: 0.36
Nodes (3): github.com/go-ozzo/ozzo-dbx.Rows, QueryExecutor[T], RowScanner

### Community 27 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 28 - "go_pkg_database_sql"
Cohesion: 0.29
Nodes (4): go_pkg_database_sql, NullTime, contextKey, TransactionalContext

### Community 29 - "manage_db.go"
Cohesion: 0.48
Nodes (4): github.com/go-ozzo/ozzo-dbx.Params, createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `contextKey`, `malformedQueryBindingTarget`, `MapIterator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `PortfolioAllocation`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `ExternalAsset`, `reflect.Type`, `infra_db.go`, `Adapter`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **Why does `Asset` connect `PortfolioAllocation` to `github.com/gin-gonic/gin.Context`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `ExternalAsset`, `context.Context`, `AllocationPlan`, `asset_integration_test.go`, `rdbms_query.go`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `App`, `rdbms_adapter.go`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **Are the 74 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestGetAllocationPlans()` and `TestPostAllocationPlanForInsertion()`) actually correct?**
  _`deferCloseResponseBody()` has 74 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `contextKey` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.050853548966756514 - nodes in this community are weakly interconnected._
- **Should `PortfolioAllocation` be split into smaller, more focused modules?**
  _Cohesion score 0.05049088359046283 - nodes in this community are weakly interconnected._