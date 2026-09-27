# Graph Report - go  (2026-09-26)

## Corpus Check
- cluster-only mode — file stats not available

## Summary
- 1050 nodes · 3047 edges · 38 communities (28 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 220 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `da880bc5`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- Community 0
- Community 1
- Community 2
- Community 3
- Community 4
- Community 5
- Community 6
- Community 7
- Community 8
- Community 9
- Community 10
- Community 11
- Community 12
- Community 13
- Community 14
- Community 15
- Community 16
- Community 17
- Community 18
- Community 19
- Community 20
- Community 21
- Community 22
- Community 23
- Community 24
- Community 25
- Community 26
- Community 27
- Community 28
- Community 29
- Community 30
- Community 31
- Community 35
- Community 36
- Community 37

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 75 edges
2. `App` - 35 edges
3. `PortfolioAllocation` - 32 edges
4. `PropagateAsAppErrorWithNewMessage()` - 31 edges
5. `Asset` - 30 edges
6. `NewMapTree()` - 27 edges
7. `BuildCleanupFunctionBuilder()` - 26 edges
8. `IsZeroValue()` - 24 edges
9. `AllocationPlan` - 23 edges
10. `PlannedAllocation` - 23 edges

## Surprising Connections (you probably didn't know these)
- `AggregateAndMapToPortfolioHistoryDTSs()` --references--> `PortfolioAllocation`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio_allocation.go
- `aggregateHistoryAsDTSMap()` --references--> `PortfolioAllocation`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio_allocation.go
- `MapToAnalysisOptionsDTS()` --references--> `AnalysisOptions`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio.go
- `MapToDivergenceAnalysisDTS()` --references--> `DivergenceAnalysis`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/divergence.go
- `mapToObservationTimestampDTS()` --references--> `PortfolioObservationTimestamp`  [EXTRACTED]
  api/rest/model/portfolio_rest_mapping.go → domain/portfolio_allocation.go

## Import Cycles
- None detected.

## Communities (38 total, 10 thin omitted)

### Community 0 - "Community 0"
Cohesion: 0.05
Nodes (58): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), mapToAllocationPlanIdentifierDTS(), mapToAllocationPlanIdentifierDTSs(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS() (+50 more)

### Community 1 - "Community 1"
Cohesion: 0.06
Nodes (47): BuildYahooFinanceAssetIntegrationClient(), TestQuoteAssetLastClosePrice_IAU(), TestSearchAssets_IAU(), go_pkg_context, go_pkg_errors, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model (+39 more)

### Community 2 - "Community 2"
Cohesion: 0.05
Nodes (44): BuildAllocationPlanRESTController(), BuildAssetRESTController(), BuildDivergenceAnalysisRESTController(), BuildPortfolioAllocationRESTController(), BuildPortfolioRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAnalysisConfigurationAppService() (+36 more)

### Community 3 - "Community 3"
Cohesion: 0.07
Nodes (57): YahooFinanceAssetIntegrationService, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+49 more)

### Community 4 - "Community 4"
Cohesion: 0.06
Nodes (41): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, github.com/go-playground/universal-translator.Translator, reflect.Kind, reflect.StructField, reflect.Type, reflect.Value (+33 more)

### Community 5 - "Community 5"
Cohesion: 0.05
Nodes (42): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_nhatthm_httpmock, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres (+34 more)

### Community 6 - "Community 6"
Cohesion: 0.10
Nodes (32): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+24 more)

### Community 7 - "Community 7"
Cohesion: 0.08
Nodes (32): AllocationHierarchy, appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData(), readValidationData(), stripRootFromBranches() (+24 more)

### Community 8 - "Community 8"
Cohesion: 0.07
Nodes (26): assetRowScanner(), go_pkg_database_sql, database/sql.Row, database/sql.Rows, database/sql.Tx, github.com/go-ozzo/ozzo-dbx.Query, BuildILikeSubstringPattern(), T (+18 more)

### Community 9 - "Community 9"
Cohesion: 0.08
Nodes (34): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), TestGetKnownAssetsRejectsExcessSearchTerms(), TestGetKnownAssetsReturnsEmptyForNonblankZeroTermSearch(), TestParseAssetTextSearch(), malformedQueryBindingTarget (+26 more)

### Community 10 - "Community 10"
Cohesion: 0.14
Nodes (36): testing.T, TestGetAllocationPlans(), TestPostAllocationPlanValidation_ChildlessHierarchyBranches(), TestPostAllocationPlanValidation_DuplicateHierarchicalIds(), TestPostAllocationPlanValidation_EmptyDetails(), TestPostAllocationPlanValidation_EmptyHierarchicalId(), TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches(), TestPostAllocationPlanValidation_MissingDetails() (+28 more)

### Community 11 - "Community 11"
Cohesion: 0.09
Nodes (18): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchyLevel, AllocationStructure, HierarchicalId, go_pkg_database_sql_driver (+10 more)

### Community 12 - "Community 12"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "Community 13"
Cohesion: 0.13
Nodes (30): net/http.Response, getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch() (+22 more)

