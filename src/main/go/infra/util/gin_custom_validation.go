package util

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// CustomValidationErrorsBuilder provides a fluent API for building custom validation errors.
type CustomValidationErrorsBuilder struct {
	errors []validator.FieldError
}

// CustomValidationError adds a validation error to the builder.
// Parameters:
//   - targetStruct: The struct or pointer to struct that contains the field
//   - fieldNamespace: The namespace of the field (e.g., "User.Address.Street" or just "Email")
//   - tag: The validation tag that failed (e.g., "required", "min", "max", "email")
//   - param: Additional parameter for the validation tag
//   - value: The actual value that failed validation
//
// Returns:
//   - The builder instance for method chaining
//
// Example:
//
//	// Create a validation error for a missing required field
//	errors := BuildCustomValidationErrorsBuilder().
//	    CustomValidationError(
//	        user,           // Target struct
//	        "Email",        // Field namespace
//	        "required",     // Validation tag
//	        "",             // Parameter
//	        "",             // Value that failed
//	    )
//
//	// Create a validation error for a minimum value constraint
//	errors.CustomValidationError(
//	    user,           // Target struct
//	    "Age",          // Field namespace
//	    "min",          // Validation tag
//	    "18",           // Parameter - minimum age is 18
//	    16,             // Value that failed - 16 is less than minimum 18
//	)
//
//	// Create a validation error for a nested field with max length
//	errors.CustomValidationError(
//	    user,                  // Target struct
//	    "Address.Street",      // Field namespace (nested)
//	    "max",                 // Validation tag
//	    "100",                 // Parameter - max length is 100
//	    "This street name is way too long...",  // Value that failed
//	)
//
//	// Build and use the validation errors
//	validationErrors := errors.Build()
//	RespondWithCustomValidationErrors(ctx, validationErrors, user)
func (b *CustomValidationErrorsBuilder) CustomValidationError(
	targetStruct interface{},
	fieldNamespace string,
	tag string,
	param string,
	value interface{},
) *CustomValidationErrorsBuilder {
	error := buildCustomValidationError(targetStruct, fieldNamespace, tag, param, value)
	b.errors = append(b.errors, error)
	return b
}

// Build creates and returns the validator.ValidationErrors collection.
// Returns:
//   - validator.ValidationErrors containing all the errors added to the builder
func (b *CustomValidationErrorsBuilder) Build() validator.ValidationErrors {
	return buildCustomValidationErrors(b.errors...)
}

// BuildCustomValidationErrorsBuilder creates a new instance of the builder.
// This function creates and returns a new builder for creating custom validation errors
// with a fluent interface.
//
// Returns:
//   - A new CustomValidationErrorsBuilder instance
//
// Example:
//
//	// Create a builder and add validation errors
//	builder := BuildCustomValidationErrorsBuilder()
//	builder.CustomValidationError(
//	    user,           // Target struct
//	    "Email",        // Field namespace
//	    "required",     // Validation tag
//	    "",             // Parameter
//	    "",             // Value that failed
//	)
//
//	// Build the validation errors collection
//	validationErrors := builder.Build()
//
//	// Or create validation errors in a fluent style
//	validationErrors := BuildCustomValidationErrorsBuilder().
//	    CustomValidationError(
//	        portfolioDTS,
//	        "Id",
//	        "required",
//	        "Portfolio ID is required for update",
//	        nil,
//	    ).
//	    Build()
//
//	// Send response with the validation errors
//	RespondWithCustomValidationErrors(context, validationErrors, portfolioDTS)
func BuildCustomValidationErrorsBuilder() *CustomValidationErrorsBuilder {
	return &CustomValidationErrorsBuilder{
		errors: []validator.FieldError{},
	}
}

// buildCustomValidationError creates a custom validation error that implements validator.FieldError.
//
// Parameters:
//   - targetStruct: The struct or pointer to struct that contains the field
//   - fieldNamespace: The namespace of the field
//   - tag: The validation tag that failed
//   - param: Additional parameter for the validation tag
//   - value: The actual value that failed validation
//
// Returns:
//   - validator.FieldError that can be used with validation error handling functions
func buildCustomValidationError(
	targetStruct interface{},
	fieldNamespace,
	tag,
	param string,
	value interface{},
) validator.FieldError {
	// Get struct name using reflection
	structName := GetStructName(targetStruct)

	// Parse the field namespace to get the full namespace and field name
	namespace, fieldName := getNamespaceInfo(structName, fieldNamespace)

	// Create and return the custom validation error
	return CustomFieldError{
		field:           fieldName,
		structField:     fieldName,
		namespace:       namespace,
		structNamespace: namespace,
		tagValue:        tag,
		paramValue:      param,
		valueValue:      value,
	}
}

// buildCustomValidationErrors creates a collection of validation errors that implements validator.ValidationErrors.
//
// Parameters:
//   - errors: A variadic list of validator.FieldError objects
//
// Returns:
//   - validator.ValidationErrors that can be used with validation error handling functions
func buildCustomValidationErrors(errors ...validator.FieldError) validator.ValidationErrors {
	// CustomValidationErrors satisfies the validator.ValidationErrors interface
	return validator.ValidationErrors(customValidationErrors(errors))
}

// RespondWithCustomValidationErrors takes custom validation errors and sends a standardized
// validation error response using Gin. This is useful when you want to send validation errors
// that were generated programmatically rather than from binding/validation.
//
// Parameters:
//   - context: The Gin context used to send the HTTP response
//   - validationErrors: Custom validation errors created with buildCustomValidationError
//   - targetStruct: The struct that contains field information (used for JSON field names)
//
// Example:
//
//	// Create custom validation errors
//	errors := []validator.FieldError{
//	    util.buildCustomValidationError("User", "Email", "required", "", ""),
//	    util.buildCustomValidationError("User", "Age", "min", "18", 16),
//	}
//
//	// Send HTTP response with these errors
//	util.RespondWithCustomValidationErrors(ctx, errors, userObject)
func RespondWithCustomValidationErrors(
	context *gin.Context,
	validationErrors validator.ValidationErrors,
	targetStruct interface{},
) {
	// Format the error messages using the existing formatValidationErrorMessages function
	errorMessages := formatValidationErrorMessages(validationErrors, targetStruct)

	// Send the validation error response using the existing function
	sendValidationErrorResponse(context, errorMessages)
}
