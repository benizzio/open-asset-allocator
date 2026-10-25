# Graph Report - go  (2026-11-01)

## Corpus Check
- 132 files · ~74,714 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1248 nodes · 3915 edges · 58 communities (47 shown, 11 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 292 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `177bf273`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- portfolio_allocation_pagination_integration_test.go
- T
- IsZeroValue
- .buildAppComponents
- reflect.Type
- Portfolio
- deferCloseResponseBody
- BuildAppError
- go_pkg_fmt
- PortfolioAllocation
- go_pkg_net_http
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- PropagateAsAppErrorWithNewMessage
- NewOrderedMapIterator
- context.Context
- asset_rest_model.go
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- AssetDomService
- AllocationPlan
- golang_ext_util_concurrent.go
- portfolio_observation_ownership_migration_integration_test.go
- assert_json_extension.go
- infra.go
- Adapter
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- portfolio_rest_mapping.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- SQLTransactionalContext
- go_pkg_strings
- PortfolioRESTController
- testing.T
- App
- asset_integration_test.go
- assert_db.go
- CustomValidationErrorsBuilder
- go_pkg_github_com_golang_glog
- asset_integration_client_yahoo_model.go
- http_client.go
- HierarchicalId
- ForeignKeyConstraintViolationError
- go_pkg_github_com_benizzio_open_asset_allocator_infra_util_urlutil
- ObservationPageCursor
- pagination_cursor.go
- PortfolioAllocationRESTController
- joinAny

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `PropagateAsAppErrorWithNewMessage()` - 44 edges
3. `BuildCleanupFunctionBuilder()` - 44 edges
4. `App` - 35 edges
5. `PortfolioAllocation` - 34 edges
6. `Asset` - 30 edges
7. `ExecuteDBQuery()` - 29 edges
8. `IsZeroValue()` - 29 edges
9. `PortfolioObservationTimestamp` - 27 edges
10. `PortfolioAllocationDomService` - 27 edges

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

## Communities (58 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.13
Nodes (24): validateAllocationPlanPlannedDate(), BuildAssetRESTController(), readAllocationPlanPaginationCursor(), readObservationPaginationCursor(), respondWithPaginationCursorValidationError(), bindAndValidatePortfolioSnapshot(), readHistoryObservationTimestampId(), go_pkg_github_com_go_playground_validator_v10 (+16 more)

### Community 1 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.12
Nodes (44): time.Time, NullTime, allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture(), encodeTestCursor(), getAllocationPlanPage() (+36 more)

### Community 2 - "T"
Cohesion: 0.10
Nodes (17): database/sql.Result, github.com/go-ozzo/ozzo-dbx.Query, QueryBuilder, T, withParams(), processParamsForPostgreSQL(), processSQL(), T (+9 more)

### Community 3 - "IsZeroValue"
Cohesion: 0.05
Nodes (49): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), validateAssetIdentification(), AllocationHierarchy, AllocationHierarchyLevel (+41 more)

### Community 4 - ".buildAppComponents"
Cohesion: 0.22
Nodes (11): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository() (+3 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "Portfolio"
Cohesion: 0.07
Nodes (36): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+28 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.13
Nodes (45): database/sql.NullString, StringPointerToNullString(), StringToNullString(), assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestGetAllocationPlans() (+37 more)

### Community 8 - "BuildAppError"
Cohesion: 0.20
Nodes (12): decodePaginationCursorPayload(), ExternalAsset, ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppError() (+4 more)

### Community 9 - "go_pkg_fmt"
Cohesion: 0.15
Nodes (7): PortfolioNotFoundError, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, go_pkg_time

### Community 10 - "PortfolioAllocation"
Cohesion: 0.14
Nodes (16): BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, buildSelectedExternalAsset() (+8 more)

### Community 11 - "go_pkg_net_http"
Cohesion: 0.33
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_net_http

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.10
Nodes (25): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra (+17 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (23): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+15 more)

### Community 15 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.19
Nodes (11): Asset, buildAssetInsertValue(), buildAssetInsertValues(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, ToPointerSlice() (+3 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "context.Context"
Cohesion: 0.20
Nodes (8): classifyObservationWriteError(), preparePortfolioAllocationTempInserts(), PortfolioObservationTimestamp, context.Context, BuildUniqueConstraintViolationError(), IsUniqueConstraintViolation(), ToSQLTransactionalContext(), PortfolioAllocationRDBMSRepository

### Community 18 - "asset_rest_model.go"
Cohesion: 0.19
Nodes (19): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+11 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.13
Nodes (37): github.com/go-ozzo/ozzo-dbx.Params, TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPostAsset(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), TestPostAssetWithoutExternalData() (+29 more)

### Community 21 - "AssetDomService"
Cohesion: 0.16
Nodes (8): AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), AssetDomService, TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), AssetIntegrationServicesPerSource

### Community 22 - "AllocationPlan"
Cohesion: 0.10
Nodes (33): BuildAllocationPlanRESTController(), AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation() (+25 more)

### Community 23 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 24 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.05
Nodes (39): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os (+31 more)

### Community 27 - "Adapter"
Cohesion: 0.06
Nodes (27): BuildYahooFinanceAssetIntegrationClient(), assetRowScanner(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), database/sql.DB, database/sql.Rows, github.com/gin-gonic/gin.Engine, github.com/go-ozzo/ozzo-dbx.Builder (+19 more)

### Community 28 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.31
Nodes (7): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 29 - "portfolio_rest_mapping.go"
Cohesion: 0.14
Nodes (27): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+19 more)

