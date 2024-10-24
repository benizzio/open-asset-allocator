package rest

import (
	"github.com/benizzio/open-asset-allocator/application"
	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

type PortfolioSliceJSON struct {
	AssetName        string `json:"assetName"`
	AssetTicker      string `json:"assetTicker"`
	Class            string `json:"class"`
	CashReserve      bool   `json:"cashReserve"`
	TotalMarketValue int    `json:"totalMarketValue"`
}

type PortfolioAtTimeJSON struct {
	TimeFrameTag string               `json:"timeFrameTag"`
	Slices       []PortfolioSliceJSON `json:"slices"`
}

type PortfolioRESTController struct {
	portfolioHistoryService *application.PortfolioHistoryService
}

func (controller *PortfolioRESTController) BuildRoutes() []infra.RESTRoute {
	return []infra.RESTRoute{
		{Method: http.MethodGet, Path: "/api/portfolio", Handlers: gin.HandlersChain{controller.getPortfolioHistory}},
	}
}

// TODO: properly handle errors and clean code
func (controller *PortfolioRESTController) getPortfolioHistory(context *gin.Context) {

	var portfolioHistory, err = controller.portfolioHistoryService.GetPortfolioHistory()
	glog.Infof("Portfolio history: %v", portfolioHistory)
	if err != nil {
		glog.Error("Error getting portfolio history: ", err)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Error getting portfolio history"})
	}

	var portfolioSlicesPerTimeFrameTag = make(map[string][]PortfolioSliceJSON)
	for _, portfolioSlice := range portfolioHistory {

		var sliceJSON = PortfolioSliceJSON{
			AssetName:        portfolioSlice.Asset.Name,
			AssetTicker:      portfolioSlice.Asset.Ticker,
			Class:            portfolioSlice.Class,
			CashReserve:      portfolioSlice.CashReserve,
			TotalMarketValue: portfolioSlice.TotalMarketValue,
		}

		var sliceAggregation = portfolioSlicesPerTimeFrameTag[portfolioSlice.TimeFrameTag]

		if sliceAggregation == nil {
			sliceAggregation = make([]PortfolioSliceJSON, 0)
		}

		sliceAggregation = append(sliceAggregation, sliceJSON)
		portfolioSlicesPerTimeFrameTag[portfolioSlice.TimeFrameTag] = sliceAggregation
	}

	var portfoliohistory = make([]PortfolioAtTimeJSON, 0)
	for timeFrameTag, slices := range portfolioSlicesPerTimeFrameTag {
		portfolioSnapshot := PortfolioAtTimeJSON{
			TimeFrameTag: timeFrameTag,
			Slices:       slices,
		}
		portfoliohistory = append(portfoliohistory, portfolioSnapshot)
	}

	context.IndentedJSON(http.StatusOK, portfoliohistory)
}

func BuildPortfolioRESTController(portfolioHistoryService *application.PortfolioHistoryService) *PortfolioRESTController {
	return &PortfolioRESTController{portfolioHistoryService}
}
