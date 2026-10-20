# Graph Report - go  (2026-10-21)

## Corpus Check
- 123 files · ~65,483 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1125 nodes · 3446 edges · 55 communities (43 shown, 12 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 261 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `4f8a15da`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- AllocationPlanDomService
- Adapter
- IsZeroValue
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- reflect.Type
- golang_ext_util_concurrent.go
- deferCloseResponseBody
- testing.T
- BuildAppError
- asset_integration_test.go
- infra.go
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- App
- NewOrderedMapIterator
- ExternalAsset
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- asset_rest_model.go
- go_pkg_github_com_golang_glog
- rdbms_query.go
- context.Context
- assert_json_extension.go
- ExternalAssetQuote
- T
- portfolio_rest_mapping.go
- AllocationStructure
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- manage_db.go
- PotentialDivergence
- Asset
- GinServer
- asset_integration_client_yahoo.go
- portfolio_divergence_analysis_service.go
- assert_db.go
- PropagateAsAppErrorWithNewMessage
- asset_integration_client_yahoo_model.go
- portfolio_allocation_controller.go
- http_client.go
- AllocationDomService
- parseAssetTextSearch
- ForeignKeyConstraintViolationError
- portfolio_allocation_repository_model.go

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
- `BuildAssetRESTController()` --references--> `AssetDomService`  [EXTRACTED]
  api/rest/asset_controller.go → domain/service/asset_domain_service.go
- `DivergenceAnalysisRESTController` --references--> `PortfolioDivergenceAnalysisAppService`  [EXTRACTED]
  api/rest/divergence_analysis_controller.go → application/portfolio_divergence_analysis_service.go
- `BuildDivergenceAnalysisRESTController()` --references--> `PortfolioDivergenceAnalysisAppService`  [EXTRACTED]
  api/rest/divergence_analysis_controller.go → application/portfolio_divergence_analysis_service.go

## Import Cycles
- None detected.

## Communities (55 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (30): validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), bindAndValidatePortfolioSnapshot(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.HandlersChain (+22 more)

### Community 1 - "AllocationPlanDomService"
Cohesion: 0.17
Nodes (18): BuildAllocationPlanRESTController(), BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, AllocationPlanRepository, BuildAllocationPlanRepository() (+10 more)

### Community 2 - "Adapter"
Cohesion: 0.09
Nodes (22): database/sql.DB, database/sql.Result, database/sql.Stmt, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.DB, createBulkInsertPreparedStatement(), executeBulkInsertPreparedStatement(), Adapter (+14 more)

### Community 3 - "IsZeroValue"
Cohesion: 0.05
Nodes (44): cleanNilAllocations(), mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AllocationHierarchy, PlannedAllocation, AssetsPerTicker (+36 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.24
Nodes (11): go_pkg_context, go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_errors, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms (+3 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (43): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+35 more)

### Community 6 - "golang_ext_util_concurrent.go"
Cohesion: 0.17
Nodes (25): go_pkg_runtime, context.CancelFunc, sync.WaitGroup, I, buildPingContext(), concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount() (+17 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.12
Nodes (33): assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds(), TestPostAllocationPlanUpdatesOwnedMetadata() (+25 more)

### Community 8 - "testing.T"
Cohesion: 0.11
Nodes (33): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), go_pkg_encoding_json, testing.T, postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound() (+25 more)

### Community 9 - "BuildAppError"
Cohesion: 0.24
Nodes (6): classifyObservationWriteError(), BuildAppError(), BuildUniqueConstraintViolationError(), IsUniqueConstraintViolation(), ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 10 - "asset_integration_test.go"
Cohesion: 0.15
Nodes (21): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+13 more)

### Community 11 - "infra.go"
Cohesion: 0.06
Nodes (34): go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_path_filepath, go_pkg_sync, github.com/nhatthm/httpmock.Server (+26 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.15
Nodes (19): TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert (+11 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.14
Nodes (41): TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRejectsWhitespaceTags(), assertPersistedObservationAllocation() (+33 more)

### Community 15 - "App"
Cohesion: 0.10
Nodes (22): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os, go_pkg_os_signal (+14 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "ExternalAsset"
Cohesion: 0.26
Nodes (10): mapToExternalAsset(), mapToExternalAssetData(), ExternalAsset, AssetExternalSource, mapToExternalAsset(), mapToExternalAssets(), YahooFinanceSearchQuoteDTS, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration (+2 more)

### Community 18 - "PortfolioAllocation"
Cohesion: 0.16
Nodes (8): AllocationHierarchyLevel, preparePortfolioAllocationTempInserts(), PortfolioAllocation, PortfolioAllocationRepository, PortfolioObservationTimestamp, BuildPortfolioAllocationDomService(), PortfolioAllocationDomService, PortfolioAllocationRDBMSRepository

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.21
Nodes (24): database/sql.NullString, StringPointerToNullString(), StringToNullString(), TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote() (+16 more)

### Community 21 - "asset_rest_model.go"
Cohesion: 0.25
Nodes (16): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+8 more)

### Community 22 - "go_pkg_github_com_golang_glog"
Cohesion: 0.29
Nodes (5): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, ConfigLogger(), main()

### Community 23 - "rdbms_query.go"
Cohesion: 0.09
Nodes (23): assetRowScanner(), database/sql.Row, database/sql.Rows, github.com/go-ozzo/ozzo-dbx.Query, BuildILikeSubstringPattern(), T, withParams(), processParamsForPostgreSQL() (+15 more)

### Community 24 - "context.Context"
Cohesion: 0.19
Nodes (16): allocationIterationMappingContextValue, contextKey, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue(), getHierarchySubIterationContextValue() (+8 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "ExternalAssetQuote"
Cohesion: 0.33
Nodes (5): ExternalAssetQuote, mapToExternalAssetQuote(), findPersistedExternalAssetQuoteMatch(), golang.org/x/text/currency.Unit, BuildAppErrorFormatted()

### Community 28 - "portfolio_rest_mapping.go"
Cohesion: 0.07
Nodes (54): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+46 more)

### Community 29 - "AllocationStructure"
Cohesion: 0.10
Nodes (15): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationStructure, HierarchicalId, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, database/sql/driver.Value (+7 more)

### Community 40 - "manage_db.go"
Cohesion: 0.48
Nodes (4): github.com/go-ozzo/ozzo-dbx.Params, createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 41 - "PotentialDivergence"
Cohesion: 0.18
Nodes (13): divergenceAnalysisContextValue, getDivergenceAnalysisContextValue(), getPotentialDivergenceMapContextValue(), buildPotentialDivergence(), calculateCurrentDivergenceValuesFromReferencedPlan(), calculateDivergenceValue(), checkAndGeneratePotentialDivergencesOnHierarchy(), generateAndAttachPotentialDivergenceForPlannedAllocation() (+5 more)

### Community 42 - "Asset"
Cohesion: 0.16
Nodes (11): Asset, AssetIntegrationService, AssetRepository, buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetDomService(), collectIntegrationServices(), AssetDomService (+3 more)

### Community 43 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 44 - "asset_integration_client_yahoo.go"
Cohesion: 0.23
Nodes (9): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceAssetIntegrationClient, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_net_url, PropagateAsAppError() (+1 more)

### Community 45 - "portfolio_divergence_analysis_service.go"
Cohesion: 0.29
Nodes (5): BuildPortfolioDivergenceAnalysisAppService(), PortfolioDivergenceAnalysisAppService, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService

### Community 46 - "assert_db.go"
Cohesion: 0.50
Nodes (8): github.com/go-ozzo/ozzo-dbx.NullStringMap, assertAllRows(), AssertDBWithQuery(), assertRowColumn(), assertSingleRow(), executeQueryForMultipleRows(), validateRowCount(), AssertableNullStringMap

### Community 47 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.15
Nodes (10): Portfolio, github.com/go-ozzo/ozzo-dbx.Rows, PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, ToPointerSlice(), AssetRDBMSRepository (+2 more)

### Community 48 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.43
Nodes (7): extractLastClose(), YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS, YahooFinanceChartQuoteIndicatorDTS

### Community 49 - "portfolio_allocation_controller.go"
Cohesion: 0.18
Nodes (14): BuildAssetRESTController(), BuildDivergenceAnalysisRESTController(), BuildPortfolioRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation (+6 more)

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 51 - "AllocationDomService"
Cohesion: 0.47
Nodes (3): AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 52 - "parseAssetTextSearch"
Cohesion: 0.47
Nodes (4): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch(), TestParseAssetTextSearch()

### Community 54 - "portfolio_allocation_repository_model.go"
Cohesion: 0.90
Nodes (4): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), portfolioAllocationJoinedRowDTS

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 94 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `AllocationPlanDomService`, `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `infra.go`, `GinServer`, `portfolio_allocation_controller.go`, `ExternalAsset`, `go_pkg_github_com_golang_glog`?**
  _High betweenness centrality (0.059) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `IsZeroValue`, `BuildAppError`, `PropagateAsAppErrorWithNewMessage`, `PortfolioAllocation`, `ExecuteDBQuery`, `asset_rest_model.go`, `portfolio_allocation_repository_model.go`, `rdbms_query.go`, `parseAssetTextSearch`, `ExternalAssetQuote`, `portfolio_rest_mapping.go`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `AllocationStructure` to `IsZeroValue`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 87 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 87 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.08087431693989071 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.0858974358974359 - nodes in this community are weakly interconnected._