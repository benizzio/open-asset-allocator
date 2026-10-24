package inttest

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestHistoricalReadQueryParameterPresence verifies encoded keys and empty selectors through existing HTTP endpoints.
// For example, an encoded empty pageCursor selects a page, while a malformed encoded cursor yields one validation JSON.
// Authored by: GPT-6.1 Sol
func TestHistoricalReadQueryParameterPresence(t *testing.T) {
	var portfolio = insertTestPortfolio(t, "Historical read query parameter presence")
	for _, endpoint := range []string{
		allocationPlanEndpoint(portfolio.Id),
		observationEndpoint(portfolio.Id, "history"),
		observationEndpoint(portfolio.Id, "history/observation"),
	} {
		t.Run(endpoint, func(t *testing.T) {
			for _, query := range []string{"page%43ursor=", "pageCursor"} {
				t.Run(query, func(t *testing.T) {
					var response = getPaginationResponse(t, endpoint+"?"+query)
					defer closePaginationResponse(t, response)
					require.Equal(t, http.StatusOK, response.StatusCode)
					var body, err = io.ReadAll(response.Body)
					require.NoError(t, err)
					assert.JSONEq(t, `{"items":[],"nextCursor":null}`, string(body))
				})
			}
			for _, query := range []string{"page%43ursor=%ZZ", "pageCursor=&page%43ursor="} {
				t.Run(query, func(t *testing.T) {
					var response = getPaginationResponse(t, endpoint+"?"+query)
					assertPaginationValidationResponse(t, response)
				})
			}
		})
	}
	for _, query := range []string{"observationTimestampId=", "observationTimestamp%49d="} {
		t.Run(query, func(t *testing.T) {
			var endpoint = fmt.Sprintf("%s?pageCursor=&%s", observationEndpoint(portfolio.Id, "history"), query)
			var response = getPaginationResponse(t, endpoint)
			assertPaginationValidationResponse(t, response)
		})
	}
}
