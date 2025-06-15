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
			Path:     specificPortfolioPath + "/divergence/options",
			Handlers: gin.HandlersChain{controller.GetDivergenceAnalysisOptions},
		},
		{
			Method:   http.MethodGet,
			Path:     specificPortfolioPath + "/divergence/:" + timeFrameTagParam + "/allocation-plan/:" + planIdParam,
			Handlers: gin.HandlersChain{controller.GetDivergenceAnalysis},
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

// TODO: refactor to use domain.PortfolioObservationTimestamp instead of TimeFrameTag
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

func BuildDivergenceAnalysisRESTController(
	portfiolioAnalysisConfigService *application.PortfolioAnalysisConfigurationAppService,
	portfiolioAnalysisService *application.PortfolioDivergenceAnalysisAppService,
) *DivergenceAnalysisRESTController {
	return &DivergenceAnalysisRESTController{
		portfiolioAnalysisConfigService,
		portfiolioAnalysisService,
	}
}
