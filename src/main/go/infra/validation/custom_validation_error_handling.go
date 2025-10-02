package validation

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/benizzio/open-asset-allocator/infra/json"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/go-playground/validator/v10"
)

// MapValidationErrorsToMessages maps validation errors from the validator library to human-readable messages.
// It uses JSON property names for field identification.
//
// Parameters:
//   - inputError: The error returned by the validator library (can be nil or non-validation error)
//   - targetStruct: The struct or pointer to struct that was validated
//
// Returns:
//   - A slice of human-readable error messages, or nil if the input error is not validation errors
//
// Co-authored by: GitHub Copilot
func MapValidationErrorsToMessages(inputError error, targetStruct interface{}) []string {
	var validationErrors = asValidationErrors(inputError)
	if validationErrors == nil {
		return nil
	}
	return FormatValidationErrorMessages(validationErrors, targetStruct)
}

// asValidationErrors attempts to cast validation errors from the given error.
// Returns nil if the error is not a validation error.
//
// Authored by: GitHub Copilot
func asValidationErrors(inputError error) validator.ValidationErrors {
	var validationErrors validator.ValidationErrors
	if errors.As(inputError, &validationErrors) {
		return validationErrors
	}
	return nil
}

// FormatValidationErrorMessages converts validation errors into human-readable messages.
//
// Authored by: GitHub Copilot
func FormatValidationErrorMessages(validationErrors validator.ValidationErrors, targetStruct interface{}) []string {

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
	jsonFieldName := json.GetJSONFieldName(namespace, fieldName, structType)

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
