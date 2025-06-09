package inttest

import (
	"encoding/json"
	"fmt"
	restmodel "github.com/benizzio/open-asset-allocator/api/rest/model"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"
)

func TestGetPortfolio(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"id":1,
			"name":"My Portfolio Example",
			"allocationStructure": {
				"hierarchy": [
					{
						"name":"Assets",
						"field":"assetTicker"
					},
					{
						"name":"Classes",
						"field":"class"
					}
				]
			}
		}
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetPortfolios(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"id":1,
				"name":"My Portfolio Example",
				"allocationStructure": {
					"hierarchy": [
						{
							"name":"Assets",
							"field":"assetTicker"
						},
						{
							"name":"Classes",
							"field":"class"
						}
					]
				}
			},
			{
				"id":2,
				"name":"Test Portfolio 2",
				"allocationStructure": {
					"hierarchy": [
						{
							"name":"Assets",
							"field":"assetTicker"
						}
					]
				}
			}
		]	
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetPortfolioAllocationHistory(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/history")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"timeFrameTag":"202503",
				"observationTimestamp" : {
					"id": 2,
					"observationTimeTag": "202503",
					"observationTimestamp": "2025-03-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					}
				],
				"totalMarketValue":10000
			},
			{
				"timeFrameTag":"202501",
				"observationTimestamp" : {
					"id": 1,
					"observationTimeTag": "202501",
					"observationTimestamp": "2025-01-01T00:00:00Z"
				},
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					},
					{
						"assetId": 2,
						"assetName":"iShares 0-5 Year TIPS Bond ETF",
						"assetTicker":"ARCA:STIP",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":8000
					},
					{
						"assetId": 3,
						"assetName":"iShares 7-10 Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:IEF",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":6000
					},
					{
						"assetId": 4,
						"assetName":"iShares 20+ Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:TLT",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":3000
					},
					{
						"assetId": 5,	
						"assetName":"iShares Short Treasury Bond ETF",
						"assetTicker":"NasdaqGM:SHV",
						"class":"STOCKS",
						"cashReserve":true,
						"totalMarketValue":9000
					},
					{
 						"assetId": 7,
						"assetName":"SPDR S\u0026P 500 ETF Trust",
						"assetTicker":"ARCA:SPY",
						"class":"STOCKS",
						"cashReserve":false,
						"totalMarketValue":8000
					},
					{
						"assetId": 6,
						"assetName":"iShares Msci Brazil ETF",
						"assetTicker":"ARCA:EWZ",
						"class":"STOCKS",
						"cashReserve":false,
						"totalMarketValue":1000
					}
				],
				"totalMarketValue":45000
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestGetPortfolioAllocationHistoryForTimeFrame(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLprefix + "/portfolio/1/history?timeFrameTag=202503")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"timeFrameTag":"202503",
				"allocations":[
					{
						"assetId": 1,
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					}
				],
				"totalMarketValue":10000
			}
		]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

func TestPostPortfolio(t *testing.T) {

	var testPortfolioName = "Test Portfolio creation"

	var postPortfolioJSON = `
		{
			"name":"` + testPortfolioName + `"
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLprefix+"/portfolio",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)

	t.Cleanup(
		inttestutil.CreateDBCleanupDeferable(
			"DELETE FROM portfolio WHERE name='%s'",
			testPortfolioName,
		),
	)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"name":"` + testPortfolioName + `",
			"allocationStructure": {
				"hierarchy": [
					{
						"name":"Assets",
						"field":"assetTicker"
					},
					{
						"name":"Classes",
						"field":"class"
					}
				]
			}
		}
	`

	inttestutil.AssertJSONEqualIgnoringFields(t, expectedResponseJSON, string(body), "id")

	var actualPortfolioDTS restmodel.PortfolioDTS
	err = json.Unmarshal(body, &actualPortfolioDTS)
	assert.NoError(t, err)
	assert.NotNil(t, actualPortfolioDTS.Id)
	assert.NotZero(t, *actualPortfolioDTS.Id)

	assertPersistedPortfolioFromDTS(
		t,
		actualPortfolioDTS,
		`{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}`,
	)
}

func TestPostPortfolioWithAllocationStructure(t *testing.T) {

	var testPortfolioName = "Test Portfolio creation with allocation structure"

	var allocationStructureJSONFragment = `
		"allocationStructure": {
			"hierarchy": [
				{
					"name":"Classes",
					"field":"class"
				}
			]
		}
	`

	var postPortfolioJSON = `
		{
			"name":"` + testPortfolioName + `",
			` + allocationStructureJSONFragment + `
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLprefix+"/portfolio",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)

	t.Cleanup(
		inttestutil.CreateDBCleanupDeferable(
			"DELETE FROM portfolio WHERE name='%s'",
			testPortfolioName,
		),
	)

	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"name":"` + testPortfolioName + `",
			` + allocationStructureJSONFragment + `
		}
	`

	inttestutil.AssertJSONEqualIgnoringFields(t, expectedResponseJSON, string(body), "id")

	var actualPortfolioDTS restmodel.PortfolioDTS
	err = json.Unmarshal(body, &actualPortfolioDTS)
	assert.NoError(t, err)
	assert.NotNil(t, actualPortfolioDTS.Id)
	assert.NotZero(t, *actualPortfolioDTS.Id)

	assertPersistedPortfolioFromDTS(
		t,
		actualPortfolioDTS,
		`{"hierarchy": [{"name": "Classes", "field": "class"}]}`,
	)
}

