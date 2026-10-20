# Graph Report - go  (2026-10-21)

## Corpus Check
- 123 files · ~65,484 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1125 nodes · 3446 edges · 56 communities (44 shown, 12 thin omitted)
- Extraction: 92% EXTRACTED · 8% INFERRED · 0% AMBIGUOUS · INFERRED: 261 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `0ad59c75`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- .buildAppComponents
- Adapter
- IsZeroValue
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- reflect.Type
- golang_ext_util_concurrent.go
- deferCloseResponseBody
- portfolio_integration_test.go
- go_pkg_strings
- testing.T
- infra_db.go
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- asset_integration_service_yahoo.go
- NewOrderedMapIterator
- ExternalAsset
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- ExecuteDBQuery
- asset_rest_model.go
- infra.go
- T
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
- go_pkg_github_com_go_ozzo_ozzo_dbx
- portfolio_divergence_analysis_service.go
- Asset
- GinServer
- App
- Portfolio
- assert_db.go
- PropagateAsAppErrorWithNewMessage
- asset_integration_client_yahoo_model.go
- go_pkg_time
- http_client.go
- portfolio_observation_ownership_migration_integration_test.go
- parseAssetTextSearch
- ForeignKeyConstraintViolationError
- go_pkg_github_com_shopspring_decimal
- joinAny

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

## Communities (56 total, 12 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.08
Nodes (29): validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), bindAndValidatePortfolioSnapshot(), BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService() (+21 more)

### Community 1 - ".buildAppComponents"
Cohesion: 0.23
Nodes (14): BuildAllocationPlanRESTController(), BuildPortfolioAllocationRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), PortfolioAllocationManagementAppService, BuildAllocationPlanRepository(), BuildAllocationRepository() (+6 more)

### Community 2 - "Adapter"
Cohesion: 0.13
Nodes (9): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey (+1 more)

### Community 3 - "IsZeroValue"
Cohesion: 0.06
Nodes (44): cleanNilAllocations(), mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AllocationHierarchy, AllocationPlanRepository, PlannedAllocation (+36 more)

### Community 4 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.27
Nodes (12): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin (+4 more)

### Community 5 - "reflect.Type"
Cohesion: 0.05
Nodes (48): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_github_com_go_playground_validator_v10, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type (+40 more)

### Community 6 - "golang_ext_util_concurrent.go"
Cohesion: 0.19
Nodes (24): go_pkg_runtime, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount(), closeConcurrentResultChannels() (+16 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.15
Nodes (28): net/http.Response, assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds() (+20 more)

### Community 8 - "portfolio_integration_test.go"
Cohesion: 0.19
Nodes (17): postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios(), TestPostPortfolioFailureWithNameExceedingMaxLength(), TestPostPortfolioFailureWithoutMandatoryFields() (+9 more)

### Community 9 - "go_pkg_strings"
Cohesion: 0.15
Nodes (13): assetRowScanner(), go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_lib_pq, go_pkg_strings, database/sql.Rows, database/sql.Stmt (+5 more)

### Community 10 - "testing.T"
Cohesion: 0.13
Nodes (29): testing.T, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+21 more)

### Community 11 - "infra_db.go"
Cohesion: 0.08
Nodes (26): go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_sync, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB (+18 more)

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.09
Nodes (23): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), malformedQueryBindingTarget, go_pkg_encoding_json, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra (+15 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.14
Nodes (43): database/sql.NullString, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), TestPostAllocationPlanRequiresExplicitValidPlannedDate() (+35 more)

### Community 15 - "asset_integration_service_yahoo.go"
Cohesion: 0.13
Nodes (18): YahooFinanceAssetIntegrationService, BuildYahooFinanceAssetIntegrationService(), BuildYahooFinanceAssetIntegrationClient(), YahooFinanceAssetIntegrationClient, TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient (+10 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "ExternalAsset"
Cohesion: 0.70
Nodes (4): ExternalAsset, mapToExternalAsset(), mapToExternalAssets(), YahooFinanceSearchQuoteDTS

### Community 18 - "PortfolioAllocation"
Cohesion: 0.20
Nodes (7): AllocationHierarchyLevel, preparePortfolioAllocationTempInserts(), PortfolioAllocation, PortfolioAllocationRepository, PortfolioObservationTimestamp, BuildPortfolioAllocationDomService(), PortfolioAllocationDomService

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "ExecuteDBQuery"
Cohesion: 0.18
Nodes (28): go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_net_url, go_pkg_strconv, TestGetExternalAssetsFailure(), TestGetExternalAssetsSuccess(), TestGetKnownAssetsIncludesPersistedExternalData(), TestPutAsset(), TestPutAssetClearsExternalDataWhenNull() (+20 more)

### Community 21 - "asset_rest_model.go"
Cohesion: 0.23
Nodes (18): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+10 more)

### Community 22 - "infra.go"
Cohesion: 0.15
Nodes (10): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_net, go_pkg_path_filepath, contextKey (+2 more)

