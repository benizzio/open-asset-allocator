package langext

import (
	"context"
	"errors"
	"slices"
	"testing"
	"time"
)

// TestFlatMapConcurrently_EmptyInput verifies that FlatMapConcurrently returns an empty slice
// without error when no input values are provided.
//
// Authored by: OpenCode
func TestFlatMapConcurrently_EmptyInput(t *testing.T) {
	var results, err = FlatMapConcurrently([]int{}, func(input int) ([]string, error) {
		return []string{string(rune(input))}, nil
	})

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("Expected empty result slice, got %#v", results)
	}
}

// TestFlatMapConcurrently_AggregatesResults verifies that FlatMapConcurrently collects and
// flattens results from multiple concurrent executions.
//
// Authored by: OpenCode
func TestFlatMapConcurrently_AggregatesResults(t *testing.T) {
	var results, err = FlatMapConcurrently([]int{1, 2, 3}, func(input int) ([]int, error) {
		return []int{input, input * 10}, nil
	})

	if err != nil {
		t.Fatalf("Expected nil error, got %v", err)
	}

	slices.Sort(results)
	var expected = []int{1, 2, 3, 10, 20, 30}
	if !slices.Equal(results, expected) {
		t.Fatalf("Expected %#v, got %#v", expected, results)
	}
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

	if !errors.Is(err, expectedErr) {
		t.Fatalf("Expected error %v, got %v", expectedErr, err)
	}
	if results != nil {
		t.Fatalf("Expected nil results on error, got %#v", results)
	}
}
