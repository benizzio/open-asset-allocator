package rest

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	urlinfra "github.com/benizzio/open-asset-allocator/infra/util/url"
	"github.com/benizzio/open-asset-allocator/infra/validation"
	"github.com/benizzio/open-asset-allocator/langext"
)

type PortfolioAllocationRESTController struct {
	portfolioAllocationDomService           *service.PortfolioAllocationDomService
	portfolioAllocationManagementAppService *application.PortfolioAllocationManagementAppService
}

func (controller *PortfolioAllocationRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/history",
			Handlers: gin.HandlersChain{controller.getPortfolioAllocationHistory},
		},
		{
			Method:   http.MethodPost,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/history",
			Handlers: gin.HandlersChain{controller.postPortfolioAllocationHistory},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/history/observation",
			Handlers: gin.HandlersChain{controller.getAvailableHistoryObservations},
		},
	}
}

// getPortfolioAllocationHistory returns the legacy history list or an opt-in cursor page.
// Co-authored by: benizzio, GPT-6.1 Sol and OpenCode
func (controller *PortfolioAllocationRESTController) getPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	response, err := controller.buildPortfolioAllocationHistoryResponse(context, portfolioId)
	if context.IsAborted() || gininfra.HandleAPIError(context, "Error getting portfolio history", err) {
		return
	}
	context.JSON(http.StatusOK, response)
}

// buildPortfolioAllocationHistoryResponse validates read options and selects complete history or a cursor page.
// Authored by: GPT-6.1 Sol
func (controller *PortfolioAllocationRESTController) buildPortfolioAllocationHistoryResponse(
	context *gin.Context,
	portfolioId int64,
) (any, error) {
	var observationTimestampId, err = readHistoryObservationTimestampId(context)
	if err != nil {
		return nil, err
	}
	var pageCursor, isPaginated = readObservationPaginationCursor(context, portfolioId)
	if context.IsAborted() {
		return nil, nil
	}
	if !isPaginated {
		return controller.buildPortfolioAllocationHistoryListResponse(portfolioId, observationTimestampId)
	}
	if urlinfra.HasQueryParameter(context.Request.URL.Query(), observationTimestampIdParam) {
		respondWithPaginationCursorValidationError(context, "cannot be combined with observationTimestampId")
		return nil, nil
	}
	return controller.buildPaginatedPortfolioAllocationHistoryResponse(portfolioId, pageCursor)
}

// readHistoryObservationTimestampId parses the optional legacy single-observation selector.
// Authored by: GPT-6.1 Sol
func readHistoryObservationTimestampId(context *gin.Context) (int64, error) {
	var value = context.Query(observationTimestampIdParam)
	if value == "" {
		return 0, nil
	}
	var id, err = langext.ParseInt64(value)
	if err != nil {
		return 0, infra.PropagateAsAppErrorWithNewMessage(
			err, getObservationTimestampIdErrorMessage, readHistoryObservationTimestampId,
		)
	}
	return id, nil
}

// buildPortfolioAllocationHistoryListResponse loads and aggregates the legacy history selection.
// Authored by: GPT-6.1 Sol
func (controller *PortfolioAllocationRESTController) buildPortfolioAllocationHistoryListResponse(
	portfolioId int64,
	observationTimestampId int64,
) ([]*model.PortfolioSnapshotDTS, error) {
	var history, err = controller.getPortfolioAllocationHistoryUpstack(portfolioId, observationTimestampId)
	if err != nil {
		if observationTimestampId != 0 {
			return nil, infra.PropagateAsAppErrorWithNewMessage(
				err,
				"getting portfolio history for observation id "+strconv.FormatInt(observationTimestampId, 10),
				controller,
			)
		}
		return nil, err
	}
	return model.AggregateAndMapToPortfolioHistoryDTSs(history), nil
}

// buildPaginatedPortfolioAllocationHistoryResponse maps complete snapshots and their continuation boundary.
// Co-authored by: GPT-6.1 Sol and OpenCode
func (controller *PortfolioAllocationRESTController) buildPaginatedPortfolioAllocationHistoryResponse(
	portfolioId int64,
	cursor *domain.ObservationPageCursor,
) (*model.PaginatedResponse[*model.PortfolioSnapshotDTS], error) {
	var page, err = controller.portfolioAllocationDomService.
		GetPortfolioAllocationsByObservationPage(portfolioId, cursor)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "getting paginated portfolio history", controller)
	}
	nextCursor, err := encodeObservationPaginationCursor(portfolioId, page.NextCursor)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "encoding next portfolio-history cursor", controller)
	}
	return &model.PaginatedResponse[*model.PortfolioSnapshotDTS]{
		Items:      model.AggregateAndMapToPortfolioHistoryDTSs(page.Allocations),
		NextCursor: nextCursor,
	}, nil
}

