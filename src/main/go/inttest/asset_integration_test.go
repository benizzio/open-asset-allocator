package inttest

import (
	"io"
	"net/http"
	"testing"

	"github.com/benizzio/open-asset-allocator/inttest/infra"
	"github.com/stretchr/testify/assert"
)

// TestGetKnownAssets tests the GET /api/asset endpoint to retrieve all known assets.
//
// Authored by: GitHub Copilot
func TestGetKnownAssets(t *testing.T) {

	response, err := http.Get(infra.TestAPIURLPrefix + "/asset")
	assert.NoError(t, err)

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
				"id": 4,
				"name": "iShares 20+ Year Treasury Bond ETF",
				"ticker": "NasdaqGM:TLT"
			},
			{
				"id": 5,
				"name": "iShares Short Treasury Bond ETF",
				"ticker": "NasdaqGM:SHV"
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
			}
		]
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetAssetByIdOrTicker(t *testing.T) {

	t.Run(
		"TestGetAssetById",
		func(t *testing.T) {

			// Test retrieving a valid asset
			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/1")
			assert.NoError(t, err)

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

// TestGetAssetByIdNotFound tests the GET /api/asset/{id} endpoint with a non-existent asset ID.
//
// Authored by: GitHub Copilot
func TestGetAssetByIdNotFound(t *testing.T) {

	// Test with a non-existent asset ID
	response, err := http.Get(infra.TestAPIURLPrefix + "/asset/999")
	assert.NoError(t, err)

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
					"id": 4,
					"name": "iShares 20+ Year Treasury Bond ETF",
					"ticker": "NasdaqGM:TLT"
				},
				{
					"id": 5,
					"name": "iShares Short Treasury Bond ETF",
					"ticker": "NasdaqGM:SHV"
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
				}
			]
		`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)
}
