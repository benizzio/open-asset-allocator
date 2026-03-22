package inttest

import (
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/benizzio/open-asset-allocator/domain"
	"github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
)

// insertTestAsset inserts a test asset into the database and registers a cleanup function.
// Returns the persisted asset with its generated ID.
//
// Authored by: GitHub Copilot
func insertTestAsset(t *testing.T, ticker string, name string) domain.Asset {

	var insertAssetSQL = `
		INSERT INTO asset (ticker, name)
		VALUES ({:ticker}, {:name})
	`

	err := inttestinfra.ExecuteDBQuery(insertAssetSQL, dbx.Params{"ticker": ticker, "name": name})
	assert.NoError(t, err)

	var testAsset domain.Asset
	err = inttestinfra.FetchWithDBQuery(
		"SELECT * FROM asset WHERE ticker = {:ticker}",
		dbx.Params{"ticker": ticker},
		func(rows *dbx.Rows) error {
			return rows.ScanStruct(&testAsset)
		},
	)
	assert.NoError(t, err)

	t.Cleanup(
		inttestutil.CreateDBCleanupFunction(
			t,
			"DELETE FROM asset WHERE id={:id}",
			dbx.Params{"id": testAsset.Id},
		),
	)

	assert.NotZero(t, testAsset)
	assert.NotZero(t, testAsset.Id)

	return testAsset
}

// assertPersistedAsset asserts that the asset with the given ID has the expected ticker and name
// in the database.
//
// Authored by: GitHub Copilot
func assertPersistedAsset(t *testing.T, assetId int64, expectedTicker string, expectedName string) {

	var assetIdString = strconv.FormatInt(assetId, 10)
	var assetNullStringMap = dbx.NullStringMap{
		"id":     util.StringToNullString(assetIdString),
		"ticker": util.StringToNullString(expectedTicker),
		"name":   util.StringToNullString(expectedName),
	}

	inttestutil.AssertDBWithQuery(
		t,
		"SELECT * FROM asset WHERE id="+assetIdString,
		assetNullStringMap,
	)
}

// putAsset sends a PUT request to the /api/asset endpoint with the given JSON body.
//
// Authored by: GitHub Copilot
func putAsset(t *testing.T, putAssetJSON string) *http.Response {

	request, err := http.NewRequest(
		http.MethodPut,
		inttestinfra.TestAPIURLPrefix+"/asset",
		strings.NewReader(putAssetJSON),
	)
	assert.NoError(t, err)

	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	assert.NoError(t, err)

	return response
}
