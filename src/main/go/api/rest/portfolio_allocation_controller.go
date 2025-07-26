package rest

import (
	"fmt"
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/util"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PortfolioAllocationRESTController struct {
	portfolioDomService                     *service.PortfolioDomService
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
			Handlers: gin.HandlersChain{controller.postPortfolioAllocationHistrory},
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
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var timeFrameTagParamValue = context.Query(timeFrameTagParam)

	var observationTimestampIdParamValue = context.Query(observationTimestampIdParam)
	var observationTimestampId int
	if !langext.IsZeroValue(observationTimestampIdParamValue) {
		observationTimestampId, err = strconv.Atoi(observationTimestampIdParamValue)
		if infra.HandleAPIError(context, getObservationTimestampIdErrorMessage, err) {
			return
		}
	}

	portfolioHistory, err := controller.getPortfolioAllocationHistoryUpstack(
		portfolioId,
		timeFrameTagParamValue,
		observationTimestampId,
	)
	if err != nil {
		var errorDetail string
		if !langext.IsZeroValue(timeFrameTagParamValue) {
			errorDetail = fmt.Sprintf(" for time frame tag %s", timeFrameTagParamValue)
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
	portfolioId int,
	timeFrameTagParamValue string,
	observationTimestampId int,
) ([]*domain.PortfolioAllocation, error) {

	var portfolioHistory []*domain.PortfolioAllocation
	var err error

	if !langext.IsZeroValue(timeFrameTagParamValue) {
		var timeFrameTag = domain.TimeFrameTag(timeFrameTagParamValue)
		portfolioHistory, err = controller.portfolioDomService.FindPortfolioAllocations(portfolioId, timeFrameTag)
	} else if !langext.IsZeroValue(observationTimestampId) {
		portfolioHistory, err = controller.portfolioDomService.FindPortfolioAllocationsByObservationTimestamp(
			portfolioId,
			observationTimestampId,
		)
	} else {
		portfolioHistory, err = controller.portfolioDomService.GetPortfolioAllocationHistory(portfolioId)
	}

	return portfolioHistory, err
}

func (controller *PortfolioAllocationRESTController) getAvailableHistoryObservations(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	availableTimestamps, err := controller.portfolioDomService.GetAvailableObservationTimestamps(portfolioId, 10)
	if infra.HandleAPIError(context, "Error getting available observation timestamps", err) {
		return
	}

	var availableTimestampsDTS = model.MapToObservationTimestampDTSs(availableTimestamps)

	context.JSON(http.StatusOK, availableTimestampsDTS)
}

func (controller *PortfolioAllocationRESTController) getAvailablePortfolioAllocationClasses(context *gin.Context) {

	portfolioIdParamValue := context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	availableClasses, err := controller.portfolioDomService.FindAvailablePortfolioAllocationClasses(portfolioId)
	if infra.HandleAPIError(context, "Error getting available portfolio allocation classes", err) {
		return
	}

	context.JSON(http.StatusOK, availableClasses)
}

// TODO confirm json format in front end
func (controller *PortfolioAllocationRESTController) postPortfolioAllocationHistrory(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var portfolioSnapshotDTS model.PortfolioSnapshotDTS
	valid, err := util.BindAndValidateJSONWithInvalidResponse(context, &portfolioSnapshotDTS)
	if err != nil {
		infra.HandleAPIError(context, bindPortfolioSnapshotErrorMessage, err)
		return
	}

	if !valid {
		return
	}

	var portfolioAllocations = model.MapToPortfolioAllocations(
		portfolioSnapshotDTS.Allocations,
		portfolioSnapshotDTS.ObservationTimestamp.Id,
	)

	err = controller.portfolioAllocationManagementAppService.MergePortfolioAllocations(
		portfolioId,
		portfolioAllocations,
	)

	if infra.HandleAPIError(context, "Error merging portfolio allocations", err) {
		return
	}

	context.Status(http.StatusNoContent)
}

func BuildPortfolioAllocationRESTController(
	portfolioDomService *service.PortfolioDomService,
	portfolioAllocationManagementAppService *application.PortfolioAllocationManagementAppService,
) *PortfolioAllocationRESTController {
	return &PortfolioAllocationRESTController{
		portfolioDomService,
		portfolioAllocationManagementAppService,
	}
}
