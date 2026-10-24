# Graph Report - go  (2026-11-01)

## Corpus Check
- 131 files · ~73,945 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1231 nodes · 3848 edges · 58 communities (44 shown, 14 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 288 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `2cd3c8e9`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- portfolio_allocation_pagination_integration_test.go
- Adapter
- IsZeroValue
- .buildAppComponents
- reflect.Type
- context.Context
- deferCloseResponseBody
- testing.T
- HierarchicalId
- asset_integration_test.go
- go_pkg_net_http
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- ForeignKeyConstraintViolationError
- NewOrderedMapIterator
- PropagateAsAppErrorWithNewMessage
- AllocationPlanRDBMSRepository
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- BuildAppError
- allocation_plan_rest_model.go
- golang_ext_util_concurrent.go
- AllocationPlan
- assert_json_extension.go
- infra.go
- GinServer
- logging.go
- portfolio_rest_mapping.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- Asset
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- asset_integration_service_yahoo.go
- FetchWithDBQuery
- App
- PlanType
- BuildCleanupFunctionBuilder
- CustomValidationErrorsBuilder
- Portfolio
- PortfolioAllocationRESTController
- http_client.go
- T
- RESTRoute
- joinAny
- go_pkg_fmt
- go_pkg_github_com_benizzio_open_asset_allocator_infra_util_urlutil
- go_pkg_time
- pagination_cursor.go

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `BuildCleanupFunctionBuilder()` - 44 edges
3. `PropagateAsAppErrorWithNewMessage()` - 37 edges
4. `App` - 35 edges
5. `PortfolioAllocation` - 34 edges
6. `Asset` - 30 edges
7. `ExecuteDBQuery()` - 29 edges
8. `IsZeroValue()` - 29 edges
9. `PortfolioObservationTimestamp` - 27 edges
10. `PortfolioAllocationDomService` - 27 edges

## Surprising Connections (you probably didn't know these)
- `readPaginationCursor()` --calls--> `RawQueryHasParameter()`  [EXTRACTED]
  api/rest/pagination_cursor.go → infra/util/url/query.go
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `AllocationPlanRESTController` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go

## Import Cycles
- None detected.

## Communities (58 total, 14 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.15
Nodes (22): validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), respondWithPaginationCursorValidationError(), bindAndValidatePortfolioSnapshot(), github.com/gin-gonic/gin.Context, HandleAPIError() (+14 more)

### Community 1 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.10
Nodes (46): go_pkg_net_url, net/url.Values, time.Time, NullTime, HasQueryParameter(), RawQueryHasParameter(), allocationPlanEndpoint(), allocationPlanPageNames() (+38 more)

### Community 2 - "Adapter"
Cohesion: 0.05
Nodes (35): database/sql.DB, database/sql.Result, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, buildPingContext(), Adapter (+27 more)

### Community 3 - "IsZeroValue"
Cohesion: 0.06
Nodes (46): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+38 more)

### Community 4 - ".buildAppComponents"
Cohesion: 0.18
Nodes (11): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), BuildAllocationDomService() (+3 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.05
Nodes (57): BuildDivergenceAnalysisRESTController(), BuildPortfolioAllocationRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext() (+49 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.11
Nodes (34): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+26 more)

### Community 8 - "testing.T"
Cohesion: 0.11
Nodes (33): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestParseAssetTextSearch(), testing.T, TestBuildILikeSubstringPattern(), TestPostAssetWithoutExternalData(), postPortfolioForValidationFailure(), putForValidationFailure() (+25 more)

### Community 9 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 10 - "asset_integration_test.go"
Cohesion: 0.13
Nodes (17): go_pkg_encoding_json, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+9 more)

