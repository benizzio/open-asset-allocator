package inttest

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	"github.com/stretchr/testify/assert"
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
						"totalMarketValue":"10000"
					}
				],
				"totalMarketValue":"10000"
			},
			{
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
						"totalMarketValue":"10000"
					},
					{
						"assetId": 2,
						"assetName":"iShares 0-5 Year TIPS Bond ETF",
						"assetTicker":"ARCA:STIP",
						"class":"BONDS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"80",
						"totalMarketValue":"8000"
					},
					{
						"assetId": 3,
						"assetName":"iShares 7-10 Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:IEF",
						"class":"BONDS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"60",
						"totalMarketValue":"6000"
					},
					{
						"assetId": 4,
						"assetName":"iShares 20+ Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:TLT",
						"class":"BONDS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"30",
						"totalMarketValue":"3000"
					},
					{
						"assetId": 5,	
						"assetName":"iShares Short Treasury Bond ETF",
						"assetTicker":"NasdaqGM:SHV",
						"class":"STOCKS",
						"cashReserve":true,
						"assetMarketPrice":"100",
						"assetQuantity":"80",
						"totalMarketValue":"9000"
					},
					{
 						"assetId": 7,
						"assetName":"SPDR S\u0026P 500 ETF Trust",
						"assetTicker":"ARCA:SPY",
						"class":"STOCKS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"90",
						"totalMarketValue":"8000"
					},
					{
						"assetId": 6,
						"assetName":"iShares Msci Brazil ETF",
						"assetTicker":"ARCA:EWZ",
						"class":"STOCKS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"10",
						"totalMarketValue":"1000"
					}
				],
				"totalMarketValue":"45000"
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
						"totalMarketValue":"10000"
					}
				],
				"totalMarketValue":"10000"
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
					"totalMarketValue": "15000"
				},
				null,
				{
					"assetId": 2,
					"assetName": "This name should not affect asset record 2",
					"assetTicker": "TTSNAAR2:TEST",
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": "100.00",
					"assetMarketPrice": "100.00",
					"totalMarketValue": "10000"
				},
				{
					"assetName": "New Asset",
					"assetTicker": "Test:NEW",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": "20",
					"assetMarketPrice": "100",
					"totalMarketValue": "2000"
				}
			]
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(postPortfolioSnapshotJSON),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

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

	// Verify that the response body is empty as expected for 204 No Content
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	// Verify that the data was correctly persisted in the portfolio_allocation_fact table
	var portfolioIdString = strconv.Itoa(1)

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
			"observation_time_id":   inttestutil.NotNullAssertableNullString(),
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
			"observation_time_id":   inttestutil.NotNullAssertableNullString(),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202505"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-05-01T00:00:00Z"),
		},
		{
			"portfolio_id":          inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":              inttestutil.NotNullAssertableNullString(),
			"class":                 inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"asset_quantity":        inttestutil.ToAssertableNullString("20.00000000"),
			"asset_market_price":    inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":    inttestutil.ToAssertableNullString("2000"),
			"name":                  inttestutil.ToAssertableNullString("New Asset"),
			"ticker":                inttestutil.ToAssertableNullString("Test:NEW"),
			"observation_time_id":   inttestutil.NotNullAssertableNullString(),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202505"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-05-01T00:00:00Z"),
		},
	}

	inttestutil.AssertDBWithQueryMultipleRows(t, allocationHistoryQuery, expectedRecords)
}

