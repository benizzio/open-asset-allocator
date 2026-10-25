# Graph Report - go  (2026-10-25)

## Corpus Check
- 130 files · ~73,269 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1228 nodes · 3850 edges · 64 communities (52 shown, 12 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 288 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `98d818c1`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- AllocationDomService
- T
- AllocationPlanDomService
- go_pkg_strings
- reflect.Type
- context.Context
- deferCloseResponseBody
- testing.T
- portfolio_allocation_pagination_integration_test.go
- asset_integration_test.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- HierarchicalId
- NewOrderedMapIterator
- PropagateAsAppErrorWithNewMessage
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- IsZeroValue
- golang_ext_util_concurrent.go
- BuildQuery
- AllocationPlan
- assert_json_extension.go
- infra.go
- GinServer
- ExternalAsset
- asset_rest_model.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- ObservationPageCursor
- portfolio_observation_ownership_migration_integration_test.go
- portfolio_rest_mapping.go
- pagination_cursor.go
- App
- YahooFinanceAssetIntegrationClient
- BuildCleanupFunctionBuilder
- Portfolio
- Adapter
- go_pkg_net_http
- go_pkg_time
- go_pkg_github_com_shopspring_decimal
- FetchWithDBQuery
- ForeignKeyConstraintViolationError
- setupAndRunTests
- asset_integration_client_yahoo_model.go
- .buildAppComponents
- RESTRoute
- Asset
- CustomValidationErrorsBuilder
- MapToPortfolioDTS
- mapPortfolioAllocationRows
- PortfolioAllocationRESTController
- parseAssetTextSearch

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
- `allocationPlanCursorFromPayload()` --references--> `AllocationPlanPageCursor`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/allocation_plan.go

## Import Cycles
- None detected.

## Communities (64 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.14
Nodes (21): validateAllocationPlanPlannedDate(), bindAndValidatePortfolioSnapshot(), readHistoryObservationTimestampId(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, HandleAPIError(), handleDomainError(), handleInfrastructureError() (+13 more)

### Community 1 - "AllocationDomService"
Cohesion: 0.38
Nodes (4): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 2 - "T"
Cohesion: 0.11
Nodes (17): github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, QueryBuilder, T, withParams(), processParamsForPostgreSQL(), processSQL(), T (+9 more)

### Community 3 - "AllocationPlanDomService"
Cohesion: 0.06
Nodes (44): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+36 more)

### Community 4 - "go_pkg_strings"
Cohesion: 0.14
Nodes (15): assetRowScanner(), go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_lib_pq, go_pkg_strings, database/sql.Row, database/sql.Rows (+7 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (41): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+33 more)

### Community 6 - "context.Context"
Cohesion: 0.09
Nodes (38): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+30 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.12
Nodes (33): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+25 more)

### Community 8 - "testing.T"
Cohesion: 0.12
Nodes (34): TestGetKnownAssetsRejectsExcessSearchTerms(), testing.T, TestBuildILikeSubstringPattern(), TestPostAsset(), TestPostAssetWithDuplicateTicker(), TestPostAssetWithoutExternalData(), postAsset(), postPortfolioForValidationFailure() (+26 more)

