# Graph Report - go  (2026-10-19)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 1076 nodes · 3213 edges · 41 communities (31 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 237 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `9d3531c6`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- .buildAppComponents
- Portfolio
- allocation_plan_domain_service.go
- portfolio_allocation_repository_model.go
- reflect.Type
- context.Context
- deferCloseResponseBody
- PortfolioAnalysisConfigurationAppService
- asset_integration_test.go
- logFanOutConsumer
- golang_ext_tree_test.go
- go_pkg_testing
- BuildCleanupFunctionBuilder
- NewOrderedMapIterator
- testing.T
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- AssetDomService
- Adapter
- golang_ext_util_concurrent.go
- assert_json_extension.go
- BuildAppError
- AllocationPlan
- AllocationStructure
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- App
- PropagateAsAppErrorWithNewMessage
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- portfolio_allocation_integration_test.go
- go_pkg_database_sql

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 79 edges
2. `App` - 35 edges
3. `BuildCleanupFunctionBuilder()` - 34 edges
4. `PortfolioAllocation` - 32 edges
5. `PropagateAsAppErrorWithNewMessage()` - 31 edges
6. `Asset` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 26 edges
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

## Communities (41 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.05
Nodes (71): validateAllocationPlanPlannedDate(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+63 more)

### Community 1 - ".buildAppComponents"
Cohesion: 0.18
Nodes (11): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), BuildAllocationDomService() (+3 more)

### Community 2 - "Portfolio"
Cohesion: 0.31
Nodes (5): Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioRDBMSRepository

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.07
Nodes (39): BuildAllocationPlanRESTController(), cleanNilAllocations(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService() (+31 more)

### Community 4 - "portfolio_allocation_repository_model.go"
Cohesion: 0.90
Nodes (4): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), portfolioAllocationJoinedRowDTS

### Community 5 - "reflect.Type"
Cohesion: 0.05
Nodes (46): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_github_com_go_playground_validator_v10, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type (+38 more)

### Community 6 - "context.Context"
Cohesion: 0.10
Nodes (33): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+25 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.13
Nodes (26): go_pkg_io, TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches() (+18 more)

### Community 9 - "PortfolioAnalysisConfigurationAppService"
Cohesion: 0.40
Nodes (5): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, AnalysisOptions, DivergenceAnalysisRESTController

### Community 10 - "asset_integration_test.go"
Cohesion: 0.14
Nodes (20): go_pkg_net_url, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+12 more)

### Community 11 - "logFanOutConsumer"
Cohesion: 0.10
Nodes (19): go_pkg_github_com_nhatthm_httpmock, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry() (+11 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.12
Nodes (18): TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require (+10 more)

### Community 14 - "BuildCleanupFunctionBuilder"
Cohesion: 0.23
Nodes (28): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), TestPostAsset(), assertPersistedAssetWithExternalData() (+20 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "testing.T"
Cohesion: 0.13
Nodes (30): net/http.Response, testing.T, TestPostAssetWithoutExternalData(), postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio() (+22 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.13
Nodes (14): MapToPortfolioObservationTimestamp(), BuildPortfolioAllocationRESTController(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, preparePortfolioAllocationTempInserts() (+6 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.27
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 21 - "AssetDomService"
Cohesion: 0.15
Nodes (8): BuildAssetRESTController(), AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), findPersistedExternalAssetQuoteMatch(), AssetDomService, AssetIntegrationServicesPerSource

### Community 23 - "Adapter"
Cohesion: 0.05
Nodes (33): assetRowScanner(), database/sql.DB, database/sql.Result, database/sql.Row, database/sql.Rows, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query (+25 more)

### Community 24 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "BuildAppError"
Cohesion: 0.10
Nodes (26): YahooFinanceAssetIntegrationService, ExternalAssetQuote, BuildYahooFinanceAssetIntegrationService(), extractLastClose(), mapToExternalAsset(), mapToExternalAssetQuote(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL() (+18 more)

### Community 28 - "AllocationPlan"
Cohesion: 0.09
Nodes (34): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+26 more)

### Community 29 - "AllocationStructure"
Cohesion: 0.09
Nodes (18): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, go_pkg_database_sql_driver (+10 more)

### Community 43 - "App"
Cohesion: 0.07
Nodes (27): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os (+19 more)

### Community 47 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.18
Nodes (12): Asset, buildAssetInsertValue(), buildAssetInsertValues(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T (+4 more)

### Community 49 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.06
Nodes (55): malformedQueryBindingTarget, go_pkg_context, go_pkg_errors, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain (+47 more)

### Community 50 - "portfolio_allocation_integration_test.go"
Cohesion: 0.13
Nodes (15): go_pkg_encoding_json, go_pkg_strconv, postPortfolioAllocationForValidationFailure(), TestGetAvailableHistoryObservations(), TestGetPortfolioAllocationHistory(), TestGetPortfolioAllocationHistoryForObservationTimestamp(), TestGetPortfolioAllocationHistoryForObservationTimestampNoneFound(), TestGetPortfolioAllocationHistoryWithMultiplePortfoliosAndManyObservations() (+7 more)

### Community 52 - "go_pkg_database_sql"
Cohesion: 0.23
Nodes (10): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_database_sql, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, database/sql.NullString, StringPointerToNullString() (+2 more)

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `MapIterator`, `contextKey`, `ErrorResponse` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `.buildAppComponents`, `reflect.Type`, `logFanOutConsumer`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Adapter`?**
  _High betweenness centrality (0.077) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `App`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `github.com/gin-gonic/gin.Context`, `allocation_plan_domain_service.go`, `portfolio_allocation_repository_model.go`, `PortfolioAllocation`, `ExecuteDBQuery`, `AssetDomService`, `Adapter`, `AllocationPlan`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **Are the 77 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 77 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `MapIterator` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.05201465201465202 - nodes in this community are weakly interconnected._
- **Should `allocation_plan_domain_service.go` be split into smaller, more focused modules?**
  _Cohesion score 0.0672316384180791 - nodes in this community are weakly interconnected._