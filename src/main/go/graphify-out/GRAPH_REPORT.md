# Graph Report - go  (2026-11-01)

## Corpus Check
- 130 files · ~73,180 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1214 nodes · 3778 edges · 60 communities (49 shown, 11 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 283 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `0eab1a51`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- portfolio_allocation_pagination_integration_test.go
- Adapter
- AllocationPlanDomService
- RepositoryRDBMSAdapter
- CustomFieldError
- context.Context
- testing.T
- insertTestPortfolio
- AllocationStructure
- net/http.Response
- go_pkg_github_com_golang_glog
- golang_ext_tree_test.go
- go_pkg_testing
- deferCloseResponseBody
- go_pkg_errors
- NewOrderedMapIterator
- Asset
- PropagateAsAppErrorWithNewMessage
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- error.go
- allocation_plan_rest_model.go
- T
- AllocationPlan
- assert_json_extension.go
- infra.go
- GinServer
- go_pkg_fmt
- PortfolioAllocation
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- rdbms_query.go
- asset_integration_service_yahoo.go
- Configuration
- App
- ObservationPageCursor
- AssertDBWithQueryMultipleRows
- custom_deep_validation.go
- reflect.Type
- portfolio_observation_ownership_migration_integration_test.go
- http_client.go
- hierarchical_id_test.go
- golang_ext_util_struct.go
- custom_validation_error_handling.go
- rdbms_adapter.go
- asset_integration_client_yahoo_model.go
- PortfolioAllocationRDBMSRepository
- go_pkg_database_sql
- go_pkg_reflect
- PaginatedResponse

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
- `allocationPlanCursorFromPayload()` --references--> `AllocationPlanPageCursor`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/historical_read_pagination.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `AllocationPlanAssetDTS` --references--> `ParseableInt64`  [EXTRACTED]
  api/rest/model/allocation_plan_rest_model.go → langext/golang_ext_parseable_int64.go

## Import Cycles
- None detected.

## Communities (60 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (63): allocationPlanCursorFromPayload(), validateAllocationPlanPlannedDate(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS() (+55 more)

### Community 1 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.12
Nodes (41): go_pkg_encoding_json, time.Time, allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture(), encodeTestCursor(), getAllocationPlanPage() (+33 more)

### Community 2 - "Adapter"
Cohesion: 0.15
Nodes (6): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, Adapter, SQLTransactionalContext, withTransaction()

### Community 3 - "AllocationPlanDomService"
Cohesion: 0.07
Nodes (34): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+26 more)

### Community 4 - "RepositoryRDBMSAdapter"
Cohesion: 0.20
Nodes (7): BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildPortfolioRepository(), go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, RepositoryRDBMSAdapter, AllocationRDBMSRepository, PortfolioRDBMSRepository

### Community 5 - "CustomFieldError"
Cohesion: 0.13
Nodes (3): github.com/go-playground/universal-translator.Translator, reflect.Kind, CustomFieldError

### Community 6 - "context.Context"
Cohesion: 0.06
Nodes (58): allocationIterationMappingContextValue, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue(), getDivergenceAnalysisContextValue() (+50 more)

### Community 7 - "testing.T"
Cohesion: 0.12
Nodes (32): testing.T, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds(), TestPostAllocationPlanUpdatesOwnedMetadata() (+24 more)

