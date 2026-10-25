package rest

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	gininfra "github.com/benizzio/open-asset-allocator/infra/gin"
	"github.com/benizzio/open-asset-allocator/infra/validation"
	"github.com/benizzio/open-asset-allocator/langext"
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

// getAllocationPlans returns the legacy plan list or an opt-in complete-plan cursor page.
// Co-authored by: GPT-6.1 Sol, benizzio and OpenCode
func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	response, err := controller.buildAllocationPlansResponse(context, portfolioId)
	if context.IsAborted() || gininfra.HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}
	context.JSON(http.StatusOK, response)
}

// buildAllocationPlansResponse selects the requested read mode without writing a successful response.
// Authored by: GPT-6.1 Sol
func (controller *AllocationPlanRESTController) buildAllocationPlansResponse(
	context *gin.Context,
	portfolioId int64,
) (any, error) {
	var planType = allocation.AssetAllocationPlan
	var pageCursor, isPaginated = readAllocationPlanPaginationCursor(
		context,
		portfolioId,
		planType.String(),
	)
	if context.IsAborted() {
		return nil, nil
	}
	if isPaginated {
		return controller.buildPaginatedAllocationPlansResponse(portfolioId, planType, pageCursor)
	}
	return controller.buildAllocationPlansListResponse(portfolioId, planType)
}

// buildAllocationPlansListResponse loads and maps the legacy unpaginated plan list.
// Authored by: GPT-6.1 Sol
func (controller *AllocationPlanRESTController) buildAllocationPlansListResponse(
	portfolioId int64,
	planType allocation.PlanType,
) ([]*model.AllocationPlanDTS, error) {
	var plans, err = controller.allocationPlanService.GetAllocationPlans(portfolioId, &planType)
	if err != nil {
		return nil, err
	}
	return model.MapToAllocationPlanDTSs(plans), nil
}

// buildPaginatedAllocationPlansResponse loads complete plans and maps their continuation boundary.
// Authored by: GPT-6.1 Sol
func (controller *AllocationPlanRESTController) buildPaginatedAllocationPlansResponse(
	portfolioId int64,
	planType allocation.PlanType,
	cursor *domain.AllocationPlanPageCursor,
) (*model.PaginatedResponse[*model.AllocationPlanDTS], error) {
	var page, err = controller.allocationPlanService.GetAllocationPlansPage(
		portfolioId, planType, cursor,
	)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "getting paginated allocation plans", controller)
	}
	nextCursor, err := encodeAllocationPlanPaginationCursor(portfolioId, planType.String(), page.NextCursor)
	if err != nil {
		return nil, infra.PropagateAsAppErrorWithNewMessage(err, "encoding next allocation-plan cursor", controller)
	}
	return &model.PaginatedResponse[*model.AllocationPlanDTS]{
		Items:      model.MapToAllocationPlanDTSs(page.Items),
		NextCursor: nextCursor,
	}, nil
}

// postAssetAllocationPlan validates the required plan date and persists the supplied plan details.
//
// Co-authored by: OpenCode and benizzio
func (controller *AllocationPlanRESTController) postAssetAllocationPlan(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var allocationPlanDTS model.AllocationPlanDTS
	valid, err := gininfra.BindAndValidateJSONWithInvalidResponse(context, &allocationPlanDTS)
	if gininfra.HandleAPIError(context, "Error binding allocation plan", err) || !valid {
		return
	}
	if !validateAllocationPlanPlannedDate(context, &allocationPlanDTS) {
		return
	}

	var cleanedPlannedAllocations = langext.CleanNilPointersInSlice(allocationPlanDTS.Details)
	allocationPlanDTS.Details = cleanedPlannedAllocations

	allocationPlan, err := model.MapToAllocationPlan(&allocationPlanDTS, portfolioId, allocation.AssetAllocationPlan)
	if gininfra.HandleAPIError(context, "Error mapping allocation plan", err) {
		return
	}

	err = controller.allocationPlanManagementAppService.PersistAllocationPlan(allocationPlan)
	if gininfra.HandleAPIError(context, "Error persisting allocation plan", err) {
		return
	}

	context.Status(http.StatusNoContent)
}

// validateAllocationPlanPlannedDate rejects dates that cannot represent a planned calendar date.
//
// Authored by: OpenCode
func validateAllocationPlanPlannedDate(
	context *gin.Context,
	allocationPlanDTS *model.AllocationPlanDTS,
) bool {
	plannedDate, err := time.Parse(time.DateOnly, string(allocationPlanDTS.PlannedDate))
	if err == nil && !plannedDate.IsZero() && plannedDate.Year() >= 1 {
		return true
	}

	var validationErrorsBuilder = validation.BuildCustomValidationErrorsBuilder()
	validationErrorsBuilder.CustomValidationError(
		allocationPlanDTS,
		"PlannedDate",
		"custom",
		"must be a valid non-zero YYYY-MM-DD date",
		allocationPlanDTS.PlannedDate,
	)
	var validationErrors = validationErrorsBuilder.Build()
	gininfra.RespondWithCustomValidationErrors(context, validationErrors, allocationPlanDTS)
	return false
}

func BuildAllocationPlanRESTController(
	allocationPlanService *service.AllocationPlanDomService,
	allocationPlanManagementAppService *application.AllocationPlanManagementAppService,
) *AllocationPlanRESTController {
	return &AllocationPlanRESTController{allocationPlanService, allocationPlanManagementAppService}
}
