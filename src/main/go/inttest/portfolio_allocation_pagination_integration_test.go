package inttest

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"testing"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/domain"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// observationPaginationFixture groups isolated persisted observations and their expected database order.
// Authored by: OpenCode
type observationPaginationFixture struct {
	portfolio domain.Portfolio
	prefix    string
	ordered   []observationPaginationRecord
}

// observationPaginationRecord is one persisted observation used by pagination assertions.
// Authored by: OpenCode
type observationPaginationRecord struct {
	id        int64
	tag       string
	timestamp time.Time
}

// observationPageResponse mirrors the opt-in observation-header envelope.
// Authored by: OpenCode
type observationPageResponse struct {
	Items      []observationPageItem `json:"items"`
	NextCursor *string               `json:"nextCursor"`
}

// observationPageItem contains the header fields used in ordering assertions.
// Authored by: OpenCode
type observationPageItem struct {
	Id        int64     `json:"id"`
	TimeTag   string    `json:"timeTag"`
	Timestamp time.Time `json:"timestamp"`
}

// historyPageResponse mirrors the opt-in complete-history envelope.
// Authored by: OpenCode
type historyPageResponse struct {
	Items      []historyPageItem `json:"items"`
	NextCursor *string           `json:"nextCursor"`
}

// historyPageItem contains the snapshot projection used in persistence-read assertions.
// Authored by: OpenCode
type historyPageItem struct {
	ObservationTimestamp observationPageItem     `json:"observationTimestamp"`
	Allocations          []historyPageAllocation `json:"allocations"`
	TotalMarketValue     string                  `json:"totalMarketValue"`
}

// historyPageAllocation projects an allocation market value used in total checks.
// Authored by: OpenCode
type historyPageAllocation struct {
	TotalMarketValue string `json:"totalMarketValue"`
}

// paginationValidationResponse mirrors the standard API validation error used for malformed cursors.
// Authored by: OpenCode
type paginationValidationResponse struct {
	ErrorMessage string   `json:"errorMessage"`
	Details      []string `json:"details"`
}

// TestPortfolioObservationPaginationCoversCardinalitiesAndLegacyReads verifies paginated header and snapshot reads,
// stable ordering, complete snapshot details, look-ahead behavior, and unchanged unpaginated contracts.
// Authored by: OpenCode
func TestPortfolioObservationPaginationCoversCardinalitiesAndLegacyReads(t *testing.T) {
	var fixtures = make([]observationPaginationFixture, 0, 6)
	for _, count := range []int{0, 1, 10, 11, 20, 21} {
		fixtures = append(fixtures, createObservationPaginationFixture(t, count))
	}

	for _, fixture := range fixtures {
		t.Run(fmt.Sprintf("records_%d", len(fixture.ordered)), func(t *testing.T) {
			verifyLegacyObservationReads(t, fixture)
			verifyObservationPages(t, fixture)
		})
	}

	var fixtureWithTwentyOne = fixtures[len(fixtures)-1]
	var firstPage = getObservationPage(t, fixtureWithTwentyOne.portfolio.Id, nil)
	require.NotNil(t, firstPage.NextCursor)
	var newTag = fixtureWithTwentyOne.prefix + "inserted-after-first-page"
	var insertedTimestamp = time.Date(2055, time.January, 1, 0, 0, 0, 0, time.UTC)
	insertObservationPaginationRecord(t, fixtureWithTwentyOne, newTag, insertedTimestamp)

	var continuingPage = getObservationPage(t, fixtureWithTwentyOne.portfolio.Id, firstPage.NextCursor)
	assert.NotContains(t, observationPageTags(continuingPage), newTag)
	var freshTraversal = getObservationPage(t, fixtureWithTwentyOne.portfolio.Id, nil)
	require.NotEmpty(t, freshTraversal.Items)
	assert.Equal(t, newTag, freshTraversal.Items[0].TimeTag)

	var malformedResponse = getPaginationResponse(
		t,
		observationEndpoint(fixtureWithTwentyOne.portfolio.Id, "history/observation")+"?pageCursor=not-a-cursor",
	)
	assertPaginationValidationResponse(t, malformedResponse)
	var malformedEncodingResponse = getPaginationResponse(
		t,
		observationEndpoint(fixtureWithTwentyOne.portfolio.Id, "history/observation")+
			"?pageCursor=%ZZ",
	)
	assert.Equal(t, http.StatusBadRequest, malformedEncodingResponse.StatusCode)
	closePaginationResponse(t, malformedEncodingResponse)

	var duplicateResponse = getPaginationResponse(
		t,
		observationEndpoint(fixtureWithTwentyOne.portfolio.Id, "history/observation")+
			"?pageCursor=&pageCursor=",
	)
	assert.Equal(t, http.StatusBadRequest, duplicateResponse.StatusCode)
	closePaginationResponse(t, duplicateResponse)

	var wrongPortfolioResponse = getPaginationResponse(
		t,
		observationEndpoint(fixtures[0].portfolio.Id, "history/observation")+
			"?pageCursor="+url.QueryEscape(*firstPage.NextCursor),
	)
	assert.Equal(t, http.StatusBadRequest, wrongPortfolioResponse.StatusCode)
	closePaginationResponse(t, wrongPortfolioResponse)

	var fullHistoryFirstPage = getHistoryPage(t, fixtureWithTwentyOne.portfolio.Id, nil)
	require.NotEmpty(t, fullHistoryFirstPage.Items)
	var conflictingObservationId = fullHistoryFirstPage.Items[0].ObservationTimestamp.Id
	var conflictingResponse = getPaginationResponse(
		t,
		fmt.Sprintf(
			"%s?observationTimestampId=%d&pageCursor=",
			observationEndpoint(fixtureWithTwentyOne.portfolio.Id, "history"),
			conflictingObservationId,
		),
	)
	assertPaginationValidationResponse(t, conflictingResponse)
}

