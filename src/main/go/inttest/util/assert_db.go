package util

import (
	"database/sql"
	inttestinfra "github.com/benizzio/open-asset-allocator/inttest/infra"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/stretchr/testify/assert"
	"testing"
)

// AssertQuery executes the given SQL query and compares the result with the expected value.
func AssertQuery(t *testing.T, sql string, expected dbx.NullStringMap) {

	var actual = dbx.NullStringMap{}

	err := inttestinfra.DatabaseConnection.NewQuery(sql).One(actual)
	if err != nil {
		t.Fatalf("Error executing query: %v", err)
	}

	assert.Equal(t, expected, actual, "Query result does not match expected value")
}

func ToNullString(str string) sql.NullString {
	return ToNullStringFromPointer(&str)
}

func ToNullStringFromPointer(str *string) sql.NullString {
	if str == nil {
		return sql.NullString{Valid: false}
	} else {
		return sql.NullString{String: *str, Valid: true}
	}
}
