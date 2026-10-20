package inttest

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

func TestGetPortfolioAllocationHistory(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history")
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
	defer deferCloseResponseBody(response)

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
	defer deferCloseResponseBody(response)

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

// TestGetPortfolioAllocationHistoryWithProjectedExternalAsset verifies that the
// history endpoint exposes the first persisted external asset entry for each
// allocation.
//
// Authored by: OpenCode
func TestGetPortfolioAllocationHistoryWithProjectedExternalAsset(t *testing.T) {
	var originalExternalData = capturePersistedAssetExternalData(t, 1)

	err := inttestinfra.ExecuteDBQuery(
		`UPDATE asset SET external_data = '{"data":[{"source":"YAHOO_FINANCE","ticker":"BIL","exchangeId":"PCX"},{"source":"YAHOO_FINANCE","ticker":"BIL-SECOND","exchangeId":"SECOND"}]}'::jsonb WHERE id = 1`,
		nil,
	)
	assert.NoError(t, err)

	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			1,
			originalExternalData,
		).Build(t),
	)

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history?observationTimestampId=2")
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
						"externalAsset": {
							"source": "YAHOO_FINANCE",
							"ticker": "BIL",
							"exchangeId": "PCX"
						},
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

// TestGetPortfolioAllocationHistoryOmitsExternalAssetWhenNull verifies that the
// history endpoint omits the external asset block when no persisted external
// data exists for the asset.
//
// Authored by: OpenCode
func TestGetPortfolioAllocationHistoryOmitsExternalAssetWhenNull(t *testing.T) {
	var originalExternalData = capturePersistedAssetExternalData(t, 1)

	err := inttestinfra.ExecuteDBQuery(`UPDATE asset SET external_data = NULL WHERE id = 1`, nil)
	assert.NoError(t, err)

	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			1,
			originalExternalData,
		).Build(t),
	)

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history?observationTimestampId=2")
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
	defer deferCloseResponseBody(response)

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

