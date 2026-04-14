package integration

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/util/http/httpclient"
)

const yahooFinanceSearchQuotesCount = "5"
const yahooFinanceUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) " +
	"AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"

// yahooFinanceDefaultOptions is the default set of RequestOption functions applied to all
// Yahoo Finance API requests. Currently sets the User-Agent header to a browser-like value,
// as Yahoo Finance's CDN blocks requests with non-browser user-agents.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
var yahooFinanceDefaultOptions = []httpclient.RequestOption{
	httpclient.WithHeader("User-Agent", yahooFinanceUserAgent),
}

// YahooFinanceAssetIntegrationClient is an HTTP client for the Yahoo Finance API.
// Provides methods to query Yahoo Finance endpoints and return their responses as typed DTSs.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceAssetIntegrationClient struct {
	config infra.YahooFinanceConfiguration
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
//	var client = BuildYahooFinanceAssetIntegrationClient(infra.YahooFinanceConfiguration{
//	    SearchURL: "https://query2.finance.yahoo.com/v1/finance/search",
//	    ChartURL:  "https://query2.finance.yahoo.com/v8/finance/chart/",
//	})
//	response, err := client.SearchAssets(context.Background(), "AAPL")
//	if err != nil {
//	    // handle error
//	}
//	for _, quote := range response.Quotes {
//	    fmt.Println(quote.Symbol, quote.Exchange)
//	}
//
// Co-authored by: OpenCode and benizzio
func (client *YahooFinanceAssetIntegrationClient) SearchAssets(
	searchContext context.Context,
	queryValue string,
) (*YahooFinanceSearchResponseDTS, error) {

	var requestURL, err = buildSearchAssetsURL(client.config.SearchURL, queryValue)
	if err != nil {
		return nil, infra.PropagateAsAppError(err, client)
	}

	var searchResponse, getErr = httpclient.ExecuteGetJSON[YahooFinanceSearchResponseDTS](
		searchContext,
		requestURL,
		yahooFinanceDefaultOptions...,
	)
	if getErr != nil {
		return nil, infra.PropagateAsAppError(getErr, client)
	}

	return searchResponse, nil
}

// buildSearchAssetsURL constructs the full Yahoo Finance search URL with the given query value
// and hardcoded default parameters.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func buildSearchAssetsURL(searchURL string, queryValue string) (string, error) {

	var parsedURL, err = url.Parse(searchURL)
	if err != nil {
		return "", fmt.Errorf("error parsing Yahoo Finance search URL %s: %w", searchURL, err)
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
//	var client = BuildYahooFinanceAssetIntegrationClient(infra.YahooFinanceConfiguration{
//	    SearchURL: "https://query2.finance.yahoo.com/v1/finance/search",
//	    ChartURL:  "https://query2.finance.yahoo.com/v8/finance/chart/",
//	})
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

	var requestURL, err = buildQuoteAssetLastClosePriceURL(client.config.ChartURL, ticker)
	if err != nil {
		return nil, infra.PropagateAsAppError(err, client)
	}

	var chartResponse, getErr = httpclient.ExecuteGetJSON[YahooFinanceChartResponseDTS](
		context.Background(),
		requestURL,
		yahooFinanceDefaultOptions...,
	)
	if getErr != nil {
		return nil, infra.PropagateAsAppError(getErr, client)
	}

	return chartResponse, nil
}

// buildQuoteAssetLastClosePriceURL constructs the full Yahoo Finance chart URL for the given ticker
// with hardcoded default parameters.
//
// Co-authored by: OpenCode and benizzio
func buildQuoteAssetLastClosePriceURL(chartURL string, ticker string) (string, error) {

	var parsedURL, err = url.Parse(chartURL)
	if err != nil {
		return "", fmt.Errorf("error parsing Yahoo Finance chart URL %s: %w", chartURL, err)
	}

	parsedURL.Path = strings.TrimRight(parsedURL.Path, "/") + "/" + url.PathEscape(ticker)

	var queryParams = parsedURL.Query()
	queryParams.Set("interval", "1d")
	queryParams.Set("events", "history")

	parsedURL.RawQuery = queryParams.Encode()

	return parsedURL.String(), nil
}

// BuildYahooFinanceAssetIntegrationClient creates a new YahooFinanceAssetIntegrationClient instance.
//
// Parameters:
//   - config: the Yahoo Finance endpoint configuration used by the client
//
// Returns:
//   - *YahooFinanceAssetIntegrationClient: the new client instance
//
// Example:
//
//	var client = integration.BuildYahooFinanceAssetIntegrationClient(infra.YahooFinanceConfiguration{
//	    SearchURL: "https://query2.finance.yahoo.com/v1/finance/search",
//	    ChartURL:  "https://query2.finance.yahoo.com/v8/finance/chart/",
//	})
//	response, err := client.SearchAssets(context.Background(), "AAPL")
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func BuildYahooFinanceAssetIntegrationClient(
	config infra.YahooFinanceConfiguration,
) *YahooFinanceAssetIntegrationClient {
	return &YahooFinanceAssetIntegrationClient{config: config}
}
