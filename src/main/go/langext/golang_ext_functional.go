// Package langext provides Go language extensions not available in the standard library.
// This file defines generic functional type aliases for common function signatures.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
package langext

// SliceFromItem is a functional type that takes a single input item and produces a slice
// of results along with a potential error. Useful as a parameter type for higher-order
// functions that fan out a single input into multiple outputs.
//
// Type parameters:
//   - I: the input item type
//   - R: the result element type
//
// Example:
//
//	var fetchAssets langext.SliceFromItem[string, Asset] = func(url string) ([]Asset, error) {
//	    return httpClient.FetchAssets(url)
//	}
//	results, err := fetchAssets("https://api.example.com/assets")
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type SliceFromItem[I any, R any] func(I) ([]R, error)