### Community 40 - "SQLTransactionalContext"
Cohesion: 0.23
Nodes (5): SQLTransactionalContext, withTransaction(), contextKey, TransactionalContext, AllocationPlanRDBMSRepository

### Community 41 - "go_pkg_strings"
Cohesion: 0.20
Nodes (10): go_pkg_database_sql, go_pkg_github_com_lib_pq, go_pkg_strings, database/sql.Row, database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), prepareBulkInsertValues() (+2 more)

### Community 42 - "PortfolioRESTController"
Cohesion: 0.21
Nodes (9): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService, PortfolioDTS (+1 more)

### Community 43 - "testing.T"
Cohesion: 0.10
Nodes (39): testing.T, capturedPaginationQuery, TestGetDivergenceAnalysisOptions(), TestGetDivergenceAnalysisV2(), TestGetDivergenceAnalysisV2FullDivergence(), TestGetDivergenceAnalysisV2WhenHistoryHighestLevelHasExtraAllocation(), TestGetDivergenceAnalysisV2WhenPlanHighestLevelHasExtraAllocation(), TestGetDivergenceAnalysisV2WhenPlanLowestLevelHasDifferentRanges() (+31 more)

### Community 44 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 45 - "asset_integration_test.go"
Cohesion: 0.19
Nodes (14): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+6 more)

### Community 46 - "assert_db.go"
Cohesion: 0.50
Nodes (8): github.com/go-ozzo/ozzo-dbx.NullStringMap, assertAllRows(), AssertDBWithQuery(), assertRowColumn(), assertSingleRow(), executeQueryForMultipleRows(), validateRowCount(), AssertableNullStringMap

### Community 47 - "CustomValidationErrorsBuilder"
Cohesion: 0.36
Nodes (4): buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 48 - "go_pkg_github_com_golang_glog"
Cohesion: 0.33
Nodes (3): go_pkg_flag, go_pkg_github_com_golang_glog, go_pkg_net

### Community 49 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.16
Nodes (16): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceChartResponseDTS (+8 more)

### Community 51 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 53 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 54 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 57 - "ObservationPageCursor"
Cohesion: 0.30
Nodes (7): addObservationCursorFilter(), buildPortfolioAllocationsByObservationPage(), buildPortfolioObservationPage(), T, ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage

### Community 59 - "pagination_cursor.go"
Cohesion: 0.15
Nodes (18): encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor(), readPaginationCursor(), validatePaginationCursorScope(), go_pkg_bytes, go_pkg_encoding_base64, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_url (+10 more)

### Community 60 - "PortfolioAllocationRESTController"
Cohesion: 0.15
Nodes (5): PaginatedResponse, T, github.com/gin-gonic/gin.HandlersChain, RESTRoute, PortfolioAllocationRESTController

### Community 62 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 105 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Adapter` connect `Adapter` to `T`, `SQLTransactionalContext`, `go_pkg_strings`, `testing.T`, `App`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Why does `App` connect `App` to `.buildAppComponents`, `reflect.Type`, `go_pkg_fmt`, `go_pkg_net_http`, `go_pkg_testing`, `go_pkg_github_com_golang_glog`, `infra.go`, `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`?**
  _High betweenness centrality (0.047) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `IsZeroValue`, `PortfolioAllocation`, `asset_rest_model.go`, `BuildCleanupFunctionBuilder`, `AssetDomService`, `AllocationPlan`, `Adapter`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.12804878048780488 - nodes in this community are weakly interconnected._
- **Should `portfolio_allocation_pagination_integration_test.go` be split into smaller, more focused modules?**
  _Cohesion score 0.12025901942645699 - nodes in this community are weakly interconnected._