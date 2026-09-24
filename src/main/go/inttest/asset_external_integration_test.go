package inttest

import (
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
)

const yahooFinanceExpectedUserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36"
const yahooFinanceSearchRequestURI = "/v1/finance/search?enableCb=false&enableCulturalAssets=false&enableFuzzyQuery=false&enableNavLinks=false&enableResearchReports=false&listsCount=0&newsCount=0&q=IAU&quotesCount=10"

// TestGetExternalAssetsSuccess verifies successful GET /api/external-asset responses using the
// shared Yahoo Finance mock server.
//
// Authored by: GitHub Copilot
func TestGetExternalAssetsSuccess(t *testing.T) {

	t.Run("ReturnsExternalAssets", func(t *testing.T) {
		var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)

		yahooFinanceMockServer.ExpectGet(yahooFinanceSearchRequestURI).
			WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
			Return(`
				{
					"quotes": [
						{
							"symbol": "IAU",
							"exchange": "PCX",
							"longname": "iShares Gold Trust",
							"exchDisp": "NYSEArca"
						}
					]
				}
			`)

		var statusCode, responseBody = getExternalAssets(t, "query=IAU")

		assert.Equal(t, http.StatusOK, statusCode)
		assert.JSONEq(t, `
			[
				{
					"source": "YAHOO_FINANCE",
					"ticker": "IAU",
					"exchangeId": "PCX",
					"name": "iShares Gold Trust",
					"exchangeName": "NYSEArca"
				}
			]
		`, responseBody)
	})

	t.Run("ReturnsEmptyArrayWhenYahooReturnsNoQuotes", func(t *testing.T) {
		var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)

		yahooFinanceMockServer.ExpectGet(yahooFinanceSearchRequestURI).
			WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
			Return(`{"quotes": []}`)

		var statusCode, responseBody = getExternalAssets(t, "query=IAU")

		assert.Equal(t, http.StatusOK, statusCode)
		assert.JSONEq(t, `[]`, responseBody)
	})

	// The provider may return more results than requested; the API still caps the response.
	// Authored by: OpenCode
	t.Run("ReturnsOnlyTheFirstTenExternalAssets", func(t *testing.T) {
		var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)

		yahooFinanceMockServer.ExpectGet(yahooFinanceSearchRequestURI).
			WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
			Return(`{"quotes":[
				{"symbol":"A01","exchange":"PCX"},{"symbol":"A02","exchange":"PCX"},
				{"symbol":"A03","exchange":"PCX"},{"symbol":"A04","exchange":"PCX"},
				{"symbol":"A05","exchange":"PCX"},{"symbol":"A06","exchange":"PCX"},
				{"symbol":"A07","exchange":"PCX"},{"symbol":"A08","exchange":"PCX"},
				{"symbol":"A09","exchange":"PCX"},{"symbol":"A10","exchange":"PCX"},
				{"symbol":"A11","exchange":"PCX"}
			]}`)

		var statusCode, responseBody = getExternalAssets(t, "query=IAU")
		assert.Equal(t, http.StatusOK, statusCode)
		assert.JSONEq(t, `[
			{"source":"YAHOO_FINANCE","ticker":"A01","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A02","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A03","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A04","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A05","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A06","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A07","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A08","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A09","exchangeId":"PCX"},
			{"source":"YAHOO_FINANCE","ticker":"A10","exchangeId":"PCX"}
		]`, responseBody)
	})
}

// TestGetExternalAssetsFailure verifies failed GET /api/external-asset responses caused by
// external service errors and request validation errors.
//
// Authored by: GitHub Copilot
func TestGetExternalAssetsFailure(t *testing.T) {

	t.Run("ReturnsInternalServerErrorWhenYahooReturnsNon200", func(t *testing.T) {
		var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)

		yahooFinanceMockServer.ExpectGet(yahooFinanceSearchRequestURI).
			WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
			ReturnCode(http.StatusTooManyRequests).
			Return(`{"error":"rate limited"}`)

		var statusCode, responseBody = getExternalAssets(t, "query=IAU")

		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.JSONEq(t, `
			{
				"errorMessage": "Internal server error"
			}
		`, responseBody)
	})

	t.Run("ReturnsInternalServerErrorWhenYahooReturnsInvalidJSON", func(t *testing.T) {
		var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)

		yahooFinanceMockServer.ExpectGet(yahooFinanceSearchRequestURI).
			WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
			Return(`{"quotes": [`)

		var statusCode, responseBody = getExternalAssets(t, "query=IAU")

		assert.Equal(t, http.StatusInternalServerError, statusCode)
		assert.JSONEq(t, `
			{
				"errorMessage": "Internal server error"
			}
		`, responseBody)
	})

	t.Run("ValidationFailsWhenQueryIsMissing", func(t *testing.T) {
		_ = inttestinfra.SetupYahooFinanceMockTest(t)

		var statusCode, responseBody = getExternalAssets(t, "")

		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.JSONEq(t, `
			{
				"errorMessage": "Validation failed",
				"details": [
					"Field 'query' failed validation: is required"
				]
			}
		`, responseBody)
	})

	t.Run("ValidationFailsWhenQueryIsEmpty", func(t *testing.T) {
		_ = inttestinfra.SetupYahooFinanceMockTest(t)

		var statusCode, responseBody = getExternalAssets(t, "query=")

		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.JSONEq(t, `
			{
				"errorMessage": "Validation failed",
				"details": [
					"Field 'query' failed validation: is required"
				]
			}
		`, responseBody)
	})

	t.Run("ValidationFailsWhenQueryExceedsMaxLength", func(t *testing.T) {
		_ = inttestinfra.SetupYahooFinanceMockTest(t)

		var oversizedQuery = strings.Repeat("A", 101)
		var statusCode, responseBody = getExternalAssets(t, "query="+oversizedQuery)

		assert.Equal(t, http.StatusBadRequest, statusCode)
		assert.JSONEq(t, `
			{
				"errorMessage": "Validation failed",
				"details": [
					"Field 'query' failed validation: must not exceed 100"
				]
			}
		`, responseBody)
	})
}