### Community 8 - "insertTestPortfolio"
Cohesion: 0.16
Nodes (18): postAssetForValidationFailure(), TestPostAsset(), TestPostAssetFailureWithMissingRequiredFields(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithoutExternalData(), postAsset(), putForValidationFailure(), TestPostPortfolio() (+10 more)

### Community 9 - "AllocationStructure"
Cohesion: 0.09
Nodes (16): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationStructure, HierarchicalId, go_pkg_database_sql_driver, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext (+8 more)

### Community 10 - "net/http.Response"
Cohesion: 0.25
Nodes (8): net/http.Response, getAssetsByTextSearch(), TestGetKnownAssetsRejectsInvalidTextSearch(), TestGetKnownAssetsWithTextSearch(), TestPutAssetFailureWithoutId(), TestPutAssetFailureWithoutRequiredFields(), TestPutAssetFailureWithZeroId(), putAsset()

### Community 11 - "go_pkg_github_com_golang_glog"
Cohesion: 0.21
Nodes (9): go_pkg_github_com_golang_glog, processParamsForPostgreSQL(), T, IsSlice(), SingleRowScanner, SQLTransactionalQueryBuilder, SQLTransactionalQueryBuilder[T], SQLTransactionalQueryExecutor (+1 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.17
Nodes (21): TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require (+13 more)

### Community 14 - "deferCloseResponseBody"
Cohesion: 0.10
Nodes (37): TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetDivergenceAnalysisOptions(), TestGetDivergenceAnalysisV2(), TestGetDivergenceAnalysisV2FullDivergence(), TestGetDivergenceAnalysisV2WhenHistoryHighestLevelHasExtraAllocation() (+29 more)

### Community 15 - "go_pkg_errors"
Cohesion: 0.16
Nodes (8): classifyObservationWriteError(), go_pkg_errors, go_pkg_github_com_lib_pq, BuildUniqueConstraintViolationError(), ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), IsUniqueConstraintViolation(), ParseableInt

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "Asset"
Cohesion: 0.23
Nodes (9): Asset, assetRowScanner(), buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetRDBMSRepository(), BuildQuery(), ToPointerSlice(), AssetRDBMSRepository (+1 more)

### Community 18 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.32
Nodes (4): BuildAppError(), PropagateAsAppErrorWithNewMessage(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.14
Nodes (35): github.com/go-ozzo/ozzo-dbx.Params, TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPostAssetWithDuplicateTicker(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted() (+27 more)

### Community 21 - "error.go"
Cohesion: 0.24
Nodes (8): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), golang.org/x/text/currency.Unit, BuildAppErrorFormatted(), logError(), newAppError(), UniqueConstraintViolationError

### Community 22 - "allocation_plan_rest_model.go"
Cohesion: 0.33
Nodes (14): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+6 more)

### Community 23 - "T"
Cohesion: 0.20
Nodes (9): github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, T, withParams(), QueryBuilder, QueryBuilder[T], QueryExecutor, QueryExecutor[T] (+1 more)

### Community 24 - "AllocationPlan"
Cohesion: 0.18
Nodes (13): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), GetPlanType(), PlanType, AllocationPlan, AllocationPlanOwnership, PlannedAllocation, buildAllocationPlanFromRow() (+5 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.06
Nodes (38): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os (+30 more)

### Community 27 - "GinServer"
Cohesion: 0.18
Nodes (6): go_pkg_net, go_pkg_path_filepath, github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "go_pkg_fmt"
Cohesion: 0.20
Nodes (4): PortfolioNotFoundError, go_pkg_flag, go_pkg_fmt, UnexpectedRowsAffectedError

### Community 29 - "PortfolioAllocation"
Cohesion: 0.06
Nodes (51): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+43 more)

### Community 40 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.05
Nodes (46): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildDivergenceAnalysisRESTController(), BuildPortfolioRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, contextKey, BuildPortfolioAnalysisConfigurationAppService() (+38 more)

### Community 41 - "rdbms_query.go"
Cohesion: 0.29
Nodes (6): database/sql.Row, database/sql.Rows, BuildILikeSubstringPattern(), processSQL(), ReturningIntIdRowScanner(), ReturningIntIdSingleRowScanner()

### Community 42 - "asset_integration_service_yahoo.go"
Cohesion: 0.17
Nodes (13): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceSearchQuoteDTS (+5 more)

### Community 43 - "Configuration"
Cohesion: 0.21
Nodes (12): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, Configuration, RDBMSConfiguration, YahooFinanceConfiguration, ReadConfig() (+4 more)

