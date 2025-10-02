package util

import (
	"database/sql"
	"testing"

	infrautil "github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
)

// AssertDBWithQuery executes the given SQL query and compares the result with the expected value.
func AssertDBWithQuery(t *testing.T, sql string, expected dbx.NullStringMap) {

	var actual = dbx.NullStringMap{}

	err := inttestinfra.DatabaseConnection.NewQuery(sql).One(actual)
	if err != nil {
		t.Fatalf("Error executing query: %v", err)
	}

	assert.Equal(t, expected, actual, "Query result does not match expected value")
}

// AssertableNullString represents a NullString with an optional custom assertion function.
//
// If assertFunction is provided, it will be used to compare the actual and expected values
// instead of using standard equality comparison.
//
// Authored by: GitHub Copilot
type AssertableNullString struct {
	sql.NullString
	assertFunction func(t *testing.T, actual sql.NullString)
}

// AssertableNullStringMap represents a map where all values are AssertableNullString.
//
// This type provides better type safety and clearer intent when defining expected
// database query results with custom validation requirements. Each value can have
// its own custom assertion function or use standard equality comparison.
//
// Authored by: GitHub Copilot
type AssertableNullStringMap map[string]AssertableNullString

// AssertDBWithQueryMultipleRows executes the given SQL query and compares the results with the expected rows.
//
// Parameters:
//   - t: The testing context
//   - query: The SQL query to execute
//   - expected: A slice of expected rows as AssertableNullStringMaps, where each value
//     can have custom assertion logic or use standard equality comparison
//
// The method will fail the test if:
//   - The query execution fails
//   - The number of returned rows doesn't match the expected count
//   - Any row doesn't match the corresponding expected row (using custom or standard assertions)
//
// Authored by: GitHub Copilot
func AssertDBWithQueryMultipleRows(t *testing.T, query string, expected []AssertableNullStringMap) {
	actual := executeQueryForMultipleRows(t, query)
	validateRowCount(t, expected, actual)
	assertAllRows(t, expected, actual)
}

// executeQueryForMultipleRows executes the SQL query and returns the result rows.
//
// Authored by: GitHub Copilot
func executeQueryForMultipleRows(t *testing.T, query string) []dbx.NullStringMap {
	var actual []dbx.NullStringMap
	err := inttestinfra.DatabaseConnection.NewQuery(query).All(&actual)
	if err != nil {
		t.Fatalf("Error executing query: %v", err)
	}
	return actual
}

// validateRowCount ensures the number of actual rows matches the expected count.
//
// Authored by: GitHub Copilot
func validateRowCount(t *testing.T, expected []AssertableNullStringMap, actual []dbx.NullStringMap) {
	if len(actual) != len(expected) {
		t.Fatalf("Expected %d rows, but got %d rows. \nActual data was %v", len(expected), len(actual), actual)
	}
}

// assertAllRows validates each row against its expected values.
//
// Authored by: GitHub Copilot
func assertAllRows(t *testing.T, expected []AssertableNullStringMap, actual []dbx.NullStringMap) {
	for i, expectedRow := range expected {
		assertSingleRow(t, i, expectedRow, actual[i])
	}
}

// assertSingleRow validates a single row against its expected values.
//
// Authored by: GitHub Copilot
func assertSingleRow(t *testing.T, rowIndex int, expectedRow AssertableNullStringMap, actualRow dbx.NullStringMap) {
	for key, expectedValue := range expectedRow {
		assertRowColumn(t, rowIndex, key, expectedValue, actualRow)
	}
}

// assertRowColumn validates a single field within a row.
//
// Authored by: GitHub Copilot
func assertRowColumn(
	t *testing.T,
	rowIndex int,
	key string,
	expectedValue AssertableNullString,
	actualRow dbx.NullStringMap,
) {

	actualValue, exists := actualRow[key]
	if !exists {
		t.Fatalf("Row %d: Expected key '%s' not found in actual result", rowIndex, key)
	}

	if expectedValue.assertFunction != nil {
		expectedValue.assertFunction(t, actualValue)
	} else {
		assert.Equal(t, expectedValue.NullString, actualValue, "Row %d, key '%s': value does not match", rowIndex, key)
	}
}

// ToAssertableNullString converts a simple string to an AssertableNullString with standard equality assertion.
//
// This is a convenience function for creating AssertableNullString instances from regular strings
// without custom assertion logic. The resulting AssertableNullString will use standard equality
// comparison when used in assertions.
//
// Parameters:
//   - str: The string value to convert
//
// Returns:
//   - AssertableNullString with the provided string value and no custom assertion function
//
// Authored by: GitHub Copilot
func ToAssertableNullString(str string) AssertableNullString {
	return AssertableNullString{
		NullString: infrautil.StringToNullString(str),
	}
}

// ToAssertableNullStringWithAssertion converts a string to an AssertableNullString with a custom assertion function.
//
// This is a convenience function for creating AssertableNullString instances with custom validation logic.
// The provided assertion function will be used instead of standard equality comparison when validating
// database query results.
//
// Parameters:
//   - str: The string value to convert
//   - assertFunc: The custom assertion function to use for validation
//
// Returns:
//   - AssertableNullString with the provided string value and custom assertion function
//
// Authored by: GitHub Copilot
func ToAssertableNullStringWithAssertion(
	assertFunc func(t *testing.T, actual sql.NullString),
) AssertableNullString {
	return AssertableNullString{
		assertFunction: assertFunc,
	}
}

// NotNullAssertableNullString creates an AssertableNullString that asserts the value is not null or empty.
//
// This is a convenience function for creating AssertableNullString instances that specifically
// validate that the database field is neither null nor an empty string. The resulting
// AssertableNullString will use a custom assertion function to enforce this condition.
//
// Returns:
//   - AssertableNullString with a custom assertion function that checks for non-null and non-empty values
func NotNullAssertableNullString() AssertableNullString {
	return ToAssertableNullStringWithAssertion(
		func(t *testing.T, actual sql.NullString) {
			assert.NotEmpty(t, actual.String)
			assert.True(t, actual.Valid)
		},
	)
}

// NotNullValueCapturingAssertableNullString creates an AssertableNullString that asserts the value is not null or empty,
// and captures the actual value into the provided pointer.
//
// This is a convenience function for creating AssertableNullString instances that specifically
// validate that the database field is neither null nor an empty string, while also capturing
// the actual value for further use. The resulting AssertableNullString will use a custom
// assertion function to enforce this condition and store the value.
//
// Parameters:
//   - capturingPointer: Pointer to a string where the actual value will be stored if the assertion passes
//
// Returns:
//   - AssertableNullString with a custom assertion function that checks for non-null and non-empty values,
//     and captures the actual value
func NotNullValueCapturingAssertableNullString(capturingPointer *string) AssertableNullString {
	return ToAssertableNullStringWithAssertion(
		func(t *testing.T, actual sql.NullString) {
			assert.NotEmpty(t, actual.String)
			assert.True(t, actual.Valid)
			*capturingPointer = actual.String
		},
	)
}

// NullAssertableNullString creates an AssertableNullString that asserts the value is null.
//
// This is a convenience function for creating AssertableNullString instances that specifically
// validate that the database field is null. The resulting AssertableNullString will use a custom
// assertion function to enforce this condition.
//
// Returns:
//   - AssertableNullString with a custom assertion function that checks for null values
func NullAssertableNullString() AssertableNullString {
	return ToAssertableNullStringWithAssertion(
		func(t *testing.T, actual sql.NullString) {
			assert.False(t, actual.Valid)
			assert.Empty(t, actual.String)
		},
	)
}