// createObservationPaginationFixture creates one portfolio and the requested observations with deterministic timestamps.
// Authored by: OpenCode
func createObservationPaginationFixture(t *testing.T, count int) observationPaginationFixture {
	t.Helper()
	var prefix = fmt.Sprintf("pr9-observation-%d-", count)
	var portfolio = insertTestPortfolio(t, "PR 9 observation pagination "+fmt.Sprint(count))
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery("DELETE FROM portfolio_allocation_fact WHERE portfolio_id = {:portfolioId}", dbx.Params{
				"portfolioId": portfolio.Id,
			}).
			AddCleanupQuery("DELETE FROM portfolio_allocation_obs_time WHERE portfolio_id = {:portfolioId}", dbx.Params{
				"portfolioId": portfolio.Id,
			}).
			Build(t),
	)

	var priorities = observationPaginationPriorities(count)
	var baseTimestamp = time.Date(2040, time.January, 1, 0, 0, 0, 123456000, time.UTC)
	var fixture = observationPaginationFixture{portfolio: portfolio, prefix: prefix}
	for index, priority := range priorities {
		var tag = fmt.Sprintf("%s%02d", prefix, index)
		var timestamp = baseTimestamp.Add(time.Duration(priority) * time.Microsecond)
		insertObservationPaginationRecord(t, fixture, tag, timestamp)
	}

	var queryError = inttestinfra.FetchWithDBQuery(
		`SELECT id, observation_time_tag, observation_timestamp
		 FROM portfolio_allocation_obs_time
		 WHERE portfolio_id = {:portfolioId} AND observation_time_tag LIKE {:tagPrefix}`,
		dbx.Params{"portfolioId": portfolio.Id, "tagPrefix": prefix + "%"},
		func(rows *dbx.Rows) error {
			var record observationPaginationRecord
			if scanError := rows.Scan(&record.id, &record.tag, &record.timestamp); scanError != nil {
				return scanError
			}
			fixture.ordered = append(fixture.ordered, record)
			return nil
		},
	)
	require.NoError(t, queryError)
	sort.Slice(fixture.ordered, func(left, right int) bool {
		if fixture.ordered[left].timestamp.Equal(fixture.ordered[right].timestamp) {
			return fixture.ordered[left].id > fixture.ordered[right].id
		}
		return fixture.ordered[left].timestamp.After(fixture.ordered[right].timestamp)
	})
	return fixture
}

