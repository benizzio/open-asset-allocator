package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
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

	planTypeRef, success := getQueryParams(context)
	if !success {
		return
	}

	allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(planTypeRef)
	if infra.HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}

	var allocationPlansDTS = model.MapAllocationPlans(allocationPlans)

	//TODO change for JSON call
	context.IndentedJSON(http.StatusOK, allocationPlansDTS)
}

func getQueryParams(context *gin.Context) (*allocation.PlanType, bool) {
	var planTypeParam = context.Query("planType")
	planTypeRef, success := getPlanTypeRef(context, planTypeParam)
	return planTypeRef, success
}

func getPlanTypeRef(context *gin.Context, planTypeParam string) (
	*allocation.PlanType,
	bool,
) {
	var planTypeRef *allocation.PlanType
	if planTypeParam != "" {
		planType, err := allocation.GetPlanType(planTypeParam)
		if infra.HandleAPIError(context, "Error converting query parameter", err) {
			return nil, false
		}
		planTypeRef = &planType
	}
	return planTypeRef, true
}

func BuildAllocationRESTController(allocationPlanService *application.AllocationPlanService) *AllocationRESTController {
	return &AllocationRESTController{allocationPlanService}
}
