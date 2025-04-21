package inttest

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetAllocationPlans(t *testing.T) {

	response, err := http.Get(testAPIURLprefix + "/portfolio/1/allocation-plan")
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
						"structuralId":["NasdaqGM:SHV","STOCKS"],
						"cashReserve":true,
						"sliceSizePercentage":"0.5"
					},
					{
						"structuralId":[null,"BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.6"
					},
					{
						"structuralId":["ARCA:SPY","STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.45"
					},
					{
						"structuralId":[null,"STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.4"
					},
					{
						"structuralId":["ARCA:BIL","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.4"
					},
					{
						"structuralId":["NasdaqGM:IEF","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.3"
					},
					{
						"structuralId":["NasdaqGM:TLT","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.2"
					},
					{
						"structuralId":["ARCA:STIP","BONDS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.1"
					},
					{
						"structuralId":["ARCA:EWZ","STOCKS"],
						"cashReserve":false,
						"sliceSizePercentage":"0.05"
					}
				]
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}
