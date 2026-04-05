//go:build extinttest

// Package extinttest contains external integration tests that verify connectivity and
// contract compliance with live external services. These tests hit real external APIs
// and are gated behind the "extinttest" build tag to prevent them from running during
// standard test execution.
//
// Run with: go test -count=1 -tags=extinttest ./extinttest/...
//
// Authored by: GitHub Copilot (claude-opus-4.6)
package extinttest

import (
	"testing"

	"github.com/benizzio/open-asset-allocator/domain/infra/integration"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const iauTicker = "IAU"
const iauExpectedCurrency = "USD"

// TestSearchAssets_IAU verifies that the Yahoo Finance search API returns results
// for the IAU ticker (iShares Gold Trust) and that the response structure matches
// the expected DTS format.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func TestSearchAssets_IAU(t *testing.T) {

	var client = integration.BuildYahooFinanceAssetIntegrationClient()

	var searchResponse, err = client.SearchAssets(iauTicker)

	require.NoError(t, err, "SearchAssets should not return an error")
	require.NotNil(t, searchResponse, "SearchAssets response should not be nil")
	require.NotEmpty(t, searchResponse.Quotes, "SearchAssets should return at least one quote")

	var iauQuoteFound = false
	for _, quote := range searchResponse.Quotes {
		if quote.Symbol == iauTicker {
			iauQuoteFound = true
			assert.NotEmpty(t, quote.Exchange, "IAU quote should have a non-empty Exchange")
			assert.NotEmpty(t, quote.LongName, "IAU quote should have a non-empty LongName")
			break
		}
	}

	assert.True(t, iauQuoteFound, "SearchAssets results should contain a quote with Symbol IAU")
}

// TestQuoteAssetLastClosePrice_IAU verifies that the Yahoo Finance chart API returns
// valid last close price data for the IAU ticker (iShares Gold Trust).
// Asserts structural correctness and that the returned quote data has positive values.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func TestQuoteAssetLastClosePrice_IAU(t *testing.T) {

	var client = integration.BuildYahooFinanceAssetIntegrationClient()

	var chartResponse, err = client.QuoteAssetLastClosePrice(iauTicker)

	require.NoError(t, err, "QuoteAssetLastClosePrice should not return an error")
	require.NotNil(t, chartResponse, "QuoteAssetLastClosePrice response should not be nil")
	require.NotEmpty(t, chartResponse.Chart.Result, "Chart response should contain at least one result")

	var result = chartResponse.Chart.Result[0]

	// Verify metadata
	assert.Equal(t, iauTicker, result.Meta.Symbol, "Chart result Meta.Symbol should be IAU")
	assert.Equal(t, iauExpectedCurrency, result.Meta.Currency, "Chart result Meta.Currency should be USD")
	assert.NotEmpty(t, result.Meta.ExchangeName, "Chart result Meta.ExchangeName should not be empty")

	// Verify timestamps
	assert.NotEmpty(t, result.Timestamps, "Chart result should contain at least one timestamp")

	// Verify indicator data
	require.NotEmpty(t, result.Indicators.Quote, "Chart result should contain at least one quote indicator")
	require.NotEmpty(t, result.Indicators.Quote[0].Close, "Quote indicator should contain at least one close price")

	var lastClosePrice = result.Indicators.Quote[0].Close[len(result.Indicators.Quote[0].Close)-1]
	assert.Greater(t, lastClosePrice, float64(0), "Last close price should be greater than zero")
}
