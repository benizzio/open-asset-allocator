package inttest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"testing"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"

	"github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// TestGetKnownAssets tests the GET /api/asset endpoint to retrieve all known assets.
//
// Authored by: GitHub Copilot
func TestGetKnownAssets(t *testing.T) {

	response, err := http.Get(infra.TestAPIURLPrefix + "/asset")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		[
			{
				"id": 1,
				"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
				"ticker": "ARCA:BIL"
			},
			{
				"id": 6,
				"name": "iShares Msci Brazil ETF",
				"ticker": "ARCA:EWZ"
			},
			{
				"id": 7,
				"name": "SPDR S&P 500 ETF Trust",
				"ticker": "ARCA:SPY"
			},
			{
				"id": 2,
				"name": "iShares 0-5 Year TIPS Bond ETF",
				"ticker": "ARCA:STIP"
			},
			{
				"id": 3,
				"name": "iShares 7-10 Year Treasury Bond ETF",
				"ticker": "NasdaqGM:IEF"
			},
			{
				"id": 5,
				"name": "iShares Short Treasury Bond ETF",
				"ticker": "NasdaqGM:SHV"
			},
			{
				"id": 4,
				"name": "iShares 20+ Year Treasury Bond ETF",
				"ticker": "NasdaqGM:TLT"
			}
		]
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetKnownAssetsIncludesPersistedExternalData verifies that GET /api/asset returns the
// complete external data payload persisted for an asset.
//
// Authored by: OpenCode
func TestGetKnownAssetsIncludesPersistedExternalData(t *testing.T) {
	const persistedExternalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"BIL-PERSISTED","exchangeId":"PCX-PERSISTED"},{"source":"YAHOO_FINANCE","ticker":"BIL-SECOND","exchangeId":"SECOND"}]}`

	var assetId int64 = 1
	var originalExternalData = capturePersistedAssetExternalData(t, assetId)
	t.Cleanup(
		addAssetExternalDataRestoreCleanup(
			inttestutil.BuildCleanupFunctionBuilder(),
			assetId,
			originalExternalData,
		).Build(t),
	)

	var err = infra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{
			"externalData": persistedExternalDataJSON,
			"id":           assetId,
		},
	)
	assert.NoError(t, err)

	response, err := http.Get(infra.TestAPIURLPrefix + "/asset")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualAssets []json.RawMessage
	err = json.Unmarshal(body, &actualAssets)
	assert.NoError(t, err)

	var persistedAssetJSON json.RawMessage
	for _, assetJSON := range actualAssets {
		var assetIdentifier struct {
			Id int64 `json:"id"`
		}
		err = json.Unmarshal(assetJSON, &assetIdentifier)
		assert.NoError(t, err)
		if assetIdentifier.Id == assetId {
			persistedAssetJSON = assetJSON
			break
		}
	}

	assert.NotEmpty(t, persistedAssetJSON)
	var expectedAssetJSON = `
		{
			"id": 1,
			"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
			"ticker": "ARCA:BIL",
			"externalData": {
				"data": [
					{
						"source": "YAHOO_FINANCE",
						"ticker": "BIL-PERSISTED",
						"exchangeId": "PCX-PERSISTED"
					},
					{
						"source": "YAHOO_FINANCE",
						"ticker": "BIL-SECOND",
						"exchangeId": "SECOND"
					}
				]
			}
		}
	`
	assert.JSONEq(t, expectedAssetJSON, string(persistedAssetJSON))
}

// TestGetKnownAssetsWithTextSearch verifies case-insensitive terms, cross-field matching, exact
// quoted phrases, and unmatched quote handling for GET /api/asset.
//
// Authored by: OpenCode
func TestGetKnownAssetsWithTextSearch(t *testing.T) {
	var crossFieldAsset = insertTestAsset(t, "TEST:CROSS-FIELD", "Alpha Search Asset")

	var testCases = []struct {
		name             string
		textSearch       string
		expectedResponse string
	}{
		{
			name:       "case insensitive terms in any order",
			textSearch: "BLOOMBERG spdr",
			expectedResponse: `
				[
					{
						"id": 1,
						"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
						"ticker": "ARCA:BIL"
					}
				]
			`,
		},
		{
			name:       "terms can match ticker and name independently",
			textSearch: "CROSS alpha",
			expectedResponse: fmt.Sprintf(`
				[
					{
						"id": %d,
						"name": "Alpha Search Asset",
						"ticker": "TEST:CROSS-FIELD"
					}
				]
			`, crossFieldAsset.Id),
		},
		{
			name:       "quoted phrase preserves order",
			textSearch: `"SPDR Bloomberg"`,
			expectedResponse: `
				[
					{
						"id": 1,
						"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
						"ticker": "ARCA:BIL"
					}
				]
			`,
		},
		{
			name:             "quoted phrase with reversed order does not match",
			textSearch:       `"Bloomberg SPDR"`,
			expectedResponse: `[]`,
		},
		{
			name:       "unmatched quote is treated as ordinary terms",
			textSearch: `"spdr bloomberg`,
			expectedResponse: `
				[
					{
						"id": 1,
						"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
						"ticker": "ARCA:BIL"
					}
				]
			`,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			response, err := getAssetsByTextSearch(t, testCase.textSearch)
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusOK, response.StatusCode)
			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.JSONEq(t, testCase.expectedResponse, string(body))
		})
	}
}

