package inttest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"testing"

	restmodel "github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
)

func insertTestPortfolio(t *testing.T, testPortfolioNameBefore string) domain.Portfolio {

	var insertPortfolioSQL = `
		INSERT INTO portfolio (name, allocation_structure)
		VALUES (
			'%s',
		    '{"hierarchy": [{"name": "Assets", "field": "assetTicker"}, {"name": "Classes", "field": "class"}]}'
		)
	`
	var formattedInsertPortfolioSQL = fmt.Sprintf(insertPortfolioSQL, testPortfolioNameBefore)

	err := inttestinfra.ExecuteDBQuery(formattedInsertPortfolioSQL)
	assert.NoError(t, err)

	var testPortFolio domain.Portfolio
	err = inttestinfra.FetchWithDBQuery(
		fmt.Sprintf("SELECT * FROM portfolio WHERE name = '%s'", testPortfolioNameBefore),
		func(rows *dbx.Rows) error {
			return rows.ScanStruct(&testPortFolio)
		},
	)
	assert.NoError(t, err)

	t.Cleanup(
		inttestutil.CreateDBCleanupFunction(
			"DELETE FROM portfolio WHERE id='%d'",
			testPortFolio.Id,
		),
	)

	assert.NotZero(t, testPortFolio)
	assert.NotZero(t, testPortFolio.Id)

	return testPortFolio
}

func assertPersistedPortfolioFromDTS(
	t *testing.T,
	actualPortfolioDTS restmodel.PortfolioDTS,
	actualPortFolioAllocationStructure string,
) {
	assertPersistedPortfolioFromAttributes(
		t,
		int64(*actualPortfolioDTS.Id),
		actualPortfolioDTS.Name,
		actualPortFolioAllocationStructure,
	)
}

func assertPersistedPortfolioFromAttributes(
	t *testing.T,
	actualPortfolioID int64,
	actualPortfolioName string,
	actualPortFolioAllocationStructure string,
) {

	var portfolioIdString = strconv.FormatInt(actualPortfolioID, 10)
	var portfolioNullStringMap = dbx.NullStringMap{
		"id":                   util.ToNullString(portfolioIdString),
		"name":                 util.ToNullString(actualPortfolioName),
		"allocation_structure": util.ToNullString(actualPortFolioAllocationStructure),
	}

	inttestutil.AssertDBWithQuery(
		t,
		"SELECT * FROM portfolio WHERE id="+portfolioIdString,
		portfolioNullStringMap,
	)
}

func putPortfolio(t *testing.T, putPortfolioJSON string) *http.Response {

	request, err := http.NewRequest(
		http.MethodPut,
		inttestinfra.TestAPIURLPrefix+"/portfolio",
		strings.NewReader(putPortfolioJSON),
	)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	return response
}
