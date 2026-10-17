# Graph Report - go  (2026-10-17)

## Corpus Check
- 119 files · ~59,697 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1074 nodes · 3201 edges · 57 communities (45 shown, 12 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 236 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `04895648`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- Portfolio
- .buildAppComponents
- allocation_plan_domain_service.go
- PropagateAsAppErrorWithNewMessage
- reflect.Type
- context.Context
- deferCloseResponseBody
- GinServer
- portfolio_rest_mapping.go
- asset_integration_test.go
- infra_db.go
- golang_ext_tree_test.go
- go_pkg_testing
- BuildCleanupFunctionBuilder
- go_pkg_fmt
- NewOrderedMapIterator
- testing.T
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- Asset
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- Adapter
- golang_ext_util_concurrent.go
- assert_json_extension.go
- ExternalAsset
- AllocationDomService
- AllocationPlan
- AllocationStructure
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- portfolio_allocation_controller.go
- allocation_plan_rest_model.go
- asset_rest_model.go
- Configuration
- go_pkg_database_sql
- App
- rdbms_adapter.go
- rdbms_query.go
- PortfolioAllocationRDBMSRepository
- go_pkg_github_com_golang_glog
- portfolio_allocation_integration_test.go
- T
- hierarchical_id_test.go
- manage_db.go
- QueryExecutor[T]
- AllocationPlanIdentifier
- .InsertAsset

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 78 edges
2. `App` - 35 edges
3. `BuildCleanupFunctionBuilder()` - 33 edges
4. `PortfolioAllocation` - 32 edges
5. `PropagateAsAppErrorWithNewMessage()` - 31 edges
6. `Asset` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 26 edges
9. `PlannedAllocation` - 23 edges
10. `AllocationPlan` - 23 edges

## Surprising Connections (you probably didn't know these)
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `BuildAllocationPlanRESTController()` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `AssetRESTController` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `DivergenceAnalysisRESTController` --references--> `PortfolioDivergenceAnalysisAppService`  [EXTRACTED]
  api/rest/divergence_analysis_controller.go → application/portfolio_divergence_analysis_service.go

## Import Cycles
- None detected.

## Communities (57 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (31): validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.HandlersChain, HandleAPIError() (+23 more)

### Community 1 - "Portfolio"
Cohesion: 0.27
Nodes (5): Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioRDBMSRepository

### Community 2 - ".buildAppComponents"
Cohesion: 0.18
Nodes (14): BuildAllocationPlanRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository() (+6 more)

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.07
Nodes (38): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+30 more)

### Community 4 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.23
Nodes (8): BuildAppError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), ToSQLTransactionalContext(), ToPointerSlice(), AllocationPlanRDBMSRepository, AssetRDBMSRepository, S

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.07
Nodes (44): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+36 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.14
Nodes (25): TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches(), TestPostAllocationPlanValidation_MissingDetails() (+17 more)

### Community 8 - "GinServer"
Cohesion: 0.18
Nodes (6): go_pkg_net, go_pkg_path_filepath, github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 9 - "portfolio_rest_mapping.go"
Cohesion: 0.18
Nodes (23): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS(), mapToExternalAsset(), mapToExternalAssetData() (+15 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.15
Nodes (20): go_pkg_net_url, net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets() (+12 more)

### Community 11 - "infra_db.go"
Cohesion: 0.07
Nodes (29): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os, go_pkg_sync, github.com/nhatthm/httpmock.Server (+21 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.09
Nodes (26): TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx (+18 more)

### Community 14 - "BuildCleanupFunctionBuilder"
Cohesion: 0.23
Nodes (28): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), TestPostAsset(), assertPersistedAssetWithExternalData() (+20 more)

### Community 15 - "go_pkg_fmt"
Cohesion: 0.13
Nodes (19): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceChartResponseDTS (+11 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "testing.T"
Cohesion: 0.14
Nodes (29): testing.T, TestPostAssetWithoutExternalData(), postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios() (+21 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.18
Nodes (11): MapToPortfolioObservationTimestamp(), BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, PortfolioAllocation, PortfolioAllocationRepository (+3 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.27
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 21 - "Asset"
Cohesion: 0.16
Nodes (10): BuildAssetRESTController(), Asset, AssetIntegrationService, AssetRepository, buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetDomService(), collectIntegrationServices() (+2 more)

### Community 22 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.39
Nodes (6): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 23 - "Adapter"
Cohesion: 0.14
Nodes (8): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, Adapter, SQLTransactionalContext, withTransaction(), contextKey, TransactionalContext

### Community 24 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "ExternalAsset"
Cohesion: 0.22
Nodes (12): MapToExternalAssetQuoteDTS(), ExternalAsset, ExternalAssetQuote, AssetExternalSource, extractLastClose(), mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), go_pkg_golang_org_x_text_currency (+4 more)

### Community 27 - "AllocationDomService"
Cohesion: 0.38
Nodes (4): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 28 - "AllocationPlan"
Cohesion: 0.14
Nodes (13): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), GetPlanType(), PlanType, AllocationPlan, PlannedAllocation, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow() (+5 more)

### Community 29 - "AllocationStructure"
Cohesion: 0.09
Nodes (18): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, go_pkg_database_sql_driver (+10 more)

### Community 40 - "portfolio_allocation_controller.go"
Cohesion: 0.25
Nodes (11): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin (+3 more)

### Community 41 - "allocation_plan_rest_model.go"
Cohesion: 0.29
Nodes (15): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+7 more)

### Community 42 - "asset_rest_model.go"
Cohesion: 0.30
Nodes (14): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+6 more)

