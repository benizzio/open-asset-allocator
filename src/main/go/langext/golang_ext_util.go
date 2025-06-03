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
