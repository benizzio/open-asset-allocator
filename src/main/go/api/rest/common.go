package rest

const (
	portfolioPath              = "/api/portfolio"
	portfolioIdParam           = "portfolioId"
	specificPortfolioPath      = portfolioPath + "/:" + portfolioIdParam
	timeFrameTagParam          = "timeFrameTag"
	planIdParam                = "planId"
	getPortfolioIdErrorMessage = "Error getting portfolioId url parameter"
	bindPortfolioErrorMessage  = "Error binding portfolio from request body"
)