func (controller *PortfolioAllocationRESTController) getPortfolioAllocationHistoryUpstack(
	portfolioId int64,
	observationTimestampId int64,
) ([]*domain.PortfolioAllocation, error) {

	var portfolioHistory []*domain.PortfolioAllocation
	var err error

	if !langext.IsZeroValue(observationTimestampId) {
		portfolioHistory, err = controller.portfolioAllocationDomService.FindPortfolioAllocationsByObservationTimestamp(
			portfolioId,
			observationTimestampId,
		)
	} else {
		portfolioHistory, err = controller.portfolioAllocationDomService.GetPortfolioAllocationHistory(portfolioId)
	}

	return portfolioHistory, err
}

// getAvailableHistoryObservations returns the legacy summary list or an opt-in cursor page.
// Co-authored by: GPT-6.1 Sol and OpenCode
func (controller *PortfolioAllocationRESTController) getAvailableHistoryObservations(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}
	response, err := controller.buildAvailableHistoryObservationsResponse(context, portfolioId)
	if context.IsAborted() || gininfra.HandleAPIError(context, "Error getting available observation timestamps", err) {
		return
	}
	context.JSON(http.StatusOK, response)
}

// buildAvailableHistoryObservationsResponse selects legacy observation summaries or an opt-in page.
// Authored by: GPT-6.1 Sol
func (controller *PortfolioAllocationRESTController) buildAvailableHistoryObservationsResponse(
	context *gin.Context,
	portfolioId int64,
) (any, error) {
	var pageCursor, isPaginated = readObservationPaginationCursor(context, portfolioId)
	if context.IsAborted() {
		return nil, nil
	}
	if isPaginated {
		return controller.buildPaginatedHistoryObservationsResponse(portfolioId, pageCursor)
	}
	return controller.buildHistoryObservationsListResponse(portfolioId)
}

// buildHistoryObservationsListResponse loads and maps the legacy ten most recent observation headers.
// Authored by: GPT-6.1 Sol
func (controller *PortfolioAllocationRESTController) buildHistoryObservationsListResponse(
	portfolioId int64,
) ([]*model.PortfolioObservationTimestampDTS, error) {
	var availableTimestamps, err = controller.portfolioAllocationDomService.GetAvailableObservationTimestamps(
		portfolioId,
		10,
	)
	if err != nil {
		return nil, err
	}
	return model.MapToPortfolioObservationTimestampDTSs(availableTimestamps), nil
}

// buildPaginatedHistoryObservationsResponse maps observation headers and their continuation boundary.
// Authored by: GPT-6.1 Sol
func (controller *PortfolioAllocationRESTController) buildPaginatedHistoryObservationsResponse(
	portfolioId int64,
	cursor *domain.ObservationPageCursor,
) (*model.PaginatedResponse[*model.PortfolioObservationTimestampDTS], error) {
	var page, err = controller.portfolioAllocationDomService.
		GetAvailableObservationTimestampsPage(portfolioId, cursor)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "getting paginated observation timestamps", controller)
	}
	nextCursor, err := encodeObservationPaginationCursor(portfolioId, page.NextCursor)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "encoding next observation cursor", controller)
	}
	return &model.PaginatedResponse[*model.PortfolioObservationTimestampDTS]{
		Items:      model.MapToPortfolioObservationTimestampDTSs(page.Items),
		NextCursor: nextCursor,
	}, nil
}

// bindAndValidatePortfolioSnapshot binds a snapshot and responds to malformed RFC3339 timestamps.
// Authored by: OpenCode
func bindAndValidatePortfolioSnapshot(
	context *gin.Context,
	portfolioSnapshotDTS *model.PortfolioSnapshotDTS,
) (bool, error) {
	var valid, err = gininfra.BindAndValidateJSONWithInvalidResponse(context, portfolioSnapshotDTS)
	if err != nil {
		var timestampParseError *time.ParseError
		if errors.As(err, &timestampParseError) {
			var validationErrorsBuilder = validation.BuildCustomValidationErrorsBuilder()
			validationErrorsBuilder.CustomValidationErrorFullNamespace(
				portfolioSnapshotDTS,
				"observationTimestamp.timestamp",
				"custom",
				"must be a valid RFC3339 timestamp",
				timestampParseError.Value,
			)
			gininfra.RespondWithCustomValidationErrors(
				context,
				validationErrorsBuilder.Build(),
				portfolioSnapshotDTS,
			)
			return false, nil
		}
	}

	return valid, err
}

