# Graph Report - go  (2026-10-25)

## Corpus Check
- 131 files · ~73,995 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1243 nodes · 3911 edges · 58 communities (46 shown, 12 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 292 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `4aff0ed3`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- AllocationPlan
- T
- error.go
- hierarchical_id_test.go
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
- AllocationStructure
- NewOrderedMapIterator
- Adapter
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- golang_ext_util_concurrent.go
- infra_db.go
- Asset
- IsZeroValue
- assert_json_extension.go
- allocation_plan_rest_model.go
- GinServer
- AllocationPlanIdentifier
- asset_rest_model.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- .buildAppComponents
- portfolio_observation_ownership_migration_integration_test.go
- portfolio_rest_mapping.go
- pagination_cursor.go
- App
- AssetDomService
- BuildCleanupFunctionBuilder
- go_pkg_github_com_benizzio_open_asset_allocator_infra
- PlannedAllocation
- database/sql.Rows
- PortfolioAllocationRESTController
- go_pkg_context
- HasQueryParameter
- rdbms_error.go
- allocation_plan_management_service.go
- AssetRESTController
- CustomValidationErrorsBuilder
- PropagateAsAppErrorWithNewMessage

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
- `AllocationPlanRESTController` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go

## Import Cycles
- None detected.

## Communities (58 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.15
Nodes (19): validateAllocationPlanPlannedDate(), readAllocationPlanPaginationCursor(), readObservationPaginationCursor(), respondWithPaginationCursorValidationError(), bindAndValidatePortfolioSnapshot(), readHistoryObservationTimestampId(), github.com/gin-gonic/gin.Context, HandleAPIError() (+11 more)

### Community 1 - "AllocationPlan"
Cohesion: 0.20
Nodes (12): GetPlanType(), PlanType, AllocationPlan, AllocationPlanOwnership, AllocationPlanPage, AllocationPlanPageCursor, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow() (+4 more)

### Community 2 - "T"
Cohesion: 0.11
Nodes (16): github.com/go-ozzo/ozzo-dbx.Query, QueryBuilder, T, withParams(), processParamsForPostgreSQL(), processSQL(), T, QueryBuilder[T] (+8 more)

### Community 3 - "error.go"
Cohesion: 0.27
Nodes (7): classifyObservationWriteError(), BuildAppErrorFormatted(), BuildUniqueConstraintViolationError(), logError(), newAppError(), IsUniqueConstraintViolation(), UniqueConstraintViolationError

### Community 4 - "hierarchical_id_test.go"
Cohesion: 0.24
Nodes (9): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, database/sql.NullString, StringPointerToNullString(), StringToNullString() (+1 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.06
Nodes (49): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+41 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.11
Nodes (34): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+26 more)

### Community 8 - "testing.T"
Cohesion: 0.12
Nodes (31): go_pkg_encoding_json, testing.T, postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios() (+23 more)

### Community 9 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.12
Nodes (45): go_pkg_encoding_base64, time.Time, NullTime, allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities(), createAllocationPlanPaginationFixture(), encodeTestCursor() (+37 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.16
Nodes (18): getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch(), TestGetKnownAssetsWithTextSearch() (+10 more)

### Community 11 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.15
Nodes (15): BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), go_pkg_database_sql, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain (+7 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.13
Nodes (23): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_io (+15 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.16
Nodes (21): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+13 more)

### Community 15 - "AllocationStructure"
Cohesion: 0.11
Nodes (15): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, database/sql/driver.Value (+7 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "Adapter"
Cohesion: 0.13
Nodes (9): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey (+1 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.10
Nodes (22): BuildPortfolioAllocationRESTController(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, buildSelectedExternalAsset(), mapPortfolioAllocationRow() (+14 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.16
Nodes (27): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource() (+19 more)

### Community 21 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 22 - "infra_db.go"
Cohesion: 0.07
Nodes (27): go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M (+19 more)

### Community 23 - "Asset"
Cohesion: 0.16
Nodes (11): Asset, buildAssetInsertValue(), buildAssetInsertValues(), BuildQuery(), BuildQueryInTransaction(), T, BuildILikeSubstringPattern(), TestBuildILikeSubstringPattern() (+3 more)

### Community 24 - "IsZeroValue"
Cohesion: 0.06
Nodes (41): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+33 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "allocation_plan_rest_model.go"
Cohesion: 0.22
Nodes (16): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+8 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "AllocationPlanIdentifier"
Cohesion: 0.31
Nodes (6): mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), AllocationPlanIdentifier, AnalysisOptions, AllocationPlanIdentifierDTS

### Community 29 - "asset_rest_model.go"
Cohesion: 0.21
Nodes (18): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+10 more)

### Community 40 - ".buildAppComponents"
Cohesion: 0.19
Nodes (9): BuildPortfolioRESTController(), AllocationRepository, BuildYahooFinanceAssetIntegrationClient(), BuildAllocationDomService(), AllocationDomService, TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), ReadConfig() (+1 more)

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 42 - "portfolio_rest_mapping.go"
Cohesion: 0.14
Nodes (27): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), MapToDivergenceAnalysisDTS(), mapToObservationTimestampDTS(), MapToPortfolio(), MapToPortfolioAllocation() (+19 more)

### Community 43 - "pagination_cursor.go"
Cohesion: 0.20
Nodes (14): encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor(), readPaginationCursor(), validatePaginationCursorScope(), go_pkg_bytes, go_pkg_net_url, RawQueryHasParameter() (+6 more)

### Community 44 - "App"
Cohesion: 0.13
Nodes (16): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os, go_pkg_os_signal, go_pkg_syscall, os.Signal, Configuration (+8 more)

### Community 45 - "AssetDomService"
Cohesion: 0.07
Nodes (32): YahooFinanceAssetIntegrationService, BuildAssetRESTController(), ExternalAsset, ExternalAssetQuote, AssetIntegrationService, AssetRepository, AssetExternalSource, BuildYahooFinanceAssetIntegrationService() (+24 more)

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.22
Nodes (30): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), TestPostAsset(), assertPersistedAssetWithExternalData() (+22 more)

### Community 47 - "go_pkg_github_com_benizzio_open_asset_allocator_infra"
Cohesion: 0.17
Nodes (15): BuildAllocationPlanRESTController(), BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service (+7 more)

### Community 48 - "PlannedAllocation"
Cohesion: 0.47
Nodes (3): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), PlannedAllocation

