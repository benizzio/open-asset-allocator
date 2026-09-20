package rdbms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestBuildCaseInsensitiveSubstringPattern verifies wildcard escaping for reusable PostgreSQL
// substring search patterns.
//
// Authored by: OpenCode
func TestBuildCaseInsensitiveSubstringPattern(t *testing.T) {
	assert.Equal(
		t,
		`%A\_B\%C\\D%`,
		BuildCaseInsensitiveSubstringPattern(`A_B%C\D`),
	)
}
