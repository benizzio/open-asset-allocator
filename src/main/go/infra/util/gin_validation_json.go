package util

import (
	"errors"
	"fmt"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// BindAndValidateJSONWithInvalidResponse binds the request body to the provided struct and validates it.
// It handles validation errors by returning appropriate HTTP responses.
// This implementation performs shallow validation via Gin binding and then deep validation
// of nested structures and arrays using custom reflection-based navigation.
//
// Authored by: GitHub Copilot
func BindAndValidateJSONWithInvalidResponse(context *gin.Context, bindingTarget interface{}) (bool, error) {

	var allErrorMessages []string

	// Phase 1: Bind the JSON to the target struct (includes shallow validation)
	bindingErr := context.ShouldBindJSON(bindingTarget)
	if bindingErr != nil {
		// Collect binding validation errors
		if validationErrors := extractValidationErrors(bindingErr); validationErrors != nil {
			bindingErrorMessages := formatValidationErrorMessages(validationErrors, bindingTarget)
			allErrorMessages = append(allErrorMessages, bindingErrorMessages...)
		} else {
			// Handle non-validation binding errors (malformed JSON, type mismatches, etc.)
			return false, bindingErr
		}
	}

	// Phase 2: Perform deep validation by navigating through the bound structure
	deepValidationErrors := performDeepValidation(bindingTarget)
	allErrorMessages = append(allErrorMessages, deepValidationErrors...)

	// If any validation errors were found, send combined response
	if len(allErrorMessages) > 0 {
		sendValidationErrorResponse(context, allErrorMessages)
		return false, nil
	}

	return true, nil
}

// extractValidationErrors attempts to extract validation errors from the given error.
// Returns nil if the error is not a validation error.
//
// Authored by: GitHub Copilot
func extractValidationErrors(inputError error) validator.ValidationErrors {
	var validationErrors validator.ValidationErrors
	if errors.As(inputError, &validationErrors) {
		return validationErrors
	}
	return nil
}

// formatValidationErrorMessages converts validation errors into human-readable messages.
//
// Authored by: GitHub Copilot
func formatValidationErrorMessages(validationErrors validator.ValidationErrors, targetStruct interface{}) []string {

	errorMessages := make([]string, 0, len(validationErrors))

	structType := langext.GetStructType(targetStruct)

	for _, validationError := range validationErrors {
		message := formatErrorMessage(validationError, structType)
		errorMessages = append(errorMessages, message)
	}

	return errorMessages
}

// formatErrorMessage formats a single validation error into a human-readable message.
//
// Authored by: GitHub Copilot
func formatErrorMessage(validationError validator.FieldError, structType reflect.Type) string {
	// Extract needed information from validation error
	namespace := validationError.Namespace()
	fieldName := validationError.Field()
	jsonFieldName := getJSONFieldName(namespace, fieldName, structType)

	return fmt.Sprintf(
		"Field '%s' failed validation: %s",
		jsonFieldName,
		formatValidationError(validationError),
	)
}

// formatValidationError formats validation error into readable messages.
//
// Co-authored by: GitHub Copilot
func formatValidationError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "is required"
	case "min":
		return fmt.Sprintf("must be at least %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("must not exceed %s", fieldError.Param())
	case "custom":
		return fieldError.Param()
	default:
		return fieldError.Tag()
	}
}

// performDeepValidation performs deep validation by navigating through nested structures and slices
// to validate fields that may have been missed by shallow validation.
//
// Authored by: GitHub Copilot
func performDeepValidation(target interface{}) []string {

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
		if validationErrors := extractValidationErrors(err); validationErrors != nil {
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

// buildFieldPath constructs the field path for error reporting using JSON field names
//
// Authored by: GitHub Copilot
func buildFieldPath(parentPath string, fieldType reflect.StructField) string {

	// Use langext utility to extract JSON field name
	fieldName := langext.ExtractJSONFieldName(fieldType)

	if parentPath == "" {
		return fieldName
	}
	return parentPath + "." + fieldName
}