// observationPaginationPriorities returns timestamps whose rank-ten boundary includes tied records at microsecond precision.
// Authored by: OpenCode
func observationPaginationPriorities(count int) []int {
	var priorities = []int{17, 2, 25, 20, 18, 3, 17, 7, 21, 5, 0, 24, 17, 6, 19, 1, 17, 8, 22, 4, 23}
	if count == 21 {
		return priorities
	}
	var result = make([]int, count)
	for index := range result {
		result[index] = (index * 7) % 31
	}
	return result
}

// insertObservationPaginationRecord inserts one owned observation and its single allocation fixture.
// Authored by: OpenCode
func insertObservationPaginationRecord(
	t *testing.T,
	fixture observationPaginationFixture,
	tag string,
	timestamp time.Time,
) {
	t.Helper()
	var err = inttestinfra.ExecuteDBQuery(
		`INSERT INTO portfolio_allocation_obs_time (portfolio_id, observation_time_tag, observation_timestamp)
		 VALUES ({:portfolioId}, {:tag}, {:timestamp})`,
		dbx.Params{"portfolioId": fixture.portfolio.Id, "tag": tag, "timestamp": timestamp},
	)
	require.NoError(t, err)
	err = inttestinfra.ExecuteDBQuery(
		`INSERT INTO portfolio_allocation_fact (
			asset_id, "class", cash_reserve, asset_quantity, asset_market_price,
			total_market_value, portfolio_id, observation_time_id
		)
		SELECT 1, 'BONDS', FALSE, 1, 100, 100, observation.portfolio_id, observation.id
		FROM portfolio_allocation_obs_time observation
		WHERE observation.portfolio_id = {:portfolioId} AND observation.observation_time_tag = {:tag}`,
		dbx.Params{"portfolioId": fixture.portfolio.Id, "tag": tag},
	)
	require.NoError(t, err)
	if tag == fixture.prefix+"12" {
		err = inttestinfra.ExecuteDBQuery(
			`INSERT INTO portfolio_allocation_fact (
				asset_id, "class", cash_reserve, asset_quantity, asset_market_price,
				total_market_value, portfolio_id, observation_time_id
			)
			SELECT 1, 'PR9-EXTRA-' || detail_index::text, FALSE, 1, 10, 10,
				observation.portfolio_id, observation.id
			FROM portfolio_allocation_obs_time observation
		CROSS JOIN generate_series(1, 12) AS detail_index
			WHERE observation.portfolio_id = {:portfolioId} AND observation.observation_time_tag = {:tag}`,
			dbx.Params{"portfolioId": fixture.portfolio.Id, "tag": tag},
		)
		require.NoError(t, err)
	}
}

