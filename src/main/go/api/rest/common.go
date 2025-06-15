package rest

const (
	portfolioPath         = "/api/portfolio"
	portfolioIdParam      = "portfolioId"
	specificPortfolioPath = portfolioPath + "/:" + portfolioIdParam
	// Deprecated: use observationTimestampIdParam
	timeFrameTagParam           = "timeFrameTag"
	observationTimestampIdParam = "observationTimestampId"
	planIdParam                 = "planId"
	getPortfolioIdErrorMessage  = "Error getting portfolioId url parameter"
	bindPortfolioErrorMessage   = "Error binding portfolio from request body"
)
