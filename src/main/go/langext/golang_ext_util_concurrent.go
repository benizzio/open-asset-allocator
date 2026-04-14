package langext

import (
	"context"
	"runtime"
	"sync"
)

// concurrentSliceResult holds the output of a single concurrent slice-producing operation.
//
// Co-authored by: OpenCode and benizzio
type concurrentSliceResult[R any] struct {
	results []R
}

// FlatMapConcurrently applies a SliceFromItem function concurrently to each element of the input
// slice and flattens all returned slices into one result slice. Result ordering is not guaranteed.
//
// Example:
//
//	urls := []string{"https://api1.com", "https://api2.com"}
//	results, err := langext.FlatMapConcurrently(urls, func(url string) ([]Response, error) {
//		return fetchFromAPI(url)
//	})
//
// Co-authored by: OpenCode and benizzio
func FlatMapConcurrently[I any, R any](inputs []I, sliceFromItem SliceFromItem[I, R]) ([]R, error) {
	var sliceFromItemWithContext SliceFromItemWithContext[I, R] = func(_ context.Context, input I) ([]R, error) {
		return sliceFromItem(input)
	}

	return FlatMapConcurrentlyCtx(context.Background(), inputs, sliceFromItemWithContext)
}

// FlatMapConcurrentlyCtx applies a context-aware slice-producing function concurrently to each
// input, cancels remaining work when the context is done or the first error occurs, and flattens
// all successful result slices into one output slice. Result ordering is not guaranteed.
//
// Example:
//
//	results, err := langext.FlatMapConcurrentlyCtx(
//		ctx,
//		urls,
//		func(ctx context.Context, url string) ([]Response, error) {
//		return fetchFromAPI(ctx, url)
//		},
//	)
//
// Co-authored by: OpenCode and benizzio
func FlatMapConcurrentlyCtx[I any, R any](
	ctx context.Context,
	inputs []I,
	sliceFromItem SliceFromItemWithContext[I, R],
) ([]R, error) {

	if len(inputs) == 0 {
		return make([]R, 0), nil
	}

	var workContext, cancel = context.WithCancel(ctx)
	defer cancel()

	var inputChannel = buildConcurrentInputChannel(workContext, inputs)
	var resultChannel = make(chan concurrentSliceResult[R], len(inputs))
	var firstErrChannel = startFlatMapWorkers(
		workContext,
		buildFlatMapWorkerCount(len(inputs)),
		inputChannel,
		sliceFromItem,
		resultChannel,
		cancel,
	)

	var aggregated = collectConcurrentResults(resultChannel)
	var firstErr = readConcurrentError(firstErrChannel)

	if firstErr != nil {
		return nil, firstErr
	}

	if err := ctx.Err(); err != nil {
		return nil, err
	}

	return aggregated, nil
}

// startFlatMapWorkers starts the worker pool and closes coordination channels when processing ends.
//
// Authored by: OpenCode
func startFlatMapWorkers[I any, R any](
	workContext context.Context,
	workerCount int,
	inputChannel <-chan I,
	sliceFromItem SliceFromItemWithContext[I, R],
	resultChannel chan<- concurrentSliceResult[R],
	cancel context.CancelFunc,
) <-chan error {
	var firstErrChannel = make(chan error, 1)
	var waitGroup sync.WaitGroup

	for range workerCount {
		waitGroup.Add(1)
		go runFlatMapWorker(
			workContext,
			&waitGroup,
			inputChannel,
			sliceFromItem,
			resultChannel,
			firstErrChannel,
			cancel,
		)
	}

	go closeConcurrentResultChannels(&waitGroup, resultChannel, firstErrChannel)

	return firstErrChannel
}

// buildFlatMapWorkerCount returns the bounded worker count used by FlatMapConcurrentlyCtx.
//
// Co-authored by: OpenCode and benizzio
func buildFlatMapWorkerCount(inputCount int) int {
	if inputCount == 0 {
		return 0
	}

	var workerCount = runtime.GOMAXPROCS(0)
	if workerCount < 1 {
		workerCount = 1
	}

	if inputCount < workerCount {
		return inputCount
	}

	return workerCount
}

