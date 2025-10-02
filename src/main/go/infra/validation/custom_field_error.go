package validation

import (
	"fmt"
	"reflect"

	universalTranslator "github.com/go-playground/universal-translator"
)

// CustomFieldError implements validator.FieldError interface for custom validation errors.
// It provides a way to create validation errors programmatically rather than relying
// on automatic validation from the validator package.
//
// This struct stores all necessary information about a validation error:
// - Field information (name, namespace)
// - Validation tag that failed
// - Parameter for the validation (if applicable)
// - Value that failed validation
//
// This struct is used by the CustomValidationErrorsBuilder to create validation errors
// that can be sent to clients via RespondWithCustomValidationErrors.
//
// Usage:
//
//	Typically, you shouldn't create this struct directly. Instead, use the
//	CustomValidationErrorsBuilder through the BuildCustomValidationErrorsBuilder() function:
//
//	  validationErrors := BuildCustomValidationErrorsBuilder().
//	    CustomValidationError(
//	      user,           // Target struct
//	      "Email",        // Field namespace
//	      "required",     // Validation tag
//	      "",             // Parameter
//	      "",             // Value that failed
//	    ).Build()
//
// Authored by: GitHub Copilot
type CustomFieldError struct {
	field           string
	structField     string
	namespace       string
	structNamespace string
	tagValue        string
	paramValue      string
	valueValue      interface{}
}

func (e CustomFieldError) Tag() string             { return e.tagValue }
func (e CustomFieldError) ActualTag() string       { return e.tagValue }
func (e CustomFieldError) Namespace() string       { return e.namespace }
func (e CustomFieldError) StructNamespace() string { return e.structNamespace }
func (e CustomFieldError) Field() string           { return e.field }
func (e CustomFieldError) StructField() string     { return e.structField }
func (e CustomFieldError) Value() interface{}      { return e.valueValue }
func (e CustomFieldError) Param() string           { return e.paramValue }

func (e CustomFieldError) Kind() reflect.Kind {
	if e.valueValue == nil {
		return reflect.Invalid
	}
	return reflect.TypeOf(e.valueValue).Kind()
}

func (e CustomFieldError) Type() reflect.Type {
	if e.valueValue == nil {
		return nil
	}
	return reflect.TypeOf(e.valueValue)
}

func (e CustomFieldError) Error() string {
	return fmt.Sprintf(
		"Key: '%s' Error:Field validation for '%s' failed on the '%s' tag",
		e.namespace,
		e.field,
		e.tagValue,
	)
}

func (e CustomFieldError) Translate(_ universalTranslator.Translator) string {
	return e.Error()
}
