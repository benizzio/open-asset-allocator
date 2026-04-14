package rest

const (
	portfolioIdParam                      = "portfolioId"
	observationTimestampIdParam           = "observationTimestampId"
	planIdParam                           = "planId"
	assetIdOrTickerParam                  = "assetIdOrTicker"
	externalAssetQueryParam               = "query"
	externalAssetSourceParam              = "externalAssetSource"
	getPortfolioIdErrorMessage            = "Error getting portfolioId url parameter"
	getObservationTimestampIdErrorMessage = "Error getting observationTimestampId url parameter"
	getPlanIdErrorMessage                 = "Error getting planId url parameter"
	bindPortfolioErrorMessage             = "Error binding portfolio from request body"
	bindPortfolioSnapshotErrorMessage     = "Error binding portfolio snapshot from request body"
	bindAssetErrorMessage                 = "Error binding asset from request body"
)