// TestPortfolioHistoryOrderingIsChronologicalAndDeterministic verifies chronological observation ordering,
// ID tie-breaking, limit-boundary selection, repeated-read stability, and ID-based reads outside the list limit.
//
// Authored by: OpenCode
func TestPortfolioHistoryOrderingIsChronologicalAndDeterministic(t *testing.T) {
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact
				 WHERE portfolio_id = 1 AND observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE 'pr5-ordering-%'
				 )`,
				nil,
			).
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE 'pr5-ordering-%'`,
				nil,
			).
			Build(t),
	)

	var insertObservationsSQL = `
		INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		VALUES
			(1, 'pr5-ordering-march', '2040-03-01T00:00:00Z'),
			(1, 'pr5-ordering-feb-1', '2040-02-01T00:00:00Z'),
			(1, 'pr5-ordering-april', '2040-04-01T00:00:00Z'),
			(1, 'pr5-ordering-feb-2', '2040-02-01T00:00:00Z'),
			(1, 'pr5-ordering-january', '2040-01-01T00:00:00Z'),
			(1, 'pr5-ordering-feb-3', '2040-02-01T00:00:00Z'),
			(1, 'pr5-ordering-may', '2040-05-01T00:00:00Z'),
			(1, 'pr5-ordering-feb-4', '2040-02-01T00:00:00Z'),
			(1, 'pr5-ordering-june', '2040-06-01T00:00:00Z'),
			(1, 'pr5-ordering-july', '2040-07-01T00:00:00Z'),
			(1, 'pr5-ordering-august', '2040-08-01T00:00:00Z'),
			(1, 'pr5-ordering-september', '2040-09-01T00:00:00Z'),
			(1, 'pr5-ordering-october', '2040-10-01T00:00:00Z')
	`
	var err = inttestinfra.ExecuteDBQuery(insertObservationsSQL, nil)
	require.NoError(t, err)

	var insertAllocationsSQL = `
		INSERT INTO portfolio_allocation_fact (
			asset_id, "class", cash_reserve, asset_quantity, asset_market_price,
			total_market_value, portfolio_id, observation_time_id
		)
		SELECT 1, 'BONDS', FALSE, 1, 100, 100, 1, id
		FROM portfolio_allocation_obs_time
		WHERE portfolio_id = 1 AND observation_time_tag LIKE 'pr5-ordering-%'
	`
	err = inttestinfra.ExecuteDBQuery(insertAllocationsSQL, nil)
	require.NoError(t, err)

	var insertAdditionalOctoberAllocationsSQL = `
		INSERT INTO portfolio_allocation_fact (
			asset_id, "class", cash_reserve, asset_quantity, asset_market_price,
			total_market_value, portfolio_id, observation_time_id
		)
		SELECT fixture.asset_id, fixture.class, fixture.cash_reserve, 1, fixture.total_market_value,
			fixture.total_market_value, 1, observation.id
		FROM (VALUES
			(2, 'STOCKS', FALSE, 200),
			(3, 'STOCKS', TRUE, 300)
		) AS fixture(asset_id, class, cash_reserve, total_market_value)
		CROSS JOIN portfolio_allocation_obs_time observation
		WHERE observation.portfolio_id = 1 AND observation.observation_time_tag = 'pr5-ordering-october'
	`
	err = inttestinfra.ExecuteDBQuery(insertAdditionalOctoberAllocationsSQL, nil)
	require.NoError(t, err)

	var observationIds = make(map[string]int64)
	err = inttestinfra.FetchWithDBQuery(
		`SELECT id, observation_time_tag FROM portfolio_allocation_obs_time
		 WHERE portfolio_id = 1 AND observation_time_tag LIKE 'pr5-ordering-%'`,
		nil,
		func(rows *dbx.Rows) error {
			var id int64
			var tag string
			var scanError error
			if scanError = rows.Scan(&id, &tag); scanError != nil {
				return scanError
			}
			observationIds[tag] = id
			return nil
		},
	)
	require.NoError(t, err)
	require.Len(t, observationIds, 13)

	type observationProjection struct {
		Id        int64  `json:"id"`
		TimeTag   string `json:"timeTag"`
		Timestamp string `json:"timestamp"`
	}
	type allocationProjection struct {
		CashReserve      bool   `json:"cashReserve"`
		TotalMarketValue string `json:"totalMarketValue"`
	}
	type snapshotProjection struct {
		ObservationTimestamp observationProjection  `json:"observationTimestamp"`
		Allocations          []allocationProjection `json:"allocations"`
		TotalMarketValue     string                 `json:"totalMarketValue"`
	}

	var expectedTags = []string{
		"pr5-ordering-october",
		"pr5-ordering-september",
		"pr5-ordering-august",
		"pr5-ordering-july",
		"pr5-ordering-june",
		"pr5-ordering-may",
		"pr5-ordering-april",
		"pr5-ordering-march",
		"pr5-ordering-feb-4",
		"pr5-ordering-feb-3",
	}
	var timestampByTag = map[string]string{
		"pr5-ordering-october":   "2040-10-01T00:00:00Z",
		"pr5-ordering-september": "2040-09-01T00:00:00Z",
		"pr5-ordering-august":    "2040-08-01T00:00:00Z",
		"pr5-ordering-july":      "2040-07-01T00:00:00Z",
		"pr5-ordering-june":      "2040-06-01T00:00:00Z",
		"pr5-ordering-may":       "2040-05-01T00:00:00Z",
		"pr5-ordering-april":     "2040-04-01T00:00:00Z",
		"pr5-ordering-march":     "2040-03-01T00:00:00Z",
		"pr5-ordering-feb-4":     "2040-02-01T00:00:00Z",
		"pr5-ordering-feb-3":     "2040-02-01T00:00:00Z",
		"pr5-ordering-feb-2":     "2040-02-01T00:00:00Z",
	}
	var expectedObservations = make([]observationProjection, 0, len(expectedTags))
	for _, tag := range expectedTags {
		expectedObservations = append(expectedObservations, observationProjection{
			Id:        observationIds[tag],
			TimeTag:   tag,
			Timestamp: timestampByTag[tag],
		})
	}
	expectedObservationsJSON, err := json.Marshal(expectedObservations)
	require.NoError(t, err)

	var readAttempt int
	for readAttempt = 0; readAttempt < 2; readAttempt++ {
		availableResponse, requestError := http.Get(
			inttestinfra.TestAPIURLPrefix + "/portfolio/1/history/observation",
		)
		require.NoError(t, requestError)
		availableBody, readError := io.ReadAll(availableResponse.Body)
		closeError := availableResponse.Body.Close()
		require.NoError(t, readError)
		require.NoError(t, closeError)
		require.Equal(t, http.StatusOK, availableResponse.StatusCode)
		assert.JSONEq(t, string(expectedObservationsJSON), string(availableBody))

		historyResponse, requestError := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/history")
		require.NoError(t, requestError)
		historyBody, readError := io.ReadAll(historyResponse.Body)
		closeError = historyResponse.Body.Close()
		require.NoError(t, readError)
		require.NoError(t, closeError)
		require.Equal(t, http.StatusOK, historyResponse.StatusCode)

		var snapshots []snapshotProjection
		require.NoError(t, json.Unmarshal(historyBody, &snapshots))
		require.Len(t, snapshots, len(expectedObservations))
		var actualHistoryObservations = make([]observationProjection, 0, len(snapshots))
		for _, snapshot := range snapshots {
			actualHistoryObservations = append(actualHistoryObservations, snapshot.ObservationTimestamp)
			if snapshot.ObservationTimestamp.TimeTag == "pr5-ordering-october" {
				assert.Equal(t, "600", snapshot.TotalMarketValue)
				assert.Equal(t, []allocationProjection{
					{CashReserve: false, TotalMarketValue: "100"},
					{CashReserve: true, TotalMarketValue: "300"},
					{CashReserve: false, TotalMarketValue: "200"},
				}, snapshot.Allocations)
			} else {
				assert.Equal(t, "100", snapshot.TotalMarketValue)
				assert.Equal(t, []allocationProjection{{CashReserve: false, TotalMarketValue: "100"}}, snapshot.Allocations)
			}
		}
		actualHistoryObservationsJSON, marshalError := json.Marshal(actualHistoryObservations)
		require.NoError(t, marshalError)
		assert.JSONEq(t, string(expectedObservationsJSON), string(actualHistoryObservationsJSON))
	}

	assert.Greater(t, observationIds["pr5-ordering-feb-4"], observationIds["pr5-ordering-feb-3"])
	assert.Greater(t, observationIds["pr5-ordering-feb-3"], observationIds["pr5-ordering-feb-2"])

	var olderObservationId = observationIds["pr5-ordering-january"]
	olderResponse, err := http.Get(
		inttestinfra.TestAPIURLPrefix + "/portfolio/1/history?observationTimestampId=" + strconv.FormatInt(olderObservationId, 10),
	)
	require.NoError(t, err)
	defer deferCloseResponseBody(olderResponse)
	require.Equal(t, http.StatusOK, olderResponse.StatusCode)
	olderBody, err := io.ReadAll(olderResponse.Body)
	require.NoError(t, err)
	var olderSnapshots []snapshotProjection
	require.NoError(t, json.Unmarshal(olderBody, &olderSnapshots))
	require.Len(t, olderSnapshots, 1)
	assert.Equal(t, observationProjection{
		Id:        olderObservationId,
		TimeTag:   "pr5-ordering-january",
		Timestamp: "2040-01-01T00:00:00Z",
	}, olderSnapshots[0].ObservationTimestamp)
	assert.Equal(t, "100", olderSnapshots[0].TotalMarketValue)
}

