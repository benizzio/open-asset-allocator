package rest

import (
	"fmt"
	"net/http"
	"slices"
	"strconv"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	"github.com/benizzio/open-asset-allocator/infra/validation"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
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
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/allocation-classes",
			Handlers: gin.HandlersChain{controller.getAvailablePortfolioAllocationClasses},
		},
	}
}

func (controller *PortfolioAllocationRESTController) getPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var observationTimestampIdParamValue = context.Query(observationTimestampIdParam)
	var observationTimestampId int64
	if !langext.IsZeroValue(observationTimestampIdParamValue) {
		observationTimestampId, err = langext.ParseInt64(observationTimestampIdParamValue)
		if infra.HandleAPIError(context, getObservationTimestampIdErrorMessage, err) {
			return
		}
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
		infra.HandleAPIError(
			context,
			fmt.Sprintf("Error getting portfolio history %s", errorDetail),
			err,
		)
		return
	}

	var aggregatedPortfoliohistoryDTS = model.AggregateAndMapToPortfolioHistoryDTSs(portfolioHistory)

	context.JSON(http.StatusOK, aggregatedPortfoliohistoryDTS)
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

func (controller *PortfolioAllocationRESTController) getAvailableHistoryObservations(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	availableTimestamps, err := controller.portfolioAllocationDomService.GetAvailableObservationTimestamps(
		portfolioId,
		10,
	)
	if infra.HandleAPIError(context, "Error getting available observation timestamps", err) {
		return
	}

	var availableTimestampsDTS = model.MapToPortfolioObservationTimestampDTSs(availableTimestamps)

	context.JSON(http.StatusOK, availableTimestampsDTS)
}

// Deprecated: Use PortfolioRESTController.getAvailablePortfolioAllocationClasses instead.
// This endpoint only returns classes from portfolio_allocation_fact table.
// The new endpoint returns classes from both portfolio_allocation_fact and planned_allocation tables.
//
// TODO make a unified version of this that get from portfolio AND from allocation plan
func (controller *PortfolioAllocationRESTController) getAvailablePortfolioAllocationClasses(context *gin.Context) {

	portfolioIdParamValue := context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	availableClasses, err := controller.portfolioAllocationDomService.FindAvailablePortfolioAllocationClasses(portfolioId)
	if infra.HandleAPIError(context, "Error getting available portfolio allocation classes", err) {
		return
	}

	context.JSON(http.StatusOK, availableClasses)
}

func (controller *PortfolioAllocationRESTController) postPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var portfolioSnapshotDTS model.PortfolioSnapshotDTS
	valid, err := gininfra.BindAndValidateJSONWithInvalidResponse(context, &portfolioSnapshotDTS)
	if infra.HandleAPIError(context, bindPortfolioSnapshotErrorMessage, err) || !valid {
		return
	}

	if !controller.validateCleanPortfolioAllocationHistory(context, &portfolioSnapshotDTS) {
		return
	}

	var portfolioAllocations = model.MapToPortfolioAllocations(
		portfolioSnapshotDTS.Allocations,
		int64(portfolioSnapshotDTS.ObservationTimestamp.Id),
	)

	var observationTimestamp = model.MapToPortfolioObservationTimestamp(portfolioSnapshotDTS.ObservationTimestamp)

	err = controller.portfolioAllocationManagementAppService.MergePortfolioAllocations(
		portfolioId,
		observationTimestamp,
		portfolioAllocations,
	)

	if infra.HandleAPIError(context, "Error merging portfolio allocations", err) {
		return
	}

	context.Status(http.StatusNoContent)
}

func (controller *PortfolioAllocationRESTController) validateCleanPortfolioAllocationHistory(
	context *gin.Context,
	portfolioSnapshotDTS *model.PortfolioSnapshotDTS,
) bool {

	var cleanAllocations = slices.DeleteFunc(
		portfolioSnapshotDTS.Allocations,
		func(allocation *model.PortfolioAllocationDTS) bool {
			if allocation == nil {
				return true
			}
			return false
		},
	)
	portfolioSnapshotDTS.Allocations = cleanAllocations

	var validationErrorsBuilder = validation.BuildCustomValidationErrorsBuilder()
	for index, allocation := range portfolioSnapshotDTS.Allocations {

		var isAssetIdentified = langext.IsZeroValue(allocation.AssetId) &&
			langext.IsZeroValue(allocation.AssetTicker) && langext.IsZeroValue(allocation.AssetName)
		if isAssetIdentified {
			validationErrorsBuilder.CustomValidationErrorFullNamespace(
				portfolioSnapshotDTS,
				"allocations["+strconv.Itoa(index)+"].assetId",
				"custom",
				"if assetId is not provided, assetTicker or assetName must be provided",
				nil,
			)
		}
	}

	var validationErrors = validationErrorsBuilder.Build()
	if len(validationErrors) > 0 {
		gininfra.RespondWithCustomValidationErrors(context, validationErrors, portfolioSnapshotDTS)
		return false
	}

	return true
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
