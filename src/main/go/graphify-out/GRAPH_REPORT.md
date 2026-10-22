# Graph Report - go  (2026-10-22)

## Corpus Check
- 123 files · ~66,412 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1130 nodes · 3469 edges · 59 communities (47 shown, 12 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 264 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `ab1b7285`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- Portfolio
- Adapter
- allocation_plan_domain_service.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- reflect.Type
- portfolio_divergence_analysis_context.go
- testing.T
- insertTestPortfolio
- createBulkInsertPreparedStatement
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- HierarchicalId
- NewOrderedMapIterator
- context.Context
- PortfolioAllocationDomService
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- BindAndValidateJSONWithInvalidResponse
- golang_ext_util_concurrent.go
- T
- AllocationPlan
- assert_json_extension.go
- logFanOutConsumer
- GinServer
- portfolio_rest_mapping.go
- Asset
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- .buildAppComponents
- database/sql.Tx
- Configuration
- PortfolioAllocation
- App
- portfolio_divergence_analysis_service.go
- AssertDBWithQueryMultipleRows
- T
- PotentialDivergence
- go_pkg_net_http
- go_pkg_fmt
- golang_ext_custom_slice_test.go
- CustomValidationErrorsBuilder
- ForeignKeyConstraintViolationError
- T
- AllocationDomService
- RESTRoute
- QueryBuilder
- BuildAssetDomService

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
- `mapToAllocationPlanAssetDTS()` --references--> `Asset`  [EXTRACTED]
  api/rest/model/allocation_plan_rest_model.go → domain/asset.go

## Import Cycles
- None detected.

## Communities (59 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.17
Nodes (15): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), github.com/gin-gonic/gin.Context, HandleAPIError(), handleDomainError(), handleInfrastructureError(), SendDataNotFoundResponse() (+7 more)

### Community 1 - "Portfolio"
Cohesion: 0.20
Nodes (10): BuildDivergenceAnalysisRESTController(), PortfolioAnalysisConfigurationAppService, buildDivergenceAnalysis(), PortfolioDivergenceAnalysisAppService, AnalysisOptions, Portfolio, PortfolioRepository, BuildPortfolioDomService() (+2 more)

### Community 2 - "Adapter"
Cohesion: 0.14
Nodes (8): database/sql.DB, database/sql.Result, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey, TransactionalContext

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.07
Nodes (39): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationStructure, AllocationPlanRepository (+31 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.19
Nodes (15): assetRowScanner(), go_pkg_context, go_pkg_database_sql, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext (+7 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (43): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+35 more)

### Community 6 - "portfolio_divergence_analysis_context.go"
Cohesion: 0.19
Nodes (11): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), potentialDivergencesPerHierarchicalId (+3 more)

### Community 7 - "testing.T"
Cohesion: 0.13
Nodes (41): testing.T, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+33 more)

### Community 8 - "insertTestPortfolio"
Cohesion: 0.21
Nodes (14): TestPostAsset(), putForValidationFailure(), TestPostPortfolio(), TestPostPortfolioWithAllocationStructure(), TestPutPortfolio(), TestPutPortfolioFailureWithNameExceedingMaxLength(), TestPutPortfolioFailureWithoutMandatoryFields(), TestPutPortfolioWithAllocationStructure() (+6 more)

### Community 9 - "createBulkInsertPreparedStatement"
Cohesion: 0.40
Nodes (5): database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), prepareBulkInsertValues(), quoteIdentifiers()

### Community 10 - "asset_integration_test.go"
Cohesion: 0.15
Nodes (21): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+13 more)

### Community 11 - "infra.go"
Cohesion: 0.13
Nodes (12): go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_net, go_pkg_path_filepath, go_pkg_time (+4 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.18
Nodes (16): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require (+8 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (22): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+14 more)

### Community 15 - "HierarchicalId"
Cohesion: 0.14
Nodes (7): HierarchicalId, go_pkg_database_sql_driver, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "context.Context"
Cohesion: 0.05
Nodes (48): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, BuildAllocationPlanRepository() (+40 more)

### Community 18 - "PortfolioAllocationDomService"
Cohesion: 0.21
Nodes (5): AllocationHierarchyLevel, PortfolioAllocationRepository, PortfolioObservationTimestamp, BuildPortfolioAllocationDomService(), PortfolioAllocationDomService

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.18
Nodes (27): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+19 more)

### Community 21 - "BindAndValidateJSONWithInvalidResponse"
Cohesion: 0.22
Nodes (10): validateAllocationPlanPlannedDate(), MapToPortfolioObservationTimestamp(), bindAndValidatePortfolioSnapshot(), go_pkg_github_com_go_playground_validator_v10, sendValidationErrorResponse(), BindAndValidateJSONWithInvalidResponse(), deduplicateValidationMessages(), validator.ValidationErrors (+2 more)

