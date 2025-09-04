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
						"sliceSizePercentage":"0.5"
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
						"sliceSizePercentage":"0.45"
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
						"sliceSizePercentage":"0.4"
					},
					{
						"id":4,
						"hierarchicalId":["NasdaqGM:IEF","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.3"
					},
					{
						"id":5,
						"hierarchicalId":["NasdaqGM:TLT","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.2"
					},
					{
						"id":6,
						"hierarchicalId":["ARCA:STIP","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.1"
					},
					{
						"id":8,
						"hierarchicalId":["ARCA:EWZ","STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.05"
					}
				]
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}