### Community 43 - "Configuration"
Cohesion: 0.19
Nodes (12): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), Configuration, RDBMSConfiguration, YahooFinanceConfiguration, ReadConfig(), TestReadConfigAPIOnly() (+4 more)

### Community 44 - "go_pkg_database_sql"
Cohesion: 0.22
Nodes (8): go_pkg_database_sql, processParamsForPostgreSQL(), T, NullTime, IsSlice(), SQLTransactionalQueryBuilder, SQLTransactionalQueryBuilder[T], SQLTransactionalQueryExecutor

### Community 45 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 46 - "rdbms_adapter.go"
Cohesion: 0.21
Nodes (9): database/sql.Stmt, database/sql.Tx, buildPingContext(), BuildQueryInTransaction(), createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), T, prepareBulkInsertValues() (+1 more)

### Community 47 - "rdbms_query.go"
Cohesion: 0.17
Nodes (9): assetRowScanner(), database/sql.Row, database/sql.Rows, BuildILikeSubstringPattern(), processSQL(), ReturningIntIdRowScanner(), ReturningIntIdSingleRowScanner(), SingleRowScanner (+1 more)

### Community 48 - "PortfolioAllocationRDBMSRepository"
Cohesion: 0.29
Nodes (6): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), preparePortfolioAllocationTempInserts(), portfolioAllocationJoinedRowDTS, PortfolioAllocationRDBMSRepository

### Community 49 - "go_pkg_github_com_golang_glog"
Cohesion: 0.29
Nodes (5): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, ConfigLogger(), main()

### Community 50 - "portfolio_allocation_integration_test.go"
Cohesion: 0.20
Nodes (11): postPortfolioAllocationForValidationFailure(), TestGetAvailableHistoryObservations(), TestGetPortfolioAllocationHistory(), TestGetPortfolioAllocationHistoryForObservationTimestamp(), TestGetPortfolioAllocationHistoryForObservationTimestampNoneFound(), TestGetPortfolioAllocationHistoryWithMultiplePortfoliosAndManyObservations(), TestPostPortfolioAllocationHistoryValidation_AssetNameExceedsMaxLength(), TestPostPortfolioAllocationHistoryValidation_AssetTickerExceedsMaxLength() (+3 more)

### Community 51 - "T"
Cohesion: 0.38
Nodes (6): github.com/go-ozzo/ozzo-dbx.Query, T, withParams(), QueryBuilder, QueryBuilder[T], QueryExecutor

### Community 52 - "hierarchical_id_test.go"
Cohesion: 0.27
Nodes (8): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), database/sql.NullString, StringPointerToNullString(), StringToNullString(), ValueToString()

### Community 54 - "manage_db.go"
Cohesion: 0.48
Nodes (4): github.com/go-ozzo/ozzo-dbx.Params, createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 56 - "AllocationPlanIdentifier"
Cohesion: 0.47
Nodes (5): mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), AllocationPlanIdentifier, AnalysisOptions, AllocationPlanIdentifierDTS

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `.buildAppComponents`, `reflect.Type`, `portfolio_allocation_controller.go`, `GinServer`, `infra_db.go`, `Configuration`, `go_pkg_fmt`, `go_pkg_github_com_golang_glog`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Adapter`, `AllocationPlan`?**
  _High betweenness centrality (0.055) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `allocation_plan_domain_service.go`, `PropagateAsAppErrorWithNewMessage`, `allocation_plan_rest_model.go`, `asset_rest_model.go`, `rdbms_query.go`, `PortfolioAllocationRDBMSRepository`, `PortfolioAllocation`, `ExecuteDBQuery`, `.InsertAsset`, `AllocationPlan`?**
  _High betweenness centrality (0.045) - this node is a cross-community bridge._
- **Why does `PlannedAllocation` connect `AllocationPlan` to `allocation_plan_domain_service.go`, `PropagateAsAppErrorWithNewMessage`, `context.Context`, `allocation_plan_rest_model.go`, `Asset`, `ExternalAsset`, `AllocationStructure`?**
  _High betweenness centrality (0.036) - this node is a cross-community bridge._
- **Are the 76 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 76 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.07834101382488479 - nodes in this community are weakly interconnected._
- **Should `allocation_plan_domain_service.go` be split into smaller, more focused modules?**
  _Cohesion score 0.06775956284153005 - nodes in this community are weakly interconnected._