// verifyLegacyObservationReads checks that requests without pageCursor retain their existing array responses and limits.
// Authored by: OpenCode
func verifyLegacyObservationReads(t *testing.T, fixture observationPaginationFixture) {
	t.Helper()
	var headersResponse = getPaginationResponse(
		t,
		observationEndpoint(fixture.portfolio.Id, "history/observation"),
	)
	require.Equal(t, http.StatusOK, headersResponse.StatusCode)
	var headersBody, err = io.ReadAll(headersResponse.Body)
	require.NoError(t, err)
	require.NoError(t, headersResponse.Body.Close())
	var headers []observationPageItem
	require.NoError(t, json.Unmarshal(headersBody, &headers))
	assert.NotNil(t, headers)
	assert.Len(t, headers, min(len(fixture.ordered), 10))
	assert.NotContains(t, string(headersBody), `"items"`)

	var historyResponse = getPaginationResponse(
		t,
		observationEndpoint(fixture.portfolio.Id, "history"),
	)
	require.Equal(t, http.StatusOK, historyResponse.StatusCode)
	var historyBody, historyError = io.ReadAll(historyResponse.Body)
	require.NoError(t, historyError)
	require.NoError(t, historyResponse.Body.Close())
	var history []historyPageItem
	require.NoError(t, json.Unmarshal(historyBody, &history))
	assert.NotNil(t, history)
	assert.Len(t, history, min(len(fixture.ordered), 10))
	assert.NotContains(t, string(historyBody), `"nextCursor"`)
	if len(fixture.ordered) > 10 {
		var detailObservationId = fixture.ordered[len(fixture.ordered)-1].id
		var detailResponse = getPaginationResponse(
			t,
			fmt.Sprintf(
				"%s?observationTimestampId=%d",
				observationEndpoint(fixture.portfolio.Id, "history"),
				detailObservationId,
			),
		)
		require.Equal(t, http.StatusOK, detailResponse.StatusCode)
		var detailBody, detailError = io.ReadAll(detailResponse.Body)
		require.NoError(t, detailError)
		require.NoError(t, detailResponse.Body.Close())
		var detailHistory []historyPageItem
		require.NoError(t, json.Unmarshal(detailBody, &detailHistory))
		require.Len(t, detailHistory, 1)
		assert.Equal(t, detailObservationId, detailHistory[0].ObservationTimestamp.Id)
	}
	for _, snapshot := range history {
		if snapshot.ObservationTimestamp.TimeTag == fixture.prefix+"12" {
			assert.Len(t, snapshot.Allocations, 13)
			assert.Equal(t, "220", snapshot.TotalMarketValue)
		} else {
			assert.Len(t, snapshot.Allocations, 1)
			assert.Equal(t, "100", snapshot.TotalMarketValue)
		}
	}
}

// verifyObservationPages traverses headers and full snapshots together and checks every parent appears exactly once.
// Authored by: OpenCode
func verifyObservationPages(t *testing.T, fixture observationPaginationFixture) {
	t.Helper()
	var cursor *string
	var collectedHeaderIDs = make([]int64, 0, len(fixture.ordered))
	var collectedHistoryIDs = make([]int64, 0, len(fixture.ordered))
	for pageIndex := 0; ; pageIndex++ {
		require.Less(t, pageIndex, 10, "pagination must terminate")
		var headerPage = getObservationPage(t, fixture.portfolio.Id, cursor)
		var historyPage = getHistoryPage(t, fixture.portfolio.Id, cursor)
		require.NotNil(t, headerPage.Items)
		require.NotNil(t, historyPage.Items)
		require.LessOrEqual(t, len(headerPage.Items), 10)
		require.LessOrEqual(t, len(historyPage.Items), 10)
		assert.Equal(t, headerPage.NextCursor, historyPage.NextCursor)
		assert.Equal(t, len(headerPage.Items), len(historyPage.Items))
		for index, item := range headerPage.Items {
			var snapshot = historyPage.Items[index]
			assert.Equal(t, item.Id, snapshot.ObservationTimestamp.Id)
			assert.Equal(t, item.TimeTag, snapshot.ObservationTimestamp.TimeTag)
			assert.True(t, item.Timestamp.Equal(snapshot.ObservationTimestamp.Timestamp))
			if item.TimeTag == fixture.prefix+"12" {
				assert.Len(t, snapshot.Allocations, 13)
				assert.Equal(t, "220", snapshot.TotalMarketValue)
			} else {
				assert.Len(t, snapshot.Allocations, 1)
				assert.Equal(t, "100", snapshot.TotalMarketValue)
				assert.Equal(t, "100", snapshot.Allocations[0].TotalMarketValue)
			}
			collectedHeaderIDs = append(collectedHeaderIDs, item.Id)
			collectedHistoryIDs = append(collectedHistoryIDs, snapshot.ObservationTimestamp.Id)
		}
		cursor = headerPage.NextCursor
		if cursor == nil {
			break
		}
	}

	var expectedIDs = make([]int64, len(fixture.ordered))
	for index, record := range fixture.ordered {
		expectedIDs[index] = record.id
	}
	assert.Equal(t, expectedIDs, collectedHeaderIDs)
	assert.Equal(t, expectedIDs, collectedHistoryIDs)
	assert.Len(t, uniqueInt64s(collectedHeaderIDs), len(fixture.ordered))
}

