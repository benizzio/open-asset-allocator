# Graph Report - go  (2026-11-01)

## Corpus Check
- 131 files · ~73,965 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1233 nodes · 3854 edges · 54 communities (41 shown, 13 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 288 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `92fec34d`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- portfolio_allocation_pagination_integration_test.go
- Adapter
- IsZeroValue
- .buildAppComponents
- reflect.Type
- context.Context
- deferCloseResponseBody
- testing.T
- go_pkg_fmt
- PortfolioAllocation
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- Asset
- NewOrderedMapIterator
- PropagateAsAppErrorWithNewMessage
- T
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- T
- AllocationPlan
- golang_ext_util_concurrent.go
- portfolio_observation_ownership_migration_integration_test.go
- assert_json_extension.go
- infra.go
- go_pkg_database_sql
- database/sql.NullString
- portfolio_rest_mapping.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- go_pkg_context
- rdbms_query.go
- AllocationDomService
- insertTestPortfolio
- App
- BuildCleanupFunctionBuilder
- CustomValidationErrorsBuilder
- Portfolio
- PortfolioAllocationRESTController
- T
- rdbms_adapter.go
- go_pkg_github_com_benizzio_open_asset_allocator_infra_util_urlutil
- PortfolioAllocationRDBMSRepository
- pagination_cursor.go

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
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `allocationPlanCursorFromPayload()` --references--> `AllocationPlanPageCursor`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/allocation_plan.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go

## Import Cycles
- None detected.

## Communities (54 total, 13 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.12
Nodes (25): validateAllocationPlanPlannedDate(), MapToPortfolio(), respondWithPaginationCursorValidationError(), bindAndValidatePortfolioSnapshot(), cleanNilAllocations(), readHistoryObservationTimestampId(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context (+17 more)

### Community 1 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.11
Nodes (44): go_pkg_net_url, net/url.Values, time.Time, HasQueryParameter(), RawQueryHasParameter(), allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities() (+36 more)

### Community 2 - "Adapter"
Cohesion: 0.16
Nodes (6): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, Adapter, SQLTransactionalContext, withTransaction()

### Community 3 - "IsZeroValue"
Cohesion: 0.06
Nodes (42): AllocationHierarchy, AllocationPlanRepository, AssetExternalSource, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+34 more)

### Community 4 - ".buildAppComponents"
Cohesion: 0.15
Nodes (13): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildPortfolioDivergenceAnalysisAppService(), PortfolioDivergenceAnalysisAppService, BuildAllocationPlanRepository(), BuildAllocationRepository(), go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation (+5 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.12
Nodes (29): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+21 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.11
Nodes (33): assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds(), TestPostAllocationPlanUpdatesOwnedMetadata() (+25 more)

### Community 8 - "testing.T"
Cohesion: 0.09
Nodes (42): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), net/http.Response, testing.T, TestBuildILikeSubstringPattern() (+34 more)

### Community 9 - "go_pkg_fmt"
Cohesion: 0.05
Nodes (27): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationStructure, HierarchicalId, PortfolioNotFoundError, go_pkg_flag (+19 more)

### Community 10 - "PortfolioAllocation"
Cohesion: 0.12
Nodes (16): BuildPortfolioAllocationRESTController(), BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AllocationHierarchyLevel, AssetsPerTicker, buildSelectedExternalAsset() (+8 more)

### Community 11 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.11
Nodes (25): allocationPlanCursorFromPayload(), BuildAllocationPlanRESTController(), BuildAssetRESTController(), observationCursorFromPayload(), validateAssetIdentification(), BuildPortfolioRESTController(), assetRowScanner(), buildAssetInsertValue() (+17 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.17
Nodes (19): go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_io, go_pkg_net_http (+11 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (22): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+14 more)

### Community 15 - "Asset"
Cohesion: 0.23
Nodes (19): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+11 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.16
Nodes (13): BuildAssetRDBMSRepository(), classifyObservationWriteError(), BuildAppError(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), RepositoryRDBMSAdapter, IsUniqueConstraintViolation() (+5 more)

### Community 18 - "T"
Cohesion: 0.31
Nodes (5): QueryBuilder, T, QueryBuilder[T], QueryExecutor[T], RowScanner

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.18
Nodes (25): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange(), TestGetExternalAssetQuoteRejectsMismatchedProviderTicker() (+17 more)

### Community 21 - "T"
Cohesion: 0.27
Nodes (7): processParamsForPostgreSQL(), processSQL(), T, IsSlice(), SQLTransactionalQueryBuilder, SQLTransactionalQueryBuilder[T], SQLTransactionalQueryExecutor

### Community 22 - "AllocationPlan"
Cohesion: 0.08
Nodes (37): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+29 more)

### Community 23 - "golang_ext_util_concurrent.go"
Cohesion: 0.19
Nodes (24): go_pkg_runtime, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount(), closeConcurrentResultChannels() (+16 more)

### Community 24 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (9): database/sql.Tx, migratedObservation, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations() (+1 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.06
Nodes (38): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os (+30 more)

### Community 27 - "go_pkg_database_sql"
Cohesion: 0.32
Nodes (6): go_pkg_database_sql, github.com/go-ozzo/ozzo-dbx.Query, withParams(), contextKey, QueryExecutor, TransactionalContext

### Community 28 - "database/sql.NullString"
Cohesion: 0.38
Nodes (6): TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), database/sql.NullString, StringPointerToNullString(), StringToNullString(), ValueToString()

