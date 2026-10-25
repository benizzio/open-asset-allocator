package inttest

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"testing"
	"time"

	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/benizzio/open-asset-allocator/domain"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	inttestutil "github.com/benizzio/open-asset-allocator/inttest/util"
)

// allocationPlanPaginationFixture groups isolated plans and their expected database order.
// Authored by: OpenCode
type allocationPlanPaginationFixture struct {
	portfolio domain.Portfolio
	prefix    string
	ordered   []allocationPlanPaginationRecord
}

// allocationPlanPaginationRecord is one persisted plan used by pagination assertions.
// Authored by: OpenCode
type allocationPlanPaginationRecord struct {
	id          int64
	name        string
	plannedDate time.Time
}

// allocationPlanPageResponse mirrors the opt-in allocation-plan envelope.
// Authored by: OpenCode
type allocationPlanPageResponse struct {
	Items      []allocationPlanPageItem `json:"items"`
	NextCursor *string                  `json:"nextCursor"`
}

// allocationPlanPageItem projects fields used in plan-order and completeness assertions.
// Authored by: OpenCode
type allocationPlanPageItem struct {
	Id          int64                       `json:"id"`
	Name        string                      `json:"name"`
	Type        string                      `json:"type"`
	PlannedDate string                      `json:"plannedDate"`
	Details     []allocationPlanPageDetails `json:"details"`
}

// allocationPlanPageDetails projects hierarchy values used to verify complete plan reads.
// Authored by: OpenCode
type allocationPlanPageDetails struct {
	HierarchicalId []*string `json:"hierarchicalId"`
}

// TestAllocationPlanPaginationCoversCardinalitiesAndLegacyReads verifies complete-plan paging, stable keyset order,
// type and portfolio scope, cursor validation, look-ahead behavior, and the existing unbounded array response.
// Authored by: OpenCode
func TestAllocationPlanPaginationCoversCardinalitiesAndLegacyReads(t *testing.T) {
	var fixtures = make([]allocationPlanPaginationFixture, 0, 6)
	for _, count := range []int{0, 1, 10, 11, 20, 21} {
		fixtures = append(fixtures, createAllocationPlanPaginationFixture(t, count))
	}
	for _, fixture := range fixtures {
		t.Run(fmt.Sprintf("plans_%d", len(fixture.ordered)), func(t *testing.T) {
			verifyLegacyAllocationPlanRead(t, fixture)
			verifyAllocationPlanPages(t, fixture)
		})
	}

	var fixtureWithTwentyOne = fixtures[len(fixtures)-1]
	var firstPage = getAllocationPlanPage(t, fixtureWithTwentyOne.portfolio.Id, nil)
	require.NotNil(t, firstPage.NextCursor)
	var insertedName = fixtureWithTwentyOne.prefix + "inserted-after-first-page"
	insertAllocationPlanPaginationRecord(
		t,
		fixtureWithTwentyOne,
		insertedName,
		time.Date(2055, time.January, 1, 0, 0, 0, 0, time.UTC),
		"ALLOCATION_PLAN",
	)
	var continuingPage = getAllocationPlanPage(
		t,
		fixtureWithTwentyOne.portfolio.Id,
		firstPage.NextCursor,
	)
	assert.NotContains(t, allocationPlanPageNames(continuingPage), insertedName)
	var freshPage = getAllocationPlanPage(t, fixtureWithTwentyOne.portfolio.Id, nil)
	require.NotEmpty(t, freshPage.Items)
	assert.Equal(t, insertedName, freshPage.Items[0].Name)

	var malformedResponse = getPaginationResponse(
		t,
		allocationPlanEndpoint(fixtureWithTwentyOne.portfolio.Id)+"?pageCursor=not-a-cursor",
	)
	assert.Equal(t, http.StatusBadRequest, malformedResponse.StatusCode)
	closePaginationResponse(t, malformedResponse)
	var malformedEncodingResponse = getPaginationResponse(
		t,
		allocationPlanEndpoint(fixtureWithTwentyOne.portfolio.Id)+"?pageCursor=%ZZ",
	)
	assert.Equal(t, http.StatusBadRequest, malformedEncodingResponse.StatusCode)
	closePaginationResponse(t, malformedEncodingResponse)

	var duplicateResponse = getPaginationResponse(
		t,
		allocationPlanEndpoint(fixtureWithTwentyOne.portfolio.Id)+"?pageCursor=&pageCursor=",
	)
	assert.Equal(t, http.StatusBadRequest, duplicateResponse.StatusCode)
	closePaginationResponse(t, duplicateResponse)

	for _, cursor := range invalidAllocationPlanCursorTokens(fixtureWithTwentyOne.portfolio.Id) {
		var invalidResponse = getPaginationResponse(
			t,
			allocationPlanEndpoint(fixtureWithTwentyOne.portfolio.Id)+"?pageCursor="+url.QueryEscape(cursor),
		)
		assert.Equal(t, http.StatusBadRequest, invalidResponse.StatusCode)
		closePaginationResponse(t, invalidResponse)
	}

	var observationCursor = encodeTestCursor(map[string]any{
		"version": 1, "portfolioId": fixtureWithTwentyOne.portfolio.Id,
		"collection": "observations", "timestamp": "2040-01-01T00:00:00.123456Z", "id": 1,
	})
	var wrongCollectionResponse = getPaginationResponse(
		t,
		allocationPlanEndpoint(fixtureWithTwentyOne.portfolio.Id)+"?pageCursor="+url.QueryEscape(observationCursor),
	)
	assert.Equal(t, http.StatusBadRequest, wrongCollectionResponse.StatusCode)
	closePaginationResponse(t, wrongCollectionResponse)

	var otherPortfolioResponse = getPaginationResponse(
		t,
		allocationPlanEndpoint(fixtures[0].portfolio.Id)+"?pageCursor="+
			url.QueryEscape(*firstPage.NextCursor),
	)
	assert.Equal(t, http.StatusBadRequest, otherPortfolioResponse.StatusCode)
	closePaginationResponse(t, otherPortfolioResponse)
}

