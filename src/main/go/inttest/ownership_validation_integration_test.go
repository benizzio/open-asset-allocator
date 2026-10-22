package inttest

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// TestPostPortfolioHistoryOwnershipValidation verifies domain ownership failures retain the existing API contract.
// Run with go test -count=1 -run '^TestPostPortfolioHistoryOwnershipValidation$' ./inttest.
// Authored by: GPT-6.1 Sol
func TestPostPortfolioHistoryOwnershipValidation(t *testing.T) {
	var testCases = []struct {
		name          string
		portfolioId   int64
		observationId int
		detail        string
	}{
		{"missing observation", 1, 999999, "Observation does not exist in the requested portfolio"},
		{"foreign observation", 1, 7, "Observation does not exist in the requested portfolio"},
		{"missing portfolio", 999999, 0, "Portfolio does not exist"},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Cleanup(inttestutil.BuildCleanupFunctionBuilder().
				AddCleanupQuery(`DELETE FROM portfolio_allocation_fact WHERE observation_time_id IN (
					SELECT id FROM portfolio_allocation_obs_time WHERE observation_time_tag = 'ownership-rejected')`, nil).
				AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time
					WHERE observation_time_tag = 'ownership-rejected'`, nil).Build(t))
			var body = portfolioHistoryUpdateJSON(testCase.observationId, "ownership-rejected", "2025-03-01T00:00:00Z",
				`{"assetId":1,"class":"BONDS","cashReserve":false,"totalMarketValue":"10000"}`)
			response, err := postPortfolioHistoryRequest(testCase.portfolioId, body)
			require.NoError(t, err)
			defer deferCloseResponseBody(response)
			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
			responseBody, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			assert.JSONEq(t, fmt.Sprintf(`{"errorMessage":"Observation validation failed","details":[%q]}`,
				testCase.detail), string(responseBody))
			assertPersistedObservationMetadata(t, 7, 3, "202503", "2025-03-01T00:00:00.000000")
			inttestutil.AssertDBWithQueryMultipleRows(t,
				`SELECT COUNT(*)::text AS count FROM portfolio_allocation_obs_time
				 WHERE observation_time_tag = 'ownership-rejected'`,
				[]inttestutil.AssertableNullStringMap{{"count": inttestutil.ToAssertableNullString("0")}})
		})
	}
}

// TestPostPortfolioHistoryRequiresNonblankTag verifies required, whitespace, and length rules for creates and updates.
// Run with go test -count=1 -run '^TestPostPortfolioHistoryRequiresNonblankTag$' ./inttest.
// Authored by: GPT-6 Luna
func TestPostPortfolioHistoryRequiresNonblankTag(t *testing.T) {
	var invalidTags = []struct {
		name           string
		tag            string
		omitTag        bool
		nullTag        bool
		expectedDetail string
	}{
		{name: "missing", omitTag: true, expectedDetail: "Field 'observationTimestamp.timeTag' failed validation: is required"},
		{name: "null", nullTag: true, expectedDetail: "Field 'observationTimestamp.timeTag' failed validation: is required"},
		{name: "empty", expectedDetail: "Field 'observationTimestamp.timeTag' failed validation: is required"},
		{name: "ascii-whitespace", tag: " \t\r\n", expectedDetail: "Field 'observationTimestamp.timeTag' failed validation: must not contain only whitespace"},
		{name: "unicode-whitespace", tag: "\u00a0\u2003", expectedDetail: "Field 'observationTimestamp.timeTag' failed validation: must not contain only whitespace"},
		{name: "over-100-characters", tag: strings.Repeat("x", 101), expectedDetail: "Field 'observationTimestamp.timeTag' failed validation: must not exceed 100"},
	}
	var fixtureTimestamp = "2048-03-04T05:06:07.123456Z"
	var fixtureTimestampSQL = "2048-03-04 05:06:07.123456+00:00"

	for _, observationId := range []int{0, 2} {
		var currentObservationId = observationId
		var mode = "create"
		if currentObservationId != 0 {
			mode = "update"
		}
		for _, invalidTag := range invalidTags {
			var currentInvalidTag = invalidTag
			t.Run(mode+"/"+currentInvalidTag.name, func(t *testing.T) {
				var suffix = fmt.Sprintf("%s_%s", mode, currentInvalidTag.name)
				var allocationClass = "PR7_TAG_VALIDATION_" + suffix
				var assetTicker = "PR7:TAG:" + suffix
				var assetName = "PR7 tag validation " + suffix
				var cleanupParams = dbx.Params{
					"class":     allocationClass,
					"ticker":    assetTicker,
					"timestamp": fixtureTimestampSQL,
				}
				var cleanupBuilder = inttestutil.BuildCleanupFunctionBuilder().AddCleanupQuery(
					`DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1 AND class = {:class}`,
					cleanupParams,
				)
				if currentObservationId == 2 {
					cleanupBuilder.AddCleanupQuery(
						`DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1 AND observation_time_id = 2`, nil,
					)
				}
				cleanupBuilder.
					AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1
						AND observation_timestamp = {:timestamp} AND id <> 2`, cleanupParams).
					AddCleanupQuery(`DELETE FROM asset WHERE ticker = {:ticker}`, cleanupParams)
				if currentObservationId == 2 {
					cleanupBuilder.
						AddCleanupQuery(`INSERT INTO portfolio_allocation_fact
							(asset_id, class, cash_reserve, asset_quantity, asset_market_price, total_market_value,
							 portfolio_id, observation_time_id)
							SELECT 1, 'BONDS', FALSE, 100.00009, 100, 10000, 1, 2
							WHERE EXISTS (SELECT 1 FROM portfolio_allocation_obs_time WHERE id = 2 AND portfolio_id = 1)`, nil).
						AddCleanupQuery(`UPDATE portfolio_allocation_obs_time
							SET observation_time_tag = '202503', observation_timestamp = '2025-03-01 00:00:00'
							WHERE id = 2 AND portfolio_id = 1`, nil)
				}
				t.Cleanup(cleanupBuilder.Build(t))

				var tagProperty string
				if currentInvalidTag.nullTag {
					tagProperty = `"timeTag":null,`
				} else if !currentInvalidTag.omitTag {
					tagProperty = fmt.Sprintf(`"timeTag":%q,`, currentInvalidTag.tag)
				}
				var allocationJSON = fmt.Sprintf(
					`{"assetName":%q,"assetTicker":%q,"class":%q,"cashReserve":false,"assetQuantity":"2","assetMarketPrice":"123","totalMarketValue":"246"}`,
					assetName,
					assetTicker,
					allocationClass,
				)
				var requestBody = fmt.Sprintf(
					`{"observationTimestamp":{"id":%d,%s"timestamp":%q},"allocations":[%s]}`,
					currentObservationId,
					tagProperty,
					fixtureTimestamp,
					allocationJSON,
				)

				if currentObservationId == 2 {
					assertPersistedObservationMetadata(t, 2, 1, "202503", "2025-03-01T00:00:00.000000")
					assertPersistedObservationAllocation(t, 2, 1)
				}
				response, err := postPortfolioHistoryRequest(1, requestBody)
				require.NoError(t, err)
				defer deferCloseResponseBody(response)
				var responseBody, readError = io.ReadAll(response.Body)
				require.NoError(t, readError)
				assert.Equal(t, http.StatusBadRequest, response.StatusCode, "%s", responseBody)
				assert.JSONEq(t, fmt.Sprintf(`{"errorMessage":"Validation failed","details":[%q]}`,
					currentInvalidTag.expectedDetail), string(responseBody))

				if currentObservationId == 2 {
					assertPersistedObservationMetadata(t, 2, 1, "202503", "2025-03-01T00:00:00.000000")
					assertPersistedObservationAllocation(t, 2, 1)
				} else {
					assertPortfolioHistoryTagValidationCount(t,
						`SELECT COUNT(*) FROM portfolio_allocation_obs_time
						 WHERE portfolio_id = 1 AND observation_timestamp = {:timestamp}`,
						dbx.Params{"timestamp": fixtureTimestampSQL},
						0,
					)
				}
				assertPortfolioHistoryTagValidationCount(t,
					`SELECT COUNT(*) FROM portfolio_allocation_fact
					 WHERE portfolio_id = 1 AND class = {:class}`,
					cleanupParams,
					0,
				)
				assertPortfolioHistoryTagValidationCount(t,
					`SELECT COUNT(*) FROM asset WHERE ticker = {:ticker}`,
					cleanupParams,
					0,
				)
			})
		}
	}
}