### Community 29 - "portfolio_rest_mapping.go"
Cohesion: 0.12
Nodes (29): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+21 more)

### Community 40 - "go_pkg_context"
Cohesion: 0.07
Nodes (35): YahooFinanceAssetIntegrationService, ExternalAsset, ExternalAssetQuote, AssetIntegrationService, AssetRepository, BuildYahooFinanceAssetIntegrationService(), extractLastClose(), mapToExternalAsset() (+27 more)

### Community 41 - "rdbms_query.go"
Cohesion: 0.20
Nodes (6): go_pkg_github_com_lib_pq, database/sql.Row, BuildILikeSubstringPattern(), ReturningIntIdSingleRowScanner(), SingleRowScanner, SQLTransactionalQueryExecutor[T]

### Community 42 - "AllocationDomService"
Cohesion: 0.47
Nodes (3): AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 43 - "insertTestPortfolio"
Cohesion: 0.23
Nodes (13): TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans(), putForValidationFailure(), TestPostPortfolio(), TestPostPortfolioWithAllocationStructure(), TestPutPortfolio(), TestPutPortfolioFailureWithNameExceedingMaxLength(), TestPutPortfolioFailureWithoutMandatoryFields(), TestPutPortfolioWithAllocationStructure() (+5 more)

### Community 44 - "App"
Cohesion: 0.06
Nodes (29): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_net (+21 more)

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.23
Nodes (28): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), assertPersistedAssetWithExternalData(), TestPostPortfolioAllocationHistoryAllowsSameTagAcrossPortfolios() (+20 more)

### Community 47 - "CustomValidationErrorsBuilder"
Cohesion: 0.36
Nodes (4): buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 48 - "Portfolio"
Cohesion: 0.22
Nodes (6): BuildPortfolioRepository(), Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioRDBMSRepository

### Community 54 - "rdbms_adapter.go"
Cohesion: 0.23
Nodes (9): go_pkg_github_com_golang_glog, database/sql.Stmt, buildPingContext(), BuildQueryInTransaction(), createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), T, prepareBulkInsertValues() (+1 more)

### Community 57 - "PortfolioAllocationRDBMSRepository"
Cohesion: 0.21
Nodes (10): addObservationCursorFilter(), buildPortfolioAllocationHistoryPage(), BuildPortfolioAllocationRepository(), buildPortfolioObservationPage(), T, preparePortfolioAllocationTempInserts(), ObservationPageCursor, PortfolioAllocationHistoryPage (+2 more)

### Community 59 - "pagination_cursor.go"
Cohesion: 0.29
Nodes (13): decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor(), readPaginationCursor(), validatePaginationCursorScope() (+5 more)

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 104 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **13 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `.buildAppComponents`, `reflect.Type`, `go_pkg_context`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `rdbms_adapter.go`, `AllocationPlan`, `infra.go`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `IsZeroValue`, `go_pkg_context`, `PortfolioAllocation`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `PropagateAsAppErrorWithNewMessage`, `ExecuteDBQuery`, `AllocationPlan`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `App`, `rdbms_adapter.go`?**
  _High betweenness centrality (0.031) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.12367864693446089 - nodes in this community are weakly interconnected._
- **Should `portfolio_allocation_pagination_integration_test.go` be split into smaller, more focused modules?**
  _Cohesion score 0.11100832562442182 - nodes in this community are weakly interconnected._