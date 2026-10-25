# Graph Report - go  (2026-10-25)

## Corpus Check
- 131 files · ~73,616 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1239 nodes · 3888 edges · 54 communities (42 shown, 12 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 290 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `d35d91ca`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- AllocationPlan
- Adapter
- IsZeroValue
- rdbms_adapter.go
- reflect.Type
- context.Context
- deferCloseResponseBody
- testing.T
- portfolio_allocation_pagination_integration_test.go
- BuildCleanupFunctionBuilder
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- HierarchicalId
- NewOrderedMapIterator
- PlannedAllocation
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- logFanOutConsumer
- PropagateAsAppErrorWithNewMessage
- allocation_plan_rest_model.go
- assert_json_extension.go
- go_pkg_fmt
- GinServer
- BuildAppError
- asset_rest_model.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- PortfolioObservationTimestamp
- portfolio_observation_ownership_migration_integration_test.go
- portfolio_rest_mapping.go
- pagination_cursor.go
- App
- ExternalAsset
- AssertDBWithQueryMultipleRows
- go_pkg_encoding_json
- go_pkg_context
- FetchWithDBQuery
- ForeignKeyConstraintViolationError
- AllocationPlanDomService
- RESTRoute
- CustomValidationErrorsBuilder
- mapPortfolioAllocationRows

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
- `AssetRESTController` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go

## Import Cycles
- None detected.

## Communities (54 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.13
Nodes (24): validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), bindAndValidatePortfolioSnapshot(), readHistoryObservationTimestampId(), github.com/gin-gonic/gin.Context, HandleAPIError() (+16 more)

### Community 1 - "AllocationPlan"
Cohesion: 0.18
Nodes (13): allocationPlanCursorFromPayload(), GetPlanType(), PlanType, AllocationPlan, AllocationPlanOwnership, AllocationPlanPage, AllocationPlanPageCursor, buildAllocationPlanFromRow() (+5 more)

### Community 2 - "Adapter"
Cohesion: 0.05
Nodes (29): assetRowScanner(), database/sql.DB, database/sql.Result, database/sql.Rows, github.com/go-ozzo/ozzo-dbx.Builder, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Executor, github.com/go-ozzo/ozzo-dbx.Query (+21 more)

### Community 3 - "IsZeroValue"
Cohesion: 0.05
Nodes (48): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+40 more)

### Community 4 - "rdbms_adapter.go"
Cohesion: 0.11
Nodes (18): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_lib_pq (+10 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.05
Nodes (68): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+60 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.11
Nodes (34): go_pkg_io, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+26 more)

### Community 8 - "testing.T"
Cohesion: 0.12
Nodes (33): testing.T, TestGetExternalAssetsFailure(), TestGetExternalAssetsSuccess(), getExternalAssets(), capturedPaginationQuery, explainPaginationQuery(), TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans(), buildPaginationCaptureAdapter() (+25 more)

### Community 9 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.10
Nodes (46): go_pkg_net_url, net/url.Values, time.Time, NullTime, HasQueryParameter(), RawQueryHasParameter(), allocationPlanEndpoint(), allocationPlanPageNames() (+38 more)

### Community 10 - "BuildCleanupFunctionBuilder"
Cohesion: 0.15
Nodes (24): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+16 more)

### Community 11 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.22
Nodes (12): go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin (+4 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.25
Nodes (10): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_net_http, go_pkg_net_http_httptest (+2 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.16
Nodes (23): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+15 more)

### Community 15 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PlannedAllocation"
Cohesion: 0.20
Nodes (8): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), PlannedAllocation, SQLTransactionalContext, ToSQLTransactionalContext(), contextKey, TransactionalContext, AllocationPlanRDBMSRepository

### Community 18 - "PortfolioAllocation"
Cohesion: 0.17
Nodes (9): BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, PortfolioAllocation, PortfolioAllocationRepository, BuildPortfolioAllocationDomService() (+1 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.27
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 22 - "logFanOutConsumer"
Cohesion: 0.10
Nodes (17): go_pkg_github_com_nhatthm_httpmock, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry(), setupAndRunTests() (+9 more)

### Community 23 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.22
Nodes (10): Asset, buildAssetInsertValue(), buildAssetInsertValues(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, ToPointerSlice() (+2 more)

### Community 24 - "allocation_plan_rest_model.go"
Cohesion: 0.18
Nodes (18): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+10 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "go_pkg_fmt"
Cohesion: 0.15
Nodes (10): go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_net, go_pkg_path_filepath, RequestOption, ConfigLogger() (+2 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "BuildAppError"
Cohesion: 0.16
Nodes (13): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), classifyObservationWriteError(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppError(), BuildAppErrorFormatted() (+5 more)

### Community 29 - "asset_rest_model.go"
Cohesion: 0.20
Nodes (19): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+11 more)

### Community 40 - "PortfolioObservationTimestamp"
Cohesion: 0.19
Nodes (10): addObservationCursorFilter(), buildPortfolioAllocationHistoryPage(), buildPortfolioObservationPage(), T, preparePortfolioAllocationTempInserts(), ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage (+2 more)

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 42 - "portfolio_rest_mapping.go"
Cohesion: 0.15
Nodes (27): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+19 more)

### Community 43 - "pagination_cursor.go"
Cohesion: 0.15
Nodes (19): PaginatedResponse, T, decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor() (+11 more)

### Community 44 - "App"
Cohesion: 0.09
Nodes (24): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_github_com_benizzio_open_asset_allocator_infra_util (+16 more)

### Community 45 - "ExternalAsset"
Cohesion: 0.16
Nodes (17): YahooFinanceAssetIntegrationService, ExternalAsset, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient (+9 more)

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.20
Nodes (29): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+21 more)

### Community 51 - "go_pkg_context"
Cohesion: 0.15
Nodes (14): go_pkg_context, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_shopspring_decimal, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_golang_org_x_text_currency, go_pkg_slices (+6 more)

### Community 52 - "FetchWithDBQuery"
Cohesion: 0.31
Nodes (7): github.com/go-ozzo/ozzo-dbx.Params, FetchWithDBQuery(), assertPortfolioHistoryTagValidationCount(), TestPortfolioHistoryOrderingIsChronologicalAndDeterministic(), createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 56 - "AllocationPlanDomService"
Cohesion: 0.05
Nodes (41): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildDivergenceAnalysisRESTController(), BuildPortfolioRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService (+33 more)

### Community 59 - "CustomValidationErrorsBuilder"
Cohesion: 0.31
Nodes (5): go_pkg_github_com_go_playground_validator_v10, buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 61 - "mapPortfolioAllocationRows"
Cohesion: 0.52
Nodes (4): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), portfolioAllocationJoinedRowDTS

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 102 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `reflect.Type`, `context.Context`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `go_pkg_context`, `AllocationPlanDomService`, `go_pkg_fmt`, `GinServer`?**
  _High betweenness centrality (0.050) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `testing.T`, `rdbms_adapter.go`, `App`?**
  _High betweenness centrality (0.044) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `AllocationPlan`, `Adapter`, `IsZeroValue`, `PlannedAllocation`, `PortfolioAllocation`, `mapPortfolioAllocationRows`, `ExecuteDBQuery`, `allocation_plan_rest_model.go`, `AllocationPlanDomService`, `BuildAppError`, `asset_rest_model.go`?**
  _High betweenness centrality (0.028) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.13240418118466898 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.05182443151771549 - nodes in this community are weakly interconnected._