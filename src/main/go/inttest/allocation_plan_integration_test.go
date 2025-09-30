package inttest

import (
	"io"
	"net/http"
	"strings"
	"testing"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	"github.com/stretchr/testify/assert"
)

func TestGetAllocationPlans(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-plan")
	assert.NoError(t, err)

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

func TestPostAllocationPlanForInsertion(t *testing.T) {

	var newAllocationPlanJSON = `
		{
			"name":"New Allocation Plan Test DELETE",
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

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio/1/allocation-plan",
		"application/json",
		strings.NewReader(newAllocationPlanJSON),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNoContent, response.StatusCode)

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`
				DELETE FROM planned_allocation 
				WHERE allocation_plan_id IN (SELECT id FROM allocation_plan WHERE name LIKE '%%DELETE')
				`,
			).
			AddCleanupQuery(`DELETE FROM allocation_plan WHERE name LIKE '%%DELETE'`).
			AddCleanupQuery(`DELETE FROM asset WHERE name LIKE '%%DELETE'`).
			Build(),
	)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.Empty(t, string(body))

	var actualAllocationPlanQuery = `
		SELECT
		    ap.id,
		    ap.name,
		    ap.type,
		    ap.planned_execution_date,
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

// TODO test updating an allocation plan and some planned allocations - test if asset name of already existing asset is not overwritten

// TODO test updating an allocation plan with deletion of planned allocations

// TODO test validation: simple controller validations (required fields, string lengths, etc.)

// TODO test validation: domain unique hierarchical ids

// TODO test validation: domain sum of slice size percentages inside parent (or top level) in hierarchy <= 100%
