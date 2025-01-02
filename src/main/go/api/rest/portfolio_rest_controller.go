package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PortfolioRESTController struct {
	portfolioHistoryService *application.PortfolioHistoryService
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
			Path:     "/api/portfolio/:portfolioId",
			Handlers: gin.HandlersChain{controller.getPortfolio},
		},
		{
			Method: http.MethodGet,
			//TODO change to /api/portfolio/:portfolioId/history and handle parameter
			Path:     "/api/portfolio/history",
			Handlers: gin.HandlersChain{controller.getPortfolioAllocationHistory},
		},
	}
}

func (controller *PortfolioRESTController) getPortfolios(context *gin.Context) {

	portfolios, err := controller.portfolioHistoryService.GetPortfolios()
	if infra.HandleAPIError(context, "Error getting portfolios", err) {
		return
	}

	portfolioDTSs := model.MapPortfolios(portfolios)

	context.IndentedJSON(http.StatusOK, portfolioDTSs)
}

func (controller *PortfolioRESTController) getPortfolio(context *gin.Context) {

	var portfolioIdParam = context.Param("portfolioId")
	portfolioId, err := strconv.Atoi(portfolioIdParam)
	if infra.HandleAPIError(context, "Error getting portfolioId url parameter", err) {
		return
	}

	portfolio, err := controller.portfolioHistoryService.GetPortfolio(portfolioId)
	if infra.HandleAPIError(context, "Error getting portfolio", err) {
		return
	}

	portfolioDTS := model.MapPortfolio(portfolio)

	context.IndentedJSON(http.StatusOK, portfolioDTS)
}

// TODO has to be selected on the context of a portfolio
func (controller *PortfolioRESTController) getPortfolioAllocationHistory(context *gin.Context) {

	portfolioHistory, err := controller.portfolioHistoryService.GetPortfolioAllocationHistory()
	if infra.HandleAPIError(context, "Error getting portfolio history", err) {
		return
	}

	var aggregatedPortfoliohistoryDTS = model.AggregateAndMapPortfolioHistory(portfolioHistory)

	//TODO change for JSON call
	context.IndentedJSON(http.StatusOK, aggregatedPortfoliohistoryDTS)
}

func BuildPortfolioRESTController(portfolioHistoryService *application.PortfolioHistoryService) *PortfolioRESTController {
	return &PortfolioRESTController{portfolioHistoryService}
}
