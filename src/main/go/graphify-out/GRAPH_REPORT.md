# Graph Report - go  (2026-10-17)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 1074 nodes · 3197 edges · 48 communities (38 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 236 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `f7071086`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- PortfolioAllocation
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- allocation_plan_domain_service.go
- Portfolio
- reflect.Type
- portfolio_divergence_analysis_service.go
- deferCloseResponseBody
- GinServer
- portfolio_rest_mapping.go
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- BuildCleanupFunctionBuilder
- Asset
- NewOrderedMapIterator
- testing.T
- Adapter
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- .buildAppComponents
- go_pkg_fmt
- rdbms_query.go
- golang_ext_util_concurrent.go
- assert_json_extension.go
- BuildAppError
- rdbms_adapter.go
- context.Context
- AllocationStructure
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- go_pkg_net_http
- App
- go_pkg_github_com_go_ozzo_ozzo_dbx
- Configuration
- asset_integration_client_yahoo_model.go
- go_pkg_time
- http_client.go
- asset_integration_client_yahoo.go

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 78 edges
2. `App` - 35 edges
3. `BuildCleanupFunctionBuilder()` - 33 edges
4. `PortfolioAllocation` - 32 edges
5. `PropagateAsAppErrorWithNewMessage()` - 31 edges
6. `Asset` - 30 edges
7. `NewMapTree()` - 27 edges
8. `IsZeroValue()` - 26 edges
9. `AllocationPlan` - 23 edges
10. `PlannedAllocation` - 23 edges

## Surprising Connections (you probably didn't know these)
- `ValueJsonColumn()` --calls--> `IsNilPointer()`  [EXTRACTED]
  infra/rdbms/sqlext/sqlext_utils.go → langext/golang_ext_util.go
- `MapToPortfolio()` --references--> `Portfolio`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio.go
- `MapToPortfolioDTS()` --references--> `Portfolio`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio.go
- `MapToPortfolioDTS()` --calls--> `ParseableInt64`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → langext/golang_ext_parseable_int64.go
- `MapToPortfolioDTSs()` --references--> `Portfolio`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio.go

## Import Cycles
- None detected.

## Communities (48 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (31): MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.HandlersChain, HandleAPIError(), handleDomainError() (+23 more)

### Community 1 - "PortfolioAllocation"
Cohesion: 0.16
Nodes (11): mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), preparePortfolioAllocationTempInserts(), PortfolioAllocation, PortfolioObservationTimestamp (+3 more)

### Community 2 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.43
Nodes (5): go_pkg_context, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.08
Nodes (32): cleanNilAllocations(), AllocationHierarchy, appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData(), readValidationData() (+24 more)

### Community 4 - "Portfolio"
Cohesion: 0.12
Nodes (13): BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), Portfolio, BuildQuery(), RepositoryRDBMSAdapter (+5 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "portfolio_divergence_analysis_service.go"
Cohesion: 0.08
Nodes (34): BuildDivergenceAnalysisRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext() (+26 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.09
Nodes (37): go_pkg_io, TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches() (+29 more)

### Community 8 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 9 - "portfolio_rest_mapping.go"
Cohesion: 0.07
Nodes (57): MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS(), mapToPlannedAllocationDTSs() (+49 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.15
Nodes (20): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+12 more)

### Community 11 - "infra.go"
Cohesion: 0.06
Nodes (35): go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_os, go_pkg_path_filepath, go_pkg_sync (+27 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.11
Nodes (21): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), TestParseAssetTextSearch(), go_pkg_database_sql, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra (+13 more)

### Community 14 - "BuildCleanupFunctionBuilder"
Cohesion: 0.23
Nodes (28): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), TestPostAsset(), assertPersistedAssetWithExternalData() (+20 more)

### Community 15 - "Asset"
Cohesion: 0.15
Nodes (27): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+19 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "testing.T"
Cohesion: 0.12
Nodes (32): go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, testing.T, TestAllocationPlanDateMigrationBackfillsLegacyRows(), postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio() (+24 more)

### Community 18 - "Adapter"
Cohesion: 0.14
Nodes (8): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, Adapter, SQLTransactionalContext, withTransaction(), contextKey, TransactionalContext

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.27
Nodes (20): TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource(), TestGetExternalAssetQuoteRejectsMismatchedProviderExchange() (+12 more)

### Community 21 - ".buildAppComponents"
Cohesion: 0.08
Nodes (29): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildPortfolioAllocationRESTController(), BuildPortfolioRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService(), BuildPortfolioDivergenceAnalysisAppService() (+21 more)

### Community 22 - "go_pkg_fmt"
Cohesion: 0.15
Nodes (10): go_pkg_encoding_json, go_pkg_errors, go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_net, go_pkg_strconv (+2 more)

### Community 23 - "rdbms_query.go"
Cohesion: 0.08
Nodes (24): assetRowScanner(), database/sql.Row, database/sql.Rows, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, BuildILikeSubstringPattern(), T (+16 more)

### Community 24 - "golang_ext_util_concurrent.go"
Cohesion: 0.18
Nodes (25): go_pkg_runtime, go_pkg_slices, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "BuildAppError"
Cohesion: 0.19
Nodes (11): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), golang.org/x/text/currency.Unit, BuildAppError(), BuildAppErrorFormatted(), BuildUniqueConstraintViolationError(), logError() (+3 more)

