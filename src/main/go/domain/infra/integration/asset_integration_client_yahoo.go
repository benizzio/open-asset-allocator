package integration

import (
	"fmt"
	"net/url"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/util/http/httpclient"
)

const yahooFinanceSearchURL = "https://query2.finance.yahoo.com/v1/finance/search"
const yahooFinanceSearchQuotesCount = "5"
const yahooFinanceChartURL = "https://query2.finance.yahoo.com/v8/finance/chart/"

// YahooFinanceAssetIntegrationClient is an HTTP client for the Yahoo Finance API.
// Provides methods to query Yahoo Finance endpoints and return their responses as typed DTSs.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceAssetIntegrationClient struct {
}

// SearchAssets queries the Yahoo Finance search API for assets matching the given query value.
// Builds the request URL with hardcoded default parameters, executes the GET request,
// and decodes the JSON response into a YahooFinanceSearchResponseDTS.
//
// Parameters:
//   - queryValue: the search term to query (mapped to the "q" query parameter)
//
// Returns:
//   - *YahooFinanceSearchResponseDTS: the decoded search response containing matching quotes
//   - error: an AppError if the request fails, returns a non-200 status, or decoding fails
//
// Example:
//
//	var client = &YahooFinanceAssetIntegrationClient{}
//	response, err := client.SearchAssets("AAPL")
//	if err != nil {
//	    // handle error
//	}
//	for _, quote := range response.Quotes {
//	    fmt.Println(quote.Symbol, quote.Exchange)
//	}
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func (client *YahooFinanceAssetIntegrationClient) SearchAssets(
	queryValue string,
) (*YahooFinanceSearchResponseDTS, error) {

	var requestURL, err = buildSearchAssetsURL(queryValue)
	if err != nil {
		return nil, infra.PropagateAsAppError(err, client)
	}

	var searchResponse, getErr = httpclient.ExecuteGetJSON[YahooFinanceSearchResponseDTS](requestURL)
	if getErr != nil {
		return nil, infra.PropagateAsAppError(getErr, client)
	}

	return searchResponse, nil
}

// buildSearchAssetsURL constructs the full Yahoo Finance search URL with the given query value
// and hardcoded default parameters.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func buildSearchAssetsURL(queryValue string) (string, error) {

	var parsedURL, err = url.Parse(yahooFinanceSearchURL)
	if err != nil {
		return "", fmt.Errorf("error parsing Yahoo Finance search URL: %w", err)
	}

	var queryParams = parsedURL.Query()
	queryParams.Set("q", queryValue)
	queryParams.Set("enableFuzzyQuery", "false")
	queryParams.Set("quotesCount", yahooFinanceSearchQuotesCount)
	queryParams.Set("newsCount", "0")
	queryParams.Set("listsCount", "0")
	queryParams.Set("enableNavLinks", "false")
	queryParams.Set("enableResearchReports", "false")
	queryParams.Set("enableCulturalAssets", "false")
	queryParams.Set("enableCb", "false")

	parsedURL.RawQuery = queryParams.Encode()

	return parsedURL.String(), nil
}

// QuoteAssetLastClosePrice queries the Yahoo Finance chart API for the last close price data
// of the asset identified by the given ticker. Returns the full chart response DTS containing
// metadata, timestamps, and indicator data.
//
// Parameters:
//   - ticker: the asset ticker symbol (e.g., "AAPL")
//
// Returns:
//   - *YahooFinanceChartResponseDTS: the decoded chart response
//   - error: an AppError if the request fails, returns a non-200 status, or decoding fails
//
// Example:
//
//	var client = &YahooFinanceAssetIntegrationClient{}
//	response, err := client.QuoteAssetLastClosePrice("AAPL")
//	if err != nil {
//	    // handle error
//	}
//	var result = response.Chart.Result[0]
//	fmt.Println(result.Meta.Symbol, result.Meta.Currency)
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func (client *YahooFinanceAssetIntegrationClient) QuoteAssetLastClosePrice(
	ticker string,
) (*YahooFinanceChartResponseDTS, error) {

	var requestURL, err = buildQuoteAssetLastClosePriceURL(ticker)
	if err != nil {
		return nil, infra.PropagateAsAppError(err, client)
	}

	var chartResponse, getErr = httpclient.ExecuteGetJSON[YahooFinanceChartResponseDTS](requestURL)
	if getErr != nil {
		return nil, infra.PropagateAsAppError(getErr, client)
	}

	return chartResponse, nil
}

// buildQuoteAssetLastClosePriceURL constructs the full Yahoo Finance chart URL for the given ticker
// with hardcoded default parameters.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func buildQuoteAssetLastClosePriceURL(ticker string) (string, error) {

	var parsedURL, err = url.Parse(yahooFinanceChartURL + ticker)
	if err != nil {
		return "", fmt.Errorf("error parsing Yahoo Finance chart URL: %w", err)
	}

	var queryParams = parsedURL.Query()
	queryParams.Set("interval", "1d")
	queryParams.Set("events", "history")

	parsedURL.RawQuery = queryParams.Encode()

	return parsedURL.String(), nil
}
