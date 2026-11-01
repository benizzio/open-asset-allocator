# Graph Report - go  (2026-11-01)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 1247 nodes · 3915 edges · 57 communities (47 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 292 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f670338e`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- IsZeroValue
- reflect.Type
- Portfolio
- Adapter
- infra.go
- AllocationPlan
- portfolio_allocation_pagination_integration_test.go
- deferCloseResponseBody
- testing.T
- go_pkg_testing
- github.com/gin-gonic/gin.Context
- BuildCleanupFunctionBuilder
- portfolio_rest_mapping.go
- T
- golang_ext_tree_test.go
- PortfolioAllocation
- golang_ext_util_concurrent.go
- PropagateAsAppErrorWithNewMessage
- NewOrderedMapIterator
- portfolio_allocation_integration_test.go
- asset_rest_model.go
- pagination_cursor.go
- context.Context
- asset_integration_client_yahoo_model.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- go_pkg_fmt
- BuildAppError
- PortfolioAllocationRESTController
- golang_ext_util_struct_unify_pointers_test.go
- go_pkg_net_http
- HierarchicalId
- AssetDomService
- go_pkg_strings
- SQLTransactionalContext
- App
- asset_integration_test.go
- .buildAppComponents
- PortfolioRESTController
- ObservationPageCursor
- assert_json_extension.go
- portfolio_observation_ownership_migration_integration_test.go
- assert_db.go
- http_client.go
- CustomValidationErrorsBuilder
- ForeignKeyConstraintViolationError
- go_pkg_github_com_golang_glog
- joinAny
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `BuildCleanupFunctionBuilder()` - 44 edges
3. `PropagateAsAppErrorWithNewMessage()` - 44 edges
4. `App` - 35 edges
5. `PortfolioAllocation` - 34 edges
6. `Asset` - 30 edges
7. `IsZeroValue()` - 29 edges
8. `ExecuteDBQuery()` - 29 edges
9. `NewMapTree()` - 27 edges
10. `PortfolioAllocationDomService` - 27 edges

## Surprising Connections (you probably didn't know these)
- `IsNilPointer()` --calls--> `ValueJsonColumn()`  [EXTRACTED]
  langext/golang_ext_util.go → infra/rdbms/sqlext/sqlext_utils.go
- `validateAssetIdentification()` --references--> `CustomValidationErrorsBuilder`  [EXTRACTED]
  api/rest/portfolio_allocation_controller.go → infra/validation/custom_validation_error_builder.go
- `AllocationHierarchyLevel` --references--> `buildHierarchySubIterationContext()`  [EXTRACTED]
  domain/allocation.go → application/portfolio_divergence_analysis_context.go
- `AllocationHierarchyLevel` --references--> `getHierarchySubIterationContextValue()`  [EXTRACTED]
  domain/allocation.go → application/portfolio_divergence_analysis_context.go
- `appendLevelDescription()` --references--> `CustomSlice`  [EXTRACTED]
  domain/service/allocation_plan_domain_service.go → langext/golang_ext_custom_slice.go

## Import Cycles
- None detected.

## Communities (57 total, 10 thin omitted)

### Community 0 - "IsZeroValue"
Cohesion: 0.05
Nodes (49): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), validateAssetIdentification(), AllocationHierarchy, AllocationHierarchyLevel (+41 more)

### Community 1 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 2 - "Portfolio"
Cohesion: 0.07
Nodes (36): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+28 more)

### Community 3 - "Adapter"
Cohesion: 0.06
Nodes (27): BuildYahooFinanceAssetIntegrationClient(), assetRowScanner(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), database/sql.DB, database/sql.Rows, github.com/gin-gonic/gin.Engine, github.com/go-ozzo/ozzo-dbx.Builder (+19 more)

### Community 4 - "infra.go"
Cohesion: 0.05
Nodes (39): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os (+31 more)

### Community 5 - "AllocationPlan"
Cohesion: 0.10
Nodes (33): BuildAllocationPlanRESTController(), AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation() (+25 more)

### Community 6 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.12
Nodes (44): time.Time, NullTime, allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture(), encodeTestCursor(), getAllocationPlanPage() (+36 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.13
Nodes (45): database/sql.NullString, StringPointerToNullString(), StringToNullString(), assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestGetAllocationPlans() (+37 more)

### Community 8 - "testing.T"
Cohesion: 0.10
Nodes (39): testing.T, capturedPaginationQuery, TestGetDivergenceAnalysisOptions(), TestGetDivergenceAnalysisV2(), TestGetDivergenceAnalysisV2FullDivergence(), TestGetDivergenceAnalysisV2WhenHistoryHighestLevelHasExtraAllocation(), TestGetDivergenceAnalysisV2WhenPlanHighestLevelHasExtraAllocation(), TestGetDivergenceAnalysisV2WhenPlanLowestLevelHasDifferentRanges() (+31 more)

### Community 9 - "go_pkg_testing"
Cohesion: 0.10
Nodes (25): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra (+17 more)

### Community 10 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.13
Nodes (24): validateAllocationPlanPlannedDate(), BuildAssetRESTController(), readAllocationPlanPaginationCursor(), readObservationPaginationCursor(), respondWithPaginationCursorValidationError(), bindAndValidatePortfolioSnapshot(), readHistoryObservationTimestampId(), go_pkg_github_com_go_playground_validator_v10 (+16 more)

### Community 11 - "BuildCleanupFunctionBuilder"
Cohesion: 0.13
Nodes (37): github.com/go-ozzo/ozzo-dbx.Params, TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPostAsset(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), TestPostAssetWithoutExternalData() (+29 more)

### Community 12 - "portfolio_rest_mapping.go"
Cohesion: 0.14
Nodes (27): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+19 more)