// TestAllocationPlanPaginationPreservesVersionOneTokensAndRejectsForeignFields verifies old flat JSON tokens
// still select the next plan and rejects observation fields even when their values are empty or null.
// For example, run go test ./inttest -run TestAllocationPlanPaginationPreservesVersionOneTokensAndRejectsForeignFields.
// Authored by: GPT-6.1 Sol
func TestAllocationPlanPaginationPreservesVersionOneTokensAndRejectsForeignFields(t *testing.T) {
	var fixture = createAllocationPlanPaginationFixture(t, 11)
	var boundary = fixture.ordered[9]
	var remaining = fixture.ordered[10]
	var legacyPayload = fmt.Sprintf(
		`{"version":1,"portfolioId":%d,"collection":"allocation-plans","plannedDate":%q,"id":%d,"planType":"ALLOCATION_PLAN"}`,
		fixture.portfolio.Id, boundary.plannedDate.Format(time.DateOnly), boundary.id,
	)
	var legacyToken = base64.RawURLEncoding.EncodeToString([]byte(legacyPayload))
	var page = getAllocationPlanPage(t, fixture.portfolio.Id, &legacyToken)
	require.Len(t, page.Items, 1)
	require.Len(t, page.Items[0].Details, 2)
	for _, detail := range page.Items[0].Details {
		require.Len(t, detail.HierarchicalId, 1)
		require.NotNil(t, detail.HierarchicalId[0])
	}
	// Detail order is outside this cursor contract; normalize the projection before comparing JSON.
	sort.Slice(page.Items[0].Details, func(left, right int) bool {
		return *page.Items[0].Details[left].HierarchicalId[0] < *page.Items[0].Details[right].HierarchicalId[0]
	})
	var pageJSON, err = json.Marshal(page)
	require.NoError(t, err)
	assert.JSONEq(t, fmt.Sprintf(
		`{"items":[{"id":%d,"name":%q,"type":"ALLOCATION_PLAN","plannedDate":%q,"details":[{"hierarchicalId":[%q]},{"hierarchicalId":[%q]}]}],"nextCursor":null}`,
		remaining.id, remaining.name, remaining.plannedDate.Format(time.DateOnly),
		remaining.name+"-first", remaining.name+"-second",
	), string(pageJSON))

	for _, foreignField := range []string{
		`"timestamp":"2040-01-01T00:00:00Z"`, `"timestamp":""`, `"timestamp":null`,
	} {
		t.Run(foreignField, func(t *testing.T) {
			var mixedPayload = strings.TrimSuffix(legacyPayload, "}") + "," + foreignField + "}"
			var token = base64.RawURLEncoding.EncodeToString([]byte(mixedPayload))
			var response = getPaginationResponse(
				t, allocationPlanEndpoint(fixture.portfolio.Id)+"?pageCursor="+url.QueryEscape(token),
			)
			assertPaginationValidationResponse(t, response)
		})
	}
}

