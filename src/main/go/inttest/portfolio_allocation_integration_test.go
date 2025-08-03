package inttest

import (
	"database/sql"
	"fmt"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
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

// TestPostPortfolioAllocationHistoryInsertOnly tests the successful creation of portfolio allocation history,
// only inserting records.
//
// This test verifies that the POST endpoint correctly accepts a portfolio snapshot
// and returns HTTP 204 No Content as expected, without any response body.
// It also verifies that the data is correctly persisted in the portfolio_allocation_fact table.
//
// Co-authored by: GitHub Copilot
func TestPostPortfolioAllocationHistoryInsertOnly(t *testing.T) {

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"timeTag": "202505",
				"timestamp": "2025-05-01T00:00:00Z"
			},
			"allocations": [
				{
					"assetId": 1,
					"assetName": "This name should not affect asset record",
					"assetTicker": "TTSNAAR:TEST",
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": "150.00",
					"assetMarketPrice": "100.00",
					"totalMarketValue": 15000
				},
				{
					"assetId": 2,
					"assetName": "This name should not affect asset record 2",
					"assetTicker": "TTSNAAR2:TEST",
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": "100.00",
					"assetMarketPrice": "100.00",
					"totalMarketValue": 10000
				},
				{
					"assetName": "New Asset",
					"assetTicker": "Test:NEW",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": "20",
					"assetMarketPrice": "100",
					"totalMarketValue": 2000
				}
			]
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(postPortfolioSnapshotJSON),
	)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
				DELETE FROM portfolio_allocation_fact 
				WHERE observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag = '202505'
				)`,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:NEW'").
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE observation_time_tag = '202505'").
			Build(),
	)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// Verify that the response body is empty as expected for 204 No Content
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	// Verify that the data was correctly persisted in the portfolio_allocation_fact table
	var portfolioIdString = strconv.Itoa(1)

	var notNullAssertion = inttestutil.ToAssertableNullStringWithAssertion(
		func(t *testing.T, actual sql.NullString) {
			assert.NotEmpty(t, actual.String)
			assert.True(t, actual.Valid)
		},
	)

	// Define expected records for the test case
	var expectedRecords = []inttestutil.AssertableNullStringMap{
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              inttestutil.ToAssertableNullString("1"),
			"class":                 inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("150.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("15000"),
			"name":                  inttestutil.ToAssertableNullString("SPDR Bloomberg 1-3 Month T-Bill ETF"),
			"ticker":                inttestutil.ToAssertableNullString("ARCA:BIL"),
			"observation_time_id":   notNullAssertion,
			"observation_time_tag":  inttestutil.ToAssertableNullString("202505"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-05-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              inttestutil.ToAssertableNullString("2"),
			"class":                 inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("100.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"name":                  inttestutil.ToAssertableNullString("iShares 0-5 Year TIPS Bond ETF"),
			"ticker":                inttestutil.ToAssertableNullString("ARCA:STIP"),
			"observation_time_id":   notNullAssertion,
			"observation_time_tag":  inttestutil.ToAssertableNullString("202505"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-05-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              notNullAssertion,
			"class":                 inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("20.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("2000"),
			"name":                  inttestutil.ToAssertableNullString("New Asset"),
			"ticker":                inttestutil.ToAssertableNullString("Test:NEW"),
			"observation_time_id":   notNullAssertion,
			"observation_time_tag":  inttestutil.ToAssertableNullString("202505"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-05-01T00:00:00Z"),
		},
	}

	var allocationHistoryQuery = fmt.Sprintf(
		`
			SELECT 
			    p.portfolio_id, 
			    p.asset_id, 
			    p.class, 
			    p.cash_reserve, 
			    p.asset_quantity, 
			    p.asset_market_price, 
			    p.total_market_value, 
			    a.ticker,
			    a.name,
			    p.observation_time_id,
			    o.observation_time_tag,
				o.observation_timestamp
			FROM portfolio_allocation_fact p 
			JOIN asset a ON p.asset_id = a.id
			JOIN portfolio_allocation_obs_time o ON p.observation_time_id = o.id
			WHERE p.portfolio_id=%s AND o.observation_time_tag='%s' 
			ORDER BY p.asset_id`,
		portfolioIdString,
		"202505",
	)

	inttestutil.AssertDBWithQueryMultipleRows(t, allocationHistoryQuery, expectedRecords)
}

func TestPostPortfolioAllocationHistoryFullMerge(t *testing.T) {

	var insertAllocationHistorySQL = `
		INSERT INTO portfolio_allocation_fact (
			asset_id,
			"class",
			cash_reserve,
			asset_quantity,
			asset_market_price,
			total_market_value,
			time_frame_tag,
			portfolio_id,
			observation_time_id
		)
		VALUES (
			   1,
			   'BONDS',
			   FALSE,
			   1,
			   100,
			   100,
			   '202504',
			   1,
			   3
			),
		    (
			   7,
			   'STOCKS',
			   FALSE,
			   10,
			   9,
			   90,
			   '202504',
			   1,
			   3
			)
		;
	`

	err := inttestinfra.ExecuteDBQuery(insertAllocationHistorySQL)
	assert.NoError(t, err)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
				DELETE FROM portfolio_allocation_fact 
				WHERE observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE id = 3
				)`,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:NEW'").
			Build(),
	)

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": "3"
			},
			"allocations": [
				{
					"assetId": "1",
					"assetName": "This name should not affect asset record",
					"assetTicker": "TTSNAAR:TEST",
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": 150.00,
					"assetMarketPrice": 10.00,
					"totalMarketValue": 1500
				},
				{
					"assetId": "1",
					"assetName": "This name should not affect asset record 2",
					"assetTicker": "TTSNAAR2:TEST",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": 1.00,
					"assetMarketPrice": 2.00,
					"totalMarketValue": 4
				},
				{
					"assetId": "6",
					"assetName": "New Asset",
					"assetTicker": "Test:NEW",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": "20",
					"assetMarketPrice": "100",
					"totalMarketValue": 2000
				},
				{
					"assetName": "New Asset Repeat 1",
					"assetTicker": "Test:NEW",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": "20",
					"assetMarketPrice": "100",
					"totalMarketValue": 2000
				},
				{
					"assetName": "New Asset Repeat 2",
					"assetTicker": "Test:NEW",
					"class": "TEST",
					"cashReserve": false,
					"assetQuantity": "30",
					"assetMarketPrice": "100",
					"totalMarketValue": 3000
				}
			]
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(postPortfolioSnapshotJSON),
	)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	var portfolioIdString = strconv.Itoa(1)

	var notNullAssertion = inttestutil.ToAssertableNullStringWithAssertion(
		func(t *testing.T, actual sql.NullString) {
			assert.NotEmpty(t, actual.String)
			assert.True(t, actual.Valid)
		},
	)

	var expectedRecords = []inttestutil.AssertableNullStringMap{
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              inttestutil.ToAssertableNullString("1"),
			"class":                 inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("150.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("10.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("1500"),
			"name":                  inttestutil.ToAssertableNullString("SPDR Bloomberg 1-3 Month T-Bill ETF"),
			"ticker":                inttestutil.ToAssertableNullString("ARCA:BIL"),
			"observation_time_id":   inttestutil.ToAssertableNullString("3"),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202504"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-04-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              inttestutil.ToAssertableNullString("1"),
			"class":                 inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("1.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("2.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("4"),
			"name":                  inttestutil.ToAssertableNullString("SPDR Bloomberg 1-3 Month T-Bill ETF"),
			"ticker":                inttestutil.ToAssertableNullString("ARCA:BIL"),
			"observation_time_id":   inttestutil.ToAssertableNullString("3"),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202504"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-04-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              inttestutil.ToAssertableNullString("6"),
			"class":                 inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("20.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("2000"),
			"name":                  inttestutil.ToAssertableNullString("iShares Msci Brazil ETF"),
			"ticker":                inttestutil.ToAssertableNullString("ARCA:EWZ"),
			"observation_time_id":   inttestutil.ToAssertableNullString("3"),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202504"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-04-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              notNullAssertion,
			"class":                 inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("20.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("2000"),
			"name":                  inttestutil.ToAssertableNullString("New Asset Repeat 2"),
			"ticker":                inttestutil.ToAssertableNullString("Test:NEW"),
			"observation_time_id":   inttestutil.ToAssertableNullString("3"),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202504"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-04-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              notNullAssertion,
			"class":                 inttestutil.ToAssertableNullString("TEST"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("30.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("3000"),
			"name":                  inttestutil.ToAssertableNullString("New Asset Repeat 2"),
			"ticker":                inttestutil.ToAssertableNullString("Test:NEW"),
			"observation_time_id":   inttestutil.ToAssertableNullString("3"),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202504"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-04-01T00:00:00Z"),
		},
	}

	var allocationHistoryQuery = fmt.Sprintf(
		`
			SELECT 
			    p.portfolio_id, 
			    p.asset_id, 
			    p.class, 
			    p.cash_reserve, 
			    p.asset_quantity, 
			    p.asset_market_price, 
			    p.total_market_value, 
			    a.ticker,
			    a.name,
			    p.observation_time_id,
			    o.observation_time_tag,
				o.observation_timestamp
			FROM portfolio_allocation_fact p 
			JOIN asset a ON p.asset_id = a.id
			JOIN portfolio_allocation_obs_time o ON p.observation_time_id = o.id
			WHERE p.portfolio_id=%s AND o.id=%d
			ORDER BY p.asset_id, p.class`,
		portfolioIdString,
		3,
	)

	inttestutil.AssertDBWithQueryMultipleRows(t, allocationHistoryQuery, expectedRecords)
}

// TODO test field validation errors