### Community 14 - "Community 14"
Cohesion: 0.19
Nodes (24): go_pkg_runtime, context.CancelFunc, sync.WaitGroup, I, concurrentSliceResult, buildConcurrentInputChannel(), buildFlatMapWorkerCount(), closeConcurrentResultChannels() (+16 more)

### Community 15 - "Community 15"
Cohesion: 0.12
Nodes (20): go_pkg_encoding_json, go_pkg_regexp, RequestOption, CloseResponseBody(), DecodeJSONResponse(), ExecuteGet(), ExecuteGetJSON(), T (+12 more)

### Community 16 - "Community 16"
Cohesion: 0.13
Nodes (9): database/sql.DB, database/sql.Result, github.com/go-ozzo/ozzo-dbx.DB, buildPingContext(), Adapter, SQLTransactionalContext, withTransaction(), contextKey (+1 more)

### Community 17 - "Community 17"
Cohesion: 0.26
Nodes (24): github.com/go-ozzo/ozzo-dbx.NullStringMap, TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(), assertPersistedAssetWithExternalData(), TestPostPortfolioAllocationHistoryFullMerge(), TestPostPortfolioAllocationHistoryInsertEmptyZeroTimestamp() (+16 more)

### Community 18 - "Community 18"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 19 - "Community 19"
Cohesion: 0.18
Nodes (9): BuildAssetRDBMSRepository(), BuildUniqueConstraintViolationError(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), IsUniqueConstraintViolation(), ToPointerSlice(), AssetRDBMSRepository, PortfolioRDBMSRepository (+1 more)

### Community 20 - "Community 20"
Cohesion: 0.21
Nodes (11): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), GetPlanType(), PlanType, AllocationPlan, PlannedAllocation, buildAllocationPlanFromRow(), buildPlannedAllocationFromRow() (+3 more)

### Community 21 - "Community 21"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 22 - "Community 22"
Cohesion: 0.23
Nodes (4): github.com/gin-gonic/gin.Engine, net/http.Server, GinServer, GinServerRESTController

### Community 23 - "Community 23"
Cohesion: 0.27
Nodes (11): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), addAssetExternalDataRestoreCleanup(), capturePersistedAssetExternalData(), ExecuteDBQuery(), TestGetPortfolioAllocationHistoryOmitsExternalAssetWhenNull(), TestGetPortfolioAllocationHistoryWithProjectedExternalAsset(), TestPostPortfolioAllocationHistoryDoesNotOverwriteExistingAssetExternalData() (+3 more)

### Community 24 - "Community 24"
Cohesion: 0.29
Nodes (5): BuildAppError(), BuildQueryInTransaction(), T, ToSQLTransactionalContext(), AllocationPlanRDBMSRepository

### Community 25 - "Community 25"
Cohesion: 0.24
Nodes (12): database/sql.NullString, StringPointerToNullString(), StringToNullString(), putForValidationFailure(), TestPutPortfolio(), TestPutPortfolioFailureWithNameExceedingMaxLength(), TestPutPortfolioFailureWithoutMandatoryFields(), TestPutPortfolioWithAllocationStructure() (+4 more)

### Community 26 - "Community 26"
Cohesion: 0.27
Nodes (7): BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildPortfolioAllocationRepository(), BuildPortfolioRepository(), github.com/go-ozzo/ozzo-dbx.Rows, RepositoryRDBMSAdapter, AllocationRDBMSRepository

### Community 27 - "Community 27"
Cohesion: 0.22
Nodes (7): demoStringer, TestCustomSlice_PrettyString_Empty(), TestCustomSlice_PrettyString_Int(), TestCustomSlice_PrettyString_Single(), TestCustomSlice_PrettyString_String(), TestCustomSlice_PrettyString_Stringer(), TestCustomSlice_PrettyString_Struct()

## Knowledge Gaps
- **13 isolated node(s):** `contextKey`, `MapIterator`, `ErrorResponse`, `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS` (+8 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 89 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `Community 1` to `Community 2`, `Community 4`, `Community 5`, `Community 16`, `Community 22`?**
  _High betweenness centrality (0.061) - this node is a cross-community bridge._
- **Why does `Asset` connect `Community 3` to `Community 2`, `Community 7`, `Community 8`, `Community 13`, `Community 19`, `Community 20`, `Community 24`?**
  _High betweenness centrality (0.042) - this node is a cross-community bridge._
- **Are the 73 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestGetAllocationPlans()` and `TestPostAllocationPlanForInsertion()`) actually correct?**
  _`deferCloseResponseBody()` has 73 INFERRED edges - model-reasoned connections that need verification._
- **What connects `contextKey`, `MapIterator`, `ErrorResponse` to the rest of the system?**
  _13 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `Community 0` be split into smaller, more focused modules?**
  _Cohesion score 0.05467856325783574 - nodes in this community are weakly interconnected._
- **Should `Community 1` be split into smaller, more focused modules?**
  _Cohesion score 0.06134371957156767 - nodes in this community are weakly interconnected._
- **Should `Community 2` be split into smaller, more focused modules?**
  _Cohesion score 0.05333333333333334 - nodes in this community are weakly interconnected._