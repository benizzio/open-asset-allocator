# Graph Report - go  (2026-10-24)

## Corpus Check
- 129 files · ~72,484 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1210 nodes · 3774 edges · 54 communities (43 shown, 11 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 283 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `3c8e29f6`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- Portfolio
- Adapter
- AllocationPlanDomService
- go_pkg_strings
- reflect.Type
- portfolio_divergence_analysis_context.go
- testing.T
- portfolio_integration_test.go
- portfolio_allocation_pagination_integration_test.go
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- HierarchicalId
- NewOrderedMapIterator
- context.Context
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- IsZeroValue
- golang_ext_util_concurrent.go
- .buildAppComponents
- portfolio_rest_mapping.go
- assert_json_extension.go
- infra_db.go
- GinServer
- error.go
- Asset
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- ObservationPageCursor
- portfolio_observation_ownership_migration_integration_test.go
- ExternalAssetQuote
- pagination_cursor.go
- App
- go_pkg_context
- AssertDBWithQueryMultipleRows
- PaginatedResponse
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- http_client.go
- golang_ext_custom_slice_test.go
- rdbms_error.go
- AllocationPlanManagementAppService
- AssetDomService

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `BuildCleanupFunctionBuilder()` - 44 edges
3. `PropagateAsAppErrorWithNewMessage()` - 37 edges
4. `App` - 35 edges
5. `PortfolioAllocation` - 34 edges
6. `Asset` - 30 edges
7. `IsZeroValue()` - 30 edges
8. `ExecuteDBQuery()` - 29 edges
9. `PortfolioAllocationDomService` - 27 edges
10. `NewMapTree()` - 27 edges

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

## Communities (54 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.26
Nodes (8): github.com/gin-gonic/gin.Context, HandleAPIError(), handleDomainError(), handleInfrastructureError(), SendDataNotFoundResponse(), BindAndValidateQueryWithInvalidResponse(), ParseInt64(), AssetRESTController

### Community 1 - "Portfolio"
Cohesion: 0.16
Nodes (13): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), BuildPortfolioRESTController(), AllocationRepository, Portfolio, PortfolioRepository, BuildAllocationDomService() (+5 more)

### Community 2 - "Adapter"
Cohesion: 0.06
Nodes (26): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, buildPingContext(), Adapter, T, withParams() (+18 more)

### Community 3 - "AllocationPlanDomService"
Cohesion: 0.06
Nodes (45): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+37 more)

### Community 4 - "go_pkg_strings"
Cohesion: 0.11
Nodes (18): assetRowScanner(), go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_lib_pq, go_pkg_strings, database/sql.Row (+10 more)

### Community 5 - "reflect.Type"
Cohesion: 0.05
Nodes (48): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_github_com_go_playground_validator_v10, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type (+40 more)

### Community 6 - "portfolio_divergence_analysis_context.go"
Cohesion: 0.07
Nodes (36): BuildDivergenceAnalysisRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext(), buildDivergenceAnalysisContext() (+28 more)

### Community 7 - "testing.T"
Cohesion: 0.10
Nodes (49): testing.T, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+41 more)

### Community 8 - "portfolio_integration_test.go"
Cohesion: 0.24
Nodes (15): postPortfolioForValidationFailure(), putForValidationFailure(), TestPostPortfolio(), TestPostPortfolioFailureWithNameExceedingMaxLength(), TestPostPortfolioFailureWithoutMandatoryFields(), TestPostPortfolioWithAllocationStructure(), TestPutPortfolio(), TestPutPortfolioFailureWithNameExceedingMaxLength() (+7 more)

