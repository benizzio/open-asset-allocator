package langext

import (
	"reflect"
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
