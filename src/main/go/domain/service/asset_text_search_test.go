package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestParseAssetTextSearch verifies ordinary terms, quoted phrases, and unmatched quote handling.
//
// Authored by: OpenCode
func TestParseAssetTextSearch(t *testing.T) {
	var testCases = []struct {
		name       string
		textSearch string
		expected   []string
	}{
		{
			name:       "unquoted terms",
			textSearch: "spdr bloomberg",
			expected:   []string{"spdr", "bloomberg"},
		},
		{
			name:       "quoted phrase",
			textSearch: `"spdr bloomberg"`,
			expected:   []string{"spdr bloomberg"},
		},
		{
			name:       "mixed terms and phrase",
			textSearch: `etf "spdr bloomberg" bond`,
			expected:   []string{"etf", "spdr bloomberg", "bond"},
		},
		{
			name:       "unmatched quote is removed",
			textSearch: `"spdr bloomberg`,
			expected:   []string{"spdr", "bloomberg"},
		},
		{
			name:       "unmatched quote does not split a term",
			textSearch: `spdr"bloomberg`,
			expected:   []string{"spdrbloomberg"},
		},
		{
			name:       "phrase whitespace is normalized",
			textSearch: `"  spdr   bloomberg  "`,
			expected:   []string{"spdr bloomberg"},
		},
		{
			name:       "blank search",
			textSearch: "  \t\n",
			expected:   []string{},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, parseAssetTextSearch(testCase.textSearch))
		})
	}
}
