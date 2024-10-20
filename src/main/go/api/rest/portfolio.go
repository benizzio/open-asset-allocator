package rest

import (
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PortfolioRESTController struct {
	portfolioHistoryService *application.PortfolioHistoryService
}

func (controller *PortfolioRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     "/api/portfolio/history",
			Handlers: gin.HandlersChain{controller.getPortfolioHistory},
		},
	}
}

func (controller *PortfolioRESTController) getPortfolioHistory(context *gin.Context) {

	portfolioHistory, err := controller.portfolioHistoryService.GetPortfolioHistory()
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