// TestPostPortfolioAllocationHistoryInsertOnly tests the successful creation of portfolio allocation history,
// only inserting records.
//
// This test verifies that the POST endpoint correctly accepts a portfolio snapshot
// and returns HTTP 204 No Content as expected, without any response body.
// It also verifies that the data is correctly persisted in the portfolio_allocation_fact table.
//
// Co-authored by: GitHub Copilot and OpenCode
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
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
				DELETE FROM portfolio_allocation_fact 
				WHERE portfolio_id = 1 AND observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag = '202505'
				)`,
				nil,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:NEW'", nil).
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag = '202505'", nil).
			Build(t),
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
			    a.external_data,
			    p.observation_time_id,
			    o.observation_time_tag,
				o.observation_timestamp
			FROM portfolio_allocation_fact p 
			JOIN asset a ON p.asset_id = a.id
			JOIN portfolio_allocation_obs_time o ON p.portfolio_id = o.portfolio_id AND p.observation_time_id = o.id
		WHERE p.portfolio_id=%s AND o.portfolio_id=p.portfolio_id AND o.observation_time_tag='%s'
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
			"external_data":         inttestutil.NullAssertableNullString(),
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
			"total_market_value":    inttestutil.ToAssertableNullString("10000"),
			"name":                  inttestutil.ToAssertableNullString("iShares 0-5 Year TIPS Bond ETF"),
			"ticker":                inttestutil.ToAssertableNullString("ARCA:STIP"),
			"external_data":         inttestutil.NullAssertableNullString(),
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
			"external_data":         inttestutil.NullAssertableNullString(),
			"observation_time_id":   inttestutil.NotNullAssertableNullString(),
			"observation_time_tag":  inttestutil.ToAssertableNullString("202505"),
			"observation_timestamp": inttestutil.ToAssertableNullString("2025-05-01T00:00:00Z"),
		},
	}

	inttestutil.AssertDBWithQueryMultipleRows(t, allocationHistoryQuery, expectedRecords)
}

// TestPostPortfolioAllocationHistoryInsertExplicitTimestamp verifies that an explicit offset-bearing
// observation timestamp is persisted as the same UTC instant.
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func TestPostPortfolioAllocationHistoryInsertExplicitTimestamp(t *testing.T) {

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": 0,
				"timeTag": "2025-05-01T02:00:00+02:00",
				"timestamp": "2025-05-01T02:00:00+02:00"
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
				WHERE portfolio_id = 1 AND observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE '%T%'
				)`,
				nil,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:NEW'", nil).
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE '%T%'", nil).
			Build(t),
	)

	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
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
			"external_data":       inttestutil.NullAssertableNullString(),
			"observation_time_id": inttestutil.NotNullAssertableNullString(),
			"observation_time_tag": inttestutil.ToAssertableNullStringWithAssertion(
				func(t *testing.T, actual sql.NullString) {
					assert.Equal(t, "2025-05-01T02:00:00+02:00", actual.String)
				},
			),
			"observation_timestamp": inttestutil.ToAssertableNullStringWithAssertion(
				func(t *testing.T, actual sql.NullString) {
					dateTime, dateErr := time.Parse(time.RFC3339Nano, actual.String)
					assert.NoError(t, dateErr)
					assert.Equal(t, "2025-05-01T00:00:00Z", dateTime.Format(time.RFC3339))
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
			    a.external_data,
			    p.observation_time_id,
			    o.observation_time_tag,
				o.observation_timestamp
			FROM portfolio_allocation_fact p 
			JOIN asset a ON p.asset_id = a.id
			JOIN portfolio_allocation_obs_time o ON p.portfolio_id = o.portfolio_id AND p.observation_time_id = o.id
			WHERE p.portfolio_id=%s AND o.observation_time_tag LIKE '%s' 
			ORDER BY p.asset_id`,
		portfolioIdString,
		"%T%",
	)

	inttestutil.AssertDBWithQueryMultipleRows(t, allocationHistoryQuery, expectedRecords)
}

// TestPostPortfolioAllocationHistoryFullMerge verifies that posting a complete
// snapshot merges the target observation while leaving data from other
// observations unchanged.
//
// Co-authored by: OpenCode and Igor Benicio de Mesquita
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
			   9
			),
		    (
			   7,
			   'STOCKS',
			   FALSE,
			   10,
			   9,
			   90,
			   1,
			   9
			)
		;
	`

	err := inttestinfra.ExecuteDBQuery(insertAllocationHistorySQL, nil)
	assert.NoError(t, err)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
				DELETE FROM portfolio_allocation_fact 
				WHERE portfolio_id = 1 AND observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND id = 9
				)`,
				nil,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:NEW'", nil).
			Build(t),
	)

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": "9",
				"timeTag": "202504",
				"timestamp": "2025-04-01T00:00:00Z"
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
	assert.NoError(t, err)
	assert.NotNil(t, response)

	if response == nil {
		return
	}
	defer deferCloseResponseBody(response)

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
			"external_data":         inttestutil.NullAssertableNullString(),
			"observation_time_id":   inttestutil.ToAssertableNullString("9"),
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
			"external_data":         inttestutil.NullAssertableNullString(),
			"observation_time_id":   inttestutil.ToAssertableNullString("9"),
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
			"external_data":         inttestutil.NullAssertableNullString(),
			"observation_time_id":   inttestutil.ToAssertableNullString("9"),
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
			"external_data":         inttestutil.NullAssertableNullString(),
			"observation_time_id":   inttestutil.ToAssertableNullString("9"),
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
			"external_data":         inttestutil.NullAssertableNullString(),
			"observation_time_id":   inttestutil.ToAssertableNullString("9"),
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
			    a.external_data,
			    p.observation_time_id,
			    o.observation_time_tag,
				o.observation_timestamp
			FROM portfolio_allocation_fact p 
			JOIN asset a ON p.asset_id = a.id
			JOIN portfolio_allocation_obs_time o ON p.portfolio_id = o.portfolio_id AND p.observation_time_id = o.id
			WHERE p.portfolio_id=%s AND o.id=%d
			ORDER BY p.asset_id, p.class`,
		portfolioIdString,
		9,
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
		9,
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

// TestPostPortfolioAllocationValidations verifies existing portfolio snapshot validation cases with explicit timestamps.
//
// Co-authored by: OpenCode and benizzio
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
						"id": "9",
						"timeTag": "202504",
						"timestamp": "2025-04-01T00:00:00Z"
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
						"id": "9",
						"timeTag": "202504",
						"timestamp": "2025-04-01T00:00:00Z"
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
						"id": "9",
						"timeTag": "202504",
						"timestamp": "2025-04-01T00:00:00Z"
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
						"id": "9",
						"timeTag": "202504",
						"timestamp": "2025-04-01T00:00:00Z"
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

// TestPostPortfolioAllocationRequiresExplicitTimestamp rejects missing and invalid observation timestamps on creates and updates.
// It asserts directly against the database that no observation or allocation is partially persisted.
//
// Co-authored by: OpenCode
func TestPostPortfolioAllocationRequiresExplicitTimestamp(t *testing.T) {
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1 AND (class = 'INVALID_TIMESTAMP_TEST' OR observation_time_id IN (SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE 'invalid-observation-date-create-%'))`,
				nil,
			).
			AddCleanupQuery(
				`UPDATE portfolio_allocation_obs_time
				 SET observation_time_tag = '202503', observation_timestamp = '2025-03-01 00:00:00'
					 WHERE id = 2 AND portfolio_id = 1`,
				nil,
			).
			AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE 'invalid-observation-date-create-%'`, nil).
			Build(t),
	)

	var invalidTimestamps = []struct {
		name     string
		property string
	}{
		{name: "missing"},
		{name: "null", property: `"timestamp":null,`},
		{name: "empty", property: `"timestamp":"",`},
		{name: "malformed", property: `"timestamp":"not-a-timestamp",`},
		{name: "zero", property: `"timestamp":"0001-01-01T00:00:00Z",`},
	}

	for _, invalidTimestamp := range invalidTimestamps {
		for _, update := range []bool{false, true} {
			var observationId = 0
			var recordClass = "INVALID_TIMESTAMP_TEST"
			var tag = "invalid-observation-date-create-" + invalidTimestamp.name
			if update {
				observationId = 2
				tag = "invalid-observation-date-update-" + invalidTimestamp.name
			}

			var requestBody = fmt.Sprintf(`{
				"observationTimestamp": {
					"id":%d,
					%s"timeTag":%q
				},
				"allocations":[{
					"assetId":1,
					"class":%q,
					"cashReserve":false,
					"totalMarketValue":"123"
				}]
			}`, observationId, invalidTimestamp.property, tag, recordClass)
			var responseBody = postPortfolioAllocationForValidationFailure(t, requestBody)
			assert.NotEmpty(t, responseBody)
			if invalidTimestamp.name == "malformed" {
				assert.JSONEq(
					t,
					`{"errorMessage":"Validation failed","details":["Field 'observationTimestamp.timestamp' failed validation: must be a valid RFC3339 timestamp"]}`,
					string(responseBody),
				)
			}

			inttestutil.AssertDBWithQueryMultipleRows(
				t,
				`SELECT count(*) AS fact_count FROM portfolio_allocation_fact WHERE class = 'INVALID_TIMESTAMP_TEST'`,
				[]inttestutil.AssertableNullStringMap{{"fact_count": inttestutil.ToAssertableNullString("0")}},
			)
			inttestutil.AssertDBWithQueryMultipleRows(
				t,
				`SELECT count(*) AS observation_count FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag LIKE 'invalid-observation-date-create-%'`,
				[]inttestutil.AssertableNullStringMap{{"observation_count": inttestutil.ToAssertableNullString("0")}},
			)
			if update {
				inttestutil.AssertDBWithQueryMultipleRows(
					t,
					`SELECT observation_time_tag,
					        to_char(observation_timestamp, 'YYYY-MM-DD"T"HH24:MI:SS') AS observation_timestamp
					 FROM portfolio_allocation_obs_time
					 WHERE id = 2 AND portfolio_id = 1`,
					[]inttestutil.AssertableNullStringMap{{
						"observation_time_tag":  inttestutil.ToAssertableNullString("202503"),
						"observation_timestamp": inttestutil.ToAssertableNullString("2025-03-01T00:00:00"),
					}},
				)
			}
		}
	}

	var missingNewObservationTag = postPortfolioAllocationForValidationFailure(t, `{
		"observationTimestamp":{"id":0,"timestamp":"2025-03-05T00:00:00Z"},
		"allocations":[{"assetId":1,"class":"INVALID_TIMESTAMP_TEST","cashReserve":false,"totalMarketValue":"123"}]
	}`)
	assert.JSONEq(
		t,
		`{"errorMessage":"Validation failed","details":["Field 'observationTimestamp.timeTag' failed validation: is required for a new observation"]}`,
		string(missingNewObservationTag),
	)
}

