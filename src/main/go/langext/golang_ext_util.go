package langext

import (
	"reflect"
)

// IsZeroValue checks if a value equals its zero value.
//
// Authored by: GitHub Copilot
func IsZeroValue[T any](value T) bool {
	return reflect.ValueOf(value).IsZero()
}

// IsNilPointer checks if a value is a typed nil pointer.
//
// Authored by: OpenCode
func IsNilPointer(value any) bool {
	if value == nil {
		return false
	}

	var valueRef = reflect.ValueOf(value)
	return valueRef.Kind() == reflect.Ptr && valueRef.IsNil()
}

// UnwrapType removes all pointer indirection layers from a reflect.Type.
//
// Co-authored by: OpenCode and GitHub Copilot
func UnwrapType(fieldType reflect.Type) reflect.Type {
	for fieldType != nil && fieldType.Kind() == reflect.Ptr {
		fieldType = fieldType.Elem()
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
