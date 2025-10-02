package validation

import (
	"fmt"
	"reflect"

	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/go-playground/validator/v10"
)

// DeepValidate performs deep validation by navigating through nested structures and slices
// to validate fields that may have been missed by shallow validation.
//
// Authored by: GitHub Copilot
func DeepValidate(target interface{}) []string {
	var errorMessages []string
	validateValueDeep(reflect.ValueOf(target), "", &errorMessages)
	return errorMessages
}

// validateValueDeep recursively validates a reflect.Value and its nested structures
//
// Co-authored by: GitHub Copilot
func validateValueDeep(value reflect.Value, fieldPath string, errorMessages *[]string) {

	// Dereference all pointer levels using langext.DereferenceValue
	actualValue, isNil := langext.DereferenceValue(value)
	if isNil {
		return
	}

	switch actualValue.Kind() {
	case reflect.Struct:
		validateStructDeep(actualValue, fieldPath, errorMessages)
	case reflect.Slice, reflect.Array:
		validateSliceDeep(actualValue, fieldPath, errorMessages)
	default:
		return
	}
}

// validateStructDeep navigates through all fields in a struct to find nested structures and slices
// It delegates actual field validation to the validator library
//
// Authored by: GitHub Copilot
func validateStructDeep(structValue reflect.Value, fieldPath string, errorMessages *[]string) {

	// Use validator library to validate this struct's fields
	validateStructWithValidator(structValue, fieldPath, errorMessages)

	// Navigate through fields to find nested structures and slices
	structType := structValue.Type()
	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structType.Field(i)

		// Skip unexported fields
		if !fieldType.IsExported() {
			continue
		}

		// Build field path for error reporting
		currentPath := buildFieldPath(fieldPath, fieldType)

		// Recursively validate nested structures and slices
		validateValueDeep(field, currentPath, errorMessages)
	}
}

// validateStructWithValidator uses the validator library to validate struct fields
// and formats errors consistently with the rest of the validation system
//
// Authored by: GitHub Copilot
func validateStructWithValidator(structValue reflect.Value, fieldPath string, errorMessages *[]string) {

	// Create validator instance
	validate := validator.New()

	// Validate the struct using the validator library
	if err := validate.Struct(structValue.Interface()); err != nil {
		if validationErrors := asValidationErrors(err); validationErrors != nil {
			// Format validation errors and add them to the error messages
			structType := structValue.Type()

			for _, validationError := range validationErrors {
				// Build the field path using JSON property names
				errorFieldPath := buildValidationErrorPathWithJSON(fieldPath, validationError, structType)
				errorMessage := fmt.Sprintf(
					"Field '%s' failed validation: %s",
					errorFieldPath,
					formatValidationError(validationError),
				)
				*errorMessages = append(*errorMessages, errorMessage)
			}
		}
	}
}

// buildValidationErrorPathWithJSON constructs the field path for validation errors from nested structs
// using JSON property names instead of Go field names
//
// Authored by: GitHub Copilot
func buildValidationErrorPathWithJSON(
	basePath string,
	validationError validator.FieldError,
	structType reflect.Type,
) string {

	fieldName := validationError.Field()

	// Find the struct field to get its JSON tag
	field, found := structType.FieldByName(fieldName)
	if found {
		// Use langext utility to extract JSON field name
		fieldName = langext.ExtractJSONFieldName(field)
	}

	// If we have a base path, combine it with the JSON field name
	if basePath != "" {
		return basePath + "." + fieldName
	}

	return fieldName
}

// validateSliceDeep validates all elements in a slice or array by recursively validating each element
//
// Authored by: GitHub Copilot
func validateSliceDeep(sliceValue reflect.Value, fieldPath string, errorMessages *[]string) {

	for i := 0; i < sliceValue.Len(); i++ {
		element := sliceValue.Index(i)
		elementPath := fmt.Sprintf("%s[%d]", fieldPath, i)
		validateValueDeep(element, elementPath, errorMessages)
	}
}