// postPortfolioAllocationHistory validates and persists a portfolio snapshot with the submitted observation instant.
//
// Co-authored by: OpenCode and benizzio
func (controller *PortfolioAllocationRESTController) postPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var portfolioSnapshotDTS model.PortfolioSnapshotDTS
	valid, err := bindAndValidatePortfolioSnapshot(context, &portfolioSnapshotDTS)
	if gininfra.HandleAPIError(context, bindPortfolioSnapshotErrorMessage, err) || !valid {
		return
	}

	if !controller.validateCleanPortfolioAllocationHistory(context, &portfolioSnapshotDTS) {
		return
	}

	var portfolioAllocations = model.MapToPortfolioAllocations(
		portfolioSnapshotDTS.Allocations,
		int64(portfolioSnapshotDTS.ObservationTimestamp.Id),
	)

	var observationTimestamp = model.MapToPortfolioObservationTimestamp(
		portfolioSnapshotDTS.ObservationTimestamp,
	)

	err = controller.portfolioAllocationManagementAppService.MergePortfolioAllocations(
		portfolioId,
		observationTimestamp,
		portfolioAllocations,
	)

	if gininfra.HandleAPIError(context, "Error merging portfolio allocations", err) {
		return
	}

	context.Status(http.StatusNoContent)
}

// validateCleanPortfolioAllocationHistory rejects zero timestamps, whitespace-only tags, and invalid allocation fields.
//
// Co-authored by: GPT-6 Luna, GPT-6.1 Sol, OpenCode and benizzio
func (controller *PortfolioAllocationRESTController) validateCleanPortfolioAllocationHistory(
	context *gin.Context,
	portfolioSnapshotDTS *model.PortfolioSnapshotDTS,
) bool {

	cleanNilAllocations(portfolioSnapshotDTS)

	var validationErrorsBuilder = validation.BuildCustomValidationErrorsBuilder()
	if portfolioSnapshotDTS.ObservationTimestamp.Timestamp.IsZero() {
		validationErrorsBuilder.CustomValidationErrorFullNamespace(
			portfolioSnapshotDTS,
			"observationTimestamp.timestamp",
			"custom",
			"must not be the zero timestamp",
			portfolioSnapshotDTS.ObservationTimestamp.Timestamp,
		)
	}
	if !langext.IsZeroValue(portfolioSnapshotDTS.ObservationTimestamp.TimeTag) &&
		langext.IsEmptyString(portfolioSnapshotDTS.ObservationTimestamp.TimeTag) {
		validationErrorsBuilder.CustomValidationErrorFullNamespace(
			portfolioSnapshotDTS,
			"observationTimestamp.timeTag",
			"custom",
			"must not contain only whitespace",
			portfolioSnapshotDTS.ObservationTimestamp.TimeTag,
		)
	}

	for index, allocation := range portfolioSnapshotDTS.Allocations {
		validateAssetIdentification(validationErrorsBuilder, allocation, portfolioSnapshotDTS, index)
	}

	var validationErrors = validationErrorsBuilder.Build()
	if len(validationErrors) > 0 {
		gininfra.RespondWithCustomValidationErrors(context, validationErrors, portfolioSnapshotDTS)
		return false
	}

	return true
}

func cleanNilAllocations(portfolioSnapshotDTS *model.PortfolioSnapshotDTS) {
	var cleanAllocations = langext.CleanNilPointersInSlice(portfolioSnapshotDTS.Allocations)
	portfolioSnapshotDTS.Allocations = cleanAllocations
}

func validateAssetIdentification(
	validationErrorsBuilder *validation.CustomValidationErrorsBuilder,
	allocation *model.PortfolioAllocationDTS,
	portfolioSnapshotDTS *model.PortfolioSnapshotDTS,
	allocationIndex int,
) {

	var isAssetIdentified = (allocation.AssetId == nil || langext.IsZeroValue(*allocation.AssetId)) &&
		langext.IsZeroValue(allocation.AssetTicker) && langext.IsZeroValue(allocation.AssetName)

	if isAssetIdentified {
		validationErrorsBuilder.CustomValidationErrorFullNamespace(
			portfolioSnapshotDTS,
			"allocations["+strconv.Itoa(allocationIndex)+"].assetId",
			"custom",
			"if assetId is not provided, assetTicker and assetName must be provided",
			nil,
		)
	}
}

func BuildPortfolioAllocationRESTController(
	portfolioAllocationDomService *service.PortfolioAllocationDomService,
	portfolioAllocationManagementAppService *application.PortfolioAllocationManagementAppService,
) *PortfolioAllocationRESTController {
	return &PortfolioAllocationRESTController{
		portfolioAllocationDomService,
		portfolioAllocationManagementAppService,
	}
}
