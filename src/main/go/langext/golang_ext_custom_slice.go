package langext

import (
	"fmt"
	"strings"
)

// CustomSlice is a named, generic slice type that provides convenience
// formatting helpers for a compact representation.
//
// It complements Go's default slice formatting by offering a bracketless,
// comma-separated representation via PrettyString().
//
// Usage:
//
//	// Using the type with its PrettyString() method
//	var s = CustomSlice[string]{"a", "b", "c"}
//	// Produces: a, b, c
//	_ = s.PrettyString()
//
// Authored by: GitHub Copilot
type CustomSlice[T any] []T

// PrettyString returns a bracketless, comma-separated string for the slice.
// A single space is included after each comma (", ").
//
// Examples:
//
//	CustomSlice[int]{1, 2, 3}.PrettyString()   -> "1, 2, 3"
//	CustomSlice[string]{"a"}.PrettyString()    -> "a"
//	CustomSlice[string]{}.PrettyString()       -> ""
//
// Formatting rules:
//   - Elements are rendered using fmt.Sprint, so any fmt.Stringer will be honored.
//   - No surrounding brackets are included.
//   - Elements are separated by a comma followed by a space (", ").
//
// Authored by: GitHub Copilot
func (slice CustomSlice[T]) PrettyString() string {
	return joinAny(slice, ", ")
}

// joinAny joins any slice into a string using the given separator, rendering each
// element via fmt.Sprint. It avoids multiple allocations by using strings.Builder.
//
// Authored by: GitHub Copilot
func joinAny[T any](elements []T, separator string) string {
	if len(elements) == 0 {
		return ""
	}

	var builder strings.Builder

	// Pre-size buffer heuristically: assume ~4 runes per element + separators.
	// This is a best-effort optimization that doesn't affect correctness.
	// Avoids overflows for very large slices.
	var estimatedSize = len(elements)*4 + (len(elements)-1)*len(separator)
	if estimatedSize > 0 {
		builder.Grow(estimatedSize)
	}

	for index, value := range elements {
		if index > 0 {
			_, _ = builder.WriteString(separator)
		}
		_, _ = builder.WriteString(fmt.Sprint(value))
	}

	return builder.String()
}