func postPortfolioAllocationForValidationFailure(t *testing.T, postPortfolioJSON string) []byte {

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/history",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
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
		INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		VALUES 
			(2, 'test_obs_1', '2025-10-01 00:00:00'::TIMESTAMP),
			(2, 'test_obs_2', '2025-10-02 00:00:00'::TIMESTAMP),
			(2, 'test_obs_3', '2025-10-03 00:00:00'::TIMESTAMP),
			(2, 'test_obs_4', '2025-10-04 00:00:00'::TIMESTAMP),
			(2, 'test_obs_5', '2025-10-05 00:00:00'::TIMESTAMP),
			(2, 'test_obs_6', '2025-10-06 00:00:00'::TIMESTAMP),
			(2, 'test_obs_7', '2025-10-07 00:00:00'::TIMESTAMP),
			(2, 'test_obs_8', '2025-10-08 00:00:00'::TIMESTAMP),
			(2, 'test_obs_9', '2025-10-09 00:00:00'::TIMESTAMP),
			(2, 'test_obs_10', '2025-10-10 00:00:00'::TIMESTAMP),
			(2, 'test_obs_11', '2025-10-11 00:00:00'::TIMESTAMP),
			(2, 'test_obs_12', '2025-10-12 00:00:00'::TIMESTAMP),
			(2, 'test_obs_13', '2025-10-13 00:00:00'::TIMESTAMP),
			(2, 'test_obs_14', '2025-10-14 00:00:00'::TIMESTAMP),
			(2, 'test_obs_15', '2025-10-15 00:00:00'::TIMESTAMP)
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
		WHERE portfolio_id = 2 AND observation_time_tag ~ '^test_obs_'
		;
	`

	err := inttestinfra.ExecuteDBQuery(setupSQL, nil)
	assert.NoError(t, err)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact 
				WHERE portfolio_id = 2 AND observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 2 AND observation_time_tag ~ '^test_obs_'
				)`,
				nil,
			).
			AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 2 AND observation_time_tag ~ '^test_obs_'`, nil).
			Build(t),
	)

	// Act: Get portfolio history for portfolio 2 with the default limit of 10 observations
	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/2/history")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
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
	defer deferCloseResponseBody(response1)
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
	defer deferCloseResponseBody(response)
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
	defer deferCloseResponseBody(response)
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
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
		"errorMessage": "Validation failed",
		"details": ["Field 'allocations[0].assetName' failed validation: must not exceed 100"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// TestPostPortfolioAllocationHistoryInsertOnlyWithExternalAsset verifies that a
// new asset persists its external asset payload and that the history endpoint
// projects that persisted value back in the response.
//
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryInsertOnlyWithExternalAsset(t *testing.T) {

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"timeTag": "202509",
				"timestamp": "2025-09-01T00:00:00Z"
			},
			"allocations": [
				{
					"assetName": "External Asset",
					"assetTicker": "Test:EXT",
					"externalAsset": {
						"source": "YAHOO_FINANCE",
						"ticker": "IAU",
						"exchangeId": "PCX",
						"name": "Ignored Name",
						"exchangeName": "Ignored Exchange"
					},
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
		inttestinfra.TestAPIURLPrefix+"/portfolio/2/history",
		"application/json",
		strings.NewReader(postPortfolioSnapshotJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 2 AND observation_time_id IN (SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 2 AND observation_time_tag = '202509')`,
				nil,
			).
			AddCleanupQuery("DELETE FROM asset WHERE ticker = 'Test:EXT'", nil).
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 2 AND observation_time_tag = '202509'", nil).
			Build(t),
	)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	var persistedAssetId string
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		"SELECT id, ticker, name, external_data FROM asset WHERE ticker = 'Test:EXT'",
		[]inttestutil.AssertableNullStringMap{
			{
				"id":     inttestutil.NotNullValueCapturingAssertableNullString(&persistedAssetId),
				"ticker": inttestutil.ToAssertableNullString("Test:EXT"),
				"name":   inttestutil.ToAssertableNullString("External Asset"),
				"external_data": inttestutil.ToAssertableNullStringWithAssertion(
					func(t *testing.T, actual sql.NullString) {
						assert.True(t, actual.Valid)
						assert.JSONEq(t, testExternalAssetDataJSON, actual.String)
					},
				),
			},
		},
	)

	var historyQuery = dbx.Params{"ticker": "Test:EXT", "timeTag": "202509"}
	var persistedObservationTimestampId string
	err = inttestinfra.FetchWithDBQuery(
		`SELECT o.id FROM portfolio_allocation_fact p JOIN asset a ON a.id = p.asset_id JOIN portfolio_allocation_obs_time o ON o.portfolio_id = p.portfolio_id AND o.id = p.observation_time_id WHERE p.portfolio_id = 2 AND a.ticker = {:ticker} AND o.observation_time_tag = {:timeTag}`,
		historyQuery,
		func(rows *dbx.Rows) error {
			return rows.Scan(&persistedObservationTimestampId)
		},
	)
	assert.NoError(t, err)

	response, err = http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/2/history?observationTimestampId=" + persistedObservationTimestampId)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err = io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expectedResponseJSON = `
		[
			{
				"observationTimestamp" : {
					"id": ` + persistedObservationTimestampId + `,
					"timeTag": "202509",
					"timestamp": "2025-09-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": ` + persistedAssetId + `,
						"assetName":"External Asset",
						"assetTicker":"Test:EXT",
						"externalAsset": {
							"source": "YAHOO_FINANCE",
							"ticker": "IAU",
							"exchangeId": "PCX"
						},
						"class":"STOCKS",
						"cashReserve":false,
						"assetMarketPrice":"100",
						"assetQuantity":"20",
						"totalMarketValue":"2000"
					}
				],
				"totalMarketValue":"2000"
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}

