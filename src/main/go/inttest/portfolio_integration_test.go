package inttest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	restmodel "github.com/benizzio/open-asset-allocator/api/rest/model"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	"github.com/stretchr/testify/assert"
)

func TestGetPortfolio(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1")
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

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio")
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
			},
			{
				"id":3,
				"name":"Set difference test portfolio",
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
		inttestinfra.TestAPIURLPrefix+"/portfolio",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)

	t.Cleanup(
		inttestutil.CreateDBCleanupFunction(
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
		inttestinfra.TestAPIURLPrefix+"/portfolio",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)

	t.Cleanup(
		inttestutil.CreateDBCleanupFunction(
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
	actualResponseJSONNullFields := string(postPortfolioForValidationFailure(t, postPortfolioJSONNullFields))
	actualResponseJSONEmptyFields := string(postPortfolioForValidationFailure(t, postPortfolioJSONEmptyFields))

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

func postPortfolioForValidationFailure(t *testing.T, postPortfolioJSON string) []byte {

	response, err := http.Post(
		inttestinfra.TestAPIURLPrefix+"/portfolio",
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
	var testPortfolioIdString = strconv.FormatInt(testPortFolio.Id, 10)

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
	var testPortfolioIdString = strconv.FormatInt(testPortfolio.Id, 10)

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

// TestGetAvailablePortfolioAllocationClasses tests the unified endpoint that retrieves
// allocation classes from both portfolio_allocation_fact and planned_allocation tables.
// Portfolio 1 has "BONDS" and "STOCKS" classes in portfolio_allocation_fact table,
// and "BONDS", "STOCKS", and "COMMODITIES" in planned_allocation table.
// The endpoint should return all three unique classes.
//
// Authored by: GitHub Copilot
func TestGetAvailablePortfolioAllocationClasses(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/1/allocation-classes")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		["BONDS", "COMMODITIES", "STOCKS"]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetAvailablePortfolioAllocationClassesNoneFound tests when no allocation classes are found
// for a portfolio that has neither portfolio_allocation_fact nor planned_allocation data.
//
// Authored by: GitHub Copilot
func TestGetAvailablePortfolioAllocationClassesNoneFound(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/2/allocation-classes")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetAvailablePortfolioAllocationClassesFromPlansOnly tests the unified endpoint
// when portfolio has classes from both allocation history and allocation plans.
// Portfolio 3 has "A_TEST_CLASS", "BONDS", and "STOCKS" classes.
//
// Authored by: GitHub Copilot
func TestGetAvailablePortfolioAllocationClassesFromPlansOnly(t *testing.T) {

	response, err := http.Get(inttestinfra.TestAPIURLPrefix + "/portfolio/3/allocation-classes")
	assert.NoError(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	// Portfolio 3 has allocation history with "A_TEST_CLASS", "BONDS", and "STOCKS"
	// And allocation plans with "BONDS" and "STOCKS"
	var expectedResponseJSON = `
		["A_TEST_CLASS", "BONDS", "STOCKS"]
	`

	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}
