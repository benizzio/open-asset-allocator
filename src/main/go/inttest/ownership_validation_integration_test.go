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

// TestPostPortfolioHistoryRejectsWhitespaceTags exercises trimmed-string emptiness on creates and updates.
// Run with go test -count=1 -run '^TestPostPortfolioHistoryRejectsWhitespaceTags$' ./inttest.
// Authored by: GPT-6.1 Sol
func TestPostPortfolioHistoryRejectsWhitespaceTags(t *testing.T) {
	for _, observationId := range []int{0, 2} {
		for _, tag := range []string{" \t\r\n", "\u00a0\u2003"} {
			t.Run(fmt.Sprintf("id=%d/tag=%q", observationId, tag), func(t *testing.T) {
				registerObservationFixtureRestore(t, 2, 1, "202503", "2025-03-01 00:00:00")
				t.Cleanup(inttestutil.BuildCleanupFunctionBuilder().
					AddCleanupQuery(`DELETE FROM portfolio_allocation_fact WHERE observation_time_id IN (
						SELECT id FROM portfolio_allocation_obs_time WHERE portfolio_id = 1
						AND observation_time_tag = {:tag} AND id <> 2)`, dbx.Params{"tag": tag}).
					AddCleanupQuery(`DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = 1
						AND observation_time_tag = {:tag} AND id <> 2`, dbx.Params{"tag": tag}).Build(t))
				var body = portfolioHistoryUpdateJSON(observationId, tag, "2025-03-01T00:00:00Z",
					`{"assetId":1,"class":"BONDS","cashReserve":false,"totalMarketValue":"10000"}`)
				response, err := postPortfolioHistoryRequest(1, body)
				require.NoError(t, err)
				defer deferCloseResponseBody(response)
				assert.Equal(t, http.StatusBadRequest, response.StatusCode)
				responseBody, err := io.ReadAll(response.Body)
				require.NoError(t, err)
				assert.JSONEq(t, `{"errorMessage":"Validation failed","details":[
					"Field 'observationTimestamp.timeTag' failed validation: must not contain only whitespace"
				]}`, string(responseBody))
				assertPersistedObservationMetadata(t, 2, 1, "202503", "2025-03-01T00:00:00.000000")
				assertPersistedObservationAllocation(t, 2, 1)
			})
		}
	}
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
