package inttest

import (
	"io"
	"net/http"
	"net/url"
	"strconv"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

const yahooFinanceQuoteRequestURI = "/v8/finance/chart/IAU?events=history&interval=1d"

// TestGetExternalAssetQuoteSuccess verifies quote responses using either a local asset ticker or
// ID and the exact external asset association persisted for that asset.
//
// Authored by: OpenCode
func TestGetExternalAssetQuoteSuccess(t *testing.T) {
	const assetTicker = "TEST:QUOTE-LOCAL"
	const externalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU","exchangeId":"PCX"}]}`
	var asset = insertTestAsset(t, assetTicker, "Quote test asset")
	var originalExternalData = capturePersistedAssetExternalData(t, asset.Id)
	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			asset.Id,
			originalExternalData,
		).Build(t),
	)

	err := inttestinfra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{"externalData": externalDataJSON, "id": asset.Id},
	)
	require.NoError(t, err)

	const chartResponseJSON = `{"chart":{"result":[{"meta":{"symbol":"IAU","exchangeName":"PCX","currency":"USD"},"timestamp":[1704067200,1704153600],"indicators":{"quote":[{"close":[20.15,null]}]}}]}}`
	const expectedQuoteJSON = `{
		"ticker":"IAU",
		"exchangeId":"PCX",
		"currency":"USD",
		"lastCloseQuote":"20.15",
		"lastCloseDate":"2024-01-01T00:00:00Z"
	}`
	var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)
	yahooFinanceMockServer.ExpectGet(yahooFinanceQuoteRequestURI).
		WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
		Return(chartResponseJSON)

	var statusCode, responseBody = getExternalAssetQuote(t, assetTicker, "YAHOO_FINANCE", "PCX", "IAU")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.JSONEq(t, expectedQuoteJSON, responseBody)

	yahooFinanceMockServer.ExpectGet(yahooFinanceQuoteRequestURI).
		WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
		Return(chartResponseJSON)
	statusCode, responseBody = getExternalAssetQuote(t, strconv.FormatInt(asset.Id, 10), "YAHOO_FINANCE", "PCX", "IAU")
	assert.Equal(t, http.StatusOK, statusCode)
	assert.JSONEq(t, expectedQuoteJSON, responseBody)
}

// TestGetExternalAssetQuoteNotFound verifies that the endpoint returns 404 without calling Yahoo
// when the local asset or exact persisted external asset association is missing.
//
// Authored by: OpenCode
func TestGetExternalAssetQuoteNotFound(t *testing.T) {
	_ = inttestinfra.SetupYahooFinanceMockTest(t)

	const associatedAssetTicker = "TEST:QUOTE-NOT-MATCHED"
	const associatedExternalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU","exchangeId":"PCX"}]}`
	var associatedAsset = insertTestAsset(t, associatedAssetTicker, "Quote association test asset")
	var originalExternalData = capturePersistedAssetExternalData(t, associatedAsset.Id)
	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			associatedAsset.Id,
			originalExternalData,
		).Build(t),
	)
	err := inttestinfra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{"externalData": associatedExternalDataJSON, "id": associatedAsset.Id},
	)
	require.NoError(t, err)

	t.Run("local asset does not exist", func(t *testing.T) {
		var statusCode, responseBody = getExternalAssetQuote(
			t,
			"TEST:QUOTE-MISSING",
			"YAHOO_FINANCE",
			"PCX",
			"IAU",
		)
		assert.Equal(t, http.StatusNotFound, statusCode)
		assert.JSONEq(t, `{"errorMessage":"Data not found","details":["External asset quote with identifier TEST:QUOTE-MISSING/YAHOO_FINANCE/PCX/IAU not found"]}`, responseBody)
	})

	t.Run("asset has no external data", func(t *testing.T) {
		var statusCode, responseBody = getExternalAssetQuote(
			t,
			"ARCA:BIL",
			"YAHOO_FINANCE",
			"PCX",
			"IAU",
		)
		assert.Equal(t, http.StatusNotFound, statusCode)
		assert.JSONEq(t, `{"errorMessage":"Data not found","details":["External asset quote with identifier ARCA:BIL/YAHOO_FINANCE/PCX/IAU not found"]}`, responseBody)
	})

	t.Run("external exchange identifier does not match", func(t *testing.T) {
		var statusCode, responseBody = getExternalAssetQuote(
			t,
			associatedAssetTicker,
			"YAHOO_FINANCE",
			"NMS",
			"IAU",
		)
		assert.Equal(t, http.StatusNotFound, statusCode)
		assert.JSONEq(t, `{"errorMessage":"Data not found","details":["External asset quote with identifier TEST:QUOTE-NOT-MATCHED/YAHOO_FINANCE/NMS/IAU not found"]}`, responseBody)
	})

	t.Run("external ticker does not match", func(t *testing.T) {
		var statusCode, responseBody = getExternalAssetQuote(
			t,
			associatedAssetTicker,
			"YAHOO_FINANCE",
			"PCX",
			"OTHER",
		)
		assert.Equal(t, http.StatusNotFound, statusCode)
		assert.JSONEq(t, `{"errorMessage":"Data not found","details":["External asset quote with identifier TEST:QUOTE-NOT-MATCHED/YAHOO_FINANCE/PCX/OTHER not found"]}`, responseBody)
	})

	// Asset ID 1 exists but has no matching external data.
	t.Run("numeric asset ID without association is not found", func(t *testing.T) {
		var statusCode, responseBody = getExternalAssetQuote(t, "1", "YAHOO_FINANCE", "PCX", "IAU")
		assert.Equal(t, http.StatusNotFound, statusCode)
		assert.JSONEq(t, `{"errorMessage":"Data not found","details":["External asset quote with identifier 1/YAHOO_FINANCE/PCX/IAU not found"]}`, responseBody)
	})
}

