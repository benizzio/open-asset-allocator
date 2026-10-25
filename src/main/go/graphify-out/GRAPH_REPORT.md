# Graph Report - go  (2026-10-25)

## Corpus Check
- 131 files · ~73,617 words
- Verdict: corpus is large enough that graph structure adds value.
- Unclassified: 2 file(s) not represented in the graph (top: .toml 1, (none) 1)

## Summary
- 1239 nodes · 3891 edges · 65 communities (55 shown, 10 thin omitted)
- Extraction: 93% EXTRACTED · 7% INFERRED · 0% AMBIGUOUS · INFERRED: 290 edges (avg confidence: 0.85)
- Token cost: 0 input · 0 output

## Graph Freshness
- Built from commit: `19002a41`
- Run `git rev-parse HEAD` and compare to check if the graph is stale.
- Run `graphify update .` after code changes (no API cost).

## Community Hubs (Navigation)
- github.com/gin-gonic/gin.Context
- AllocationPlan
- Adapter
- error.go
- go_pkg_database_sql
- CustomFieldError
- context.Context
- deferCloseResponseBody
- testing.T
- portfolio_allocation_pagination_integration_test.go
- asset_integration_test.go
- go_pkg_github_com_benizzio_open_asset_allocator_langext
- golang_ext_tree_test.go
- go_pkg_testing
- portfolio_allocation_integration_test.go
- HierarchicalId
- NewOrderedMapIterator
- SQLTransactionalContext
- PortfolioAllocation
- golang_ext_util_struct_unify_pointers_test.go
- BuildCleanupFunctionBuilder
- AssetDomService
- logFanOutConsumer
- PropagateAsAppErrorWithNewMessage
- allocation_plan_domain_service.go
- assert_json_extension.go
- infra.go
- AllocationHierarchy
- BuildAppError
- asset_rest_model.go
- base.go
- CustomSliceTable[T]
- Go Cognitive Complexity Limit
- Go Coding Standards
- github.com/benizzio/open-asset-allocator
- Backend Linting Standards
- External Integration Tests
- Portfolio
- portfolio_observation_ownership_migration_integration_test.go
- portfolio_rest_mapping.go
- pagination_cursor.go
- App
- GinServer
- AssertDBWithQueryMultipleRows
- portfolio_allocation_controller.go
- IsZeroValue
- reflect.Type
- go_pkg_encoding_json
- go_pkg_fmt
- AllocationPlanDomService
- ForeignKeyConstraintViolationError
- custom_deep_validation.go
- golang_ext_util_struct.go
- .buildAppComponents
- AllocationPlanRESTController
- custom_validation_error_handling.go
- CustomValidationErrorsBuilder
- golang_ext_util_slice.go
- PortfolioAllocationRDBMSRepository
- go_pkg_reflect
- parseAssetTextSearch
- joinAny

