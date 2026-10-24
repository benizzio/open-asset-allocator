package inttest

import (
	"strings"
	"testing"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans seeds enough records to inspect keyset index use.
// Authored by: OpenCode
func TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans(t *testing.T) {
	var portfolio = insertTestPortfolio(t, "PR 9 pagination query plan inspection")
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery("DELETE FROM portfolio_allocation_fact WHERE portfolio_id = {:portfolioId}", dbx.Params{
				"portfolioId": portfolio.Id,
			}).
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = {:portfolioId}", dbx.Params{
				"portfolioId": portfolio.Id,
			}).
			AddCleanupQuery(
				`DELETE FROM planned_allocation WHERE allocation_plan_id IN (
					SELECT id FROM allocation_plan WHERE portfolio_id = {:portfolioId}
				)`,
				dbx.Params{"portfolioId": portfolio.Id},
			).
			AddCleanupQuery("DELETE FROM allocation_plan WHERE portfolio_id = {:portfolioId}", dbx.Params{
				"portfolioId": portfolio.Id,
			}).
			Build(t),
	)

	var seedSize = 12000
	var seedObservationsSQL = `
		INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		SELECT {:portfolioId}, 'pr9-explain-observation-' || observation_number::text,
			TIMESTAMPTZ '2040-01-01 00:00:00Z' + observation_number * INTERVAL '1 second'
		FROM generate_series(1, {:seedSize}) AS observation_number
	`
	var err = inttestinfra.ExecuteDBQuery(seedObservationsSQL, dbx.Params{
		"portfolioId": portfolio.Id,
		"seedSize":    seedSize,
	})
	require.NoError(t, err)
	err = inttestinfra.ExecuteDBQuery(
		`INSERT INTO portfolio_allocation_fact (
			asset_id, "class", cash_reserve, asset_quantity, asset_market_price,
			total_market_value, portfolio_id, observation_time_id
		)
		SELECT 1, 'BONDS', FALSE, 1, 100, 100, observation.portfolio_id, observation.id
		FROM portfolio_allocation_obs_time observation
		WHERE observation.portfolio_id = {:portfolioId}`,
		dbx.Params{"portfolioId": portfolio.Id},
	)
	require.NoError(t, err)
	err = inttestinfra.ExecuteDBQuery(
		`INSERT INTO allocation_plan (portfolio_id, name, type, planned_date)
		SELECT {:portfolioId}, 'pr9-explain-plan-' || plan_number::text, 'ALLOCATION_PLAN',
			DATE '2040-01-01' + plan_number
		FROM generate_series(1, {:seedSize}) AS plan_number`,
		dbx.Params{"portfolioId": portfolio.Id, "seedSize": seedSize},
	)
	require.NoError(t, err)
	err = inttestinfra.ExecuteDBQuery(
		`INSERT INTO planned_allocation (
			allocation_plan_id, hierarchical_id, cash_reserve, slice_size_percentage
		)
		SELECT id, ARRAY['EXPLAIN'], FALSE, 0.5
		FROM allocation_plan WHERE portfolio_id = {:portfolioId}`,
		dbx.Params{"portfolioId": portfolio.Id},
	)
	require.NoError(t, err)
	for _, tableName := range []string{
		"portfolio_allocation_obs_time",
		"portfolio_allocation_fact",
		"allocation_plan",
		"planned_allocation",
	} {
		err = inttestinfra.ExecuteDBQuery("ANALYZE "+tableName, nil)
		require.NoError(t, err)
	}

	var observationPlan = explainPaginationQuery(t, `
		SELECT observation.id
		FROM portfolio_allocation_obs_time observation
		WHERE observation.portfolio_id = {:portfolioId}
		  AND (observation.observation_timestamp, observation.id) <
		      ({:cursorTimestamp}, {:cursorId})
		  AND EXISTS (
			SELECT 1 FROM portfolio_allocation_fact fact
			WHERE fact.portfolio_id = observation.portfolio_id
			  AND fact.observation_time_id = observation.id
		  )
		ORDER BY observation.observation_timestamp DESC, observation.id DESC
		LIMIT 11`,
		dbx.Params{
			"portfolioId":     portfolio.Id,
			"cursorTimestamp": "2040-01-01 01:40:00Z",
			"cursorId":        0,
		},
	)
	var allocationPlanPlan = explainPaginationQuery(t, `
		SELECT plan.id
		FROM allocation_plan plan
		WHERE plan.portfolio_id = {:portfolioId}
		  AND plan.type = 'ALLOCATION_PLAN'
		  AND (plan.planned_date, plan.id) < ({:cursorDate}, {:cursorId})
		  AND EXISTS (
			SELECT 1 FROM planned_allocation detail WHERE detail.allocation_plan_id = plan.id
		  )
		ORDER BY plan.planned_date DESC, plan.id DESC
		LIMIT 11`,
		dbx.Params{
			"portfolioId": portfolio.Id,
			"cursorDate":  time.Date(2040, time.April, 10, 0, 0, 0, 0, time.UTC),
			"cursorId":    0,
		},
	)
	t.Logf("Observation keyset query plan:\n%s", observationPlan)
	t.Logf("Allocation-plan keyset query plan:\n%s", allocationPlanPlan)
	assert.Contains(t, observationPlan, "portfolio_allocation_obs_time_page_idx")
	assert.Contains(t, observationPlan, "portfolio_allocation_fact_observation_page_idx")
	assert.Contains(t, allocationPlanPlan, "allocation_plan_page_idx")
}

// explainPaginationQuery returns an analyzed execution plan for one parameterized keyset query.
// Authored by: OpenCode
func explainPaginationQuery(t *testing.T, query string, params dbx.Params) string {
	t.Helper()
	var planLines = make([]string, 0)
	var err = inttestinfra.FetchWithDBQuery(
		"EXPLAIN (ANALYZE, BUFFERS) "+query,
		params,
		func(rows *dbx.Rows) error {
			var planLine string
			if scanError := rows.Scan(&planLine); scanError != nil {
				return scanError
			}
			planLines = append(planLines, planLine)
			return nil
		},
	)
	require.NoError(t, err)
	return strings.Join(planLines, "\n")
}
