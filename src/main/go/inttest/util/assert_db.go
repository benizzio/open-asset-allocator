package util

import (
	"database/sql"
	infrautil "github.com/benizzio/open-asset-allocator/infra/util"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"testing"
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
// TODO clean
func AssertDBWithQueryMultipleRows(t *testing.T, query string, expected []AssertableNullStringMap) {

	var actual []dbx.NullStringMap

	err := inttestinfra.DatabaseConnection.NewQuery(query).All(&actual)
	if err != nil {
		t.Fatalf("Error executing query: %v", err)
	}

	if len(actual) != len(expected) {
		t.Fatalf("Expected %d rows, but got %d rows", len(expected), len(actual))
	}

	for i, expectedRow := range expected {
		actualRow := actual[i]

		for key, expectedValue := range expectedRow {
			actualValue, exists := actualRow[key]
			if !exists {
				t.Fatalf("Row %d: Expected key '%s' not found in actual result", i, key)
			}

			// Check if the expected value has a custom assertion function
			if expectedValue.assertFunction != nil {
				// Use custom assertion function
				expectedValue.assertFunction(t, actualValue)
			} else {
				// Use standard equality assertion
				assert.Equal(t, expectedValue.NullString, actualValue, "Row %d, key '%s': value does not match", i, key)
			}
		}
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
		NullString: infrautil.ToNullString(str),
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