func TestPostPortfolioFailureWithoutMandatoryFields(t *testing.T) {

	var postPortfolioJSONNullFields = `
		{
			"name": null
		}
	`

	var postPortfolioJSONEmptyFields = `
		{
			"name": ""
		}
	`
	actualResponseJSONNullFields := string(postForValidationFailure(t, postPortfolioJSONNullFields))
	actualResponseJSONEmptyFields := string(postForValidationFailure(t, postPortfolioJSONEmptyFields))

	var expectedResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'name' failed validation: is required"
			]
		}
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSONNullFields)
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSONEmptyFields)
}

func postForValidationFailure(t *testing.T, postPortfolioJSON string) []byte {

	response, err := http.Post(
		inttestinfra.TestAPIURLprefix+"/portfolio",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
	return body
}

func TestPutPortfolio(t *testing.T) {

	var testPortfolioNameBefore = "This Test Portfolio will be updated"
	var testPortfolioNameAfter = "Test Portfolio update"

	testPortFolio := insertTestPortfolio(t, testPortfolioNameBefore)
	var testPortfolioIdString = strconv.Itoa(testPortFolio.Id)

	var putPortfolioJSON = `
		{
			"id":` + testPortfolioIdString + `,
			"name":"` + testPortfolioNameAfter + `"
		}
	`

	response := putPortfolio(t, putPortfolioJSON)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"id":` + testPortfolioIdString + `,
			"name":"` + testPortfolioNameAfter + `",
			"allocationStructure": {
				"hierarchy": [
					{
						"name":"Assets",
						"field":"assetTicker"
					},
					{
						"name":"Classes",
						"field":"class"
					}
				]
			}
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))

	var portfolioDTS restmodel.PortfolioDTS
	err = json.Unmarshal(body, &portfolioDTS)
	assert.NoError(t, err)

	assertPersistedPortfolioFromDTS(
		t,
		portfolioDTS,
		`{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}`,
	)
}

func TestPutPortfolioFailureWithoutMandatoryFields(t *testing.T) {

	var testPortfolioName = "This Test Portfolio will be updated"
	testPortFolio := insertTestPortfolio(t, testPortfolioName)

	var putPortfolioJSONNullFields = `
		{
			"id": %d,
			"name": null
		}
	`

	var putPortfolioJSONEmptyFields = `
		{
			"id": %d,
			"name": ""
		}
	`

	var putPortfolioJSONNoFields = `
		{
			"id": %d
		}
	`

	var putPortfolioJSONNoId = `
		{
			"name": "Portfolio without ID"
		}
	`

	var actualResponseJSONNullFields = string(
		putForValidationFailure(
			t,
			fmt.Sprintf(putPortfolioJSONNullFields, testPortFolio.Id),
		),
	)

	var actualResponseJSONEmptyFields = string(
		putForValidationFailure(
			t,
			fmt.Sprintf(putPortfolioJSONEmptyFields, testPortFolio.Id),
		),
	)

	var actualResponseJSONNoFields = string(
		putForValidationFailure(
			t,
			fmt.Sprintf(putPortfolioJSONNoFields, testPortFolio.Id),
		),
	)

	var actualResponseJSONNoId = string(
		putForValidationFailure(
			t,
			putPortfolioJSONNoId,
		),
	)

	var expectedNoNameResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'name' failed validation: is required"
			]
		}
	`

	var expectedNoIdResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'id' failed validation: is required"
			]
		}
	`

	assert.JSONEq(t, expectedNoNameResponseJSON, actualResponseJSONNullFields)
	assert.JSONEq(t, expectedNoNameResponseJSON, actualResponseJSONEmptyFields)
	assert.JSONEq(t, expectedNoNameResponseJSON, actualResponseJSONNoFields)
	assert.JSONEq(t, expectedNoIdResponseJSON, actualResponseJSONNoId)

	assertPersistedPortfolioFromAttributes(
		t,
		testPortFolio.Id,
		testPortfolioName,
		`{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}`,
	)
}

func TestPutPortfolioWithAllocationStructure(t *testing.T) {

	var testPortfolioNameBefore = "This Test Portfolio will be updated"
	var testPortfolioNameAfter = "Test Portfolio update"

	testPortfolio := insertTestPortfolio(t, testPortfolioNameBefore)
	var testPortfolioIdString = strconv.Itoa(testPortfolio.Id)

	var allocationStructureJSONFragment = `
		"allocationStructure": {
			"hierarchy": [
				{
					"name":"Classes",
					"field":"class"
				}
			]
		}
	`

	var putPortfolioJSON = `
		{
			"id":"` + testPortfolioIdString + `",
			"name":"` + testPortfolioNameAfter + `",
			` + allocationStructureJSONFragment + `
		}
	`

	response := putPortfolio(t, putPortfolioJSON)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"id":` + testPortfolioIdString + `,
			"name":"` + testPortfolioNameAfter + `",
			"allocationStructure": {
				"hierarchy": [
					{
						"name":"Assets",
						"field":"assetTicker"
					},
					{
						"name":"Classes",
						"field":"class"
					}
				]
			}
		}
	`
	t.Log(string(body))
	assert.JSONEq(t, expectedResponseJSON, string(body))

	var portfolioDTS restmodel.PortfolioDTS
	err = json.Unmarshal(body, &portfolioDTS)
	assert.NoError(t, err)

	assertPersistedPortfolioFromDTS(
		t,
		portfolioDTS,
		`{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}`,
	)
}

func putForValidationFailure(t *testing.T, putPortfolioJSON string) []byte {

	response := putPortfolio(t, putPortfolioJSON)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
	return body
}
