# Graph Report - go  (2026-11-01)

## Corpus Check
- 132 files · ~74,313 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1244 nodes · 3895 edges · 63 communities (52 shown, 11 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 290 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `598579bd`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- portfolio_allocation_pagination_integration_test.go
- Adapter
- AllocationPlanDomService
- .buildAppComponents
- reflect.Type
- context.Context
- deferCloseResponseBody
- FetchWithDBQuery
- go_pkg_fmt
- PortfolioAllocation
- go_pkg_net_http
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- PropagateAsAppErrorWithNewMessage
- NewOrderedMapIterator
- BuildAppError
- Asset
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- BindAndValidateJSONWithInvalidResponse
- AllocationPlan
- golang_ext_util_concurrent.go
- portfolio_observation_ownership_migration_integration_test.go
- assert_json_extension.go
- infra.go
- GinServer
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- portfolio_rest_mapping.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- go_pkg_time
- go_pkg_strings
- PortfolioRESTController
- testing.T
- App
- asset_integration_service_yahoo.go
- BuildCleanupFunctionBuilder
- BuildCustomValidationErrorsBuilder
- Portfolio
- error.go
- PortfolioAllocationRDBMSRepository
- http_client.go
- allocation_plan_domain_service.go
- HierarchicalId
- ForeignKeyConstraintViolationError
- go_pkg_github_com_benizzio_open_asset_allocator_infra_util_urlutil
- IsZeroValue
- ObservationPageCursor
- AllocationStructure
- pagination_cursor.go
- AllocationPlanRESTController
- golang_ext_util_slice.go
- joinAny

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `BuildCleanupFunctionBuilder()` - 44 edges
3. `PropagateAsAppErrorWithNewMessage()` - 42 edges
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
- `allocationPlanCursorFromPayload()` --references--> `AllocationPlanPageCursor`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/allocation_plan.go

## Import Cycles
- None detected.

## Communities (63 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.23
Nodes (9): validateAllocationPlanPlannedDate(), readHistoryObservationTimestampId(), github.com/gin-gonic/gin.Context, HandleAPIError(), handleDomainError(), handleInfrastructureError(), SendDataNotFoundResponse(), ParseInt64() (+1 more)

### Community 1 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.11
Nodes (44): go_pkg_net_url, net/url.Values, time.Time, HasQueryParameter(), RawQueryHasParameter(), allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities() (+36 more)

### Community 2 - "Adapter"
Cohesion: 0.06
Nodes (25): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, buildPingContext(), Adapter, QueryBuilder, T (+17 more)

### Community 3 - "AllocationPlanDomService"
Cohesion: 0.20
Nodes (7): AllocationPlanRepository, BuildAllocationPlanDomService(), AllocationPlanDomService, buildExternalAssetQuoteValidationError(), BuildAppErrorFormattedUnconverted(), BuildDomainValidationError(), AppError

### Community 4 - ".buildAppComponents"
Cohesion: 0.14
Nodes (17): BuildAllocationPlanRESTController(), BuildDivergenceAnalysisRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildPortfolioDivergenceAnalysisAppService(), BuildPortfolioAllocationManagementAppService() (+9 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.10
Nodes (32): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+24 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.15
Nodes (28): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+20 more)

### Community 8 - "FetchWithDBQuery"
Cohesion: 0.23
Nodes (10): github.com/go-ozzo/ozzo-dbx.Params, TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), postAsset(), FetchWithDBQuery(), assertPortfolioHistoryTagValidationCount(), TestPortfolioHistoryOrderingIsChronologicalAndDeterministic(), createDBCleanupFunctionMulti() (+2 more)

### Community 9 - "go_pkg_fmt"
Cohesion: 0.13
Nodes (10): PortfolioNotFoundError, go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_flag, go_pkg_fmt, database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement() (+2 more)

### Community 10 - "PortfolioAllocation"
Cohesion: 0.20
Nodes (7): BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), PortfolioAllocation, PortfolioObservationTimestamp, PortfolioAllocationDomService