### Community 49 - "database/sql.Rows"
Cohesion: 0.28
Nodes (7): assetRowScanner(), database/sql.Rows, github.com/go-ozzo/ozzo-dbx.Builder, github.com/go-ozzo/ozzo-dbx.Executor, ReturningIntIdRowScanner(), newPaginationCaptureBuilder(), paginationQueryCaptureExecutor

### Community 51 - "go_pkg_context"
Cohesion: 0.10
Nodes (19): go_pkg_context, go_pkg_database_sql_driver, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_golang_glog (+11 more)

### Community 53 - "rdbms_error.go"
Cohesion: 0.25
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 56 - "allocation_plan_management_service.go"
Cohesion: 0.60
Nodes (3): BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, TransactionManager

### Community 57 - "AssetRESTController"
Cohesion: 0.17
Nodes (6): github.com/gin-gonic/gin.HandlersChain, BindAndValidateQueryWithInvalidResponse(), TestBindAndValidateQueryWithInvalidResponse_MalformedQuery(), RESTRoute, DeepValidate(), AssetRESTController

### Community 59 - "CustomValidationErrorsBuilder"
Cohesion: 0.24
Nodes (6): go_pkg_github_com_go_playground_validator_v10, deduplicateValidationMessages(), buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 61 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.19
Nodes (8): decodePaginationCursorPayload(), preparePortfolioAllocationTempInserts(), github.com/go-ozzo/ozzo-dbx.Rows, BuildAppError(), PropagateAsAppErrorWithNewMessage(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository, PortfolioAllocationRDBMSRepository

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 102 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `reflect.Type`, `.buildAppComponents`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `go_pkg_github_com_benizzio_open_asset_allocator_infra`, `Adapter`, `go_pkg_context`, `infra_db.go`, `GinServer`?**
  _High betweenness centrality (0.053) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `go_pkg_context`, `App`, `go_pkg_testing`?**
  _High betweenness centrality (0.041) - this node is a cross-community bridge._
- **Why does `GinServer` connect `GinServer` to `go_pkg_context`, `App`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `T` be split into smaller, more focused modules?**
  _Cohesion score 0.11363636363636363 - nodes in this community are weakly interconnected._
- **Should `reflect.Type` be split into smaller, more focused modules?**
  _Cohesion score 0.061343204653622425 - nodes in this community are weakly interconnected._