// TestGetKnownAssetsWithTextSearchLimit verifies that text searches are capped at 20 results while
// an empty or whitespace-only textSearch parameter continues to return the complete asset list.
//
// Authored by: OpenCode
func TestGetKnownAssetsWithTextSearchLimit(t *testing.T) {
	const matchingAssetCount = 21
	const initialAssetCount = 7

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery("DELETE FROM asset WHERE ticker LIKE 'TEST:TEXTSEARCH-LIMIT-%'", nil).
			Build(t),
	)

	for index := 1; index <= matchingAssetCount; index++ {
		var ticker = fmt.Sprintf("TEST:TEXTSEARCH-LIMIT-%02d", index)
		var name = fmt.Sprintf("Text Search Limit Asset %02d", index)
		err := infra.ExecuteDBQuery(
			"INSERT INTO asset (ticker, name) VALUES ({:ticker}, {:name})",
			dbx.Params{"ticker": ticker, "name": name},
		)
		assert.NoError(t, err)
	}

	response, err := getAssetsByTextSearch(t, "textsearch-limit")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)

	var searchedAssets []struct {
		Ticker string `json:"ticker"`
	}
	err = json.Unmarshal(body, &searchedAssets)
	assert.NoError(t, err)
	assert.Len(t, searchedAssets, 20)
	assert.Equal(t, "TEST:TEXTSEARCH-LIMIT-01", searchedAssets[0].Ticker)
	assert.Equal(t, "TEST:TEXTSEARCH-LIMIT-20", searchedAssets[19].Ticker)

	blankSearchResponse, err := getAssetsByTextSearch(t, "")
	assert.NoError(t, err)
	defer deferCloseResponseBody(blankSearchResponse)

	assert.Equal(t, http.StatusOK, blankSearchResponse.StatusCode)
	blankSearchBody, err := io.ReadAll(blankSearchResponse.Body)
	assert.NoError(t, err)

	var allAssets []json.RawMessage
	err = json.Unmarshal(blankSearchBody, &allAssets)
	assert.NoError(t, err)
	assert.Len(t, allAssets, initialAssetCount+matchingAssetCount)

	whitespaceSearchResponse, err := getAssetsByTextSearch(t, " \t ")
	assert.NoError(t, err)
	defer deferCloseResponseBody(whitespaceSearchResponse)

	assert.Equal(t, http.StatusOK, whitespaceSearchResponse.StatusCode)
	whitespaceSearchBody, err := io.ReadAll(whitespaceSearchResponse.Body)
	assert.NoError(t, err)

	var whitespaceSearchAssets []json.RawMessage
	err = json.Unmarshal(whitespaceSearchBody, &whitespaceSearchAssets)
	assert.NoError(t, err)
	assert.Len(t, whitespaceSearchAssets, initialAssetCount+matchingAssetCount)
}

