# Graph Report - go  (2026-10-24)

## Corpus Check
- 130 files · ~73,246 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1226 nodes · 3844 edges · 64 communities (54 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 288 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `5fb65f9c`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- ParseableInt64
- Adapter
- allocation_plan_domain_service.go
- go_pkg_fmt
- reflect.Type
- portfolio_divergence_analysis_context.go
- deferCloseResponseBody
- portfolio_integration_test.go
- portfolio_allocation_pagination_integration_test.go
- testing.T
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- HierarchicalId
- NewOrderedMapIterator
- context.Context
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- IsZeroValue
- golang_ext_util_concurrent.go
- PropagateAsAppErrorWithNewMessage
- AllocationPlan
- assert_json_extension.go
- infra.go
- GinServer
- error.go
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
- Configuration
- go_pkg_time
- BuildCleanupFunctionBuilder
- Portfolio
- portfolio_rest_model.go
- portfolio_allocation_controller.go
- http_client.go
- App
- FetchWithDBQuery
- ForeignKeyConstraintViolationError
- BuildAppErrorFormattedUnconverted
- AllocationStructure
- RepositoryRDBMSAdapter
- golang_ext_custom_slice.go
- .buildAppComponents
- BuildCustomValidationErrorsBuilder
- golang_ext_util_slice.go
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
- `appendLevelDescription()` --references--> `CustomSlice`  [EXTRACTED]
  domain/service/allocation_plan_domain_service.go → langext/golang_ext_custom_slice.go
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `AllocationPlanRESTController` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `validateAllocationPlanPlannedDate()` --calls--> `BuildCustomValidationErrorsBuilder()`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → infra/validation/custom_validation_error_builder.go

## Import Cycles
- None detected.

## Communities (64 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.13
Nodes (21): validateAllocationPlanPlannedDate(), MapToAssetDTS(), MapToAssetDTSs(), readHistoryObservationTimestampId(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, HandleAPIError(), handleDomainError() (+13 more)

### Community 1 - "ParseableInt64"
Cohesion: 0.13
Nodes (12): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService, github.com/gin-gonic/gin.HandlersChain (+4 more)

### Community 2 - "Adapter"
Cohesion: 0.07
Nodes (20): database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, buildPingContext(), Adapter, T, withParams(), processParamsForPostgreSQL() (+12 more)

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.32
Nodes (12): AllocationHierarchy, appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData(), readValidationData(), validateHierarchySize() (+4 more)

### Community 4 - "go_pkg_fmt"
Cohesion: 0.19
Nodes (10): go_pkg_database_sql_driver, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_golang_glog, go_pkg_net, database/sql.Stmt, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement() (+2 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "portfolio_divergence_analysis_context.go"
Cohesion: 0.08
Nodes (30): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+22 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.11
Nodes (34): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+26 more)

### Community 8 - "portfolio_integration_test.go"
Cohesion: 0.14
Nodes (21): go_pkg_encoding_json, go_pkg_strconv, postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios() (+13 more)

### Community 9 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.13
Nodes (40): time.Time, allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture(), encodeTestCursor(), getAllocationPlanPage(), insertAllocationPlanPaginationRecord() (+32 more)

### Community 10 - "testing.T"
Cohesion: 0.13
Nodes (27): testing.T, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+19 more)

### Community 11 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.37
Nodes (6): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.13
Nodes (22): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util (+14 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.18
Nodes (21): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+13 more)

### Community 15 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "context.Context"
Cohesion: 0.17
Nodes (11): preparePortfolioAllocationTempInserts(), context.Context, database/sql.DB, BuildAppError(), SQLTransactionalContext, ToSQLTransactionalContext(), withTransaction(), contextKey (+3 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.18
Nodes (8): mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, PortfolioAllocation, PortfolioAllocationRepository, PortfolioObservationTimestamp, BuildPortfolioAllocationDomService(), PortfolioAllocationDomService

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.21
Nodes (22): go_pkg_net_url, TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource() (+14 more)

### Community 21 - "IsZeroValue"
Cohesion: 0.23
Nodes (6): AllocationPlanRepository, BuildAllocationPlanDomService(), AllocationPlanDomService, T, IsZeroValue(), ParseableInt

### Community 22 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 23 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.14
Nodes (15): Asset, buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetRDBMSRepository(), classifyObservationWriteError(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery() (+7 more)

### Community 24 - "AllocationPlan"
Cohesion: 0.10
Nodes (31): allocationPlanCursorFromPayload(), AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation() (+23 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.06
Nodes (37): go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_path_filepath, go_pkg_sync (+29 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "error.go"
Cohesion: 0.15
Nodes (11): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), buildExternalAssetQuoteValidationError(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppErrorFormatted(), BuildDomainValidationError() (+3 more)

### Community 29 - "asset_rest_model.go"
Cohesion: 0.09
Nodes (33): MapToAsset(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs(), MapToExternalAssetQuoteDTS(), mapToExternalAsset() (+25 more)

### Community 40 - "ObservationPageCursor"
Cohesion: 0.36
Nodes (5): buildPortfolioAllocationHistoryPage(), buildPortfolioObservationPage(), ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.16
Nodes (16): assetRowScanner(), go_pkg_database_sql, go_pkg_github_com_lib_pq, database/sql.Row, database/sql.Rows, database/sql.Tx, ReturningIntIdRowScanner(), ReturningIntIdSingleRowScanner() (+8 more)

### Community 42 - "portfolio_rest_mapping.go"
Cohesion: 0.20
Nodes (17): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToObservationTimestampDTS(), MapToPortfolioAllocation(), mapToPortfolioAllocationDTS(), MapToPortfolioAllocations() (+9 more)

### Community 43 - "pagination_cursor.go"
Cohesion: 0.16
Nodes (19): PaginatedResponse, T, decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor() (+11 more)

### Community 44 - "Configuration"
Cohesion: 0.15
Nodes (14): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_os, Configuration, RDBMSConfiguration, YahooFinanceConfiguration (+6 more)

### Community 45 - "go_pkg_time"
Cohesion: 0.23
Nodes (8): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), YahooFinanceAssetIntegrationClient, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, go_pkg_time

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.19
Nodes (33): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+25 more)

### Community 47 - "Portfolio"
Cohesion: 0.20
Nodes (11): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildPortfolioDivergenceAnalysisAppService(), PortfolioDivergenceAnalysisAppService, AnalysisOptions, Portfolio, PortfolioRepository (+3 more)

### Community 48 - "portfolio_rest_model.go"
Cohesion: 0.20
Nodes (11): mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS(), mapToPotentialDivergenceDTS(), mapToPotentialDivergenceDTSs(), AllocationPlanIdentifier, AllocationPlanIdentifierDTS (+3 more)

### Community 49 - "portfolio_allocation_controller.go"
Cohesion: 0.34
Nodes (7): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 51 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 52 - "FetchWithDBQuery"
Cohesion: 0.23
Nodes (10): github.com/go-ozzo/ozzo-dbx.Params, TestPostAssetWithDuplicateTicker(), explainPaginationQuery(), TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans(), FetchWithDBQuery(), assertPortfolioHistoryTagValidationCount(), TestPortfolioHistoryOrderingIsChronologicalAndDeterministic(), createDBCleanupFunctionMulti() (+2 more)

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 54 - "BuildAppErrorFormattedUnconverted"
Cohesion: 0.27
Nodes (4): DomainValidationError, BuildAppErrorFormattedUnconverted(), AppError, IsEmptyString()

### Community 55 - "AllocationStructure"
Cohesion: 0.42
Nodes (8): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, AllocationHierarchyLevelDTS, AllocationStructureDTS

### Community 56 - "RepositoryRDBMSAdapter"
Cohesion: 0.28
Nodes (7): BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), github.com/go-ozzo/ozzo-dbx.Rows, RepositoryRDBMSAdapter, AllocationRDBMSRepository

### Community 57 - "golang_ext_custom_slice.go"
Cohesion: 0.25
Nodes (6): stripRootFromBranches(), CustomSlice[T], CustomSlice, CustomSliceTable, T, joinAny()

### Community 58 - ".buildAppComponents"
Cohesion: 0.14
Nodes (16): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, AssetIntegrationService (+8 more)

### Community 59 - "BuildCustomValidationErrorsBuilder"
Cohesion: 0.33
Nodes (5): buildCustomValidationError(), BuildCustomValidationErrorsBuilder(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 60 - "golang_ext_util_slice.go"
Cohesion: 0.36
Nodes (7): cleanNilAllocations(), CleanNilPointersInSlice(), DereferenceSliceContent(), T, IsSlice(), ReverseSlice(), SliceContainsZeroValue()

### Community 61 - "mapPortfolioAllocationRows"
Cohesion: 0.52
Nodes (4): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), portfolioAllocationJoinedRowDTS

### Community 62 - "PortfolioAllocationRESTController"
Cohesion: 0.43
Nodes (3): net/url.Values, HasQueryParameter(), PortfolioAllocationRESTController

### Community 63 - "parseAssetTextSearch"
Cohesion: 0.47
Nodes (4): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch(), TestParseAssetTextSearch()

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 100 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `go_pkg_fmt`, `.buildAppComponents`, `reflect.Type`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Configuration`, `go_pkg_time`, `portfolio_allocation_controller.go`, `infra.go`, `GinServer`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `github.com/gin-gonic/gin.Context`, `portfolio_observation_ownership_migration_integration_test.go`, `PortfolioAllocation`, `mapPortfolioAllocationRows`, `ExecuteDBQuery`, `AllocationPlan`, `.buildAppComponents`, `error.go`, `asset_rest_model.go`, `parseAssetTextSearch`?**
  _High betweenness centrality (0.040) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `context.Context`, `App`, `go_pkg_fmt`, `Configuration`?**
  _High betweenness centrality (0.033) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.13213213213213212 - nodes in this community are weakly interconnected._
- **Should `ParseableInt64` be split into smaller, more focused modules?**
  _Cohesion score 0.12987012987012986 - nodes in this community are weakly interconnected._