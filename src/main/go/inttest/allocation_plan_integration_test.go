package inttest

import (
	"io"
	"net/http"
	"testing"

	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
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
