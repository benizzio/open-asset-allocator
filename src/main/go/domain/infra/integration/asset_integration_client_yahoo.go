package integration

import (
	"fmt"
	"net/url"

	"github.com/benizzio/open-asset-allocator/infra"
	"github.com/benizzio/open-asset-allocator/infra/util/http/httpclient"
)

const yahooFinanceSearchURL = "https://query2.finance.yahoo.com/v1/finance/search"
const yahooFinanceSearchQuotesCount = "5"

// ================================================
// TYPES
// ================================================

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

// YahooFinanceAssetIntegrationClient is an HTTP client for the Yahoo Finance API.
// Provides methods to query Yahoo Finance endpoints and return their responses as typed DTSs.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type YahooFinanceAssetIntegrationClient struct {
}

// ================================================
// CLIENT FUNCTIONS
// ================================================

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

	resp, err := httpclient.ExecuteGet(requestURL)
	if err != nil {
		return nil, infra.PropagateAsAppError(err, client)
	}
	defer httpclient.CloseResponseBody(resp)

	var searchResponse, decodeErr = httpclient.DecodeJSONResponse[YahooFinanceSearchResponseDTS](resp)
	if decodeErr != nil {
		return nil, infra.PropagateAsAppError(decodeErr, client)
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