// getObservationPage requests one paginated observation-header response using the returned opaque cursor unchanged.
// Authored by: OpenCode
func getObservationPage(t *testing.T, portfolioId int64, cursor *string) observationPageResponse {
	t.Helper()
	var endpoint = observationEndpoint(portfolioId, "history/observation") + "?pageCursor="
	if cursor != nil {
		endpoint += url.QueryEscape(*cursor)
	}
	var response = getPaginationResponse(t, endpoint)
	defer closePaginationResponse(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode)
	var body, err = io.ReadAll(response.Body)
	require.NoError(t, err)
	var page observationPageResponse
	require.NoError(t, json.Unmarshal(body, &page))
	assert.Contains(t, string(body), `"nextCursor"`)
	if page.NextCursor == nil {
		assert.Contains(t, string(body), `"nextCursor":null`)
	}
	return page
}

// getHistoryPage requests one paginated full-history response using the returned opaque cursor unchanged.
// Authored by: OpenCode
func getHistoryPage(t *testing.T, portfolioId int64, cursor *string) historyPageResponse {
	t.Helper()
	var endpoint = observationEndpoint(portfolioId, "history") + "?pageCursor="
	if cursor != nil {
		endpoint += url.QueryEscape(*cursor)
	}
	var response = getPaginationResponse(t, endpoint)
	defer closePaginationResponse(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode)
	var body, err = io.ReadAll(response.Body)
	require.NoError(t, err)
	var page historyPageResponse
	require.NoError(t, json.Unmarshal(body, &page))
	assert.Contains(t, string(body), `"nextCursor"`)
	if page.NextCursor == nil {
		assert.Contains(t, string(body), `"nextCursor":null`)
	}
	return page
}

// observationEndpoint builds an existing observation endpoint URL for one portfolio.
// Authored by: OpenCode
func observationEndpoint(portfolioId int64, suffix string) string {
	return fmt.Sprintf("%s/portfolio/%d/%s", inttestinfra.TestAPIURLPrefix, portfolioId, suffix)
}

// getPaginationResponse performs an HTTP GET for one pagination integration assertion.
// Authored by: OpenCode
func getPaginationResponse(t *testing.T, endpoint string) *http.Response {
	t.Helper()
	response, err := http.Get(endpoint)
	require.NoError(t, err)
	return response
}

// closePaginationResponse closes an HTTP response body and reports cleanup failures to the test.
// Authored by: OpenCode
func closePaginationResponse(t *testing.T, response *http.Response) {
	t.Helper()
	assert.NoError(t, response.Body.Close())
}

// assertPaginationValidationResponse checks that invalid cursors produce one standard validation JSON document.
// Authored by: OpenCode
func assertPaginationValidationResponse(t *testing.T, response *http.Response) {
	t.Helper()
	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	var body, err = io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())
	var validationResponse paginationValidationResponse
	require.NoError(t, json.Unmarshal(body, &validationResponse))
	assert.Equal(t, "Validation failed", validationResponse.ErrorMessage)
	assert.NotEmpty(t, validationResponse.Details)
}

// observationPageTags returns page tags for save-refresh and insertion-boundary assertions.
// Authored by: OpenCode
func observationPageTags(page observationPageResponse) []string {
	var tags = make([]string, len(page.Items))
	for index, item := range page.Items {
		tags[index] = item.TimeTag
	}
	return tags
}

// uniqueInt64s returns the distinct IDs in an already traversed page sequence.
// Authored by: OpenCode
func uniqueInt64s(values []int64) map[int64]struct{} {
	var unique = make(map[int64]struct{}, len(values))
	for _, value := range values {
		unique[value] = struct{}{}
	}
	return unique
}
