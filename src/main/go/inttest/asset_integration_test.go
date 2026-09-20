package inttest

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"

	"github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// TestGetKnownAssets tests the GET /api/asset endpoint to retrieve all known assets.
//
// Authored by: GitHub Copilot
func TestGetKnownAssets(t *testing.T) {

	response, err := http.Get(infra.TestAPIURLPrefix + "/asset")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"id": 1,
				"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
				"ticker": "ARCA:BIL"
			},
			{
				"id": 6,
				"name": "iShares Msci Brazil ETF",
				"ticker": "ARCA:EWZ"
			},
			{
				"id": 7,
				"name": "SPDR S&P 500 ETF Trust",
				"ticker": "ARCA:SPY"
			},
			{
				"id": 2,
				"name": "iShares 0-5 Year TIPS Bond ETF",
				"ticker": "ARCA:STIP"
			},
			{
				"id": 3,
				"name": "iShares 7-10 Year Treasury Bond ETF",
				"ticker": "NasdaqGM:IEF"
			},
			{
				"id": 5,
				"name": "iShares Short Treasury Bond ETF",
				"ticker": "NasdaqGM:SHV"
			},
			{
				"id": 4,
				"name": "iShares 20+ Year Treasury Bond ETF",
				"ticker": "NasdaqGM:TLT"
			}
		]
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetKnownAssetsIncludesPersistedExternalData verifies that GET /api/asset returns the
// complete external data payload persisted for an asset.
//
// Authored by: OpenCode
func TestGetKnownAssetsIncludesPersistedExternalData(t *testing.T) {
	const persistedExternalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"BIL-PERSISTED","exchangeId":"PCX-PERSISTED"},{"source":"YAHOO_FINANCE","ticker":"BIL-SECOND","exchangeId":"SECOND"}]}`

	var assetId int64 = 1
	var originalExternalData = capturePersistedAssetExternalData(t, assetId)
	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			assetId,
			originalExternalData,
		).Build(t),
	)

	err := infra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{
			"externalData": persistedExternalDataJSON,
			"id":           assetId,
		},
	)
	assert.NoError(t, err)

	response, err := http.Get(infra.TestAPIURLPrefix + "/asset")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualAssets []json.RawMessage
	err = json.Unmarshal(body, &actualAssets)
	assert.NoError(t, err)

	var persistedAssetJSON json.RawMessage
	for _, assetJSON := range actualAssets {
		var assetIdentifier struct {
			Id int64 `json:"id"`
		}
		err = json.Unmarshal(assetJSON, &assetIdentifier)
		assert.NoError(t, err)
		if assetIdentifier.Id == assetId {
			persistedAssetJSON = assetJSON
			break
		}
	}

	assert.NotEmpty(t, persistedAssetJSON)
	var expectedAssetJSON = `
		{
			"id": 1,
			"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
			"ticker": "ARCA:BIL",
			"externalData": {
				"data": [
					{
						"source": "YAHOO_FINANCE",
						"ticker": "BIL-PERSISTED",
						"exchangeId": "PCX-PERSISTED"
					},
					{
						"source": "YAHOO_FINANCE",
						"ticker": "BIL-SECOND",
						"exchangeId": "SECOND"
					}
				]
			}
		}
	`
	assert.JSONEq(t, expectedAssetJSON, string(persistedAssetJSON))
}

func TestGetAssetByIdOrTicker(t *testing.T) {

	t.Run(
		"TestGetAssetById",
		func(t *testing.T) {

			// Test retrieving a valid asset
			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/1")
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusOK, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var actualResponseJSON = string(body)
			var expectedResponseJSON = `
				{
					"id": 1,
					"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"ticker": "ARCA:BIL"
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)

	t.Run(
		"TestGetAssetByTicker",
		func(t *testing.T) {

			// Test retrieving a valid asset
			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/ARCA:BIL")
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusOK, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var actualResponseJSON = string(body)
			var expectedResponseJSON = `
				{
					"id": 1,
					"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"ticker": "ARCA:BIL"
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)
}

// TestPutAsset tests the PUT /api/asset endpoint to update an existing asset's ticker and name.
//
// Authored by: GitHub Copilot
func TestPutAsset(t *testing.T) {

	var testTickerBefore = "TEST:BEFORE"
	var testNameBefore = "Test Asset Before Update"
	var testTickerAfter = "TEST:AFTER"
	var testNameAfter = "Test Asset After Update"

	var testAsset = insertTestAsset(t, testTickerBefore, testNameBefore)
	var testAssetIdString = strconv.FormatInt(testAsset.Id, 10)

	var putAssetJSON = `
		{
			"id":` + testAssetIdString + `,
			"ticker":"` + testTickerAfter + `",
			"name":"` + testNameAfter + `"
		}
	`

	response := putAsset(t, putAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"id":` + testAssetIdString + `,
			"ticker":"` + testTickerAfter + `",
			"name":"` + testNameAfter + `"
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))

	assertPersistedAsset(t, testAsset.Id, testTickerAfter, testNameAfter)
}

// TestPutAssetFailureWithoutId tests the PUT /api/asset endpoint returns a validation error
// when the asset ID is missing from the request body.
//
// Authored by: GitHub Copilot
func TestPutAssetFailureWithoutId(t *testing.T) {

	var putAssetJSONNoId = `
		{
			"ticker": "TEST:NOID",
			"name": "Asset Without ID"
		}
	`

	response := putAsset(t, putAssetJSONNoId)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'id' failed validation: is required"
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}

// TestPutAssetFailureWithZeroId tests the PUT /api/asset endpoint returns a validation error
// when the asset ID is zero in the request body.
//
// Authored by: GitHub Copilot
func TestPutAssetFailureWithZeroId(t *testing.T) {

	var putAssetJSONZeroId = `
		{
			"id": 0,
			"ticker": "TEST:ZEROID",
			"name": "Asset With Zero ID"
		}
	`

	response := putAsset(t, putAssetJSONZeroId)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'id' failed validation: is required"
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}

// TestPutAssetFailureWithoutRequiredFields tests the PUT /api/asset endpoint returns validation
// errors when required fields (name, ticker) are missing from the request body.
//
// Authored by: GitHub Copilot
func TestPutAssetFailureWithoutRequiredFields(t *testing.T) {

	t.Run(
		"WithoutName",
		func(t *testing.T) {

			var putAssetJSON = `
				{
					"id": 1,
					"ticker": "TEST:NONAME"
				}
			`

			response := putAsset(t, putAssetJSON)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'name' failed validation: is required"
					]
				}
			`

			assert.JSONEq(t, expectedResponseJSON, string(body))
		},
	)

	t.Run(
		"WithoutTicker",
		func(t *testing.T) {

			var putAssetJSON = `
				{
					"id": 1,
					"name": "Asset Without Ticker"
				}
			`

			response := putAsset(t, putAssetJSON)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'ticker' failed validation: is required"
					]
				}
			`

			assert.JSONEq(t, expectedResponseJSON, string(body))
		},
	)

	t.Run(
		"WithoutNameAndTicker",
		func(t *testing.T) {

			var putAssetJSON = `
				{
					"id": 1
				}
			`

			response := putAsset(t, putAssetJSON)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'name' failed validation: is required",
						"Field 'ticker' failed validation: is required"
					]
				}
			`

			assert.JSONEq(t, expectedResponseJSON, string(body))
		},
	)
}