### Community 9 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.13
Nodes (38): allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture(), encodeTestCursor(), getAllocationPlanPage(), insertAllocationPlanPaginationRecord(), invalidAllocationPlanCursorTokens() (+30 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.19
Nodes (18): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetKnownAssetsRejectsInvalidTextSearch(), TestGetKnownAssetsWithTextSearch(), TestPostAsset(), TestPostAssetFailureWithMissingRequiredFields(), TestPostAssetRejectsNonZeroId() (+10 more)

### Community 11 - "infra.go"
Cohesion: 0.13
Nodes (11): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_net, go_pkg_path_filepath (+3 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.11
Nodes (23): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra (+15 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.31
Nodes (15): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+7 more)

### Community 15 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "context.Context"
Cohesion: 0.19
Nodes (11): preparePortfolioAllocationTempInserts(), context.Context, BuildAppError(), PropagateAsAppErrorWithNewMessage(), ToSQLTransactionalContext(), buildDatabaseConnection(), runFlywayTestcontainer(), SetupTestInfra() (+3 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.13
Nodes (14): BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows() (+6 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.17
Nodes (29): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPostAssetWithDuplicateTicker(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+21 more)

### Community 21 - "IsZeroValue"
Cohesion: 0.19
Nodes (12): validateAllocationPlanPlannedDate(), bindAndValidatePortfolioSnapshot(), sendValidationErrorResponse(), BindAndValidateJSONWithInvalidResponse(), deduplicateValidationMessages(), validator.ValidationErrors, RespondWithCustomValidationErrors(), BuildCustomValidationErrorsBuilder() (+4 more)

### Community 22 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 23 - ".buildAppComponents"
Cohesion: 0.11
Nodes (17): BuildAssetRESTController(), BuildPortfolioAllocationManagementAppService(), BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), github.com/go-ozzo/ozzo-dbx.Rows (+9 more)

### Community 24 - "portfolio_rest_mapping.go"
Cohesion: 0.06
Nodes (63): allocationPlanCursorFromPayload(), AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation() (+55 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra_db.go"
Cohesion: 0.08
Nodes (25): go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer (+17 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "error.go"
Cohesion: 0.27
Nodes (7): classifyObservationWriteError(), BuildAppErrorFormatted(), BuildUniqueConstraintViolationError(), logError(), newAppError(), IsUniqueConstraintViolation(), UniqueConstraintViolationError

### Community 29 - "Asset"
Cohesion: 0.15
Nodes (21): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+13 more)

### Community 40 - "ObservationPageCursor"
Cohesion: 0.42
Nodes (4): ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage, buildPortfolioAllocationHistoryPage()

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 42 - "ExternalAssetQuote"
Cohesion: 0.15
Nodes (17): YahooFinanceAssetIntegrationService, ExternalAssetQuote, BuildYahooFinanceAssetIntegrationService(), extractLastClose(), mapToExternalAssetQuote(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient (+9 more)

### Community 43 - "pagination_cursor.go"
Cohesion: 0.16
Nodes (20): decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor(), hasQueryParameter(), rawQueryHasParameter() (+12 more)

### Community 44 - "App"
Cohesion: 0.10
Nodes (23): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os (+15 more)

### Community 45 - "go_pkg_context"
Cohesion: 0.36
Nodes (5): go_pkg_context, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, go_pkg_time

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.19
Nodes (30): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+22 more)

### Community 49 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.24
Nodes (13): go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin (+5 more)

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 51 - "golang_ext_custom_slice_test.go"
Cohesion: 0.22
Nodes (7): demoStringer, TestCustomSlice_PrettyString_Empty(), TestCustomSlice_PrettyString_Int(), TestCustomSlice_PrettyString_Single(), TestCustomSlice_PrettyString_String(), TestCustomSlice_PrettyString_Stringer(), TestCustomSlice_PrettyString_Struct()

### Community 53 - "rdbms_error.go"
Cohesion: 0.25
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 56 - "AllocationPlanManagementAppService"
Cohesion: 0.18
Nodes (7): BuildAllocationPlanRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, github.com/gin-gonic/gin.HandlersChain, RESTRoute, TransactionManager, AllocationPlanRESTController

### Community 58 - "AssetDomService"
Cohesion: 0.14
Nodes (13): ExternalAsset, AssetIntegrationService, AssetRepository, mapToExternalAsset(), mapToExternalAssets(), YahooFinanceSearchQuoteDTS, BuildAssetDomService(), collectIntegrationServices() (+5 more)

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 100 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `reflect.Type`, `infra.go`, `go_pkg_context`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `.buildAppComponents`, `infra_db.go`, `GinServer`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `AllocationPlanDomService`, `go_pkg_strings`, `context.Context`, `PortfolioAllocation`, `BuildCleanupFunctionBuilder`, `.buildAppComponents`, `portfolio_rest_mapping.go`, `AssetDomService`, `error.go`?**
  _High betweenness centrality (0.036) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.05902980713033314 - nodes in this community are weakly interconnected._
- **Should `AllocationPlanDomService` be split into smaller, more focused modules?**
  _Cohesion score 0.05921325051759834 - nodes in this community are weakly interconnected._
- **Should `go_pkg_strings` be split into smaller, more focused modules?**
  _Cohesion score 0.1111111111111111 - nodes in this community are weakly interconnected._