### Community 11 - "go_pkg_net_http"
Cohesion: 0.28
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_github_com_golang_glog, go_pkg_net (+1 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.13
Nodes (20): malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_io (+12 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.18
Nodes (21): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+13 more)

### Community 15 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.17
Nodes (10): PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, BuildILikeSubstringPattern(), TestBuildILikeSubstringPattern(), ToPointerSlice(), AssetRDBMSRepository (+2 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "BuildAppError"
Cohesion: 0.24
Nodes (6): classifyObservationWriteError(), BuildAppError(), BuildUniqueConstraintViolationError(), IsUniqueConstraintViolation(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 18 - "Asset"
Cohesion: 0.06
Nodes (42): BuildAssetRESTController(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+34 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.25
Nodes (21): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+13 more)

### Community 21 - "BindAndValidateJSONWithInvalidResponse"
Cohesion: 0.26
Nodes (9): bindAndValidatePortfolioSnapshot(), go_pkg_github_com_go_playground_validator_v10, sendValidationErrorResponse(), BindAndValidateJSONWithInvalidResponse(), deduplicateValidationMessages(), validator.ValidationErrors, RespondWithCustomValidationErrors(), BindAndValidateQueryWithInvalidResponse() (+1 more)

### Community 22 - "AllocationPlan"
Cohesion: 0.16
Nodes (13): GetPlanType(), PlanType, AllocationPlan, AllocationPlanOwnership, AllocationPlanPage, AllocationPlanPageCursor, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow() (+5 more)

### Community 23 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 24 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.38
Nodes (9): database/sql.Tx, migratedObservation, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations() (+1 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.06
Nodes (38): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_path_filepath (+30 more)

### Community 27 - "GinServer"
Cohesion: 0.21
Nodes (5): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController, GinServerConfiguration

### Community 28 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.39
Nodes (6): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 29 - "portfolio_rest_mapping.go"
Cohesion: 0.10
Nodes (43): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+35 more)

### Community 40 - "go_pkg_time"
Cohesion: 0.39
Nodes (4): go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_shopspring_decimal, go_pkg_time

### Community 41 - "go_pkg_strings"
Cohesion: 0.13
Nodes (17): BuildAllocationPlanRepository(), assetRowScanner(), go_pkg_github_com_lib_pq, go_pkg_strings, database/sql.Row, database/sql.Rows, github.com/go-ozzo/ozzo-dbx.Builder, github.com/go-ozzo/ozzo-dbx.Executor (+9 more)

### Community 42 - "PortfolioRESTController"
Cohesion: 0.16
Nodes (7): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService, github.com/gin-gonic/gin.HandlersChain, RESTRoute, PortfolioRESTController

### Community 43 - "testing.T"
Cohesion: 0.07
Nodes (51): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), testing.T, ValueToString(), TestGetExternalAssetsFailure(), TestGetExternalAssetsSuccess() (+43 more)

### Community 44 - "App"
Cohesion: 0.12
Nodes (18): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os, go_pkg_os_signal, go_pkg_syscall, os.Signal, Configuration (+10 more)

### Community 45 - "asset_integration_service_yahoo.go"
Cohesion: 0.24
Nodes (9): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), BuildYahooFinanceAssetIntegrationClient(), YahooFinanceAssetIntegrationClient, TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient (+1 more)

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.20
Nodes (32): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+24 more)

### Community 47 - "BuildCustomValidationErrorsBuilder"
Cohesion: 0.33
Nodes (5): buildCustomValidationError(), BuildCustomValidationErrorsBuilder(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 48 - "Portfolio"
Cohesion: 0.23
Nodes (8): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioDTS

### Community 49 - "error.go"
Cohesion: 0.11
Nodes (19): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, golang.org/x/text/currency.Unit (+11 more)

### Community 50 - "PortfolioAllocationRDBMSRepository"
Cohesion: 0.29
Nodes (6): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), preparePortfolioAllocationTempInserts(), portfolioAllocationJoinedRowDTS, PortfolioAllocationRDBMSRepository

### Community 51 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 52 - "allocation_plan_domain_service.go"
Cohesion: 0.23
Nodes (15): AllocationHierarchy, appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData(), readValidationData(), stripRootFromBranches() (+7 more)

### Community 53 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 54 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 56 - "IsZeroValue"
Cohesion: 0.19
Nodes (8): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), PlannedAllocation, AssetsPerTicker, T, IsZeroValue(), IsEmptyString(), ParseableInt

### Community 57 - "ObservationPageCursor"
Cohesion: 0.24
Nodes (9): addObservationCursorFilter(), buildPortfolioAllocationHistoryPage(), buildPortfolioObservationPage(), T, ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioAllocationRepository, PortfolioObservationPage (+1 more)

### Community 58 - "AllocationStructure"
Cohesion: 0.42
Nodes (8): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, AllocationHierarchyLevelDTS, AllocationStructureDTS

### Community 59 - "pagination_cursor.go"
Cohesion: 0.18
Nodes (17): decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor(), readPaginationCursor(), respondWithPaginationCursorValidationError() (+9 more)

### Community 60 - "AllocationPlanRESTController"
Cohesion: 0.31
Nodes (4): allocationPlanCursorFromPayload(), PaginatedResponse, T, AllocationPlanRESTController

### Community 61 - "golang_ext_util_slice.go"
Cohesion: 0.36
Nodes (7): cleanNilAllocations(), CleanNilPointersInSlice(), DereferenceSliceContent(), T, IsSlice(), ReverseSlice(), SliceContainsZeroValue()

### Community 62 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 105 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `.buildAppComponents`, `reflect.Type`, `go_pkg_time`, `go_pkg_net_http`, `asset_integration_service_yahoo.go`, `infra.go`, `GinServer`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `go_pkg_fmt`, `App`, `go_pkg_strings`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Why does `GinServer` connect `GinServer` to `go_pkg_net_http`, `App`?**
  _High betweenness centrality (0.031) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `portfolio_allocation_pagination_integration_test.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11100832562442182 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.060496067755595885 - nodes in this community are weakly interconnected._