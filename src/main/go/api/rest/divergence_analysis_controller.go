package rest

import (
	"net/http"
	"strconv"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
)

type DivergenceAnalysisRESTController struct {
	portfolioAnalysisConfigService     *application.PortfolioAnalysisConfigurationAppService
	portfolioDivergenceAnalysisService *application.PortfolioDivergenceAnalysisAppService
}

func (controller *DivergenceAnalysisRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/divergence/options",
			Handlers: gin.HandlersChain{controller.getDivergenceAnalysisOptions},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/v2/portfolio/:" + portfolioIdParam + "/divergence/:" + observationTimestampIdParam + "/allocation-plan/:" + planIdParam,
			Handlers: gin.HandlersChain{controller.GetDivergenceAnalysis},
		},
	}
}

func (controller *DivergenceAnalysisRESTController) getDivergenceAnalysisOptions(context *gin.Context) {

	portfolioIdParam := context.Param(portfolioIdParam)
	portfolioId, err := strconv.ParseInt(portfolioIdParam, 10, 64)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	analysisOptions, err := controller.portfolioAnalysisConfigService.GetDivergenceAnalysisOptions(portfolioId)
	if infra.HandleAPIError(context, "Error getting divergence analysis options", err) {
		return
	}

	var analysisOptionsDTS = model.MapToAnalysisOptionsDTS(analysisOptions)

	context.JSON(http.StatusOK, analysisOptionsDTS)
}

func (controller *DivergenceAnalysisRESTController) GetDivergenceAnalysis(context *gin.Context) {

	portfolioIdParamValue := context.Param(portfolioIdParam)
	portfolioId, err := strconv.ParseInt(portfolioIdParamValue, 10, 64)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var observationTimestampIdParamValue = context.Param(observationTimestampIdParam)
	observationTimestampId, err := strconv.ParseInt(observationTimestampIdParamValue, 10, 64)
	if infra.HandleAPIError(context, getObservationTimestampIdErrorMessage, err) {
		return
	}

	planIdParamValue := context.Param(planIdParam)
	planId, err := strconv.ParseInt(planIdParamValue, 10, 64)
	if infra.HandleAPIError(context, getPlanIdErrorMessage, err) {
		return
	}

	analysis, err := controller.portfolioDivergenceAnalysisService.GeneratePortfolioDivergenceAnalysis(
		portfolioId,
		observationTimestampId,
		planId,
	)
	if infra.HandleAPIError(context, "Error generating portfolio divergence analysis", err) {
		return
	}

	var analysisDTS = model.MapToDivergenceAnalysisDTS(analysis)

	context.JSON(http.StatusOK, analysisDTS)
}

func BuildDivergenceAnalysisRESTController(
	portfolioAnalysisConfigService *application.PortfolioAnalysisConfigurationAppService,
	portfolioAnalysisService *application.PortfolioDivergenceAnalysisAppService,
) *DivergenceAnalysisRESTController {
	return &DivergenceAnalysisRESTController{
		portfolioAnalysisConfigService,
		portfolioAnalysisService,
	}
}
