package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
			Handlers: gin.HandlersChain{controller.GetDivergenceAnalysisOptions},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/divergence/:" + timeFrameTagParam + "/allocation-plan/:" + planIdParam,
			Handlers: gin.HandlersChain{controller.GetDivergenceAnalysis},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/v2/portfolio/:" + portfolioIdParam + "/divergence/:" + observationTimestampIdParam + "/allocation-plan/:" + planIdParam,
			Handlers: gin.HandlersChain{controller.GetDivergenceAnalysisNew},
		},
	}
}

func (controller *DivergenceAnalysisRESTController) GetDivergenceAnalysisOptions(context *gin.Context) {

	portfolioIdParam := context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParam)
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

// Deprecated: use GetDivergenceAnalysisNew
func (controller *DivergenceAnalysisRESTController) GetDivergenceAnalysis(context *gin.Context) {

	portfolioIdParam := context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParam)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var timeFrameTagParam = domain.TimeFrameTag(context.Param(timeFrameTagParam))

	planIdParam := context.Param(planIdParam)
	planId, err := strconv.Atoi(planIdParam)
	if infra.HandleAPIError(context, "Error getting planId url parameter", err) {
		return
	}

	analysis, err := controller.portfolioDivergenceAnalysisService.GeneratePortfolioDivergenceAnalysis(
		portfolioId,
		timeFrameTagParam,
		planId,
	)
	if infra.HandleAPIError(context, "Error generating portfolio divergence analysis", err) {
		return
	}

	var analysisDTS = model.MapToDivergenceAnalysisDTS(analysis)

	context.JSON(http.StatusOK, analysisDTS)
}

// TODO: rename when GetDivergenceAnalysis is removed
func (controller *DivergenceAnalysisRESTController) GetDivergenceAnalysisNew(context *gin.Context) {

	portfolioIdParamValue := context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParamValue)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	var observationTimestampIdParamValue = context.Param(observationTimestampIdParam)
	observationTimestampId, err := strconv.Atoi(observationTimestampIdParamValue)
	if infra.HandleAPIError(context, "Error getting observationTimestampId url parameter", err) {
		return
	}

	planIdParamValue := context.Param(planIdParam)
	planId, err := strconv.Atoi(planIdParamValue)
	if infra.HandleAPIError(context, "Error getting planId url parameter", err) {
		return
	}

	analysis, err := controller.portfolioDivergenceAnalysisService.GeneratePortfolioDivergenceAnalysisNew(
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
	portfiolioAnalysisConfigService *application.PortfolioAnalysisConfigurationAppService,
	portfiolioAnalysisService *application.PortfolioDivergenceAnalysisAppService,
) *DivergenceAnalysisRESTController {
	return &DivergenceAnalysisRESTController{
		portfiolioAnalysisConfigService,
		portfiolioAnalysisService,
	}
}
