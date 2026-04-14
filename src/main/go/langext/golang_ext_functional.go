// Package langext provides Go language extensions not available in the standard library.
// This file defines generic functional type aliases for common function signatures.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
package langext

import "context"

// SliceFromItem is a functional type that takes a single input item and produces a slice
// of results along with a potential error. Useful as a parameter type for higher-order
// functions that fan out a single input into multiple outputs.
//
// Type parameters:
//   - I: the input item type
//   - R: the result element type
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type SliceFromItem[I any, R any] func(I) ([]R, error)

// SliceFromItemWithContext is a functional type that takes a context and a single input item and
// produces a slice of results along with a potential error.
//
// Type parameters:
//   - I: the input item type
//   - R: the result element type
//
// Authored by: OpenCode
type SliceFromItemWithContext[I any, R any] func(context.Context, I) ([]R, error)
