package rest

import (
	"errors"
	"fmt"
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
// Authored by: OpenCode
func (controller *PortfolioAllocationRESTController) getPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var observationTimestampIdParamValue = context.Query(observationTimestampIdParam)
	var observationTimestampId int64
	if !langext.IsZeroValue(observationTimestampIdParamValue) {
		observationTimestampId, err = langext.ParseInt64(observationTimestampIdParamValue)
		if gininfra.HandleAPIError(context, getObservationTimestampIdErrorMessage, err) {
			return
		}
	}

	if controller.respondWithPaginatedPortfolioAllocationHistory(context, portfolioId) {
		return
	}

	portfolioHistory, err := controller.getPortfolioAllocationHistoryUpstack(
		portfolioId,
		observationTimestampId,
	)
	if err != nil {
		var errorDetail string
		if !langext.IsZeroValue(observationTimestampId) {
			errorDetail = fmt.Sprintf(" for observation id %d", observationTimestampId)
		}
		gininfra.HandleAPIError(
			context,
			fmt.Sprintf("Error getting portfolio history %s", errorDetail),
			err,
		)
		return
	}

	var aggregatedPortfoliohistoryDTS = model.AggregateAndMapToPortfolioHistoryDTSs(portfolioHistory)

	context.JSON(http.StatusOK, aggregatedPortfoliohistoryDTS)
}

// respondWithPaginatedPortfolioAllocationHistory handles opt-in keyset paging and reports whether it was selected.
// Authored by: OpenCode
func (controller *PortfolioAllocationRESTController) respondWithPaginatedPortfolioAllocationHistory(
	context *gin.Context,
	portfolioId int64,
) bool {
	pageCursor, isPaginated := readPaginationCursor(
		context,
		portfolioId,
		observationCursorCollection,
		"",
	)
	if !isPaginated {
		return false
	}
	if context.IsAborted() {
		return true
	}
	if hasQueryParameter(context.Request.URL.Query(), observationTimestampIdParam) {
		respondWithPaginationCursorValidationError(context, "cannot be combined with observationTimestampId")
		return true
	}
	portfolioHistoryPage, err := controller.portfolioAllocationDomService.
		GetPortfolioAllocationHistoryPage(portfolioId, observationCursorFromPayload(pageCursor))
	if gininfra.HandleAPIError(context, "Error getting paginated portfolio history", err) {
		return true
	}
	nextCursor, err := encodeObservationPaginationCursor(portfolioId, portfolioHistoryPage.NextCursor)
	if gininfra.HandleAPIError(context, "Error encoding next portfolio-history cursor", err) {
		return true
	}
	var aggregatedPortfolioHistory = model.AggregateAndMapToPortfolioHistoryDTSs(
		portfolioHistoryPage.Allocations,
	)
	context.JSON(http.StatusOK, model.PaginatedResponse[*model.PortfolioSnapshotDTS]{
		Items:      aggregatedPortfolioHistory,
		NextCursor: nextCursor,
	})
	return true
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

// observationCursorFromPayload converts a validated REST cursor into the observation domain boundary.
// Authored by: OpenCode
func observationCursorFromPayload(cursor *parsedPaginationCursor) *domain.ObservationPageCursor {
	if cursor == nil {
		return nil
	}
	return &domain.ObservationPageCursor{Timestamp: cursor.Timestamp, Id: cursor.Id}
}

// getAvailableHistoryObservations returns the legacy summary list or an opt-in cursor page.
// Authored by: OpenCode
func (controller *PortfolioAllocationRESTController) getAvailableHistoryObservations(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}
	pageCursor, isPaginated := readPaginationCursor(
		context,
		portfolioId,
		observationCursorCollection,
		"",
	)
	if isPaginated {
		if context.IsAborted() {
			return
		}
		availableTimestampsPage, pageError := controller.portfolioAllocationDomService.
			GetAvailableObservationTimestampsPage(portfolioId, observationCursorFromPayload(pageCursor))
		if gininfra.HandleAPIError(context, "Error getting paginated observation timestamps", pageError) {
			return
		}
		nextCursor, cursorError := encodeObservationPaginationCursor(portfolioId, availableTimestampsPage.NextCursor)
		if gininfra.HandleAPIError(context, "Error encoding next observation cursor", cursorError) {
			return
		}
		context.JSON(http.StatusOK, model.PaginatedResponse[*model.PortfolioObservationTimestampDTS]{
			Items:      model.MapToPortfolioObservationTimestampDTSs(availableTimestampsPage.Items),
			NextCursor: nextCursor,
		})
		return
	}

	availableTimestamps, err := controller.portfolioAllocationDomService.GetAvailableObservationTimestamps(
		portfolioId,
		10,
	)
	if gininfra.HandleAPIError(context, "Error getting available observation timestamps", err) {
		return
	}

	var availableTimestampsDTS = model.MapToPortfolioObservationTimestampDTSs(availableTimestamps)

	context.JSON(http.StatusOK, availableTimestampsDTS)
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