// TestPostPortfolioHistoryAcceptsValidTagBoundaries verifies character-count limits and verbatim tag persistence.
// Run with go test -count=1 -run '^TestPostPortfolioHistoryAcceptsValidTagBoundaries$' ./inttest.
// Authored by: GPT-6 Luna
func TestPostPortfolioHistoryAcceptsValidTagBoundaries(t *testing.T) {
	var validTags = []struct {
		name string
		tag  string
	}{
		{name: "ascii-100", tag: strings.Repeat("a", 100)},
		{name: "unicode-100", tag: strings.Repeat("界", 100)},
		{name: "nonblank-padded", tag: " padded tag "},
	}
	for _, observationId := range []int{0, 2} {
		var currentObservationId = observationId
		var mode = "create"
		if currentObservationId != 0 {
			mode = "update"
		}
		for _, validTag := range validTags {
			var currentValidTag = validTag
			t.Run(mode+"/"+currentValidTag.name, func(t *testing.T) {
				var fixtureTimestamp = "2048-03-04T05:06:07.123456Z"
				var cleanupParams = dbx.Params{"tag": currentValidTag.tag}
				var cleanupBuilder = inttestutil.BuildCleanupFunctionBuilder()
				if currentObservationId == 0 {
					cleanupBuilder.
						AddCleanupQuery(`DELETE FROM portfolio_allocation_fact WHERE portfolio_id = 1 AND observation_time_id IN (
							SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag = {:tag})`, cleanupParams).
						AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1
							AND observation_time_tag = {:tag}`, cleanupParams)
				} else {
					cleanupBuilder.
						AddCleanupQuery(`UPDATE portfolio_allocation_obs_time
							SET observation_time_tag = '202503', observation_timestamp = '2025-03-01 00:00:00'
							WHERE id = 2 AND portfolio_id = 1`, nil)
				}
				t.Cleanup(cleanupBuilder.Build(t))

				var body = portfolioHistoryUpdateJSON(currentObservationId, currentValidTag.tag, fixtureTimestamp,
					`{"assetId":1,"class":"BONDS","cashReserve":false,"assetQuantity":"100.00009","assetMarketPrice":"100","totalMarketValue":"10000"}`)
				response, err := postPortfolioHistoryRequest(1, body)
				require.NoError(t, err)
				defer deferCloseResponseBody(response)
				assert.Equal(t, http.StatusNoContent, response.StatusCode)

				if currentObservationId == 0 {
					var persistedPortfolioId string
					var persistedTag string
					var persistedTimestamp string
					var matchingObservations = 0
					var readError = inttestinfra.FetchWithDBQuery(
						`SELECT portfolio_id::text AS portfolio_id, observation_time_tag,
						        to_char(observation_timestamp, 'YYYY-MM-DD"T"HH24:MI:SS.US') AS observation_timestamp
						 FROM portfolio_allocation_obs_time WHERE portfolio_id = 1 AND observation_time_tag = {:tag}`,
						cleanupParams,
						func(rows *dbx.Rows) error {
							matchingObservations++
							return rows.Scan(&persistedPortfolioId, &persistedTag, &persistedTimestamp)
						},
					)
					require.NoError(t, readError)
					assert.Equal(t, 1, matchingObservations)
					assert.Equal(t, "1", persistedPortfolioId)
					assert.Equal(t, currentValidTag.tag, persistedTag)
					assert.Equal(t, "2048-03-04T05:06:07.123456", persistedTimestamp)
				} else {
					assertPersistedObservationMetadata(t, 2, 1, currentValidTag.tag, "2048-03-04T05:06:07.123456")
					assertPersistedObservationAllocation(t, 2, 1)
				}
			})
		}
	}
}

// assertPortfolioHistoryTagValidationCount checks a parameterized persistence count in observation-tag tests.
// Example: assertPortfolioHistoryTagValidationCount(t, query, dbx.Params{"tag": "Quarter-end"}, 0).
// Authored by: GPT-6 Luna
func assertPortfolioHistoryTagValidationCount(t *testing.T, query string, params dbx.Params, expected int64) {
	t.Helper()
	var actual int64
	var err = inttestinfra.FetchWithDBQuery(query, params, func(rows *dbx.Rows) error {
		return rows.Scan(&actual)
	})
	require.NoError(t, err)
	assert.Equal(t, expected, actual)
}

// TestPostAllocationPlanOwnershipValidation verifies unknown plans and invalid persisted detail IDs are domain errors.
// Run with go test -count=1 -run '^TestPostAllocationPlanOwnershipValidation$' ./inttest.
// Authored by: GPT-6.1 Sol
func TestPostAllocationPlanOwnershipValidation(t *testing.T) {
	var details = `[
		{"id":30,"hierarchicalId":[null,"BONDS"],"sliceSizePercentage":"0.5"},
		{"id":31,"hierarchicalId":[null,"STOCKS"],"sliceSizePercentage":"0.5"},
		{"id":32,"hierarchicalId":["ARCA:BIL","BONDS"],"sliceSizePercentage":"1.0","asset":{"id":1}},
		{"id":33,"hierarchicalId":["ARCA:SPY","STOCKS"],"sliceSizePercentage":"1.0","asset":{"id":7}}
	]`
	var invalidDetailsMessage = "One or more planned allocation detail IDs do not belong to this plan"
	var testCases = []struct {
		name    string
		planId  int
		details string
		message string
	}{
		{"missing plan", 999999, details, "Allocation plan does not exist in the requested portfolio"},
		{"create with persisted details", 0, details, "New allocation plans cannot reuse existing detail IDs"},
		{"missing detail", 6, strings.Replace(details, `"id":33`, `"id":999999`, 1), invalidDetailsMessage},
		{"duplicate detail", 6, strings.Replace(details, `"id":33`, `"id":32`, 1), invalidDetailsMessage},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			registerAllocationPlanFixtureRestore(t)
			t.Cleanup(inttestutil.BuildCleanupFunctionBuilder().
				AddCleanupQuery(`DELETE FROM planned_allocation WHERE allocation_plan_id IN (
					SELECT id FROM allocation_plan WHERE portfolio_id = 5 AND name = 'Rejected plan' AND id <> 6)`, nil).
				AddCleanupQuery(`DELETE FROM allocation_plan WHERE portfolio_id = 5
					AND name = 'Rejected plan' AND id <> 6`, nil).Build(t))
			var body = fmt.Sprintf(`{"id":%d,"name":"Rejected plan","plannedDate":"2025-06-02","details":%s}`,
				testCase.planId, testCase.details)
			response, err := postAllocationPlanRequest(5, body)
			require.NoError(t, err)
			defer deferCloseResponseBody(response)
			assert.Equal(t, http.StatusBadRequest, response.StatusCode)
			responseBody, err := io.ReadAll(response.Body)
			require.NoError(t, err)
			assert.JSONEq(t, fmt.Sprintf(`{"errorMessage":"Allocation plan validation failed","details":[%q]}`,
				testCase.message), string(responseBody))
			assertPersistedAllocationPlanMetadata(t, 6, "Update Allocation Plan Fixture", "2025-06-01")
			assertPersistedAllocationPlanDetailIds(t, 6)
			inttestutil.AssertDBWithQueryMultipleRows(t,
				`SELECT COUNT(*)::text AS count FROM allocation_plan WHERE portfolio_id = 5 AND name = 'Rejected plan'`,
				[]inttestutil.AssertableNullStringMap{{"count": inttestutil.ToAssertableNullString("0")}})
		})
	}
}

// TestPostAllocationPlanRejectsWhitespaceName checks Unicode-trimmed emptiness before plan persistence.
// Run with go test -count=1 -run '^TestPostAllocationPlanRejectsWhitespaceName$' ./inttest.
// Authored by: GPT-6.1 Sol
func TestPostAllocationPlanRejectsWhitespaceName(t *testing.T) {
	t.Cleanup(func() {
		var err = inttestinfra.ExecuteDBQuery(`DELETE FROM planned_allocation WHERE allocation_plan_id IN (
			SELECT id FROM allocation_plan WHERE portfolio_id = 5 AND btrim(name, U&'\00A0\2003') = '')`, nil)
		assert.NoError(t, err)
		err = inttestinfra.ExecuteDBQuery(`DELETE FROM allocation_plan WHERE portfolio_id = 5
			AND btrim(name, U&'\00A0\2003') = ''`, nil)
		assert.NoError(t, err)
	})
	var body = `{"name":"\u00a0\u2003","plannedDate":"2025-06-01","details":[
		{"hierarchicalId":[null,"BONDS"],"sliceSizePercentage":"1.0"},
		{"hierarchicalId":["ARCA:BIL","BONDS"],"sliceSizePercentage":"1.0","asset":{"id":1}}
	]}`
	response, err := postAllocationPlanRequest(5, body)
	require.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	responseBody, err := io.ReadAll(response.Body)
	require.NoError(t, err)
	assert.JSONEq(t, `{"errorMessage":"Allocation plan validation failed","details":[
		"Allocation plan name is required and must not be blank"
	]}`, string(responseBody))
	inttestutil.AssertDBWithQueryMultipleRows(t,
		`SELECT COUNT(*)::text AS count FROM allocation_plan WHERE btrim(name, U&'\00A0\2003') = ''`,
		[]inttestutil.AssertableNullStringMap{{"count": inttestutil.ToAssertableNullString("0")}})
}
