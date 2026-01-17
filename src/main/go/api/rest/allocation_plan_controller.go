package rest

import (
	"net/http"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
)

type AllocationPlanRESTController struct {
	allocationPlanService              *service.AllocationPlanDomService
	allocationPlanManagementAppService *application.AllocationPlanManagementAppService
}

func (controller *AllocationPlanRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/allocation-plan",
			Handlers: gin.HandlersChain{controller.getAllocationPlans},
		},
		{
			Method:   http.MethodPost,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/allocation-plan",
			Handlers: gin.HandlersChain{controller.postAssetAllocationPlan},
		},
	}
}

func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)

	var planType = allocation.AssetAllocationPlan
	allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(
		portfolioId,
		&planType,
	)
	if HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}

	var allocationPlansDTS = model.MapToAllocationPlanDTSs(allocationPlans)

	context.JSON(http.StatusOK, allocationPlansDTS)
}

func (controller *AllocationPlanRESTController) postAssetAllocationPlan(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var allocationPlanDTS model.AllocationPlanDTS
	valid, err := gininfra.BindAndValidateJSONWithInvalidResponse(context, &allocationPlanDTS)
	if HandleAPIError(context, "Error binding allocation plan", err) || !valid {
		return
	}

	var cleanedPlannedAllocations = langext.CleanNilPointersInSlice(allocationPlanDTS.Details)
	allocationPlanDTS.Details = cleanedPlannedAllocations

	allocationPlan, err := model.MapToAllocationPlan(&allocationPlanDTS, portfolioId, allocation.AssetAllocationPlan)
	if HandleAPIError(context, "Error mapping allocation plan", err) {
		return
	}

	err = controller.allocationPlanManagementAppService.PersistAllocationPlan(allocationPlan)
	if HandleAPIError(context, "Error persisting allocation plan", err) {
		return
	}

	context.Status(http.StatusNoContent)
}

func BuildAllocationPlanRESTController(
	allocationPlanService *service.AllocationPlanDomService,
	allocationPlanManagementAppService *application.AllocationPlanManagementAppService,
) *AllocationPlanRESTController {
	return &AllocationPlanRESTController{allocationPlanService, allocationPlanManagementAppService}
}
