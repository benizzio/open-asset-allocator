package util

import (
	"errors"
	"fmt"
	"github.com/benizzio/open-asset-allocator/api/rest/model"
	"github.com/benizzio/open-asset-allocator/langext"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"reflect"
)

// BindAndValidateJSONWithInvalidResponse binds the request body to the provided struct and validates it.
// It handles validation validationErrors by returning appropriate HTTP responses.
//
// Authored by: GitHub Copilot
func BindAndValidateJSONWithInvalidResponse(context *gin.Context, bindingTarget interface{}) (bool, error) {
	if err := context.ShouldBindJSON(bindingTarget); err == nil {
		return true, nil
	} else if validationErrors := extractValidationErrors(err); validationErrors != nil {
		// Handle validation validationErrors
		errorMessages := formatValidationErrorMessages(validationErrors, bindingTarget)
		sendValidationErrorResponse(context, errorMessages)
		return false, nil
	} else {
		// Handle non-validation validationErrors
		return false, err
	}
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

// formatValidationError formats validation validationErrors into readable messages.
//
// Authored by: GitHub Copilot
func formatValidationError(fieldError validator.FieldError) string {
	switch fieldError.Tag() {
	case "required":
		return "is required"
	case "min":
		return fmt.Sprintf("must be at least %s", fieldError.Param())
	case "max":
		return fmt.Sprintf("must not exceed %s", fieldError.Param())
	default:
		return fieldError.Tag()
	}
}

// sendValidationErrorResponse sends a standardized HTTP response for validation validationErrors.
//
// Authored by: GitHub Copilot
func sendValidationErrorResponse(context *gin.Context, errorMessages []string) {
	context.JSON(
		http.StatusBadRequest, model.ErrorResponse{
			ErrorMessage: "Validation failed",
			Details:      errorMessages,
		},
	)
}
