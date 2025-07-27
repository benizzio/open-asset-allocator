package inttest

import (
	"fmt"
	"github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestGetPortfolioAllocationHistory(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history")
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
						"assetMarketPrice":"100",
						"assetQuantity":"100.00009",
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
						"assetMarketPrice":"100",
						"assetQuantity":"100",
						"totalMarketValue":10000
					},
					{
						"assetId": 2,
						"assetName":"iShares 0-5 Year TIPS Bond ETF",
						"assetTicker":"ARCA:STIP",
						"class":"BONDS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"80",
						"totalMarketValue":8000
					},
					{
						"assetId": 3,
						"assetName":"iShares 7-10 Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:IEF",
						"class":"BONDS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"60",
						"totalMarketValue":6000
					},
					{
						"assetId": 4,
						"assetName":"iShares 20+ Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:TLT",
						"class":"BONDS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"30",
						"totalMarketValue":3000
					},
					{
						"assetId": 5,	
						"assetName":"iShares Short Treasury Bond ETF",
						"assetTicker":"NasdaqGM:SHV",
						"class":"STOCKS",
						"cashReserve":true,
						"assetMarketPrice":"100",
						"assetQuantity":"80",
						"totalMarketValue":9000
					},
					{
 						"assetId": 7,
						"assetName":"SPDR S\u0026P 500 ETF Trust",
						"assetTicker":"ARCA:SPY",
						"class":"STOCKS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"90",
						"totalMarketValue":8000
					},
					{
						"assetId": 6,
						"assetName":"iShares Msci Brazil ETF",
						"assetTicker":"ARCA:EWZ",
						"class":"STOCKS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"10",
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

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history?timeFrameTag=202503")
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
						"assetMarketPrice":"100",
						"assetQuantity":"100.00009",
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

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history?observationTimestampId=2")
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
						"assetMarketPrice":"100",
						"assetQuantity":"100.00009",
						"totalMarketValue":10000
					}
				],
				"totalMarketValue":10000
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetPortfolioAllocationHistoryForObservationTimestampNoneFound(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/2/history?observationTimestampId=2")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[]
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
	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history/observation")
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

func TestGetAvailablePortfolioAllocationClasses(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-classes")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		["BONDS", "STOCKS"]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetAvailablePortfolioAllocationClassesNoneFound(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/2/allocation-classes")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestPostPortfolioAllocationHistorySuccess tests the successful creation of portfolio allocation history.
//
// This test verifies that the POST endpoint correctly accepts a portfolio snapshot
// and returns HTTP 204 No Content as expected, without any response body.
// It also verifies that the data is correctly persisted in the portfolio_allocation_fact table.
//
// Authored by: GitHub Copilot
func TestPostPortfolioAllocationHistorySuccess(t *testing.T) {

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": 3
			},
			"allocations": [
				{
					"assetId": 1,
					"assetName": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"assetTicker": "ARCA:BIL",
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": "150.00",
					"assetMarketPrice": "100.00",
					"totalMarketValue": 15000
				},
				{
					"assetId": 2,
					"assetName": "iShares 0-5 Year TIPS Bond ETF",
					"assetTicker": "ARCA:STIP",
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": "100.00",
					"assetMarketPrice": "100.00",
					"totalMarketValue": 10000
				}
			],
			"totalMarketValue": 25000
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(postPortfolioSnapshotJSON),
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// Verify that the response body is empty as expected for 204 No Content
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, body)

	// Verify that the data was correctly persisted in the portfolio_allocation_fact table
	assertPortfolioAllocationFactPersisted(t, 1, 1, "BONDS", false, "150.00000000", "100.00000000", 15000, 3)
	assertPortfolioAllocationFactPersisted(t, 1, 2, "BONDS", false, "100.00000000", "100.00000000", 10000, 3)
}

// assertPortfolioAllocationFactPersisted verifies that a portfolio allocation record
// was correctly inserted into the portfolio_allocation_fact table with the expected values.
//
// Authored by: GitHub Copilot
func assertPortfolioAllocationFactPersisted(
	t *testing.T,
	portfolioId int,
	assetId int,
	class string,
	cashReserve bool,
	assetQuantity string,
	assetMarketPrice string,
	totalMarketValue int64,
	observationTimeId int,
) {

	var portfolioIdString = strconv.Itoa(portfolioId)
	var assetIdString = strconv.Itoa(assetId)
	var cashReserveString = strconv.FormatBool(cashReserve)
	var totalMarketValueString = strconv.FormatInt(totalMarketValue, 10)
	var observationTimeIdString = strconv.Itoa(observationTimeId)

	var expectedRecord = dbx.NullStringMap{
		"portfolio_id":        util.ToNullString(portfolioIdString),
		"asset_id":            util.ToNullString(assetIdString),
		"class":               util.ToNullString(class),
		"cash_reserve":        util.ToNullString(cashReserveString),
		"asset_quantity":      util.ToNullString(assetQuantity),
		"asset_market_price":  util.ToNullString(assetMarketPrice),
		"total_market_value":  util.ToNullString(totalMarketValueString),
		"observation_time_id": util.ToNullString(observationTimeIdString),
	}

	var query = fmt.Sprintf(
		"SELECT portfolio_id, asset_id, class, cash_reserve, asset_quantity, asset_market_price, total_market_value, observation_time_id FROM portfolio_allocation_fact WHERE portfolio_id=%s AND asset_id=%s AND observation_time_id=%s",
		portfolioIdString,
		assetIdString,
		observationTimeIdString,
	)

	inttestutil.AssertDBWithQuery(t, query, expectedRecord)
}
