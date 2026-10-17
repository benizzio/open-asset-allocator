# Graph Report - go  (2026-10-17)

## Corpus Check
- 119 files · ~59,852 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1075 nodes · 3201 edges · 40 communities (29 shown, 11 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 236 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `aad78f79`
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
- testing.T
- App
- portfolio_rest_mapping.go
- asset_integration_test.go
- T
- golang_ext_tree_test.go
- go_pkg_testing
- assert_db.go
- asset_rest_model.go
- NewOrderedMapIterator
- portfolio_integration_test.go
- golang_ext_custom_slice_test.go
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- AssetDomService
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- Adapter
- golang_ext_util_concurrent.go
- assert_json_extension.go
- BuildAppError
- AllocationDomService
- AllocationPlanRDBMSRepository
- HierarchicalId
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests

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

## Communities (40 total, 11 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.07
Nodes (34): BuildAllocationPlanRESTController(), validateAllocationPlanPlannedDate(), MapToPortfolio(), MapToPortfolioDTS(), MapToPortfolioDTSs(), go_pkg_github_com_go_playground_validator_v10, github.com/gin-gonic/gin.Context, github.com/gin-gonic/gin.HandlersChain (+26 more)

### Community 1 - "Portfolio"
Cohesion: 0.22
Nodes (8): BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, Portfolio, PortfolioRepository, BuildPortfolioDomService(), PortfolioDomService, TransactionManager, PortfolioRDBMSRepository

### Community 2 - ".buildAppComponents"
Cohesion: 0.29
Nodes (8): BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), github.com/go-ozzo/ozzo-dbx.Rows, RepositoryRDBMSAdapter, AllocationRDBMSRepository

### Community 3 - "allocation_plan_domain_service.go"
Cohesion: 0.07
Nodes (41): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), cleanNilAllocations(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure (+33 more)

### Community 4 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.18
Nodes (13): Asset, buildAssetInsertValue(), buildAssetInsertValues(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T (+5 more)

### Community 5 - "reflect.Type"
Cohesion: 0.06
Nodes (42): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+34 more)

### Community 6 - "context.Context"
Cohesion: 0.05
Nodes (57): BuildDivergenceAnalysisRESTController(), BuildPortfolioAllocationRESTController(), allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, buildAllocationIterationContext() (+49 more)

### Community 7 - "testing.T"
Cohesion: 0.13
Nodes (40): go_pkg_io, testing.T, TestAllocationPlanDateMigrationBackfillsLegacyRows(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails() (+32 more)

### Community 8 - "App"
Cohesion: 0.06
Nodes (36): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_github_com_nhatthm_httpmock (+28 more)

### Community 9 - "portfolio_rest_mapping.go"
Cohesion: 0.07
Nodes (56): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+48 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.14
Nodes (22): go_pkg_net_url, net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets() (+14 more)

### Community 11 - "T"
Cohesion: 0.31
Nodes (4): T, NewSliceIterator(), SliceIterator, SliceIterator[T]

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.11
Nodes (17): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), TestParseAssetTextSearch(), go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_stretchr_testify_assert (+9 more)

### Community 14 - "assert_db.go"
Cohesion: 0.21
Nodes (28): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+20 more)

### Community 15 - "asset_rest_model.go"
Cohesion: 0.08
Nodes (38): YahooFinanceAssetIntegrationService, MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS() (+30 more)

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "portfolio_integration_test.go"
Cohesion: 0.15
Nodes (21): TestPostAsset(), postPortfolioForValidationFailure(), putForValidationFailure(), TestGetAvailablePortfolioAllocationClasses(), TestGetAvailablePortfolioAllocationClassesNoneFound(), TestGetPortfolio(), TestGetPortfolios(), TestPostPortfolio() (+13 more)

### Community 18 - "golang_ext_custom_slice_test.go"
Cohesion: 0.22
Nodes (7): demoStringer, TestCustomSlice_PrettyString_Empty(), TestCustomSlice_PrettyString_Int(), TestCustomSlice_PrettyString_Single(), TestCustomSlice_PrettyString_String(), TestCustomSlice_PrettyString_Stringer(), TestCustomSlice_PrettyString_Struct()

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.18
Nodes (27): go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_require, github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPostAssetWithDuplicateTicker(), getExternalAssetQuote() (+19 more)

### Community 21 - "AssetDomService"
Cohesion: 0.14
Nodes (10): BuildAssetRESTController(), AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), findPersistedExternalAssetQuoteMatch(), AssetDomService, TestGetKnownAssetsRejectsExcessSearchTerms() (+2 more)

### Community 22 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.06
Nodes (54): assetRowScanner(), malformedQueryBindingTarget, go_pkg_context, go_pkg_database_sql, go_pkg_database_sql_driver, go_pkg_encoding_json, go_pkg_errors, go_pkg_flag (+46 more)

### Community 23 - "Adapter"
Cohesion: 0.05
Nodes (33): database/sql.DB, database/sql.Result, database/sql.Stmt, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Query, RDBMSConfiguration, buildPingContext() (+25 more)

### Community 24 - "golang_ext_util_concurrent.go"
Cohesion: 0.09
Nodes (34): go_pkg_runtime, go_pkg_slices, go_pkg_sync, context.CancelFunc, sync.Mutex, sync.WaitGroup, testing.TB, I (+26 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "BuildAppError"
Cohesion: 0.23
Nodes (10): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), golang.org/x/text/currency.Unit, DomainValidationError, BuildAppError(), BuildAppErrorFormatted(), logError() (+2 more)

### Community 27 - "AllocationDomService"
Cohesion: 0.38
Nodes (4): BuildPortfolioRESTController(), AllocationRepository, BuildAllocationDomService(), AllocationDomService

### Community 29 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

## Knowledge Gaps
- **15 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `contextKey`, `github.com/benizzio/open-asset-allocator` (+10 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 91 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **11 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `.buildAppComponents`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `Adapter`?**
  _High betweenness centrality (0.048) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `allocation_plan_domain_service.go`, `context.Context`, `portfolio_rest_mapping.go`, `asset_rest_model.go`, `BuildCleanupFunctionBuilder`, `AssetDomService`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`?**
  _High betweenness centrality (0.046) - this node is a cross-community bridge._
- **Why does `IsZeroValue()` connect `github.com/gin-gonic/gin.Context` to `Portfolio`, `allocation_plan_domain_service.go`, `reflect.Type`, `context.Context`, `portfolio_rest_mapping.go`, `golang_ext_tree_test.go`, `BuildAppError`?**
  _High betweenness centrality (0.038) - this node is a cross-community bridge._
- **Are the 76 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 76 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _15 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.07259615384615385 - nodes in this community are weakly interconnected._
- **Should `allocation_plan_domain_service.go` be split into smaller, more focused modules?**
  _Cohesion score 0.06758832565284179 - nodes in this community are weakly interconnected._