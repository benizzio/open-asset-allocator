package inttest

import (
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetDivergenceAnalysisOptions(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/divergence/options")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	t.Log(actualResponseJSON)
	var expectedResponseJSON = `
		{
			"availableHistory":["202503", "202501"],
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

func TestGetDivergenceAnalysis(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/divergence/202501/allocation-plan/1")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"portfolioId":1,
			"timeFrameTag":"202501",
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