// TestPostPortfolioAllocationHistoryDoesNotOverwriteExistingAssetExternalData
// verifies that posting a portfolio allocation for an existing asset preserves
// the asset's already persisted external data.
//
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryDoesNotOverwriteExistingAssetExternalData(t *testing.T) {
	var originalExternalData = capturePersistedAssetExternalData(t, 1)

	err := inttestinfra.ExecuteDBQuery(
		`UPDATE asset SET external_data = '{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU","exchangeId":"PCX"}]}'::jsonb WHERE id = 1`,
		nil,
	)
	assert.NoError(t, err)

	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder().
				AddCleanupQuery(
					"DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 2 AND observation_time_id = 8 AND asset_id = 1",
					nil,
				),
			1,
			originalExternalData,
		).Build(t),
	)

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": "8",
				"timeTag": "202504",
				"timestamp": "2025-04-01T00:00:00Z"
			},
			"allocations": [
				{
					"assetId": "1",
					"assetName": "Ignored",
					"assetTicker": "Ignored",
					"externalAsset": {
						"source": "YAHOO_FINANCE",
						"ticker": "SHOULD_NOT_PERSIST",
						"exchangeId": "SHOULD_NOT_PERSIST"
					},
					"class": "BONDS",
					"cashReserve": false,
					"assetQuantity": "1",
					"assetMarketPrice": "100",
					"totalMarketValue": "100"
				}
			]
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/2/history",
		"application/json",
		strings.NewReader(postPortfolioSnapshotJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	assertPersistedAssetWithExternalData(
		t,
		1,
		"ARCA:BIL",
		"SPDR Bloomberg 1-3 Month T-Bill ETF",
		new(testExternalAssetDataJSON),
	)
}

