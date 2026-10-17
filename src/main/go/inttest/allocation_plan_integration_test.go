package inttest

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// TestGetAllocationPlans verifies that the allocation-plan API returns the initialized portfolio plan.
// Usage: go test -count=1 -run '^TestGetAllocationPlans$' ./inttest
// Co-authored by: OpenCode and Igor Benicio de Mesquita
func TestGetAllocationPlans(t *testing.T) {

	var response, err = http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-plan")
	require.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		[
			{
				"id":1,
				"name":"60/40 Portfolio Classic - Example",
				"type":"ALLOCATION_PLAN",
				"plannedDate":"2025-01-01",
				"details":[
					{
						"id":7,
						"hierarchicalId":["NasdaqGM:SHV","STOCKS"],
						"cashReserve":true,
						"sliceSizePercentage":"0.5",
						"asset":{ "id":5, "name":"iShares Short Treasury Bond ETF", "ticker":"NasdaqGM:SHV" }
					},
					{
						"id":2,
						"hierarchicalId":[null,"BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.6"
					},
					{
						"id":9,
						"hierarchicalId":["ARCA:SPY","STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.45",
						"asset":{ "id":7, "name":"SPDR S&P 500 ETF Trust", "ticker":"ARCA:SPY" }
					},
					{
						"id":1,
						"hierarchicalId":[null,"STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.4"
					},
					{
						"id":3,
						"hierarchicalId":["ARCA:BIL","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.4",
						"asset":{ "id":1, "name":"SPDR Bloomberg 1-3 Month T-Bill ETF", "ticker":"ARCA:BIL" }
					},
					{
						"id":4,
						"hierarchicalId":["NasdaqGM:IEF","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.3",
						"asset":{ "id":3, "name":"iShares 7-10 Year Treasury Bond ETF", "ticker":"NasdaqGM:IEF" }
					},
					{
						"id":5,
						"hierarchicalId":["NasdaqGM:TLT","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.2",
						"asset":{ "id":4, "name":"iShares 20+ Year Treasury Bond ETF", "ticker":"NasdaqGM:TLT" }
					},
					{
						"id":6,
						"hierarchicalId":["ARCA:STIP","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.1",
						"asset":{ "id":2, "name":"iShares 0-5 Year TIPS Bond ETF", "ticker":"ARCA:STIP" }
					},
					{
						"id":8,
						"hierarchicalId":["ARCA:EWZ","STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.05",
						"asset":{ "id":6, "name":"iShares Msci Brazil ETF", "ticker":"ARCA:EWZ" }
					}
				]
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}

// TestAllocationPlanReadsOrderByPlannedDateAndId verifies planned-date ordering and its ID tie-breaker
// on both complete-plan and plan-identifier read APIs. It seeds distinct dates and equal-date records directly.
//
// Co-authored by: OpenCode
func TestAllocationPlanReadsOrderByPlannedDateAndId(t *testing.T) {
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM planned_allocation WHERE allocation_plan_id IN (SELECT id FROM allocation_plan WHERE name LIKE 'Historical date ordering test%')`,
				nil,
			).
			AddCleanupQuery(`DELETE FROM allocation_plan WHERE name LIKE 'Historical date ordering test%'`, nil).
			Build(t),
	)

	var insertPlansSQL = `
		INSERT INTO allocation_plan (portfolio_id, name, type, planned_date)
		VALUES
			(1, 'Historical date ordering test future', 'ALLOCATION_PLAN', '2099-01-01'),
			(1, 'Historical date ordering test tie first', 'ALLOCATION_PLAN', '2040-05-15'),
			(1, 'Historical date ordering test tie second', 'ALLOCATION_PLAN', '2040-05-15'),
			(1, 'Historical date ordering test past', 'ALLOCATION_PLAN', '2001-01-01')
	`
	err := inttestinfra.ExecuteDBQuery(insertPlansSQL, nil)
	require.NoError(t, err)

	var insertDetailsSQL = `
		INSERT INTO planned_allocation (
			allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value
		)
		SELECT id, '{NULL, "BONDS"}', NULL, FALSE, 1, NULL
		FROM allocation_plan
		WHERE name LIKE 'Historical date ordering test%'
	`
	err = inttestinfra.ExecuteDBQuery(insertDetailsSQL, nil)
	require.NoError(t, err)

	var assertNameOrder = func(responseBody string) {
		var orderedNames = []string{
			"Historical date ordering test future",
			"Historical date ordering test tie second",
			"Historical date ordering test tie first",
			"60/40 Portfolio Classic - Example",
			"Historical date ordering test past",
		}
		var previousIndex = -1
		for _, name := range orderedNames {
			var nameIndex = strings.Index(responseBody, name)
			assert.Greater(t, nameIndex, previousIndex, "unexpected order for %q", name)
			previousIndex = nameIndex
		}
	}

	planResponse, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-plan")
	require.NoError(t, err)
	defer deferCloseResponseBody(planResponse)
	require.Equal(t, http.StatusOK, planResponse.StatusCode)
	planBody, err := io.ReadAll(planResponse.Body)
	require.NoError(t, err)
	assertNameOrder(string(planBody))
	assert.Contains(t, string(planBody), `"plannedDate":"2040-05-15"`)
	assert.Contains(t, string(planBody), `"plannedDate":"2099-01-01"`)
	assert.Contains(t, string(planBody), `"plannedDate":"2001-01-01"`)

	optionsResponse, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/divergence/options")
	require.NoError(t, err)
	defer deferCloseResponseBody(optionsResponse)
	require.Equal(t, http.StatusOK, optionsResponse.StatusCode)
	optionsBody, err := io.ReadAll(optionsResponse.Body)
	require.NoError(t, err)
	assertNameOrder(string(optionsBody))
	assert.Contains(t, string(optionsBody), `"plannedDate":"2040-05-15"`)
	assert.Contains(t, string(optionsBody), `"plannedDate":"2099-01-01"`)
	assert.Contains(t, string(optionsBody), `"plannedDate":"2001-01-01"`)
}

// TestAllocationPlanDateMigrationBackfillsLegacyRows executes the migration against a temporary legacy table.
// It verifies creation-date backfill, CURRENT_DATE fallback, and the migration's NOT NULL/no-default contract.
//
// Co-authored by: OpenCode
func TestAllocationPlanDateMigrationBackfillsLegacyRows(t *testing.T) {
	var transaction, err = inttestinfra.DatabaseConnection.DB().Begin()
	require.NoError(t, err)
	t.Cleanup(func() {
		assert.NoError(t, transaction.Rollback())
	})

	_, err = transaction.Exec(`
		CREATE TEMPORARY TABLE allocation_plan (
			id BIGINT PRIMARY KEY,
			create_timestamp TIMESTAMP
		) ON COMMIT DROP
	`)
	require.NoError(t, err)

	_, err = transaction.Exec(`
		INSERT INTO allocation_plan (id, create_timestamp)
		VALUES (1, '2021-02-03 18:45:00'), (2, NULL)
	`)
	require.NoError(t, err)

	var migrationPath = filepath.Join(
		"..", "..", "flyway", "sql", "migration-[15]-allocation_plan_planned_date.sql",
	)
	migrationSQL, err := os.ReadFile(migrationPath)
	require.NoError(t, err)
	for _, statement := range strings.Split(string(migrationSQL), ";") {
		if strings.TrimSpace(statement) == "" {
			continue
		}
		_, err = transaction.Exec(statement)
		require.NoError(t, err)
	}

	var actualCreationDate, actualNullTimestampDate, currentDate string
	err = transaction.QueryRow(`
		SELECT
			(SELECT planned_date::text FROM allocation_plan WHERE id = 1),
			(SELECT planned_date::text FROM allocation_plan WHERE id = 2),
			CURRENT_DATE::text
	`).Scan(&actualCreationDate, &actualNullTimestampDate, &currentDate)
	require.NoError(t, err)
	assert.Equal(t, "2021-02-03", actualCreationDate)
	assert.Equal(t, currentDate, actualNullTimestampDate)

	var nullable string
	var defaultValue sql.NullString
	err = transaction.QueryRow(`
		SELECT is_nullable, column_default
		FROM information_schema.columns
		WHERE table_schema LIKE 'pg_temp_%'
		  AND table_name = 'allocation_plan'
		  AND column_name = 'planned_date'
	`).Scan(&nullable, &defaultValue)
	require.NoError(t, err)
	assert.Equal(t, "NO", nullable)
	assert.False(t, defaultValue.Valid)
}

// TestPostAllocationPlanForInsertion verifies a plan with an explicit plannedDate persists and is returned by reads.
//
// Co-authored by: OpenCode
func TestPostAllocationPlanForInsertion(t *testing.T) {

	var newAllocationPlanJSON = `
		{
			"name":"New Allocation Plan Test DELETE",
			"plannedDate":"2025-08-19",
			"details":[
				{
					"hierarchicalId":[null,"STOCKS"],
					"cashReserve":false,
					"sliceSizePercentage":"0.6",
					"asset":null
				},
				{
					"hierarchicalId":[null,"BONDS"],
					"sliceSizePercentage":"0.4"
				},
				{
					"hierarchicalId":["ARCA:SPY","STOCKS"],
					"cashReserve":false,
					"sliceSizePercentage":"0.5",
					"asset":{ "id":7, "name":"SPDR S&P 500 ETF Trust", "ticker":"ARCA:SPY" }
				},
				null,
				{
					"hierarchicalId":["TEST:ALTBOND","STOCKS"],
					"cashReserve":true,
					"sliceSizePercentage":"0.5",
					"asset":{ "name":"Test Cash Reserve Bond DELETE", "ticker":"TEST:ALTBOND" }
				},
				{
					"hierarchicalId":["ARCA:BIL","BONDS"],
					"cashReserve":false,
					"sliceSizePercentage":"0.4",
					"asset":{ "id":1, "name":"SPDR Bloomberg 1-3 Month T-Bill ETF", "ticker":"ARCA:BIL" }
				},
				{
					"hierarchicalId":["TEST:ALTBOND2","BONDS"],
					"cashReserve":false,
					"sliceSizePercentage":"0.6",
					"asset":{ "name":"Test Bond 2 DELETE", "ticker":"TEST:ALTBOND2" }
				}
			]
		}
	`

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
				DELETE FROM planned_allocation
				WHERE allocation_plan_id IN (SELECT id FROM allocation_plan WHERE name LIKE '%DELETE')
				`,
				nil,
			).
			AddCleanupQuery(`DELETE FROM allocation_plan WHERE name LIKE '%DELETE'`, nil).
			AddCleanupQuery(`DELETE FROM asset WHERE name LIKE '%DELETE'`, nil).
			Build(t),
	)

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
		"application/json",
		strings.NewReader(newAllocationPlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	readResponse, readErr := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-plan")
	assert.NoError(t, readErr)
	defer deferCloseResponseBody(readResponse)
	readBody, readErr := io.ReadAll(readResponse.Body)
	assert.NoError(t, readErr)
	assert.Contains(t, string(readBody), `"plannedDate":"2025-08-19"`)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	var actualAllocationPlanQuery = `
		SELECT
		    ap.id,
		    ap.name,
		    ap.type,
		    ap.planned_execution_date,
	    ap.planned_date::text AS planned_date,
		    ap.portfolio_id
		FROM allocation_plan ap
		WHERE ap.name LIKE '%DELETE'
	`

	var allocationPlanIdString string

	var expectedRecords = []inttestutil.AssertableNullStringMap{
		{
			"id":                     inttestutil.NotNullValueCapturingAssertableNullString(&allocationPlanIdString),
			"name":                   inttestutil.ToAssertableNullString("New Allocation Plan Test DELETE"),
			"type":                   inttestutil.ToAssertableNullString("ALLOCATION_PLAN"),
			"planned_execution_date": inttestutil.NullAssertableNullString(),
			"planned_date":           inttestutil.ToAssertableNullString("2025-08-19"),
			"portfolio_id":           inttestutil.ToAssertableNullString("1"),
		},
	}

	inttestutil.AssertDBWithQueryMultipleRows(t, actualAllocationPlanQuery, expectedRecords)

	var actualPlannedAllocationsQuery = `
		SELECT
		    pa.id,
		    pa.allocation_plan_id,
		    pa.hierarchical_id,
		    pa.cash_reserve,
		    pa.slice_size_percentage,
		    pa.total_market_value,
		    pa.asset_id,
		    ass.ticker AS asset_ticker,
		    ass.name AS asset_name
		FROM planned_allocation pa
		JOIN allocation_plan ap ON pa.allocation_plan_id = ap.id
		LEFT JOIN asset ass ON ass.id = pa.asset_id
		WHERE ap.name LIKE '%DELETE'
		ORDER BY pa.hierarchical_id ASC
	`

	expectedRecords = []inttestutil.AssertableNullStringMap{
		{
			"id":                    inttestutil.NotNullAssertableNullString(),
			"allocation_plan_id":    inttestutil.ToAssertableNullString(allocationPlanIdString),
			"hierarchical_id":       inttestutil.ToAssertableNullString("{ARCA:BIL,BONDS}"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"slice_size_percentage": inttestutil.ToAssertableNullString("0.40000"),
			"total_market_value":    inttestutil.NullAssertableNullString(),
			"asset_id":              inttestutil.ToAssertableNullString("1"),
			"asset_ticker":          inttestutil.ToAssertableNullString("ARCA:BIL"),
			"asset_name":            inttestutil.ToAssertableNullString("SPDR Bloomberg 1-3 Month T-Bill ETF"),
		},
		{
			"id":                    inttestutil.NotNullAssertableNullString(),
			"allocation_plan_id":    inttestutil.ToAssertableNullString(allocationPlanIdString),
			"hierarchical_id":       inttestutil.ToAssertableNullString("{ARCA:SPY,STOCKS}"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"slice_size_percentage": inttestutil.ToAssertableNullString("0.50000"),
			"total_market_value":    inttestutil.NullAssertableNullString(),
			"asset_id":              inttestutil.ToAssertableNullString("7"),
			"asset_ticker":          inttestutil.ToAssertableNullString("ARCA:SPY"),
			"asset_name":            inttestutil.ToAssertableNullString("SPDR S&P 500 ETF Trust"),
		},
		{
			"id":                    inttestutil.NotNullAssertableNullString(),
			"allocation_plan_id":    inttestutil.ToAssertableNullString(allocationPlanIdString),
			"hierarchical_id":       inttestutil.ToAssertableNullString("{TEST:ALTBOND,STOCKS}"),
			"cash_reserve":          inttestutil.ToAssertableNullString("true"),
			"slice_size_percentage": inttestutil.ToAssertableNullString("0.50000"),
			"total_market_value":    inttestutil.NullAssertableNullString(),
			"asset_id":              inttestutil.NotNullAssertableNullString(),
			"asset_ticker":          inttestutil.ToAssertableNullString("TEST:ALTBOND"),
			"asset_name":            inttestutil.ToAssertableNullString("Test Cash Reserve Bond DELETE"),
		},
		{
			"id":                    inttestutil.NotNullAssertableNullString(),
			"allocation_plan_id":    inttestutil.ToAssertableNullString(allocationPlanIdString),
			"hierarchical_id":       inttestutil.ToAssertableNullString("{TEST:ALTBOND2,BONDS}"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"slice_size_percentage": inttestutil.ToAssertableNullString("0.60000"),
			"total_market_value":    inttestutil.NullAssertableNullString(),
			"asset_id":              inttestutil.NotNullAssertableNullString(),
			"asset_ticker":          inttestutil.ToAssertableNullString("TEST:ALTBOND2"),
			"asset_name":            inttestutil.ToAssertableNullString("Test Bond 2 DELETE"),
		},
		{
			"id":                    inttestutil.NotNullAssertableNullString(),
			"allocation_plan_id":    inttestutil.ToAssertableNullString(allocationPlanIdString),
			"hierarchical_id":       inttestutil.ToAssertableNullString("{NULL,BONDS}"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"slice_size_percentage": inttestutil.ToAssertableNullString("0.40000"),
			"total_market_value":    inttestutil.NullAssertableNullString(),
			"asset_id":              inttestutil.NullAssertableNullString(),
			"asset_ticker":          inttestutil.NullAssertableNullString(),
			"asset_name":            inttestutil.NullAssertableNullString(),
		},
		{
			"id":                    inttestutil.NotNullAssertableNullString(),
			"allocation_plan_id":    inttestutil.ToAssertableNullString(allocationPlanIdString),
			"hierarchical_id":       inttestutil.ToAssertableNullString("{NULL,STOCKS}"),
			"cash_reserve":          inttestutil.ToAssertableNullString("false"),
			"slice_size_percentage": inttestutil.ToAssertableNullString("0.60000"),
			"total_market_value":    inttestutil.NullAssertableNullString(),
			"asset_id":              inttestutil.NullAssertableNullString(),
			"asset_ticker":          inttestutil.NullAssertableNullString(),
			"asset_name":            inttestutil.NullAssertableNullString(),
		},
	}

	inttestutil.AssertDBWithQueryMultipleRows(t, actualPlannedAllocationsQuery, expectedRecords)
}

// TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName updates a pre-seeded plan
// and asserts an existing asset name is not overwritten, slice sizes are updated, and a new
// planned allocation is inserted as part of the update operation. It also updates the plan name
// and verifies the rename works, then cleans up to restore the original fixture state.
//
// Co-authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName(t *testing.T) {

	// 1) Fetch the pre-seeded plan id
	var fetchPlanSQL = `
		SELECT ap.id, ap.name
		FROM allocation_plan ap
		WHERE ap.portfolio_id = 5 AND ap.name = 'Update Allocation Plan Fixture'
	`
	var allocationPlanIdString string
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		fetchPlanSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":   inttestutil.NotNullValueCapturingAssertableNullString(&allocationPlanIdString),
				"name": inttestutil.ToAssertableNullString("Update Allocation Plan Fixture"),
			},
		},
	)

	// 2) Fetch planned allocation ids for this plan so update can match existing rows
	var idBil, idSpy, idBondsTop, idStocksTop string
	var fetchPlannedAllocationsSQL = `
		SELECT pa.id, pa.hierarchical_id
		FROM planned_allocation pa
		WHERE pa.allocation_plan_id = ` + allocationPlanIdString + `
		ORDER BY pa.hierarchical_id
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		fetchPlannedAllocationsSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idBil),
				"hierarchical_id": inttestutil.ToAssertableNullString("{ARCA:BIL,BONDS}"),
			},
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idSpy),
				"hierarchical_id": inttestutil.ToAssertableNullString("{ARCA:SPY,STOCKS}"),
			},
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idBondsTop),
				"hierarchical_id": inttestutil.ToAssertableNullString("{NULL,BONDS}"),
			},
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idStocksTop),
				"hierarchical_id": inttestutil.ToAssertableNullString("{NULL,STOCKS}"),
			},
		},
	)

	// Register cleanup to restore fixture state after update
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM planned_allocation WHERE allocation_plan_id = {:allocationPlanId} AND hierarchical_id = '{TEST:ALTBOND,BONDS}'`,
				dbx.Params{"allocationPlanId": allocationPlanIdString},
			).
			AddCleanupQuery(`UPDATE planned_allocation SET slice_size_percentage = 0.5 WHERE id = {:id}`, dbx.Params{"id": idBil}).
			AddCleanupQuery(`UPDATE planned_allocation SET slice_size_percentage = 0.5 WHERE id = {:id}`, dbx.Params{"id": idSpy}).
			AddCleanupQuery(`UPDATE planned_allocation SET slice_size_percentage = 0.5 WHERE id = {:id}`, dbx.Params{"id": idBondsTop}).
			AddCleanupQuery(`UPDATE planned_allocation SET slice_size_percentage = 0.5 WHERE id = {:id}`, dbx.Params{"id": idStocksTop}).
			AddCleanupQuery(
				`UPDATE allocation_plan SET name = 'Update Allocation Plan Fixture', planned_date = '2025-06-01' WHERE id = {:id}`,
				dbx.Params{"id": allocationPlanIdString},
			).
			AddCleanupQuery(`DELETE FROM asset WHERE name LIKE '%DELETE'`, nil).
			Build(t),
	)

	// 3) Update the plan: change some slices, send a different name for existing asset id=1,
	// insert a new planned allocation for NasdaqGM:TLT under BONDS, and update plan name
	var updatedPlanName = "Update Allocation Plan Fixture Updated"
	var updatePlanJSON = fmt.Sprintf(
		`{
			"id": %s,
			"name":"%s",
			"plannedDate":"2041-06-15",
			"details":[
				{ "id": %s, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.4" },
				{ "id": %s, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.6" },
				{
					"id": %s,
					"hierarchicalId":["ARCA:BIL","BONDS"],
					"cashReserve":false,
					"sliceSizePercentage":"0.9",
					"asset":{ "id":1, "name":"SHOULD NOT OVERWRITE", "ticker":"ARCA:BIL" }
				},
				{
					"id": %s,
					"hierarchicalId":["ARCA:SPY","STOCKS"],
					"cashReserve":false,
					"sliceSizePercentage":"1.0",
					"asset":{ "id":7, "name":"SPDR S&P 500 ETF Trust", "ticker":"ARCA:SPY" }
				},
				{
					"hierarchicalId":["TEST:ALTBOND","BONDS"],
					"cashReserve":true,
					"sliceSizePercentage":"0.1",
					"asset":{ "name":"Test Cash Reserve Bond DELETE", "ticker":"TEST:ALTBOND" }
				}
			]
		}`,
		allocationPlanIdString, updatedPlanName, idBondsTop, idStocksTop, idBil, idSpy,
	)

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// 4) Assert asset with id=1 kept its original name
	var assetAssertSQL = `
		SELECT a.id, a.ticker, a.name FROM asset a WHERE a.id = 1
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		assetAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":     inttestutil.ToAssertableNullString("1"),
				"ticker": inttestutil.ToAssertableNullString("ARCA:BIL"),
				"name":   inttestutil.ToAssertableNullString("SPDR Bloomberg 1-3 Month T-Bill ETF"),
			},
		},
	)

	// 5) Assert the plan name was updated
	var planAssertSQL = `
		SELECT ap.id, ap.name, ap.planned_date::text AS planned_date FROM allocation_plan ap WHERE ap.id = ` + allocationPlanIdString + `
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		planAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":           inttestutil.ToAssertableNullString(allocationPlanIdString),
				"name":         inttestutil.ToAssertableNullString(updatedPlanName),
				"planned_date": inttestutil.ToAssertableNullString("2041-06-15"),
			},
		},
	)
	var updatedPlanResponse, getPlanError = http.Get(
		inttestinfra.TestAPIURLPrefix + "/portfolio/5/allocation-plan",
	)
	require.NoError(t, getPlanError)
	defer deferCloseResponseBody(updatedPlanResponse)
	var updatedPlanBody, readPlanError = io.ReadAll(updatedPlanResponse.Body)
	require.NoError(t, readPlanError)
	assert.Contains(t, string(updatedPlanBody), `"plannedDate":"2041-06-15"`)

	// 6) Assert planned allocation rows reflect the update and insertion
	var plannedAllocationsAssertSQL = `
		SELECT 
		    pa.hierarchical_id, 
		    pa.slice_size_percentage, 
		    pa.asset_id, 
		    pa.cash_reserve, 
		    pa.total_market_value,
			ass.ticker AS asset_ticker,
		    ass.name AS asset_name
		FROM planned_allocation pa
		LEFT JOIN asset ass ON ass.id = pa.asset_id
		WHERE pa.allocation_plan_id = ` + allocationPlanIdString + `
		ORDER BY pa.hierarchical_id
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		plannedAllocationsAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{ARCA:BIL,BONDS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.90000"),
				"asset_id":              inttestutil.ToAssertableNullString("1"),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{ARCA:SPY,STOCKS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("1.00000"),
				"asset_id":              inttestutil.ToAssertableNullString("7"),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{TEST:ALTBOND,BONDS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.10000"),
				"cash_reserve":          inttestutil.ToAssertableNullString("true"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
				"asset_id":              inttestutil.NotNullAssertableNullString(),
				"asset_ticker":          inttestutil.ToAssertableNullString("TEST:ALTBOND"),
				"asset_name":            inttestutil.ToAssertableNullString("Test Cash Reserve Bond DELETE"),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{NULL,BONDS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.40000"),
				"asset_id":              inttestutil.NullAssertableNullString(),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{NULL,STOCKS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.60000"),
				"asset_id":              inttestutil.NullAssertableNullString(),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
		},
	)
}

// TestPostAllocationPlanForUpdate_DoesNotOverwriteExistingAssetName updates a pre-seeded plan
// and asserts an existing asset name is not overwritten, slice sizes are updated, and a new
// planned allocation is inserted as part of the update operation. It also updates the plan name
// and verifies the rename works, then cleans up to restore the original fixture state.
//
// Co-authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanForUpdate_DeletesPlannedAllocationAndKeepsAsset(t *testing.T) {

	// 1) Fetch the pre-seeded plan id
	var fetchPlanSQL = `
		SELECT ap.id, ap.name
		FROM allocation_plan ap
		WHERE ap.portfolio_id = 5 AND ap.name = 'Update Allocation Plan Fixture'
	`
	var allocationPlanIdString string
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		fetchPlanSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":   inttestutil.NotNullValueCapturingAssertableNullString(&allocationPlanIdString),
				"name": inttestutil.ToAssertableNullString("Update Allocation Plan Fixture"),
			},
		},
	)

	// 2) Fetch planned allocation ids for this plan so update can match existing rows
	var idBil, idSpy, idBondsTop, idStocksTop string
	var fetchPlannedAllocationsSQL = `
		SELECT pa.id, pa.hierarchical_id
		FROM planned_allocation pa
		WHERE pa.allocation_plan_id = ` + allocationPlanIdString + `
		ORDER BY pa.hierarchical_id
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		fetchPlannedAllocationsSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idBil),
				"hierarchical_id": inttestutil.ToAssertableNullString("{ARCA:BIL,BONDS}"),
			},
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idSpy),
				"hierarchical_id": inttestutil.ToAssertableNullString("{ARCA:SPY,STOCKS}"),
			},
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idBondsTop),
				"hierarchical_id": inttestutil.ToAssertableNullString("{NULL,BONDS}"),
			},
			{
				"id":              inttestutil.NotNullValueCapturingAssertableNullString(&idStocksTop),
				"hierarchical_id": inttestutil.ToAssertableNullString("{NULL,STOCKS}"),
			},
		},
	)

	// Cleanup: reinsert the deleted planned allocation for ARCA:SPY, delete new EWZ, and keep fixture stable
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM planned_allocation WHERE allocation_plan_id = 6 AND hierarchical_id = '{"ARCA:EWZ", "STOCKS"}'`,
				nil,
			).
			AddCleanupQuery(
				`INSERT INTO planned_allocation (id, allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage, total_market_value)
				 VALUES (33, 6, '{"ARCA:SPY", "STOCKS"}', 7, FALSE, 0.5, NULL)
				 ON CONFLICT (id) DO UPDATE SET
				 	allocation_plan_id = EXCLUDED.allocation_plan_id,
				 	hierarchical_id = EXCLUDED.hierarchical_id,
				 	asset_id = EXCLUDED.asset_id,
				 	cash_reserve = EXCLUDED.cash_reserve,
				 	slice_size_percentage = EXCLUDED.slice_size_percentage,
				 	total_market_value = EXCLUDED.total_market_value`,
				nil,
			).
			Build(t),
	)

	// 3) Update the plan omitting ARCA:SPY planned allocation to trigger deletion
	// Adding a different asset (EWZ) under STOCKS to satisfy the childless validation
	var updatePlanJSON = fmt.Sprintf(
		`{
			"id": %s,
			"name":"Update Allocation Plan Fixture",
			"plannedDate":"2025-06-01",
			"details":[
				{ "id": %s, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
				{ "id": %s, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" },
				{ "id": %s, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"1.0", "cashReserve":false },
				{ "hierarchicalId":["ARCA:EWZ","STOCKS"], "sliceSizePercentage":"1.0", "cashReserve":false, "asset":{"id":6} }
			]
		}`,
		allocationPlanIdString, idBondsTop, idStocksTop, idBil,
	)

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// 4) Assert planned allocations no longer contain ARCA:SPY row and
	//    that the remaining planned allocations keep the same asset_id values
	//    as in the initial fixture (BIL=1, top-level rows have NULL).
	var plannedAllocationsAssertSQL = `
		SELECT pa.hierarchical_id, pa.asset_id
		FROM planned_allocation pa
		WHERE pa.allocation_plan_id = ` + allocationPlanIdString + `
		ORDER BY pa.hierarchical_id
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		plannedAllocationsAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"hierarchical_id": inttestutil.ToAssertableNullString("{ARCA:BIL,BONDS}"),
				"asset_id":        inttestutil.ToAssertableNullString("1"),
			},
			{
				"hierarchical_id": inttestutil.ToAssertableNullString("{ARCA:EWZ,STOCKS}"),
				"asset_id":        inttestutil.ToAssertableNullString("6"),
			},
			{
				"hierarchical_id": inttestutil.ToAssertableNullString("{NULL,BONDS}"),
				"asset_id":        inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id": inttestutil.ToAssertableNullString("{NULL,STOCKS}"),
				"asset_id":        inttestutil.NullAssertableNullString(),
			},
		},
	)

	// 5) Assert SPY asset still exists unchanged
	var assetAssertSQL = `
		SELECT a.id, a.ticker, a.name FROM asset a WHERE a.id = 7
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		assetAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":     inttestutil.ToAssertableNullString("7"),
				"ticker": inttestutil.ToAssertableNullString("ARCA:SPY"),
				"name":   inttestutil.ToAssertableNullString("SPDR S&P 500 ETF Trust"),
			},
		},
	)
}

// TestPostAllocationPlanForUpdate_ChangesHierarchicalId verifies a detail update carries the existing planned date.
//
// Co-authored by: OpenCode
func TestPostAllocationPlanForUpdate_ChangesHierarchicalId(t *testing.T) {

	var allocationPlanIdString = "6"
	var idBil = "32"
	var idSpy = "33"
	var idBondsTop = "30"
	var idStocksTop = "31"

	// Register cleanup to restore fixture state after update
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
					DELETE FROM planned_allocation 
					WHERE allocation_plan_id = {:allocationPlanId} AND hierarchical_id = '{TEST:ALTBOND,BONDS2}'
				`,
				dbx.Params{"allocationPlanId": allocationPlanIdString},
			).
			AddCleanupQuery(
				`
					UPDATE planned_allocation 
					SET slice_size_percentage = 0.5, hierarchical_id = '{"ARCA:BIL", "BONDS"}' 
					WHERE id = {:id}
				`,
				dbx.Params{"id": idBil},
			).
			AddCleanupQuery(
				`
					UPDATE planned_allocation 
					SET slice_size_percentage = 0.5, hierarchical_id = '{"ARCA:SPY", "STOCKS"}'
					WHERE id = {:id}
				`,
				dbx.Params{"id": idSpy},
			).
			AddCleanupQuery(
				`
					UPDATE planned_allocation 
					SET slice_size_percentage = 0.5, hierarchical_id = '{NULL, "BONDS"}' 
					WHERE id = {:id}
				`,
				dbx.Params{"id": idBondsTop},
			).
			AddCleanupQuery(
				`
					UPDATE planned_allocation 
					SET slice_size_percentage = 0.5, hierarchical_id = '{NULL, "STOCKS"}' 
					WHERE id = {:id}
				`,
				dbx.Params{"id": idStocksTop},
			).
			AddCleanupQuery(
				`UPDATE allocation_plan SET name = 'Update Allocation Plan Fixture' WHERE id = {:id}`,
				dbx.Params{"id": allocationPlanIdString},
			).
			AddCleanupQuery(`DELETE FROM asset WHERE name LIKE '%DELETE'`, nil).
			Build(t),
	)

	// 1) Update the plan: change some slices, send a different name for existing asset id=1,
	// insert a new planned allocation for NasdaqGM:TLT under BONDS, and update plan name
	var updatedPlanName = "Update Allocation Plan Fixture Updated"
	var updatePlanJSON = fmt.Sprintf(
		`{
			"id": %s,
			"name":"%s",
			"plannedDate":"2025-06-01",
			"details":[
				{ "id": %s, "hierarchicalId":[null,"BONDS2"],  "sliceSizePercentage":"0.4" },
				{ "id": %s, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.6" },
				{
					"id": %s,
					"hierarchicalId":["ARCA:BIL","BONDS2"],
					"cashReserve":false,
					"sliceSizePercentage":"0.9",
					"asset":{ "id":1, "name":"SHOULD NOT OVERWRITE", "ticker":"ARCA:BIL" }
				},
				{
					"id": %s,
					"hierarchicalId":["ARCA:SPY","STOCKS"],
					"cashReserve":false,
					"sliceSizePercentage":"1.0",
					"asset":{ "id":7, "name":"SPDR S&P 500 ETF Trust", "ticker":"ARCA:SPY" }
				},
				{
					"hierarchicalId":["TEST:ALTBOND","BONDS2"],
					"cashReserve":true,
					"sliceSizePercentage":"0.1",
					"asset":{ "name":"Test Cash Reserve Bond DELETE", "ticker":"TEST:ALTBOND" }
				}
			]
		}`,
		allocationPlanIdString, updatedPlanName, idBondsTop, idStocksTop, idBil, idSpy,
	)

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	// 2) Assert asset with id=1 kept its original name
	var assetAssertSQL = `
		SELECT a.id, a.ticker, a.name FROM asset a WHERE a.id = 1
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		assetAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":     inttestutil.ToAssertableNullString("1"),
				"ticker": inttestutil.ToAssertableNullString("ARCA:BIL"),
				"name":   inttestutil.ToAssertableNullString("SPDR Bloomberg 1-3 Month T-Bill ETF"),
			},
		},
	)

	// 3) Assert the plan name was updated
	var planAssertSQL = `
		SELECT ap.id, ap.name FROM allocation_plan ap WHERE ap.id = ` + allocationPlanIdString + `
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		planAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":   inttestutil.ToAssertableNullString(allocationPlanIdString),
				"name": inttestutil.ToAssertableNullString(updatedPlanName),
			},
		},
	)

	// 4) Assert planned allocation rows reflect the update and insertion
	var plannedAllocationsAssertSQL = `
		SELECT 
		    pa.hierarchical_id, 
		    pa.slice_size_percentage, 
		    pa.asset_id, 
		    pa.cash_reserve, 
		    pa.total_market_value,
			ass.ticker AS asset_ticker,
		    ass.name AS asset_name
		FROM planned_allocation pa
		LEFT JOIN asset ass ON ass.id = pa.asset_id
		WHERE pa.allocation_plan_id = ` + allocationPlanIdString + `
		ORDER BY pa.hierarchical_id
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		plannedAllocationsAssertSQL,
		[]inttestutil.AssertableNullStringMap{
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{ARCA:BIL,BONDS2}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.90000"),
				"asset_id":              inttestutil.ToAssertableNullString("1"),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{ARCA:SPY,STOCKS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("1.00000"),
				"asset_id":              inttestutil.ToAssertableNullString("7"),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{TEST:ALTBOND,BONDS2}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.10000"),
				"cash_reserve":          inttestutil.ToAssertableNullString("true"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
				"asset_id":              inttestutil.NotNullAssertableNullString(),
				"asset_ticker":          inttestutil.ToAssertableNullString("TEST:ALTBOND"),
				"asset_name":            inttestutil.ToAssertableNullString("Test Cash Reserve Bond DELETE"),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{NULL,BONDS2}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.40000"),
				"asset_id":              inttestutil.NullAssertableNullString(),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
			{
				"hierarchical_id":       inttestutil.ToAssertableNullString("{NULL,STOCKS}"),
				"slice_size_percentage": inttestutil.ToAssertableNullString("0.60000"),
				"asset_id":              inttestutil.NullAssertableNullString(),
				"cash_reserve":          inttestutil.ToAssertableNullString("false"),
				"total_market_value":    inttestutil.NullAssertableNullString(),
			},
		},
	)
}

// Test validation: missing required field 'name'
// Expects 400 with message "Field 'name' failed validation: is required"
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_MissingName(t *testing.T) {

	var payload = `{
		"plannedDate": "2025-01-01",
		"details": [
			{ "hierarchicalId": [null, "BONDS"], "sliceSizePercentage": "1.0" }
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'name' failed validation: is required"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// Test validation: missing required field 'details'
// Expects 400 with message "Field 'details' failed validation: is required"
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_MissingDetails(t *testing.T) {

	var payload = `{
		"name": "Plan without details",
		"plannedDate": "2025-01-01"
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'details' failed validation: is required"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// Test validation: empty 'details' array (min length violation)
// Expects 400 with message "Field 'details' failed validation: must be at least 1"
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_EmptyDetails(t *testing.T) {

	var payload = `{
		"name": "Plan with empty details",
		"plannedDate": "2025-01-01",
		"details": []
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'details' failed validation: must be at least 1"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// Test validation: missing required field 'hierarchicalId' in a planned allocation
// Expects 400 with message "Field 'details[0].hierarchicalId' failed validation: is required"
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_MissingHierarchicalId(t *testing.T) {

	var payload = `{
		"name": "Plan with invalid detail",
		"plannedDate": "2025-01-01",
		"details": [
			{ "sliceSizePercentage": "1.0" }
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'details[0].hierarchicalId' failed validation: is required"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// Test validation: empty 'hierarchicalId' array (min length violation)
// Expects 400 with message "Field 'details[0].hierarchicalId' failed validation: must be at least 1"
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_EmptyHierarchicalId(t *testing.T) {

	var payload = `{
		"name": "Plan with empty hierarchical id",
		"plannedDate": "2025-01-01",
		"details": [
			{ "hierarchicalId": [], "sliceSizePercentage": "1.0" }
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'details[0].hierarchicalId' failed validation: must be at least 1"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// Test validation (domain): duplicate hierarchical ids within the same plan should be rejected.
// Currently NOT implemented, so this test is expected to FAIL (receives 204/500 instead of 400).
// Establishes desired error response contract for future implementation: a non-field-specific message
// listing the duplicated hierarchical IDs found in the request.
//
// Co-authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_DuplicateHierarchicalIds(t *testing.T) {

	// Top-level sums: 0.5 BONDS + 0.5 STOCKS = 1.0
	// BONDS children: BIL 0.5 + BIL (duplicate) 0.5 = 1.0
	// STOCKS children: SPY 1.0
	var updatePlanJSON = `
		{
			"id":6,
			"name":"Update Allocation Plan Fixture",
			"plannedDate":"2025-06-01",
			"details":[
				{ "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
				{ "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" },
				{ "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"0.5", "cashReserve":false },
				{ "id": 33, "hierarchicalId":["ARCA:SPY","STOCKS"], "sliceSizePercentage":"1.0", "cashReserve":false },
				{ "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"0.5", "cashReserve":false }
			]
		}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
		"errorMessage": "Allocation plan validation failed",
		"details": ["Planned allocations contain duplicated hierarchical IDs: ARCA:BIL|BONDS"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// Test validation (domain): child slice sizes within a parent must not exceed 100%.
//
// Co-authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_PercentageSumExceedsParentLimit(t *testing.T) {

	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"0.7", "cashReserve":false },
                { "id": 33, "hierarchicalId":["ARCA:SPY","STOCKS"], "sliceSizePercentage":"1.0", "cashReserve":false },
                { 
					"hierarchicalId":["NasdaqGM:TLT","BONDS"], 
					"sliceSizePercentage":"0.5", 
					"cashReserve":false, 
					"asset": {"id": 4, "ticker": "NasdaqGM:TLT", "name": "iShares 20+ Year Treasury Bond ETF"} 
				}
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
        "errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations slice sizes exceed 100% within hierarchy level(s): Classes = BONDS (120%)"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_PercentageSumBelowParentLimit keeps the plan date valid while checking hierarchy totals.
//
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_PercentageSumBelowParentLimit(t *testing.T) {

	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"0.7", "cashReserve":false },
                { "id": 33, "hierarchicalId":["ARCA:SPY","STOCKS"], "sliceSizePercentage":"1.0", "cashReserve":false },
                { 
					"hierarchicalId":["NasdaqGM:TLT","BONDS"], 
					"sliceSizePercentage":"0.1", 
					"cashReserve":false, 
					"asset": {"id": 4, "ticker": "NasdaqGM:TLT", "name": "iShares 20+ Year Treasury Bond ETF"} 
				}
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
        "errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations slice sizes sum to less than 100% within hierarchy level(s): Classes = BONDS (80%)"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// Test validation (domain): top-level slice sizes must not exceed 100%.
//
// Co-authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_TopLevelPercentageSumExceedsLimit(t *testing.T) {

	// Keep children within 1.0 to isolate the top-level violation
	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.7" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.6" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"],   "sliceSizePercentage":"1.0", "cashReserve":false },
                { "id": 33, "hierarchicalId":["ARCA:SPY","STOCKS"],  "sliceSizePercentage":"1.0", "cashReserve":false }
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
       	"errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations slice sizes exceed 100% within hierarchy level(s): Classes (TOP) (130%)"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_TopLevelPercentageSumBelowLimit isolates the allocation-total validation with a valid plan date.
//
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_TopLevelPercentageSumBelowLimit(t *testing.T) {

	// Keep children within 1.0 to isolate the top-level violation
	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.7" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.1" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"],   "sliceSizePercentage":"1.0", "cashReserve":false },
                { "id": 33, "hierarchicalId":["ARCA:SPY","STOCKS"],  "sliceSizePercentage":"1.0", "cashReserve":false }
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
       	"errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations slice sizes sum to less than 100% within hierarchy level(s): Classes (TOP) (80%)"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_NameExceedsMaxLength tests that posting an allocation plan
// with a name exceeding the max length (100 characters) returns a validation error.
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_NameExceedsMaxLength(t *testing.T) {

	var longName = strings.Repeat("a", 101) // 101 characters exceeds max=100

	var payload = `{
		"name": "` + longName + `",
		"plannedDate": "2025-01-01",
		"details": [
			{ "hierarchicalId": [null, "TEST"], "sliceSizePercentage": "1.0" }
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'name' failed validation: must not exceed 100"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_TypeExceedsMaxLength tests that posting an allocation plan
// with a type exceeding the max length (50 characters) returns a validation error.
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_TypeExceedsMaxLength(t *testing.T) {

	var longType = strings.Repeat("T", 51) // 51 characters exceeds max=50

	var payload = `{
		"name": "Test Plan",
		"type": "` + longType + `",
		"plannedDate": "2025-01-01",
		"details": [
			{ "hierarchicalId": [null, "TEST"], "sliceSizePercentage": "1.0" }
		]
	}`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
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
		"details": ["Field 'type' failed validation: must not exceed 50"]
	}`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches tests that posting an allocation plan
// with a hierarchical id that has a different number of levels than the portfolio hierarchy
// returns a validation error.
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_InvalidSizeHierarchyBranches(t *testing.T) {

	// Portfolio 5 has a 2-level hierarchy: Classes -> Assets
	// This test sends a hierarchical id with 3 levels (invalid size)
	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"1.0", "cashReserve":false },
                { "hierarchicalId":["ARCA:SPY","STOCKS"], "sliceSizePercentage":"1.0", "cashReserve":false },
                { "hierarchicalId":["ARCA:EWZ","STOCKS", "EXTRA_LEVEL"], "sliceSizePercentage":"1.0", "cashReserve":false }
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
        "errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations contain hierarchy branches with invalid size: \nEXTRA_LEVEL -> STOCKS -> ARCA:EWZ\n for portfolio hierarchy: Classes -> Assets"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_MissingParentHierarchyBranches tests that posting an allocation plan
// with a hierarchical id that has a non-null child level but missing parent level (null at a higher level)
// returns a validation error.
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_MissingParentHierarchyBranches(t *testing.T) {

	// Portfolio 5 has a 2-level hierarchy: Classes -> Assets
	// This test sends a hierarchical id with a specific asset but null class (missing parent)
	// The orphan asset ARCA:SPY has no class (null) - this is invalid because an asset must belong to a class
	// Note: slice sizes are crafted to sum to exactly 100% at the top level to isolate the validation
	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"1.0" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"1.0", "cashReserve":false },
                { "hierarchicalId":["ARCA:SPY",null], "sliceSizePercentage":"0.0", "cashReserve":false }
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
        "errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations contain hierarchy branches with missing parent levels: \n -> ARCA:SPY\n for portfolio hierarchy: Classes -> Assets"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_ChildlessHierarchyBranches tests that posting an allocation plan
// with a parent level (class) that has no corresponding child planned allocations (assets)
// returns a validation error.
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_ChildlessHierarchyBranches(t *testing.T) {

	// Portfolio 5 has a 2-level hierarchy: Classes -> Assets
	// This test sends only top-level planned allocations (class) without child asset allocations for STOCKS
	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" },
                { "id": 32, "hierarchicalId":["ARCA:BIL","BONDS"], "sliceSizePercentage":"1.0", "cashReserve":false }
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var expected = `{
        "errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations contain hierarchy branches with missing child levels: \nSTOCKS\n for portfolio hierarchy: Classes -> Assets"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanValidation_MultipleChildlessHierarchyBranches tests that posting an allocation plan
// with multiple parent levels (classes) that have no corresponding child planned allocations (assets)
// returns a validation error listing all incomplete branches.
//
// Authored by: GitHub Copilot
// Co-authored by: OpenCode
func TestPostAllocationPlanValidation_MultipleChildlessHierarchyBranches(t *testing.T) {

	// Portfolio 5 has a 2-level hierarchy: Classes -> Assets
	// This test sends only top-level planned allocations without any child asset allocations
	var updatePlanJSON = `
		{
            "id":6,
            "name":"Update Allocation Plan Fixture",
            "plannedDate":"2025-06-01",
            "details":[
                { "id": 30, "hierarchicalId":[null,"BONDS"],  "sliceSizePercentage":"0.5" },
                { "id": 31, "hierarchicalId":[null,"STOCKS"], "sliceSizePercentage":"0.5" }
            ]
        }
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/5/allocation-plan",
		"application/json",
		strings.NewReader(updatePlanJSON),
	)
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	// Both BONDS and STOCKS branches are childless (root node stripped from output)
	var expected = `{
        "errorMessage": "Allocation plan validation failed",
        "details": ["Planned allocations contain hierarchy branches with missing child levels: \nBONDS\nSTOCKS\n for portfolio hierarchy: Classes -> Assets"]
    }`
	assert.JSONEq(t, expected, string(body))
}

// TestPostAllocationPlanRequiresExplicitValidPlannedDate rejects missing and invalid dates on both creates and updates.
// It also verifies rejected requests leave existing plans unchanged and create no plan rows.
//
// Co-authored by: OpenCode
func TestPostAllocationPlanRequiresExplicitValidPlannedDate(t *testing.T) {
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(`DELETE FROM planned_allocation WHERE allocation_plan_id IN (SELECT id FROM allocation_plan WHERE name LIKE 'Invalid planned date create%')`, nil).
			AddCleanupQuery(`DELETE FROM allocation_plan WHERE name LIKE 'Invalid planned date create%'`, nil).
			AddCleanupQuery(`UPDATE planned_allocation SET slice_size_percentage = 0.5 WHERE allocation_plan_id = 6`, nil).
			AddCleanupQuery(`UPDATE allocation_plan SET name = 'Update Allocation Plan Fixture', planned_date = '2025-06-01' WHERE id = 6`, nil).
			Build(t),
	)

	var columnContract = `
		SELECT data_type, is_nullable, column_default
		FROM information_schema.columns
		WHERE table_name = 'allocation_plan' AND column_name = 'planned_date'
	`
	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		columnContract,
		[]inttestutil.AssertableNullStringMap{{
			"data_type":      inttestutil.ToAssertableNullString("date"),
			"is_nullable":    inttestutil.ToAssertableNullString("NO"),
			"column_default": inttestutil.NullAssertableNullString(),
		}},
	)

	var invalidDates = []struct {
		name     string
		property string
	}{
		{name: "missing"},
		{name: "null", property: `"plannedDate":null`},
		{name: "empty", property: `"plannedDate":""`},
		{name: "malformed", property: `"plannedDate":"17/10/2026"`},
		{name: "zero", property: `"plannedDate":"0001-01-01"`},
		{name: "year-zero", property: `"plannedDate":"0000-01-01"`},
	}

	for _, invalidDate := range invalidDates {
		for _, update := range []bool{false, true} {
			var dateProperty = ""
			if invalidDate.property != "" {
				dateProperty = invalidDate.property + ","
			}

			var requestURL = inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-plan"
			var requestBody string
			if update {
				requestURL = inttestinfra.TestAPIURLPrefix + "/portfolio/5/allocation-plan"
				requestBody = fmt.Sprintf(`{
					"id":6,
					"name":"Invalid planned date update %s",
					%s
					"details":[
						{"id":30,"hierarchicalId":[null,"BONDS"],"sliceSizePercentage":"0.4"},
						{"id":31,"hierarchicalId":[null,"STOCKS"],"sliceSizePercentage":"0.6"},
						{"id":32,"hierarchicalId":["ARCA:BIL","BONDS"],"sliceSizePercentage":"0.5"},
						{"id":33,"hierarchicalId":["ARCA:SPY","STOCKS"],"sliceSizePercentage":"0.5"}
					]
				}`, invalidDate.name, dateProperty)
			} else {
				requestBody = fmt.Sprintf(`{
					"name":"Invalid planned date create %s",
					%s
					"details":[{"hierarchicalId":[null,"BONDS"],"sliceSizePercentage":"1.0"}]
				}`, invalidDate.name, dateProperty)
			}

			response, err := http.Post(requestURL, "application/json", strings.NewReader(requestBody))
			assert.NoError(t, err)
			if response != nil {
				defer deferCloseResponseBody(response)
				assert.Equal(t, http.StatusBadRequest, response.StatusCode, "%s update=%t", invalidDate.name, update)
			}

			if !update {
				inttestutil.AssertDBWithQueryMultipleRows(
					t,
					`SELECT count(*) AS plan_count FROM allocation_plan WHERE name LIKE 'Invalid planned date create%'`,
					[]inttestutil.AssertableNullStringMap{{"plan_count": inttestutil.ToAssertableNullString("0")}},
				)
			} else {
				inttestutil.AssertDBWithQueryMultipleRows(
					t,
					`SELECT name, planned_date::text AS planned_date FROM allocation_plan WHERE id = 6`,
					[]inttestutil.AssertableNullStringMap{{
						"name":         inttestutil.ToAssertableNullString("Update Allocation Plan Fixture"),
						"planned_date": inttestutil.ToAssertableNullString("2025-06-01"),
					}},
				)
				inttestutil.AssertDBWithQueryMultipleRows(
					t,
					`SELECT count(*) AS unchanged_detail_count
					 FROM planned_allocation
					 WHERE allocation_plan_id = 6 AND slice_size_percentage = 0.5`,
					[]inttestutil.AssertableNullStringMap{{
						"unchanged_detail_count": inttestutil.ToAssertableNullString("4"),
					}},
				)
			}
		}
	}
}
