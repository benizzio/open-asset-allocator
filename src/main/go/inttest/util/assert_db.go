package util

import (
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

// AssertDBWithQueryMultipleRows executes the given SQL query and compares the results with the expected rows.
//
// Parameters:
//   - t: The testing context
//   - sql: The SQL query to execute
//   - expected: A slice of expected rows as NullStringMaps
//
// The method will fail the test if:
//   - The query execution fails
//   - The number of returned rows doesn't match the expected count
//   - Any row doesn't match the corresponding expected row
//
// Authored by: GitHub Copilot
func AssertDBWithQueryMultipleRows(t *testing.T, sql string, expected []dbx.NullStringMap) {

	var actual []dbx.NullStringMap

	err := inttestinfra.DatabaseConnection.NewQuery(sql).All(&actual)
	if err != nil {
		t.Fatalf("Error executing query: %v", err)
	}

	assert.Equal(t, expected, actual, "Query results do not match expected rows")
}
