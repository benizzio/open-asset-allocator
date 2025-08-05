package rest

const (
	// Deprecated: use static strings for paths and parameters
	portfolioPath    = "/api/portfolio"
	portfolioIdParam = "portfolioId"
	// Deprecated: use static strings for paths and parameters
	specificPortfolioPath = portfolioPath + "/:" + portfolioIdParam
	// Deprecated: use observationTimestampIdParam
	timeFrameTagParam                     = "timeFrameTag"
	observationTimestampIdParam           = "observationTimestampId"
	planIdParam                           = "planId"
	assetIdOrTickerParam                  = "assetIdOrTicker"
	getPortfolioIdErrorMessage            = "Error getting portfolioId url parameter"
	getObservationTimestampIdErrorMessage = "Error getting observationTimestampId url parameter"
	getPlanIdErrorMessage                 = "Error getting planId url parameter"
	bindPortfolioErrorMessage             = "Error binding portfolio from request body"
	bindPortfolioSnapshotErrorMessage     = "Error binding portfolio snapshot from request body"
)
