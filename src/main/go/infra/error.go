package infra

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
	"reflect"
)

// AppError represents an error that was handled or created in the app boundary
// (inside the layered architecture, originating from internal or library code).
type AppError struct {
	Message string
	Cause   error
}

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

func (appError *AppError) Error() string {
	return appError.Message
}

func (appError *AppError) String() string {
	return appError.Message + ". Cause: " + appError.Cause.Error()
}

func BuildAppError(message string, origin any) error {
	return newAppError(message, nil, origin)
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

func HandleAPIError(context *gin.Context, message string, cause error) bool {
	var handle = cause != nil
	if handle {
		glog.Error(message, ": ", cause)
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": message})
	}
	return handle
}
