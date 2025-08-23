package langext

import (
	"fmt"
	"testing"
)

// TestOrderedMapIteratorStringKeys tests the OrderedMapIterator with string keys - basic functionality.
//
// This test verifies the iterator's size functionality with string keys.
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorStringKeys(t *testing.T) {

	var data = map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
		"david":   4,
	}

	var iterator = NewOrderedMapIterator(data)

	// Test Size
	if iterator.Size() != 4 {
		t.Errorf("Expected size 4, got %d", iterator.Size())
	}
}

// TestOrderedMapIteratorStringKeysIteration tests complete key iteration with string keys.
//
// This test verifies that NextKey() iteration:
// - Processes keys in sorted order
// - Returns correct indices
// - Consumes all elements and reaches end state
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorStringKeysIteration(t *testing.T) {

	var data = map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
		"david":   4,
	}

	var expectedKeys = []string{"alice", "bob", "charlie", "david"}

	// Test iteration with NextKey() only
	var keyIterator = NewOrderedMapIterator(data)
	var iteratedKeys []string
	for i := 0; keyIterator.HasNext(); i++ {
		var key, iterIndex = keyIterator.NextKey()

		if iterIndex != i {
			t.Errorf("Expected index %d, got %d", i, iterIndex)
		}

		if key != expectedKeys[i] {
			t.Errorf("Expected key %s, got %s", expectedKeys[i], key)
		}

		iteratedKeys = append(iteratedKeys, key)
	}

	// Verify all keys were consumed and iterator reached end state
	if len(iteratedKeys) != len(expectedKeys) {
		t.Errorf("Expected to iterate %d keys, but iterated %d", len(expectedKeys), len(iteratedKeys))
	}

	if keyIterator.HasNext() {
		t.Error("Expected HasNext() to return false after consuming all keys")
	}
}

// TestOrderedMapIteratorStringValuesIteration tests complete value iteration with string keys.
//
// This test verifies that NextValue() iteration:
// - Processes values in correct order (corresponding to sorted keys)
// - Returns correct indices
// - Consumes all elements and reaches end state
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorStringValuesIteration(t *testing.T) {

	var data = map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
		"david":   4,
	}

	var expectedValues = []int{1, 2, 3, 4}

	// Test iteration with NextValue() only
	var valueIterator = NewOrderedMapIterator(data)
	var iteratedValues []int
	for i := 0; valueIterator.HasNext(); i++ {
		var value, iterIndex = valueIterator.NextValue()

		if iterIndex != i {
			t.Errorf("Expected index %d, got %d", i, iterIndex)
		}

		if value != expectedValues[i] {
			t.Errorf("Expected value %d, got %d", expectedValues[i], value)
		}

		iteratedValues = append(iteratedValues, value)
	}

	// Verify all values were consumed and iterator reached end state
	if len(iteratedValues) != len(expectedValues) {
		t.Errorf("Expected to iterate %d values, but iterated %d", len(expectedValues), len(iteratedValues))
	}

	if valueIterator.HasNext() {
		t.Error("Expected HasNext() to return false after consuming all values")
	}
}

// TestOrderedMapIteratorIntKeys tests the OrderedMapIterator with integer keys.
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorIntKeys(t *testing.T) {

	var data = map[int]string{
		30: "thirty",
		10: "ten",
		20: "twenty",
		5:  "five",
	}

	var expectedKeys = []int{5, 10, 20, 30}
	var expectedValues = []string{"five", "ten", "twenty", "thirty"}

	// Test key iteration
	var keyIterator = NewOrderedMapIterator(data)
	var iteratedKeys []int
	for i := 0; keyIterator.HasNext(); i++ {
		var key, _ = keyIterator.NextKey()

		if key != expectedKeys[i] {
			t.Errorf("Expected key %d, got %d", expectedKeys[i], key)
		}

		iteratedKeys = append(iteratedKeys, key)
	}

	// Verify all keys were consumed and iterator reached end state
	if len(iteratedKeys) != len(expectedKeys) {
		t.Errorf("Expected to iterate %d keys, but iterated %d", len(expectedKeys), len(iteratedKeys))
	}

	if keyIterator.HasNext() {
		t.Error("Expected HasNext() to return false after consuming all keys")
	}

	// Test value iteration
	var valueIterator = NewOrderedMapIterator(data)
	var iteratedValues []string
	for i := 0; valueIterator.HasNext(); i++ {
		var value, _ = valueIterator.NextValue()

		if value != expectedValues[i] {
			t.Errorf("Expected value %s, got %s", expectedValues[i], value)
		}

		iteratedValues = append(iteratedValues, value)
	}

	// Verify all values were consumed and iterator reached end state
	if len(iteratedValues) != len(expectedValues) {
		t.Errorf("Expected to iterate %d values, but iterated %d", len(expectedValues), len(iteratedValues))
	}

	if valueIterator.HasNext() {
		t.Error("Expected HasNext() to return false after consuming all values")
	}
}

