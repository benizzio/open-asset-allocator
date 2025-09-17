package langext

import (
	"encoding/json"
	"testing"
)

// TestParseInt64 tests the ParseInt64 utility function.
//
// Authored by: GitHub Copilot
func TestParseInt64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int64
		hasError bool
	}{
		{
			name:     "positive number",
			input:    "123456789",
			expected: 123456789,
			hasError: false,
		},
		{
			name:     "negative number",
			input:    "-987654321",
			expected: -987654321,
			hasError: false,
		},
		{
			name:     "zero",
			input:    "0",
			expected: 0,
			hasError: false,
		},
		{
			name:     "large int64 value",
			input:    "9223372036854775807",
			expected: 9223372036854775807,
			hasError: false,
		},
		{
			name:     "invalid string",
			input:    "not-a-number",
			expected: 0,
			hasError: true,
		},
		{
			name:     "empty string",
			input:    "",
			expected: 0,
			hasError: true,
		},
		{
			name:     "overflow",
			input:    "9223372036854775808", // Max int64 + 1
			expected: 0,
			hasError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseInt64(tt.input)

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