// TestGetKnownAssetsRejectsInvalidTextSearch verifies request-length and parsed-term limits and
// confirms that nonblank input producing no terms returns an empty result.
//
// Authored by: OpenCode
func TestGetKnownAssetsRejectsInvalidTextSearch(t *testing.T) {
	t.Run("request length exceeds maximum", func(t *testing.T) {
		var response, err = getAssetsByTextSearch(t, strings.Repeat("a", 101))
		assert.NoError(t, err)
		defer deferCloseResponseBody(response)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{
			"errorMessage": "Validation failed",
			"details": ["Field 'textSearch' failed validation: must not exceed 100"]
		}`, string(body))
	})

	t.Run("parsed term count exceeds maximum", func(t *testing.T) {
		var textSearch = strings.TrimSpace(strings.Repeat("a ", 21))
		var response, err = getAssetsByTextSearch(t, textSearch)
		assert.NoError(t, err)
		defer deferCloseResponseBody(response)

		assert.Equal(t, http.StatusBadRequest, response.StatusCode)
		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `{
			"errorMessage": "Asset text search validation failed",
			"details": ["Text search must not exceed 20 terms"]
		}`, string(body))
	})

	t.Run("nonblank input producing no terms", func(t *testing.T) {
		var response, err = getAssetsByTextSearch(t, `""`)
		assert.NoError(t, err)
		defer deferCloseResponseBody(response)

		assert.Equal(t, http.StatusOK, response.StatusCode)
		body, err := io.ReadAll(response.Body)
		assert.NoError(t, err)
		assert.JSONEq(t, `[]`, string(body))
	})
}

// getAssetsByTextSearch sends an encoded textSearch request to the known-assets endpoint.
//
// Authored by: OpenCode
func getAssetsByTextSearch(t *testing.T, textSearch string) (*http.Response, error) {
	t.Helper()

	var queryValues = url.Values{}
	queryValues.Set("textSearch", textSearch)
	return http.Get(infra.TestAPIURLPrefix + "/asset?" + queryValues.Encode())
}

func TestGetAssetByIdOrTicker(t *testing.T) {

	t.Run(
		"TestGetAssetById",
		func(t *testing.T) {

			// Test retrieving a valid asset
			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/1")
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusOK, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var actualResponseJSON = string(body)
			var expectedResponseJSON = `
				{
					"id": 1,
					"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"ticker": "ARCA:BIL"
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)

	t.Run(
		"TestGetAssetByTicker",
		func(t *testing.T) {

			// Test retrieving a valid asset
			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/ARCA:BIL")
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusOK, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var actualResponseJSON = string(body)
			var expectedResponseJSON = `
				{
					"id": 1,
					"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"ticker": "ARCA:BIL"
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)
}

// TestPostAsset creates an asset with multiple external data entries and verifies that transient
// external asset display fields are excluded from both the response and persistence.
//
// Authored by: OpenCode
func TestPostAsset(t *testing.T) {
	var testTicker = "TEST:POST-ASSET"
	var testName = "Test Asset Created"
	var externalDataRequestJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU-CREATED","exchangeId":"PCX-CREATED","name":"Transient Name","exchangeName":"Transient Exchange"},{"source":"YAHOO_FINANCE","ticker":"BIL-CREATED","exchangeId":"SECOND-CREATED"}]}`
	var persistedExternalDataJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU-CREATED","exchangeId":"PCX-CREATED"},{"source":"YAHOO_FINANCE","ticker":"BIL-CREATED","exchangeId":"SECOND-CREATED"}]}`

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery("DELETE FROM asset WHERE ticker={:ticker}", dbx.Params{"ticker": testTicker}).
			Build(t),
	)

	var postAssetJSON = `
		{
			"name":"` + testName + `",
			"ticker":"` + testTicker + `",
			"externalData":` + externalDataRequestJSON + `
		}
	`
	response := postAsset(t, postAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"name":"` + testName + `",
			"ticker":"` + testTicker + `",
			"externalData":` + persistedExternalDataJSON + `
		}
	`
	inttestutil.AssertJSONEqualIgnoringFields(t, expectedResponseJSON, string(body), "id")

	var responseAsset struct {
		Id *int64 `json:"id"`
	}
	err = json.Unmarshal(body, &responseAsset)
	assert.NoError(t, err)
	if assert.NotNil(t, responseAsset.Id) {
		assert.NotZero(t, *responseAsset.Id)
		assertPersistedAssetWithExternalData(
			t,
			*responseAsset.Id,
			testTicker,
			testName,
			&persistedExternalDataJSON,
		)
	}
}

// TestPostAssetWithoutExternalData creates an asset without external data and verifies that the
// optional database column remains NULL.
//
// Authored by: OpenCode
func TestPostAssetWithoutExternalData(t *testing.T) {
	var testTicker = "TEST:POST-ASSET-NO-EXTERNAL"
	var testName = "Test Asset Without External Data"

	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery("DELETE FROM asset WHERE ticker={:ticker}", dbx.Params{"ticker": testTicker}).
			Build(t),
	)

	var postAssetJSON = `
		{
			"name":"` + testName + `",
			"ticker":"` + testTicker + `"
		}
	`
	response := postAsset(t, postAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusCreated, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"name":"` + testName + `",
			"ticker":"` + testTicker + `"
		}
	`
	inttestutil.AssertJSONEqualIgnoringFields(t, expectedResponseJSON, string(body), "id")

	var responseAsset struct {
		Id *int64 `json:"id"`
	}
	err = json.Unmarshal(body, &responseAsset)
	assert.NoError(t, err)
	if assert.NotNil(t, responseAsset.Id) {
		assert.NotZero(t, *responseAsset.Id)
		assertPersistedAsset(t, *responseAsset.Id, testTicker, testName)
	}
}

// TestPostAssetRejectsNonZeroId verifies that POST does not accept a client-supplied asset ID and
// does not insert the rejected asset.
//
// Authored by: OpenCode
func TestPostAssetRejectsNonZeroId(t *testing.T) {
	var testTicker = "TEST:POST-ASSET-ID"
	var postAssetJSON = `
		{
			"id":999999,
			"name":"Asset With Client ID",
			"ticker":"` + testTicker + `"
		}
	`

	response := postAsset(t, postAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	var expectedResponseJSON = `
		{
			"errorMessage":"Validation failed",
			"details":["Field 'id' failed validation: Asset ID must be omitted or zero for creation"]
		}
	`
	assert.JSONEq(t, expectedResponseJSON, string(body))

	var assetCount int
	err = infra.FetchWithDBQuery(
		"SELECT COUNT(*) FROM asset WHERE ticker = {:ticker}",
		dbx.Params{"ticker": testTicker},
		func(rows *dbx.Rows) error {
			return rows.Scan(&assetCount)
		},
	)
	assert.NoError(t, err)
	assert.Zero(t, assetCount)
}

// TestPostAssetFailureWithMissingRequiredFields verifies validation for asset fields and nested
// external asset fields.
//
// Authored by: OpenCode
func TestPostAssetFailureWithMissingRequiredFields(t *testing.T) {
	t.Run(
		"WithoutName",
		func(t *testing.T) {
			var postAssetJSON = `{"ticker":"TEST:POST-ASSET-NAME"}`
			var actualResponseJSON = string(postAssetForValidationFailure(t, postAssetJSON))
			var expectedResponseJSON = `
				{
					"errorMessage":"Validation failed",
					"details":["Field 'name' failed validation: is required"]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)

	t.Run(
		"WithoutTicker",
		func(t *testing.T) {
			var postAssetJSON = `{"name":"Asset Without Ticker"}`
			var actualResponseJSON = string(postAssetForValidationFailure(t, postAssetJSON))
			var expectedResponseJSON = `
				{
					"errorMessage":"Validation failed",
					"details":["Field 'ticker' failed validation: is required"]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)

	t.Run(
		"WithoutExternalSource",
		func(t *testing.T) {
			var postAssetJSON = `
				{
					"name":"Asset Without External Source",
					"ticker":"TEST:POST-ASSET-SOURCE",
					"externalData":{"data":[{"ticker":"IAU","exchangeId":"PCX"}]}
				}
			`
			var actualResponseJSON = string(postAssetForValidationFailure(t, postAssetJSON))
			var expectedResponseJSON = `
				{
					"errorMessage":"Validation failed",
					"details":["Field 'externalData.data[0].source' failed validation: is required"]
				}
			`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)
}

// postAssetForValidationFailure sends an invalid asset creation request and returns its response
// body after asserting the expected bad-request status.
//
// Authored by: OpenCode
func postAssetForValidationFailure(t *testing.T, postAssetJSON string) []byte {
	t.Helper()

	var response = postAsset(t, postAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)
	return body
}

// TestPostAssetWithDuplicateTicker verifies that a duplicate asset ticker returns HTTP 409 and
// does not create an additional record.
//
// Authored by: OpenCode
func TestPostAssetWithDuplicateTicker(t *testing.T) {
	var testTicker = "TEST:POST-ASSET-DUPLICATE"
	insertTestAsset(t, testTicker, "Existing Asset")

	var postAssetJSON = `
		{
			"name":"Duplicate Asset",
			"ticker":"` + testTicker + `"
		}
	`
	response := postAsset(t, postAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusConflict, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	var expectedResponseJSON = `
		{
			"errorMessage":"Asset already exists",
			"details":["Asset with ticker ` + testTicker + ` already exists"]
		}
	`
	assert.JSONEq(t, expectedResponseJSON, string(body))

	var assetCount int
	err = infra.FetchWithDBQuery(
		"SELECT COUNT(*) FROM asset WHERE ticker = {:ticker}",
		dbx.Params{"ticker": testTicker},
		func(rows *dbx.Rows) error {
			return rows.Scan(&assetCount)
		},
	)
	assert.NoError(t, err)
	assert.Equal(t, 1, assetCount)
}

// TestPutAsset tests the PUT /api/asset endpoint to replace an existing asset's ticker, name, and
// complete external data payload.
//
// Co-authored by: GitHub Copilot and OpenCode
func TestPutAsset(t *testing.T) {

	var testTickerBefore = "TEST:BEFORE"
	var testNameBefore = "Test Asset Before Update"
	var testTickerAfter = "TEST:AFTER"
	var testNameAfter = "Test Asset After Update"
	var testExternalDataBeforeJSON = testExternalAssetDataJSON
	var testExternalDataAfterRequestJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU-AFTER","exchangeId":"PCX-AFTER","name":"Transient Name","exchangeName":"Transient Exchange"},{"source":"YAHOO_FINANCE","ticker":"BIL-AFTER","exchangeId":"SECOND-AFTER"}]}`
	var testExternalDataAfterJSON = `{"data":[{"source":"YAHOO_FINANCE","ticker":"IAU-AFTER","exchangeId":"PCX-AFTER"},{"source":"YAHOO_FINANCE","ticker":"BIL-AFTER","exchangeId":"SECOND-AFTER"}]}`

	var testAsset = insertTestAsset(t, testTickerBefore, testNameBefore)
	var testAssetIdString = strconv.FormatInt(testAsset.Id, 10)

	var err = infra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{
			"externalData": testExternalDataBeforeJSON,
			"id":           testAsset.Id,
		},
	)
	assert.NoError(t, err)

	var putAssetJSON = `
		{
			"id":` + testAssetIdString + `,
			"ticker":"` + testTickerAfter + `",
			"name":"` + testNameAfter + `",
			"externalData":` + testExternalDataAfterRequestJSON + `
		}
	`

	response := putAsset(t, putAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"id":` + testAssetIdString + `,
			"ticker":"` + testTickerAfter + `",
			"name":"` + testNameAfter + `",
			"externalData":` + testExternalDataAfterJSON + `
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))

	assertPersistedAssetWithExternalData(
		t,
		testAsset.Id,
		testTickerAfter,
		testNameAfter,
		&testExternalDataAfterJSON,
	)
}

// TestPutAssetClearsExternalDataWhenOmitted verifies that a PUT request without externalData
// clears the existing persisted external data.
//
// Authored by: OpenCode
func TestPutAssetClearsExternalDataWhenOmitted(t *testing.T) {
	var testAsset = insertTestAsset(t, "TEST:CLEAR-OMITTED", "Test Asset Clear Omitted")
	var testExternalDataJSON = testExternalAssetDataJSON

	var err = infra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{
			"externalData": testExternalDataJSON,
			"id":           testAsset.Id,
		},
	)
	assert.NoError(t, err)

	var putAssetJSON = `
		{
			"id":` + strconv.FormatInt(testAsset.Id, 10) + `,
			"ticker":"TEST:CLEAR-OMITTED-AFTER",
			"name":"Test Asset Clear Omitted After"
		}
	`
	response := putAsset(t, putAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `
		{
			"id":`+strconv.FormatInt(testAsset.Id, 10)+`,
			"ticker":"TEST:CLEAR-OMITTED-AFTER",
			"name":"Test Asset Clear Omitted After"
		}
	`, string(body))

	assertPersistedAsset(t, testAsset.Id, "TEST:CLEAR-OMITTED-AFTER", "Test Asset Clear Omitted After")
}

// TestPutAssetClearsExternalDataWhenNull verifies that a PUT request with a null externalData
// value clears the existing persisted external data.
//
// Authored by: OpenCode
func TestPutAssetClearsExternalDataWhenNull(t *testing.T) {
	var testAsset = insertTestAsset(t, "TEST:CLEAR-NULL", "Test Asset Clear Null")
	var testExternalDataJSON = testExternalAssetDataJSON

	var err = infra.ExecuteDBQuery(
		"UPDATE asset SET external_data = {:externalData}::jsonb WHERE id = {:id}",
		dbx.Params{
			"externalData": testExternalDataJSON,
			"id":           testAsset.Id,
		},
	)
	assert.NoError(t, err)

	var putAssetJSON = `
		{
			"id":` + strconv.FormatInt(testAsset.Id, 10) + `,
			"ticker":"TEST:CLEAR-NULL-AFTER",
			"name":"Test Asset Clear Null After",
			"externalData":null
		}
	`
	response := putAsset(t, putAssetJSON)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusOK, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.JSONEq(t, `
		{
			"id":`+strconv.FormatInt(testAsset.Id, 10)+`,
			"ticker":"TEST:CLEAR-NULL-AFTER",
			"name":"Test Asset Clear Null After"
		}
	`, string(body))

	assertPersistedAsset(t, testAsset.Id, "TEST:CLEAR-NULL-AFTER", "Test Asset Clear Null After")
}

// TestPutAssetFailureWithoutId tests the PUT /api/asset endpoint returns a validation error
// when the asset ID is missing from the request body.
//
// Authored by: GitHub Copilot
func TestPutAssetFailureWithoutId(t *testing.T) {

	var putAssetJSONNoId = `
		{
			"ticker": "TEST:NOID",
			"name": "Asset Without ID"
		}
	`

	response := putAsset(t, putAssetJSONNoId)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'id' failed validation: is required"
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}

// TestPutAssetFailureWithZeroId tests the PUT /api/asset endpoint returns a validation error
// when the asset ID is zero in the request body.
//
// Authored by: GitHub Copilot
func TestPutAssetFailureWithZeroId(t *testing.T) {

	var putAssetJSONZeroId = `
		{
			"id": 0,
			"ticker": "TEST:ZEROID",
			"name": "Asset With Zero ID"
		}
	`

	response := putAsset(t, putAssetJSONZeroId)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var expectedResponseJSON = `
		{
			"errorMessage": "Validation failed",
			"details": [
				"Field 'id' failed validation: is required"
			]
		}
	`

	assert.JSONEq(t, expectedResponseJSON, string(body))
}

// TestPutAssetFailureWithoutRequiredFields tests the PUT /api/asset endpoint returns validation
// errors when required fields (name, ticker) are missing from the request body.
//
// Authored by: GitHub Copilot
func TestPutAssetFailureWithoutRequiredFields(t *testing.T) {

	t.Run(
		"WithoutName",
		func(t *testing.T) {

			var putAssetJSON = `
				{
					"id": 1,
					"ticker": "TEST:NONAME"
				}
			`

			response := putAsset(t, putAssetJSON)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'name' failed validation: is required"
					]
				}
			`

			assert.JSONEq(t, expectedResponseJSON, string(body))
		},
	)

	t.Run(
		"WithoutTicker",
		func(t *testing.T) {

			var putAssetJSON = `
				{
					"id": 1,
					"name": "Asset Without Ticker"
				}
			`

			response := putAsset(t, putAssetJSON)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'ticker' failed validation: is required"
					]
				}
			`

			assert.JSONEq(t, expectedResponseJSON, string(body))
		},
	)

	t.Run(
		"WithoutNameAndTicker",
		func(t *testing.T) {

			var putAssetJSON = `
				{
					"id": 1
				}
			`

			response := putAsset(t, putAssetJSON)
			defer deferCloseResponseBody(response)

			assert.Equal(t, http.StatusBadRequest, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			var expectedResponseJSON = `
				{
					"errorMessage": "Validation failed",
					"details": [
						"Field 'name' failed validation: is required",
						"Field 'ticker' failed validation: is required"
					]
				}
			`

			assert.JSONEq(t, expectedResponseJSON, string(body))
		},
	)
}

// TestGetAssetByIdNotFound tests the GET /api/asset/{id} endpoint with a non-existent asset ID.
//
// Authored by: GitHub Copilot
func TestGetAssetByIdNotFound(t *testing.T) {

	// Test with a non-existent asset ID
	response, err := http.Get(infra.TestAPIURLPrefix + "/asset/999")
	assert.NoError(t, err)
	defer deferCloseResponseBody(response)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)

	body, err := io.ReadAll(response.Body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body)

	var actualResponseJSON = string(body)
	var expectedResponseJSON = `
		{
			"errorMessage": "Data not found",
			"details": [
				"Asset with identifier 999 not found"
			]
		}
	`
	assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
}

// TestGetAssetByIdInvalidId tests the GET /api/asset/{id} endpoint with invalid ID formats.
//
// Authored by: GitHub Copilot
func TestGetAssetByIdInvalidId(t *testing.T) {

	// Test case: Empty ID (Gin redirects /api/asset/ to /api/asset)
	t.Run(
		"InvalidIdEmpty",
		func(t *testing.T) {

			response, err := http.Get(infra.TestAPIURLPrefix + "/asset/")
			assert.NoError(t, err)
			defer deferCloseResponseBody(response)

			// Gin redirects /api/asset/ to /api/asset and returns all assets
			assert.Equal(t, http.StatusOK, response.StatusCode)

			body, err := io.ReadAll(response.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)

			// Should return the complete list of assets (same as GET /api/asset)
			var actualResponseJSON = string(body)
			var expectedResponseJSON = `
			[
				{
					"id": 1,
					"name": "SPDR Bloomberg 1-3 Month T-Bill ETF",
					"ticker": "ARCA:BIL"
				},
				{
					"id": 6,
					"name": "iShares Msci Brazil ETF",
					"ticker": "ARCA:EWZ"
				},
				{
					"id": 7,
					"name": "SPDR S&P 500 ETF Trust",
					"ticker": "ARCA:SPY"
				},
				{
					"id": 2,
					"name": "iShares 0-5 Year TIPS Bond ETF",
					"ticker": "ARCA:STIP"
				},
				{
					"id": 3,
					"name": "iShares 7-10 Year Treasury Bond ETF",
					"ticker": "NasdaqGM:IEF"
				},
				{
					"id": 5,
					"name": "iShares Short Treasury Bond ETF",
					"ticker": "NasdaqGM:SHV"
				},
				{
					"id": 4,
					"name": "iShares 20+ Year Treasury Bond ETF",
					"ticker": "NasdaqGM:TLT"
				}
			]
		`
			assert.JSONEq(t, expectedResponseJSON, actualResponseJSON)
		},
	)
}
