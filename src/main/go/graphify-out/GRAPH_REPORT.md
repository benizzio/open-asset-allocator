# Graph Report - go  (2026-11-01)

## Corpus Check
- 124 files · ~67,040 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1134 nodes · 3473 edges · 48 communities (37 shown, 11 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 264 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `e172b503`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- BuildCleanupFunctionBuilder
- Adapter
- allocation_plan_domain_service.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- reflect.Type
- context.Context
- deferCloseResponseBody
- portfolio_integration_test.go
- go_pkg_strings
- testing.T
- go_pkg_github_com_golang_glog
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- Asset
- NewOrderedMapIterator
- PropagateAsAppErrorWithNewMessage
- PortfolioObservationTimestamp
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- ExternalAssetQuote
- golang_ext_util_concurrent.go
- portfolio_observation_ownership_migration_integration_test.go
- AllocationPlan
- assert_json_extension.go
- infra.go
- GinServer
- PortfolioNotFoundError
- PortfolioAllocation
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- .buildAppComponents
- assetRowScanner
- asset_integration_client_yahoo_model.go
- App
- AssertDBWithQueryMultipleRows
- go_pkg_net_http
- go_pkg_encoding_json
- ForeignKeyConstraintViolationError

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `BuildCleanupFunctionBuilder()` - 41 edges
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
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `AssetRESTController` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `AllocationPlanAssetDTS` --references--> `ParseableInt64`  [EXTRACTED]
  api/rest/model/allocation_plan_rest_model.go → langext/golang_ext_parseable_int64.go

## Import Cycles
- None detected.

## Communities (48 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (47): validateAllocationPlanPlannedDate(), mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs() (+39 more)

### Community 1 - "BuildCleanupFunctionBuilder"
Cohesion: 0.19
Nodes (16): github.com/go-ozzo/ozzo-dbx.Params, TestPostAsset(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), TestPostAssetWithoutExternalData(), postAsset(), FetchWithDBQuery(), assertPortfolioHistoryTagValidationCount() (+8 more)

### Community 2 - "Adapter"
Cohesion: 0.12
Nodes (9): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey (+1 more)

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.06
Nodes (41): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+33 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.39
Nodes (6): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.07
Nodes (41): BuildDivergenceAnalysisRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext(), buildDivergenceAnalysisContext() (+33 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.15
Nodes (27): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+19 more)

### Community 8 - "portfolio_integration_test.go"
Cohesion: 0.14
Nodes (22): go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, database/sql.NullString, StringPointerToNullString(), StringToNullString(), postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses() (+14 more)

### Community 9 - "go_pkg_strings"
Cohesion: 0.11
Nodes (17): go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_lib_pq, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency (+9 more)

### Community 10 - "testing.T"
Cohesion: 0.10
Nodes (31): go_pkg_io, go_pkg_net_url, testing.T, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker() (+23 more)

### Community 11 - "go_pkg_github_com_golang_glog"
Cohesion: 0.33
Nodes (3): go_pkg_flag, go_pkg_github_com_golang_glog, go_pkg_net

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.16
Nodes (13): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require (+5 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.16
Nodes (23): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+15 more)

### Community 15 - "Asset"
Cohesion: 0.25
Nodes (8): Asset, buildAssetInsertValue(), buildAssetInsertValues(), classifyObservationWriteError(), BuildUniqueConstraintViolationError(), BuildQueryInTransaction(), T, IsUniqueConstraintViolation()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.16
Nodes (13): BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioRepository(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), RepositoryRDBMSAdapter, BuildILikeSubstringPattern(), TestBuildILikeSubstringPattern() (+5 more)

### Community 18 - "PortfolioObservationTimestamp"
Cohesion: 0.25
Nodes (6): preparePortfolioAllocationTempInserts(), PortfolioObservationTimestamp, BuildAppError(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository, PortfolioAllocationRDBMSRepository

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.25
Nodes (21): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+13 more)

### Community 21 - "ExternalAssetQuote"
Cohesion: 0.29
Nodes (5): ExternalAssetQuote, mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppErrorFormatted()

### Community 22 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 23 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.09
Nodes (25): database/sql.Tx, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, T, withParams(), processParamsForPostgreSQL(), processSQL(), T (+17 more)

### Community 24 - "AllocationPlan"
Cohesion: 0.07
Nodes (35): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+27 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.06
Nodes (37): go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_path_filepath, go_pkg_sync (+29 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 29 - "PortfolioAllocation"
Cohesion: 0.06
Nodes (59): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+51 more)

### Community 40 - ".buildAppComponents"
Cohesion: 0.10
Nodes (19): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, AssetIntegrationService (+11 more)

### Community 41 - "assetRowScanner"
Cohesion: 0.67
Nodes (3): assetRowScanner(), database/sql.Rows, ReturningIntIdRowScanner()

### Community 42 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.16
Nodes (16): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceChartResponseDTS (+8 more)

### Community 44 - "App"
Cohesion: 0.10
Nodes (23): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os (+15 more)

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.22
Nodes (27): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), assertPersistedAssetWithExternalData(), TestPostAllocationPlanRejectsWhitespaceName() (+19 more)

### Community 49 - "go_pkg_net_http"
Cohesion: 0.24
Nodes (11): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin (+3 more)

### Community 50 - "go_pkg_encoding_json"
Cohesion: 0.21
Nodes (10): go_pkg_encoding_json, go_pkg_strconv, RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T (+2 more)

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 96 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `.buildAppComponents`, `go_pkg_strings`, `go_pkg_github_com_golang_glog`, `go_pkg_net_http`, `infra.go`, `GinServer`?**
  _High betweenness centrality (0.051) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `allocation_plan_domain_service.go`, `.buildAppComponents`, `assetRowScanner`, `PropagateAsAppErrorWithNewMessage`, `ExecuteDBQuery`, `ExternalAssetQuote`, `AllocationPlan`, `PortfolioAllocation`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `AllocationPlan` to `go_pkg_strings`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.05651176133103844 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.12 - nodes in this community are weakly interconnected._