// TestPostPortfolioAllocationHistoryValidation_InvalidExternalAssetFields
// verifies that incomplete external asset payloads are rejected with the
// current validation error structure.
//
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryValidation_InvalidExternalAssetFields(t *testing.T) {

	var postPortfolioSnapshotJSON = `
		{
			"observationTimestamp": {
				"id": "9",
				"timeTag": "202504",
				"timestamp": "2025-04-01T00:00:00Z"
			},
			"allocations": [
				{
					"assetName": "External Asset",
					"assetTicker": "Test:VALIDATION",
					"externalAsset": {
						"ticker": "IAU"
					},
					"class": "STOCKS",
					"totalMarketValue": "1000"
				}
			]
		}
	`

	actualResponseJSON := string(postPortfolioAllocationForValidationFailure(t, postPortfolioSnapshotJSON))

	var expectedResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'allocations[0].externalAsset.source' failed validation: is required",
				"Field 'allocations[0].externalAsset.exchangeId' failed validation: is required"
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestPostPortfolioAllocationHistoryUpdatesObservationMetadata verifies metadata updates retain the observation and
// allocation identities while persisting tag and timestamp changes in the owning portfolio.
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryUpdatesObservationMetadata(t *testing.T) {
	registerObservationFixtureRestore(t, 2, 1, "202503", "2025-03-01 00:00:00")

	var tagUpdate = portfolioHistoryUpdateJSON(
		2,
		"202503-edited",
		"2025-03-01T00:00:00Z",
		`{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"100.00009","assetMarketPrice":"100","totalMarketValue":"10000"}`,
	)
	response, err := postPortfolioHistoryRequest(1, tagUpdate)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	require.Equal(t, http.StatusNoContent, response.StatusCode)
	assertPersistedObservationMetadata(t, 2, 1, "202503-edited", "2025-03-01T00:00:00.000000")
	assertPersistedObservationAllocation(t, 2, 1)

	var timestampUpdate = portfolioHistoryUpdateJSON(
		2,
		"202503-edited",
		"2025-03-02T12:34:56.123456Z",
		`{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"100.00009","assetMarketPrice":"100","totalMarketValue":"10000"}`,
	)
	response, err = postPortfolioHistoryRequest(1, timestampUpdate)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	require.Equal(t, http.StatusNoContent, response.StatusCode)
	assertPersistedObservationMetadata(t, 2, 1, "202503-edited", "2025-03-02T12:34:56.123456")
	assertPersistedObservationAllocation(t, 2, 1)

	var omittedTagUpdate = `{
		"observationTimestamp":{"id":2,"timestamp":"2025-03-02T12:34:56.123456Z"},
		"allocations":[{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"100.00009","assetMarketPrice":"100","totalMarketValue":"10000"}]
	}`
	response, err = postPortfolioHistoryRequest(1, omittedTagUpdate)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	require.Equal(t, http.StatusNoContent, response.StatusCode)
	assertPersistedObservationMetadata(t, 2, 1, "202503-edited", "2025-03-02T12:34:56.123456")
	assertPersistedObservationAllocation(t, 2, 1)
}

// TestPostPortfolioAllocationHistoryScopesTagsAndUpdatesToPortfolio verifies portfolio-local tag uniqueness and
// rejects attempts to update an observation owned by another portfolio.
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryScopesTagsAndUpdatesToPortfolio(t *testing.T) {
	var duplicateTagRequest = portfolioHistoryUpdateJSON(
		2,
		"202501",
		"2025-03-03T00:00:00Z",
		`{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"100.00009","assetMarketPrice":"100","totalMarketValue":"10000"}`,
	)
	response, err := postPortfolioHistoryRequest(1, duplicateTagRequest)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusConflict, response.StatusCode)
	var duplicateTagResponse, readError = io.ReadAll(response.Body)
	require.NoError(t, readError)
	assert.JSONEq(
		t,
		`{"errorMessage":"Observation tag already exists in this portfolio.","details":["Choose a different observationTimestamp.timeTag."]}`,
		string(duplicateTagResponse),
	)

	var foreignObservationRequest = portfolioHistoryUpdateJSON(
		7,
		"foreign-edit",
		"2025-03-04T00:00:00Z",
		`{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"1","assetMarketPrice":"100","totalMarketValue":"100"}`,
	)
	response, err = postPortfolioHistoryRequest(1, foreignObservationRequest)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	var foreignObservationResponse, foreignReadError = io.ReadAll(response.Body)
	require.NoError(t, foreignReadError)
	assert.JSONEq(
		t,
		`{"errorMessage":"Observation validation failed","details":["Observation does not exist in the requested portfolio"]}`,
		string(foreignObservationResponse),
	)

	assertPersistedObservationMetadata(t, 2, 1, "202503", "2025-03-01T00:00:00.000000")
	assertPersistedObservationMetadata(t, 7, 3, "202503", "2025-03-01T00:00:00.000000")
}