// buildConcurrentInputChannel creates the worker input channel and feeds it until the parent
// context is done or all items are sent.
//

// Authored by: OpenCode
func buildConcurrentInputChannel[I any](ctx context.Context, inputs []I) <-chan I {
	var inputChannel = make(chan I)

	go func() {
		defer close(inputChannel)

		for _, input := range inputs {
			select {
			case <-ctx.Done():
				return
			case inputChannel <- input:
			}
		}
	}()

	return inputChannel
}

// runFlatMapWorker drains the shared input channel until work is complete or an error occurs.
//
// Authored by: OpenCode
func runFlatMapWorker[I any, R any](
	workContext context.Context,
	waitGroup *sync.WaitGroup,
	inputChannel <-chan I,
	sliceFromItem SliceFromItemWithContext[I, R],
	resultChannel chan<- concurrentSliceResult[R],
	firstErrChannel chan<- error,
	cancel context.CancelFunc,
) {
	defer waitGroup.Done()

	for {
		var input, ok = readConcurrentInput(workContext, inputChannel)
		if !ok {
			return
		}

		var err = executeSliceFromItemConcurrently(workContext, input, sliceFromItem, resultChannel)
		if shouldStopFlatMapWorker(workContext, err, firstErrChannel, cancel) {
			return
		}
	}
}

// readConcurrentInput reads the next worker input or stops when the work context is done.
//
// Authored by: OpenCode
func readConcurrentInput[I any](ctx context.Context, inputChannel <-chan I) (I, bool) {
	select {
	case <-ctx.Done():
		var zeroValue I
		return zeroValue, false
	case input, ok := <-inputChannel:
		return input, ok
	}
}

// shouldStopFlatMapWorker decides whether a worker should stop after processing an item.
//
// Authored by: OpenCode
func shouldStopFlatMapWorker(
	workContext context.Context,
	err error,
	firstErrChannel chan<- error,
	cancel context.CancelFunc,
) bool {
	if err == nil {
		return false
	}

	if workContext.Err() != nil {
		return true
	}

	select {
	case firstErrChannel <- err:
		cancel()
	default:
	}

	return true
}

// closeConcurrentResultChannels closes worker coordination channels after all workers finish.
//
// Authored by: OpenCode
func closeConcurrentResultChannels[R any](
	waitGroup *sync.WaitGroup,
	resultChannel chan<- concurrentSliceResult[R],
	firstErrChannel chan<- error,
) {
	waitGroup.Wait()
	close(resultChannel)
	close(firstErrChannel)
}

// executeSliceFromItemConcurrently runs a single context-aware slice-producing operation and
// sends the successful result to the channel.
//
// Co-authored by: OpenCode and benizzio
func executeSliceFromItemConcurrently[I any, R any](
	ctx context.Context,
	input I,
	sliceFromItem SliceFromItemWithContext[I, R],
	resultChannel chan<- concurrentSliceResult[R],
) error {
	var results, err = sliceFromItem(ctx, input)
	if err != nil {
		return err
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	case resultChannel <- concurrentSliceResult[R]{results: results}:
		return nil
	}
}

// readConcurrentError returns the first worker error if one was reported.
//
// Authored by: OpenCode
func readConcurrentError(firstErrChannel <-chan error) error {
	for err := range firstErrChannel {
		return err
	}

	return nil
}

// collectConcurrentResults drains the result channel and aggregates the collected result slices.
//
// Co-authored by: OpenCode and benizzio
func collectConcurrentResults[R any](resultChannel <-chan concurrentSliceResult[R]) []R {
	var aggregated = make([]R, 0)

	for result := range resultChannel {
		aggregated = append(aggregated, result.results...)
	}

	return aggregated
}
