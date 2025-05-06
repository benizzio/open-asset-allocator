package rest

const (
	portfolioPath              = "/api/portfolio"
	portfolioIdParam           = "portfolioId"
	specificPortfolioPath      = portfolioPath + "/:" + portfolioIdParam
	timeFrameTagParam          = "timeFrameTag"
	planIdParam                = "planId"
	getPortfolioIdErrorMessage = "Error getting portfolioId url parameter"
)