### Community 23 - "T"
Cohesion: 0.09
Nodes (20): database/sql.Row, github.com/go-ozzo/ozzo-dbx.Query, github.com/go-ozzo/ozzo-dbx.Rows, T, withParams(), processParamsForPostgreSQL(), processSQL(), ReturningIntIdSingleRowScanner() (+12 more)

### Community 24 - "context.Context"
Cohesion: 0.16
Nodes (17): allocationIterationMappingContextValue, contextKey, buildAllocationIterationContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue(), getHierarchySubIterationContextValue(), getPotentialDivergenceMapContextValue() (+9 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "ExternalAssetQuote"
Cohesion: 0.43
Nodes (5): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), golang.org/x/text/currency.Unit, BuildAppErrorFormatted()

### Community 28 - "portfolio_rest_mapping.go"
Cohesion: 0.07
Nodes (54): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+46 more)

### Community 29 - "AllocationStructure"
Cohesion: 0.10
Nodes (15): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationStructure, HierarchicalId, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, database/sql/driver.Value (+7 more)

### Community 40 - "go_pkg_github_com_go_ozzo_ozzo_dbx"
Cohesion: 0.29
Nodes (7): go_pkg_github_com_go_ozzo_ozzo_dbx, github.com/go-ozzo/ozzo-dbx.Params, FetchWithDBQuery(), TestPortfolioHistoryOrderingIsChronologicalAndDeterministic(), createDBCleanupFunctionMulti(), CleanupFunctionBuilder, testSQLParamsPair

### Community 41 - "portfolio_divergence_analysis_service.go"
Cohesion: 0.17
Nodes (16): divergenceAnalysisContextValue, buildDivergenceAnalysisContext(), getDivergenceAnalysisContextValue(), attachToRootIfTopLevel(), buildAndAttachPotentialDivergenceIfNotExists(), buildDivergenceAnalysis(), buildPotentialDivergence(), calculateCurrentDivergenceValuesFromReferencedPlan() (+8 more)

### Community 42 - "Asset"
Cohesion: 0.13
Nodes (13): BuildAssetRESTController(), Asset, AssetIntegrationService, AssetRepository, buildAssetInsertValue(), buildAssetInsertValues(), BuildAssetDomService(), collectIntegrationServices() (+5 more)

### Community 43 - "GinServer"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 44 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 45 - "Portfolio"
Cohesion: 0.22
Nodes (7): BuildPortfolioDivergenceAnalysisAppService(), PortfolioDivergenceAnalysisAppService, Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, PortfolioRDBMSRepository

### Community 46 - "assert_db.go"
Cohesion: 0.50
Nodes (8): github.com/go-ozzo/ozzo-dbx.NullStringMap, assertAllRows(), AssertDBWithQuery(), assertRowColumn(), assertSingleRow(), executeQueryForMultipleRows(), validateRowCount(), AssertableNullStringMap

### Community 47 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.11
Nodes (16): classifyObservationWriteError(), BuildAppError(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, IsUniqueConstraintViolation() (+8 more)

### Community 48 - "asset_integration_client_yahoo_model.go"
Cohesion: 0.23
Nodes (10): buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), YahooFinanceChartResponseDTS, YahooFinanceChartResultDTS, PropagateAsAppError(), YahooFinanceChartDTS, YahooFinanceChartIndicatorsDTS, YahooFinanceChartMetaDTS (+2 more)

### Community 49 - "go_pkg_time"
Cohesion: 0.28
Nodes (6): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_time, DivergenceAnalysisRESTController

### Community 50 - "http_client.go"
Cohesion: 0.43
Nodes (7): RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T, WithHeader()

### Community 51 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.38
Nodes (9): database/sql.Tx, migratedObservation, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema(), seedLegacyObservations() (+1 more)

### Community 52 - "parseAssetTextSearch"
Cohesion: 0.47
Nodes (4): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch(), TestParseAssetTextSearch()

### Community 54 - "go_pkg_github_com_shopspring_decimal"
Cohesion: 0.39
Nodes (7): AssetExternalSource, buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, portfolioAllocationJoinedRowDTS

### Community 55 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 94 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **12 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `.buildAppComponents`, `Adapter`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `reflect.Type`, `infra_db.go`, `GinServer`, `asset_integration_service_yahoo.go`, `go_pkg_time`, `infra.go`?**
  _High betweenness centrality (0.058) - this node is a cross-community bridge._
- **Why does `Asset` connect `Asset` to `IsZeroValue`, `go_pkg_strings`, `PropagateAsAppErrorWithNewMessage`, `PortfolioAllocation`, `ExecuteDBQuery`, `asset_rest_model.go`, `go_pkg_github_com_shopspring_decimal`, `parseAssetTextSearch`, `portfolio_rest_mapping.go`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Why does `HierarchicalId` connect `AllocationStructure` to `go_pkg_strings`, `IsZeroValue`?**
  _High betweenness centrality (0.035) - this node is a cross-community bridge._
- **Are the 87 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 87 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.08416130917592052 - nodes in this community are weakly interconnected._
- **Should `Adapter` be split into smaller, more focused modules?**
  _Cohesion score 0.12666666666666668 - nodes in this community are weakly interconnected._