### Community 44 - "App"
Cohesion: 0.21
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), App

### Community 45 - "ObservationPageCursor"
Cohesion: 0.28
Nodes (6): AllocationPlanPage, AllocationPlanPageCursor, ObservationPageCursor, PortfolioAllocationHistoryPage, PortfolioObservationPage, buildPortfolioAllocationHistoryPage()

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.22
Nodes (27): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), assertPersistedAssetWithExternalData(), TestPostAllocationPlanRejectsWhitespaceName() (+19 more)

### Community 47 - "custom_deep_validation.go"
Cohesion: 0.33
Nodes (11): go_pkg_github_com_go_playground_validator_v10, reflect.Value, buildValidationErrorPathWithJSON(), DeepValidate(), validator.FieldError, isDirectValidationErrorForStruct(), validateSliceDeep(), validateStructDeep() (+3 more)

### Community 48 - "reflect.Type"
Cohesion: 0.42
Nodes (11): reflect.StructField, reflect.Type, buildJSONFieldPath(), GetJSONFieldName(), mapIndexedNamespacePart(), mapNamespacePart(), parseNamespace(), buildFieldPath() (+3 more)

### Community 49 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.38
Nodes (9): database/sql.Tx, migratedObservation, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations() (+1 more)

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 51 - "hierarchical_id_test.go"
Cohesion: 0.27
Nodes (8): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), database/sql.NullString, StringPointerToNullString(), StringToNullString(), ValueToString()

### Community 52 - "golang_ext_util_struct.go"
Cohesion: 0.29
Nodes (7): deepCompleteReflective(), DeepCompleteStruct(), GetStructName(), GetStructNamespaceDescription(), GetStructType(), T, StructString()

### Community 53 - "custom_validation_error_handling.go"
Cohesion: 0.39
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, asValidationErrors(), formatErrorMessage(), formatValidationError(), FormatValidationErrorMessages(), validator.FieldError, validator.ValidationErrors, MapValidationErrorsToMessages()

### Community 54 - "rdbms_adapter.go"
Cohesion: 0.36
Nodes (7): database/sql.Stmt, BuildQueryInTransaction(), createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), T, prepareBulkInsertValues(), quoteIdentifiers()

### Community 55 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.52
Nodes (6): YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS

### Community 56 - "PortfolioAllocationRDBMSRepository"
Cohesion: 0.48
Nodes (3): BuildPortfolioAllocationRepository(), preparePortfolioAllocationTempInserts(), PortfolioAllocationRDBMSRepository

### Community 57 - "go_pkg_database_sql"
Cohesion: 0.29
Nodes (4): go_pkg_database_sql, NullTime, contextKey, TransactionalContext

### Community 58 - "go_pkg_reflect"
Cohesion: 0.40
Nodes (3): go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, IsNilPointer()

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 102 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `RepositoryRDBMSAdapter`, `context.Context`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Configuration`, `go_pkg_github_com_golang_glog`, `go_pkg_errors`, `golang_ext_util_struct.go`, `infra.go`, `GinServer`, `PortfolioAllocation`?**
  _High betweenness centrality (0.063) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `github.com/gin-gonic/gin.Context`, `AllocationPlanDomService`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `go_pkg_errors`, `PropagateAsAppErrorWithNewMessage`, `BuildCleanupFunctionBuilder`, `allocation_plan_rest_model.go`, `AllocationPlan`, `PortfolioAllocation`?**
  _High betweenness centrality (0.039) - this node is a cross-community bridge._
- **Why does `PortfolioAllocation` connect `PortfolioAllocation` to `github.com/gin-gonic/gin.Context`, `context.Context`, `ObservationPageCursor`, `Asset`, `PortfolioAllocationRDBMSRepository`?**
  _High betweenness centrality (0.027) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.05541368743615935 - nodes in this community are weakly interconnected._
- **Should `portfolio_allocation_pagination_integration_test.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11515151515151516 - nodes in this community are weakly interconnected._