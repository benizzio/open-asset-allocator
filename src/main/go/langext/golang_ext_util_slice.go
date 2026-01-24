package langext

import (
	"reflect"
	"slices"
)

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

// SliceContainsZeroValue checks if any element in the slice has a zero value.
//
// Iterates through the slice and checks each element using IsZeroValue.
// Returns true as soon as any zero value element is found.
//
// Parameters:
//   - slice: The slice to check for zero value elements
//
// Returns:
//   - bool: true if any element is a zero value, false otherwise
//
// Example:
//
//	values := []int{1, 0, 3}
//	SliceContainsZeroValue(values) // returns true (0 is zero value for int)
//
//	names := []string{"a", "b", "c"}
//	SliceContainsZeroValue(names) // returns false
//
// Authored by: GitHub Copilot
func SliceContainsZeroValue[T any](slice []T) bool {

	for _, element := range slice {
		if IsZeroValue(element) {
			return true
		}
	}

	return false
}

// ToPointerSlice converts a slice of values to a slice of pointers.
//
// Each element in the input slice is converted to a pointer to that element.
//
// Parameters:
//   - slice: A slice of values of type S
//
// Returns:
//   - []*S: A slice of pointers to the values
//
// Authored by: GitHub Copilot
func ToPointerSlice[S any](slice []S) []*S {

	result := make([]*S, len(slice))

	for index, value := range slice {
		result[index] = &value
	}

	return result
}

// CleanNilPointersInSlice removes nil pointers from a slice of pointers.
//
// Uses slices.DeleteFunc to filter out nil elements from the input slice.
//
// Parameters:
//   - slice: A slice of pointers to type T
//
// Returns:
//   - []*T: A new slice with nil pointers removed
//
// Example:
//
//	a, b := 1, 2
//	pointers := []*int{&a, nil, &b, nil}
//	clean := CleanNilPointersInSlice(pointers)
//	// clean will be []*int{&a, &b}
//
// Authored by: GitHub Copilot
func CleanNilPointersInSlice[T any](slice []*T) []*T {

	var cleanSlice = slices.DeleteFunc(
		slice,
		func(item *T) bool {
			return item == nil
		},
	)

	return cleanSlice
}

// DereferenceSliceContent transforms a slice of pointers to a slice of values.
//
// Each pointer in the input slice is dereferenced to obtain its underlying value.
// Nil pointers are skipped and not included in the resulting slice.
//
// Parameters:
//   - slice: A slice of pointers to type T
//
// Returns:
//   - []T: A slice containing the dereferenced values (nil pointers excluded)
//
// Example:
//
//	a, b := 1, 2
//	pointers := []*int{&a, nil, &b}
//	values := DereferenceSliceContent(pointers)
//	// values will be []int{1, 2}
//
// Authored by: GitHub Copilot
func DereferenceSliceContent[T any](slice []*T) []T {

	result := make([]T, 0, len(slice))

	for _, ptr := range slice {

		if ptr != nil {
			result = append(result, *ptr)
		} else {
			var zeroValue T
			result = append(result, zeroValue)
		}
	}

	return result
}

// ReverseSlice returns a new slice with the elements in reverse order.
//
// This is a pure function: the original slice is not modified.
// Internally uses slices.Reverse on a copy of the input slice.
//
// Parameters:
//   - slice: The slice to reverse
//
// Returns:
//   - []T: A new slice with elements in reverse order
//
// Example:
//
//	original := []int{1, 2, 3}
//	reversed := ReverseSlice(original)
//	// reversed will be []int{3, 2, 1}
//	// original remains []int{1, 2, 3}
//
// Authored by: GitHub Copilot
func ReverseSlice[T any](slice []T) []T {

	result := make([]T, len(slice))
	copy(result, slice)
	slices.Reverse(result)

	return result
}
