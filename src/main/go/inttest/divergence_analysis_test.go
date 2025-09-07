package inttest

import (
	"io"
	"net/http"
	"testing"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	"github.com/stretchr/testify/assert"
)

func TestGetDivergenceAnalysisOptions(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/divergence/options")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"availableObservedHistory":[
				{ "id": 2, "timeTag": "202503", "timestamp": "2025-03-01T00:00:00Z" },
				{ "id": 1, "timeTag": "202501", "timestamp": "2025-01-01T00:00:00Z" }
			],	
			"availablePlans":[
				{
					"id":1,
					"name":"60/40 Portfolio Classic - Example"
				}
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetDivergenceAnalysisV2(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/v2/portfolio/1/divergence/1/allocation-plan/1")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"portfolioId":1,
			"observationTimestamp":  {
				"id": 1,
				"timeTag": "202501",
				"timestamp": "2025-01-01T00:00:00Z"
			},
			"allocationPlanId":1,
			"portfolioTotalMarketValue":45000,
			"root":[
				{
					"hierarchyLevelKey":"BONDS",
					"hierarchicalId":"BONDS",
					"totalMarketValue":27000,
					"totalMarketValueDivergence":0,
					"depth":0,
					"internalDivergences":[
						{
							"hierarchyLevelKey":"ARCA:BIL",
							"hierarchicalId":"ARCA:BIL|BONDS",
							"totalMarketValue":10000,
							"totalMarketValueDivergence":-800,
							"depth":1
						},
						{
							"hierarchyLevelKey":"ARCA:STIP",
							"hierarchicalId":"ARCA:STIP|BONDS",
							"totalMarketValue":8000,
							"totalMarketValueDivergence":5300,
							"depth":1
						},
						{
							"hierarchyLevelKey":"NasdaqGM:IEF",
							"hierarchicalId":"NasdaqGM:IEF|BONDS",
							"totalMarketValue":6000,
							"totalMarketValueDivergence":-2100,
							"depth":1
						},
						{
							"hierarchyLevelKey":"NasdaqGM:TLT",
							"hierarchicalId":"NasdaqGM:TLT|BONDS",
							"totalMarketValue":3000,
							"totalMarketValueDivergence":-2400,
							"depth":1
						}
					]
				},
				{
					"hierarchyLevelKey":"STOCKS",
					"hierarchicalId":"STOCKS",
					"totalMarketValue":18000,
					"totalMarketValueDivergence":0,
					"depth":0,
					"internalDivergences":[
						{
							"hierarchyLevelKey":"NasdaqGM:SHV",
							"hierarchicalId":"NasdaqGM:SHV|STOCKS",
							"totalMarketValue":9000,
							"totalMarketValueDivergence":0,
							"depth":1
						},
						{
							"hierarchyLevelKey":"ARCA:SPY",
							"hierarchicalId":"ARCA:SPY|STOCKS",
							"totalMarketValue":8000,
							"totalMarketValueDivergence":-100,
							"depth":1
						},
						{
							"hierarchyLevelKey":"ARCA:EWZ",
							"hierarchicalId":"ARCA:EWZ|STOCKS",
							"totalMarketValue":1000,
							"totalMarketValueDivergence":100,
							"depth":1
						}
					]
				}
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetDivergenceAnalysisV2WhenPlanHasBiggerRange tests divergence analysis when the allocation plan
// has a bigger range of assets than what's observed in the current portfolio state.
//
// Authored by: GitHub Copilot
// TODO test expectations are wrong, fix
func TestGetDivergenceAnalysisV2WhenPlanHasBiggerRange(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/v2/portfolio/1/divergence/2/allocation-plan/1")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"portfolioId":1,
			"observationTimestamp":  {
				"id": 2,
				"timeTag": "202503",
				"timestamp": "2025-03-01T00:00:00Z"
			},
			"allocationPlanId":1,
			"portfolioTotalMarketValue":10000,
			"root":[
				{
					"hierarchyLevelKey":"BONDS",
					"hierarchicalId":"BONDS",
					"totalMarketValue":10000,
					"totalMarketValueDivergence":4000,
					"depth":0,
					"internalDivergences":[
						{
							"hierarchyLevelKey":"ARCA:BIL",
							"hierarchicalId":"ARCA:BIL|BONDS",
							"totalMarketValue":10000,
							"totalMarketValueDivergence":6000,
							"depth":1
						},
						{
							"hierarchyLevelKey":"ARCA:STIP",
							"hierarchicalId":"ARCA:STIP|BONDS",
							"totalMarketValue":0,
							"totalMarketValueDivergence":-1000,
							"depth":1
						},
						{
							"hierarchyLevelKey":"NasdaqGM:IEF",
							"hierarchicalId":"NasdaqGM:IEF|BONDS",
							"totalMarketValue":0,
							"totalMarketValueDivergence":-3000,
							"depth":1
						},
						{
							"hierarchyLevelKey":"NasdaqGM:TLT",
							"hierarchicalId":"NasdaqGM:TLT|BONDS",
							"totalMarketValue":0,
							"totalMarketValueDivergence":-2000,
							"depth":1
						}
					]
				},
				{
					"hierarchyLevelKey":"STOCKS",
					"hierarchicalId":"STOCKS",
					"totalMarketValue":0,
					"totalMarketValueDivergence":-500,
					"depth":0,
					"internalDivergences":[
						{
							"hierarchyLevelKey":"ARCA:EWZ",
							"hierarchicalId":"ARCA:EWZ|STOCKS",
							"totalMarketValue":0,
							"totalMarketValueDivergence":0,
							"depth":1
						},
						{
							"hierarchyLevelKey":"ARCA:SPY",
							"hierarchicalId":"ARCA:SPY|STOCKS",
							"totalMarketValue":0,
							"totalMarketValueDivergence":0,
							"depth":1
						},
						{
							"hierarchyLevelKey":"NasdaqGM:SHV",
							"hierarchicalId":"NasdaqGM:SHV|STOCKS",
							"totalMarketValue":0,
							"totalMarketValueDivergence":0,
							"depth":1
						}
					]
				}
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}
