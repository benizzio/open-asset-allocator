package langext

import (
	"sort"
)

// OrderedMapIterator provides ordered iteration over map entries using a sorted slice of keys.
//
// The iterator maintains the original map and a sorted slice of keys to ensure consistent
// iteration order. Keys are sorted using Go's default comparison for the key type.
//
// Type parameters:
//   - K: The key type (must be comparable and sortable)
//   - V: The value type
//
// Authored by: GitHub Copilot
type OrderedMapIterator[K comparable, V any] struct {
	index       int
	orderedKeys []K
	sourceMap   map[K]V
}

// KeyValue represents a key-value pair from the map.
//
// Authored by: GitHub Copilot
type KeyValue[K comparable, V any] struct {
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
func NewOrderedMapIterator[K comparable, V any](sourceMap map[K]V) *OrderedMapIterator[K, V] {

	var orderedKeys = make([]K, 0, len(sourceMap))
	for key := range sourceMap {
		orderedKeys = append(orderedKeys, key)
	}

	sort.Slice(
		orderedKeys, func(i, j int) bool {
			return compareKeys(orderedKeys[i], orderedKeys[j])
		},
	)

	return &OrderedMapIterator[K, V]{
		index:       -1,
		orderedKeys: orderedKeys,
		sourceMap:   sourceMap,
	}
}

// compareKeys provides generic comparison for different key types.
//
// This function handles the most common comparable types used as map keys.
// For custom types, you may need to create a specialized constructor.
//
// Parameters:
//   - a, b: keys to compare
//
// Returns:
//   - bool: true if a < b
//
// Authored by: GitHub Copilot
func compareKeys[K comparable](a, b K) bool {
	switch any(a).(type) {
	case string:
		return any(a).(string) < any(b).(string)
	case int:
		return any(a).(int) < any(b).(int)
	case int8:
		return any(a).(int8) < any(b).(int8)
	case int16:
		return any(a).(int16) < any(b).(int16)
	case int32:
		return any(a).(int32) < any(b).(int32)
	case int64:
		return any(a).(int64) < any(b).(int64)
	case uint:
		return any(a).(uint) < any(b).(uint)
	case uint8:
		return any(a).(uint8) < any(b).(uint8)
	case uint16:
		return any(a).(uint16) < any(b).(uint16)
	case uint32:
		return any(a).(uint32) < any(b).(uint32)
	case uint64:
		return any(a).(uint64) < any(b).(uint64)
	case float32:
		return any(a).(float32) < any(b).(float32)
	case float64:
		return any(a).(float64) < any(b).(float64)
	default:
		// For other comparable types, we'll need to convert to string for comparison
		// This is a fallback and may not provide meaningful ordering for all types
		return false
	}
}
