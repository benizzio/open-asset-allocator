package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AllocationRESTController struct {
	allocationPlanService *application.AllocationPlanService
}

func (controller *AllocationRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/allocation/plan",
			Handlers: gin.HandlersChain{controller.getAllocationPlans},
		},
	}
}

func (controller *AllocationRESTController) getAllocationPlans(context *gin.Context) {

	allocationPlans, err := controller.allocationPlanService.GetAllocationPlans()
	if infra.HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}

	var allocationPlansDTS = model.MapAllocationPlans(allocationPlans)

	//TODO change for JSON call
	context.IndentedJSON(http.StatusOK, allocationPlansDTS)
}

func BuildAllocationRESTController(allocationPlanService *application.AllocationPlanService) *AllocationRESTController {
	return &AllocationRESTController{allocationPlanService}
}