// TestPostPortfolioAllocationHistoryAllowsSameTagAcrossPortfolios verifies tags are unique per owner, not globally.
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryAllowsSameTagAcrossPortfolios(t *testing.T) {
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact
				 WHERE portfolio_id = 3 AND observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 3 AND observation_time_tag = '202501'
				 )`,
				nil,
			).
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 3 AND observation_time_tag = '202501'`,
				nil,
			).
			Build(t),
	)

	var createRequest = portfolioHistoryUpdateJSON(
		0,
		"202501",
		"2025-01-01T00:00:00Z",
		`{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"1","assetMarketPrice":"100","totalMarketValue":"100"}`,
	)
	response, err := postPortfolioHistoryRequest(3, createRequest)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	require.Equal(t, http.StatusNoContent, response.StatusCode)

	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		`SELECT portfolio_id::text AS portfolio_id, observation_time_tag, COUNT(*)::text AS observation_count
		 FROM portfolio_allocation_obs_time
		 WHERE observation_time_tag = '202501'
		 GROUP BY portfolio_id, observation_time_tag
		 ORDER BY portfolio_id`,
		[]inttestutil.AssertableNullStringMap{
			{
				"portfolio_id":         inttestutil.ToAssertableNullString("1"),
				"observation_time_tag": inttestutil.ToAssertableNullString("202501"),
				"observation_count":    inttestutil.ToAssertableNullString("1"),
			},
			{
				"portfolio_id":         inttestutil.ToAssertableNullString("3"),
				"observation_time_tag": inttestutil.ToAssertableNullString("202501"),
				"observation_count":    inttestutil.ToAssertableNullString("1"),
			},
		},
	)
}

