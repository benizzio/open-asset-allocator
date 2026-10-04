# Graph Report - go  (2026-10-04)

## Corpus Check
- 119 files · ~57,555 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1063 nodes · 3134 edges · 42 communities (33 shown, 9 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 233 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `0769c512`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- joinAny
- PortfolioAllocation
- asset_rest_model.go
- reflect.Type
- infra.go
- context.Context
- allocation_plan_domain_service.go
- parseAssetTextSearch
- asset_integration_test.go
- deferCloseResponseBody
- AllocationStructure
- golang_ext_tree_test.go
- golang_ext_util_concurrent.go
- assert_json_extension.go
- Adapter
- assert_db.go
- NewOrderedMapIterator
- Portfolio
- AllocationPlan
- golang_ext_util_struct_unify_pointers_test.go
- GinServer
- BuildCleanupFunctionBuilder
- BuildAppError
- portfolio_integration_test.go
- .buildAppComponents
- testing.T
- App
- base.go
- CustomSliceTable[T]
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Go Lint Configuration
- External Integration Tests
- Asset
- database/sql.NullString
- AllocationPlanManagementAppService
- AllocationDomService
- PropagateAsAppErrorWithNewMessage

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 76 edges
2. `App` - 35 edges
3. `PortfolioAllocation` - 32 edges
4. `PropagateAsAppErrorWithNewMessage()` - 32 edges
5. `Asset` - 31 edges
6. `BuildCleanupFunctionBuilder()` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 24 edges
9. `PlannedAllocation` - 23 edges
10. `AllocationPlan` - 23 edges

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

## Communities (42 total, 9 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.09
Nodes (27): BuildDivergenceAnalysisRESTController(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.HandlersChain, HandleAPIError(), handleDomainError() (+19 more)

### Community 1 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

### Community 2 - "PortfolioAllocation"
Cohesion: 0.07
Nodes (47): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+39 more)

### Community 3 - "asset_rest_model.go"
Cohesion: 0.08
Nodes (41): YahooFinanceAssetIntegrationService, MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+33 more)

### Community 4 - "reflect.Type"
Cohesion: 0.05
Nodes (47): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_github_com_go_playground_validator_v10, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type (+39 more)

### Community 5 - "infra.go"
Cohesion: 0.05
Nodes (39): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os (+31 more)

### Community 6 - "context.Context"
Cohesion: 0.08
Nodes (39): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+31 more)

### Community 7 - "allocation_plan_domain_service.go"
Cohesion: 0.09
Nodes (32): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+24 more)

### Community 8 - "parseAssetTextSearch"
Cohesion: 0.83
Nodes (3): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch()

### Community 9 - "asset_integration_test.go"
Cohesion: 0.12
Nodes (24): malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_io (+16 more)

### Community 10 - "deferCloseResponseBody"
Cohesion: 0.13
Nodes (24): TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches(), TestPostAllocationPlanValidation_MissingDetails(), TestPostAllocationPlanValidation_MissingHierarchicalId() (+16 more)

### Community 11 - "AllocationStructure"
Cohesion: 0.11
Nodes (15): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, database/sql/driver.Value (+7 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 14 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 15 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 16 - "Adapter"
Cohesion: 0.06
Nodes (26): database/sql.DB, database/sql.Result, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, buildPingContext(), Adapter, T (+18 more)

### Community 17 - "assert_db.go"
Cohesion: 0.26
Nodes (24): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), assertPersistedAssetWithExternalData(), TestPostPortfolioAllocationHistoryFullMerge(), TestPostPortfolioAllocationHistoryInsertEmptyZeroTimestamp() (+16 more)

### Community 18 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 19 - "Portfolio"
Cohesion: 0.27
Nodes (5): BuildPortfolioDivergenceAnalysisAppService(), Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService

### Community 20 - "AllocationPlan"
Cohesion: 0.12
Nodes (27): MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS(), mapToPlannedAllocationDTSs() (+19 more)

### Community 21 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 22 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 23 - "BuildCleanupFunctionBuilder"
Cohesion: 0.18
Nodes (27): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPostAssetWithoutExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote() (+19 more)

### Community 24 - "BuildAppError"
Cohesion: 0.26
Nodes (6): BuildAllocationPlanRepository(), BuildAppError(), BuildQueryInTransaction(), T, ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 25 - "portfolio_integration_test.go"
Cohesion: 0.11
Nodes (28): net/http.Response, postAssetForValidationFailure(), TestPostAsset(), TestPostAssetFailureWithMissingRequiredFields(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), postAsset(), FetchWithDBQuery() (+20 more)

### Community 26 - ".buildAppComponents"
Cohesion: 0.26
Nodes (9): BuildAssetRESTController(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), github.com/go-ozzo/ozzo-dbx.Rows, RepositoryRDBMSAdapter, AllocationRDBMSRepository (+1 more)

### Community 27 - "testing.T"
Cohesion: 0.09
Nodes (31): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), testing.T, TestBindAndValidateQueryWithInvalidResponse_MalformedQuery(), TestBuildILikeSubstringPattern() (+23 more)

### Community 28 - "App"
Cohesion: 0.05
Nodes (56): BuildYahooFinanceAssetIntegrationClient(), assetRowScanner(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_context, go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_errors (+48 more)

### Community 38 - "Asset"
Cohesion: 0.17
Nodes (9): Asset, AssetIntegrationService, AssetRepository, buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetDomService(), collectIntegrationServices(), AssetDomService (+1 more)

### Community 44 - "database/sql.NullString"
Cohesion: 0.38
Nodes (6): TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), database/sql.NullString, StringPointerToNullString(), StringToNullString(), ValueToString()

### Community 46 - "AllocationPlanManagementAppService"
Cohesion: 0.38
Nodes (6): BuildAllocationPlanRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), TransactionManager, AllocationPlanRESTController

### Community 47 - "AllocationDomService"
Cohesion: 0.38
Nodes (4): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 51 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.22
Nodes (8): BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), IsUniqueConstraintViolation(), BuildILikeSubstringPattern(), ToPointerSlice(), AssetRDBMSRepository, S

## Knowledge Gaps
- **13 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 89 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **9 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `reflect.Type`, `infra.go`, `Adapter`, `GinServer`, `.buildAppComponents`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `PortfolioAllocation`, `asset_rest_model.go`, `PropagateAsAppErrorWithNewMessage`, `AllocationPlan`, `BuildCleanupFunctionBuilder`, `BuildAppError`, `App`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `App`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **Are the 74 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestGetAllocationPlans()` and `TestPostAllocationPlanForInsertion()`) actually correct?**
  _`deferCloseResponseBody()` has 74 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.09225589225589226 - nodes in this community are weakly interconnected._
- **Should `PortfolioAllocation` be split into smaller, more focused modules?**
  _Cohesion score 0.06881287726358148 - nodes in this community are weakly interconnected._