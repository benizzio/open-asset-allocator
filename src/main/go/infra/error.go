package infra

import (
	"fmt"
	"reflect"

	"github.com/golang/glog"
)

// ======================================================================
// Error Types
// ======================================================================

// AppError represents an error that was handled or created in the app boundary
// (inside the layered architecture, originating from internal or library code).
type AppError struct {
	Message string
	Cause   error
}

func (appError *AppError) Error() string {
	return appError.Message
}

func (appError *AppError) String() string {
	return appError.Message + ". Cause: " + appError.Cause.Error()
}

type DomainValidationError struct {
	Message string
	Causes  []*AppError
}

func (domError *DomainValidationError) Error() string {
	return domError.Message
}

// ======================================================================
// Error API
// ======================================================================

func BuildAppError(message string, origin any) error {
	return newAppError(message, nil, origin)
}

func BuildAppErrorFormatted(origin any, message string, params ...any) error {
	return newAppError(fmt.Sprintf(message, params...), nil, origin)
}

func BuildAppErrorFormattedUnconverted(origin any, message string, params ...any) *AppError {
	return newAppError(fmt.Sprintf(message, params...), nil, origin)
}

func PropagateAsAppError(cause error, origin any) error {
	return PropagateAsAppErrorWithNewMessage(cause, cause.Error(), origin)
}

func PropagateAsAppErrorWithNewMessage(cause error, message string, origin any) error {
	if cause != nil {
		return newAppError(message, cause, origin)
	}
	return nil
}

func BuildDomainValidationError(message string, causes []*AppError) error {
	return &DomainValidationError{Message: message, Causes: causes}
}

// ======================================================================
// Internal functionality
// ======================================================================

func newAppError(message string, cause error, originType any) *AppError {
	logError(message, cause, originType)
	return &AppError{Message: message, Cause: cause}
}

func logError(message string, cause error, origin any) {
	var errorLog = "Error in " + reflect.TypeOf(origin).String() + ": " + message
	if cause != nil {
		errorLog += ". Cause: " + cause.Error()
	}
	glog.Error(errorLog)
}
