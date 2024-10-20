package model

import "github.com/benizzio/open-asset-allocator/domain"

// ================================================
// TYPES
// ================================================

type PortfolioSliceDTS struct {
	AssetName        string `json:"assetName"`
	AssetTicker      string `json:"assetTicker"`
	Class            string `json:"class"`
	CashReserve      bool   `json:"cashReserve"`
	TotalMarketValue int    `json:"totalMarketValue"`
}

type PortfolioAtTimeDTS struct {
	TimeFrameTag     domain.TimeFrameTag `json:"timeFrameTag"`
	Slices           []PortfolioSliceDTS `json:"slices"`
	TotalMarketValue int                 `json:"totalMarketValue"`
}

type portfolioSlicesPerTimeFrameMap map[domain.TimeFrameTag][]PortfolioSliceDTS

func (aggregationMap portfolioSlicesPerTimeFrameMap) getOrBuild(timeFrameTag domain.TimeFrameTag) []PortfolioSliceDTS {
	var sliceAggregation = aggregationMap[timeFrameTag]
	if sliceAggregation == nil {
		sliceAggregation = make([]PortfolioSliceDTS, 0)
	}
	return sliceAggregation
}

func (aggregationMap portfolioSlicesPerTimeFrameMap) aggregate(
	timeFrameTag domain.TimeFrameTag,
	sliceDTS PortfolioSliceDTS,
) {
	var sliceAggregation = aggregationMap.getOrBuild(timeFrameTag)
	sliceAggregation = append(sliceAggregation, sliceDTS)
	aggregationMap[timeFrameTag] = sliceAggregation
}

func (aggregationMap portfolioSlicesPerTimeFrameMap) getAggregatedMarketValue(timeFrame domain.TimeFrameTag) int {
	var sliceAggregation = aggregationMap[timeFrame]
	var totalMarketValue = 0
	for _, slice := range sliceAggregation {
		totalMarketValue += slice.TotalMarketValue
	}
	return totalMarketValue
}

// ================================================
// MAPPING FUNCTIONS
// ================================================

func AggregateAndMapPortfolioHistory(portfolioHistory []domain.PortfolioSliceAtTime) []PortfolioAtTimeDTS {
	portfolioSlicesPerTimeFrame := aggregateHistoryAsDTSMap(portfolioHistory)
	aggregatedPortfoliohistory := buildHistoryDTS(portfolioSlicesPerTimeFrame)
	return aggregatedPortfoliohistory
}

func aggregateHistoryAsDTSMap(portfolioHistory []domain.PortfolioSliceAtTime) portfolioSlicesPerTimeFrameMap {

	var portfolioSlicesPerTimeFrame = make(portfolioSlicesPerTimeFrameMap)
	for _, portfolioSlice := range portfolioHistory {
		var aggregationTimeFrame = portfolioSlice.TimeFrameTag
		var sliceJSON = sliceAtTimeToSliceDTS(portfolioSlice)
		portfolioSlicesPerTimeFrame.aggregate(aggregationTimeFrame, sliceJSON)
	}

	return portfolioSlicesPerTimeFrame
}

func sliceAtTimeToSliceDTS(portfolioSlice domain.PortfolioSliceAtTime) PortfolioSliceDTS {
	return PortfolioSliceDTS{
		AssetName:        portfolioSlice.Asset.Name,
		AssetTicker:      portfolioSlice.Asset.Ticker,
		Class:            portfolioSlice.Class,
		CashReserve:      portfolioSlice.CashReserve,
		TotalMarketValue: portfolioSlice.TotalMarketValue,
	}
}

func buildHistoryDTS(portfolioSlicesPerTimeFrame portfolioSlicesPerTimeFrameMap) []PortfolioAtTimeDTS {

	var aggregatedPortfoliohistory = make([]PortfolioAtTimeDTS, 0)
	for timeFrameTag, slices := range portfolioSlicesPerTimeFrame {
		var totalMarketValue = portfolioSlicesPerTimeFrame.getAggregatedMarketValue(timeFrameTag)
		portfolioSnapshot := buildSnapshotDTS(timeFrameTag, slices, totalMarketValue)
		aggregatedPortfoliohistory = append(aggregatedPortfoliohistory, portfolioSnapshot)
	}

	return aggregatedPortfoliohistory
}

func buildSnapshotDTS(
	timeFrameTag domain.TimeFrameTag,
	slices []PortfolioSliceDTS,
	totalMarketValue int,
) PortfolioAtTimeDTS {
	return PortfolioAtTimeDTS{
		TimeFrameTag:     timeFrameTag,
		Slices:           slices,
		TotalMarketValue: totalMarketValue,
	}
}
