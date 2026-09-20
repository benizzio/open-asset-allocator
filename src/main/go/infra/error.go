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

// UniqueConstraintViolationError represents an application error caused by a violated database
// uniqueness constraint. Its transport-specific mapping is handled by the API boundary.
//
// Authored by: OpenCode
type UniqueConstraintViolationError struct {
	ConstraintName string
	Message        string
	Details        []string
}

// Error returns the uniqueness-constraint violation message.
//
// Authored by: OpenCode
func (constraintError *UniqueConstraintViolationError) Error() string {
	return constraintError.Message
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

// BuildUniqueConstraintViolationError creates an error for a violated database uniqueness
// constraint with an optional list of details.
//
// Example:
//
//	return BuildUniqueConstraintViolationError(
//		"asset_ticker_uk",
//		"Asset already exists",
//		[]string{"Ticker is already registered"},
//	)
//
// Authored by: OpenCode
func BuildUniqueConstraintViolationError(
	constraintName string,
	message string,
	details []string,
) error {
	return &UniqueConstraintViolationError{
		ConstraintName: constraintName,
		Message:        message,
		Details:        details,
	}
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
