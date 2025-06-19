package rest

const (
	// Deprecated: use static strings for paths and parameters
	portfolioPath    = "/api/portfolio"
	portfolioIdParam = "portfolioId"
	// Deprecated: use static strings for paths and parameters
	specificPortfolioPath = portfolioPath + "/:" + portfolioIdParam
	// Deprecated: use observationTimestampIdParam
	timeFrameTagParam           = "timeFrameTag"
	observationTimestampIdParam = "observationTimestampId"
	planIdParam                 = "planId"
	getPortfolioIdErrorMessage  = "Error getting portfolioId url parameter"
	bindPortfolioErrorMessage   = "Error binding portfolio from request body"
)
