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
// Authored by: benizzio and OpenCode
func (controller *AllocationPlanRESTController) getAllocationPlans(context *gin.Context) {

	var portfolioIdParamValue = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParamValue)
	if gininfra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var planType = allocation.AssetAllocationPlan
	pageCursor, isPaginated := readPaginationCursor(
		context,
		portfolioId,
		allocationPlanCollection,
		planType.String(),
	)
	if isPaginated {
		if context.IsAborted() {
			return
		}
		allocationPlansPage, pageError := controller.allocationPlanService.GetAllocationPlansPage(
			portfolioId,
			planType,
			allocationPlanCursorFromPayload(pageCursor),
		)
		if gininfra.HandleAPIError(context, "Error getting paginated allocation plans", pageError) {
			return
		}
		nextCursor, cursorError := encodeAllocationPlanPaginationCursor(
			portfolioId,
			planType.String(),
			allocationPlansPage.NextCursor,
		)
		if gininfra.HandleAPIError(context, "Error encoding next allocation-plan cursor", cursorError) {
			return
		}
		context.JSON(http.StatusOK, model.PaginatedResponse[*model.AllocationPlanDTS]{
			Items:      model.MapToAllocationPlanDTSs(allocationPlansPage.Items),
			NextCursor: nextCursor,
		})
		return
	}
	allocationPlans, err := controller.allocationPlanService.GetAllocationPlans(
		portfolioId,
		&planType,
	)
	if gininfra.HandleAPIError(context, "Error getting allocation plans", err) {
		return
	}

	var allocationPlansDTS = model.MapToAllocationPlanDTSs(allocationPlans)

	context.JSON(http.StatusOK, allocationPlansDTS)
}

// allocationPlanCursorFromPayload converts a validated REST cursor into the allocation-plan domain boundary.
// Authored by: OpenCode
func allocationPlanCursorFromPayload(cursor *parsedPaginationCursor) *domain.AllocationPlanPageCursor {
	if cursor == nil {
		return nil
	}
	return &domain.AllocationPlanPageCursor{PlannedDate: cursor.PlannedDate, Id: cursor.Id}
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
