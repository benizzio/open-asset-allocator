package inttest

import (
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetPortfolioAllocationHistory(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/history")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"timeFrameTag":"202503",
				"observationTimestamp" : {
					"id": 2,
					"timeTag": "202503",
					"timestamp": "2025-03-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					}
				],
				"totalMarketValue":10000
			},
			{
				"timeFrameTag":"202501",
				"observationTimestamp" : {
					"id": 1,
					"timeTag": "202501",
					"timestamp": "2025-01-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					},
					{
						"assetId": 2,
						"assetName":"iShares 0-5 Year TIPS Bond ETF",
						"assetTicker":"ARCA:STIP",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":8000
					},
					{
						"assetId": 3,
						"assetName":"iShares 7-10 Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:IEF",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":6000
					},
					{
						"assetId": 4,
						"assetName":"iShares 20+ Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:TLT",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":3000
					},
					{
						"assetId": 5,	
						"assetName":"iShares Short Treasury Bond ETF",
						"assetTicker":"NasdaqGM:SHV",
						"class":"STOCKS",
						"cashReserve":true,
						"totalMarketValue":9000
					},
					{
 						"assetId": 7,
						"assetName":"SPDR S\u0026P 500 ETF Trust",
						"assetTicker":"ARCA:SPY",
						"class":"STOCKS",
						"cashReserve":false,
						"totalMarketValue":8000
					},
					{
						"assetId": 6,
						"assetName":"iShares Msci Brazil ETF",
						"assetTicker":"ARCA:EWZ",
						"class":"STOCKS",
						"cashReserve":false,
						"totalMarketValue":1000
					}
				],
				"totalMarketValue":45000
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetPortfolioAllocationHistoryForTimeFrame(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/history?timeFrameTag=202503")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"timeFrameTag":"202503",
				"observationTimestamp" : {
					"id": 2,
					"timeTag": "202503",
					"timestamp": "2025-03-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					}
				],
				"totalMarketValue":10000
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetPortfolioAllocationHistoryForObservationTimestamp(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/history?observationTimestampId=2")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"timeFrameTag":"202503",
				"observationTimestamp" : {
					"id": 2,
					"timeTag": "202503",
					"timestamp": "2025-03-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					}
				],
				"totalMarketValue":10000
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetAvailableHistoryObservations tests the retrieval of available history observations for a portfolio.
//
// This test verifies that the API correctly returns the list of available observation timestamps
// for a given portfolio, including their IDs, time tags, and timestamps.
//
// Authored by: GitHub Copilot
func TestGetAvailableHistoryObservations(t *testing.T) {

	// Call the API endpoint to get available history observations
	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/history/observation")
	assert.NoError(t, err)

	// Verify successful response status code
	assert.Equal(t, http.StatusOK, response.StatusCode)

	// Read and validate response body
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	// Check the actual response against expected response
	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"id": 2,
				"timeTag": "202503",
				"timestamp": "2025-03-01T00:00:00Z"
			},
			{
				"id": 1,
				"timeTag": "202501",
				"timestamp": "2025-01-01T00:00:00Z"
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}