// TestOrderedMapIteratorCurrentAndPointers tests Current and pointer methods.
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorCurrentAndPointers(t *testing.T) {

	var data = map[string]int{
		"b": 2,
		"a": 1,
	}

	var iterator = NewOrderedMapIterator(data)

	// Test HasNext before any iteration
	if !iterator.HasNext() {
		t.Error("Expected HasNext to be true initially")
	}

	// Move to first element using NextKey
	var key1, index1 = iterator.NextKey()
	if key1 != "a" || index1 != 0 {
		t.Errorf("Expected (a, 0), got (%s, %d)", key1, index1)
	}

	// Test Current methods
	var currentKV, currentIndex = iterator.Current()
	if currentKV.Key != "a" || currentKV.Value != 1 || currentIndex != 0 {
		t.Errorf("Expected Current to return (a, 1, 0), got (%s, %d, %d)", currentKV.Key, currentKV.Value, currentIndex)
	}

	// Test CurrentPointer
	var currentPtr, currentPtrIndex = iterator.CurrentPointer()
	if currentPtr.Key != "a" || currentPtr.Value != 1 || currentPtrIndex != 0 {
		t.Errorf(
			"Expected CurrentPointer to return (a, 1, 0), got (%s, %d, %d)",
			currentPtr.Key,
			currentPtr.Value,
			currentPtrIndex,
		)
	}

	// Test NextKeyPointer
	var nextKeyPtr, nextKeyIndex = iterator.NextKeyPointer()
	if *nextKeyPtr != "b" || nextKeyIndex != 1 {
		t.Errorf("Expected NextKeyPointer to return (b, 1), got (%s, %d)", *nextKeyPtr, nextKeyIndex)
	}

	// Should not have next now
	if iterator.HasNext() {
		t.Error("Expected HasNext to be false after iterating through all elements")
	}
}

// TestOrderedMapIteratorPointerMethods tests the pointer-based methods.
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorPointerMethods(t *testing.T) {

	var data = map[string]int{
		"b": 2,
		"a": 1,
	}

	// Test NextKeyPointer
	var keyIterator = NewOrderedMapIterator(data)
	var keyPtr1, index1 = keyIterator.NextKeyPointer()
	if *keyPtr1 != "a" || index1 != 0 {
		t.Errorf("Expected NextKeyPointer to return (a, 0), got (%s, %d)", *keyPtr1, index1)
	}

	// Test NextValuePointer
	var valueIterator = NewOrderedMapIterator(data)
	var valuePtr1, index2 = valueIterator.NextValuePointer()
	if *valuePtr1 != 1 || index2 != 0 {
		t.Errorf("Expected NextValuePointer to return (1, 0), got (%d, %d)", *valuePtr1, index2)
	}
}

// TestOrderedMapIteratorEmptyMap tests behavior with an empty map.
//
// Authored by: GitHub Copilot
func TestOrderedMapIteratorEmptyMap(t *testing.T) {

	var data = map[string]int{}
	var iterator = NewOrderedMapIterator(data)

	if iterator.Size() != 0 {
		t.Errorf("Expected size 0, got %d", iterator.Size())
	}

	if iterator.HasNext() {
		t.Error("Expected HasNext to be false for empty map")
	}
}

// ExampleOrderedMapIterator demonstrates basic usage of OrderedMapIterator.
//
// Authored by: GitHub Copilot
func ExampleOrderedMapIterator() {

	var data = map[string]int{
		"charlie": 3,
		"alice":   1,
		"bob":     2,
	}

	var iterator = NewOrderedMapIterator(data)

	fmt.Printf("Map size: %d\n", iterator.Size())

	// Example 1: Iterate through keys only
	var keyIterator = NewOrderedMapIterator(data)
	for keyIterator.HasNext() {
		var key, index = keyIterator.NextKey()
		fmt.Printf("Index %d: Key = %s\n", index, key)
	}

	// Example 2: Iterate through values only
	var valueIterator = NewOrderedMapIterator(data)
	for valueIterator.HasNext() {
		var value, index = valueIterator.NextValue()
		fmt.Printf("Index %d: Value = %d\n", index, value)
	}

	// Output:
	// Map size: 3
	// Index 0: Key = alice
	// Index 1: Key = bob
	// Index 2: Key = charlie
	// Index 0: Value = 1
	// Index 1: Value = 2
	// Index 2: Value = 3
}

// ExampleNewOrderedMapIterator_intKeys demonstrates usage with integer keys.
//
// Authored by: GitHub Copilot
func ExampleNewOrderedMapIterator_intKeys() {

	var data = map[int]string{
		30: "thirty",
		10: "ten",
		20: "twenty",
	}

	var iterator = NewOrderedMapIterator(data)

	for iterator.HasNext() {
		var key, index = iterator.NextKey()
		fmt.Printf("Index %d: Key = %d\n", index, key)
	}

	// Output:
	// Index 0: Key = 10
	// Index 1: Key = 20
	// Index 2: Key = 30
}