### Community 9 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.11
Nodes (43): go_pkg_net_url, net/url.Values, HasQueryParameter(), RawQueryHasParameter(), allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture() (+35 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.16
Nodes (18): getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch(), TestGetKnownAssetsWithTextSearch() (+10 more)

### Community 11 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.29
Nodes (7): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.14
Nodes (19): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra (+11 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (22): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+14 more)

### Community 15 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.17
Nodes (9): classifyObservationWriteError(), preparePortfolioAllocationTempInserts(), BuildAppError(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), IsUniqueConstraintViolation(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository (+1 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.17
Nodes (10): BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, PortfolioAllocation, PortfolioAllocationRepository, PortfolioObservationTimestamp (+2 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.28
Nodes (19): database/sql.NullString, StringPointerToNullString(), StringToNullString(), TestGetKnownAssetsIncludesPersistedExternalData(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+11 more)

### Community 21 - "IsZeroValue"
Cohesion: 0.32
Nodes (5): validateAssetIdentification(), T, IsZeroValue(), IsEmptyString(), ParseableInt

### Community 22 - "golang_ext_util_concurrent.go"
Cohesion: 0.09
Nodes (33): go_pkg_runtime, go_pkg_slices, go_pkg_sync, context.CancelFunc, sync.Mutex, sync.WaitGroup, testing.TB, I (+25 more)

### Community 23 - "BuildQuery"
Cohesion: 0.21
Nodes (7): BuildQuery(), BuildQueryInTransaction(), T, BuildILikeSubstringPattern(), ToPointerSlice(), AssetRDBMSRepository, S

### Community 24 - "AllocationPlan"
Cohesion: 0.08
Nodes (40): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+32 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.09
Nodes (23): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_net (+15 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "ExternalAsset"
Cohesion: 0.36
Nodes (7): ExternalAsset, ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppErrorFormatted()

### Community 29 - "asset_rest_model.go"
Cohesion: 0.27
Nodes (15): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+7 more)

### Community 40 - "ObservationPageCursor"
Cohesion: 0.29
Nodes (7): addObservationCursorFilter(), buildPortfolioAllocationHistoryPage(), buildPortfolioObservationPage(), T, ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.38
Nodes (9): database/sql.Tx, migratedObservation, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations() (+1 more)

### Community 42 - "portfolio_rest_mapping.go"
Cohesion: 0.19
Nodes (22): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS(), mapToObservationTimestampDTS(), MapToPortfolioAllocation() (+14 more)

### Community 43 - "pagination_cursor.go"
Cohesion: 0.16
Nodes (19): allocationPlanCursorFromPayload(), PaginatedResponse, T, decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor() (+11 more)

### Community 44 - "App"
Cohesion: 0.10
Nodes (23): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os (+15 more)

### Community 45 - "YahooFinanceAssetIntegrationClient"
Cohesion: 0.21
Nodes (10): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceSearchQuoteDTS (+2 more)

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.22
Nodes (29): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), assertPersistedAssetWithExternalData(), TestPostAllocationPlanRejectsWhitespaceName() (+21 more)

### Community 47 - "Portfolio"
Cohesion: 0.29
Nodes (5): Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioRDBMSRepository

### Community 48 - "Adapter"
Cohesion: 0.13
Nodes (9): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey (+1 more)

### Community 49 - "go_pkg_net_http"
Cohesion: 0.26
Nodes (10): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin (+2 more)

### Community 50 - "go_pkg_time"
Cohesion: 0.28
Nodes (5): go_pkg_encoding_json, go_pkg_strconv, go_pkg_time, RequestOption, WithHeader()

### Community 51 - "go_pkg_github_com_shopspring_decimal"
Cohesion: 0.20
Nodes (8): mapToExternalAsset(), mapToExternalAssetData(), ExternalAssetData, AssetExternalSource, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, T, ScanJsonColumn()

### Community 52 - "FetchWithDBQuery"
Cohesion: 0.23
Nodes (10): github.com/go-ozzo/ozzo-dbx.Params, TestPostAssetRejectsNonZeroId(), explainPaginationQuery(), TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans(), FetchWithDBQuery(), assertPortfolioHistoryTagValidationCount(), TestPortfolioHistoryOrderingIsChronologicalAndDeterministic(), createDBCleanupFunctionMulti() (+2 more)

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 54 - "setupAndRunTests"
Cohesion: 0.29
Nodes (9): go_pkg_github_com_nhatthm_httpmock, github.com/nhatthm/httpmock.Server, testing.M, setupAndRunTests(), TestMain(), BuildAndStartYahooFinanceMockServer(), GetYahooFinanceMockServer(), resetYahooFinanceMockServer() (+1 more)

### Community 55 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.52
Nodes (6): YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS

### Community 56 - ".buildAppComponents"
Cohesion: 0.14
Nodes (17): BuildAllocationPlanRESTController(), BuildDivergenceAnalysisRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildPortfolioAllocationManagementAppService(), BuildAllocationPlanRepository() (+9 more)

### Community 58 - "Asset"
Cohesion: 0.16
Nodes (10): BuildAssetRESTController(), Asset, AssetIntegrationService, AssetRepository, buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetDomService(), collectIntegrationServices() (+2 more)

### Community 59 - "CustomValidationErrorsBuilder"
Cohesion: 0.36
Nodes (4): buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 60 - "MapToPortfolioDTS"
Cohesion: 0.67
Nodes (4): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), PortfolioDTS

### Community 61 - "mapPortfolioAllocationRows"
Cohesion: 0.90
Nodes (4): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), portfolioAllocationJoinedRowDTS

### Community 63 - "parseAssetTextSearch"
Cohesion: 0.83
Nodes (3): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch()

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 101 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `reflect.Type`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Adapter`, `go_pkg_net_http`, `go_pkg_time`, `.buildAppComponents`, `infra.go`, `GinServer`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **Why does `IsZeroValue()` connect `IsZeroValue` to `github.com/gin-gonic/gin.Context`, `AllocationPlanDomService`, `reflect.Type`, `golang_ext_tree_test.go`, `Portfolio`, `PortfolioAllocation`, `AllocationPlan`, `ExternalAsset`, `mapPortfolioAllocationRows`, `PortfolioAllocationRESTController`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `go_pkg_strings`, `PropagateAsAppErrorWithNewMessage`, `PortfolioAllocation`, `go_pkg_github_com_shopspring_decimal`, `ExecuteDBQuery`, `mapPortfolioAllocationRows`, `BuildQuery`, `AllocationPlan`, `asset_rest_model.go`?**
  _High betweenness centrality (0.032) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.14082503556187767 - nodes in this community are weakly interconnected._
- **Should `T` be split into smaller, more focused modules?**
  _Cohesion score 0.10873440285204991 - nodes in this community are weakly interconnected._