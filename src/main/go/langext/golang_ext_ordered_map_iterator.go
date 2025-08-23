package langext

import (
	"cmp"
	"sort"
)

// OrderedMapIterator provides ordered iteration over map entries using a sorted slice of keys.
//
// The iterator maintains the original map and a sorted slice of keys to ensure consistent
// iteration order. Keys are sorted using Go's built-in cmp.Compare function.
//
// Type parameters:
//   - K: The key type (must be ordered - supports <, <=, >=, > operators)
//   - V: The value type
//
// Authored by: GitHub Copilot
// TODO add an interface to this
type OrderedMapIterator[K cmp.Ordered, V any] struct {
	index       int
	orderedKeys []K
	sourceMap   map[K]V
}

// KeyValue represents a key-value pair from the map.
//
// Authored by: GitHub Copilot
type KeyValue[K cmp.Ordered, V any] struct {
	Key   K
	Value V
}

// HasNext returns true if there are more elements to iterate over.
//
// Returns:
//   - bool: true if there are more elements, false otherwise
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) HasNext() bool {
	return iterator.index < len(iterator.orderedKeys)-1
}

// NextKeyPointer advances the iterator and returns a pointer to the current key and its index.
//
// Returns:
//   - *K: pointer to the current key
//   - int: the current index in the ordered iteration
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) NextKeyPointer() (*K, int) {

	iterator.index++
	var key = &iterator.orderedKeys[iterator.index]
	var resultIndex = iterator.index

	return key, resultIndex
}

// NextValuePointer advances the iterator and returns a pointer to the current value and its index.
//
// Returns:
//   - *V: pointer to the current value
//   - int: the current index in the ordered iteration
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) NextValuePointer() (*V, int) {

	iterator.index++
	var key = iterator.orderedKeys[iterator.index]
	var value = iterator.sourceMap[key]
	var resultIndex = iterator.index

	return &value, resultIndex
}

// NextKey advances the iterator and returns the current key and its index.
//
// Returns:
//   - K: the current key
//   - int: the current index in the ordered iteration
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) NextKey() (K, int) {

	var keyPointer, index = iterator.NextKeyPointer()
	return *keyPointer, index
}

// NextValue advances the iterator and returns the current value and its index.
//
// Returns:
//   - V: the current value
//   - int: the current index in the ordered iteration
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) NextValue() (V, int) {

	var valuePointer, index = iterator.NextValuePointer()
	return *valuePointer, index
}

// CurrentPointer returns a pointer to the current KeyValue pair and its index without advancing.
//
// Returns:
//   - *KeyValue[K, V]: pointer to the current key-value pair
//   - int: the current index in the ordered iteration
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) CurrentPointer() (*KeyValue[K, V], int) {
	var key = iterator.orderedKeys[iterator.index]
	var value = iterator.sourceMap[key]
	var result = &KeyValue[K, V]{Key: key, Value: value}

	return result, iterator.index
}

// Current returns the current KeyValue pair and its index without advancing.
//
// Returns:
//   - KeyValue[K, V]: the current key-value pair
//   - int: the current index in the ordered iteration
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) Current() (KeyValue[K, V], int) {
	var pointer, index = iterator.CurrentPointer()
	return *pointer, index
}

// Size returns the total number of elements in the map.
//
// Returns:
//   - int: the number of key-value pairs in the map
//
// Authored by: GitHub Copilot
func (iterator *OrderedMapIterator[K, V]) Size() int {
	return len(iterator.orderedKeys)
}

// NewOrderedMapIterator creates a new OrderedMapIterator from the provided map.
//
// The iterator will sort the keys and iterate through them in ascending order.
// The original map is not modified.
//
// Parameters:
//   - sourceMap: the map to iterate over
//
// Returns:
//   - *OrderedMapIterator[K, V]: a new iterator positioned before the first element
//
// Usage example:
//
//	data := map[string]int{"charlie": 3, "alice": 1, "bob": 2}
//	iterator := NewOrderedMapIterator(data)
//	for iterator.HasNext() {
//	    key, index := iterator.NextKey()
//	    value, _ := iterator.NextValue() // Note: NextValue() advances iterator again
//	    fmt.Printf("Index %d: %s = %d\n", index, key, value)
//	}
//
//	// Or iterate through keys only:
//	iterator = NewOrderedMapIterator(data)
//	for iterator.HasNext() {
//	    key, index := iterator.NextKey()
//	    fmt.Printf("Key at index %d: %s\n", index, key)
//	}
//
//	// Or iterate through values only:
//	iterator = NewOrderedMapIterator(data)
//	for iterator.HasNext() {
//	    value, index := iterator.NextValue()
//	    fmt.Printf("Value at index %d: %d\n", index, value)
//	}
//
// Authored by: GitHub Copilot
func NewOrderedMapIterator[K cmp.Ordered, V any](sourceMap map[K]V) *OrderedMapIterator[K, V] {

	var orderedKeys = make([]K, 0, len(sourceMap))
	for key := range sourceMap {
		orderedKeys = append(orderedKeys, key)
	}

	sort.Slice(
		orderedKeys, func(i, j int) bool {
			return cmp.Compare(orderedKeys[i], orderedKeys[j]) < 0
		},
	)

	return &OrderedMapIterator[K, V]{
		index:       -1,
		orderedKeys: orderedKeys,
		sourceMap:   sourceMap,
	}
}
