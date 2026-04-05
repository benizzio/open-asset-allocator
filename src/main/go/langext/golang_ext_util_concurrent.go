package langext

// concurrentSliceResult holds the output of a single concurrent slice-producing operation,
// bundling the result slice with any error that occurred during execution.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
type concurrentSliceResult[R any] struct {
	results []R
	err     error
}

// FlatMapConcurrently applies a SliceFromItem function concurrently to each element of the input
// slice, collects all results into a single aggregated slice, and returns the first error
// encountered from any goroutine. Each SliceFromItem invocation may return a slice of results,
// which are flattened into the final output.
//
// Parameters:
//   - inputs: the slice of input values to process concurrently
//   - sliceFromItem: a function that takes a single input and returns a result slice and an error
//
// Returns:
//   - []R: the flattened aggregation of all result slices
//   - error: the first error encountered, or nil if all goroutines succeeded
//
// Example:
//
//	urls := []string{"https://api1.com", "https://api2.com"}
//	results, err := langext.FlatMapConcurrently(urls, func(url string) ([]Response, error) {
//	    return fetchFromAPI(url)
//	})
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func FlatMapConcurrently[I any, R any](inputs []I, sliceFromItem SliceFromItem[I, R]) ([]R, error) {

	var resultChannel = make(chan concurrentSliceResult[R], len(inputs))

	for _, input := range inputs {
		go executeSliceFromItemConcurrently(input, sliceFromItem, resultChannel)
	}

	return collectConcurrentResults(resultChannel, len(inputs))
}

// executeSliceFromItemConcurrently runs a single SliceFromItem operation and sends the result
// to the channel.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func executeSliceFromItemConcurrently[I any, R any](input I, sliceFromItem SliceFromItem[I, R], resultChannel chan<- concurrentSliceResult[R]) {
	var results, err = sliceFromItem(input)
	resultChannel <- concurrentSliceResult[R]{results: results, err: err}
}

// collectConcurrentResults reads the expected number of results from the channel,
// aggregating result slices or returning the first error encountered.
//
// Authored by: GitHub Copilot (claude-opus-4.6)
func collectConcurrentResults[R any](resultChannel <-chan concurrentSliceResult[R], expectedCount int) ([]R, error) {

	var aggregated = make([]R, 0)

	for range expectedCount {
		var result = <-resultChannel
		if result.err != nil {
			return nil, result.err
		}
		aggregated = append(aggregated, result.results...)
	}

	return aggregated, nil
}
