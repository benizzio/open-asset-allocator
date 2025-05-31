package util

import (
	"fmt"
	universalTranslator "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// CustomFieldError implements validator.FieldError interface
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

// customValidationErrors is a custom implementation of validator.ValidationErrors
type customValidationErrors validator.ValidationErrors

// Error implements the error interface
func (errors customValidationErrors) Error() string {
	if len(errors) == 0 {
		return ""
	}
	return errors[0].Error()
}
