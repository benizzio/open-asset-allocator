package rest

import (
	"fmt"
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type PortfolioRESTController struct {
	portfolioDomService *service.PortfolioDomService
}

func (controller *PortfolioRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{
			Method:   http.MethodGet,
			Path:     portfolioPath,
			Handlers: gin.HandlersChain{controller.getPortfolios},
		},
		{
			Method:   http.MethodGet,
			Path:     specificPortfolioPath,
			Handlers: gin.HandlersChain{controller.getPortfolio},
		},
		{
			Method:   http.MethodGet,
			Path:     specificPortfolioPath + "/history",
			Handlers: gin.HandlersChain{controller.getPortfolioAllocationHistory},
		},
		{
			Method:   http.MethodPost,
			Path:     portfolioPath,
			Handlers: gin.HandlersChain{controller.postPortfolio},
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
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
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
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	portfolioHistory, err := controller.portfolioDomService.GetPortfolioAllocationHistory(portfolioId)
	if infra.HandleAPIError(context, "Error getting portfolio history", err) {
		return
	}

	var aggregatedPortfoliohistoryDTS = model.AggregateAndMapPortfolioHistory(portfolioHistory)

	context.JSON(http.StatusOK, aggregatedPortfoliohistoryDTS)
}

func (controller *PortfolioRESTController) postPortfolio(context *gin.Context) {

	var portfolioDTS model.PortfolioDTS
	if err := context.ShouldBindJSON(&portfolioDTS); infra.HandleAPIError(context, "Error binding portfolio", err) {
		return
	}

	var portfolio = model.MapPortfolioDTS(&portfolioDTS)
	//TODO test code, remove
	//print the portfolioDTS
	fmt.Println(portfolio)

	//portfolio, err := controller.portfolioDomService.CreatePortfolio(portfolioDTS)
	//if infra.HandleAPIError(context, "Error creating portfolio", err) {
	//	return
	//}
	//

	//TODO test code, fix with return from persistence
	var id = 999
	portfolioDTS.Id = &id
	portfolioDTS.AllocationStructure = &model.AllocationStructureDTS{
		Hierarchy: []model.AllocationHierarchyLevelDTS{
			{
				Name:  "Assets",
				Field: "assetTicker",
			},
			{
				Name:  "Classes",
				Field: "class",
			},
		},
	}
	context.JSON(http.StatusCreated, portfolioDTS)
}

func BuildPortfolioRESTController(portfolioDomService *service.PortfolioDomService) *PortfolioRESTController {
	return &PortfolioRESTController{
		portfolioDomService,
	}
}