### Community 27 - "rdbms_adapter.go"
Cohesion: 0.25
Nodes (8): database/sql.Stmt, buildPingContext(), BuildQueryInTransaction(), createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), T, prepareBulkInsertValues(), quoteIdentifiers()

### Community 28 - "context.Context"
Cohesion: 0.31
Nodes (4): context.Context, PropagateAsAppErrorWithNewMessage(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 29 - "AllocationStructure"
Cohesion: 0.09
Nodes (18): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, go_pkg_database_sql_driver (+10 more)

### Community 40 - "go_pkg_net_http"
Cohesion: 0.30
Nodes (9): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_net_http (+1 more)

### Community 41 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 42 - "go_pkg_github_com_go_ozzo_ozzo_dbx"
Cohesion: 0.39
Nodes (5): go_pkg_github_com_go_ozzo_ozzo_dbx, github.com/go-ozzo/ozzo-dbx.Params, createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 43 - "Configuration"
Cohesion: 0.19
Nodes (12): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), Configuration, RDBMSConfiguration, YahooFinanceConfiguration, ReadConfig(), TestReadConfigAPIOnly() (+4 more)

### Community 44 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.26
Nodes (9): buildSearchAssetsURL(), YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, PropagateAsAppError(), YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS (+1 more)

### Community 45 - "go_pkg_time"
Cohesion: 0.19
Nodes (10): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), YahooFinanceAssetIntegrationClient, PortfolioAllocationRepository, BuildPortfolioAllocationDomService(), go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_shopspring_decimal (+2 more)

### Community 46 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 47 - "asset_integration_client_yahoo.go"
Cohesion: 0.50
Nodes (3): buildQuoteAssetLastClosePriceURL(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_net_url

## Knowledge Gaps
- **15 isolated node(s):** `contextKey`, `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `MapIterator`, `contextKey` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `go_pkg_net_http`, `GinServer`, `infra.go`, `Configuration`, `go_pkg_time`, `Adapter`, `.buildAppComponents`, `go_pkg_fmt`?**
  _High betweenness centrality (0.066) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `PortfolioAllocation`, `allocation_plan_domain_service.go`, `Portfolio`, `portfolio_rest_mapping.go`, `ExecuteDBQuery`, `.buildAppComponents`, `rdbms_query.go`, `BuildAppError`, `context.Context`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `Configuration`, `rdbms_adapter.go`, `App`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Are the 76 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 76 INFERRED edges - model-reasoned connections that need verification._
- **What connects `contextKey`, `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.07667900581702802 - nodes in this community are weakly interconnected._
- **Should `allocation_plan_domain_service.go` be split into smaller, more focused modules?**
  _Cohesion score 0.0783673469387755 - nodes in this community are weakly interconnected._