### Community 11 - "go_pkg_net_http"
Cohesion: 0.23
Nodes (11): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_github_com_golang_glog (+3 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.20
Nodes (14): TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_io (+6 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (22): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+14 more)

### Community 15 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.20
Nodes (9): PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, BuildILikeSubstringPattern(), ToPointerSlice(), AssetRDBMSRepository, PortfolioRDBMSRepository (+1 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.26
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 21 - "BuildAppError"
Cohesion: 0.16
Nodes (13): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), classifyObservationWriteError(), github.com/shopspring/decimal.Decimal, golang.org/x/text/currency.Unit, BuildAppError(), BuildAppErrorFormatted() (+5 more)

### Community 22 - "allocation_plan_rest_model.go"
Cohesion: 0.29
Nodes (15): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+7 more)

### Community 23 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 24 - "AllocationPlan"
Cohesion: 0.27
Nodes (9): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), AllocationPlan, PlannedAllocation, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow(), mapPlannedAllocationRow(), mapPlannedAllocationRows() (+1 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.06
Nodes (36): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os, go_pkg_path_filepath (+28 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "logging.go"
Cohesion: 0.33
Nodes (4): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, ConfigLogger(), main()

### Community 29 - "portfolio_rest_mapping.go"
Cohesion: 0.12
Nodes (30): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+22 more)

### Community 40 - "Asset"
Cohesion: 0.06
Nodes (43): BuildAssetRESTController(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+35 more)

### Community 41 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.21
Nodes (13): assetRowScanner(), go_pkg_context, go_pkg_database_sql, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext (+5 more)

### Community 42 - "asset_integration_service_yahoo.go"
Cohesion: 0.19
Nodes (11): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), BuildYahooFinanceAssetIntegrationClient(), YahooFinanceAssetIntegrationClient, TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient (+3 more)

### Community 43 - "FetchWithDBQuery"
Cohesion: 0.20
Nodes (12): github.com/go-ozzo/ozzo-dbx.Params, TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), postAsset(), explainPaginationQuery(), TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans(), FetchWithDBQuery(), assertPortfolioHistoryTagValidationCount() (+4 more)

### Community 44 - "App"
Cohesion: 0.13
Nodes (16): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, Configuration, RDBMSConfiguration (+8 more)

### Community 45 - "PlanType"
Cohesion: 0.21
Nodes (7): allocationPlanCursorFromPayload(), GetPlanType(), PlanType, AllocationPlanOwnership, AllocationPlanPage, AllocationPlanPageCursor, buildAllocationPlanPage()

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.19
Nodes (33): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+25 more)

### Community 47 - "CustomValidationErrorsBuilder"
Cohesion: 0.31
Nodes (5): go_pkg_github_com_go_playground_validator_v10, buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 48 - "Portfolio"
Cohesion: 0.24
Nodes (8): BuildAllocationPlanRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, TransactionManager

### Community 49 - "PortfolioAllocationRESTController"
Cohesion: 0.31
Nodes (3): readHistoryObservationTimestampId(), ParseInt64(), PortfolioAllocationRESTController

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 53 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

### Community 54 - "go_pkg_fmt"
Cohesion: 0.15
Nodes (10): PortfolioNotFoundError, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_strings, database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement() (+2 more)

### Community 57 - "go_pkg_time"
Cohesion: 0.36
Nodes (3): go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_shopspring_decimal, go_pkg_time

### Community 59 - "pagination_cursor.go"
Cohesion: 0.15
Nodes (18): PaginatedResponse, T, decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor() (+10 more)

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 103 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **14 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `.buildAppComponents`, `reflect.Type`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `asset_integration_service_yahoo.go`, `go_pkg_net_http`, `go_pkg_time`, `infra.go`, `GinServer`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `App`, `go_pkg_fmt`?**
  _High betweenness centrality (0.037) - this node is a cross-community bridge._
- **Why does `GinServer` connect `GinServer` to `go_pkg_net_http`, `App`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.14714714714714713 - nodes in this community are weakly interconnected._
- **Should `portfolio_allocation_pagination_integration_test.go` be split into smaller, more focused modules?**
  _Cohesion score 0.10040816326530612 - nodes in this community are weakly interconnected._