### Community 13 - "T"
Cohesion: 0.10
Nodes (17): database/sql.Result, github.com/go-ozzo/ozzo-dbx.Query, QueryBuilder, T, withParams(), processParamsForPostgreSQL(), processSQL(), T (+9 more)

### Community 14 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 15 - "PortfolioAllocation"
Cohesion: 0.14
Nodes (16): BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, buildSelectedExternalAsset() (+8 more)

### Community 16 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.19
Nodes (11): Asset, buildAssetInsertValue(), buildAssetInsertValues(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, ToPointerSlice() (+3 more)

### Community 18 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 19 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (23): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+15 more)

### Community 20 - "asset_rest_model.go"
Cohesion: 0.19
Nodes (19): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+11 more)

### Community 21 - "pagination_cursor.go"
Cohesion: 0.15
Nodes (18): encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor(), readPaginationCursor(), validatePaginationCursorScope(), go_pkg_bytes, go_pkg_encoding_base64, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_url (+10 more)

### Community 22 - "context.Context"
Cohesion: 0.20
Nodes (8): classifyObservationWriteError(), preparePortfolioAllocationTempInserts(), PortfolioObservationTimestamp, context.Context, BuildUniqueConstraintViolationError(), IsUniqueConstraintViolation(), ToSQLTransactionalContext(), PortfolioAllocationRDBMSRepository

### Community 23 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.16
Nodes (16): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceChartResponseDTS (+8 more)

### Community 24 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.31
Nodes (7): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 25 - "go_pkg_fmt"
Cohesion: 0.15
Nodes (7): PortfolioNotFoundError, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, go_pkg_time

### Community 26 - "BuildAppError"
Cohesion: 0.20
Nodes (12): decodePaginationCursorPayload(), ExternalAsset, ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppError() (+4 more)

### Community 27 - "PortfolioAllocationRESTController"
Cohesion: 0.15
Nodes (5): PaginatedResponse, T, github.com/gin-gonic/gin.HandlersChain, RESTRoute, PortfolioAllocationRESTController

### Community 28 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 29 - "go_pkg_net_http"
Cohesion: 0.33
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_net_http

### Community 30 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 31 - "AssetDomService"
Cohesion: 0.16
Nodes (8): AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), AssetDomService, TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), AssetIntegrationServicesPerSource

### Community 32 - "go_pkg_strings"
Cohesion: 0.20
Nodes (10): go_pkg_database_sql, go_pkg_github_com_lib_pq, go_pkg_strings, database/sql.Row, database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), prepareBulkInsertValues() (+2 more)

### Community 33 - "SQLTransactionalContext"
Cohesion: 0.23
Nodes (5): SQLTransactionalContext, withTransaction(), contextKey, TransactionalContext, AllocationPlanRDBMSRepository

### Community 34 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 35 - "asset_integration_test.go"
Cohesion: 0.19
Nodes (14): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+6 more)

### Community 36 - ".buildAppComponents"
Cohesion: 0.22
Nodes (11): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository() (+3 more)

### Community 37 - "PortfolioRESTController"
Cohesion: 0.21
Nodes (9): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService, PortfolioDTS (+1 more)

### Community 38 - "ObservationPageCursor"
Cohesion: 0.30
Nodes (7): addObservationCursorFilter(), buildPortfolioAllocationsByObservationPage(), buildPortfolioObservationPage(), T, ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage

### Community 39 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 40 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 41 - "assert_db.go"
Cohesion: 0.50
Nodes (8): github.com/go-ozzo/ozzo-dbx.NullStringMap, assertAllRows(), AssertDBWithQuery(), assertRowColumn(), assertSingleRow(), executeQueryForMultipleRows(), validateRowCount(), AssertableNullStringMap

### Community 42 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 43 - "CustomValidationErrorsBuilder"
Cohesion: 0.36
Nodes (4): buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 44 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 45 - "go_pkg_github_com_golang_glog"
Cohesion: 0.33
Nodes (3): go_pkg_flag, go_pkg_github_com_golang_glog, go_pkg_net

### Community 46 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

## Knowledge Gaps
- **17 isolated node(s):** `MapIterator`, `contextKey`, `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `paginationCursorValidationDTS` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 104 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `reflect.Type`, `Adapter`, `infra.go`, `.buildAppComponents`, `go_pkg_testing`, `go_pkg_github_com_golang_glog`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `go_pkg_fmt`, `go_pkg_net_http`?**
  _High betweenness centrality (0.037) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `IsZeroValue`, `Adapter`, `AllocationPlan`, `BuildCleanupFunctionBuilder`, `PortfolioAllocation`, `asset_rest_model.go`, `AssetDomService`?**
  _High betweenness centrality (0.030) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `MapIterator`, `contextKey`, `AssetSearchQueryDTS` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `IsZeroValue` be split into smaller, more focused modules?**
  _Cohesion score 0.05427905427905428 - nodes in this community are weakly interconnected._
- **Should `reflect.Type` be split into smaller, more focused modules?**
  _Cohesion score 0.061343204653622425 - nodes in this community are weakly interconnected._
- **Should `Portfolio` be split into smaller, more focused modules?**
  _Cohesion score 0.06836158192090395 - nodes in this community are weakly interconnected._