package integration

// asset_integration_client_yahoo_model.go contains the Data Transfer Structures (DTS)
// that map to Yahoo Finance API JSON responses. These types are used by the
// YahooFinanceAssetIntegrationClient to decode search and chart endpoint responses.
//
// Authored by: GitHub Copilot (claude-opus-4.6)

// YahooFinanceSearchQuoteDTS represents a single quote result from the Yahoo Finance search API.
// Fields map to the JSON response keys returned by the /v1/finance/search endpoint.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceSearchQuoteDTS struct {
	Symbol   string `json:"symbol"`
	Exchange string `json:"exchange"`
	LongName string `json:"longname"`
	ExchDisp string `json:"exchDisp"`
}

// YahooFinanceSearchResponseDTS represents the top-level response from the Yahoo Finance search API.
// Contains the list of quote results matching the search query.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceSearchResponseDTS struct {
	Quotes []YahooFinanceSearchQuoteDTS `json:"quotes"`
}

// YahooFinanceChartQuoteIndicatorDTS represents the OHLCV quote indicator data
// from the Yahoo Finance chart API response.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceChartQuoteIndicatorDTS struct {
	Close []float64 `json:"close"`
}

// YahooFinanceChartIndicatorsDTS represents the indicators section
// of the Yahoo Finance chart API response.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceChartIndicatorsDTS struct {
	Quote []YahooFinanceChartQuoteIndicatorDTS `json:"quote"`
}

// YahooFinanceChartMetaDTS represents the metadata section
// of a Yahoo Finance chart result, including symbol, exchange, and currency.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceChartMetaDTS struct {
	Symbol       string `json:"symbol"`
	ExchangeName string `json:"exchangeName"`
	Currency     string `json:"currency"`
}

// YahooFinanceChartResultDTS represents a single result entry
// from the Yahoo Finance chart API response.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceChartResultDTS struct {
	Meta       YahooFinanceChartMetaDTS       `json:"meta"`
	Timestamps []int64                        `json:"timestamp"`
	Indicators YahooFinanceChartIndicatorsDTS `json:"indicators"`
}

// YahooFinanceChartDTS represents the chart wrapper object
// in the Yahoo Finance chart API response.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceChartDTS struct {
	Result []YahooFinanceChartResultDTS `json:"result"`
}

// YahooFinanceChartResponseDTS represents the top-level response from the Yahoo Finance chart API.
// Contains chart results with metadata, timestamps, and indicator data for the requested asset.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceChartResponseDTS struct {
	Chart YahooFinanceChartDTS `json:"chart"`
}
