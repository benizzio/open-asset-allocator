package langext

import (
	"context"
	"errors"
	"runtime"
	"slices"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestFlatMapConcurrently_EmptyInput verifies that FlatMapConcurrently returns an empty slice
// without error when no input values are provided.
//
// Authored by: OpenCode
func TestFlatMapConcurrently_EmptyInput(t *testing.T) {
	var results, err = FlatMapConcurrently([]int{}, func(input int) ([]string, error) {
		return []string{string(rune(input))}, nil
	})

	require.NoError(t, err)
	assert.Empty(t, results)
}

// TestFlatMapConcurrently_AggregatesResults verifies that FlatMapConcurrently collects and
// flattens results from multiple concurrent executions.
//
// Authored by: OpenCode
func TestFlatMapConcurrently_AggregatesResults(t *testing.T) {
	var results, err = FlatMapConcurrently([]int{1, 2, 3}, func(input int) ([]int, error) {
		return []int{input, input * 10}, nil
	})

	require.NoError(t, err)

	slices.Sort(results)
	var expected = []int{1, 2, 3, 10, 20, 30}
	assert.True(t, slices.Equal(results, expected), "expected %#v, got %#v", expected, results)
}

// TestFlatMapConcurrentlyCtx_ReturnsFirstError verifies that FlatMapConcurrentlyCtx returns the
// first worker error and cancels remaining work.
//
// Authored by: OpenCode
func TestFlatMapConcurrentlyCtx_ReturnsFirstError(t *testing.T) {
	var expectedErr = errors.New("boom")

	var results, err = FlatMapConcurrentlyCtx(context.Background(), []int{1, 2, 3}, func(ctx context.Context, input int) ([]int, error) {
		if input == 2 {
			return nil, expectedErr
		}

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(50 * time.Millisecond):
			return []int{input}, nil
		}
	})

	require.ErrorIs(t, err, expectedErr)
	assert.Nil(t, results)
}

// TestFlatMapConcurrently_DoesNotLeakProducerOnWorkerError verifies that the producer goroutine
// observes worker cancellation on the FlatMapConcurrently path as well.
//
// Authored by: OpenCode
func TestFlatMapConcurrently_DoesNotLeakProducerOnWorkerError(t *testing.T) {
	var expectedErr = errors.New("boom")

	var inputs = make([]int, 256)
	for index := range inputs {
		inputs[index] = index
	}

	var baselineGoroutineCount = runtime.NumGoroutine()

	for range 20 {
		_, err := FlatMapConcurrently(inputs, func(input int) ([]int, error) {
			if input == 0 {
				return nil, expectedErr
			}

			time.Sleep(10 * time.Millisecond)
			return []int{input}, nil
		})

		require.ErrorIs(t, err, expectedErr)
	}

	var currentGoroutineCount = baselineGoroutineCount
	var deadline = time.Now().Add(500 * time.Millisecond)
	for time.Now().Before(deadline) {
		runtime.GC()
		currentGoroutineCount = runtime.NumGoroutine()
		if currentGoroutineCount <= baselineGoroutineCount+4 {
			assert.LessOrEqual(t, currentGoroutineCount, baselineGoroutineCount+4)
			return
		}

		time.Sleep(10 * time.Millisecond)
	}

	t.Fatalf(
		"Expected no leaked producer goroutines, baseline=%d current=%d",
		baselineGoroutineCount,
		currentGoroutineCount,
	)
}
