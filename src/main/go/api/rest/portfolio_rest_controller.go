package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const (
	portfolioIdParam  = "portfolioId"
	timeFrameTagParam = "timeFrameTag"
	planIdParam       = "planId"
)

type PortfolioRESTController struct {
	portfolioDomService       *service.PortfolioDomService
	portfiolioAnalysisService *application.PortfolioAnalysisAppService
}

func (controller *PortfolioRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio",
			Handlers: gin.HandlersChain{controller.getPortfolios},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam,
			Handlers: gin.HandlersChain{controller.getPortfolio},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/history",
			Handlers: gin.HandlersChain{controller.getPortfolioAllocationHistory},
		},
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/:" + portfolioIdParam + "/divergence/:" + timeFrameTagParam + "/allocation-plan/:" + planIdParam,
			Handlers: gin.HandlersChain{controller.GetDivercenceAnalysis},
		},
	}
}

func (controller *PortfolioRESTController) getPortfolios(context *gin.Context) {

	portfolios, err := controller.portfolioDomService.GetPortfolios()
	if infra.HandleAPIError(context, "Error getting portfolios", err) {
		return
	}

	portfolioDTSs := model.MapPortfolios(portfolios)

	context.JSON(http.StatusOK, portfolioDTSs)
}

func (controller *PortfolioRESTController) getPortfolio(context *gin.Context) {

	var portfolioIdParam = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParam)
	if infra.HandleAPIError(context, "Error getting portfolioId url parameter", err) {
		return
	}

	portfolio, err := controller.portfolioDomService.GetPortfolio(portfolioId)
	if infra.HandleAPIError(context, "Error getting portfolio", err) {
		return
	}

	portfolioDTS := model.MapPortfolio(portfolio)

	context.JSON(http.StatusOK, portfolioDTS)
}

func (controller *PortfolioRESTController) getPortfolioAllocationHistory(context *gin.Context) {

	var portfolioIdParam = context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParam)
	if infra.HandleAPIError(context, "Error getting portfolioId url parameter", err) {
		return
	}

	portfolioHistory, err := controller.portfolioDomService.GetPortfolioAllocationHistory(portfolioId)
	if infra.HandleAPIError(context, "Error getting portfolio history", err) {
		return
	}

	var aggregatedPortfoliohistoryDTS = model.AggregateAndMapPortfolioHistory(portfolioHistory)

	context.JSON(http.StatusOK, aggregatedPortfoliohistoryDTS)
}

func (controller *PortfolioRESTController) GetDivercenceAnalysis(context *gin.Context) {

	portfolioIdParam := context.Param(portfolioIdParam)
	portfolioId, err := strconv.Atoi(portfolioIdParam)
	if infra.HandleAPIError(context, "Error getting portfolioId url parameter", err) {
		return
	}

	var timeFrameTagParam = domain.TimeFrameTag(context.Param(timeFrameTagParam))

	planIdParam := context.Param(planIdParam)
	planId, err := strconv.Atoi(planIdParam)
	if infra.HandleAPIError(context, "Error getting planId url parameter", err) {
		return
	}

	analysis, err := controller.portfiolioAnalysisService.GeneratePortfolioDivergenceAnalysis(
		portfolioId,
		timeFrameTagParam,
		planId,
	)
	if infra.HandleAPIError(context, "Error generating portfolio divergence analysis", err) {
		return
	}

	context.JSON(http.StatusOK, analysis)
}

func BuildPortfolioRESTController(
	portfolioDomService *service.PortfolioDomService,
	portfiolioAnalysisService *application.PortfolioAnalysisAppService,
) *PortfolioRESTController {
	return &PortfolioRESTController{portfolioDomService, portfiolioAnalysisService}
}
