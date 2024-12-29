package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AllocationPlanRESTController struct {
	allocationPlanService *application.AllocationPlanService
}

func (controller *AllocationPlanRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:portfolioId/allocation-plan",
			Handlers: gin.HandlersChain{controller.getAllocationPlans},
		},
	}
}

func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {

	var portfolioIdParam = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParam)

	var planType = allocation.AssetAllocationPlan
	allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(
		portfolioId,
		&planType,
	)
	if infra.HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}

	var allocationPlansDTS = model.MapAllocationPlans(allocationPlans)

	context.JSON(http.StatusOK, allocationPlansDTS)
}

func BuildAllocationPlanRESTController(allocationPlanService *application.AllocationPlanService) *AllocationPlanRESTController {
	return &AllocationPlanRESTController{allocationPlanService}
}
