package langext

import "strings"

// IsEmptyString reports whether trimming Unicode whitespace leaves the string's zero value.
// For example, IsEmptyString(" \t\n") returns true, while IsEmptyString(" tag ") returns false.
// Authored by: GPT-6.1 Sol
func IsEmptyString(value string) bool {
	return IsZeroValue(strings.TrimSpace(value))
}
