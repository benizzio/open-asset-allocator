package inttest

import (
	"strings"
	"testing"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/domain/allocation"
	"github.com/benizzio/open-asset-allocator/domain/infra/repository"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans seeds enough records to inspect keyset index use.
// Run with go test -run TestHistoricalReadPaginationIndexesHaveRepresentativeQueryPlans ./inttest.
// TODO: Remove this temporary, explicitly approved white-box EXPLAIN test at a later date.
// Co-authored by: GPT-6.1 Sol and OpenCode
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

	var adapter = buildPaginationCaptureAdapter(t)
	var observations = repository.BuildPortfolioAllocationRepository(adapter)
	var allocationPlans = repository.BuildAllocationPlanRepository(adapter)
	var observationCursor = &domain.ObservationPageCursor{
		Timestamp: time.Date(2040, time.January, 1, 1, 40, 0, 0, time.UTC),
	}
	_, err = observations.FindAvailableObservationTimestampsPage(portfolio.Id, observationCursor)
	require.NoError(t, err)
	var observationPlan = explainPaginationQuery(t, takePaginationQuery(t))
	t.Logf("Observation-header production query plan:\n%s", observationPlan)
	assert.Contains(t, observationPlan, "portfolio_allocation_obs_time_page_idx")
	assert.Contains(t, observationPlan, "portfolio_allocation_fact_observation_page_idx")

	_, err = observations.FindPortfolioAllocationsByObservationPage(portfolio.Id, observationCursor)
	require.NoError(t, err)
	var historyPlan = explainPaginationQuery(t, takePaginationQuery(t))
	t.Logf("Allocation-history production query plan:\n%s", historyPlan)
	assert.Contains(t, historyPlan, "portfolio_allocation_obs_time_page_idx")
	assert.Contains(t, historyPlan, "portfolio_allocation_fact_observation_page_idx")

	_, err = allocationPlans.GetAllocationPlansPage(portfolio.Id, allocation.AssetAllocationPlan,
		&domain.AllocationPlanPageCursor{
			PlannedDate: time.Date(2040, time.April, 10, 0, 0, 0, 0, time.UTC),
		})
	require.NoError(t, err)
	var allocationPlanPlan = explainPaginationQuery(t, takePaginationQuery(t))
	t.Logf("Allocation-plan production query plan:\n%s", allocationPlanPlan)
	assert.Contains(t, allocationPlanPlan, "allocation_plan_page_idx")
}

// explainPaginationQuery analyzes the exact executable SQL and bindings captured from a production read.
// Co-authored by: GPT-6.1 Sol and OpenCode
func explainPaginationQuery(t *testing.T, query capturedPaginationQuery) string {
	t.Helper()
	var planLines = make([]string, 0)
	var rows, err = inttestinfra.DatabaseConnection.DB().Query(
		"EXPLAIN (ANALYZE, BUFFERS) "+query.sql, query.args...)
	require.NoError(t, err)
	defer rows.Close()
	for rows.Next() {
		var planLine string
		require.NoError(t, rows.Scan(&planLine))
		planLines = append(planLines, planLine)
	}
	require.NoError(t, rows.Err())
	return strings.Join(planLines, "\n")
}
