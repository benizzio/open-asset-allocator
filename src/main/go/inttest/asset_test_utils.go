package inttest

import (
	"database/sql"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/domain"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

const testExternalAssetDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU","exchangeId":"PCX"}]}`

// capturePersistedAssetExternalData retrieves the current external_data JSON payload for the
// given asset so tests can restore the original database state in cleanup.
//
// Authored by: OpenCode
func capturePersistedAssetExternalData(t *testing.T, assetId int64) sql.NullString {
	t.Helper()

	var externalData sql.NullString
	var found bool
	err := inttestinfra.FetchWithDBQuery(
		"SELECT external_data FROM asset WHERE id = {:id}",
		dbx.Params{"id": assetId},
		func(rows *dbx.Rows) error {
			found = true
			return rows.Scan(&externalData)
		},
	)
	require.NoError(t, err)
	require.True(t, found)

	return externalData
}

// addAssetExternalDataRestoreCleanup appends the cleanup query needed to restore an asset's
// original external_data payload after a test mutates it.
//
// Authored by: OpenCode
func addAssetExternalDataRestoreCleanup(
	builder *inttestutil.CleanupFunctionBuilder,
	assetId int64,
	originalExternalData sql.NullString,
) *inttestutil.CleanupFunctionBuilder {

	var params = dbx.Params{"id": assetId}
	if !originalExternalData.Valid {
		return builder.AddCleanupQuery(
			"UPDATE asset SET external_data = NULL WHERE id = {:id}",
			params,
		)
	}

	params["externalData"] = originalExternalData.String
	return builder.AddCleanupQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		params,
	)
}

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
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery("DELETE FROM asset WHERE id={:id}", dbx.Params{"id": testAsset.Id}).
			Build(t),
	)

	assert.NotZero(t, testAsset)
	assert.NotZero(t, testAsset.Id)

	return testAsset
}

// assertPersistedAsset asserts that the asset with the given ID has the expected ticker and name
// in the database.
//
// Co-authored by: OpenCode, GitHub Copilot and benizzio
func assertPersistedAsset(t *testing.T, assetId int64, expectedTicker string, expectedName string) {
	assertPersistedAssetWithExternalData(t, assetId, expectedTicker, expectedName, nil)
}

// assertPersistedAssetWithExternalData asserts that the asset with the given ID has the expected
// ticker, name, and optional external data payload in the database.
//
// Authored by: OpenCode
func assertPersistedAssetWithExternalData(
	t *testing.T,
	assetId int64,
	expectedTicker string,
	expectedName string,
	expectedExternalData *string,
) {

	var assetIdString = strconv.FormatInt(assetId, 10)
	var expectedExternalDataAssertion = inttestutil.NullAssertableNullString()
	if expectedExternalData != nil {
		expectedExternalDataJSON := *expectedExternalData
		expectedExternalDataAssertion = inttestutil.ToAssertableNullStringWithAssertion(
			func(t *testing.T, actual sql.NullString) {
				assert.True(t, actual.Valid)
				assert.JSONEq(t, expectedExternalDataJSON, actual.String)
			},
		)
	}

	inttestutil.AssertDBWithQueryMultipleRows(
		t,
		"SELECT * FROM asset WHERE id="+assetIdString,
		[]inttestutil.AssertableNullStringMap{
			{
				"id":            inttestutil.ToAssertableNullString(assetIdString),
				"ticker":        inttestutil.ToAssertableNullString(expectedTicker),
				"name":          inttestutil.ToAssertableNullString(expectedName),
				"external_data": expectedExternalDataAssertion,
			},
		},
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

// getExternalAssets sends a GET request to the /api/external-asset endpoint with the given raw
// query string and returns the response status code and body.
//
// Co-authored by: OpenCode and benizzio
func getExternalAssets(t *testing.T, rawQuery string) (int, string) {
	t.Helper()

	var requestURL = inttestinfra.TestAPIURLPrefix + "/external-asset"
	if rawQuery != "" {
		requestURL += "?" + rawQuery
	}

	var response, err = http.Get(requestURL)
	require.NoError(t, err)

	var responseBodyBytes, readErr = io.ReadAll(response.Body)
	require.NoError(t, readErr)

	err = response.Body.Close()
	require.NoError(t, err)

	return response.StatusCode, string(responseBodyBytes)
}
