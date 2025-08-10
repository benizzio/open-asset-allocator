package rest

import (
	"net/http"
	"strconv"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
)

type AllocationPlanRESTController struct {
	allocationPlanService *service.AllocationPlanDomService
}

func (controller *AllocationPlanRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/allocation-plan",
			Handlers: gin.HandlersChain{controller.getAllocationPlans},
		},
	}
}

func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)

	var planType = allocation.AssetAllocationPlan
	allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(
		portfolioId,
		&planType,
	)
	if infra.HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}

	var allocationPlansDTS = model.MapToAllocationPlanDTSs(allocationPlans)

	context.JSON(http.StatusOK, allocationPlansDTS)
}

func BuildAllocationPlanRESTController(allocationPlanService *service.AllocationPlanDomService) *AllocationPlanRESTController {
	return &AllocationPlanRESTController{allocationPlanService}
}