// createAllocationPlanPaginationFixture creates one isolated portfolio and the requested plans with tied dates.
// Authored by: OpenCode
func createAllocationPlanPaginationFixture(t *testing.T, count int) allocationPlanPaginationFixture {
	t.Helper()
	var prefix = fmt.Sprintf("pr9-plan-%d-", count)
	var portfolio = insertTestPortfolio(t, "PR 9 allocation-plan pagination "+fmt.Sprint(count))
	t.Cleanup(
		inttestutil.BuildCleanupFunctionBuilder().
			AddCleanupQuery(
				`DELETE FROM planned_allocation WHERE allocation_plan_id IN (
					SELECT id FROM allocation_plan WHERE portfolio_id = {:portfolioId}
				)`,
				dbx.Params{"portfolioId": portfolio.Id},
			).
			AddCleanupQuery("DELETE FROM allocation_plan WHERE portfolio_id = {:portfolioId}", dbx.Params{
				"portfolioId": portfolio.Id,
			}).
			Build(t),
	)

	var priorities = allocationPlanPaginationPriorities(count)
	var baseDate = time.Date(2040, time.May, 15, 0, 0, 0, 0, time.UTC)
	var fixture = allocationPlanPaginationFixture{portfolio: portfolio, prefix: prefix}
	for index, priority := range priorities {
		var planName = fmt.Sprintf("%s%02d", prefix, index)
		var plannedDate = baseDate.AddDate(0, 0, priority)
		insertAllocationPlanPaginationRecord(t, fixture, planName, plannedDate, "ALLOCATION_PLAN")
	}
	if count == 21 {
		insertAllocationPlanPaginationRecord(
			t,
			fixture,
			prefix+"execution-plan-excluded",
			baseDate.AddDate(0, 0, 100),
			"EXECUTION_PLAN",
		)
	}

	var queryError = inttestinfra.FetchWithDBQuery(
		`SELECT id, name, planned_date
		 FROM allocation_plan
		 WHERE portfolio_id = {:portfolioId} AND type = 'ALLOCATION_PLAN' AND name LIKE {:namePrefix}`,
		dbx.Params{"portfolioId": portfolio.Id, "namePrefix": prefix + "%"},
		func(rows *dbx.Rows) error {
			var record allocationPlanPaginationRecord
			if scanError := rows.Scan(&record.id, &record.name, &record.plannedDate); scanError != nil {
				return scanError
			}
			fixture.ordered = append(fixture.ordered, record)
			return nil
		},
	)
	require.NoError(t, queryError)
	sort.Slice(fixture.ordered, func(left, right int) bool {
		if fixture.ordered[left].plannedDate.Equal(fixture.ordered[right].plannedDate) {
			return fixture.ordered[left].id > fixture.ordered[right].id
		}
		return fixture.ordered[left].plannedDate.After(fixture.ordered[right].plannedDate)
	})
	return fixture
}