// TestGetAssetByIdNotFound tests the GET /api/asset/{id} endpoint with a non-existent asset ID.
//
// Authored by: GitHub Copilot
func TestGetAssetByIdNotFound(t *testing.T) {

	// Test with a non-existent asset ID
	response, err := http.Get(infra.TestAPIURLPrefix + "/asset/999")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"errorMessage": "Data not found",
			"details": [
				"Asset with identifier 999 not found"
			]
		}
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetAssetByIdInvalidId tests the GET /api/asset/{id} endpoint with invalid ID formats.
//
// Authored by: GitHub Copilot
func TestGetAssetByIdInvalidId(t *testing.T) {

	// Test case: Empty ID (Gin redirects /api/asset/ to /api/asset)
	t.Run(
		"InvalidIdEmpty",
		func(t *testing.T) {

			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/")
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			// Gin redirects /api/asset/ to /api/asset and returns all assets
			assert.Equal(t, http.StatusOK, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			// Should return the complete list of assets (same as GET /api/asset)
			var actualResponseJSON = string(body)
			var expectedResponseJSON = `
			[
				{
					"id": 1,
					"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"ticker": "ARCA:BIL"
				},
				{
					"id": 6,
					"name": "iShares Msci Brazil ETF",
					"ticker": "ARCA:EWZ"
				},
				{
					"id": 7,
					"name": "SPDR S&P 500 ETF Trust",
					"ticker": "ARCA:SPY"
				},
				{
					"id": 2,
					"name": "iShares 0-5 Year TIPS Bond ETF",
					"ticker": "ARCA:STIP"
				},
				{
					"id": 3,
					"name": "iShares 7-10 Year Treasury Bond ETF",
					"ticker": "NasdaqGM:IEF"
				},
				{
					"id": 5,
					"name": "iShares Short Treasury Bond ETF",
					"ticker": "NasdaqGM:SHV"
				},
				{
					"id": 4,
					"name": "iShares 20+ Year Treasury Bond ETF",
					"ticker": "NasdaqGM:TLT"
				}
			]
		`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)
}
