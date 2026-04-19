package gin

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	"github.com/benizzio/open-asset-allocator/infra/validation"
)

// BindAndValidateJSONWithInvalidResponse binds the request body to the provided struct and validates it.
// It handles validation errors by returning appropriate HTTP responses.
// This implementation performs shallow validation via Gin binding and then deep validation
// of nested structures and arrays using custom reflection-based navigation.
//
// Co-authored by: GitHub Copilot and OpenCode
func BindAndValidateJSONWithInvalidResponse(context *gin.Context, bindingTarget interface{}) (bool, error) {

	var allErrorMessages []string

	// Phase 1: Bind the JSON to the target struct (includes shallow validation)
	bindingErr := context.ShouldBindJSON(bindingTarget)
	if bindingErr != nil {

		// Collect binding validation errors
		var validationErrors = validation.MapValidationErrorsToMessages(bindingErr, bindingTarget)
		if validationErrors == nil {
			// Handle non-validation binding errors (malformed JSON, type mismatches, etc.)
			return false, bindingErr
		}

		allErrorMessages = append(allErrorMessages, validationErrors...)
	}

	// Phase 2: Perform deep validation by navigating through the bound structure
	deepValidationErrors := validation.DeepValidate(bindingTarget)
	allErrorMessages = append(allErrorMessages, deepValidationErrors...)
	allErrorMessages = deduplicateValidationMessages(allErrorMessages)

	// If any validation errors were found, send combined response
	if len(allErrorMessages) > 0 {
		sendValidationErrorResponse(context, allErrorMessages)
		return false, nil
	}

	return true, nil
}

// deduplicateValidationMessages preserves the original message order while removing duplicate
// validation messages emitted by the shallow and deep validation passes.
//
// Authored by: OpenCode
func deduplicateValidationMessages(messages []string) []string {
	var uniqueMessages = make([]string, 0, len(messages))
	var seenMessages = make(map[string]struct{}, len(messages))
	for _, message := range messages {
		if _, alreadySeen := seenMessages[message]; alreadySeen {
			continue
		}

		seenMessages[message] = struct{}{}
		uniqueMessages = append(uniqueMessages, message)
	}

	return uniqueMessages
}

// RespondWithCustomValidationErrors takes custom validation errors and sends a standardized
// validation error response using Gin. This is useful when you want to send validation errors
// that were generated programmatically rather than from binding/validation.
//
// Parameters:
//   - context: The Gin context used to send the HTTP response
//   - validationErrors: Validation errors created with CustomValidationErrorsBuilder
//   - targetStruct: The struct that contains field information (used for JSON field names)
//
// Example:
//
//	// Create validation errors using the builder pattern
//	validationErrors := BuildCustomValidationErrorsBuilder().
//	    CustomValidationError(
//	        user,           // Target struct
//	        "Email",        // Field namespace
//	        "required",     // Validation tag
//	        "",             // Parameter
//	        "",             // Value that failed
//	    ).
//	    CustomValidationError(
//	        user,           // Target struct
//	        "Age",          // Field namespace
//	        "min",          // Validation tag
//	        "18",           // Parameter
//	        16,             // Value that failed
//	    ).
//	    Build()
//
//	// Send HTTP response with these errors
//	gininfra.RespondWithCustomValidationErrors(ctx, validationErrors, user)
//
// Authored by: GitHub Copilot
func RespondWithCustomValidationErrors(
	context *gin.Context,
	validationErrors validator.ValidationErrors,
	targetStruct interface{},
) {
	// Format the error messages using the existing formatValidationErrorMessages function
	errorMessages := validation.FormatValidationErrorMessages(validationErrors, targetStruct)
	// Send the validation error response using the existing function
	sendValidationErrorResponse(context, errorMessages)
}