// allocationPlanPaginationPriorities returns date offsets with four equal dates spanning the rank-ten boundary.
// Authored by: OpenCode
func allocationPlanPaginationPriorities(count int) []int {
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

// insertAllocationPlanPaginationRecord inserts one plan and two distinguishable hierarchy details.
// Authored by: OpenCode
func insertAllocationPlanPaginationRecord(
	t *testing.T,
	fixture allocationPlanPaginationFixture,
	name string,
	plannedDate time.Time,
	planType string,
) {
	t.Helper()
	var err = inttestinfra.ExecuteDBQuery(
		`INSERT INTO allocation_plan (portfolio_id, name, type, planned_date)
		 VALUES ({:portfolioId}, {:name}, {:planType}, {:plannedDate})`,
		dbx.Params{
			"portfolioId": fixture.portfolio.Id,
			"name":        name,
			"planType":    planType,
			"plannedDate": plannedDate,
		},
	)
	require.NoError(t, err)
	for _, suffix := range []string{"first", "second"} {
		err = inttestinfra.ExecuteDBQuery(
			`INSERT INTO planned_allocation (
				allocation_plan_id, hierarchical_id, asset_id, cash_reserve, slice_size_percentage
			)
			SELECT id, ARRAY[{:hierarchicalId}], NULL, FALSE, 0.5
			FROM allocation_plan
			WHERE portfolio_id = {:portfolioId} AND name = {:name}`,
			dbx.Params{
				"hierarchicalId": name + "-" + suffix,
				"portfolioId":    fixture.portfolio.Id,
				"name":           name,
			},
		)
		require.NoError(t, err)
	}
}

// verifyLegacyAllocationPlanRead confirms absent pageCursor keeps the complete unbounded JSON-array response.
// Authored by: OpenCode
func verifyLegacyAllocationPlanRead(t *testing.T, fixture allocationPlanPaginationFixture) {
	t.Helper()
	var response = getPaginationResponse(t, allocationPlanEndpoint(fixture.portfolio.Id))
	require.Equal(t, http.StatusOK, response.StatusCode)
	var body, err = io.ReadAll(response.Body)
	require.NoError(t, err)
	require.NoError(t, response.Body.Close())
	var plans []allocationPlanPageItem
	require.NoError(t, json.Unmarshal(body, &plans))
	assert.NotNil(t, plans)
	assert.Len(t, plans, len(fixture.ordered))
	assert.NotContains(t, string(body), `"nextCursor"`)
	for _, plan := range plans {
		assert.Equal(t, "ALLOCATION_PLAN", plan.Type)
		assert.Len(t, plan.Details, 2)
	}
}

// verifyAllocationPlanPages traverses plan pages and checks IDs, dates, detail completeness, and terminal cursors.
// Authored by: OpenCode
func verifyAllocationPlanPages(t *testing.T, fixture allocationPlanPaginationFixture) {
	t.Helper()
	var cursor *string
	var collectedIDs = make([]int64, 0, len(fixture.ordered))
	var returnedPlans = make([]allocationPlanPageItem, 0, len(fixture.ordered))
	for pageIndex := 0; ; pageIndex++ {
		require.Less(t, pageIndex, 10, "pagination must terminate")
		var page = getAllocationPlanPage(t, fixture.portfolio.Id, cursor)
		require.NotNil(t, page.Items)
		require.LessOrEqual(t, len(page.Items), 10)
		for _, item := range page.Items {
			collectedIDs = append(collectedIDs, item.Id)
			returnedPlans = append(returnedPlans, item)
			assert.Equal(t, "ALLOCATION_PLAN", item.Type)
			assert.Len(t, item.Details, 2)
			var actualHierarchy = make([]string, 0, len(item.Details))
			for _, detail := range item.Details {
				require.Len(t, detail.HierarchicalId, 1)
				require.NotNil(t, detail.HierarchicalId[0])
				actualHierarchy = append(actualHierarchy, *detail.HierarchicalId[0])
			}
			assert.ElementsMatch(t, []string{item.Name + "-first", item.Name + "-second"}, actualHierarchy)
		}
		cursor = page.NextCursor
		if cursor == nil {
			break
		}
	}
	var expectedIDs = make([]int64, len(fixture.ordered))
	var expectedDateByID = make(map[int64]string, len(fixture.ordered))
	for index, record := range fixture.ordered {
		expectedIDs[index] = record.id
		expectedDateByID[record.id] = record.plannedDate.Format(time.DateOnly)
	}
	for _, plan := range returnedPlans {
		assert.Equal(t, expectedDateByID[plan.Id], plan.PlannedDate)
	}
	assert.Equal(t, expectedIDs, collectedIDs)
	assert.Len(t, uniqueInt64s(collectedIDs), len(fixture.ordered))
}

// getAllocationPlanPage requests one existing-endpoint page using the opaque continuation cursor.
// Authored by: OpenCode
func getAllocationPlanPage(t *testing.T, portfolioId int64, cursor *string) allocationPlanPageResponse {
	t.Helper()
	var endpoint = allocationPlanEndpoint(portfolioId) + "?pageCursor="
	if cursor != nil {
		endpoint += url.QueryEscape(*cursor)
	}
	var response = getPaginationResponse(t, endpoint)
	defer closePaginationResponse(t, response)
	require.Equal(t, http.StatusOK, response.StatusCode)
	var body, err = io.ReadAll(response.Body)
	require.NoError(t, err)
	var page allocationPlanPageResponse
	require.NoError(t, json.Unmarshal(body, &page))
	assert.Contains(t, string(body), `"nextCursor"`)
	if page.NextCursor == nil {
		assert.Contains(t, string(body), `"nextCursor":null`)
	}
	return page
}

// allocationPlanEndpoint builds the existing allocation-plan endpoint URL for one portfolio.
// Authored by: OpenCode
func allocationPlanEndpoint(portfolioId int64) string {
	return fmt.Sprintf("%s/portfolio/%d/allocation-plan", inttestinfra.TestAPIURLPrefix, portfolioId)
}

// allocationPlanPageNames returns page plan names for insertion-boundary assertions.
// Authored by: OpenCode
func allocationPlanPageNames(page allocationPlanPageResponse) []string {
	var names = make([]string, len(page.Items))
	for index, plan := range page.Items {
		names[index] = plan.Name
	}
	return names
}

// invalidAllocationPlanCursorTokens builds validly encoded but invalid plan cursors for validation coverage.
// Authored by: OpenCode
func invalidAllocationPlanCursorTokens(portfolioId int64) []string {
	var base = map[string]any{
		"version": 1, "portfolioId": portfolioId, "collection": "allocation-plans",
		"plannedDate": "2040-05-15", "id": 1, "planType": "ALLOCATION_PLAN",
	}
	var invalidPayloads = []map[string]any{
		{"version": 2, "portfolioId": portfolioId, "collection": "allocation-plans", "plannedDate": "2040-05-15", "id": 1, "planType": "ALLOCATION_PLAN"},
		{"version": 1, "portfolioId": portfolioId + 1, "collection": "allocation-plans", "plannedDate": "2040-05-15", "id": 1, "planType": "ALLOCATION_PLAN"},
		{"version": 1, "portfolioId": portfolioId, "collection": "unknown", "plannedDate": "2040-05-15", "id": 1, "planType": "ALLOCATION_PLAN"},
		{"version": 1, "portfolioId": portfolioId, "collection": "allocation-plans", "plannedDate": "2040-02-30", "id": 1, "planType": "ALLOCATION_PLAN"},
		{"version": 1, "portfolioId": portfolioId, "collection": "allocation-plans", "plannedDate": "2040-05-15", "id": 1, "planType": "EXECUTION_PLAN"},
	}
	var tokens = make([]string, 0, len(invalidPayloads)+1)
	for _, payload := range invalidPayloads {
		tokens = append(tokens, encodeTestCursor(payload))
	}
	var withUnexpectedField = make(map[string]any, len(base)+1)
	for key, value := range base {
		withUnexpectedField[key] = value
	}
	withUnexpectedField["unexpected"] = true
	tokens = append(tokens, encodeTestCursor(withUnexpectedField))
	return tokens
}

// encodeTestCursor constructs an encoded cursor payload for black-box API validation tests.
// Authored by: OpenCode
func encodeTestCursor(payload map[string]any) string {
	encodedPayload, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}
	return base64.RawURLEncoding.EncodeToString(encodedPayload)
}
