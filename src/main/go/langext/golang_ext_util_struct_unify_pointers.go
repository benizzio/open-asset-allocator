package langext

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// UnifyStructPointers processes a slice of any struct type and ensures that all pointers
// to equal values across the structs are unified to point to the same memory location.
// This is useful for reducing memory usage and improving equality comparisons.
//
// Parameters:
//   - items: A slice of any struct type to process
//
// Example:
//
//	type Person struct {
//	  Name *string
//	  Age  int
//	}
//	people := []Person{...}
//	UnifyStructPointers(people)
//
// After calling this function, any Person structs that had different pointers
// to equal Name values will now share the same pointer.
//
// Authored by: GitHub Copilot
func UnifyStructPointers[T any](items []T) {

	if len(items) <= 1 {
		return
	}

	// Get type information
	var itemType = reflect.TypeOf(items).Elem()

	// Process each field in the struct
	for j := 0; j < itemType.NumField(); j++ {
		unifyPointersForField(items, itemType, j)
	}
}

// unifyPointersForField processes a single field across all structs in the slice,
// unifying any pointers to equal values.
//
// Parameters:
//   - items: The slice of structs to process
//   - itemType: The reflect.Type of the structs in the slice
//   - fieldIndex: The index of the field to process
//
// Authored by: GitHub Copilot
func unifyPointersForField[T any](items []T, itemType reflect.Type, fieldIndex int) {

	var field = itemType.Field(fieldIndex)

	// Skip non-pointer fields
	if field.Type.Kind() != reflect.Ptr {
		return
	}

	// For each field, we'll create a map of value -> pointer
	var valueMap = make(map[string]interface{})

	// Process all items in the slice
	for i := range items {
		processItemField(items, i, fieldIndex, valueMap)
	}
}

// processItemField handles the unification of a specific field in a specific item.
//
// Parameters:
//   - items: The slice of structs to process
//   - itemIndex: The index of the item to process
//   - fieldIndex: The index of the field to process
//   - valueMap: The map of value keys to pointer values
//
// Authored by: GitHub Copilot
func processItemField[T any](items []T, itemIndex int, fieldIndex int, valueMap map[string]interface{}) {

	var itemValue = reflect.ValueOf(&items[itemIndex]).Elem()
	var fieldValue = itemValue.Field(fieldIndex)

	// Skip nil pointers
	if fieldValue.IsNil() {
		return
	}

	// Get the actual value that the pointer points to
	var pointedValue = fieldValue.Elem()

	// Create a unique key for the value
	var valueKey = createValueKey(pointedValue)
	if valueKey == "" {
		return // Skip if we couldn't create a key
	}

	// Check if we've seen this value before
	if existingPtr, found := valueMap[valueKey]; found {
		// Replace the pointer with the existing one
		fieldValue.Set(reflect.ValueOf(existingPtr))
	} else {
		// Store this pointer for future reference
		valueMap[valueKey] = fieldValue.Interface()
	}
}

// createValueKey generates a string key that uniquely identifies a value.
// Returns an empty string if the type is not supported.
//
// Parameters:
//   - value: The reflect.Value to create a key for
//
// Returns:
//   - A string key that uniquely identifies the value, or an empty string if the type is not supported
//
// Authored by: GitHub Copilot
func createValueKey(value reflect.Value) string {

	var valueKey string

	switch value.Kind() {
	case reflect.String:
		valueKey = "s:" + value.String()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		valueKey = fmt.Sprintf("i:%d", value.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		valueKey = fmt.Sprintf("u:%d", value.Uint())
	case reflect.Float32, reflect.Float64:
		// Use %f for consistent representation
		valueKey = fmt.Sprintf("f:%f", value.Float())
	case reflect.Bool:
		valueKey = fmt.Sprintf("b:%t", value.Bool())
	case reflect.Struct:
		// For structs, use JSON serialization
		jsonBytes, err := json.Marshal(value.Interface())
		if err != nil {
			return "" // Skip if can't marshal
		}
		valueKey = "struct:" + string(jsonBytes)
	default:
		// Skip unsupported types
		return ""
	}

	return valueKey
}
