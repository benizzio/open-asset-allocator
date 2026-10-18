# Graph Report - go  (2026-10-18)

## Corpus Check
- 119 files · ~59,731 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1075 nodes · 3205 edges · 47 communities (37 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 236 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `8aca179e`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- ExternalAsset
- http_client.go
- allocation_plan_domain_service.go
- asset_integration_client_yahoo_model.go
- reflect.Type
- context.Context
- deferCloseResponseBody
- GinServer
- portfolio_rest_mapping.go
- asset_integration_test.go
- logFanOutConsumer
- golang_ext_tree_test.go
- go_pkg_testing
- assert_db.go
- YahooFinanceAssetIntegrationClient
- NewOrderedMapIterator
- testing.T
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- AssetDomService
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- Adapter
- golang_ext_util_concurrent.go
- assert_json_extension.go
- BuildAppError
- PropagateAsAppErrorWithNewMessage
- AllocationStructure
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- portfolio_allocation_controller.go
- asset_rest_model.go
- Configuration
- App
- rdbms_adapter.go
- go_pkg_strings
- portfolio_allocation_integration_test.go
- database/sql.NullString

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

## Communities (47 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.06
Nodes (43): BuildAllocationPlanRESTController(), validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), bindAndValidatePortfolioSnapshot(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService (+35 more)

### Community 1 - "ExternalAsset"
Cohesion: 0.39
Nodes (8): MapToExternalAssetDTS(), MapToExternalAssetDTSs(), mapToExternalAsset(), mapToExternalAssetData(), ExternalAsset, AssetExternalSource, go_pkg_golang_org_x_text_currency, ExternalAssetDTS

### Community 2 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.08
Nodes (34): cleanNilAllocations(), AllocationHierarchy, AllocationPlanRepository, appendLevelDescription(), BuildAllocationPlanDomService(), AllocationPlanDomService, readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData() (+26 more)

### Community 4 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.52
Nodes (6): YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (41): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+33 more)

### Community 6 - "context.Context"
Cohesion: 0.08
Nodes (40): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+32 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.14
Nodes (25): TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches(), TestPostAllocationPlanValidation_MissingDetails() (+17 more)

### Community 8 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 9 - "portfolio_rest_mapping.go"
Cohesion: 0.15
Nodes (28): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+20 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.15
Nodes (21): go_pkg_net_url, net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets() (+13 more)

### Community 11 - "logFanOutConsumer"
Cohesion: 0.10
Nodes (18): go_pkg_github_com_nhatthm_httpmock, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry() (+10 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.09
Nodes (25): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_infra_util (+17 more)

### Community 14 - "assert_db.go"
Cohesion: 0.23
Nodes (26): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), assertPersistedAssetWithExternalData(), TestPostPortfolioAllocationHistoryFullMerge() (+18 more)

### Community 15 - "YahooFinanceAssetIntegrationClient"
Cohesion: 0.21
Nodes (10): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, YahooFinanceSearchQuoteDTS (+2 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "testing.T"
Cohesion: 0.14
Nodes (28): testing.T, TestPostAsset(), postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios() (+20 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.07
Nodes (33): BuildDivergenceAnalysisRESTController(), BuildPortfolioAllocationRESTController(), BuildPortfolioRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations() (+25 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.19
Nodes (26): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestPostAssetRejectsNonZeroId(), TestPostAssetWithDuplicateTicker(), TestPutAsset(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource() (+18 more)

### Community 21 - "AssetDomService"
Cohesion: 0.15
Nodes (8): BuildAssetRESTController(), AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), findPersistedExternalAssetQuoteMatch(), AssetDomService, AssetIntegrationServicesPerSource

### Community 22 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.41
Nodes (5): go_pkg_context, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 23 - "Adapter"
Cohesion: 0.05
Nodes (33): database/sql.DB, database/sql.Result, database/sql.Stmt, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, buildPingContext() (+25 more)

### Community 24 - "golang_ext_util_concurrent.go"
Cohesion: 0.19
Nodes (24): go_pkg_runtime, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount(), closeConcurrentResultChannels() (+16 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "BuildAppError"
Cohesion: 0.21
Nodes (11): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), golang.org/x/text/currency.Unit, DomainValidationError, BuildAppError(), BuildAppErrorFormatted(), BuildUniqueConstraintViolationError() (+3 more)

### Community 28 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.08
Nodes (38): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+30 more)

### Community 29 - "AllocationStructure"
Cohesion: 0.09
Nodes (17): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, go_pkg_database_sql_driver (+9 more)

### Community 40 - "portfolio_allocation_controller.go"
Cohesion: 0.33
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin, go_pkg_net_http

### Community 42 - "asset_rest_model.go"
Cohesion: 0.29
Nodes (13): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetQuoteDTS(), ExternalAssetData (+5 more)

### Community 43 - "Configuration"
Cohesion: 0.16
Nodes (14): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_os, Configuration, RDBMSConfiguration, YahooFinanceConfiguration (+6 more)

### Community 45 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 47 - "rdbms_adapter.go"
Cohesion: 0.12
Nodes (13): assetRowScanner(), go_pkg_database_sql, go_pkg_errors, go_pkg_github_com_lib_pq, go_pkg_github_com_shopspring_decimal, go_pkg_time, database/sql.Row, database/sql.Rows (+5 more)

### Community 49 - "go_pkg_strings"
Cohesion: 0.10
Nodes (18): go_pkg_flag, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres (+10 more)

### Community 50 - "portfolio_allocation_integration_test.go"
Cohesion: 0.22
Nodes (10): postPortfolioAllocationForValidationFailure(), TestGetAvailableHistoryObservations(), TestGetPortfolioAllocationHistory(), TestGetPortfolioAllocationHistoryForObservationTimestamp(), TestGetPortfolioAllocationHistoryForObservationTimestampNoneFound(), TestPostPortfolioAllocationHistoryValidation_AssetNameExceedsMaxLength(), TestPostPortfolioAllocationHistoryValidation_AssetTickerExceedsMaxLength(), TestPostPortfolioAllocationHistoryValidation_ClassExceedsMaxLength() (+2 more)

### Community 52 - "database/sql.NullString"
Cohesion: 0.32
Nodes (7): TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), database/sql.NullString, StringPointerToNullString(), StringToNullString(), ValueToString(), assertPersistedPortfolioFromAttributes()

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `reflect.Type`, `portfolio_allocation_controller.go`, `GinServer`, `Configuration`, `rdbms_adapter.go`, `go_pkg_strings`, `PortfolioAllocation`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Adapter`?**
  _High betweenness centrality (0.057) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `allocation_plan_domain_service.go`, `asset_rest_model.go`, `rdbms_adapter.go`, `PortfolioAllocation`, `BuildCleanupFunctionBuilder`, `AssetDomService`?**
  _High betweenness centrality (0.052) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `Configuration`, `App`, `rdbms_adapter.go`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Are the 76 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 76 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.05822784810126582 - nodes in this community are weakly interconnected._
- **Should `allocation_plan_domain_service.go` be split into smaller, more focused modules?**
  _Cohesion score 0.07686932215234102 - nodes in this community are weakly interconnected._