func TestPostPortfolioAllocationHistoryInsertEmptyZeroTimestamp(t *testing.T) {

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": 0
			},
			"allocations": [
				{
					"assetName": "New Asset",
					"assetTicker": "Test:NEW",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": "20",
					"assetMarketPrice": "100",
					"totalMarketValue": "2000"
				}
			]
		}
	`

	var testTime = time.Now()
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
					SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag LIKE '%%T%%'
				)`,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:NEW'").
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE observation_time_tag LIKE '%%T%%'").
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

	// Define expected records for the test case
	var expectedRecords = []inttestutil.AssertableNullStringMap{
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.NotNullAssertableNullString(),
			"class":               inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("20.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("2000"),
			"name":                inttestutil.ToAssertableNullString("New Asset"),
			"ticker":              inttestutil.ToAssertableNullString("Test:NEW"),
			"observation_time_id": inttestutil.NotNullAssertableNullString(),
			"observation_time_tag": inttestutil.ToAssertableNullStringWithAssertion(
				func(t *testing.T, actual sql.NullString) {
					dateTime, dateErr := time.Parse(time.RFC3339, actual.String)
					assert.NoError(t, dateErr)
					assert.WithinDuration(t, testTime, dateTime, time.Second)
				},
			),
			"observation_timestamp": inttestutil.ToAssertableNullStringWithAssertion(
				func(t *testing.T, actual sql.NullString) {
					dateTime, dateErr := time.Parse(time.RFC3339Nano, actual.String)
					assert.NoError(t, dateErr)
					assert.WithinDuration(t, testTime, dateTime, time.Second)
				},
			),
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
			WHERE p.portfolio_id=%s AND o.observation_time_tag LIKE '%s' 
			ORDER BY p.asset_id`,
		portfolioIdString,
		"%T%",
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
					"totalMarketValue": "1500"
				},
				{
					"assetId": "1",
					"assetName": "This name should not affect asset record 2",
					"assetTicker": "TTSNAAR2:TEST",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": 1.00,
					"assetMarketPrice": 2.00,
					"totalMarketValue": "4"
				},
				{
					"assetId": "6",
					"assetName": "New Asset",
					"assetTicker": "Test:NEW",
					"class": "STOCKS",
					"cashReserve": false,
					"assetQuantity": "20",
					"assetMarketPrice": "100",
					"totalMarketValue": "2000"
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
	assert.NotNil(t, response)

	if response == nil {
		return
	}

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	var portfolioIdString = strconv.Itoa(1)

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
			"asset_id":              inttestutil.NotNullAssertableNullString(),
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
			"asset_id":              inttestutil.NotNullAssertableNullString(),
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

	var allocationHistoryOtherObservationsQuery = fmt.Sprintf(
		`
			SELECT
				p.portfolio_id, 
			    p.asset_id, 
			    p.class, 
			    p.cash_reserve, 
			    p.asset_quantity, 
			    p.asset_market_price, 
			    p.total_market_value, 
			    p.observation_time_id
			FROM portfolio_allocation_fact p
			WHERE p.portfolio_id=%s AND p.observation_time_id!=%d
			ORDER BY p.observation_time_id, p.asset_id
		`,
		portfolioIdString,
		3,
	)

	var expectedAllocationsOtherObservationsResult = []inttestutil.AssertableNullStringMap{
		// Records from observation_time_id = 1 (initial data from 202501)
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("1"),
			"class":               inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("100.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("10000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("2"),
			"class":               inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("80.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("8000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("3"),
			"class":               inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("60.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("6000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("4"),
			"class":               inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("30.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("3000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("5"),
			"class":               inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("true"),
			"asset_quantity":      inttestutil.ToAssertableNullString("80.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("9000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("6"),
			"class":               inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("10.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("1000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("7"),
			"class":               inttestutil.ToAssertableNullString("STOCKS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("90.00000000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("8000"),
			"observation_time_id": inttestutil.ToAssertableNullString("1"),
		},
		// Record from observation_time_id = 2 (initial data from 202503)
		{
			"portfolio_id":        inttestutil.ToAssertableNullString(portfolioIdString),
			"asset_id":            inttestutil.ToAssertableNullString("1"),
			"class":               inttestutil.ToAssertableNullString("BONDS"),
			"cash_reserve":        inttestutil.ToAssertableNullString("false"),
			"asset_quantity":      inttestutil.ToAssertableNullString("100.00009000"),
			"asset_market_price":  inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value":  inttestutil.ToAssertableNullString("10000"),
			"observation_time_id": inttestutil.ToAssertableNullString("2"),
		},
	}

	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		allocationHistoryOtherObservationsQuery,
		expectedAllocationsOtherObservationsResult,
	)

}

func TestPostPortfolioAllocationValidations(t *testing.T) {

	t.Run(
		"TestPostPortfolioAllocationWithoutObservationAndAllocationFields",
		func(t *testing.T) {

			var postPortfolioSnapshotJSON = `
				{
					"allocations": [
						{
						}
					]
				}
			`

			actualResponseJSONNullFields := string(
				postPortfolioAllocationForValidationFailure(
					t,
					postPortfolioSnapshotJSON,
				),
			)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'observationTimestamp' failed validation: is required",
						"Field 'allocations[0].class' failed validation: is required",
						"Field 'allocations[0].totalMarketValue' failed validation: is required"
					]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
		},
	)

	t.Run(
		"TestPostPortfolioAllocationWithInvalidAssetFields",
		func(t *testing.T) {

			var postPortfolioSnapshotJSON = `
				{
					"observationTimestamp": {
						"id": "3"
					},
					"allocations": [
						{
							"class": "TEST",
							"totalMarketValue": "10"
						}
					]
				}
			`

			actualResponseJSONNullFields := string(
				postPortfolioAllocationForValidationFailure(
					t,
					postPortfolioSnapshotJSON,
				),
			)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'allocations[0].assetId' failed validation: if assetId is not provided, assetTicker and assetName must be provided"
					]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
		},
	)

	t.Run(
		"TestPostPortfolioAllocationWithInvalidEmptyAssetFields",
		func(t *testing.T) {

			var postPortfolioSnapshotJSON = `
				{
					"observationTimestamp": {
						"id": "3"
					},
					"allocations": [
						{
							"assetId": "",
							"assetName": "",
							"assetTicker": "",
							"class": "TEST",
							"totalMarketValue": "10"
						}
					]
				}
			`

			actualResponseJSONNullFields := string(
				postPortfolioAllocationForValidationFailure(
					t,
					postPortfolioSnapshotJSON,
				),
			)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'allocations[0].assetId' failed validation: if assetId is not provided, assetTicker and assetName must be provided"
					]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
		},
	)

	t.Run(
		"TestPostPortfolioAllocationWithEmptyAllocationFields",
		func(t *testing.T) {

			var postPortfolioSnapshotJSON = `
				{
					"allocations": [
						{
							"class": "",
							"totalMarketValue": "0"
						}
					]
				}
			`

			actualResponseJSONNullFields := string(
				postPortfolioAllocationForValidationFailure(
					t,
					postPortfolioSnapshotJSON,
				),
			)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'observationTimestamp' failed validation: is required",
						"Field 'allocations[0].class' failed validation: is required"
					]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
		},
	)

	t.Run(
		"TestPostPortfolioAllocationWithoutAllocations",
		func(t *testing.T) {

			var postPortfolioSnapshotJSON = `
				{
					"observationTimestamp": {
						"id": "3"
					}
				}
			`

			actualResponseJSONNullFields := string(
				postPortfolioAllocationForValidationFailure(
					t,
					postPortfolioSnapshotJSON,
				),
			)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'allocations' failed validation: is required"
					]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
		},
	)

	t.Run(
		"TestPostPortfolioAllocationWithEmptyAllocations",
		func(t *testing.T) {

			var postPortfolioSnapshotJSON = `
				{
					"observationTimestamp": {
						"id": "3"
					},
					"allocations": []
				}
			`

			actualResponseJSONNullFields := string(
				postPortfolioAllocationForValidationFailure(
					t,
					postPortfolioSnapshotJSON,
				),
			)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'allocations' failed validation: must be at least 1"
					]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
		},
	)
}

func postPortfolioAllocationForValidationFailure(t *testing.T, postPortfolioJSON string) []byte {

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
	return body
}

// TestGetPortfolioAllocationHistoryWithMultiplePortfoliosAndManyObservations tests the fix for the issue where
// portfolio history data was being skipped when there are more than 10 observations across multiple portfolios.
//
// This test verifies that when the LIMIT clause in the availableObservationTimestampsComplement CTE is applied,
// it correctly filters by portfolio_id so that data for a specific portfolio is not incorrectly skipped.
//
// Authored by: GitHub Copilot
func TestGetPortfolioAllocationHistoryWithMultiplePortfoliosAndManyObservations(t *testing.T) {

	// Setup: Create 15 observation timestamps (more than the default limit of 10)
	// for portfolio 2, ensuring they all get newer timestamps than portfolio 1
	var setupSQL = `
		-- Insert 15 observation timestamps for testing
		INSERT INTO portfolio_allocation_obs_time (observation_time_tag, observation_timestamp)
		VALUES 
			('test_obs_1', '2025-10-01 00:00:00'::TIMESTAMP),
			('test_obs_2', '2025-10-02 00:00:00'::TIMESTAMP),
			('test_obs_3', '2025-10-03 00:00:00'::TIMESTAMP),
			('test_obs_4', '2025-10-04 00:00:00'::TIMESTAMP),
			('test_obs_5', '2025-10-05 00:00:00'::TIMESTAMP),
			('test_obs_6', '2025-10-06 00:00:00'::TIMESTAMP),
			('test_obs_7', '2025-10-07 00:00:00'::TIMESTAMP),
			('test_obs_8', '2025-10-08 00:00:00'::TIMESTAMP),
			('test_obs_9', '2025-10-09 00:00:00'::TIMESTAMP),
			('test_obs_10', '2025-10-10 00:00:00'::TIMESTAMP),
			('test_obs_11', '2025-10-11 00:00:00'::TIMESTAMP),
			('test_obs_12', '2025-10-12 00:00:00'::TIMESTAMP),
			('test_obs_13', '2025-10-13 00:00:00'::TIMESTAMP),
			('test_obs_14', '2025-10-14 00:00:00'::TIMESTAMP),
			('test_obs_15', '2025-10-15 00:00:00'::TIMESTAMP)
		;

		-- Add portfolio allocations for portfolio 2 with all 15 observation timestamps
		INSERT INTO portfolio_allocation_fact (
			asset_id, "class", cash_reserve, asset_quantity, asset_market_price,
			total_market_value, portfolio_id, observation_time_id
		)
		SELECT 
			1, 'BONDS', FALSE, 100, 100, 10000, 2,
			id
		FROM portfolio_allocation_obs_time 
		WHERE observation_time_tag ~ '^test_obs_'
		;
	`

	err := inttestinfra.ExecuteDBQuery(setupSQL)
	assert.NoError(t, err)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact 
				WHERE observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag ~ '^test_obs_'
				)`,
			).
			AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time WHERE observation_time_tag ~ '^test_obs_'`).
			Build(),
	)

	// Act: Get portfolio history for portfolio 2 with the default limit of 10 observations
	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/2/history")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	// Assert: Verify that we get exactly 10 observations for portfolio 2
	// (the most recent 10 based on the LIMIT)
	var actualResponseJSON = string(body)

	// The response should contain 10 snapshots with observation timestamps from test_obs_15 down to test_obs_6
	// This verifies that the CTE correctly filters by portfolio_id and doesn't skip data
	assert.Contains(t, actualResponseJSON, "test_obs_15")
	assert.Contains(t, actualResponseJSON, "test_obs_14")
	assert.Contains(t, actualResponseJSON, "test_obs_13")
	assert.Contains(t, actualResponseJSON, "test_obs_12")
	assert.Contains(t, actualResponseJSON, "test_obs_11")
	assert.Contains(t, actualResponseJSON, "test_obs_10")
	assert.Contains(t, actualResponseJSON, "test_obs_9")
	assert.Contains(t, actualResponseJSON, "test_obs_8")
	assert.Contains(t, actualResponseJSON, "test_obs_7")
	assert.Contains(t, actualResponseJSON, "test_obs_6")

	// test_obs_1 through test_obs_5 should NOT be in the response (beyond the LIMIT of 10)
	// Use specific format to avoid substring matching issues (e.g., test_obs_1 appearing in test_obs_10)
	assert.NotContains(t, actualResponseJSON, "\"test_obs_5\"")
	assert.NotContains(t, actualResponseJSON, "\"test_obs_4\"")
	assert.NotContains(t, actualResponseJSON, "\"test_obs_3\"")
	assert.NotContains(t, actualResponseJSON, "\"test_obs_2\"")
	assert.NotContains(t, actualResponseJSON, "\"test_obs_1\"")

	// Additional verification: Ensure portfolio 1 history is still working correctly
	// and hasn't been affected by the changes
	response1, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, response1.StatusCode)

	body1, err := io.ReadAll(response1.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body1)

	// Portfolio 1 should still have its original 2 observations
	var actualResponse1JSON = string(body1)
	assert.Contains(t, actualResponse1JSON, "202503")
	assert.Contains(t, actualResponse1JSON, "202501")
}

// TestPostPortfolioAllocationHistoryValidation_ClassExceedsMaxLength tests that posting
// a portfolio allocation with a class exceeding the max length (100 characters) returns a validation error.
//
// Authored by: GitHub Copilot
func TestPostPortfolioAllocationHistoryValidation_ClassExceedsMaxLength(t *testing.T) {

	var longClass = strings.Repeat("C", 101) // 101 characters exceeds max=100

	var payload = `{
		"observationTimestamp": {
			"timeTag": "TEST_LONG_CLASS",
			"timestamp": "2025-12-01T00:00:00Z"
		},
		"allocations": [
			{
				"assetTicker": "TEST:TICKER",
				"assetName": "Test Asset",
				"class": "` + longClass + `",
				"totalMarketValue": "1000"
			}
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(payload),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
		"errorMessage": "Validation failed",
		"details": ["Field 'allocations[0].class' failed validation: must not exceed 100"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// TestPostPortfolioAllocationHistoryValidation_AssetTickerExceedsMaxLength tests that posting
// a portfolio allocation with an asset ticker exceeding the max length (40 characters) returns a validation error.
//
// Authored by: GitHub Copilot
func TestPostPortfolioAllocationHistoryValidation_AssetTickerExceedsMaxLength(t *testing.T) {

	var longTicker = strings.Repeat("T", 41) // 41 characters exceeds max=40

	var payload = `{
		"observationTimestamp": {
			"timeTag": "TEST_LONG_TICKER",
			"timestamp": "2025-12-01T00:00:00Z"
		},
		"allocations": [
			{
				"assetTicker": "` + longTicker + `",
				"assetName": "Test Asset",
				"class": "STOCKS",
				"totalMarketValue": "1000"
			}
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(payload),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
		"errorMessage": "Validation failed",
		"details": ["Field 'allocations[0].assetTicker' failed validation: must not exceed 40"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// TestPostPortfolioAllocationHistoryValidation_AssetNameExceedsMaxLength tests that posting
// a portfolio allocation with an asset name exceeding the max length (100 characters) returns a validation error.
//
// Authored by: GitHub Copilot
func TestPostPortfolioAllocationHistoryValidation_AssetNameExceedsMaxLength(t *testing.T) {

	var longName = strings.Repeat("N", 101) // 101 characters exceeds max=100

	var payload = `{
		"observationTimestamp": {
			"timeTag": "TEST_LONG_NAME",
			"timestamp": "2025-12-01T00:00:00Z"
		},
		"allocations": [
			{
				"assetTicker": "TEST:TICKER",
				"assetName": "` + longName + `",
				"class": "STOCKS",
				"totalMarketValue": "1000"
			}
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(payload),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
		"errorMessage": "Validation failed",
		"details": ["Field 'allocations[0].assetName' failed validation: must not exceed 100"]
	}`
	assert.JSONEq(t, expected, string(body))
}
