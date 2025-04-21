package inttest

import (
	"github.com/benizzio/open-asset-allocator/inttest/util"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

func TestGetPortfolio(t *testing.T) {

	response, err := http.Get(util.TestAPIURLprefix + "/portfolio/1")
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

	response, err := http.Get(util.TestAPIURLprefix + "/portfolio")
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
