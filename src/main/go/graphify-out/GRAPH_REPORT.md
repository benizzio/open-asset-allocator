# Graph Report - go  (2026-10-22)

## Corpus Check
- 123 files · ~65,726 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1128 nodes · 3457 edges · 52 communities (42 shown, 10 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 261 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `30675898`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- AllocationPlanDomService
- context.Context
- allocation_plan_domain_service.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- reflect.Type
- portfolio_divergence_analysis_context.go
- deferCloseResponseBody
- testing.T
- go_pkg_strings
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- asset_integration_client_yahoo_model.go
- NewOrderedMapIterator
- PropagateAsAppErrorWithNewMessage
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- BuildAppError
- golang_ext_util_concurrent.go
- Adapter
- AllocationPlan
- assert_json_extension.go
- logFanOutConsumer
- GinServer
- portfolio_rest_mapping.go
- asset_rest_model.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- go_pkg_github_com_stretchr_testify_require
- portfolio_observation_ownership_migration_integration_test.go
- go_pkg_time
- PlannedAllocation
- App
- go_pkg_encoding_json
- BuildCleanupFunctionBuilder
- PortfolioAllocationRDBMSRepository
- yahoo_finance_mock.go
- go_pkg_net_http
- http_client.go
- ForeignKeyConstraintViolationError

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 89 edges
2. `BuildCleanupFunctionBuilder()` - 40 edges
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
- `AssetRESTController` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `MapToAllocationPlanDTSs()` --references--> `AllocationPlan`  [EXTRACTED]
  api/rest/model/allocation_plan_rest_model.go → domain/allocation_plan.go
- `mapToAllocationPlanDTS()` --references--> `AllocationPlan`  [EXTRACTED]
  api/rest/model/allocation_plan_rest_model.go → domain/allocation_plan.go
- `mapToPlannedAllocationDTSs()` --references--> `AllocationPlan`  [EXTRACTED]
  api/rest/model/allocation_plan_rest_model.go → domain/allocation_plan.go

## Import Cycles
- None detected.

## Communities (52 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (35): validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), bindAndValidatePortfolioSnapshot(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.HandlersChain (+27 more)

### Community 1 - "AllocationPlanDomService"
Cohesion: 0.05
Nodes (41): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildDivergenceAnalysisRESTController(), BuildPortfolioRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService (+33 more)

### Community 2 - "context.Context"
Cohesion: 0.19
Nodes (8): context.Context, database/sql.DB, SQLTransactionalContext, ToSQLTransactionalContext(), withTransaction(), contextKey, TransactionalContext, AllocationPlanRDBMSRepository

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.08
Nodes (33): cleanNilAllocations(), AllocationHierarchy, appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData(), readValidationData() (+25 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.21
Nodes (13): assetRowScanner(), go_pkg_context, go_pkg_database_sql, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext (+5 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (41): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+33 more)

### Community 6 - "portfolio_divergence_analysis_context.go"
Cohesion: 0.08
Nodes (30): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+22 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.16
Nodes (26): assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds(), TestPostAllocationPlanUpdatesOwnedMetadata() (+18 more)

### Community 8 - "testing.T"
Cohesion: 0.15
Nodes (27): testing.T, postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios(), TestPostPortfolio() (+19 more)

### Community 9 - "go_pkg_strings"
Cohesion: 0.15
Nodes (12): go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_golang_glog, go_pkg_net, go_pkg_path_filepath, go_pkg_strings, database/sql.Stmt (+4 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.12
Nodes (29): go_pkg_net_url, net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets() (+21 more)

### Community 11 - "infra.go"
Cohesion: 0.10
Nodes (22): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, testing.M, contextKey (+14 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.12
Nodes (17): TestParseAssetTextSearch(), go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_stretchr_testify_assert, go_pkg_io, go_pkg_testing, TestBuildILikeSubstringPattern(), TestGetExternalAssetsFailure() (+9 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (21): TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRejectsWhitespaceTags(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest(), registerObservationFixtureRestore() (+13 more)

### Community 15 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.23
Nodes (10): buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, PropagateAsAppError(), YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS (+2 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.22
Nodes (11): Asset, buildAssetInsertValue(), buildAssetInsertValues(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, BuildILikeSubstringPattern() (+3 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.16
Nodes (9): MapToPortfolioObservationTimestamp(), BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, preparePortfolioAllocationTempInserts(), PortfolioAllocation, PortfolioAllocationRepository, PortfolioObservationTimestamp, BuildPortfolioAllocationDomService() (+1 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.21
Nodes (21): go_pkg_github_com_go_ozzo_ozzo_dbx, github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange(), TestGetExternalAssetQuoteRejectsMismatchedProviderTicker() (+13 more)

### Community 21 - "BuildAppError"
Cohesion: 0.18
Nodes (13): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, time.Time, BuildAppError(), BuildAppErrorFormatted() (+5 more)

### Community 22 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 23 - "Adapter"
Cohesion: 0.07
Nodes (21): database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, buildPingContext(), Adapter, T, withParams() (+13 more)

### Community 24 - "AllocationPlan"
Cohesion: 0.20
Nodes (10): GetPlanType(), PlanType, AllocationPlan, AllocationPlanOwnership, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow(), mapPlannedAllocationRow(), mapPlannedAllocationRows() (+2 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "logFanOutConsumer"
Cohesion: 0.15
Nodes (9): go_pkg_sync, sync.Mutex, testing.TB, logFanOutConsumer, BuildDeferRegistry(), AttachDatabaseLogsTo(), newLogFanOutConsumer(), testcontainers.Log (+1 more)

### Community 27 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 28 - "portfolio_rest_mapping.go"
Cohesion: 0.11
Nodes (41): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+33 more)

### Community 29 - "asset_rest_model.go"
Cohesion: 0.06
Nodes (41): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets() (+33 more)

### Community 40 - "go_pkg_github_com_stretchr_testify_require"
Cohesion: 0.24
Nodes (9): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_github_com_stretchr_testify_require, database/sql.NullString, StringPointerToNullString(), StringToNullString() (+1 more)

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.44
Nodes (8): database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations(), TestPortfolioObservationOwnershipMigration()

### Community 42 - "go_pkg_time"
Cohesion: 0.33
Nodes (5): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), YahooFinanceAssetIntegrationClient, go_pkg_github_com_shopspring_decimal, go_pkg_time

### Community 43 - "PlannedAllocation"
Cohesion: 0.33
Nodes (6): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), PlannedAllocation, AssetsPerTicker

### Community 44 - "App"
Cohesion: 0.10
Nodes (23): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os (+15 more)

### Community 45 - "go_pkg_encoding_json"
Cohesion: 0.33
Nodes (4): go_pkg_encoding_json, go_pkg_strconv, TestParseableInt64_UnmarshalJSON(), TestParseInt64()

### Community 46 - "BuildCleanupFunctionBuilder"
Cohesion: 0.23
Nodes (28): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), TestPostAllocationPlanRejectsWhitespaceName(), TestPostPortfolioAllocationHistoryAllowsSameTagAcrossPortfolios() (+20 more)

### Community 47 - "PortfolioAllocationRDBMSRepository"
Cohesion: 0.24
Nodes (7): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), classifyObservationWriteError(), IsUniqueConstraintViolation(), portfolioAllocationJoinedRowDTS, PortfolioAllocationRDBMSRepository

### Community 48 - "yahoo_finance_mock.go"
Cohesion: 0.43
Nodes (6): go_pkg_github_com_nhatthm_httpmock, github.com/nhatthm/httpmock.Server, BuildAndStartYahooFinanceMockServer(), GetYahooFinanceMockServer(), resetYahooFinanceMockServer(), SetYahooFinanceMockServer()

### Community 49 - "go_pkg_net_http"
Cohesion: 0.26
Nodes (10): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin (+2 more)

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 94 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `AllocationPlanDomService`, `context.Context`, `allocation_plan_domain_service.go`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `PlannedAllocation`, `PortfolioAllocationRDBMSRepository`, `PortfolioAllocation`, `ExecuteDBQuery`, `BuildAppError`, `AllocationPlan`, `portfolio_rest_mapping.go`, `asset_rest_model.go`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **Why does `App` connect `App` to `AllocationPlanDomService`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `go_pkg_strings`, `go_pkg_time`, `infra.go`, `go_pkg_net_http`, `Adapter`, `GinServer`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `asset_rest_model.go` to `PlannedAllocation`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 87 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 87 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.0710085933966531 - nodes in this community are weakly interconnected._
- **Should `AllocationPlanDomService` be split into smaller, more focused modules?**
  _Cohesion score 0.054274084124830396 - nodes in this community are weakly interconnected._