package rdbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBuildILikeSubstringPattern verifies wildcard escaping for reusable PostgreSQL ILIKE
// substring search patterns.
//
// Authored by: OpenCode
func TestBuildILikeSubstringPattern(t *testing.T) {
	assert.Equal(
		t,
		`%A\_B\%C\\D%`,
		BuildILikeSubstringPattern(`A_B%C\D`),
	)
}
