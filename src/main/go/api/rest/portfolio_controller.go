package rest

import (
	"net/http"

	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain/service"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/util"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
)

type PortfolioRESTController struct {
	portfolioDomService *service.PortfolioDomService
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
			Method:   http.MethodPost,
			Path:     "/api/portfolio",
			Handlers: gin.HandlersChain{controller.postPortfolio},
		},
		{
			Method:   http.MethodPut,
			Path:     "/api/portfolio",
			Handlers: gin.HandlersChain{controller.putPortfolio},
		},
	}
}

func (controller *PortfolioRESTController) getPortfolios(context *gin.Context) {

	portfolios, err := controller.portfolioDomService.GetPortfolios()
	if infra.HandleAPIError(context, "Error getting portfolios", err) {
		return
	}

	portfolioDTSs := model.MapToPortfolioDTSs(portfolios)

	context.JSON(http.StatusOK, portfolioDTSs)
}

func (controller *PortfolioRESTController) getPortfolio(context *gin.Context) {

	var portfolioIdParam = context.Param(portfolioIdParam)
	portfolioId, err := langext.ParseInt64(portfolioIdParam)
	if infra.HandleAPIError(context, getPortfolioIdErrorMessage, err) {
		return
	}

	portfolio, err := controller.portfolioDomService.GetPortfolio(portfolioId)
	if infra.HandleAPIError(context, "Error getting portfolio", err) {
		return
	}

	portfolioDTS := model.MapToPortfolioDTS(portfolio)

	context.JSON(http.StatusOK, portfolioDTS)
}

func (controller *PortfolioRESTController) postPortfolio(context *gin.Context) {

	var portfolioDTS model.PortfolioDTS
	valid, err := util.BindAndValidateJSONWithInvalidResponse(context, &portfolioDTS)
	if err != nil {
		infra.HandleAPIError(context, bindPortfolioErrorMessage, err)
		return
	}
	if !valid {
		return
	}

	var portfolio = model.MapToPortfolio(&portfolioDTS)
	persistedPortfolio, err := controller.portfolioDomService.PersistPortfolio(portfolio)
	if infra.HandleAPIError(context, "Error creating portfolio", err) {
		return
	}

	var responseBody = model.MapToPortfolioDTS(persistedPortfolio)
	context.JSON(http.StatusCreated, responseBody)
}

func (controller *PortfolioRESTController) putPortfolio(context *gin.Context) {

	var portfolioDTS model.PortfolioDTS
	valid, err := util.BindAndValidateJSONWithInvalidResponse(context, &portfolioDTS)
	if err != nil {
		infra.HandleAPIError(context, bindPortfolioErrorMessage, err)
		return
	}
	if !valid {
		return
	}

	if portfolioDTS.Id == nil || langext.IsZeroValue(portfolioDTS.Id) {
		var validationErrors = util.BuildCustomValidationErrorsBuilder().
			CustomValidationError(
				portfolioDTS,
				"Id",
				"required",
				"Portfolio ID is required for update",
				nil,
			).
			Build()
		util.RespondWithCustomValidationErrors(context, validationErrors, portfolioDTS)
		return
	}

	var portfolio = model.MapToPortfolio(&portfolioDTS)
	persistedPortfolio, err := controller.portfolioDomService.PersistPortfolio(portfolio)
	if infra.HandleAPIError(context, "Error updating portfolio", err) {
		return
	}

	var responseBody = model.MapToPortfolioDTS(persistedPortfolio)
	context.JSON(http.StatusOK, responseBody)
}

func BuildPortfolioRESTController(portfolioDomService *service.PortfolioDomService) *PortfolioRESTController {
	return &PortfolioRESTController{
		portfolioDomService,
	}
}