### Community 22 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 23 - "T"
Cohesion: 0.19
Nodes (10): BuildQueryInTransaction(), T, processParamsForPostgreSQL(), T, IsSlice(), SingleRowScanner, SQLTransactionalQueryBuilder, SQLTransactionalQueryBuilder[T] (+2 more)

### Community 24 - "AllocationPlan"
Cohesion: 0.11
Nodes (33): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+25 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "logFanOutConsumer"
Cohesion: 0.10
Nodes (18): go_pkg_github_com_nhatthm_httpmock, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry() (+10 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "portfolio_rest_mapping.go"
Cohesion: 0.18
Nodes (22): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS(), mapToObservationTimestampDTS(), MapToPortfolioAllocation() (+14 more)

### Community 29 - "Asset"
Cohesion: 0.06
Nodes (44): BuildAssetRESTController(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+36 more)

### Community 40 - ".buildAppComponents"
Cohesion: 0.21
Nodes (10): BuildAllocationPlanRESTController(), BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), BuildPortfolioDivergenceAnalysisAppService(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService (+2 more)

### Community 41 - "database/sql.Tx"
Cohesion: 0.36
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 42 - "Configuration"
Cohesion: 0.14
Nodes (17): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), BuildYahooFinanceAssetIntegrationClient(), YahooFinanceAssetIntegrationClient, TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_net_url (+9 more)

### Community 43 - "PortfolioAllocation"
Cohesion: 0.19
Nodes (11): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, PortfolioAllocation, T, IsZeroValue() (+3 more)

### Community 44 - "App"
Cohesion: 0.18
Nodes (10): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel() (+2 more)

### Community 45 - "portfolio_divergence_analysis_service.go"
Cohesion: 0.29
Nodes (10): getDivergenceAnalysisContextValue(), getPotentialDivergenceMapContextValue(), buildPotentialDivergence(), calculateCurrentDivergenceValuesFromReferencedPlan(), calculateDivergenceValue(), checkAndGeneratePotentialDivergencesOnHierarchy(), generateAndAttachPotentialDivergenceForPlannedAllocation(), generatePotentialDivergencesFromAllocationPlanSetDifference() (+2 more)

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.20
Nodes (29): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+21 more)

### Community 47 - "T"
Cohesion: 0.26
Nodes (7): github.com/go-ozzo/ozzo-dbx.Query, T, withParams(), processSQL(), QueryExecutor, QueryExecutor[T], RowScanner

### Community 48 - "PotentialDivergence"
Cohesion: 0.29
Nodes (6): getAllocationIterationContextValue(), getHierarchySubIterationContextValue(), attachToRootIfTopLevel(), buildAndAttachPotentialDivergenceIfNotExists(), DivergenceAnalysis, PotentialDivergence

### Community 49 - "go_pkg_net_http"
Cohesion: 0.27
Nodes (10): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin (+2 more)

### Community 50 - "go_pkg_fmt"
Cohesion: 0.11
Nodes (16): TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_encoding_json, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_root, RequestOption, ConfigLogger() (+8 more)

### Community 51 - "golang_ext_custom_slice_test.go"
Cohesion: 0.22
Nodes (7): demoStringer, TestCustomSlice_PrettyString_Empty(), TestCustomSlice_PrettyString_Int(), TestCustomSlice_PrettyString_Single(), TestCustomSlice_PrettyString_String(), TestCustomSlice_PrettyString_Stringer(), TestCustomSlice_PrettyString_Struct()

### Community 52 - "CustomValidationErrorsBuilder"
Cohesion: 0.36
Nodes (4): buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 55 - "AllocationDomService"
Cohesion: 0.38
Nodes (4): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 57 - "QueryBuilder"
Cohesion: 0.48
Nodes (3): github.com/go-ozzo/ozzo-dbx.DB, QueryBuilder, QueryBuilder[T]

### Community 58 - "BuildAssetDomService"
Cohesion: 0.40
Nodes (4): AssetRepository, BuildAssetDomService(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch()

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 94 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `.buildAppComponents`, `Configuration`, `infra.go`, `go_pkg_testing`, `go_pkg_net_http`, `GinServer`?**
  _High betweenness centrality (0.054) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `PortfolioAllocation`, `context.Context`, `BuildCleanupFunctionBuilder`, `AllocationPlan`?**
  _High betweenness centrality (0.047) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `HierarchicalId` to `AllocationPlan`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.1383399209486166 - nodes in this community are weakly interconnected._
- **Should `allocation_plan_domain_service.go` be split into smaller, more focused modules?**
  _Cohesion score 0.06927551560021153 - nodes in this community are weakly interconnected._