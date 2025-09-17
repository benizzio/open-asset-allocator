package langext

import (
	"encoding/json"
	"testing"
)

// TestParseableInt64_UnmarshalJSON tests the JSON unmarshaling functionality of ParseableInt64.
//
// Authored by: GitHub Copilot
func TestParseableInt64_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected ParseableInt64
		hasError bool
	}{
		{
			name:     "numeric value",
			input:    `123456789`,
			expected: ParseableInt64(123456789),
			hasError: false,
		},
		{
			name:     "string numeric value",
			input:    `"987654321"`,
			expected: ParseableInt64(987654321),
			hasError: false,
		},
		{
			name:     "zero value",
			input:    `0`,
			expected: ParseableInt64(0),
			hasError: false,
		},
		{
			name:     "string zero value",
			input:    `"0"`,
			expected: ParseableInt64(0),
			hasError: false,
		},
		{
			name:     "empty string",
			input:    `""`,
			expected: ParseableInt64(0),
			hasError: false,
		},
		{
			name:     "large int64 value",
			input:    `9223372036854775807`,
			expected: ParseableInt64(9223372036854775807),
			hasError: false,
		},
		{
			name:     "negative value",
			input:    `-123456789`,
			expected: ParseableInt64(-123456789),
			hasError: false,
		},
		{
			name:     "invalid string",
			input:    `"not-a-number"`,
			expected: ParseableInt64(0),
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result ParseableInt64
			err := json.Unmarshal([]byte(tt.input), &result)

			if tt.hasError {
				if err == nil {
					t.Errorf("Expected error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("Expected %d, got %d", tt.expected, result)
				}
			}
		})
	}
}