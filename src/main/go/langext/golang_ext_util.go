package langext

import (
	"reflect"
	"slices"
)

func ToPointerSlice[S any](slice []S) []*S {
	result := make([]*S, len(slice))
	for index, value := range slice {
		result[index] = &value
	}
	return result
}

// IsZeroValue checks if a value equals its zero value.
//
// Authored by: GitHub Copilot
func IsZeroValue[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}

// UnwrapType removes pointer indirection if present.
//
// Authored by: GitHub Copilot
func UnwrapType(fieldType reflect.Type) reflect.Type {
	if fieldType.Kind() == reflect.Ptr {
		return fieldType.Elem()
	}
	return fieldType
}

// DereferenceValue dereferences all pointer levels from a reflect.Value, returning the final
// non-pointer value and whether any nil pointer was encountered during dereferencing.
//
// This function handles nested pointers of arbitrary depth (e.g., ***SomeStruct) by recursively
// dereferencing until it reaches a non-pointer value or encounters a nil pointer.
//
// Parameters:
//   - value: The reflect.Value to dereference (may be a pointer of any depth)
//
// Returns:
//   - reflect.Value: The final dereferenced value (non-pointer type)
//   - bool: true if any nil pointer was encountered during dereferencing, false otherwise
//
// Example:
//
//	var data ***MyStruct = &&&MyStruct{Name: "test"}
//	finalValue, isNil := DereferenceValue(reflect.ValueOf(data))
//	// finalValue will be of type MyStruct, isNil will be false
//
// Authored by: GitHub Copilot
func DereferenceValue(value reflect.Value) (reflect.Value, bool) {

	for value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return value, true
		}
		value = value.Elem()
	}

	return value, false
}

// IsSlice checks if the provided value is a slice type.
//
// Uses reflection to determine if the parameter is a slice, which is required
// for automatic conversion to pq.Array for PostgreSQL compatibility.
//
// Parameters:
//   - value: The value to check
//
// Returns:
//   - bool: true if the value is a slice, false otherwise
//
// Authored by: GitHub Copilot
func IsSlice(value any) bool {

	if value == nil {
		return false
	}

	var valueType = UnwrapType(reflect.TypeOf(value))
	return valueType.Kind() == reflect.Slice
}

func CleanNilPointersInSlice[T any](slice []*T) []*T {
	var cleanSlice = slices.DeleteFunc(
		slice,
		func(item *T) bool {
			return item == nil
		},
	)
	return cleanSlice
}