// TestGetExternalAssetQuoteRejectsInvalidSource verifies unsupported source values are rejected
// before any provider request is made.
//
// Authored by: OpenCode
func TestGetExternalAssetQuoteRejectsInvalidSource(t *testing.T) {
	_ = inttestinfra.SetupYahooFinanceMockTest(t)

	var statusCode, responseBody = getExternalAssetQuote(t, "ARCA:BIL", "UNSUPPORTED", "PCX", "IAU")
	assert.Equal(t, http.StatusBadRequest, statusCode)
	assert.JSONEq(t, `{"errorMessage":"Invalid AssetExternalSource UNSUPPORTED"}`, responseBody)
}

// TestGetExternalAssetQuoteRejectsMismatchedProviderTicker verifies Yahoo responses for a different
// ticker are treated as provider errors rather than returned as the requested asset's quote.
//
// Authored by: OpenCode
func TestGetExternalAssetQuoteRejectsMismatchedProviderTicker(t *testing.T) {
	const assetTicker = "TEST:QUOTE-MISMATCH"
	const externalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU","exchangeId":"PCX"}]}`
	var asset = insertTestAsset(t, assetTicker, "Quote mismatch test asset")
	var originalExternalData = capturePersistedAssetExternalData(t, asset.Id)
	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			asset.Id,
			originalExternalData,
		).Build(t),
	)

	err := inttestinfra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{"externalData": externalDataJSON, "id": asset.Id},
	)
	require.NoError(t, err)

	var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)
	yahooFinanceMockServer.ExpectGet(yahooFinanceQuoteRequestURI).
		WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
		Return(`{"chart":{"result":[{"meta":{"symbol":"WRONG","exchangeName":"PCX","currency":"USD"},"timestamp":[1704067200],"indicators":{"quote":[{"close":[20.15]}]}}]}}`)

	var statusCode, responseBody = getExternalAssetQuote(t, assetTicker, "YAHOO_FINANCE", "PCX", "IAU")
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.JSONEq(t, `{"errorMessage":"Internal server error"}`, responseBody)
}

// TestGetExternalAssetQuoteRejectsMismatchedProviderExchange verifies Yahoo responses for a
// different exchange are treated as provider errors rather than returned for the requested asset.
//
// Authored by: OpenCode
func TestGetExternalAssetQuoteRejectsMismatchedProviderExchange(t *testing.T) {
	const assetTicker = "TEST:QUOTE-MISMATCH-EXCHANGE"
	const externalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU","exchangeId":"PCX"}]}`
	var asset = insertTestAsset(t, assetTicker, "Quote exchange mismatch test asset")
	var originalExternalData = capturePersistedAssetExternalData(t, asset.Id)
	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			asset.Id,
			originalExternalData,
		).Build(t),
	)

	err := inttestinfra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{"externalData": externalDataJSON, "id": asset.Id},
	)
	require.NoError(t, err)

	var yahooFinanceMockServer = inttestinfra.SetupYahooFinanceMockTest(t)
	yahooFinanceMockServer.ExpectGet(yahooFinanceQuoteRequestURI).
		WithHeader("User-Agent", yahooFinanceExpectedUserAgent).
		Return(`{"chart":{"result":[{"meta":{"symbol":"IAU","exchangeName":"NMS","currency":"USD"},"timestamp":[1704067200],"indicators":{"quote":[{"close":[20.15]}]}}]}}`)

	var statusCode, responseBody = getExternalAssetQuote(t, assetTicker, "YAHOO_FINANCE", "PCX", "IAU")
	assert.Equal(t, http.StatusInternalServerError, statusCode)
	assert.JSONEq(t, `{"errorMessage":"Internal server error"}`, responseBody)
}

// getExternalAssetQuote requests a quote by local asset ID or ticker and external identifiers,
// returning its status code and response body.
//
// Authored by: OpenCode
func getExternalAssetQuote(
	t *testing.T,
	assetIdOrTicker string,
	source string,
	exchangeId string,
	externalTicker string,
) (int, string) {
	t.Helper()

	var requestURL = inttestinfra.TestAPIURLPrefix +
		"/asset/" + url.PathEscape(assetIdOrTicker) +
		"/external-asset/" + url.PathEscape(source) +
		"/" + url.PathEscape(exchangeId) +
		"/" + url.PathEscape(externalTicker) +
		"/quote"
	var response, err = http.Get(requestURL)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)

	var responseBody, readErr = io.ReadAll(response.Body)
	require.NoError(t, readErr)
	return response.StatusCode, string(responseBody)
}