// TestPostPortfolioAllocationHistoryRollsBackMetadataAndAssets verifies a persistence failure rolls back an earlier
// metadata update and asset insertion within the same transaction.
// Authored by: OpenCode
func TestPostPortfolioAllocationHistoryRollsBackMetadataAndAssets(t *testing.T) {
	registerObservationFixtureRestore(t, 2, 1, "202503", "2025-03-01 00:00:00")
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(`DELETE FROM asset WHERE ticker = 'Test:ROLLBACK'`, nil).
			AddCleanupQuery(
				`DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1 AND class LIKE 'ROLLBACK_TEST%'`,
				nil,
			).
			Build(t),
	)

	var requestBody = `{
		"observationTimestamp":{"id":2,"timeTag":"rollback-attempt","timestamp":"2025-03-05T00:00:00Z"},
		"allocations":[
			{"assetName":"Rollback Asset","assetTicker":"Test:ROLLBACK","class":"ROLLBACK_TEST_NEW","cashReserve":false,"assetQuantity":"1","assetMarketPrice":"100","totalMarketValue":"100"},
			{"assetId":999999,"class":"ROLLBACK_TEST_INVALID","cashReserve":false,"assetQuantity":"1","assetMarketPrice":"100","totalMarketValue":"100"}
		]
	}`
	response, err := postPortfolioHistoryRequest(1, requestBody)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusInternalServerError, response.StatusCode)

	assertPersistedObservationMetadata(t, 2, 1, "202503", "2025-03-01T00:00:00.000000")
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		`SELECT COUNT(*)::text AS asset_count FROM asset WHERE ticker = 'Test:ROLLBACK'`,
		[]inttestutil.AssertableNullStringMap{{"asset_count": inttestutil.ToAssertableNullString("0")}},
	)
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		`SELECT COUNT(*)::text AS fact_count FROM portfolio_allocation_fact WHERE portfolio_id = 1 AND class LIKE 'ROLLBACK_TEST%'`,
		[]inttestutil.AssertableNullStringMap{{"fact_count": inttestutil.ToAssertableNullString("0")}},
	)
}

// registerObservationFixtureRestore restores an owned observation after a test mutates fixture metadata.
// Authored by: OpenCode
func registerObservationFixtureRestore(t *testing.T, observationId, portfolioId int64, tag, timestamp string) {
	t.Helper()
	t.Cleanup(func() {
		var err = inttestinfra.ExecuteDBQuery(
			`UPDATE portfolio_allocation_obs_time
			 SET observation_time_tag = {:tag}, observation_timestamp = {:timestamp}
			 WHERE id = {:id} AND portfolio_id = {:portfolioId}`,
			dbx.Params{
				"id":          observationId,
				"portfolioId": portfolioId,
				"tag":         tag,
				"timestamp":   timestamp,
			},
		)
		assert.NoError(t, err)
	})
}

// assertPersistedObservationMetadata checks the persisted identity, owner, tag, and timestamp at PostgreSQL precision.
// Authored by: OpenCode
func assertPersistedObservationMetadata(
	t *testing.T,
	observationId, portfolioId int64,
	tag, timestamp string,
) {
	t.Helper()
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		`SELECT id::text AS id, portfolio_id::text AS portfolio_id, observation_time_tag,
		        to_char(observation_timestamp, 'YYYY-MM-DD"T"HH24:MI:SS.US') AS observation_timestamp
		 FROM portfolio_allocation_obs_time WHERE id = `+strconv.FormatInt(observationId, 10),
		[]inttestutil.AssertableNullStringMap{{
			"id":                    inttestutil.ToAssertableNullString(strconv.FormatInt(observationId, 10)),
			"portfolio_id":          inttestutil.ToAssertableNullString(strconv.FormatInt(portfolioId, 10)),
			"observation_time_tag":  inttestutil.ToAssertableNullString(tag),
			"observation_timestamp": inttestutil.ToAssertableNullString(timestamp),
		}},
	)
}

// assertPersistedObservationAllocation checks that a metadata-only edit leaves the portfolio facts unchanged.
// Authored by: OpenCode
func assertPersistedObservationAllocation(t *testing.T, observationId, portfolioId int64) {
	t.Helper()
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		`SELECT COUNT(*)::text AS fact_count,
		        SUM(asset_quantity)::numeric(18,8)::text AS asset_quantity,
		        SUM(asset_market_price)::numeric(18,8)::text AS asset_market_price,
		        SUM(total_market_value)::text AS total_market_value
		 FROM portfolio_allocation_fact
		 WHERE observation_time_id = `+strconv.FormatInt(observationId, 10)+
			` AND portfolio_id = `+strconv.FormatInt(portfolioId, 10),
		[]inttestutil.AssertableNullStringMap{{
			"fact_count":         inttestutil.ToAssertableNullString("1"),
			"asset_quantity":     inttestutil.ToAssertableNullString("100.00009000"),
			"asset_market_price": inttestutil.ToAssertableNullString("100.00000000"),
			"total_market_value": inttestutil.ToAssertableNullString("10000"),
		}},
	)
}

// portfolioHistoryUpdateJSON builds a full snapshot request with one allocation for observation update tests.
// Authored by: OpenCode
func portfolioHistoryUpdateJSON(observationId int, tag, timestamp, allocationJSON string) string {
	return fmt.Sprintf(
		`{"observationTimestamp":{"id":%d,"timeTag":%q,"timestamp":%q},"allocations":[%s]}`,
		observationId,
		tag,
		timestamp,
		allocationJSON,
	)
}

// postPortfolioHistoryRequest posts a JSON snapshot to the portfolio history endpoint.
// Authored by: OpenCode
func postPortfolioHistoryRequest(portfolioId int64, requestBody string) (*http.Response, error) {
	var request, err = http.NewRequest(
		http.MethodPost,
		inttestinfra.TestAPIURLPrefix+"/portfolio/"+strconv.FormatInt(portfolioId, 10)+"/history",
		strings.NewReader(requestBody),
	)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	return http.DefaultClient.Do(request)
}
