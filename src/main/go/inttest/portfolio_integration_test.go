package inttest

import (
	"encoding/json"
	restmodel "github.com/benizzio/open-asset-allocator/api/rest/model"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
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
				"allocations":[
					{
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
				"allocations":[
					{
						"assetName":"SPDR Bloomberg 1-3 Month T-Bill ETF",
						"assetTicker":"ARCA:BIL",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":10000
					},
					{
						"assetName":"iShares 0-5 Year TIPS Bond ETF",
						"assetTicker":"ARCA:STIP",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":8000
					},
					{
						"assetName":"iShares 7-10 Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:IEF",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":6000
					},
					{
						"assetName":"iShares 20+ Year Treasury Bond ETF",
						"assetTicker":"NasdaqGM:TLT",
						"class":"BONDS",
						"cashReserve":false,
						"totalMarketValue":3000
					},
					{
						"assetName":"iShares Short Treasury Bond ETF",
						"assetTicker":"NasdaqGM:SHV",
						"class":"STOCKS",
						"cashReserve":true,
						"totalMarketValue":9000
					},
					{
						"assetName":"SPDR S\u0026P 500 ETF Trust",
						"assetTicker":"ARCA:SPY",
						"class":"STOCKS",
						"cashReserve":false,
						"totalMarketValue":8000
					},
					{
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

func TestPostPortfolio(t *testing.T) {

	var postPortfolioJSON = `
		{
			"name":"Test Portfolio creation"
		}
	`

	response, err := http.Post(
		inttestinfra.TestAPIURLprefix+"/portfolio",
		"application/json",
		strings.NewReader(postPortfolioJSON),
	)
	assert.NoError(t, err)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"name":"Test Portfolio creation",
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

	var portfolioDTS restmodel.PortfolioDTS
	err = json.Unmarshal(body, &portfolioDTS)
	assert.NoError(t, err)

	var portfolioIdString = strconv.Itoa(*portfolioDTS.Id)
	var portfolioNullStringMap = dbx.NullStringMap{
		"id":                   inttestutil.ToNullString(portfolioIdString),
		"name":                 inttestutil.ToNullString("Test Portfolio creation"),
		"allocation_structure": inttestutil.ToNullString(`{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}`),
	}

	inttestutil.AssertQuery(
		t,
		"SELECT * FROM portfolio WHERE id="+portfolioIdString,
		portfolioNullStringMap,
	)

	// TODO create solution to clean test data
}