## God Nodes (most connected - your core abstractions)
1. `deferCloseResponseBody()` - 90 edges
2. `BuildCleanupFunctionBuilder()` - 44 edges
3. `PropagateAsAppErrorWithNewMessage()` - 42 edges
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
- `AllocationPlanRESTController` --references--> `AllocationPlanManagementAppService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → application/allocation_plan_management_service.go
- `AllocationPlanRESTController` --references--> `AllocationPlanDomService`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/service/allocation_plan_domain_service.go
- `allocationPlanCursorFromPayload()` --references--> `AllocationPlanPageCursor`  [EXTRACTED]
  api/rest/allocation_plan_controller.go → domain/allocation_plan.go

## Import Cycles
- None detected.

## Communities (65 total, 10 thin omitted)

### Community 0 - "github.com/gin-gonic/gin.Context"
Cohesion: 0.13
Nodes (21): validateAllocationPlanPlannedDate(), MapToExternalAssetQuoteDTS(), bindAndValidatePortfolioSnapshot(), readHistoryObservationTimestampId(), github.com/gin-gonic/gin.Context, HandleAPIError(), handleDomainError(), handleInfrastructureError() (+13 more)

### Community 1 - "AllocationPlan"
Cohesion: 0.09
Nodes (35): AllocationPlanDTS, MapToAllocationPlan(), mapToAllocationPlanAssetDTS(), mapToAllocationPlanDTS(), MapToAllocationPlanDTSs(), mapToAssetFromAllocationPlan(), mapToPlannedAllocation(), mapToPlannedAllocationDTS() (+27 more)

### Community 2 - "Adapter"
Cohesion: 0.06
Nodes (27): assetRowScanner(), database/sql.Result, database/sql.Rows, github.com/go-ozzo/ozzo-dbx.Builder, github.com/go-ozzo/ozzo-dbx.DB, github.com/go-ozzo/ozzo-dbx.Executor, github.com/go-ozzo/ozzo-dbx.Query, buildPingContext() (+19 more)

### Community 3 - "error.go"
Cohesion: 0.17
Nodes (9): buildExternalAssetQuoteValidationError(), DomainValidationError, BuildAppErrorFormattedUnconverted(), BuildDomainValidationError(), AppError, logError(), newAppError(), UniqueConstraintViolationError (+1 more)

### Community 4 - "go_pkg_database_sql"
Cohesion: 0.16
Nodes (11): TestHierarchicalIdIsTopLevel_Empty(), TestHierarchicalIdParentLevelId_Empty(), TestHierarchicalIdValue_AllNonNil(), TestHierarchicalIdValue_WithNilLevels(), go_pkg_database_sql, go_pkg_github_com_lib_pq, database/sql.Row, BuildILikeSubstringPattern() (+3 more)

### Community 5 - "CustomFieldError"
Cohesion: 0.13
Nodes (3): github.com/go-playground/universal-translator.Translator, reflect.Kind, CustomFieldError

### Community 6 - "context.Context"
Cohesion: 0.05
Nodes (67): allocationIterationMappingContextValue, contextKey, divergenceAnalysisContextValue, buildAllocationIterationContext(), buildDivergenceAnalysisContext(), buildHierarchySubIterationContext(), buildPotentialDivergenceMapContext(), getAllocationIterationContextValue() (+59 more)

### Community 7 - "deferCloseResponseBody"
Cohesion: 0.16
Nodes (26): assertPersistedAllocationPlanDetailIds(), assertPersistedAllocationPlanMetadata(), postAllocationPlanRequest(), registerAllocationPlanFixtureRestore(), TestAllocationPlanReadsOrderByPlannedDateAndId(), TestGetAllocationPlans(), TestPostAllocationPlanRejectsForeignPlanAndDetailIds(), TestPostAllocationPlanUpdatesOwnedMetadata() (+18 more)

### Community 8 - "testing.T"
Cohesion: 0.08
Nodes (44): net/http.Response, testing.T, TestGetExternalAssetsFailure(), TestGetExternalAssetsSuccess(), TestPostAsset(), getExternalAssets(), capturedPaginationQuery, TestGetDivergenceAnalysisOptions() (+36 more)

### Community 9 - "portfolio_allocation_pagination_integration_test.go"
Cohesion: 0.11
Nodes (45): go_pkg_net_url, net/url.Values, time.Time, HasQueryParameter(), RawQueryHasParameter(), allocationPlanEndpoint(), allocationPlanPageNames(), allocationPlanPaginationPriorities() (+37 more)

### Community 10 - "asset_integration_test.go"
Cohesion: 0.16
Nodes (18): getAssetsByTextSearch(), postAssetForValidationFailure(), TestGetAssetByIdInvalidId(), TestGetAssetByIdNotFound(), TestGetAssetByIdOrTicker(), TestGetKnownAssets(), TestGetKnownAssetsRejectsInvalidTextSearch(), TestGetKnownAssetsWithTextSearch() (+10 more)

### Community 11 - "go_pkg_github_com_benizzio_open_asset_allocator_langext"
Cohesion: 0.28
Nodes (7): go_pkg_context, go_pkg_errors, go_pkg_github_com_benizzio_open_asset_allocator_domain, go_pkg_github_com_benizzio_open_asset_allocator_infra, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms, go_pkg_github_com_benizzio_open_asset_allocator_infra_util_http_httpclient, go_pkg_github_com_benizzio_open_asset_allocator_langext

### Community 12 - "golang_ext_tree_test.go"
Cohesion: 0.12
Nodes (28): MapTreeNode, T, NewMapTree(), sortBranches(), TestAddBranch_CreatesFullPath(), TestAddBranch_EmptyBranch(), TestAddBranch_ReusesExistingNodes(), TestAddBranchBreakingOnZeroValues_BasicBranch() (+20 more)

### Community 13 - "go_pkg_testing"
Cohesion: 0.23
Nodes (11): malformedQueryBindingTarget, go_pkg_github_com_benizzio_open_asset_allocator_inttest_infra, go_pkg_github_com_benizzio_open_asset_allocator_inttest_util, go_pkg_github_com_go_ozzo_ozzo_dbx, go_pkg_github_com_stretchr_testify_assert, go_pkg_github_com_stretchr_testify_require, go_pkg_io, go_pkg_net_http (+3 more)

### Community 14 - "portfolio_allocation_integration_test.go"
Cohesion: 0.17
Nodes (22): TestPostPortfolioHistoryAcceptsValidTagBoundaries(), TestPostPortfolioHistoryOwnershipValidation(), TestPostPortfolioHistoryRequiresNonblankTag(), assertPersistedObservationAllocation(), assertPersistedObservationMetadata(), portfolioHistoryUpdateJSON(), postPortfolioAllocationForValidationFailure(), postPortfolioHistoryRequest() (+14 more)

### Community 15 - "HierarchicalId"
Cohesion: 0.17
Nodes (5): HierarchicalId, database/sql/driver.Value, BuildNullStringSlice(), NullStringSlice, ValueJsonColumn()

### Community 16 - "NewOrderedMapIterator"
Cohesion: 0.14
Nodes (17): go_pkg_cmp, go_pkg_sort, K, NewOrderedMapIterator(), ExampleNewOrderedMapIterator_intKeys(), ExampleOrderedMapIterator(), TestOrderedMapIteratorCurrent(), TestOrderedMapIteratorEmptyMap() (+9 more)

### Community 17 - "SQLTransactionalContext"
Cohesion: 0.20
Nodes (7): database/sql.DB, SQLTransactionalContext, ToSQLTransactionalContext(), withTransaction(), contextKey, TransactionalContext, AllocationPlanRDBMSRepository

### Community 18 - "PortfolioAllocation"
Cohesion: 0.12
Nodes (17): BuildPortfolioAllocationRESTController(), PortfolioAllocationManagementAppService, mapNewAssetsPerTickerFromPortfolioAllocations(), replacePersistedAssetsOnPortfolioAllocations(), AssetsPerTicker, addObservationCursorFilter(), buildPortfolioAllocationHistoryPage(), buildPortfolioObservationPage() (+9 more)

### Community 19 - "golang_ext_util_struct_unify_pointers_test.go"
Cohesion: 0.22
Nodes (15): go_pkg_github_com_okhomin_gohashcode, T, processItemField(), createBoolPointer(), createFloatPointer(), createIntPointer(), createStringPointer(), TestUnifyStructPointersBasicStringPointers() (+7 more)

### Community 20 - "BuildCleanupFunctionBuilder"
Cohesion: 0.17
Nodes (29): github.com/go-ozzo/ozzo-dbx.Params, TestGetKnownAssetsIncludesPersistedExternalData(), TestGetKnownAssetsWithTextSearchLimit(), TestPutAssetClearsExternalDataWhenNull(), TestPutAssetClearsExternalDataWhenOmitted(), getExternalAssetQuote(), TestGetExternalAssetQuoteNotFound(), TestGetExternalAssetQuoteRejectsInvalidSource() (+21 more)

### Community 21 - "AssetDomService"
Cohesion: 0.13
Nodes (10): BuildAssetRESTController(), AssetIntegrationService, AssetRepository, BuildAssetDomService(), collectIntegrationServices(), findPersistedExternalAssetQuoteMatch(), AssetDomService, TestGetKnownAssetsRejectsExcessSearchTerms() (+2 more)

### Community 22 - "logFanOutConsumer"
Cohesion: 0.10
Nodes (17): go_pkg_github_com_nhatthm_httpmock, github.com/nhatthm/httpmock.Server, sync.Mutex, testing.M, testing.TB, logFanOutConsumer, BuildDeferRegistry(), setupAndRunTests() (+9 more)

### Community 23 - "PropagateAsAppErrorWithNewMessage"
Cohesion: 0.19
Nodes (11): Asset, buildAssetInsertValue(), buildAssetInsertValues(), PropagateAsAppErrorWithNewMessage(), BuildQuery(), BuildQueryInTransaction(), T, ToPointerSlice() (+3 more)

### Community 24 - "allocation_plan_domain_service.go"
Cohesion: 0.25
Nodes (14): appendLevelDescription(), readPlannedAllocationChildlessHierarchyBranchesValidationData(), readPlannedAllocationForRepeatedValidationData(), readPlannedAllocationForSliceSizeTotalsValidationData(), readPlannedAllocationHierarchicalBranchValidationData(), readValidationData(), stripRootFromBranches(), validateHierarchySize() (+6 more)

### Community 25 - "assert_json_extension.go"
Cohesion: 0.42
Nodes (9): go_pkg_regexp, extractArrayIndexInfo(), handleArrayIndexPath(), handleFinalRemoval(), handleRegularFieldPath(), isArrayIndexNotation(), processArrayElements(), removeNestedField() (+1 more)

### Community 26 - "infra.go"
Cohesion: 0.10
Nodes (16): go_pkg_flag, go_pkg_github_com_benizzio_open_asset_allocator_root, go_pkg_github_com_golang_glog, go_pkg_github_com_moby_moby_api_types_container, go_pkg_github_com_testcontainers_testcontainers_go, go_pkg_github_com_testcontainers_testcontainers_go_modules_postgres, go_pkg_github_com_testcontainers_testcontainers_go_wait, go_pkg_net (+8 more)

### Community 27 - "AllocationHierarchy"
Cohesion: 0.22
Nodes (11): mapToAllocationHierarchyLevelDTSs(), mapToAllocationHierarchyLevels(), mapToAllocationStructure(), mapToAllocationStructureDTS(), AllocationHierarchy, AllocationHierarchyLevel, AllocationStructure, T (+3 more)

### Community 28 - "BuildAppError"
Cohesion: 0.24
Nodes (9): ExternalAssetQuote, extractLastClose(), mapToExternalAssetQuote(), classifyObservationWriteError(), golang.org/x/text/currency.Unit, BuildAppError(), BuildAppErrorFormatted(), BuildUniqueConstraintViolationError() (+1 more)

### Community 29 - "asset_rest_model.go"
Cohesion: 0.22
Nodes (17): MapToAsset(), MapToAssetDTS(), MapToAssetDTSs(), MapToAssets(), mapToDomainExternalAssetData(), mapToExternalAssetDataDTS(), MapToExternalAssetDTS(), MapToExternalAssetDTSs() (+9 more)

### Community 40 - "Portfolio"
Cohesion: 0.20
Nodes (8): BuildPortfolioRESTController(), AllocationRepository, Portfolio, PortfolioRepository, BuildAllocationDomService(), AllocationDomService, BuildPortfolioDomService(), PortfolioDomService

### Community 41 - "portfolio_observation_ownership_migration_integration_test.go"
Cohesion: 0.29
Nodes (10): go_pkg_github_com_benizzio_open_asset_allocator_infra_util, go_pkg_os, database/sql.Tx, assertExpectedMigrationFailure(), assertMigratedObservationOwnership(), assertMigrationConstraints(), assertObservationSequenceAfterMigration(), createLegacyObservationSchema() (+2 more)

### Community 42 - "portfolio_rest_mapping.go"
Cohesion: 0.14
Nodes (28): AggregateAndMapToPortfolioHistoryDTSs(), aggregateHistoryAsDTSMap(), buildHistoryDTS(), buildSnapshotDTS(), MapToAnalysisOptionsDTS(), MapToDivergenceAnalysisDTS(), mapToObservationTimestampDTS(), MapToPortfolio() (+20 more)

### Community 43 - "pagination_cursor.go"
Cohesion: 0.15
Nodes (19): PaginatedResponse, T, decodePaginationCursor(), decodePaginationCursorBoundary(), decodePaginationCursorPayload(), encodeAllocationPlanPaginationCursor(), encodeObservationPaginationCursor(), encodePaginationCursor() (+11 more)

### Community 44 - "App"
Cohesion: 0.20
Nodes (9): go_pkg_github_com_benizzio_open_asset_allocator_api_rest, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_anticorruption, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_repository, go_pkg_os_signal, go_pkg_syscall, os.Signal, buildStopChannel(), buildStopContext() (+1 more)

### Community 45 - "GinServer"
Cohesion: 0.07
Nodes (33): YahooFinanceAssetIntegrationService, ExternalAsset, BuildYahooFinanceAssetIntegrationService(), mapToExternalAsset(), mapToExternalAssets(), buildQuoteAssetLastClosePriceURL(), buildSearchAssetsURL(), BuildYahooFinanceAssetIntegrationClient() (+25 more)

### Community 46 - "AssertDBWithQueryMultipleRows"
Cohesion: 0.19
Nodes (30): database/sql.NullString, github.com/go-ozzo/ozzo-dbx.NullStringMap, StringPointerToNullString(), StringToNullString(), TestPostAllocationPlanForInsertion(), TestPostAllocationPlanForUpdate_ChangesHierarchicalId(), TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(), TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName() (+22 more)

### Community 47 - "portfolio_allocation_controller.go"
Cohesion: 0.34
Nodes (7): go_pkg_github_com_benizzio_open_asset_allocator_api_rest_model, go_pkg_github_com_benizzio_open_asset_allocator_application, go_pkg_github_com_benizzio_open_asset_allocator_domain_allocation, go_pkg_github_com_benizzio_open_asset_allocator_domain_service, go_pkg_github_com_benizzio_open_asset_allocator_infra_gin, go_pkg_github_com_benizzio_open_asset_allocator_infra_validation, go_pkg_github_com_gin_gonic_gin

### Community 48 - "IsZeroValue"
Cohesion: 0.22
Nodes (6): mapNewAssetsPerTickerFromPlannedAllocations(), replacePersistedAssetsOnPlannedAllocations(), PlannedAllocation, T, IsZeroValue(), ParseableInt

### Community 49 - "reflect.Type"
Cohesion: 0.42
Nodes (11): reflect.StructField, reflect.Type, buildJSONFieldPath(), GetJSONFieldName(), mapIndexedNamespacePart(), mapNamespacePart(), parseNamespace(), buildFieldPath() (+3 more)

### Community 50 - "go_pkg_encoding_json"
Cohesion: 0.38
Nodes (4): go_pkg_encoding_json, go_pkg_strconv, RequestOption, WithHeader()

### Community 51 - "go_pkg_fmt"
Cohesion: 0.15
Nodes (13): go_pkg_database_sql_driver, go_pkg_fmt, go_pkg_github_com_benizzio_open_asset_allocator_domain_infra_integration, go_pkg_github_com_benizzio_open_asset_allocator_infra_rdbms_sqlext, go_pkg_github_com_shopspring_decimal, go_pkg_golang_org_x_text_currency, go_pkg_slices, go_pkg_time (+5 more)

### Community 52 - "AllocationPlanDomService"
Cohesion: 0.31
Nodes (9): BuildDivergenceAnalysisRESTController(), BuildPortfolioAnalysisConfigurationAppService(), PortfolioAnalysisConfigurationAppService, BuildPortfolioDivergenceAnalysisAppService(), PortfolioDivergenceAnalysisAppService, AllocationPlanRepository, BuildAllocationPlanDomService(), AllocationPlanDomService (+1 more)

### Community 53 - "ForeignKeyConstraintViolationError"
Cohesion: 0.29
Nodes (3): ForeignKeyConstraintViolationError, ClassifyForeignKeyViolation(), UnexpectedRowsAffectedError

### Community 54 - "custom_deep_validation.go"
Cohesion: 0.42
Nodes (9): reflect.Value, buildValidationErrorPathWithJSON(), validator.FieldError, isDirectValidationErrorForStruct(), validateSliceDeep(), validateStructDeep(), validateStructWithValidator(), validateValueDeep() (+1 more)

### Community 55 - "golang_ext_util_struct.go"
Cohesion: 0.29
Nodes (7): deepCompleteReflective(), DeepCompleteStruct(), GetStructName(), GetStructNamespaceDescription(), GetStructType(), T, StructString()

### Community 56 - ".buildAppComponents"
Cohesion: 0.18
Nodes (13): BuildAllocationPlanRESTController(), BuildAllocationPlanManagementAppService(), AllocationPlanManagementAppService, BuildPortfolioAllocationManagementAppService(), BuildAllocationPlanRepository(), BuildAllocationRepository(), BuildAssetRDBMSRepository(), BuildPortfolioAllocationRepository() (+5 more)

### Community 57 - "AllocationPlanRESTController"
Cohesion: 0.21
Nodes (4): allocationPlanCursorFromPayload(), github.com/gin-gonic/gin.HandlersChain, RESTRoute, AllocationPlanRESTController

### Community 58 - "custom_validation_error_handling.go"
Cohesion: 0.39
Nodes (8): go_pkg_github_com_benizzio_open_asset_allocator_infra_json, asValidationErrors(), formatErrorMessage(), formatValidationError(), FormatValidationErrorMessages(), validator.FieldError, validator.ValidationErrors, MapValidationErrorsToMessages()

### Community 59 - "CustomValidationErrorsBuilder"
Cohesion: 0.31
Nodes (5): go_pkg_github_com_go_playground_validator_v10, buildCustomValidationError(), CustomValidationErrorsBuilder, validator.FieldError, validator.ValidationErrors

### Community 60 - "golang_ext_util_slice.go"
Cohesion: 0.36
Nodes (7): cleanNilAllocations(), CleanNilPointersInSlice(), DereferenceSliceContent(), T, IsSlice(), ReverseSlice(), SliceContainsZeroValue()

### Community 61 - "PortfolioAllocationRDBMSRepository"
Cohesion: 0.26
Nodes (6): buildSelectedExternalAsset(), mapPortfolioAllocationRow(), mapPortfolioAllocationRows(), preparePortfolioAllocationTempInserts(), portfolioAllocationJoinedRowDTS, PortfolioAllocationRDBMSRepository

### Community 62 - "go_pkg_reflect"
Cohesion: 0.29
Nodes (4): go_pkg_github_com_go_playground_universal_translator, go_pkg_reflect, IsNilPointer(), TestFindStructFieldByNameOrJSONName()

### Community 63 - "parseAssetTextSearch"
Cohesion: 0.47
Nodes (4): appendAssetTextSearchPhrase(), appendAssetTextSearchTerms(), parseAssetTextSearch(), TestParseAssetTextSearch()

### Community 64 - "joinAny"
Cohesion: 0.50
Nodes (3): CustomSlice[T], T, joinAny()

## Knowledge Gaps
- **17 isolated node(s):** `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse`, `paginationCursorValidationDTS`, `contextKey` (+12 more)
  These have ≤1 connection - possible missing edges or undocumented components. (Counts symbols only; 102 node(s) total have ≤1 connection when file, concept and rationale nodes are included.)
- **10 thin communities (<3 nodes) omitted from report** — run `graphify query` to explore isolated nodes.

## Suggested Questions
_Questions this graph is uniquely positioned to answer:_

- **Why does `App` connect `App` to `Adapter`, `portfolio_observation_ownership_migration_integration_test.go`, `go_pkg_github_com_benizzio_open_asset_allocator_langext`, `GinServer`, `portfolio_allocation_controller.go`, `go_pkg_fmt`, `golang_ext_util_struct.go`, `.buildAppComponents`, `infra.go`?**
  _High betweenness centrality (0.049) - this node is a cross-community bridge._
- **Why does `Adapter` connect `Adapter` to `testing.T`, `App`, `GinServer`, `SQLTransactionalContext`, `go_pkg_fmt`?**
  _High betweenness centrality (0.043) - this node is a cross-community bridge._
- **Why does `Asset` connect `PropagateAsAppErrorWithNewMessage` to `AllocationPlan`, `Adapter`, `IsZeroValue`, `PortfolioAllocation`, `PortfolioAllocationRDBMSRepository`, `BuildCleanupFunctionBuilder`, `AssetDomService`, `asset_rest_model.go`, `parseAssetTextSearch`?**
  _High betweenness centrality (0.028) - this node is a cross-community bridge._
- **Are the 88 inferred relationships involving `deferCloseResponseBody()` (e.g. with `TestAllocationPlanReadsOrderByPlannedDateAndId()` and `TestGetAllocationPlans()`) actually correct?**
  _`deferCloseResponseBody()` has 88 INFERRED edges - model-reasoned connections that need verification._
- **What connects `AssetSearchQueryDTS`, `ExternalAssetSearchQueryDTS`, `ErrorResponse` to the rest of the system?**
  _17 weakly-connected nodes found - possible documentation gaps or missing edges._
- **Should `github.com/gin-gonic/gin.Context` be split into smaller, more focused modules?**
  _Cohesion score 0.1329268292682927 - nodes in this community are weakly interconnected._
- **Should `AllocationPlan` be split into smaller, more focused modules?**
  _Cohesion score 0.08549019607843138 - nodes in this community are weakly interconnected._